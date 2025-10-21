// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavepointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Savepoint) *ListSavepointsResponseBody
	GetData() []*Savepoint
	SetErrorCode(v string) *ListSavepointsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSavepointsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListSavepointsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListSavepointsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListSavepointsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSavepointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSavepointsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListSavepointsResponseBody
	GetTotalSize() *int32
}

type ListSavepointsResponseBody struct {
	// 	- If the value of success was true, a list of savepoints was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Savepoint `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSavepointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSavepointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSavepointsResponseBody) GetData() []*Savepoint {
	return s.Data
}

func (s *ListSavepointsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSavepointsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSavepointsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListSavepointsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListSavepointsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSavepointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSavepointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSavepointsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListSavepointsResponseBody) SetData(v []*Savepoint) *ListSavepointsResponseBody {
	s.Data = v
	return s
}

func (s *ListSavepointsResponseBody) SetErrorCode(v string) *ListSavepointsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSavepointsResponseBody) SetErrorMessage(v string) *ListSavepointsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSavepointsResponseBody) SetHttpCode(v int32) *ListSavepointsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListSavepointsResponseBody) SetPageIndex(v int32) *ListSavepointsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListSavepointsResponseBody) SetPageSize(v int32) *ListSavepointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSavepointsResponseBody) SetRequestId(v string) *ListSavepointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSavepointsResponseBody) SetSuccess(v bool) *ListSavepointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSavepointsResponseBody) SetTotalSize(v int32) *ListSavepointsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListSavepointsResponseBody) Validate() error {
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
