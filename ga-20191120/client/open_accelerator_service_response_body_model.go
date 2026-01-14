// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAcceleratorServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenAcceleratorServiceResponseBody
	GetCode() *string
	SetMessage(v string) *OpenAcceleratorServiceResponseBody
	GetMessage() *string
	SetOrderId(v int64) *OpenAcceleratorServiceResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *OpenAcceleratorServiceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *OpenAcceleratorServiceResponseBody
	GetSuccess() *string
}

type OpenAcceleratorServiceResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 208257****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B49B60F6-F6C8-49E5-B06B-E33E3A469A92
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true:*	- The call was successful.
	//
	// 	- **false:*	- The call failed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenAcceleratorServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenAcceleratorServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenAcceleratorServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenAcceleratorServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenAcceleratorServiceResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *OpenAcceleratorServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenAcceleratorServiceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *OpenAcceleratorServiceResponseBody) SetCode(v string) *OpenAcceleratorServiceResponseBody {
	s.Code = &v
	return s
}

func (s *OpenAcceleratorServiceResponseBody) SetMessage(v string) *OpenAcceleratorServiceResponseBody {
	s.Message = &v
	return s
}

func (s *OpenAcceleratorServiceResponseBody) SetOrderId(v int64) *OpenAcceleratorServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenAcceleratorServiceResponseBody) SetRequestId(v string) *OpenAcceleratorServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenAcceleratorServiceResponseBody) SetSuccess(v string) *OpenAcceleratorServiceResponseBody {
	s.Success = &v
	return s
}

func (s *OpenAcceleratorServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
