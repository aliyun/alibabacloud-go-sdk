// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DeleteServiceAddressRequest
	GetAddress() *string
	SetAddressType(v string) *DeleteServiceAddressRequest
	GetAddressType() *string
	SetRegionId(v string) *DeleteServiceAddressRequest
	GetRegionId() *string
	SetSagId(v string) *DeleteServiceAddressRequest
	GetSagId() *string
	SetSn(v string) *DeleteServiceAddressRequest
	GetSn() *string
}

type DeleteServiceAddressRequest struct {
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
	// The ID of the region where the SAG instance is deployed.
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

func (s DeleteServiceAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceAddressRequest) GetAddress() *string {
	return s.Address
}

func (s *DeleteServiceAddressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *DeleteServiceAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteServiceAddressRequest) GetSagId() *string {
	return s.SagId
}

func (s *DeleteServiceAddressRequest) GetSn() *string {
	return s.Sn
}

func (s *DeleteServiceAddressRequest) SetAddress(v string) *DeleteServiceAddressRequest {
	s.Address = &v
	return s
}

func (s *DeleteServiceAddressRequest) SetAddressType(v string) *DeleteServiceAddressRequest {
	s.AddressType = &v
	return s
}

func (s *DeleteServiceAddressRequest) SetRegionId(v string) *DeleteServiceAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceAddressRequest) SetSagId(v string) *DeleteServiceAddressRequest {
	s.SagId = &v
	return s
}

func (s *DeleteServiceAddressRequest) SetSn(v string) *DeleteServiceAddressRequest {
	s.Sn = &v
	return s
}

func (s *DeleteServiceAddressRequest) Validate() error {
	return dara.Validate(s)
}
