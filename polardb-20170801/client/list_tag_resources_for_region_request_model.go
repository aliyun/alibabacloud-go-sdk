// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesForRegionRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagResourcesForRegionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagResourcesForRegionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTagResourcesForRegionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTagResourcesForRegionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListTagResourcesForRegionRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *ListTagResourcesForRegionRequest
	GetResourceType() *string
}

type ListTagResourcesForRegionRequest struct {
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
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
	// example:
	//
	// aicluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagResourcesForRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForRegionRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForRegionRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesForRegionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagResourcesForRegionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagResourcesForRegionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesForRegionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagResourcesForRegionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListTagResourcesForRegionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesForRegionRequest) SetNextToken(v string) *ListTagResourcesForRegionRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetOwnerAccount(v string) *ListTagResourcesForRegionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetOwnerId(v int64) *ListTagResourcesForRegionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetRegionId(v string) *ListTagResourcesForRegionRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetResourceOwnerAccount(v string) *ListTagResourcesForRegionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetResourceOwnerId(v int64) *ListTagResourcesForRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) SetResourceType(v string) *ListTagResourcesForRegionRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesForRegionRequest) Validate() error {
	return dara.Validate(s)
}
