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
	// The instance ID of the cloud service instance to which you want to attach the EIPs.
	//
	// You can enter the instance ID of a NAT gateway instance or a secondary elastic network interface controller (NIC).
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-hp3akk9irtd69jad****
	BindedInstanceId *string `json:"BindedInstanceId,omitempty" xml:"BindedInstanceId,omitempty"`
	// The type of the instance to which you want to attach the EIPs. Valid values:
	//
	// - **Nat**: NAT gateway.
	//
	// - **NetworkInterface**: secondary elastic network interface controller (NIC).
	//
	// This parameter is required.
	//
	// example:
	//
	// Nat
	BindedInstanceType *string `json:"BindedInstanceType,omitempty" xml:"BindedInstanceType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of EIPs that you want to associate with the cloud service instance.
	//
	// You must specify at least one EIP and can specify up to 50 EIPs.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The association mode. Set the value to **MULTI_BINDED**, which indicates the multi-EIP to ENI mode.
	//
	// This parameter is required only when **InstanceType*	- is set to **NetworkInterface**.
	//
	// example:
	//
	// MULTI_BINDED
	Mode    *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the EIPs that you want to associate with the cloud service instances.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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
