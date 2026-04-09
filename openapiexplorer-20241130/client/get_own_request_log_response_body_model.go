// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOwnRequestLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogInfo(v *GetOwnRequestLogResponseBodyLogInfo) *GetOwnRequestLogResponseBody
	GetLogInfo() *GetOwnRequestLogResponseBodyLogInfo
	SetRequestId(v string) *GetOwnRequestLogResponseBody
	GetRequestId() *string
}

type GetOwnRequestLogResponseBody struct {
	// The detailed information about the log of the API call.
	LogInfo *GetOwnRequestLogResponseBodyLogInfo `json:"logInfo,omitempty" xml:"logInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetOwnRequestLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBody) GetLogInfo() *GetOwnRequestLogResponseBodyLogInfo {
	return s.LogInfo
}

func (s *GetOwnRequestLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOwnRequestLogResponseBody) SetLogInfo(v *GetOwnRequestLogResponseBodyLogInfo) *GetOwnRequestLogResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBody) SetRequestId(v string) *GetOwnRequestLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOwnRequestLogResponseBody) Validate() error {
	if s.LogInfo != nil {
		if err := s.LogInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOwnRequestLogResponseBodyLogInfo struct {
	// The authentication information.
	AuthenticationInfo *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo `json:"authenticationInfo,omitempty" xml:"authenticationInfo,omitempty" type:"Struct"`
	// The basic information about the log of the API call.
	BasicInfo *GetOwnRequestLogResponseBodyLogInfoBasicInfo `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" type:"Struct"`
	// The information about the caller.
	CallerInfo *GetOwnRequestLogResponseBodyLogInfoCallerInfo `json:"callerInfo,omitempty" xml:"callerInfo,omitempty" type:"Struct"`
	// The information about the request parameters.
	Parameters []*GetOwnRequestLogResponseBodyLogInfoParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// The information that is returned for the request.
	Responses *GetOwnRequestLogResponseBodyLogInfoResponses `json:"responses,omitempty" xml:"responses,omitempty" type:"Struct"`
}

func (s GetOwnRequestLogResponseBodyLogInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfo) GetAuthenticationInfo() *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	return s.AuthenticationInfo
}

func (s *GetOwnRequestLogResponseBodyLogInfo) GetBasicInfo() *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	return s.BasicInfo
}

func (s *GetOwnRequestLogResponseBodyLogInfo) GetCallerInfo() *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	return s.CallerInfo
}

func (s *GetOwnRequestLogResponseBodyLogInfo) GetParameters() []*GetOwnRequestLogResponseBodyLogInfoParameters {
	return s.Parameters
}

func (s *GetOwnRequestLogResponseBodyLogInfo) GetResponses() *GetOwnRequestLogResponseBodyLogInfoResponses {
	return s.Responses
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetAuthenticationInfo(v *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.AuthenticationInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetBasicInfo(v *GetOwnRequestLogResponseBodyLogInfoBasicInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.BasicInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetCallerInfo(v *GetOwnRequestLogResponseBodyLogInfoCallerInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.CallerInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetParameters(v []*GetOwnRequestLogResponseBodyLogInfoParameters) *GetOwnRequestLogResponseBodyLogInfo {
	s.Parameters = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetResponses(v *GetOwnRequestLogResponseBodyLogInfoResponses) *GetOwnRequestLogResponseBodyLogInfo {
	s.Responses = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) Validate() error {
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

type GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo struct {
	// The authentication type. Valid values:
	//
	// 	- AK: includes a permanent AccessKey pair, a temporary AccessKey pair, and a STS token.
	//
	// 	- PRIVATEKEY: an AccessKey pair for an asymmetric cryptography algorithm.
	//
	// 	- BEARETOKEN: an authentication mechanism that is widely used in the OAuth 2.0 framework and cloud services.
	//
	// 	- CUSTOM_SPI: an efficient and secure authentication method that is suitable for the delivery and management of Software as a Service (SaaS) services in Alibaba Cloud Marketplace.
	//
	// 	- Anonymous: anonymous access.
	//
	// 	- DPS: an authentication method that is similar to AK. Its signature algorithm is different from that of Alibaba Cloud services and is exclusive to specific products.
	//
	// example:
	//
	// AK
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
	// The signature algorithm. Valid values:
	//
	// 	- HMAC-SHA1
	//
	// 	- HMAC-SHA256
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

func (s GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) GetSignatureMethod() *string {
	return s.SignatureMethod
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) GetSignatureVersion() *string {
	return s.SignatureVersion
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetAuthenticationType(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.AuthenticationType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureMethod(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureMethod = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureVersion(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureVersion = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfo struct {
	// The error message returned if the operator does not have the required permissions.
	AccessDeniedDetail *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty" type:"Struct"`
	// The name of the API.
	//
	// example:
	//
	// RunInstances
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// The information about the API documentation.
	ApiDoc *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc `json:"apiDoc,omitempty" xml:"apiDoc,omitempty" type:"Struct"`
	// The API style. Valid values: roa and rpc.
	//
	// example:
	//
	// rpc
	ApiStyle *string `json:"apiStyle,omitempty" xml:"apiStyle,omitempty"`
	// The version of the API.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The endpoint of the service region.
	//
	// example:
	//
	// ecs.cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The error code in the log. This parameter is left empty if no error is reported in the API call.
	//
	// example:
	//
	// IncorrectStatus.TransitRouter
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message in the log. This parameter is left empty if no error is reported in the API call.
	//
	// example:
	//
	// The resource is not in a valid state for the operation.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The time when the gateway receives the request. Indicate the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-01-21T07:43:06Z
	GatewayProcessTime *string `json:"gatewayProcessTime,omitempty" xml:"gatewayProcessTime,omitempty"`
	// The HTTP request method.
	//
	// example:
	//
	// GET
	HttpMethod *string `json:"httpMethod,omitempty" xml:"httpMethod,omitempty"`
	// The HTTP status code in the log.
	//
	// example:
	//
	// 400
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
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
	// The product name, which includes the Chinese name and English name.
	ProductName *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName `json:"productName,omitempty" xml:"productName,omitempty" type:"Struct"`
	// The service region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The duration from when the gateway receives the request to when the client receives a response. Unit: milliseconds.
	//
	// example:
	//
	// 188
	RequestDuration *string `json:"requestDuration,omitempty" xml:"requestDuration,omitempty"`
	// The time when the request is initiated. Indicate the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-01-21T07:43:06Z
	SdkRequestTime *string `json:"sdkRequestTime,omitempty" xml:"sdkRequestTime,omitempty"`
	// The throttling result. Valid values: FC.PASS: The task is not blocked by throttling. FC.DENY: The task is blocked by throttling.
	//
	// example:
	//
	// FC.PASS
	ThrottlingResult *string `json:"throttlingResult,omitempty" xml:"throttlingResult,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetAccessDeniedDetail() *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetApi() *string {
	return s.Api
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetApiDoc() *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	return s.ApiDoc
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetApiStyle() *string {
	return s.ApiStyle
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetGatewayProcessTime() *string {
	return s.GatewayProcessTime
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetProduct() *string {
	return s.Product
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetProductName() *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName {
	return s.ProductName
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetRequestDuration() *string {
	return s.RequestDuration
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetSdkRequestTime() *string {
	return s.SdkRequestTime
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) GetThrottlingResult() *string {
	return s.ThrottlingResult
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetAccessDeniedDetail(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApi(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Api = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiDoc(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiDoc = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiStyle(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiStyle = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiVersion(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiVersion = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetEndpoint(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Endpoint = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetErrorCode(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorCode = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetErrorMessage(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetGatewayProcessTime(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.GatewayProcessTime = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetHttpMethod(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpMethod = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetHttpStatusCode(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetLogRequestId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.LogRequestId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetProduct(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Product = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetProductName(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ProductName = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetRegionId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.RegionId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetRequestDuration(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.RequestDuration = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetSdkRequestTime(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.SdkRequestTime = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetThrottlingResult(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ThrottlingResult = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) Validate() error {
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

type GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail struct {
	// The operation that the operator does not have permissions to perform.
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
	// The ID of the Alibaba Cloud account to which the current identity belongs.
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
	// The information after encoding, which can be used for troubleshooting. You can call the DecodeDiagnosticMessage operation of Resource Access Management (RAM) for further diagnostics.
	//
	// example:
	//
	// -
	EncodedDiagnosticMessage *string `json:"encodedDiagnosticMessage,omitempty" xml:"encodedDiagnosticMessage,omitempty"`
	// The cause of the permission-related error.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"noPermissionType,omitempty" xml:"noPermissionType,omitempty"`
	// The type of the policy that causes the permission-related error.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthAction(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetNoPermissionType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetPolicyType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc struct {
	// The documentation URL on the international site (alibabacloud.com).
	//
	// example:
	//
	// https://api.alibabacloud.com/document/Ecs/2014-05-26/RunInstances
	AlibabacloudSite *string `json:"alibabacloudSite,omitempty" xml:"alibabacloudSite,omitempty"`
	// The documentation URL on the China site (aliyun.com).
	//
	// example:
	//
	// https://api.aliyun.com/document/Ecs/2014-05-26/RunInstances
	AliyunSite *string `json:"aliyunSite,omitempty" xml:"aliyunSite,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) GetAlibabacloudSite() *string {
	return s.AlibabacloudSite
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) GetAliyunSite() *string {
	return s.AliyunSite
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAlibabacloudSite(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AlibabacloudSite = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAliyunSite(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AliyunSite = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName struct {
	// The product name in Chinese.
	CnName *string `json:"cnName,omitempty" xml:"cnName,omitempty"`
	// The product name in English.
	//
	// example:
	//
	// Elastic Compute Service
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) GetCnName() *string {
	return s.CnName
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) GetEnName() *string {
	return s.EnName
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) SetCnName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.CnName = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) SetEnName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.EnName = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoCallerInfo struct {
	// The account ID of the caller.
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
	// The type of the caller. Valid values:
	//
	// 1.  customer: an Alibaba Cloud account
	//
	// 2.  sub: a RAM user
	//
	// 3.  AssumedRoleUser: a user that uses a temporary Security Token Service (STS) token
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
	// The information about the user agent.
	//
	// example:
	//
	// AlibabaCloud API Workbench
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoCallerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoCallerInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) GetCallerAccountId() *string {
	return s.CallerAccountId
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) GetCallerIp() *string {
	return s.CallerIp
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) GetCallerType() *string {
	return s.CallerType
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) GetUserAgent() *string {
	return s.UserAgent
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerAccountId(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerAccountId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerIp(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerIp = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerType(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetMasterAccountId(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.MasterAccountId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetUserAgent(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.UserAgent = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoParameters struct {
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

func (s GetOwnRequestLogResponseBodyLogInfoParameters) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoParameters) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) GetName() *string {
	return s.Name
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) GetRequired() *bool {
	return s.Required
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) GetType() *string {
	return s.Type
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) GetValue() interface{} {
	return s.Value
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetName(v string) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Name = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetRequired(v bool) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Required = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetType(v string) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Type = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetValue(v interface{}) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Value = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) Validate() error {
	return dara.Validate(s)
}

type GetOwnRequestLogResponseBodyLogInfoResponses struct {
	// The response body.
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

func (s GetOwnRequestLogResponseBodyLogInfoResponses) String() string {
	return dara.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoResponses) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) GetResponseBody() *string {
	return s.ResponseBody
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) GetResponseBodyFormat() *string {
	return s.ResponseBodyFormat
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) SetResponseBody(v string) *GetOwnRequestLogResponseBodyLogInfoResponses {
	s.ResponseBody = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) SetResponseBodyFormat(v string) *GetOwnRequestLogResponseBodyLogInfoResponses {
	s.ResponseBodyFormat = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) Validate() error {
	return dara.Validate(s)
}
