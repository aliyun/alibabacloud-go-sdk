// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnCrossAccountAuthorizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeVpnCrossAccountAuthorizationsRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeVpnCrossAccountAuthorizationsRequest
	GetOwnerAccount() *string
	SetPageNumber(v int32) *DescribeVpnCrossAccountAuthorizationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpnCrossAccountAuthorizationsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpnCrossAccountAuthorizationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpnCrossAccountAuthorizationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpnCrossAccountAuthorizationsRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DescribeVpnCrossAccountAuthorizationsRequest
	GetVpnConnectionId() *string
}

type DescribeVpnCrossAccountAuthorizationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region to which the IPsec-VPN connection belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w2jpkhi2eeop6q6****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DescribeVpnCrossAccountAuthorizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnCrossAccountAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetClientToken(v string) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetOwnerAccount(v string) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetPageNumber(v int32) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetPageSize(v int32) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetRegionId(v string) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetResourceOwnerAccount(v string) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetResourceOwnerId(v int64) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) SetVpnConnectionId(v string) *DescribeVpnCrossAccountAuthorizationsRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DescribeVpnCrossAccountAuthorizationsRequest) Validate() error {
	return dara.Validate(s)
}
