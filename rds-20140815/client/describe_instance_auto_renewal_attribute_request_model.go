// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewalAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest
	GetResourceOwnerId() *int64
	SetProxyId(v string) *DescribeInstanceAutoRenewalAttributeRequest
	GetProxyId() *string
}

type DescribeInstanceAutoRenewalAttributeRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProxyId              *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeInstanceAutoRenewalAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewalAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetClientToken(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetDBInstanceId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageNumber(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetPageSize(v int32) *DescribeInstanceAutoRenewalAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetRegionId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerAccount(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetResourceOwnerId(v int64) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) SetProxyId(v string) *DescribeInstanceAutoRenewalAttributeRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeInstanceAutoRenewalAttributeRequest) Validate() error {
	return dara.Validate(s)
}
