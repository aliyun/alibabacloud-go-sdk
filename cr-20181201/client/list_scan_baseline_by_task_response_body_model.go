// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanBaselineByTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListScanBaselineByTaskResponseBody
	GetCode() *int32
	SetIsSuccess(v bool) *ListScanBaselineByTaskResponseBody
	GetIsSuccess() *bool
	SetPageNo(v int32) *ListScanBaselineByTaskResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanBaselineByTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScanBaselineByTaskResponseBody
	GetRequestId() *string
	SetScanBaselines(v []*ListScanBaselineByTaskResponseBodyScanBaselines) *ListScanBaselineByTaskResponseBody
	GetScanBaselines() []*ListScanBaselineByTaskResponseBodyScanBaselines
	SetTotalCount(v int32) *ListScanBaselineByTaskResponseBody
	GetTotalCount() *int32
}

type ListScanBaselineByTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request was successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5259118F-79E2-57E9-9AEA-551586F4FAED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scanned baseline risks.
	ScanBaselines []*ListScanBaselineByTaskResponseBodyScanBaselines `json:"ScanBaselines,omitempty" xml:"ScanBaselines,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanBaselineByTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScanBaselineByTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListScanBaselineByTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListScanBaselineByTaskResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanBaselineByTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanBaselineByTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScanBaselineByTaskResponseBody) GetScanBaselines() []*ListScanBaselineByTaskResponseBodyScanBaselines {
	return s.ScanBaselines
}

func (s *ListScanBaselineByTaskResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScanBaselineByTaskResponseBody) SetCode(v int32) *ListScanBaselineByTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetIsSuccess(v bool) *ListScanBaselineByTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetPageNo(v int32) *ListScanBaselineByTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetPageSize(v int32) *ListScanBaselineByTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetRequestId(v string) *ListScanBaselineByTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetScanBaselines(v []*ListScanBaselineByTaskResponseBodyScanBaselines) *ListScanBaselineByTaskResponseBody {
	s.ScanBaselines = v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetTotalCount(v int32) *ListScanBaselineByTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) Validate() error {
	if s.ScanBaselines != nil {
		for _, item := range s.ScanBaselines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScanBaselineByTaskResponseBodyScanBaselines struct {
	// The category to which the baseline risk belongs.
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// Suggestions about how to fix the baseline risk.
	BaselineDetailAdvice *string `json:"BaselineDetailAdvice,omitempty" xml:"BaselineDetailAdvice,omitempty"`
	// The description of the baseline risk.
	BaselineDetailDescription *string `json:"BaselineDetailDescription,omitempty" xml:"BaselineDetailDescription,omitempty"`
	// The path and content of the baseline risk.
	//
	// example:
	//
	// usr/local/www/project/environments/dev/common/config/paramsxxx
	BaselineDetailPrompt *string `json:"BaselineDetailPrompt,omitempty" xml:"BaselineDetailPrompt,omitempty"`
	// The number of baseline risks.
	//
	// example:
	//
	// 1
	BaselineItemCount *int32 `json:"BaselineItemCount,omitempty" xml:"BaselineItemCount,omitempty"`
	// The name of the baseline risk.
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The key of the baseline name.
	//
	// example:
	//
	// ak_leak
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The severity of the baseline risk.
	//
	// example:
	//
	// high
	BaselineNameLevel *string `json:"BaselineNameLevel,omitempty" xml:"BaselineNameLevel,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1695090008000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time of the first scan.
	//
	// example:
	//
	// 2024-04-10 15:33:26
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// High risk quantity.
	//
	// example:
	//
	// 1
	HighRiskItemCount *int32 `json:"HighRiskItemCount,omitempty" xml:"HighRiskItemCount,omitempty"`
	// Low risk quantity.
	//
	// example:
	//
	// 1
	LowRiskItemCount *int32 `json:"LowRiskItemCount,omitempty" xml:"LowRiskItemCount,omitempty"`
	// Medium risk quantity.
	//
	// example:
	//
	// 1
	MiddleRiskItemCount *int32 `json:"MiddleRiskItemCount,omitempty" xml:"MiddleRiskItemCount,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 2328fcaa-f28a-405d-a357-asdvfrew23
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1684220824226
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScanBaselineByTaskResponseBodyScanBaselines) String() string {
	return dara.Prettify(s)
}

func (s ListScanBaselineByTaskResponseBodyScanBaselines) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineClassAlias() *string {
	return s.BaselineClassAlias
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineDetailAdvice() *string {
	return s.BaselineDetailAdvice
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineDetailDescription() *string {
	return s.BaselineDetailDescription
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineDetailPrompt() *string {
	return s.BaselineDetailPrompt
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineItemCount() *int32 {
	return s.BaselineItemCount
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineNameAlias() *string {
	return s.BaselineNameAlias
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineNameKey() *string {
	return s.BaselineNameKey
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetBaselineNameLevel() *string {
	return s.BaselineNameLevel
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetFirstScanTime() *int64 {
	return s.FirstScanTime
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetHighRiskItemCount() *int32 {
	return s.HighRiskItemCount
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetLowRiskItemCount() *int32 {
	return s.LowRiskItemCount
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetMiddleRiskItemCount() *int32 {
	return s.MiddleRiskItemCount
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetScanTaskId() *string {
	return s.ScanTaskId
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineClassAlias(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineClassAlias = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailAdvice(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailAdvice = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailDescription(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailDescription = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailPrompt(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailPrompt = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameAlias(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameAlias = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameKey(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameKey = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameLevel(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameLevel = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetCreateTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.CreateTime = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetFirstScanTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.FirstScanTime = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetHighRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.HighRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetLowRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.LowRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetMiddleRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.MiddleRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetScanTaskId(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetUpdateTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.UpdateTime = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) Validate() error {
	return dara.Validate(s)
}
