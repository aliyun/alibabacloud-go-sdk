// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpsecServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpsecServerId(v []*string) *ListIpsecServersRequest
	GetIpsecServerId() []*string
	SetIpsecServerName(v string) *ListIpsecServersRequest
	GetIpsecServerName() *string
	SetMaxResults(v int32) *ListIpsecServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpsecServersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpsecServersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpsecServersRequest
	GetResourceGroupId() *string
	SetVpnGatewayId(v string) *ListIpsecServersRequest
	GetVpnGatewayId() *string
}

type ListIpsecServersRequest struct {
	// The ID of the IPsec server.
	//
	// example:
	//
	// iss-bp1bo3xuvcxo7ixll****
	IpsecServerId []*string `json:"IpsecServerId,omitempty" xml:"IpsecServerId,omitempty" type:"Repeated"`
	// The name of the IPsec server.
	//
	// The name must be 1 to 100 characters in length.
	//
	// example:
	//
	// test
	IpsecServerName *string `json:"IpsecServerName,omitempty" xml:"IpsecServerName,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **20**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of **NextToken**.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the IPsec server is created.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the IPsec server belongs.
	//
	// The IPsec server and its associated VPN gateway belong to the same resource group. You can call [DescribeVpnGateway](https://help.aliyun.com/document_detail/2794055.html) to query the ID of the resource group to which the VPN gateway belongs.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the VPN gateway.
	//
	// example:
	//
	// vpn-bp1q8bgx4xnkm2ogj****
	VpnGatewayId *string `json:"VpnGatewayId,omitempty" xml:"VpnGatewayId,omitempty"`
}

func (s ListIpsecServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpsecServersRequest) GoString() string {
	return s.String()
}

func (s *ListIpsecServersRequest) GetIpsecServerId() []*string {
	return s.IpsecServerId
}

func (s *ListIpsecServersRequest) GetIpsecServerName() *string {
	return s.IpsecServerName
}

func (s *ListIpsecServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpsecServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpsecServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpsecServersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpsecServersRequest) GetVpnGatewayId() *string {
	return s.VpnGatewayId
}

func (s *ListIpsecServersRequest) SetIpsecServerId(v []*string) *ListIpsecServersRequest {
	s.IpsecServerId = v
	return s
}

func (s *ListIpsecServersRequest) SetIpsecServerName(v string) *ListIpsecServersRequest {
	s.IpsecServerName = &v
	return s
}

func (s *ListIpsecServersRequest) SetMaxResults(v int32) *ListIpsecServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpsecServersRequest) SetNextToken(v string) *ListIpsecServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpsecServersRequest) SetRegionId(v string) *ListIpsecServersRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpsecServersRequest) SetResourceGroupId(v string) *ListIpsecServersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpsecServersRequest) SetVpnGatewayId(v string) *ListIpsecServersRequest {
	s.VpnGatewayId = &v
	return s
}

func (s *ListIpsecServersRequest) Validate() error {
	return dara.Validate(s)
}
