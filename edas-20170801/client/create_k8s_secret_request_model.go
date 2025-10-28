// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64Encoded(v bool) *CreateK8sSecretRequest
	GetBase64Encoded() *bool
	SetCertId(v string) *CreateK8sSecretRequest
	GetCertId() *string
	SetCertRegionId(v string) *CreateK8sSecretRequest
	GetCertRegionId() *string
	SetClusterId(v string) *CreateK8sSecretRequest
	GetClusterId() *string
	SetData(v string) *CreateK8sSecretRequest
	GetData() *string
	SetName(v string) *CreateK8sSecretRequest
	GetName() *string
	SetNamespace(v string) *CreateK8sSecretRequest
	GetNamespace() *string
	SetType(v string) *CreateK8sSecretRequest
	GetType() *string
}

type CreateK8sSecretRequest struct {
	// Specifies whether the data has been encoded in Base64. Valid values: true and false.
	//
	// example:
	//
	// true
	Base64Encoded *bool `json:"Base64Encoded,omitempty" xml:"Base64Encoded,omitempty"`
	// The certificate ID provided by Alibaba Cloud Certificate Management Service.
	//
	// example:
	//
	// 6651483
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The region in which the certificate is stored.
	//
	// example:
	//
	// cn-hangzhou
	CertRegionId *string `json:"CertRegionId,omitempty" xml:"CertRegionId,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// 712082c3-****-****-9217-a947b5cde6ee
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
	// The Secret type. Valid values:
	//
	// 	- Opaque: user-defined data
	//
	// 	- kubernetes.io/tls: Transport Layer Security (TLS) certificate
	//
	// example:
	//
	// Opaque
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateK8sSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateK8sSecretRequest) GetBase64Encoded() *bool {
	return s.Base64Encoded
}

func (s *CreateK8sSecretRequest) GetCertId() *string {
	return s.CertId
}

func (s *CreateK8sSecretRequest) GetCertRegionId() *string {
	return s.CertRegionId
}

func (s *CreateK8sSecretRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateK8sSecretRequest) GetData() *string {
	return s.Data
}

func (s *CreateK8sSecretRequest) GetName() *string {
	return s.Name
}

func (s *CreateK8sSecretRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateK8sSecretRequest) GetType() *string {
	return s.Type
}

func (s *CreateK8sSecretRequest) SetBase64Encoded(v bool) *CreateK8sSecretRequest {
	s.Base64Encoded = &v
	return s
}

func (s *CreateK8sSecretRequest) SetCertId(v string) *CreateK8sSecretRequest {
	s.CertId = &v
	return s
}

func (s *CreateK8sSecretRequest) SetCertRegionId(v string) *CreateK8sSecretRequest {
	s.CertRegionId = &v
	return s
}

func (s *CreateK8sSecretRequest) SetClusterId(v string) *CreateK8sSecretRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateK8sSecretRequest) SetData(v string) *CreateK8sSecretRequest {
	s.Data = &v
	return s
}

func (s *CreateK8sSecretRequest) SetName(v string) *CreateK8sSecretRequest {
	s.Name = &v
	return s
}

func (s *CreateK8sSecretRequest) SetNamespace(v string) *CreateK8sSecretRequest {
	s.Namespace = &v
	return s
}

func (s *CreateK8sSecretRequest) SetType(v string) *CreateK8sSecretRequest {
	s.Type = &v
	return s
}

func (s *CreateK8sSecretRequest) Validate() error {
	return dara.Validate(s)
}
