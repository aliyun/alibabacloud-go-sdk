// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakeFunctionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeFunctionResponseBody
	GetErrorMessage() *string
	SetFunctionList(v []*DLFunction) *ListDataLakeFunctionResponseBody
	GetFunctionList() []*DLFunction
	SetMaxResults(v int32) *ListDataLakeFunctionResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataLakeFunctionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeFunctionResponseBody
	GetSuccess() *bool
}

type ListDataLakeFunctionResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FunctionList []*DLFunction `json:"FunctionList,omitempty" xml:"FunctionList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakeFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeFunctionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeFunctionResponseBody) GetFunctionList() []*DLFunction {
	return s.FunctionList
}

func (s *ListDataLakeFunctionResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeFunctionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeFunctionResponseBody) SetErrorCode(v string) *ListDataLakeFunctionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetErrorMessage(v string) *ListDataLakeFunctionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetFunctionList(v []*DLFunction) *ListDataLakeFunctionResponseBody {
	s.FunctionList = v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetMaxResults(v int32) *ListDataLakeFunctionResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetNextToken(v string) *ListDataLakeFunctionResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetRequestId(v string) *ListDataLakeFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) SetSuccess(v bool) *ListDataLakeFunctionResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
