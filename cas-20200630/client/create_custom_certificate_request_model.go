// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiPassthrough(v *CreateCustomCertificateRequestApiPassthrough) *CreateCustomCertificateRequest
	GetApiPassthrough() *CreateCustomCertificateRequestApiPassthrough
	SetCsr(v string) *CreateCustomCertificateRequest
	GetCsr() *string
	SetEnableCrl(v int64) *CreateCustomCertificateRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateCustomCertificateRequest
	GetImmediately() *int32
	SetParentIdentifier(v string) *CreateCustomCertificateRequest
	GetParentIdentifier() *string
	SetResourceGroupId(v string) *CreateCustomCertificateRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateCustomCertificateRequestTags) *CreateCustomCertificateRequest
	GetTags() []*CreateCustomCertificateRequestTags
	SetValidity(v string) *CreateCustomCertificateRequest
	GetValidity() *string
	SetCustomIdentifier(v string) *CreateCustomCertificateRequest
	GetCustomIdentifier() *string
}

type CreateCustomCertificateRequest struct {
	// Pass-through parameters.
	ApiPassthrough *CreateCustomCertificateRequestApiPassthrough `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty" type:"Struct"`
	// The content of the CSR. You can generate a CSR using tools such as OpenSSL or Keytool. For more information, see [Create a CSR file](https://help.aliyun.com/document_detail/42218.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// vbIgMQIhAKHDWD6/WAMbtezAt4bysJ/BZIDz1jPWuUR5GV4TJ/mS
	//
	// -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// Specifies whether to include a CRL address.
	//
	// - 0 - No
	//
	// - 1 - Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Obtain the certificate immediately.
	//
	// - 0 - Issue the certificate asynchronously.
	//
	// - 1 - Issue the certificate immediately.
	//
	// - 2 - Issue the certificate immediately and return the CA certificate chain.
	//
	// example:
	//
	// 0
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The identifier of the CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ed4068c-6f1b-6deb-8e32-3f8439a851cb
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group. You can obtain this ID by calling the [ListResources](https://help.aliyun.com/document_detail/2716559.html) operation.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tags []*CreateCustomCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate. This period cannot exceed the validity period of the instance. You can use relative time or absolute time.
	//
	// Relative time: Supports years, months, and days.
	//
	// - Year - y
	//
	// - Month - m
	//
	// - Day - d
	//
	// Absolute time: Uses GMT. Format: `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"`
	//
	// - Specify the end time - $NotAfter
	//
	// - Specify the start and end times - $NotBefore/$NotAfter
	//
	// This parameter is required.
	//
	// example:
	//
	// 相对时间：
	//
	// ● 1y
	//
	// ● 3m
	//
	// ● 7d
	//
	// 绝对时间：
	//
	// ● 2006-01-02T15:04:05Z
	//
	// ● 2006-01-02T15:04:05Z/2023-03-09T17:48:13Z
	Validity *string `json:"Validity,omitempty" xml:"Validity,omitempty"`
	// A custom identifier.
	//
	// example:
	//
	// XXX068c-6f1b-6deb-8e32-3f8439a8XXX
	CustomIdentifier *string `json:"customIdentifier,omitempty" xml:"customIdentifier,omitempty"`
}

func (s CreateCustomCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequest) GetApiPassthrough() *CreateCustomCertificateRequestApiPassthrough {
	return s.ApiPassthrough
}

func (s *CreateCustomCertificateRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateCustomCertificateRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateCustomCertificateRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateCustomCertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateCustomCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCustomCertificateRequest) GetTags() []*CreateCustomCertificateRequestTags {
	return s.Tags
}

func (s *CreateCustomCertificateRequest) GetValidity() *string {
	return s.Validity
}

func (s *CreateCustomCertificateRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *CreateCustomCertificateRequest) SetApiPassthrough(v *CreateCustomCertificateRequestApiPassthrough) *CreateCustomCertificateRequest {
	s.ApiPassthrough = v
	return s
}

func (s *CreateCustomCertificateRequest) SetCsr(v string) *CreateCustomCertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetEnableCrl(v int64) *CreateCustomCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetImmediately(v int32) *CreateCustomCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetParentIdentifier(v string) *CreateCustomCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetResourceGroupId(v string) *CreateCustomCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetTags(v []*CreateCustomCertificateRequestTags) *CreateCustomCertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateCustomCertificateRequest) SetValidity(v string) *CreateCustomCertificateRequest {
	s.Validity = &v
	return s
}

func (s *CreateCustomCertificateRequest) SetCustomIdentifier(v string) *CreateCustomCertificateRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *CreateCustomCertificateRequest) Validate() error {
	if s.ApiPassthrough != nil {
		if err := s.ApiPassthrough.Validate(); err != nil {
			return err
		}
	}
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

type CreateCustomCertificateRequestApiPassthrough struct {
	// The certificate extensions.
	Extensions *CreateCustomCertificateRequestApiPassthroughExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Struct"`
	// The custom serial number of the certificate. Must be a long integer.
	//
	// example:
	//
	// 16889526086333
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The certificate subject.
	Subject *CreateCustomCertificateRequestApiPassthroughSubject `json:"Subject,omitempty" xml:"Subject,omitempty" type:"Struct"`
}

func (s CreateCustomCertificateRequestApiPassthrough) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthrough) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthrough) GetExtensions() *CreateCustomCertificateRequestApiPassthroughExtensions {
	return s.Extensions
}

func (s *CreateCustomCertificateRequestApiPassthrough) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateCustomCertificateRequestApiPassthrough) GetSubject() *CreateCustomCertificateRequestApiPassthroughSubject {
	return s.Subject
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetExtensions(v *CreateCustomCertificateRequestApiPassthroughExtensions) *CreateCustomCertificateRequestApiPassthrough {
	s.Extensions = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetSerialNumber(v string) *CreateCustomCertificateRequestApiPassthrough {
	s.SerialNumber = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthrough) SetSubject(v *CreateCustomCertificateRequestApiPassthroughSubject) *CreateCustomCertificateRequestApiPassthrough {
	s.Subject = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthrough) Validate() error {
	if s.Extensions != nil {
		if err := s.Extensions.Validate(); err != nil {
			return err
		}
	}
	if s.Subject != nil {
		if err := s.Subject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomCertificateRequestApiPassthroughExtensions struct {
	// If an extension is critical, its name is included in the criticals list.
	Criticals []*string `json:"Criticals,omitempty" xml:"Criticals,omitempty" type:"Repeated"`
	// The extended key usages.
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The key usage.
	KeyUsage *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" type:"Struct"`
	// The subject alternative names (SANs) of the certificate.
	SubjectAlternativeNames []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensions) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) GetCriticals() []*string {
	return s.Criticals
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) GetExtendedKeyUsages() []*string {
	return s.ExtendedKeyUsages
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) GetKeyUsage() *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	return s.KeyUsage
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) GetSubjectAlternativeNames() []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames {
	return s.SubjectAlternativeNames
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetCriticals(v []*string) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.Criticals = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetExtendedKeyUsages(v []*string) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.ExtendedKeyUsages = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetKeyUsage(v *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.KeyUsage = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) SetSubjectAlternativeNames(v []*CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) *CreateCustomCertificateRequestApiPassthroughExtensions {
	s.SubjectAlternativeNames = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensions) Validate() error {
	if s.KeyUsage != nil {
		if err := s.KeyUsage.Validate(); err != nil {
			return err
		}
	}
	if s.SubjectAlternativeNames != nil {
		for _, item := range s.SubjectAlternativeNames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage struct {
	// Content commitment. Formerly known as NonRepudiation. Allows the certificate key to be used for content commitment.
	//
	// example:
	//
	// false
	ContentCommitment *bool `json:"ContentCommitment,omitempty" xml:"ContentCommitment,omitempty"`
	// Data encipherment.
	//
	// example:
	//
	// false
	DataEncipherment *bool `json:"DataEncipherment,omitempty" xml:"DataEncipherment,omitempty"`
	// When KeyAgreement is true, this marks that the certificate key can only be used for decryption.
	//
	// example:
	//
	// false
	DecipherOnly *bool `json:"DecipherOnly,omitempty" xml:"DecipherOnly,omitempty"`
	// Digital signature. Allows the private key of the certificate to be used for digital signatures and the public key to be used to verify digital signatures.
	//
	// example:
	//
	// true
	DigitalSignature *bool `json:"DigitalSignature,omitempty" xml:"DigitalSignature,omitempty"`
	// When KeyAgreement is true, this marks that the certificate key can only be used for encryption.
	//
	// example:
	//
	// false
	EncipherOnly *bool `json:"EncipherOnly,omitempty" xml:"EncipherOnly,omitempty"`
	// Key agreement.
	//
	// example:
	//
	// false
	KeyAgreement *bool `json:"KeyAgreement,omitempty" xml:"KeyAgreement,omitempty"`
	// Key encipherment. Allows the certificate key to be used to encrypt other keys.
	//
	// example:
	//
	// false
	KeyEncipherment *bool `json:"KeyEncipherment,omitempty" xml:"KeyEncipherment,omitempty"`
	// Non-repudiation. This has been renamed to ContentCommitment in the X.509 standard.
	//
	// example:
	//
	// false
	NonRepudiation *bool `json:"NonRepudiation,omitempty" xml:"NonRepudiation,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetContentCommitment() *bool {
	return s.ContentCommitment
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetDataEncipherment() *bool {
	return s.DataEncipherment
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetDecipherOnly() *bool {
	return s.DecipherOnly
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetDigitalSignature() *bool {
	return s.DigitalSignature
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetEncipherOnly() *bool {
	return s.EncipherOnly
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetKeyAgreement() *bool {
	return s.KeyAgreement
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetKeyEncipherment() *bool {
	return s.KeyEncipherment
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) GetNonRepudiation() *bool {
	return s.NonRepudiation
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetContentCommitment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.ContentCommitment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDataEncipherment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DataEncipherment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDecipherOnly(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DecipherOnly = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetDigitalSignature(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.DigitalSignature = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetEncipherOnly(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.EncipherOnly = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetKeyAgreement(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.KeyAgreement = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetKeyEncipherment(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.KeyEncipherment = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) SetNonRepudiation(v bool) *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage {
	s.NonRepudiation = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage) Validate() error {
	return dara.Validate(s)
}

type CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames struct {
	// The following values are allowed:
	//
	// - rfc822Name - Email address
	//
	// - dNSName - Domain name
	//
	// - uniformResourceIdentifier - Uniform Resource Identifier (URI)
	//
	// - iPAddress - IP address
	//
	// This parameter is required.
	//
	// example:
	//
	// dNSName
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A value that matches the specified Type.
	//
	// example:
	//
	// rfc822Name:
	//
	// example.aliyundoc.com
	//
	// dNSName:
	//
	// learn.aliyundoc.com
	//
	// uniformResourceIdentifier:
	//
	// acs:ecs:regionid:15619224785*****:instance/i-bp1bzvz55uz27hf*****
	//
	// iPAddress:
	//
	// 127.0.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) GetType() *string {
	return s.Type
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) GetValue() *string {
	return s.Value
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) SetType(v string) *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames {
	s.Type = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) SetValue(v string) *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames {
	s.Value = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughExtensionsSubjectAlternativeNames) Validate() error {
	return dara.Validate(s)
}

type CreateCustomCertificateRequestApiPassthroughSubject struct {
	// The common name of the certificate user.
	//
	// example:
	//
	// 张三
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code. Use the two-letter country code from ISO 3166-1. For more information, see [ISO](https://www.iso.org/obp/ui/#search/code/).
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The custom subject properties of the certificate.
	CustomAttributes []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
	// The name of the city where the organization is located. Chinese characters and letters are supported.
	//
	// example:
	//
	// 杭州市
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// XXX公司
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch within the organization.
	//
	// example:
	//
	// XXX部门
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The province or state where the organization is located.
	//
	// example:
	//
	// 浙江省
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughSubject) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughSubject) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetCountry() *string {
	return s.Country
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetCustomAttributes() []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes {
	return s.CustomAttributes
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetLocality() *string {
	return s.Locality
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetOrganization() *string {
	return s.Organization
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) GetState() *string {
	return s.State
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCommonName(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.CommonName = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCountry(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Country = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetCustomAttributes(v []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.CustomAttributes = v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetLocality(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Locality = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetOrganization(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.Organization = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetOrganizationUnit(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) SetState(v string) *CreateCustomCertificateRequestApiPassthroughSubject {
	s.State = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubject) Validate() error {
	if s.CustomAttributes != nil {
		for _, item := range s.CustomAttributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes struct {
	// The key of the custom property. It must comply with industry standards. Examples:
	//
	// - 2.5.4.6: Country code
	//
	// - 2.5.4.10: Organization
	//
	// - 2.5.4.11: Organizational unit name
	//
	// - 2.5.4.12: Title
	//
	// - 2.5.4.3: Common name
	//
	// - 2.5.4.9: Street
	//
	// - 2.5.4.5: Serial number
	//
	// - 2.5.4.7: Locality
	//
	// - 2.5.4.8: State or province
	//
	// - 1.3.6.1.4.1.37244.1.1: Matter certificate - Node ID
	//
	// - 1.3.6.1.4.1.37244.1.5: Matter certificate - Fabric ID
	//
	// - 1.3.6.1.4.1.37244.2.1: Matter certificate Vendor ID (VID)
	//
	// - 1.3.6.1.4.1.37244.2.2: Matter certificate Product ID (PID)
	//
	// example:
	//
	// 2.5.4.3
	ObjectIdentifier *string `json:"ObjectIdentifier,omitempty" xml:"ObjectIdentifier,omitempty"`
	// The value of the custom property.
	//
	// example:
	//
	// Aliyun
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) GetObjectIdentifier() *string {
	return s.ObjectIdentifier
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) GetValue() *string {
	return s.Value
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) SetObjectIdentifier(v string) *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes {
	s.ObjectIdentifier = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) SetValue(v string) *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes {
	s.Value = &v
	return s
}

func (s *CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes) Validate() error {
	return dara.Validate(s)
}

type CreateCustomCertificateRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomCertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateCustomCertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateCustomCertificateRequestTags) SetKey(v string) *CreateCustomCertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateCustomCertificateRequestTags) SetValue(v string) *CreateCustomCertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateCustomCertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
