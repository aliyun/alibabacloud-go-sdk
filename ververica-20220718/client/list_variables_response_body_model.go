// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVariablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Variable) *ListVariablesResponseBody
	GetData() []*Variable
	SetErrorCode(v string) *ListVariablesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListVariablesResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListVariablesResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListVariablesResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListVariablesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVariablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVariablesResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListVariablesResponseBody
	GetTotalSize() *int32
}

type ListVariablesResponseBody struct {
	// 	- If the value of success was true, a list of variables was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Variable `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// CBC799F0-ABCF-1D30-8A4F-882ED4DD****
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

func (s ListVariablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVariablesResponseBody) GetData() []*Variable {
	return s.Data
}

func (s *ListVariablesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListVariablesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListVariablesResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListVariablesResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListVariablesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVariablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVariablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVariablesResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListVariablesResponseBody) SetData(v []*Variable) *ListVariablesResponseBody {
	s.Data = v
	return s
}

func (s *ListVariablesResponseBody) SetErrorCode(v string) *ListVariablesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListVariablesResponseBody) SetErrorMessage(v string) *ListVariablesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListVariablesResponseBody) SetHttpCode(v int32) *ListVariablesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListVariablesResponseBody) SetPageIndex(v int32) *ListVariablesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListVariablesResponseBody) SetPageSize(v int32) *ListVariablesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVariablesResponseBody) SetRequestId(v string) *ListVariablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVariablesResponseBody) SetSuccess(v bool) *ListVariablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListVariablesResponseBody) SetTotalSize(v int32) *ListVariablesResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListVariablesResponseBody) Validate() error {
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
