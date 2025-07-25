// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVpcIpamServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *OpenVpcIpamServiceRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *OpenVpcIpamServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenVpcIpamServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenVpcIpamServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenVpcIpamServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenVpcIpamServiceRequest
	GetResourceOwnerId() *int64
}

type OpenVpcIpamServiceRequest struct {
	// Client token, used to ensure the idempotence of requests.
	//
	// Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// > If not specified, the system automatically uses the RequestId of the API request as the ClientToken identifier. The RequestId may differ for each API request.
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

func (s OpenVpcIpamServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenVpcIpamServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenVpcIpamServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *OpenVpcIpamServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenVpcIpamServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenVpcIpamServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenVpcIpamServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenVpcIpamServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenVpcIpamServiceRequest) SetClientToken(v string) *OpenVpcIpamServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetOwnerAccount(v string) *OpenVpcIpamServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetOwnerId(v int64) *OpenVpcIpamServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetRegionId(v string) *OpenVpcIpamServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetResourceOwnerAccount(v string) *OpenVpcIpamServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) SetResourceOwnerId(v int64) *OpenVpcIpamServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenVpcIpamServiceRequest) Validate() error {
	return dara.Validate(s)
}
