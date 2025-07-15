// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeSslVpnClientCertsRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeSslVpnClientCertsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSslVpnClientCertsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSslVpnClientCertsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSslVpnClientCertsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSslVpnClientCertsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSslVpnClientCertsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSslVpnClientCertsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSslVpnClientCertsRequest
	GetResourceOwnerId() *int64
	SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertsRequest
	GetSslVpnClientCertId() *string
	SetSslVpnServerId(v string) *DescribeSslVpnClientCertsRequest
	GetSslVpnServerId() *string
}

type DescribeSslVpnClientCertsRequest struct {
	// The name of the SSL client certificate.
	//
	// example:
	//
	// cert1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the SSL client certificate.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the SSL client certificate belongs.
	//
	// The SSL client certificate and its associated SSL server belong to the same resource group. You can call the [DescribeSslVpnServers](https://help.aliyun.com/document_detail/2794078.html) operation to query the ID of the resource group to which the SSL server belongs.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-bp1n8wcf134yl0osr****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
}

func (s DescribeSslVpnClientCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSslVpnClientCertsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSslVpnClientCertsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSslVpnClientCertsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSslVpnClientCertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSslVpnClientCertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSslVpnClientCertsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSslVpnClientCertsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSslVpnClientCertsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSslVpnClientCertsRequest) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *DescribeSslVpnClientCertsRequest) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *DescribeSslVpnClientCertsRequest) SetName(v string) *DescribeSslVpnClientCertsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetOwnerAccount(v string) *DescribeSslVpnClientCertsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetOwnerId(v int64) *DescribeSslVpnClientCertsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetPageNumber(v int32) *DescribeSslVpnClientCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetPageSize(v int32) *DescribeSslVpnClientCertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetRegionId(v string) *DescribeSslVpnClientCertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetResourceGroupId(v string) *DescribeSslVpnClientCertsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetResourceOwnerAccount(v string) *DescribeSslVpnClientCertsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetResourceOwnerId(v int64) *DescribeSslVpnClientCertsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetSslVpnClientCertId(v string) *DescribeSslVpnClientCertsRequest {
	s.SslVpnClientCertId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) SetSslVpnServerId(v string) *DescribeSslVpnClientCertsRequest {
	s.SslVpnServerId = &v
	return s
}

func (s *DescribeSslVpnClientCertsRequest) Validate() error {
	return dara.Validate(s)
}
