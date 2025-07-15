// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpv4GatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateIpv4GatewayAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateIpv4GatewayAttributeRequest
	GetDryRun() *bool
	SetIpv4GatewayDescription(v string) *UpdateIpv4GatewayAttributeRequest
	GetIpv4GatewayDescription() *string
	SetIpv4GatewayId(v string) *UpdateIpv4GatewayAttributeRequest
	GetIpv4GatewayId() *string
	SetIpv4GatewayName(v string) *UpdateIpv4GatewayAttributeRequest
	GetIpv4GatewayName() *string
	SetOwnerAccount(v string) *UpdateIpv4GatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateIpv4GatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateIpv4GatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateIpv4GatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpv4GatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type UpdateIpv4GatewayAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The new description of the IPv4 gateway.
	//
	// example:
	//
	// new
	Ipv4GatewayDescription *string `json:"Ipv4GatewayDescription,omitempty" xml:"Ipv4GatewayDescription,omitempty"`
	// The ID of the IPv4 gateway whose name or description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	// The new name of the IPv4 gateway.
	//
	// example:
	//
	// newname
	Ipv4GatewayName *string `json:"Ipv4GatewayName,omitempty" xml:"Ipv4GatewayName,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv4 gateway whose name or description you want to modify.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-6
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpv4GatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpv4GatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpv4GatewayAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateIpv4GatewayAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateIpv4GatewayAttributeRequest) GetIpv4GatewayDescription() *string {
	return s.Ipv4GatewayDescription
}

func (s *UpdateIpv4GatewayAttributeRequest) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *UpdateIpv4GatewayAttributeRequest) GetIpv4GatewayName() *string {
	return s.Ipv4GatewayName
}

func (s *UpdateIpv4GatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateIpv4GatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpv4GatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateIpv4GatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpv4GatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpv4GatewayAttributeRequest) SetClientToken(v string) *UpdateIpv4GatewayAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetDryRun(v bool) *UpdateIpv4GatewayAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetIpv4GatewayDescription(v string) *UpdateIpv4GatewayAttributeRequest {
	s.Ipv4GatewayDescription = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetIpv4GatewayId(v string) *UpdateIpv4GatewayAttributeRequest {
	s.Ipv4GatewayId = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetIpv4GatewayName(v string) *UpdateIpv4GatewayAttributeRequest {
	s.Ipv4GatewayName = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetOwnerAccount(v string) *UpdateIpv4GatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetOwnerId(v int64) *UpdateIpv4GatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetRegionId(v string) *UpdateIpv4GatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetResourceOwnerAccount(v string) *UpdateIpv4GatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) SetResourceOwnerId(v int64) *UpdateIpv4GatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpv4GatewayAttributeRequest) Validate() error {
	return dara.Validate(s)
}
