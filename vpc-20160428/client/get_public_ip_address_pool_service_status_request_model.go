// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicIpAddressPoolServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetPublicIpAddressPoolServiceStatusRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *GetPublicIpAddressPoolServiceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetPublicIpAddressPoolServiceStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetPublicIpAddressPoolServiceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetPublicIpAddressPoolServiceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetPublicIpAddressPoolServiceStatusRequest
	GetResourceOwnerId() *int64
}

type GetPublicIpAddressPoolServiceStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655442455
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IP address pool.
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

func (s GetPublicIpAddressPoolServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublicIpAddressPoolServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetClientToken(v string) *GetPublicIpAddressPoolServiceStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetOwnerAccount(v string) *GetPublicIpAddressPoolServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetOwnerId(v int64) *GetPublicIpAddressPoolServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetRegionId(v string) *GetPublicIpAddressPoolServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetResourceOwnerAccount(v string) *GetPublicIpAddressPoolServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) SetResourceOwnerId(v int64) *GetPublicIpAddressPoolServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPublicIpAddressPoolServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
