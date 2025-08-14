// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTraficMatchRuleFromTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkRuleIds(v []*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetTrafficMarkRuleIds() []*string
	SetTrafficMarkingPolicyId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
}

type RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic classification rule.
	TrafficMarkRuleIds []*string `json:"TrafficMarkRuleIds,omitempty" xml:"TrafficMarkRuleIds,omitempty" type:"Repeated"`
	// The ID of the traffic marking policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-d33hdczo3qo8ta****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetTrafficMarkRuleIds() []*string {
	return s.TrafficMarkRuleIds
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetClientToken(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetDryRun(v bool) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkRuleIds(v []*string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkRuleIds = v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
