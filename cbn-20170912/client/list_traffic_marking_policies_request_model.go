// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficMarkingPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTrafficMarkingPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrafficMarkingPoliciesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTrafficMarkingPoliciesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTrafficMarkingPoliciesRequest
	GetResourceOwnerId() *int64
	SetTrafficMarkingPolicyDescription(v string) *ListTrafficMarkingPoliciesRequest
	GetTrafficMarkingPolicyDescription() *string
	SetTrafficMarkingPolicyId(v string) *ListTrafficMarkingPoliciesRequest
	GetTrafficMarkingPolicyId() *string
	SetTrafficMarkingPolicyName(v string) *ListTrafficMarkingPoliciesRequest
	GetTrafficMarkingPolicyName() *string
	SetTransitRouterId(v string) *ListTrafficMarkingPoliciesRequest
	GetTransitRouterId() *string
}

type ListTrafficMarkingPoliciesRequest struct {
	// The number of entries to return on each page. Valid values: **1*	- to **100**. Default value: **20**.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. Valid values:
	//
	// 	- If this is your first query or no subsequent query is to be sent, ignore this parameter.
	//
	// 	- If a next query is to be sent, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the traffic marking policy.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	TrafficMarkingPolicyDescription *string `json:"TrafficMarkingPolicyDescription,omitempty" xml:"TrafficMarkingPolicyDescription,omitempty"`
	// The ID of the traffic marking policy.
	//
	// example:
	//
	// tm-iz5egnyitxiroq****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
	// The name of the traffic marking policy.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TrafficMarkingPolicyName *string `json:"TrafficMarkingPolicyName,omitempty" xml:"TrafficMarkingPolicyName,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1rmwxnk221e3fas****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListTrafficMarkingPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficMarkingPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListTrafficMarkingPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrafficMarkingPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrafficMarkingPoliciesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTrafficMarkingPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTrafficMarkingPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTrafficMarkingPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTrafficMarkingPoliciesRequest) GetTrafficMarkingPolicyDescription() *string {
	return s.TrafficMarkingPolicyDescription
}

func (s *ListTrafficMarkingPoliciesRequest) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *ListTrafficMarkingPoliciesRequest) GetTrafficMarkingPolicyName() *string {
	return s.TrafficMarkingPolicyName
}

func (s *ListTrafficMarkingPoliciesRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListTrafficMarkingPoliciesRequest) SetMaxResults(v int32) *ListTrafficMarkingPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetNextToken(v string) *ListTrafficMarkingPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetOwnerId(v int64) *ListTrafficMarkingPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetResourceOwnerAccount(v string) *ListTrafficMarkingPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetResourceOwnerId(v int64) *ListTrafficMarkingPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyDescription(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyDescription = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyId(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTrafficMarkingPolicyName(v string) *ListTrafficMarkingPoliciesRequest {
	s.TrafficMarkingPolicyName = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) SetTransitRouterId(v string) *ListTrafficMarkingPoliciesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListTrafficMarkingPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
