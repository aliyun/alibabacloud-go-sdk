// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineCheckResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineResult(v []*DescribeImageBaselineCheckResultResponseBodyBaselineResult) *DescribeImageBaselineCheckResultResponseBody
	GetBaselineResult() []*DescribeImageBaselineCheckResultResponseBodyBaselineResult
	SetPageInfo(v *DescribeImageBaselineCheckResultResponseBodyPageInfo) *DescribeImageBaselineCheckResultResponseBody
	GetPageInfo() *DescribeImageBaselineCheckResultResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageBaselineCheckResultResponseBody
	GetRequestId() *string
}

type DescribeImageBaselineCheckResultResponseBody struct {
	// An array that consists of the check results of image baselines.
	BaselineResult []*DescribeImageBaselineCheckResultResponseBodyBaselineResult `json:"BaselineResult,omitempty" xml:"BaselineResult,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageBaselineCheckResultResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageBaselineCheckResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckResultResponseBody) GetBaselineResult() []*DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	return s.BaselineResult
}

func (s *DescribeImageBaselineCheckResultResponseBody) GetPageInfo() *DescribeImageBaselineCheckResultResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageBaselineCheckResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageBaselineCheckResultResponseBody) SetBaselineResult(v []*DescribeImageBaselineCheckResultResponseBodyBaselineResult) *DescribeImageBaselineCheckResultResponseBody {
	s.BaselineResult = v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBody) SetPageInfo(v *DescribeImageBaselineCheckResultResponseBodyPageInfo) *DescribeImageBaselineCheckResultResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBody) SetRequestId(v string) *DescribeImageBaselineCheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBody) Validate() error {
	if s.BaselineResult != nil {
		for _, item := range s.BaselineResult {
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

type DescribeImageBaselineCheckResultResponseBodyBaselineResult struct {
	// The key of the image baseline type.
	//
	// example:
	//
	// identification
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// The number of baseline check items.
	//
	// example:
	//
	// 5
	BaselineItemCount *int32 `json:"BaselineItemCount,omitempty" xml:"BaselineItemCount,omitempty"`
	// The alias of the image baseline.
	//
	// example:
	//
	// Identity authentication
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The key of the image baseline.
	//
	// example:
	//
	// identification
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The severity of the image baseline. Valid values:
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
	BaselineNameLevel *string `json:"BaselineNameLevel,omitempty" xml:"BaselineNameLevel,omitempty"`
	// The timestamp generated when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1610304058366
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The number of high-risk images that are affected.
	//
	// example:
	//
	// 1
	HighRiskItemCount *int32 `json:"HighRiskItemCount,omitempty" xml:"HighRiskItemCount,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1610304058301
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of low-risk images that are affected.
	//
	// example:
	//
	// 10
	LowRiskItemCount *int32 `json:"LowRiskItemCount,omitempty" xml:"LowRiskItemCount,omitempty"`
	// The number of medium-risk images that are affected.
	//
	// example:
	//
	// 1
	MiddleRiskItemCount *int32 `json:"MiddleRiskItemCount,omitempty" xml:"MiddleRiskItemCount,omitempty"`
	// The status of the baseline risks. Valid values:
	//
	// 	- **0**: unfixed
	//
	// 	- **1**: fixed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeImageBaselineCheckResultResponseBodyBaselineResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckResultResponseBodyBaselineResult) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetBaselineClassAlias() *string {
	return s.BaselineClassAlias
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetBaselineItemCount() *int32 {
	return s.BaselineItemCount
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetBaselineNameAlias() *string {
	return s.BaselineNameAlias
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetBaselineNameLevel() *string {
	return s.BaselineNameLevel
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetHighRiskItemCount() *int32 {
	return s.HighRiskItemCount
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetLowRiskItemCount() *int32 {
	return s.LowRiskItemCount
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetMiddleRiskItemCount() *int32 {
	return s.MiddleRiskItemCount
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetBaselineClassAlias(v string) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.BaselineClassAlias = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetBaselineItemCount(v int32) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.BaselineItemCount = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetBaselineNameAlias(v string) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.BaselineNameAlias = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetBaselineNameKey(v string) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.BaselineNameKey = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetBaselineNameLevel(v string) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.BaselineNameLevel = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetFirstScanTime(v int64) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.FirstScanTime = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetHighRiskItemCount(v int32) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.HighRiskItemCount = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetLastScanTime(v int64) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetLowRiskItemCount(v int32) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.LowRiskItemCount = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetMiddleRiskItemCount(v int32) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.MiddleRiskItemCount = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) SetStatus(v int32) *DescribeImageBaselineCheckResultResponseBodyBaselineResult {
	s.Status = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyBaselineResult) Validate() error {
	return dara.Validate(s)
}

type DescribeImageBaselineCheckResultResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageBaselineCheckResultResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineCheckResultResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) SetCount(v int32) *DescribeImageBaselineCheckResultResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageBaselineCheckResultResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageBaselineCheckResultResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageBaselineCheckResultResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageBaselineCheckResultResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
