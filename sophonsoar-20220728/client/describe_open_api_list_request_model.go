// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribeOpenApiListRequest
	GetApiName() *string
	SetApiVersion(v string) *DescribeOpenApiListRequest
	GetApiVersion() *string
	SetLang(v string) *DescribeOpenApiListRequest
	GetLang() *string
	SetPopCode(v string) *DescribeOpenApiListRequest
	GetPopCode() *string
	SetRoleFor(v int64) *DescribeOpenApiListRequest
	GetRoleFor() *int64
	SetRoleType(v string) *DescribeOpenApiListRequest
	GetRoleType() *string
}

type DescribeOpenApiListRequest struct {
	// The operation that you want to perform.
	//
	// example:
	//
	// DescribePopApiItemList
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The version number of the API.
	//
	// >  You can call the [DescribeGroupProductions](~~DescribeGroupProductions~~) operation to query the version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-10-01
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The POP code of the Alibaba Cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sas
	PopCode *string `json:"PopCode,omitempty" xml:"PopCode,omitempty"`
	// The ID of the user who switches from the current view to the destination view by using the management account.
	//
	// example:
	//
	// 137602xxx8718726
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Default value: 0. Valid values:
	//
	// 	- 0: the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeOpenApiListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiListRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeOpenApiListRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *DescribeOpenApiListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOpenApiListRequest) GetPopCode() *string {
	return s.PopCode
}

func (s *DescribeOpenApiListRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeOpenApiListRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeOpenApiListRequest) SetApiName(v string) *DescribeOpenApiListRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetApiVersion(v string) *DescribeOpenApiListRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetLang(v string) *DescribeOpenApiListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetPopCode(v string) *DescribeOpenApiListRequest {
	s.PopCode = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetRoleFor(v int64) *DescribeOpenApiListRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOpenApiListRequest) SetRoleType(v string) *DescribeOpenApiListRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeOpenApiListRequest) Validate() error {
	return dara.Validate(s)
}
