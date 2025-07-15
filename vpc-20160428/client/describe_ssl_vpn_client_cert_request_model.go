// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSslVpnClientCertRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSslVpnClientCertRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSslVpnClientCertRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSslVpnClientCertRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSslVpnClientCertRequest
	GetResourceOwnerId() *int64
	SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertRequest
	GetSslVpnClientCertId() *string
}

type DescribeSslVpnClientCertRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SSL client certificate. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL client certificate that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsc-bp17r58rjf5r1gjyr****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
}

func (s DescribeSslVpnClientCertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSslVpnClientCertRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSslVpnClientCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientCertRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSslVpnClientCertRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSslVpnClientCertRequest) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DescribeSslVpnClientCertRequest) SetOwnerAccount(v string) *DescribeSslVpnClientCertRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) SetOwnerId(v int64) *DescribeSslVpnClientCertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) SetRegionId(v string) *DescribeSslVpnClientCertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) SetResourceOwnerAccount(v string) *DescribeSslVpnClientCertRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) SetResourceOwnerId(v int64) *DescribeSslVpnClientCertRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertRequest {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DescribeSslVpnClientCertRequest) Validate() error {
	return dara.Validate(s)
}
