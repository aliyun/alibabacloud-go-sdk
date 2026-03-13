// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadModuleAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v map[string]*string) *UploadModuleAdvanceRequest
	GetCode() map[string]*string
	SetModuleId(v string) *UploadModuleAdvanceRequest
	GetModuleId() *string
	SetModuleName(v string) *UploadModuleAdvanceRequest
	GetModuleName() *string
	SetNamespaceName(v string) *UploadModuleAdvanceRequest
	GetNamespaceName() *string
	SetUrlObject(v io.Reader) *UploadModuleAdvanceRequest
	GetUrlObject() io.Reader
}

type UploadModuleAdvanceRequest struct {
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
	UrlObject io.Reader `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadModuleAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadModuleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadModuleAdvanceRequest) GetCode() map[string]*string {
	return s.Code
}

func (s *UploadModuleAdvanceRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *UploadModuleAdvanceRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *UploadModuleAdvanceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UploadModuleAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *UploadModuleAdvanceRequest) SetCode(v map[string]*string) *UploadModuleAdvanceRequest {
	s.Code = v
	return s
}

func (s *UploadModuleAdvanceRequest) SetModuleId(v string) *UploadModuleAdvanceRequest {
	s.ModuleId = &v
	return s
}

func (s *UploadModuleAdvanceRequest) SetModuleName(v string) *UploadModuleAdvanceRequest {
	s.ModuleName = &v
	return s
}

func (s *UploadModuleAdvanceRequest) SetNamespaceName(v string) *UploadModuleAdvanceRequest {
	s.NamespaceName = &v
	return s
}

func (s *UploadModuleAdvanceRequest) SetUrlObject(v io.Reader) *UploadModuleAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *UploadModuleAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
