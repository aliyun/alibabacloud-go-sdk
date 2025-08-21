// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeNamesShrink(v string) *DeleteRenderingInstanceSettingsShrinkRequest
	GetAttributeNamesShrink() *string
	SetRenderingInstanceId(v string) *DeleteRenderingInstanceSettingsShrinkRequest
	GetRenderingInstanceId() *string
}

type DeleteRenderingInstanceSettingsShrinkRequest struct {
	AttributeNamesShrink *string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DeleteRenderingInstanceSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceSettingsShrinkRequest) GetAttributeNamesShrink() *string {
	return s.AttributeNamesShrink
}

func (s *DeleteRenderingInstanceSettingsShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DeleteRenderingInstanceSettingsShrinkRequest) SetAttributeNamesShrink(v string) *DeleteRenderingInstanceSettingsShrinkRequest {
	s.AttributeNamesShrink = &v
	return s
}

func (s *DeleteRenderingInstanceSettingsShrinkRequest) SetRenderingInstanceId(v string) *DeleteRenderingInstanceSettingsShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DeleteRenderingInstanceSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
