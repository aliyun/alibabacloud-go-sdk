// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeImageBaselineCheckSummaryRequest
	GetClusterId() *string
	SetCriteria(v string) *DescribeImageBaselineCheckSummaryRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageBaselineCheckSummaryRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageBaselineCheckSummaryRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeImageBaselineCheckSummaryRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageBaselineCheckSummaryRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeImageBaselineCheckSummaryRequest
	GetRiskLevel() *string
	SetScanRange(v []*string) *DescribeImageBaselineCheckSummaryRequest
	GetScanRange() []*string
}

type DescribeImageBaselineCheckSummaryRequest struct {
	// The ID of the specified container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the ID of the container cluster.
	//
	// example:
	//
	// c60b77fe62093480db6164a3c2fa5****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The search condition for the image baseline.
	//
	// example:
	//
	// Unauthorized access
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **BaselineNameAlias**: baseline name
	//
	// 	- **BaselineClassAlias**: baseline category
	//
	// example:
	//
	// BaselineNameAlias
	CriteriaType *string `json:"CriteriaType,omitempty" xml:"CriteriaType,omitempty"`
	// The number of the page to return. Pages start from page **1**. Default value: **1**.
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
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The severity of the image baseline that you want to query. Separate multiple severities with commas (,). Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
}

func (s DescribeImageBaselineCheckSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageBaselineCheckSummaryRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetClusterId(v string) *DescribeImageBaselineCheckSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetCriteria(v string) *DescribeImageBaselineCheckSummaryRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetCriteriaType(v string) *DescribeImageBaselineCheckSummaryRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetCurrentPage(v int32) *DescribeImageBaselineCheckSummaryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetLang(v string) *DescribeImageBaselineCheckSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetPageSize(v int32) *DescribeImageBaselineCheckSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetRiskLevel(v string) *DescribeImageBaselineCheckSummaryRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) SetScanRange(v []*string) *DescribeImageBaselineCheckSummaryRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageBaselineCheckSummaryRequest) Validate() error {
	return dara.Validate(s)
}
