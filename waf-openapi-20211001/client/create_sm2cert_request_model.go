// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSM2CertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *CreateSM2CertRequest
	GetCertName() *string
	SetEncryptCertificate(v string) *CreateSM2CertRequest
	GetEncryptCertificate() *string
	SetEncryptPrivateKey(v string) *CreateSM2CertRequest
	GetEncryptPrivateKey() *string
	SetInstanceId(v string) *CreateSM2CertRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateSM2CertRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateSM2CertRequest
	GetResourceManagerResourceGroupId() *string
	SetSignCertificate(v string) *CreateSM2CertRequest
	GetSignCertificate() *string
	SetSignPrivateKey(v string) *CreateSM2CertRequest
	GetSignPrivateKey() *string
}

type CreateSM2CertRequest struct {
	// The name of the SM certificate.
	//
	// example:
	//
	// test-sm2
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The content of the SM certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// ***
	//
	// -----END CERTIFICATE-----
	EncryptCertificate *string `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	// The private key of the SM certificate.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	//
	// ***
	//
	// -----END PRIVATE KEY-----
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The content of the signing certificate for the SM certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// ***
	//
	// -----END CERTIFICATE-----
	SignCertificate *string `json:"SignCertificate,omitempty" xml:"SignCertificate,omitempty"`
	// The private key of the signing certificate for the SM certificate.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	//
	// ***
	//
	// -----END PRIVATE KEY-----
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
}

func (s CreateSM2CertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSM2CertRequest) GoString() string {
	return s.String()
}

func (s *CreateSM2CertRequest) GetCertName() *string {
	return s.CertName
}

func (s *CreateSM2CertRequest) GetEncryptCertificate() *string {
	return s.EncryptCertificate
}

func (s *CreateSM2CertRequest) GetEncryptPrivateKey() *string {
	return s.EncryptPrivateKey
}

func (s *CreateSM2CertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSM2CertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSM2CertRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateSM2CertRequest) GetSignCertificate() *string {
	return s.SignCertificate
}

func (s *CreateSM2CertRequest) GetSignPrivateKey() *string {
	return s.SignPrivateKey
}

func (s *CreateSM2CertRequest) SetCertName(v string) *CreateSM2CertRequest {
	s.CertName = &v
	return s
}

func (s *CreateSM2CertRequest) SetEncryptCertificate(v string) *CreateSM2CertRequest {
	s.EncryptCertificate = &v
	return s
}

func (s *CreateSM2CertRequest) SetEncryptPrivateKey(v string) *CreateSM2CertRequest {
	s.EncryptPrivateKey = &v
	return s
}

func (s *CreateSM2CertRequest) SetInstanceId(v string) *CreateSM2CertRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSM2CertRequest) SetRegionId(v string) *CreateSM2CertRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSM2CertRequest) SetResourceManagerResourceGroupId(v string) *CreateSM2CertRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateSM2CertRequest) SetSignCertificate(v string) *CreateSM2CertRequest {
	s.SignCertificate = &v
	return s
}

func (s *CreateSM2CertRequest) SetSignPrivateKey(v string) *CreateSM2CertRequest {
	s.SignPrivateKey = &v
	return s
}

func (s *CreateSM2CertRequest) Validate() error {
	return dara.Validate(s)
}
