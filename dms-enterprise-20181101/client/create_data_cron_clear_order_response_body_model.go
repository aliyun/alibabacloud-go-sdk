// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCronClearOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateDataCronClearOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateDataCronClearOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataCronClearOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataCronClearOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataCronClearOrderResponseBody
	GetSuccess() *bool
}

type CreateDataCronClearOrderResponseBody struct {
	// The ID of the ticket.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataCronClearOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCronClearOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCronClearOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataCronClearOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataCronClearOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataCronClearOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataCronClearOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataCronClearOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataCronClearOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetErrorCode(v string) *CreateDataCronClearOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetErrorMessage(v string) *CreateDataCronClearOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetRequestId(v string) *CreateDataCronClearOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) SetSuccess(v bool) *CreateDataCronClearOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataCronClearOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
