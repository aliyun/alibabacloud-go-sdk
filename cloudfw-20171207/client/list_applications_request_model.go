// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v int64) *ListApplicationsRequest
	GetApplicationId() *int64
	SetApplicationName(v string) *ListApplicationsRequest
	GetApplicationName() *string
	SetCategoryId(v int64) *ListApplicationsRequest
	GetCategoryId() *int64
	SetCategoryName(v string) *ListApplicationsRequest
	GetCategoryName() *string
	SetDescription(v string) *ListApplicationsRequest
	GetDescription() *string
	SetGroupName(v string) *ListApplicationsRequest
	GetGroupName() *string
	SetLang(v string) *ListApplicationsRequest
	GetLang() *string
	SetPageNo(v int64) *ListApplicationsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationsRequest
	GetPageSize() *int64
	SetRisk(v string) *ListApplicationsRequest
	GetRisk() *string
	SetTechnology(v string) *ListApplicationsRequest
	GetTechnology() *string
}

type ListApplicationsRequest struct {
	// The application ID. Exact match is supported.
	//
	// example:
	//
	// 10012150000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application name. Fuzzy match is supported.
	//
	// example:
	//
	// DingTalk
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application category ID. Exact match is supported.
	//
	// example:
	//
	// 8
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The application category name. Fuzzy match is supported.
	//
	// example:
	//
	// Office collaboration
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The application description. Fuzzy match is supported.
	//
	// example:
	//
	// Office
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application group name. Fuzzy match is supported.
	//
	// example:
	//
	// Dingtalk
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The language of the returned information. Valid values:
	//
	// - **zh**: Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// This parameter determines whether the application name, category name, technology type, vendor, and description are returned in Chinese or English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than or equal to 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level. Exact match is supported. Valid values:
	//
	// - **Information**: Informational.
	//
	// - **Low**: Low.
	//
	// - **Medium**: Medium.
	//
	// - **High**: High.
	//
	// - **Critical**: Critical.
	//
	// If an invalid value is passed, this filter condition is ignored.
	//
	// example:
	//
	// Low
	Risk *string `json:"Risk,omitempty" xml:"Risk,omitempty"`
	// The technology type name. Fuzzy match is supported.
	//
	// example:
	//
	// Browser-based
	Technology *string `json:"Technology,omitempty" xml:"Technology,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListApplicationsRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListApplicationsRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListApplicationsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListApplicationsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListApplicationsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationsRequest) GetRisk() *string {
	return s.Risk
}

func (s *ListApplicationsRequest) GetTechnology() *string {
	return s.Technology
}

func (s *ListApplicationsRequest) SetApplicationId(v int64) *ListApplicationsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsRequest) SetApplicationName(v string) *ListApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsRequest) SetCategoryId(v int64) *ListApplicationsRequest {
	s.CategoryId = &v
	return s
}

func (s *ListApplicationsRequest) SetCategoryName(v string) *ListApplicationsRequest {
	s.CategoryName = &v
	return s
}

func (s *ListApplicationsRequest) SetDescription(v string) *ListApplicationsRequest {
	s.Description = &v
	return s
}

func (s *ListApplicationsRequest) SetGroupName(v string) *ListApplicationsRequest {
	s.GroupName = &v
	return s
}

func (s *ListApplicationsRequest) SetLang(v string) *ListApplicationsRequest {
	s.Lang = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNo(v int64) *ListApplicationsRequest {
	s.PageNo = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int64) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetRisk(v string) *ListApplicationsRequest {
	s.Risk = &v
	return s
}

func (s *ListApplicationsRequest) SetTechnology(v string) *ListApplicationsRequest {
	s.Technology = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
