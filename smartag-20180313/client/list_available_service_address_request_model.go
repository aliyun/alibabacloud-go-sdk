// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableServiceAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressType(v string) *ListAvailableServiceAddressRequest
	GetAddressType() *string
	SetRegionId(v string) *ListAvailableServiceAddressRequest
	GetRegionId() *string
	SetSagId(v string) *ListAvailableServiceAddressRequest
	GetSagId() *string
	SetSn(v string) *ListAvailableServiceAddressRequest
	GetSn() *string
}

type ListAvailableServiceAddressRequest struct {
	// The type of service address. Valid values:
	//
	// 	- **ProbeTask**: probes the source IP address.
	//
	// 	- **RemoteWeb**: probes the IP address for remote logon.
	//
	// > If you do not specify a value, all service IP addresses are queried.
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

func (s ListAvailableServiceAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableServiceAddressRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableServiceAddressRequest) GetAddressType() *string {
	return s.AddressType
}

func (s *ListAvailableServiceAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAvailableServiceAddressRequest) GetSagId() *string {
	return s.SagId
}

func (s *ListAvailableServiceAddressRequest) GetSn() *string {
	return s.Sn
}

func (s *ListAvailableServiceAddressRequest) SetAddressType(v string) *ListAvailableServiceAddressRequest {
	s.AddressType = &v
	return s
}

func (s *ListAvailableServiceAddressRequest) SetRegionId(v string) *ListAvailableServiceAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ListAvailableServiceAddressRequest) SetSagId(v string) *ListAvailableServiceAddressRequest {
	s.SagId = &v
	return s
}

func (s *ListAvailableServiceAddressRequest) SetSn(v string) *ListAvailableServiceAddressRequest {
	s.Sn = &v
	return s
}

func (s *ListAvailableServiceAddressRequest) Validate() error {
	return dara.Validate(s)
}
