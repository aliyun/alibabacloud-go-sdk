// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSystemPropertyTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableAuto(v bool) *CreateSystemPropertyTemplateShrinkRequest
	GetEnableAuto() *bool
	SetFilePath(v string) *CreateSystemPropertyTemplateShrinkRequest
	GetFilePath() *string
	SetSystemPropertyInfoShrink(v string) *CreateSystemPropertyTemplateShrinkRequest
	GetSystemPropertyInfoShrink() *string
	SetTemplateName(v string) *CreateSystemPropertyTemplateShrinkRequest
	GetTemplateName() *string
}

type CreateSystemPropertyTemplateShrinkRequest struct {
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// example:
	//
	// https://filepath****.com
	FilePath                 *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	SystemPropertyInfoShrink *string `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty"`
	TemplateName             *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateSystemPropertyTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSystemPropertyTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSystemPropertyTemplateShrinkRequest) GetEnableAuto() *bool {
	return s.EnableAuto
}

func (s *CreateSystemPropertyTemplateShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateSystemPropertyTemplateShrinkRequest) GetSystemPropertyInfoShrink() *string {
	return s.SystemPropertyInfoShrink
}

func (s *CreateSystemPropertyTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSystemPropertyTemplateShrinkRequest) SetEnableAuto(v bool) *CreateSystemPropertyTemplateShrinkRequest {
	s.EnableAuto = &v
	return s
}

func (s *CreateSystemPropertyTemplateShrinkRequest) SetFilePath(v string) *CreateSystemPropertyTemplateShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *CreateSystemPropertyTemplateShrinkRequest) SetSystemPropertyInfoShrink(v string) *CreateSystemPropertyTemplateShrinkRequest {
	s.SystemPropertyInfoShrink = &v
	return s
}

func (s *CreateSystemPropertyTemplateShrinkRequest) SetTemplateName(v string) *CreateSystemPropertyTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSystemPropertyTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
