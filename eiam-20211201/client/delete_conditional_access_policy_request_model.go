// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConditionalAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditionalAccessPolicyId(v string) *DeleteConditionalAccessPolicyRequest
	GetConditionalAccessPolicyId() *string
	SetInstanceId(v string) *DeleteConditionalAccessPolicyRequest
	GetInstanceId() *string
}

type DeleteConditionalAccessPolicyRequest struct {
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

func (s DeleteConditionalAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConditionalAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteConditionalAccessPolicyRequest) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *DeleteConditionalAccessPolicyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteConditionalAccessPolicyRequest) SetConditionalAccessPolicyId(v string) *DeleteConditionalAccessPolicyRequest {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *DeleteConditionalAccessPolicyRequest) SetInstanceId(v string) *DeleteConditionalAccessPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteConditionalAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
