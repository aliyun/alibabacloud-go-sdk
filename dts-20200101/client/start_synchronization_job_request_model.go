// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *StartSynchronizationJobRequest
	GetAccountId() *string
	SetOwnerId(v string) *StartSynchronizationJobRequest
	GetOwnerId() *string
	SetRegionId(v string) *StartSynchronizationJobRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StartSynchronizationJobRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *StartSynchronizationJobRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *StartSynchronizationJobRequest
	GetSynchronizationJobId() *string
}

type StartSynchronizationJobRequest struct {
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
	// 	- The default value is **Forward**.
	//
	// 	- You can set this parameter to **Reverse*	- to start the reverse synchronization task only if the topology is two-way synchronization.
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
	// dtsf19100l2186****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s StartSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *StartSynchronizationJobRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StartSynchronizationJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartSynchronizationJobRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StartSynchronizationJobRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *StartSynchronizationJobRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *StartSynchronizationJobRequest) SetAccountId(v string) *StartSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetOwnerId(v string) *StartSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetRegionId(v string) *StartSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetResourceGroupId(v string) *StartSynchronizationJobRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetSynchronizationDirection(v string) *StartSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetSynchronizationJobId(v string) *StartSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *StartSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}
