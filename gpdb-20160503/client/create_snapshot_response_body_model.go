// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchId(v string) *CreateSnapshotResponseBody
	GetBranchId() *string
	SetLsn(v string) *CreateSnapshotResponseBody
	GetLsn() *string
	SetRequestId(v string) *CreateSnapshotResponseBody
	GetRequestId() *string
	SetTimestamp(v string) *CreateSnapshotResponseBody
	GetTimestamp() *string
}

type CreateSnapshotResponseBody struct {
	// The branch ID to which the snapshot belongs.
	//
	// example:
	//
	// br-xxxx
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The LSN for the snapshot. You must specify either this parameter or SnapshotTimestamp. If this parameter is specified, the snapshot is created based on the specified LSN.
	//
	// example:
	//
	// 0/3522648
	Lsn *string `json:"Lsn,omitempty" xml:"Lsn,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The actual point in time that corresponds to the created snapshot.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *CreateSnapshotResponseBody) GetLsn() *string {
	return s.Lsn
}

func (s *CreateSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSnapshotResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *CreateSnapshotResponseBody) SetBranchId(v string) *CreateSnapshotResponseBody {
	s.BranchId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetLsn(v string) *CreateSnapshotResponseBody {
	s.Lsn = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetTimestamp(v string) *CreateSnapshotResponseBody {
	s.Timestamp = &v
	return s
}

func (s *CreateSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
