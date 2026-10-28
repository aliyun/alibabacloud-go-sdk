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
	SetApiVersion(v string) *GetAddonSchemaRequest
	GetApiVersion() *string
	SetEnvironmentType(v string) *GetAddonSchemaRequest
	GetEnvironmentType() *string
	SetVersion(v string) *GetAddonSchemaRequest
	GetVersion() *string
}

type GetAddonSchemaRequest struct {
	// The language environment. Default value: zh. Valid values: zh, en.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// The add-on catalog version. Valid values: v1, v2.
	//
	// example:
	//
	// v2
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The environment type. Valid values: CS (container), ECS.
	//
	// example:
	//
	// CS
	EnvironmentType *string `json:"environmentType,omitempty" xml:"environmentType,omitempty"`
	// The component version.
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

func (s *GetAddonSchemaRequest) GetApiVersion() *string {
	return s.ApiVersion
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

func (s *GetAddonSchemaRequest) SetApiVersion(v string) *GetAddonSchemaRequest {
	s.ApiVersion = &v
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
