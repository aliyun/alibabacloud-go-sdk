// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaNodeStatistics interface {
	dara.Model
	String() string
	GoString() string
	SetActualMinHyperNodeNum(v int64) *QuotaNodeStatistics
	GetActualMinHyperNodeNum() *int64
	SetActualMinNodeNum(v int64) *QuotaNodeStatistics
	GetActualMinNodeNum() *int64
	SetAllocatedHyperNodeDetails(v []*AllocatedHyperNodeDetail) *QuotaNodeStatistics
	GetAllocatedHyperNodeDetails() []*AllocatedHyperNodeDetail
	SetAllocatedHyperNodeNum(v int64) *QuotaNodeStatistics
	GetAllocatedHyperNodeNum() *int64
	SetAllocatedNodeNum(v int64) *QuotaNodeStatistics
	GetAllocatedNodeNum() *int64
	SetEmptyNodeNum(v int64) *QuotaNodeStatistics
	GetEmptyNodeNum() *int64
}

type QuotaNodeStatistics struct {
	// The guaranteed minimum number of hyper nodes available in the quota.
	ActualMinHyperNodeNum *int64 `json:"ActualMinHyperNodeNum,omitempty" xml:"ActualMinHyperNodeNum,omitempty"`
	// The guaranteed minimum number of nodes available in the quota.
	ActualMinNodeNum *int64 `json:"ActualMinNodeNum,omitempty" xml:"ActualMinNodeNum,omitempty"`
	// The details of an allocated hyper node.
	AllocatedHyperNodeDetails []*AllocatedHyperNodeDetail `json:"AllocatedHyperNodeDetails,omitempty" xml:"AllocatedHyperNodeDetails,omitempty" type:"Repeated"`
	// The number of hyper nodes currently allocated from the quota.
	AllocatedHyperNodeNum *int64 `json:"AllocatedHyperNodeNum,omitempty" xml:"AllocatedHyperNodeNum,omitempty"`
	// The number of nodes currently allocated from the quota.
	AllocatedNodeNum *int64 `json:"AllocatedNodeNum,omitempty" xml:"AllocatedNodeNum,omitempty"`
	// The number of allocated nodes currently idle.
	EmptyNodeNum *int64 `json:"EmptyNodeNum,omitempty" xml:"EmptyNodeNum,omitempty"`
}

func (s QuotaNodeStatistics) String() string {
	return dara.Prettify(s)
}

func (s QuotaNodeStatistics) GoString() string {
	return s.String()
}

func (s *QuotaNodeStatistics) GetActualMinHyperNodeNum() *int64 {
	return s.ActualMinHyperNodeNum
}

func (s *QuotaNodeStatistics) GetActualMinNodeNum() *int64 {
	return s.ActualMinNodeNum
}

func (s *QuotaNodeStatistics) GetAllocatedHyperNodeDetails() []*AllocatedHyperNodeDetail {
	return s.AllocatedHyperNodeDetails
}

func (s *QuotaNodeStatistics) GetAllocatedHyperNodeNum() *int64 {
	return s.AllocatedHyperNodeNum
}

func (s *QuotaNodeStatistics) GetAllocatedNodeNum() *int64 {
	return s.AllocatedNodeNum
}

func (s *QuotaNodeStatistics) GetEmptyNodeNum() *int64 {
	return s.EmptyNodeNum
}

func (s *QuotaNodeStatistics) SetActualMinHyperNodeNum(v int64) *QuotaNodeStatistics {
	s.ActualMinHyperNodeNum = &v
	return s
}

func (s *QuotaNodeStatistics) SetActualMinNodeNum(v int64) *QuotaNodeStatistics {
	s.ActualMinNodeNum = &v
	return s
}

func (s *QuotaNodeStatistics) SetAllocatedHyperNodeDetails(v []*AllocatedHyperNodeDetail) *QuotaNodeStatistics {
	s.AllocatedHyperNodeDetails = v
	return s
}

func (s *QuotaNodeStatistics) SetAllocatedHyperNodeNum(v int64) *QuotaNodeStatistics {
	s.AllocatedHyperNodeNum = &v
	return s
}

func (s *QuotaNodeStatistics) SetAllocatedNodeNum(v int64) *QuotaNodeStatistics {
	s.AllocatedNodeNum = &v
	return s
}

func (s *QuotaNodeStatistics) SetEmptyNodeNum(v int64) *QuotaNodeStatistics {
	s.EmptyNodeNum = &v
	return s
}

func (s *QuotaNodeStatistics) Validate() error {
	if s.AllocatedHyperNodeDetails != nil {
		for _, item := range s.AllocatedHyperNodeDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
