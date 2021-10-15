package entities

import (
	"context"
	"fmt"
	ants "github.com/panjf2000/ants/v2"
	"testing"
)

func TestEntity(t *testing.T) {

	coroutinePool, err := ants.NewPool(500)
	if nil != err {
		panic(err)
	}

	tag := "test"
	mgr := NewEntityManager(context.Background(), coroutinePool)

	enty1, err := NewEntity(context.Background(), mgr, "", "abcd", "tomas", &tag, 001)
	enty2, err := NewEntity(context.Background(), mgr, "", "abcd", "tomas", &tag, 001)

	fmt.Println(enty1, enty2, err)
}