// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanExecutedHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ScheduledPlanExecutedInfo) *ListScheduledPlanExecutedHistoryResponseBody
	GetData() []*ScheduledPlanExecutedInfo
	SetErrorCode(v string) *ListScheduledPlanExecutedHistoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListScheduledPlanExecutedHistoryResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListScheduledPlanExecutedHistoryResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListScheduledPlanExecutedHistoryResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScheduledPlanExecutedHistoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduledPlanExecutedHistoryResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody
	GetTotalSize() *int32
}

type ListScheduledPlanExecutedHistoryResponseBody struct {
	Data []*ScheduledPlanExecutedInfo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success   *bool  `json:"success,omitempty" xml:"success,omitempty"`
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListScheduledPlanExecutedHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanExecutedHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetData() []*ScheduledPlanExecutedInfo {
	return s.Data
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetData(v []*ScheduledPlanExecutedInfo) *ListScheduledPlanExecutedHistoryResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetErrorCode(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetErrorMessage(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetHttpCode(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetPageIndex(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetPageSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetRequestId(v string) *ListScheduledPlanExecutedHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetSuccess(v bool) *ListScheduledPlanExecutedHistoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) SetTotalSize(v int32) *ListScheduledPlanExecutedHistoryResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListScheduledPlanExecutedHistoryResponseBody) Validate() error {
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
