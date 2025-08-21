// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeforeTime(v string) *ListUserMessageShrinkRequest
	GetBeforeTime() *string
	SetUserInfoShrink(v string) *ListUserMessageShrinkRequest
	GetUserInfoShrink() *string
	SetLimit(v int32) *ListUserMessageShrinkRequest
	GetLimit() *int32
}

type ListUserMessageShrinkRequest struct {
	// example:
	//
	// 2022-07-27 14:06:55.984
	BeforeTime *string `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s ListUserMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListUserMessageShrinkRequest) GetBeforeTime() *string {
	return s.BeforeTime
}

func (s *ListUserMessageShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *ListUserMessageShrinkRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUserMessageShrinkRequest) SetBeforeTime(v string) *ListUserMessageShrinkRequest {
	s.BeforeTime = &v
	return s
}

func (s *ListUserMessageShrinkRequest) SetUserInfoShrink(v string) *ListUserMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *ListUserMessageShrinkRequest) SetLimit(v int32) *ListUserMessageShrinkRequest {
	s.Limit = &v
	return s
}

func (s *ListUserMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
