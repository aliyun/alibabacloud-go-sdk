// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindedInstanceId(v string) *AssociateEipAddressBatchRequest
	GetBindedInstanceId() *string
	SetBindedInstanceType(v string) *AssociateEipAddressBatchRequest
	GetBindedInstanceType() *string
	SetClientToken(v string) *AssociateEipAddressBatchRequest
	GetClientToken() *string
	SetInstanceIds(v []*string) *AssociateEipAddressBatchRequest
	GetInstanceIds() []*string
	SetMode(v string) *AssociateEipAddressBatchRequest
	GetMode() *string
	SetOwnerId(v int64) *AssociateEipAddressBatchRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AssociateEipAddressBatchRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AssociateEipAddressBatchRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateEipAddressBatchRequest
	GetResourceOwnerId() *int64
}

type AssociateEipAddressBatchRequest struct {
	// The ID of the instance with which you want to associate the EIPs.
	//
	// The instance can be a NAT gateway or a secondary ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-hp3akk9irtd69jad****
	BindedInstanceId *string `json:"BindedInstanceId,omitempty" xml:"BindedInstanceId,omitempty"`
	// The type of the instance with which you want to associate the EIPs. Valid values:
	//
	// 	- **Nat**: NAT gateway
	//
	// 	- **NetworkInterface**: secondary ENI
	//
	// This parameter is required.
	//
	// example:
	//
	// Nat
	BindedInstanceType *string `json:"BindedInstanceType,omitempty" xml:"BindedInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The EIPs to be associated with the instance.
	//
	// You must enter at least one EIP. You can enter up to 50 EIPs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The association mode. Set the value to **MULTI_BINDED**, which specifies the Multi-EIP-to-ENI mode.
	//
	// This parameter is required only when **BindedInstanceType*	- is set to **NetworkInterface**.
	//
	// example:
	//
	// MULTI_BINDED
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the EIPs belong. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s AssociateEipAddressBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressBatchRequest) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressBatchRequest) GetBindedInstanceId() *string {
	return s.BindedInstanceId
}

func (s *AssociateEipAddressBatchRequest) GetBindedInstanceType() *string {
	return s.BindedInstanceType
}

func (s *AssociateEipAddressBatchRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateEipAddressBatchRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AssociateEipAddressBatchRequest) GetMode() *string {
	return s.Mode
}

func (s *AssociateEipAddressBatchRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateEipAddressBatchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateEipAddressBatchRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateEipAddressBatchRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateEipAddressBatchRequest) SetBindedInstanceId(v string) *AssociateEipAddressBatchRequest {
	s.BindedInstanceId = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetBindedInstanceType(v string) *AssociateEipAddressBatchRequest {
	s.BindedInstanceType = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetClientToken(v string) *AssociateEipAddressBatchRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetInstanceIds(v []*string) *AssociateEipAddressBatchRequest {
	s.InstanceIds = v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetMode(v string) *AssociateEipAddressBatchRequest {
	s.Mode = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetOwnerId(v int64) *AssociateEipAddressBatchRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetRegionId(v string) *AssociateEipAddressBatchRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetResourceOwnerAccount(v string) *AssociateEipAddressBatchRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) SetResourceOwnerId(v int64) *AssociateEipAddressBatchRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateEipAddressBatchRequest) Validate() error {
	return dara.Validate(s)
}
