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
	// The new disk type. This parameter is valid only during disk migration between different dedicated block storage clusters. Only cloud_essd (ESSD) is supported.
	//
	// Default value: empty, which indicates that the disk type is not changed.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: performs a dry run without executing the request. The system checks whether required parameters are specified, the request format is valid, service limits are met, and ECS resources are sufficient. If the request fails the check, the corresponding error is returned. If the request passes the check, the DryRunOperation error code is returned.
	//
	// - false: executes the request. If the request passes the check, a 2xx HTTP status code is returned and the disk is migrated.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new ESSD performance level. This parameter is valid only during disk migration between different dedicated block storage clusters. Valid values:
	//
	// - PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// Default value: empty, which indicates that the performance level is not changed.
	//
	// example:
	//
	// PL1
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the dedicated block storage cluster. Before you migrate a disk into a dedicated block storage cluster, you must purchase and create the cluster, and ensure that the disk and the destination cluster are in the same zone.
	//
	// - To migrate a disk to a dedicated block storage cluster, you must specify `StorageClusterId`.
	//
	// - To migrate a disk to a public cloud block storage cluster, you must leave `StorageClusterId` empty.
	//
	// Default value: empty, which indicates that the disk is migrated to a public cloud block storage cluster.
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
