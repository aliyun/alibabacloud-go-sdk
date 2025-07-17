// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTableFromCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RemoveTableFromCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RemoveTableFromCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RemoveTableFromCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveTableFromCategoryResponseBody
	GetSuccess() *bool
}

type RemoveTableFromCategoryResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 19DA51A9-AC3E-5C36-8351-07EBCD2B89A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveTableFromCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTableFromCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTableFromCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RemoveTableFromCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RemoveTableFromCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTableFromCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveTableFromCategoryResponseBody) SetErrorCode(v string) *RemoveTableFromCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RemoveTableFromCategoryResponseBody) SetErrorMessage(v string) *RemoveTableFromCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RemoveTableFromCategoryResponseBody) SetRequestId(v string) *RemoveTableFromCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTableFromCategoryResponseBody) SetSuccess(v bool) *RemoveTableFromCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveTableFromCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
