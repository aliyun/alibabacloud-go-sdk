// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListSnapshotsResponseBody
	GetNextPageToken() *string
	SetSnapshots(v []*Snapshot) *ListSnapshotsResponseBody
	GetSnapshots() []*Snapshot
}

type ListSnapshotsResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string     `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Snapshots     []*Snapshot `json:"snapshots,omitempty" xml:"snapshots,omitempty" type:"Repeated"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListSnapshotsResponseBody) GetSnapshots() []*Snapshot {
	return s.Snapshots
}

func (s *ListSnapshotsResponseBody) SetNextPageToken(v string) *ListSnapshotsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetSnapshots(v []*Snapshot) *ListSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}
