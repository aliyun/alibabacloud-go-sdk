// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListPartitionsRequest
	GetMaxResults() *int64
	SetPageToken(v string) *ListPartitionsRequest
	GetPageToken() *string
	SetPartitionNamePattern(v string) *ListPartitionsRequest
	GetPartitionNamePattern() *string
}

type ListPartitionsRequest struct {
	// example:
	//
	// 1000
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// partition
	PartitionNamePattern *string `json:"partitionNamePattern,omitempty" xml:"partitionNamePattern,omitempty"`
}

func (s ListPartitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListPartitionsRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListPartitionsRequest) GetPartitionNamePattern() *string {
	return s.PartitionNamePattern
}

func (s *ListPartitionsRequest) SetMaxResults(v int64) *ListPartitionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPartitionsRequest) SetPageToken(v string) *ListPartitionsRequest {
	s.PageToken = &v
	return s
}

func (s *ListPartitionsRequest) SetPartitionNamePattern(v string) *ListPartitionsRequest {
	s.PartitionNamePattern = &v
	return s
}

func (s *ListPartitionsRequest) Validate() error {
	return dara.Validate(s)
}
