// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterRuleAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetClientToken() *string
	SetDestinationCidrBlock(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetDestinationCidrBlock() *string
	SetDestinationPortRange(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetDestinationPortRange() *string
	SetDryRun(v bool) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetOwnerId() *int64
	SetPriority(v int32) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetPriority() *int32
	SetProtocol(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetProtocol() *string
	SetRegionId(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetResourceOwnerId() *int64
	SetRuleAction(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetRuleAction() *string
	SetSourceCidrBlock(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetSourceCidrBlock() *string
	SetSourcePortRange(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetSourcePortRange() *string
	SetTrafficMirrorFilterRuleId(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest
	GetTrafficMirrorFilterRuleId() *string
}

type UpdateTrafficMirrorFilterRuleAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The destination CIDR block of the network traffic for the inbound or outbound rule to be modified.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The destination port range of the network traffic for the inbound or outbound rule to be modified.
	//
	// > If **Protocol*	- is set to **ICMP**, the port range cannot be modified.
	//
	// example:
	//
	// -1/-1
	DestinationPortRange *string `json:"DestinationPortRange,omitempty" xml:"DestinationPortRange,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the configuration of the inbound or outbound rule is modified.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the inbound or outbound rule to be modified. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Protocol Type of the network traffic to be mirrored by the inbound or outbound rule. Valid values:
	//
	// - **ALL**: all protocols.
	//
	// - **ICMP**: Internet Control Message Protocol.
	//
	// - **TCP**: Transmission Control Protocol.
	//
	// - **UDP**: User Datagram Protocol.
	//
	// example:
	//
	// ICMP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the traffic mirror.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query the most recent region list. For more information about regions that support traffic mirroring, see [Traffic mirroring overview](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The collection policy of the inbound or outbound rule to be modified. Valid values:
	//
	// - **accept**: collects network traffic.
	//
	// - **drop**: does not collect network traffic.
	//
	// example:
	//
	// accept
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The source CIDR block of the network traffic for the inbound or outbound rule to be modified.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrBlock *string `json:"SourceCidrBlock,omitempty" xml:"SourceCidrBlock,omitempty"`
	// The source port range of the network traffic for the inbound or outbound rule to be modified.
	//
	// > If **Protocol*	- is set to **ICMP**, the port range cannot be modified.
	//
	// example:
	//
	// 22/40
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The instance ID of the inbound or outbound rule of the traffic mirroring filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmr-j6c89rzmtd3hhdugq****
	TrafficMirrorFilterRuleId *string `json:"TrafficMirrorFilterRuleId,omitempty" xml:"TrafficMirrorFilterRuleId,omitempty"`
}

func (s UpdateTrafficMirrorFilterRuleAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterRuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetDestinationPortRange() *string {
	return s.DestinationPortRange
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetRuleAction() *string {
	return s.RuleAction
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetSourceCidrBlock() *string {
	return s.SourceCidrBlock
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) GetTrafficMirrorFilterRuleId() *string {
	return s.TrafficMirrorFilterRuleId
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetClientToken(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetDestinationCidrBlock(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetDestinationPortRange(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.DestinationPortRange = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetDryRun(v bool) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetOwnerAccount(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetOwnerId(v int64) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetPriority(v int32) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetProtocol(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetRegionId(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetResourceOwnerId(v int64) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetRuleAction(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.RuleAction = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetSourceCidrBlock(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.SourceCidrBlock = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetSourcePortRange(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.SourcePortRange = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) SetTrafficMirrorFilterRuleId(v string) *UpdateTrafficMirrorFilterRuleAttributeRequest {
	s.TrafficMirrorFilterRuleId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeRequest) Validate() error {
	return dara.Validate(s)
}
