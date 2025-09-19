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
	// An array that consists of the check results of image baselines.
	BaselineResultSummary []*DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary `json:"BaselineResultSummary,omitempty" xml:"BaselineResultSummary,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeImageBaselineCheckSummaryResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	return dara.Validate(s)
}

type DescribeImageBaselineCheckSummaryResponseBodyBaselineResultSummary struct {
	// The category of the baseline.
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
	// The name of the baseline.
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
	// 1626628760000
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The number of images on which **high*	- baseline risks are detected.
	//
	// example:
	//
	// 15
	HighRiskImage *int32 `json:"HighRiskImage,omitempty" xml:"HighRiskImage,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1626628760000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of images on which **low*	- baseline risks are detected.
	//
	// example:
	//
	// 0
	LowRiskImage *int32 `json:"LowRiskImage,omitempty" xml:"LowRiskImage,omitempty"`
	// The number of images on which **medium*	- baseline risks are detected.
	//
	// example:
	//
	// 0
	MiddleRiskImage *int32 `json:"MiddleRiskImage,omitempty" xml:"MiddleRiskImage,omitempty"`
	// The status of the baseline risks. Valid values:
	//
	// 	- **0**: unfixed
	//
	// 	- **1**: fixed
	//
	// 	- **2**: pending verification
	//
	// 	- **3**: fixing failed
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
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
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
