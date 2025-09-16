// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderInfo(v *CreateInstanceResponseBodyOrderInfo) *CreateInstanceResponseBody
	GetOrderInfo() *CreateInstanceResponseBodyOrderInfo
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceResponseBody
	GetSuccess() *bool
}

type CreateInstanceResponseBody struct {
	OrderInfo *CreateInstanceResponseBodyOrderInfo `json:"OrderInfo,omitempty" xml:"OrderInfo,omitempty" type:"Struct"`
	// example:
	//
	// 67F33190-946B-1105-B6A1-E2DF0426DD51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetOrderInfo() *CreateInstanceResponseBodyOrderInfo {
	return s.OrderInfo
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceResponseBody) SetOrderInfo(v *CreateInstanceResponseBodyOrderInfo) *CreateInstanceResponseBody {
	s.OrderInfo = v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceResponseBodyOrderInfo struct {
	// example:
	//
	// f-cn-zvp2q0zik06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 210406354694567
	OrderId           *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	StorageInstanceId *string `json:"StorageInstanceId,omitempty" xml:"StorageInstanceId,omitempty"`
	StorageOrderId    *int64  `json:"StorageOrderId,omitempty" xml:"StorageOrderId,omitempty"`
}

func (s CreateInstanceResponseBodyOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyOrderInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyOrderInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyOrderInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateInstanceResponseBodyOrderInfo) GetStorageInstanceId() *string {
	return s.StorageInstanceId
}

func (s *CreateInstanceResponseBodyOrderInfo) GetStorageOrderId() *int64 {
	return s.StorageOrderId
}

func (s *CreateInstanceResponseBodyOrderInfo) SetInstanceId(v string) *CreateInstanceResponseBodyOrderInfo {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) SetOrderId(v int64) *CreateInstanceResponseBodyOrderInfo {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) SetStorageInstanceId(v string) *CreateInstanceResponseBodyOrderInfo {
	s.StorageInstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) SetStorageOrderId(v int64) *CreateInstanceResponseBodyOrderInfo {
	s.StorageOrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyOrderInfo) Validate() error {
	return dara.Validate(s)
}
