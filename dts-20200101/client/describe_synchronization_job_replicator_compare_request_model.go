// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobReplicatorCompareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *DescribeSynchronizationJobReplicatorCompareRequest
	GetSynchronizationJobId() *string
}

type DescribeSynchronizationJobReplicatorCompareRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the data synchronization instance resides. For more information, see [List of supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// 	- This parameter is required only when the topology of the data synchronization instance is two-way synchronization.
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	// The ID of the data synchronization instance. You can call the [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetAccountId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetClientToken(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetOwnerId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetRegionId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetResourceGroupId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) Validate() error {
	return dara.Validate(s)
}
