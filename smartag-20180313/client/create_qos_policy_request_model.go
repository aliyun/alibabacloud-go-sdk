// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosPolicyRequest
	GetDescription() *string
	SetDestCidr(v string) *CreateQosPolicyRequest
	GetDestCidr() *string
	SetDestPortRange(v string) *CreateQosPolicyRequest
	GetDestPortRange() *string
	SetDpiGroupIds(v []*string) *CreateQosPolicyRequest
	GetDpiGroupIds() []*string
	SetDpiSignatureIds(v []*string) *CreateQosPolicyRequest
	GetDpiSignatureIds() []*string
	SetEndTime(v string) *CreateQosPolicyRequest
	GetEndTime() *string
	SetIpProtocol(v string) *CreateQosPolicyRequest
	GetIpProtocol() *string
	SetName(v string) *CreateQosPolicyRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateQosPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateQosPolicyRequest
	GetOwnerId() *int64
	SetPriority(v int32) *CreateQosPolicyRequest
	GetPriority() *int32
	SetQosId(v string) *CreateQosPolicyRequest
	GetQosId() *string
	SetRegionId(v string) *CreateQosPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateQosPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateQosPolicyRequest
	GetResourceOwnerId() *int64
	SetSourceCidr(v string) *CreateQosPolicyRequest
	GetSourceCidr() *string
	SetSourcePortRange(v string) *CreateQosPolicyRequest
	GetSourcePortRange() *string
	SetStartTime(v string) *CreateQosPolicyRequest
	GetStartTime() *string
}

type CreateQosPolicyRequest struct {
	// The description of the traffic classification rule.
	//
	// The description must be 1 to 512 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// Specify the value of this parameter in CIDR notation. Example: 192.168.10.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.20.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples:
	//
	// 	- **1/200**: a port range from 1 to 200
	//
	// 	- **80/80**: port 80
	//
	// 	- **-1/-1**: all ports
	//
	// This parameter is required.
	//
	// example:
	//
	// 80/80
	DestPortRange *string `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	// example:
	//
	// 20
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The time when the traffic classification rule expires.
	//
	// Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ss+0800` format. The time must be in UTC+8.
	//
	// example:
	//
	// 2022-09-14T16:41:33+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the protocol that applies to the traffic classification rule.
	//
	// The supported protocols provided in this topic are for reference only. The actual protocols in the console shall prevail.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the traffic classification rule.
	//
	// The name must be 2 to 100 characters in length, and can contain letters, digits, underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The priority of the traffic throttling policy to which the traffic classification rule belongs.
	//
	// Valid values: **1 to 3**. A smaller value indicates a higher priority.
	//
	// This parameter is required.
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
	// The ID of the region to which the QoS policy belongs.
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
	// Specify the value of this parameter in CIDR notation. Example: 192.168.1.0/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.10.10.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// Valid values: **1*	- to **65535*	- and **-1**.
	//
	// Examples:
	//
	// 	- **1/200**: a port range from 1 to 200
	//
	// 	- **80/80**: port 80
	//
	// 	- **-1/-1**: all ports
	//
	// This parameter is required.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The time when the traffic classification rule takes effect.
	//
	// Specify the time in the ISO 8601 standard in the `YYYY-MM-DDThh:mm:ss+0800` format. The time must be in UTC+8.
	//
	// example:
	//
	// 2022-07-14T16:41:33+0800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQosPolicyRequest) GetDestCidr() *string {
	return s.DestCidr
}

func (s *CreateQosPolicyRequest) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *CreateQosPolicyRequest) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *CreateQosPolicyRequest) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *CreateQosPolicyRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateQosPolicyRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateQosPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateQosPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateQosPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateQosPolicyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateQosPolicyRequest) GetQosId() *string {
	return s.QosId
}

func (s *CreateQosPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateQosPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateQosPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateQosPolicyRequest) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *CreateQosPolicyRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateQosPolicyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateQosPolicyRequest) SetDescription(v string) *CreateQosPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateQosPolicyRequest) SetDestCidr(v string) *CreateQosPolicyRequest {
	s.DestCidr = &v
	return s
}

func (s *CreateQosPolicyRequest) SetDestPortRange(v string) *CreateQosPolicyRequest {
	s.DestPortRange = &v
	return s
}

func (s *CreateQosPolicyRequest) SetDpiGroupIds(v []*string) *CreateQosPolicyRequest {
	s.DpiGroupIds = v
	return s
}

func (s *CreateQosPolicyRequest) SetDpiSignatureIds(v []*string) *CreateQosPolicyRequest {
	s.DpiSignatureIds = v
	return s
}

func (s *CreateQosPolicyRequest) SetEndTime(v string) *CreateQosPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateQosPolicyRequest) SetIpProtocol(v string) *CreateQosPolicyRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateQosPolicyRequest) SetName(v string) *CreateQosPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateQosPolicyRequest) SetOwnerAccount(v string) *CreateQosPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateQosPolicyRequest) SetOwnerId(v int64) *CreateQosPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetPriority(v int32) *CreateQosPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateQosPolicyRequest) SetQosId(v string) *CreateQosPolicyRequest {
	s.QosId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetRegionId(v string) *CreateQosPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetResourceOwnerAccount(v string) *CreateQosPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateQosPolicyRequest) SetResourceOwnerId(v int64) *CreateQosPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetSourceCidr(v string) *CreateQosPolicyRequest {
	s.SourceCidr = &v
	return s
}

func (s *CreateQosPolicyRequest) SetSourcePortRange(v string) *CreateQosPolicyRequest {
	s.SourcePortRange = &v
	return s
}

func (s *CreateQosPolicyRequest) SetStartTime(v string) *CreateQosPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
