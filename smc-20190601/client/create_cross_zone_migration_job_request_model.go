// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrossZoneMigrationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateCrossZoneMigrationJobRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreateCrossZoneMigrationJobRequest
	GetClientToken() *string
	SetDisk(v []*CreateCrossZoneMigrationJobRequestDisk) *CreateCrossZoneMigrationJobRequest
	GetDisk() []*CreateCrossZoneMigrationJobRequestDisk
	SetInstanceId(v string) *CreateCrossZoneMigrationJobRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *CreateCrossZoneMigrationJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCrossZoneMigrationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCrossZoneMigrationJobRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCrossZoneMigrationJobRequest
	GetResourceOwnerAccount() *string
	SetTargetInstanceType(v string) *CreateCrossZoneMigrationJobRequest
	GetTargetInstanceType() *string
	SetTargetVSwitchId(v string) *CreateCrossZoneMigrationJobRequest
	GetTargetVSwitchId() *string
	SetTargetZoneId(v string) *CreateCrossZoneMigrationJobRequest
	GetTargetZoneId() *string
}

type CreateCrossZoneMigrationJobRequest struct {
	// Specifies whether to enable automatic payment. Valid values:
	//
	// 	- **true*	- (default): enables automatic payment. Make sure that you have sufficient balance within your account.
	//
	// 	- **false**: disables automatic payment. In this case, you must manually pay for the instance. For more information, see [Renew a subscription instance](https://help.aliyun.com/document_detail/85052.html).
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk list.
	Disk []*CreateCrossZoneMigrationJobRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1ff25rzvnul6kr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the destination Alibaba Cloud region.
	//
	// For example, if you want to migrate the source server to the China (Hangzhou) region, set this parameter to `cn-hangzhou`. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the new instance.
	//
	// example:
	//
	// ecs.g7.large
	TargetInstanceType *string `json:"TargetInstanceType,omitempty" xml:"TargetInstanceType,omitempty"`
	// The vSwitch ID of the destination Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// vsw-bp1mxqnssl8nafltcx32e
	TargetVSwitchId *string `json:"TargetVSwitchId,omitempty" xml:"TargetVSwitchId,omitempty"`
	// The ID of the destination zone.
	//
	// example:
	//
	// cn-hangzhou-i
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s CreateCrossZoneMigrationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossZoneMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateCrossZoneMigrationJobRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateCrossZoneMigrationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCrossZoneMigrationJobRequest) GetDisk() []*CreateCrossZoneMigrationJobRequestDisk {
	return s.Disk
}

func (s *CreateCrossZoneMigrationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCrossZoneMigrationJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCrossZoneMigrationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCrossZoneMigrationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCrossZoneMigrationJobRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCrossZoneMigrationJobRequest) GetTargetInstanceType() *string {
	return s.TargetInstanceType
}

func (s *CreateCrossZoneMigrationJobRequest) GetTargetVSwitchId() *string {
	return s.TargetVSwitchId
}

func (s *CreateCrossZoneMigrationJobRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *CreateCrossZoneMigrationJobRequest) SetAutoPay(v bool) *CreateCrossZoneMigrationJobRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetClientToken(v string) *CreateCrossZoneMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetDisk(v []*CreateCrossZoneMigrationJobRequestDisk) *CreateCrossZoneMigrationJobRequest {
	s.Disk = v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetInstanceId(v string) *CreateCrossZoneMigrationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetOwnerId(v int64) *CreateCrossZoneMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetRegionId(v string) *CreateCrossZoneMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetResourceGroupId(v string) *CreateCrossZoneMigrationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetResourceOwnerAccount(v string) *CreateCrossZoneMigrationJobRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetTargetInstanceType(v string) *CreateCrossZoneMigrationJobRequest {
	s.TargetInstanceType = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetTargetVSwitchId(v string) *CreateCrossZoneMigrationJobRequest {
	s.TargetVSwitchId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) SetTargetZoneId(v string) *CreateCrossZoneMigrationJobRequest {
	s.TargetZoneId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCrossZoneMigrationJobRequestDisk struct {
	// The disk category. A value of cloud_essd indicates enhanced SSD (ESSD).
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp1eeplkn4j29wf7irpb
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s CreateCrossZoneMigrationJobRequestDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateCrossZoneMigrationJobRequestDisk) GoString() string {
	return s.String()
}

func (s *CreateCrossZoneMigrationJobRequestDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateCrossZoneMigrationJobRequestDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *CreateCrossZoneMigrationJobRequestDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateCrossZoneMigrationJobRequestDisk) SetCategory(v string) *CreateCrossZoneMigrationJobRequestDisk {
	s.Category = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequestDisk) SetDiskId(v string) *CreateCrossZoneMigrationJobRequestDisk {
	s.DiskId = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequestDisk) SetPerformanceLevel(v string) *CreateCrossZoneMigrationJobRequestDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateCrossZoneMigrationJobRequestDisk) Validate() error {
	return dara.Validate(s)
}
