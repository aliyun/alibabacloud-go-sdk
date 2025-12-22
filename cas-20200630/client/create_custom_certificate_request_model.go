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
	// The passthrough parameters.
	ApiPassthrough *CreateCustomCertificateRequestApiPassthrough `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty" type:"Struct"`
	// The content of the CSR. You can generate a CSR by using the OpenSSL tool or the Keytool tool. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html)
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
	// include the CRL address.
	//
	// - 0- No
	//
	// - 1- Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to immediately issue the certificate. Valid values:
	//
	// 	- 0: asynchronously issues the certificate.
	//
	// 	- 1: immediately issues the certificate.
	//
	// 	- 2: immediately issues the certificate and returns the certificate chain.
	//
	// example:
	//
	// 0
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The identifier of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ed4068c-6f1b-6deb-8e32-3f8439a851cb
	ParentIdentifier *string                               `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	ResourceGroupId  *string                               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags             []*CreateCustomCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate. The value cannot exceed the validity period of the certificate instance. Relative time and absolute time are supported.
	//
	// Units of relative time: year, month, and day.
	//
	// 	- Use y to specify years.
	//
	// 	- Use m to specify months.
	//
	// 	- Use d to specify days.
	//
	// Absolute time: Use Greenwich Mean Time (GMT). Format: `yyyy-MM-dd\\"T\\"HH:mm:ss\\"Z\\"`
	//
	// 	- Format of the end time: $NotAfter
	//
	// 	- Format of the start time and end time: $NotBefore/$NotAfter
	//
	// This parameter is required.
	//
	// example:
	//
	// Relative time:
	//
	//  ● 1y
	//
	//  ● 3m
	//
	//  ● 7d
	//
	// Absolute time:
	//
	// ● 2006-01-02T15:04:05Z
	//
	// ● 2006-01-02T15:04:05Z/2023-03-09T17:48:13Z
	Validity         *string `json:"Validity,omitempty" xml:"Validity,omitempty"`
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
	// The extensions of the certificate.
	Extensions *CreateCustomCertificateRequestApiPassthroughExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Struct"`
	// The serial number MUST be a positive integer assigned by the CA to each certificate.
	//
	// example:
	//
	// 16889526086333
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The name of the entity that uses the certificate.
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
	// If it is a necessary parameter, the critical list contains the parameter name.
	Criticals []*string `json:"Criticals,omitempty" xml:"Criticals,omitempty" type:"Repeated"`
	// The extended key usage.
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The key usage.
	KeyUsage *CreateCustomCertificateRequestApiPassthroughExtensionsKeyUsage `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" type:"Struct"`
	// The aliases of the entities.
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
	// The original name of the parameter is NonRepudiation.
	//
	// example:
	//
	// false
	ContentCommitment *bool `json:"ContentCommitment,omitempty" xml:"ContentCommitment,omitempty"`
	// Specifies whether the key can be used for data encryption.
	//
	// example:
	//
	// false
	DataEncipherment *bool `json:"DataEncipherment,omitempty" xml:"DataEncipherment,omitempty"`
	// Specifies whether the key can be used only for data decryption.
	//
	// example:
	//
	// false
	DecipherOnly *bool `json:"DecipherOnly,omitempty" xml:"DecipherOnly,omitempty"`
	// Specifies whether the key can be used for digital signing. If you set this parameter to true, the private key of the certificate can be used to generate digital signatures, and the public key of the certificate can be used to verify digital signatures.
	//
	// example:
	//
	// true
	DigitalSignature *bool `json:"DigitalSignature,omitempty" xml:"DigitalSignature,omitempty"`
	// Specifies whether the key can be used only for data encryption.
	//
	// example:
	//
	// false
	EncipherOnly *bool `json:"EncipherOnly,omitempty" xml:"EncipherOnly,omitempty"`
	// Specifies whether the key can be used for key agreement.
	//
	// example:
	//
	// false
	KeyAgreement *bool `json:"KeyAgreement,omitempty" xml:"KeyAgreement,omitempty"`
	// Specifies whether the key can be used for data encipherment.
	//
	// example:
	//
	// false
	KeyEncipherment *bool `json:"KeyEncipherment,omitempty" xml:"KeyEncipherment,omitempty"`
	// Specifies whether the key can be used for non-repudiation. This parameter is renamed ContentCommitment in the X.509 standard.
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
	// The type of the alias. Valid values:
	//
	// 	- rfc822Name: email address
	//
	// 	- dNSName: domain name
	//
	// 	- uniformResourceIdentifier: URI
	//
	// 	- iPAddress: IP address
	//
	// This parameter is required.
	//
	// example:
	//
	// dNSName
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The alias that meets the requirement of a specified type.
	//
	// example:
	//
	// rfc822Name:
	//
	// exmaple@certqa.cn
	//
	// dNSName:
	//
	// www.certqa.cn
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
	// Bob
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country. The value is an alpha-2 country code that complies with the ISO 3166-1 standard. For more information about country codes, visit <https://www.iso.org/obp/ui/#search/code/>.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Customize the Subject attributes of the certificate.
	CustomAttributes []*CreateCustomCertificateRequestApiPassthroughSubjectCustomAttributes `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty" type:"Repeated"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// XXX company
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization.
	//
	// example:
	//
	// XXX department
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The name of the province or state in which the organization associated with the certificate is located.
	//
	// example:
	//
	// Zhejiang
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
	// Custom attribute type as:
	//
	// - 2.5.4.6 : country
	//
	// - 2.5.4.10 : organization
	//
	// - 2.5.4.11 : organizational unit
	//
	// - 2.5.4.12 : title
	//
	// - 2.5.4.3 : common name
	//
	// - 2.5.4.9 : street
	//
	// - 2.5.4.5 : serial number
	//
	// - 2.5.4.7 : locality
	//
	// - 2.5.4.8 : state
	//
	// - 1.3.6.1.4.1.37244.1.1 : Matter Operational Certificate - Node ID
	//
	// - 1.3.6.1.4.1.37244.1.5 : Matter Operational Certificate - Fabric ID
	//
	// - 1.3.6.1.4.1.37244.2.1 : Matter Device Attestation Certificate Vender ID (VID)
	//
	// - 1.3.6.1.4.1.37244.2.2 : Matter Device Attestation Certificate Product ID (PID).
	//
	// example:
	//
	// 2.5.4.3
	ObjectIdentifier *string `json:"ObjectIdentifier,omitempty" xml:"ObjectIdentifier,omitempty"`
	// Custom attribute value.
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
