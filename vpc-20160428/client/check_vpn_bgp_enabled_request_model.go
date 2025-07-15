// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVpnBgpEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CheckVpnBgpEnabledRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CheckVpnBgpEnabledRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *CheckVpnBgpEnabledRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CheckVpnBgpEnabledRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckVpnBgpEnabledRequest
	GetResourceOwnerId() *int64
}

type CheckVpnBgpEnabledRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// The region ID of the IPsec-VPN connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CheckVpnBgpEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckVpnBgpEnabledRequest) GoString() string {
	return s.String()
}

func (s *CheckVpnBgpEnabledRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckVpnBgpEnabledRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckVpnBgpEnabledRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckVpnBgpEnabledRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckVpnBgpEnabledRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckVpnBgpEnabledRequest) SetClientToken(v string) *CheckVpnBgpEnabledRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckVpnBgpEnabledRequest) SetOwnerAccount(v string) *CheckVpnBgpEnabledRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckVpnBgpEnabledRequest) SetRegionId(v string) *CheckVpnBgpEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *CheckVpnBgpEnabledRequest) SetResourceOwnerAccount(v string) *CheckVpnBgpEnabledRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckVpnBgpEnabledRequest) SetResourceOwnerId(v int64) *CheckVpnBgpEnabledRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckVpnBgpEnabledRequest) Validate() error {
	return dara.Validate(s)
}
