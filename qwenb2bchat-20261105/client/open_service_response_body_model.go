// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v *OpenServiceResponseBodyError) *OpenServiceResponseBody
	GetError() *OpenServiceResponseBodyError
	SetOrderId(v string) *OpenServiceResponseBody
	GetOrderId() *string
	SetSpInstanceId(v string) *OpenServiceResponseBody
	GetSpInstanceId() *string
	SetStatus(v string) *OpenServiceResponseBody
	GetStatus() *string
}

type OpenServiceResponseBody struct {
	// The details of a non-2xx error response, which are mutually exclusive with the status, orderId, and spInstanceId of a normal response.
	Error *OpenServiceResponseBodyError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The order ID associated with the service activation process. Returned when an order exists.
	//
	// example:
	//
	// order-example-001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The service instance ID returned after successful activation and settlement.
	//
	// example:
	//
	// sp-example-001
	SpInstanceId *string `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	// The test or activation result. TEST_OK indicates that the call validation succeeded.
	//
	// example:
	//
	// TEST_OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OpenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBody) GetError() *OpenServiceResponseBodyError {
	return s.Error
}

func (s *OpenServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenServiceResponseBody) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *OpenServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *OpenServiceResponseBody) SetError(v *OpenServiceResponseBodyError) *OpenServiceResponseBody {
	s.Error = v
	return s
}

func (s *OpenServiceResponseBody) SetOrderId(v string) *OpenServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenServiceResponseBody) SetSpInstanceId(v string) *OpenServiceResponseBody {
	s.SpInstanceId = &v
	return s
}

func (s *OpenServiceResponseBody) SetStatus(v string) *OpenServiceResponseBody {
	s.Status = &v
	return s
}

func (s *OpenServiceResponseBody) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenServiceResponseBodyError struct {
	// The error code recognizable by programs.
	//
	// This parameter is required.
	//
	// example:
	//
	// INVALID_REQUEST
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// This parameter is required.
	//
	// example:
	//
	// invalid request
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The correlation ID of the current call.
	//
	// This parameter is required.
	//
	// example:
	//
	// open-service-001
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s OpenServiceResponseBodyError) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponseBodyError) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBodyError) GetCode() *string {
	return s.Code
}

func (s *OpenServiceResponseBodyError) GetMessage() *string {
	return s.Message
}

func (s *OpenServiceResponseBodyError) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenServiceResponseBodyError) SetCode(v string) *OpenServiceResponseBodyError {
	s.Code = &v
	return s
}

func (s *OpenServiceResponseBodyError) SetMessage(v string) *OpenServiceResponseBodyError {
	s.Message = &v
	return s
}

func (s *OpenServiceResponseBodyError) SetRequestId(v string) *OpenServiceResponseBodyError {
	s.RequestId = &v
	return s
}

func (s *OpenServiceResponseBodyError) Validate() error {
	return dara.Validate(s)
}
