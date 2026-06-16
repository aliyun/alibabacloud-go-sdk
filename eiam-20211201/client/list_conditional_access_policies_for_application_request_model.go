// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConditionalAccessPoliciesForApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListConditionalAccessPoliciesForApplicationRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListConditionalAccessPoliciesForApplicationRequest
	GetInstanceId() *string
}

type ListConditionalAccessPoliciesForApplicationRequest struct {
	// The ID of the application that is associated with the conditional access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_11111
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListConditionalAccessPoliciesForApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConditionalAccessPoliciesForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListConditionalAccessPoliciesForApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListConditionalAccessPoliciesForApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListConditionalAccessPoliciesForApplicationRequest) SetApplicationId(v string) *ListConditionalAccessPoliciesForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationRequest) SetInstanceId(v string) *ListConditionalAccessPoliciesForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListConditionalAccessPoliciesForApplicationRequest) Validate() error {
	return dara.Validate(s)
}
