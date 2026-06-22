// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPeekMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPeekonly(v bool) *PeekMessageRequest
	GetPeekonly() *bool
}

type PeekMessageRequest struct {
	Peekonly *bool `json:"peekonly,omitempty" xml:"peekonly,omitempty"`
}

func (s PeekMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s PeekMessageRequest) GoString() string {
	return s.String()
}

func (s *PeekMessageRequest) GetPeekonly() *bool {
	return s.Peekonly
}

func (s *PeekMessageRequest) SetPeekonly(v bool) *PeekMessageRequest {
	s.Peekonly = &v
	return s
}

func (s *PeekMessageRequest) Validate() error {
	return dara.Validate(s)
}
