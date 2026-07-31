// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *ModifyInstanceSpecRequestSystemDisk) *ModifyInstanceSpecRequest
	GetSystemDisk() *ModifyInstanceSpecRequestSystemDisk
	SetTemporary(v *ModifyInstanceSpecRequestTemporary) *ModifyInstanceSpecRequest
	GetTemporary() *ModifyInstanceSpecRequestTemporary
	SetAllowMigrateAcrossZone(v bool) *ModifyInstanceSpecRequest
	GetAllowMigrateAcrossZone() *bool
	SetAsync(v bool) *ModifyInstanceSpecRequest
	GetAsync() *bool
	SetClientToken(v string) *ModifyInstanceSpecRequest
	GetClientToken() *string
	SetDisk(v []*ModifyInstanceSpecRequestDisk) *ModifyInstanceSpecRequest
	GetDisk() []*ModifyInstanceSpecRequestDisk
	SetDryRun(v bool) *ModifyInstanceSpecRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyInstanceSpecRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyInstanceSpecRequest
	GetInstanceType() *string
	SetInternetMaxBandwidthIn(v int32) *ModifyInstanceSpecRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *ModifyInstanceSpecRequest
	GetInternetMaxBandwidthOut() *int32
	SetModifyMode(v string) *ModifyInstanceSpecRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifyInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceSpecRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceSpecRequest struct {
	SystemDisk *ModifyInstanceSpecRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	Temporary  *ModifyInstanceSpecRequestTemporary  `json:"Temporary,omitempty" xml:"Temporary,omitempty" type:"Struct"`
	// Specifies whether cross-cluster instance type upgrades are supported.
	//
	// example:
	//
	// false
	AllowMigrateAcrossZone *bool `json:"AllowMigrateAcrossZone,omitempty" xml:"AllowMigrateAcrossZone,omitempty"`
	// Specifies whether to submit an asynchronous request. Valid values:
	//
	// example:
	//
	// false
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// > This parameter is not publicly available.
	Disk []*ModifyInstanceSpecRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target instance type of the instance. For more information, see [Instance family](https://help.aliyun.com/document_detail/25378.html). You can also invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) to query the most recent instance type list.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s (Megabit per second). Valid values:
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s (Megabit per second). Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) GetSystemDisk() *ModifyInstanceSpecRequestSystemDisk {
	return s.SystemDisk
}

func (s *ModifyInstanceSpecRequest) GetTemporary() *ModifyInstanceSpecRequestTemporary {
	return s.Temporary
}

func (s *ModifyInstanceSpecRequest) GetAllowMigrateAcrossZone() *bool {
	return s.AllowMigrateAcrossZone
}

func (s *ModifyInstanceSpecRequest) GetAsync() *bool {
	return s.Async
}

func (s *ModifyInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceSpecRequest) GetDisk() []*ModifyInstanceSpecRequestDisk {
	return s.Disk
}

func (s *ModifyInstanceSpecRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSpecRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyInstanceSpecRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *ModifyInstanceSpecRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyInstanceSpecRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceSpecRequest) SetSystemDisk(v *ModifyInstanceSpecRequestSystemDisk) *ModifyInstanceSpecRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetTemporary(v *ModifyInstanceSpecRequestTemporary) *ModifyInstanceSpecRequest {
	s.Temporary = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetAllowMigrateAcrossZone(v bool) *ModifyInstanceSpecRequest {
	s.AllowMigrateAcrossZone = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetAsync(v bool) *ModifyInstanceSpecRequest {
	s.Async = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetClientToken(v string) *ModifyInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetDisk(v []*ModifyInstanceSpecRequestDisk) *ModifyInstanceSpecRequest {
	s.Disk = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetDryRun(v bool) *ModifyInstanceSpecRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceType(v string) *ModifyInstanceSpecRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInternetMaxBandwidthIn(v int32) *ModifyInstanceSpecRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInternetMaxBandwidthOut(v int32) *ModifyInstanceSpecRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetModifyMode(v string) *ModifyInstanceSpecRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Temporary != nil {
		if err := s.Temporary.Validate(); err != nil {
			return err
		}
	}
	if s.Disk != nil {
		for _, item := range s.Disk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceSpecRequestSystemDisk struct {
	// The new system disk category. Valid values:
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ModifyInstanceSpecRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyInstanceSpecRequestSystemDisk) SetCategory(v string) *ModifyInstanceSpecRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyInstanceSpecRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceSpecRequestTemporary struct {
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyInstanceSpecRequestTemporary) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequestTemporary) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequestTemporary) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyInstanceSpecRequestTemporary) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyInstanceSpecRequestTemporary) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyInstanceSpecRequestTemporary) SetEndTime(v string) *ModifyInstanceSpecRequestTemporary {
	s.EndTime = &v
	return s
}

func (s *ModifyInstanceSpecRequestTemporary) SetInternetMaxBandwidthOut(v int32) *ModifyInstanceSpecRequestTemporary {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyInstanceSpecRequestTemporary) SetStartTime(v string) *ModifyInstanceSpecRequestTemporary {
	s.StartTime = &v
	return s
}

func (s *ModifyInstanceSpecRequestTemporary) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceSpecRequestDisk struct {
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// > This parameter is not publicly available.
	//
	// example:
	//
	// null
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s ModifyInstanceSpecRequestDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequestDisk) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequestDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyInstanceSpecRequestDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyInstanceSpecRequestDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyInstanceSpecRequestDisk) SetCategory(v string) *ModifyInstanceSpecRequestDisk {
	s.Category = &v
	return s
}

func (s *ModifyInstanceSpecRequestDisk) SetDiskId(v string) *ModifyInstanceSpecRequestDisk {
	s.DiskId = &v
	return s
}

func (s *ModifyInstanceSpecRequestDisk) SetPerformanceLevel(v string) *ModifyInstanceSpecRequestDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyInstanceSpecRequestDisk) Validate() error {
	return dara.Validate(s)
}
