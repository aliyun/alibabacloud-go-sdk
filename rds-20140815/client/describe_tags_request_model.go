// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeTagsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeTagsRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagsRequest
	GetResourceType() *string
	SetTags(v string) *DescribeTagsRequest
	GetTags() *string
	SetProxyId(v string) *DescribeTagsRequest
	GetProxyId() *string
}

type DescribeTagsRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	ProxyId              *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeTagsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagsRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeTagsRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeTagsRequest) SetClientToken(v string) *DescribeTagsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeTagsRequest) SetDBInstanceId(v string) *DescribeTagsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsRequest) SetTags(v string) *DescribeTagsRequest {
	s.Tags = &v
	return s
}

func (s *DescribeTagsRequest) SetProxyId(v string) *DescribeTagsRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
	return dara.Validate(s)
}
