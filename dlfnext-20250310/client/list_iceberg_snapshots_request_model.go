// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIcebergSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIcebergSnapshotsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListIcebergSnapshotsRequest
	GetPageToken() *string
}

type ListIcebergSnapshotsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListIcebergSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIcebergSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListIcebergSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIcebergSnapshotsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListIcebergSnapshotsRequest) SetMaxResults(v int32) *ListIcebergSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIcebergSnapshotsRequest) SetPageToken(v string) *ListIcebergSnapshotsRequest {
	s.PageToken = &v
	return s
}

func (s *ListIcebergSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
