// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConditionalAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicyId(v string) *DisableConditionalAccessPolicyRequest
	GetConditionalAccessPolicyId() *string
	SetInstanceId(v string) *DisableConditionalAccessPolicyRequest
	GetInstanceId() *string
}

type DisableConditionalAccessPolicyRequest struct {
	// Conditional Access Policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cap_11111
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableConditionalAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableConditionalAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableConditionalAccessPolicyRequest) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *DisableConditionalAccessPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableConditionalAccessPolicyRequest) SetConditionalAccessPolicyId(v string) *DisableConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *DisableConditionalAccessPolicyRequest) SetInstanceId(v string) *DisableConditionalAccessPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableConditionalAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
