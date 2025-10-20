// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListPartitionsResponseBody
	GetNextPageToken() *string
	SetPartitions(v []*Partition) *ListPartitionsResponseBody
	GetPartitions() []*Partition
}

type ListPartitionsResponseBody struct {
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string      `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Partitions    []*Partition `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s ListPartitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListPartitionsResponseBody) GetPartitions() []*Partition {
	return s.Partitions
}

func (s *ListPartitionsResponseBody) SetNextPageToken(v string) *ListPartitionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionsResponseBody) SetPartitions(v []*Partition) *ListPartitionsResponseBody {
	s.Partitions = v
	return s
}

func (s *ListPartitionsResponseBody) Validate() error {
	if s.Partitions != nil {
		for _, item := range s.Partitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
