// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateServiceAddressRequest
	GetAddress() *string
	SetAddressType(v string) *CreateServiceAddressRequest
	GetAddressType() *string
	SetRegionId(v string) *CreateServiceAddressRequest
	GetRegionId() *string
	SetSagId(v string) *CreateServiceAddressRequest
	GetSagId() *string
	SetSn(v string) *CreateServiceAddressRequest
	GetSn() *string
}

type CreateServiceAddressRequest struct {
	// The service address. Example: **192.168.1.1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.1
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The type of service address. Set the value to **ProbeTask**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProbeTask
	AddressType *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag****
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s CreateServiceAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAddressRequest) GetAddress() *string {
	return s.Address
}

func (s *CreateServiceAddressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *CreateServiceAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceAddressRequest) GetSagId() *string {
	return s.SagId
}

func (s *CreateServiceAddressRequest) GetSn() *string {
	return s.Sn
}

func (s *CreateServiceAddressRequest) SetAddress(v string) *CreateServiceAddressRequest {
	s.Address = &v
	return s
}

func (s *CreateServiceAddressRequest) SetAddressType(v string) *CreateServiceAddressRequest {
	s.AddressType = &v
	return s
}

func (s *CreateServiceAddressRequest) SetRegionId(v string) *CreateServiceAddressRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceAddressRequest) SetSagId(v string) *CreateServiceAddressRequest {
	s.SagId = &v
	return s
}

func (s *CreateServiceAddressRequest) SetSn(v string) *CreateServiceAddressRequest {
	s.Sn = &v
	return s
}

func (s *CreateServiceAddressRequest) Validate() error {
	return dara.Validate(s)
}
