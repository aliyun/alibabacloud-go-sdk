// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutopilotTuningHistoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAutopilotTuningHistoriesResponseBodyData) *ListAutopilotTuningHistoriesResponseBody
	GetData() *ListAutopilotTuningHistoriesResponseBodyData
	SetErrorCode(v string) *ListAutopilotTuningHistoriesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAutopilotTuningHistoriesResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListAutopilotTuningHistoriesResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListAutopilotTuningHistoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAutopilotTuningHistoriesResponseBody
	GetSuccess() *bool
}

type ListAutopilotTuningHistoriesResponseBody struct {
	// The tuning history list result.
	Data *ListAutopilotTuningHistoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// When success is false, this value is not empty and indicates the business error code. When success is true, this value is empty.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// When success is false, this value is not empty and indicates the business error message. When success is true, this value is empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is always 200. Use success to determine whether the business request is successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAutopilotTuningHistoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutopilotTuningHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetData() *ListAutopilotTuningHistoriesResponseBodyData {
	return s.Data
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutopilotTuningHistoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetData(v *ListAutopilotTuningHistoriesResponseBodyData) *ListAutopilotTuningHistoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetErrorCode(v string) *ListAutopilotTuningHistoriesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetErrorMessage(v string) *ListAutopilotTuningHistoriesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetHttpCode(v int32) *ListAutopilotTuningHistoriesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetRequestId(v string) *ListAutopilotTuningHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) SetSuccess(v bool) *ListAutopilotTuningHistoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutopilotTuningHistoriesResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 42
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The list of tuning history records.
	TuningHistories []*TuningHistory `json:"tuningHistories,omitempty" xml:"tuningHistories,omitempty" type:"Repeated"`
}

func (s ListAutopilotTuningHistoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAutopilotTuningHistoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) GetTuningHistories() []*TuningHistory {
	return s.TuningHistories
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) SetPageNumber(v int32) *ListAutopilotTuningHistoriesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) SetPageSize(v int32) *ListAutopilotTuningHistoriesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) SetTotalCount(v int32) *ListAutopilotTuningHistoriesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) SetTuningHistories(v []*TuningHistory) *ListAutopilotTuningHistoriesResponseBodyData {
	s.TuningHistories = v
	return s
}

func (s *ListAutopilotTuningHistoriesResponseBodyData) Validate() error {
	if s.TuningHistories != nil {
		for _, item := range s.TuningHistories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
