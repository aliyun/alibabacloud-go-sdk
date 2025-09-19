// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *DescribeImageBaselineCheckResultRequest
	GetCriteria() *string
	SetCriteriaType(v string) *DescribeImageBaselineCheckResultRequest
	GetCriteriaType() *string
	SetCurrentPage(v int32) *DescribeImageBaselineCheckResultRequest
	GetCurrentPage() *int32
	SetImageUuid(v string) *DescribeImageBaselineCheckResultRequest
	GetImageUuid() *string
	SetLang(v string) *DescribeImageBaselineCheckResultRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeImageBaselineCheckResultRequest
	GetPageSize() *int32
	SetRiskLevel(v string) *DescribeImageBaselineCheckResultRequest
	GetRiskLevel() *string
	SetScanRange(v []*string) *DescribeImageBaselineCheckResultRequest
	GetScanRange() []*string
}

type DescribeImageBaselineCheckResultRequest struct {
	// The search condition for the image baseline.
	//
	// example:
	//
	// ak_leak
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
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// a910053dd4710173ecc9e9d8931f****
	ImageUuid *string `json:"ImageUuid,omitempty" xml:"ImageUuid,omitempty"`
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
	// The severity of the image baseline that you want to query. Separate multiple severities with commas (,). By default, all valid values are used. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high,medium,low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The types of the assets that you want to scan.
	ScanRange []*string `json:"ScanRange,omitempty" xml:"ScanRange,omitempty" type:"Repeated"`
}

func (s DescribeImageBaselineCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckResultRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *DescribeImageBaselineCheckResultRequest) GetCriteriaType() *string {
	return s.CriteriaType
}

func (s *DescribeImageBaselineCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineCheckResultRequest) GetImageUuid() *string {
	return s.ImageUuid
}

func (s *DescribeImageBaselineCheckResultRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeImageBaselineCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineCheckResultRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeImageBaselineCheckResultRequest) GetScanRange() []*string {
	return s.ScanRange
}

func (s *DescribeImageBaselineCheckResultRequest) SetCriteria(v string) *DescribeImageBaselineCheckResultRequest {
	s.Criteria = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetCriteriaType(v string) *DescribeImageBaselineCheckResultRequest {
	s.CriteriaType = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetCurrentPage(v int32) *DescribeImageBaselineCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetImageUuid(v string) *DescribeImageBaselineCheckResultRequest {
	s.ImageUuid = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetLang(v string) *DescribeImageBaselineCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetPageSize(v int32) *DescribeImageBaselineCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetRiskLevel(v string) *DescribeImageBaselineCheckResultRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) SetScanRange(v []*string) *DescribeImageBaselineCheckResultRequest {
	s.ScanRange = v
	return s
}

func (s *DescribeImageBaselineCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
