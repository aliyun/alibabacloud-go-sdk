// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeSynchronizationJobStatusRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeSynchronizationJobStatusRequest
	GetClientToken() *string
	SetOwnerId(v string) *DescribeSynchronizationJobStatusRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeSynchronizationJobStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSynchronizationJobStatusRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusRequest
	GetSynchronizationJobId() *string
}

type DescribeSynchronizationJobStatusRequest struct {
	// The ID of the Alibaba Cloud account. You do not need to specify this parameter because this parameter will be removed in the future.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The resource group ID.
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
	// >  Default value: **Forward**.
	//
	// The value **Reverse*	- takes effect only if the topology of the data synchronization instance is two-way synchronization.
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

func (s DescribeSynchronizationJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeSynchronizationJobStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeSynchronizationJobStatusRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeSynchronizationJobStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSynchronizationJobStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSynchronizationJobStatusRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobStatusRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobStatusRequest) SetAccountId(v string) *DescribeSynchronizationJobStatusRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetClientToken(v string) *DescribeSynchronizationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetOwnerId(v string) *DescribeSynchronizationJobStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetRegionId(v string) *DescribeSynchronizationJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetResourceGroupId(v string) *DescribeSynchronizationJobStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
