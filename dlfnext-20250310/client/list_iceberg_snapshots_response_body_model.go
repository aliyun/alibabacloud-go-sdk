// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListIcebergSnapshotsResponseBody
	GetNextPageToken() *string
	SetSnapshots(v []*IcebergSnapshot) *ListIcebergSnapshotsResponseBody
	GetSnapshots() []*IcebergSnapshot
}

type ListIcebergSnapshotsResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string            `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Snapshots     []*IcebergSnapshot `json:"snapshots,omitempty" xml:"snapshots,omitempty" type:"Repeated"`
}

func (s ListIcebergSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIcebergSnapshotsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListIcebergSnapshotsResponseBody) GetSnapshots() []*IcebergSnapshot {
	return s.Snapshots
}

func (s *ListIcebergSnapshotsResponseBody) SetNextPageToken(v string) *ListIcebergSnapshotsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListIcebergSnapshotsResponseBody) SetSnapshots(v []*IcebergSnapshot) *ListIcebergSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListIcebergSnapshotsResponseBody) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
