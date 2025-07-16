// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPartitionSummariesRequest
	GetMaxResults() *int32
	SetPageToken(v string) *ListPartitionSummariesRequest
	GetPageToken() *string
	SetPartitionNamePattern(v string) *ListPartitionSummariesRequest
	GetPartitionNamePattern() *string
}

type ListPartitionSummariesRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	PageToken *string `json:"pageToken,omitempty" xml:"pageToken,omitempty"`
	// example:
	//
	// hh=10
	PartitionNamePattern *string `json:"partitionNamePattern,omitempty" xml:"partitionNamePattern,omitempty"`
}

func (s ListPartitionSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionSummariesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPartitionSummariesRequest) GetPageToken() *string {
	return s.PageToken
}

func (s *ListPartitionSummariesRequest) GetPartitionNamePattern() *string {
	return s.PartitionNamePattern
}

func (s *ListPartitionSummariesRequest) SetMaxResults(v int32) *ListPartitionSummariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPartitionSummariesRequest) SetPageToken(v string) *ListPartitionSummariesRequest {
	s.PageToken = &v
	return s
}

func (s *ListPartitionSummariesRequest) SetPartitionNamePattern(v string) *ListPartitionSummariesRequest {
	s.PartitionNamePattern = &v
	return s
}

func (s *ListPartitionSummariesRequest) Validate() error {
	return dara.Validate(s)
}
