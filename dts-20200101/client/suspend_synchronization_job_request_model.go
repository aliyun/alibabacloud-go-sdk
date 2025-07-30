// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *SuspendSynchronizationJobRequest
	GetAccountId() *string
	SetOwnerId(v string) *SuspendSynchronizationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *SuspendSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *SuspendSynchronizationJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *SuspendSynchronizationJobRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *SuspendSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type SuspendSynchronizationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource GroupId
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// >
	//
	// 	- Default value: **Forward**.
	//
	// 	- You can set this parameter to **Reverse*	- to pause reverse synchronization only when the topology is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the **DescribeSynchronizationJobs*	- operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsmr1q4mc2152****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s SuspendSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *SuspendSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SuspendSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SuspendSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SuspendSynchronizationJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *SuspendSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *SuspendSynchronizationJobRequest) SetAccountId(v string) *SuspendSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetOwnerId(v string) *SuspendSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetRegionId(v string) *SuspendSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetResourceGroupId(v string) *SuspendSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetSynchronizationDirection(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetSynchronizationJobId(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
