// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *ListSnapshotsRequest
	GetNamespace() *string
	SetObjectId(v string) *ListSnapshotsRequest
	GetObjectId() *string
	SetPageNumber(v int32) *ListSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSnapshotsRequest
	GetPageSize() *int32
	SetType(v string) *ListSnapshotsRequest
	GetType() *string
}

type ListSnapshotsRequest struct {
	// The namespace (project space projectId or personal space baseId).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1389623
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The unique ID of the object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8467231038932407294
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The snapshot type. Multiple values are supported. Valid values: Saved, Deployed, and UnDeployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Saved
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSnapshotsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *ListSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotsRequest) GetType() *string {
	return s.Type
}

func (s *ListSnapshotsRequest) SetNamespace(v string) *ListSnapshotsRequest {
	s.Namespace = &v
	return s
}

func (s *ListSnapshotsRequest) SetObjectId(v string) *ListSnapshotsRequest {
	s.ObjectId = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetType(v string) *ListSnapshotsRequest {
	s.Type = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
