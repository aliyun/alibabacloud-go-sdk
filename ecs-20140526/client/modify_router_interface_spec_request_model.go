// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRouterInterfaceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyRouterInterfaceSpecRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ModifyRouterInterfaceSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyRouterInterfaceSpecRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRouterInterfaceSpecRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRouterInterfaceSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRouterInterfaceSpecRequest
	GetResourceOwnerId() *int64
	SetRouterInterfaceId(v string) *ModifyRouterInterfaceSpecRequest
	GetRouterInterfaceId() *string
	SetSpec(v string) *ModifyRouterInterfaceSpecRequest
	GetSpec() *string
	SetUserCidr(v string) *ModifyRouterInterfaceSpecRequest
	GetUserCidr() *string
}

type ModifyRouterInterfaceSpecRequest struct {
	// A client token to ensure that the request is idempotent. You can generate this token from your client, but you must ensure that it is unique across requests. The token must consist of only ASCII characters and not exceed 64 characters in length.
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the router interface is located. To obtain the latest list of regions, call the [DescribeRegions](~~DescribeRegions~~) operation.
	//
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the router interface.
	//
	// This parameter is required.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The new specification of the router interface. The supported specifications vary based on the type and region of the router interface. For more information, see the pricing page of Express Connect.
	//
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The client-side CIDR block for the router interface. This parameter is used for VBR-to-VPC connections. The system assigns a gateway IP address from this CIDR block to the VBR to forward traffic to the VPC.
	UserCidr *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s ModifyRouterInterfaceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRouterInterfaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyRouterInterfaceSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyRouterInterfaceSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyRouterInterfaceSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRouterInterfaceSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRouterInterfaceSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRouterInterfaceSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRouterInterfaceSpecRequest) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *ModifyRouterInterfaceSpecRequest) GetSpec() *string {
	return s.Spec
}

func (s *ModifyRouterInterfaceSpecRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *ModifyRouterInterfaceSpecRequest) SetClientToken(v string) *ModifyRouterInterfaceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetOwnerAccount(v string) *ModifyRouterInterfaceSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetOwnerId(v int64) *ModifyRouterInterfaceSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetRegionId(v string) *ModifyRouterInterfaceSpecRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetResourceOwnerAccount(v string) *ModifyRouterInterfaceSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetResourceOwnerId(v int64) *ModifyRouterInterfaceSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetRouterInterfaceId(v string) *ModifyRouterInterfaceSpecRequest {
	s.RouterInterfaceId = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetSpec(v string) *ModifyRouterInterfaceSpecRequest {
	s.Spec = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) SetUserCidr(v string) *ModifyRouterInterfaceSpecRequest {
	s.UserCidr = &v
	return s
}

func (s *ModifyRouterInterfaceSpecRequest) Validate() error {
	return dara.Validate(s)
}
