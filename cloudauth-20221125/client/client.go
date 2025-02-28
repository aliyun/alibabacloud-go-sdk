// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EntElementVerifyRequest struct {
	EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	// example:
	//
	// ENT_2META
	InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
	// example:
	//
	// 370105*****3892
	LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	LegalPersonName   *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// example:
	//
	// 32132***328932
	LicenseNo *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	// example:
	//
	// 32198****193000
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 432***421
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// withdraw
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// 1
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyRequest) GoString() string {
	return s.String()
}

func (s *EntElementVerifyRequest) SetEntName(v string) *EntElementVerifyRequest {
	s.EntName = &v
	return s
}

func (s *EntElementVerifyRequest) SetInfoVerifyType(v string) *EntElementVerifyRequest {
	s.InfoVerifyType = &v
	return s
}

func (s *EntElementVerifyRequest) SetLegalPersonCertNo(v string) *EntElementVerifyRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *EntElementVerifyRequest) SetLegalPersonName(v string) *EntElementVerifyRequest {
	s.LegalPersonName = &v
	return s
}

func (s *EntElementVerifyRequest) SetLicenseNo(v string) *EntElementVerifyRequest {
	s.LicenseNo = &v
	return s
}

func (s *EntElementVerifyRequest) SetMerchantBizId(v string) *EntElementVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EntElementVerifyRequest) SetMerchantUserId(v string) *EntElementVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EntElementVerifyRequest) SetSceneCode(v string) *EntElementVerifyRequest {
	s.SceneCode = &v
	return s
}

func (s *EntElementVerifyRequest) SetUserAuthorization(v string) *EntElementVerifyRequest {
	s.UserAuthorization = &v
	return s
}

type EntElementVerifyResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7***B-A3DC0DE3C83E
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntElementVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponseBody) SetCode(v string) *EntElementVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetMessage(v string) *EntElementVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetRequestId(v string) *EntElementVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntElementVerifyResponseBody) SetResult(v *EntElementVerifyResponseBodyResult) *EntElementVerifyResponseBody {
	s.Result = v
	return s
}

type EntElementVerifyResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode      *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	OpenTime     *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	ReasonCode   *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntElementVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponseBodyResult) SetBizCode(v string) *EntElementVerifyResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetOpenTime(v string) *EntElementVerifyResponseBodyResult {
	s.OpenTime = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonCode(v string) *EntElementVerifyResponseBodyResult {
	s.ReasonCode = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyResponseBodyResult {
	s.ReasonDetail = &v
	return s
}

func (s *EntElementVerifyResponseBodyResult) SetStatus(v string) *EntElementVerifyResponseBodyResult {
	s.Status = &v
	return s
}

type EntElementVerifyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EntElementVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntElementVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyResponse) GoString() string {
	return s.String()
}

func (s *EntElementVerifyResponse) SetHeaders(v map[string]*string) *EntElementVerifyResponse {
	s.Headers = v
	return s
}

func (s *EntElementVerifyResponse) SetStatusCode(v int32) *EntElementVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EntElementVerifyResponse) SetBody(v *EntElementVerifyResponseBody) *EntElementVerifyResponse {
	s.Body = v
	return s
}

type EntElementVerifyV2Request struct {
	// This parameter is required.
	EntName *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ENT_2META
	InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
	// example:
	//
	// 1******************9
	LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	LegalPersonName   *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 91330106673959****
	LicenseNo *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mch7x9a2b4c8d3e5f6g1h2i3j4k5****
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000000006
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntElementVerifyV2Request) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyV2Request) GoString() string {
	return s.String()
}

func (s *EntElementVerifyV2Request) SetEntName(v string) *EntElementVerifyV2Request {
	s.EntName = &v
	return s
}

func (s *EntElementVerifyV2Request) SetInfoVerifyType(v string) *EntElementVerifyV2Request {
	s.InfoVerifyType = &v
	return s
}

func (s *EntElementVerifyV2Request) SetLegalPersonCertNo(v string) *EntElementVerifyV2Request {
	s.LegalPersonCertNo = &v
	return s
}

func (s *EntElementVerifyV2Request) SetLegalPersonName(v string) *EntElementVerifyV2Request {
	s.LegalPersonName = &v
	return s
}

func (s *EntElementVerifyV2Request) SetLicenseNo(v string) *EntElementVerifyV2Request {
	s.LicenseNo = &v
	return s
}

func (s *EntElementVerifyV2Request) SetMerchantBizId(v string) *EntElementVerifyV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *EntElementVerifyV2Request) SetMerchantUserId(v string) *EntElementVerifyV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *EntElementVerifyV2Request) SetSceneCode(v string) *EntElementVerifyV2Request {
	s.SceneCode = &v
	return s
}

func (s *EntElementVerifyV2Request) SetUserAuthorization(v string) *EntElementVerifyV2Request {
	s.UserAuthorization = &v
	return s
}

type EntElementVerifyV2ResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7***B-A3DC0DE3C83E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntElementVerifyV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntElementVerifyV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyV2ResponseBody) GoString() string {
	return s.String()
}

func (s *EntElementVerifyV2ResponseBody) SetCode(v string) *EntElementVerifyV2ResponseBody {
	s.Code = &v
	return s
}

func (s *EntElementVerifyV2ResponseBody) SetMessage(v string) *EntElementVerifyV2ResponseBody {
	s.Message = &v
	return s
}

func (s *EntElementVerifyV2ResponseBody) SetRequestId(v string) *EntElementVerifyV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntElementVerifyV2ResponseBody) SetResult(v *EntElementVerifyV2ResponseBodyResult) *EntElementVerifyV2ResponseBody {
	s.Result = v
	return s
}

type EntElementVerifyV2ResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 2018-09-25/9999-09-09
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// example:
	//
	// LegalPersonNameFlag,LegalPersonCertNoFlag
	ReasonDetail *string `json:"ReasonDetail,omitempty" xml:"ReasonDetail,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntElementVerifyV2ResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntElementVerifyV2ResponseBodyResult) SetBizCode(v string) *EntElementVerifyV2ResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetOpenTime(v string) *EntElementVerifyV2ResponseBodyResult {
	s.OpenTime = &v
	return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetReasonDetail(v string) *EntElementVerifyV2ResponseBodyResult {
	s.ReasonDetail = &v
	return s
}

func (s *EntElementVerifyV2ResponseBodyResult) SetStatus(v string) *EntElementVerifyV2ResponseBodyResult {
	s.Status = &v
	return s
}

type EntElementVerifyV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EntElementVerifyV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntElementVerifyV2Response) String() string {
	return tea.Prettify(s)
}

func (s EntElementVerifyV2Response) GoString() string {
	return s.String()
}

func (s *EntElementVerifyV2Response) SetHeaders(v map[string]*string) *EntElementVerifyV2Response {
	s.Headers = v
	return s
}

func (s *EntElementVerifyV2Response) SetStatusCode(v int32) *EntElementVerifyV2Response {
	s.StatusCode = &v
	return s
}

func (s *EntElementVerifyV2Response) SetBody(v *EntElementVerifyV2ResponseBody) *EntElementVerifyV2Response {
	s.Body = v
	return s
}

type EntRiskQueryRequest struct {
	// example:
	//
	// 32198****193000
	MerchantBizId  *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// 00
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// example:
	//
	// 91330106673959****
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	// example:
	//
	// 1000000086
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// 1
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntRiskQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s EntRiskQueryRequest) GoString() string {
	return s.String()
}

func (s *EntRiskQueryRequest) SetMerchantBizId(v string) *EntRiskQueryRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EntRiskQueryRequest) SetMerchantUserId(v string) *EntRiskQueryRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EntRiskQueryRequest) SetParamType(v string) *EntRiskQueryRequest {
	s.ParamType = &v
	return s
}

func (s *EntRiskQueryRequest) SetParamValue(v string) *EntRiskQueryRequest {
	s.ParamValue = &v
	return s
}

func (s *EntRiskQueryRequest) SetSceneCode(v string) *EntRiskQueryRequest {
	s.SceneCode = &v
	return s
}

func (s *EntRiskQueryRequest) SetUserAuthorization(v string) *EntRiskQueryRequest {
	s.UserAuthorization = &v
	return s
}

type EntRiskQueryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 8FC3D6AC-9FED-4311-8DA7-C4BF47D9F260
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntRiskQueryResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntRiskQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntRiskQueryResponseBody) GoString() string {
	return s.String()
}

func (s *EntRiskQueryResponseBody) SetCode(v string) *EntRiskQueryResponseBody {
	s.Code = &v
	return s
}

func (s *EntRiskQueryResponseBody) SetMessage(v string) *EntRiskQueryResponseBody {
	s.Message = &v
	return s
}

func (s *EntRiskQueryResponseBody) SetRequestId(v string) *EntRiskQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntRiskQueryResponseBody) SetResult(v *EntRiskQueryResponseBodyResult) *EntRiskQueryResponseBody {
	s.Result = v
	return s
}

type EntRiskQueryResponseBodyResult struct {
	// example:
	//
	// 1
	BizCode  *string                                   `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	RiskList []*EntRiskQueryResponseBodyResultRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EntRiskQueryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntRiskQueryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntRiskQueryResponseBodyResult) SetBizCode(v string) *EntRiskQueryResponseBodyResult {
	s.BizCode = &v
	return s
}

func (s *EntRiskQueryResponseBodyResult) SetRiskList(v []*EntRiskQueryResponseBodyResultRiskList) *EntRiskQueryResponseBodyResult {
	s.RiskList = v
	return s
}

func (s *EntRiskQueryResponseBodyResult) SetStatus(v string) *EntRiskQueryResponseBodyResult {
	s.Status = &v
	return s
}

type EntRiskQueryResponseBodyResultRiskList struct {
	// example:
	//
	// 92500112MA5UHU****
	CreditCode *string `json:"CreditCode,omitempty" xml:"CreditCode,omitempty"`
	EntName    *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	// example:
	//
	// 2023-02-03
	ListedDate   *string `json:"ListedDate,omitempty" xml:"ListedDate,omitempty"`
	ListedReason *string `json:"ListedReason,omitempty" xml:"ListedReason,omitempty"`
	OperationOrg *string `json:"OperationOrg,omitempty" xml:"OperationOrg,omitempty"`
	// example:
	//
	// 50011260996****
	RegNo *string `json:"RegNo,omitempty" xml:"RegNo,omitempty"`
	// example:
	//
	// 2023-02-06
	RemovedDate   *string `json:"RemovedDate,omitempty" xml:"RemovedDate,omitempty"`
	RemovedOrg    *string `json:"RemovedOrg,omitempty" xml:"RemovedOrg,omitempty"`
	RemovedReason *string `json:"RemovedReason,omitempty" xml:"RemovedReason,omitempty"`
}

func (s EntRiskQueryResponseBodyResultRiskList) String() string {
	return tea.Prettify(s)
}

func (s EntRiskQueryResponseBodyResultRiskList) GoString() string {
	return s.String()
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetCreditCode(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.CreditCode = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetEntName(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.EntName = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetListedDate(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.ListedDate = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetListedReason(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.ListedReason = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetOperationOrg(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.OperationOrg = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRegNo(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.RegNo = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedDate(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.RemovedDate = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedOrg(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.RemovedOrg = &v
	return s
}

func (s *EntRiskQueryResponseBodyResultRiskList) SetRemovedReason(v string) *EntRiskQueryResponseBodyResultRiskList {
	s.RemovedReason = &v
	return s
}

type EntRiskQueryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EntRiskQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntRiskQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s EntRiskQueryResponse) GoString() string {
	return s.String()
}

func (s *EntRiskQueryResponse) SetHeaders(v map[string]*string) *EntRiskQueryResponse {
	s.Headers = v
	return s
}

func (s *EntRiskQueryResponse) SetStatusCode(v int32) *EntRiskQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *EntRiskQueryResponse) SetBody(v *EntRiskQueryResponseBody) *EntRiskQueryResponse {
	s.Body = v
	return s
}

type EntVerifyRequest struct {
	// example:
	//
	// 321324***38293
	AccountNo      *string `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	EntName        *string `json:"EntName,omitempty" xml:"EntName,omitempty"`
	InfoVerifyType *string `json:"InfoVerifyType,omitempty" xml:"InfoVerifyType,omitempty"`
	// example:
	//
	// 370105*****3892
	LegalPersonCertNo *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	// example:
	//
	// 1300***53
	LegalPersonMobile *string `json:"LegalPersonMobile,omitempty" xml:"LegalPersonMobile,omitempty"`
	LegalPersonName   *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// example:
	//
	// 32132***328932
	LicenseNo      *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	MerchantBizId  *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// BASIC
	RiskModelVersion *string `json:"RiskModelVersion,omitempty" xml:"RiskModelVersion,omitempty"`
	RiskVerifyType   *string `json:"RiskVerifyType,omitempty" xml:"RiskVerifyType,omitempty"`
	// example:
	//
	// withdraw
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// 1
	UserAuthorization *string `json:"UserAuthorization,omitempty" xml:"UserAuthorization,omitempty"`
}

func (s EntVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyRequest) GoString() string {
	return s.String()
}

func (s *EntVerifyRequest) SetAccountNo(v string) *EntVerifyRequest {
	s.AccountNo = &v
	return s
}

func (s *EntVerifyRequest) SetEntName(v string) *EntVerifyRequest {
	s.EntName = &v
	return s
}

func (s *EntVerifyRequest) SetInfoVerifyType(v string) *EntVerifyRequest {
	s.InfoVerifyType = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonCertNo(v string) *EntVerifyRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonMobile(v string) *EntVerifyRequest {
	s.LegalPersonMobile = &v
	return s
}

func (s *EntVerifyRequest) SetLegalPersonName(v string) *EntVerifyRequest {
	s.LegalPersonName = &v
	return s
}

func (s *EntVerifyRequest) SetLicenseNo(v string) *EntVerifyRequest {
	s.LicenseNo = &v
	return s
}

func (s *EntVerifyRequest) SetMerchantBizId(v string) *EntVerifyRequest {
	s.MerchantBizId = &v
	return s
}

func (s *EntVerifyRequest) SetMerchantUserId(v string) *EntVerifyRequest {
	s.MerchantUserId = &v
	return s
}

func (s *EntVerifyRequest) SetRiskModelVersion(v string) *EntVerifyRequest {
	s.RiskModelVersion = &v
	return s
}

func (s *EntVerifyRequest) SetRiskVerifyType(v string) *EntVerifyRequest {
	s.RiskVerifyType = &v
	return s
}

func (s *EntVerifyRequest) SetSceneCode(v string) *EntVerifyRequest {
	s.SceneCode = &v
	return s
}

func (s *EntVerifyRequest) SetUserAuthorization(v string) *EntVerifyRequest {
	s.UserAuthorization = &v
	return s
}

type EntVerifyResponseBody struct {
	// example:
	//
	// Success
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7-A***B-A3DC0DE3C83E
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *EntVerifyResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBody) SetCode(v string) *EntVerifyResponseBody {
	s.Code = &v
	return s
}

func (s *EntVerifyResponseBody) SetMessage(v string) *EntVerifyResponseBody {
	s.Message = &v
	return s
}

func (s *EntVerifyResponseBody) SetRequestId(v string) *EntVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *EntVerifyResponseBody) SetResult(v *EntVerifyResponseBodyResult) *EntVerifyResponseBody {
	s.Result = v
	return s
}

type EntVerifyResponseBodyResult struct {
	RiskVerifyResult *EntVerifyResponseBodyResultRiskVerifyResult `json:"RiskVerifyResult,omitempty" xml:"RiskVerifyResult,omitempty" type:"Struct"`
}

func (s EntVerifyResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResult) SetRiskVerifyResult(v *EntVerifyResponseBodyResultRiskVerifyResult) *EntVerifyResponseBodyResult {
	s.RiskVerifyResult = v
	return s
}

type EntVerifyResponseBodyResultRiskVerifyResult struct {
	// example:
	//
	// true
	Found        *bool                                                      `json:"Found,omitempty" xml:"Found,omitempty"`
	ModelResults []*EntVerifyResponseBodyResultRiskVerifyResultModelResults `json:"ModelResults,omitempty" xml:"ModelResults,omitempty" type:"Repeated"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResult) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetFound(v bool) *EntVerifyResponseBodyResultRiskVerifyResult {
	s.Found = &v
	return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResult) SetModelResults(v []*EntVerifyResponseBodyResultRiskVerifyResultModelResults) *EntVerifyResponseBodyResultRiskVerifyResult {
	s.ModelResults = v
	return s
}

type EntVerifyResponseBodyResultRiskVerifyResultModelResults struct {
	// example:
	//
	// model_1
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponseBodyResultRiskVerifyResultModelResults) GoString() string {
	return s.String()
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetModelName(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
	s.ModelName = &v
	return s
}

func (s *EntVerifyResponseBodyResultRiskVerifyResultModelResults) SetResult(v string) *EntVerifyResponseBodyResultRiskVerifyResultModelResults {
	s.Result = &v
	return s
}

type EntVerifyResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EntVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EntVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s EntVerifyResponse) GoString() string {
	return s.String()
}

func (s *EntVerifyResponse) SetHeaders(v map[string]*string) *EntVerifyResponse {
	s.Headers = v
	return s
}

func (s *EntVerifyResponse) SetStatusCode(v int32) *EntVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *EntVerifyResponse) SetBody(v *EntVerifyResponseBody) *EntVerifyResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 企业要素核验
//
// @param request - EntElementVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntElementVerifyResponse
func (client *Client) EntElementVerifyWithOptions(request *EntElementVerifyRequest, runtime *util.RuntimeOptions) (_result *EntElementVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		query["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.InfoVerifyType)) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntElementVerify"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EntElementVerifyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EntElementVerifyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 企业要素核验
//
// @param request - EntElementVerifyRequest
//
// @return EntElementVerifyResponse
func (client *Client) EntElementVerify(request *EntElementVerifyRequest) (_result *EntElementVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntElementVerifyResponse{}
	_body, _err := client.EntElementVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 企业要素验证V2
//
// @param request - EntElementVerifyV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntElementVerifyV2Response
func (client *Client) EntElementVerifyV2WithOptions(request *EntElementVerifyV2Request, runtime *util.RuntimeOptions) (_result *EntElementVerifyV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		query["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.InfoVerifyType)) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntElementVerifyV2"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EntElementVerifyV2Response{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EntElementVerifyV2Response{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 企业要素验证V2
//
// @param request - EntElementVerifyV2Request
//
// @return EntElementVerifyV2Response
func (client *Client) EntElementVerifyV2(request *EntElementVerifyV2Request) (_result *EntElementVerifyV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntElementVerifyV2Response{}
	_body, _err := client.EntElementVerifyV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 企业经营异常查询
//
// @param request - EntRiskQueryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntRiskQueryResponse
func (client *Client) EntRiskQueryWithOptions(request *EntRiskQueryRequest, runtime *util.RuntimeOptions) (_result *EntRiskQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ParamType)) {
		query["ParamType"] = request.ParamType
	}

	if !tea.BoolValue(util.IsUnset(request.ParamValue)) {
		query["ParamValue"] = request.ParamValue
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntRiskQuery"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EntRiskQueryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EntRiskQueryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 企业经营异常查询
//
// @param request - EntRiskQueryRequest
//
// @return EntRiskQueryResponse
func (client *Client) EntRiskQuery(request *EntRiskQueryRequest) (_result *EntRiskQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntRiskQueryResponse{}
	_body, _err := client.EntRiskQueryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 企业认证
//
// @param request - EntVerifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EntVerifyResponse
func (client *Client) EntVerifyWithOptions(request *EntVerifyRequest, runtime *util.RuntimeOptions) (_result *EntVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountNo)) {
		query["AccountNo"] = request.AccountNo
	}

	if !tea.BoolValue(util.IsUnset(request.EntName)) {
		query["EntName"] = request.EntName
	}

	if !tea.BoolValue(util.IsUnset(request.InfoVerifyType)) {
		query["InfoVerifyType"] = request.InfoVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonCertNo)) {
		query["LegalPersonCertNo"] = request.LegalPersonCertNo
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonMobile)) {
		query["LegalPersonMobile"] = request.LegalPersonMobile
	}

	if !tea.BoolValue(util.IsUnset(request.LegalPersonName)) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.RiskModelVersion)) {
		query["RiskModelVersion"] = request.RiskModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RiskVerifyType)) {
		query["RiskVerifyType"] = request.RiskVerifyType
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserAuthorization)) {
		query["UserAuthorization"] = request.UserAuthorization
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EntVerify"),
		Version:     tea.String("2022-11-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EntVerifyResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EntVerifyResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 企业认证
//
// @param request - EntVerifyRequest
//
// @return EntVerifyResponse
func (client *Client) EntVerify(request *EntVerifyRequest) (_result *EntVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EntVerifyResponse{}
	_body, _err := client.EntVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
