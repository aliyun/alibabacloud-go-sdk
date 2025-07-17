// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ListTagValuesRequest
	GetKey() *string
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListTagValuesRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListTagValuesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTagValuesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListTagValuesRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
}

type ListTagValuesRequest struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// ESS
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The token that determines the start point of the next query. If this parameter is empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Auto Scaling resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The type of the Auto Scaling resource. Set the value to scalinggroup. This indicates that the tag is added to a scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// scalinggroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagValuesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagValuesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int32) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
