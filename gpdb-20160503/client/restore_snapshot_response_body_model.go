// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *RestoreSnapshotResponseBody
	GetBranchId() *string
	SetProjectId(v string) *RestoreSnapshotResponseBody
	GetProjectId() *string
	SetRequestId(v string) *RestoreSnapshotResponseBody
	GetRequestId() *string
}

type RestoreSnapshotResponseBody struct {
	// The ID of the restored branch.
	//
	// example:
	//
	// br-restore
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The Supabase project ID.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreSnapshotResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *RestoreSnapshotResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *RestoreSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreSnapshotResponseBody) SetBranchId(v string) *RestoreSnapshotResponseBody {
	s.BranchId = &v
	return s
}

func (s *RestoreSnapshotResponseBody) SetProjectId(v string) *RestoreSnapshotResponseBody {
	s.ProjectId = &v
	return s
}

func (s *RestoreSnapshotResponseBody) SetRequestId(v string) *RestoreSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
