package session

import (
	"testing"

	"github.com/XiaoMi/pegasus-go-client/idl/base"
	"github.com/XiaoMi/pegasus-go-client/idl/replication"
	"github.com/XiaoMi/pegasus-go-client/idl/rrdb"
	"github.com/stretchr/testify/assert"
)

func TestCodec_Marshal(t *testing.T) {
	expected := []byte{
		0x54, 0x48, 0x46, 0x54, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x4a, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x80, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x26,
		0x52, 0x50, 0x43, 0x5f, 0x43, 0x4d, 0x5f, 0x51,
		0x55, 0x45, 0x52, 0x59, 0x5f, 0x50, 0x41, 0x52,
		0x54, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
		0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f, 0x42, 0x59,
		0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x00, 0x00,
		0x00, 0x01, 0x0c, 0x00, 0x01, 0x0b, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04, 0x74, 0x65, 0x6d, 0x70,
		0x0f, 0x00, 0x02, 0x08, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00,
	}
	arg := rrdb.NewMetaQueryCfgArgs()
	arg.Query = replication.NewQueryCfgRequest()
	arg.Query.AppName = "temp"
	arg.Query.PartitionIndices = []int32{}

	r := &rpcCall{
		args:  arg,
		name:  "RPC_CM_QUERY_PARTITION_CONFIG_BY_INDEX",
		gpid:  &base.Gpid{0, 0},
		seqId: 1,
	}

	actual, _ := PegasusCodec{}.Marshal(r)
	assert.Equal(t, expected, actual)
}

func TestCodec_UnmarshalErrorCode(t *testing.T) {
	recvBytes := []byte{
		0x54, 0x48, 0x46,
	}

	r := &rpcCall{}
	err := PegasusCodec{}.Unmarshal(recvBytes, r)
	assert.NotNil(t, err)
}