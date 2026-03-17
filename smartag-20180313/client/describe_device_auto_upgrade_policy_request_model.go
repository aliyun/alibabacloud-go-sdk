// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceAutoUpgradePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDeviceAutoUpgradePolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDeviceAutoUpgradePolicyRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetSmartAGId() *string
	SetVersionType(v string) *DescribeDeviceAutoUpgradePolicyRequest
	GetVersionType() *string
}

type DescribeDeviceAutoUpgradePolicyRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x022502****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-kxe2cv7hot7qrv****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The type of the software for which you want to query the automatic upgrade policy. Valid values:
	//
	// 	- **Device**: The operating system that is run by the SAG instance.
	//
	// 	- **Dpi**: The signature database that is used by the SAG instance.
	//
	// example:
	//
	// Device
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s DescribeDeviceAutoUpgradePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceAutoUpgradePolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetOwnerAccount(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetOwnerId(v int64) *DescribeDeviceAutoUpgradePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetRegionId(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetResourceOwnerAccount(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetResourceOwnerId(v int64) *DescribeDeviceAutoUpgradePolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetSerialNumber(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetSmartAGId(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) SetVersionType(v string) *DescribeDeviceAutoUpgradePolicyRequest {
	s.VersionType = &v
	return s
}

func (s *DescribeDeviceAutoUpgradePolicyRequest) Validate() error {
	return dara.Validate(s)
}
