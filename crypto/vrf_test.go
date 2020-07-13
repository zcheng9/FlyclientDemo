package crypto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/zcheng9/FlyclientDemo/common"
	"github.com/zcheng9/FlyclientDemo/common/hexutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testPriv1 = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232031"

func TestEvaluate(t *testing.T) {
	key1, err1 := HexToECDSA(testPriv1)
	assert.NoError(t, err1)
	seed := []byte{1, 2}
	vrf, nizk, err := VRFProve(key1, seed)
	fmt.Println(hexutil.Encode(vrf[:]) )
	fmt.Println(hexutil.Encode(nizk[:]) )
	fmt.Println(err)

	res := append(nizk, vrf...)

	fmt.Println(VRFVerify(&key1.PublicKey, seed, res))
}


func BytesToHash(b []byte) common.Hash {
	var a common.Hash
	a.SetBytes(b)
	return a
}
func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func TestVRFProve(t *testing.T) {
	seed := IntToBytes(1)
	fmt.Println(hashToInt(seed))
	fmt.Println(hashToCurve(seed))
}