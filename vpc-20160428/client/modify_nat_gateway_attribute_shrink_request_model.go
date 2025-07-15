// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetDescription() *string
	SetEipBindMode(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetEipBindMode() *string
	SetEnableSessionLog(v bool) *ModifyNatGatewayAttributeShrinkRequest
	GetEnableSessionLog() *bool
	SetIcmpReplyEnabled(v bool) *ModifyNatGatewayAttributeShrinkRequest
	GetIcmpReplyEnabled() *bool
	SetLogDeliveryShrink(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetLogDeliveryShrink() *string
	SetName(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetName() *string
	SetNatGatewayId(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNatGatewayAttributeShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNatGatewayAttributeShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNatGatewayAttributeShrinkRequest
	GetResourceOwnerId() *int64
}

type ModifyNatGatewayAttributeShrinkRequest struct {
	// The description of the NAT gateway.
	//
	// The description must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mode in which the NAT gateway is associated with an elastic IP address (EIP). You can leave this parameter empty. If you want to specify a value for this parameter, set the value to **NAT**, which indicates that the NAT gateway is associated with the EIP in NAT mode.
	//
	// **
	//
	// **Description**
	//
	// 	- If EipBindMode is set to MULTI_BINDED when the NAT gateway is created, you can change the value of this parameter from **MULTI_BINDED*	- to **NAT**. If EipBindMode is set to NAT when the NAT gateway is created, you cannot change the value of this parameter from **NAT*	- to **MULTI_BINDED**. For more information about **MULTI_BINDED**, see [CreateNatGateway](https://help.aliyun.com/document_detail/120219.html).
	//
	// 	- When the mode in which the NAT gateway is associated with an EIP is being changed, a transient connection that lasts a few seconds may occur. If the number of EIPs with which the NAT gateway is associated increases, the transient connection lasts longer. You can change the mode only for a NAT gateway that is associated with up to five EIPs. We recommend that you change the mode during off-peak hours.
	//
	// 	- After the mode is changed to **NAT**, the Internet NAT gateway is compatible with the IPv4 gateway. However, if you associate an EIP with the NAT gateway, the EIP occupies one private IP address on the vSwitch of the NAT gateway. Make sure that the vSwitch has sufficient private IP addresses. Otherwise, the EIP fails to be associated with the NAT gateway.
	//
	// example:
	//
	// NAT
	EipBindMode      *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	EnableSessionLog *bool   `json:"EnableSessionLog,omitempty" xml:"EnableSessionLog,omitempty"`
	// Specifies whether to enable the Internet Control Message Protocol (ICMP) non-retrieval feature. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	IcmpReplyEnabled  *bool   `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	LogDeliveryShrink *string `json:"LogDelivery,omitempty" xml:"LogDelivery,omitempty"`
	// The name of the NAT gateway.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nat123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-2ze0dcn4mq31qx2jc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s ModifyNatGatewayAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetEipBindMode() *string {
	return s.EipBindMode
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetEnableSessionLog() *bool {
	return s.EnableSessionLog
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetLogDeliveryShrink() *string {
	return s.LogDeliveryShrink
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNatGatewayAttributeShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetDescription(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetEipBindMode(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.EipBindMode = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetEnableSessionLog(v bool) *ModifyNatGatewayAttributeShrinkRequest {
	s.EnableSessionLog = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetIcmpReplyEnabled(v bool) *ModifyNatGatewayAttributeShrinkRequest {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetLogDeliveryShrink(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.LogDeliveryShrink = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetName(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetNatGatewayId(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetOwnerAccount(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetOwnerId(v int64) *ModifyNatGatewayAttributeShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetRegionId(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetResourceOwnerAccount(v string) *ModifyNatGatewayAttributeShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) SetResourceOwnerId(v int64) *ModifyNatGatewayAttributeShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNatGatewayAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
