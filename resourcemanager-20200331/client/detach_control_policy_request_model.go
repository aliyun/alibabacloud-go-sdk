// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DetachControlPolicyRequest
	GetPolicyId() *string
	SetTargetId(v string) *DetachControlPolicyRequest
	GetTargetId() *string
}

type DetachControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-jExXAqIYkwHN****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The ID of the object from which you want to detach the access control policy. Access control policies can be attached to the following objects:
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

func (s DetachControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DetachControlPolicyRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *DetachControlPolicyRequest) SetPolicyId(v string) *DetachControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DetachControlPolicyRequest) SetTargetId(v string) *DetachControlPolicyRequest {
	s.TargetId = &v
	return s
}

func (s *DetachControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
