// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkRuleIds(v []*string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetTrafficMarkRuleIds() []*string
	SetTrafficMarkingPolicyId(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
	GetTrafficMarkingPolicyId() *string
}

type RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GoString() string {
	return s.String()
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetTrafficMarkRuleIds() []*string {
	return s.TrafficMarkRuleIds
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetClientToken(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetDryRun(v bool) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerAccount(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetOwnerId(v int64) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerAccount(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetResourceOwnerId(v int64) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkRuleIds(v []*string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkRuleIds = v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) SetTrafficMarkingPolicyId(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) Validate() error {
	return dara.Validate(s)
}
