// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAccessGatewayVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *UpdateSmartAccessGatewayVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateSmartAccessGatewayVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateSmartAccessGatewayVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateSmartAccessGatewayVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmartAccessGatewayVersionRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *UpdateSmartAccessGatewayVersionRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *UpdateSmartAccessGatewayVersionRequest
	GetSmartAGId() *string
	SetVersionCode(v string) *UpdateSmartAccessGatewayVersionRequest
	GetVersionCode() *string
	SetVersionType(v string) *UpdateSmartAccessGatewayVersionRequest
	GetVersionType() *string
}

type UpdateSmartAccessGatewayVersionRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG device.
	//
	// example:
	//
	// sag233****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0ovhf732a9j*******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The version to which you want to upgrade the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
	// The type of software that you want to upgrade. Valid values:
	//
	// 	- **Device**: The operating system run by the SAG device.
	//
	// 	- **Dpi**: The signature database used by the SAG device.
	//
	// example:
	//
	// Device
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s UpdateSmartAccessGatewayVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAccessGatewayVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetVersionCode() *string {
	return s.VersionCode
}

func (s *UpdateSmartAccessGatewayVersionRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetOwnerAccount(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetOwnerId(v int64) *UpdateSmartAccessGatewayVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetRegionId(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetResourceOwnerAccount(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetResourceOwnerId(v int64) *UpdateSmartAccessGatewayVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetSerialNumber(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.SerialNumber = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetSmartAGId(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetVersionCode(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) SetVersionType(v string) *UpdateSmartAccessGatewayVersionRequest {
	s.VersionType = &v
	return s
}

func (s *UpdateSmartAccessGatewayVersionRequest) Validate() error {
	return dara.Validate(s)
}
