// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportNacosConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ImportNacosConfigRequest
	GetAcceptLanguage() *string
	SetFileUrl(v string) *ImportNacosConfigRequest
	GetFileUrl() *string
	SetInstanceId(v string) *ImportNacosConfigRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *ImportNacosConfigRequest
	GetNamespaceId() *string
	SetPolicy(v string) *ImportNacosConfigRequest
	GetPolicy() *string
}

type ImportNacosConfigRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// This parameter is required.
	//
	// example:
	//
	// http://mse-shared.oss-xxx.aliyuncs.com/cfg/import/xxxx/2021/01/11/xxxx.zip
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The ID of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// mse_prepaid_public_cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The policy.
	//
	// example:
	//
	// f5cdc80a-****-8094-282f5650fc00
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The policy.
	//
	// Valid values:
	//
	// 	- ABORT
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- OVERWRITE
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- SKIP
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// overwrite
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ImportNacosConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportNacosConfigRequest) GoString() string {
	return s.String()
}

func (s *ImportNacosConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ImportNacosConfigRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImportNacosConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportNacosConfigRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ImportNacosConfigRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ImportNacosConfigRequest) SetAcceptLanguage(v string) *ImportNacosConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ImportNacosConfigRequest) SetFileUrl(v string) *ImportNacosConfigRequest {
	s.FileUrl = &v
	return s
}

func (s *ImportNacosConfigRequest) SetInstanceId(v string) *ImportNacosConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportNacosConfigRequest) SetNamespaceId(v string) *ImportNacosConfigRequest {
	s.NamespaceId = &v
	return s
}

func (s *ImportNacosConfigRequest) SetPolicy(v string) *ImportNacosConfigRequest {
	s.Policy = &v
	return s
}

func (s *ImportNacosConfigRequest) Validate() error {
	return dara.Validate(s)
}
