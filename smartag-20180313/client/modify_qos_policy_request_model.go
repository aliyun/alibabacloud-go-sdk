// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyQosPolicyRequest
	GetDescription() *string
	SetDestCidr(v string) *ModifyQosPolicyRequest
	GetDestCidr() *string
	SetDestPortRange(v string) *ModifyQosPolicyRequest
	GetDestPortRange() *string
	SetDpiGroupIds(v []*string) *ModifyQosPolicyRequest
	GetDpiGroupIds() []*string
	SetDpiSignatureIds(v []*string) *ModifyQosPolicyRequest
	GetDpiSignatureIds() []*string
	SetEndTime(v string) *ModifyQosPolicyRequest
	GetEndTime() *string
	SetIpProtocol(v string) *ModifyQosPolicyRequest
	GetIpProtocol() *string
	SetName(v string) *ModifyQosPolicyRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyQosPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyQosPolicyRequest
	GetOwnerId() *int64
	SetPriority(v int32) *ModifyQosPolicyRequest
	GetPriority() *int32
	SetQosId(v string) *ModifyQosPolicyRequest
	GetQosId() *string
	SetQosPolicyId(v string) *ModifyQosPolicyRequest
	GetQosPolicyId() *string
	SetRegionId(v string) *ModifyQosPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyQosPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyQosPolicyRequest
	GetResourceOwnerId() *int64
	SetSourceCidr(v string) *ModifyQosPolicyRequest
	GetSourceCidr() *string
	SetSourcePortRange(v string) *ModifyQosPolicyRequest
	GetSourcePortRange() *string
	SetStartTime(v string) *ModifyQosPolicyRequest
	GetStartTime() *string
}

type ModifyQosPolicyRequest struct {
	// The description of the stream classification rule.
	//
	// The description must be 1 to 512 characters in length. It must start with a letter and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// The destination CIDR block must be in CIDR format. Example: 192.168.10.0/24.
	//
	// example:
	//
	// 10.10.20.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Examples of how to specify a port range:
	//
	// - **1/200**: ports 1 through 200.
	//
	// - **80/80**: port 80.
	//
	// - **-1/-1**: all ports.
	//
	// example:
	//
	// 300/400
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// The list of application group IDs.
	//
	// example:
	//
	// 20
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// The list of application IDs.
	//
	// example:
	//
	// 1
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The time when the stream classification rule expires.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC+8. Format: YYYY-MM-DDThh:mm:ss+0800.
	//
	// example:
	//
	// 2019-09-14T16:41:33+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The protocol.
	//
	// For a list of supported protocols, see the information in the console.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the stream classification rule.
	//
	// The name must be 2 to 100 characters in length. It must start with a letter and can contain digits, underscores (_), and hyphens (-).
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the throttling rule to which the stream classification rule belongs.
	//
	// Valid values: 1 to **3**. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the stream classification rule in the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qospy-427m9fo6wkh****
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	// The ID of the region where the QoS policy is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source CIDR block.
	//
	// The source CIDR block must be in CIDR format. Example: 192.168.1.0/24.
	//
	// example:
	//
	// 10.10.10.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// Valid values: **-1*	- or **1*	- to **65535**.
	//
	// Examples of how to specify a port range:
	//
	// - **1/200**: ports 1 through 200.
	//
	// - **80/80**: port 80.
	//
	// - **-1/-1**: all ports.
	//
	// example:
	//
	// 1/200
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The time when the stream classification rule takes effect.
	//
	// Specify the time in the ISO 8601 standard. The time must be in UTC+8. Format: YYYY-MM-DDThh:mm:ss+0800.
	//
	// example:
	//
	// 2019-07-14T16:41:33+0800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyQosPolicyRequest) GetDestCidr() *string {
	return s.DestCidr
}

func (s *ModifyQosPolicyRequest) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *ModifyQosPolicyRequest) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *ModifyQosPolicyRequest) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *ModifyQosPolicyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyQosPolicyRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyQosPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyQosPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyQosPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyQosPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifyQosPolicyRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyQosPolicyRequest) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *ModifyQosPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyQosPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyQosPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyQosPolicyRequest) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *ModifyQosPolicyRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifyQosPolicyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyQosPolicyRequest) SetDescription(v string) *ModifyQosPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetDestCidr(v string) *ModifyQosPolicyRequest {
	s.DestCidr = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetDestPortRange(v string) *ModifyQosPolicyRequest {
	s.DestPortRange = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetDpiGroupIds(v []*string) *ModifyQosPolicyRequest {
	s.DpiGroupIds = v
	return s
}

func (s *ModifyQosPolicyRequest) SetDpiSignatureIds(v []*string) *ModifyQosPolicyRequest {
	s.DpiSignatureIds = v
	return s
}

func (s *ModifyQosPolicyRequest) SetEndTime(v string) *ModifyQosPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetIpProtocol(v string) *ModifyQosPolicyRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetName(v string) *ModifyQosPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetOwnerAccount(v string) *ModifyQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetOwnerId(v int64) *ModifyQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetPriority(v int32) *ModifyQosPolicyRequest {
	s.Priority = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetQosId(v string) *ModifyQosPolicyRequest {
	s.QosId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetQosPolicyId(v string) *ModifyQosPolicyRequest {
	s.QosPolicyId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetRegionId(v string) *ModifyQosPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetResourceOwnerAccount(v string) *ModifyQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetResourceOwnerId(v int64) *ModifyQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetSourceCidr(v string) *ModifyQosPolicyRequest {
	s.SourceCidr = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetSourcePortRange(v string) *ModifyQosPolicyRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetStartTime(v string) *ModifyQosPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
