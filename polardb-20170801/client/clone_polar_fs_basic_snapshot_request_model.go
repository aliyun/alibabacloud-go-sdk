// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolarFsBasicSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ClonePolarFsBasicSnapshotRequest
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *ClonePolarFsBasicSnapshotRequest
	GetPolarFsInstanceId() *string
	SetSourcePath(v string) *ClonePolarFsBasicSnapshotRequest
	GetSourcePath() *string
	SetTargetPath(v string) *ClonePolarFsBasicSnapshotRequest
	GetTargetPath() *string
}

type ClonePolarFsBasicSnapshotRequest struct {
	// The ID of the database cluster.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the Polarlakebase instance.
	//
	// example:
	//
	// pfs-test*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The source path of the file resource. This parameter is empty if the type is local.
	//
	// example:
	//
	// /test
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The destination path.
	//
	// example:
	//
	// /testclone
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
}

func (s ClonePolarFsBasicSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s ClonePolarFsBasicSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ClonePolarFsBasicSnapshotRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ClonePolarFsBasicSnapshotRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ClonePolarFsBasicSnapshotRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ClonePolarFsBasicSnapshotRequest) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ClonePolarFsBasicSnapshotRequest) SetDBClusterId(v string) *ClonePolarFsBasicSnapshotRequest {
	s.DBClusterId = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotRequest) SetPolarFsInstanceId(v string) *ClonePolarFsBasicSnapshotRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotRequest) SetSourcePath(v string) *ClonePolarFsBasicSnapshotRequest {
	s.SourcePath = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotRequest) SetTargetPath(v string) *ClonePolarFsBasicSnapshotRequest {
	s.TargetPath = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
