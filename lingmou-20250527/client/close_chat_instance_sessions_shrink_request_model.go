// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseChatInstanceSessionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionIdsShrink(v string) *CloseChatInstanceSessionsShrinkRequest
	GetSessionIdsShrink() *string
}

type CloseChatInstanceSessionsShrinkRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// ["8C9F2D4E-7A6B-4F1C-9E53-0B2C8D3F9A4B"]
	SessionIdsShrink *string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty"`
}

func (s CloseChatInstanceSessionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseChatInstanceSessionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CloseChatInstanceSessionsShrinkRequest) GetSessionIdsShrink() *string {
	return s.SessionIdsShrink
}

func (s *CloseChatInstanceSessionsShrinkRequest) SetSessionIdsShrink(v string) *CloseChatInstanceSessionsShrinkRequest {
	s.SessionIdsShrink = &v
	return s
}

func (s *CloseChatInstanceSessionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
