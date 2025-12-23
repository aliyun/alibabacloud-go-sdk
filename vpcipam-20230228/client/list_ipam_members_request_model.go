// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListIpamMembersRequest
	GetMaxResults() *int32
	SetMemberIds(v []*string) *ListIpamMembersRequest
	GetMemberIds() []*string
	SetNextToken(v string) *ListIpamMembersRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamMembersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamMembersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpamMembersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListIpamMembersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamMembersRequest
	GetResourceOwnerId() *int64
}

type ListIpamMembersRequest struct {
	// example:
	//
	// 20
	MaxResults *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	MemberIds  []*string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty" type:"Repeated"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListIpamMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamMembersRequest) GoString() string {
	return s.String()
}

func (s *ListIpamMembersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamMembersRequest) GetMemberIds() []*string {
	return s.MemberIds
}

func (s *ListIpamMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamMembersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamMembersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamMembersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamMembersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamMembersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamMembersRequest) SetMaxResults(v int32) *ListIpamMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamMembersRequest) SetMemberIds(v []*string) *ListIpamMembersRequest {
	s.MemberIds = v
	return s
}

func (s *ListIpamMembersRequest) SetNextToken(v string) *ListIpamMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamMembersRequest) SetOwnerAccount(v string) *ListIpamMembersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamMembersRequest) SetOwnerId(v int64) *ListIpamMembersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamMembersRequest) SetRegionId(v string) *ListIpamMembersRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamMembersRequest) SetResourceOwnerAccount(v string) *ListIpamMembersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamMembersRequest) SetResourceOwnerId(v int64) *ListIpamMembersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamMembersRequest) Validate() error {
	return dara.Validate(s)
}
