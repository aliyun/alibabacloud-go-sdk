// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenApiInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiName(v string) *DescribeOpenApiInfoRequest
	GetApiName() *string
	SetApiVersion(v string) *DescribeOpenApiInfoRequest
	GetApiVersion() *string
	SetLang(v string) *DescribeOpenApiInfoRequest
	GetLang() *string
	SetPopCode(v string) *DescribeOpenApiInfoRequest
	GetPopCode() *string
	SetRoleFor(v int64) *DescribeOpenApiInfoRequest
	GetRoleFor() *int64
	SetRoleType(v string) *DescribeOpenApiInfoRequest
	GetRoleType() *string
}

type DescribeOpenApiInfoRequest struct {
	// The operation that you want to perform.
	//
	// This parameter is required.
	//
	// example:
	//
	// DescribePopApiItemList
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The version number of the API.
	//
	// >  You can call the [DescribeGroupProductions](~~DescribeGroupProductions~~) operation to query version numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-03
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese.
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
	// 1592757xxx002956
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// 	- 0 (default): the view of the current Alibaba Cloud account.
	//
	// 	- 1: the view of all accounts for the enterprise.
	//
	// example:
	//
	// 0
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeOpenApiInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenApiInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenApiInfoRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeOpenApiInfoRequest) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *DescribeOpenApiInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOpenApiInfoRequest) GetPopCode() *string {
	return s.PopCode
}

func (s *DescribeOpenApiInfoRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeOpenApiInfoRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeOpenApiInfoRequest) SetApiName(v string) *DescribeOpenApiInfoRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetApiVersion(v string) *DescribeOpenApiInfoRequest {
	s.ApiVersion = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetLang(v string) *DescribeOpenApiInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetPopCode(v string) *DescribeOpenApiInfoRequest {
	s.PopCode = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetRoleFor(v int64) *DescribeOpenApiInfoRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) SetRoleType(v string) *DescribeOpenApiInfoRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeOpenApiInfoRequest) Validate() error {
	return dara.Validate(s)
}
