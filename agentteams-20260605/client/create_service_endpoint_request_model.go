// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *CreateServiceEndpointRequest
	GetCertIdentifier() *string
	SetClientToken(v string) *CreateServiceEndpointRequest
	GetClientToken() *string
	SetComponent(v string) *CreateServiceEndpointRequest
	GetComponent() *string
	SetDomain(v string) *CreateServiceEndpointRequest
	GetDomain() *string
	SetInstanceId(v string) *CreateServiceEndpointRequest
	GetInstanceId() *string
	SetResourceName(v string) *CreateServiceEndpointRequest
	GetResourceName() *string
}

type CreateServiceEndpointRequest struct {
	// example:
	//
	// cert-001
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WORKER
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ins-001
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s CreateServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceEndpointRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *CreateServiceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceEndpointRequest) GetComponent() *string {
	return s.Component
}

func (s *CreateServiceEndpointRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateServiceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateServiceEndpointRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateServiceEndpointRequest) SetCertIdentifier(v string) *CreateServiceEndpointRequest {
	s.CertIdentifier = &v
	return s
}

func (s *CreateServiceEndpointRequest) SetClientToken(v string) *CreateServiceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceEndpointRequest) SetComponent(v string) *CreateServiceEndpointRequest {
	s.Component = &v
	return s
}

func (s *CreateServiceEndpointRequest) SetDomain(v string) *CreateServiceEndpointRequest {
	s.Domain = &v
	return s
}

func (s *CreateServiceEndpointRequest) SetInstanceId(v string) *CreateServiceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateServiceEndpointRequest) SetResourceName(v string) *CreateServiceEndpointRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateServiceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
