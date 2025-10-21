// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ScheduledPlan) *ListScheduledPlanResponseBody
	GetData() []*ScheduledPlan
	SetErrorCode(v string) *ListScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListScheduledPlanResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListScheduledPlanResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListScheduledPlanResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduledPlanResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListScheduledPlanResponseBody
	GetTotalSize() *int32
}

type ListScheduledPlanResponseBody struct {
	Data []*ScheduledPlan `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 4
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanResponseBody) GetData() []*ScheduledPlan {
	return s.Data
}

func (s *ListScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListScheduledPlanResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListScheduledPlanResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduledPlanResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListScheduledPlanResponseBody) SetData(v []*ScheduledPlan) *ListScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduledPlanResponseBody) SetErrorCode(v string) *ListScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetErrorMessage(v string) *ListScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetHttpCode(v int32) *ListScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetPageIndex(v int32) *ListScheduledPlanResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetPageSize(v int32) *ListScheduledPlanResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetRequestId(v string) *ListScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetSuccess(v bool) *ListScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduledPlanResponseBody) SetTotalSize(v int32) *ListScheduledPlanResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListScheduledPlanResponseBody) Validate() error {
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
