// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemPropertyTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAuto(v bool) *ModifySystemPropertyTemplateShrinkRequest
	GetEnableAuto() *bool
	SetFilePath(v string) *ModifySystemPropertyTemplateShrinkRequest
	GetFilePath() *string
	SetSystemPropertyInfoShrink(v string) *ModifySystemPropertyTemplateShrinkRequest
	GetSystemPropertyInfoShrink() *string
	SetTemplateId(v string) *ModifySystemPropertyTemplateShrinkRequest
	GetTemplateId() *string
	SetTemplateName(v string) *ModifySystemPropertyTemplateShrinkRequest
	GetTemplateName() *string
}

type ModifySystemPropertyTemplateShrinkRequest struct {
	// Specifies whether to automatically generate preset system properties.
	//
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// The URL of the property template file. The system synchronously parses the file. If the file format is invalid, a parsing error is returned.
	//
	// > File template format: `{ "properties":{"key1":"value1", "key2":"value2"}}`.
	//
	// example:
	//
	// https://filepath****.com
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The information about the system property template.
	SystemPropertyInfoShrink *string `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty"`
	// The ID of the property template.
	//
	// example:
	//
	// ap-angyvganxlf****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ModifySystemPropertyTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemPropertyTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifySystemPropertyTemplateShrinkRequest) GetEnableAuto() *bool {
	return s.EnableAuto
}

func (s *ModifySystemPropertyTemplateShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ModifySystemPropertyTemplateShrinkRequest) GetSystemPropertyInfoShrink() *string {
	return s.SystemPropertyInfoShrink
}

func (s *ModifySystemPropertyTemplateShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifySystemPropertyTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifySystemPropertyTemplateShrinkRequest) SetEnableAuto(v bool) *ModifySystemPropertyTemplateShrinkRequest {
	s.EnableAuto = &v
	return s
}

func (s *ModifySystemPropertyTemplateShrinkRequest) SetFilePath(v string) *ModifySystemPropertyTemplateShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *ModifySystemPropertyTemplateShrinkRequest) SetSystemPropertyInfoShrink(v string) *ModifySystemPropertyTemplateShrinkRequest {
	s.SystemPropertyInfoShrink = &v
	return s
}

func (s *ModifySystemPropertyTemplateShrinkRequest) SetTemplateId(v string) *ModifySystemPropertyTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifySystemPropertyTemplateShrinkRequest) SetTemplateName(v string) *ModifySystemPropertyTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifySystemPropertyTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
