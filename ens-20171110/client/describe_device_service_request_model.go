// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeDeviceServiceRequest
	GetAppId() *string
	SetEnsRegionId(v string) *DescribeDeviceServiceRequest
	GetEnsRegionId() *string
	SetInstanceId(v string) *DescribeDeviceServiceRequest
	GetInstanceId() *string
	SetOrderId(v string) *DescribeDeviceServiceRequest
	GetOrderId() *string
	SetRegionId(v string) *DescribeDeviceServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *DescribeDeviceServiceRequest
	GetServiceId() *string
}

type DescribeDeviceServiceRequest struct {
	// The ID of the application.
	//
	// example:
	//
	// a2bac6f4-75dc-455e-8389-2dc8e47526d3
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter does not take effect.
	//
	// example:
	//
	// cn-chongqing-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5sg1owx0g4ojy66ab2tez77r2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2661b1dd-3453-418d-8182-bb34f79e8d3c
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-chongqing-11
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID
	//
	// example:
	//
	// s-cxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DescribeDeviceServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceServiceRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeDeviceServiceRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeDeviceServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeviceServiceRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeDeviceServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeviceServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeDeviceServiceRequest) SetAppId(v string) *DescribeDeviceServiceRequest {
	s.AppId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetEnsRegionId(v string) *DescribeDeviceServiceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetInstanceId(v string) *DescribeDeviceServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetOrderId(v string) *DescribeDeviceServiceRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetRegionId(v string) *DescribeDeviceServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) SetServiceId(v string) *DescribeDeviceServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DescribeDeviceServiceRequest) Validate() error {
	return dara.Validate(s)
}
