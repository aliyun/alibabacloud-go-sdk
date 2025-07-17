// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLReviewOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateSQLReviewOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateSQLReviewOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateSQLReviewOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateSQLReviewOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSQLReviewOrderResponseBody
	GetSuccess() *bool
}

type CreateSQLReviewOrderResponseBody struct {
	// The result of the ticket creation task.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSQLReviewOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLReviewOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateSQLReviewOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateSQLReviewOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateSQLReviewOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSQLReviewOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSQLReviewOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateSQLReviewOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetErrorCode(v string) *CreateSQLReviewOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetErrorMessage(v string) *CreateSQLReviewOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetRequestId(v string) *CreateSQLReviewOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) SetSuccess(v bool) *CreateSQLReviewOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSQLReviewOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
