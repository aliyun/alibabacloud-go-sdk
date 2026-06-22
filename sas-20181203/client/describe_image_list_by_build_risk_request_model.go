// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListByBuildRiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageListByBuildRiskRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageListByBuildRiskRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageListByBuildRiskRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeImageListByBuildRiskRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageListByBuildRiskRequest
	GetPageSize() *int32
	SetRiskKey(v string) *DescribeImageListByBuildRiskRequest
	GetRiskKey() *string
	SetRiskLevel(v string) *DescribeImageListByBuildRiskRequest
	GetRiskLevel() *string
	SetStatus(v int32) *DescribeImageListByBuildRiskRequest
	GetStatus() *int32
}

type DescribeImageListByBuildRiskRequest struct {
	// The value of the condition parameter.
	//
	// example:
	//
	// sas
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The name of the condition parameter. Valid values:
	//
	// - **RepoNamespace**: namespace.
	//
	// - **RepoName**: repository name.
	//
	// example:
	//
	// RepoNamespace
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The page number of the current page in paging query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page in paging query. Default value: 20. If you leave this parameter empty, 20 entries are returned.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The key of the build risk rule. You can call the [DescribeImageBuildRiskList](~~~~) operation to obtain the RiskKey.
	//
	// example:
	//
	// no_user
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The risk level. Valid values:
	//
	// - **high**
	//
	// - **medium**
	//
	// - **low**.
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the alert event to query. Valid values:
	//
	// - **0**: Unhandled.
	//
	// - **1**: Ignored.
	//
	// - **2**: False positive.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageListByBuildRiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageListByBuildRiskRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageListByBuildRiskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageListByBuildRiskRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageListByBuildRiskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageListByBuildRiskRequest) GetRiskKey() *string {
	return s.RiskKey
}

func (s *DescribeImageListByBuildRiskRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageListByBuildRiskRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageListByBuildRiskRequest) SetCriteria(v string) *DescribeImageListByBuildRiskRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetCriteriaType(v string) *DescribeImageListByBuildRiskRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetCurrentPage(v int32) *DescribeImageListByBuildRiskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetLang(v string) *DescribeImageListByBuildRiskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetPageSize(v int32) *DescribeImageListByBuildRiskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetRiskKey(v string) *DescribeImageListByBuildRiskRequest {
	s.RiskKey = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetRiskLevel(v string) *DescribeImageListByBuildRiskRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) SetStatus(v int32) *DescribeImageListByBuildRiskRequest {
	s.Status = &v
	return s
}

func (s *DescribeImageListByBuildRiskRequest) Validate() error {
	return dara.Validate(s)
}
