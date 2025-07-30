// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ResetSynchronizationJobRequest
	GetAccountId() *string
	SetOwnerId(v string) *ResetSynchronizationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *ResetSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ResetSynchronizationJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ResetSynchronizationJobRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *ResetSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type ResetSynchronizationJobRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
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
	// 	- You can set this parameter to **Reverse*	- to stop reverse synchronization only when the topology is two-way synchronization.
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
	// dtsm761239l27m****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s ResetSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ResetSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ResetSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ResetSynchronizationJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ResetSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ResetSynchronizationJobRequest) SetAccountId(v string) *ResetSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetOwnerId(v string) *ResetSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetRegionId(v string) *ResetSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetResourceGroupId(v string) *ResetSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationDirection(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationJobId(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
