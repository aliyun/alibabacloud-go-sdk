// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCheckColumnResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDataCheckColumnResultsResponseBodyData) *ListDataCheckColumnResultsResponseBody
	GetData() []*ListDataCheckColumnResultsResponseBodyData
	SetErrCode(v string) *ListDataCheckColumnResultsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDataCheckColumnResultsResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListDataCheckColumnResultsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDataCheckColumnResultsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDataCheckColumnResultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCheckColumnResultsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListDataCheckColumnResultsResponseBody
	GetTotalCount() *int32
}

type ListDataCheckColumnResultsResponseBody struct {
	// The response data.
	Data []*ListDataCheckColumnResultsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// None
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// not supported.pos 3222, line 112, column 14, token IDENTIFIER settings
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The page number that indicates the requested page.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 016D6CE5-51C6-5767-A8F9-D2818FC56509
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCheckColumnResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckColumnResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCheckColumnResultsResponseBody) GetData() []*ListDataCheckColumnResultsResponseBodyData {
	return s.Data
}

func (s *ListDataCheckColumnResultsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDataCheckColumnResultsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDataCheckColumnResultsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDataCheckColumnResultsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataCheckColumnResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCheckColumnResultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCheckColumnResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCheckColumnResultsResponseBody) SetData(v []*ListDataCheckColumnResultsResponseBodyData) *ListDataCheckColumnResultsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetErrCode(v string) *ListDataCheckColumnResultsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetErrMessage(v string) *ListDataCheckColumnResultsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetPageIndex(v int32) *ListDataCheckColumnResultsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetPageSize(v int32) *ListDataCheckColumnResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetRequestId(v string) *ListDataCheckColumnResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetSuccess(v bool) *ListDataCheckColumnResultsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) SetTotalCount(v int32) *ListDataCheckColumnResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCheckColumnResultsResponseBodyData struct {
	// The actual difference.
	//
	// example:
	//
	// 0.00
	ActualThreshold *string `json:"actualThreshold,omitempty" xml:"actualThreshold,omitempty"`
	// The execution result. Valid values:
	//
	// - 0: no record.
	//
	// - 1: passed.
	//
	// - 2: failed.
	//
	// example:
	//
	// 0
	CheckResult *int32 `json:"checkResult,omitempty" xml:"checkResult,omitempty"`
	// The comparison rule.
	//
	// example:
	//
	// SUM
	CheckRule *string `json:"checkRule,omitempty" xml:"checkRule,omitempty"`
	// The alias of the destination.
	//
	// example:
	//
	// dst
	DstAlias *string `json:"dstAlias,omitempty" xml:"dstAlias,omitempty"`
	// The field name of the destination.
	//
	// example:
	//
	// amount
	DstColumnName *string `json:"dstColumnName,omitempty" xml:"dstColumnName,omitempty"`
	// The field type of the destination.
	//
	// example:
	//
	// decimal(38,18)
	DstColumnType *string `json:"dstColumnType,omitempty" xml:"dstColumnType,omitempty"`
	// The metric key of the destination.
	//
	// example:
	//
	// dst_sum_amount
	DstMetricColumn *string `json:"dstMetricColumn,omitempty" xml:"dstMetricColumn,omitempty"`
	// The result value of the destination field.
	//
	// example:
	//
	// 1000.00
	DstResult *string `json:"dstResult,omitempty" xml:"dstResult,omitempty"`
	// The expected threshold.
	//
	// example:
	//
	// 0.00
	ExpectThreshold *string `json:"expectThreshold,omitempty" xml:"expectThreshold,omitempty"`
	// The validation result. Valid values:
	//
	// - 0: inconsistent.
	//
	// - 1: consistent.
	//
	// - 2: manually repaired.
	//
	// example:
	//
	// 0
	IsConsistent *int32 `json:"isConsistent,omitempty" xml:"isConsistent,omitempty"`
	// The alias of the source.
	//
	// example:
	//
	// src
	SrcAlias *string `json:"srcAlias,omitempty" xml:"srcAlias,omitempty"`
	// The field name of the source.
	//
	// example:
	//
	// amount
	SrcColumnName *string `json:"srcColumnName,omitempty" xml:"srcColumnName,omitempty"`
	// The field type of the source.
	//
	// example:
	//
	// decimal(38,18)
	SrcColumnType *string `json:"srcColumnType,omitempty" xml:"srcColumnType,omitempty"`
	// The metric key of the source.
	//
	// example:
	//
	// src_sum_amount
	SrcMetricColumn *string `json:"srcMetricColumn,omitempty" xml:"srcMetricColumn,omitempty"`
	// The result value of the source field.
	//
	// example:
	//
	// 1000.00
	SrcResult *string `json:"srcResult,omitempty" xml:"srcResult,omitempty"`
	// The step ID.
	//
	// example:
	//
	// 1
	StepId *int64 `json:"stepId,omitempty" xml:"stepId,omitempty"`
}

func (s ListDataCheckColumnResultsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataCheckColumnResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetActualThreshold() *string {
	return s.ActualThreshold
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetCheckResult() *int32 {
	return s.CheckResult
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetCheckRule() *string {
	return s.CheckRule
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetDstAlias() *string {
	return s.DstAlias
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetDstColumnName() *string {
	return s.DstColumnName
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetDstColumnType() *string {
	return s.DstColumnType
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetDstMetricColumn() *string {
	return s.DstMetricColumn
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetDstResult() *string {
	return s.DstResult
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetExpectThreshold() *string {
	return s.ExpectThreshold
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetIsConsistent() *int32 {
	return s.IsConsistent
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetSrcAlias() *string {
	return s.SrcAlias
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetSrcColumnName() *string {
	return s.SrcColumnName
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetSrcColumnType() *string {
	return s.SrcColumnType
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetSrcMetricColumn() *string {
	return s.SrcMetricColumn
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetSrcResult() *string {
	return s.SrcResult
}

func (s *ListDataCheckColumnResultsResponseBodyData) GetStepId() *int64 {
	return s.StepId
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetActualThreshold(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.ActualThreshold = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetCheckResult(v int32) *ListDataCheckColumnResultsResponseBodyData {
	s.CheckResult = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetCheckRule(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.CheckRule = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetDstAlias(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.DstAlias = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetDstColumnName(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.DstColumnName = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetDstColumnType(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.DstColumnType = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetDstMetricColumn(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.DstMetricColumn = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetDstResult(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.DstResult = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetExpectThreshold(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.ExpectThreshold = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetIsConsistent(v int32) *ListDataCheckColumnResultsResponseBodyData {
	s.IsConsistent = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetSrcAlias(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.SrcAlias = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetSrcColumnName(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.SrcColumnName = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetSrcColumnType(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.SrcColumnType = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetSrcMetricColumn(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.SrcMetricColumn = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetSrcResult(v string) *ListDataCheckColumnResultsResponseBodyData {
	s.SrcResult = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) SetStepId(v int64) *ListDataCheckColumnResultsResponseBodyData {
	s.StepId = &v
	return s
}

func (s *ListDataCheckColumnResultsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
