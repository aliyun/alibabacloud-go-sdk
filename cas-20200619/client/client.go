// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateSSLCertificateRequest struct {
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s CreateSSLCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateRequest) SetCertificate(v string) *CreateSSLCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *CreateSSLCertificateRequest) SetPrivateKey(v string) *CreateSSLCertificateRequest {
	s.PrivateKey = &v
	return s
}

type CreateSSLCertificateResponseBody struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSSLCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateResponseBody) SetCertIdentifier(v string) *CreateSSLCertificateResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *CreateSSLCertificateResponseBody) SetRequestId(v string) *CreateSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

type CreateSSLCertificateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSSLCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateResponse) SetHeaders(v map[string]*string) *CreateSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateSSLCertificateResponse) SetBody(v *CreateSSLCertificateResponseBody) *CreateSSLCertificateResponse {
	s.Body = v
	return s
}

type CreateSSLCertificateWithNameRequest struct {
	CertName    *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s CreateSSLCertificateWithNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateWithNameRequest) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateWithNameRequest) SetCertName(v string) *CreateSSLCertificateWithNameRequest {
	s.CertName = &v
	return s
}

func (s *CreateSSLCertificateWithNameRequest) SetCertificate(v string) *CreateSSLCertificateWithNameRequest {
	s.Certificate = &v
	return s
}

func (s *CreateSSLCertificateWithNameRequest) SetPrivateKey(v string) *CreateSSLCertificateWithNameRequest {
	s.PrivateKey = &v
	return s
}

type CreateSSLCertificateWithNameResponseBody struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSSLCertificateWithNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateWithNameResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateWithNameResponseBody) SetCertIdentifier(v string) *CreateSSLCertificateWithNameResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *CreateSSLCertificateWithNameResponseBody) SetRequestId(v string) *CreateSSLCertificateWithNameResponseBody {
	s.RequestId = &v
	return s
}

type CreateSSLCertificateWithNameResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSSLCertificateWithNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSSLCertificateWithNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSSLCertificateWithNameResponse) GoString() string {
	return s.String()
}

func (s *CreateSSLCertificateWithNameResponse) SetHeaders(v map[string]*string) *CreateSSLCertificateWithNameResponse {
	s.Headers = v
	return s
}

func (s *CreateSSLCertificateWithNameResponse) SetBody(v *CreateSSLCertificateWithNameResponseBody) *CreateSSLCertificateWithNameResponse {
	s.Body = v
	return s
}

type DeleteSSLCertificateRequest struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
}

func (s DeleteSSLCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSSLCertificateRequest) SetCertIdentifier(v string) *DeleteSSLCertificateRequest {
	s.CertIdentifier = &v
	return s
}

type DeleteSSLCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSSLCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSSLCertificateResponseBody) SetRequestId(v string) *DeleteSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSSLCertificateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSSLCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteSSLCertificateResponse) SetHeaders(v map[string]*string) *DeleteSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteSSLCertificateResponse) SetBody(v *DeleteSSLCertificateResponseBody) *DeleteSSLCertificateResponse {
	s.Body = v
	return s
}

type DescribeSSLCertificateCountRequest struct {
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
}

func (s DescribeSSLCertificateCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateCountRequest) SetSearchValue(v string) *DescribeSSLCertificateCountRequest {
	s.SearchValue = &v
	return s
}

type DescribeSSLCertificateCountResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSSLCertificateCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateCountResponseBody) SetRequestId(v string) *DescribeSSLCertificateCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSSLCertificateCountResponseBody) SetTotalCount(v int32) *DescribeSSLCertificateCountResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSSLCertificateCountResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSSLCertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSSLCertificateCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateCountResponse) SetHeaders(v map[string]*string) *DescribeSSLCertificateCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSSLCertificateCountResponse) SetBody(v *DescribeSSLCertificateCountResponseBody) *DescribeSSLCertificateCountResponse {
	s.Body = v
	return s
}

type DescribeSSLCertificateListRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	SearchValue *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribeSSLCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateListRequest) SetCurrentPage(v int32) *DescribeSSLCertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSSLCertificateListRequest) SetSearchValue(v string) *DescribeSSLCertificateListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeSSLCertificateListRequest) SetShowSize(v int32) *DescribeSSLCertificateListRequest {
	s.ShowSize = &v
	return s
}

type DescribeSSLCertificateListResponseBody struct {
	CertMetaList []*DescribeSSLCertificateListResponseBodyCertMetaList `json:"CertMetaList,omitempty" xml:"CertMetaList,omitempty" type:"Repeated"`
	CurrentPage  *int32                                                `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount    *int32                                                `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize     *int32                                                `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount   *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSSLCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateListResponseBody) SetCertMetaList(v []*DescribeSSLCertificateListResponseBodyCertMetaList) *DescribeSSLCertificateListResponseBody {
	s.CertMetaList = v
	return s
}

func (s *DescribeSSLCertificateListResponseBody) SetCurrentPage(v int32) *DescribeSSLCertificateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBody) SetPageCount(v int32) *DescribeSSLCertificateListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBody) SetRequestId(v string) *DescribeSSLCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBody) SetShowSize(v int32) *DescribeSSLCertificateListResponseBody {
	s.ShowSize = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBody) SetTotalCount(v int32) *DescribeSSLCertificateListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSSLCertificateListResponseBodyCertMetaList struct {
	AfterDate      *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate     *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CommonName     *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Issuer         *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	KeySize        *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Md5            *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Sans           *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNo       *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Sha2           *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm  *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
}

func (s DescribeSSLCertificateListResponseBodyCertMetaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateListResponseBodyCertMetaList) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetAfterDate(v int64) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.AfterDate = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetAlgorithm(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.Algorithm = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetBeforeDate(v int64) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetCertIdentifier(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetCertName(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.CertName = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetCommonName(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.CommonName = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetIssuer(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.Issuer = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetKeySize(v int32) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.KeySize = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetMd5(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.Md5 = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetSans(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.Sans = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetSerialNo(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.SerialNo = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetSha2(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.Sha2 = &v
	return s
}

func (s *DescribeSSLCertificateListResponseBodyCertMetaList) SetSignAlgorithm(v string) *DescribeSSLCertificateListResponseBodyCertMetaList {
	s.SignAlgorithm = &v
	return s
}

type DescribeSSLCertificateListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSSLCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSSLCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateListResponse) SetHeaders(v map[string]*string) *DescribeSSLCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSSLCertificateListResponse) SetBody(v *DescribeSSLCertificateListResponseBody) *DescribeSSLCertificateListResponse {
	s.Body = v
	return s
}

type DescribeSSLCertificateMatchDomainListRequest struct {
	Algorithm   *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	MatchType   *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribeSSLCertificateMatchDomainListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateMatchDomainListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateMatchDomainListRequest) SetAlgorithm(v string) *DescribeSSLCertificateMatchDomainListRequest {
	s.Algorithm = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListRequest) SetCurrentPage(v int32) *DescribeSSLCertificateMatchDomainListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListRequest) SetDomain(v string) *DescribeSSLCertificateMatchDomainListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListRequest) SetMatchType(v string) *DescribeSSLCertificateMatchDomainListRequest {
	s.MatchType = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListRequest) SetShowSize(v int32) *DescribeSSLCertificateMatchDomainListRequest {
	s.ShowSize = &v
	return s
}

type DescribeSSLCertificateMatchDomainListResponseBody struct {
	CertMetaList []*DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList `json:"CertMetaList,omitempty" xml:"CertMetaList,omitempty" type:"Repeated"`
	CurrentPage  *int32                                                           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageCount    *int32                                                           `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	RequestId    *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShowSize     *int32                                                           `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	TotalCount   *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSSLCertificateMatchDomainListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateMatchDomainListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetCertMetaList(v []*DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.CertMetaList = v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetCurrentPage(v int32) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetPageCount(v int32) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetRequestId(v string) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetShowSize(v int32) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.ShowSize = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBody) SetTotalCount(v int32) *DescribeSSLCertificateMatchDomainListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList struct {
	AfterDate       *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm       *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate      *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertIdentifier  *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CommonName      *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	DomainMatchCert *bool   `json:"DomainMatchCert,omitempty" xml:"DomainMatchCert,omitempty"`
	Issuer          *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	KeySize         *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Sans            *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNo        *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Sha2            *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm   *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
}

func (s DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetAfterDate(v int64) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.AfterDate = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetAlgorithm(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.Algorithm = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetBeforeDate(v int64) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetCertIdentifier(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetCertName(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.CertName = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetCommonName(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.CommonName = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetDomainMatchCert(v bool) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.DomainMatchCert = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetIssuer(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.Issuer = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetKeySize(v int32) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.KeySize = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetMd5(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.Md5 = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetSans(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.Sans = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetSerialNo(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.SerialNo = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetSha2(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.Sha2 = &v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList) SetSignAlgorithm(v string) *DescribeSSLCertificateMatchDomainListResponseBodyCertMetaList {
	s.SignAlgorithm = &v
	return s
}

type DescribeSSLCertificateMatchDomainListResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSSLCertificateMatchDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSSLCertificateMatchDomainListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificateMatchDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificateMatchDomainListResponse) SetHeaders(v map[string]*string) *DescribeSSLCertificateMatchDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSSLCertificateMatchDomainListResponse) SetBody(v *DescribeSSLCertificateMatchDomainListResponseBody) *DescribeSSLCertificateMatchDomainListResponse {
	s.Body = v
	return s
}

type DescribeSSLCertificatePrivateKeyRequest struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	EncryptedCode  *string `json:"EncryptedCode,omitempty" xml:"EncryptedCode,omitempty"`
}

func (s DescribeSSLCertificatePrivateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePrivateKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePrivateKeyRequest) SetCertIdentifier(v string) *DescribeSSLCertificatePrivateKeyRequest {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeSSLCertificatePrivateKeyRequest) SetEncryptedCode(v string) *DescribeSSLCertificatePrivateKeyRequest {
	s.EncryptedCode = &v
	return s
}

type DescribeSSLCertificatePrivateKeyResponseBody struct {
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignPrivateKey    *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	X509PrivateKey    *string `json:"X509PrivateKey,omitempty" xml:"X509PrivateKey,omitempty"`
}

func (s DescribeSSLCertificatePrivateKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePrivateKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePrivateKeyResponseBody) SetEncryptPrivateKey(v string) *DescribeSSLCertificatePrivateKeyResponseBody {
	s.EncryptPrivateKey = &v
	return s
}

func (s *DescribeSSLCertificatePrivateKeyResponseBody) SetRequestId(v string) *DescribeSSLCertificatePrivateKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSSLCertificatePrivateKeyResponseBody) SetSignPrivateKey(v string) *DescribeSSLCertificatePrivateKeyResponseBody {
	s.SignPrivateKey = &v
	return s
}

func (s *DescribeSSLCertificatePrivateKeyResponseBody) SetX509PrivateKey(v string) *DescribeSSLCertificatePrivateKeyResponseBody {
	s.X509PrivateKey = &v
	return s
}

type DescribeSSLCertificatePrivateKeyResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSSLCertificatePrivateKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSSLCertificatePrivateKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePrivateKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePrivateKeyResponse) SetHeaders(v map[string]*string) *DescribeSSLCertificatePrivateKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSSLCertificatePrivateKeyResponse) SetBody(v *DescribeSSLCertificatePrivateKeyResponseBody) *DescribeSSLCertificatePrivateKeyResponse {
	s.Body = v
	return s
}

type DescribeSSLCertificatePublicKeyDetailRequest struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
}

func (s DescribeSSLCertificatePublicKeyDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePublicKeyDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePublicKeyDetailRequest) SetCertIdentifier(v string) *DescribeSSLCertificatePublicKeyDetailRequest {
	s.CertIdentifier = &v
	return s
}

type DescribeSSLCertificatePublicKeyDetailResponseBody struct {
	CertificateInfo    *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo `json:"CertificateInfo,omitempty" xml:"CertificateInfo,omitempty" type:"Struct"`
	EncryptCertificate *string                                                           `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	RequestId          *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignCertificate    *string                                                           `json:"SignCertificate,omitempty" xml:"SignCertificate,omitempty"`
	X509Certificate    *string                                                           `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s DescribeSSLCertificatePublicKeyDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePublicKeyDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBody) SetCertificateInfo(v *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) *DescribeSSLCertificatePublicKeyDetailResponseBody {
	s.CertificateInfo = v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBody) SetEncryptCertificate(v string) *DescribeSSLCertificatePublicKeyDetailResponseBody {
	s.EncryptCertificate = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBody) SetRequestId(v string) *DescribeSSLCertificatePublicKeyDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBody) SetSignCertificate(v string) *DescribeSSLCertificatePublicKeyDetailResponseBody {
	s.SignCertificate = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBody) SetX509Certificate(v string) *DescribeSSLCertificatePublicKeyDetailResponseBody {
	s.X509Certificate = &v
	return s
}

type DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo struct {
	AfterDate      *int64  `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	BeforeDate     *int64  `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertName       *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CommonName     *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Issuer         *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	KeySize        *int32  `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	Md5            *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Sans           *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	SerialNo       *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	Sha2           *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	SignAlgorithm  *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
}

func (s DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetAfterDate(v int64) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.AfterDate = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetAlgorithm(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.Algorithm = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetBeforeDate(v int64) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.BeforeDate = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetCertIdentifier(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetCertName(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.CertName = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetCommonName(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.CommonName = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetIssuer(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.Issuer = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetKeySize(v int32) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.KeySize = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetMd5(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.Md5 = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetSans(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.Sans = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetSerialNo(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.SerialNo = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetSha2(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.Sha2 = &v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo) SetSignAlgorithm(v string) *DescribeSSLCertificatePublicKeyDetailResponseBodyCertificateInfo {
	s.SignAlgorithm = &v
	return s
}

type DescribeSSLCertificatePublicKeyDetailResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSSLCertificatePublicKeyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSSLCertificatePublicKeyDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSSLCertificatePublicKeyDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeSSLCertificatePublicKeyDetailResponse) SetHeaders(v map[string]*string) *DescribeSSLCertificatePublicKeyDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeSSLCertificatePublicKeyDetailResponse) SetBody(v *DescribeSSLCertificatePublicKeyDetailResponseBody) *DescribeSSLCertificatePublicKeyDetailResponse {
	s.Body = v
	return s
}

type UploadSSLCertificateRequest struct {
	CertName           *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	Certificate        *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	EncryptCertificate *string `json:"EncryptCertificate,omitempty" xml:"EncryptCertificate,omitempty"`
	EncryptPrivateKey  *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	PrivateKey         *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	SignCertificate    *string `json:"SignCertificate,omitempty" xml:"SignCertificate,omitempty"`
	SignPrivateKey     *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
}

func (s UploadSSLCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadSSLCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadSSLCertificateRequest) SetCertName(v string) *UploadSSLCertificateRequest {
	s.CertName = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetCertificate(v string) *UploadSSLCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetEncryptCertificate(v string) *UploadSSLCertificateRequest {
	s.EncryptCertificate = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetEncryptPrivateKey(v string) *UploadSSLCertificateRequest {
	s.EncryptPrivateKey = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetPrivateKey(v string) *UploadSSLCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetSignCertificate(v string) *UploadSSLCertificateRequest {
	s.SignCertificate = &v
	return s
}

func (s *UploadSSLCertificateRequest) SetSignPrivateKey(v string) *UploadSSLCertificateRequest {
	s.SignPrivateKey = &v
	return s
}

type UploadSSLCertificateResponseBody struct {
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadSSLCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadSSLCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSSLCertificateResponseBody) SetCertIdentifier(v string) *UploadSSLCertificateResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *UploadSSLCertificateResponseBody) SetRequestId(v string) *UploadSSLCertificateResponseBody {
	s.RequestId = &v
	return s
}

type UploadSSLCertificateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadSSLCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadSSLCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadSSLCertificateResponse) GoString() string {
	return s.String()
}

func (s *UploadSSLCertificateResponse) SetHeaders(v map[string]*string) *UploadSSLCertificateResponse {
	s.Headers = v
	return s
}

func (s *UploadSSLCertificateResponse) SetBody(v *UploadSSLCertificateResponseBody) *UploadSSLCertificateResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("cas.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cas.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cas.aliyuncs.com"),
		"cn-beijing":                  tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cas.aliyuncs.com"),
		"cn-fujian":                   tea.String("cas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cas.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cas.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cas.aliyuncs.com"),
		"eu-west-1":                   tea.String("cas.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cas.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cas.aliyuncs.com"),
		"us-east-1":                   tea.String("cas.aliyuncs.com"),
		"us-west-1":                   tea.String("cas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSSLCertificateWithOptions(request *CreateSSLCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateSSLCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSSLCertificate"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSSLCertificate(request *CreateSSLCertificateRequest) (_result *CreateSSLCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSSLCertificateResponse{}
	_body, _err := client.CreateSSLCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSSLCertificateWithNameWithOptions(request *CreateSSLCertificateWithNameRequest, runtime *util.RuntimeOptions) (_result *CreateSSLCertificateWithNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSSLCertificateWithName"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSSLCertificateWithNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSSLCertificateWithName(request *CreateSSLCertificateWithNameRequest) (_result *CreateSSLCertificateWithNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSSLCertificateWithNameResponse{}
	_body, _err := client.CreateSSLCertificateWithNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSSLCertificateWithOptions(request *DeleteSSLCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteSSLCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSSLCertificate"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSSLCertificate(request *DeleteSSLCertificateRequest) (_result *DeleteSSLCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSSLCertificateResponse{}
	_body, _err := client.DeleteSSLCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSSLCertificateCountWithOptions(request *DescribeSSLCertificateCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSSLCertificateCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSSLCertificateCount"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSSLCertificateCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSSLCertificateCount(request *DescribeSSLCertificateCountRequest) (_result *DescribeSSLCertificateCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSSLCertificateCountResponse{}
	_body, _err := client.DescribeSSLCertificateCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSSLCertificateListWithOptions(request *DescribeSSLCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeSSLCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		query["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSSLCertificateList"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSSLCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSSLCertificateList(request *DescribeSSLCertificateListRequest) (_result *DescribeSSLCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSSLCertificateListResponse{}
	_body, _err := client.DescribeSSLCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSSLCertificateMatchDomainListWithOptions(request *DescribeSSLCertificateMatchDomainListRequest, runtime *util.RuntimeOptions) (_result *DescribeSSLCertificateMatchDomainListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Algorithm)) {
		query["Algorithm"] = request.Algorithm
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.MatchType)) {
		query["MatchType"] = request.MatchType
	}

	if !tea.BoolValue(util.IsUnset(request.ShowSize)) {
		query["ShowSize"] = request.ShowSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSSLCertificateMatchDomainList"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSSLCertificateMatchDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSSLCertificateMatchDomainList(request *DescribeSSLCertificateMatchDomainListRequest) (_result *DescribeSSLCertificateMatchDomainListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSSLCertificateMatchDomainListResponse{}
	_body, _err := client.DescribeSSLCertificateMatchDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSSLCertificatePrivateKeyWithOptions(request *DescribeSSLCertificatePrivateKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeSSLCertificatePrivateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedCode)) {
		query["EncryptedCode"] = request.EncryptedCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSSLCertificatePrivateKey"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSSLCertificatePrivateKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSSLCertificatePrivateKey(request *DescribeSSLCertificatePrivateKeyRequest) (_result *DescribeSSLCertificatePrivateKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSSLCertificatePrivateKeyResponse{}
	_body, _err := client.DescribeSSLCertificatePrivateKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSSLCertificatePublicKeyDetailWithOptions(request *DescribeSSLCertificatePublicKeyDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeSSLCertificatePublicKeyDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertIdentifier)) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSSLCertificatePublicKeyDetail"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSSLCertificatePublicKeyDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSSLCertificatePublicKeyDetail(request *DescribeSSLCertificatePublicKeyDetailRequest) (_result *DescribeSSLCertificatePublicKeyDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSSLCertificatePublicKeyDetailResponse{}
	_body, _err := client.DescribeSSLCertificatePublicKeyDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadSSLCertificateWithOptions(request *UploadSSLCertificateRequest, runtime *util.RuntimeOptions) (_result *UploadSSLCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertName)) {
		query["CertName"] = request.CertName
	}

	if !tea.BoolValue(util.IsUnset(request.Certificate)) {
		query["Certificate"] = request.Certificate
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptCertificate)) {
		query["EncryptCertificate"] = request.EncryptCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptPrivateKey)) {
		query["EncryptPrivateKey"] = request.EncryptPrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.SignCertificate)) {
		query["SignCertificate"] = request.SignCertificate
	}

	if !tea.BoolValue(util.IsUnset(request.SignPrivateKey)) {
		query["SignPrivateKey"] = request.SignPrivateKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadSSLCertificate"),
		Version:     tea.String("2020-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadSSLCertificate(request *UploadSSLCertificateRequest) (_result *UploadSSLCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadSSLCertificateResponse{}
	_body, _err := client.UploadSSLCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
