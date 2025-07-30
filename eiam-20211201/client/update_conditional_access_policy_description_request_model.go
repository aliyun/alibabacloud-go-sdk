// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConditionalAccessPolicyDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateConditionalAccessPolicyDescriptionRequest
	GetClientToken() *string
	SetConditionalAccessPolicyId(v string) *UpdateConditionalAccessPolicyDescriptionRequest
	GetConditionalAccessPolicyId() *string
	SetDescription(v string) *UpdateConditionalAccessPolicyDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateConditionalAccessPolicyDescriptionRequest
	GetInstanceId() *string
}

type UpdateConditionalAccessPolicyDescriptionRequest struct {
	// Idp client token.
	//
	// example:
	//
	// client-examplexxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Conditional Access Policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cap_11111
	ConditionalAccessPolicyId *string `json:"ConditionalAccessPolicyId,omitempty" xml:"ConditionalAccessPolicyId,omitempty"`
	// Description of the conditional access policy
	//
	// This parameter is required.
	//
	// example:
	//
	// Test Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateConditionalAccessPolicyDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConditionalAccessPolicyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) GetConditionalAccessPolicyId() *string {
	return s.ConditionalAccessPolicyId
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) SetClientToken(v string) *UpdateConditionalAccessPolicyDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) SetConditionalAccessPolicyId(v string) *UpdateConditionalAccessPolicyDescriptionRequest {
	s.ConditionalAccessPolicyId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) SetDescription(v string) *UpdateConditionalAccessPolicyDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) SetInstanceId(v string) *UpdateConditionalAccessPolicyDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConditionalAccessPolicyDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
