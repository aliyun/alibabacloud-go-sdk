// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest
	GetResourceOwnerId() *int64
	SetTrafficQosPolicyDescription(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetTrafficQosPolicyDescription() *string
	SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetTrafficQosPolicyId() *string
	SetTrafficQosPolicyName(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetTrafficQosPolicyName() *string
	SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetTransitRouterAttachmentId() *string
	SetTransitRouterId(v string) *ListCenInterRegionTrafficQosPoliciesRequest
	GetTransitRouterId() *string
}

type ListCenInterRegionTrafficQosPoliciesRequest struct {
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
	// The description of the QoS policy.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// desctest
	TrafficQosPolicyDescription *string `json:"TrafficQosPolicyDescription,omitempty" xml:"TrafficQosPolicyDescription,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-rnghap5gc8155x****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The name of the QoS policy.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// nametest
	TrafficQosPolicyName *string `json:"TrafficQosPolicyName,omitempty" xml:"TrafficQosPolicyName,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-r6g0m3epjehw57****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-bp1rmwxnk221e3fas****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetTrafficQosPolicyDescription() *string {
	return s.TrafficQosPolicyDescription
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetTrafficQosPolicyName() *string {
	return s.TrafficQosPolicyName
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetMaxResults(v int32) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetNextToken(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetResourceOwnerAccount(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetResourceOwnerId(v int64) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyDescription(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTrafficQosPolicyName(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TrafficQosPolicyName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosPoliciesRequest {
	s.TransitRouterId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
