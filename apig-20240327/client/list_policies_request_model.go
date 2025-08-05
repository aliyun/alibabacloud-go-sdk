// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *ListPoliciesRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *ListPoliciesRequest
	GetAttachResourceType() *string
	SetEnvironmentId(v string) *ListPoliciesRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *ListPoliciesRequest
	GetGatewayId() *string
	SetWithAttachments(v bool) *ListPoliciesRequest
	GetWithAttachments() *bool
	SetWithSystemPolicy(v bool) *ListPoliciesRequest
	GetWithSystemPolicy() *bool
}

type ListPoliciesRequest struct {
	// example:
	//
	// api-cq7l5s5lhtgi6qasrdc0
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-cq2fm65lhtgm9nrrl7ag
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// false
	WithAttachments  *bool `json:"withAttachments,omitempty" xml:"withAttachments,omitempty"`
	WithSystemPolicy *bool `json:"withSystemPolicy,omitempty" xml:"withSystemPolicy,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *ListPoliciesRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *ListPoliciesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListPoliciesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListPoliciesRequest) GetWithAttachments() *bool {
	return s.WithAttachments
}

func (s *ListPoliciesRequest) GetWithSystemPolicy() *bool {
	return s.WithSystemPolicy
}

func (s *ListPoliciesRequest) SetAttachResourceId(v string) *ListPoliciesRequest {
	s.AttachResourceId = &v
	return s
}

func (s *ListPoliciesRequest) SetAttachResourceType(v string) *ListPoliciesRequest {
	s.AttachResourceType = &v
	return s
}

func (s *ListPoliciesRequest) SetEnvironmentId(v string) *ListPoliciesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListPoliciesRequest) SetGatewayId(v string) *ListPoliciesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListPoliciesRequest) SetWithAttachments(v bool) *ListPoliciesRequest {
	s.WithAttachments = &v
	return s
}

func (s *ListPoliciesRequest) SetWithSystemPolicy(v bool) *ListPoliciesRequest {
	s.WithSystemPolicy = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
