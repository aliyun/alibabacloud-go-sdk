// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceConfigurationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigurationShrink(v string) *DeleteRenderingInstanceConfigurationShrinkRequest
	GetConfigurationShrink() *string
	SetRenderingInstanceId(v string) *DeleteRenderingInstanceConfigurationShrinkRequest
	GetRenderingInstanceId() *string
}

type DeleteRenderingInstanceConfigurationShrinkRequest struct {
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DeleteRenderingInstanceConfigurationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceConfigurationShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *DeleteRenderingInstanceConfigurationShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DeleteRenderingInstanceConfigurationShrinkRequest) SetConfigurationShrink(v string) *DeleteRenderingInstanceConfigurationShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationShrinkRequest) SetRenderingInstanceId(v string) *DeleteRenderingInstanceConfigurationShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
