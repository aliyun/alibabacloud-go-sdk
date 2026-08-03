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
	// The language of the response. Valid values:
	//
	// - **zh-CN*	- (default): Chinese.
	//
	// - **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The list of access points.
	//
	// This parameter is required.
	ApList []*CreateHighReliablePhysicalConnectionRequestApList `json:"ApList,omitempty" xml:"ApList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of advanced device capabilities.
	DeviceAdvancedCapacity []*string `json:"DeviceAdvancedCapacity,omitempty" xml:"DeviceAdvancedCapacity,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the instance. The system checks the required parameters, request format, and instance status. If the check fails, the error code `DRYRUN.FAIL` is returned along with the corresponding error list. If the check succeeds, the code `DRYRUN.SUCCESS` is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, the instance is created.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The zone redundancy mode. Valid values:
	//
	// - **MultiApMultiDevice**: Maximum disaster recovery. This mode uses two different access points and two different devices, providing maximum disaster recovery.
	//
	// - **MultiApSingleDevice**: Enhanced disaster recovery. This mode uses two different access points and one device, providing enhanced disaster recovery.
	//
	// - **SingleApMultiDevice**: Development and testing. This mode uses one access point and two devices. It is recommended only for development and testing of non-critical workloads.
	//
	// - **SingleApMultiConnection**: High-bandwidth load balancing. This mode is available only to users in the whitelist. It uses one access point, one device, and multiple physical ports. Contact your account manager if you need this mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// MultiApMultiDevice
	HighReliableType *string `json:"HighReliableType,omitempty" xml:"HighReliableType,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port type of the Express Connect circuit. Valid values:
	//
	// - **1000Base-LX**: GE single-mode optical port (10 km).
	//
	// - **10GBase-LR**: 10 GE single-mode optical port (10 km).
	//
	// - **40GBase-LR**: 40 GE single-mode optical port.
	//
	// - **100GBase-LR**: 100 GE single-mode optical port.
	//
	//
	//
	// > 40GBase-LR and 100GBase-LR are subject to the actual port availability in the backend. Contact your account manager for port availability details.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000Base-LX
	PortType *string `json:"PortType,omitempty" xml:"PortType,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Express Connect circuit belongs.
	//
	// example:
	//
	// rg-acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
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
	// The ID of the access point where the Express Connect circuit is located.
	//
	// > When **HighReliableType*	- is set to **MultiApMultiDevice*	- or **MultiApSingleDevice**, you must specify two different access points. When **HighReliableType*	- is set to **SingleApMultiDevice*	- or **SingleApMultiConnection**, you must specify one access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-cn-beijing-ft-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The bandwidth of the shared Express Connect circuits. Unit: Mbit/s.
	//
	// Valid values: 50, 100, 200, 300, 400, 500, 1000, 2000, 4000, 5000, 8000, and 10000.
	//
	// example:
	//
	// 50
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The circuit code provided by the connectivity provider for the Express Connect circuit.
	//
	// example:
	//
	// longtel001
	CircuitCode *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	// The description of the Express Connect circuit.
	//
	// The description must be 2 to 256 characters in length, and must start with a letter or Chinese character, but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The connectivity provider of the Express Connect circuit. Valid values:
	//
	// - **CT**: China Telecom.
	//
	// - **CU**: China Unicom.
	//
	// - **CM**: China Mobile.
	//
	// - **CO**: Other Chinese carriers.
	//
	// - **Equinix**: Equinix.
	//
	// - **Other**: Other carriers outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// CT
	LineOperator *string `json:"LineOperator,omitempty" xml:"LineOperator,omitempty"`
	// The name of the Express Connect circuit.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or Chinese character. It can contain digits, underscores (_), and hyphens (-), but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The optical module model supported by the Express Connect circuit access point. Valid values:
	//
	// - 1000Base-LX :
	//
	//   - `SFP-GE-LR-SM1310,10KM`
	//
	//   - `SFP-GE-ER-SM1310,40KM`
	//
	//   - `SFP-GE-ZR-SM1550,80KM`
	//
	// - 10GBase-LR :
	//
	//   - `SFP-10G-LR-SM1310,10KM`
	//
	//   - `SFP-10G-ER-SM1550,40KM`
	//
	//   - `SFP-10G-ZR-SM1550,80KM`
	//
	// - 40GBase-LR ：
	//
	//   - `QSFP-40G-LR4-WDM1300,10KM`
	//
	//   - `QSFP-40G-ER4-WDM1300,40KM`
	//
	//   - `QSFP-40G-ZR4-WDM1300,80KM`
	//
	// - 100GBase-LR ：
	//
	//   - `QSFP28-100G-LR4-WDM1300,10KM`
	//
	//   - `QSFP28-100G-ER4-WDM1300,40KM`
	//
	//   - `QSFP28-100G-ZR4-WDM1300,80KM`
	//
	// example:
	//
	// SFP-GE-LR-SM1310,10KM
	OpticalModuleModel *string `json:"OpticalModuleModel,omitempty" xml:"OpticalModuleModel,omitempty"`
	// The geographical location of the on-premises data center.
	//
	// example:
	//
	// XX Street
	PeerLocation *string `json:"PeerLocation,omitempty" xml:"PeerLocation,omitempty"`
	// The number of ports. This parameter is required only when **HighReliableType*	- is set to **SingleApMultiConnection**. Valid values: 2 to 16.
	//
	// example:
	//
	// 2
	PortNum *int32 `json:"PortNum,omitempty" xml:"PortNum,omitempty"`
	// The region ID of the Express Connect circuit.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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

func (s *CreateHighReliablePhysicalConnectionRequestApList) GetOpticalModuleModel() *string {
	return s.OpticalModuleModel
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

func (s *CreateHighReliablePhysicalConnectionRequestApList) SetOpticalModuleModel(v string) *CreateHighReliablePhysicalConnectionRequestApList {
	s.OpticalModuleModel = &v
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length, and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, and must start with a letter or Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
