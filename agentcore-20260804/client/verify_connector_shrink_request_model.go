// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConnectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *VerifyConnectorShrinkRequest
	GetBodyShrink() *string
}

type VerifyConnectorShrinkRequest struct {
	// The validation request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyConnectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *VerifyConnectorShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *VerifyConnectorShrinkRequest) SetBodyShrink(v string) *VerifyConnectorShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *VerifyConnectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
