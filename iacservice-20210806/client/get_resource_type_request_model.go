// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetResourceTypeRequest
	GetAcceptLanguage() *string
	SetFilterReadOnly(v bool) *GetResourceTypeRequest
	GetFilterReadOnly() *bool
	SetTerraformProviderVersion(v string) *GetResourceTypeRequest
	GetTerraformProviderVersion() *string
}

type GetResourceTypeRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// Default value: zh-CN.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// Specifies whether to filter out read-only properties.
	//
	// example:
	//
	// false
	FilterReadOnly *bool `json:"filterReadOnly,omitempty" xml:"filterReadOnly,omitempty"`
	// The Terraform provider version. If left empty, the latest version is used by default.
	//
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s GetResourceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetResourceTypeRequest) GetFilterReadOnly() *bool {
	return s.FilterReadOnly
}

func (s *GetResourceTypeRequest) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GetResourceTypeRequest) SetAcceptLanguage(v string) *GetResourceTypeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetResourceTypeRequest) SetFilterReadOnly(v bool) *GetResourceTypeRequest {
	s.FilterReadOnly = &v
	return s
}

func (s *GetResourceTypeRequest) SetTerraformProviderVersion(v string) *GetResourceTypeRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetResourceTypeRequest) Validate() error {
	return dara.Validate(s)
}
