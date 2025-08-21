// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageId(v int64) *ReadMessageShrinkRequest
	GetMessageId() *int64
	SetUserInfoShrink(v string) *ReadMessageShrinkRequest
	GetUserInfoShrink() *string
}

type ReadMessageShrinkRequest struct {
	// example:
	//
	// 12345
	MessageId *int64 `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s ReadMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReadMessageShrinkRequest) GetMessageId() *int64 {
	return s.MessageId
}

func (s *ReadMessageShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ReadMessageShrinkRequest) SetMessageId(v int64) *ReadMessageShrinkRequest {
	s.MessageId = &v
	return s
}

func (s *ReadMessageShrinkRequest) SetUserInfoShrink(v string) *ReadMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ReadMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
