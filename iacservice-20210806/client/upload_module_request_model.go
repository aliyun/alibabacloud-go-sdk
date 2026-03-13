// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v map[string]*string) *UploadModuleRequest
	GetCode() map[string]*string
	SetModuleId(v string) *UploadModuleRequest
	GetModuleId() *string
	SetModuleName(v string) *UploadModuleRequest
	GetModuleName() *string
	SetNamespaceName(v string) *UploadModuleRequest
	GetNamespaceName() *string
	SetUrl(v string) *UploadModuleRequest
	GetUrl() *string
}

type UploadModuleRequest struct {
	Code map[string]*string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// mod-kw1018ogp2m3qp22b3k31d
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// RegistryModule-test0ef88
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// test_namespace
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// url
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadModuleRequest) GoString() string {
	return s.String()
}

func (s *UploadModuleRequest) GetCode() map[string]*string {
	return s.Code
}

func (s *UploadModuleRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *UploadModuleRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *UploadModuleRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UploadModuleRequest) GetUrl() *string {
	return s.Url
}

func (s *UploadModuleRequest) SetCode(v map[string]*string) *UploadModuleRequest {
	s.Code = v
	return s
}

func (s *UploadModuleRequest) SetModuleId(v string) *UploadModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *UploadModuleRequest) SetModuleName(v string) *UploadModuleRequest {
	s.ModuleName = &v
	return s
}

func (s *UploadModuleRequest) SetNamespaceName(v string) *UploadModuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *UploadModuleRequest) SetUrl(v string) *UploadModuleRequest {
	s.Url = &v
	return s
}

func (s *UploadModuleRequest) Validate() error {
	return dara.Validate(s)
}
