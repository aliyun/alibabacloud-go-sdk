// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateServerCertificateRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateServerCertificateRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateServerCertificateRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateServerCertificateRequest
	GetCommonName() *string
	SetCountry(v string) *CreateServerCertificateRequest
	GetCountry() *string
	SetCustomIdentifier(v string) *CreateServerCertificateRequest
	GetCustomIdentifier() *string
	SetDays(v int32) *CreateServerCertificateRequest
	GetDays() *int32
	SetDomain(v string) *CreateServerCertificateRequest
	GetDomain() *string
	SetEnableCrl(v int64) *CreateServerCertificateRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateServerCertificateRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateServerCertificateRequest
	GetLocality() *string
	SetMonths(v int32) *CreateServerCertificateRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateServerCertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateServerCertificateRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateServerCertificateRequest
	GetParentIdentifier() *string
	SetResourceGroupId(v string) *CreateServerCertificateRequest
	GetResourceGroupId() *string
	SetState(v string) *CreateServerCertificateRequest
	GetState() *string
	SetTags(v []*CreateServerCertificateRequestTags) *CreateServerCertificateRequest
	GetTags() []*CreateServerCertificateRequestTags
	SetYears(v int32) *CreateServerCertificateRequest
	GetYears() *int32
}

type CreateServerCertificateRequest struct {
	// The expiration time of the server certificate. This value is a UNIX timestamp in seconds.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must be specified together or left empty together.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The algorithm is in the `<encryption algorithm>_<key length>` format. Valid values:
	//
	// - **RSA_1024**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: The corresponding signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server certificate must be the same as the encryption algorithm of the subordinate CA certificate, but the key length can be different. For example, if the key algorithm of the subordinate CA certificate is RSA_2048, the key algorithm of the server certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// > Call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to query the key algorithm of the subordinate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the server certificate. This value is a UNIX timestamp in seconds. The default value is the time when you call this operation.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must be specified together or left empty together.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the certificate user. For a server authentication (ServerAuth) certificate, the user is the server. Enter the domain name or IP address that is bound to the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country code, such as CN or US.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// A custom identifier. This key must be unique.
	//
	// example:
	//
	// ****6bb538d538c70c01f81dg3****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the server certificate, in days. The **Days**, **BeforeTime**, and **AfterTime*	- parameters cannot all be empty. The **BeforeTime*	- and **AfterTime*	- parameters must be specified together or left empty together. The following rules describe how to set these parameters:
	//
	// - If you specify **Days**, the **BeforeTime*	- and **AfterTime*	- parameters are optional.
	//
	// - If you do not specify **Days**, you must specify both **BeforeTime*	- and **AfterTime**.
	//
	// > 	- If you specify **Days**, **BeforeTime**, and **AfterTime*	- at the same time, the value of **Days*	- determines the validity period of the server certificate.
	//
	// - The validity period of the server certificate cannot exceed the validity period of the subordinate CA certificate. You can call [DescribeCACertificate](https://help.aliyun.com/document_detail/465954.html) to view the validity period of the subordinate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names and IP addresses for the server certificate. This information lets you apply the certificate to multiple domain names and IP addresses.
	//
	// Separate multiple domain names or IP addresses with a comma (,).
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to include the Certificate Revocation List (CRL) address.
	//
	// 0: No
	//
	// 1: Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the digital certificate immediately.
	//
	// - **0**: No. This is the default value.
	//
	// - **1**: Returns the certificate.
	//
	// - **2**: Returns the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The city where the organization is located. Chinese and English characters are supported. The default value is the city of the organization that is associated with the subordinate CA certificate that issues this certificate.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the certificate, in months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. The default value is Alibaba Inc.
	//
	// example:
	//
	// 阿里云
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. The default value is Alibaba Cloud CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the subordinate CA certificate that issues this certificate.
	//
	// > Call [DescribeCACertificateList](https://help.aliyun.com/document_detail/465957.html) to query the unique identifier of the subordinate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 271ae6bb538d538c70c01f81dg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group. Call the [ListResources](https://help.aliyun.com/document_detail/2716559.html) operation to get this ID.
	//
	// example:
	//
	// test
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The province or state where the organization is located. Chinese and English characters are supported. The default value is the province or state of the organization that is associated with the subordinate CA certificate that issues this certificate.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags.
	Tags []*CreateServerCertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the certificate, in years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateServerCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateServerCertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateServerCertificateRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateServerCertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateServerCertificateRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateServerCertificateRequest) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *CreateServerCertificateRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateServerCertificateRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateServerCertificateRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateServerCertificateRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateServerCertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateServerCertificateRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateServerCertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateServerCertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateServerCertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateServerCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerCertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateServerCertificateRequest) GetTags() []*CreateServerCertificateRequestTags {
	return s.Tags
}

func (s *CreateServerCertificateRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateServerCertificateRequest) SetAfterTime(v int64) *CreateServerCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetAlgorithm(v string) *CreateServerCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateServerCertificateRequest) SetBeforeTime(v int64) *CreateServerCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCommonName(v string) *CreateServerCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCountry(v string) *CreateServerCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCustomIdentifier(v string) *CreateServerCertificateRequest {
	s.CustomIdentifier = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDays(v int32) *CreateServerCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDomain(v string) *CreateServerCertificateRequest {
	s.Domain = &v
	return s
}

func (s *CreateServerCertificateRequest) SetEnableCrl(v int64) *CreateServerCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateServerCertificateRequest) SetImmediately(v int32) *CreateServerCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateServerCertificateRequest) SetLocality(v string) *CreateServerCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateServerCertificateRequest) SetMonths(v int32) *CreateServerCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganization(v string) *CreateServerCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganizationUnit(v string) *CreateServerCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateServerCertificateRequest) SetParentIdentifier(v string) *CreateServerCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateServerCertificateRequest) SetResourceGroupId(v string) *CreateServerCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerCertificateRequest) SetState(v string) *CreateServerCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateServerCertificateRequest) SetTags(v []*CreateServerCertificateRequestTags) *CreateServerCertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateServerCertificateRequest) SetYears(v int32) *CreateServerCertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateServerCertificateRequest) Validate() error {
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

type CreateServerCertificateRequestTags struct {
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
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServerCertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateServerCertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateServerCertificateRequestTags) SetKey(v string) *CreateServerCertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateServerCertificateRequestTags) SetValue(v string) *CreateServerCertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateServerCertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
