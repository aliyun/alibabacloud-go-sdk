// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySynchronizationObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ModifySynchronizationObjectRequest
	GetAccountId() *string
	SetOwnerId(v string) *ModifySynchronizationObjectRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifySynchronizationObjectRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifySynchronizationObjectRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ModifySynchronizationObjectRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *ModifySynchronizationObjectRequest
	GetSynchronizationJobId() *string
	SetSynchronizationObjects(v string) *ModifySynchronizationObjectRequest
	GetSynchronizationObjects() *string
}

type ModifySynchronizationObjectRequest struct {
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
	// 	- This parameter is required only when the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the DescribeSynchronizationJobs operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtskfq1149w254****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// This parameter is required.
	SynchronizationObjects *string `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
}

func (s ModifySynchronizationObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySynchronizationObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ModifySynchronizationObjectRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifySynchronizationObjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySynchronizationObjectRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySynchronizationObjectRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ModifySynchronizationObjectRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ModifySynchronizationObjectRequest) GetSynchronizationObjects() *string {
	return s.SynchronizationObjects
}

func (s *ModifySynchronizationObjectRequest) SetAccountId(v string) *ModifySynchronizationObjectRequest {
	s.AccountId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetOwnerId(v string) *ModifySynchronizationObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetRegionId(v string) *ModifySynchronizationObjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetResourceGroupId(v string) *ModifySynchronizationObjectRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetSynchronizationDirection(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetSynchronizationJobId(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetSynchronizationObjects(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationObjects = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) Validate() error {
	return dara.Validate(s)
}
