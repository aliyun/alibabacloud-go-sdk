// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcIpamServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetVpcIpamServiceStatusRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *GetVpcIpamServiceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetVpcIpamServiceStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetVpcIpamServiceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetVpcIpamServiceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetVpcIpamServiceStatusRequest
	GetResourceOwnerId() *int64
}

type GetVpcIpamServiceStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted.
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

func (s GetVpcIpamServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcIpamServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVpcIpamServiceStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetVpcIpamServiceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetVpcIpamServiceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetVpcIpamServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcIpamServiceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetVpcIpamServiceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetVpcIpamServiceStatusRequest) SetClientToken(v string) *GetVpcIpamServiceStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetOwnerAccount(v string) *GetVpcIpamServiceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetOwnerId(v int64) *GetVpcIpamServiceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetRegionId(v string) *GetVpcIpamServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetResourceOwnerAccount(v string) *GetVpcIpamServiceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) SetResourceOwnerId(v int64) *GetVpcIpamServiceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVpcIpamServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
