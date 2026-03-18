// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolList(v []*ListPolicyAttachmentRequestIpPortProtocolList) *ListPolicyAttachmentRequest
	GetIpPortProtocolList() []*ListPolicyAttachmentRequestIpPortProtocolList
	SetPageNo(v int64) *ListPolicyAttachmentRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListPolicyAttachmentRequest
	GetPageSize() *int64
	SetPolicyId(v string) *ListPolicyAttachmentRequest
	GetPolicyId() *string
	SetPolicyType(v string) *ListPolicyAttachmentRequest
	GetPolicyType() *string
	SetPortVersion(v string) *ListPolicyAttachmentRequest
	GetPortVersion() *string
}

type ListPolicyAttachmentRequest struct {
	// The protected objects.
	IpPortProtocolList []*ListPolicyAttachmentRequestIpPortProtocolList `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// f38f6520-92b7-451e-b520-9ab3********
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The type of the policy. Valid values:
	//
	// 	- **default**: the default mitigation policies.
	//
	// 	- **l3**: IP-specific mitigation policies.
	//
	// 	- **l4**: port-specific mitigation policies.
	//
	// example:
	//
	// l3
	PolicyType  *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s ListPolicyAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentRequest) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentRequest) GetIpPortProtocolList() []*ListPolicyAttachmentRequestIpPortProtocolList {
	return s.IpPortProtocolList
}

func (s *ListPolicyAttachmentRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListPolicyAttachmentRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPolicyAttachmentRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListPolicyAttachmentRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListPolicyAttachmentRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *ListPolicyAttachmentRequest) SetIpPortProtocolList(v []*ListPolicyAttachmentRequestIpPortProtocolList) *ListPolicyAttachmentRequest {
	s.IpPortProtocolList = v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPageNo(v int64) *ListPolicyAttachmentRequest {
	s.PageNo = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPageSize(v int64) *ListPolicyAttachmentRequest {
	s.PageSize = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPolicyId(v string) *ListPolicyAttachmentRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPolicyType(v string) *ListPolicyAttachmentRequest {
	s.PolicyType = &v
	return s
}

func (s *ListPolicyAttachmentRequest) SetPortVersion(v string) *ListPolicyAttachmentRequest {
	s.PortVersion = &v
	return s
}

func (s *ListPolicyAttachmentRequest) Validate() error {
	if s.IpPortProtocolList != nil {
		for _, item := range s.IpPortProtocolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPolicyAttachmentRequestIpPortProtocolList struct {
	// The IP address of the protected object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47.118.172.***
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
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
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListPolicyAttachmentRequestIpPortProtocolList) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyAttachmentRequestIpPortProtocolList) GoString() string {
	return s.String()
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) GetIp() *string {
	return s.Ip
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) GetPort() *int32 {
	return s.Port
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) GetPortRange() *string {
	return s.PortRange
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) GetProtocol() *string {
	return s.Protocol
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetIp(v string) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Ip = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetPort(v int32) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Port = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetPortRange(v string) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.PortRange = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) SetProtocol(v string) *ListPolicyAttachmentRequestIpPortProtocolList {
	s.Protocol = &v
	return s
}

func (s *ListPolicyAttachmentRequestIpPortProtocolList) Validate() error {
	return dara.Validate(s)
}
