// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransitRouterMulticastDomainVSwitchesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetCenId() *string
	SetMaxResults(v int32) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetResourceOwnerId() *int64
	SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *ListTransitRouterMulticastDomainVSwitchesRequest
	GetVpcId() *string
}

type ListTransitRouterMulticastDomainVSwitchesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cen-44m0p68spvlrqq****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchIds           []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-p0wr1cd4fcuxy3ui0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListTransitRouterMulticastDomainVSwitchesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransitRouterMulticastDomainVSwitchesRequest) GoString() string {
	return s.String()
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetCenId() *string {
	return s.CenId
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetCenId(v string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.CenId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetMaxResults(v int32) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetNextToken(v string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetOwnerAccount(v string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetOwnerId(v int64) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetResourceOwnerAccount(v string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetResourceOwnerId(v int64) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetVSwitchIds(v []*string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.VSwitchIds = v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) SetVpcId(v string) *ListTransitRouterMulticastDomainVSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *ListTransitRouterMulticastDomainVSwitchesRequest) Validate() error {
	return dara.Validate(s)
}
