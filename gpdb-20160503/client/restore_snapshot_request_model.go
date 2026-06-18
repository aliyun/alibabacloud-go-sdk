// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestoreSnapshotRequest
	GetClientToken() *string
	SetFinalizeRestore(v bool) *RestoreSnapshotRequest
	GetFinalizeRestore() *bool
	SetProjectId(v string) *RestoreSnapshotRequest
	GetProjectId() *string
	SetRegionId(v string) *RestoreSnapshotRequest
	GetRegionId() *string
	SetRestoredBranchName(v string) *RestoreSnapshotRequest
	GetRestoredBranchName() *string
	SetRestoredLsn(v string) *RestoreSnapshotRequest
	GetRestoredLsn() *string
	SetTargetBranchId(v string) *RestoreSnapshotRequest
	GetTargetBranchId() *string
}

type RestoreSnapshotRequest struct {
	// The idempotence token. Ensures that repeated requests do not execute the same operation more than once.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to complete the restoration immediately. Default value: false.
	//
	// example:
	//
	// false
	FinalizeRestore *bool `json:"FinalizeRestore,omitempty" xml:"FinalizeRestore,omitempty"`
	// The Supabase project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. Specifies the region in which to query or perform the operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the restored branch. If not specified, the backend generates a name automatically.
	//
	// example:
	//
	// restore_preview
	RestoredBranchName *string `json:"RestoredBranchName,omitempty" xml:"RestoredBranchName,omitempty"`
	// The snapshot LSN used for restoration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0/3522648
	RestoredLsn *string `json:"RestoredLsn,omitempty" xml:"RestoredLsn,omitempty"`
	// The target branch ID. If not specified, the backend selects the target branch based on the restoration process.
	//
	// example:
	//
	// br-main
	TargetBranchId *string `json:"TargetBranchId,omitempty" xml:"TargetBranchId,omitempty"`
}

func (s RestoreSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreSnapshotRequest) GoString() string {
	return s.String()
}

func (s *RestoreSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestoreSnapshotRequest) GetFinalizeRestore() *bool {
	return s.FinalizeRestore
}

func (s *RestoreSnapshotRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *RestoreSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreSnapshotRequest) GetRestoredBranchName() *string {
	return s.RestoredBranchName
}

func (s *RestoreSnapshotRequest) GetRestoredLsn() *string {
	return s.RestoredLsn
}

func (s *RestoreSnapshotRequest) GetTargetBranchId() *string {
	return s.TargetBranchId
}

func (s *RestoreSnapshotRequest) SetClientToken(v string) *RestoreSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *RestoreSnapshotRequest) SetFinalizeRestore(v bool) *RestoreSnapshotRequest {
	s.FinalizeRestore = &v
	return s
}

func (s *RestoreSnapshotRequest) SetProjectId(v string) *RestoreSnapshotRequest {
	s.ProjectId = &v
	return s
}

func (s *RestoreSnapshotRequest) SetRegionId(v string) *RestoreSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreSnapshotRequest) SetRestoredBranchName(v string) *RestoreSnapshotRequest {
	s.RestoredBranchName = &v
	return s
}

func (s *RestoreSnapshotRequest) SetRestoredLsn(v string) *RestoreSnapshotRequest {
	s.RestoredLsn = &v
	return s
}

func (s *RestoreSnapshotRequest) SetTargetBranchId(v string) *RestoreSnapshotRequest {
	s.TargetBranchId = &v
	return s
}

func (s *RestoreSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
