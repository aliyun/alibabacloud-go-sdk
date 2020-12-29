// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateDVOrderAuditRequest struct {
	SourceIp         *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Domain           *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainVerifyType *string `json:"DomainVerifyType,omitempty" xml:"DomainVerifyType,omitempty"`
	Username         *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile           *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Province         *string `json:"Province,omitempty" xml:"Province,omitempty"`
	City             *string `json:"City,omitempty" xml:"City,omitempty"`
}

func (s CreateDVOrderAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDVOrderAuditRequest) GoString() string {
	return s.String()
}

func (s *CreateDVOrderAuditRequest) SetSourceIp(v string) *CreateDVOrderAuditRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetLang(v string) *CreateDVOrderAuditRequest {
	s.Lang = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetInstanceId(v string) *CreateDVOrderAuditRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetDomain(v string) *CreateDVOrderAuditRequest {
	s.Domain = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetDomainVerifyType(v string) *CreateDVOrderAuditRequest {
	s.DomainVerifyType = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetUsername(v string) *CreateDVOrderAuditRequest {
	s.Username = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetEmail(v string) *CreateDVOrderAuditRequest {
	s.Email = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetMobile(v string) *CreateDVOrderAuditRequest {
	s.Mobile = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetProvince(v string) *CreateDVOrderAuditRequest {
	s.Province = &v
	return s
}

func (s *CreateDVOrderAuditRequest) SetCity(v string) *CreateDVOrderAuditRequest {
	s.City = &v
	return s
}

type CreateDVOrderAuditResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDVOrderAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDVOrderAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDVOrderAuditResponseBody) SetRequestId(v string) *CreateDVOrderAuditResponseBody {
	s.RequestId = &v
	return s
}

type CreateDVOrderAuditResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDVOrderAuditResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDVOrderAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDVOrderAuditResponse) GoString() string {
	return s.String()
}

func (s *CreateDVOrderAuditResponse) SetHeaders(v map[string]*string) *CreateDVOrderAuditResponse {
	s.Headers = v
	return s
}

func (s *CreateDVOrderAuditResponse) SetBody(v *CreateDVOrderAuditResponseBody) *CreateDVOrderAuditResponse {
	s.Body = v
	return s
}

type CreateUserCertificateRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Cert     *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s CreateUserCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateUserCertificateRequest) SetName(v string) *CreateUserCertificateRequest {
	s.Name = &v
	return s
}

func (s *CreateUserCertificateRequest) SetCert(v string) *CreateUserCertificateRequest {
	s.Cert = &v
	return s
}

func (s *CreateUserCertificateRequest) SetKey(v string) *CreateUserCertificateRequest {
	s.Key = &v
	return s
}

func (s *CreateUserCertificateRequest) SetSourceIp(v string) *CreateUserCertificateRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateUserCertificateRequest) SetLang(v string) *CreateUserCertificateRequest {
	s.Lang = &v
	return s
}

type CreateUserCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CertId    *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
}

func (s CreateUserCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserCertificateResponseBody) SetRequestId(v string) *CreateUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserCertificateResponseBody) SetCertId(v int64) *CreateUserCertificateResponseBody {
	s.CertId = &v
	return s
}

type CreateUserCertificateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateUserCertificateResponse) SetHeaders(v map[string]*string) *CreateUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateUserCertificateResponse) SetBody(v *CreateUserCertificateResponseBody) *CreateUserCertificateResponse {
	s.Body = v
	return s
}

type DeleteUserCertificateRequest struct {
	CertId   *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteUserCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateRequest) SetCertId(v int64) *DeleteUserCertificateRequest {
	s.CertId = &v
	return s
}

func (s *DeleteUserCertificateRequest) SetSourceIp(v string) *DeleteUserCertificateRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserCertificateRequest) SetLang(v string) *DeleteUserCertificateRequest {
	s.Lang = &v
	return s
}

type DeleteUserCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponseBody) SetRequestId(v string) *DeleteUserCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserCertificateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponse) SetHeaders(v map[string]*string) *DeleteUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserCertificateResponse) SetBody(v *DeleteUserCertificateResponseBody) *DeleteUserCertificateResponse {
	s.Body = v
	return s
}

type DescribeDVOrderResultRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDVOrderResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDVOrderResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDVOrderResultRequest) SetSourceIp(v string) *DescribeDVOrderResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDVOrderResultRequest) SetLang(v string) *DescribeDVOrderResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDVOrderResultRequest) SetInstanceId(v string) *DescribeDVOrderResultRequest {
	s.InstanceId = &v
	return s
}

type DescribeDVOrderResultResponseBody struct {
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	CheckName   *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
	PrivateKey  *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CheckType   *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	CheckValue  *string `json:"CheckValue,omitempty" xml:"CheckValue,omitempty"`
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
}

func (s DescribeDVOrderResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDVOrderResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDVOrderResultResponseBody) SetOrderStatus(v string) *DescribeDVOrderResultResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetCheckName(v string) *DescribeDVOrderResultResponseBody {
	s.CheckName = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetPrivateKey(v string) *DescribeDVOrderResultResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetRequestId(v string) *DescribeDVOrderResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetCheckType(v string) *DescribeDVOrderResultResponseBody {
	s.CheckType = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetCheckValue(v string) *DescribeDVOrderResultResponseBody {
	s.CheckValue = &v
	return s
}

func (s *DescribeDVOrderResultResponseBody) SetCertificate(v string) *DescribeDVOrderResultResponseBody {
	s.Certificate = &v
	return s
}

type DescribeDVOrderResultResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDVOrderResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDVOrderResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDVOrderResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDVOrderResultResponse) SetHeaders(v map[string]*string) *DescribeDVOrderResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDVOrderResultResponse) SetBody(v *DescribeDVOrderResultResponseBody) *DescribeDVOrderResultResponse {
	s.Body = v
	return s
}

type DescribeOrderInstanceListRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StartIndex *int32  `json:"StartIndex,omitempty" xml:"StartIndex,omitempty"`
}

func (s DescribeOrderInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrderInstanceListRequest) SetSourceIp(v string) *DescribeOrderInstanceListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOrderInstanceListRequest) SetLang(v string) *DescribeOrderInstanceListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOrderInstanceListRequest) SetStartIndex(v int32) *DescribeOrderInstanceListRequest {
	s.StartIndex = &v
	return s
}

type DescribeOrderInstanceListResponseBody struct {
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderList  []*DescribeOrderInstanceListResponseBodyOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
}

func (s DescribeOrderInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrderInstanceListResponseBody) SetTotalCount(v int32) *DescribeOrderInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBody) SetRequestId(v string) *DescribeOrderInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBody) SetOrderList(v []*DescribeOrderInstanceListResponseBodyOrderList) *DescribeOrderInstanceListResponseBody {
	s.OrderList = v
	return s
}

type DescribeOrderInstanceListResponseBodyOrderList struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CertType   *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Source     *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeOrderInstanceListResponseBodyOrderList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInstanceListResponseBodyOrderList) GoString() string {
	return s.String()
}

func (s *DescribeOrderInstanceListResponseBodyOrderList) SetStatus(v string) *DescribeOrderInstanceListResponseBodyOrderList {
	s.Status = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBodyOrderList) SetCertType(v string) *DescribeOrderInstanceListResponseBodyOrderList {
	s.CertType = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBodyOrderList) SetInstanceId(v string) *DescribeOrderInstanceListResponseBodyOrderList {
	s.InstanceId = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBodyOrderList) SetSource(v string) *DescribeOrderInstanceListResponseBodyOrderList {
	s.Source = &v
	return s
}

func (s *DescribeOrderInstanceListResponseBodyOrderList) SetId(v int64) *DescribeOrderInstanceListResponseBodyOrderList {
	s.Id = &v
	return s
}

type DescribeOrderInstanceListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOrderInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOrderInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrderInstanceListResponse) SetHeaders(v map[string]*string) *DescribeOrderInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrderInstanceListResponse) SetBody(v *DescribeOrderInstanceListResponseBody) *DescribeOrderInstanceListResponse {
	s.Body = v
	return s
}

type DescribeUserCertificateDetailRequest struct {
	CertId   *int64  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserCertificateDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateDetailRequest) SetCertId(v int64) *DescribeUserCertificateDetailRequest {
	s.CertId = &v
	return s
}

func (s *DescribeUserCertificateDetailRequest) SetSourceIp(v string) *DescribeUserCertificateDetailRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserCertificateDetailRequest) SetLang(v string) *DescribeUserCertificateDetailRequest {
	s.Lang = &v
	return s
}

type DescribeUserCertificateDetailResponseBody struct {
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Issuer      *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	Expired     *bool   `json:"Expired,omitempty" xml:"Expired,omitempty"`
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	City        *string `json:"City,omitempty" xml:"City,omitempty"`
	EndDate     *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Province    *string `json:"Province,omitempty" xml:"Province,omitempty"`
	BuyInAliyun *bool   `json:"BuyInAliyun,omitempty" xml:"BuyInAliyun,omitempty"`
	Common      *string `json:"Common,omitempty" xml:"Common,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Sans        *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	Country     *string `json:"Country,omitempty" xml:"Country,omitempty"`
	Cert        *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeUserCertificateDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateDetailResponseBody) SetFingerprint(v string) *DescribeUserCertificateDetailResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetRequestId(v string) *DescribeUserCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetIssuer(v string) *DescribeUserCertificateDetailResponseBody {
	s.Issuer = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetExpired(v bool) *DescribeUserCertificateDetailResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetOrgName(v string) *DescribeUserCertificateDetailResponseBody {
	s.OrgName = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetCity(v string) *DescribeUserCertificateDetailResponseBody {
	s.City = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetEndDate(v string) *DescribeUserCertificateDetailResponseBody {
	s.EndDate = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetProvince(v string) *DescribeUserCertificateDetailResponseBody {
	s.Province = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetBuyInAliyun(v bool) *DescribeUserCertificateDetailResponseBody {
	s.BuyInAliyun = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetCommon(v string) *DescribeUserCertificateDetailResponseBody {
	s.Common = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetName(v string) *DescribeUserCertificateDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetStartDate(v string) *DescribeUserCertificateDetailResponseBody {
	s.StartDate = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetSans(v string) *DescribeUserCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetCountry(v string) *DescribeUserCertificateDetailResponseBody {
	s.Country = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetCert(v string) *DescribeUserCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetId(v int64) *DescribeUserCertificateDetailResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeUserCertificateDetailResponseBody) SetKey(v string) *DescribeUserCertificateDetailResponseBody {
	s.Key = &v
	return s
}

type DescribeUserCertificateDetailResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserCertificateDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeUserCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCertificateDetailResponse) SetBody(v *DescribeUserCertificateDetailResponseBody) *DescribeUserCertificateDetailResponse {
	s.Body = v
	return s
}

type DescribeUserCertificateListRequest struct {
	ShowSize    *int32  `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeUserCertificateListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateListRequest) SetShowSize(v int32) *DescribeUserCertificateListRequest {
	s.ShowSize = &v
	return s
}

func (s *DescribeUserCertificateListRequest) SetCurrentPage(v int32) *DescribeUserCertificateListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserCertificateListRequest) SetSourceIp(v string) *DescribeUserCertificateListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserCertificateListRequest) SetLang(v string) *DescribeUserCertificateListRequest {
	s.Lang = &v
	return s
}

type DescribeUserCertificateListResponseBody struct {
	TotalCount      *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CurrentPage     *int32                                                    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CertificateList []*DescribeUserCertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	ShowSize        *int32                                                    `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s DescribeUserCertificateListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateListResponseBody) SetTotalCount(v int32) *DescribeUserCertificateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserCertificateListResponseBody) SetRequestId(v string) *DescribeUserCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserCertificateListResponseBody) SetCurrentPage(v int32) *DescribeUserCertificateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserCertificateListResponseBody) SetCertificateList(v []*DescribeUserCertificateListResponseBodyCertificateList) *DescribeUserCertificateListResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribeUserCertificateListResponseBody) SetShowSize(v int32) *DescribeUserCertificateListResponseBody {
	s.ShowSize = &v
	return s
}

type DescribeUserCertificateListResponseBodyCertificateList struct {
	StartDate   *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Province    *string `json:"province,omitempty" xml:"province,omitempty"`
	Sans        *string `json:"sans,omitempty" xml:"sans,omitempty"`
	Common      *string `json:"common,omitempty" xml:"common,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Country     *string `json:"country,omitempty" xml:"country,omitempty"`
	Issuer      *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	BuyInAliyun *bool   `json:"buyInAliyun,omitempty" xml:"buyInAliyun,omitempty"`
	EndDate     *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	Expired     *bool   `json:"expired,omitempty" xml:"expired,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	City        *string `json:"city,omitempty" xml:"city,omitempty"`
	OrgName     *string `json:"orgName,omitempty" xml:"orgName,omitempty"`
}

func (s DescribeUserCertificateListResponseBodyCertificateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateListResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetStartDate(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.StartDate = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetProvince(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Province = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetSans(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetCommon(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Common = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetId(v int64) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Id = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetCountry(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Country = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetIssuer(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Issuer = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetBuyInAliyun(v bool) *DescribeUserCertificateListResponseBodyCertificateList {
	s.BuyInAliyun = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetEndDate(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.EndDate = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetExpired(v bool) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Expired = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetFingerprint(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Fingerprint = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetName(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.Name = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetCity(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.City = &v
	return s
}

func (s *DescribeUserCertificateListResponseBodyCertificateList) SetOrgName(v string) *DescribeUserCertificateListResponseBodyCertificateList {
	s.OrgName = &v
	return s
}

type DescribeUserCertificateListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserCertificateListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateListResponse) SetHeaders(v map[string]*string) *DescribeUserCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCertificateListResponse) SetBody(v *DescribeUserCertificateListResponseBody) *DescribeUserCertificateListResponse {
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

func (client *Client) CreateDVOrderAuditWithOptions(request *CreateDVOrderAuditRequest, runtime *util.RuntimeOptions) (_result *CreateDVOrderAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDVOrderAuditResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDVOrderAudit"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDVOrderAudit(request *CreateDVOrderAuditRequest) (_result *CreateDVOrderAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDVOrderAuditResponse{}
	_body, _err := client.CreateDVOrderAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserCertificateWithOptions(request *CreateUserCertificateRequest, runtime *util.RuntimeOptions) (_result *CreateUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserCertificate"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserCertificate(request *CreateUserCertificateRequest) (_result *CreateUserCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserCertificateResponse{}
	_body, _err := client.CreateUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserCertificateWithOptions(request *DeleteUserCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteUserCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserCertificate"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserCertificate(request *DeleteUserCertificateRequest) (_result *DeleteUserCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserCertificateResponse{}
	_body, _err := client.DeleteUserCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDVOrderResultWithOptions(request *DescribeDVOrderResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDVOrderResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDVOrderResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDVOrderResult"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDVOrderResult(request *DescribeDVOrderResultRequest) (_result *DescribeDVOrderResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDVOrderResultResponse{}
	_body, _err := client.DescribeDVOrderResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOrderInstanceListWithOptions(request *DescribeOrderInstanceListRequest, runtime *util.RuntimeOptions) (_result *DescribeOrderInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOrderInstanceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOrderInstanceList"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOrderInstanceList(request *DescribeOrderInstanceListRequest) (_result *DescribeOrderInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrderInstanceListResponse{}
	_body, _err := client.DescribeOrderInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserCertificateDetailWithOptions(request *DescribeUserCertificateDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeUserCertificateDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserCertificateDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserCertificateDetail"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserCertificateDetail(request *DescribeUserCertificateDetailRequest) (_result *DescribeUserCertificateDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserCertificateDetailResponse{}
	_body, _err := client.DescribeUserCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserCertificateListWithOptions(request *DescribeUserCertificateListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserCertificateListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserCertificateListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserCertificateList"), tea.String("2018-07-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserCertificateList(request *DescribeUserCertificateListRequest) (_result *DescribeUserCertificateListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserCertificateListResponse{}
	_body, _err := client.DescribeUserCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
