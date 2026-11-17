// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTrafficMirrorSessionRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTrafficMirrorSessionRequest
	GetDryRun() *bool
	SetEnabled(v bool) *CreateTrafficMirrorSessionRequest
	GetEnabled() *bool
	SetOwnerAccount(v string) *CreateTrafficMirrorSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTrafficMirrorSessionRequest
	GetOwnerId() *int64
	SetPacketLength(v int32) *CreateTrafficMirrorSessionRequest
	GetPacketLength() *int32
	SetPriority(v int32) *CreateTrafficMirrorSessionRequest
	GetPriority() *int32
	SetRegionId(v string) *CreateTrafficMirrorSessionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTrafficMirrorSessionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateTrafficMirrorSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTrafficMirrorSessionRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTrafficMirrorSessionRequestTag) *CreateTrafficMirrorSessionRequest
	GetTag() []*CreateTrafficMirrorSessionRequestTag
	SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorFilterId() *string
	SetTrafficMirrorSessionDescription(v string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorSessionDescription() *string
	SetTrafficMirrorSessionName(v string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorSessionName() *string
	SetTrafficMirrorSourceIds(v []*string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorSourceIds() []*string
	SetTrafficMirrorTargetId(v string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorTargetId() *string
	SetTrafficMirrorTargetType(v string) *CreateTrafficMirrorSessionRequest
	GetTrafficMirrorTargetType() *string
	SetVirtualNetworkId(v int32) *CreateTrafficMirrorSessionRequest
	GetVirtualNetworkId() *int32
}

type CreateTrafficMirrorSessionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the traffic mirror session. The system checks whether the required parameters are specified, the request format, and the limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the traffic mirror session is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable the traffic mirror session. Valid values:
	//
	// - **false*	- (default): disables the traffic mirror session.
	//
	// - **true**: enables the traffic mirror session.
	//
	// example:
	//
	// false
	Enabled      *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The length of the original mirrored packet, excluding the length of the VXLAN packet. Default value: **1500**. Valid values: **64*	- to **8500**. Unit: bytes.
	//
	// - The value of this parameter affects the length of the packet received by the traffic mirror destination. For more information, see the mirrored packet length and MTU limits in [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// - This parameter is available only in specific regions. For more information, see the description of the mirrored packet length parameter in [Create and manage traffic mirroring](https://help.aliyun.com/document_detail/207514.html).
	//
	// example:
	//
	// 1500
	PacketLength *int32 `json:"PacketLength,omitempty" xml:"PacketLength,omitempty"`
	// The priority of the traffic mirror session. Valid values: **1*	- to **32766**.
	//
	// A smaller value indicates a higher priority. The priorities of traffic mirror sessions created by the same Alibaba Cloud account in the same region must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the traffic mirror session. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about the regions that support traffic mirroring, see [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the traffic mirroring belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateTrafficMirrorSessionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The instance ID of the traffic mirror filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The description of the traffic mirror session.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a trafficmirrorsession.
	TrafficMirrorSessionDescription *string `json:"TrafficMirrorSessionDescription,omitempty" xml:"TrafficMirrorSessionDescription,omitempty"`
	// The name of the traffic mirror session.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	TrafficMirrorSessionName *string `json:"TrafficMirrorSessionName,omitempty" xml:"TrafficMirrorSessionName,omitempty"`
	// The instance ID of the traffic mirror source. Currently, network interface controllers (NICs) are supported as traffic mirror sources. The default value of **N*	- is **1**, which means that you can add only one traffic mirror source to a traffic mirror session.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-j6c2fp57q8rr47rp****
	TrafficMirrorSourceIds []*string `json:"TrafficMirrorSourceIds,omitempty" xml:"TrafficMirrorSourceIds,omitempty" type:"Repeated"`
	// The instance ID of the traffic mirror destination. Currently, network interface controllers (NICs) and internal-facing SLB instances are supported as traffic mirror destinations.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-j6c8znm5l1yt4sox****
	TrafficMirrorTargetId *string `json:"TrafficMirrorTargetId,omitempty" xml:"TrafficMirrorTargetId,omitempty"`
	// The type of the traffic mirror destination. Valid values:
	//
	// - **NetworkInterface**: network interface controllers (NICs).
	//
	// - **SLB**: internal-facing SLB instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// NetworkInterface
	TrafficMirrorTargetType *string `json:"TrafficMirrorTargetType,omitempty" xml:"TrafficMirrorTargetType,omitempty"`
	// The VXLAN network identifier (VNI) that is used to distinguish different mirrored data. Valid values: **0*	- to **16777215**.
	//
	// You can use the VNI to distinguish the mirrored data of different sessions at the traffic mirror destination. You can specify a custom VNI or let the system randomly allocate the value. If you want the system to randomly allocate the VNI, do not specify this parameter.
	//
	// example:
	//
	// 1
	VirtualNetworkId *int32 `json:"VirtualNetworkId,omitempty" xml:"VirtualNetworkId,omitempty"`
}

func (s CreateTrafficMirrorSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorSessionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTrafficMirrorSessionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTrafficMirrorSessionRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateTrafficMirrorSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTrafficMirrorSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTrafficMirrorSessionRequest) GetPacketLength() *int32 {
	return s.PacketLength
}

func (s *CreateTrafficMirrorSessionRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTrafficMirrorSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTrafficMirrorSessionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTrafficMirrorSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTrafficMirrorSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTrafficMirrorSessionRequest) GetTag() []*CreateTrafficMirrorSessionRequestTag {
	return s.Tag
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorSessionDescription() *string {
	return s.TrafficMirrorSessionDescription
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorSessionName() *string {
	return s.TrafficMirrorSessionName
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorSourceIds() []*string {
	return s.TrafficMirrorSourceIds
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorTargetId() *string {
	return s.TrafficMirrorTargetId
}

func (s *CreateTrafficMirrorSessionRequest) GetTrafficMirrorTargetType() *string {
	return s.TrafficMirrorTargetType
}

func (s *CreateTrafficMirrorSessionRequest) GetVirtualNetworkId() *int32 {
	return s.VirtualNetworkId
}

func (s *CreateTrafficMirrorSessionRequest) SetClientToken(v string) *CreateTrafficMirrorSessionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetDryRun(v bool) *CreateTrafficMirrorSessionRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetEnabled(v bool) *CreateTrafficMirrorSessionRequest {
	s.Enabled = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetOwnerAccount(v string) *CreateTrafficMirrorSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetOwnerId(v int64) *CreateTrafficMirrorSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetPacketLength(v int32) *CreateTrafficMirrorSessionRequest {
	s.PacketLength = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetPriority(v int32) *CreateTrafficMirrorSessionRequest {
	s.Priority = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetRegionId(v string) *CreateTrafficMirrorSessionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetResourceGroupId(v string) *CreateTrafficMirrorSessionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetResourceOwnerAccount(v string) *CreateTrafficMirrorSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetResourceOwnerId(v int64) *CreateTrafficMirrorSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTag(v []*CreateTrafficMirrorSessionRequestTag) *CreateTrafficMirrorSessionRequest {
	s.Tag = v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorSessionDescription(v string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorSessionDescription = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorSessionName(v string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorSessionName = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorSourceIds(v []*string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorSourceIds = v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorTargetId(v string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorTargetId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetTrafficMirrorTargetType(v string) *CreateTrafficMirrorSessionRequest {
	s.TrafficMirrorTargetType = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) SetVirtualNetworkId(v int32) *CreateTrafficMirrorSessionRequest {
	s.VirtualNetworkId = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequest) Validate() error {
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

type CreateTrafficMirrorSessionRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTrafficMirrorSessionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorSessionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorSessionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTrafficMirrorSessionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTrafficMirrorSessionRequestTag) SetKey(v string) *CreateTrafficMirrorSessionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequestTag) SetValue(v string) *CreateTrafficMirrorSessionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTrafficMirrorSessionRequestTag) Validate() error {
	return dara.Validate(s)
}
