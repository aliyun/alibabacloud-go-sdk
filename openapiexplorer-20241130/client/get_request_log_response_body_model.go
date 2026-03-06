// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogInfo(v *GetRequestLogResponseBodyLogInfo) *GetRequestLogResponseBody
	GetLogInfo() *GetRequestLogResponseBodyLogInfo
	SetRequestId(v string) *GetRequestLogResponseBody
	GetRequestId() *string
}

type GetRequestLogResponseBody struct {
	// The details of the API request log.
	LogInfo *GetRequestLogResponseBodyLogInfo `json:"logInfo,omitempty" xml:"logInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRequestLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBody) GetLogInfo() *GetRequestLogResponseBodyLogInfo {
	return s.LogInfo
}

func (s *GetRequestLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRequestLogResponseBody) SetLogInfo(v *GetRequestLogResponseBodyLogInfo) *GetRequestLogResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetRequestLogResponseBody) SetRequestId(v string) *GetRequestLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRequestLogResponseBody) Validate() error {
	if s.LogInfo != nil {
		if err := s.LogInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRequestLogResponseBodyLogInfo struct {
	// The authentication information.
	AuthenticationInfo *GetRequestLogResponseBodyLogInfoAuthenticationInfo `json:"authenticationInfo,omitempty" xml:"authenticationInfo,omitempty" type:"Struct"`
	// The basic information about the API request.
	BasicInfo *GetRequestLogResponseBodyLogInfoBasicInfo `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" type:"Struct"`
	// The information about the caller.
	CallerInfo *GetRequestLogResponseBodyLogInfoCallerInfo `json:"callerInfo,omitempty" xml:"callerInfo,omitempty" type:"Struct"`
	// The request parameter information.
	Parameters []*GetRequestLogResponseBodyLogInfoParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// The response information that corresponds to the request.
	Responses *GetRequestLogResponseBodyLogInfoResponses `json:"responses,omitempty" xml:"responses,omitempty" type:"Struct"`
}

func (s GetRequestLogResponseBodyLogInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfo) GetAuthenticationInfo() *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	return s.AuthenticationInfo
}

func (s *GetRequestLogResponseBodyLogInfo) GetBasicInfo() *GetRequestLogResponseBodyLogInfoBasicInfo {
	return s.BasicInfo
}

func (s *GetRequestLogResponseBodyLogInfo) GetCallerInfo() *GetRequestLogResponseBodyLogInfoCallerInfo {
	return s.CallerInfo
}

func (s *GetRequestLogResponseBodyLogInfo) GetParameters() []*GetRequestLogResponseBodyLogInfoParameters {
	return s.Parameters
}

func (s *GetRequestLogResponseBodyLogInfo) GetResponses() *GetRequestLogResponseBodyLogInfoResponses {
	return s.Responses
}

func (s *GetRequestLogResponseBodyLogInfo) SetAuthenticationInfo(v *GetRequestLogResponseBodyLogInfoAuthenticationInfo) *GetRequestLogResponseBodyLogInfo {
	s.AuthenticationInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetBasicInfo(v *GetRequestLogResponseBodyLogInfoBasicInfo) *GetRequestLogResponseBodyLogInfo {
	s.BasicInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetCallerInfo(v *GetRequestLogResponseBodyLogInfoCallerInfo) *GetRequestLogResponseBodyLogInfo {
	s.CallerInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetParameters(v []*GetRequestLogResponseBodyLogInfoParameters) *GetRequestLogResponseBodyLogInfo {
	s.Parameters = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetResponses(v *GetRequestLogResponseBodyLogInfoResponses) *GetRequestLogResponseBodyLogInfo {
	s.Responses = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) Validate() error {
	if s.AuthenticationInfo != nil {
		if err := s.AuthenticationInfo.Validate(); err != nil {
			return err
		}
	}
	if s.BasicInfo != nil {
		if err := s.BasicInfo.Validate(); err != nil {
			return err
		}
	}
	if s.CallerInfo != nil {
		if err := s.CallerInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Responses != nil {
		if err := s.Responses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRequestLogResponseBodyLogInfoAuthenticationInfo struct {
	// The authentication type.
	//
	// - AK: An AccessKey pair, which can be a permanent AccessKey pair, a temporary AccessKey pair, or an STS token.
	//
	// - PRIVATEKEY: An AccessKey pair that uses asymmetric key encryption.
	//
	// - BEARERTOKEN: An identity verification mechanism that is widely used in the OAuth 2.0 framework and cloud services.
	//
	// - CUSTOM_SPI: An efficient and secure authentication method that is used for the delivery and management of Software as a Service (SaaS) products in Alibaba Cloud Marketplace.
	//
	// - Anonymous: Anonymous access.
	//
	// - DPS: Similar to an AccessKey pair, but uses a signature algorithm that is different from the official Alibaba Cloud algorithm. This method is specific to certain products.
	//
	// example:
	//
	// AK
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
	// The signature method.
	//
	// - HMAC-SHA1: The request is signed using the HMAC-SHA1 algorithm.
	//
	// - HMAC-SHA256: The request is signed using the HMAC-SHA256 algorithm.
	//
	// example:
	//
	// HMAC-SHA256
	SignatureMethod *string `json:"signatureMethod,omitempty" xml:"signatureMethod,omitempty"`
	// The signature version.
	//
	// example:
	//
	// unknown
	SignatureVersion *string `json:"signatureVersion,omitempty" xml:"signatureVersion,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoAuthenticationInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoAuthenticationInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) GetSignatureMethod() *string {
	return s.SignatureMethod
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) GetSignatureVersion() *string {
	return s.SignatureVersion
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetAuthenticationType(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.AuthenticationType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureMethod(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureMethod = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureVersion(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureVersion = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoBasicInfo struct {
	// The details of the access denied error. This field is returned only if an authentication error occurs for the specified request ID.
	AccessDeniedDetail *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty" type:"Struct"`
	// The name of the API that was called.
	//
	// example:
	//
	// RunInstances
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// The API reference information.
	ApiDoc *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc `json:"apiDoc,omitempty" xml:"apiDoc,omitempty" type:"Struct"`
	// The API style. Valid values: ROA and RPC.
	//
	// example:
	//
	// roa
	ApiStyle *string `json:"apiStyle,omitempty" xml:"apiStyle,omitempty"`
	// The API version.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The endpoint of the area where the endpoint is deployed.
	//
	// example:
	//
	// ecs.cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The error code in the log. This field is empty if no error is reported for the request.
	//
	// example:
	//
	// IncorrectStatus.TransitRouter
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message in the log. This field is empty if no error is reported for the request.
	//
	// example:
	//
	// The resource is not in a valid state for the operation.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The time when the gateway received the request. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-01-21T07:43:06Z
	GatewayProcessTime *string `json:"gatewayProcessTime,omitempty" xml:"gatewayProcessTime,omitempty"`
	// The HTTP request method. Examples: GET, PUT, and POST.
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"httpMethod,omitempty" xml:"httpMethod,omitempty"`
	// The HTTP status code in the log.
	//
	// example:
	//
	// 404
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID that you specified.
	//
	// example:
	//
	// 123E4567-E89B-12D3-A456-426614174000
	LogRequestId *string `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
	// The product code.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The product name, which includes the Chinese and English names.
	ProductName *GetRequestLogResponseBodyLogInfoBasicInfoProductName `json:"productName,omitempty" xml:"productName,omitempty" type:"Struct"`
	// The ID of the area where the endpoint is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time that elapses from when the gateway receives a request to when it returns a response. Unit: milliseconds (ms).
	//
	// example:
	//
	// 188
	RequestDuration *string `json:"requestDuration,omitempty" xml:"requestDuration,omitempty"`
	// The time when the request was initiated. The time is in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-01-21T07:43:06Z
	SdkRequestTime *string `json:"sdkRequestTime,omitempty" xml:"sdkRequestTime,omitempty"`
	// The result of the throttling check. FC.PASS: The request was not blocked by the throttling check and no rate limiting occurred. FC.DENY: The request was blocked by the throttling check and rate limiting occurred.
	//
	// example:
	//
	// FC.PASS
	ThrottlingResult *string `json:"throttlingResult,omitempty" xml:"throttlingResult,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetAccessDeniedDetail() *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetApi() *string {
	return s.Api
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetApiDoc() *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	return s.ApiDoc
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetApiStyle() *string {
	return s.ApiStyle
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetGatewayProcessTime() *string {
	return s.GatewayProcessTime
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetProduct() *string {
	return s.Product
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetProductName() *GetRequestLogResponseBodyLogInfoBasicInfoProductName {
	return s.ProductName
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetRequestDuration() *string {
	return s.RequestDuration
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetSdkRequestTime() *string {
	return s.SdkRequestTime
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) GetThrottlingResult() *string {
	return s.ThrottlingResult
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetAccessDeniedDetail(v *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApi(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Api = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiDoc(v *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiDoc = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiStyle(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiStyle = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiVersion(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiVersion = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetEndpoint(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Endpoint = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetErrorCode(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorCode = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetErrorMessage(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetGatewayProcessTime(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.GatewayProcessTime = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetHttpMethod(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpMethod = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetHttpStatusCode(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetLogRequestId(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.LogRequestId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetProduct(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Product = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetProductName(v *GetRequestLogResponseBodyLogInfoBasicInfoProductName) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ProductName = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetRegionId(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.RegionId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetRequestDuration(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.RequestDuration = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetSdkRequestTime(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.SdkRequestTime = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetThrottlingResult(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ThrottlingResult = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ApiDoc != nil {
		if err := s.ApiDoc.Validate(); err != nil {
			return err
		}
	}
	if s.ProductName != nil {
		if err := s.ProductName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail struct {
	// The specific operation that is denied.
	//
	// example:
	//
	// openapiexplorer:GetRequestLog
	AuthAction *string `json:"authAction,omitempty" xml:"authAction,omitempty"`
	// The identity.
	//
	// example:
	//
	// 205618123456123456
	AuthPrincipalDisplayName *string `json:"authPrincipalDisplayName,omitempty" xml:"authPrincipalDisplayName,omitempty"`
	// The UID of the Alibaba Cloud account to which the current identity belongs.
	//
	// example:
	//
	// 1001234561234567
	AuthPrincipalOwnerId *string `json:"authPrincipalOwnerId,omitempty" xml:"authPrincipalOwnerId,omitempty"`
	// The identity type of the operator.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"authPrincipalType,omitempty" xml:"authPrincipalType,omitempty"`
	// The encoded diagnostic message. To obtain more diagnostic information, call the DecodeDiagnosticMessage operation of RAM.
	//
	// example:
	//
	// -
	EncodedDiagnosticMessage *string `json:"encodedDiagnosticMessage,omitempty" xml:"encodedDiagnosticMessage,omitempty"`
	// The reason why the access is denied.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"noPermissionType,omitempty" xml:"noPermissionType,omitempty"`
	// The type of the policy that denies the access.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthAction(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetNoPermissionType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetPolicyType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoBasicInfoApiDoc struct {
	// The link to the API reference on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// https://api.alibabacloud.com/document/Ecs/2014-05-26/RunInstances
	AlibabacloudSite *string `json:"alibabacloudSite,omitempty" xml:"alibabacloudSite,omitempty"`
	// The link to the API reference on the Alibaba Cloud China Website (www\\.aliyun.com).
	//
	// example:
	//
	// https://api.aliyun.com/document/Ecs/2014-05-26/RunInstances
	AliyunSite *string `json:"aliyunSite,omitempty" xml:"aliyunSite,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) GetAlibabacloudSite() *string {
	return s.AlibabacloudSite
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) GetAliyunSite() *string {
	return s.AliyunSite
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAlibabacloudSite(v string) *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AlibabacloudSite = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAliyunSite(v string) *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AliyunSite = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoBasicInfoProductName struct {
	// The Chinese name of the product.
	//
	// example:
	//
	// 云服务器 ECS
	CnName *string `json:"cnName,omitempty" xml:"cnName,omitempty"`
	// The English name of the product.
	//
	// example:
	//
	// Elastic Compute Service
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoProductName) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoProductName) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) GetCnName() *string {
	return s.CnName
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) GetEnName() *string {
	return s.EnName
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) SetCnName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.CnName = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) SetEnName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.EnName = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoCallerInfo struct {
	// The ID of the caller\\"s account.
	//
	// example:
	//
	// 241009849925897811
	CallerAccountId *string `json:"callerAccountId,omitempty" xml:"callerAccountId,omitempty"`
	// The IP address of the caller.
	//
	// example:
	//
	// 100.68.xxx.xxx
	CallerIp *string `json:"callerIp,omitempty" xml:"callerIp,omitempty"`
	// The type of the caller.
	//
	// 1. customer: Alibaba Cloud account
	//
	// 2. sub: RAM user
	//
	// 3. AssumedRoleUser: A temporary identity that is created using a Security Token Service (STS) token.
	//
	// example:
	//
	// sub
	CallerType *string `json:"callerType,omitempty" xml:"callerType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1973374733454118
	MasterAccountId *string `json:"masterAccountId,omitempty" xml:"masterAccountId,omitempty"`
	// The user agent.
	//
	// example:
	//
	// AlibabaCloud API Workbench
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoCallerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoCallerInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) GetCallerAccountId() *string {
	return s.CallerAccountId
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) GetCallerIp() *string {
	return s.CallerIp
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) GetCallerType() *string {
	return s.CallerType
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) GetUserAgent() *string {
	return s.UserAgent
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerAccountId(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerAccountId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerIp(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerIp = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerType(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetMasterAccountId(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.MasterAccountId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetUserAgent(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.UserAgent = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoParameters struct {
	// The name of the request parameter.
	//
	// example:
	//
	// InstanceType
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the request parameter is required.
	//
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// The type of the request parameter.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the request parameter.
	//
	// example:
	//
	// ecs.g6.large
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoParameters) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoParameters) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoParameters) GetName() *string {
	return s.Name
}

func (s *GetRequestLogResponseBodyLogInfoParameters) GetRequired() *bool {
	return s.Required
}

func (s *GetRequestLogResponseBodyLogInfoParameters) GetType() *string {
	return s.Type
}

func (s *GetRequestLogResponseBodyLogInfoParameters) GetValue() interface{} {
	return s.Value
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetName(v string) *GetRequestLogResponseBodyLogInfoParameters {
	s.Name = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetRequired(v bool) *GetRequestLogResponseBodyLogInfoParameters {
	s.Required = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetType(v string) *GetRequestLogResponseBodyLogInfoParameters {
	s.Type = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetValue(v interface{}) *GetRequestLogResponseBodyLogInfoParameters {
	s.Value = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) Validate() error {
	return dara.Validate(s)
}

type GetRequestLogResponseBodyLogInfoResponses struct {
	// The response information.
	//
	// example:
	//
	// -
	ResponseBody *string `json:"responseBody,omitempty" xml:"responseBody,omitempty"`
	// The type of the response body. Valid values: JSON, XML, and HTML.
	//
	// example:
	//
	// JSON
	ResponseBodyFormat *string `json:"responseBodyFormat,omitempty" xml:"responseBodyFormat,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoResponses) String() string {
	return dara.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoResponses) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoResponses) GetResponseBody() *string {
	return s.ResponseBody
}

func (s *GetRequestLogResponseBodyLogInfoResponses) GetResponseBodyFormat() *string {
	return s.ResponseBodyFormat
}

func (s *GetRequestLogResponseBodyLogInfoResponses) SetResponseBody(v string) *GetRequestLogResponseBodyLogInfoResponses {
	s.ResponseBody = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoResponses) SetResponseBodyFormat(v string) *GetRequestLogResponseBodyLogInfoResponses {
	s.ResponseBodyFormat = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoResponses) Validate() error {
	return dara.Validate(s)
}
