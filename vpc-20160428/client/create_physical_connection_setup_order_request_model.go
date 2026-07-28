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
	// The access point ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-cn-beijing-ft-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// Specifies whether to enable automatic payment. Valid values:
	//
	// - **false*	- (default): disables automatic payment.
	//
	// - **true**: enables automatic payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may vary for each API request.
	//
	// example:
	//
	// 318BB676-0A2B-43A0-9AD8-F1D34E93750F
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The carrier that provides the physical connection. Valid values:
	//
	// - **CT**: China Telecom
	//
	// - **CU**: China Unicom
	//
	// - **CM**: China Mobile
	//
	// - **CO**: other carriers in the Chinese mainland
	//
	// - **Equinix**: Equinix
	//
	// - **Other**: other carriers outside the Chinese mainland
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
	// - **100Base-T**: 100M Ethernet port.
	//
	// - **1000Base-T*	- (default): 1 GE port.
	//
	// - **1000Base-LX**: GE single-mode optical port (10 km).
	//
	// - **10GBase-T**: 10 GE port.
	//
	// - **10GBase-LR**: 10 GE single-mode optical port (10 km).
	//
	// - **40GBase-LR**: 40 GE single-mode optical port.
	//
	// - **100GBase-LR**: 100 GE single-mode optical port.
	//
	// > 40GBase-LR and 100GBase-LR ports are created based on the actual port availability. Contact your account manager for details.
	//
	// example:
	//
	// 100Base-T
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant Express Connect circuit. The circuit must be in the **Allocated**, **Confirmed**, or **Enabled*	- state.
	//
	// example:
	//
	// pc-bp10zsv5ntp****
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
