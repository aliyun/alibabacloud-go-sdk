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
	// The description of the hosted connection.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Default value: 45104. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including required parameters, request syntax, and instance status. If the request fails the dry run, an error code is returned. If the request passes the dry run, `DRYRUN.SUCCESS` is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the hosted connection.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The payer for the hosted connection. Valid values:
	//
	// 	- **PayByPhysicalConnectionOwner**: The partner pays for the hosted connection.
	//
	// 	- **PayByVirtualPhysicalConnectionOwner**: The tenant pays for the hosted connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// PayByVirtualPhysicalConnectionOwner
	OrderMode *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	// The ID of the Express Connect circuit over which the hosted connection is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1ciz7ekd2grn1as****
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// The region ID of the hosted connection.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the hosted connection belongs.
	//
	// example:
	//
	// rg-aekzjty2chzuqky
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The bandwidth value of the hosted connection.
	//
	// Valid values: **50M**, **100M**, **200M**, **300M**, **400M**, **500M**, **1G**, **2G**, **5G**, **8G**, and **10G**.
	//
	// >  **2G**, **5G**, **8G**, and **10G*	- are unavailable by default. If you want to use these bandwidth values, contact your account manager.
	//
	// **M*	- indicates Mbit/s and **G*	- indicates Gbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50M
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The tags.
	Tag []*CreateVirtualPhysicalConnectionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// CBCE910E-D396-4944-8****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The virtual local area network (VLAN) ID of the hosted connection. Valid values: **0*	- to **2999**.
	//
	// 	- If the VLAN ID is set to **0**, it indicates that the switch port of the virtual border router (VBR) is a Layer 3 router interface instead of a VLAN interface. When a Layer 3 router interface is used, each Express Connect circuit corresponds to a VBR.
	//
	// 	- If the VLAN ID is set to a value from **1*	- to **2999**, the switch port of the VBR is a Layer 3 VLAN subinterface. When a Layer 3 VLAN subinterface is used, each VLAN ID corresponds to one VBR. In this case, the Express Connect circuit with which the VBR is associated can be used to connect to virtual private clouds (VPCs) that belong to different Alibaba Cloud accounts. VBRs in different VLANs are isolated from each other at Layer 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	VlanId *int64 `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The Alibaba Cloud account ID of the tenant.
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
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The key must start with a letter but cannot start with `aliyun` or `acs:`. The key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The value can be up to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The value must start with a letter but cannot start with `aliyun` or `acs:`. The value cannot contain `http://` or `https://`.
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
