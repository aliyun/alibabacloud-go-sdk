// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceByTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstanceByTagsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeDBInstanceByTagsRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDBInstanceByTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstanceByTagsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstanceByTagsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceByTagsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDBInstanceByTagsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstanceByTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceByTagsRequest
	GetResourceOwnerId() *int64
	SetProxyId(v string) *DescribeDBInstanceByTagsRequest
	GetProxyId() *string
}

type DescribeDBInstanceByTagsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// example:
	//
	// rm-uf6w**********
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **30 to 100**. Default value: **30**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// A deprecated parameter.
	//
	// example:
	//
	// None
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstanceByTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstanceByTagsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceByTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstanceByTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstanceByTagsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceByTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceByTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceByTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstanceByTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceByTagsRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstanceByTagsRequest) SetClientToken(v string) *DescribeDBInstanceByTagsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetDBInstanceId(v string) *DescribeDBInstanceByTagsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetOwnerAccount(v string) *DescribeDBInstanceByTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetOwnerId(v int64) *DescribeDBInstanceByTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetPageNumber(v int32) *DescribeDBInstanceByTagsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetPageSize(v int32) *DescribeDBInstanceByTagsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetRegionId(v string) *DescribeDBInstanceByTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetResourceOwnerAccount(v string) *DescribeDBInstanceByTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceByTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) SetProxyId(v string) *DescribeDBInstanceByTagsRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstanceByTagsRequest) Validate() error {
	return dara.Validate(s)
}
