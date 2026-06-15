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
	// The new category of the disk. This parameter is valid only when you migrate a disk between different dedicated block storage clusters. The only valid value is `cloud_essd` (ESSD disk).
	//
	// Default value: An empty string. If you leave this parameter empty, the category of the disk remains unchanged.
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
	// - `true`: Performs a dry run. The system checks the request for required parameters, format, service limits, and inventory. The system returns an error if the check fails, or the `DryRunOperation` error code if the check succeeds.
	//
	// - `false`: Sends the request. If the request passes the check, the system returns a 2xx HTTP status code and migrates the disk.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new performance level of the ESSD disk. This parameter is valid only when you migrate a disk between different dedicated block storage clusters. Valid values:
	//
	// - `PL0`: A maximum of 10,000 random read/write IOPS per disk.
	//
	// - `PL1`: A maximum of 50,000 random read/write IOPS per disk.
	//
	// Default value: An empty string. If you leave this parameter empty, the performance level of the disk remains unchanged.
	//
	// example:
	//
	// PL1
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the dedicated block storage cluster.
	//
	// - To migrate the disk to a dedicated block storage cluster, specify `StorageClusterId`.
	//
	// - To migrate the disk to a public cloud block storage cluster, leave `StorageClusterId` empty.
	//
	// Default value: An empty string. If you leave this parameter empty, the disk is migrated to a public cloud block storage cluster.
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
