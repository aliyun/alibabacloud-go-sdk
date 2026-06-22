// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBuildRiskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageBuildRiskListRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageBuildRiskListRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageBuildRiskListRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeImageBuildRiskListRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageBuildRiskListRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeImageBuildRiskListRequest
	GetRiskLevel() *string
}

type DescribeImageBuildRiskListRequest struct {
	// The **rule name*	- or **category name*	- of the build risk. You can call the [DescribeImageBuildRiskList](~~~~) operation to obtain the value. Valid values:
	//
	// - If **CriteriaType*	- is set to **RiskKeyName**, the value is the **rule name*	- of the build risk.
	//
	// - If **CriteriaType*	- is set to **RiskClassName**, the value is the **category name*	- of the build risk.
	//
	// example:
	//
	// no_user
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The query type of the build risk. Valid values:
	//
	// - **RiskKeyName**: build risk rule name.
	//
	// - **RiskClassName**: build risk category name.
	//
	// example:
	//
	// RiskKeyName
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The page number of the current page when paging is used. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content in the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page when paging is used. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level. Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeImageBuildRiskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBuildRiskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBuildRiskListRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageBuildRiskListRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageBuildRiskListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBuildRiskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBuildRiskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBuildRiskListRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageBuildRiskListRequest) SetCriteria(v string) *DescribeImageBuildRiskListRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) SetCriteriaType(v string) *DescribeImageBuildRiskListRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) SetCurrentPage(v int32) *DescribeImageBuildRiskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) SetLang(v string) *DescribeImageBuildRiskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) SetPageSize(v int32) *DescribeImageBuildRiskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) SetRiskLevel(v string) *DescribeImageBuildRiskListRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageBuildRiskListRequest) Validate() error {
	return dara.Validate(s)
}
