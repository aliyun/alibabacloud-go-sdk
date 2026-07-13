// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *UpdateServiceEndpointRequest
	GetCertIdentifier() *string
	SetClientToken(v string) *UpdateServiceEndpointRequest
	GetClientToken() *string
	SetDomain(v string) *UpdateServiceEndpointRequest
	GetDomain() *string
	SetEndpointId(v string) *UpdateServiceEndpointRequest
	GetEndpointId() *string
	SetInstanceId(v string) *UpdateServiceEndpointRequest
	GetInstanceId() *string
}

type UpdateServiceEndpointRequest struct {
	// example:
	//
	// 22584627-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// api.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mep-abc123
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateServiceEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceEndpointRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateServiceEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateServiceEndpointRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateServiceEndpointRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateServiceEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServiceEndpointRequest) SetCertIdentifier(v string) *UpdateServiceEndpointRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateServiceEndpointRequest) SetClientToken(v string) *UpdateServiceEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceEndpointRequest) SetDomain(v string) *UpdateServiceEndpointRequest {
	s.Domain = &v
	return s
}

func (s *UpdateServiceEndpointRequest) SetEndpointId(v string) *UpdateServiceEndpointRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateServiceEndpointRequest) SetInstanceId(v string) *UpdateServiceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateServiceEndpointRequest) Validate() error {
	return dara.Validate(s)
}
