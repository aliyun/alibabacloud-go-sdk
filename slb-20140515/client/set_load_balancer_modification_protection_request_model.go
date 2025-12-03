// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerModificationProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerModificationProtectionRequest
	GetLoadBalancerId() *string
	SetModificationProtectionReason(v string) *SetLoadBalancerModificationProtectionRequest
	GetModificationProtectionReason() *string
	SetModificationProtectionStatus(v string) *SetLoadBalancerModificationProtectionRequest
	GetModificationProtectionStatus() *string
	SetOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerModificationProtectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest
	GetResourceOwnerId() *int64
}

type SetLoadBalancerModificationProtectionRequest struct {
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08e*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The reason why the configuration read-only mode is enabled. The value must be 1 to 80 characters in length. It must start with a letter and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  This parameter is valid only if the **ModificationProtectionStatus*	- parameter is set to **ConsoleProtection**.
	//
	// example:
	//
	// Configuration change
	ModificationProtectionReason *string `json:"ModificationProtectionReason,omitempty" xml:"ModificationProtectionReason,omitempty"`
	// Specifies whether to enable the configuration read-only mode. Valid values:
	//
	// 	- **NonProtection**: disables the configuration read-only mode. After you disable the configuration read-only mode, the value of **ModificationProtectionReason*	- is cleared.
	//
	// 	- **ConsoleProtection**: enables the configuration read-only mode.
	//
	// >  If you set this parameter to **ConsoleProtection**, you cannot use the CLB console to modify instance configurations. However, you can call API operations to modify instance configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// ConsoleProtection
	ModificationProtectionStatus *string `json:"ModificationProtectionStatus,omitempty" xml:"ModificationProtectionStatus,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerModificationProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerModificationProtectionRequest) GetModificationProtectionReason() *string {
	return s.ModificationProtectionReason
}

func (s *SetLoadBalancerModificationProtectionRequest) GetModificationProtectionStatus() *string {
	return s.ModificationProtectionStatus
}

func (s *SetLoadBalancerModificationProtectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerModificationProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerModificationProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerModificationProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerModificationProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerModificationProtectionRequest) SetLoadBalancerId(v string) *SetLoadBalancerModificationProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetModificationProtectionReason(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ModificationProtectionReason = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetModificationProtectionStatus(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ModificationProtectionStatus = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetRegionId(v string) *SetLoadBalancerModificationProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerModificationProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) SetResourceOwnerId(v int64) *SetLoadBalancerModificationProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionRequest) Validate() error {
	return dara.Validate(s)
}
