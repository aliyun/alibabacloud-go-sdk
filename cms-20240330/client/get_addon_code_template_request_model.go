// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonCodeTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *GetAddonCodeTemplateRequest
	GetAliyunLang() *string
	SetEnvironmentType(v string) *GetAddonCodeTemplateRequest
	GetEnvironmentType() *string
	SetVersion(v string) *GetAddonCodeTemplateRequest
	GetVersion() *string
}

type GetAddonCodeTemplateRequest struct {
	// example:
	//
	// zh
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
	// example:
	//
	// Client
	EnvironmentType *string `json:"environmentType,omitempty" xml:"environmentType,omitempty"`
	// example:
	//
	// 0.1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAddonCodeTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAddonCodeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetAddonCodeTemplateRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *GetAddonCodeTemplateRequest) GetEnvironmentType() *string {
	return s.EnvironmentType
}

func (s *GetAddonCodeTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *GetAddonCodeTemplateRequest) SetAliyunLang(v string) *GetAddonCodeTemplateRequest {
	s.AliyunLang = &v
	return s
}

func (s *GetAddonCodeTemplateRequest) SetEnvironmentType(v string) *GetAddonCodeTemplateRequest {
	s.EnvironmentType = &v
	return s
}

func (s *GetAddonCodeTemplateRequest) SetVersion(v string) *GetAddonCodeTemplateRequest {
	s.Version = &v
	return s
}

func (s *GetAddonCodeTemplateRequest) Validate() error {
	return dara.Validate(s)
}
