// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64Encoded(v bool) *UpdateK8sSecretRequest
	GetBase64Encoded() *bool
	SetCertId(v string) *UpdateK8sSecretRequest
	GetCertId() *string
	SetCertRegionId(v string) *UpdateK8sSecretRequest
	GetCertRegionId() *string
	SetClusterId(v string) *UpdateK8sSecretRequest
	GetClusterId() *string
	SetData(v string) *UpdateK8sSecretRequest
	GetData() *string
	SetName(v string) *UpdateK8sSecretRequest
	GetName() *string
	SetNamespace(v string) *UpdateK8sSecretRequest
	GetNamespace() *string
	SetType(v string) *UpdateK8sSecretRequest
	GetType() *string
}

type UpdateK8sSecretRequest struct {
	// Specifies whether the data has been encoded in Base64.
	//
	// example:
	//
	// false
	Base64Encoded *bool `json:"Base64Encoded,omitempty" xml:"Base64Encoded,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 6650277
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The region ID of the certificate.
	//
	// example:
	//
	// cn-hangzhou
	CertRegionId *string `json:"CertRegionId,omitempty" xml:"CertRegionId,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// 9c28bbb9-****-44b3-b953-54ef8a2d0be2
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data of the Secret. The value must be a JSON array that contains the following information:
	//
	// 	- Key: Secret key
	//
	// 	- Value: Secret value
	//
	// example:
	//
	// [{"Key":"name","Value":"william"},{"Key":"age","Value":"12"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The name of the Secret. The name must start with a letter, and can contain digits, letters, and hyphens (-). It can be up to 63 characters in length.
	//
	// example:
	//
	// my-secret
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace of the Kubernetes cluster.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of the Secret. Valid values:
	//
	// 	- Opaque: user-defined data type
	//
	// 	- kubernetes.io/tls: Transport Layer Security (TLS) certificate type
	//
	// example:
	//
	// Opaque
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateK8sSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateK8sSecretRequest) GetBase64Encoded() *bool {
	return s.Base64Encoded
}

func (s *UpdateK8sSecretRequest) GetCertId() *string {
	return s.CertId
}

func (s *UpdateK8sSecretRequest) GetCertRegionId() *string {
	return s.CertRegionId
}

func (s *UpdateK8sSecretRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateK8sSecretRequest) GetData() *string {
	return s.Data
}

func (s *UpdateK8sSecretRequest) GetName() *string {
	return s.Name
}

func (s *UpdateK8sSecretRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateK8sSecretRequest) GetType() *string {
	return s.Type
}

func (s *UpdateK8sSecretRequest) SetBase64Encoded(v bool) *UpdateK8sSecretRequest {
	s.Base64Encoded = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetCertId(v string) *UpdateK8sSecretRequest {
	s.CertId = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetCertRegionId(v string) *UpdateK8sSecretRequest {
	s.CertRegionId = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetClusterId(v string) *UpdateK8sSecretRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetData(v string) *UpdateK8sSecretRequest {
	s.Data = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetName(v string) *UpdateK8sSecretRequest {
	s.Name = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetNamespace(v string) *UpdateK8sSecretRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateK8sSecretRequest) SetType(v string) *UpdateK8sSecretRequest {
	s.Type = &v
	return s
}

func (s *UpdateK8sSecretRequest) Validate() error {
	return dara.Validate(s)
}
