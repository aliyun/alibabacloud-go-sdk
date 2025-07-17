// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcCorrectOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v []*int64) *CreateProcCorrectOrderResponseBody
	GetCreateOrderResult() []*int64
	SetErrorCode(v string) *CreateProcCorrectOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateProcCorrectOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateProcCorrectOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProcCorrectOrderResponseBody
	GetSuccess() *bool
}

type CreateProcCorrectOrderResponseBody struct {
	CreateOrderResult []*int64 `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Repeated"`
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
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6E23D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProcCorrectOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProcCorrectOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProcCorrectOrderResponseBody) GetCreateOrderResult() []*int64 {
	return s.CreateOrderResult
}

func (s *CreateProcCorrectOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProcCorrectOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateProcCorrectOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProcCorrectOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProcCorrectOrderResponseBody) SetCreateOrderResult(v []*int64) *CreateProcCorrectOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateProcCorrectOrderResponseBody) SetErrorCode(v string) *CreateProcCorrectOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProcCorrectOrderResponseBody) SetErrorMessage(v string) *CreateProcCorrectOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProcCorrectOrderResponseBody) SetRequestId(v string) *CreateProcCorrectOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProcCorrectOrderResponseBody) SetSuccess(v bool) *CreateProcCorrectOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProcCorrectOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
