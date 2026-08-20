// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiPolicyReference interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyAttachmentId(v string) *HttpApiPolicyReference
	GetPolicyAttachmentId() *string
	SetPolicyId(v string) *HttpApiPolicyReference
	GetPolicyId() *string
}

type HttpApiPolicyReference struct {
	// The policy attachment ID.
	//
	// example:
	//
	// pa-3c8ayyy
	PolicyAttachmentId *string `json:"policyAttachmentId,omitempty" xml:"policyAttachmentId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pol-9f2exxx
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s HttpApiPolicyReference) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyReference) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyReference) GetPolicyAttachmentId() *string {
	return s.PolicyAttachmentId
}

func (s *HttpApiPolicyReference) GetPolicyId() *string {
	return s.PolicyId
}

func (s *HttpApiPolicyReference) SetPolicyAttachmentId(v string) *HttpApiPolicyReference {
	s.PolicyAttachmentId = &v
	return s
}

func (s *HttpApiPolicyReference) SetPolicyId(v string) *HttpApiPolicyReference {
	s.PolicyId = &v
	return s
}

func (s *HttpApiPolicyReference) Validate() error {
	return dara.Validate(s)
}
