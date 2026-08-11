// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachFromPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolListShrink(v string) *DetachFromPolicyShrinkRequest
	GetIpPortProtocolListShrink() *string
	SetPolicyType(v string) *DetachFromPolicyShrinkRequest
	GetPolicyType() *string
	SetPortVersion(v string) *DetachFromPolicyShrinkRequest
	GetPortVersion() *string
}

type DetachFromPolicyShrinkRequest struct {
	// The list of protected objects.
	//
	// This parameter is required.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The policy type. Valid values:
	//
	// - **default**: default mitigation policy.
	//
	// - **l3**: IP-specific mitigation policy.
	//
	// - **l4**: port-specific mitigation policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// l3
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The version of the port-specific mitigation policy. Valid values:
	//
	// - **Not specified**: dissociates the default surf anti-DDoS engine policy.
	//
	// - **2**: dissociates the new stream anti-DDoS engine policy.
	//
	// > Only port-specific mitigation policies support this parameter.
	//
	// example:
	//
	// 2
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s DetachFromPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachFromPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachFromPolicyShrinkRequest) GetIpPortProtocolListShrink() *string {
	return s.IpPortProtocolListShrink
}

func (s *DetachFromPolicyShrinkRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachFromPolicyShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *DetachFromPolicyShrinkRequest) SetIpPortProtocolListShrink(v string) *DetachFromPolicyShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *DetachFromPolicyShrinkRequest) SetPolicyType(v string) *DetachFromPolicyShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachFromPolicyShrinkRequest) SetPortVersion(v string) *DetachFromPolicyShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *DetachFromPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
