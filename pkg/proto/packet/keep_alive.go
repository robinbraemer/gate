package packet

import (
	"go.minekube.com/gate/pkg/proto"
	"go.minekube.com/gate/pkg/proto/util"
	"io"
)

type KeepAlive struct {
	RandomId int64
}

func (k *KeepAlive) Encode(c *proto.PacketContext, wr io.Writer) error {
	if c.Protocol.GreaterEqual(proto.Minecraft_1_12_2) {
		return util.WriteInt64(wr, k.RandomId)
	} else if c.Protocol.GreaterEqual(proto.Minecraft_1_8) {
		return util.WriteVarInt(wr, int(k.RandomId))
	}
	return util.WriteInt32(wr, int32(k.RandomId))
}

func (k *KeepAlive) Decode(c *proto.PacketContext, r io.Reader) (err error) {
	rd := &countingReader{Reader: r}
	if c.Protocol.GreaterEqual(proto.Minecraft_1_12_2) {
		k.RandomId, err = util.ReadInt64(rd)
	} else if c.Protocol.GreaterEqual(proto.Minecraft_1_8) {
		var id int
		id, err = util.ReadVarInt(rd)
		k.RandomId = int64(id)
	} else {
		var id int32
		id, err = util.ReadInt32(rd)
		k.RandomId = int64(id)
	}
	return
}

type countingReader struct {
	io.Reader
	n int
}

func (cr *countingReader) Read(p []byte) (n int, err error) {
	n, err = cr.Reader.Read(p)
	cr.n += n
	return
}

var _ proto.Packet = (*KeepAlive)(nil)
