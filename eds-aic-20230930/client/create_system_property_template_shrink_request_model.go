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
	// Specifies whether to automatically generate preset system properties.
	//
	// example:
	//
	// true
	EnableAuto *bool `json:"EnableAuto,omitempty" xml:"EnableAuto,omitempty"`
	// The URL of the property template file. The API parses the file synchronously. An error is returned if the file format is invalid.
	//
	// > The file must be in the following format: `{ "properties":{"key1":"value1"}}`.
	//
	// example:
	//
	// https://filepath****.com
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The information about the system property template.
	SystemPropertyInfoShrink *string `json:"SystemPropertyInfo,omitempty" xml:"SystemPropertyInfo,omitempty"`
	// The name of the template. The name must meet the following requirements:
	//
	// - Be 2 to 32 characters in length.
	//
	// - Start with an uppercase or lowercase letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - Contain letters, digits, colons (:), underscores (_), and hyphens (-). Periods (.) are not supported.
	//
	// example:
	//
	// Template1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
