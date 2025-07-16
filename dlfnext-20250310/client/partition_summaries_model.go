// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPartitionSummaries interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *PartitionSummaries
	GetNextPageToken() *string
	SetPartitions(v []*PartitionSummary) *PartitionSummaries
	GetPartitions() []*PartitionSummary
}

type PartitionSummaries struct {
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	// Current page of partition profiles
	Partitions []*PartitionSummary `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s PartitionSummaries) String() string {
	return dara.Prettify(s)
}

func (s PartitionSummaries) GoString() string {
	return s.String()
}

func (s *PartitionSummaries) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *PartitionSummaries) GetPartitions() []*PartitionSummary {
	return s.Partitions
}

func (s *PartitionSummaries) SetNextPageToken(v string) *PartitionSummaries {
	s.NextPageToken = &v
	return s
}

func (s *PartitionSummaries) SetPartitions(v []*PartitionSummary) *PartitionSummaries {
	s.Partitions = v
	return s
}

func (s *PartitionSummaries) Validate() error {
	return dara.Validate(s)
}
