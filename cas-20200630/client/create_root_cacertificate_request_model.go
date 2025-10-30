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
	// The key algorithm of the root CA certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
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
	// The encryption algorithm of the root CA certificate must be consistent with the **encryption algorithm*	- of the private root CA instance that you purchase. For example, if the **encryption algorithm*	- of the private root CA instance that you purchase is **RSA**, the key algorithm of the root CA certificate must be **RSA_1024**, **RSA_2048**, or **RSA_4096**.
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
	// Alibaba
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country or region in which the organization is located. You can enter an alpha-2 code. For example, you can use **CN*	- to indicate China and use **US*	- to indicate the United States.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The name of the organization that is associated with the root CA certificate. You can enter the name of your enterprise or company. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The value can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhejiang
	State *string                               `json:"State,omitempty" xml:"State,omitempty"`
	Tags  []*CreateRootCACertificateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the root CA certificate. Unit: years.
	//
	// >  We recommend that you set this parameter to a value from 5 to 10.
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
