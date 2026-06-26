// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPartitionsByNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPartitions(v []*Partition) *ListPartitionsByNamesResponseBody
	GetPartitions() []*Partition
}

type ListPartitionsByNamesResponseBody struct {
	// 分区。
	Partitions []*Partition `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
}

func (s ListPartitionsByNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPartitionsByNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionsByNamesResponseBody) GetPartitions() []*Partition {
	return s.Partitions
}

func (s *ListPartitionsByNamesResponseBody) SetPartitions(v []*Partition) *ListPartitionsByNamesResponseBody {
	s.Partitions = v
	return s
}

func (s *ListPartitionsByNamesResponseBody) Validate() error {
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
