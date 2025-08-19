// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *AttachControlPolicyRequest
	GetPolicyId() *string
	SetTargetId(v string) *AttachControlPolicyRequest
	GetTargetId() *string
}

type AttachControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object to which you want to attach the access control policy. Access control policies can be attached to the following objects:
	//
	// 	- Root folder
	//
	// 	- Subfolders of the Root folder
	//
	// 	- Members
	//
	// This parameter is required.
	//
	// example:
	//
	// fd-ZDNPiT****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s AttachControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *AttachControlPolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *AttachControlPolicyRequest) SetPolicyId(v string) *AttachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *AttachControlPolicyRequest) SetTargetId(v string) *AttachControlPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *AttachControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
