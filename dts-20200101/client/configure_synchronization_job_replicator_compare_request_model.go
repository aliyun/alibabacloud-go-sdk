// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureSynchronizationJobReplicatorCompareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetAccountId() *string
	SetClientToken(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetClientToken() *string
	SetOwnerId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetOwnerId() *string
	SetRegionId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetResourceGroupId() *string
	SetSynchronizationDirection(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetSynchronizationDirection() *string
	SetSynchronizationJobId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetSynchronizationJobId() *string
	SetSynchronizationReplicatorCompareEnable(v bool) *ConfigureSynchronizationJobReplicatorCompareRequest
	GetSynchronizationReplicatorCompareEnable() *bool
}

type ConfigureSynchronizationJobReplicatorCompareRequest struct {
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
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the data synchronization instance. You can call the [DescribeSynchronizationJobs](https://help.aliyun.com/document_detail/49454.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// Specifies whether to enable image matching. Valid values:
	//
	// 	- **true**: enables image matching
	//
	// 	- **false**: disables image matching
	//
	// example:
	//
	// true
	SynchronizationReplicatorCompareEnable *bool `json:"SynchronizationReplicatorCompareEnable,omitempty" xml:"SynchronizationReplicatorCompareEnable,omitempty"`
}

func (s ConfigureSynchronizationJobReplicatorCompareRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) GetSynchronizationReplicatorCompareEnable() *bool {
	return s.SynchronizationReplicatorCompareEnable
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetAccountId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.AccountId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetClientToken(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetOwnerId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetRegionId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetResourceGroupId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationReplicatorCompareEnable(v bool) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationReplicatorCompareEnable = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) Validate() error {
	return dara.Validate(s)
}
