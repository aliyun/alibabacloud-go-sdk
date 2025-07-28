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
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// example:
	//
	// https://filepath****.com
	FilePath                 *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	SystemPropertyInfoShrink *string `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty"`
	// example:
	//
	// ap-angyvganxlf****
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
