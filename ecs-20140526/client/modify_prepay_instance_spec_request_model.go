// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *ModifyPrepayInstanceSpecRequestSystemDisk) *ModifyPrepayInstanceSpecRequest
	GetSystemDisk() *ModifyPrepayInstanceSpecRequestSystemDisk
	SetAutoPay(v bool) *ModifyPrepayInstanceSpecRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyPrepayInstanceSpecRequest
	GetClientToken() *string
	SetDisk(v []*ModifyPrepayInstanceSpecRequestDisk) *ModifyPrepayInstanceSpecRequest
	GetDisk() []*ModifyPrepayInstanceSpecRequestDisk
	SetEndTime(v string) *ModifyPrepayInstanceSpecRequest
	GetEndTime() *string
	SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest
	GetInstanceId() *string
	SetInstanceType(v string) *ModifyPrepayInstanceSpecRequest
	GetInstanceType() *string
	SetMigrateAcrossZone(v bool) *ModifyPrepayInstanceSpecRequest
	GetMigrateAcrossZone() *bool
	SetModifyMode(v string) *ModifyPrepayInstanceSpecRequest
	GetModifyMode() *string
	SetOperatorType(v string) *ModifyPrepayInstanceSpecRequest
	GetOperatorType() *string
	SetOwnerAccount(v string) *ModifyPrepayInstanceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPrepayInstanceSpecRequest
	GetOwnerId() *int64
	SetRebootTime(v string) *ModifyPrepayInstanceSpecRequest
	GetRebootTime() *string
	SetRebootWhenFinished(v bool) *ModifyPrepayInstanceSpecRequest
	GetRebootWhenFinished() *bool
	SetRegionId(v string) *ModifyPrepayInstanceSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPrepayInstanceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPrepayInstanceSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyPrepayInstanceSpecRequest struct {
	SystemDisk *ModifyPrepayInstanceSpecRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether automatic payment is enabled when you upgrade the instance type. Valid values:
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// >This parameter is not publicly available.
	Disk []*ModifyPrepayInstanceSpecRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// The end time of the temporary specification change. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-01-01T12:05Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The target instance type for the Upgrade/Downgrade. For more information, see [Instance family](https://help.aliyun.com/document_detail/25378.html) or invoke [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g5.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Specifies whether cross-cluster instance type upgrades are supported. Valid values:
	//
	// example:
	//
	// false
	MigrateAcrossZone *bool `json:"MigrateAcrossZone,omitempty" xml:"MigrateAcrossZone,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The type of the operation. Valid values:
	//
	// example:
	//
	// upgrade
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The restart time of the instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-01-01T12:05Z
	RebootTime *string `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
	// Specifies whether to immediately restart the instance after the specification change is complete. Valid values:
	//
	// example:
	//
	// false
	RebootWhenFinished *bool `json:"RebootWhenFinished,omitempty" xml:"RebootWhenFinished,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) GetSystemDisk() *ModifyPrepayInstanceSpecRequestSystemDisk {
	return s.SystemDisk
}

func (s *ModifyPrepayInstanceSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyPrepayInstanceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPrepayInstanceSpecRequest) GetDisk() []*ModifyPrepayInstanceSpecRequestDisk {
	return s.Disk
}

func (s *ModifyPrepayInstanceSpecRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyPrepayInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayInstanceSpecRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyPrepayInstanceSpecRequest) GetMigrateAcrossZone() *bool {
	return s.MigrateAcrossZone
}

func (s *ModifyPrepayInstanceSpecRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyPrepayInstanceSpecRequest) GetOperatorType() *string {
	return s.OperatorType
}

func (s *ModifyPrepayInstanceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPrepayInstanceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPrepayInstanceSpecRequest) GetRebootTime() *string {
	return s.RebootTime
}

func (s *ModifyPrepayInstanceSpecRequest) GetRebootWhenFinished() *bool {
	return s.RebootWhenFinished
}

func (s *ModifyPrepayInstanceSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPrepayInstanceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPrepayInstanceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPrepayInstanceSpecRequest) SetSystemDisk(v *ModifyPrepayInstanceSpecRequestSystemDisk) *ModifyPrepayInstanceSpecRequest {
	s.SystemDisk = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetAutoPay(v bool) *ModifyPrepayInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetClientToken(v string) *ModifyPrepayInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetDisk(v []*ModifyPrepayInstanceSpecRequestDisk) *ModifyPrepayInstanceSpecRequest {
	s.Disk = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetEndTime(v string) *ModifyPrepayInstanceSpecRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceType(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceType = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetMigrateAcrossZone(v bool) *ModifyPrepayInstanceSpecRequest {
	s.MigrateAcrossZone = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetModifyMode(v string) *ModifyPrepayInstanceSpecRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetOperatorType(v string) *ModifyPrepayInstanceSpecRequest {
	s.OperatorType = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetOwnerAccount(v string) *ModifyPrepayInstanceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetOwnerId(v int64) *ModifyPrepayInstanceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetRebootTime(v string) *ModifyPrepayInstanceSpecRequest {
	s.RebootTime = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetRebootWhenFinished(v bool) *ModifyPrepayInstanceSpecRequest {
	s.RebootWhenFinished = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetRegionId(v string) *ModifyPrepayInstanceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetResourceOwnerAccount(v string) *ModifyPrepayInstanceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetResourceOwnerId(v int64) *ModifyPrepayInstanceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
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

type ModifyPrepayInstanceSpecRequestSystemDisk struct {
	// The new category of the system disk. Valid values:
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyPrepayInstanceSpecRequestSystemDisk) SetCategory(v string) *ModifyPrepayInstanceSpecRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type ModifyPrepayInstanceSpecRequestDisk struct {
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// >This parameter is not publicly available.
	//
	// example:
	//
	// null
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestDisk) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestDisk) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestDisk) GetCategory() *string {
	return s.Category
}

func (s *ModifyPrepayInstanceSpecRequestDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyPrepayInstanceSpecRequestDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyPrepayInstanceSpecRequestDisk) SetCategory(v string) *ModifyPrepayInstanceSpecRequestDisk {
	s.Category = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestDisk) SetDiskId(v string) *ModifyPrepayInstanceSpecRequestDisk {
	s.DiskId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestDisk) SetPerformanceLevel(v string) *ModifyPrepayInstanceSpecRequestDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestDisk) Validate() error {
	return dara.Validate(s)
}
