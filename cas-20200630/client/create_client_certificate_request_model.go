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
	SetBeforeTime(v int64) *CreateClientCertificateRequest
	GetBeforeTime() *int64
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
	// The expiration time of the client certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// 	- **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_384**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_512**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate certificate authority (CA) certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the client certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// > You can call the [DescribeCACertificate] operation to query the key algorithm of an intermediate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the client certificate user. In most cases, the user of a client certificate is an individual, a company, an organization, or an application. We recommend that you enter the common name of a user. Examples: Bob, Alibaba, Alibaba Cloud password platform, and Tmall Genie.
	//
	// example:
	//
	// aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country in which the organization is located. Default value: CN.
	//
	// example:
	//
	// CN
	Country          *string `json:"Country,omitempty" xml:"Country,omitempty"`
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the client certificate. Unit: day. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters at the same time, the validity period of the client certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
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
	// Specifies whether to return the certificate. Valid values:
	//
	// 	- **0**: does not return the certificate. This is the default value.
	//
	// 	- **1**: returns the certificate.
	//
	// 	- **2**: returns the certificate and the certificate chain of the certificate.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the client certificate. Unit: months.
	//
	// example:
	//
	// 1
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// > You can call the [DescribeCACertificateList] operation to query the unique identifier of an intermediate CA certificate.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// 	- **1**: an email address
	//
	// 	- **6**: a Uniform Resource Identifier (URI)
	//
	// example:
	//
	// 1
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string                               `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*CreateClientCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the client certificate. Unit: years.
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

func (s *CreateClientCertificateRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
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

func (s *CreateClientCertificateRequest) SetBeforeTime(v int64) *CreateClientCertificateRequest {
	s.BeforeTime = &v
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
