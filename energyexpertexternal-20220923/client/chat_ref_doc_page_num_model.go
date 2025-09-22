// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRefDocPageNum interface {
	dara.Model
	String() string
	GoString() string
	SetNum(v int32) *ChatRefDocPageNum
	GetNum() *int32
	SetPos(v [][]*ChatRefDocPostion) *ChatRefDocPageNum
	GetPos() [][]*ChatRefDocPostion
}

type ChatRefDocPageNum struct {
	Num *int32                 `json:"num,omitempty" xml:"num,omitempty"`
	Pos [][]*ChatRefDocPostion `json:"pos,omitempty" xml:"pos,omitempty" type:"Repeated"`
}

func (s ChatRefDocPageNum) String() string {
	return dara.Prettify(s)
}

func (s ChatRefDocPageNum) GoString() string {
	return s.String()
}

func (s *ChatRefDocPageNum) GetNum() *int32 {
	return s.Num
}

func (s *ChatRefDocPageNum) GetPos() [][]*ChatRefDocPostion {
	return s.Pos
}

func (s *ChatRefDocPageNum) SetNum(v int32) *ChatRefDocPageNum {
	s.Num = &v
	return s
}

func (s *ChatRefDocPageNum) SetPos(v [][]*ChatRefDocPostion) *ChatRefDocPageNum {
	s.Pos = v
	return s
}

func (s *ChatRefDocPageNum) Validate() error {
	return dara.Validate(s)
}
