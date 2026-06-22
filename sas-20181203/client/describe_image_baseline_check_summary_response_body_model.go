// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineResultSummary(v []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) *DescribeImageBaselineCheckSummaryResponseBody
	GetBaselineResultSummary() []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary
	SetPageInfo(v *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) *DescribeImageBaselineCheckSummaryResponseBody
	GetPageInfo() *DescribeImageBaselineCheckSummaryResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageBaselineCheckSummaryResponseBody
	GetRequestId() *string
}

type DescribeImageBaselineCheckSummaryResponseBody struct {
	// The details of the image baseline check list.
	BaselineResultSummary []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary `json:"BaselineResultSummary,omitempty" xml:"BaselineResultSummary,omitempty" type:"Repeated"`
	// The paging information displayed on the page in a paged query.
	PageInfo *DescribeImageBaselineCheckSummaryResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 5BD95679-D63A-4151-97D0-188432F4A57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageBaselineCheckSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) GetBaselineResultSummary() []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	return s.BaselineResultSummary
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) GetPageInfo() *DescribeImageBaselineCheckSummaryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) SetBaselineResultSummary(v []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) *DescribeImageBaselineCheckSummaryResponseBody {
	s.BaselineResultSummary = v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) SetPageInfo(v *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) *DescribeImageBaselineCheckSummaryResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) SetRequestId(v string) *DescribeImageBaselineCheckSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBody) Validate() error {
	if s.BaselineResultSummary != nil {
		for _, item := range s.BaselineResultSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary struct {
	// The baseline category.
	//
	// example:
	//
	// Unauthorized access
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// The keyword of the baseline category.
	//
	// example:
	//
	// hc_image_exploit
	BaselineClassKey *string `json:"BaselineClassKey,omitempty" xml:"BaselineClassKey,omitempty"`
	// The baseline name.
	//
	// example:
	//
	// Unauthorized access
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The keyword of the baseline name.
	//
	// example:
	//
	// hc_image_exploit
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The risk level of the baseline. Valid values:
	//
	// - **high**: high risk
	//
	// - **medium**: medium risk
	//
	// - **low**: low risk.
	//
	// example:
	//
	// high
	BaselineNameLevel *string `json:"BaselineNameLevel,omitempty" xml:"BaselineNameLevel,omitempty"`
	// The timestamp of the first scan, in milliseconds.
	//
	// example:
	//
	// 1626628760000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The number of images with a **high*	- risk level that have baseline risk issues.
	//
	// example:
	//
	// 15
	HighRiskImage *int32 `json:"HighRiskImage,omitempty" xml:"HighRiskImage,omitempty"`
	// The timestamp of the most recent scan, in milliseconds.
	//
	// example:
	//
	// 1626628760000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of images with a **low*	- risk level that have baseline risk issues.
	//
	// example:
	//
	// 0
	LowRiskImage *int32 `json:"LowRiskImage,omitempty" xml:"LowRiskImage,omitempty"`
	// The number of images with a **medium*	- risk level that have baseline risk issues.
	//
	// example:
	//
	// 0
	MiddleRiskImage *int32 `json:"MiddleRiskImage,omitempty" xml:"MiddleRiskImage,omitempty"`
	// The fix status of the baseline risk. Valid values:
	//
	// - **0**: Unfixed.
	//
	// - **1**: Fixed.
	//
	// - **2**: Pending verification.
	//
	// - **3**: Fix failed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetBaselineClassAlias() *string {
	return s.BaselineClassAlias
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetBaselineClassKey() *string {
	return s.BaselineClassKey
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetBaselineNameAlias() *string {
	return s.BaselineNameAlias
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetBaselineNameLevel() *string {
	return s.BaselineNameLevel
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetHighRiskImage() *int32 {
	return s.HighRiskImage
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetLowRiskImage() *int32 {
	return s.LowRiskImage
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetMiddleRiskImage() *int32 {
	return s.MiddleRiskImage
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetBaselineClassAlias(v string) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.BaselineClassAlias = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetBaselineClassKey(v string) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.BaselineClassKey = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetBaselineNameAlias(v string) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.BaselineNameAlias = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetBaselineNameKey(v string) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.BaselineNameKey = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetBaselineNameLevel(v string) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.BaselineNameLevel = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetFirstScanTime(v int64) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetHighRiskImage(v int32) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.HighRiskImage = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetLastScanTime(v int64) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetLowRiskImage(v int32) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.LowRiskImage = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetMiddleRiskImage(v int32) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.MiddleRiskImage = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) SetStatus(v int32) *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary {
	s.Status = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary) Validate() error {
	return dara.Validate(s)
}

type DescribeImageBaselineCheckSummaryResponseBodyPageInfo struct {
	// The number of entries on the current page in a paged query.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The current page number in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of image baseline check results per page in a paged query. Default value: **20**, which indicates that 20 image baseline check results are displayed per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of query results.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageBaselineCheckSummaryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckSummaryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) SetCount(v int32) *DescribeImageBaselineCheckSummaryResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageBaselineCheckSummaryResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageBaselineCheckSummaryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageBaselineCheckSummaryResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageBaselineCheckSummaryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
