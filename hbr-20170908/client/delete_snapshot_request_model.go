// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DeleteSnapshotRequest
	GetClientId() *string
	SetForce(v bool) *DeleteSnapshotRequest
	GetForce() *bool
	SetInstanceId(v string) *DeleteSnapshotRequest
	GetInstanceId() *string
	SetSnapshotId(v string) *DeleteSnapshotRequest
	GetSnapshotId() *string
	SetSourceType(v string) *DeleteSnapshotRequest
	GetSourceType() *string
	SetToken(v string) *DeleteSnapshotRequest
	GetToken() *string
	SetVaultId(v string) *DeleteSnapshotRequest
	GetVaultId() *string
}

type DeleteSnapshotRequest struct {
	// The ID of the Cloud Backup client. If you delete a backup snapshot for Elastic Compute Service (ECS) instances, you must specify one of the ClientId and **InstanceId*	- parameters.
	//
	// example:
	//
	// c-*********************
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprected.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the ECS instance. If you delete a backup snapshot for ECS instances, you must specify one of the InstanceId and **ClientId*	- parameters.
	//
	// example:
	//
	// i-*********************
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the backup snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-*********************
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the backup source. Valid values:
	//
	// 	- **ECS_FILE**: backup snapshots for ECS files
	//
	// 	- **OSS**: backup snapshots for Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: backup snapshots for Apsara File Storage NAS (NAS) file systems
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The token that you want to delete.
	//
	// example:
	//
	// 02WJDOE7
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DeleteSnapshotRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteSnapshotRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DeleteSnapshotRequest) GetToken() *string {
	return s.Token
}

func (s *DeleteSnapshotRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *DeleteSnapshotRequest) SetClientId(v string) *DeleteSnapshotRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetForce(v bool) *DeleteSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSnapshotRequest) SetInstanceId(v string) *DeleteSnapshotRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSourceType(v string) *DeleteSnapshotRequest {
	s.SourceType = &v
	return s
}

func (s *DeleteSnapshotRequest) SetToken(v string) *DeleteSnapshotRequest {
	s.Token = &v
	return s
}

func (s *DeleteSnapshotRequest) SetVaultId(v string) *DeleteSnapshotRequest {
	s.VaultId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
