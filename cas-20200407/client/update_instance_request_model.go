// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReissue(v string) *UpdateInstanceRequest
	GetAutoReissue() *string
	SetCertificateName(v string) *UpdateInstanceRequest
	GetCertificateName() *string
	SetCity(v string) *UpdateInstanceRequest
	GetCity() *string
	SetCompanyId(v int64) *UpdateInstanceRequest
	GetCompanyId() *int64
	SetContactIdList(v []*int64) *UpdateInstanceRequest
	GetContactIdList() []*int64
	SetCountryCode(v string) *UpdateInstanceRequest
	GetCountryCode() *string
	SetCsr(v string) *UpdateInstanceRequest
	GetCsr() *string
	SetDomain(v string) *UpdateInstanceRequest
	GetDomain() *string
	SetGenerateCsrMethod(v string) *UpdateInstanceRequest
	GetGenerateCsrMethod() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetKeyAlgorithm(v string) *UpdateInstanceRequest
	GetKeyAlgorithm() *string
	SetProvince(v string) *UpdateInstanceRequest
	GetProvince() *string
	SetResourceGroupId(v string) *UpdateInstanceRequest
	GetResourceGroupId() *string
	SetTags(v []*UpdateInstanceRequestTags) *UpdateInstanceRequest
	GetTags() []*UpdateInstanceRequestTags
	SetValidationMethod(v string) *UpdateInstanceRequest
	GetValidationMethod() *string
}

type UpdateInstanceRequest struct {
	// Whether to enable automatic management.
	//
	// - enable: enabled
	//
	// - disable: disabled
	//
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// The name of the instance. When issuing a certificate, this value serves as the default certificate name.
	//
	// example:
	//
	// 123
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The city where the company or organization of the certificate requester is located. Required when generating a CSR for a DV certificate. Default value: Beijing.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The company information ID. Required for OV and EV certificates. Otherwise, you cannot call ApplyCertificate to request a certificate.
	//
	// example:
	//
	// 44211
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The list of contact IDs. You must specify at least one contact. Otherwise, you cannot call ApplyCertificate to request a certificate.
	ContactIdList []*int64 `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty" type:"Repeated"`
	// The country or region code of the certificate organization. For example, CN represents China, and US represents the United States. Required when generating a CSR for a DV certificate. Default value: CN.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The CSR content. You can generate a CSR using OpenSSL or Keytool. For more information, see [How to create a CSR file](https://help.aliyun.com/document_detail/42218.html).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The domain name to bind to the certificate. Requirements are as follows:
	//
	// - Supports single domain names or wildcard domain names (for example, `*.aliyundoc.com`).
	//
	// - Supports multiple domain names. Separate multiple domain names with commas (,). The first domain determines whether a free domain is included.
	//
	// 	Notice:
	//
	// If you bind multiple domain names to the certificate, this parameter is required. This parameter and the **Csr*	- parameter cannot both be empty. If you set both parameters, the system uses the **CN*	- field value from the **Csr*	- as the domain name for the certificate.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The CSR generation method. Default value: online.
	//
	// - online: The system generates the CSR. The Csr field is ignored.
	//
	// - upload: You upload the CSR. The Csr field is required.
	//
	// example:
	//
	// online
	GenerateCsrMethod *string `json:"GenerateCsrMethod,omitempty" xml:"GenerateCsrMethod,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-68n1mm16****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The certificate algorithm. Default value: RSA_2048.
	//
	// - **RSA_2048**
	//
	// - **RSA_3072**
	//
	// - **RSA_4096**
	//
	// - **ECC_256**
	//
	// - **SM2**
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The province or region where the company is located. Required when generating a CSR for a DV certificate. Default value: Beijing.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-ae******4wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// A list of tags.
	Tags []*UpdateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The certificate validation method.
	//
	// - DNS: DNS validation using TXT or CNAME records.
	//
	// - HTTP: File-based validation.
	//
	// example:
	//
	// DNS
	ValidationMethod *string `json:"ValidationMethod,omitempty" xml:"ValidationMethod,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAutoReissue() *string {
	return s.AutoReissue
}

func (s *UpdateInstanceRequest) GetCertificateName() *string {
	return s.CertificateName
}

func (s *UpdateInstanceRequest) GetCity() *string {
	return s.City
}

func (s *UpdateInstanceRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *UpdateInstanceRequest) GetContactIdList() []*int64 {
	return s.ContactIdList
}

func (s *UpdateInstanceRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *UpdateInstanceRequest) GetCsr() *string {
	return s.Csr
}

func (s *UpdateInstanceRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateInstanceRequest) GetGenerateCsrMethod() *string {
	return s.GenerateCsrMethod
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *UpdateInstanceRequest) GetProvince() *string {
	return s.Province
}

func (s *UpdateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateInstanceRequest) GetTags() []*UpdateInstanceRequestTags {
	return s.Tags
}

func (s *UpdateInstanceRequest) GetValidationMethod() *string {
	return s.ValidationMethod
}

func (s *UpdateInstanceRequest) SetAutoReissue(v string) *UpdateInstanceRequest {
	s.AutoReissue = &v
	return s
}

func (s *UpdateInstanceRequest) SetCertificateName(v string) *UpdateInstanceRequest {
	s.CertificateName = &v
	return s
}

func (s *UpdateInstanceRequest) SetCity(v string) *UpdateInstanceRequest {
	s.City = &v
	return s
}

func (s *UpdateInstanceRequest) SetCompanyId(v int64) *UpdateInstanceRequest {
	s.CompanyId = &v
	return s
}

func (s *UpdateInstanceRequest) SetContactIdList(v []*int64) *UpdateInstanceRequest {
	s.ContactIdList = v
	return s
}

func (s *UpdateInstanceRequest) SetCountryCode(v string) *UpdateInstanceRequest {
	s.CountryCode = &v
	return s
}

func (s *UpdateInstanceRequest) SetCsr(v string) *UpdateInstanceRequest {
	s.Csr = &v
	return s
}

func (s *UpdateInstanceRequest) SetDomain(v string) *UpdateInstanceRequest {
	s.Domain = &v
	return s
}

func (s *UpdateInstanceRequest) SetGenerateCsrMethod(v string) *UpdateInstanceRequest {
	s.GenerateCsrMethod = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetKeyAlgorithm(v string) *UpdateInstanceRequest {
	s.KeyAlgorithm = &v
	return s
}

func (s *UpdateInstanceRequest) SetProvince(v string) *UpdateInstanceRequest {
	s.Province = &v
	return s
}

func (s *UpdateInstanceRequest) SetResourceGroupId(v string) *UpdateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateInstanceRequest) SetTags(v []*UpdateInstanceRequestTags) *UpdateInstanceRequest {
	s.Tags = v
	return s
}

func (s *UpdateInstanceRequest) SetValidationMethod(v string) *UpdateInstanceRequest {
	s.ValidationMethod = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateInstanceRequestTags struct {
	// The tag key of the instance. Valid values for N: **1*	- to **20**. If you specify this value, it cannot be an empty string.
	//
	// It can contain up to 64 characters. It cannot start with `aliyun` or `acs:`, and it cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the instance. Valid values for N: **1*	- to **20**. If you specify this value, it can be an empty string.
	//
	// It can contain up to 128 characters. It cannot start with `aliyun` or `acs:`, and it cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateInstanceRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateInstanceRequestTags) SetTagKey(v string) *UpdateInstanceRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateInstanceRequestTags) SetTagValue(v string) *UpdateInstanceRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
