// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighReliablePhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateHighReliablePhysicalConnectionRequest
	GetAcceptLanguage() *string
	SetApList(v []*CreateHighReliablePhysicalConnectionRequestApList) *CreateHighReliablePhysicalConnectionRequest
	GetApList() []*CreateHighReliablePhysicalConnectionRequestApList
	SetClientToken(v string) *CreateHighReliablePhysicalConnectionRequest
	GetClientToken() *string
	SetDeviceAdvancedCapacity(v []*string) *CreateHighReliablePhysicalConnectionRequest
	GetDeviceAdvancedCapacity() []*string
	SetDryRun(v string) *CreateHighReliablePhysicalConnectionRequest
	GetDryRun() *string
	SetHighReliableType(v string) *CreateHighReliablePhysicalConnectionRequest
	GetHighReliableType() *string
	SetOwnerAccount(v string) *CreateHighReliablePhysicalConnectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateHighReliablePhysicalConnectionRequest
	GetOwnerId() *int64
	SetPortType(v string) *CreateHighReliablePhysicalConnectionRequest
	GetPortType() *string
	SetRegionId(v string) *CreateHighReliablePhysicalConnectionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateHighReliablePhysicalConnectionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateHighReliablePhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHighReliablePhysicalConnectionRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateHighReliablePhysicalConnectionRequestTag) *CreateHighReliablePhysicalConnectionRequest
	GetTag() []*CreateHighReliablePhysicalConnectionRequestTag
}

type CreateHighReliablePhysicalConnectionRequest struct {
	// The language to display the results. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The access points.
	//
	// This parameter is required.
	ApList []*CreateHighReliablePhysicalConnectionRequestApList `json:"ApList,omitempty" xml:"ApList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The advanced features of the device.
	DeviceAdvancedCapacity []*string `json:"DeviceAdvancedCapacity,omitempty" xml:"DeviceAdvancedCapacity,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The high availability mode. Valid values:
	//
	// - **MultiApMultiDevice*	- : This mode supports two access points and two devices, and provides the maximum disaster recovery capability.
	//
	// - **MultiApSingleDevice*	- : This mode supports two access points and one device, and provides robust disaster recovery capability.
	//
	// - **SingleApMultiDevice*	- : This mode supports one access point and two devices, and is recommended for non-critical business test and development.
	//
	// - **SingleApMultiConnection*	- : This mode supports one access point, one device, and multiple physical ports. Only users in the whitelist can use this mode. To use this mode, contact your account manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// MultiApMultiDevice
	HighReliableType *string `json:"HighReliableType,omitempty" xml:"HighReliableType,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1000Base-T
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateHighReliablePhysicalConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateHighReliablePhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetApList() []*CreateHighReliablePhysicalConnectionRequestApList {
	return s.ApList
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetDeviceAdvancedCapacity() []*string {
	return s.DeviceAdvancedCapacity
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetHighReliableType() *string {
	return s.HighReliableType
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetPortType() *string {
	return s.PortType
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHighReliablePhysicalConnectionRequest) GetTag() []*CreateHighReliablePhysicalConnectionRequestTag {
	return s.Tag
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetAcceptLanguage(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetApList(v []*CreateHighReliablePhysicalConnectionRequestApList) *CreateHighReliablePhysicalConnectionRequest {
	s.ApList = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetClientToken(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetDeviceAdvancedCapacity(v []*string) *CreateHighReliablePhysicalConnectionRequest {
	s.DeviceAdvancedCapacity = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetDryRun(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetHighReliableType(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.HighReliableType = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetOwnerAccount(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetOwnerId(v int64) *CreateHighReliablePhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetPortType(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.PortType = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetRegionId(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetResourceGroupId(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetResourceOwnerAccount(v string) *CreateHighReliablePhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetResourceOwnerId(v int64) *CreateHighReliablePhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) SetTag(v []*CreateHighReliablePhysicalConnectionRequestTag) *CreateHighReliablePhysicalConnectionRequest {
	s.Tag = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequest) Validate() error {
	if s.ApList != nil {
		for _, item := range s.ApList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateHighReliablePhysicalConnectionRequestApList struct {
	// The ID of the access point that is associated with the Express Connect circuit.
	//
	// > Two access points must be specified when **HighReliableType*	- is set to **MultiApMultiDevice*	- or **MultiApSingleDevice**. One access point must be specified when **HighReliableType*	- is set to **SingleApMultiDevice*	- or **SingleApMultiConnection**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-cn-beijing-ft-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The maximum bandwidth of the hosted connection. Unit: Mbit/s.
	//
	// Valid values: 50, 100, 200, 300, 400, 500, 1000, 2000, 4000, 5000, 8000, and 10000.
	//
	// example:
	//
	// 50
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The circuit code of the Express Connect circuit, which is provided by the connectivity provider.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The description of the Express Connect circuit.
	//
	// The description must be 2 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter but cannot start with `http://` or` https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The geographical location of the data center.
	//
	// example:
	//
	// ram-test
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The number of ports. Valid values: 2 to 16. This parameter is required only when **HighReliableType*	- is set to **SingleApMultiConnection**.
	//
	// example:
	//
	// 2
	PortNum *int32 `json:"PortNum,omitempty" xml:"PortNum,omitempty"`
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
	// The type of the Express Connect circuit. Default value: **VPC**.
	//
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionRequestApList) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionRequestApList) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetDescription() *string {
	return s.Description
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetLineOperator() *string {
	return s.LineOperator
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetName() *string {
	return s.Name
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetPeerLocation() *string {
	return s.PeerLocation
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetPortNum() *int32 {
	return s.PortNum
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetType() *string {
	return s.Type
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetAccessPointId(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.AccessPointId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetBandwidth(v int64) *CreateHighReliablePhysicalConnectionRequestApList {
	s.Bandwidth = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetCircuitCode(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.CircuitCode = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetDescription(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.Description = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetLineOperator(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.LineOperator = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetName(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.Name = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetPeerLocation(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.PeerLocation = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetPortNum(v int32) *CreateHighReliablePhysicalConnectionRequestApList {
	s.PortNum = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetRegionId(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.RegionId = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetType(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.Type = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestApList) Validate() error {
	return dara.Validate(s)
}

type CreateHighReliablePhysicalConnectionRequestTag struct {
	// The key of tag N to add to the resource. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. Valid values of N: 1 to 20. The tag value cannot be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateHighReliablePhysicalConnectionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateHighReliablePhysicalConnectionRequestTag) SetKey(v string) *CreateHighReliablePhysicalConnectionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestTag) SetValue(v string) *CreateHighReliablePhysicalConnectionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionRequestTag) Validate() error {
	return dara.Validate(s)
}
