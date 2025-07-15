// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcClassicLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableVpcClassicLinkRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DisableVpcClassicLinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableVpcClassicLinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableVpcClassicLinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableVpcClassicLinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableVpcClassicLinkRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DisableVpcClassicLinkRequest
	GetVpcId() *string
}

type DisableVpcClassicLinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC for which you want to disable ClassicLink.
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
	// The ID of the VPC for which you want to disable ClassicLink.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1m7v25emi1h5mtc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DisableVpcClassicLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcClassicLinkRequest) GoString() string {
	return s.String()
}

func (s *DisableVpcClassicLinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableVpcClassicLinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableVpcClassicLinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableVpcClassicLinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableVpcClassicLinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableVpcClassicLinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableVpcClassicLinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DisableVpcClassicLinkRequest) SetClientToken(v string) *DisableVpcClassicLinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetOwnerAccount(v string) *DisableVpcClassicLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetOwnerId(v int64) *DisableVpcClassicLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetRegionId(v string) *DisableVpcClassicLinkRequest {
	s.RegionId = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetResourceOwnerAccount(v string) *DisableVpcClassicLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetResourceOwnerId(v int64) *DisableVpcClassicLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) SetVpcId(v string) *DisableVpcClassicLinkRequest {
	s.VpcId = &v
	return s
}

func (s *DisableVpcClassicLinkRequest) Validate() error {
	return dara.Validate(s)
}
