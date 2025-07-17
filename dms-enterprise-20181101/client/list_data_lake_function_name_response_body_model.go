// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakeFunctionNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeFunctionNameResponseBody
	GetErrorMessage() *string
	SetFunctionNameList(v []*string) *ListDataLakeFunctionNameResponseBody
	GetFunctionNameList() []*string
	SetMaxResults(v int32) *ListDataLakeFunctionNameResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionNameResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataLakeFunctionNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeFunctionNameResponseBody
	GetSuccess() *bool
}

type ListDataLakeFunctionNameResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage     *string   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FunctionNameList []*string `json:"FunctionNameList,omitempty" xml:"FunctionNameList,omitempty" type:"Repeated"`
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
	// F1C78D32-1AFD-58AD-9DD2-C8A0896969DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakeFunctionNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeFunctionNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeFunctionNameResponseBody) GetFunctionNameList() []*string {
	return s.FunctionNameList
}

func (s *ListDataLakeFunctionNameResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionNameResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeFunctionNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeFunctionNameResponseBody) SetErrorCode(v string) *ListDataLakeFunctionNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetErrorMessage(v string) *ListDataLakeFunctionNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetFunctionNameList(v []*string) *ListDataLakeFunctionNameResponseBody {
	s.FunctionNameList = v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetMaxResults(v int32) *ListDataLakeFunctionNameResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetNextToken(v string) *ListDataLakeFunctionNameResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetRequestId(v string) *ListDataLakeFunctionNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) SetSuccess(v bool) *ListDataLakeFunctionNameResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeFunctionNameResponseBody) Validate() error {
	return dara.Validate(s)
}
