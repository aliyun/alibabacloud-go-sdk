// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateOrderResult(v *CreateOrderResponseBodyCreateOrderResult) *CreateOrderResponseBody
	GetCreateOrderResult() *CreateOrderResponseBodyCreateOrderResult
	SetErrorCode(v string) *CreateOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrderResponseBody
	GetSuccess() *bool
}

type CreateOrderResponseBody struct {
	// The ID of the ticket.
	CreateOrderResult *CreateOrderResponseBodyCreateOrderResult `json:"CreateOrderResult,omitempty" xml:"CreateOrderResult,omitempty" type:"Struct"`
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
	// 427688B8-ADFB-4C4E-9D45-EF5C1FD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetCreateOrderResult() *CreateOrderResponseBodyCreateOrderResult {
	return s.CreateOrderResult
}

func (s *CreateOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrderResponseBody) SetCreateOrderResult(v *CreateOrderResponseBodyCreateOrderResult) *CreateOrderResponseBody {
	s.CreateOrderResult = v
	return s
}

func (s *CreateOrderResponseBody) SetErrorCode(v string) *CreateOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrderResponseBody) SetErrorMessage(v string) *CreateOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	if s.CreateOrderResult != nil {
		if err := s.CreateOrderResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrderResponseBodyCreateOrderResult struct {
	OrderIds []*int64 `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s CreateOrderResponseBodyCreateOrderResult) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBodyCreateOrderResult) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBodyCreateOrderResult) GetOrderIds() []*int64 {
	return s.OrderIds
}

func (s *CreateOrderResponseBodyCreateOrderResult) SetOrderIds(v []*int64) *CreateOrderResponseBodyCreateOrderResult {
	s.OrderIds = v
	return s
}

func (s *CreateOrderResponseBodyCreateOrderResult) Validate() error {
	return dara.Validate(s)
}
