// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateSubCACertificateRequest
	GetAlgorithm() *string
	SetClientToken(v string) *CreateSubCACertificateRequest
	GetClientToken() *string
	SetCommonName(v string) *CreateSubCACertificateRequest
	GetCommonName() *string
	SetCountryCode(v string) *CreateSubCACertificateRequest
	GetCountryCode() *string
	SetCrlDay(v int32) *CreateSubCACertificateRequest
	GetCrlDay() *int32
	SetEnableCrl(v bool) *CreateSubCACertificateRequest
	GetEnableCrl() *bool
	SetExtendedKeyUsages(v []*string) *CreateSubCACertificateRequest
	GetExtendedKeyUsages() []*string
	SetLocality(v string) *CreateSubCACertificateRequest
	GetLocality() *string
	SetOrganization(v string) *CreateSubCACertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateSubCACertificateRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateSubCACertificateRequest
	GetParentIdentifier() *string
	SetPathLenConstraint(v int32) *CreateSubCACertificateRequest
	GetPathLenConstraint() *int32
	SetResourceGroupId(v string) *CreateSubCACertificateRequest
	GetResourceGroupId() *string
	SetState(v string) *CreateSubCACertificateRequest
	GetState() *string
	SetTags(v []*CreateSubCACertificateRequestTags) *CreateSubCACertificateRequest
	GetTags() []*CreateSubCACertificateRequestTags
	SetYears(v int32) *CreateSubCACertificateRequest
	GetYears() *int32
}

type CreateSubCACertificateRequest struct {
	// The type of the key algorithm of the intermediate CA. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// 	- **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of an intermediate CA certificate must be consistent with the encryption algorithm of a root CA certificate. The length of the keys can be different. For example, if the key algorithm of the root CA certificate is **RSA_2048**, the key algorithm of the intermediate CA certificate must be **RSA_1024**, **RSA_2048**, or **RSA_4096**.
	//
	// > You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) operation to query the key algorithm of a root CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	Algorithm   *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The common name or abbreviation of the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 or alpha-3 code. For example, you can use **CN*	- to indicate China and use **US*	- to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"*	- section in [Manage company profiles](https://help.aliyun.com/document_detail/198289.html).
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// CRL validity period: 1-365 days
	//
	// example:
	//
	// 30
	CrlDay *int32 `json:"CrlDay,omitempty" xml:"CrlDay,omitempty"`
	// Enable Crl Service.
	//
	// - 0- No
	//
	// - 1- Yes
	//
	// example:
	//
	// 1
	EnableCrl *bool `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// The extended key usages of the certificate.
	ExtendedKeyUsages []*string `json:"ExtendedKeyUsages,omitempty" xml:"ExtendedKeyUsages,omitempty" type:"Repeated"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the intermediate CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Maizi Technology
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate.
	//
	// > You can call the [DescribeCACertificateList] operation to query the unique identifiers of all CA certificates.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The path length constraint of the certificate. Default value: 0.
	//
	// example:
	//
	// 0
	PathLenConstraint *int32  `json:"PathLenConstraint,omitempty" xml:"PathLenConstraint,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the province or state in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string                              `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*CreateSubCACertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the intermediate CA certificate. Unit: years.
	//
	// We recommend that you set this parameter to 5 to 10.
	//
	// > The validity period of the intermediate CA certificate cannot exceed the validity period of the root CA certificate. You can call the [DescribeCACertificate]operation to query the validity period of a root CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateSubCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCACertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateSubCACertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSubCACertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateSubCACertificateRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CreateSubCACertificateRequest) GetCrlDay() *int32 {
	return s.CrlDay
}

func (s *CreateSubCACertificateRequest) GetEnableCrl() *bool {
	return s.EnableCrl
}

func (s *CreateSubCACertificateRequest) GetExtendedKeyUsages() []*string {
	return s.ExtendedKeyUsages
}

func (s *CreateSubCACertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateSubCACertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateSubCACertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateSubCACertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateSubCACertificateRequest) GetPathLenConstraint() *int32 {
	return s.PathLenConstraint
}

func (s *CreateSubCACertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateSubCACertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateSubCACertificateRequest) GetTags() []*CreateSubCACertificateRequestTags {
	return s.Tags
}

func (s *CreateSubCACertificateRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateSubCACertificateRequest) SetAlgorithm(v string) *CreateSubCACertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetClientToken(v string) *CreateSubCACertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCommonName(v string) *CreateSubCACertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCountryCode(v string) *CreateSubCACertificateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetCrlDay(v int32) *CreateSubCACertificateRequest {
	s.CrlDay = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetEnableCrl(v bool) *CreateSubCACertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetExtendedKeyUsages(v []*string) *CreateSubCACertificateRequest {
	s.ExtendedKeyUsages = v
	return s
}

func (s *CreateSubCACertificateRequest) SetLocality(v string) *CreateSubCACertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetOrganization(v string) *CreateSubCACertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetOrganizationUnit(v string) *CreateSubCACertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetParentIdentifier(v string) *CreateSubCACertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetPathLenConstraint(v int32) *CreateSubCACertificateRequest {
	s.PathLenConstraint = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetResourceGroupId(v string) *CreateSubCACertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetState(v string) *CreateSubCACertificateRequest {
	s.State = &v
	return s
}

func (s *CreateSubCACertificateRequest) SetTags(v []*CreateSubCACertificateRequestTags) *CreateSubCACertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateSubCACertificateRequest) SetYears(v int32) *CreateSubCACertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateSubCACertificateRequest) Validate() error {
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

type CreateSubCACertificateRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSubCACertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCACertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateSubCACertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateSubCACertificateRequestTags) SetKey(v string) *CreateSubCACertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateSubCACertificateRequestTags) SetValue(v string) *CreateSubCACertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateSubCACertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
