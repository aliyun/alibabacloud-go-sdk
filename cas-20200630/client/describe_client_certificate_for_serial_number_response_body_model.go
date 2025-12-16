// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateForSerialNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList) *DescribeClientCertificateForSerialNumberResponseBody
	GetCertificateList() []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList
	SetRequestId(v string) *DescribeClientCertificateForSerialNumberResponseBody
	GetRequestId() *string
}

type DescribeClientCertificateForSerialNumberResponseBody struct {
	CertificateList []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) GetCertificateList() []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	return s.CertificateList
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) SetCertificateList(v []*DescribeClientCertificateForSerialNumberResponseBodyCertificateList) *DescribeClientCertificateForSerialNumberResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) SetRequestId(v string) *DescribeClientCertificateForSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBody) Validate() error {
	if s.CertificateList != nil {
		for _, item := range s.CertificateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClientCertificateForSerialNumberResponseBodyCertificateList struct {
	// example:
	//
	// 2022-08-23T16:15Z
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 2021-10-28T16:15Z
	BeforeDate      *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// 4096
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Md5          *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// ISSUE
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE-----  ...... -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeClientCertificateForSerialNumberResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetAfterDate() *string {
	return s.AfterDate
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetBeforeDate() *string {
	return s.BeforeDate
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetLocality() *string {
	return s.Locality
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetMd5() *string {
	return s.Md5
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetOrganization() *string {
	return s.Organization
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetSans() *string {
	return s.Sans
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetSha2() *string {
	return s.Sha2
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetState() *string {
	return s.State
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) GetYears() *int32 {
	return s.Years
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetAfterDate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetAlgorithm(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetBeforeDate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetCertificateType(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetCommonName(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetCountryCode(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetIdentifier(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetKeySize(v int32) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetLocality(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetMd5(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetOrganization(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetOrganizationUnit(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetParentIdentifier(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSans(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSerialNumber(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSha2(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSignAlgorithm(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetState(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetStatus(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetSubjectDN(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetX509Certificate(v string) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) SetYears(v int32) *DescribeClientCertificateForSerialNumberResponseBodyCertificateList {
	s.Years = &v
	return s
}

func (s *DescribeClientCertificateForSerialNumberResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
