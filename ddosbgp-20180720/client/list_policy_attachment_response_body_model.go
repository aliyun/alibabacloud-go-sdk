// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentList(v []*ListPolicyAttachmentResponseBodyAttachmentList) *ListPolicyAttachmentResponseBody
	GetAttachmentList() []*ListPolicyAttachmentResponseBodyAttachmentList
	SetRequestId(v string) *ListPolicyAttachmentResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListPolicyAttachmentResponseBody
	GetTotal() *int64
}

type ListPolicyAttachmentResponseBody struct {
	// The records of attachments to the mitigation policy.
	AttachmentList []*ListPolicyAttachmentResponseBodyAttachmentList `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of attachments to the mitigation policy.
	//
	// example:
	//
	// 28
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPolicyAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponseBody) GetAttachmentList() []*ListPolicyAttachmentResponseBodyAttachmentList {
	return s.AttachmentList
}

func (s *ListPolicyAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyAttachmentResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListPolicyAttachmentResponseBody) SetAttachmentList(v []*ListPolicyAttachmentResponseBodyAttachmentList) *ListPolicyAttachmentResponseBody {
	s.AttachmentList = v
	return s
}

func (s *ListPolicyAttachmentResponseBody) SetRequestId(v string) *ListPolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyAttachmentResponseBody) SetTotal(v int64) *ListPolicyAttachmentResponseBody {
	s.Total = &v
	return s
}

func (s *ListPolicyAttachmentResponseBody) Validate() error {
	if s.AttachmentList != nil {
		for _, item := range s.AttachmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyAttachmentResponseBodyAttachmentList struct {
	// The IP address of the protected object.
	//
	// example:
	//
	// 147.139.183.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The UID of the member to which the IP address of the protected object belongs.
	//
	// example:
	//
	// 177699790631****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// 1b43f44e-65e1-411a-b0c0-d6c1********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test**
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The description of the policy.
	//
	// example:
	//
	// test
	PolicyRemark *string `json:"PolicyRemark,omitempty" xml:"PolicyRemark,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The port number of the protected object.
	//
	// example:
	//
	// 8*
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The protocol type of the protected object. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// example:
	//
	// udp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region to which the IP address of the protected object belongs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListPolicyAttachmentResponseBodyAttachmentList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentResponseBodyAttachmentList) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetIp() *string {
	return s.Ip
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetMemberUid() *string {
	return s.MemberUid
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPolicyRemark() *string {
	return s.PolicyRemark
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPort() *int32 {
	return s.Port
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetPortRange() *string {
	return s.PortRange
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) GetRegion() *string {
	return s.Region
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetIp(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Ip = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetMemberUid(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.MemberUid = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyId(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyName(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyName = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyRemark(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyRemark = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPolicyType(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPort(v int32) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Port = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetPortRange(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.PortRange = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetProtocol(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) SetRegion(v string) *ListPolicyAttachmentResponseBodyAttachmentList {
	s.Region = &v
	return s
}

func (s *ListPolicyAttachmentResponseBodyAttachmentList) Validate() error {
	return dara.Validate(s)
}
