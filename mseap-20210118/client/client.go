// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateLicenseRequest struct {
	// example:
	//
	// P20211118170645000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// esp.bookkeeping_course
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// B09YICKLVHNR7ZQR
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseNo   *string `json:"LicenseNo,omitempty" xml:"LicenseNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yunMarket
	LicensePublisher *string `json:"LicensePublisher,omitempty" xml:"LicensePublisher,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) SetBizId(v string) *ActivateLicenseRequest {
	s.BizId = &v
	return s
}

func (s *ActivateLicenseRequest) SetBizType(v string) *ActivateLicenseRequest {
	s.BizType = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseCode(v string) *ActivateLicenseRequest {
	s.LicenseCode = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicenseNo(v string) *ActivateLicenseRequest {
	s.LicenseNo = &v
	return s
}

func (s *ActivateLicenseRequest) SetLicensePublisher(v string) *ActivateLicenseRequest {
	s.LicensePublisher = &v
	return s
}

type ActivateLicenseResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 88EDA98E-6BE7-55DA-9EB6-B6444B882C59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) SetData(v bool) *ActivateLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

type ActivateLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetStatusCode(v int32) *ActivateLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
	s.Body = v
	return s
}

type CallbackTaskRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// MPC
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid     *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// orderId
	//
	// example:
	//
	// 1672369049358
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// outTaskId
	//
	// example:
	//
	// 1
	OutTaskId    *string `json:"OutTaskId,omitempty" xml:"OutTaskId,omitempty"`
	PrincipalKey *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// taskData
	//
	// example:
	//
	// {\\"result\\":\\"SUCCESS\\",\\"message\\":\\"null\\",\\"taskId\\":\\"8cbc97d8-9b2b-4c2f-862f-983ea5dbedc2\\"}
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
	// taskId
	//
	// example:
	//
	// 2559492
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// taskType
	//
	// example:
	//
	// PATENT_QUERY
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// 1
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s CallbackTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CallbackTaskRequest) GoString() string {
	return s.String()
}

func (s *CallbackTaskRequest) SetAliyunKp(v string) *CallbackTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *CallbackTaskRequest) SetApiType(v string) *CallbackTaskRequest {
	s.ApiType = &v
	return s
}

func (s *CallbackTaskRequest) SetBid(v string) *CallbackTaskRequest {
	s.Bid = &v
	return s
}

func (s *CallbackTaskRequest) SetBizCode(v string) *CallbackTaskRequest {
	s.BizCode = &v
	return s
}

func (s *CallbackTaskRequest) SetLang(v string) *CallbackTaskRequest {
	s.Lang = &v
	return s
}

func (s *CallbackTaskRequest) SetOrderId(v string) *CallbackTaskRequest {
	s.OrderId = &v
	return s
}

func (s *CallbackTaskRequest) SetOriginalRequest(v string) *CallbackTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *CallbackTaskRequest) SetOutTaskId(v string) *CallbackTaskRequest {
	s.OutTaskId = &v
	return s
}

func (s *CallbackTaskRequest) SetPrincipalKey(v string) *CallbackTaskRequest {
	s.PrincipalKey = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskData(v string) *CallbackTaskRequest {
	s.TaskData = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskId(v string) *CallbackTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CallbackTaskRequest) SetTaskType(v string) *CallbackTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CallbackTaskRequest) SetUserAccessKeyId(v string) *CallbackTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *CallbackTaskRequest) SetUserBid(v string) *CallbackTaskRequest {
	s.UserBid = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerParentId(v int64) *CallbackTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerSecurityTransport(v bool) *CallbackTaskRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *CallbackTaskRequest) SetUserCallerType(v string) *CallbackTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *CallbackTaskRequest) SetUserClientIp(v string) *CallbackTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *CallbackTaskRequest) SetUserKp(v string) *CallbackTaskRequest {
	s.UserKp = &v
	return s
}

func (s *CallbackTaskRequest) SetUserMfaPresent(v bool) *CallbackTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *CallbackTaskRequest) SetUserSecurityToken(v string) *CallbackTaskRequest {
	s.UserSecurityToken = &v
	return s
}

type CallbackTaskResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// bohai-web-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: vpc-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// True
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 56B009EB-AAA5-52C9-B05F-AF425E3885E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CallbackTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CallbackTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CallbackTaskResponseBody) SetAllowRetry(v bool) *CallbackTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *CallbackTaskResponseBody) SetAppName(v string) *CallbackTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *CallbackTaskResponseBody) SetDynamicCode(v string) *CallbackTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetDynamicMessage(v string) *CallbackTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CallbackTaskResponseBody) SetErrorCode(v string) *CallbackTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetErrorMsg(v string) *CallbackTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CallbackTaskResponseBody) SetHttpStatusCode(v int32) *CallbackTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CallbackTaskResponseBody) SetModule(v bool) *CallbackTaskResponseBody {
	s.Module = &v
	return s
}

func (s *CallbackTaskResponseBody) SetRequestId(v string) *CallbackTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CallbackTaskResponseBody) SetSuccess(v bool) *CallbackTaskResponseBody {
	s.Success = &v
	return s
}

type CallbackTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallbackTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CallbackTaskResponse) GoString() string {
	return s.String()
}

func (s *CallbackTaskResponse) SetHeaders(v map[string]*string) *CallbackTaskResponse {
	s.Headers = v
	return s
}

func (s *CallbackTaskResponse) SetStatusCode(v int32) *CallbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CallbackTaskResponse) SetBody(v *CallbackTaskResponseBody) *CallbackTaskResponse {
	s.Body = v
	return s
}

type DescribeAgreementStatusRequest struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s DescribeAgreementStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusRequest) SetAgreementCode(v string) *DescribeAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

type DescribeAgreementStatusResponseBody struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
	// example:
	//
	// 6254E13A-A79F-5786-9C75-7590727342C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeAgreementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponseBody) SetAgreementCode(v string) *DescribeAgreementStatusResponseBody {
	s.AgreementCode = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetRequestId(v string) *DescribeAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetStatus(v int32) *DescribeAgreementStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAgreementStatusResponseBody) SetUserId(v string) *DescribeAgreementStatusResponseBody {
	s.UserId = &v
	return s
}

type DescribeAgreementStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgreementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgreementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgreementStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgreementStatusResponse) SetHeaders(v map[string]*string) *DescribeAgreementStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgreementStatusResponse) SetStatusCode(v int32) *DescribeAgreementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgreementStatusResponse) SetBody(v *DescribeAgreementStatusResponseBody) *DescribeAgreementStatusResponse {
	s.Body = v
	return s
}

type GenerateUploadFilePolicyForPartnerRequest struct {
	// example:
	//
	// esp.website
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 2024/06/25/website_partner_third_party_leads_02.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// website_partner_third_party_leads
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetBizType(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetFileName(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetFileType(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.FileType = &v
	return s
}

type GenerateUploadFilePolicyForPartnerResponseBody struct {
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// live
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *GenerateUploadFilePolicyForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 6254E13A-A79F-5786-9C75-7590727342C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetAllowRetry(v bool) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetAppName(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetDynamicCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetDynamicMessage(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorArgs(v []interface{}) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorMsg(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetHttpStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetModule(v *GenerateUploadFilePolicyForPartnerResponseBodyModule) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetSuccess(v bool) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.Success = &v
	return s
}

type GenerateUploadFilePolicyForPartnerResponseBodyModule struct {
	// example:
	//
	// LTAI5tQPEXqDVu7794Bvw2xM
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// XXXXXXX
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	// example:
	//
	// 1719112842
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// website_partner_leads/website/xxxxxx/xxxxxx
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// example:
	//
	// //xx-xxx-partner.oss-cn-zhangjiakou.aliyuncs.com/
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// https://msea-website-partner.oss-cn-zhangjiakou.aliyuncs.com/website_xxxx_party_leads/website/xxxx/xxxx/2024/06/25/website_partner_third_party_leads_01?Expires=1719868413&OSSAccessKeyId=LTAI5tAnyDDDDD&Signature=XXXX
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// example:
	//
	// qQb34p8lIXcSFtog2y0H08bC0OI=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetAccessId(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetEncodedPolicy(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetExpireTime(v int64) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetFileDir(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetHost(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetOssUrl(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.OssUrl = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetSignature(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.Signature = &v
	return s
}

type GenerateUploadFilePolicyForPartnerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadFilePolicyForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetHeaders(v map[string]*string) *GenerateUploadFilePolicyForPartnerResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponse) SetBody(v *GenerateUploadFilePolicyForPartnerResponseBody) *GenerateUploadFilePolicyForPartnerResponse {
	s.Body = v
	return s
}

type GetNodeByFlowIdRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// MPC
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 180
	FlowId *int64 `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// lang
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetNodeByFlowIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByFlowIdRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdRequest) SetAliyunKp(v string) *GetNodeByFlowIdRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetApiType(v string) *GetNodeByFlowIdRequest {
	s.ApiType = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetBid(v string) *GetNodeByFlowIdRequest {
	s.Bid = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetFlowId(v int64) *GetNodeByFlowIdRequest {
	s.FlowId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetLang(v string) *GetNodeByFlowIdRequest {
	s.Lang = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetOriginalRequest(v string) *GetNodeByFlowIdRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserAccessKeyId(v string) *GetNodeByFlowIdRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserBid(v string) *GetNodeByFlowIdRequest {
	s.UserBid = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerParentId(v int64) *GetNodeByFlowIdRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerSecurityTransport(v bool) *GetNodeByFlowIdRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserCallerType(v string) *GetNodeByFlowIdRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserClientIp(v string) *GetNodeByFlowIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserKp(v string) *GetNodeByFlowIdRequest {
	s.UserKp = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserMfaPresent(v bool) *GetNodeByFlowIdRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetNodeByFlowIdRequest) SetUserSecurityToken(v string) *GetNodeByFlowIdRequest {
	s.UserSecurityToken = &v
	return s
}

type GetNodeByFlowIdResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// gatewayprood
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: vpc-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 200,131
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 8F30A673-46F0-5774-9D25-B45A29DB626E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeByFlowIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByFlowIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdResponseBody) SetAllowRetry(v bool) *GetNodeByFlowIdResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetAppName(v string) *GetNodeByFlowIdResponseBody {
	s.AppName = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetDynamicCode(v string) *GetNodeByFlowIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetDynamicMessage(v string) *GetNodeByFlowIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetErrorCode(v string) *GetNodeByFlowIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetErrorMsg(v string) *GetNodeByFlowIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetHttpStatusCode(v int32) *GetNodeByFlowIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetModule(v string) *GetNodeByFlowIdResponseBody {
	s.Module = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetRequestId(v string) *GetNodeByFlowIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeByFlowIdResponseBody) SetSuccess(v bool) *GetNodeByFlowIdResponseBody {
	s.Success = &v
	return s
}

type GetNodeByFlowIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeByFlowIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeByFlowIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByFlowIdResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByFlowIdResponse) SetHeaders(v map[string]*string) *GetNodeByFlowIdResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByFlowIdResponse) SetStatusCode(v int32) *GetNodeByFlowIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByFlowIdResponse) SetBody(v *GetNodeByFlowIdResponseBody) *GetNodeByFlowIdResponse {
	s.Body = v
	return s
}

type GetNodeByTemplateIdRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// example:
	//
	// 13
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// true
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetNodeByTemplateIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdRequest) SetAliyunKp(v string) *GetNodeByTemplateIdRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetApiType(v string) *GetNodeByTemplateIdRequest {
	s.ApiType = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetBid(v string) *GetNodeByTemplateIdRequest {
	s.Bid = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetLang(v string) *GetNodeByTemplateIdRequest {
	s.Lang = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetOriginalRequest(v string) *GetNodeByTemplateIdRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetTemplateId(v int64) *GetNodeByTemplateIdRequest {
	s.TemplateId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserAccessKeyId(v string) *GetNodeByTemplateIdRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserBid(v string) *GetNodeByTemplateIdRequest {
	s.UserBid = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerParentId(v int64) *GetNodeByTemplateIdRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerSecurityTransport(v bool) *GetNodeByTemplateIdRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserCallerType(v string) *GetNodeByTemplateIdRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserClientIp(v string) *GetNodeByTemplateIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserKp(v string) *GetNodeByTemplateIdRequest {
	s.UserKp = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserMfaPresent(v bool) *GetNodeByTemplateIdRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetNodeByTemplateIdRequest) SetUserSecurityToken(v string) *GetNodeByTemplateIdRequest {
	s.UserSecurityToken = &v
	return s
}

type GetNodeByTemplateIdResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// qdhxgcagent01
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// zxdfghjklasdfghjkl
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 220,116
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 53D045B1-466F-5165-B3BB-42E36F02BA86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNodeByTemplateIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdResponseBody) SetAllowRetry(v bool) *GetNodeByTemplateIdResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetAppName(v string) *GetNodeByTemplateIdResponseBody {
	s.AppName = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetDynamicCode(v string) *GetNodeByTemplateIdResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetDynamicMessage(v string) *GetNodeByTemplateIdResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetErrorCode(v string) *GetNodeByTemplateIdResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetErrorMsg(v string) *GetNodeByTemplateIdResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetHttpStatusCode(v int32) *GetNodeByTemplateIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetModule(v string) *GetNodeByTemplateIdResponseBody {
	s.Module = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetRequestId(v string) *GetNodeByTemplateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeByTemplateIdResponseBody) SetSuccess(v bool) *GetNodeByTemplateIdResponseBody {
	s.Success = &v
	return s
}

type GetNodeByTemplateIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeByTemplateIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByTemplateIdResponse) SetHeaders(v map[string]*string) *GetNodeByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByTemplateIdResponse) SetStatusCode(v int32) *GetNodeByTemplateIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByTemplateIdResponse) SetBody(v *GetNodeByTemplateIdResponseBody) *GetNodeByTemplateIdResponse {
	s.Body = v
	return s
}

type GetProxyByTypeRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// type
	//
	// example:
	//
	// 2
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// 1
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetProxyByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProxyByTypeRequest) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeRequest) SetAliyunKp(v string) *GetProxyByTypeRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetApiType(v string) *GetProxyByTypeRequest {
	s.ApiType = &v
	return s
}

func (s *GetProxyByTypeRequest) SetBid(v string) *GetProxyByTypeRequest {
	s.Bid = &v
	return s
}

func (s *GetProxyByTypeRequest) SetLang(v string) *GetProxyByTypeRequest {
	s.Lang = &v
	return s
}

func (s *GetProxyByTypeRequest) SetOriginalRequest(v string) *GetProxyByTypeRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetProxyByTypeRequest) SetType(v int32) *GetProxyByTypeRequest {
	s.Type = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserAccessKeyId(v string) *GetProxyByTypeRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserBid(v string) *GetProxyByTypeRequest {
	s.UserBid = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerParentId(v int64) *GetProxyByTypeRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerSecurityTransport(v bool) *GetProxyByTypeRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserCallerType(v string) *GetProxyByTypeRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserClientIp(v string) *GetProxyByTypeRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserKp(v string) *GetProxyByTypeRequest {
	s.UserKp = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserMfaPresent(v bool) *GetProxyByTypeRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetProxyByTypeRequest) SetUserSecurityToken(v string) *GetProxyByTypeRequest {
	s.UserSecurityToken = &v
	return s
}

type GetProxyByTypeResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// voldemort-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 118.113.245.10:3128
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// CEC1731F-0DA9-5E7D-AE53-7E4D76201C48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProxyByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProxyByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeResponseBody) SetAllowRetry(v bool) *GetProxyByTypeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetAppName(v string) *GetProxyByTypeResponseBody {
	s.AppName = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetDynamicCode(v string) *GetProxyByTypeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetDynamicMessage(v string) *GetProxyByTypeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetErrorCode(v string) *GetProxyByTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetErrorMsg(v string) *GetProxyByTypeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetHttpStatusCode(v int32) *GetProxyByTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetModule(v string) *GetProxyByTypeResponseBody {
	s.Module = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetRequestId(v string) *GetProxyByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProxyByTypeResponseBody) SetSuccess(v bool) *GetProxyByTypeResponseBody {
	s.Success = &v
	return s
}

type GetProxyByTypeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProxyByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProxyByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProxyByTypeResponse) GoString() string {
	return s.String()
}

func (s *GetProxyByTypeResponse) SetHeaders(v map[string]*string) *GetProxyByTypeResponse {
	s.Headers = v
	return s
}

func (s *GetProxyByTypeResponse) SetStatusCode(v int32) *GetProxyByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProxyByTypeResponse) SetBody(v *GetProxyByTypeResponseBody) *GetProxyByTypeResponse {
	s.Body = v
	return s
}

type GetRedisValueRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// MPC
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// key
	//
	// example:
	//
	// 106.14.34.208
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// timeout
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
	// example:
	//
	// {   \\"cust_id\\":\\"1111111\\",   \\"cust_name\\":\\"aa\\" }
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRedisValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRedisValueRequest) GoString() string {
	return s.String()
}

func (s *GetRedisValueRequest) SetAliyunKp(v string) *GetRedisValueRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetRedisValueRequest) SetApiType(v string) *GetRedisValueRequest {
	s.ApiType = &v
	return s
}

func (s *GetRedisValueRequest) SetBid(v string) *GetRedisValueRequest {
	s.Bid = &v
	return s
}

func (s *GetRedisValueRequest) SetKey(v string) *GetRedisValueRequest {
	s.Key = &v
	return s
}

func (s *GetRedisValueRequest) SetLang(v string) *GetRedisValueRequest {
	s.Lang = &v
	return s
}

func (s *GetRedisValueRequest) SetOriginalRequest(v string) *GetRedisValueRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetRedisValueRequest) SetTimeout(v int64) *GetRedisValueRequest {
	s.Timeout = &v
	return s
}

func (s *GetRedisValueRequest) SetUserAccessKeyId(v string) *GetRedisValueRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetRedisValueRequest) SetUserBid(v string) *GetRedisValueRequest {
	s.UserBid = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerParentId(v int64) *GetRedisValueRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerSecurityTransport(v bool) *GetRedisValueRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetRedisValueRequest) SetUserCallerType(v string) *GetRedisValueRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetRedisValueRequest) SetUserClientIp(v string) *GetRedisValueRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetRedisValueRequest) SetUserKp(v string) *GetRedisValueRequest {
	s.UserKp = &v
	return s
}

func (s *GetRedisValueRequest) SetUserMfaPresent(v bool) *GetRedisValueRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetRedisValueRequest) SetUserSecurityToken(v string) *GetRedisValueRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *GetRedisValueRequest) SetValue(v string) *GetRedisValueRequest {
	s.Value = &v
	return s
}

type GetRedisValueResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// bohai-web-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: eleme-zb
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// zxdfghjklasdfghjkl
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 107,72
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 0320C9F4-5EDC-5355-9D7E-DF4CF6C2B3BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRedisValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRedisValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetRedisValueResponseBody) SetAllowRetry(v bool) *GetRedisValueResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetRedisValueResponseBody) SetAppName(v string) *GetRedisValueResponseBody {
	s.AppName = &v
	return s
}

func (s *GetRedisValueResponseBody) SetDynamicCode(v string) *GetRedisValueResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetDynamicMessage(v string) *GetRedisValueResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetRedisValueResponseBody) SetErrorCode(v string) *GetRedisValueResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetErrorMsg(v string) *GetRedisValueResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRedisValueResponseBody) SetHttpStatusCode(v int32) *GetRedisValueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRedisValueResponseBody) SetModule(v string) *GetRedisValueResponseBody {
	s.Module = &v
	return s
}

func (s *GetRedisValueResponseBody) SetRequestId(v string) *GetRedisValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRedisValueResponseBody) SetSuccess(v bool) *GetRedisValueResponseBody {
	s.Success = &v
	return s
}

type GetRedisValueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRedisValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRedisValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRedisValueResponse) GoString() string {
	return s.String()
}

func (s *GetRedisValueResponse) SetHeaders(v map[string]*string) *GetRedisValueResponse {
	s.Headers = v
	return s
}

func (s *GetRedisValueResponse) SetStatusCode(v int32) *GetRedisValueResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRedisValueResponse) SetBody(v *GetRedisValueResponseBody) *GetRedisValueResponse {
	s.Body = v
	return s
}

type GetVariableRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// example:
	//
	// 17
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s GetVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVariableRequest) GoString() string {
	return s.String()
}

func (s *GetVariableRequest) SetAliyunKp(v string) *GetVariableRequest {
	s.AliyunKp = &v
	return s
}

func (s *GetVariableRequest) SetApiType(v string) *GetVariableRequest {
	s.ApiType = &v
	return s
}

func (s *GetVariableRequest) SetBid(v string) *GetVariableRequest {
	s.Bid = &v
	return s
}

func (s *GetVariableRequest) SetLang(v string) *GetVariableRequest {
	s.Lang = &v
	return s
}

func (s *GetVariableRequest) SetOriginalRequest(v string) *GetVariableRequest {
	s.OriginalRequest = &v
	return s
}

func (s *GetVariableRequest) SetTemplateId(v int64) *GetVariableRequest {
	s.TemplateId = &v
	return s
}

func (s *GetVariableRequest) SetUserAccessKeyId(v string) *GetVariableRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetVariableRequest) SetUserBid(v string) *GetVariableRequest {
	s.UserBid = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerParentId(v int64) *GetVariableRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerSecurityTransport(v bool) *GetVariableRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *GetVariableRequest) SetUserCallerType(v string) *GetVariableRequest {
	s.UserCallerType = &v
	return s
}

func (s *GetVariableRequest) SetUserClientIp(v string) *GetVariableRequest {
	s.UserClientIp = &v
	return s
}

func (s *GetVariableRequest) SetUserKp(v string) *GetVariableRequest {
	s.UserKp = &v
	return s
}

func (s *GetVariableRequest) SetUserMfaPresent(v bool) *GetVariableRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *GetVariableRequest) SetUserSecurityToken(v string) *GetVariableRequest {
	s.UserSecurityToken = &v
	return s
}

type GetVariableResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// voldemort-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: eleme-zb-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 207,155
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// D28419C9-335E-50A7-BD7D-ACF250A825E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVariableResponseBody) GoString() string {
	return s.String()
}

func (s *GetVariableResponseBody) SetAllowRetry(v bool) *GetVariableResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetVariableResponseBody) SetAppName(v string) *GetVariableResponseBody {
	s.AppName = &v
	return s
}

func (s *GetVariableResponseBody) SetDynamicCode(v string) *GetVariableResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetVariableResponseBody) SetDynamicMessage(v string) *GetVariableResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetVariableResponseBody) SetErrorCode(v string) *GetVariableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetVariableResponseBody) SetErrorMsg(v string) *GetVariableResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetVariableResponseBody) SetHttpStatusCode(v int32) *GetVariableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVariableResponseBody) SetModule(v string) *GetVariableResponseBody {
	s.Module = &v
	return s
}

func (s *GetVariableResponseBody) SetRequestId(v string) *GetVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVariableResponseBody) SetSuccess(v bool) *GetVariableResponseBody {
	s.Success = &v
	return s
}

type GetVariableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVariableResponse) GoString() string {
	return s.String()
}

func (s *GetVariableResponse) SetHeaders(v map[string]*string) *GetVariableResponse {
	s.Headers = v
	return s
}

func (s *GetVariableResponse) SetStatusCode(v int32) *GetVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVariableResponse) SetBody(v *GetVariableResponseBody) *GetVariableResponse {
	s.Body = v
	return s
}

type IdentifyCodeRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// {\\"engine\\": \\"MySQL\\", \\"instanceId\\": \\"rm-2zes07949gc0febg6\\", \\"userId\\": \\"1204765431532768\\", \\"previousExistConfig\\": False, \\"engineVersion\\": \\"8.0\\", \\"autoResourceOptimize\\": 0, \\"dasProOn\\": False}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 1551278
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// example:
	//
	// CBWP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// true
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// true
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s IdentifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s IdentifyCodeRequest) GoString() string {
	return s.String()
}

func (s *IdentifyCodeRequest) SetAliyunKp(v string) *IdentifyCodeRequest {
	s.AliyunKp = &v
	return s
}

func (s *IdentifyCodeRequest) SetApiType(v string) *IdentifyCodeRequest {
	s.ApiType = &v
	return s
}

func (s *IdentifyCodeRequest) SetBid(v string) *IdentifyCodeRequest {
	s.Bid = &v
	return s
}

func (s *IdentifyCodeRequest) SetData(v string) *IdentifyCodeRequest {
	s.Data = &v
	return s
}

func (s *IdentifyCodeRequest) SetLabel(v string) *IdentifyCodeRequest {
	s.Label = &v
	return s
}

func (s *IdentifyCodeRequest) SetLang(v string) *IdentifyCodeRequest {
	s.Lang = &v
	return s
}

func (s *IdentifyCodeRequest) SetOriginalRequest(v string) *IdentifyCodeRequest {
	s.OriginalRequest = &v
	return s
}

func (s *IdentifyCodeRequest) SetType(v string) *IdentifyCodeRequest {
	s.Type = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserAccessKeyId(v string) *IdentifyCodeRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserBid(v string) *IdentifyCodeRequest {
	s.UserBid = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerParentId(v int64) *IdentifyCodeRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerSecurityTransport(v bool) *IdentifyCodeRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserCallerType(v string) *IdentifyCodeRequest {
	s.UserCallerType = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserClientIp(v string) *IdentifyCodeRequest {
	s.UserClientIp = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserKp(v string) *IdentifyCodeRequest {
	s.UserKp = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserMfaPresent(v bool) *IdentifyCodeRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *IdentifyCodeRequest) SetUserSecurityToken(v string) *IdentifyCodeRequest {
	s.UserSecurityToken = &v
	return s
}

type IdentifyCodeResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// baasamlservice
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: lazada-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// 1234567890
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 230,94
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 8F30A673-46F0-5774-9D25-B45A29DB626E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IdentifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IdentifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *IdentifyCodeResponseBody) SetAllowRetry(v bool) *IdentifyCodeResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetAppName(v string) *IdentifyCodeResponseBody {
	s.AppName = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetDynamicCode(v string) *IdentifyCodeResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetDynamicMessage(v string) *IdentifyCodeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetErrorCode(v string) *IdentifyCodeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetErrorMsg(v string) *IdentifyCodeResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetHttpStatusCode(v int32) *IdentifyCodeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetModule(v string) *IdentifyCodeResponseBody {
	s.Module = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetRequestId(v string) *IdentifyCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *IdentifyCodeResponseBody) SetSuccess(v bool) *IdentifyCodeResponseBody {
	s.Success = &v
	return s
}

type IdentifyCodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IdentifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IdentifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s IdentifyCodeResponse) GoString() string {
	return s.String()
}

func (s *IdentifyCodeResponse) SetHeaders(v map[string]*string) *IdentifyCodeResponse {
	s.Headers = v
	return s
}

func (s *IdentifyCodeResponse) SetStatusCode(v int32) *IdentifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *IdentifyCodeResponse) SetBody(v *IdentifyCodeResponseBody) *IdentifyCodeResponse {
	s.Body = v
	return s
}

type PullRpaModelRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// templateId
	//
	// example:
	//
	// 17
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s PullRpaModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PullRpaModelRequest) GoString() string {
	return s.String()
}

func (s *PullRpaModelRequest) SetAliyunKp(v string) *PullRpaModelRequest {
	s.AliyunKp = &v
	return s
}

func (s *PullRpaModelRequest) SetApiType(v string) *PullRpaModelRequest {
	s.ApiType = &v
	return s
}

func (s *PullRpaModelRequest) SetBid(v string) *PullRpaModelRequest {
	s.Bid = &v
	return s
}

func (s *PullRpaModelRequest) SetLang(v string) *PullRpaModelRequest {
	s.Lang = &v
	return s
}

func (s *PullRpaModelRequest) SetOriginalRequest(v string) *PullRpaModelRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PullRpaModelRequest) SetTemplateId(v int64) *PullRpaModelRequest {
	s.TemplateId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserAccessKeyId(v string) *PullRpaModelRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserBid(v string) *PullRpaModelRequest {
	s.UserBid = &v
	return s
}

func (s *PullRpaModelRequest) SetUserCallerParentId(v int64) *PullRpaModelRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PullRpaModelRequest) SetUserCallerType(v string) *PullRpaModelRequest {
	s.UserCallerType = &v
	return s
}

func (s *PullRpaModelRequest) SetUserClientIp(v string) *PullRpaModelRequest {
	s.UserClientIp = &v
	return s
}

func (s *PullRpaModelRequest) SetUserKp(v string) *PullRpaModelRequest {
	s.UserKp = &v
	return s
}

func (s *PullRpaModelRequest) SetUserMfaPresent(v bool) *PullRpaModelRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PullRpaModelRequest) SetUserSecurityToken(v string) *PullRpaModelRequest {
	s.UserSecurityToken = &v
	return s
}

type PullRpaModelResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// gatewayprood
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: vpc-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 100008
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// {}
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 06055768-6BC0-5FE7-BDFF-BD4D79537035
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PullRpaModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullRpaModelResponseBody) GoString() string {
	return s.String()
}

func (s *PullRpaModelResponseBody) SetAllowRetry(v bool) *PullRpaModelResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PullRpaModelResponseBody) SetAppName(v string) *PullRpaModelResponseBody {
	s.AppName = &v
	return s
}

func (s *PullRpaModelResponseBody) SetDynamicCode(v string) *PullRpaModelResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetDynamicMessage(v string) *PullRpaModelResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PullRpaModelResponseBody) SetErrorCode(v string) *PullRpaModelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetErrorMsg(v string) *PullRpaModelResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PullRpaModelResponseBody) SetHttpStatusCode(v int32) *PullRpaModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PullRpaModelResponseBody) SetModule(v string) *PullRpaModelResponseBody {
	s.Module = &v
	return s
}

func (s *PullRpaModelResponseBody) SetRequestId(v string) *PullRpaModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PullRpaModelResponseBody) SetSuccess(v bool) *PullRpaModelResponseBody {
	s.Success = &v
	return s
}

type PullRpaModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullRpaModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullRpaModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PullRpaModelResponse) GoString() string {
	return s.String()
}

func (s *PullRpaModelResponse) SetHeaders(v map[string]*string) *PullRpaModelResponse {
	s.Headers = v
	return s
}

func (s *PullRpaModelResponse) SetStatusCode(v int32) *PullRpaModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PullRpaModelResponse) SetBody(v *PullRpaModelResponseBody) *PullRpaModelResponse {
	s.Body = v
	return s
}

type PullTaskRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// openAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid     *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1672369049358
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	PrincipalKey    *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// taskType
	//
	// example:
	//
	// PATENT_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerSecurityTransport
	//
	// example:
	//
	// 1
	UserCallerSecurityTransport *bool `json:"UserCallerSecurityTransport,omitempty" xml:"UserCallerSecurityTransport,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s PullTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PullTaskRequest) GoString() string {
	return s.String()
}

func (s *PullTaskRequest) SetAliyunKp(v string) *PullTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *PullTaskRequest) SetApiType(v string) *PullTaskRequest {
	s.ApiType = &v
	return s
}

func (s *PullTaskRequest) SetBid(v string) *PullTaskRequest {
	s.Bid = &v
	return s
}

func (s *PullTaskRequest) SetBizCode(v string) *PullTaskRequest {
	s.BizCode = &v
	return s
}

func (s *PullTaskRequest) SetLang(v string) *PullTaskRequest {
	s.Lang = &v
	return s
}

func (s *PullTaskRequest) SetOrderId(v string) *PullTaskRequest {
	s.OrderId = &v
	return s
}

func (s *PullTaskRequest) SetOriginalRequest(v string) *PullTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PullTaskRequest) SetPrincipalKey(v string) *PullTaskRequest {
	s.PrincipalKey = &v
	return s
}

func (s *PullTaskRequest) SetTaskType(v string) *PullTaskRequest {
	s.TaskType = &v
	return s
}

func (s *PullTaskRequest) SetUserAccessKeyId(v string) *PullTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PullTaskRequest) SetUserBid(v string) *PullTaskRequest {
	s.UserBid = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerParentId(v int64) *PullTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerSecurityTransport(v bool) *PullTaskRequest {
	s.UserCallerSecurityTransport = &v
	return s
}

func (s *PullTaskRequest) SetUserCallerType(v string) *PullTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *PullTaskRequest) SetUserClientIp(v string) *PullTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *PullTaskRequest) SetUserKp(v string) *PullTaskRequest {
	s.UserKp = &v
	return s
}

func (s *PullTaskRequest) SetUserMfaPresent(v bool) *PullTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PullTaskRequest) SetUserSecurityToken(v string) *PullTaskRequest {
	s.UserSecurityToken = &v
	return s
}

type PullTaskResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// voldemort-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *PullTaskResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 9831C9A6-3423-52C2-B0E5-5AE01D92C097
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PullTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PullTaskResponseBody) SetAllowRetry(v bool) *PullTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PullTaskResponseBody) SetAppName(v string) *PullTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *PullTaskResponseBody) SetDynamicCode(v string) *PullTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PullTaskResponseBody) SetDynamicMessage(v string) *PullTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PullTaskResponseBody) SetErrorCode(v string) *PullTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PullTaskResponseBody) SetErrorMsg(v string) *PullTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PullTaskResponseBody) SetHttpStatusCode(v int32) *PullTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PullTaskResponseBody) SetModule(v *PullTaskResponseBodyModule) *PullTaskResponseBody {
	s.Module = v
	return s
}

func (s *PullTaskResponseBody) SetRequestId(v string) *PullTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PullTaskResponseBody) SetSuccess(v bool) *PullTaskResponseBody {
	s.Success = &v
	return s
}

type PullTaskResponseBodyModule struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 1649470201045
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	OutTaskId    *string `json:"OutTaskId,omitempty" xml:"OutTaskId,omitempty"`
	PrincipalKey *string `json:"PrincipalKey,omitempty" xml:"PrincipalKey,omitempty"`
	// example:
	//
	// {\\"result\\":\\"SUCCESS\\",\\"message\\":\\"null\\",\\"taskId\\":\\"d8800bab-88b6-4c60-9e4f-ed38dbbdd9b3\\"}
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
	// example:
	//
	// 704614
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// PATENT_QUERY
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s PullTaskResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s PullTaskResponseBodyModule) GoString() string {
	return s.String()
}

func (s *PullTaskResponseBodyModule) SetBizCode(v string) *PullTaskResponseBodyModule {
	s.BizCode = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetOrderId(v string) *PullTaskResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetOutTaskId(v string) *PullTaskResponseBodyModule {
	s.OutTaskId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetPrincipalKey(v string) *PullTaskResponseBodyModule {
	s.PrincipalKey = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskData(v string) *PullTaskResponseBodyModule {
	s.TaskData = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskId(v string) *PullTaskResponseBodyModule {
	s.TaskId = &v
	return s
}

func (s *PullTaskResponseBodyModule) SetTaskType(v string) *PullTaskResponseBodyModule {
	s.TaskType = &v
	return s
}

type PullTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PullTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PullTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PullTaskResponse) GoString() string {
	return s.String()
}

func (s *PullTaskResponse) SetHeaders(v map[string]*string) *PullTaskResponse {
	s.Headers = v
	return s
}

func (s *PullTaskResponse) SetStatusCode(v int32) *PullTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PullTaskResponse) SetBody(v *PullTaskResponseBody) *PullTaskResponse {
	s.Body = v
	return s
}

type PushRpaTaskRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// public
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// modelId
	//
	// example:
	//
	// 1951
	ModelId *int64 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// request
	//
	// example:
	//
	// {}
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// status
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// taskId
	//
	// example:
	//
	// 833812
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// templateId
	//
	// example:
	//
	// 26
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s PushRpaTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskRequest) GoString() string {
	return s.String()
}

func (s *PushRpaTaskRequest) SetAliyunKp(v string) *PushRpaTaskRequest {
	s.AliyunKp = &v
	return s
}

func (s *PushRpaTaskRequest) SetApiType(v string) *PushRpaTaskRequest {
	s.ApiType = &v
	return s
}

func (s *PushRpaTaskRequest) SetBid(v string) *PushRpaTaskRequest {
	s.Bid = &v
	return s
}

func (s *PushRpaTaskRequest) SetLang(v string) *PushRpaTaskRequest {
	s.Lang = &v
	return s
}

func (s *PushRpaTaskRequest) SetModelId(v int64) *PushRpaTaskRequest {
	s.ModelId = &v
	return s
}

func (s *PushRpaTaskRequest) SetName(v string) *PushRpaTaskRequest {
	s.Name = &v
	return s
}

func (s *PushRpaTaskRequest) SetOriginalRequest(v string) *PushRpaTaskRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PushRpaTaskRequest) SetRequest(v string) *PushRpaTaskRequest {
	s.Request = &v
	return s
}

func (s *PushRpaTaskRequest) SetStatus(v int32) *PushRpaTaskRequest {
	s.Status = &v
	return s
}

func (s *PushRpaTaskRequest) SetTaskId(v int64) *PushRpaTaskRequest {
	s.TaskId = &v
	return s
}

func (s *PushRpaTaskRequest) SetTemplateId(v int64) *PushRpaTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserAccessKeyId(v string) *PushRpaTaskRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserBid(v string) *PushRpaTaskRequest {
	s.UserBid = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserCallerParentId(v int64) *PushRpaTaskRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserCallerType(v string) *PushRpaTaskRequest {
	s.UserCallerType = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserClientIp(v string) *PushRpaTaskRequest {
	s.UserClientIp = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserKp(v string) *PushRpaTaskRequest {
	s.UserKp = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserMfaPresent(v bool) *PushRpaTaskRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *PushRpaTaskRequest) SetUserSecurityToken(v string) *PushRpaTaskRequest {
	s.UserSecurityToken = &v
	return s
}

type PushRpaTaskResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// itl-material
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: lazada-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// 11111111111111111111111
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 58.23.71.83:3128
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// ECE5E7EF-6898-5E24-97A1-B96C73BDE26C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushRpaTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PushRpaTaskResponseBody) SetAllowRetry(v bool) *PushRpaTaskResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetAppName(v string) *PushRpaTaskResponseBody {
	s.AppName = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetDynamicCode(v string) *PushRpaTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetDynamicMessage(v string) *PushRpaTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetErrorCode(v string) *PushRpaTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetErrorMsg(v string) *PushRpaTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetHttpStatusCode(v int32) *PushRpaTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetModule(v string) *PushRpaTaskResponseBody {
	s.Module = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetRequestId(v string) *PushRpaTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushRpaTaskResponseBody) SetSuccess(v bool) *PushRpaTaskResponseBody {
	s.Success = &v
	return s
}

type PushRpaTaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushRpaTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushRpaTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskResponse) GoString() string {
	return s.String()
}

func (s *PushRpaTaskResponse) SetHeaders(v map[string]*string) *PushRpaTaskResponse {
	s.Headers = v
	return s
}

func (s *PushRpaTaskResponse) SetStatusCode(v int32) *PushRpaTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *PushRpaTaskResponse) SetBody(v *PushRpaTaskResponseBody) *PushRpaTaskResponse {
	s.Body = v
	return s
}

type PushRpaTaskDetailRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// MPC
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// inputData
	//
	// example:
	//
	// http://wssq.sbj.cnipa.gov.cn:9080/tmsve/wssqsy_getCayzDl.xhtml
	InputData *string `json:"InputData,omitempty" xml:"InputData,omitempty"`
	// inputHtml
	//
	// example:
	//
	// 1
	InputHtml *string `json:"InputHtml,omitempty" xml:"InputHtml,omitempty"`
	// inputScreenshot
	//
	// example:
	//
	// 1
	InputScreenshot *string `json:"InputScreenshot,omitempty" xml:"InputScreenshot,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// modelDetailId
	//
	// example:
	//
	// 6
	ModelDetailId *int64 `json:"ModelDetailId,omitempty" xml:"ModelDetailId,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// outputData
	OutputData *string `json:"OutputData,omitempty" xml:"OutputData,omitempty"`
	// outputHtml
	//
	// example:
	//
	// <div class=\\"photobox\\" id=\\"Layer3\\" style=\\"visibility: visible
	OutputHtml *string `json:"OutputHtml,omitempty" xml:"OutputHtml,omitempty"`
	// outputScreenshot
	//
	// example:
	//
	// http://dbu-nap-p-test.oss-cn-beijing.aliyuncs.com/202301/20230110/5782089/1673334129101-d111874e-f181-4d1c-8edc-83e789975329.jpg?Expires=1675926129&OSSAccessKeyId=hObpgEXoca42qH3V&Signature=------
	OutputScreenshot *string `json:"OutputScreenshot,omitempty" xml:"OutputScreenshot,omitempty"`
	// status
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// taskDetailId
	//
	// example:
	//
	// 1
	TaskDetailId *int64 `json:"TaskDetailId,omitempty" xml:"TaskDetailId,omitempty"`
	// taskId
	//
	// example:
	//
	// 5596654
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
}

func (s PushRpaTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailRequest) SetAliyunKp(v string) *PushRpaTaskDetailRequest {
	s.AliyunKp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetApiType(v string) *PushRpaTaskDetailRequest {
	s.ApiType = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetBid(v string) *PushRpaTaskDetailRequest {
	s.Bid = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputData(v string) *PushRpaTaskDetailRequest {
	s.InputData = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputHtml(v string) *PushRpaTaskDetailRequest {
	s.InputHtml = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetInputScreenshot(v string) *PushRpaTaskDetailRequest {
	s.InputScreenshot = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetLang(v string) *PushRpaTaskDetailRequest {
	s.Lang = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetModelDetailId(v int64) *PushRpaTaskDetailRequest {
	s.ModelDetailId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOriginalRequest(v string) *PushRpaTaskDetailRequest {
	s.OriginalRequest = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputData(v string) *PushRpaTaskDetailRequest {
	s.OutputData = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputHtml(v string) *PushRpaTaskDetailRequest {
	s.OutputHtml = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetOutputScreenshot(v string) *PushRpaTaskDetailRequest {
	s.OutputScreenshot = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetStatus(v int32) *PushRpaTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetTaskDetailId(v int64) *PushRpaTaskDetailRequest {
	s.TaskDetailId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetTaskId(v int64) *PushRpaTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserAccessKeyId(v string) *PushRpaTaskDetailRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserBid(v string) *PushRpaTaskDetailRequest {
	s.UserBid = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserCallerParentId(v int64) *PushRpaTaskDetailRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserCallerType(v string) *PushRpaTaskDetailRequest {
	s.UserCallerType = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserClientIp(v string) *PushRpaTaskDetailRequest {
	s.UserClientIp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserKp(v string) *PushRpaTaskDetailRequest {
	s.UserKp = &v
	return s
}

func (s *PushRpaTaskDetailRequest) SetUserSecurityToken(v string) *PushRpaTaskDetailRequest {
	s.UserSecurityToken = &v
	return s
}

type PushRpaTaskDetailResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// voldemort-aliyun-com
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// can not find env: vpc-sg-pre
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	//
	// example:
	//
	// Success. Request Success.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 207,155
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 0320C9F4-5EDC-5355-9D7E-DF4CF6C2B3BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PushRpaTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailResponseBody) SetAllowRetry(v bool) *PushRpaTaskDetailResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetAppName(v string) *PushRpaTaskDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetDynamicCode(v string) *PushRpaTaskDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetDynamicMessage(v string) *PushRpaTaskDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetErrorCode(v string) *PushRpaTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetErrorMsg(v string) *PushRpaTaskDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetHttpStatusCode(v int32) *PushRpaTaskDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetModule(v string) *PushRpaTaskDetailResponseBody {
	s.Module = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetRequestId(v string) *PushRpaTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushRpaTaskDetailResponseBody) SetSuccess(v bool) *PushRpaTaskDetailResponseBody {
	s.Success = &v
	return s
}

type PushRpaTaskDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushRpaTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushRpaTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s PushRpaTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *PushRpaTaskDetailResponse) SetHeaders(v map[string]*string) *PushRpaTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *PushRpaTaskDetailResponse) SetStatusCode(v int32) *PushRpaTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *PushRpaTaskDetailResponse) SetBody(v *PushRpaTaskDetailResponseBody) *PushRpaTaskDetailResponse {
	s.Body = v
	return s
}

type SendNotificationForPartnerRequest struct {
	// example:
	//
	// DMP
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// MESSAGE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// dmp_website_xxx
	NotifycationEventType *string            `json:"NotifycationEventType,omitempty" xml:"NotifycationEventType,omitempty"`
	ParamMap              map[string]*string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// example:
	//
	// 12312212312
	SendTarget       *string   `json:"SendTarget,omitempty" xml:"SendTarget,omitempty"`
	SmartSubChannels []*string `json:"SmartSubChannels,omitempty" xml:"SmartSubChannels,omitempty" type:"Repeated"`
	// example:
	//
	// 5b29647n-a172-4ccd-ba33-73669896c287
	TrackId *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s SendNotificationForPartnerRequest) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationForPartnerRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerRequest) SetBizId(v string) *SendNotificationForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetChannelType(v string) *SendNotificationForPartnerRequest {
	s.ChannelType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetNotifyType(v string) *SendNotificationForPartnerRequest {
	s.NotifyType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetNotifycationEventType(v string) *SendNotificationForPartnerRequest {
	s.NotifycationEventType = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetParamMap(v map[string]*string) *SendNotificationForPartnerRequest {
	s.ParamMap = v
	return s
}

func (s *SendNotificationForPartnerRequest) SetSendTarget(v string) *SendNotificationForPartnerRequest {
	s.SendTarget = &v
	return s
}

func (s *SendNotificationForPartnerRequest) SetSmartSubChannels(v []*string) *SendNotificationForPartnerRequest {
	s.SmartSubChannels = v
	return s
}

func (s *SendNotificationForPartnerRequest) SetTrackId(v string) *SendNotificationForPartnerRequest {
	s.TrackId = &v
	return s
}

type SendNotificationForPartnerShrinkRequest struct {
	// example:
	//
	// DMP
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 1
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// example:
	//
	// MESSAGE
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// example:
	//
	// dmp_website_xxx
	NotifycationEventType *string `json:"NotifycationEventType,omitempty" xml:"NotifycationEventType,omitempty"`
	ParamMapShrink        *string `json:"ParamMap,omitempty" xml:"ParamMap,omitempty"`
	// example:
	//
	// 12312212312
	SendTarget             *string `json:"SendTarget,omitempty" xml:"SendTarget,omitempty"`
	SmartSubChannelsShrink *string `json:"SmartSubChannels,omitempty" xml:"SmartSubChannels,omitempty"`
	// example:
	//
	// 5b29647n-a172-4ccd-ba33-73669896c287
	TrackId *string `json:"TrackId,omitempty" xml:"TrackId,omitempty"`
}

func (s SendNotificationForPartnerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationForPartnerShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerShrinkRequest) SetBizId(v string) *SendNotificationForPartnerShrinkRequest {
	s.BizId = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetChannelType(v string) *SendNotificationForPartnerShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetNotifyType(v string) *SendNotificationForPartnerShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetNotifycationEventType(v string) *SendNotificationForPartnerShrinkRequest {
	s.NotifycationEventType = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetParamMapShrink(v string) *SendNotificationForPartnerShrinkRequest {
	s.ParamMapShrink = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetSendTarget(v string) *SendNotificationForPartnerShrinkRequest {
	s.SendTarget = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetSmartSubChannelsShrink(v string) *SendNotificationForPartnerShrinkRequest {
	s.SmartSubChannelsShrink = &v
	return s
}

func (s *SendNotificationForPartnerShrinkRequest) SetTrackId(v string) *SendNotificationForPartnerShrinkRequest {
	s.TrackId = &v
	return s
}

type SendNotificationForPartnerResponseBody struct {
	// example:
	//
	// 11111111111111111111111
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 0A011920166449C2FAAE8D179E1704C5
	MsgId *string `json:"MsgId,omitempty" xml:"MsgId,omitempty"`
	// example:
	//
	// 1940A84F-6D90-5764-9119-6279970C6FF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendNotificationForPartnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerResponseBody) SetErrorMsg(v string) *SendNotificationForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetMsgId(v string) *SendNotificationForPartnerResponseBody {
	s.MsgId = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetRequestId(v string) *SendNotificationForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendNotificationForPartnerResponseBody) SetSuccess(v bool) *SendNotificationForPartnerResponseBody {
	s.Success = &v
	return s
}

type SendNotificationForPartnerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendNotificationForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendNotificationForPartnerResponse) String() string {
	return tea.Prettify(s)
}

func (s SendNotificationForPartnerResponse) GoString() string {
	return s.String()
}

func (s *SendNotificationForPartnerResponse) SetHeaders(v map[string]*string) *SendNotificationForPartnerResponse {
	s.Headers = v
	return s
}

func (s *SendNotificationForPartnerResponse) SetStatusCode(v int32) *SendNotificationForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SendNotificationForPartnerResponse) SetBody(v *SendNotificationForPartnerResponseBody) *SendNotificationForPartnerResponse {
	s.Body = v
	return s
}

type SetRedisValueRequest struct {
	// aliyunKp
	//
	// example:
	//
	// 1
	AliyunKp *string `json:"AliyunKp,omitempty" xml:"AliyunKp,omitempty"`
	// apiType
	//
	// example:
	//
	// part_config_data
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// bid
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// example:
	//
	// 1684967696495902
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// lang
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// originalRequest
	//
	// example:
	//
	// 1
	OriginalRequest *string `json:"OriginalRequest,omitempty" xml:"OriginalRequest,omitempty"`
	// requestId
	//
	// example:
	//
	// F864A883-AD76-53D5-9A24-A6DAD5177697
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// timeout
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// userAccessKeyId
	//
	// example:
	//
	// 1
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// userBid
	//
	// example:
	//
	// 1
	UserBid *string `json:"UserBid,omitempty" xml:"UserBid,omitempty"`
	// userCallerParentId
	//
	// example:
	//
	// 1
	UserCallerParentId *int64 `json:"UserCallerParentId,omitempty" xml:"UserCallerParentId,omitempty"`
	// userCallerType
	//
	// example:
	//
	// 1
	UserCallerType *string `json:"UserCallerType,omitempty" xml:"UserCallerType,omitempty"`
	// userClientIp
	//
	// example:
	//
	// 1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// userKp
	//
	// example:
	//
	// 1
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// userMfaPresent
	//
	// example:
	//
	// 1
	UserMfaPresent *bool `json:"UserMfaPresent,omitempty" xml:"UserMfaPresent,omitempty"`
	// userSecurityToken
	//
	// example:
	//
	// 1
	UserSecurityToken *string `json:"UserSecurityToken,omitempty" xml:"UserSecurityToken,omitempty"`
	// example:
	//
	// 259200000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetRedisValueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetRedisValueRequest) GoString() string {
	return s.String()
}

func (s *SetRedisValueRequest) SetAliyunKp(v string) *SetRedisValueRequest {
	s.AliyunKp = &v
	return s
}

func (s *SetRedisValueRequest) SetApiType(v string) *SetRedisValueRequest {
	s.ApiType = &v
	return s
}

func (s *SetRedisValueRequest) SetBid(v string) *SetRedisValueRequest {
	s.Bid = &v
	return s
}

func (s *SetRedisValueRequest) SetKey(v string) *SetRedisValueRequest {
	s.Key = &v
	return s
}

func (s *SetRedisValueRequest) SetLang(v string) *SetRedisValueRequest {
	s.Lang = &v
	return s
}

func (s *SetRedisValueRequest) SetOriginalRequest(v string) *SetRedisValueRequest {
	s.OriginalRequest = &v
	return s
}

func (s *SetRedisValueRequest) SetRequestId(v string) *SetRedisValueRequest {
	s.RequestId = &v
	return s
}

func (s *SetRedisValueRequest) SetTimeout(v int64) *SetRedisValueRequest {
	s.Timeout = &v
	return s
}

func (s *SetRedisValueRequest) SetUserAccessKeyId(v string) *SetRedisValueRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *SetRedisValueRequest) SetUserBid(v string) *SetRedisValueRequest {
	s.UserBid = &v
	return s
}

func (s *SetRedisValueRequest) SetUserCallerParentId(v int64) *SetRedisValueRequest {
	s.UserCallerParentId = &v
	return s
}

func (s *SetRedisValueRequest) SetUserCallerType(v string) *SetRedisValueRequest {
	s.UserCallerType = &v
	return s
}

func (s *SetRedisValueRequest) SetUserClientIp(v string) *SetRedisValueRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetRedisValueRequest) SetUserKp(v string) *SetRedisValueRequest {
	s.UserKp = &v
	return s
}

func (s *SetRedisValueRequest) SetUserMfaPresent(v bool) *SetRedisValueRequest {
	s.UserMfaPresent = &v
	return s
}

func (s *SetRedisValueRequest) SetUserSecurityToken(v string) *SetRedisValueRequest {
	s.UserSecurityToken = &v
	return s
}

func (s *SetRedisValueRequest) SetValue(v string) *SetRedisValueRequest {
	s.Value = &v
	return s
}

type SetRedisValueResponseBody struct {
	// allowRetry
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// appName
	//
	// example:
	//
	// cloudquery
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamicCode
	//
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamicMessage
	//
	// example:
	//
	// 10.151.12.0/24,100.104.36.0/26,47.102.181.0/24,100.104.52.0/24,47.101.109.0/24,120.55.129.0/24,11.115.103.0/24,47.102.234.0/24
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// errorCode
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// errorMsg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// module
	//
	// example:
	//
	// 71,135
	Module *bool `json:"Module,omitempty" xml:"Module,omitempty"`
	// requestId
	//
	// example:
	//
	// 195BABE2-7105-5C16-ABCE-2D0997CCE2E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetRedisValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetRedisValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetRedisValueResponseBody) SetAllowRetry(v bool) *SetRedisValueResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *SetRedisValueResponseBody) SetAppName(v string) *SetRedisValueResponseBody {
	s.AppName = &v
	return s
}

func (s *SetRedisValueResponseBody) SetDynamicCode(v string) *SetRedisValueResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetDynamicMessage(v string) *SetRedisValueResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SetRedisValueResponseBody) SetErrorCode(v string) *SetRedisValueResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetErrorMsg(v string) *SetRedisValueResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SetRedisValueResponseBody) SetHttpStatusCode(v int32) *SetRedisValueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SetRedisValueResponseBody) SetModule(v bool) *SetRedisValueResponseBody {
	s.Module = &v
	return s
}

func (s *SetRedisValueResponseBody) SetRequestId(v string) *SetRedisValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetRedisValueResponseBody) SetSuccess(v bool) *SetRedisValueResponseBody {
	s.Success = &v
	return s
}

type SetRedisValueResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRedisValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRedisValueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetRedisValueResponse) GoString() string {
	return s.String()
}

func (s *SetRedisValueResponse) SetHeaders(v map[string]*string) *SetRedisValueResponse {
	s.Headers = v
	return s
}

func (s *SetRedisValueResponse) SetStatusCode(v int32) *SetRedisValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRedisValueResponse) SetBody(v *SetRedisValueResponseBody) *SetRedisValueResponse {
	s.Body = v
	return s
}

type UpdateAgreementStatusRequest struct {
	// example:
	//
	// 10aa40008e081ad7b1fb50bffc3a70b1
	AgreementCode *string `json:"AgreementCode,omitempty" xml:"AgreementCode,omitempty"`
}

func (s UpdateAgreementStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusRequest) SetAgreementCode(v string) *UpdateAgreementStatusRequest {
	s.AgreementCode = &v
	return s
}

type UpdateAgreementStatusResponseBody struct {
	// example:
	//
	// 6BDB1964-A6E9-5946-89A4-523D35645BB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgreementStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusResponseBody) SetRequestId(v string) *UpdateAgreementStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAgreementStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgreementStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgreementStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgreementStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgreementStatusResponse) SetHeaders(v map[string]*string) *UpdateAgreementStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgreementStatusResponse) SetStatusCode(v int32) *UpdateAgreementStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgreementStatusResponse) SetBody(v *UpdateAgreementStatusResponseBody) *UpdateAgreementStatusResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mseap"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *util.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseCode)) {
		query["LicenseCode"] = request.LicenseCode
	}

	if !tea.BoolValue(util.IsUnset(request.LicenseNo)) {
		query["LicenseNo"] = request.LicenseNo
	}

	if !tea.BoolValue(util.IsUnset(request.LicensePublisher)) {
		query["LicensePublisher"] = request.LicensePublisher
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateLicense"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ActivateLicenseRequest
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务回调
//
// @param request - CallbackTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallbackTaskResponse
func (client *Client) CallbackTaskWithOptions(request *CallbackTaskRequest, runtime *util.RuntimeOptions) (_result *CallbackTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.OutTaskId)) {
		query["OutTaskId"] = request.OutTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalKey)) {
		query["PrincipalKey"] = request.PrincipalKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskData)) {
		query["TaskData"] = request.TaskData
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CallbackTask"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CallbackTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务回调
//
// @param request - CallbackTaskRequest
//
// @return CallbackTaskResponse
func (client *Client) CallbackTask(request *CallbackTaskRequest) (_result *CallbackTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CallbackTaskResponse{}
	_body, _err := client.CallbackTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeAgreementStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAgreementStatusResponse
func (client *Client) DescribeAgreementStatusWithOptions(request *DescribeAgreementStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeAgreementStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementCode)) {
		query["AgreementCode"] = request.AgreementCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAgreementStatus"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAgreementStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAgreementStatusRequest
//
// @return DescribeAgreementStatusResponse
func (client *Client) DescribeAgreementStatus(request *DescribeAgreementStatusRequest) (_result *DescribeAgreementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAgreementStatusResponse{}
	_body, _err := client.DescribeAgreementStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴生成上传文件策略
//
// @param request - GenerateUploadFilePolicyForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadFilePolicyForPartnerResponse
func (client *Client) GenerateUploadFilePolicyForPartnerWithOptions(request *GenerateUploadFilePolicyForPartnerRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadFilePolicyForPartnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		query["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadFilePolicyForPartner"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadFilePolicyForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴生成上传文件策略
//
// @param request - GenerateUploadFilePolicyForPartnerRequest
//
// @return GenerateUploadFilePolicyForPartnerResponse
func (client *Client) GenerateUploadFilePolicyForPartner(request *GenerateUploadFilePolicyForPartnerRequest) (_result *GenerateUploadFilePolicyForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadFilePolicyForPartnerResponse{}
	_body, _err := client.GenerateUploadFilePolicyForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取node节点通过流程id
//
// @param request - GetNodeByFlowIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeByFlowIdResponse
func (client *Client) GetNodeByFlowIdWithOptions(request *GetNodeByFlowIdRequest, runtime *util.RuntimeOptions) (_result *GetNodeByFlowIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.FlowId)) {
		query["FlowId"] = request.FlowId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeByFlowId"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeByFlowIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取node节点通过流程id
//
// @param request - GetNodeByFlowIdRequest
//
// @return GetNodeByFlowIdResponse
func (client *Client) GetNodeByFlowId(request *GetNodeByFlowIdRequest) (_result *GetNodeByFlowIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeByFlowIdResponse{}
	_body, _err := client.GetNodeByFlowIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取node节点通过模版id
//
// @param request - GetNodeByTemplateIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeByTemplateIdResponse
func (client *Client) GetNodeByTemplateIdWithOptions(request *GetNodeByTemplateIdRequest, runtime *util.RuntimeOptions) (_result *GetNodeByTemplateIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNodeByTemplateId"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeByTemplateIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取node节点通过模版id
//
// @param request - GetNodeByTemplateIdRequest
//
// @return GetNodeByTemplateIdResponse
func (client *Client) GetNodeByTemplateId(request *GetNodeByTemplateIdRequest) (_result *GetNodeByTemplateIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeByTemplateIdResponse{}
	_body, _err := client.GetNodeByTemplateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代理
//
// @param request - GetProxyByTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProxyByTypeResponse
func (client *Client) GetProxyByTypeWithOptions(request *GetProxyByTypeRequest, runtime *util.RuntimeOptions) (_result *GetProxyByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProxyByType"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProxyByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代理
//
// @param request - GetProxyByTypeRequest
//
// @return GetProxyByTypeResponse
func (client *Client) GetProxyByType(request *GetProxyByTypeRequest) (_result *GetProxyByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProxyByTypeResponse{}
	_body, _err := client.GetProxyByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取redis值
//
// @param request - GetRedisValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRedisValueResponse
func (client *Client) GetRedisValueWithOptions(request *GetRedisValueRequest, runtime *util.RuntimeOptions) (_result *GetRedisValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRedisValue"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRedisValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取redis值
//
// @param request - GetRedisValueRequest
//
// @return GetRedisValueResponse
func (client *Client) GetRedisValue(request *GetRedisValueRequest) (_result *GetRedisValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRedisValueResponse{}
	_body, _err := client.GetRedisValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取变量
//
// @param request - GetVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVariableResponse
func (client *Client) GetVariableWithOptions(request *GetVariableRequest, runtime *util.RuntimeOptions) (_result *GetVariableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVariable"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取变量
//
// @param request - GetVariableRequest
//
// @return GetVariableResponse
func (client *Client) GetVariable(request *GetVariableRequest) (_result *GetVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVariableResponse{}
	_body, _err := client.GetVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 识别验证码
//
// @param request - IdentifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IdentifyCodeResponse
func (client *Client) IdentifyCodeWithOptions(request *IdentifyCodeRequest, runtime *util.RuntimeOptions) (_result *IdentifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Label)) {
		query["Label"] = request.Label
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IdentifyCode"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IdentifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 识别验证码
//
// @param request - IdentifyCodeRequest
//
// @return IdentifyCodeResponse
func (client *Client) IdentifyCode(request *IdentifyCodeRequest) (_result *IdentifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IdentifyCodeResponse{}
	_body, _err := client.IdentifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 拉取协议变更识别模型
//
// @param request - PullRpaModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullRpaModelResponse
func (client *Client) PullRpaModelWithOptions(request *PullRpaModelRequest, runtime *util.RuntimeOptions) (_result *PullRpaModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullRpaModel"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PullRpaModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 拉取协议变更识别模型
//
// @param request - PullRpaModelRequest
//
// @return PullRpaModelResponse
func (client *Client) PullRpaModel(request *PullRpaModelRequest) (_result *PullRpaModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PullRpaModelResponse{}
	_body, _err := client.PullRpaModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RPA拉取任务
//
// @param request - PullTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PullTaskResponse
func (client *Client) PullTaskWithOptions(request *PullTaskRequest, runtime *util.RuntimeOptions) (_result *PullTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.BizCode)) {
		query["BizCode"] = request.BizCode
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalKey)) {
		query["PrincipalKey"] = request.PrincipalKey
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerSecurityTransport)) {
		query["UserCallerSecurityTransport"] = request.UserCallerSecurityTransport
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PullTask"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PullTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RPA拉取任务
//
// @param request - PullTaskRequest
//
// @return PullTaskResponse
func (client *Client) PullTask(request *PullTaskRequest) (_result *PullTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PullTaskResponse{}
	_body, _err := client.PullTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送RPA运行时任务
//
// @param request - PushRpaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushRpaTaskResponse
func (client *Client) PushRpaTaskWithOptions(request *PushRpaTaskRequest, runtime *util.RuntimeOptions) (_result *PushRpaTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.Request)) {
		query["Request"] = request.Request
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushRpaTask"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushRpaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送RPA运行时任务
//
// @param request - PushRpaTaskRequest
//
// @return PushRpaTaskResponse
func (client *Client) PushRpaTask(request *PushRpaTaskRequest) (_result *PushRpaTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushRpaTaskResponse{}
	_body, _err := client.PushRpaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推送运行时任务明细
//
// @param request - PushRpaTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushRpaTaskDetailResponse
func (client *Client) PushRpaTaskDetailWithOptions(request *PushRpaTaskDetailRequest, runtime *util.RuntimeOptions) (_result *PushRpaTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.InputData)) {
		query["InputData"] = request.InputData
	}

	if !tea.BoolValue(util.IsUnset(request.InputHtml)) {
		query["InputHtml"] = request.InputHtml
	}

	if !tea.BoolValue(util.IsUnset(request.InputScreenshot)) {
		query["InputScreenshot"] = request.InputScreenshot
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ModelDetailId)) {
		query["ModelDetailId"] = request.ModelDetailId
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.OutputData)) {
		query["OutputData"] = request.OutputData
	}

	if !tea.BoolValue(util.IsUnset(request.OutputHtml)) {
		query["OutputHtml"] = request.OutputHtml
	}

	if !tea.BoolValue(util.IsUnset(request.OutputScreenshot)) {
		query["OutputScreenshot"] = request.OutputScreenshot
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskDetailId)) {
		query["TaskDetailId"] = request.TaskDetailId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PushRpaTaskDetail"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PushRpaTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推送运行时任务明细
//
// @param request - PushRpaTaskDetailRequest
//
// @return PushRpaTaskDetailResponse
func (client *Client) PushRpaTaskDetail(request *PushRpaTaskDetailRequest) (_result *PushRpaTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PushRpaTaskDetailResponse{}
	_body, _err := client.PushRpaTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴发送消息通知
//
// @param tmpReq - SendNotificationForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendNotificationForPartnerResponse
func (client *Client) SendNotificationForPartnerWithOptions(tmpReq *SendNotificationForPartnerRequest, runtime *util.RuntimeOptions) (_result *SendNotificationForPartnerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendNotificationForPartnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParamMap)) {
		request.ParamMapShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParamMap, tea.String("ParamMap"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SmartSubChannels)) {
		request.SmartSubChannelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SmartSubChannels, tea.String("SmartSubChannels"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		query["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.NotifycationEventType)) {
		query["NotifycationEventType"] = request.NotifycationEventType
	}

	if !tea.BoolValue(util.IsUnset(request.ParamMapShrink)) {
		query["ParamMap"] = request.ParamMapShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SendTarget)) {
		query["SendTarget"] = request.SendTarget
	}

	if !tea.BoolValue(util.IsUnset(request.SmartSubChannelsShrink)) {
		query["SmartSubChannels"] = request.SmartSubChannelsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TrackId)) {
		query["TrackId"] = request.TrackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendNotificationForPartner"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendNotificationForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴发送消息通知
//
// @param request - SendNotificationForPartnerRequest
//
// @return SendNotificationForPartnerResponse
func (client *Client) SendNotificationForPartner(request *SendNotificationForPartnerRequest) (_result *SendNotificationForPartnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendNotificationForPartnerResponse{}
	_body, _err := client.SendNotificationForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// redis设置
//
// @param request - SetRedisValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRedisValueResponse
func (client *Client) SetRedisValueWithOptions(request *SetRedisValueRequest, runtime *util.RuntimeOptions) (_result *SetRedisValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunKp)) {
		query["AliyunKp"] = request.AliyunKp
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		query["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.Bid)) {
		query["Bid"] = request.Bid
	}

	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalRequest)) {
		query["OriginalRequest"] = request.OriginalRequest
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserBid)) {
		query["UserBid"] = request.UserBid
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerParentId)) {
		query["UserCallerParentId"] = request.UserCallerParentId
	}

	if !tea.BoolValue(util.IsUnset(request.UserCallerType)) {
		query["UserCallerType"] = request.UserCallerType
	}

	if !tea.BoolValue(util.IsUnset(request.UserClientIp)) {
		query["UserClientIp"] = request.UserClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.UserKp)) {
		query["UserKp"] = request.UserKp
	}

	if !tea.BoolValue(util.IsUnset(request.UserMfaPresent)) {
		query["UserMfaPresent"] = request.UserMfaPresent
	}

	if !tea.BoolValue(util.IsUnset(request.UserSecurityToken)) {
		query["UserSecurityToken"] = request.UserSecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetRedisValue"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetRedisValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// redis设置
//
// @param request - SetRedisValueRequest
//
// @return SetRedisValueResponse
func (client *Client) SetRedisValue(request *SetRedisValueRequest) (_result *SetRedisValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetRedisValueResponse{}
	_body, _err := client.SetRedisValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAgreementStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgreementStatusResponse
func (client *Client) UpdateAgreementStatusWithOptions(request *UpdateAgreementStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateAgreementStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgreementCode)) {
		query["AgreementCode"] = request.AgreementCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAgreementStatus"),
		Version:     tea.String("2021-01-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAgreementStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAgreementStatusRequest
//
// @return UpdateAgreementStatusResponse
func (client *Client) UpdateAgreementStatus(request *UpdateAgreementStatusRequest) (_result *UpdateAgreementStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgreementStatusResponse{}
	_body, _err := client.UpdateAgreementStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
