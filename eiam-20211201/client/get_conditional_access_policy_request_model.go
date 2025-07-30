// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConditionalAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicyId(v string) *GetConditionalAccessPolicyRequest
	GetConditionalAccessPolicyId() *string
	SetInstanceId(v string) *GetConditionalAccessPolicyRequest
	GetInstanceId() *string
}

type GetConditionalAccessPolicyRequest struct {
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

func (s GetConditionalAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConditionalAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetConditionalAccessPolicyRequest) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *GetConditionalAccessPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetConditionalAccessPolicyRequest) SetConditionalAccessPolicyId(v string) *GetConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *GetConditionalAccessPolicyRequest) SetInstanceId(v string) *GetConditionalAccessPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConditionalAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
