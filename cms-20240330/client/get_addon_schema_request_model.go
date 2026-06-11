// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetAddonSchemaRequest
	GetAliyunLang() *string
	SetEnvironmentType(v string) *GetAddonSchemaRequest
	GetEnvironmentType() *string
	SetVersion(v string) *GetAddonSchemaRequest
	GetVersion() *string
}

type GetAddonSchemaRequest struct {
	// The language of the response. Valid values: \\`zh\\` and \\`en\\`. The default value is \\`zh\\`.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// The environment type. Valid values: \\`CS\\` (container) and \\`ECS\\`.
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"environmentType,omitempty" xml:"environmentType,omitempty"`
	// The version of the component.
	//
	// example:
	//
	// 0.1.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAddonSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetAddonSchemaRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *GetAddonSchemaRequest) GetVersion() *string {
	return s.Version
}

func (s *GetAddonSchemaRequest) SetAliyunLang(v string) *GetAddonSchemaRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetAddonSchemaRequest) SetEnvironmentType(v string) *GetAddonSchemaRequest {
	s.EnvironmentType = &v
	return s
}

func (s *GetAddonSchemaRequest) SetVersion(v string) *GetAddonSchemaRequest {
	s.Version = &v
	return s
}

func (s *GetAddonSchemaRequest) Validate() error {
	return dara.Validate(s)
}
