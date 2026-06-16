// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListCertResponseBody
	GetCurrentPage() *int32
	SetList(v []*ListCertResponseBodyList) *ListCertResponseBody
	GetList() []*ListCertResponseBodyList
	SetMaxResults(v int32) *ListCertResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCertResponseBody
	GetNextToken() *string
	SetPageCount(v int32) *ListCertResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *ListCertResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCertResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListCertResponseBody
	GetTotalCount() *int64
}

type ListCertResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of certificates.
	List []*ListCertResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A token to retrieve the next page of results. If this value is empty, all results have been returned.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The page size.
	//
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of certificates.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCertResponseBody) GetList() []*ListCertResponseBodyList {
	return s.List
}

func (s *ListCertResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCertResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCertResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCertResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertResponseBody) SetCurrentPage(v int32) *ListCertResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertResponseBody) SetList(v []*ListCertResponseBodyList) *ListCertResponseBody {
	s.List = v
	return s
}

func (s *ListCertResponseBody) SetMaxResults(v int32) *ListCertResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCertResponseBody) SetNextToken(v string) *ListCertResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCertResponseBody) SetPageCount(v int32) *ListCertResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListCertResponseBody) SetRequestId(v string) *ListCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertResponseBody) SetShowSize(v int32) *ListCertResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertResponseBody) SetTotalCount(v int64) *ListCertResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertResponseBodyList struct {
	// The expiration time of the certificate.
	//
	// example:
	//
	// 2024-05-13 12:59:45
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The expiration time of the client certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must be both left empty or both specified.
	//
	// example:
	//
	// 1728921600000
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The public key algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The alias of the certificate.
	//
	// example:
	//
	// test
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The issuance time of the certificate.
	//
	// example:
	//
	// 2026-05-19
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// > The **BeforeTime*	- and **AfterTime*	- parameters must be both left empty or both specified.
	//
	// example:
	//
	// 1728921600000
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - `free`: Free certificate.
	//
	// - `cas`: Alibaba Cloud Security certificate.
	//
	// - `upload`: A user-uploaded certificate.
	//
	// example:
	//
	// Server
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The primary domain name of the certificate.
	//
	// example:
	//
	// www.kfsjn.xyz
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// A unique, user-defined identifier for the certificate.
	//
	// example:
	//
	// ***b86sca4384811e0b5e8707e68***
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// A JSON string containing extended attributes.
	//
	// example:
	//
	// {"appId":"APP_PFHMIGUHKDUW6S3N7ZL2"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The ID of the data source to which the certificate order belongs.
	//
	// example:
	//
	// 1806958
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 1ef539a8-1e1f-6b88-8c11-21cf01a203e9
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// Specifies if the private key is exportable. Valid values:
	//
	// - `true`: The private key is exportable.
	//
	// - `false`: The private key is not exportable.
	//
	// example:
	//
	// true
	KeyExportable *bool `json:"KeyExportable,omitempty" xml:"KeyExportable,omitempty"`
	// The organization specified in the certificate.
	//
	// example:
	//
	// test
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The organizational unit (OU) specified in the certificate.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The certificate serial number.
	//
	// example:
	//
	// 3a3ee3c3597d675e
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - `ISSUE`: Issued.
	//
	// - `REVOKE`: Revoked.
	//
	// example:
	//
	// complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) of the certificate subject.
	//
	// example:
	//
	// SubjectDn
	SubjectDn *string `json:"SubjectDn,omitempty" xml:"SubjectDn,omitempty"`
	// The tags of the certificate.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListCertResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListCertResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListCertResponseBodyList) GetAfterDate() *string {
	return s.AfterDate
}

func (s *ListCertResponseBodyList) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *ListCertResponseBodyList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCertResponseBodyList) GetAliasName() *string {
	return s.AliasName
}

func (s *ListCertResponseBodyList) GetBeforeDate() *string {
	return s.BeforeDate
}

func (s *ListCertResponseBodyList) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *ListCertResponseBodyList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListCertResponseBodyList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertResponseBodyList) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *ListCertResponseBodyList) GetExtra() *string {
	return s.Extra
}

func (s *ListCertResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *ListCertResponseBodyList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListCertResponseBodyList) GetKeyExportable() *bool {
	return s.KeyExportable
}

func (s *ListCertResponseBodyList) GetOrganization() *string {
	return s.Organization
}

func (s *ListCertResponseBodyList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *ListCertResponseBodyList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListCertResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *ListCertResponseBodyList) GetSubjectDn() *string {
	return s.SubjectDn
}

func (s *ListCertResponseBodyList) GetTags() []*string {
	return s.Tags
}

func (s *ListCertResponseBodyList) SetAfterDate(v string) *ListCertResponseBodyList {
	s.AfterDate = &v
	return s
}

func (s *ListCertResponseBodyList) SetAfterTime(v int64) *ListCertResponseBodyList {
	s.AfterTime = &v
	return s
}

func (s *ListCertResponseBodyList) SetAlgorithm(v string) *ListCertResponseBodyList {
	s.Algorithm = &v
	return s
}

func (s *ListCertResponseBodyList) SetAliasName(v string) *ListCertResponseBodyList {
	s.AliasName = &v
	return s
}

func (s *ListCertResponseBodyList) SetBeforeDate(v string) *ListCertResponseBodyList {
	s.BeforeDate = &v
	return s
}

func (s *ListCertResponseBodyList) SetBeforeTime(v int64) *ListCertResponseBodyList {
	s.BeforeTime = &v
	return s
}

func (s *ListCertResponseBodyList) SetCertificateType(v string) *ListCertResponseBodyList {
	s.CertificateType = &v
	return s
}

func (s *ListCertResponseBodyList) SetCommonName(v string) *ListCertResponseBodyList {
	s.CommonName = &v
	return s
}

func (s *ListCertResponseBodyList) SetCustomIdentifier(v string) *ListCertResponseBodyList {
	s.CustomIdentifier = &v
	return s
}

func (s *ListCertResponseBodyList) SetExtra(v string) *ListCertResponseBodyList {
	s.Extra = &v
	return s
}

func (s *ListCertResponseBodyList) SetId(v int64) *ListCertResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListCertResponseBodyList) SetIdentifier(v string) *ListCertResponseBodyList {
	s.Identifier = &v
	return s
}

func (s *ListCertResponseBodyList) SetKeyExportable(v bool) *ListCertResponseBodyList {
	s.KeyExportable = &v
	return s
}

func (s *ListCertResponseBodyList) SetOrganization(v string) *ListCertResponseBodyList {
	s.Organization = &v
	return s
}

func (s *ListCertResponseBodyList) SetOrganizationUnit(v string) *ListCertResponseBodyList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListCertResponseBodyList) SetSerialNumber(v string) *ListCertResponseBodyList {
	s.SerialNumber = &v
	return s
}

func (s *ListCertResponseBodyList) SetStatus(v string) *ListCertResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListCertResponseBodyList) SetSubjectDn(v string) *ListCertResponseBodyList {
	s.SubjectDn = &v
	return s
}

func (s *ListCertResponseBodyList) SetTags(v []*string) *ListCertResponseBodyList {
	s.Tags = v
	return s
}

func (s *ListCertResponseBodyList) Validate() error {
	return dara.Validate(s)
}
