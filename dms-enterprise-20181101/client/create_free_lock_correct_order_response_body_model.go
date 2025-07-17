// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFreeLockCorrectOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateFreeLockCorrectOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateFreeLockCorrectOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateFreeLockCorrectOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateFreeLockCorrectOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateFreeLockCorrectOrderResponseBody
	GetSuccess() *bool
}

type CreateFreeLockCorrectOrderResponseBody struct {
	// The ID of the ticket.
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
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
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFreeLockCorrectOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFreeLockCorrectOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFreeLockCorrectOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateFreeLockCorrectOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateFreeLockCorrectOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateFreeLockCorrectOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFreeLockCorrectOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateFreeLockCorrectOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetErrorCode(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetErrorMessage(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetRequestId(v string) *CreateFreeLockCorrectOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) SetSuccess(v bool) *CreateFreeLockCorrectOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFreeLockCorrectOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
