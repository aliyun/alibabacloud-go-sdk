// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionSetupOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetAccessPointId() *string
	SetAutoPay(v bool) *CreatePhysicalConnectionSetupOrderRequest
	GetAutoPay() *bool
	SetClientToken(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetClientToken() *string
	SetLineOperator(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetLineOperator() *string
	SetOwnerAccount(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePhysicalConnectionSetupOrderRequest
	GetOwnerId() *int64
	SetPortType(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetPortType() *string
	SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetRedundantPhysicalConnectionId() *string
	SetRegionId(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreatePhysicalConnectionSetupOrderRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhysicalConnectionSetupOrderRequest
	GetResourceOwnerId() *int64
}

type CreatePhysicalConnectionSetupOrderRequest struct {
	// The ID of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-cn-beijing-ft-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Specifies whether to enable automatic payments. Valid values:
	//
	// 	- **false*	- (default): disables automatic payment.
	//
	// 	- **true**
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 318BB676-0A2B-43A0-9AD8-F1D34E93750F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CT**: China Telecom
	//
	// 	- **CU**: China Unicom
	//
	// 	- **CM**: China Mobile
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland
	//
	// 	- **Equinix**: Equinix
	//
	// 	- **Other**: other connectivity providers outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// CT
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100Base-T**: 100 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-T*	- (default): 1,000 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-LX**: 1,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **10GBase-T**: 10,000 Mbit/s copper Ethernet port
	//
	// 	- **10GBase-LR**: 10,000 Mbit/s single-mode optical port (10 km)
	//
	// 	- **40GBase-LR**: 40,000 Mbit/s single-mode optical port
	//
	// 	- **100GBase-LR**: 100,000 Mbit/s single-mode optical port
	//
	// >  Whether 40GBase-LR and 100GBase-LR ports can be created depends on resource supplies. For more information, contact your account manager.
	//
	// example:
	//
	// 100Base-T
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant physical connection. The redundant physical connection must be in the **Allocated**, **Confirmed**, or **Enabled*	- state.
	//
	// example:
	//
	// pc-bp10zsv5ntp****
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreatePhysicalConnectionSetupOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionSetupOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetLineOperator() *string {
	return s.LineOperator
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhysicalConnectionSetupOrderRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetAccessPointId(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetAutoPay(v bool) *CreatePhysicalConnectionSetupOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetClientToken(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetLineOperator(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.LineOperator = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetOwnerAccount(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetOwnerId(v int64) *CreatePhysicalConnectionSetupOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetPortType(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.PortType = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetRegionId(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetResourceOwnerAccount(v string) *CreatePhysicalConnectionSetupOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) SetResourceOwnerId(v int64) *CreatePhysicalConnectionSetupOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionSetupOrderRequest) Validate() error {
	return dara.Validate(s)
}
