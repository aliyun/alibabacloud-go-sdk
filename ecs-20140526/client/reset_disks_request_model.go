// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisk(v []*ResetDisksRequestDisk) *ResetDisksRequest
	GetDisk() []*ResetDisksRequestDisk
	SetDryRun(v bool) *ResetDisksRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ResetDisksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetDisksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ResetDisksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ResetDisksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetDisksRequest
	GetResourceOwnerId() *int64
}

type ResetDisksRequest struct {
	// The list of cloud disks.
	//
	// This parameter is required.
	Disk []*ResetDisksRequestDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: performs a dry run without actually rolling back the cloud disks. The system checks whether required parameters are specified, whether the request format is valid, and whether resource status constraints are met. If the check fails, the corresponding error message is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - false: performs a dry run and sends the request. If the check succeeds, the cloud disk rollback operation is initiated.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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

func (s ResetDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksRequest) GoString() string {
	return s.String()
}

func (s *ResetDisksRequest) GetDisk() []*ResetDisksRequestDisk {
	return s.Disk
}

func (s *ResetDisksRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResetDisksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetDisksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetDisksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDisksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetDisksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetDisksRequest) SetDisk(v []*ResetDisksRequestDisk) *ResetDisksRequest {
	s.Disk = v
	return s
}

func (s *ResetDisksRequest) SetDryRun(v bool) *ResetDisksRequest {
	s.DryRun = &v
	return s
}

func (s *ResetDisksRequest) SetOwnerAccount(v string) *ResetDisksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetDisksRequest) SetOwnerId(v int64) *ResetDisksRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetDisksRequest) SetRegionId(v string) *ResetDisksRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDisksRequest) SetResourceOwnerAccount(v string) *ResetDisksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetDisksRequest) SetResourceOwnerId(v int64) *ResetDisksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetDisksRequest) Validate() error {
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

type ResetDisksRequestDisk struct {
	// The ID of the cloud disk to be rolled back. Valid values of N: 1 to 10.
	//
	// example:
	//
	// d-j6cf7l0ewidb78lq****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The snapshot ID that corresponds to the specified cloud disk in the instance snapshot. Valid values of N: 1 to 10.
	//
	// example:
	//
	// s-j6cdofbycydvg7ey****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDisksRequestDisk) String() string {
	return dara.Prettify(s)
}

func (s ResetDisksRequestDisk) GoString() string {
	return s.String()
}

func (s *ResetDisksRequestDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *ResetDisksRequestDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetDisksRequestDisk) SetDiskId(v string) *ResetDisksRequestDisk {
	s.DiskId = &v
	return s
}

func (s *ResetDisksRequestDisk) SetSnapshotId(v string) *ResetDisksRequestDisk {
	s.SnapshotId = &v
	return s
}

func (s *ResetDisksRequestDisk) Validate() error {
	return dara.Validate(s)
}
