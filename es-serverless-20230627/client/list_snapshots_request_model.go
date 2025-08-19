// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSnapshotsRequest
	GetPageSize() *int32
	SetRepository(v string) *ListSnapshotsRequest
	GetRepository() *string
	SetSnapshot(v string) *ListSnapshotsRequest
	GetSnapshot() *string
}

type ListSnapshotsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// aliyun_auto_snapshot
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// example:
	//
	// qingning
	Snapshot *string `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotsRequest) GetRepository() *string {
	return s.Repository
}

func (s *ListSnapshotsRequest) GetSnapshot() *string {
	return s.Snapshot
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetRepository(v string) *ListSnapshotsRequest {
	s.Repository = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshot(v string) *ListSnapshotsRequest {
	s.Snapshot = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
