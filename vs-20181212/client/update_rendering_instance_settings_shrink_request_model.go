// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *UpdateRenderingInstanceSettingsShrinkRequest
	GetRenderingInstanceId() *string
	SetSettingsShrink(v string) *UpdateRenderingInstanceSettingsShrinkRequest
	GetSettingsShrink() *string
}

type UpdateRenderingInstanceSettingsShrinkRequest struct {
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	SettingsShrink      *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s UpdateRenderingInstanceSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceSettingsShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpdateRenderingInstanceSettingsShrinkRequest) GetSettingsShrink() *string {
	return s.SettingsShrink
}

func (s *UpdateRenderingInstanceSettingsShrinkRequest) SetRenderingInstanceId(v string) *UpdateRenderingInstanceSettingsShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsShrinkRequest) SetSettingsShrink(v string) *UpdateRenderingInstanceSettingsShrinkRequest {
	s.SettingsShrink = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
