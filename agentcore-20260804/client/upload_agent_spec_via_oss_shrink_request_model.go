// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSpecViaOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *UploadAgentSpecViaOssShrinkRequest
	GetBodyShrink() *string
}

type UploadAgentSpecViaOssShrinkRequest struct {
	// The request body.
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAgentSpecViaOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSpecViaOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadAgentSpecViaOssShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *UploadAgentSpecViaOssShrinkRequest) SetBodyShrink(v string) *UploadAgentSpecViaOssShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *UploadAgentSpecViaOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}
