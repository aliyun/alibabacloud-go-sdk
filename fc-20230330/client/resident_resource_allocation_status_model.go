// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResidentResourceAllocationStatus interface {
	dara.Model
	String() string
	GoString() string
	SetLastAllocatedTime(v string) *ResidentResourceAllocationStatus
	GetLastAllocatedTime() *string
	SetLastAllocation(v []*ResidentResourceAllocation) *ResidentResourceAllocationStatus
	GetLastAllocation() []*ResidentResourceAllocation
}

type ResidentResourceAllocationStatus struct {
	LastAllocatedTime *string                       `json:"lastAllocatedTime,omitempty" xml:"lastAllocatedTime,omitempty"`
	LastAllocation    []*ResidentResourceAllocation `json:"lastAllocation" xml:"lastAllocation" type:"Repeated"`
}

func (s ResidentResourceAllocationStatus) String() string {
	return dara.Prettify(s)
}

func (s ResidentResourceAllocationStatus) GoString() string {
	return s.String()
}

func (s *ResidentResourceAllocationStatus) GetLastAllocatedTime() *string {
	return s.LastAllocatedTime
}

func (s *ResidentResourceAllocationStatus) GetLastAllocation() []*ResidentResourceAllocation {
	return s.LastAllocation
}

func (s *ResidentResourceAllocationStatus) SetLastAllocatedTime(v string) *ResidentResourceAllocationStatus {
	s.LastAllocatedTime = &v
	return s
}

func (s *ResidentResourceAllocationStatus) SetLastAllocation(v []*ResidentResourceAllocation) *ResidentResourceAllocationStatus {
	s.LastAllocation = v
	return s
}

func (s *ResidentResourceAllocationStatus) Validate() error {
	if s.LastAllocation != nil {
		for _, item := range s.LastAllocation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
