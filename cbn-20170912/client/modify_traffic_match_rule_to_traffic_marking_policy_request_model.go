// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficMatchRuleToTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
	SetTrafficMatchRuleDescription(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMatchRuleDescription() *string
	SetTrafficMatchRuleId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMatchRuleId() *string
	SetTrafficMatchRuleName(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
	GetTrafficMatchRuleName() *string
}

type ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and sends the request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic marking policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-u9nxup5kww5po8****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	// The description of the traffic classification rule.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// descriptiontest
	TrafficMatchRuleDescription *string `json:"TrafficMatchRuleDescription,omitempty" xml:"TrafficMatchRuleDescription,omitempty"`
	// The ID of the traffic classification rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// tm-rule-fa9kgq1e90rmhc****
	TrafficMatchRuleId *string `json:"TrafficMatchRuleId,omitempty" xml:"TrafficMatchRuleId,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficMatchRuleName *string `json:"TrafficMatchRuleName,omitempty" xml:"TrafficMatchRuleName,omitempty"`
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMatchRuleDescription() *string {
	return s.TrafficMatchRuleDescription
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMatchRuleId() *string {
	return s.TrafficMatchRuleId
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) GetTrafficMatchRuleName() *string {
	return s.TrafficMatchRuleName
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetClientToken(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetDryRun(v bool) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetOwnerId(v int64) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRuleDescription(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRuleDescription = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRuleId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRuleId = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) SetTrafficMatchRuleName(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest {
	s.TrafficMatchRuleName = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
