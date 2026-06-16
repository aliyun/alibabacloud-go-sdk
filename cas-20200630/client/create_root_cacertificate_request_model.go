// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRootCACertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *CreateRootCACertificateRequest
	GetAlgorithm() *string
	SetClientToken(v string) *CreateRootCACertificateRequest
	GetClientToken() *string
	SetCommonName(v string) *CreateRootCACertificateRequest
	GetCommonName() *string
	SetCountryCode(v string) *CreateRootCACertificateRequest
	GetCountryCode() *string
	SetLocality(v string) *CreateRootCACertificateRequest
	GetLocality() *string
	SetOrganization(v string) *CreateRootCACertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateRootCACertificateRequest
	GetOrganizationUnit() *string
	SetResourceGroupId(v string) *CreateRootCACertificateRequest
	GetResourceGroupId() *string
	SetState(v string) *CreateRootCACertificateRequest
	GetState() *string
	SetTags(v []*CreateRootCACertificateRequestTags) *CreateRootCACertificateRequest
	GetTags() []*CreateRootCACertificateRequestTags
	SetYears(v int32) *CreateRootCACertificateRequest
	GetYears() *int32
}

type CreateRootCACertificateRequest struct {
	// The key algorithm of the root CA certificate. The key algorithm is in the `<encryption algorithm>_<key length>` format. Valid values:
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
	// The encryption algorithm of the root CA certificate must be the same as the **Certificate Algorithm*	- of the private root CA that you purchased. For example, if you set **Certificate Algorithm*	- to **RSA*	- when you purchase a private root CA, the key algorithm of the root CA certificate must be **RSA_1024**, **RSA_2048**, or **RSA_4096**.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// A client token to ensure the idempotence of the request.
	//
	// Generate a unique value for this parameter from your client. The token supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 3838B684-3075-582B-9A45-8C99104029DF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The common name or abbreviation of the organization. Supports Chinese characters and letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The two-letter uppercase code of the country or region where the organization is located. For example, **CN*	- indicates China and **US*	- indicates the United States.
	//
	// For more information about country codes, see the **Country codes*	- section in [Manage company information](https://help.aliyun.com/document_detail/198289.html).
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The name of the city where the organization is located. Supports Chinese characters and letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization for the root CA certificate. This is typically your company or enterprise name. Supports Chinese characters and letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. Supports Chinese characters and letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// <props="intl">The name of the province or state where the organization is located. Supports Chinese characters and letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags.
	Tags []*CreateRootCACertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the root CA certificate. Unit: years.
	//
	// > Set the validity period to 5 to 10 years.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateRootCACertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRootCACertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateRootCACertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRootCACertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateRootCACertificateRequest) GetCountryCode() *string {
	return s.CountryCode
}

func (s *CreateRootCACertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateRootCACertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateRootCACertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateRootCACertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRootCACertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateRootCACertificateRequest) GetTags() []*CreateRootCACertificateRequestTags {
	return s.Tags
}

func (s *CreateRootCACertificateRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateRootCACertificateRequest) SetAlgorithm(v string) *CreateRootCACertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetClientToken(v string) *CreateRootCACertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetCommonName(v string) *CreateRootCACertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetCountryCode(v string) *CreateRootCACertificateRequest {
	s.CountryCode = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetLocality(v string) *CreateRootCACertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetOrganization(v string) *CreateRootCACertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetOrganizationUnit(v string) *CreateRootCACertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetResourceGroupId(v string) *CreateRootCACertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetState(v string) *CreateRootCACertificateRequest {
	s.State = &v
	return s
}

func (s *CreateRootCACertificateRequest) SetTags(v []*CreateRootCACertificateRequestTags) *CreateRootCACertificateRequest {
	s.Tags = v
	return s
}

func (s *CreateRootCACertificateRequest) SetYears(v int32) *CreateRootCACertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateRootCACertificateRequest) Validate() error {
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

type CreateRootCACertificateRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// runtime
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRootCACertificateRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateRootCACertificateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateRootCACertificateRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateRootCACertificateRequestTags) SetKey(v string) *CreateRootCACertificateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateRootCACertificateRequestTags) SetValue(v string) *CreateRootCACertificateRequestTags {
	s.Value = &v
	return s
}

func (s *CreateRootCACertificateRequestTags) Validate() error {
	return dara.Validate(s)
}
