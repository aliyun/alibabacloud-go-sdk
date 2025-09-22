// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRefDocPostion interface {
	dara.Model
	String() string
	GoString() string
	SetX(v int32) *ChatRefDocPostion
	GetX() *int32
	SetY(v int32) *ChatRefDocPostion
	GetY() *int32
}

type ChatRefDocPostion struct {
	X *int32 `json:"x,omitempty" xml:"x,omitempty"`
	Y *int32 `json:"y,omitempty" xml:"y,omitempty"`
}

func (s ChatRefDocPostion) String() string {
	return dara.Prettify(s)
}

func (s ChatRefDocPostion) GoString() string {
	return s.String()
}

func (s *ChatRefDocPostion) GetX() *int32 {
	return s.X
}

func (s *ChatRefDocPostion) GetY() *int32 {
	return s.Y
}

func (s *ChatRefDocPostion) SetX(v int32) *ChatRefDocPostion {
	s.X = &v
	return s
}

func (s *ChatRefDocPostion) SetY(v int32) *ChatRefDocPostion {
	s.Y = &v
	return s
}

func (s *ChatRefDocPostion) Validate() error {
	return dara.Validate(s)
}
