// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatInstanceSessionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionIdsShrink(v string) *QueryChatInstanceSessionsShrinkRequest
	GetSessionIdsShrink() *string
}

type QueryChatInstanceSessionsShrinkRequest struct {
	// if can be null:
	// true
	//
	// example:
	//
	// ["8C9F2D4E-7A6B-4F1C-9E53-0B2C8D3F9A4B"]
	SessionIdsShrink *string `json:"sessionIds,omitempty" xml:"sessionIds,omitempty"`
}

func (s QueryChatInstanceSessionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryChatInstanceSessionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryChatInstanceSessionsShrinkRequest) GetSessionIdsShrink() *string {
	return s.SessionIdsShrink
}

func (s *QueryChatInstanceSessionsShrinkRequest) SetSessionIdsShrink(v string) *QueryChatInstanceSessionsShrinkRequest {
	s.SessionIdsShrink = &v
	return s
}

func (s *QueryChatInstanceSessionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
