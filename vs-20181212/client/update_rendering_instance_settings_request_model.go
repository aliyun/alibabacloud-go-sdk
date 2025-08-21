// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *UpdateRenderingInstanceSettingsRequest
	GetRenderingInstanceId() *string
	SetSettings(v []*UpdateRenderingInstanceSettingsRequestSettings) *UpdateRenderingInstanceSettingsRequest
	GetSettings() []*UpdateRenderingInstanceSettingsRequestSettings
}

type UpdateRenderingInstanceSettingsRequest struct {
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string                                           `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	Settings            []*UpdateRenderingInstanceSettingsRequestSettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s UpdateRenderingInstanceSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceSettingsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpdateRenderingInstanceSettingsRequest) GetSettings() []*UpdateRenderingInstanceSettingsRequestSettings {
	return s.Settings
}

func (s *UpdateRenderingInstanceSettingsRequest) SetRenderingInstanceId(v string) *UpdateRenderingInstanceSettingsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsRequest) SetSettings(v []*UpdateRenderingInstanceSettingsRequestSettings) *UpdateRenderingInstanceSettingsRequest {
	s.Settings = v
	return s
}

func (s *UpdateRenderingInstanceSettingsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRenderingInstanceSettingsRequestSettings struct {
	// example:
	//
	// navbar.hide
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// example:
	//
	// 1
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s UpdateRenderingInstanceSettingsRequestSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceSettingsRequestSettings) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceSettingsRequestSettings) GetAttributeName() *string {
	return s.AttributeName
}

func (s *UpdateRenderingInstanceSettingsRequestSettings) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *UpdateRenderingInstanceSettingsRequestSettings) SetAttributeName(v string) *UpdateRenderingInstanceSettingsRequestSettings {
	s.AttributeName = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsRequestSettings) SetAttributeValue(v string) *UpdateRenderingInstanceSettingsRequestSettings {
	s.AttributeValue = &v
	return s
}

func (s *UpdateRenderingInstanceSettingsRequestSettings) Validate() error {
	return dara.Validate(s)
}
