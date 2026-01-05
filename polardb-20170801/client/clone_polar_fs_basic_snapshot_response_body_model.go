// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolarFsBasicSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ClonePolarFsBasicSnapshotResponseBody
	GetDBClusterId() *string
	SetPolarFsInstanceId(v string) *ClonePolarFsBasicSnapshotResponseBody
	GetPolarFsInstanceId() *string
	SetRequestId(v string) *ClonePolarFsBasicSnapshotResponseBody
	GetRequestId() *string
	SetSourcePath(v string) *ClonePolarFsBasicSnapshotResponseBody
	GetSourcePath() *string
	SetTargetPath(v string) *ClonePolarFsBasicSnapshotResponseBody
	GetTargetPath() *string
}

type ClonePolarFsBasicSnapshotResponseBody struct {
	// example:
	//
	// pc-bp150t3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// pfs-2ze0i7*****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// /test
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// example:
	//
	// /testclone
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
}

func (s ClonePolarFsBasicSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClonePolarFsBasicSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ClonePolarFsBasicSnapshotResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ClonePolarFsBasicSnapshotResponseBody) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *ClonePolarFsBasicSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClonePolarFsBasicSnapshotResponseBody) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ClonePolarFsBasicSnapshotResponseBody) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ClonePolarFsBasicSnapshotResponseBody) SetDBClusterId(v string) *ClonePolarFsBasicSnapshotResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponseBody) SetPolarFsInstanceId(v string) *ClonePolarFsBasicSnapshotResponseBody {
	s.PolarFsInstanceId = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponseBody) SetRequestId(v string) *ClonePolarFsBasicSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponseBody) SetSourcePath(v string) *ClonePolarFsBasicSnapshotResponseBody {
	s.SourcePath = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponseBody) SetTargetPath(v string) *ClonePolarFsBasicSnapshotResponseBody {
	s.TargetPath = &v
	return s
}

func (s *ClonePolarFsBasicSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
