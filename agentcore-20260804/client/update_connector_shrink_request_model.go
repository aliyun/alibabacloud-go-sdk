// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UpdateConnectorShrinkRequest
	GetBodyShrink() *string
}

type UpdateConnectorShrinkRequest struct {
	// The update request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectorShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UpdateConnectorShrinkRequest) SetBodyShrink(v string) *UpdateConnectorShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UpdateConnectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
