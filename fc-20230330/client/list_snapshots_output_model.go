// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSnapshotsOutput
	GetNextToken() *string
	SetSnapshots(v []*Snapshot) *ListSnapshotsOutput
	GetSnapshots() []*Snapshot
}

type ListSnapshotsOutput struct {
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	Snapshots []*Snapshot `json:"snapshots" xml:"snapshots" type:"Repeated"`
}

func (s ListSnapshotsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsOutput) GoString() string {
	return s.String()
}

func (s *ListSnapshotsOutput) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSnapshotsOutput) GetSnapshots() []*Snapshot {
	return s.Snapshots
}

func (s *ListSnapshotsOutput) SetNextToken(v string) *ListSnapshotsOutput {
	s.NextToken = &v
	return s
}

func (s *ListSnapshotsOutput) SetSnapshots(v []*Snapshot) *ListSnapshotsOutput {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsOutput) Validate() error {
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
