// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNatGatewayAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyNatGatewayAttributeRequest
	GetDescription() *string
	SetEipBindMode(v string) *ModifyNatGatewayAttributeRequest
	GetEipBindMode() *string
	SetEnableSessionLog(v bool) *ModifyNatGatewayAttributeRequest
	GetEnableSessionLog() *bool
	SetIcmpReplyEnabled(v bool) *ModifyNatGatewayAttributeRequest
	GetIcmpReplyEnabled() *bool
	SetLogDelivery(v *ModifyNatGatewayAttributeRequestLogDelivery) *ModifyNatGatewayAttributeRequest
	GetLogDelivery() *ModifyNatGatewayAttributeRequestLogDelivery
	SetName(v string) *ModifyNatGatewayAttributeRequest
	GetName() *string
	SetNatGatewayId(v string) *ModifyNatGatewayAttributeRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *ModifyNatGatewayAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyNatGatewayAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNatGatewayAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNatGatewayAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyNatGatewayAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyNatGatewayAttributeRequest struct {
	// The description of the NAT gateway that you want to modify.
	//
	// The description must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The EIP attach pattern of the NAT gateway. Valid values: The value can be empty. If the value is not empty, only **NAT*	- is allowed, which indicates that the attach pattern is EIP Normal pattern.
	//
	//
	// > - You can only change the pattern from **MULTI_BINDED*	- to **NAT**. You cannot change the pattern from **NAT*	- to **MULTI_BINDED**. For more information about the **MULTI_BINDED*	- pattern, see [CreateNatGateway](https://help.aliyun.com/document_detail/120219.html).
	//
	// - During the EIP attach pattern switchover procedure, network connectivity may experience second-level transient connections (the transient connection duration increases as the number of EIPs increases. Currently, configuration changes are supported for NAT gateways with up to 5 EIPs attached). Execute the switchover during off-peak hours.
	//
	// - After the EIP attach pattern is changed to **NAT**, the Internet NAT gateway is compatible with the IPv4 gateway. However, attaching a public EIP occupies a private IP in the vSwitch where the NAT gateway resides. Make sure that sufficient private IP addresses are available in the vSwitch. If no available idle private IP addresses exist in the vSwitch, new EIPs cannot be attached.
	//
	// example:
	//
	// NAT
	EipBindMode *string `json:"EipBindMode,omitempty" xml:"EipBindMode,omitempty"`
	// Specifies whether to enable session logging. Valid values:
	//
	// - **true**: Session logging is enabled.
	//
	// - **false**: Session logging is disabled.
	//
	// example:
	//
	// true
	EnableSessionLog *bool `json:"EnableSessionLog,omitempty" xml:"EnableSessionLog,omitempty"`
	// Specifies whether to enable ICMP echo reply. Valid values:
	//
	// - **true*	- (default): Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	IcmpReplyEnabled *bool `json:"IcmpReplyEnabled,omitempty" xml:"IcmpReplyEnabled,omitempty"`
	// The session log configuration.
	LogDelivery *ModifyNatGatewayAttributeRequestLogDelivery `json:"LogDelivery,omitempty" xml:"LogDelivery,omitempty" type:"Struct"`
	// The name of the NAT gateway that you want to modify.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// nat123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the NAT gateway that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-2ze0dcn4mq31qx2jc****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the NAT gateway that you want to modify.
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

func (s ModifyNatGatewayAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyNatGatewayAttributeRequest) GetEipBindMode() *string {
	return s.EipBindMode
}

func (s *ModifyNatGatewayAttributeRequest) GetEnableSessionLog() *bool {
	return s.EnableSessionLog
}

func (s *ModifyNatGatewayAttributeRequest) GetIcmpReplyEnabled() *bool {
	return s.IcmpReplyEnabled
}

func (s *ModifyNatGatewayAttributeRequest) GetLogDelivery() *ModifyNatGatewayAttributeRequestLogDelivery {
	return s.LogDelivery
}

func (s *ModifyNatGatewayAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyNatGatewayAttributeRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ModifyNatGatewayAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyNatGatewayAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNatGatewayAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNatGatewayAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNatGatewayAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyNatGatewayAttributeRequest) SetDescription(v string) *ModifyNatGatewayAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetEipBindMode(v string) *ModifyNatGatewayAttributeRequest {
	s.EipBindMode = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetEnableSessionLog(v bool) *ModifyNatGatewayAttributeRequest {
	s.EnableSessionLog = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetIcmpReplyEnabled(v bool) *ModifyNatGatewayAttributeRequest {
	s.IcmpReplyEnabled = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetLogDelivery(v *ModifyNatGatewayAttributeRequestLogDelivery) *ModifyNatGatewayAttributeRequest {
	s.LogDelivery = v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetName(v string) *ModifyNatGatewayAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetNatGatewayId(v string) *ModifyNatGatewayAttributeRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetOwnerAccount(v string) *ModifyNatGatewayAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetOwnerId(v int64) *ModifyNatGatewayAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetRegionId(v string) *ModifyNatGatewayAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetResourceOwnerAccount(v string) *ModifyNatGatewayAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) SetResourceOwnerId(v int64) *ModifyNatGatewayAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequest) Validate() error {
	if s.LogDelivery != nil {
		if err := s.LogDelivery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyNatGatewayAttributeRequestLogDelivery struct {
	// The type of session log delivery. Valid values: **sls**, which indicates Simple Log Service.
	//
	// example:
	//
	// sls
	LogDeliveryType *string `json:"LogDeliveryType,omitempty" xml:"LogDeliveryType,omitempty"`
	// The destination address for session log delivery. Value: acs:log:${regionName}:${projectOwnerAliUid}:project/${projectName}/logstore/${logstoreName}
	//
	// example:
	//
	// acs:log:cn-hangzhou:0000:project/nat_session_log_project/logstore/session_log_test
	LogDestination *string `json:"LogDestination,omitempty" xml:"LogDestination,omitempty"`
}

func (s ModifyNatGatewayAttributeRequestLogDelivery) String() string {
	return dara.Prettify(s)
}

func (s ModifyNatGatewayAttributeRequestLogDelivery) GoString() string {
	return s.String()
}

func (s *ModifyNatGatewayAttributeRequestLogDelivery) GetLogDeliveryType() *string {
	return s.LogDeliveryType
}

func (s *ModifyNatGatewayAttributeRequestLogDelivery) GetLogDestination() *string {
	return s.LogDestination
}

func (s *ModifyNatGatewayAttributeRequestLogDelivery) SetLogDeliveryType(v string) *ModifyNatGatewayAttributeRequestLogDelivery {
	s.LogDeliveryType = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequestLogDelivery) SetLogDestination(v string) *ModifyNatGatewayAttributeRequestLogDelivery {
	s.LogDestination = &v
	return s
}

func (s *ModifyNatGatewayAttributeRequestLogDelivery) Validate() error {
	return dara.Validate(s)
}
