// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVpnAttachmentRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteVpnAttachmentRequest
	GetOwnerAccount() *string
	SetRegionId(v string) *DeleteVpnAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteVpnAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVpnAttachmentRequest
	GetResourceOwnerId() *int64
	SetVpnConnectionId(v string) *DeleteVpnAttachmentRequest
	GetVpnConnectionId() *string
}

type DeleteVpnAttachmentRequest struct {
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
	// The ID of the IPsec-VPN connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// vco-p0w7gtr14m09r9lkr****
	VpnConnectionId *string `json:"VpnConnectionId,omitempty" xml:"VpnConnectionId,omitempty"`
}

func (s DeleteVpnAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnAttachmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpnAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVpnAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteVpnAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpnAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVpnAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVpnAttachmentRequest) GetVpnConnectionId() *string {
	return s.VpnConnectionId
}

func (s *DeleteVpnAttachmentRequest) SetClientToken(v string) *DeleteVpnAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) SetOwnerAccount(v string) *DeleteVpnAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) SetRegionId(v string) *DeleteVpnAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) SetResourceOwnerAccount(v string) *DeleteVpnAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) SetResourceOwnerId(v int64) *DeleteVpnAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) SetVpnConnectionId(v string) *DeleteVpnAttachmentRequest {
	s.VpnConnectionId = &v
	return s
}

func (s *DeleteVpnAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
