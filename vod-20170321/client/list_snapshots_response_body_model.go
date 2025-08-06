// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaSnapshot(v *ListSnapshotsResponseBodyMediaSnapshot) *ListSnapshotsResponseBody
	GetMediaSnapshot() *ListSnapshotsResponseBodyMediaSnapshot
	SetRequestId(v string) *ListSnapshotsResponseBody
	GetRequestId() *string
}

type ListSnapshotsResponseBody struct {
	// The information about the snapshot.
	MediaSnapshot *ListSnapshotsResponseBodyMediaSnapshot `json:"MediaSnapshot,omitempty" xml:"MediaSnapshot,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetMediaSnapshot() *ListSnapshotsResponseBodyMediaSnapshot {
	return s.MediaSnapshot
}

func (s *ListSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotsResponseBody) SetMediaSnapshot(v *ListSnapshotsResponseBodyMediaSnapshot) *ListSnapshotsResponseBody {
	s.MediaSnapshot = v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSnapshotsResponseBodyMediaSnapshot struct {
	// The time when the snapshot job was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-20T12:23:45Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the snapshot job.
	//
	// example:
	//
	// ad90a501b1b9472374ad005046****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The rule used to generate snapshot URLs.
	//
	// example:
	//
	// http://example.aliyundoc.com/snapshot/sample{SnapshotCount}.jpg
	Regular *string `json:"Regular,omitempty" xml:"Regular,omitempty"`
	// The details of the snapshot.
	Snapshots *ListSnapshotsResponseBodyMediaSnapshotSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The total number of snapshots.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSnapshotsResponseBodyMediaSnapshot) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshot) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) GetJobId() *string {
	return s.JobId
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) GetRegular() *string {
	return s.Regular
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) GetSnapshots() *ListSnapshotsResponseBodyMediaSnapshotSnapshots {
	return s.Snapshots
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) GetTotal() *int64 {
	return s.Total
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetCreationTime(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetJobId(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.JobId = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetRegular(v string) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Regular = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetSnapshots(v *ListSnapshotsResponseBodyMediaSnapshotSnapshots) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) SetTotal(v int64) *ListSnapshotsResponseBodyMediaSnapshot {
	s.Total = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshot) Validate() error {
	return dara.Validate(s)
}

type ListSnapshotsResponseBodyMediaSnapshotSnapshots struct {
	Snapshot []*ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshots) GetSnapshot() []*ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot {
	return s.Snapshot
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshots) SetSnapshot(v []*ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) *ListSnapshotsResponseBodyMediaSnapshotSnapshots {
	s.Snapshot = v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshots) Validate() error {
	return dara.Validate(s)
}

type ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot struct {
	// The index of the snapshot.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The URL of the snapshot.
	//
	// example:
	//
	// http://example.aliyundoc.com/snapshot/sample00001****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) GetIndex() *int64 {
	return s.Index
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) GetUrl() *string {
	return s.Url
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) SetIndex(v int64) *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot {
	s.Index = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) SetUrl(v string) *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot {
	s.Url = &v
	return s
}

func (s *ListSnapshotsResponseBodyMediaSnapshotSnapshotsSnapshot) Validate() error {
	return dara.Validate(s)
}
