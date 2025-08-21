// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationShrink(v string) *UpdateRenderingInstanceConfigurationShrinkRequest
	GetConfigurationShrink() *string
	SetRenderingInstanceId(v string) *UpdateRenderingInstanceConfigurationShrinkRequest
	GetRenderingInstanceId() *string
}

type UpdateRenderingInstanceConfigurationShrinkRequest struct {
	// This parameter is required.
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *UpdateRenderingInstanceConfigurationShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpdateRenderingInstanceConfigurationShrinkRequest) SetConfigurationShrink(v string) *UpdateRenderingInstanceConfigurationShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationShrinkRequest) SetRenderingInstanceId(v string) *UpdateRenderingInstanceConfigurationShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
