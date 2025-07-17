// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceId(v string) *CreatePolicyAttachmentRequest
	GetAttachResourceId() *string
	SetAttachResourceType(v string) *CreatePolicyAttachmentRequest
	GetAttachResourceType() *string
	SetEnvironmentId(v string) *CreatePolicyAttachmentRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *CreatePolicyAttachmentRequest
	GetGatewayId() *string
	SetPolicyId(v string) *CreatePolicyAttachmentRequest
	GetPolicyId() *string
}

type CreatePolicyAttachmentRequest struct {
	// Attached resource ID
	//
	// This parameter is required.
	//
	// example:
	//
	// api-cu07jj6m1hkokaus***
	AttachResourceId *string `json:"attachResourceId,omitempty" xml:"attachResourceId,omitempty"`
	// Attached resource type, such as HttpApi, GatewayRoute, Operation, GatewayService, GatewayServicePort, Gateway, Domain
	//
	// This parameter is required.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// Environment ID
	//
	// This parameter is required.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// Gateway instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-cq7l5s5lhtgi6qas***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// Policy ID
	//
	// This parameter is required.
	//
	// example:
	//
	// p-cq787etlhtghrptjg***
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
}

func (s CreatePolicyAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyAttachmentRequest) GetAttachResourceId() *string {
	return s.AttachResourceId
}

func (s *CreatePolicyAttachmentRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *CreatePolicyAttachmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreatePolicyAttachmentRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreatePolicyAttachmentRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyAttachmentRequest) SetAttachResourceId(v string) *CreatePolicyAttachmentRequest {
	s.AttachResourceId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetAttachResourceType(v string) *CreatePolicyAttachmentRequest {
	s.AttachResourceType = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetEnvironmentId(v string) *CreatePolicyAttachmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetGatewayId(v string) *CreatePolicyAttachmentRequest {
	s.GatewayId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) SetPolicyId(v string) *CreatePolicyAttachmentRequest {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
