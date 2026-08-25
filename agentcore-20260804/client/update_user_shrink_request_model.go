// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateUserShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateUserShrinkRequest
	GetClientToken() *string
}

type UpdateUserShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateUserShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateUserShrinkRequest) SetBodyShrink(v string) *UpdateUserShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateUserShrinkRequest) SetClientToken(v string) *UpdateUserShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
