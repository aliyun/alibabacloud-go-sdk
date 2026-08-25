// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateCredentialShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateCredentialShrinkRequest
	GetClientToken() *string
}

type UpdateCredentialShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCredentialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateCredentialShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialShrinkRequest) SetBodyShrink(v string) *UpdateCredentialShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateCredentialShrinkRequest) SetClientToken(v string) *UpdateCredentialShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
