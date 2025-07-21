// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *TagResourceRequest
	GetCertificateId() *string
	SetKeyId(v string) *TagResourceRequest
	GetKeyId() *string
	SetSecretName(v string) *TagResourceRequest
	GetSecretName() *string
	SetTags(v string) *TagResourceRequest
	GetTags() *string
}

type TagResourceRequest struct {
	// The ID of the certificate.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 770dbe42-e146-43d1-a55a-1355db86****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddf****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the secret.
	//
	// >  You can configure only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// MyDbC****
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// One or more tags that you want to add. The value is in the array format.
	//
	// Tag attributes:
	//
	// 	- TagKey: the tag key.
	//
	// 	- TagValue: the tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"TagKey":"S1key1","TagValue":"S1val1"},{"TagKey":"S1key2","TagValue":"S2val2"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *TagResourceRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *TagResourceRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *TagResourceRequest) GetTags() *string {
	return s.Tags
}

func (s *TagResourceRequest) SetCertificateId(v string) *TagResourceRequest {
	s.CertificateId = &v
	return s
}

func (s *TagResourceRequest) SetKeyId(v string) *TagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *TagResourceRequest) SetSecretName(v string) *TagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *TagResourceRequest) SetTags(v string) *TagResourceRequest {
	s.Tags = &v
	return s
}

func (s *TagResourceRequest) Validate() error {
	return dara.Validate(s)
}
