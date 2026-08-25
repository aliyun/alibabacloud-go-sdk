// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelConnectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateModelConnectionShrinkRequest
	GetBodyShrink() *string
	SetClientToken(v string) *UpdateModelConnectionShrinkRequest
	GetClientToken() *string
}

type UpdateModelConnectionShrinkRequest struct {
	BodyShrink  *string `json:"body,omitempty" xml:"body,omitempty"`
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateModelConnectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelConnectionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateModelConnectionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelConnectionShrinkRequest) SetBodyShrink(v string) *UpdateModelConnectionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateModelConnectionShrinkRequest) SetClientToken(v string) *UpdateModelConnectionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelConnectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
