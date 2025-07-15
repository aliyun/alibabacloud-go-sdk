// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhysicalConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCircuitCode(v string) *ModifyPhysicalConnectionAttributeRequest
	GetCircuitCode() *string
	SetClientToken(v string) *ModifyPhysicalConnectionAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyPhysicalConnectionAttributeRequest
	GetDescription() *string
	SetLineOperator(v string) *ModifyPhysicalConnectionAttributeRequest
	GetLineOperator() *string
	SetName(v string) *ModifyPhysicalConnectionAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest
	GetOwnerId() *int64
	SetPeerLocation(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPeerLocation() *string
	SetPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPhysicalConnectionId() *string
	SetPortType(v string) *ModifyPhysicalConnectionAttributeRequest
	GetPortType() *string
	SetRedundantPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetRedundantPhysicalConnectionId() *string
	SetRegionId(v string) *ModifyPhysicalConnectionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest
	GetResourceOwnerId() *int64
	SetBandwidth(v int32) *ModifyPhysicalConnectionAttributeRequest
	GetBandwidth() *int32
}

type ModifyPhysicalConnectionAttributeRequest struct {
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// efefe566754h
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Express Connect circuit.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// The description of the Express Connect circuit.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// example:
	//
	// CT
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The name of the Express Connect circuit.
	//
	// The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Name
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The geographical location of the data center.
	//
	// example:
	//
	// XX Number, XX Road, XX Town, XX District, Hangzhou City, Zhejiang Province.
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-119mfjzm******
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// 	- **100Base-T**: 100 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-T*	- (default): 1,000 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-LX**: 1,000 Mbit/s single-mode optical port (10 kilometers)
	//
	// 	- **10GBase-T**: 10,000 Mbit/s copper Ethernet port
	//
	// 	- **10GBase-LR**: 10,000 Mbit/s single-mode optical port (10 kilometers)
	//
	// 	- **40GBase-LR**: 40,000 Mbit/s single-mode optical port
	//
	// 	- **100GBase-LR**: 100,000 Mbit/s single-mode optical port
	//
	// >  To use ports 40GBase-LR and 100GBase-LR, you must first contact your account manager.
	//
	// example:
	//
	// 1000Base-LX
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant Express Connect circuit. The redundant Express Connect circuit must be in the **Allocated**, **Confirmed**, or **Enabled*	- state.
	//
	// example:
	//
	// pc-119mfjzm7
	RedundantPhysicalConnectionId *string `json:"RedundantPhysicalConnectionId,omitempty" xml:"RedundantPhysicalConnectionId,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The bandwidth value for the connection over the Express Connect circuit. Unit: Mbit/s. Valid values: 2 to 10240.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
}

func (s ModifyPhysicalConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhysicalConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetLineOperator() *string {
	return s.LineOperator
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetPortType() *string {
	return s.PortType
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPhysicalConnectionAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetCircuitCode(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.CircuitCode = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetClientToken(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetDescription(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetLineOperator(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.LineOperator = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetName(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPeerLocation(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PeerLocation = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetPortType(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.PortType = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetRedundantPhysicalConnectionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetRegionId(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetResourceOwnerAccount(v string) *ModifyPhysicalConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetResourceOwnerId(v int64) *ModifyPhysicalConnectionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) SetBandwidth(v int32) *ModifyPhysicalConnectionAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyPhysicalConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
