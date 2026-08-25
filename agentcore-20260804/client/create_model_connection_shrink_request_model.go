// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelConnectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *CreateModelConnectionShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *CreateModelConnectionShrinkRequest
	GetClientToken() *string
}

type CreateModelConnectionShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateModelConnectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *CreateModelConnectionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateModelConnectionShrinkRequest) SetBodyShrink(v string) *CreateModelConnectionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *CreateModelConnectionShrinkRequest) SetClientToken(v string) *CreateModelConnectionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModelConnectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
