// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAffinity(v int64) *CreateDeploymentSetRequest
	GetAffinity() *int64
	SetClientToken(v string) *CreateDeploymentSetRequest
	GetClientToken() *string
	SetDeploymentSetName(v string) *CreateDeploymentSetRequest
	GetDeploymentSetName() *string
	SetDescription(v string) *CreateDeploymentSetRequest
	GetDescription() *string
	SetDomain(v string) *CreateDeploymentSetRequest
	GetDomain() *string
	SetGranularity(v string) *CreateDeploymentSetRequest
	GetGranularity() *string
	SetGroupCount(v int64) *CreateDeploymentSetRequest
	GetGroupCount() *int64
	SetOnUnableToRedeployFailedInstance(v string) *CreateDeploymentSetRequest
	GetOnUnableToRedeployFailedInstance() *string
	SetOwnerAccount(v string) *CreateDeploymentSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDeploymentSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDeploymentSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateDeploymentSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDeploymentSetRequest
	GetResourceOwnerId() *int64
	SetStrategy(v string) *CreateDeploymentSetRequest
	GetStrategy() *string
	SetType(v string) *CreateDeploymentSetRequest
	GetType() *string
}

type CreateDeploymentSetRequest struct {
	// The affinity level of the deployment set. This level determines how instances are distributed within the set. The value must be an integer from 1 to 10. Default value: 1.
	//
	// example:
	//
	// 3
	Affinity *int64 `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// A client-generated token that you can use to ensure request idempotence. The token must be unique across requests.
	//
	// The **ClientToken*	- value must be an ASCII string of up to 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the deployment set. The name must be 2 to 128 characters long and start with a letter. It can contain digits, colons (:), underscores (_), and hyphens (-). The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDeploymentSetName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// The description of the deployment set. The description must be 2 to 256 characters long and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of partitions in the deployment set group. Valid values: 1 to 7.
	//
	// Default value: 3.
	//
	// > This parameter is valid only when `Strategy` is set to `AvailabilityGroup`.
	//
	// example:
	//
	// 1
	GroupCount *int64 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The policy for an instance that fails to be redeployed after a failover due to insufficient resources. Valid values:
	//
	// - CancelMembershipAndStart: Removes the instance from the deployment set and starts the instance immediately after failover.
	//
	// - KeepStopped: Keeps the instance in the deployment set and in the Stopped state.
	//
	// Default value: CancelMembershipAndStart.
	//
	// example:
	//
	// CancelMembershipAndStart
	OnUnableToRedeployFailedInstance *string `json:"OnUnableToRedeployFailedInstance,omitempty" xml:"OnUnableToRedeployFailedInstance,omitempty"`
	OwnerAccount                     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region for the deployment set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The deployment strategy. Valid values:
	//
	// - Availability: High availability strategy.
	//
	// - AvailabilityGroup: High availability strategy for deployment set groups.
	//
	// - LowLatency: Low-latency strategy.
	//
	// Default value: Availability.
	//
	// example:
	//
	// Availability
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The deployment granularity. Valid values:
	//
	// - host: Spreads instances across different hosts.
	//
	// - sw: Spreads instances across different switches.
	//
	// - rack: Spreads instances across different racks.
	//
	// Default value: host.
	//
	// example:
	//
	// host
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentSetRequest) GetAffinity() *int64 {
	return s.Affinity
}

func (s *CreateDeploymentSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDeploymentSetRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *CreateDeploymentSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDeploymentSetRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDeploymentSetRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *CreateDeploymentSetRequest) GetGroupCount() *int64 {
	return s.GroupCount
}

func (s *CreateDeploymentSetRequest) GetOnUnableToRedeployFailedInstance() *string {
	return s.OnUnableToRedeployFailedInstance
}

func (s *CreateDeploymentSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDeploymentSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDeploymentSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDeploymentSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDeploymentSetRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateDeploymentSetRequest) GetType() *string {
	return s.Type
}

func (s *CreateDeploymentSetRequest) SetAffinity(v int64) *CreateDeploymentSetRequest {
	s.Affinity = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetClientToken(v string) *CreateDeploymentSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetDeploymentSetName(v string) *CreateDeploymentSetRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetDescription(v string) *CreateDeploymentSetRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetDomain(v string) *CreateDeploymentSetRequest {
	s.Domain = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetGranularity(v string) *CreateDeploymentSetRequest {
	s.Granularity = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetGroupCount(v int64) *CreateDeploymentSetRequest {
	s.GroupCount = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetOnUnableToRedeployFailedInstance(v string) *CreateDeploymentSetRequest {
	s.OnUnableToRedeployFailedInstance = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetOwnerAccount(v string) *CreateDeploymentSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetOwnerId(v int64) *CreateDeploymentSetRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetRegionId(v string) *CreateDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetResourceOwnerAccount(v string) *CreateDeploymentSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetResourceOwnerId(v int64) *CreateDeploymentSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetStrategy(v string) *CreateDeploymentSetRequest {
	s.Strategy = &v
	return s
}

func (s *CreateDeploymentSetRequest) SetType(v string) *CreateDeploymentSetRequest {
	s.Type = &v
	return s
}

func (s *CreateDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
