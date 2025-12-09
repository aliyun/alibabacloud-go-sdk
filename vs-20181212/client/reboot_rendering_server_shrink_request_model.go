// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingServerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceIdsShrink(v string) *RebootRenderingServerShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type RebootRenderingServerShrinkRequest struct {
	// This parameter is required.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s RebootRenderingServerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *RebootRenderingServerShrinkRequest) SetRenderingInstanceIdsShrink(v string) *RebootRenderingServerShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *RebootRenderingServerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
