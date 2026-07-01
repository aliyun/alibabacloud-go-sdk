// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Event) *GetEventsResponseBody
	GetData() []*Event
	SetErrorCode(v string) *GetEventsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetEventsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetEventsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *GetEventsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *GetEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEventsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *GetEventsResponseBody
	GetTotalSize() *int32
}

type GetEventsResponseBody struct {
	// A list of runtime events matching the filter criteria. Returned only when the request is successful (`success` is `true`).
	Data []*Event `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The business error code. Returned only when the request fails (that is, `success` is `false`).
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The business error message. Returned only when the request fails (that is, `success` is `false`).
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code. This field always returns `200`. To confirm the request\\"s success, check the `success` parameter.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries on this page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 4
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventsResponseBody) GetData() []*Event {
	return s.Data
}

func (s *GetEventsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetEventsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetEventsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetEventsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetEventsResponseBody) SetData(v []*Event) *GetEventsResponseBody {
	s.Data = v
	return s
}

func (s *GetEventsResponseBody) SetErrorCode(v string) *GetEventsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetEventsResponseBody) SetErrorMessage(v string) *GetEventsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetEventsResponseBody) SetHttpCode(v int32) *GetEventsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetEventsResponseBody) SetPageIndex(v int32) *GetEventsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetEventsResponseBody) SetPageSize(v int32) *GetEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetEventsResponseBody) SetRequestId(v string) *GetEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventsResponseBody) SetSuccess(v bool) *GetEventsResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventsResponseBody) SetTotalSize(v int32) *GetEventsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *GetEventsResponseBody) Validate() error {
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
