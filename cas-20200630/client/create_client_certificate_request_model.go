// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateClientCertificateRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateClientCertificateRequest
	GetAlgorithm() *string
	SetAliasName(v string) *CreateClientCertificateRequest
	GetAliasName() *string
	SetBeforeTime(v int64) *CreateClientCertificateRequest
	GetBeforeTime() *int64
	SetClientToken(v string) *CreateClientCertificateRequest
	GetClientToken() *string
	SetCommonName(v string) *CreateClientCertificateRequest
	GetCommonName() *string
	SetCountry(v string) *CreateClientCertificateRequest
	GetCountry() *string
	SetCustomIdentifier(v string) *CreateClientCertificateRequest
	GetCustomIdentifier() *string
	SetDays(v int32) *CreateClientCertificateRequest
	GetDays() *int32
	SetEnableCrl(v int64) *CreateClientCertificateRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateClientCertificateRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateClientCertificateRequest
	GetLocality() *string
	SetMonths(v int32) *CreateClientCertificateRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateClientCertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateClientCertificateRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateClientCertificateRequest
	GetParentIdentifier() *string
	SetResourceGroupId(v string) *CreateClientCertificateRequest
	GetResourceGroupId() *string
	SetSanType(v int32) *CreateClientCertificateRequest
	GetSanType() *int32
	SetSanValue(v string) *CreateClientCertificateRequest
	GetSanValue() *string
	SetState(v string) *CreateClientCertificateRequest
	GetState() *string
	SetTags(v []*CreateClientCertificateRequestTags) *CreateClientCertificateRequest
	GetTags() []*CreateClientCertificateRequestTags
	SetYears(v int32) *CreateClientCertificateRequest
	GetYears() *int32
}

type CreateClientCertificateRequest struct {
	// The expiration time of the client certificate in UNIX timestamp format. The unit is seconds.
	//
	// > **BeforeTime*	- and **AfterTime*	- must be specified together or left empty together.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm for the client certificate. The format is `<encryption algorithm>_<key length>`. Valid values:
	//
	// - **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: The signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: The signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must be the same as the subordinate CA certificate. The key length can be different. For example, if the subordinate CA certificate uses the RSA_2048 key algorithm, the client certificate must use RSA_1024, RSA_2048, or RSA_4096.
	//
	// > Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to find the key algorithm of the subordinate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Set the name of the issued certificate.
	//
	// example:
	//
	// cert-name
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The issuance time of the client certificate in UNIX timestamp format. The unit is seconds. The default value is the time when you call this operation.
	//
	// > **BeforeTime*	- and **AfterTime*	- must be specified together or left empty together.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// Used to ensure request idempotence. The client generates this parameter value, which must be unique across different requests. It can contain a maximum of 64 ASCII characters and must not include any non-ASCII characters.
	//
	// example:
	//
	// XXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the certificate user. For a client authentication (ClientAuth) certificate, the user is typically an individual, a company, an organization, or an application. Specify the common name of the user, such as John Doe, Alibaba, Alibaba Cloud Cryptography Platform, or Tmall Genie.
	//
	// example:
	//
	// aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code. Default: CN.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// A custom identifier. This is a unique key.
	//
	// example:
	//
	// ****6bb538d538c70c01f81jh2****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the client certificate in days. The **Days**, **BeforeTime**, or **AfterTime*	- parameters cannot all be empty. The **BeforeTime*	- and **AfterTime*	- parameters must be set together or left empty. The parameters are configured as follows:
	//
	// - If you set the **Days*	- parameter, the **BeforeTime*	- and **AfterTime*	- parameters are optional.
	//
	// - If you do not set the **Days*	- parameter, you must set both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// > 	- If you set the **Days**, **BeforeTime**, and **AfterTime*	- parameters, the value of the **Days*	- parameter takes precedence.
	//
	// - The validity period of the client certificate cannot exceed the validity period of the subordinate CA certificate. To view the validity period of the subordinate CA certificate, you can call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html).
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Specifies whether to include the Certificate Revocation List (CRL) address.
	//
	// Valid values: 0 (No) and 1 (Yes).
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the digital certificate immediately.
	//
	// - **0**: No. This is the default value.
	//
	// - **1**: Yes, return the certificate.
	//
	// - **2**: Yes, return the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city where the organization is located. The default value is the city of the subordinate CA that issues the certificate.
	//
	// example:
	//
	// 杭州市
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the certificate in months.
	//
	// example:
	//
	// 1
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default: Alibaba Inc.
	//
	// example:
	//
	// 阿里云
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default: Alibaba Cloud CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the subordinate CA certificate that issues this certificate.
	//
	// > Call DescribeCACertificateList to query the unique identifier of the subordinate CA certificate.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of Subject Alternative Name (SAN) extension for the client certificate. Valid values:
	//
	// - **1**: Email
	//
	// - **6**: Uniform Resource Identifier (URI)
	//
	// example:
	//
	// 1
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The extension information for the client certificate. To enter multiple extensions, separate them with commas (,).
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// Specify the province or state of the certificate organization. The value can contain letters. The default value is the province or state of the organization for the intermediate CA that issued the certificate.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags.
	Tags []*CreateClientCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate in years.
	//
	// example:
	//
	// 5
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateClientCertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateClientCertificateRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *CreateClientCertificateRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateClientCertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateClientCertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateClientCertificateRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateClientCertificateRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *CreateClientCertificateRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateClientCertificateRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateClientCertificateRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateClientCertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateClientCertificateRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateClientCertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateClientCertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateClientCertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateClientCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClientCertificateRequest) GetSanType() *int32 {
	return s.SanType
}

func (s *CreateClientCertificateRequest) GetSanValue() *string {
	return s.SanValue
}

func (s *CreateClientCertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateClientCertificateRequest) GetTags() []*CreateClientCertificateRequestTags {
	return s.Tags
}

func (s *CreateClientCertificateRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateClientCertificateRequest) SetAfterTime(v int64) *CreateClientCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateClientCertificateRequest) SetAlgorithm(v string) *CreateClientCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateClientCertificateRequest) SetAliasName(v string) *CreateClientCertificateRequest {
	s.AliasName = &v
	return s
}

func (s *CreateClientCertificateRequest) SetBeforeTime(v int64) *CreateClientCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateClientCertificateRequest) SetClientToken(v string) *CreateClientCertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClientCertificateRequest) SetCommonName(v string) *CreateClientCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateClientCertificateRequest) SetCountry(v string) *CreateClientCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateClientCertificateRequest) SetCustomIdentifier(v string) *CreateClientCertificateRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *CreateClientCertificateRequest) SetDays(v int32) *CreateClientCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateClientCertificateRequest) SetEnableCrl(v int64) *CreateClientCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateClientCertificateRequest) SetImmediately(v int32) *CreateClientCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateClientCertificateRequest) SetLocality(v string) *CreateClientCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateClientCertificateRequest) SetMonths(v int32) *CreateClientCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateClientCertificateRequest) SetOrganization(v string) *CreateClientCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateClientCertificateRequest) SetOrganizationUnit(v string) *CreateClientCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateClientCertificateRequest) SetParentIdentifier(v string) *CreateClientCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateClientCertificateRequest) SetResourceGroupId(v string) *CreateClientCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClientCertificateRequest) SetSanType(v int32) *CreateClientCertificateRequest {
	s.SanType = &v
	return s
}

func (s *CreateClientCertificateRequest) SetSanValue(v string) *CreateClientCertificateRequest {
	s.SanValue = &v
	return s
}

func (s *CreateClientCertificateRequest) SetState(v string) *CreateClientCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateClientCertificateRequest) SetTags(v []*CreateClientCertificateRequestTags) *CreateClientCertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateClientCertificateRequest) SetYears(v int32) *CreateClientCertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateClientCertificateRequest) Validate() error {
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

type CreateClientCertificateRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// account
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClientCertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateClientCertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateClientCertificateRequestTags) SetKey(v string) *CreateClientCertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateClientCertificateRequestTags) SetValue(v string) *CreateClientCertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateClientCertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
