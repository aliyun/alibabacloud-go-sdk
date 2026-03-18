// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachToPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpPortProtocolListShrink(v string) *AttachToPolicyShrinkRequest
	GetIpPortProtocolListShrink() *string
	SetPolicyId(v string) *AttachToPolicyShrinkRequest
	GetPolicyId() *string
	SetPortVersion(v string) *AttachToPolicyShrinkRequest
	GetPortVersion() *string
}

type AttachToPolicyShrinkRequest struct {
	// The protected objects.
	//
	// This parameter is required.
	IpPortProtocolListShrink *string `json:"IpPortProtocolList,omitempty" xml:"IpPortProtocolList,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cd8b4d70-e4e0-413a-b390-e71d********
	PolicyId    *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	PortVersion *string `json:"PortVersion,omitempty" xml:"PortVersion,omitempty"`
}

func (s AttachToPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachToPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachToPolicyShrinkRequest) GetIpPortProtocolListShrink() *string {
	return s.IpPortProtocolListShrink
}

func (s *AttachToPolicyShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachToPolicyShrinkRequest) GetPortVersion() *string {
	return s.PortVersion
}

func (s *AttachToPolicyShrinkRequest) SetIpPortProtocolListShrink(v string) *AttachToPolicyShrinkRequest {
	s.IpPortProtocolListShrink = &v
	return s
}

func (s *AttachToPolicyShrinkRequest) SetPolicyId(v string) *AttachToPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachToPolicyShrinkRequest) SetPortVersion(v string) *AttachToPolicyShrinkRequest {
	s.PortVersion = &v
	return s
}

func (s *AttachToPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
