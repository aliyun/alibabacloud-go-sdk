// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosPolicyResponseBody
	GetDescription() *string
	SetDestCidr(v string) *CreateQosPolicyResponseBody
	GetDestCidr() *string
	SetDestPortRange(v string) *CreateQosPolicyResponseBody
	GetDestPortRange() *string
	SetDpiGroupIds(v *CreateQosPolicyResponseBodyDpiGroupIds) *CreateQosPolicyResponseBody
	GetDpiGroupIds() *CreateQosPolicyResponseBodyDpiGroupIds
	SetDpiSignatureIds(v *CreateQosPolicyResponseBodyDpiSignatureIds) *CreateQosPolicyResponseBody
	GetDpiSignatureIds() *CreateQosPolicyResponseBodyDpiSignatureIds
	SetEndTime(v string) *CreateQosPolicyResponseBody
	GetEndTime() *string
	SetIpProtocol(v string) *CreateQosPolicyResponseBody
	GetIpProtocol() *string
	SetName(v string) *CreateQosPolicyResponseBody
	GetName() *string
	SetPriority(v int32) *CreateQosPolicyResponseBody
	GetPriority() *int32
	SetQosId(v string) *CreateQosPolicyResponseBody
	GetQosId() *string
	SetQosPolicyId(v string) *CreateQosPolicyResponseBody
	GetQosPolicyId() *string
	SetRequestId(v string) *CreateQosPolicyResponseBody
	GetRequestId() *string
	SetSourceCidr(v string) *CreateQosPolicyResponseBody
	GetSourceCidr() *string
	SetSourcePortRange(v string) *CreateQosPolicyResponseBody
	GetSourcePortRange() *string
	SetStartTime(v string) *CreateQosPolicyResponseBody
	GetStartTime() *string
}

type CreateQosPolicyResponseBody struct {
	// The description of the traffic classification rule.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 10.10.10.0/24
	DestCidr *string `json:"DestCidr,omitempty" xml:"DestCidr,omitempty"`
	// The destination port range.
	//
	// example:
	//
	// 80/80
	DestPortRange   *string                                     `json:"DestPortRange,omitempty" xml:"DestPortRange,omitempty"`
	DpiGroupIds     *CreateQosPolicyResponseBodyDpiGroupIds     `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Struct"`
	DpiSignatureIds *CreateQosPolicyResponseBodyDpiSignatureIds `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Struct"`
	// The time when the traffic classification rule expires.
	//
	// example:
	//
	// 2022-09-14T16:41:33+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The protocol that applies to the traffic classification rule.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The name of the traffic classification rule.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The priority of the traffic throttling policy to which the traffic classification rule belongs.
	//
	// example:
	//
	// 3
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-xitd8690ucu8ro****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the traffic classification rule.
	//
	// example:
	//
	// qospy-xhwhyuo43l********
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 97862812-2C7E-4D25-B0D5-B26DAC7FA293
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 10.10.10.0/24
	SourceCidr *string `json:"SourceCidr,omitempty" xml:"SourceCidr,omitempty"`
	// The source port range.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The time when the traffic classification rule takes effect.
	//
	// example:
	//
	// 2022-07-14T16:41:33+0800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateQosPolicyResponseBody) GetDestCidr() *string {
	return s.DestCidr
}

func (s *CreateQosPolicyResponseBody) GetDestPortRange() *string {
	return s.DestPortRange
}

func (s *CreateQosPolicyResponseBody) GetDpiGroupIds() *CreateQosPolicyResponseBodyDpiGroupIds {
	return s.DpiGroupIds
}

func (s *CreateQosPolicyResponseBody) GetDpiSignatureIds() *CreateQosPolicyResponseBodyDpiSignatureIds {
	return s.DpiSignatureIds
}

func (s *CreateQosPolicyResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateQosPolicyResponseBody) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateQosPolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateQosPolicyResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateQosPolicyResponseBody) GetQosId() *string {
	return s.QosId
}

func (s *CreateQosPolicyResponseBody) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *CreateQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQosPolicyResponseBody) GetSourceCidr() *string {
	return s.SourceCidr
}

func (s *CreateQosPolicyResponseBody) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *CreateQosPolicyResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateQosPolicyResponseBody) SetDescription(v string) *CreateQosPolicyResponseBody {
	s.Description = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetDestCidr(v string) *CreateQosPolicyResponseBody {
	s.DestCidr = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetDestPortRange(v string) *CreateQosPolicyResponseBody {
	s.DestPortRange = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetDpiGroupIds(v *CreateQosPolicyResponseBodyDpiGroupIds) *CreateQosPolicyResponseBody {
	s.DpiGroupIds = v
	return s
}

func (s *CreateQosPolicyResponseBody) SetDpiSignatureIds(v *CreateQosPolicyResponseBodyDpiSignatureIds) *CreateQosPolicyResponseBody {
	s.DpiSignatureIds = v
	return s
}

func (s *CreateQosPolicyResponseBody) SetEndTime(v string) *CreateQosPolicyResponseBody {
	s.EndTime = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetIpProtocol(v string) *CreateQosPolicyResponseBody {
	s.IpProtocol = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetName(v string) *CreateQosPolicyResponseBody {
	s.Name = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetPriority(v int32) *CreateQosPolicyResponseBody {
	s.Priority = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetQosId(v string) *CreateQosPolicyResponseBody {
	s.QosId = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetQosPolicyId(v string) *CreateQosPolicyResponseBody {
	s.QosPolicyId = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetRequestId(v string) *CreateQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetSourceCidr(v string) *CreateQosPolicyResponseBody {
	s.SourceCidr = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetSourcePortRange(v string) *CreateQosPolicyResponseBody {
	s.SourcePortRange = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetStartTime(v string) *CreateQosPolicyResponseBody {
	s.StartTime = &v
	return s
}

func (s *CreateQosPolicyResponseBody) Validate() error {
	if s.DpiGroupIds != nil {
		if err := s.DpiGroupIds.Validate(); err != nil {
			return err
		}
	}
	if s.DpiSignatureIds != nil {
		if err := s.DpiSignatureIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQosPolicyResponseBodyDpiGroupIds struct {
	DpiGroupId []*string `json:"DpiGroupId,omitempty" xml:"DpiGroupId,omitempty" type:"Repeated"`
}

func (s CreateQosPolicyResponseBodyDpiGroupIds) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyResponseBodyDpiGroupIds) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyResponseBodyDpiGroupIds) GetDpiGroupId() []*string {
	return s.DpiGroupId
}

func (s *CreateQosPolicyResponseBodyDpiGroupIds) SetDpiGroupId(v []*string) *CreateQosPolicyResponseBodyDpiGroupIds {
	s.DpiGroupId = v
	return s
}

func (s *CreateQosPolicyResponseBodyDpiGroupIds) Validate() error {
	return dara.Validate(s)
}

type CreateQosPolicyResponseBodyDpiSignatureIds struct {
	DpiSignatureId []*string `json:"DpiSignatureId,omitempty" xml:"DpiSignatureId,omitempty" type:"Repeated"`
}

func (s CreateQosPolicyResponseBodyDpiSignatureIds) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyResponseBodyDpiSignatureIds) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyResponseBodyDpiSignatureIds) GetDpiSignatureId() []*string {
	return s.DpiSignatureId
}

func (s *CreateQosPolicyResponseBodyDpiSignatureIds) SetDpiSignatureId(v []*string) *CreateQosPolicyResponseBodyDpiSignatureIds {
	s.DpiSignatureId = v
	return s
}

func (s *CreateQosPolicyResponseBodyDpiSignatureIds) Validate() error {
	return dara.Validate(s)
}
