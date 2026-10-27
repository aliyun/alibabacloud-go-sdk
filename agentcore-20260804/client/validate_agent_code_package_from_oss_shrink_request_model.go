// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAgentCodePackageFromOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ValidateAgentCodePackageFromOssShrinkRequest
	GetBodyShrink() *string
}

type ValidateAgentCodePackageFromOssShrinkRequest struct {
	// The request body.
	//
	// This parameter is required.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateAgentCodePackageFromOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ValidateAgentCodePackageFromOssShrinkRequest) SetBodyShrink(v string) *ValidateAgentCodePackageFromOssShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}
