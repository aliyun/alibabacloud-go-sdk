// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ResetDiskRequest
	GetDiskId() *string
	SetDryRun(v bool) *ResetDiskRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ResetDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ResetDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ResetDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ResetDiskRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *ResetDiskRequest
	GetSnapshotId() *string
}

type ResetDiskRequest struct {
	// The ID of the cloud disk to be rolled back.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp199lyny9b3****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - true: performs a dry run without actually rolling back the cloud disk. The system checks whether required parameters are specified, whether the request format is valid, and whether resource status constraints are met. If the check fails, the corresponding error message is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// - false: performs a dry run and sends the request. If the check succeeds, the cloud disk rollback operation is initiated.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the snapshot to use for rolling back the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp199lyny9b3****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResetDiskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ResetDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ResetDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ResetDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ResetDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ResetDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetDiskRequest) SetDiskId(v string) *ResetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResetDiskRequest) SetDryRun(v bool) *ResetDiskRequest {
	s.DryRun = &v
	return s
}

func (s *ResetDiskRequest) SetOwnerAccount(v string) *ResetDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetDiskRequest) SetOwnerId(v int64) *ResetDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetDiskRequest) SetResourceOwnerAccount(v string) *ResetDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetDiskRequest) SetResourceOwnerId(v int64) *ResetDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetDiskRequest) SetSnapshotId(v string) *ResetDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetDiskRequest) Validate() error {
	return dara.Validate(s)
}
