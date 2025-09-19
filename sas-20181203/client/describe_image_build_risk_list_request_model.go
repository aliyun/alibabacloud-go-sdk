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
	// The **rule name*	- or **type name*	- of the risk. You can call the [DescribeImageBuildRiskList](~~~~) operation to obtain the name. Optional parameters:
	//
	// 	- If **CriteriaType*	- is set to **RiskKeyName**, you must specify a **rule name*	- for this parameter.
	//
	// 	- If **CriteriaType*	- is set to**RiskClassName**, you must specify a **type name*	- for this parameter.
	//
	// example:
	//
	// no_user
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The query type.of the risk. Valid values:
	//
	// 	- **RiskKeyName**: the rule name of the risk
	//
	// 	- **RiskClassName**: the type name of the risk
	//
	// example:
	//
	// RiskKeyName
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
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
