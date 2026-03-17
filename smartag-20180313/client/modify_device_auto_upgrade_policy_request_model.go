// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceAutoUpgradePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetCronExpression() *string
	SetDuration(v int32) *ModifyDeviceAutoUpgradePolicyRequest
	GetDuration() *int32
	SetOwnerAccount(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDeviceAutoUpgradePolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDeviceAutoUpgradePolicyRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetSmartAGId() *string
	SetTimeZone(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetTimeZone() *string
	SetUpgradeType(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetUpgradeType() *string
	SetVersionType(v string) *ModifyDeviceAutoUpgradePolicyRequest
	GetVersionType() *string
}

type ModifyDeviceAutoUpgradePolicyRequest struct {
	// The time when upgrades start. Valid values: Set the parameter in a cron expression.
	//
	// For example, you can use `0 0 4 1/1 	- ?` to specify 04:00:00 (UTC+8) on the first day of each month.
	//
	// example:
	//
	// 0 0 4 1/1 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The time period during which upgrades are performed. Valid values: **30 to 120**.
	//
	// Unit: minutes.
	//
	// example:
	//
	// 30
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the SAG device.
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
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The time zone. Valid values:
	//
	// **Asia/Shanghai**: UTC+8 (Beijing)
	//
	// **Asia/Hong_Kong**: UTC+8 (Hong Kong)
	//
	// **Asia/Tokyo**: UTC+9 (Tokyo)
	//
	// **Australia/Sydney**: UTC+10 (Sydney)
	//
	// **Asia/Kuala_Lumpur**: UTC+8 (Kuala Lumpur)
	//
	// **Europe/Berli**: UTC+1 (Berlin)
	//
	// **Asia/Singapore**: UTC+8 (Singapore)
	//
	// **Asia/Jakarta**: UTC+7 (Jakarta)
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The update type. Valid values:
	//
	// 	- **scheduled**: scheduled upgrade.
	//
	// 	- **boot**: automatic upgrade upon device startup.
	//
	// 	- **manual**: manual upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// scheduled
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	// The type of software for which you want to modify the upgrade policy. Valid values:
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

func (s ModifyDeviceAutoUpgradePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceAutoUpgradePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetUpgradeType() *string {
	return s.UpgradeType
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) GetVersionType() *string {
	return s.VersionType
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetCronExpression(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.CronExpression = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetDuration(v int32) *ModifyDeviceAutoUpgradePolicyRequest {
	s.Duration = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetOwnerAccount(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetOwnerId(v int64) *ModifyDeviceAutoUpgradePolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetRegionId(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetResourceOwnerAccount(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetResourceOwnerId(v int64) *ModifyDeviceAutoUpgradePolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetSerialNumber(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.SerialNumber = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetSmartAGId(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetTimeZone(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.TimeZone = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetUpgradeType(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.UpgradeType = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) SetVersionType(v string) *ModifyDeviceAutoUpgradePolicyRequest {
	s.VersionType = &v
	return s
}

func (s *ModifyDeviceAutoUpgradePolicyRequest) Validate() error {
	return dara.Validate(s)
}
