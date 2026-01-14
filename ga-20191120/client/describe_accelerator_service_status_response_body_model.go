// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAcceleratorServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAcceleratorServiceStatusResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeAcceleratorServiceStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAcceleratorServiceStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAcceleratorServiceStatusResponseBody
	GetStatus() *string
	SetSuccess(v string) *DescribeAcceleratorServiceStatusResponseBody
	GetSuccess() *string
}

type DescribeAcceleratorServiceStatusResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
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
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the GA instance. Valid values:
	//
	// 	- Released: The instance was released due to overdue payments.
	//
	// 	- Expired: The instance expired due to overdue payments.
	//
	// 	- NotOpened: The instance is not activated.
	//
	// 	- Normal: The instance is activated.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAcceleratorServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAcceleratorServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAcceleratorServiceStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAcceleratorServiceStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAcceleratorServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAcceleratorServiceStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAcceleratorServiceStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeAcceleratorServiceStatusResponseBody) SetCode(v string) *DescribeAcceleratorServiceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponseBody) SetMessage(v string) *DescribeAcceleratorServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponseBody) SetRequestId(v string) *DescribeAcceleratorServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponseBody) SetStatus(v string) *DescribeAcceleratorServiceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponseBody) SetSuccess(v string) *DescribeAcceleratorServiceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAcceleratorServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
