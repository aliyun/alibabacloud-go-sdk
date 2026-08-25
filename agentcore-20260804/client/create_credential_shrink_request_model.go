// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateCredentialShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateCredentialShrinkRequest
	GetClientToken() *string
}

type CreateCredentialShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 暂不支持
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCredentialShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateCredentialShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialShrinkRequest) SetBodyShrink(v string) *CreateCredentialShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateCredentialShrinkRequest) SetClientToken(v string) *CreateCredentialShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialShrinkRequest) Validate() error {
	return dara.Validate(s)
}
