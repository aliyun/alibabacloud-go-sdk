// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveryAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamId(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetIpamId() *string
	SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetIpamResourceDiscoveryId() *string
	SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest
	GetResourceOwnerId() *int64
}

type ListIpamResourceDiscoveryAssociationsRequest struct {
	// The ID of the IPAM.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The ID of resource discovery instance.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The maximum number of entries on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first or only query, this parameter is left empty.
	//
	// 	- If a next query is to be sent, the returned value is the value of NextToken that was returned last time this operation was called.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
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

func (s ListIpamResourceDiscoveryAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveryAssociationsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetIpamId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.IpamId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetIpamResourceDiscoveryId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetMaxResults(v int32) *ListIpamResourceDiscoveryAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetNextToken(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetRegionId(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveryAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) SetResourceOwnerId(v int64) *ListIpamResourceDiscoveryAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveryAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
