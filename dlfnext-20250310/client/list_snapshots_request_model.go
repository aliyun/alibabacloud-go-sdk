// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSnapshotsRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListSnapshotsRequest
	GetPageToken() *string
}

type ListSnapshotsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSnapshotsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListSnapshotsRequest) SetMaxResults(v int32) *ListSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageToken(v string) *ListSnapshotsRequest {
	s.PageToken = &v
	return s
}

func (s *ListSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
