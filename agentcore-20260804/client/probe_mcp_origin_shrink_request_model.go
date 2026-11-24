// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeMcpOriginShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ProbeMcpOriginShrinkRequest
	GetBodyShrink() *string
}

type ProbeMcpOriginShrinkRequest struct {
	// The scope of the origin server configuration must not exceed that of the creation interface. The control fields for this validation only restrict the validation scope and do not extend the origin server access configuration.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProbeMcpOriginShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginShrinkRequest) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ProbeMcpOriginShrinkRequest) SetBodyShrink(v string) *ProbeMcpOriginShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ProbeMcpOriginShrinkRequest) Validate() error {
	return dara.Validate(s)
}
