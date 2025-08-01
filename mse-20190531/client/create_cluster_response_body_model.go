// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateClusterResponseBody
	GetErrorCode() *string
	SetInstanceId(v string) *CreateClusterResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateClusterResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreateClusterResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateClusterResponseBody
	GetSuccess() *bool
}

type CreateClusterResponseBody struct {
	// Error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Return message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Order ID.
	//
	// example:
	//
	// 20574710974****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// dc63-465d-8ef5-20dc18af****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request result, with the following values:
	//
	// - `true`: Request succeeded.
	//
	// - `false`: Request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateClusterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateClusterResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateClusterResponseBody) SetErrorCode(v string) *CreateClusterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateClusterResponseBody) SetInstanceId(v string) *CreateClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateClusterResponseBody) SetMessage(v string) *CreateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *CreateClusterResponseBody) SetOrderId(v string) *CreateClusterResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetSuccess(v bool) *CreateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
