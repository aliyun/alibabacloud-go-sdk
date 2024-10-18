// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckSmsVerifyCodeRequest struct {
	// The verification policy for uppercase and lowercase letters of the verification code. Valid values:
	//
	// 	- 1: The verification policy does not distinguish uppercase and lowercase letters.
	//
	// 	- 2: The verification policy distinguishes uppercase and lowercase letters.
	//
	// example:
	//
	// 1
	CaseAuthPolicy *int64 `json:"CaseAuthPolicy,omitempty" xml:"CaseAuthPolicy,omitempty"`
	// The country code of the phone number. Default value: 86.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 12123231
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 18653529399
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The verification service name. If this parameter is not specified, the default service is used. The name can be up to 20 characters in length.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s CheckSmsVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeRequest) SetCaseAuthPolicy(v int64) *CheckSmsVerifyCodeRequest {
	s.CaseAuthPolicy = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetCountryCode(v string) *CheckSmsVerifyCodeRequest {
	s.CountryCode = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetOutId(v string) *CheckSmsVerifyCodeRequest {
	s.OutId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetOwnerId(v int64) *CheckSmsVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetPhoneNumber(v string) *CheckSmsVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetResourceOwnerAccount(v string) *CheckSmsVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetResourceOwnerId(v int64) *CheckSmsVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetSchemeName(v string) *CheckSmsVerifyCodeRequest {
	s.SchemeName = &v
	return s
}

func (s *CheckSmsVerifyCodeRequest) SetVerifyCode(v string) *CheckSmsVerifyCodeRequest {
	s.VerifyCode = &v
	return s
}

type CheckSmsVerifyCodeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [Response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11174283.0.0.70c5616bkj38Wa).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *CheckSmsVerifyCodeResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSmsVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *CheckSmsVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetCode(v string) *CheckSmsVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetMessage(v string) *CheckSmsVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetModel(v *CheckSmsVerifyCodeResponseBodyModel) *CheckSmsVerifyCodeResponseBody {
	s.Model = v
	return s
}

func (s *CheckSmsVerifyCodeResponseBody) SetSuccess(v bool) *CheckSmsVerifyCodeResponseBody {
	s.Success = &v
	return s
}

type CheckSmsVerifyCodeResponseBodyModel struct {
	// The external ID.
	//
	// example:
	//
	// 1212312
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- PASS: The verification is successful.
	//
	// 	- UNKNOWN: The verification failed.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s CheckSmsVerifyCodeResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsVerifyCodeResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponseBodyModel) SetOutId(v string) *CheckSmsVerifyCodeResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *CheckSmsVerifyCodeResponseBodyModel) SetVerifyResult(v string) *CheckSmsVerifyCodeResponseBodyModel {
	s.VerifyResult = &v
	return s
}

type CheckSmsVerifyCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSmsVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSmsVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckSmsVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *CheckSmsVerifyCodeResponse) SetHeaders(v map[string]*string) *CheckSmsVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *CheckSmsVerifyCodeResponse) SetStatusCode(v int32) *CheckSmsVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSmsVerifyCodeResponse) SetBody(v *CheckSmsVerifyCodeResponseBody) *CheckSmsVerifyCodeResponse {
	s.Body = v
	return s
}

type CreateSchemeConfigRequest struct {
	// The package name. This parameter is required when Platform is set to Android. The name must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun.android
	AndroidPackageName *string `json:"AndroidPackageName,omitempty" xml:"AndroidPackageName,omitempty"`
	// The package signature. This parameter is required when Platform is set to Android. The signature must be 32 characters in length and can contain digits and letters.
	//
	// example:
	//
	// dfsfaawklll1****olkweklk***
	AndroidPackageSign *string `json:"AndroidPackageSign,omitempty" xml:"AndroidPackageSign,omitempty"`
	// The app name, which can be up to 20 characters in length and can contain letters.
	//
	// example:
	//
	// Alibaba Cloud Communications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The reserved field. HTML5 apps are not supported.
	//
	// example:
	//
	// -
	H5Origin *string `json:"H5Origin,omitempty" xml:"H5Origin,omitempty"`
	// The reserved field. HTML5 apps are not supported.
	//
	// example:
	//
	// -
	H5Url *string `json:"H5Url,omitempty" xml:"H5Url,omitempty"`
	// The bundle ID. This parameter is required when OsType is set to iOS. The bundle ID must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun.ios
	IosBundleId *string `json:"IosBundleId,omitempty" xml:"IosBundleId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The app platform.
	//
	// Valid values:
	//
	// 	- Android
	//
	// 	- iOS
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name, which can be up to 10 characters in length and can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
}

func (s CreateSchemeConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigRequest) SetAndroidPackageName(v string) *CreateSchemeConfigRequest {
	s.AndroidPackageName = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetAndroidPackageSign(v string) *CreateSchemeConfigRequest {
	s.AndroidPackageSign = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetAppName(v string) *CreateSchemeConfigRequest {
	s.AppName = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetH5Origin(v string) *CreateSchemeConfigRequest {
	s.H5Origin = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetH5Url(v string) *CreateSchemeConfigRequest {
	s.H5Url = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetIosBundleId(v string) *CreateSchemeConfigRequest {
	s.IosBundleId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetOwnerId(v int64) *CreateSchemeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetPlatform(v string) *CreateSchemeConfigRequest {
	s.Platform = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetResourceOwnerAccount(v string) *CreateSchemeConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetResourceOwnerId(v int64) *CreateSchemeConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetSchemeName(v string) *CreateSchemeConfigRequest {
	s.SchemeName = &v
	return s
}

type CreateSchemeConfigResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11186623.0.0.5c3a662fbgeAuk).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned results.
	Model *CreateSchemeConfigResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B95B36EC-8108-4479-D3AA-2BB27F9B155A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchemeConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponseBody) SetCode(v string) *CreateSchemeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetMessage(v string) *CreateSchemeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetModel(v *CreateSchemeConfigResponseBodyModel) *CreateSchemeConfigResponseBody {
	s.Model = v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetRequestId(v string) *CreateSchemeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetSuccess(v bool) *CreateSchemeConfigResponseBody {
	s.Success = &v
	return s
}

type CreateSchemeConfigResponseBodyModel struct {
	// The service code.
	//
	// example:
	//
	// FA100000168468035
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s CreateSchemeConfigResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeConfigResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponseBodyModel) SetSchemeCode(v string) *CreateSchemeConfigResponseBodyModel {
	s.SchemeCode = &v
	return s
}

type CreateSchemeConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchemeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchemeConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSchemeConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponse) SetHeaders(v map[string]*string) *CreateSchemeConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSchemeConfigResponse) SetStatusCode(v int32) *CreateSchemeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchemeConfigResponse) SetBody(v *CreateSchemeConfigResponseBody) *CreateSchemeConfigResponse {
	s.Body = v
	return s
}

type CreateVerifySchemeRequest struct {
	// The app name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud Communications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The verification type. You can select multiple types only when the phone number verification is supported. Separate multiple types with commas (,).
	//
	// 	- **1**: phone number verification
	//
	// 	- **2**: SMS verification
	//
	// example:
	//
	// 1,2
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The bundle ID. This parameter is required when OsType is set to iOS. The bundle ID must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// example.aliyundoc.com
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The channel code of China Mobile.
	//
	// example:
	//
	// 1
	CmApiCode *int64 `json:"CmApiCode,omitempty" xml:"CmApiCode,omitempty"`
	// The channel code of China Telecom.
	//
	// example:
	//
	// 3
	CtApiCode *int64 `json:"CtApiCode,omitempty" xml:"CtApiCode,omitempty"`
	// The channel code of China Unicom.
	//
	// example:
	//
	// 2
	CuApiCode *int64 `json:"CuApiCode,omitempty" xml:"CuApiCode,omitempty"`
	// The email address that receives the key.
	//
	// example:
	//
	// username@aliyundoc.com
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HmAppIdentifier *string `json:"HmAppIdentifier,omitempty" xml:"HmAppIdentifier,omitempty"`
	HmPackageName   *string `json:"HmPackageName,omitempty" xml:"HmPackageName,omitempty"`
	HmSignName      *string `json:"HmSignName,omitempty" xml:"HmSignName,omitempty"`
	// The IP address whitelist.
	//
	// example:
	//
	// 139.9.167.181
	//
	// 122.112.210.205
	//
	// 139.9.172.0/24
	IpWhiteList *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	// The source URL of the HTML5 app page. We recommend that you specify this parameter as a domain name.
	//
	// example:
	//
	// https://h5.minexiot.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The type of the operating system for the terminal. Valid values: iOS and Android.
	//
	// This parameter is required.
	//
	// example:
	//
	// iOS
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name. This parameter is required when OsType is set to Android. The name must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun
	PackName *string `json:"PackName,omitempty" xml:"PackName,omitempty"`
	// The package signature. This parameter is required when OsType is set to Android. The signature must be 32 characters in length and can contain digits and letters.
	//
	// example:
	//
	// 123aliyun
	PackSign             *string `json:"PackSign,omitempty" xml:"PackSign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service type.
	//
	// example:
	//
	// 0
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The bound SMS signature. This parameter is valid only when AuthType is set to 2. The signature must be approved.
	//
	// example:
	//
	// Aliyun Test
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// The URL of the HTML5 app page.
	//
	// example:
	//
	// https://h5.minexiot.com/index.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeRequest) SetAppName(v string) *CreateVerifySchemeRequest {
	s.AppName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetAuthType(v string) *CreateVerifySchemeRequest {
	s.AuthType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetBundleId(v string) *CreateVerifySchemeRequest {
	s.BundleId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCmApiCode(v int64) *CreateVerifySchemeRequest {
	s.CmApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCtApiCode(v int64) *CreateVerifySchemeRequest {
	s.CtApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCuApiCode(v int64) *CreateVerifySchemeRequest {
	s.CuApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetEmail(v string) *CreateVerifySchemeRequest {
	s.Email = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmAppIdentifier(v string) *CreateVerifySchemeRequest {
	s.HmAppIdentifier = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmPackageName(v string) *CreateVerifySchemeRequest {
	s.HmPackageName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmSignName(v string) *CreateVerifySchemeRequest {
	s.HmSignName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetIpWhiteList(v string) *CreateVerifySchemeRequest {
	s.IpWhiteList = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOrigin(v string) *CreateVerifySchemeRequest {
	s.Origin = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOsType(v string) *CreateVerifySchemeRequest {
	s.OsType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOwnerId(v int64) *CreateVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackName(v string) *CreateVerifySchemeRequest {
	s.PackName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackSign(v string) *CreateVerifySchemeRequest {
	s.PackSign = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerAccount(v string) *CreateVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerId(v int64) *CreateVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSceneType(v string) *CreateVerifySchemeRequest {
	s.SceneType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSchemeName(v string) *CreateVerifySchemeRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSmsSignName(v string) *CreateVerifySchemeRequest {
	s.SmsSignName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetUrl(v string) *CreateVerifySchemeRequest {
	s.Url = &v
	return s
}

type CreateVerifySchemeResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	GateVerifySchemeDTO *CreateVerifySchemeResponseBodyGateVerifySchemeDTO `json:"GateVerifySchemeDTO,omitempty" xml:"GateVerifySchemeDTO,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1E75E69-3049-5FDB-A376-D745837CD2B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBody) SetCode(v string) *CreateVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetGateVerifySchemeDTO(v *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) *CreateVerifySchemeResponseBody {
	s.GateVerifySchemeDTO = v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetHttpStatusCode(v int64) *CreateVerifySchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetMessage(v string) *CreateVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetRequestId(v string) *CreateVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVerifySchemeResponseBody) SetSuccess(v bool) *CreateVerifySchemeResponseBody {
	s.Success = &v
	return s
}

type CreateVerifySchemeResponseBodyGateVerifySchemeDTO struct {
	// The service code.
	//
	// example:
	//
	// FC10001287****
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponseBodyGateVerifySchemeDTO) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponseBodyGateVerifySchemeDTO) SetSchemeCode(v string) *CreateVerifySchemeResponseBodyGateVerifySchemeDTO {
	s.SchemeCode = &v
	return s
}

type CreateVerifySchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponse) SetHeaders(v map[string]*string) *CreateVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySchemeResponse) SetStatusCode(v int32) *CreateVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVerifySchemeResponse) SetBody(v *CreateVerifySchemeResponseBody) *CreateVerifySchemeResponse {
	s.Body = v
	return s
}

type DeleteVerifySchemeRequest struct {
	// The user ID.
	//
	// example:
	//
	// 12345678
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000014164****
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DeleteVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeRequest) SetCustomerId(v int64) *DeleteVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerAccount(v string) *DeleteVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetSchemeCode(v string) *DeleteVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

type DeleteVerifySchemeResponseBody struct {
	// The request is successful. For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E3754956-D0B1-5947-962A-AE767D354F01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation. Valid values:
	//
	// 	- **true**: The verification service is deleted.
	//
	// 	- **false**: The verification service failed to be deleted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponseBody) SetCode(v string) *DeleteVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetMessage(v string) *DeleteVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetRequestId(v string) *DeleteVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifySchemeResponseBody) SetResult(v bool) *DeleteVerifySchemeResponseBody {
	s.Result = &v
	return s
}

type DeleteVerifySchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponse) SetHeaders(v map[string]*string) *DeleteVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifySchemeResponse) SetStatusCode(v int32) *DeleteVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVerifySchemeResponse) SetBody(v *DeleteVerifySchemeResponseBody) *DeleteVerifySchemeResponse {
	s.Body = v
	return s
}

type DescribeVerifySchemeRequest struct {
	// The user ID.
	//
	// example:
	//
	// 1234****
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000010643****
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DescribeVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeRequest) SetCustomerId(v int64) *DescribeVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerAccount(v string) *DescribeVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetResourceOwnerId(v int64) *DescribeVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVerifySchemeRequest) SetSchemeCode(v string) *DescribeVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

type DescribeVerifySchemeResponseBody struct {
	// The response code. OK indicates that the request is successful. For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0C5380A7-2032-5F7D-9614-1BF8B54D16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	SchemeQueryResultDTO *DescribeVerifySchemeResponseBodySchemeQueryResultDTO `json:"SchemeQueryResultDTO,omitempty" xml:"SchemeQueryResultDTO,omitempty" type:"Struct"`
}

func (s DescribeVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBody) SetCode(v string) *DescribeVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetMessage(v string) *DescribeVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetRequestId(v string) *DescribeVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetSchemeQueryResultDTO(v *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) *DescribeVerifySchemeResponseBody {
	s.SchemeQueryResultDTO = v
	return s
}

type DescribeVerifySchemeResponseBodySchemeQueryResultDTO struct {
	// The key generated when you create a service in the console.
	//
	// example:
	//
	// ZDMARqPkyQzWVJjB/sB/+fCp5TA4lNsRnY7rEC+HfGsOIOk1Brj8UyXFW2RBYIWqLieCSo8ZypEaEj+h9rLd3FgpXAjGYDfmOperod6jPUUwFHhBObxK+HuKVoi2jOqN7aDOlyPyGcATyq3BDdlf922JmnFLT8Hvnu4qgzzCZk0LXWTb0XVPnm5/fHUGHEA2Q+aTrGkaWcHjmTDqQ7BtvrAIIcJJkCJu4i1aeU++/0EzGWap4mcb2VhKROBs****
	AppEncryptInfo *string `json:"AppEncryptInfo,omitempty" xml:"AppEncryptInfo,omitempty"`
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) SetAppEncryptInfo(v string) *DescribeVerifySchemeResponseBodySchemeQueryResultDTO {
	s.AppEncryptInfo = &v
	return s
}

type DescribeVerifySchemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponse) SetHeaders(v map[string]*string) *DescribeVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySchemeResponse) SetStatusCode(v int32) *DescribeVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySchemeResponse) SetBody(v *DescribeVerifySchemeResponseBody) *DescribeVerifySchemeResponse {
	s.Body = v
	return s
}

type GetAuthTokenRequest struct {
	// The requested domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyundoc.com
	Origin               *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SceneCode            *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The URL of the requested web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyundoc.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthTokenRequest) SetOrigin(v string) *GetAuthTokenRequest {
	s.Origin = &v
	return s
}

func (s *GetAuthTokenRequest) SetOwnerId(v int64) *GetAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerAccount(v string) *GetAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthTokenRequest) SetResourceOwnerId(v int64) *GetAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthTokenRequest) SetSceneCode(v string) *GetAuthTokenRequest {
	s.SceneCode = &v
	return s
}

func (s *GetAuthTokenRequest) SetUrl(v string) *GetAuthTokenRequest {
	s.Url = &v
	return s
}

type GetAuthTokenResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	TokenInfo *GetAuthTokenResponseBodyTokenInfo `json:"TokenInfo,omitempty" xml:"TokenInfo,omitempty" type:"Struct"`
}

func (s GetAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBody) SetCode(v string) *GetAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetMessage(v string) *GetAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetRequestId(v string) *GetAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthTokenResponseBody) SetTokenInfo(v *GetAuthTokenResponseBodyTokenInfo) *GetAuthTokenResponseBody {
	s.TokenInfo = v
	return s
}

type GetAuthTokenResponseBodyTokenInfo struct {
	// The business authentication token.
	//
	// >  AccessToken is valid for 10 minutes and can be used repeatedly within its validity period.
	//
	// example:
	//
	// agag****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The API authentication token.
	//
	// >  JwtToken is valid for 1 hour and can be used repeatedly within its validity period.
	//
	// example:
	//
	// aweghd****
	JwtToken *string `json:"JwtToken,omitempty" xml:"JwtToken,omitempty"`
}

func (s GetAuthTokenResponseBodyTokenInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponseBodyTokenInfo) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetAccessToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.AccessToken = &v
	return s
}

func (s *GetAuthTokenResponseBodyTokenInfo) SetJwtToken(v string) *GetAuthTokenResponseBodyTokenInfo {
	s.JwtToken = &v
	return s
}

type GetAuthTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthTokenResponse) SetHeaders(v map[string]*string) *GetAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthTokenResponse) SetStatusCode(v int32) *GetAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthTokenResponse) SetBody(v *GetAuthTokenResponseBody) *GetAuthTokenResponse {
	s.Body = v
	return s
}

type GetAuthorizationUrlRequest struct {
	// The authorization end date, which is in the yyyy-MM-dd format. This parameter is required for services of contract type.
	//
	// example:
	//
	// 2020–12–28
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the authorization scenario. You can view the ID of the authorization scenario on the **Authorization Scenario Management*	- page in the **Phone Number Verification Service console**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 234****
	SchemeId *int64 `json:"SchemeId,omitempty" xml:"SchemeId,omitempty"`
}

func (s GetAuthorizationUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlRequest) SetEndDate(v string) *GetAuthorizationUrlRequest {
	s.EndDate = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetPhoneNo(v string) *GetAuthorizationUrlRequest {
	s.PhoneNo = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerAccount(v string) *GetAuthorizationUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetResourceOwnerId(v int64) *GetAuthorizationUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAuthorizationUrlRequest) SetSchemeId(v int64) *GetAuthorizationUrlRequest {
	s.SchemeId = &v
	return s
}

type GetAuthorizationUrlResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetAuthorizationUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBody) SetCode(v string) *GetAuthorizationUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetData(v *GetAuthorizationUrlResponseBodyData) *GetAuthorizationUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetMessage(v string) *GetAuthorizationUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthorizationUrlResponseBody) SetRequestId(v string) *GetAuthorizationUrlResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthorizationUrlResponseBodyData struct {
	// The authorization URL.
	//
	// example:
	//
	// https://render.****.com/p/s/web-call-minapp/auth-bao?page=commauth/index&token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJiaXpUeXBlIjoiQ29tbUF1dGgiLCJiaXpJZCI6IjVmNWZjNjAzZDQzMTQ0MWZiYTZiNjYzM2QyMjIyNzU0IiwiZXhwIjoxNjA4MTkxODQxfQ.5IvBj2nKgr60APtotaIB13vtPVrdsPQ6avIfWxte1pA&_env=prod
	AuthorizationUrl *string `json:"AuthorizationUrl,omitempty" xml:"AuthorizationUrl,omitempty"`
}

func (s GetAuthorizationUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponseBodyData) SetAuthorizationUrl(v string) *GetAuthorizationUrlResponseBodyData {
	s.AuthorizationUrl = &v
	return s
}

type GetAuthorizationUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationUrlResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationUrlResponse) SetHeaders(v map[string]*string) *GetAuthorizationUrlResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationUrlResponse) SetStatusCode(v int32) *GetAuthorizationUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationUrlResponse) SetBody(v *GetAuthorizationUrlResponseBody) *GetAuthorizationUrlResponse {
	s.Body = v
	return s
}

type GetFusionAuthTokenRequest struct {
	// The bundle ID of the app. This parameter is required when Platform is set to iOS.
	//
	// example:
	//
	// com.example.test
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The validity period of the token. Unit: seconds. Valid values: 900 to 43200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 900
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name of the app. This parameter is required when Platform is set to Android.
	//
	// example:
	//
	// com.example.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The package signature of the app. This parameter is required when Platform is set to Android.
	//
	// example:
	//
	// 47fcc************************278
	PackageSign *string `json:"PackageSign,omitempty" xml:"PackageSign,omitempty"`
	// The platform type. Valid values: Android and iOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FA1000*************201
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s GetFusionAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFusionAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenRequest) SetBundleId(v string) *GetFusionAuthTokenRequest {
	s.BundleId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetDurationSeconds(v int64) *GetFusionAuthTokenRequest {
	s.DurationSeconds = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetOwnerId(v int64) *GetFusionAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPackageName(v string) *GetFusionAuthTokenRequest {
	s.PackageName = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPackageSign(v string) *GetFusionAuthTokenRequest {
	s.PackageSign = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPlatform(v string) *GetFusionAuthTokenRequest {
	s.Platform = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetResourceOwnerAccount(v string) *GetFusionAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetResourceOwnerId(v int64) *GetFusionAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetSchemeCode(v string) *GetFusionAuthTokenRequest {
	s.SchemeCode = &v
	return s
}

type GetFusionAuthTokenResponseBody struct {
	// The response code. If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The authentication code. The value of this parameter is a string.
	//
	// example:
	//
	// FKcksloqk***********jalEc+
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFusionAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFusionAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenResponseBody) SetCode(v string) *GetFusionAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetMessage(v string) *GetFusionAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetModel(v string) *GetFusionAuthTokenResponseBody {
	s.Model = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetRequestId(v string) *GetFusionAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetSuccess(v bool) *GetFusionAuthTokenResponseBody {
	s.Success = &v
	return s
}

type GetFusionAuthTokenResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFusionAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFusionAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFusionAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenResponse) SetHeaders(v map[string]*string) *GetFusionAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *GetFusionAuthTokenResponse) SetStatusCode(v int32) *GetFusionAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFusionAuthTokenResponse) SetBody(v *GetFusionAuthTokenResponseBody) *GetFusionAuthTokenResponse {
	s.Body = v
	return s
}

type GetMobileRequest struct {
	// The logon token obtained by the SDK for your app.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad5422****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 22345****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMobileRequest) GoString() string {
	return s.String()
}

func (s *GetMobileRequest) SetAccessToken(v string) *GetMobileRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMobileRequest) SetOutId(v string) *GetMobileRequest {
	s.OutId = &v
	return s
}

func (s *GetMobileRequest) SetOwnerId(v int64) *GetMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerAccount(v string) *GetMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerId(v int64) *GetMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMobileResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	GetMobileResultDTO *GetMobileResponseBodyGetMobileResultDTO `json:"GetMobileResultDTO,omitempty" xml:"GetMobileResultDTO,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponseBody) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBody) SetCode(v string) *GetMobileResponseBody {
	s.Code = &v
	return s
}

func (s *GetMobileResponseBody) SetGetMobileResultDTO(v *GetMobileResponseBodyGetMobileResultDTO) *GetMobileResponseBody {
	s.GetMobileResultDTO = v
	return s
}

func (s *GetMobileResponseBody) SetMessage(v string) *GetMobileResponseBody {
	s.Message = &v
	return s
}

func (s *GetMobileResponseBody) SetRequestId(v string) *GetMobileResponseBody {
	s.RequestId = &v
	return s
}

type GetMobileResponseBodyGetMobileResultDTO struct {
	// The phone number,
	//
	// example:
	//
	// 13900001234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetMobileResponseBodyGetMobileResultDTO) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponseBodyGetMobileResultDTO) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBodyGetMobileResultDTO) SetMobile(v string) *GetMobileResponseBodyGetMobileResultDTO {
	s.Mobile = &v
	return s
}

type GetMobileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMobileResponse) GoString() string {
	return s.String()
}

func (s *GetMobileResponse) SetHeaders(v map[string]*string) *GetMobileResponse {
	s.Headers = v
	return s
}

func (s *GetMobileResponse) SetStatusCode(v int32) *GetMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMobileResponse) SetBody(v *GetMobileResponseBody) *GetMobileResponse {
	s.Body = v
	return s
}

type GetPhoneWithTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token for phone number verification that is obtained by the JavaScript SDK. The validity period of the token is 10 minutes for China Telecom, 30 minutes for China Unicom, and 2 minutes for China Mobile. The token can be used only once.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	SpToken *string `json:"SpToken,omitempty" xml:"SpToken,omitempty"`
}

func (s GetPhoneWithTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneWithTokenRequest) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenRequest) SetOwnerId(v int64) *GetPhoneWithTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *GetPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetResourceOwnerId(v int64) *GetPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPhoneWithTokenRequest) SetSpToken(v string) *GetPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

type GetPhoneWithTokenResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetPhoneWithTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0F335F48-****-****-****-CA7914FE5D77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPhoneWithTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponseBody) SetCode(v string) *GetPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetData(v *GetPhoneWithTokenResponseBodyData) *GetPhoneWithTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetMessage(v string) *GetPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhoneWithTokenResponseBody) SetRequestId(v string) *GetPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetPhoneWithTokenResponseBodyData struct {
	// The phone number.
	//
	// example:
	//
	// 13900001234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetPhoneWithTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneWithTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponseBodyData) SetMobile(v string) *GetPhoneWithTokenResponseBodyData {
	s.Mobile = &v
	return s
}

type GetPhoneWithTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneWithTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhoneWithTokenResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponse) SetHeaders(v map[string]*string) *GetPhoneWithTokenResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneWithTokenResponse) SetStatusCode(v int32) *GetPhoneWithTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneWithTokenResponse) SetBody(v *GetPhoneWithTokenResponseBody) *GetPhoneWithTokenResponse {
	s.Body = v
	return s
}

type GetSmsAuthTokensRequest struct {
	// The ID of the iOS application. This parameter is required if OsType is set to **iOS**.
	//
	// example:
	//
	// 12345****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The validity period of the token. Unit: seconds. Valid values: 900 to 43200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 900
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The type of the operating system. Valid values: **Android*	- and **iOS**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name. This parameter is required if OsType is set to **Android**.
	//
	// example:
	//
	// com.aliqin.mytel.test
	PackageName          *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC100000134840112
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The signature. This parameter is required if OsType is set to **Android**.
	//
	// example:
	//
	// 47fcc6615485e83b4100433****
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The validity period of the SMS verification code. Unit: seconds. Default value: 180.
	//
	// example:
	//
	// 60
	SmsCodeExpire *int32 `json:"SmsCodeExpire,omitempty" xml:"SmsCodeExpire,omitempty"`
	// The code of the text message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_13987****
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
}

func (s GetSmsAuthTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensRequest) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensRequest) SetBundleId(v string) *GetSmsAuthTokensRequest {
	s.BundleId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetExpire(v int64) *GetSmsAuthTokensRequest {
	s.Expire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOsType(v string) *GetSmsAuthTokensRequest {
	s.OsType = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetPackageName(v string) *GetSmsAuthTokensRequest {
	s.PackageName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerAccount(v string) *GetSmsAuthTokensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSceneCode(v string) *GetSmsAuthTokensRequest {
	s.SceneCode = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSignName(v string) *GetSmsAuthTokensRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsCodeExpire(v int32) *GetSmsAuthTokensRequest {
	s.SmsCodeExpire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsTemplateCode(v string) *GetSmsAuthTokensRequest {
	s.SmsTemplateCode = &v
	return s
}

type GetSmsAuthTokensResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetSmsAuthTokensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSmsAuthTokensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBody) SetCode(v string) *GetSmsAuthTokensResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetData(v *GetSmsAuthTokensResponseBodyData) *GetSmsAuthTokensResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetMessage(v string) *GetSmsAuthTokensResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsAuthTokensResponseBody) SetRequestId(v string) *GetSmsAuthTokensResponseBody {
	s.RequestId = &v
	return s
}

type GetSmsAuthTokensResponseBodyData struct {
	// The business token.
	//
	// example:
	//
	// FZSMeRbLCiapIBo65NXNHvGbkGDmhs23XWTZDOtZN0g5n/kqSc2FU27Gc9YhGb6dNn9/L9ZXSYiIB6C6LMLQJjyXjRzt5v6pzZXqnjO4cSuPWYUxJvdc8l8OpucEYe8Mx17HxsHDzj0VC4D5+atcrTpJE6jQ7e2QVNjZIPMwsfxELjQS7c****
	BizToken *string `json:"BizToken,omitempty" xml:"BizToken,omitempty"`
	// The time when the token expired. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1631526326000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NSqC****
	StsAccessKeyId *string `json:"StsAccessKeyId,omitempty" xml:"StsAccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// VboZ4xbZ****
	StsAccessKeySecret *string `json:"StsAccessKeySecret,omitempty" xml:"StsAccessKeySecret,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAISiQJ1q6Ft5B2yfSjIr5DEDP/BurtW9PemMEfBrEpsOr5K17XjuDz2IHtLfXFsBusYt/U2nWpX5v4clrxIToR7SFbFY9pb6ZhazBisebDGv8HtR3TcFEjiSwapEBfe8JL4QYeQFaHwGJqEb1TDiVUAo9/TfimjWFqIKICAjYUdAP0cQgi/a0gtZr4UXHwAzvUXLnzML/2gHwf3i27LdipStxF7lHl05NbUoKTeyGKH3AGqlLVF9tite8f9NpczBvolDYfpht4RX7HazStd5yJN8KpLl6Fe8V/FxIrGXAAJv0rdbbOFq4Q1c18hOLJHAKtfsvXmlPNpsevfmpnsx****
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
}

func (s GetSmsAuthTokensResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponseBodyData) SetBizToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.BizToken = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetExpireTime(v int64) *GetSmsAuthTokensResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeyId(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeyId = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsAccessKeySecret(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsAccessKeySecret = &v
	return s
}

func (s *GetSmsAuthTokensResponseBodyData) SetStsToken(v string) *GetSmsAuthTokensResponseBodyData {
	s.StsToken = &v
	return s
}

type GetSmsAuthTokensResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsAuthTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsAuthTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSmsAuthTokensResponse) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensResponse) SetHeaders(v map[string]*string) *GetSmsAuthTokensResponse {
	s.Headers = v
	return s
}

func (s *GetSmsAuthTokensResponse) SetStatusCode(v int32) *GetSmsAuthTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsAuthTokensResponse) SetBody(v *GetSmsAuthTokensResponseBody) *GetSmsAuthTokensResponse {
	s.Body = v
	return s
}

type JyCreateVerifySchemeRequest struct {
	// This parameter is required.
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// 1
	CmApiCode *int64 `json:"CmApiCode,omitempty" xml:"CmApiCode,omitempty"`
	// example:
	//
	// 3
	CtApiCode *int64 `json:"CtApiCode,omitempty" xml:"CtApiCode,omitempty"`
	// example:
	//
	// 2
	CuApiCode *int64 `json:"CuApiCode,omitempty" xml:"CuApiCode,omitempty"`
	// This parameter is required.
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PackName             *string `json:"PackName,omitempty" xml:"PackName,omitempty"`
	PackSign             *string `json:"PackSign,omitempty" xml:"PackSign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
}

func (s JyCreateVerifySchemeRequest) String() string {
	return tea.Prettify(s)
}

func (s JyCreateVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *JyCreateVerifySchemeRequest) SetAppName(v string) *JyCreateVerifySchemeRequest {
	s.AppName = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetBundleId(v string) *JyCreateVerifySchemeRequest {
	s.BundleId = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetCmApiCode(v int64) *JyCreateVerifySchemeRequest {
	s.CmApiCode = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetCtApiCode(v int64) *JyCreateVerifySchemeRequest {
	s.CtApiCode = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetCuApiCode(v int64) *JyCreateVerifySchemeRequest {
	s.CuApiCode = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetOsType(v string) *JyCreateVerifySchemeRequest {
	s.OsType = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetOwnerId(v int64) *JyCreateVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetPackName(v string) *JyCreateVerifySchemeRequest {
	s.PackName = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetPackSign(v string) *JyCreateVerifySchemeRequest {
	s.PackSign = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetResourceOwnerAccount(v string) *JyCreateVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetResourceOwnerId(v int64) *JyCreateVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JyCreateVerifySchemeRequest) SetSchemeName(v string) *JyCreateVerifySchemeRequest {
	s.SchemeName = &v
	return s
}

type JyCreateVerifySchemeResponseBody struct {
	Code                 *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	GateVerifySchemeData *JyCreateVerifySchemeResponseBodyGateVerifySchemeData `json:"GateVerifySchemeData,omitempty" xml:"GateVerifySchemeData,omitempty" type:"Struct"`
	Message              *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JyCreateVerifySchemeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JyCreateVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *JyCreateVerifySchemeResponseBody) SetCode(v string) *JyCreateVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *JyCreateVerifySchemeResponseBody) SetGateVerifySchemeData(v *JyCreateVerifySchemeResponseBodyGateVerifySchemeData) *JyCreateVerifySchemeResponseBody {
	s.GateVerifySchemeData = v
	return s
}

func (s *JyCreateVerifySchemeResponseBody) SetMessage(v string) *JyCreateVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *JyCreateVerifySchemeResponseBody) SetRequestId(v string) *JyCreateVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

type JyCreateVerifySchemeResponseBodyGateVerifySchemeData struct {
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s JyCreateVerifySchemeResponseBodyGateVerifySchemeData) String() string {
	return tea.Prettify(s)
}

func (s JyCreateVerifySchemeResponseBodyGateVerifySchemeData) GoString() string {
	return s.String()
}

func (s *JyCreateVerifySchemeResponseBodyGateVerifySchemeData) SetSchemeCode(v string) *JyCreateVerifySchemeResponseBodyGateVerifySchemeData {
	s.SchemeCode = &v
	return s
}

type JyCreateVerifySchemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JyCreateVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JyCreateVerifySchemeResponse) String() string {
	return tea.Prettify(s)
}

func (s JyCreateVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *JyCreateVerifySchemeResponse) SetHeaders(v map[string]*string) *JyCreateVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *JyCreateVerifySchemeResponse) SetStatusCode(v int32) *JyCreateVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *JyCreateVerifySchemeResponse) SetBody(v *JyCreateVerifySchemeResponseBody) *JyCreateVerifySchemeResponse {
	s.Body = v
	return s
}

type JyQueryAppInfoBySceneCodeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s JyQueryAppInfoBySceneCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s JyQueryAppInfoBySceneCodeRequest) GoString() string {
	return s.String()
}

func (s *JyQueryAppInfoBySceneCodeRequest) SetOwnerId(v int64) *JyQueryAppInfoBySceneCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeRequest) SetResourceOwnerAccount(v string) *JyQueryAppInfoBySceneCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeRequest) SetResourceOwnerId(v int64) *JyQueryAppInfoBySceneCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeRequest) SetSceneCode(v string) *JyQueryAppInfoBySceneCodeRequest {
	s.SceneCode = &v
	return s
}

type JyQueryAppInfoBySceneCodeResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *JyQueryAppInfoBySceneCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JyQueryAppInfoBySceneCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JyQueryAppInfoBySceneCodeResponseBody) GoString() string {
	return s.String()
}

func (s *JyQueryAppInfoBySceneCodeResponseBody) SetCode(v string) *JyQueryAppInfoBySceneCodeResponseBody {
	s.Code = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBody) SetData(v *JyQueryAppInfoBySceneCodeResponseBodyData) *JyQueryAppInfoBySceneCodeResponseBody {
	s.Data = v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBody) SetMessage(v string) *JyQueryAppInfoBySceneCodeResponseBody {
	s.Message = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBody) SetRequestId(v string) *JyQueryAppInfoBySceneCodeResponseBody {
	s.RequestId = &v
	return s
}

type JyQueryAppInfoBySceneCodeResponseBodyData struct {
	CmAppId  *string `json:"CmAppId,omitempty" xml:"CmAppId,omitempty"`
	CmAppKey *string `json:"CmAppKey,omitempty" xml:"CmAppKey,omitempty"`
	CtAppId  *string `json:"CtAppId,omitempty" xml:"CtAppId,omitempty"`
	CtAppKey *string `json:"CtAppKey,omitempty" xml:"CtAppKey,omitempty"`
}

func (s JyQueryAppInfoBySceneCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s JyQueryAppInfoBySceneCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *JyQueryAppInfoBySceneCodeResponseBodyData) SetCmAppId(v string) *JyQueryAppInfoBySceneCodeResponseBodyData {
	s.CmAppId = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBodyData) SetCmAppKey(v string) *JyQueryAppInfoBySceneCodeResponseBodyData {
	s.CmAppKey = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBodyData) SetCtAppId(v string) *JyQueryAppInfoBySceneCodeResponseBodyData {
	s.CtAppId = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponseBodyData) SetCtAppKey(v string) *JyQueryAppInfoBySceneCodeResponseBodyData {
	s.CtAppKey = &v
	return s
}

type JyQueryAppInfoBySceneCodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JyQueryAppInfoBySceneCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JyQueryAppInfoBySceneCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s JyQueryAppInfoBySceneCodeResponse) GoString() string {
	return s.String()
}

func (s *JyQueryAppInfoBySceneCodeResponse) SetHeaders(v map[string]*string) *JyQueryAppInfoBySceneCodeResponse {
	s.Headers = v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponse) SetStatusCode(v int32) *JyQueryAppInfoBySceneCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *JyQueryAppInfoBySceneCodeResponse) SetBody(v *JyQueryAppInfoBySceneCodeResponseBody) *JyQueryAppInfoBySceneCodeResponse {
	s.Body = v
	return s
}

type QueryGateVerifyBillingPublicRequest struct {
	// The verification method. Valid values:
	//
	// 	- **0**: phone number verification
	//
	// 	- **1**: one-click logon
	//
	// 	- **2**: all
	//
	// 	- **3**: facial recognition
	//
	// 	- **4**: SMS verification
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The month in which the bill is generated. Specify this parameter in the YYYYMM format. Example: 202111.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202111
	Month                *string `json:"Month,omitempty" xml:"Month,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s QueryGateVerifyBillingPublicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyBillingPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetMonth(v string) *QueryGateVerifyBillingPublicRequest {
	s.Month = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetOwnerId(v int64) *QueryGateVerifyBillingPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyBillingPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyBillingPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type QueryGateVerifyBillingPublicResponseBody struct {
	// The response code. Valid values:
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The billing information about each verification service.
	Data *QueryGateVerifyBillingPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetCode(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetData(v *QueryGateVerifyBillingPublicResponseBodyData) *QueryGateVerifyBillingPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetMessage(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBody) SetRequestId(v string) *QueryGateVerifyBillingPublicResponseBody {
	s.RequestId = &v
	return s
}

type QueryGateVerifyBillingPublicResponseBodyData struct {
	// The fees generated for all verification services. Unitrogen: CNY.
	//
	// example:
	//
	// 1234
	AmountSum *string `json:"AmountSum,omitempty" xml:"AmountSum,omitempty"`
	// The details of fees.
	SceneBillingList []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList `json:"SceneBillingList,omitempty" xml:"SceneBillingList,omitempty" type:"Repeated"`
}

func (s QueryGateVerifyBillingPublicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetAmountSum(v string) *QueryGateVerifyBillingPublicResponseBodyData {
	s.AmountSum = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyData) SetSceneBillingList(v []*QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) *QueryGateVerifyBillingPublicResponseBodyData {
	s.SceneBillingList = v
	return s
}

type QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList struct {
	// The billable items.
	//
	// example:
	//
	// 74
	Add *string `json:"Add,omitempty" xml:"Add,omitempty"`
	// The fees generated for the verification service. Unitrogen: CNY.
	//
	// example:
	//
	// 1.48
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The application name.
	//
	// example:
	//
	// Aliyun
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The verification method.
	//
	// example:
	//
	// Verification of local phone number
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// The service code.
	//
	// example:
	//
	// FC100000038194004
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The service name.
	//
	// example:
	//
	// Alibaba Cloud Communications
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The unit price. Unit: CNY.
	//
	// example:
	//
	// 0.02
	SinglePrice *string `json:"SinglePrice,omitempty" xml:"SinglePrice,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAdd(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Add = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAmount(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.Amount = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetAppName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.AppName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetItemName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.ItemName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneCode(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSceneName(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SceneName = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList) SetSinglePrice(v string) *QueryGateVerifyBillingPublicResponseBodyDataSceneBillingList {
	s.SinglePrice = &v
	return s
}

type QueryGateVerifyBillingPublicResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyBillingPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyBillingPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyBillingPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyBillingPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyBillingPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) SetStatusCode(v int32) *QueryGateVerifyBillingPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyBillingPublicResponse) SetBody(v *QueryGateVerifyBillingPublicResponseBody) *QueryGateVerifyBillingPublicResponse {
	s.Body = v
	return s
}

type QueryGateVerifyStatisticPublicRequest struct {
	// The verification method. Valid values:
	//
	// 	- **1**: one-click logon
	//
	// 	- **2**: phone number verification, including the verification of the phone number used in HTML5 pages
	//
	// 	- **3**: SMS verification
	//
	// 	- **4**: facial recognition
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthenticationType *int32 `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// The end date. Specify this parameter in the YYYYMMDD format. Example: 20220106.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20220106
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **Android**
	//
	// 	- **iOS**
	//
	// example:
	//
	// Android
	OsType               *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The service code.
	//
	// example:
	//
	// FC100000038194004
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The start date. Specify this parameter in the YYYYMMDD format. Example: 20220101.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20220101
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryGateVerifyStatisticPublicRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicRequest) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicRequest) SetAuthenticationType(v int32) *QueryGateVerifyStatisticPublicRequest {
	s.AuthenticationType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetEndDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.EndDate = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOsType(v string) *QueryGateVerifyStatisticPublicRequest {
	s.OsType = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetOwnerId(v int64) *QueryGateVerifyStatisticPublicRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetResourceOwnerAccount(v string) *QueryGateVerifyStatisticPublicRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetSceneCode(v string) *QueryGateVerifyStatisticPublicRequest {
	s.SceneCode = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicRequest) SetStartDate(v string) *QueryGateVerifyStatisticPublicRequest {
	s.StartDate = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBody struct {
	// The response code. Valid values:
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the calls of Phone Number Verification Service, including the total calls, the successful calls, failed calls, unknown calls, and daily calls within the statistical date range.
	Data *QueryGateVerifyStatisticPublicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetCode(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetData(v *QueryGateVerifyStatisticPublicResponseBodyData) *QueryGateVerifyStatisticPublicResponseBody {
	s.Data = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetMessage(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBody) SetRequestId(v string) *QueryGateVerifyStatisticPublicResponseBody {
	s.RequestId = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBodyData struct {
	// The information about the daily calls.
	DayStatistic []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic `json:"DayStatistic,omitempty" xml:"DayStatistic,omitempty" type:"Repeated"`
	// The total calls.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The failed calls.
	//
	// example:
	//
	// 20
	TotalFail *int64 `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	// The successful calls.
	//
	// example:
	//
	// 0
	TotalSuccess *int64 `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	// The unknown calls.
	//
	// example:
	//
	// 0
	TotalUnknown *int64 `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetDayStatistic(v []*QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.DayStatistic = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotal(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyData) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyData {
	s.TotalUnknown = &v
	return s
}

type QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic struct {
	// The date. This field is accurate to the day. The value of this field is in the YYYYMMDD format. Example: 20220103.
	//
	// example:
	//
	// 20220103
	StatisticDateStr *string `json:"StatisticDateStr,omitempty" xml:"StatisticDateStr,omitempty"`
	// The failed calls on the day.
	//
	// example:
	//
	// 20
	TotalFail *int64 `json:"TotalFail,omitempty" xml:"TotalFail,omitempty"`
	// The successful calls on the day.
	//
	// example:
	//
	// 0
	TotalSuccess *int64 `json:"TotalSuccess,omitempty" xml:"TotalSuccess,omitempty"`
	// The unknown calls on the day.
	//
	// example:
	//
	// 0
	TotalUnknown *int64 `json:"TotalUnknown,omitempty" xml:"TotalUnknown,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetStatisticDateStr(v string) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.StatisticDateStr = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalFail(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalFail = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalSuccess(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalSuccess = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic) SetTotalUnknown(v int64) *QueryGateVerifyStatisticPublicResponseBodyDataDayStatistic {
	s.TotalUnknown = &v
	return s
}

type QueryGateVerifyStatisticPublicResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyStatisticPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyStatisticPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) SetStatusCode(v int32) *QueryGateVerifyStatisticPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) SetBody(v *QueryGateVerifyStatisticPublicResponseBody) *QueryGateVerifyStatisticPublicResponse {
	s.Body = v
	return s
}

type QuerySendDetailsRequest struct {
	// The unique ID of the business, which is provided by Alibaba Cloud.
	//
	// example:
	//
	// 1231891289318923^12
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The number of the page on which you are reading the text message. Pages start from page 1. The value of this parameter cannot exceed the maximum page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 186****9399
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The date when the text message was sent. You can query text messages that were sent within the last 30 days.
	//
	// Specify the date in the yyyyMMdd format. Example: 20181225.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20181225
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
}

func (s QuerySendDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsRequest) SetBizId(v string) *QuerySendDetailsRequest {
	s.BizId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetCurrentPage(v int64) *QuerySendDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuerySendDetailsRequest) SetOwnerId(v int64) *QuerySendDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPageSize(v int64) *QuerySendDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPhoneNumber(v string) *QuerySendDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerAccount(v string) *QuerySendDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerId(v int64) *QuerySendDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetSendDate(v string) *QuerySendDetailsRequest {
	s.SendDate = &v
	return s
}

type QuerySendDetailsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// none
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html?spm=a2c4g.419277.0.i11).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model []*QuerySendDetailsResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Repeated"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySendDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBody) SetAccessDeniedDetail(v string) *QuerySendDetailsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetCode(v string) *QuerySendDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetMessage(v string) *QuerySendDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetModel(v []*QuerySendDetailsResponseBodyModel) *QuerySendDetailsResponseBody {
	s.Model = v
	return s
}

func (s *QuerySendDetailsResponseBody) SetSuccess(v bool) *QuerySendDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySendDetailsResponseBody) SetTotalCount(v int64) *QuerySendDetailsResponseBody {
	s.TotalCount = &v
	return s
}

type QuerySendDetailsResponseBodyModel struct {
	// The content of the text message.
	//
	// example:
	//
	// 203160
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The status code returned by the carrier.
	//
	// 	- If the text message was delivered, "DELIVERED" is returned.
	//
	// 	- If the text message failed to be sent, see [Error codes](https://help.aliyun.com/document_detail/101347.html?spm=a2c4g.419277.0.i8) for more information.
	//
	// example:
	//
	// DELIVERED
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The extension field.
	//
	// example:
	//
	// 12131231
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 1390000****
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The date and time when the text message was received.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	ReceiveDate *string `json:"ReceiveDate,omitempty" xml:"ReceiveDate,omitempty"`
	// The date when the text message was sent. You can query text messages that were sent within the last 30 days.
	//
	// The date is in the yyyyMMdd format. Example: 20181225.
	//
	// example:
	//
	// 2019-01-08 16:44:13
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
	// The delivery status of the text message.
	//
	// 	- 1: A delivery receipt is to be sent.
	//
	// 	- 2: The text message failed to be sent.
	//
	// 	- 3: The text message was sent.
	//
	// example:
	//
	// 3
	SendStatus *int64 `json:"SendStatus,omitempty" xml:"SendStatus,omitempty"`
	// The code of the text message template.
	//
	// Log on to the SMS console. In the left-side navigation pane, click **Go China*	- or **Go Globe**. You can view the text message template code in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// >  The text message templates must be created on the Go Globe page and approved.
	//
	// example:
	//
	// SMS_12231****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s QuerySendDetailsResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponseBodyModel) SetContent(v string) *QuerySendDetailsResponseBodyModel {
	s.Content = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetErrCode(v string) *QuerySendDetailsResponseBodyModel {
	s.ErrCode = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetOutId(v string) *QuerySendDetailsResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetPhoneNum(v string) *QuerySendDetailsResponseBodyModel {
	s.PhoneNum = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetReceiveDate(v string) *QuerySendDetailsResponseBodyModel {
	s.ReceiveDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetSendDate(v string) *QuerySendDetailsResponseBodyModel {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetSendStatus(v int64) *QuerySendDetailsResponseBodyModel {
	s.SendStatus = &v
	return s
}

func (s *QuerySendDetailsResponseBodyModel) SetTemplateCode(v string) *QuerySendDetailsResponseBodyModel {
	s.TemplateCode = &v
	return s
}

type QuerySendDetailsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySendDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySendDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySendDetailsResponse) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsResponse) SetHeaders(v map[string]*string) *QuerySendDetailsResponse {
	s.Headers = v
	return s
}

func (s *QuerySendDetailsResponse) SetStatusCode(v int32) *QuerySendDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySendDetailsResponse) SetBody(v *QuerySendDetailsResponseBody) *QuerySendDetailsResponse {
	s.Body = v
	return s
}

type SendSmsVerifyCodeRequest struct {
	// The length of the verification code. Default value: 4. Valid values: 4 to 8.
	//
	// example:
	//
	// 4
	CodeLength *int64 `json:"CodeLength,omitempty" xml:"CodeLength,omitempty"`
	// The type of the generated verification code. Default value: 1. Valid values:
	//
	// 	- 1: digits only
	//
	// 	- 2: uppercase letters only
	//
	// 	- 3: lowercase letters only
	//
	// 	- 4: uppercase and lowercase letters
	//
	// 	- 5: digits and uppercase letters
	//
	// 	- 6: digits and lowercase letters
	//
	// 	- 7: digits and uppercase and lowercase letters
	//
	// example:
	//
	// 1
	CodeType *int64 `json:"CodeType,omitempty" xml:"CodeType,omitempty"`
	// The country code of the phone number. SMS verification codes can be sent only by using phone numbers in the Chinese mainland. Default value: 86.
	//
	// example:
	//
	// 86
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// Specifies how to handle the verification codes received earlier in a case where verification codes are sent to the same phone number for the same scenario within the validity period.
	//
	// 	- 1 (default): The latest verification code overwrites the verification codes received earlier. In this case, verification codes received earlier expire.
	//
	// 	- 2: Verification codes within their validity period are valid and can be used for verification.
	//
	// example:
	//
	// 1
	DuplicatePolicy *int64 `json:"DuplicatePolicy,omitempty" xml:"DuplicatePolicy,omitempty"`
	// The time interval. Unit: seconds. Default value: 60. This parameter specifies how often you can send a verification code.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 12358794Aqzaq
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 86130****0000
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to return a verification code.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ReturnVerifyCode *bool `json:"ReturnVerifyCode,omitempty" xml:"ReturnVerifyCode,omitempty"`
	// The verification service name. If this parameter is not specified, the default service is used. The name can be up to 20 characters in length.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The signature.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Test
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The extension code of the upstream text message. Upstream text messages are text messages sent to the communication service provider. Upstream text messages are used to customize a service, complete an inquiry, or send a request. You are charged for sending upstream text messages based on the billing standards of the service provider.
	//
	// >  The extension code is automatically generated by the system when the signature is generated. You do not need to specify the extension code. You can skip this parameter based on your business requirements. If you want to use custom extension codes, contact your account manager.
	//
	// example:
	//
	// 1213123
	SmsUpExtendCode *string `json:"SmsUpExtendCode,omitempty" xml:"SmsUpExtendCode,omitempty"`
	// The code of the text message template.
	//
	// Log on to the [SMS console](https://dysms.console.aliyun.com/dysms.htm?spm=5176.12818093.categories-n-products.ddysms.3b2816d0xml2NA#/overview). In the left-side navigation pane, click **Go China*	- or **Go Globe**. You can view the text message template code in the **Template Code*	- column on the **Message Templates*	- tab.
	//
	// >  The text message templates must be created on the Go Globe page and approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// azsq_*****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The value of the variable in the text message template. The verification code is replaced with "##code##".
	//
	// Example 1: For a system-defined template that contains variables, if the template content is "Your verification code is ${code} and valid for 5 minutes. Do not disclose the verification code to others.", specify the value of this parameter as {"code":"##code##"}
	//
	// Example 2: For a custom template, if the template content is ${content}, specify the value of this parameter as {"content":"Your verification code is ##code## and must be used within 5 minutes."}.
	//
	// >
	//
	// 	- If line breaks are required in JSON-formatted data, they must meet the relevant requirements that are specified in the standard JSON protocol.
	//
	// 	- For more information about template variables, see [SMS template specifications](https://help.aliyun.com/document_detail/108253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"code":"##code##"}
	TemplateParam *string `json:"TemplateParam,omitempty" xml:"TemplateParam,omitempty"`
	// The validity period of the verification code. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	ValidTime *int64 `json:"ValidTime,omitempty" xml:"ValidTime,omitempty"`
}

func (s SendSmsVerifyCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendSmsVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeRequest) SetCodeLength(v int64) *SendSmsVerifyCodeRequest {
	s.CodeLength = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetCodeType(v int64) *SendSmsVerifyCodeRequest {
	s.CodeType = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetCountryCode(v string) *SendSmsVerifyCodeRequest {
	s.CountryCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetDuplicatePolicy(v int64) *SendSmsVerifyCodeRequest {
	s.DuplicatePolicy = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetInterval(v int64) *SendSmsVerifyCodeRequest {
	s.Interval = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetOutId(v string) *SendSmsVerifyCodeRequest {
	s.OutId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetOwnerId(v int64) *SendSmsVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetPhoneNumber(v string) *SendSmsVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetResourceOwnerAccount(v string) *SendSmsVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetResourceOwnerId(v int64) *SendSmsVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetReturnVerifyCode(v bool) *SendSmsVerifyCodeRequest {
	s.ReturnVerifyCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSchemeName(v string) *SendSmsVerifyCodeRequest {
	s.SchemeName = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSignName(v string) *SendSmsVerifyCodeRequest {
	s.SignName = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetSmsUpExtendCode(v string) *SendSmsVerifyCodeRequest {
	s.SmsUpExtendCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetTemplateCode(v string) *SendSmsVerifyCodeRequest {
	s.TemplateCode = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetTemplateParam(v string) *SendSmsVerifyCodeRequest {
	s.TemplateParam = &v
	return s
}

func (s *SendSmsVerifyCodeRequest) SetValidTime(v int64) *SendSmsVerifyCodeRequest {
	s.ValidTime = &v
	return s
}

type SendSmsVerifyCodeResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code. If OK is returned, the request is successful. For more information, see [Response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11174283.0.0.70c5616bkj38Wa).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *SendSmsVerifyCodeResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendSmsVerifyCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendSmsVerifyCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponseBody) SetAccessDeniedDetail(v string) *SendSmsVerifyCodeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetCode(v string) *SendSmsVerifyCodeResponseBody {
	s.Code = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetMessage(v string) *SendSmsVerifyCodeResponseBody {
	s.Message = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetModel(v *SendSmsVerifyCodeResponseBodyModel) *SendSmsVerifyCodeResponseBody {
	s.Model = v
	return s
}

func (s *SendSmsVerifyCodeResponseBody) SetSuccess(v bool) *SendSmsVerifyCodeResponseBody {
	s.Success = &v
	return s
}

type SendSmsVerifyCodeResponseBodyModel struct {
	// The business ID.
	//
	// example:
	//
	// 112231421412414124123^4
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 1231231313
	OutId *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// API-reqelekrqkllkkewrlwrjlsdfsdf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification code.
	//
	// example:
	//
	// 42324
	VerifyCode *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s SendSmsVerifyCodeResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s SendSmsVerifyCodeResponseBodyModel) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetBizId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.BizId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetOutId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.OutId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetRequestId(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.RequestId = &v
	return s
}

func (s *SendSmsVerifyCodeResponseBodyModel) SetVerifyCode(v string) *SendSmsVerifyCodeResponseBodyModel {
	s.VerifyCode = &v
	return s
}

type SendSmsVerifyCodeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendSmsVerifyCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendSmsVerifyCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendSmsVerifyCodeResponse) GoString() string {
	return s.String()
}

func (s *SendSmsVerifyCodeResponse) SetHeaders(v map[string]*string) *SendSmsVerifyCodeResponse {
	s.Headers = v
	return s
}

func (s *SendSmsVerifyCodeResponse) SetStatusCode(v int32) *SendSmsVerifyCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendSmsVerifyCodeResponse) SetBody(v *SendSmsVerifyCodeResponseBody) *SendSmsVerifyCodeResponse {
	s.Body = v
	return s
}

type VerifyMobileRequest struct {
	// The token obtained by the SDK for your app.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	AccessCode *string `json:"AccessCode,omitempty" xml:"AccessCode,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 123456
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13800****00
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s VerifyMobileRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileRequest) GoString() string {
	return s.String()
}

func (s *VerifyMobileRequest) SetAccessCode(v string) *VerifyMobileRequest {
	s.AccessCode = &v
	return s
}

func (s *VerifyMobileRequest) SetOutId(v string) *VerifyMobileRequest {
	s.OutId = &v
	return s
}

func (s *VerifyMobileRequest) SetOwnerId(v int64) *VerifyMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyMobileRequest) SetPhoneNumber(v string) *VerifyMobileRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerAccount(v string) *VerifyMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyMobileRequest) SetResourceOwnerId(v int64) *VerifyMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

type VerifyMobileResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	GateVerifyResultDTO *VerifyMobileResponseBodyGateVerifyResultDTO `json:"GateVerifyResultDTO,omitempty" xml:"GateVerifyResultDTO,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyMobileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBody) SetCode(v string) *VerifyMobileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyMobileResponseBody) SetGateVerifyResultDTO(v *VerifyMobileResponseBodyGateVerifyResultDTO) *VerifyMobileResponseBody {
	s.GateVerifyResultDTO = v
	return s
}

func (s *VerifyMobileResponseBody) SetMessage(v string) *VerifyMobileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyMobileResponseBody) SetRequestId(v string) *VerifyMobileResponseBody {
	s.RequestId = &v
	return s
}

type VerifyMobileResponseBodyGateVerifyResultDTO struct {
	// The verification ID.
	//
	// example:
	//
	// 121343241
	VerifyId *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- **PASS: The input phone number is consistent with the phone number that you use.**
	//
	// 	- **REJECT: The input phone number is different from the phone number that you use.**
	//
	// 	- **UNKNOWN: The system cannot judge whether the input phone number is consistent with the phone number that you use.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyId(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyId = &v
	return s
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyResult(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyResult = &v
	return s
}

type VerifyMobileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyMobileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyMobileResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyMobileResponse) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponse) SetHeaders(v map[string]*string) *VerifyMobileResponse {
	s.Headers = v
	return s
}

func (s *VerifyMobileResponse) SetStatusCode(v int32) *VerifyMobileResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyMobileResponse) SetBody(v *VerifyMobileResponseBody) *VerifyMobileResponse {
	s.Body = v
	return s
}

type VerifyPhoneWithTokenRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1380000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token for phone number verification that is obtained by the JavaScript SDK.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad542****
	SpToken *string `json:"SpToken,omitempty" xml:"SpToken,omitempty"`
}

func (s VerifyPhoneWithTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenRequest) SetOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetPhoneNumber(v string) *VerifyPhoneWithTokenRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerAccount(v string) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetResourceOwnerId(v int64) *VerifyPhoneWithTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyPhoneWithTokenRequest) SetSpToken(v string) *VerifyPhoneWithTokenRequest {
	s.SpToken = &v
	return s
}

type VerifyPhoneWithTokenResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	GateVerify *VerifyPhoneWithTokenResponseBodyGateVerify `json:"GateVerify,omitempty" xml:"GateVerify,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBody) SetCode(v string) *VerifyPhoneWithTokenResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetGateVerify(v *VerifyPhoneWithTokenResponseBodyGateVerify) *VerifyPhoneWithTokenResponseBody {
	s.GateVerify = v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetMessage(v string) *VerifyPhoneWithTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBody) SetRequestId(v string) *VerifyPhoneWithTokenResponseBody {
	s.RequestId = &v
	return s
}

type VerifyPhoneWithTokenResponseBodyGateVerify struct {
	// The external ID.
	//
	// example:
	//
	// 12134****
	VerifyId *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- PASS: The input phone number is consistent with the phone number used in HTML5 pages.
	//
	// 	- REJECT: The input phone number is different from the phone number used in HTML5 pages.
	//
	// 	- UNKNOWN: The system cannot judge whether the input phone number is consistent with the phone number used in HTML5 pages.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponseBodyGateVerify) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyId(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyId = &v
	return s
}

func (s *VerifyPhoneWithTokenResponseBodyGateVerify) SetVerifyResult(v string) *VerifyPhoneWithTokenResponseBodyGateVerify {
	s.VerifyResult = &v
	return s
}

type VerifyPhoneWithTokenResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyPhoneWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyPhoneWithTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyPhoneWithTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyPhoneWithTokenResponse) SetHeaders(v map[string]*string) *VerifyPhoneWithTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyPhoneWithTokenResponse) SetStatusCode(v int32) *VerifyPhoneWithTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyPhoneWithTokenResponse) SetBody(v *VerifyPhoneWithTokenResponseBody) *VerifyPhoneWithTokenResponse {
	s.Body = v
	return s
}

type VerifySmsCodeRequest struct {
	// The phone number, which is used to receive SMS verification codes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1321111****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The SMS verification code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12****
	SmsCode *string `json:"SmsCode,omitempty" xml:"SmsCode,omitempty"`
	// The text message verification code. After you successfully call the corresponding API operation to send the SMS verification code, the end users receive the SMS verification code. SmsToken is returned by the SDK for SMS verification for you to verify the text message verification code. For an Android client, sendVerifyCode is called to send the verification code. For an iOS client, sendVerifyCodeWithTimeout is called to send the verification code. For more information, see [Overview](https://help.aliyun.com/document_detail/400434.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// sddsbvdbvjd****
	SmsToken *string `json:"SmsToken,omitempty" xml:"SmsToken,omitempty"`
}

func (s VerifySmsCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeRequest) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeRequest) SetPhoneNumber(v string) *VerifySmsCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsCode(v string) *VerifySmsCodeRequest {
	s.SmsCode = &v
	return s
}

func (s *VerifySmsCodeRequest) SetSmsToken(v string) *VerifySmsCodeRequest {
	s.SmsToken = &v
	return s
}

type VerifySmsCodeResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifySmsCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponseBody) SetCode(v string) *VerifySmsCodeResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetData(v bool) *VerifySmsCodeResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetMessage(v string) *VerifySmsCodeResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySmsCodeResponseBody) SetRequestId(v string) *VerifySmsCodeResponseBody {
	s.RequestId = &v
	return s
}

type VerifySmsCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySmsCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySmsCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySmsCodeResponse) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponse) SetHeaders(v map[string]*string) *VerifySmsCodeResponse {
	s.Headers = v
	return s
}

func (s *VerifySmsCodeResponse) SetStatusCode(v int32) *VerifySmsCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySmsCodeResponse) SetBody(v *VerifySmsCodeResponseBody) *VerifySmsCodeResponse {
	s.Body = v
	return s
}

type VerifyWithFusionAuthTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unified verification token that is returned by the client SDKs.
	//
	// This parameter is required.
	//
	// example:
	//
	// LD108enNdlsl*******sFLKCks1==
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
}

func (s VerifyWithFusionAuthTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyWithFusionAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenRequest) SetOwnerId(v int64) *VerifyWithFusionAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetResourceOwnerAccount(v string) *VerifyWithFusionAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetResourceOwnerId(v int64) *VerifyWithFusionAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenRequest) SetVerifyToken(v string) *VerifyWithFusionAuthTokenRequest {
	s.VerifyToken = &v
	return s
}

type VerifyWithFusionAuthTokenResponseBody struct {
	// The response code. If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned data.
	Model *VerifyWithFusionAuthTokenResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID, which is used to troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetCode(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetMessage(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetModel(v *VerifyWithFusionAuthTokenResponseBodyModel) *VerifyWithFusionAuthTokenResponseBody {
	s.Model = v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetRequestId(v string) *VerifyWithFusionAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBody) SetSuccess(v bool) *VerifyWithFusionAuthTokenResponseBody {
	s.Success = &v
	return s
}

type VerifyWithFusionAuthTokenResponseBodyModel struct {
	// The phone number, which is returned when the verification is successful.
	//
	// example:
	//
	// 180********
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The phone number score, which is generated only after the phone number scoring node is enabled and the verification is successful. The higher the score, the more risky the phone number. Valid values: 0 to 100.
	//
	// example:
	//
	// 20
	PhoneScore *int64 `json:"PhoneScore,omitempty" xml:"PhoneScore,omitempty"`
	// The verification result. Valid values: PASS and UNKNOWN.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponseBodyModel) String() string {
	return tea.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponseBodyModel) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetPhoneNumber(v string) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.PhoneNumber = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetPhoneScore(v int64) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.PhoneScore = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponseBodyModel) SetVerifyResult(v string) *VerifyWithFusionAuthTokenResponseBodyModel {
	s.VerifyResult = &v
	return s
}

type VerifyWithFusionAuthTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyWithFusionAuthTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyWithFusionAuthTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyWithFusionAuthTokenResponse) GoString() string {
	return s.String()
}

func (s *VerifyWithFusionAuthTokenResponse) SetHeaders(v map[string]*string) *VerifyWithFusionAuthTokenResponse {
	s.Headers = v
	return s
}

func (s *VerifyWithFusionAuthTokenResponse) SetStatusCode(v int32) *VerifyWithFusionAuthTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyWithFusionAuthTokenResponse) SetBody(v *VerifyWithFusionAuthTokenResponseBody) *VerifyWithFusionAuthTokenResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dypnsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Verifies SMS verification codes.
//
// @param request - CheckSmsVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSmsVerifyCodeResponse
func (client *Client) CheckSmsVerifyCodeWithOptions(request *CheckSmsVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *CheckSmsVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaseAuthPolicy)) {
		query["CaseAuthPolicy"] = request.CaseAuthPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		query["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckSmsVerifyCode"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckSmsVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies SMS verification codes.
//
// @param request - CheckSmsVerifyCodeRequest
//
// @return CheckSmsVerifyCodeResponse
func (client *Client) CheckSmsVerifyCode(request *CheckSmsVerifyCodeRequest) (_result *CheckSmsVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckSmsVerifyCodeResponse{}
	_body, _err := client.CheckSmsVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a code for a converged communication authentication service.
//
// @param request - CreateSchemeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeConfigResponse
func (client *Client) CreateSchemeConfigWithOptions(request *CreateSchemeConfigRequest, runtime *util.RuntimeOptions) (_result *CreateSchemeConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AndroidPackageName)) {
		query["AndroidPackageName"] = request.AndroidPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.AndroidPackageSign)) {
		query["AndroidPackageSign"] = request.AndroidPackageSign
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.H5Origin)) {
		query["H5Origin"] = request.H5Origin
	}

	if !tea.BoolValue(util.IsUnset(request.H5Url)) {
		query["H5Url"] = request.H5Url
	}

	if !tea.BoolValue(util.IsUnset(request.IosBundleId)) {
		query["IosBundleId"] = request.IosBundleId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchemeConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSchemeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a code for a converged communication authentication service.
//
// @param request - CreateSchemeConfigRequest
//
// @return CreateSchemeConfigResponse
func (client *Client) CreateSchemeConfig(request *CreateSchemeConfigRequest) (_result *CreateSchemeConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSchemeConfigResponse{}
	_body, _err := client.CreateSchemeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySchemeResponse
func (client *Client) CreateVerifySchemeWithOptions(request *CreateVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *CreateVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.CmApiCode)) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.CtApiCode)) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.CuApiCode)) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.HmAppIdentifier)) {
		query["HmAppIdentifier"] = request.HmAppIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.HmPackageName)) {
		query["HmPackageName"] = request.HmPackageName
	}

	if !tea.BoolValue(util.IsUnset(request.HmSignName)) {
		query["HmSignName"] = request.HmSignName
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhiteList)) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackName)) {
		query["PackName"] = request.PackName
	}

	if !tea.BoolValue(util.IsUnset(request.PackSign)) {
		query["PackSign"] = request.PackSign
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		query["SceneType"] = request.SceneType
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsSignName)) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CreateVerifySchemeRequest
//
// @return CreateVerifySchemeResponse
func (client *Client) CreateVerifyScheme(request *CreateVerifySchemeRequest) (_result *CreateVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CreateVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVerifySchemeResponse
func (client *Client) DeleteVerifySchemeWithOptions(request *DeleteVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *DeleteVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DeleteVerifySchemeRequest
//
// @return DeleteVerifySchemeResponse
func (client *Client) DeleteVerifyScheme(request *DeleteVerifySchemeRequest) (_result *DeleteVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVerifySchemeResponse{}
	_body, _err := client.DeleteVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifySchemeResponse
func (client *Client) DescribeVerifySchemeWithOptions(request *DescribeVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *DescribeVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 100 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - DescribeVerifySchemeRequest
//
// @return DescribeVerifySchemeResponse
func (client *Client) DescribeVerifyScheme(request *DescribeVerifySchemeRequest) (_result *DescribeVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVerifySchemeResponse{}
	_body, _err := client.DescribeVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the authorization token used for the authentication of the phone number verification for HTML5 pages. You can obtain AccessToken and JwtToken after a successful call.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthTokenWithOptions(request *GetAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Origin)) {
		query["Origin"] = request.Origin
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authorization token used for the authentication of the phone number verification for HTML5 pages. You can obtain AccessToken and JwtToken after a successful call.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthTokenRequest
//
// @return GetAuthTokenResponse
func (client *Client) GetAuthToken(request *GetAuthTokenRequest) (_result *GetAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthTokenResponse{}
	_body, _err := client.GetAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the URL for the Alipay account authorization.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account and obtain an Alibaba Cloud AccessKey pair. For more information, see [Process of communication authorization](https://help.aliyun.com/document_detail/196922.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthorizationUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationUrlResponse
func (client *Client) GetAuthorizationUrlWithOptions(request *GetAuthorizationUrlRequest, runtime *util.RuntimeOptions) (_result *GetAuthorizationUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNo)) {
		query["PhoneNo"] = request.PhoneNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeId)) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthorizationUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the URL for the Alipay account authorization.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account and obtain an Alibaba Cloud AccessKey pair. For more information, see [Process of communication authorization](https://help.aliyun.com/document_detail/196922.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetAuthorizationUrlRequest
//
// @return GetAuthorizationUrlResponse
func (client *Client) GetAuthorizationUrl(request *GetAuthorizationUrlRequest) (_result *GetAuthorizationUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthorizationUrlResponse{}
	_body, _err := client.GetAuthorizationUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the token that is obtained from the client SDKs.
//
// @param request - GetFusionAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFusionAuthTokenResponse
func (client *Client) GetFusionAuthTokenWithOptions(request *GetFusionAuthTokenRequest, runtime *util.RuntimeOptions) (_result *GetFusionAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.DurationSeconds)) {
		query["DurationSeconds"] = request.DurationSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.PackageSign)) {
		query["PackageSign"] = request.PackageSign
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeCode)) {
		query["SchemeCode"] = request.SchemeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFusionAuthToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFusionAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the token that is obtained from the client SDKs.
//
// @param request - GetFusionAuthTokenRequest
//
// @return GetFusionAuthTokenResponse
func (client *Client) GetFusionAuthToken(request *GetFusionAuthTokenRequest) (_result *GetFusionAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFusionAuthTokenResponse{}
	_body, _err := client.GetFusionAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration. You can call this operation only after you confirm the authorization on the authorization page provided by the SDK for one-click logon. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMobileResponse
func (client *Client) GetMobileWithOptions(request *GetMobileRequest, runtime *util.RuntimeOptions) (_result *GetMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessToken)) {
		query["AccessToken"] = request.AccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMobile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration. You can call this operation only after you confirm the authorization on the authorization page provided by the SDK for one-click logon. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetMobileRequest
//
// @return GetMobileResponse
func (client *Client) GetMobile(request *GetMobileRequest) (_result *GetMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMobileResponse{}
	_body, _err := client.GetMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon. This operation is exclusive to HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration in HTML5 pages. You can call this operation only after you confirm the authorization on the authorization page provided by the JavaScript SDK. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneWithTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhoneWithTokenResponse
func (client *Client) GetPhoneWithTokenWithOptions(request *GetPhoneWithTokenRequest, runtime *util.RuntimeOptions) (_result *GetPhoneWithTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpToken)) {
		query["SpToken"] = request.SpToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPhoneWithToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPhoneWithTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a phone number for one-click logon. This operation is exclusive to HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable only to one-click logon or registration in HTML5 pages. You can call this operation only after you confirm the authorization on the authorization page provided by the JavaScript SDK. You are prohibited from simulating or bypassing the authorization process. Alibaba Cloud reserves the right to terminate our services and take legal actions against such violations.
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetPhoneWithTokenRequest
//
// @return GetPhoneWithTokenResponse
func (client *Client) GetPhoneWithToken(request *GetPhoneWithTokenRequest) (_result *GetPhoneWithTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhoneWithTokenResponse{}
	_body, _err := client.GetPhoneWithTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the authorization token for an SMS verification code.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSmsAuthTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsAuthTokensResponse
func (client *Client) GetSmsAuthTokensWithOptions(request *GetSmsAuthTokensRequest, runtime *util.RuntimeOptions) (_result *GetSmsAuthTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.Expire)) {
		query["Expire"] = request.Expire
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsCodeExpire)) {
		query["SmsCodeExpire"] = request.SmsCodeExpire
	}

	if !tea.BoolValue(util.IsUnset(request.SmsTemplateCode)) {
		query["SmsTemplateCode"] = request.SmsTemplateCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSmsAuthTokens"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSmsAuthTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the authorization token for an SMS verification code.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - GetSmsAuthTokensRequest
//
// @return GetSmsAuthTokensResponse
func (client *Client) GetSmsAuthTokens(request *GetSmsAuthTokensRequest) (_result *GetSmsAuthTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSmsAuthTokensResponse{}
	_body, _err := client.GetSmsAuthTokensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI JyCreateVerifyScheme is deprecated, please use Dypnsapi::2017-05-25::CreateVerifyScheme instead.
//
// Summary:
//
// 创建方案号（为极意临时定制）
//
// @param request - JyCreateVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JyCreateVerifySchemeResponse
// Deprecated
func (client *Client) JyCreateVerifySchemeWithOptions(request *JyCreateVerifySchemeRequest, runtime *util.RuntimeOptions) (_result *JyCreateVerifySchemeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BundleId)) {
		query["BundleId"] = request.BundleId
	}

	if !tea.BoolValue(util.IsUnset(request.CmApiCode)) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.CtApiCode)) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.CuApiCode)) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackName)) {
		query["PackName"] = request.PackName
	}

	if !tea.BoolValue(util.IsUnset(request.PackSign)) {
		query["PackSign"] = request.PackSign
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("JyCreateVerifyScheme"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JyCreateVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI JyCreateVerifyScheme is deprecated, please use Dypnsapi::2017-05-25::CreateVerifyScheme instead.
//
// Summary:
//
// 创建方案号（为极意临时定制）
//
// @param request - JyCreateVerifySchemeRequest
//
// @return JyCreateVerifySchemeResponse
// Deprecated
func (client *Client) JyCreateVerifyScheme(request *JyCreateVerifySchemeRequest) (_result *JyCreateVerifySchemeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JyCreateVerifySchemeResponse{}
	_body, _err := client.JyCreateVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI JyQueryAppInfoBySceneCode is deprecated, please use Dypnsapi::2017-05-25::QueryAppInfoBySceneCode instead.
//
// Summary:
//
// 根据方案号查询运营商APP信（为极意临时定制）
//
// @param request - JyQueryAppInfoBySceneCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JyQueryAppInfoBySceneCodeResponse
// Deprecated
func (client *Client) JyQueryAppInfoBySceneCodeWithOptions(request *JyQueryAppInfoBySceneCodeRequest, runtime *util.RuntimeOptions) (_result *JyQueryAppInfoBySceneCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("JyQueryAppInfoBySceneCode"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JyQueryAppInfoBySceneCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI JyQueryAppInfoBySceneCode is deprecated, please use Dypnsapi::2017-05-25::QueryAppInfoBySceneCode instead.
//
// Summary:
//
// 根据方案号查询运营商APP信（为极意临时定制）
//
// @param request - JyQueryAppInfoBySceneCodeRequest
//
// @return JyQueryAppInfoBySceneCodeResponse
// Deprecated
func (client *Client) JyQueryAppInfoBySceneCode(request *JyQueryAppInfoBySceneCodeRequest) (_result *JyQueryAppInfoBySceneCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JyQueryAppInfoBySceneCodeResponse{}
	_body, _err := client.JyQueryAppInfoBySceneCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the fees generated by a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyBillingPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyBillingPublicResponse
func (client *Client) QueryGateVerifyBillingPublicWithOptions(request *QueryGateVerifyBillingPublicRequest, runtime *util.RuntimeOptions) (_result *QueryGateVerifyBillingPublicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationType)) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !tea.BoolValue(util.IsUnset(request.Month)) {
		query["Month"] = request.Month
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGateVerifyBillingPublic"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGateVerifyBillingPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the fees generated by a verification service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyBillingPublicRequest
//
// @return QueryGateVerifyBillingPublicResponse
func (client *Client) QueryGateVerifyBillingPublic(request *QueryGateVerifyBillingPublicRequest) (_result *QueryGateVerifyBillingPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGateVerifyBillingPublicResponse{}
	_body, _err := client.QueryGateVerifyBillingPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the calls of Phone Number Verification Service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyStatisticPublicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyStatisticPublicResponse
func (client *Client) QueryGateVerifyStatisticPublicWithOptions(request *QueryGateVerifyStatisticPublicRequest, runtime *util.RuntimeOptions) (_result *QueryGateVerifyStatisticPublicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationType)) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.OsType)) {
		query["OsType"] = request.OsType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryGateVerifyStatisticPublic"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryGateVerifyStatisticPublicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the calls of Phone Number Verification Service.
//
// Description:
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - QueryGateVerifyStatisticPublicRequest
//
// @return QueryGateVerifyStatisticPublicResponse
func (client *Client) QueryGateVerifyStatisticPublic(request *QueryGateVerifyStatisticPublicRequest) (_result *QueryGateVerifyStatisticPublicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryGateVerifyStatisticPublicResponse{}
	_body, _err := client.QueryGateVerifyStatisticPublicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI QuerySendDetails is deprecated
//
// Summary:
//
// Queries the delivery status of the SMS verification code. You can query only the delivery status of the SMS verification code that is sent by calling corresponding API operations.
//
// @param request - QuerySendDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySendDetailsResponse
// Deprecated
func (client *Client) QuerySendDetailsWithOptions(request *QuerySendDetailsRequest, runtime *util.RuntimeOptions) (_result *QuerySendDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SendDate)) {
		query["SendDate"] = request.SendDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySendDetails"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI QuerySendDetails is deprecated
//
// Summary:
//
// Queries the delivery status of the SMS verification code. You can query only the delivery status of the SMS verification code that is sent by calling corresponding API operations.
//
// @param request - QuerySendDetailsRequest
//
// @return QuerySendDetailsResponse
// Deprecated
func (client *Client) QuerySendDetails(request *QuerySendDetailsRequest) (_result *QuerySendDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySendDetailsResponse{}
	_body, _err := client.QuerySendDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends SMS verification codes.
//
// @param request - SendSmsVerifyCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSmsVerifyCodeResponse
func (client *Client) SendSmsVerifyCodeWithOptions(request *SendSmsVerifyCodeRequest, runtime *util.RuntimeOptions) (_result *SendSmsVerifyCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeLength)) {
		query["CodeLength"] = request.CodeLength
	}

	if !tea.BoolValue(util.IsUnset(request.CodeType)) {
		query["CodeType"] = request.CodeType
	}

	if !tea.BoolValue(util.IsUnset(request.CountryCode)) {
		query["CountryCode"] = request.CountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.DuplicatePolicy)) {
		query["DuplicatePolicy"] = request.DuplicatePolicy
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnVerifyCode)) {
		query["ReturnVerifyCode"] = request.ReturnVerifyCode
	}

	if !tea.BoolValue(util.IsUnset(request.SchemeName)) {
		query["SchemeName"] = request.SchemeName
	}

	if !tea.BoolValue(util.IsUnset(request.SignName)) {
		query["SignName"] = request.SignName
	}

	if !tea.BoolValue(util.IsUnset(request.SmsUpExtendCode)) {
		query["SmsUpExtendCode"] = request.SmsUpExtendCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateParam)) {
		query["TemplateParam"] = request.TemplateParam
	}

	if !tea.BoolValue(util.IsUnset(request.ValidTime)) {
		query["ValidTime"] = request.ValidTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendSmsVerifyCode"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendSmsVerifyCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends SMS verification codes.
//
// @param request - SendSmsVerifyCodeRequest
//
// @return SendSmsVerifyCodeResponse
func (client *Client) SendSmsVerifyCode(request *SendSmsVerifyCodeRequest) (_result *SendSmsVerifyCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendSmsVerifyCodeResponse{}
	_body, _err := client.SendSmsVerifyCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the phone number that you use.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable to only the verification of thephone number that you use. To obtain a phone number for one-click logon, call [GetMobile](https://help.aliyun.com/document_detail/189865.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyMobileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyMobileResponse
func (client *Client) VerifyMobileWithOptions(request *VerifyMobileRequest, runtime *util.RuntimeOptions) (_result *VerifyMobileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessCode)) {
		query["AccessCode"] = request.AccessCode
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyMobile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyMobileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the phone number that you use.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Getting Started](https://help.aliyun.com/document_detail/84541.html).
//
// >  This operation is applicable to only the verification of thephone number that you use. To obtain a phone number for one-click logon, call [GetMobile](https://help.aliyun.com/document_detail/189865.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 5,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyMobileRequest
//
// @return VerifyMobileResponse
func (client *Client) VerifyMobile(request *VerifyMobileRequest) (_result *VerifyMobileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyMobileResponse{}
	_body, _err := client.VerifyMobileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the phone number used in HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyPhoneWithTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyPhoneWithTokenResponse
func (client *Client) VerifyPhoneWithTokenWithOptions(request *VerifyPhoneWithTokenRequest, runtime *util.RuntimeOptions) (_result *VerifyPhoneWithTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SpToken)) {
		query["SpToken"] = request.SpToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyPhoneWithToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the phone number used in HTML5 pages.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the phone number verification feature for HTML5 pages](https://help.aliyun.com/document_detail/169786.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 1,000 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifyPhoneWithTokenRequest
//
// @return VerifyPhoneWithTokenResponse
func (client *Client) VerifyPhoneWithToken(request *VerifyPhoneWithTokenRequest) (_result *VerifyPhoneWithTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyPhoneWithTokenResponse{}
	_body, _err := client.VerifyPhoneWithTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies SMS verification codes.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifySmsCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySmsCodeResponse
func (client *Client) VerifySmsCodeWithOptions(request *VerifySmsCodeRequest, runtime *util.RuntimeOptions) (_result *VerifySmsCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.SmsCode)) {
		query["SmsCode"] = request.SmsCode
	}

	if !tea.BoolValue(util.IsUnset(request.SmsToken)) {
		query["SmsToken"] = request.SmsToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySmsCode"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies SMS verification codes.
//
// Description:
//
// ### [](#)Preparations
//
// You must register an Alibaba Cloud account, obtain an Alibaba Cloud AccessKey pair, and create a verification service. For more information, see [Use the SMS verification feature](https://help.aliyun.com/document_detail/313209.html).
//
// ### [](#qps)QPS limits
//
// You can call this operation up to 500 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - VerifySmsCodeRequest
//
// @return VerifySmsCodeResponse
func (client *Client) VerifySmsCode(request *VerifySmsCodeRequest) (_result *VerifySmsCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifySmsCodeResponse{}
	_body, _err := client.VerifySmsCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the verification token that is obtained by using the authentication token.
//
// @param request - VerifyWithFusionAuthTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyWithFusionAuthTokenResponse
func (client *Client) VerifyWithFusionAuthTokenWithOptions(request *VerifyWithFusionAuthTokenRequest, runtime *util.RuntimeOptions) (_result *VerifyWithFusionAuthTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyToken)) {
		query["VerifyToken"] = request.VerifyToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyWithFusionAuthToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyWithFusionAuthTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the verification results by using the verification token that is obtained by using the authentication token.
//
// @param request - VerifyWithFusionAuthTokenRequest
//
// @return VerifyWithFusionAuthTokenResponse
func (client *Client) VerifyWithFusionAuthToken(request *VerifyWithFusionAuthTokenRequest) (_result *VerifyWithFusionAuthTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyWithFusionAuthTokenResponse{}
	_body, _err := client.VerifyWithFusionAuthTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
