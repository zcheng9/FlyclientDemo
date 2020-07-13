package flyclientdemo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zcheng9/FlyclientDemo/mmr"
	"math/big"
	"testing"
	"time"
)

func TestBlockChain_GetProof(t *testing.T) {
	bc := NewBlockChain()
	length := 50000
	//first creat a blockchain with length of 50000,you can change the length
	for i := 0; i < length; i++ {
		b := NewBlock(uint64(i), 2, big.NewInt(10000))
		bc.InsertBlock(b)
	}
	
	//generate proof 
	start := time.Now()
	proof := bc.GetProof()
	fmt.Printf("size of blockchain:%d,size of proof:%d\n",length,len(proof.Elems))
	
	//verify proof with limited information
	fmt.Println("gen proof cost:", time.Now().Sub(start))
	start = time.Now()
	pBlocks, err := mmr.VerifyRequiredBlocks(proof, RightDif)
	assert.NoError(t, err)
	assert.True(t, proof.VerifyProof(pBlocks))

	fmt.Println("verify cost:", time.Now().Sub(start))

	//var wg sync.WaitGroup
	//fmt.Println("waiting for all goroutine ")
	//for n := 0; n < 1000; n++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//
	//	}()
	//}
	//wg.Wait()
	//fmt.Println("All goroutines finished!")
}
