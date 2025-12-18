// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *CreatePhysicalConnectionRequest
	GetAccessPointId() *string
	SetCircuitCode(v string) *CreatePhysicalConnectionRequest
	GetCircuitCode() *string
	SetClientToken(v string) *CreatePhysicalConnectionRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePhysicalConnectionRequest
	GetDescription() *string
	SetDeviceAdvancedCapacity(v []*string) *CreatePhysicalConnectionRequest
	GetDeviceAdvancedCapacity() []*string
	SetLineOperator(v string) *CreatePhysicalConnectionRequest
	GetLineOperator() *string
	SetName(v string) *CreatePhysicalConnectionRequest
	GetName() *string
	SetOpticalModuleModel(v string) *CreatePhysicalConnectionRequest
	GetOpticalModuleModel() *string
	SetOwnerAccount(v string) *CreatePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPeerLocation(v string) *CreatePhysicalConnectionRequest
	GetPeerLocation() *string
	SetPortType(v string) *CreatePhysicalConnectionRequest
	GetPortType() *string
	SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionRequest
	GetRedundantPhysicalConnectionId() *string
	SetRegionId(v string) *CreatePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePhysicalConnectionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreatePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePhysicalConnectionRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreatePhysicalConnectionRequestTag) *CreatePhysicalConnectionRequest
	GetTag() []*CreatePhysicalConnectionRequestTag
	SetType(v string) *CreatePhysicalConnectionRequest
	GetType() *string
	SetBandwidth(v int32) *CreatePhysicalConnectionRequest
	GetBandwidth() *int32
}

type CreatePhysicalConnectionRequest struct {
	// The access point ID of the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-cn-beijing-ft-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The circuit code of the Express Connect circuit. The circuit code is provided by the connectivity provider.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Express Connect circuit.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The advanced features of the device.
	DeviceAdvancedCapacity []*string `json:"DeviceAdvancedCapacity,omitempty" xml:"DeviceAdvancedCapacity,omitempty" type:"Repeated"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// 	- **CT**: China Telecom.
	//
	// 	- **CU**: China Unicom.
	//
	// 	- **CM**: China Mobile.
	//
	// 	- **CO**: other connectivity providers in the Chinese mainland.
	//
	// 	- **Equinix**: Equinix.
	//
	// 	- **Other**: other connectivity providers outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// CT
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The name of the Express Connect circuit.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OpticalModuleModel *string `json:"OpticalModuleModel,omitempty" xml:"OpticalModuleModel,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The geographical location of the data center.
	//
	// example:
	//
	// cn-hangzhou
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The port type. Valid values:
	//
	// 	- **100Base-T**: 100 Mbit/s copper Ethernet port
	//
	// 	- **1000Base-T**: 1,000 Mbit/s copper Ethernet port
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
	// >  To use ports 40GBase-LR and 100GBase-LR, you must first contact your account manager.
	//
	// example:
	//
	// 1000Base-T
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The ID of the redundant Express Connect circuit. The redundant Express Connect circuit must be in the **Allocated**, **Confirmed**, or **Enabled*	- state.
	//
	// example:
	//
	// pc-119mfjzm****
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Express Connect circuit belongs.
	//
	// example:
	//
	// rg-acfmoiyermp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*CreatePhysicalConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The type of Express Connect circuit. Default value: **VPC**.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The maximum bandwidth of the hosted connection. Unit: Mbit/s.
	//
	// Valid values: **50**, **100**, **200**, **300**, **400**, **500**, **1000**, **2000**, **4000**, **5000**, **8000**, and **10000**.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"bandwidth,omitempty" xml:"bandwidth,omitempty"`
}

func (s CreatePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreatePhysicalConnectionRequest) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *CreatePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePhysicalConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePhysicalConnectionRequest) GetDeviceAdvancedCapacity() []*string {
	return s.DeviceAdvancedCapacity
}

func (s *CreatePhysicalConnectionRequest) GetLineOperator() *string {
	return s.LineOperator
}

func (s *CreatePhysicalConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreatePhysicalConnectionRequest) GetOpticalModuleModel() *string {
	return s.OpticalModuleModel
}

func (s *CreatePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePhysicalConnectionRequest) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *CreatePhysicalConnectionRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreatePhysicalConnectionRequest) GetRedundantPhysicalConnectionId() *string {
	return s.RedundantPhysicalConnectionId
}

func (s *CreatePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePhysicalConnectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePhysicalConnectionRequest) GetTag() []*CreatePhysicalConnectionRequestTag {
	return s.Tag
}

func (s *CreatePhysicalConnectionRequest) GetType() *string {
	return s.Type
}

func (s *CreatePhysicalConnectionRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreatePhysicalConnectionRequest) SetAccessPointId(v string) *CreatePhysicalConnectionRequest {
	s.AccessPointId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetCircuitCode(v string) *CreatePhysicalConnectionRequest {
	s.CircuitCode = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetClientToken(v string) *CreatePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetDescription(v string) *CreatePhysicalConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetDeviceAdvancedCapacity(v []*string) *CreatePhysicalConnectionRequest {
	s.DeviceAdvancedCapacity = v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetLineOperator(v string) *CreatePhysicalConnectionRequest {
	s.LineOperator = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetName(v string) *CreatePhysicalConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetOpticalModuleModel(v string) *CreatePhysicalConnectionRequest {
	s.OpticalModuleModel = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetOwnerAccount(v string) *CreatePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetOwnerId(v int64) *CreatePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetPeerLocation(v string) *CreatePhysicalConnectionRequest {
	s.PeerLocation = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetPortType(v string) *CreatePhysicalConnectionRequest {
	s.PortType = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetRedundantPhysicalConnectionId(v string) *CreatePhysicalConnectionRequest {
	s.RedundantPhysicalConnectionId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetRegionId(v string) *CreatePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetResourceGroupId(v string) *CreatePhysicalConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *CreatePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetResourceOwnerId(v int64) *CreatePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetTag(v []*CreatePhysicalConnectionRequestTag) *CreatePhysicalConnectionRequest {
	s.Tag = v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetType(v string) *CreatePhysicalConnectionRequest {
	s.Type = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) SetBandwidth(v int32) *CreatePhysicalConnectionRequest {
	s.Bandwidth = &v
	return s
}

func (s *CreatePhysicalConnectionRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePhysicalConnectionRequestTag struct {
	// The key of tag to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value to add to the resource. You can specify up to 20 tag values The tag value can be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePhysicalConnectionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePhysicalConnectionRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePhysicalConnectionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePhysicalConnectionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePhysicalConnectionRequestTag) SetKey(v string) *CreatePhysicalConnectionRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePhysicalConnectionRequestTag) SetValue(v string) *CreatePhysicalConnectionRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePhysicalConnectionRequestTag) Validate() error {
	return dara.Validate(s)
}
