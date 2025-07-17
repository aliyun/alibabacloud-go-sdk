// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructSyncOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateStructSyncOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateStructSyncOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateStructSyncOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateStructSyncOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStructSyncOrderResponseBody
	GetSuccess() *bool
}

type CreateStructSyncOrderResponseBody struct {
	// The result of creating the ticket.
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
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStructSyncOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStructSyncOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStructSyncOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateStructSyncOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateStructSyncOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateStructSyncOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStructSyncOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStructSyncOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateStructSyncOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetErrorCode(v string) *CreateStructSyncOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetErrorMessage(v string) *CreateStructSyncOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetRequestId(v string) *CreateStructSyncOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) SetSuccess(v bool) *CreateStructSyncOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStructSyncOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
