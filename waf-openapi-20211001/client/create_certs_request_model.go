// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertContent(v string) *CreateCertsRequest
	GetCertContent() *string
	SetCertKey(v string) *CreateCertsRequest
	GetCertKey() *string
	SetCertName(v string) *CreateCertsRequest
	GetCertName() *string
	SetInstanceId(v string) *CreateCertsRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateCertsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateCertsRequest
	GetResourceManagerResourceGroupId() *string
}

type CreateCertsRequest struct {
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- 62EcYPWd2****s6MTXcJSfN9Z7rZ9fmxWr2BFN2XbahgnsSXM48ixZJ4krc+1M+j2kcubVpsE2cgHdj4v8H6jUz9Ji4mr7vMNS6dXv8PUkl/qoDeNGCNdyTS5NIL5ir+g92cL8IGOkjgvhlqt9vc65Cgb4mL+n5+DV9uOyTZTW/MojmlgfUekC2xiXa54nxJf17Y1TADGSbyJbsC0Q9nIrHsPl8YKkvRWvIAqYxXZ7wRwWWmv4TMxFhWRiNY7yZIo2ZUhl02S****gIEeg== -----END CERTIFICATE-----
	CertContent *string `json:"CertContent,omitempty" xml:"CertContent,omitempty"`
	// The private key that corresponds to the certificate.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- DADTPZoOHd9WtZ3U****RgNQmioPQn2bqdKHop+B/dn/4VZL7Jt8zSDGM9sTMThLyvsmLQKBgQCr+ujntC1kN6pGBj2Fw2l/EA/W3rYEce2tyhjgmG7rZ+A/jVE9fld5sQra6ZdwBcQJaiygoIYoaMF2EjRwc0qwHaluq0C15f6ujSoHh2e+D5zdmkTg/3NKNjqNv6xA2gYpinVDzFdZ9Zujxvuh9o4Vqf0YF8****K5G04RtKadOw== -----END RSA PRIVATE KEY-----
	CertKey *string `json:"CertKey,omitempty" xml:"CertKey,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// testrsa1
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// example:
	//
	// waf_v2_public_cn-20p****nw01
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
}

func (s CreateCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCertsRequest) GoString() string {
	return s.String()
}

func (s *CreateCertsRequest) GetCertContent() *string {
	return s.CertContent
}

func (s *CreateCertsRequest) GetCertKey() *string {
	return s.CertKey
}

func (s *CreateCertsRequest) GetCertName() *string {
	return s.CertName
}

func (s *CreateCertsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCertsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateCertsRequest) SetCertContent(v string) *CreateCertsRequest {
	s.CertContent = &v
	return s
}

func (s *CreateCertsRequest) SetCertKey(v string) *CreateCertsRequest {
	s.CertKey = &v
	return s
}

func (s *CreateCertsRequest) SetCertName(v string) *CreateCertsRequest {
	s.CertName = &v
	return s
}

func (s *CreateCertsRequest) SetInstanceId(v string) *CreateCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCertsRequest) SetRegionId(v string) *CreateCertsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCertsRequest) SetResourceManagerResourceGroupId(v string) *CreateCertsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateCertsRequest) Validate() error {
	return dara.Validate(s)
}
