// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataCorrectOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateDataCorrectOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateDataCorrectOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataCorrectOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataCorrectOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataCorrectOrderResponseBody
	GetSuccess() *bool
}

type CreateDataCorrectOrderResponseBody struct {
	// The IDs of the tickets.
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

func (s CreateDataCorrectOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataCorrectOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataCorrectOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateDataCorrectOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataCorrectOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataCorrectOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataCorrectOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataCorrectOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateDataCorrectOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetErrorCode(v string) *CreateDataCorrectOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetErrorMessage(v string) *CreateDataCorrectOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetRequestId(v string) *CreateDataCorrectOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) SetSuccess(v bool) *CreateDataCorrectOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataCorrectOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
