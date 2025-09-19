// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludes(v string) *CreateRestoreJobRequest
	GetIncludes() *string
	SetSnapshotHash(v string) *CreateRestoreJobRequest
	GetSnapshotHash() *string
	SetSnapshotId(v string) *CreateRestoreJobRequest
	GetSnapshotId() *string
	SetSnapshotVersion(v string) *CreateRestoreJobRequest
	GetSnapshotVersion() *string
	SetSourceType(v string) *CreateRestoreJobRequest
	GetSourceType() *string
	SetTarget(v string) *CreateRestoreJobRequest
	GetTarget() *string
	SetUuid(v string) *CreateRestoreJobRequest
	GetUuid() *string
	SetVaultId(v string) *CreateRestoreJobRequest
	GetVaultId() *string
}

type CreateRestoreJobRequest struct {
	// The directory in which the files included in the restoration task are located. This parameter is specified when you create the anti-ransomware policy. The value is a directory that requires protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["root"]
	Includes *string `json:"Includes,omitempty" xml:"Includes,omitempty"`
	// The hash value of the snapshot.
	//
	// > You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// d4b399edaad94b038e8f91873f19e3eae010ca30798fc36db3a164dd343f****
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the snapshot that you want to use for restoration.
	//
	// > You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-00023dhaatxp18mh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The version of the backup data.
	//
	// > You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 165570544****
	SnapshotVersion *string `json:"SnapshotVersion,omitempty" xml:"SnapshotVersion,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots for Elastic Compute Service (ECS) files
	//
	// 	- **File**: backup snapshots for on-premises servers
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The path to which you want to restore data.
	//
	// This parameter is required.
	//
	// example:
	//
	// /root/testfls
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The UUID of the server whose data you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-ecs-4e876cb0-09f7-43b8-82ef-4bc7a93769b5
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the backup vault that is used in the restoration task.
	//
	// > You can call the [DescribeSnapshots](~~DescribeSnapshots~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// v-0002n12wokck2q0x****
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateRestoreJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobRequest) GetIncludes() *string {
	return s.Includes
}

func (s *CreateRestoreJobRequest) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *CreateRestoreJobRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRestoreJobRequest) GetSnapshotVersion() *string {
	return s.SnapshotVersion
}

func (s *CreateRestoreJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRestoreJobRequest) GetTarget() *string {
	return s.Target
}

func (s *CreateRestoreJobRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateRestoreJobRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateRestoreJobRequest) SetIncludes(v string) *CreateRestoreJobRequest {
	s.Includes = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotHash(v string) *CreateRestoreJobRequest {
	s.SnapshotHash = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotId(v string) *CreateRestoreJobRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotVersion(v string) *CreateRestoreJobRequest {
	s.SnapshotVersion = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSourceType(v string) *CreateRestoreJobRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTarget(v string) *CreateRestoreJobRequest {
	s.Target = &v
	return s
}

func (s *CreateRestoreJobRequest) SetUuid(v string) *CreateRestoreJobRequest {
	s.Uuid = &v
	return s
}

func (s *CreateRestoreJobRequest) SetVaultId(v string) *CreateRestoreJobRequest {
	s.VaultId = &v
	return s
}

func (s *CreateRestoreJobRequest) Validate() error {
	return dara.Validate(s)
}
