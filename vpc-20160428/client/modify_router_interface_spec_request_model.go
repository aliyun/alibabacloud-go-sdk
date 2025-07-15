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
}

type ModifyRouterInterfaceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the router interface is deployed.
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
	// The ID of the router interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ri-2zeo3xzyf38r4urzd****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The specification of the router interface. Valid specifications and bandwidth values:
	//
	// 	- **Mini.2**: 2 Mbit/s
	//
	// 	- **Mini.5**: 5 Mbit/s
	//
	// 	- **Small.1**: 10 Mbit/s
	//
	// 	- **Small.2**: 20 Mbit/s
	//
	// 	- **Small.5**: 50 Mbit/s
	//
	// 	- **Middle.1**: 100 Mbit/s
	//
	// 	- **Middle.2**: 200 Mbit/s
	//
	// 	- **Middle.5**: 500 Mbit/s
	//
	// 	- **Large.1**: 1,000 Mbit/s
	//
	// 	- **Large.2**: 2,000 Mbit/s
	//
	// 	- **Large.5**: 5,000 Mbit/s
	//
	// 	- **Xlarge.1**: 10,000 Mbit/s
	//
	// >  When **Role*	- is set to **AcceptingSide**, set **Spec*	- to **Negative**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Small.1
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
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

func (s *ModifyRouterInterfaceSpecRequest) Validate() error {
	return dara.Validate(s)
}
