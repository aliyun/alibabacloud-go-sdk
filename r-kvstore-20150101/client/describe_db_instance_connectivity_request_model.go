// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbInstanceConnectivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDbInstanceConnectivityRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeDbInstanceConnectivityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDbInstanceConnectivityRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDbInstanceConnectivityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDbInstanceConnectivityRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeDbInstanceConnectivityRequest
	GetSecurityToken() *string
	SetSrcIp(v string) *DescribeDbInstanceConnectivityRequest
	GetSrcIp() *string
}

type DescribeDbInstanceConnectivityRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1r36hdqlrgt1****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The IP address of the client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 124.207.240.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s DescribeDbInstanceConnectivityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbInstanceConnectivityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbInstanceConnectivityRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDbInstanceConnectivityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDbInstanceConnectivityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDbInstanceConnectivityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDbInstanceConnectivityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDbInstanceConnectivityRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDbInstanceConnectivityRequest) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeDbInstanceConnectivityRequest) SetDBInstanceId(v string) *DescribeDbInstanceConnectivityRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetOwnerAccount(v string) *DescribeDbInstanceConnectivityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetOwnerId(v int64) *DescribeDbInstanceConnectivityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetResourceOwnerAccount(v string) *DescribeDbInstanceConnectivityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetResourceOwnerId(v int64) *DescribeDbInstanceConnectivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetSecurityToken(v string) *DescribeDbInstanceConnectivityRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) SetSrcIp(v string) *DescribeDbInstanceConnectivityRequest {
	s.SrcIp = &v
	return s
}

func (s *DescribeDbInstanceConnectivityRequest) Validate() error {
	return dara.Validate(s)
}
