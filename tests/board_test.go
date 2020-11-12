package shiko

import (
  "github.com/obsessedyouth/shiko"
  "github.com/stretchr/testify/assert"
  "testing"
  )

func TestBoard(t *testing.T){
  assert.Equal(t, "Super awesome chess engine :)", shiko.Board())
}
