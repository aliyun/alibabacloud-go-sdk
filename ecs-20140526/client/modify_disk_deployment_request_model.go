// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskCategory(v string) *ModifyDiskDeploymentRequest
	GetDiskCategory() *string
	SetDiskId(v string) *ModifyDiskDeploymentRequest
	GetDiskId() *string
	SetDryRun(v bool) *ModifyDiskDeploymentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyDiskDeploymentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskDeploymentRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *ModifyDiskDeploymentRequest
	GetPerformanceLevel() *string
	SetResourceOwnerAccount(v string) *ModifyDiskDeploymentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskDeploymentRequest
	GetResourceOwnerId() *int64
	SetStorageClusterId(v string) *ModifyDiskDeploymentRequest
	GetStorageClusterId() *string
}

type ModifyDiskDeploymentRequest struct {
	// The new disk type. This parameter is valid only when you migrate a disk between different dedicated block storage clusters. Only cloud_essd (standard SSD) is supported.
	//
	// Default value: empty, which indicates that the disk type is not changed (no Upgrade/Downgrade) during migration.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values:
	//
	// - true: performs a dry run. The system checks whether the required parameters are specified, the request format is valid, business limits are met, and ECS inventory is sufficient. If the check fails, the corresponding error is returned. If the check passes, the error code DryRunOperation is returned.
	//
	// - false: performs a normal request. After the check passes, a 2XX HTTP status code is returned and the disk migration starts immediately.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new ESSD performance level of the standard SSD. This parameter is valid only when you migrate a disk between different dedicated block storage clusters. Valid values:
	//
	// - PL0: maximum random read/write IOPS of 10,000 for a single disk.
	//
	// - PL1: maximum random read/write IOPS of 50,000 for a single disk.
	//
	// Default value: empty, which indicates that the performance level is not changed during migration.
	//
	// example:
	//
	// PL1
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The dedicated block storage cluster ID.
	//
	// - If you migrate the disk to a dedicated block storage cluster, you must specify `StorageClusterId`.
	//
	// - If you migrate the disk to a public block storage cluster, `StorageClusterId` must be empty.
	//
	// Default value: empty, which indicates that the disk is migrated to a public block storage cluster.
	//
	// example:
	//
	// dbsc-cn-c4d2uea****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s ModifyDiskDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskDeploymentRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ModifyDiskDeploymentRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyDiskDeploymentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDiskDeploymentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskDeploymentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskDeploymentRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyDiskDeploymentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskDeploymentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskDeploymentRequest) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *ModifyDiskDeploymentRequest) SetDiskCategory(v string) *ModifyDiskDeploymentRequest {
	s.DiskCategory = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetDiskId(v string) *ModifyDiskDeploymentRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetDryRun(v bool) *ModifyDiskDeploymentRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetOwnerAccount(v string) *ModifyDiskDeploymentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetOwnerId(v int64) *ModifyDiskDeploymentRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetPerformanceLevel(v string) *ModifyDiskDeploymentRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetResourceOwnerAccount(v string) *ModifyDiskDeploymentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetResourceOwnerId(v int64) *ModifyDiskDeploymentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetStorageClusterId(v string) *ModifyDiskDeploymentRequest {
	s.StorageClusterId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
