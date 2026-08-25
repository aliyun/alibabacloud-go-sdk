// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateUserShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateUserShrinkRequest
	GetClientToken() *string
}

type CreateUserShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateUserShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateUserShrinkRequest) SetBodyShrink(v string) *CreateUserShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateUserShrinkRequest) SetClientToken(v string) *CreateUserShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
