// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateVirtualPhysicalConnectionRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVirtualPhysicalConnectionRequest
	GetDryRun() *bool
	SetName(v string) *CreateVirtualPhysicalConnectionRequest
	GetName() *string
	SetOrderMode(v string) *CreateVirtualPhysicalConnectionRequest
	GetOrderMode() *string
	SetPhysicalConnectionId(v string) *CreateVirtualPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *CreateVirtualPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVirtualPhysicalConnectionRequest
	GetResourceGroupId() *string
	SetSpec(v string) *CreateVirtualPhysicalConnectionRequest
	GetSpec() *string
	SetTag(v []*CreateVirtualPhysicalConnectionRequestTag) *CreateVirtualPhysicalConnectionRequest
	GetTag() []*CreateVirtualPhysicalConnectionRequestTag
	SetToken(v string) *CreateVirtualPhysicalConnectionRequest
	GetToken() *string
	SetVlanId(v int64) *CreateVirtualPhysicalConnectionRequest
	GetVlanId() *int64
	SetVpconnAliUid(v int64) *CreateVirtualPhysicalConnectionRequest
	GetVpconnAliUid() *int64
}

type CreateVirtualPhysicalConnectionRequest struct {
	// The description of the shared Express Connect circuits.
	//
	// The description must be 2 to 256 characters in length and must start with a letter or a Chinese character, but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the shared Express Connect circuits. The system checks required parameters, request format, and instance status. If the check fails, the corresponding error is returned. If the check passes, `DRYRUN.SUCCESS` is returned.
	//
	// - **false*	- (default): sends a Normal request. After the check passes, the shared Express Connect circuits are created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the shared Express Connect circuits.
	//
	// The name must be 2 to 128 characters in length and must start with a letter or a Chinese character. It can contain digits, underscores (_), and hyphens (-), but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The payer of the shared Express Connect circuits. Valid values:
	//
	// - **PayByPhysicalConnectionOwner**: The partner pays.
	//
	// - **PayByVirtualPhysicalConnectionOwner**: The tenant pays.
	//
	// > Default value: PayByVirtualPhysicalConnectionOwner (tenant pays).
	//
	// This parameter is required.
	//
	// example:
	//
	// PayByVirtualPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	// The ID of the Express Connect circuit associated with the shared Express Connect circuits.
	//
	// > The Express Connect circuit must be in the Enabled state and must be an Express Connect circuit (shared Express Connect circuits IDs are not supported). Otherwise, ResourceNotFound.PhysicalConnectionId or IncorrectStatus.PhysicalConnection is returned. You can invoke DescribePhysicalConnections to query the status of the Express Connect circuit. The caller must be the account (partner) that owns the Express Connect circuit.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the shared Express Connect circuits.
	//
	// You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the shared Express Connect circuits belong.
	//
	// example:
	//
	// rg-aekzjty2chzuqky
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth value of the shared Express Connect circuits.
	//
	// Valid values: **50M**, **100M**, **200M**, **300M**, **400M**, **500M**, **1G**, **2G**, **5G**, **8G**, and **10G**.
	//
	// <props="china">
	//
	// > The bandwidth values **2G**, **5G**, **8G**, and **10G*	- are not available by default. To use these values, contact your account manager.
	//
	// <props="intl">
	//
	// > The bandwidth values **2G**, **5G**, **8G**, and **10G*	- are not available by default. To use these values, contact your account manager.
	//
	// Unit: **M*	- indicates Mbit/s and **G*	- indicates Gbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The list of tags.
	Tag []*CreateVirtualPhysicalConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure uniqueness across different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The VLAN ID of the shared Express Connect circuits. Valid values: **0*	- to **2999**.
	//
	// - If the VLAN ID is **0**, the physical switch port of the Virtual Border Router (VBR) does not use VLAN mode but uses Layer 3 routing interface mode. In Layer 3 routing interface mode, each Express Connect circuit corresponds to one VBR.
	//
	// - If the VLAN ID is **1*	- to **2999**, the physical switch port of the VBR uses VLAN-based Layer 3 sub-interfaces. In Layer 3 sub-interface mode, each VLAN ID corresponds to one VBR. In this case, the Express Connect circuit of the VBR can connect to VPCs under multiple accounts. VBRs in different VLANs have network isolation at Layer 2 and cannot communicate with each other.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	VlanId *int64 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The ID of the tenant\\"s Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1210123456123456
	VpconnAliUid *int64 `json:"VpconnAliUid,omitempty" xml:"VpconnAliUid,omitempty"`
}

func (s CreateVirtualPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualPhysicalConnectionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVirtualPhysicalConnectionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVirtualPhysicalConnectionRequest) GetName() *string {
	return s.Name
}

func (s *CreateVirtualPhysicalConnectionRequest) GetOrderMode() *string {
	return s.OrderMode
}

func (s *CreateVirtualPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *CreateVirtualPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirtualPhysicalConnectionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVirtualPhysicalConnectionRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateVirtualPhysicalConnectionRequest) GetTag() []*CreateVirtualPhysicalConnectionRequestTag {
	return s.Tag
}

func (s *CreateVirtualPhysicalConnectionRequest) GetToken() *string {
	return s.Token
}

func (s *CreateVirtualPhysicalConnectionRequest) GetVlanId() *int64 {
	return s.VlanId
}

func (s *CreateVirtualPhysicalConnectionRequest) GetVpconnAliUid() *int64 {
	return s.VpconnAliUid
}

func (s *CreateVirtualPhysicalConnectionRequest) SetDescription(v string) *CreateVirtualPhysicalConnectionRequest {
	s.Description = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetDryRun(v bool) *CreateVirtualPhysicalConnectionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetName(v string) *CreateVirtualPhysicalConnectionRequest {
	s.Name = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetOrderMode(v string) *CreateVirtualPhysicalConnectionRequest {
	s.OrderMode = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *CreateVirtualPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetRegionId(v string) *CreateVirtualPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetResourceGroupId(v string) *CreateVirtualPhysicalConnectionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetSpec(v string) *CreateVirtualPhysicalConnectionRequest {
	s.Spec = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetTag(v []*CreateVirtualPhysicalConnectionRequestTag) *CreateVirtualPhysicalConnectionRequest {
	s.Tag = v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetToken(v string) *CreateVirtualPhysicalConnectionRequest {
	s.Token = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetVlanId(v int64) *CreateVirtualPhysicalConnectionRequest {
	s.VlanId = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) SetVpconnAliUid(v int64) *CreateVirtualPhysicalConnectionRequest {
	s.VpconnAliUid = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequest) Validate() error {
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

type CreateVirtualPhysicalConnectionRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVirtualPhysicalConnectionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualPhysicalConnectionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVirtualPhysicalConnectionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVirtualPhysicalConnectionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVirtualPhysicalConnectionRequestTag) SetKey(v string) *CreateVirtualPhysicalConnectionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequestTag) SetValue(v string) *CreateVirtualPhysicalConnectionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionRequestTag) Validate() error {
	return dara.Validate(s)
}
