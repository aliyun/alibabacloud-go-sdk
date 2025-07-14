// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddClientIdToOIDCProviderRequest struct {
	// The client ID that you want to add.
	//
	// The client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods, (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// The client ID can be up to 128 characters in length.
	//
	// example:
	//
	// 598469743454717****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s AddClientIdToOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s AddClientIdToOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderRequest) SetClientId(v string) *AddClientIdToOIDCProviderRequest {
	s.ClientId = &v
	return s
}

func (s *AddClientIdToOIDCProviderRequest) SetOIDCProviderName(v string) *AddClientIdToOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type AddClientIdToOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *AddClientIdToOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B49B8FA7-87FC-5308-AE75-5EC9E36C360A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClientIdToOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponseBody) SetOIDCProvider(v *AddClientIdToOIDCProviderResponseBodyOIDCProvider) *AddClientIdToOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBody) SetRequestId(v string) *AddClientIdToOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type AddClientIdToOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 498469743454717****,598469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s AddClientIdToOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type AddClientIdToOIDCProviderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddClientIdToOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddClientIdToOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponse) SetHeaders(v map[string]*string) *AddClientIdToOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *AddClientIdToOIDCProviderResponse) SetStatusCode(v int32) *AddClientIdToOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponse) SetBody(v *AddClientIdToOIDCProviderResponseBody) *AddClientIdToOIDCProviderResponse {
	s.Body = v
	return s
}

type AddFingerprintToOIDCProviderRequest struct {
	// The fingerprint of the HTTPS CA certificate.
	//
	// The fingerprint can contain letters and digits.
	//
	// The fingerprint can be up to 128 characters in length.
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s AddFingerprintToOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFingerprintToOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderRequest) SetFingerprint(v string) *AddFingerprintToOIDCProviderRequest {
	s.Fingerprint = &v
	return s
}

func (s *AddFingerprintToOIDCProviderRequest) SetOIDCProviderName(v string) *AddFingerprintToOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type AddFingerprintToOIDCProviderResponseBody struct {
	// The name of the OIDC IdP.
	OIDCProvider *AddFingerprintToOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B809BBC-0E78-544A-A91A-648926412E3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponseBody) SetOIDCProvider(v *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) *AddFingerprintToOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBody) SetRequestId(v string) *AddFingerprintToOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type AddFingerprintToOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/OIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 502ef2deeb3c5b13ea4c3d5193629309e231****,902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type AddFingerprintToOIDCProviderResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFingerprintToOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponse) SetHeaders(v map[string]*string) *AddFingerprintToOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *AddFingerprintToOIDCProviderResponse) SetStatusCode(v int32) *AddFingerprintToOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponse) SetBody(v *AddFingerprintToOIDCProviderResponseBody) *AddFingerprintToOIDCProviderResponse {
	s.Body = v
	return s
}

type AddUserToGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) SetGroupName(v string) *AddUserToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AddUserToGroupRequest) SetUserPrincipalName(v string) *AddUserToGroupRequest {
	s.UserPrincipalName = &v
	return s
}

type AddUserToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponseBody) SetRequestId(v string) *AddUserToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponse) SetHeaders(v map[string]*string) *AddUserToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUserToGroupResponse) SetStatusCode(v int32) *AddUserToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToGroupResponse) SetBody(v *AddUserToGroupResponseBody) *AddUserToGroupResponse {
	s.Body = v
	return s
}

type BindMFADeviceRequest struct {
	// The first verification code.
	//
	// >  You can call the [CreateVirtualMFADevice](https://help.aliyun.com/document_detail/186179.html) operation to create an MFA device and generate a key (value of `Base32StringSeed`). Then, use the key on the Alibaba Cloud app to manually add an MFA device, and obtain the two consecutive verification codes.
	//
	// example:
	//
	// 123456
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	// The second verification code.
	//
	// >  You can call the [CreateVirtualMFADevice](https://help.aliyun.com/document_detail/186179.html) operation to create an MFA device and generate a key (value of `Base32StringSeed`). Then, use the key on the Alibaba Cloud app to manually add an MFA device, and obtain the two consecutive verification codes.
	//
	// example:
	//
	// 654321
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
	// The serial number of the MFA device.
	//
	// >  You can call the [CreateVirtualMFADevice](https://help.aliyun.com/document_detail/186179.html) operation to obtain the serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s BindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *BindMFADeviceRequest) SetAuthenticationCode1(v string) *BindMFADeviceRequest {
	s.AuthenticationCode1 = &v
	return s
}

func (s *BindMFADeviceRequest) SetAuthenticationCode2(v string) *BindMFADeviceRequest {
	s.AuthenticationCode2 = &v
	return s
}

func (s *BindMFADeviceRequest) SetSerialNumber(v string) *BindMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindMFADeviceRequest) SetUserPrincipalName(v string) *BindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

type BindMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponseBody) SetRequestId(v string) *BindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type BindMFADeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponse) SetHeaders(v map[string]*string) *BindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *BindMFADeviceResponse) SetStatusCode(v int32) *BindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindMFADeviceResponse) SetBody(v *BindMFADeviceResponseBody) *BindMFADeviceResponse {
	s.Body = v
	return s
}

type ChangePasswordRequest struct {
	// The new password that is used to log on to the console.
	//
	// The password must meet the complexity requirements. For more information, see [GetPasswordPolicy](https://help.aliyun.com/document_detail/186691.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// newpassword
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The old password that is used to log on to the console.
	//
	// This parameter is required.
	//
	// example:
	//
	// mypassword
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

type ChangePasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

type ChangePasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) SetHeaders(v map[string]*string) *ChangePasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangePasswordResponse) SetStatusCode(v int32) *ChangePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangePasswordResponse) SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse {
	s.Body = v
	return s
}

type CreateAccessKeyRequest struct {
	// The logon name of the RAM user.
	//
	// If this parameter is empty, an AccessKey pair is created for the current user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyRequest) SetUserPrincipalName(v string) *CreateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

type CreateAccessKeyResponseBody struct {
	// The information about the AccessKey pair.
	AccessKey *CreateAccessKeyResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 19DDD9F7-AFCC-4D72-8CBA-CCE5A142E7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBody) SetAccessKey(v *CreateAccessKeyResponseBodyAccessKey) *CreateAccessKeyResponseBody {
	s.AccessKey = v
	return s
}

func (s *CreateAccessKeyResponseBody) SetRequestId(v string) *CreateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccessKeyResponseBodyAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// yourAccessKeySecret
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-15T08:08:54Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAccessKeyResponseBodyAccessKey) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponseBodyAccessKey) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeyId(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeySecret(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetCreateDate(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.CreateDate = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetStatus(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.Status = &v
	return s
}

type CreateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponse) SetHeaders(v map[string]*string) *CreateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateAccessKeyResponse) SetStatusCode(v int32) *CreateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAccessKeyResponse) SetBody(v *CreateAccessKeyResponseBody) *CreateAccessKeyResponse {
	s.Body = v
	return s
}

type CreateAppSecretRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSecretRequest) SetAppId(v string) *CreateAppSecretRequest {
	s.AppId = &v
	return s
}

type CreateAppSecretResponseBody struct {
	// The details of the application secret.
	AppSecret *CreateAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EE46FC3C-3BDE-4771-B531-27B7B6EB533D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponseBody) SetAppSecret(v *CreateAppSecretResponseBodyAppSecret) *CreateAppSecretResponseBody {
	s.AppSecret = v
	return s
}

func (s *CreateAppSecretResponseBody) SetRequestId(v string) *CreateAppSecretResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSecretResponseBodyAppSecret struct {
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	// The content of the application secret. This value can be used as the client secret for open authorization (OAuth).
	//
	// example:
	//
	// ai78ZmmxnlUG1jXlBZRDFKos9DIjY4m17Q7dCpMwn1rqXsTGb1X1XmrmveMp****
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-26T02:52:31Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s CreateAppSecretResponseBodyAppSecret) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretResponseBodyAppSecret) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppId(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppId = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppSecretId(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppSecretValue(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetCreateDate(v string) *CreateAppSecretResponseBodyAppSecret {
	s.CreateDate = &v
	return s
}

type CreateAppSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponse) SetHeaders(v map[string]*string) *CreateAppSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSecretResponse) SetStatusCode(v int32) *CreateAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSecretResponse) SetBody(v *CreateAppSecretResponseBody) *CreateAppSecretResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	// The validity period of the access token.
	//
	// Valid values: 900 to 10800. Unit: seconds.
	//
	// Default value: 3600.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The application name.
	//
	// The name can be up to 64 characters in length. The name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- WebApp: a web application that interacts with a browser.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop operating system or a mobile operating system.
	//
	// 	- ServerApp: an application that accesses Alibaba Cloud services without the need of manual user logon. User provisioning is automated based on the System for Cross-Domain Identity Management (SCIM) protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The display name of the application.
	//
	// The name can be up to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts. Valid values:
	//
	// 	- true: If you do not set this parameter for applications of the NativeApp and ServerApp types, true is used.
	//
	// 	- false: If you do not set this parameter for applications of the WebApp type, false is used.
	//
	// example:
	//
	// false
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The scope of application permissions.
	//
	// For more information about the application permission scope, see [Open authorization scope](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to query the permissions that are supported by different types of applications.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// example:
	//
	// aliuid
	PredefinedScopes *string `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty"`
	// The callback URL.
	//
	// If you enter multiple callback URLs, separate them with semicolons (;).
	//
	// example:
	//
	// https://www.example.com
	RedirectUris *string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty"`
	// The validity period of the refreshed token.
	//
	// Valid values: 7200 to 31536000. Unit: seconds.
	//
	// Default value:
	//
	// 	- For applications of the WebApp and ServerApp types, if this parameter is left empty, the value 2592000 is used. The value 2592000 indicates that the validity period of the refreshed token is 30 days.
	//
	// 	- For applications of the NativeApp type, if this parameter is left empty, the value 7776000 is used. The value 7776000 indicates that the validity period of the refreshed token is 90 days.
	//
	// example:
	//
	// 2592000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// The required permission.
	//
	// You can specify one or more permissions for the `RequiredScopes` parameter. After you specify this parameter, the required permissions are automatically selected and cannot be revoked when a user grants permissions on the application.
	//
	// If you enter multiple permission scopes, separate them with semicolons (;).
	//
	// >  If the permission that you specify for the `RequiredScopes` parameter is not included in the value of the `PredefinedScopes` parameter, the permission does not take effect.
	//
	// example:
	//
	// aliuid;profile
	RequiredScopes *string `json:"RequiredScopes,omitempty" xml:"RequiredScopes,omitempty"`
	// Indicates whether a secret is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >- For applications of the WebApp and ServerApp types, this parameter is automatically set to true and cannot be changed.
	//
	// >- For applications of the NativeApp type, this parameter can be set to true or false. If you do not set this parameter, false is used. Applications of the NativeApp type run in untrusted environments and the secrets of these applications are not protected. Therefore, we recommend that you do not set this parameter to true unless otherwise specified. For more information, see [Use an application of the NativeApp type to log on to Alibaba Cloud](https://help.aliyun.com/document_detail/93697.html).
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAccessTokenValidity(v int32) *CreateApplicationRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

func (s *CreateApplicationRequest) SetAppType(v string) *CreateApplicationRequest {
	s.AppType = &v
	return s
}

func (s *CreateApplicationRequest) SetDisplayName(v string) *CreateApplicationRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationRequest) SetIsMultiTenant(v bool) *CreateApplicationRequest {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationRequest) SetPredefinedScopes(v string) *CreateApplicationRequest {
	s.PredefinedScopes = &v
	return s
}

func (s *CreateApplicationRequest) SetRedirectUris(v string) *CreateApplicationRequest {
	s.RedirectUris = &v
	return s
}

func (s *CreateApplicationRequest) SetRefreshTokenValidity(v int32) *CreateApplicationRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetRequiredScopes(v string) *CreateApplicationRequest {
	s.RequiredScopes = &v
	return s
}

func (s *CreateApplicationRequest) SetSecretRequired(v bool) *CreateApplicationRequest {
	s.SecretRequired = &v
	return s
}

type CreateApplicationResponseBody struct {
	// The information about the application.
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationResponseBodyApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *CreateApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The callback URLs.
	RedirectUris *CreateApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *CreateApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAccountId(v string) *CreateApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppId(v string) *CreateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppName(v string) *CreateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetAppType(v string) *CreateApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetCreateDate(v string) *CreateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDelegatedScope(v *CreateApplicationResponseBodyApplicationDelegatedScope) *CreateApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDisplayName(v string) *CreateApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *CreateApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetRedirectUris(v *CreateApplicationResponseBodyApplicationRedirectUris) *CreateApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *CreateApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetSecretRequired(v bool) *CreateApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateDate(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

type CreateApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *CreateApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *CreateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

type CreateApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s CreateApplicationResponseBodyApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *CreateApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// The description.
	//
	// The description can be up to 128 characters in length.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user group.
	//
	// The name can be up to 24 characters in length.
	//
	// example:
	//
	// Dev-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the RAM user group. You must specify this parameter.
	//
	// The name can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetComments(v string) *CreateGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateGroupRequest) SetDisplayName(v string) *CreateGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

type CreateGroupResponseBody struct {
	// The information about the RAM user group.
	Group *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3C38192B-7BF8-45DA-8F0A-E670EA51426C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody {
	s.Group = v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyGroup) SetComments(v string) *CreateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetCreateDate(v string) *CreateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetDisplayName(v string) *CreateGroupResponseBodyGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupId(v string) *CreateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupName(v string) *CreateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetUpdateDate(v string) *CreateGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateLoginProfileRequest struct {
	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	//
	// 	- true: forcefully enables MFA for the RAM user. The RAM user must bind an MFA device upon the next logon.
	//
	// 	- false (default): does not forcefully enable MFA for the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The password that the RAM user uses to log on to the console.
	//
	// The password must meet the complexity requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user is required to reset the password upon the next logon. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Specifies whether to enable password-based logons to the console. Valid values:
	//
	// 	- Active: Password-based logon is enabled. This is the default value.
	//
	// 	- Inactive: Password-based logon is disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileRequest) SetMFABindRequired(v bool) *CreateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPassword(v string) *CreateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPasswordResetRequired(v bool) *CreateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetStatus(v string) *CreateLoginProfileRequest {
	s.Status = &v
	return s
}

func (s *CreateLoginProfileRequest) SetUserPrincipalName(v string) *CreateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type CreateLoginProfileResponseBody struct {
	// The logon information.
	LoginProfile *CreateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 29CB303C-1F05-43A6-A6BC-EBC5A797F8DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBody) SetLoginProfile(v *CreateLoginProfileResponseBodyLoginProfile) *CreateLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *CreateLoginProfileResponseBody) SetRequestId(v string) *CreateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type CreateLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether to forcefully enable MFA for the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Indicates whether to enable password-based logons to the console.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-14T03:47:51Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateLoginProfileResponseBodyLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *CreateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetStatus(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *CreateLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *CreateLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

type CreateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponse) SetHeaders(v map[string]*string) *CreateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *CreateLoginProfileResponse) SetStatusCode(v int32) *CreateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoginProfileResponse) SetBody(v *CreateLoginProfileResponseBody) *CreateLoginProfileResponse {
	s.Body = v
	return s
}

type CreateOIDCProviderRequest struct {
	// The ID of the client, which is provided by the external IdP. If you want to specify multiple client IDs, separate the client IDs with commas (,).
	//
	// The client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// The client ID can be up to 128 characters in length.
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The description of the OIDC IdP.
	//
	// The description can be up to 256 characters in length.
	//
	// example:
	//
	// This is an OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS CA certificate, which is provided by the external IdP. If you want to specify multiple fingerprints, separate the fingerprints with commas (,).
	//
	// The fingerprint can contain letters and digits.
	//
	// The fingerprint can be up to 128 characters in length.
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer, which is provided by the external IdP. The URL of the issuer must be unique within an Alibaba Cloud account.
	//
	// The URL of the issuer must start with `https` and be in the valid URL format. The URL cannot contain query parameters that follow a question mark (`?`) or logon information that is identified by at signs (`@`). The URL cannot be a fragment URL that contains number signs (`#`).
	//
	// The URL can be up to 255 characters in length.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// The name can contain letters, digits, and special characters and cannot start or end with the special characters. The special characters are `periods, (.), hyphens (-), and underscores (_)`.``
	//
	// The name can be up to 128 characters in length.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s CreateOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderRequest) SetClientIds(v string) *CreateOIDCProviderRequest {
	s.ClientIds = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetDescription(v string) *CreateOIDCProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetFingerprints(v string) *CreateOIDCProviderRequest {
	s.Fingerprints = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetIssuanceLimitTime(v int64) *CreateOIDCProviderRequest {
	s.IssuanceLimitTime = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetIssuerUrl(v string) *CreateOIDCProviderRequest {
	s.IssuerUrl = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetOIDCProviderName(v string) *CreateOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type CreateOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *CreateOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 64B11B41-636D-51E3-A39B-C8703CD2218C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponseBody) SetOIDCProvider(v *CreateOIDCProviderResponseBodyOIDCProvider) *CreateOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *CreateOIDCProviderResponseBody) SetRequestId(v string) *CreateOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type CreateOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client.
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is an OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate.
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636613763000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s CreateOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type CreateOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponse) SetHeaders(v map[string]*string) *CreateOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateOIDCProviderResponse) SetStatusCode(v int32) *CreateOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOIDCProviderResponse) SetBody(v *CreateOIDCProviderResponseBody) *CreateOIDCProviderResponse {
	s.Body = v
	return s
}

type CreateSAMLProviderRequest struct {
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata file which is Base64-encoded.
	//
	// The file is provided by an IdP that supports Security Assertion Markup Language (SAML) 2.0.
	//
	// example:
	//
	// PD94bWwgdmVy****
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty"`
	// The name of the IdP.
	//
	// The name can be up to 128 characters in length. The name can contain letters, digits, `periods (.), hyphens (-), and underscores (_)`. The name cannot start or end with `periods (.), hyphens (-), or underscores (_)`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s CreateSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderRequest) SetDescription(v string) *CreateSAMLProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetEncodedSAMLMetadataDocument(v string) *CreateSAMLProviderRequest {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetSAMLProviderName(v string) *CreateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type CreateSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A95A763D-F6B7-5242-83EB-AB45CE67F358
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the IdP.
	SAMLProvider *CreateSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s CreateSAMLProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponseBody) SetRequestId(v string) *CreateSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSAMLProviderResponseBody) SetSAMLProvider(v *CreateSAMLProviderResponseBodySAMLProvider) *CreateSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

type CreateSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The creation time. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The update time. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:51:20Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateSAMLProviderResponseBodySAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetArn(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *CreateSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *CreateSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

type CreateSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponse) SetHeaders(v map[string]*string) *CreateSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateSAMLProviderResponse) SetStatusCode(v int32) *CreateSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSAMLProviderResponse) SetBody(v *CreateSAMLProviderResponseBody) *CreateSAMLProviderResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// The description.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user.
	//
	// The name must be 1 to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile number of the RAM user.
	//
	// Format: Country code-Mobile phone number.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateUserRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// The name is in the format of `<username>@<AccountAlias>.onaliyun.com`. `<username>` indicates the name of the RAM user. `<AccountAlias>.onaliyun.com` indicates the default domain name. For more information about how to query the default domain name, see [GetDefaultDomain](https://help.aliyun.com/document_detail/186720.html).
	//
	// The value of `UserPrincipalName` must be `1 to 128` characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The value of `<username>` must be `1 to 64` characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetComments(v string) *CreateUserRequest {
	s.Comments = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetMobilePhone(v string) *CreateUserRequest {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserRequest) SetTag(v []*CreateUserRequestTag) *CreateUserRequest {
	s.Tag = v
	return s
}

func (s *CreateUserRequest) SetUserPrincipalName(v string) *CreateUserRequest {
	s.UserPrincipalName = &v
	return s
}

type CreateUserRequestTag struct {
	// The key of the tag.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. The tag value cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestTag) GoString() string {
	return s.String()
}

func (s *CreateUserRequestTag) SetKey(v string) *CreateUserRequestTag {
	s.Key = &v
	return s
}

func (s *CreateUserRequestTag) SetValue(v string) *CreateUserRequestTag {
	s.Value = &v
	return s
}

type CreateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2BB8C44A-2862-4922-AD43-03924749173B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *CreateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUser(v *CreateUserResponseBodyUser) *CreateUserResponseBody {
	s.User = v
	return s
}

type CreateUserResponseBodyUser struct {
	// The description.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The last time when the RAM user logged on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The source of the RAM user. Valid values:
	//
	// 	- Manual: The RAM user is manually created in the RAM console.
	//
	// 	- SCIM: The RAM user is mapped by using System for Cross-domain Identity Management (SCIM).
	//
	// 	- CloudSSO: The RAM user is mapped from a CloudSSO user.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The tag value.
	Tags *CreateUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the information about the RAM user was updated.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUser) SetComments(v string) *CreateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetCreateDate(v string) *CreateUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetDisplayName(v string) *CreateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetEmail(v string) *CreateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetLastLoginDate(v string) *CreateUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetMobilePhone(v string) *CreateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetProvisionType(v string) *CreateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetTags(v *CreateUserResponseBodyUserTags) *CreateUserResponseBodyUser {
	s.Tags = v
	return s
}

func (s *CreateUserResponseBodyUser) SetUpdateDate(v string) *CreateUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserId(v string) *CreateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserPrincipalName(v string) *CreateUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

type CreateUserResponseBodyUserTags struct {
	Tag []*CreateUserResponseBodyUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateUserResponseBodyUserTags) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUserTags) SetTag(v []*CreateUserResponseBodyUserTagsTag) *CreateUserResponseBodyUserTags {
	s.Tag = v
	return s
}

type CreateUserResponseBodyUserTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// operator
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateUserResponseBodyUserTagsTag) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBodyUserTagsTag) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUserTagsTag) SetTagKey(v string) *CreateUserResponseBodyUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *CreateUserResponseBodyUserTagsTag) SetTagValue(v string) *CreateUserResponseBodyUserTagsTag {
	s.TagValue = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateVirtualMFADeviceRequest struct {
	// The name of the MFA device.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// device001
	VirtualMFADeviceName *string `json:"VirtualMFADeviceName,omitempty" xml:"VirtualMFADeviceName,omitempty"`
}

func (s CreateVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceRequest) SetVirtualMFADeviceName(v string) *CreateVirtualMFADeviceRequest {
	s.VirtualMFADeviceName = &v
	return s
}

type CreateVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C609CC97-10FE-43EB-BE32-BDC219D8F1E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA device.
	VirtualMFADevice *CreateVirtualMFADeviceResponseBodyVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Struct"`
}

func (s CreateVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBody) SetRequestId(v string) *CreateVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBody) SetVirtualMFADevice(v *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) *CreateVirtualMFADeviceResponseBody {
	s.VirtualMFADevice = v
	return s
}

type CreateVirtualMFADeviceResponseBodyVirtualMFADevice struct {
	// The key of the MFA device.
	//
	// example:
	//
	// LD3CJ23Z2VGEX6R7ZTQCOA4XL2KODS5PKH7442NLKYX2PVHSHYB7UT3TS5HA****
	Base32StringSeed *string `json:"Base32StringSeed,omitempty" xml:"Base32StringSeed,omitempty"`
	// The Base64-encoded QR code of the key.
	//
	// example:
	//
	// YXNkZmFzZDlmeW5hc2Q5OGZoODd4bXJmcThhaGU5aSBmYXNkZiBzYWRmIGFGIDRxd2VjIGEgdHEz****
	QRCodePNG *string `json:"QRCodePNG,omitempty" xml:"QRCodePNG,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseBodyVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetBase32StringSeed(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.Base32StringSeed = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetQRCodePNG(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.QRCodePNG = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseBodyVirtualMFADevice) SetSerialNumber(v string) *CreateVirtualMFADeviceResponseBodyVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

type CreateVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *CreateVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetStatusCode(v int32) *CreateVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetBody(v *CreateVirtualMFADeviceResponseBody) *CreateVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DeleteAccessKeyRequest struct {
	// The ID of the AccessKey pair that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is empty, the AccessKey pair of the current user is deleted.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyRequest) SetUserPrincipalName(v string) *DeleteAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

type DeleteAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponseBody) SetRequestId(v string) *DeleteAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponse) SetHeaders(v map[string]*string) *DeleteAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessKeyResponse) SetStatusCode(v int32) *DeleteAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessKeyResponse) SetBody(v *DeleteAccessKeyResponseBody) *DeleteAccessKeyResponse {
	s.Body = v
	return s
}

type DeleteAccessKeyInRecycleBinRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the RAM user.
	//
	// > - If you use an Alibaba Cloud account to call the operation, you must specify the parameter.
	//
	// > - If you use a RAM user to call the operation, you can leave the parameter empty. In this case, the ID of the RAM user is used by default.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinRequest) SetUserId(v string) *DeleteAccessKeyInRecycleBinRequest {
	s.UserId = &v
	return s
}

type DeleteAccessKeyInRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinResponseBody) SetRequestId(v string) *DeleteAccessKeyInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccessKeyInRecycleBinResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessKeyInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetHeaders(v map[string]*string) *DeleteAccessKeyInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetStatusCode(v int32) *DeleteAccessKeyInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponse) SetBody(v *DeleteAccessKeyInRecycleBinResponseBody) *DeleteAccessKeyInRecycleBinResponse {
	s.Body = v
	return s
}

type DeleteAppSecretRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
}

func (s DeleteAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretRequest) SetAppId(v string) *DeleteAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppSecretRequest) SetAppSecretId(v string) *DeleteAppSecretRequest {
	s.AppSecretId = &v
	return s
}

type DeleteAppSecretResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretResponseBody) SetRequestId(v string) *DeleteAppSecretResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretResponse) SetHeaders(v map[string]*string) *DeleteAppSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppSecretResponse) SetStatusCode(v int32) *DeleteAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppSecretResponse) SetBody(v *DeleteAppSecretResponseBody) *DeleteAppSecretResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

type DeleteApplicationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetGroupName(v string) *DeleteGroupRequest {
	s.GroupName = &v
	return s
}

type DeleteGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteLoginProfileRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileRequest) SetUserPrincipalName(v string) *DeleteLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type DeleteLoginProfileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponseBody) SetRequestId(v string) *DeleteLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponse) SetHeaders(v map[string]*string) *DeleteLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoginProfileResponse) SetStatusCode(v int32) *DeleteLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoginProfileResponse) SetBody(v *DeleteLoginProfileResponseBody) *DeleteLoginProfileResponse {
	s.Body = v
	return s
}

type DeleteOIDCProviderRequest struct {
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s DeleteOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderRequest) SetOIDCProviderName(v string) *DeleteOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type DeleteOIDCProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderResponseBody) SetRequestId(v string) *DeleteOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteOIDCProviderResponse) SetHeaders(v map[string]*string) *DeleteOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteOIDCProviderResponse) SetStatusCode(v int32) *DeleteOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOIDCProviderResponse) SetBody(v *DeleteOIDCProviderResponseBody) *DeleteOIDCProviderResponse {
	s.Body = v
	return s
}

type DeletePasskeyRequest struct {
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeletePasskeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePasskeyRequest) GoString() string {
	return s.String()
}

func (s *DeletePasskeyRequest) SetPasskeyId(v string) *DeletePasskeyRequest {
	s.PasskeyId = &v
	return s
}

func (s *DeletePasskeyRequest) SetUserPrincipalName(v string) *DeletePasskeyRequest {
	s.UserPrincipalName = &v
	return s
}

type DeletePasskeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8A1673AA-5DB3-5AFB-8758-AF9EC2889259
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePasskeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePasskeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePasskeyResponseBody) SetRequestId(v string) *DeletePasskeyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePasskeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePasskeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePasskeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePasskeyResponse) GoString() string {
	return s.String()
}

func (s *DeletePasskeyResponse) SetHeaders(v map[string]*string) *DeletePasskeyResponse {
	s.Headers = v
	return s
}

func (s *DeletePasskeyResponse) SetStatusCode(v int32) *DeletePasskeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePasskeyResponse) SetBody(v *DeletePasskeyResponseBody) *DeletePasskeyResponse {
	s.Body = v
	return s
}

type DeleteSAMLProviderRequest struct {
	// The name of the IdP that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s DeleteSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderRequest) SetSAMLProviderName(v string) *DeleteSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type DeleteSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSAMLProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderResponseBody) SetRequestId(v string) *DeleteSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderResponse) SetHeaders(v map[string]*string) *DeleteSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteSAMLProviderResponse) SetStatusCode(v int32) *DeleteSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSAMLProviderResponse) SetBody(v *DeleteSAMLProviderResponseBody) *DeleteSAMLProviderResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// The ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) SetUserPrincipalName(v string) *DeleteUserRequest {
	s.UserPrincipalName = &v
	return s
}

type DeleteUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserInRecycleBinRequest struct {
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinRequest) SetUserId(v string) *DeleteUserInRecycleBinRequest {
	s.UserId = &v
	return s
}

type DeleteUserInRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinResponseBody) SetRequestId(v string) *DeleteUserInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserInRecycleBinResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserInRecycleBinResponse) SetHeaders(v map[string]*string) *DeleteUserInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserInRecycleBinResponse) SetStatusCode(v int32) *DeleteUserInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserInRecycleBinResponse) SetBody(v *DeleteUserInRecycleBinResponseBody) *DeleteUserInRecycleBinResponse {
	s.Body = v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	// The serial number of the MFA device.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:ram::123456789012****:mfa/device002
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVirtualMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponseBody) SetRequestId(v string) *DeleteVirtualMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetHeaders(v map[string]*string) *DeleteVirtualMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetStatusCode(v int32) *DeleteVirtualMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualMFADeviceResponse) SetBody(v *DeleteVirtualMFADeviceResponseBody) *DeleteVirtualMFADeviceResponse {
	s.Body = v
	return s
}

type DeprovisionApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeprovisionApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationRequest) SetAppId(v string) *DeprovisionApplicationRequest {
	s.AppId = &v
	return s
}

type DeprovisionApplicationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprovisionApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationResponseBody) SetRequestId(v string) *DeprovisionApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeprovisionApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionApplicationResponse) SetHeaders(v map[string]*string) *DeprovisionApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionApplicationResponse) SetStatusCode(v int32) *DeprovisionApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionApplicationResponse) SetBody(v *DeprovisionApplicationResponseBody) *DeprovisionApplicationResponse {
	s.Body = v
	return s
}

type DeprovisionExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 417273362044613****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeprovisionExternalApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationRequest) SetAppId(v string) *DeprovisionExternalApplicationRequest {
	s.AppId = &v
	return s
}

type DeprovisionExternalApplicationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1C2-3AB045D1AE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeprovisionExternalApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationResponseBody) SetRequestId(v string) *DeprovisionExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeprovisionExternalApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeprovisionExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeprovisionExternalApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeprovisionExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeprovisionExternalApplicationResponse) SetHeaders(v map[string]*string) *DeprovisionExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeprovisionExternalApplicationResponse) SetStatusCode(v int32) *DeprovisionExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeprovisionExternalApplicationResponse) SetBody(v *DeprovisionExternalApplicationResponseBody) *DeprovisionExternalApplicationResponse {
	s.Body = v
	return s
}

type DisableVirtualMFARequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DisableVirtualMFARequest) String() string {
	return tea.Prettify(s)
}

func (s DisableVirtualMFARequest) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFARequest) SetUserPrincipalName(v string) *DisableVirtualMFARequest {
	s.UserPrincipalName = &v
	return s
}

type DisableVirtualMFAResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableVirtualMFAResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableVirtualMFAResponseBody) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFAResponseBody) SetRequestId(v string) *DisableVirtualMFAResponseBody {
	s.RequestId = &v
	return s
}

type DisableVirtualMFAResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVirtualMFAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVirtualMFAResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableVirtualMFAResponse) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFAResponse) SetHeaders(v map[string]*string) *DisableVirtualMFAResponse {
	s.Headers = v
	return s
}

func (s *DisableVirtualMFAResponse) SetStatusCode(v int32) *DisableVirtualMFAResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVirtualMFAResponse) SetBody(v *DisableVirtualMFAResponseBody) *DisableVirtualMFAResponse {
	s.Body = v
	return s
}

type GenerateCredentialReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BBCCA90A-A1F0-4B16-B355-692247197805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The generation status of the user credential report. Valid values:
	//
	// 	- STARTED: The system starts to generate the user 
	//
	// 	- INPROGRESS: The user credential report is being generated.
	//
	// 	- COMPLETED: The user credential report is generated.
	//
	// example:
	//
	// STARTED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GenerateCredentialReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCredentialReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCredentialReportResponseBody) SetRequestId(v string) *GenerateCredentialReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCredentialReportResponseBody) SetState(v string) *GenerateCredentialReportResponseBody {
	s.State = &v
	return s
}

type GenerateCredentialReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCredentialReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCredentialReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateCredentialReportResponse) SetHeaders(v map[string]*string) *GenerateCredentialReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateCredentialReportResponse) SetStatusCode(v int32) *GenerateCredentialReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCredentialReportResponse) SetBody(v *GenerateCredentialReportResponseBody) *GenerateCredentialReportResponse {
	s.Body = v
	return s
}

type GenerateGovernanceReportResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 492E130C-76D3-55D5-BE5C-C023E431369A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The report generation status. Valid values:
	//
	// 	- Started: The system starts to generate a governance report.
	//
	// 	- Progressing: The system is generating a governance report.
	//
	// 	- Completed: A governance report is generated.
	//
	// example:
	//
	// Started
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GenerateGovernanceReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateGovernanceReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateGovernanceReportResponseBody) SetRequestId(v string) *GenerateGovernanceReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateGovernanceReportResponseBody) SetState(v string) *GenerateGovernanceReportResponseBody {
	s.State = &v
	return s
}

type GenerateGovernanceReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateGovernanceReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateGovernanceReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateGovernanceReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateGovernanceReportResponse) SetHeaders(v map[string]*string) *GenerateGovernanceReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateGovernanceReportResponse) SetStatusCode(v int32) *GenerateGovernanceReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateGovernanceReportResponse) SetBody(v *GenerateGovernanceReportResponseBody) *GenerateGovernanceReportResponse {
	s.Body = v
	return s
}

type GetAccessKeyInfoInRecycleBinRequest struct {
	// The AccessKey ID of the Resource Access Management (RAM) user.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinRequest) SetUserAccessKeyId(v string) *GetAccessKeyInfoInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

type GetAccessKeyInfoInRecycleBinResponseBody struct {
	// The information about the AccessKey pair.
	AccessKey *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) SetAccessKey(v *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) *GetAccessKeyInfoInRecycleBinResponseBody {
	s.AccessKey = v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBody) SetRequestId(v string) *GetAccessKeyInfoInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyInfoInRecycleBinResponseBodyAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the AccessKey pair will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-12T10:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The time when the AccessKey pair was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-12T10:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	// Indicates whether the RAM user to which the AccessKey pair belongs is in the recycle bin. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	UserRecycled *bool `json:"UserRecycled,omitempty" xml:"UserRecycled,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetAccessKeyId(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetCreateDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.CreateDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetDeleteDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.DeleteDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetRecycleDate(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.RecycleDate = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserId(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserId = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserPrincipalName(v string) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserPrincipalName = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey) SetUserRecycled(v bool) *GetAccessKeyInfoInRecycleBinResponseBodyAccessKey {
	s.UserRecycled = &v
	return s
}

type GetAccessKeyInfoInRecycleBinResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyInfoInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyInfoInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyInfoInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetHeaders(v map[string]*string) *GetAccessKeyInfoInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetStatusCode(v int32) *GetAccessKeyInfoInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyInfoInRecycleBinResponse) SetBody(v *GetAccessKeyInfoInRecycleBinResponseBody) *GetAccessKeyInfoInRecycleBinResponse {
	s.Body = v
	return s
}

type GetAccessKeyLastUsedRequest struct {
	// The ID of the AccessKey pair that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If you do not specify this parameter, the AccessKey pair of the current user is queried.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetAccessKeyLastUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedRequest) SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) SetUserPrincipalName(v string) *GetAccessKeyLastUsedRequest {
	s.UserPrincipalName = &v
	return s
}

type GetAccessKeyLastUsedResponseBody struct {
	// The details of the time when the AccessKey pair was used for the last time.
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B29C79F6-354B-4297-A994-1338CC22A2EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBody) SetAccessKeyLastUsed(v *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) *GetAccessKeyLastUsedResponseBody {
	s.AccessKeyLastUsed = v
	return s
}

func (s *GetAccessKeyLastUsedResponseBody) SetRequestId(v string) *GetAccessKeyLastUsedResponseBody {
	s.RequestId = &v
	return s
}

type GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed struct {
	// The time when the AccessKey pair was used for the last time.
	//
	// example:
	//
	// 2020-10-16T01:37:37Z
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty"`
	// The Alibaba Cloud service that was last accessed by using the AccessKey pair.
	//
	// example:
	//
	// Ram
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetLastUsedDate(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.LastUsedDate = &v
	return s
}

func (s *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed) SetServiceName(v string) *GetAccessKeyLastUsedResponseBodyAccessKeyLastUsed {
	s.ServiceName = &v
	return s
}

type GetAccessKeyLastUsedResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessKeyLastUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessKeyLastUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponse) SetHeaders(v map[string]*string) *GetAccessKeyLastUsedResponse {
	s.Headers = v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetStatusCode(v int32) *GetAccessKeyLastUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetBody(v *GetAccessKeyLastUsedResponseBody) *GetAccessKeyLastUsedResponse {
	s.Body = v
	return s
}

type GetAccountMFAInfoResponseBody struct {
	// Indicates whether MFA devices are enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsMFAEnable *bool `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4BE83135-0B08-467C-B3A2-27B312FD0F57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountMFAInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountMFAInfoResponseBody) SetIsMFAEnable(v bool) *GetAccountMFAInfoResponseBody {
	s.IsMFAEnable = &v
	return s
}

func (s *GetAccountMFAInfoResponseBody) SetRequestId(v string) *GetAccountMFAInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountMFAInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountMFAInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountMFAInfoResponse) SetHeaders(v map[string]*string) *GetAccountMFAInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountMFAInfoResponse) SetStatusCode(v int32) *GetAccountMFAInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountMFAInfoResponse) SetBody(v *GetAccountMFAInfoResponseBody) *GetAccountMFAInfoResponse {
	s.Body = v
	return s
}

type GetAccountSecurityPracticeReportResponseBody struct {
	// The information about the security report for the Alibaba Cloud account.
	AccountSecurityPracticeInfo *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo `json:"AccountSecurityPracticeInfo,omitempty" xml:"AccountSecurityPracticeInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ABA822EE-85C2-4418-9577-A1831FC8466D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBody) SetAccountSecurityPracticeInfo(v *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) *GetAccountSecurityPracticeReportResponseBody {
	s.AccountSecurityPracticeInfo = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBody) SetRequestId(v string) *GetAccountSecurityPracticeReportResponseBody {
	s.RequestId = &v
	return s
}

type GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo struct {
	// The information about the security report for the Alibaba Cloud account.
	AccountSecurityPracticeUserInfo *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo `json:"AccountSecurityPracticeUserInfo,omitempty" xml:"AccountSecurityPracticeUserInfo,omitempty" type:"Struct"`
	// The security score of the Alibaba Cloud account.
	//
	// example:
	//
	// 63
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) SetAccountSecurityPracticeUserInfo(v *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo {
	s.AccountSecurityPracticeUserInfo = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo) SetScore(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfo {
	s.Score = &v
	return s
}

type GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo struct {
	// Indicates whether multi-factor authentication (MFA) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	BindMfa *bool `json:"BindMfa,omitempty" xml:"BindMfa,omitempty"`
	// The number of old AccessKey pairs for the Alibaba Cloud account.
	//
	// example:
	//
	// 0
	OldAkNum *int32 `json:"OldAkNum,omitempty" xml:"OldAkNum,omitempty"`
	// The number of AccessKey pairs for the Alibaba Cloud account.
	//
	// example:
	//
	// 1
	RootWithAccessKey *int32 `json:"RootWithAccessKey,omitempty" xml:"RootWithAccessKey,omitempty"`
	// The number of RAM users within the Alibaba Cloud account.
	//
	// example:
	//
	// 9
	SubUser *int32 `json:"SubUser,omitempty" xml:"SubUser,omitempty"`
	// The number of RAM users that have MFA devices bound.
	//
	// example:
	//
	// 0
	SubUserBindMfa *int32 `json:"SubUserBindMfa,omitempty" xml:"SubUserBindMfa,omitempty"`
	// The complexity level of the password for the RAM user. Valid values:
	//
	// 	- low
	//
	// 	- mid
	//
	// 	- high
	//
	// example:
	//
	// low
	SubUserPwdLevel *string `json:"SubUserPwdLevel,omitempty" xml:"SubUserPwdLevel,omitempty"`
	// The number of RAM users that use the old AccessKey pairs.
	//
	// example:
	//
	// 0
	SubUserWithOldAccessKey *int32 `json:"SubUserWithOldAccessKey,omitempty" xml:"SubUserWithOldAccessKey,omitempty"`
	// The number of Resource Access Management (RAM) users that have unused AccessKey pairs.
	//
	// example:
	//
	// 0
	SubUserWithUnusedAccessKey *int32 `json:"SubUserWithUnusedAccessKey,omitempty" xml:"SubUserWithUnusedAccessKey,omitempty"`
	// The number of AccessKey pairs that are not used for the Alibaba Cloud account.
	//
	// example:
	//
	// 0
	UnusedAkNum *int32 `json:"UnusedAkNum,omitempty" xml:"UnusedAkNum,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetBindMfa(v bool) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.BindMfa = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetOldAkNum(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.OldAkNum = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetRootWithAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.RootWithAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUser(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUser = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserBindMfa(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserBindMfa = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserPwdLevel(v string) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserPwdLevel = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithOldAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithOldAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithUnusedAccessKey(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithUnusedAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetUnusedAkNum(v int32) *GetAccountSecurityPracticeReportResponseBodyAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.UnusedAkNum = &v
	return s
}

type GetAccountSecurityPracticeReportResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountSecurityPracticeReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountSecurityPracticeReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponse) SetHeaders(v map[string]*string) *GetAccountSecurityPracticeReportResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) SetStatusCode(v int32) *GetAccountSecurityPracticeReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) SetBody(v *GetAccountSecurityPracticeReportResponseBody) *GetAccountSecurityPracticeReportResponse {
	s.Body = v
	return s
}

type GetAccountSummaryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 81313F5E-3C85-478F-BCC9-E1B70E4556DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The overview information about the Alibaba Cloud account.
	SummaryMap *GetAccountSummaryResponseBodySummaryMap `json:"SummaryMap,omitempty" xml:"SummaryMap,omitempty" type:"Struct"`
}

func (s GetAccountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBody) SetRequestId(v string) *GetAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountSummaryResponseBody) SetSummaryMap(v *GetAccountSummaryResponseBodySummaryMap) *GetAccountSummaryResponseBody {
	s.SummaryMap = v
	return s
}

type GetAccountSummaryResponseBodySummaryMap struct {
	// The maximum number of AccessKey pairs that a Resource Access Management (RAM) user can have.
	//
	// example:
	//
	// 2
	AccessKeysPerUserQuota *int32 `json:"AccessKeysPerUserQuota,omitempty" xml:"AccessKeysPerUserQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM user group.
	//
	// example:
	//
	// 5
	AttachedPoliciesPerGroupQuota *int32 `json:"AttachedPoliciesPerGroupQuota,omitempty" xml:"AttachedPoliciesPerGroupQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM role.
	//
	// example:
	//
	// 5
	AttachedPoliciesPerRoleQuota *int32 `json:"AttachedPoliciesPerRoleQuota,omitempty" xml:"AttachedPoliciesPerRoleQuota,omitempty"`
	// The maximum number of custom policies that can be added to a RAM user.
	//
	// example:
	//
	// 10
	AttachedPoliciesPerUserQuota *int32 `json:"AttachedPoliciesPerUserQuota,omitempty" xml:"AttachedPoliciesPerUserQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM user group.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerGroupQuota *int32 `json:"AttachedSystemPoliciesPerGroupQuota,omitempty" xml:"AttachedSystemPoliciesPerGroupQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM role.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerRoleQuota *int32 `json:"AttachedSystemPoliciesPerRoleQuota,omitempty" xml:"AttachedSystemPoliciesPerRoleQuota,omitempty"`
	// The maximum number of system policies that can be added to a RAM user.
	//
	// example:
	//
	// 20
	AttachedSystemPoliciesPerUserQuota *int32 `json:"AttachedSystemPoliciesPerUserQuota,omitempty" xml:"AttachedSystemPoliciesPerUserQuota,omitempty"`
	// The maximum number of network access control policies that can be configured for an Alibaba Cloud account or AccessKey pair.
	//
	// example:
	//
	// 8
	ConditionsPerAKPolicyQuota *int32 `json:"ConditionsPerAKPolicyQuota,omitempty" xml:"ConditionsPerAKPolicyQuota,omitempty"`
	// The number of RAM user groups.
	//
	// example:
	//
	// 7
	Groups *int32 `json:"Groups,omitempty" xml:"Groups,omitempty"`
	// The maximum number of RAM user groups to which a RAM user can be added.
	//
	// example:
	//
	// 5
	GroupsPerUserQuota *int32 `json:"GroupsPerUserQuota,omitempty" xml:"GroupsPerUserQuota,omitempty"`
	// The maximum number of RAM user groups that can be created.
	//
	// example:
	//
	// 50
	GroupsQuota *int32 `json:"GroupsQuota,omitempty" xml:"GroupsQuota,omitempty"`
	// The maximum number of IP addresses that can be specified in an account-level AccessKey pair-based or AccessKey pair-level policy for network access control.
	//
	// example:
	//
	// 50
	IPItemsPerAKPolicyQuota *int32 `json:"IPItemsPerAKPolicyQuota,omitempty" xml:"IPItemsPerAKPolicyQuota,omitempty"`
	// The number of virtual multi-factor authentication (MFA) devices.
	//
	// example:
	//
	// 13
	MFADevices *int32 `json:"MFADevices,omitempty" xml:"MFADevices,omitempty"`
	// The number of virtual MFA devices in use.
	//
	// example:
	//
	// 2
	MFADevicesInUse *int32 `json:"MFADevicesInUse,omitempty" xml:"MFADevicesInUse,omitempty"`
	// The number of custom policies.
	//
	// example:
	//
	// 13
	Policies *int32 `json:"Policies,omitempty" xml:"Policies,omitempty"`
	// The maximum number of custom policies that can be created.
	//
	// example:
	//
	// 1500
	PoliciesQuota *int32 `json:"PoliciesQuota,omitempty" xml:"PoliciesQuota,omitempty"`
	// The maximum length of the policy content.
	//
	// example:
	//
	// 2048
	PolicySizeQuota *int32 `json:"PolicySizeQuota,omitempty" xml:"PolicySizeQuota,omitempty"`
	// The number of RAM roles.
	//
	// example:
	//
	// 19
	Roles *int32 `json:"Roles,omitempty" xml:"Roles,omitempty"`
	// The maximum number of RAM roles that can be created.
	//
	// example:
	//
	// 1000
	RolesQuota *int32 `json:"RolesQuota,omitempty" xml:"RolesQuota,omitempty"`
	// The number of RAM users.
	//
	// example:
	//
	// 9
	Users *int32 `json:"Users,omitempty" xml:"Users,omitempty"`
	// The maximum number of RAM users that can be created.
	//
	// example:
	//
	// 1000
	UsersQuota *int32 `json:"UsersQuota,omitempty" xml:"UsersQuota,omitempty"`
	// The maximum number of policy versions.
	//
	// example:
	//
	// 5
	VersionsPerPolicyQuota *int32 `json:"VersionsPerPolicyQuota,omitempty" xml:"VersionsPerPolicyQuota,omitempty"`
	// The maximum number of virtual MFA devices that can be created.
	//
	// example:
	//
	// 1000
	VirtualMFADevicesQuota *int32 `json:"VirtualMFADevicesQuota,omitempty" xml:"VirtualMFADevicesQuota,omitempty"`
}

func (s GetAccountSummaryResponseBodySummaryMap) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponseBodySummaryMap) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAccessKeysPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AccessKeysPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerGroupQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerRoleQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetAttachedSystemPoliciesPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.AttachedSystemPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetConditionsPerAKPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.ConditionsPerAKPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroups(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Groups = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroupsPerUserQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.GroupsPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetGroupsQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.GroupsQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetIPItemsPerAKPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.IPItemsPerAKPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetMFADevices(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.MFADevices = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetMFADevicesInUse(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.MFADevicesInUse = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicies(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Policies = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPoliciesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PoliciesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetPolicySizeQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.PolicySizeQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRoles(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Roles = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetRolesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.RolesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetUsers(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.Users = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetUsersQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.UsersQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetVersionsPerPolicyQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.VersionsPerPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseBodySummaryMap) SetVirtualMFADevicesQuota(v int32) *GetAccountSummaryResponseBodySummaryMap {
	s.VirtualMFADevicesQuota = &v
	return s
}

type GetAccountSummaryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponse) SetHeaders(v map[string]*string) *GetAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetAccountSummaryResponse) SetStatusCode(v int32) *GetAccountSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountSummaryResponse) SetBody(v *GetAccountSummaryResponseBody) *GetAccountSummaryResponse {
	s.Body = v
	return s
}

type GetAppSecretRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
}

func (s GetAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *GetAppSecretRequest) SetAppId(v string) *GetAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *GetAppSecretRequest) SetAppSecretId(v string) *GetAppSecretRequest {
	s.AppSecretId = &v
	return s
}

type GetAppSecretResponseBody struct {
	// The details of the application secret.
	AppSecret *GetAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EE46FC3C-3BDE-4771-B531-27B7B6EB533D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBody) SetAppSecret(v *GetAppSecretResponseBodyAppSecret) *GetAppSecretResponseBody {
	s.AppSecret = v
	return s
}

func (s *GetAppSecretResponseBody) SetRequestId(v string) *GetAppSecretResponseBody {
	s.RequestId = &v
	return s
}

type GetAppSecretResponseBodyAppSecret struct {
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	// The content of the application secret.
	//
	// example:
	//
	// ai78ZmmxnlUG1jXlBZRDFKos9DIjY4m17Q7dCpMwn1rqXsTGb1X1XmrmveMp****
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-26T02:52:31Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s GetAppSecretResponseBodyAppSecret) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponseBodyAppSecret) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppId(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppId = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppSecretId(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppSecretValue(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetCreateDate(v string) *GetAppSecretResponseBodyAppSecret {
	s.CreateDate = &v
	return s
}

type GetAppSecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponse) SetHeaders(v map[string]*string) *GetAppSecretResponse {
	s.Headers = v
	return s
}

func (s *GetAppSecretResponse) SetStatusCode(v int32) *GetAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSecretResponse) SetBody(v *GetAppSecretResponseBody) *GetAppSecretResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetAppId(v string) *GetApplicationRequest {
	s.AppId = &v
	return s
}

type GetApplicationResponseBody struct {
	// The information about the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationResponseBodyApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- WebApp: a web application.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop or mobile operating system.
	//
	// 	- ServerApp: an application that can access Alibaba Cloud services without the need for user logon. Only applications that synchronize user information based on the System for Cross-domain Identity Management (SCIM) protocol are supported.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *GetApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The callback URL.
	RedirectUris *GetApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *GetApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAccountId(v string) *GetApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppId(v string) *GetApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppName(v string) *GetApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppType(v string) *GetApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateDate(v string) *GetApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDelegatedScope(v *GetApplicationResponseBodyApplicationDelegatedScope) *GetApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDisplayName(v string) *GetApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *GetApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRedirectUris(v *GetApplicationResponseBodyApplicationRedirectUris) *GetApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *GetApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSecretRequired(v bool) *GetApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateDate(v string) *GetApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

type GetApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetApplicationResponseBodyApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *GetApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *GetApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

type GetApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *GetApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetApplicationProvisionInfoRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationProvisionInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoRequest) SetAppId(v string) *GetApplicationProvisionInfoRequest {
	s.AppId = &v
	return s
}

type GetApplicationProvisionInfoResponseBody struct {
	// The installation information about the application.
	ApplicationProvisionInfo *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 936E1D9C-157D-45BD-8A3B-81C0716EB078
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisionInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBody) SetApplicationProvisionInfo(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) *GetApplicationProvisionInfoResponseBody {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBody) SetRequestId(v string) *GetApplicationProvisionInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// GiteePrd
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the scopes of permissions that are granted to the application.
	DelegatedScope *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAccountId(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAppId(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAppName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetCreateDate(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetDelegatedScope(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetDisplayName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetUpdateDate(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope struct {
	// The information about the scopes of permissions that are granted to the application.
	PredefinedScopes *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission scope.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission scope.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type GetApplicationProvisionInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisionInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisionInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisionInfoResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponse) SetHeaders(v map[string]*string) *GetApplicationProvisionInfoResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisionInfoResponse) SetStatusCode(v int32) *GetApplicationProvisionInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisionInfoResponse) SetBody(v *GetApplicationProvisionInfoResponseBody) *GetApplicationProvisionInfoResponse {
	s.Body = v
	return s
}

type GetCredentialReportRequest struct {
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 3501. Default value: 3501.
	//
	// example:
	//
	// 1000
	MaxItems *string `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The token that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.``
	//
	// example:
	//
	// EXAMPLE
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetCredentialReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialReportRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialReportRequest) SetMaxItems(v string) *GetCredentialReportRequest {
	s.MaxItems = &v
	return s
}

func (s *GetCredentialReportRequest) SetNextToken(v string) *GetCredentialReportRequest {
	s.NextToken = &v
	return s
}

type GetCredentialReportResponseBody struct {
	// The content of the user credential report.
	//
	// The report is Base64-encoded. After you decode the report, the credential report is in the CSV format.
	//
	// example:
	//
	// OVZWK4RMOVZW****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the user credential report was generated.
	//
	// example:
	//
	// 2020-10-19T15:06:52Z
	GeneratedTime *string `json:"GeneratedTime,omitempty" xml:"GeneratedTime,omitempty"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *string `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The parameter that is used to obtain the truncated part. This parameter takes effect only when `IsTruncated` is set to true.
	//
	// example:
	//
	// EXAMPLE
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A01826E-7601-44B0-B4DF-2B0C509836DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCredentialReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialReportResponseBody) SetContent(v string) *GetCredentialReportResponseBody {
	s.Content = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetGeneratedTime(v string) *GetCredentialReportResponseBody {
	s.GeneratedTime = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetIsTruncated(v string) *GetCredentialReportResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetNextToken(v string) *GetCredentialReportResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetRequestId(v string) *GetCredentialReportResponseBody {
	s.RequestId = &v
	return s
}

type GetCredentialReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCredentialReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCredentialReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialReportResponse) SetHeaders(v map[string]*string) *GetCredentialReportResponse {
	s.Headers = v
	return s
}

func (s *GetCredentialReportResponse) SetStatusCode(v int32) *GetCredentialReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCredentialReportResponse) SetBody(v *GetCredentialReportResponseBody) *GetCredentialReportResponse {
	s.Body = v
	return s
}

type GetDefaultDomainResponseBody struct {
	// The default domain name.
	//
	// example:
	//
	// examplecompany.onaliyun.com
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66815255-7CCE-4759-AC37-9755794C3626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDefaultDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultDomainResponseBody) SetDefaultDomainName(v string) *GetDefaultDomainResponseBody {
	s.DefaultDomainName = &v
	return s
}

func (s *GetDefaultDomainResponseBody) SetRequestId(v string) *GetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

type GetDefaultDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultDomainResponse) SetHeaders(v map[string]*string) *GetDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultDomainResponse) SetStatusCode(v int32) *GetDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultDomainResponse) SetBody(v *GetDefaultDomainResponseBody) *GetDefaultDomainResponse {
	s.Body = v
	return s
}

type GetExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetExternalApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationRequest) SetAppId(v string) *GetExternalApplicationRequest {
	s.AppId = &v
	return s
}

type GetExternalApplicationResponseBody struct {
	// The information about the external application.
	ExternalApplication *GetExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4C4D1BD-2558-5BD1-8C36-A5D7FB174A55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExternalApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBody) SetExternalApplication(v *GetExternalApplicationResponseBodyExternalApplication) *GetExternalApplicationResponseBody {
	s.ExternalApplication = v
	return s
}

func (s *GetExternalApplicationResponseBody) SetRequestId(v string) *GetExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetExternalApplicationResponseBodyExternalApplication struct {
	// The name of the application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.153082740420****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1737534146000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 407426893752729****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account for which the application is installed.
	//
	// example:
	//
	// 173082740420****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time of the application. The value is a timestamp.
	//
	// example:
	//
	// 1737534146000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetExternalApplicationResponseBodyExternalApplication) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplication) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetAppPrincipalName(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetCreateDate(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetDelegatedScope(v *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) *GetExternalApplicationResponseBodyExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetDisplayName(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetForeignAppId(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetTenantId(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.TenantId = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetUpdateDate(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.UpdateDate = &v
	return s
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) SetPredefinedScopes(v *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type GetExternalApplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponse) SetHeaders(v map[string]*string) *GetExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetExternalApplicationResponse) SetStatusCode(v int32) *GetExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalApplicationResponse) SetBody(v *GetExternalApplicationResponseBody) *GetExternalApplicationResponse {
	s.Body = v
	return s
}

type GetGovernanceItemReportRequest struct {
	// example:
	//
	// SSOLoginEnabled
	GovernanceItemType *string `json:"GovernanceItemType,omitempty" xml:"GovernanceItemType,omitempty"`
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// 1000
	MaxItems *string `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s GetGovernanceItemReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportRequest) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportRequest) SetGovernanceItemType(v string) *GetGovernanceItemReportRequest {
	s.GovernanceItemType = &v
	return s
}

func (s *GetGovernanceItemReportRequest) SetMarker(v string) *GetGovernanceItemReportRequest {
	s.Marker = &v
	return s
}

func (s *GetGovernanceItemReportRequest) SetMaxItems(v string) *GetGovernanceItemReportRequest {
	s.MaxItems = &v
	return s
}

type GetGovernanceItemReportResponseBody struct {
	ColumnsSchema *GetGovernanceItemReportResponseBodyColumnsSchema `json:"ColumnsSchema,omitempty" xml:"ColumnsSchema,omitempty" type:"Struct"`
	ColumnsValue  *GetGovernanceItemReportResponseBodyColumnsValue  `json:"ColumnsValue,omitempty" xml:"ColumnsValue,omitempty" type:"Struct"`
	// example:
	//
	// 2020-10-19T15:06:52Z
	GenerateTime *string `json:"GenerateTime,omitempty" xml:"GenerateTime,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// Number
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// 100
	MetricValue interface{} `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// example:
	//
	// F2CE9688-AA85-5F23-8C22-0EC23473405F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGovernanceItemReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBody) SetColumnsSchema(v *GetGovernanceItemReportResponseBodyColumnsSchema) *GetGovernanceItemReportResponseBody {
	s.ColumnsSchema = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetColumnsValue(v *GetGovernanceItemReportResponseBodyColumnsValue) *GetGovernanceItemReportResponseBody {
	s.ColumnsValue = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetGenerateTime(v string) *GetGovernanceItemReportResponseBody {
	s.GenerateTime = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetIsTruncated(v bool) *GetGovernanceItemReportResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMarker(v string) *GetGovernanceItemReportResponseBody {
	s.Marker = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMetricType(v string) *GetGovernanceItemReportResponseBody {
	s.MetricType = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMetricValue(v interface{}) *GetGovernanceItemReportResponseBody {
	s.MetricValue = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetRequestId(v string) *GetGovernanceItemReportResponseBody {
	s.RequestId = &v
	return s
}

type GetGovernanceItemReportResponseBodyColumnsSchema struct {
	ColumnSchema []*GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema `json:"ColumnSchema,omitempty" xml:"ColumnSchema,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsSchema) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsSchema) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchema) SetColumnSchema(v []*GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) *GetGovernanceItemReportResponseBodyColumnsSchema {
	s.ColumnSchema = v
	return s
}

type GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema struct {
	// example:
	//
	// UserPrincipalName
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// String
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
}

func (s GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) SetColumnName(v string) *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema {
	s.ColumnName = &v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) SetColumnType(v string) *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema {
	s.ColumnType = &v
	return s
}

type GetGovernanceItemReportResponseBodyColumnsValue struct {
	ColumnRow []*GetGovernanceItemReportResponseBodyColumnsValueColumnRow `json:"ColumnRow,omitempty" xml:"ColumnRow,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsValue) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsValue) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsValue) SetColumnRow(v []*GetGovernanceItemReportResponseBodyColumnsValueColumnRow) *GetGovernanceItemReportResponseBodyColumnsValue {
	s.ColumnRow = v
	return s
}

type GetGovernanceItemReportResponseBodyColumnsValueColumnRow struct {
	ColumnValue []interface{} `json:"ColumnValue,omitempty" xml:"ColumnValue,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsValueColumnRow) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsValueColumnRow) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsValueColumnRow) SetColumnValue(v []interface{}) *GetGovernanceItemReportResponseBodyColumnsValueColumnRow {
	s.ColumnValue = v
	return s
}

type GetGovernanceItemReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceItemReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceItemReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceItemReportResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponse) SetHeaders(v map[string]*string) *GetGovernanceItemReportResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceItemReportResponse) SetStatusCode(v int32) *GetGovernanceItemReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceItemReportResponse) SetBody(v *GetGovernanceItemReportResponseBody) *GetGovernanceItemReportResponse {
	s.Body = v
	return s
}

type GetGovernanceReportStatusResponseBody struct {
	GovernanceItemsStatus *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus `json:"GovernanceItemsStatus,omitempty" xml:"GovernanceItemsStatus,omitempty" type:"Struct"`
	// example:
	//
	// F2CE9688-AA85-5F23-8C22-0EC23473405A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Progressing
	WholeReportStatus *string `json:"WholeReportStatus,omitempty" xml:"WholeReportStatus,omitempty"`
}

func (s GetGovernanceReportStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBody) SetGovernanceItemsStatus(v *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) *GetGovernanceReportStatusResponseBody {
	s.GovernanceItemsStatus = v
	return s
}

func (s *GetGovernanceReportStatusResponseBody) SetRequestId(v string) *GetGovernanceReportStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBody) SetWholeReportStatus(v string) *GetGovernanceReportStatusResponseBody {
	s.WholeReportStatus = &v
	return s
}

type GetGovernanceReportStatusResponseBodyGovernanceItemsStatus struct {
	GovernanceItemStatus []*GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus `json:"GovernanceItemStatus,omitempty" xml:"GovernanceItemStatus,omitempty" type:"Repeated"`
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus) SetGovernanceItemStatus(v []*GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatus {
	s.GovernanceItemStatus = v
	return s
}

type GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus struct {
	// example:
	//
	// AccountRecentUsingAccessKey
	GovernanceItem *string `json:"GovernanceItem,omitempty" xml:"GovernanceItem,omitempty"`
	// example:
	//
	// Progressing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) SetGovernanceItem(v string) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus {
	s.GovernanceItem = &v
	return s
}

func (s *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus) SetStatus(v string) *GetGovernanceReportStatusResponseBodyGovernanceItemsStatusGovernanceItemStatus {
	s.Status = &v
	return s
}

type GetGovernanceReportStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceReportStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceReportStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGovernanceReportStatusResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceReportStatusResponse) SetHeaders(v map[string]*string) *GetGovernanceReportStatusResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceReportStatusResponse) SetStatusCode(v int32) *GetGovernanceReportStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceReportStatusResponse) SetBody(v *GetGovernanceReportStatusResponseBody) *GetGovernanceReportStatusResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetGroupName(v string) *GetGroupRequest {
	s.GroupName = &v
	return s
}

type GetGroupResponseBody struct {
	// The information about the RAM user group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 86ECEC3C-7262-4C3C-94B4-A98F7CC1F060
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) SetComments(v string) *GetGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetCreateDate(v string) *GetGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDisplayName(v string) *GetGroupResponseBodyGroup {
	s.DisplayName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupName(v string) *GetGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateDate(v string) *GetGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetLoginProfileRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *GetLoginProfileRequest) SetUserPrincipalName(v string) *GetLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type GetLoginProfileResponseBody struct {
	// The console logon configurations.
	LoginProfile *GetLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E517F18B-632C-48FC-93F1-CDCBCC6F8444
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBody) SetLoginProfile(v *GetLoginProfileResponseBodyLoginProfile) *GetLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *GetLoginProfileResponseBody) SetRequestId(v string) *GetLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type GetLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether console logon is automatically disabled if a RAM user does not log on to the console in the previous specified number of days. The number of days is specified by MaxIdleDaysForUsers. The default value is true, and you cannot change the value.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// The time of the most recent logon. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-14T07:25:25Z
	LastLoginTime *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty"`
	// Indicates whether multi-factor authentication (MFA) must be enabled. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Indicates whether console logon is enabled. Valid values:
	//
	// 	- Active: enabled.
	//
	// 	- Inactive: disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The modification time. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-14T06:56:45Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetLoginProfileResponseBodyLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetAutoDisableLoginStatus(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.AutoDisableLoginStatus = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetLastLoginTime(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.LastLoginTime = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *GetLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetStatus(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *GetLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *GetLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

type GetLoginProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponse) SetHeaders(v map[string]*string) *GetLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *GetLoginProfileResponse) SetStatusCode(v int32) *GetLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginProfileResponse) SetBody(v *GetLoginProfileResponseBody) *GetLoginProfileResponse {
	s.Body = v
	return s
}

type GetOIDCProviderRequest struct {
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s GetOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderRequest) SetOIDCProviderName(v string) *GetOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type GetOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *GetOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E5E1A300-279D-5FBD-A8CF-F4EDC20C4896
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponseBody) SetOIDCProvider(v *GetOIDCProviderResponseBodyOIDCProvider) *GetOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *GetOIDCProviderResponseBody) SetRequestId(v string) *GetOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type GetOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s GetOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *GetOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type GetOIDCProviderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponse) SetHeaders(v map[string]*string) *GetOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *GetOIDCProviderResponse) SetStatusCode(v int32) *GetOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOIDCProviderResponse) SetBody(v *GetOIDCProviderResponseBody) *GetOIDCProviderResponse {
	s.Body = v
	return s
}

type GetPasswordPolicyResponseBody struct {
	// The details of the password policy.
	PasswordPolicy *GetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BDAA8408-E67C-428B-BFF0-1B2AC05C9610
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBody) SetPasswordPolicy(v *GetPasswordPolicyResponseBodyPasswordPolicy) *GetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *GetPasswordPolicyResponseBody) SetRequestId(v string) *GetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether to disable logon after the password expires.
	//
	// example:
	//
	// false
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// The maximum number of password retries.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The validity period of the password.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in the password.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// The minimum required number of characters in a password.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Indicates whether to exclude the username from the password.
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// The policy for password history check.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether the password must contain lowercase letters.
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether the password must contain digits.
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether the password must contain special characters.
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether the password must contain uppercase letters.
	//
	// example:
	//
	// false
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordDifferentCharacter(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUserName(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

type GetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponse) SetHeaders(v map[string]*string) *GetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordPolicyResponse) SetStatusCode(v int32) *GetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordPolicyResponse) SetBody(v *GetPasswordPolicyResponseBody) *GetPasswordPolicyResponse {
	s.Body = v
	return s
}

type GetSAMLProviderRequest struct {
	// The name of the IdP.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s GetSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderRequest) SetSAMLProviderName(v string) *GetSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type GetSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BAADB995-0C7A-476D-B293-7E94568EEDFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the IdP.
	SAMLProvider *GetSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s GetSAMLProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponseBody) SetRequestId(v string) *GetSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSAMLProviderResponseBody) SetSAMLProvider(v *GetSAMLProviderResponseBodySAMLProvider) *GetSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

type GetSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata file, which is Base64 encoded.
	//
	// example:
	//
	// PD94bWwgdmVy****
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-22T02:51:20Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetSAMLProviderResponseBodySAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetArn(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetEncodedSAMLMetadataDocument(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

type GetSAMLProviderResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponse) SetHeaders(v map[string]*string) *GetSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *GetSAMLProviderResponse) SetStatusCode(v int32) *GetSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSAMLProviderResponse) SetBody(v *GetSAMLProviderResponseBody) *GetSAMLProviderResponse {
	s.Body = v
	return s
}

type GetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30C9068D-FBAA-4998-9986-8A562FED0BC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of security preferences.
	SecurityPreference *GetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBody) SetRequestId(v string) *GetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityPreferenceResponseBody) SetSecurityPreference(v *GetSecurityPreferenceResponseBodySecurityPreference) *GetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preference.
	MFAPreference *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The maximum idle periods. Unit: days.
	MaxIdleDays *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays `json:"MaxIdleDays,omitempty" xml:"MaxIdleDays,omitempty" type:"Struct"`
	// The personal information preference.
	PersonalInfoPreference *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference `json:"PersonalInfoPreference,omitempty" xml:"PersonalInfoPreference,omitempty" type:"Struct"`
	// The MFA method preference.
	VerificationPreference *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference `json:"VerificationPreference,omitempty" xml:"VerificationPreference,omitempty" type:"Struct"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetMaxIdleDays(v *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.MaxIdleDays = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetPersonalInfoPreference(v *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.PersonalInfoPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreference) SetVerificationPreference(v *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) *GetSecurityPreferenceResponseBodySecurityPreference {
	s.VerificationPreference = v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether a RAM user can use a passkey for logon. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Indicates whether RAM users can remember the multi-factor authentication (MFA) devices for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users. Unit: hours.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Indicates whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console. Valid values:
	//
	// 	- mandatory: MFA is required for all RAM users. If you use EnforceMFAForLogin, set the value to true.
	//
	// 	- independent (default): User-specific settings are applied. If you use EnforceMFAForLogin, set the value to false.
	//
	// 	- adaptive: MFA is required only for RAM users who initiated unusual logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Indicates whether to enable MFA for RAM users who initiated unusual logons. Valid values:
	//
	// 	- autonomous (default): yes. MFA is prompted for RAM users who initiated unusual logons. However, the RAM users are allowed to skip MFA.
	//
	// 	- enforceVerify: MFA is prompted for RAM users who initiated unusual logons and the RAM users cannot skip MFA.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToLoginWithPasskey(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetMFAOperationForLogin(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.MFAOperationForLogin = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetOperationForRiskLogin(v string) *GetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.OperationForRiskLogin = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *GetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays struct {
	// The maximum number of days that the AccessKey pair of a RAM user can stay unused. If an AccessKey pair is not used in the previous specified number of days, the AccessKey pair is automatically disabled on the next day. The default value is 730. You cannot change the value.
	//
	// example:
	//
	// 730
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum number of days that a RAM user can stay idle. If a RAM user for whom console logon is enabled does not log on to the console in the previous specified number of days, console logon is automatically disabled for the RAM user on the next day. Single sign-on (SSO) is not involved. The default value is 730. You cannot change the value.
	//
	// example:
	//
	// 730
	MaxIdleDaysForUsers *int32 `json:"MaxIdleDaysForUsers,omitempty" xml:"MaxIdleDaysForUsers,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForAccessKeys(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForUsers(v int32) *GetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForUsers = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference struct {
	// Indicates whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) SetAllowUserToManagePersonalDingTalk(v bool) *GetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

type GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference struct {
	// The MFA methods.
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) SetVerificationTypes(v []*string) *GetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	s.VerificationTypes = v
	return s
}

type GetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *GetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityPreferenceResponse) SetStatusCode(v int32) *GetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecurityPreferenceResponse) SetBody(v *GetSecurityPreferenceResponseBody) *GetSecurityPreferenceResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName`, `UserId`, and `UserAccessKeyId`.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName`, `UserId`, and `UserAccessKeyId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// The name is in the format of `<username>@<AccountAlias>.onaliyun.com`. `<username>` indicates the name of the RAM user. `<AccountAlias>.onaliyun.com` indicates the default domain name.
	//
	// The value of `UserPrincipalName` must be `1 to 128` characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The value of `<username>` must be `1 to 64` characters in length.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName`, `UserId`, and `UserAccessKeyId`.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserAccessKeyId(v string) *GetUserRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) SetUserPrincipalName(v string) *GetUserRequest {
	s.UserPrincipalName = &v
	return s
}

type GetUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	// The description.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The last time when the RAM user logged on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The source of the RAM user. Valid value:
	//
	// 	- Manual: The RAM user is manually created in the RAM console.
	//
	// 	- SCIM: The RAM user is mapped by using System for Cross-domain Identity Management (SCIM).
	//
	// 	- CloudSSO: The RAM user is mapped from a CloudSSO user.
	//
	// example:
	//
	// CloudSSO
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The tags.
	Tags *GetUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the information about the RAM user was updated.
	//
	// example:
	//
	// 2020-10-13T07:39:22Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user, which is the prefix of the logon name of the RAM user.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetComments(v string) *GetUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCreateDate(v string) *GetUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastLoginDate(v string) *GetUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobilePhone(v string) *GetUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *GetUserResponseBodyUser) SetProvisionType(v string) *GetUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetTags(v *GetUserResponseBodyUserTags) *GetUserResponseBodyUser {
	s.Tags = v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateDate(v string) *GetUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserPrincipalName(v string) *GetUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

type GetUserResponseBodyUserTags struct {
	Tag []*GetUserResponseBodyUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserTags) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserTags) SetTag(v []*GetUserResponseBodyUserTagsTag) *GetUserResponseBodyUserTags {
	s.Tag = v
	return s
}

type GetUserResponseBodyUserTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// operator
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetUserResponseBodyUserTagsTag) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserTagsTag) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserTagsTag) SetTagKey(v string) *GetUserResponseBodyUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *GetUserResponseBodyUserTagsTag) SetTagValue(v string) *GetUserResponseBodyUserTagsTag {
	s.TagValue = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserInRecycleBinRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserId` and `UserAccessKeyId`.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `UserId` and `UserAccessKeyId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinRequest) SetUserAccessKeyId(v string) *GetUserInRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *GetUserInRecycleBinRequest) SetUserId(v string) *GetUserInRecycleBinRequest {
	s.UserId = &v
	return s
}

type GetUserInRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *GetUserInRecycleBinResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponseBody) SetRequestId(v string) *GetUserInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInRecycleBinResponseBody) SetUser(v *GetUserInRecycleBinResponseBodyUser) *GetUserInRecycleBinResponseBody {
	s.User = v
	return s
}

type GetUserInRecycleBinResponseBodyUser struct {
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-11T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the RAM user will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-12T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserInRecycleBinResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserInRecycleBinResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponseBodyUser) SetCreateDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetDeleteDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.DeleteDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetDisplayName(v string) *GetUserInRecycleBinResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetRecycleDate(v string) *GetUserInRecycleBinResponseBodyUser {
	s.RecycleDate = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetUserId(v string) *GetUserInRecycleBinResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserInRecycleBinResponseBodyUser) SetUserPrincipalName(v string) *GetUserInRecycleBinResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

type GetUserInRecycleBinResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *GetUserInRecycleBinResponse) SetHeaders(v map[string]*string) *GetUserInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *GetUserInRecycleBinResponse) SetStatusCode(v int32) *GetUserInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInRecycleBinResponse) SetBody(v *GetUserInRecycleBinResponseBody) *GetUserInRecycleBinResponse {
	s.Body = v
	return s
}

type GetUserMFAInfoRequest struct {
	// The logon name of the RAM user. This parameter is differently set in the following scenarios:
	//
	// 	- If you use a RAM user to call this operation, this parameter can be left empty. If you do not specify this parameter, information about the MFA device that is bound to the RAM user is queried.
	//
	// 	- If you use an Alibaba Cloud account to call this operation, you must set this parameter to the logon name of the RAM user that you want to query.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) SetUserPrincipalName(v string) *GetUserMFAInfoRequest {
	s.UserPrincipalName = &v
	return s
}

type GetUserMFAInfoResponseBody struct {
	// Indicates whether the MFA device is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsMFAEnable *bool `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty"`
	// The information about the MFA device.
	MFADevice *GetUserMFAInfoResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FCF7322A-20A9-4F68-8B7F-F86958839BC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserMFAInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBody) SetIsMFAEnable(v bool) *GetUserMFAInfoResponseBody {
	s.IsMFAEnable = &v
	return s
}

func (s *GetUserMFAInfoResponseBody) SetMFADevice(v *GetUserMFAInfoResponseBodyMFADevice) *GetUserMFAInfoResponseBody {
	s.MFADevice = v
	return s
}

func (s *GetUserMFAInfoResponseBody) SetRequestId(v string) *GetUserMFAInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetUserMFAInfoResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The type of the MFA device. Valid values:
	//
	// 	- VMFA: virtual MFA device.
	//
	// 	- U2F: Universal 2nd Factor (U2F) security key.
	//
	// example:
	//
	// VMFA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetUserMFAInfoResponseBodyMFADevice) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetSerialNumber(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *GetUserMFAInfoResponseBodyMFADevice) SetType(v string) *GetUserMFAInfoResponseBodyMFADevice {
	s.Type = &v
	return s
}

type GetUserMFAInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMFAInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponse) SetHeaders(v map[string]*string) *GetUserMFAInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserMFAInfoResponse) SetStatusCode(v int32) *GetUserMFAInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMFAInfoResponse) SetBody(v *GetUserMFAInfoResponseBody) *GetUserMFAInfoResponse {
	s.Body = v
	return s
}

type GetUserSsoSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 69FC3E5E-D3D9-434B-90CA-BBA8E0551A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of user-based SSO.
	UserSsoSettings *GetUserSsoSettingsResponseBodyUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" type:"Struct"`
}

func (s GetUserSsoSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponseBody) SetRequestId(v string) *GetUserSsoSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserSsoSettingsResponseBody) SetUserSsoSettings(v *GetUserSsoSettingsResponseBodyUserSsoSettings) *GetUserSsoSettingsResponseBody {
	s.UserSsoSettings = v
	return s
}

type GetUserSsoSettingsResponseBodyUserSsoSettings struct {
	// The auxiliary domain name.
	//
	// example:
	//
	// example.com
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file, which is Base64-encoded.
	//
	// example:
	//
	// PD94bWwgdmVy****
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Indicates whether user-based SSO is enabled.
	//
	// example:
	//
	// false
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Indicates whether the Security Assertion Markup Language (SAML) SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the identity provider (IdP) for SSO must have a domain name as the suffix.
	//
	// 	- If the value of the parameter is `true`, the `<saml:NameID>` element **must*	- be in the `username@domain` format. You can set `domain` to the default domain name or the configured domain alias.
	//
	// 	- If the value of the parameter is `false`, the `<saml:NameID>` element **must*	- be in the `username` format and **cannot*	- contain the `domain` suffix.
	//
	// The default value is `true`.
	//
	// example:
	//
	// true
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s GetUserSsoSettingsResponseBodyUserSsoSettings) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsResponseBodyUserSsoSettings) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetAuxiliaryDomain(v string) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetMetadataDocument(v string) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoEnabled(v bool) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *GetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoLoginWithDomain(v bool) *GetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoLoginWithDomain = &v
	return s
}

type GetUserSsoSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserSsoSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserSsoSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponse) SetHeaders(v map[string]*string) *GetUserSsoSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetUserSsoSettingsResponse) SetStatusCode(v int32) *GetUserSsoSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserSsoSettingsResponse) SetBody(v *GetUserSsoSettingsResponseBody) *GetUserSsoSettingsResponse {
	s.Body = v
	return s
}

type GetVerificationInfoRequest struct {
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetVerificationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoRequest) SetUserPrincipalName(v string) *GetVerificationInfoRequest {
	s.UserPrincipalName = &v
	return s
}

type GetVerificationInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the email.
	SecurityEmailDevice *GetVerificationInfoResponseBodySecurityEmailDevice `json:"SecurityEmailDevice,omitempty" xml:"SecurityEmailDevice,omitempty" type:"Struct"`
	// The information about the mobile phone.
	SecurityPhoneDevice *GetVerificationInfoResponseBodySecurityPhoneDevice `json:"SecurityPhoneDevice,omitempty" xml:"SecurityPhoneDevice,omitempty" type:"Struct"`
}

func (s GetVerificationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBody) SetRequestId(v string) *GetVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVerificationInfoResponseBody) SetSecurityEmailDevice(v *GetVerificationInfoResponseBodySecurityEmailDevice) *GetVerificationInfoResponseBody {
	s.SecurityEmailDevice = v
	return s
}

func (s *GetVerificationInfoResponseBody) SetSecurityPhoneDevice(v *GetVerificationInfoResponseBodySecurityPhoneDevice) *GetVerificationInfoResponseBody {
	s.SecurityPhoneDevice = v
	return s
}

type GetVerificationInfoResponseBodySecurityEmailDevice struct {
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The status of the email. Valid values:
	//
	// 	- active: The email is activated.
	//
	// 	- pending: The email is pending for activation.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVerificationInfoResponseBodySecurityEmailDevice) String() string {
	return tea.Prettify(s)
}

func (s GetVerificationInfoResponseBodySecurityEmailDevice) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) SetEmail(v string) *GetVerificationInfoResponseBodySecurityEmailDevice {
	s.Email = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityEmailDevice) SetStatus(v string) *GetVerificationInfoResponseBodySecurityEmailDevice {
	s.Status = &v
	return s
}

type GetVerificationInfoResponseBodySecurityPhoneDevice struct {
	// The international dialing code.
	//
	// example:
	//
	// 86
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 13900001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The status of the mobile phone. Valid values:
	//
	// 	- active: The mobile phone is activated.
	//
	// 	- pending: The mobile phone is pending for activation.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVerificationInfoResponseBodySecurityPhoneDevice) String() string {
	return tea.Prettify(s)
}

func (s GetVerificationInfoResponseBodySecurityPhoneDevice) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetAreaCode(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.AreaCode = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetPhoneNumber(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.PhoneNumber = &v
	return s
}

func (s *GetVerificationInfoResponseBodySecurityPhoneDevice) SetStatus(v string) *GetVerificationInfoResponseBodySecurityPhoneDevice {
	s.Status = &v
	return s
}

type GetVerificationInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVerificationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVerificationInfoResponse) SetHeaders(v map[string]*string) *GetVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVerificationInfoResponse) SetStatusCode(v int32) *GetVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVerificationInfoResponse) SetBody(v *GetVerificationInfoResponseBody) *GetVerificationInfoResponse {
	s.Body = v
	return s
}

type ListAccessKeysRequest struct {
	// The logon name of the RAM user.
	//
	// If this parameter is empty, the AccessKey pairs of the current user are queried.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListAccessKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysRequest) SetUserPrincipalName(v string) *ListAccessKeysRequest {
	s.UserPrincipalName = &v
	return s
}

type ListAccessKeysResponseBody struct {
	// The AccessKey pairs.
	AccessKeys *ListAccessKeysResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBody) SetAccessKeys(v *ListAccessKeysResponseBodyAccessKeys) *ListAccessKeysResponseBody {
	s.AccessKeys = v
	return s
}

func (s *ListAccessKeysResponseBody) SetRequestId(v string) *ListAccessKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListAccessKeysResponseBodyAccessKeys struct {
	AccessKey []*ListAccessKeysResponseBodyAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Repeated"`
}

func (s ListAccessKeysResponseBodyAccessKeys) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeys) SetAccessKey(v []*ListAccessKeysResponseBodyAccessKeysAccessKey) *ListAccessKeysResponseBodyAccessKeys {
	s.AccessKey = v
	return s
}

type ListAccessKeysResponseBodyAccessKeysAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-13T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the AccessKey pair was updated.
	//
	// example:
	//
	// 2020-10-13T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseBodyAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetStatus(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.Status = &v
	return s
}

func (s *ListAccessKeysResponseBodyAccessKeysAccessKey) SetUpdateDate(v string) *ListAccessKeysResponseBodyAccessKeysAccessKey {
	s.UpdateDate = &v
	return s
}

type ListAccessKeysResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponse) SetHeaders(v map[string]*string) *ListAccessKeysResponse {
	s.Headers = v
	return s
}

func (s *ListAccessKeysResponse) SetStatusCode(v int32) *ListAccessKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessKeysResponse) SetBody(v *ListAccessKeysResponseBody) *ListAccessKeysResponse {
	s.Body = v
	return s
}

type ListAccessKeysInRecycleBinRequest struct {
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAccessKeysInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinRequest) SetUserId(v string) *ListAccessKeysInRecycleBinRequest {
	s.UserId = &v
	return s
}

type ListAccessKeysInRecycleBinResponseBody struct {
	// The information about the AccessKey pairs.
	AccessKeys *ListAccessKeysInRecycleBinResponseBodyAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBody) SetAccessKeys(v *ListAccessKeysInRecycleBinResponseBodyAccessKeys) *ListAccessKeysInRecycleBinResponseBody {
	s.AccessKeys = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBody) SetRequestId(v string) *ListAccessKeysInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type ListAccessKeysInRecycleBinResponseBodyAccessKeys struct {
	AccessKey []*ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Repeated"`
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeys) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeys) SetAccessKey(v []*ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) *ListAccessKeysInRecycleBinResponseBodyAccessKeys {
	s.AccessKey = v
	return s
}

type ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// LTAI*******************
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-11T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the AccessKey pair will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-12T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The time when the AccessKey pair was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetDeleteDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.DeleteDate = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey) SetRecycleDate(v string) *ListAccessKeysInRecycleBinResponseBodyAccessKeysAccessKey {
	s.RecycleDate = &v
	return s
}

type ListAccessKeysInRecycleBinResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAccessKeysInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAccessKeysInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysInRecycleBinResponse) SetHeaders(v map[string]*string) *ListAccessKeysInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *ListAccessKeysInRecycleBinResponse) SetStatusCode(v int32) *ListAccessKeysInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAccessKeysInRecycleBinResponse) SetBody(v *ListAccessKeysInRecycleBinResponseBody) *ListAccessKeysInRecycleBinResponse {
	s.Body = v
	return s
}

type ListAppSecretIdsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListAppSecretIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsRequest) SetAppId(v string) *ListAppSecretIdsRequest {
	s.AppId = &v
	return s
}

type ListAppSecretIdsResponseBody struct {
	// The details of the application secret.
	AppSecrets *ListAppSecretIdsResponseBodyAppSecrets `json:"AppSecrets,omitempty" xml:"AppSecrets,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5F2FD500-7173-47D6-BD2F-EB60879B4F16
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppSecretIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBody) SetAppSecrets(v *ListAppSecretIdsResponseBodyAppSecrets) *ListAppSecretIdsResponseBody {
	s.AppSecrets = v
	return s
}

func (s *ListAppSecretIdsResponseBody) SetRequestId(v string) *ListAppSecretIdsResponseBody {
	s.RequestId = &v
	return s
}

type ListAppSecretIdsResponseBodyAppSecrets struct {
	AppSecret []*ListAppSecretIdsResponseBodyAppSecretsAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Repeated"`
}

func (s ListAppSecretIdsResponseBodyAppSecrets) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponseBodyAppSecrets) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBodyAppSecrets) SetAppSecret(v []*ListAppSecretIdsResponseBodyAppSecretsAppSecret) *ListAppSecretIdsResponseBodyAppSecrets {
	s.AppSecret = v
	return s
}

type ListAppSecretIdsResponseBodyAppSecretsAppSecret struct {
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-26T03:18:39Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s ListAppSecretIdsResponseBodyAppSecretsAppSecret) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponseBodyAppSecretsAppSecret) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetAppId(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.AppId = &v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetAppSecretId(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetCreateDate(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.CreateDate = &v
	return s
}

type ListAppSecretIdsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppSecretIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppSecretIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponse) SetHeaders(v map[string]*string) *ListAppSecretIdsResponse {
	s.Headers = v
	return s
}

func (s *ListAppSecretIdsResponse) SetStatusCode(v int32) *ListAppSecretIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppSecretIdsResponse) SetBody(v *ListAppSecretIdsResponseBody) *ListAppSecretIdsResponse {
	s.Body = v
	return s
}

type ListApplicationProvisionInfosRequest struct {
	// The source of the applications. Valid values:
	//
	// 	- inner: The applications are from the current account.
	//
	// 	- external: The applications are from other accounts.
	//
	// example:
	//
	// external
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListApplicationProvisionInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosRequest) SetSourceType(v string) *ListApplicationProvisionInfosRequest {
	s.SourceType = &v
	return s
}

type ListApplicationProvisionInfosResponseBody struct {
	// The information about the installed applications.
	ApplicationProvisionInfos *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos `json:"ApplicationProvisionInfos,omitempty" xml:"ApplicationProvisionInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E403EBFD-C997-489D-BFC7-37C05E66D67C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBody) SetApplicationProvisionInfos(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) *ListApplicationProvisionInfosResponseBody {
	s.ApplicationProvisionInfos = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBody) SetRequestId(v string) *ListApplicationProvisionInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos struct {
	ApplicationProvisionInfo []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) SetApplicationProvisionInfo(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos {
	s.ApplicationProvisionInfo = v
	return s
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// GiteePrd
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted to the application.
	DelegatedScope *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAccountId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetCreateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDelegatedScope(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDisplayName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetUpdateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope struct {
	// The information about the permissions that are granted to the application.
	PredefinedScopes *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ListApplicationProvisionInfosResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationProvisionInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationProvisionInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationProvisionInfosResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponse) SetHeaders(v map[string]*string) *ListApplicationProvisionInfosResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationProvisionInfosResponse) SetStatusCode(v int32) *ListApplicationProvisionInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationProvisionInfosResponse) SetBody(v *ListApplicationProvisionInfosResponseBody) *ListApplicationProvisionInfosResponse {
	s.Body = v
	return s
}

type ListApplicationsResponseBody struct {
	// The information about the application.
	Applications *ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CE458B58-8C40-46F7-A9D4-CB85136B0C06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

type ListApplicationsResponseBodyApplications struct {
	Application []*ListApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) SetApplication(v []*ListApplicationsResponseBodyApplicationsApplication) *ListApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

type ListApplicationsResponseBodyApplicationsApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 441442900344560****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Valid values:
	//
	// 	- WebApp: a web application.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop or mobile operating system.
	//
	// 	- ServerApp: an application that can access Alibaba Cloud services without the need for user logon. Only applications that synchronize user information based on the System for Cross-domain Identity Management (SCIM) protocol are supported.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T09:33:22Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The callback URLs.
	RedirectUris *ListApplicationsResponseBodyApplicationsApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T09:33:22Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplication) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAccessTokenValidity(v int32) *ListApplicationsResponseBodyApplicationsApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAccountId(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AccountId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppId(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppType(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetCreateDate(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetDelegatedScope(v *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) *ListApplicationsResponseBodyApplicationsApplication {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetDisplayName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetIsMultiTenant(v bool) *ListApplicationsResponseBodyApplicationsApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetRedirectUris(v *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) *ListApplicationsResponseBodyApplicationsApplication {
	s.RedirectUris = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetRefreshTokenValidity(v int32) *ListApplicationsResponseBodyApplicationsApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetSecretRequired(v bool) *ListApplicationsResponseBodyApplicationsApplication {
	s.SecretRequired = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetUpdateDate(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.UpdateDate = &v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) SetPredefinedScopes(v *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

type ListApplicationsResponseBodyApplicationsApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) SetRedirectUri(v []*string) *ListApplicationsResponseBodyApplicationsApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListExternalApplicationsResponseBody struct {
	// The information about the external applications.
	ExternalApplications *ListExternalApplicationsResponseBodyExternalApplications `json:"ExternalApplications,omitempty" xml:"ExternalApplications,omitempty" type:"Struct"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 936E1D9C-157D-45BD-8A3B-81C0716EB077
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExternalApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBody) SetExternalApplications(v *ListExternalApplicationsResponseBodyExternalApplications) *ListExternalApplicationsResponseBody {
	s.ExternalApplications = v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetIsTruncated(v bool) *ListExternalApplicationsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetMarker(v string) *ListExternalApplicationsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetRequestId(v string) *ListExternalApplicationsResponseBody {
	s.RequestId = &v
	return s
}

type ListExternalApplicationsResponseBodyExternalApplications struct {
	ExternalApplication []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Repeated"`
}

func (s ListExternalApplicationsResponseBodyExternalApplications) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplications) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplications) SetExternalApplication(v []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) *ListExternalApplicationsResponseBodyExternalApplications {
	s.ExternalApplication = v
	return s
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication struct {
	// The name of the external application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.153082740420****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the external application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693318000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted to the external application.
	DelegatedScope *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the external application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the external application.
	//
	// example:
	//
	// 407426893752729****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account for which the external application was installed.
	//
	// example:
	//
	// 173082740420****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time of the external application. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetAppPrincipalName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetCreateDate(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetDelegatedScope(v *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetDisplayName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetForeignAppId(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetTenantId(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.TenantId = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetUpdateDate(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.UpdateDate = &v
	return s
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope struct {
	// The information about the permissions that are granted to the external application.
	PredefinedScopes *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) SetPredefinedScopes(v *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ListExternalApplicationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExternalApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponse) SetHeaders(v map[string]*string) *ListExternalApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListExternalApplicationsResponse) SetStatusCode(v int32) *ListExternalApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalApplicationsResponse) SetBody(v *ListExternalApplicationsResponseBody) *ListExternalApplicationsResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetMarker(v string) *ListGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListGroupsRequest) SetMaxItems(v int32) *ListGroupsRequest {
	s.MaxItems = &v
	return s
}

type ListGroupsResponseBody struct {
	// The information about the RAM user groups.
	Groups *ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 065527AA-2F2E-AD7C-7484-F2626CFE4934
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetGroups(v *ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetIsTruncated(v bool) *ListGroupsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponseBody) SetMarker(v string) *ListGroupsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupsResponseBodyGroups struct {
	Group []*ListGroupsResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) SetGroup(v []*ListGroupsResponseBodyGroupsGroup) *ListGroupsResponseBodyGroups {
	s.Group = v
	return s
}

type ListGroupsResponseBodyGroupsGroup struct {
	// The description.
	//
	// example:
	//
	// Dev-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-19T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// dev-team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-19T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListGroupsResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroupsGroup) SetComments(v string) *ListGroupsResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetCreateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.CreateDate = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetDisplayName(v string) *ListGroupsResponseBodyGroupsGroup {
	s.DisplayName = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroupsGroup) SetUpdateDate(v string) *ListGroupsResponseBodyGroupsGroup {
	s.UpdateDate = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListGroupsForUserRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) SetUserPrincipalName(v string) *ListGroupsForUserRequest {
	s.UserPrincipalName = &v
	return s
}

type ListGroupsForUserResponseBody struct {
	// The information about the RAM user groups.
	Groups *ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7158A935-FB5E-49A7-8E52-FDA5B2B67247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) SetGroups(v *ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetRequestId(v string) *ListGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupsForUserResponseBodyGroups struct {
	Group []*ListGroupsForUserResponseBodyGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Repeated"`
}

func (s ListGroupsForUserResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroup(v []*ListGroupsForUserResponseBodyGroupsGroup) *ListGroupsForUserResponseBodyGroups {
	s.Group = v
	return s
}

type ListGroupsForUserResponseBodyGroupsGroup struct {
	// The description.
	//
	// example:
	//
	// Test-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the RAM user was added.
	//
	// example:
	//
	// 2020-10-20T06:57:00Z
	JoinDate *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
}

func (s ListGroupsForUserResponseBodyGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetComments(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetDisplayName(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.DisplayName = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupId(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetGroupName(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroupsGroup) SetJoinDate(v string) *ListGroupsForUserResponseBodyGroupsGroup {
	s.JoinDate = &v
	return s
}

type ListGroupsForUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) SetHeaders(v map[string]*string) *ListGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForUserResponse) SetStatusCode(v int32) *ListGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForUserResponse) SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse {
	s.Body = v
	return s
}

type ListOIDCProvidersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListOIDCProvidersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOIDCProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersRequest) SetMarker(v string) *ListOIDCProvidersRequest {
	s.Marker = &v
	return s
}

func (s *ListOIDCProvidersRequest) SetMaxItems(v int32) *ListOIDCProvidersRequest {
	s.MaxItems = &v
	return s
}

type ListOIDCProvidersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The information about the OIDC IdP.
	OIDCProviders *ListOIDCProvidersResponseBodyOIDCProviders `json:"OIDCProviders,omitempty" xml:"OIDCProviders,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D2148337-B86A-57F0-8B31-EB7BE0125226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOIDCProvidersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOIDCProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBody) SetIsTruncated(v bool) *ListOIDCProvidersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetMarker(v string) *ListOIDCProvidersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetOIDCProviders(v *ListOIDCProvidersResponseBodyOIDCProviders) *ListOIDCProvidersResponseBody {
	s.OIDCProviders = v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetRequestId(v string) *ListOIDCProvidersResponseBody {
	s.RequestId = &v
	return s
}

type ListOIDCProvidersResponseBodyOIDCProviders struct {
	OIDCProvider []*ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Repeated"`
}

func (s ListOIDCProvidersResponseBodyOIDCProviders) String() string {
	return tea.Prettify(s)
}

func (s ListOIDCProvidersResponseBodyOIDCProviders) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBodyOIDCProviders) SetOIDCProvider(v []*ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) *ListOIDCProvidersResponseBodyOIDCProviders {
	s.OIDCProvider = v
	return s
}

type ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client, If you want to specify multiple client IDs, separate the client IDs with commas (,).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetArn(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Arn = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetClientIds(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetCreateDate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetDescription(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Description = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetFingerprints(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetGmtCreate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetGmtModified(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetIssuanceLimitTime(v int64) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetIssuerUrl(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetOIDCProviderName(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetUpdateDate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.UpdateDate = &v
	return s
}

type ListOIDCProvidersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOIDCProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOIDCProvidersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOIDCProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponse) SetHeaders(v map[string]*string) *ListOIDCProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListOIDCProvidersResponse) SetStatusCode(v int32) *ListOIDCProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOIDCProvidersResponse) SetBody(v *ListOIDCProvidersResponseBody) *ListOIDCProvidersResponse {
	s.Body = v
	return s
}

type ListPasskeysRequest struct {
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListPasskeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPasskeysRequest) GoString() string {
	return s.String()
}

func (s *ListPasskeysRequest) SetUserPrincipalName(v string) *ListPasskeysRequest {
	s.UserPrincipalName = &v
	return s
}

type ListPasskeysResponseBody struct {
	// The information about the passkeys.
	Passkeys []*ListPasskeysResponseBodyPasskeys `json:"Passkeys,omitempty" xml:"Passkeys,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPasskeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPasskeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponseBody) SetPasskeys(v []*ListPasskeysResponseBodyPasskeys) *ListPasskeysResponseBody {
	s.Passkeys = v
	return s
}

func (s *ListPasskeysResponseBody) SetRequestId(v string) *ListPasskeysResponseBody {
	s.RequestId = &v
	return s
}

type ListPasskeysResponseBodyPasskeys struct {
	// The time when the passkey was created. The value is a timestamp.
	//
	// example:
	//
	// 1737450279
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the passkey was last used. The value is a timestamp.
	//
	// example:
	//
	// 1737450390
	LastUseDate *string `json:"LastUseDate,omitempty" xml:"LastUseDate,omitempty"`
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The name of the passkey.
	//
	// example:
	//
	// device1
	PasskeyName *string `json:"PasskeyName,omitempty" xml:"PasskeyName,omitempty"`
}

func (s ListPasskeysResponseBodyPasskeys) String() string {
	return tea.Prettify(s)
}

func (s ListPasskeysResponseBodyPasskeys) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponseBodyPasskeys) SetCreateDate(v string) *ListPasskeysResponseBodyPasskeys {
	s.CreateDate = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetLastUseDate(v string) *ListPasskeysResponseBodyPasskeys {
	s.LastUseDate = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetPasskeyId(v string) *ListPasskeysResponseBodyPasskeys {
	s.PasskeyId = &v
	return s
}

func (s *ListPasskeysResponseBodyPasskeys) SetPasskeyName(v string) *ListPasskeysResponseBodyPasskeys {
	s.PasskeyName = &v
	return s
}

type ListPasskeysResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPasskeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPasskeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPasskeysResponse) GoString() string {
	return s.String()
}

func (s *ListPasskeysResponse) SetHeaders(v map[string]*string) *ListPasskeysResponse {
	s.Headers = v
	return s
}

func (s *ListPasskeysResponse) SetStatusCode(v int32) *ListPasskeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPasskeysResponse) SetBody(v *ListPasskeysResponseBody) *ListPasskeysResponse {
	s.Body = v
	return s
}

type ListPredefinedScopesRequest struct {
	// The type of the application. Valid values:
	//
	// 	- WebApp
	//
	// 	- NativeApp
	//
	// 	- ServerApp
	//
	// If this parameter is empty, the permissions on all types of applications are queried.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s ListPredefinedScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesRequest) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesRequest) SetAppType(v string) *ListPredefinedScopesRequest {
	s.AppType = &v
	return s
}

type ListPredefinedScopesResponseBody struct {
	// The information about application permissions.
	PredefinedScopes *ListPredefinedScopesResponseBodyPredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 955C096D-EC99-480B-AF37-3921109107D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPredefinedScopesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBody) SetPredefinedScopes(v *ListPredefinedScopesResponseBodyPredefinedScopes) *ListPredefinedScopesResponseBody {
	s.PredefinedScopes = v
	return s
}

func (s *ListPredefinedScopesResponseBody) SetRequestId(v string) *ListPredefinedScopesResponseBody {
	s.RequestId = &v
	return s
}

type ListPredefinedScopesResponseBodyPredefinedScopes struct {
	PredefinedScope []*ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListPredefinedScopesResponseBodyPredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponseBodyPredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopes) SetPredefinedScope(v []*ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) *ListPredefinedScopesResponseBodyPredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope struct {
	// The description of the permission scope.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission scope.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) SetDescription(v string) *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope) SetName(v string) *ListPredefinedScopesResponseBodyPredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ListPredefinedScopesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPredefinedScopesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPredefinedScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponse) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponse) SetHeaders(v map[string]*string) *ListPredefinedScopesResponse {
	s.Headers = v
	return s
}

func (s *ListPredefinedScopesResponse) SetStatusCode(v int32) *ListPredefinedScopesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPredefinedScopesResponse) SetBody(v *ListPredefinedScopesResponseBody) *ListPredefinedScopesResponse {
	s.Body = v
	return s
}

type ListRecentGovernanceMetricsResponseBody struct {
	// The time when the report was generated.
	//
	// example:
	//
	// 2025-02-10T02:11:23Z
	GenerateTime *string `json:"GenerateTime,omitempty" xml:"GenerateTime,omitempty"`
	// The metric values of all governance items. The value of the parameter is an array, and each row in the array contains the metric value of a governance item.
	GovernanceMetrics *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics `json:"GovernanceMetrics,omitempty" xml:"GovernanceMetrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 49846A91-C1C5-5C2B-BC64-8B0B7BADB4C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRecentGovernanceMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBody) SetGenerateTime(v string) *ListRecentGovernanceMetricsResponseBody {
	s.GenerateTime = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBody) SetGovernanceMetrics(v *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) *ListRecentGovernanceMetricsResponseBody {
	s.GovernanceMetrics = v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBody) SetRequestId(v string) *ListRecentGovernanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

type ListRecentGovernanceMetricsResponseBodyGovernanceMetrics struct {
	GovernanceMetric []*ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric `json:"GovernanceMetric,omitempty" xml:"GovernanceMetric,omitempty" type:"Repeated"`
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) String() string {
	return tea.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics) SetGovernanceMetric(v []*ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) *ListRecentGovernanceMetricsResponseBodyGovernanceMetrics {
	s.GovernanceMetric = v
	return s
}

type ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric struct {
	// The name of the governance item.
	//
	// example:
	//
	// RecentAccountLoginTimes
	GovernanceItem *string `json:"GovernanceItem,omitempty" xml:"GovernanceItem,omitempty"`
	// The type of the metric value. Valid values:
	//
	// 	- Number
	//
	// 	- String
	//
	// 	- Boolean
	//
	// example:
	//
	// Number
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The metric value. The type of the metric value is determined by `MetricType`.
	//
	// example:
	//
	// 5
	MetricValue interface{} `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) String() string {
	return tea.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetGovernanceItem(v string) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.GovernanceItem = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetMetricType(v string) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.MetricType = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric) SetMetricValue(v interface{}) *ListRecentGovernanceMetricsResponseBodyGovernanceMetricsGovernanceMetric {
	s.MetricValue = v
	return s
}

type ListRecentGovernanceMetricsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentGovernanceMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentGovernanceMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecentGovernanceMetricsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentGovernanceMetricsResponse) SetHeaders(v map[string]*string) *ListRecentGovernanceMetricsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentGovernanceMetricsResponse) SetStatusCode(v int32) *ListRecentGovernanceMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentGovernanceMetricsResponse) SetBody(v *ListRecentGovernanceMetricsResponseBody) *ListRecentGovernanceMetricsResponse {
	s.Body = v
	return s
}

type ListSAMLProvidersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListSAMLProvidersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersRequest) SetMarker(v string) *ListSAMLProvidersRequest {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersRequest) SetMaxItems(v int32) *ListSAMLProvidersRequest {
	s.MaxItems = &v
	return s
}

type ListSAMLProvidersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2D8B70D3-E194-41C9-93C5-F6A10D716D24
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about IdPs.
	SAMLProviders *ListSAMLProvidersResponseBodySAMLProviders `json:"SAMLProviders,omitempty" xml:"SAMLProviders,omitempty" type:"Struct"`
}

func (s ListSAMLProvidersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBody) SetIsTruncated(v bool) *ListSAMLProvidersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetMarker(v string) *ListSAMLProvidersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetRequestId(v string) *ListSAMLProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSAMLProvidersResponseBody) SetSAMLProviders(v *ListSAMLProvidersResponseBodySAMLProviders) *ListSAMLProvidersResponseBody {
	s.SAMLProviders = v
	return s
}

type ListSAMLProvidersResponseBodySAMLProviders struct {
	SAMLProvider []*ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Repeated"`
}

func (s ListSAMLProvidersResponseBodySAMLProviders) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponseBodySAMLProviders) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBodySAMLProviders) SetSAMLProvider(v []*ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) *ListSAMLProvidersResponseBodySAMLProviders {
	s.SAMLProvider = v
	return s
}

type ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-22T06:26:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-22T06:26:15Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetArn(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.Arn = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetCreateDate(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetDescription(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.Description = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetSAMLProviderName(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider) SetUpdateDate(v string) *ListSAMLProvidersResponseBodySAMLProvidersSAMLProvider {
	s.UpdateDate = &v
	return s
}

type ListSAMLProvidersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSAMLProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSAMLProvidersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponse) SetHeaders(v map[string]*string) *ListSAMLProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListSAMLProvidersResponse) SetStatusCode(v int32) *ListSAMLProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSAMLProvidersResponse) SetBody(v *ListSAMLProvidersResponseBody) *ListSAMLProvidersResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of PageSize, the value of IsTruncated will be true. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of resource N.
	//
	// Valid values of N: 1 to 50. If ResourceType is set to user, the resource ID is the ID of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The name of resource N.
	//
	// Valid values of N: 1 to 50. If ResourceType is set to user, the resource name is the name of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	ResourcePrincipalName []*string `json:"ResourcePrincipalName,omitempty" xml:"ResourcePrincipalName,omitempty" type:"Repeated"`
	// The type of the resource. Valid value:
	//
	// 	- user: a RAM user
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag value.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourcePrincipalName(v []*string) *ListTagResourcesRequest {
	s.ResourcePrincipalName = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The marker. This parameter is returned only if the value of IsTruncated is true. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4BE83135-0B08-467C-B3A2-27B312FD0F57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag key.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetIsTruncated(v bool) *ListTagResourcesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the resource.
	//
	// example:
	//
	// 20732900249392****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- user: a RAM user
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// operator
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUserBasicInfosRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be `true`.
	//
	// Valid values: 1 to 1000. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The status of the RAM users that you want to query. Valid values: active, freeze, and active,freeze. If you set the value to active,freeze, RAM users in both active and freeze states are queried. If you leave the parameter empty, the value active is used by default. If the Tag parameter is specified, you cannot specify the Status parameter. In this case, RAM users in both states are queried.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tag []*ListUserBasicInfosRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUserBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosRequest) SetMarker(v string) *ListUserBasicInfosRequest {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetMaxItems(v int32) *ListUserBasicInfosRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetStatus(v string) *ListUserBasicInfosRequest {
	s.Status = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetTag(v []*ListUserBasicInfosRequestTag) *ListUserBasicInfosRequest {
	s.Tag = v
	return s
}

type ListUserBasicInfosRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserBasicInfosRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosRequestTag) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosRequestTag) SetKey(v string) *ListUserBasicInfosRequestTag {
	s.Key = &v
	return s
}

func (s *ListUserBasicInfosRequestTag) SetValue(v string) *ListUserBasicInfosRequestTag {
	s.Value = &v
	return s
}

type ListUserBasicInfosResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF2B25FD-CADE-445B-BE4D-E082E0FF1A0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The basic information about the RAM users.
	UserBasicInfos *ListUserBasicInfosResponseBodyUserBasicInfos `json:"UserBasicInfos,omitempty" xml:"UserBasicInfos,omitempty" type:"Struct"`
}

func (s ListUserBasicInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBody) SetIsTruncated(v bool) *ListUserBasicInfosResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetMarker(v string) *ListUserBasicInfosResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetRequestId(v string) *ListUserBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserBasicInfosResponseBody) SetUserBasicInfos(v *ListUserBasicInfosResponseBodyUserBasicInfos) *ListUserBasicInfosResponseBody {
	s.UserBasicInfos = v
	return s
}

type ListUserBasicInfosResponseBodyUserBasicInfos struct {
	UserBasicInfo []*ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo `json:"UserBasicInfo,omitempty" xml:"UserBasicInfo,omitempty" type:"Repeated"`
}

func (s ListUserBasicInfosResponseBodyUserBasicInfos) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponseBodyUserBasicInfos) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfos) SetUserBasicInfo(v []*ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) *ListUserBasicInfosResponseBodyUserBasicInfos {
	s.UserBasicInfo = v
	return s
}

type ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the RAM user.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetDisplayName(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.DisplayName = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetStatus(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.Status = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetUserId(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.UserId = &v
	return s
}

func (s *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo) SetUserPrincipalName(v string) *ListUserBasicInfosResponseBodyUserBasicInfosUserBasicInfo {
	s.UserPrincipalName = &v
	return s
}

type ListUserBasicInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponse) SetHeaders(v map[string]*string) *ListUserBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListUserBasicInfosResponse) SetStatusCode(v int32) *ListUserBasicInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserBasicInfosResponse) SetBody(v *ListUserBasicInfosResponseBody) *ListUserBasicInfosResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 1000. Default value: 1000.
	//
	// example:
	//
	// 1000
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The status of the RAM users that you want to query. Valid values: active, freeze, and active,freeze. If you leave the parameter empty, the value active is used by default. If you specify a value for the Tag parameter, users in both states are queried.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags. A maximum number of 20 tags are supported.
	Tag []*ListUsersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetMarker(v string) *ListUsersRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersRequest) SetMaxItems(v int32) *ListUsersRequest {
	s.MaxItems = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) SetTag(v []*ListUsersRequestTag) *ListUsersRequest {
	s.Tag = v
	return s
}

type ListUsersRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUsersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequestTag) GoString() string {
	return s.String()
}

func (s *ListUsersRequestTag) SetKey(v string) *ListUsersRequestTag {
	s.Key = &v
	return s
}

func (s *ListUsersRequestTag) SetValue(v string) *ListUsersRequestTag {
	s.Value = &v
	return s
}

type ListUsersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The parameter that is used to obtain the truncated part. It takes effect only when `IsTruncated` is set to `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetIsTruncated(v bool) *ListUsersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponseBody) SetMarker(v string) *ListUsersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	User []*ListUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetUser(v []*ListUsersResponseBodyUsersUser) *ListUsersResponseBodyUsers {
	s.User = v
	return s
}

type ListUsersResponseBodyUsersUser struct {
	// The description.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The point in time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter applies only to the Alibaba Cloud China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The timestamp when the RAM user last logged on to the console.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// >  This parameter applies only to the Alibaba Cloud China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The source of the RAM user. Valid values:
	//
	// 	- Manual: The RAM user is manually created in the RAM console.
	//
	// 	- SCIM: The RAM user is mapped by using System for Cross-domain Identity Management (SCIM).
	//
	// 	- CloudSSO: The RAM user is mapped from a CloudSSO user.
	//
	// example:
	//
	// CloudSSO
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the RAM user.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags *ListUsersResponseBodyUsersUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The point in time when the information about the RAM user was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-13T09:19:49Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUsersResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUser) SetComments(v string) *ListUsersResponseBodyUsersUser {
	s.Comments = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetCreateDate(v string) *ListUsersResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetDisplayName(v string) *ListUsersResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetEmail(v string) *ListUsersResponseBodyUsersUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetLastLoginDate(v string) *ListUsersResponseBodyUsersUser {
	s.LastLoginDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetMobilePhone(v string) *ListUsersResponseBodyUsersUser {
	s.MobilePhone = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetProvisionType(v string) *ListUsersResponseBodyUsersUser {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetStatus(v string) *ListUsersResponseBodyUsersUser {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetTags(v *ListUsersResponseBodyUsersUserTags) *ListUsersResponseBodyUsersUser {
	s.Tags = v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUpdateDate(v string) *ListUsersResponseBodyUsersUser {
	s.UpdateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserId(v string) *ListUsersResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

type ListUsersResponseBodyUsersUserTags struct {
	Tag []*ListUsersResponseBodyUsersUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsersUserTags) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserTags) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserTags) SetTag(v []*ListUsersResponseBodyUsersUserTagsTag) *ListUsersResponseBodyUsersUserTags {
	s.Tag = v
	return s
}

type ListUsersResponseBodyUsersUserTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// oparator
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListUsersResponseBodyUsersUserTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserTagsTag) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserTagsTag) SetTagKey(v string) *ListUsersResponseBodyUsersUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserTagsTag) SetTagValue(v string) *ListUsersResponseBodyUsersUserTagsTag {
	s.TagValue = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersForGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.``
	//
	// When you call the operation for the first time, if the total number of returned entries exceeds the value of `MaxItems`, the entries are truncated. The system returns entries based on the value of `MaxItems` and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `Marker` is returned. In the next call, you can use the value of `Marker` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) SetGroupName(v string) *ListUsersForGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMarker(v string) *ListUsersForGroupRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMaxItems(v int32) *ListUsersForGroupRequest {
	s.MaxItems = &v
	return s
}

type ListUsersForGroupResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 789FF581-B3C8-43A8-9115-54304B46D05C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) SetIsTruncated(v bool) *ListUsersForGroupResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetMarker(v string) *ListUsersForGroupResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetRequestId(v string) *ListUsersForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetUsers(v *ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody {
	s.Users = v
	return s
}

type ListUsersForGroupResponseBodyUsers struct {
	User []*ListUsersForGroupResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsers) SetUser(v []*ListUsersForGroupResponseBodyUsersUser) *ListUsersForGroupResponseBodyUsers {
	s.User = v
	return s
}

type ListUsersForGroupResponseBodyUsersUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was added to the RAM user group. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-20T06:57:00Z
	JoinDate *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUsersForGroupResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetDisplayName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetJoinDate(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.JoinDate = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserId(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersForGroupResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

type ListUsersForGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) SetHeaders(v map[string]*string) *ListUsersForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForGroupResponse) SetStatusCode(v int32) *ListUsersForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForGroupResponse) SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse {
	s.Body = v
	return s
}

type ListUsersInRecycleBinRequest struct {
	// The filter condition that is used to query information about a specified RAM user in the recycle bin.
	//
	// You must specify this parameter in the `UserPrincipalName eq <username>@<AccountAlias>.onaliyun.com` format.
	//
	// example:
	//
	// UserPrincipalName eq test@example.onaliyun.com
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The `marker`. If part of a previous response is truncated, you can use this parameter to obtain the truncated part.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries to return. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersInRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersInRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinRequest) SetFilter(v string) *ListUsersInRecycleBinRequest {
	s.Filter = &v
	return s
}

func (s *ListUsersInRecycleBinRequest) SetMarker(v string) *ListUsersInRecycleBinRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersInRecycleBinRequest) SetMaxItems(v int32) *ListUsersInRecycleBinRequest {
	s.MaxItems = &v
	return s
}

type ListUsersInRecycleBinResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The parameter that is used to obtain the truncated part. It takes effect only when `IsTruncated` is set to `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3687BD52-49FD-585B-AB14-CD05B7C76963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM users.
	Users *ListUsersInRecycleBinResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersInRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBody) SetIsTruncated(v bool) *ListUsersInRecycleBinResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetMarker(v string) *ListUsersInRecycleBinResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetRequestId(v string) *ListUsersInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBody) SetUsers(v *ListUsersInRecycleBinResponseBodyUsers) *ListUsersInRecycleBinResponseBody {
	s.Users = v
	return s
}

type ListUsersInRecycleBinResponseBodyUsers struct {
	User []*ListUsersInRecycleBinResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersInRecycleBinResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBodyUsers) SetUser(v []*ListUsersInRecycleBinResponseBodyUsersUser) *ListUsersInRecycleBinResponseBodyUsers {
	s.User = v
	return s
}

type ListUsersInRecycleBinResponseBodyUsersUser struct {
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the RAM user will be permanently deleted from the recycle bin.
	//
	// example:
	//
	// 2020-11-15T09:12:00Z
	DeleteDate *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the RAM user was deleted and moved to the recycle bin.
	//
	// example:
	//
	// 2020-10-15T09:12:00Z
	RecycleDate *string `json:"RecycleDate,omitempty" xml:"RecycleDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUsersInRecycleBinResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersInRecycleBinResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetCreateDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetDeleteDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.DeleteDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetDisplayName(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetRecycleDate(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.RecycleDate = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetUserId(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersInRecycleBinResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersInRecycleBinResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

type ListUsersInRecycleBinResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersInRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersInRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersInRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *ListUsersInRecycleBinResponse) SetHeaders(v map[string]*string) *ListUsersInRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *ListUsersInRecycleBinResponse) SetStatusCode(v int32) *ListUsersInRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersInRecycleBinResponse) SetBody(v *ListUsersInRecycleBinResponseBody) *ListUsersInRecycleBinResponse {
	s.Body = v
	return s
}

type ListVirtualMFADevicesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.``
	//
	// When you call the operation for the first time, if the total number of returned entries exceeds the value of `MaxItems`, the entries are truncated. The system returns entries based on the value of `MaxItems` and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `Marker` is returned. In the next call, you can use the value of `Marker` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of the `IsTruncated` parameter becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxItems *int32 `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesRequest) SetMarker(v string) *ListVirtualMFADevicesRequest {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesRequest) SetMaxItems(v int32) *ListVirtualMFADevicesRequest {
	s.MaxItems = &v
	return s
}

type ListVirtualMFADevicesResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 32272612-DF82-485E-8BA9-AFA4E0C3D0BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the MFA device.
	VirtualMFADevices *ListVirtualMFADevicesResponseBodyVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBody) SetIsTruncated(v bool) *ListVirtualMFADevicesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetMarker(v string) *ListVirtualMFADevicesResponseBody {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetRequestId(v string) *ListVirtualMFADevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBody) SetVirtualMFADevices(v *ListVirtualMFADevicesResponseBodyVirtualMFADevices) *ListVirtualMFADevicesResponseBody {
	s.VirtualMFADevices = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevices struct {
	VirtualMFADevice []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" type:"Repeated"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevices) SetVirtualMFADevice(v []*ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) *ListVirtualMFADevicesResponseBodyVirtualMFADevices {
	s.VirtualMFADevice = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice struct {
	// The time when the MFA device was activated.
	//
	// example:
	//
	// 2020-10-16T06:02:09Z
	ActivateDate *string `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty"`
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::177242285274****:mfa/test
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The information of the RAM user that has an MFA device bound.
	User *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetActivateDate(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.ActivateDate = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetSerialNumber(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice) SetUser(v *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADevice {
	s.User = v
	return s
}

type ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser struct {
	// The display name of the RAM user.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@177242285274****.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetDisplayName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.DisplayName = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserId(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser) SetUserPrincipalName(v string) *ListVirtualMFADevicesResponseBodyVirtualMFADevicesVirtualMFADeviceUser {
	s.UserPrincipalName = &v
	return s
}

type ListVirtualMFADevicesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualMFADevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponse) SetHeaders(v map[string]*string) *ListVirtualMFADevicesResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetStatusCode(v int32) *ListVirtualMFADevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetBody(v *ListVirtualMFADevicesResponseBody) *ListVirtualMFADevicesResponse {
	s.Body = v
	return s
}

type ProvisionApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 407426893752729****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The permissions that are granted to the application. Separate multiple permissions with a semicolon (;).
	//
	// >  For more information about the supported permissions, see [Overview](https://help.aliyun.com/document_detail/93693.html).
	//
	// example:
	//
	// openid;aliuid
	Scopes *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
}

func (s ProvisionApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationRequest) SetAppId(v string) *ProvisionApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ProvisionApplicationRequest) SetScopes(v string) *ProvisionApplicationRequest {
	s.Scopes = &v
	return s
}

type ProvisionApplicationResponseBody struct {
	// The installation information of the application.
	ApplicationProvisionInfo *ProvisionApplicationResponseBodyApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C8AE06ED-9593-5BF9-8D4A-68D5DDCC90AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProvisionApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBody) SetApplicationProvisionInfo(v *ProvisionApplicationResponseBodyApplicationProvisionInfo) *ProvisionApplicationResponseBody {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *ProvisionApplicationResponseBody) SetRequestId(v string) *ProvisionApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ProvisionApplicationResponseBodyApplicationProvisionInfo struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// GiteePrd
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The name of the application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.177242285274****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted to the application.
	DelegatedScope *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the Alibaba Cloud account for which the application was installed.
	//
	// example:
	//
	// 177242285274****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfo) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAccountId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppPrincipalName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppPrincipalName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetCreateDate(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetDelegatedScope(v *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetDisplayName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetTenantId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.TenantId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetUpdateDate(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope struct {
	// The information about the permissions that are granted to the application.
	PredefinedScopes *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ProvisionApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ProvisionApplicationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponse) SetHeaders(v map[string]*string) *ProvisionApplicationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionApplicationResponse) SetStatusCode(v int32) *ProvisionApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionApplicationResponse) SetBody(v *ProvisionApplicationResponseBody) *ProvisionApplicationResponse {
	s.Body = v
	return s
}

type ProvisionExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 403550611646604****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The permissions that are granted to the application. Separate multiple permissions with a semicolon (;).
	//
	// >  For more information about the supported permissions, see [Overview](https://help.aliyun.com/document_detail/93693.html).
	//
	// example:
	//
	// openid;aliuid
	Scopes *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
}

func (s ProvisionExternalApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationRequest) SetAppId(v string) *ProvisionExternalApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ProvisionExternalApplicationRequest) SetScopes(v string) *ProvisionExternalApplicationRequest {
	s.Scopes = &v
	return s
}

type ProvisionExternalApplicationResponseBody struct {
	// The information about the external application.
	ExternalApplication *ProvisionExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 55535873-9A6B-5C87-853F-C7CD258826F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProvisionExternalApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBody) SetExternalApplication(v *ProvisionExternalApplicationResponseBodyExternalApplication) *ProvisionExternalApplicationResponseBody {
	s.ExternalApplication = v
	return s
}

func (s *ProvisionExternalApplicationResponseBody) SetRequestId(v string) *ProvisionExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ProvisionExternalApplicationResponseBodyExternalApplication struct {
	// The name of the application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.177242285274****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the scopes of permissions that are granted to the application.
	DelegatedScope *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 403550611646604****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account to which the external application belongs.
	//
	// example:
	//
	// 157242285274****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplication) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplication) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetAppPrincipalName(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetCreateDate(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetDelegatedScope(v *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetDisplayName(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetForeignAppId(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetTenantId(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.TenantId = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetUpdateDate(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.UpdateDate = &v
	return s
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
	// The information about the scopes of permissions that are granted to the application.
	PredefinedScopes *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) SetPredefinedScopes(v *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission scope.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission scope.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ProvisionExternalApplicationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProvisionExternalApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProvisionExternalApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ProvisionExternalApplicationResponse) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponse) SetHeaders(v map[string]*string) *ProvisionExternalApplicationResponse {
	s.Headers = v
	return s
}

func (s *ProvisionExternalApplicationResponse) SetStatusCode(v int32) *ProvisionExternalApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ProvisionExternalApplicationResponse) SetBody(v *ProvisionExternalApplicationResponseBody) *ProvisionExternalApplicationResponse {
	s.Body = v
	return s
}

type RemoveClientIdFromOIDCProviderRequest struct {
	// The client ID that you want to remove.
	//
	// The client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods, (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// The client ID can be up to 128 characters in length.
	//
	// example:
	//
	// 498469743454717****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetClientId(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.ClientId = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderRequest) SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type RemoveClientIdFromOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9A8F3B-AFA5-5C8F-999D-F97BC7CF1FC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) SetOIDCProvider(v *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) *RemoveClientIdFromOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) SetRequestId(v string) *RemoveClientIdFromOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 598469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type RemoveClientIdFromOIDCProviderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveClientIdFromOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetHeaders(v map[string]*string) *RemoveClientIdFromOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetStatusCode(v int32) *RemoveClientIdFromOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponse) SetBody(v *RemoveClientIdFromOIDCProviderResponseBody) *RemoveClientIdFromOIDCProviderResponse {
	s.Body = v
	return s
}

type RemoveFingerprintFromOIDCProviderRequest struct {
	// The fingerprint that you want to remove.
	//
	// example:
	//
	// 6938fd4d98bab03faadb97b34396831e3780****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderRequest) SetFingerprint(v string) *RemoveFingerprintFromOIDCProviderRequest {
	s.Fingerprint = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderRequest) SetOIDCProviderName(v string) *RemoveFingerprintFromOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type RemoveFingerprintFromOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C90CE971-4C7F-5D93-BD3E-2D0E79D03C01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) SetOIDCProvider(v *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) *RemoveFingerprintFromOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) SetRequestId(v string) *RemoveFingerprintFromOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 0oa4u6l8x5WoaVbd****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 5938fd4d98bab03faadb97b34396831e3780****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type RemoveFingerprintFromOIDCProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFingerprintFromOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetHeaders(v map[string]*string) *RemoveFingerprintFromOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetStatusCode(v int32) *RemoveFingerprintFromOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponse) SetBody(v *RemoveFingerprintFromOIDCProviderResponseBody) *RemoveFingerprintFromOIDCProviderResponse {
	s.Body = v
	return s
}

type RemoveUserFromGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) SetGroupName(v string) *RemoveUserFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetUserPrincipalName(v string) *RemoveUserFromGroupRequest {
	s.UserPrincipalName = &v
	return s
}

type RemoveUserFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponseBody) SetRequestId(v string) *RemoveUserFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUserFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromGroupResponse) SetStatusCode(v int32) *RemoveUserFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromGroupResponse) SetBody(v *RemoveUserFromGroupResponseBody) *RemoveUserFromGroupResponse {
	s.Body = v
	return s
}

type RestoreAccessKeyFromRecycleBinRequest struct {
	// The AccessKey ID of the RAM user.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The ID of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinRequest) SetUserAccessKeyId(v string) *RestoreAccessKeyFromRecycleBinRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinRequest) SetUserId(v string) *RestoreAccessKeyFromRecycleBinRequest {
	s.UserId = &v
	return s
}

type RestoreAccessKeyFromRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinResponseBody) SetRequestId(v string) *RestoreAccessKeyFromRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type RestoreAccessKeyFromRecycleBinResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreAccessKeyFromRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreAccessKeyFromRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreAccessKeyFromRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetHeaders(v map[string]*string) *RestoreAccessKeyFromRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetStatusCode(v int32) *RestoreAccessKeyFromRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreAccessKeyFromRecycleBinResponse) SetBody(v *RestoreAccessKeyFromRecycleBinResponseBody) *RestoreAccessKeyFromRecycleBinResponse {
	s.Body = v
	return s
}

type RestoreUserFromRecycleBinRequest struct {
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RestoreUserFromRecycleBinRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreUserFromRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinRequest) SetUserId(v string) *RestoreUserFromRecycleBinRequest {
	s.UserId = &v
	return s
}

type RestoreUserFromRecycleBinResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4507D1CD-526A-4E2B-A1E2-3AB045D1EE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreUserFromRecycleBinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreUserFromRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinResponseBody) SetRequestId(v string) *RestoreUserFromRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

type RestoreUserFromRecycleBinResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreUserFromRecycleBinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreUserFromRecycleBinResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreUserFromRecycleBinResponse) GoString() string {
	return s.String()
}

func (s *RestoreUserFromRecycleBinResponse) SetHeaders(v map[string]*string) *RestoreUserFromRecycleBinResponse {
	s.Headers = v
	return s
}

func (s *RestoreUserFromRecycleBinResponse) SetStatusCode(v int32) *RestoreUserFromRecycleBinResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreUserFromRecycleBinResponse) SetBody(v *RestoreUserFromRecycleBinResponseBody) *RestoreUserFromRecycleBinResponse {
	s.Body = v
	return s
}

type SetDefaultDomainRequest struct {
	// The default domain name.
	//
	// The default domain name is in the format of `<AccountAlias>.onaliyun.com`. `<AccountAlias>` indicates the account alias. By default, the value of AccountAlias is the ID of the Alibaba Cloud account. The default domain name must end with `.onaliyun.com`.
	//
	// The default domain name can contain up to 64 characters in length. The default domain name can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// >  The default domain name cannot start or end with a hyphen (-) and cannot contain two consecutive hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// examplecompany.onaliyun.com
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
}

func (s SetDefaultDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) SetDefaultDomainName(v string) *SetDefaultDomainRequest {
	s.DefaultDomainName = &v
	return s
}

type SetDefaultDomainResponseBody struct {
	// The default domain name.
	//
	// example:
	//
	// examplecompany.onaliyun.com
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 66815255-7CCE-4759-AC37-9755794C3626
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponseBody) SetDefaultDomainName(v string) *SetDefaultDomainResponseBody {
	s.DefaultDomainName = &v
	return s
}

func (s *SetDefaultDomainResponseBody) SetRequestId(v string) *SetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponse) SetHeaders(v map[string]*string) *SetDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultDomainResponse) SetStatusCode(v int32) *SetDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultDomainResponse) SetBody(v *SetDefaultDomainResponseBody) *SetDefaultDomainResponse {
	s.Body = v
	return s
}

type SetPasswordPolicyRequest struct {
	// Specifies whether to disable logon after the password expires. Valid values:
	//
	// 	- true: After the password expires, you cannot use the password to log on to the console. You can log on to the console only after you reset the password by using your Alibaba Cloud account or as a RAM user that has administrative rights.
	//
	// 	- false: After the password expires, you can change the password to log on to the console. This is the default value.
	//
	// example:
	//
	// false
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// The maximum number of password retries. If you enter the wrong passwords for the specified consecutive times, the account is locked for one hour.
	//
	// Valid values: 0 to 32.
	//
	// The default value is 0, which indicates that the password retries are not limited.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The validity period of the password.
	//
	// Valid values: 0 to 1095. Unit: days.
	//
	// The default value is 0, which indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in the password.
	//
	// Valid values: 0 to 8.
	//
	// The default value is 0, which indicates that no limits are imposed on the number of unique characters in a password.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// The minimum number of characters in the password.
	//
	// Valid values: 8 to 32. Default value: 8.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Specifies whether to exclude the username from the password. Valid values:
	//
	// 	- true: A password cannot contain the username.
	//
	// 	- false: A password can contain the username. This is the default value.
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// The policy for password history check.
	//
	// The previous N passwords cannot be reused. Valid values of N: 0 to 24.
	//
	// The default value is 0, which indicates that RAM users can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Specifies whether the password must contain lowercase letters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Specifies whether the password must contain digits. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Specifies whether the password must contain special characters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Specifies whether the password must contain uppercase letters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) SetHardExpire(v bool) *SetPasswordPolicyRequest {
	s.HardExpire = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttemps(v int32) *SetPasswordPolicyRequest {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordDifferentCharacter(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyRequest {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireNumbers(v bool) *SetPasswordPolicyRequest {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireSymbols(v bool) *SetPasswordPolicyRequest {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireUppercaseCharacters = &v
	return s
}

type SetPasswordPolicyResponseBody struct {
	// The details of the password policy.
	PasswordPolicy *SetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3FB5551F-B2ED-40D4-8392-1E4AC2384EFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBody) SetPasswordPolicy(v *SetPasswordPolicyResponseBodyPasswordPolicy) *SetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *SetPasswordPolicyResponseBody) SetRequestId(v string) *SetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether to disable logon after the password expires.
	//
	// example:
	//
	// false
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// The maximum number of password retries.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The validity period of the password.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in the password.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// The minimum number of characters in the password.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Indicates whether to exclude the username from the password.
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// The policy for password history check.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether the password must contain lowercase letters.
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether the password must contain digits.
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether the password must contain special characters.
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether the password must contain uppercase letters.
	//
	// example:
	//
	// false
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordDifferentCharacter(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

type SetPasswordPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponse) SetHeaders(v map[string]*string) *SetPasswordPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordPolicyResponse) SetStatusCode(v int32) *SetPasswordPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordPolicyResponse) SetBody(v *SetPasswordPolicyResponseBody) *SetPasswordPolicyResponse {
	s.Body = v
	return s
}

type SetSecurityPreferenceRequest struct {
	// Specifies whether RAM users can change their passwords. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether a RAM user can use a passkey for logon. Valid values:
	//
	// 	- true: A RAM user can use a passkey for logon. This is the default value.
	//
	// 	- false: A RAM user cannot use a passkey for logon.
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
	// Specifies whether RAM users can remember the MFA devices for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). This parameter does not take effect on API calls that are authenticated by using AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// You can specify up to 40 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// Valid values: 1 to 24. Unit: hours.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console. This parameter is used to replace EnforceMFAForLogin. EnforceMFAForLogin is still valid. However, we recommend that you use MFAOperationForLogin. Valid values:
	//
	// 	- mandatory: MFA is required for all RAM users. If you use EnforceMFAForLogin, set the value to true.
	//
	// 	- independent (default): User-specific settings are applied. If you use EnforceMFAForLogin, set the value to false.
	//
	// 	- adaptive: MFA is required only for RAM users who initiated unusual logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable MFA for RAM users who initiated unusual logons. Valid values:
	//
	// 	- autonomous (default): yes. MFA is prompted for RAM users who initiated unusual logons. However, the RAM users are allowed to skip MFA.
	//
	// 	- enforceVerify: MFA is prompted for RAM users who initiated unusual logons and the RAM users cannot skip MFA.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
	// The MFA methods.
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s SetSecurityPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetMFAOperationForLogin(v string) *SetSecurityPreferenceRequest {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetOperationForRiskLogin(v string) *SetSecurityPreferenceRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetVerificationTypes(v []*string) *SetSecurityPreferenceRequest {
	s.VerificationTypes = v
	return s
}

type SetSecurityPreferenceShrinkRequest struct {
	// Specifies whether RAM users can change their passwords. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Specifies whether a RAM user can use a passkey for logon. Valid values:
	//
	// 	- true: A RAM user can use a passkey for logon. This is the default value.
	//
	// 	- false: A RAM user cannot use a passkey for logon.
	//
	// example:
	//
	// true
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Specifies whether RAM users can manage their AccessKey pairs. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	// Specifies whether RAM users can manage their MFA devices. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	// Specifies whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
	// Specifies whether RAM users can remember the MFA devices for seven days. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask that specifies the IP addresses from which you can log on to the Alibaba Cloud Management Console. This parameter takes effect on password-based logon and single sign-on (SSO). This parameter does not take effect on API calls that are authenticated by using AccessKey pairs.
	//
	// 	- If you specify a subnet mask, RAM users can use only the IP addresses in the subnet mask to log on to the Alibaba Cloud Management Console.
	//
	// 	- If you do not specify a subnet mask, RAM users can use all IP addresses to log on to the Alibaba Cloud Management Console.
	//
	// If you need to specify multiple subnet masks, separate the subnet masks with semicolons (;). Example: 192.168.0.0/16;10.0.0.0/8.
	//
	// You can specify up to 40 subnet masks. The total length of the subnet masks can be a maximum of 512 characters.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// Valid values: 1 to 24. Unit: hours.
	//
	// Default value: 6.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Specifies whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console. This parameter is used to replace EnforceMFAForLogin. EnforceMFAForLogin is still valid. However, we recommend that you use MFAOperationForLogin. Valid values:
	//
	// 	- mandatory: MFA is required for all RAM users. If you use EnforceMFAForLogin, set the value to true.
	//
	// 	- independent (default): User-specific settings are applied. If you use EnforceMFAForLogin, set the value to false.
	//
	// 	- adaptive: MFA is required only for RAM users who initiated unusual logons.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable MFA for RAM users who initiated unusual logons. Valid values:
	//
	// 	- autonomous (default): yes. MFA is prompted for RAM users who initiated unusual logons. However, the RAM users are allowed to skip MFA.
	//
	// 	- enforceVerify: MFA is prompted for RAM users who initiated unusual logons and the RAM users cannot skip MFA.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
	// The MFA methods.
	VerificationTypesShrink *string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty"`
}

func (s SetSecurityPreferenceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceShrinkRequest {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceShrinkRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceShrinkRequest {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetLoginSessionDuration(v int32) *SetSecurityPreferenceShrinkRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetMFAOperationForLogin(v string) *SetSecurityPreferenceShrinkRequest {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetOperationForRiskLogin(v string) *SetSecurityPreferenceShrinkRequest {
	s.OperationForRiskLogin = &v
	return s
}

func (s *SetSecurityPreferenceShrinkRequest) SetVerificationTypesShrink(v string) *SetSecurityPreferenceShrinkRequest {
	s.VerificationTypesShrink = &v
	return s
}

type SetSecurityPreferenceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 17494710-B4BA-4185-BBBB-C1A6ABDE1639
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of security preferences.
	SecurityPreference *SetSecurityPreferenceResponseBodySecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBody) SetRequestId(v string) *SetSecurityPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityPreferenceResponseBody) SetSecurityPreference(v *SetSecurityPreferenceResponseBodySecurityPreference) *SetSecurityPreferenceResponseBody {
	s.SecurityPreference = v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreference struct {
	// The AccessKey pair preference.
	AccessKeyPreference *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" type:"Struct"`
	// The logon preference.
	LoginProfilePreference *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" type:"Struct"`
	// The MFA preference.
	MFAPreference *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" type:"Struct"`
	// The maximum idle periods. Unit: days.
	MaxIdleDays *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays `json:"MaxIdleDays,omitempty" xml:"MaxIdleDays,omitempty" type:"Struct"`
	// The personal information preference.
	PersonalInfoPreference *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference `json:"PersonalInfoPreference,omitempty" xml:"PersonalInfoPreference,omitempty" type:"Struct"`
	// The MFA method preference.
	VerificationPreference *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference `json:"VerificationPreference,omitempty" xml:"VerificationPreference,omitempty" type:"Struct"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetAccessKeyPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetLoginProfilePreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMFAPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MFAPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetMaxIdleDays(v *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.MaxIdleDays = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetPersonalInfoPreference(v *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.PersonalInfoPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreference) SetVerificationPreference(v *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) *SetSecurityPreferenceResponseBodySecurityPreference {
	s.VerificationPreference = v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference struct {
	// Indicates whether RAM users can manage their AccessKey pairs.
	//
	// example:
	//
	// false
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference struct {
	// Indicates whether RAM users can change their passwords.
	//
	// example:
	//
	// true
	AllowUserToChangePassword *bool `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	// Indicates whether a RAM user can use a passkey for logon.
	//
	// example:
	//
	// false
	AllowUserToLoginWithPasskey *bool `json:"AllowUserToLoginWithPasskey,omitempty" xml:"AllowUserToLoginWithPasskey,omitempty"`
	// Indicates whether RAM users can remember the MFA devices for seven days.
	//
	// example:
	//
	// false
	EnableSaveMFATicket *bool `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	// The subnet mask.
	//
	// example:
	//
	// 10.0.0.0/8
	LoginNetworkMasks *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
	// The validity period of the logon session of RAM users.
	//
	// example:
	//
	// 6
	LoginSessionDuration *int32 `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	// Indicates whether MFA is required for all RAM users when they log on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// adaptive
	MFAOperationForLogin *string `json:"MFAOperationForLogin,omitempty" xml:"MFAOperationForLogin,omitempty"`
	// Deprecated
	//
	// Indicates whether to enable MFA for RAM users who initiated unusual logons.
	//
	// example:
	//
	// autonomous
	OperationForRiskLogin *string `json:"OperationForRiskLogin,omitempty" xml:"OperationForRiskLogin,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetAllowUserToLoginWithPasskey(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.AllowUserToLoginWithPasskey = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetMFAOperationForLogin(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.MFAOperationForLogin = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference) SetOperationForRiskLogin(v string) *SetSecurityPreferenceResponseBodySecurityPreferenceLoginProfilePreference {
	s.OperationForRiskLogin = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference struct {
	// Indicates whether RAM users can manage their MFA devices.
	//
	// example:
	//
	// false
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceResponseBodySecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays struct {
	// The maximum number of days that the AccessKey pair of a RAM user can stay unused. If an AccessKey pair is not used in the previous specified number of days, the AccessKey pair is automatically disabled on the next day. The default value is 730. You cannot change the value.
	//
	// example:
	//
	// 730
	MaxIdleDaysForAccessKeys *int32 `json:"MaxIdleDaysForAccessKeys,omitempty" xml:"MaxIdleDaysForAccessKeys,omitempty"`
	// The maximum number of days that a RAM user can stay idle. If a RAM user for whom console logon is enabled does not log on to the console in the previous specified number of days, console logon is automatically disabled for the RAM user on the next day. SSO is not involved. The default value is 730. You cannot change the value.
	//
	// example:
	//
	// 730
	MaxIdleDaysForUsers *int32 `json:"MaxIdleDaysForUsers,omitempty" xml:"MaxIdleDaysForUsers,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForAccessKeys(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays) SetMaxIdleDaysForUsers(v int32) *SetSecurityPreferenceResponseBodySecurityPreferenceMaxIdleDays {
	s.MaxIdleDaysForUsers = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference struct {
	// Indicates whether RAM users can manage their personal DingTalk accounts, such as binding and unbinding of the accounts.
	//
	// example:
	//
	// true
	AllowUserToManagePersonalDingTalk *bool `json:"AllowUserToManagePersonalDingTalk,omitempty" xml:"AllowUserToManagePersonalDingTalk,omitempty"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference) SetAllowUserToManagePersonalDingTalk(v bool) *SetSecurityPreferenceResponseBodySecurityPreferencePersonalInfoPreference {
	s.AllowUserToManagePersonalDingTalk = &v
	return s
}

type SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference struct {
	// The MFA methods.
	VerificationTypes []*string `json:"VerificationTypes,omitempty" xml:"VerificationTypes,omitempty" type:"Repeated"`
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference) SetVerificationTypes(v []*string) *SetSecurityPreferenceResponseBodySecurityPreferenceVerificationPreference {
	s.VerificationTypes = v
	return s
}

type SetSecurityPreferenceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSecurityPreferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponse) SetHeaders(v map[string]*string) *SetSecurityPreferenceResponse {
	s.Headers = v
	return s
}

func (s *SetSecurityPreferenceResponse) SetStatusCode(v int32) *SetSecurityPreferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSecurityPreferenceResponse) SetBody(v *SetSecurityPreferenceResponseBody) *SetSecurityPreferenceResponse {
	s.Body = v
	return s
}

type SetUserSsoSettingsRequest struct {
	// The auxiliary domain name.
	//
	// example:
	//
	// example.com
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file, which is Base64-encoded.
	//
	// The file is provided by an identity provider (IdP) that supports Security Assertion Markup Language (SAML) 2.0.
	//
	// example:
	//
	// PD94bWwgdmVy****
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Specifies whether to enable SSO for the RAM user. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// true
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Specifies whether the SAML SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the IdP for SSO must have a domain name as the suffix.
	//
	// 	- If the value of the parameter is `true`, the `<saml:NameID>` element **must*	- be in the `username@domain` format. You can set `domain` to the default domain name or the configured domain alias.
	//
	// 	- If the value of the parameter is `false`, the `<saml:NameID>` element **must*	- be in the `username` format and **cannot*	- contain the `domain` suffix.
	//
	// Set the value to the default `true`.
	//
	// example:
	//
	// true
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s SetUserSsoSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsRequest) SetAuxiliaryDomain(v string) *SetUserSsoSettingsRequest {
	s.AuxiliaryDomain = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetMetadataDocument(v string) *SetUserSsoSettingsRequest {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetSsoEnabled(v bool) *SetUserSsoSettingsRequest {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetSsoLoginWithDomain(v bool) *SetUserSsoSettingsRequest {
	s.SsoLoginWithDomain = &v
	return s
}

type SetUserSsoSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 87F2E3F6-28A0-43F3-A77F-F7760E62F61E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of user-based SSO.
	UserSsoSettings *SetUserSsoSettingsResponseBodyUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" type:"Struct"`
}

func (s SetUserSsoSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseBody) SetRequestId(v string) *SetUserSsoSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserSsoSettingsResponseBody) SetUserSsoSettings(v *SetUserSsoSettingsResponseBodyUserSsoSettings) *SetUserSsoSettingsResponseBody {
	s.UserSsoSettings = v
	return s
}

type SetUserSsoSettingsResponseBodyUserSsoSettings struct {
	// The auxiliary domain name.
	//
	// example:
	//
	// example.com
	AuxiliaryDomain *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
	// The metadata file, which is Base64-encoded.
	//
	// example:
	//
	// PD94bWwgdmVy****
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	// Indicates whether user-based SSO is enabled.
	//
	// example:
	//
	// true
	SsoEnabled *bool `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	// Indicates whether the SAML SSO requires a domain name in the `<saml:NameID>` element of the SAML response. If yes, the username specified by the IdP for SSO must have a domain name as the suffix.
	//
	// 	- If the value of the parameter is `true`, the `<saml:NameID>` element **must*	- be in the `username@domain` format. You can set `domain` to the default domain name or the configured domain alias.
	//
	// 	- If the value of the parameter is `false`, the `<saml:NameID>` element **must*	- be in the `username` format and **cannot*	- contain the `domain` suffix.
	//
	// The default value is `true`.
	//
	// example:
	//
	// true
	SsoLoginWithDomain *bool `json:"SsoLoginWithDomain,omitempty" xml:"SsoLoginWithDomain,omitempty"`
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsResponseBodyUserSsoSettings) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetAuxiliaryDomain(v string) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetMetadataDocument(v string) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoEnabled(v bool) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsResponseBodyUserSsoSettings) SetSsoLoginWithDomain(v bool) *SetUserSsoSettingsResponseBodyUserSsoSettings {
	s.SsoLoginWithDomain = &v
	return s
}

type SetUserSsoSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserSsoSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserSsoSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponse) SetHeaders(v map[string]*string) *SetUserSsoSettingsResponse {
	s.Headers = v
	return s
}

func (s *SetUserSsoSettingsResponse) SetStatusCode(v int32) *SetUserSsoSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserSsoSettingsResponse) SetBody(v *SetUserSsoSettingsResponseBody) *SetUserSsoSettingsResponse {
	s.Body = v
	return s
}

type SetVerificationInfoRequest struct {
	// The email address.
	//
	// >  If you set `VerifyType` to `email`, you must specify this parameter.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number.
	//
	// >  If you set `VerifyType` to `sms`, you must specify this parameter.
	//
	// example:
	//
	// 86-13900001234
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	// The multi-factor authentication (MFA) method. Valid values:
	//
	// 	- sms: mobile phone.
	//
	// 	- email: email.
	//
	// example:
	//
	// sms
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SetVerificationInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoRequest) SetEmail(v string) *SetVerificationInfoRequest {
	s.Email = &v
	return s
}

func (s *SetVerificationInfoRequest) SetMobilePhone(v string) *SetVerificationInfoRequest {
	s.MobilePhone = &v
	return s
}

func (s *SetVerificationInfoRequest) SetUserPrincipalName(v string) *SetVerificationInfoRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *SetVerificationInfoRequest) SetVerifyType(v string) *SetVerificationInfoRequest {
	s.VerifyType = &v
	return s
}

type SetVerificationInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVerificationInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVerificationInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoResponseBody) SetRequestId(v string) *SetVerificationInfoResponseBody {
	s.RequestId = &v
	return s
}

type SetVerificationInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetVerificationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetVerificationInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVerificationInfoResponse) GoString() string {
	return s.String()
}

func (s *SetVerificationInfoResponse) SetHeaders(v map[string]*string) *SetVerificationInfoResponse {
	s.Headers = v
	return s
}

func (s *SetVerificationInfoResponse) SetStatusCode(v int32) *SetVerificationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SetVerificationInfoResponse) SetBody(v *SetVerificationInfoResponseBody) *SetVerificationInfoResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of resource N.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource ID is the ID of the RAM user.
	//
	// >  You must specify only one of the following parameters: `ResourceId` and `ResourcePrincipalName`.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The name of resource N.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource name is the name of the RAM user.
	//
	// >  You must specify only one of the following parameters: `ResourceId` and `ResourcePrincipalName`.
	//
	// example:
	//
	// TagResources
	ResourcePrincipalName []*string `json:"ResourcePrincipalName,omitempty" xml:"ResourcePrincipalName,omitempty" type:"Repeated"`
	// The type of the resource. Valid value:
	//
	// 	- user: a Resource Access Management (RAM) user.
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. You cannot specify empty strings as tag keys. The tag key can be up to 128 characters in length. The tag key cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourcePrincipalName(v []*string) *TagResourcesRequest {
	s.ResourcePrincipalName = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N.
	//
	// Valid values of N: 1 to 20. You cannot specify empty strings as tag keys. The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3687BD52-49FD-585B-AB14-CD05B7C76963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnbindMFADeviceRequest struct {
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UnbindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceRequest) SetUserPrincipalName(v string) *UnbindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

type UnbindMFADeviceResponseBody struct {
	// The information about the MFA device.
	MFADevice *UnbindMFADeviceResponseBodyMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A26CB3E9-1021-452A-AC57-3134B3BA0E4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindMFADeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBody) SetMFADevice(v *UnbindMFADeviceResponseBodyMFADevice) *UnbindMFADeviceResponseBody {
	s.MFADevice = v
	return s
}

func (s *UnbindMFADeviceResponseBody) SetRequestId(v string) *UnbindMFADeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnbindMFADeviceResponseBodyMFADevice struct {
	// The serial number of the MFA device.
	//
	// example:
	//
	// acs:ram::151298381312****:mfa/device001
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnbindMFADeviceResponseBodyMFADevice) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponseBodyMFADevice) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseBodyMFADevice) SetSerialNumber(v string) *UnbindMFADeviceResponseBodyMFADevice {
	s.SerialNumber = &v
	return s
}

type UnbindMFADeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponse) SetHeaders(v map[string]*string) *UnbindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindMFADeviceResponse) SetStatusCode(v int32) *UnbindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindMFADeviceResponse) SetBody(v *UnbindMFADeviceResponseBody) *UnbindMFADeviceResponse {
	s.Body = v
	return s
}

type UnbindVerificationRequest struct {
	// The email address.
	//
	// >  If you set `VerifyType` to `email`, you must specify this parameter.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number.
	//
	// >  If you set `VerifyType` to `sms`, you must specify this parameter.
	//
	// example:
	//
	// 86-13900001234
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	// The multi-factor authentication (MFA) method. Valid values:
	//
	// 	- sms: mobile phone.
	//
	// 	- email: email.
	//
	// example:
	//
	// sms
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s UnbindVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindVerificationRequest) GoString() string {
	return s.String()
}

func (s *UnbindVerificationRequest) SetEmail(v string) *UnbindVerificationRequest {
	s.Email = &v
	return s
}

func (s *UnbindVerificationRequest) SetMobilePhone(v string) *UnbindVerificationRequest {
	s.MobilePhone = &v
	return s
}

func (s *UnbindVerificationRequest) SetUserPrincipalName(v string) *UnbindVerificationRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UnbindVerificationRequest) SetVerifyType(v string) *UnbindVerificationRequest {
	s.VerifyType = &v
	return s
}

type UnbindVerificationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindVerificationResponseBody) SetRequestId(v string) *UnbindVerificationResponseBody {
	s.RequestId = &v
	return s
}

type UnbindVerificationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindVerificationResponse) GoString() string {
	return s.String()
}

func (s *UnbindVerificationResponse) SetHeaders(v map[string]*string) *UnbindVerificationResponse {
	s.Headers = v
	return s
}

func (s *UnbindVerificationResponse) SetStatusCode(v int32) *UnbindVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindVerificationResponse) SetBody(v *UnbindVerificationResponseBody) *UnbindVerificationResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- true: remove all tags from the resources.
	//
	// 	- false (default): does not remove all tags from the resources.
	//
	// > This parameter takes effect only when TagKey.N is not set in the request.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The IDs of resources.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource ID is the ID of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	//
	// example:
	//
	// UntagResources
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The names of resources.
	//
	// Valid values of N: 1 to 50. If the ResourceType parameter is set to user, the resource name is the name of the RAM user.
	//
	// > You must specify only one of the following parameters: ResourceId and ResourcePrincipalName.
	ResourcePrincipalName []*string `json:"ResourcePrincipalName,omitempty" xml:"ResourcePrincipalName,omitempty" type:"Repeated"`
	// The type of the resource. Valid value:
	//
	// 	- user: a RAM user
	//
	// example:
	//
	// user
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of resources.
	//
	// Valid values of N: 1 to 20. N must be consecutive.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourcePrincipalName(v []*string) *UntagResourcesRequest {
	s.ResourcePrincipalName = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6AC79B74-9A78-58E2-818B-1D38EDC64D67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAccessKeyRequest struct {
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// This parameter is required.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The AccessKey ID of the AccessKey pair for which you want to modify the status.
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAI*******************
	UserAccessKeyId *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
	// The logon name of the RAM user.
	//
	// If this parameter is empty, the status of the AccessKey pair for the current user is modified.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyRequest) SetStatus(v string) *UpdateAccessKeyRequest {
	s.Status = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserAccessKeyId(v string) *UpdateAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserPrincipalName(v string) *UpdateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

type UpdateAccessKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B9AF80E4-1565-42D9-9256-0B8B0D9FD3EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccessKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponseBody) SetRequestId(v string) *UpdateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAccessKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAccessKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponse) SetHeaders(v map[string]*string) *UpdateAccessKeyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAccessKeyResponse) SetStatusCode(v int32) *UpdateAccessKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAccessKeyResponse) SetBody(v *UpdateAccessKeyResponseBody) *UpdateAccessKeyResponse {
	s.Body = v
	return s
}

type UpdateApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The validity period of the access token.
	//
	// Valid values: 900 to 10800. Unit: seconds.
	//
	// example:
	//
	// 3600
	NewAccessTokenValidity *int32 `json:"NewAccessTokenValidity,omitempty" xml:"NewAccessTokenValidity,omitempty"`
	// The display name.
	//
	// example:
	//
	// NewApp
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// Specifies whether the application can be installed by using other Alibaba Cloud accounts. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	NewIsMultiTenant *bool `json:"NewIsMultiTenant,omitempty" xml:"NewIsMultiTenant,omitempty"`
	// The permission that is granted on the application.
	//
	// For more information about the application permission scope, see [OAuth scopes](https://help.aliyun.com/document_detail/93693.html). You can also call the [ListPredefinedScopes](https://help.aliyun.com/document_detail/187206.html) operation to query the permissions that are supported by different types of applications.
	//
	// If you enter multiple permissions, separate them with semicolons (;).
	//
	// The new value of this parameter overwrites the original value, and the permission specified by the new value takes effect. For example, if the original value is `/acs/ccc`, and the new value is `/acs/alidns`, `/acs/alidns` takes effect. If you want to retain the original permission and the `/acs/alidns` permission, set the value to `/acs/ccc;/acs/alidns`.
	//
	// example:
	//
	// openid
	NewPredefinedScopes *string `json:"NewPredefinedScopes,omitempty" xml:"NewPredefinedScopes,omitempty"`
	// The callback URL.
	//
	// If you enter multiple callback URLs, separate them with semicolons (;).
	//
	// example:
	//
	// https://www.example.com
	NewRedirectUris *string `json:"NewRedirectUris,omitempty" xml:"NewRedirectUris,omitempty"`
	// The validity period of the refresh token.
	//
	// Valid values: 7200 to 31536000. Unit: seconds.
	//
	// example:
	//
	// 7776000
	NewRefreshTokenValidity *int32 `json:"NewRefreshTokenValidity,omitempty" xml:"NewRefreshTokenValidity,omitempty"`
	// The required permission.
	//
	// You can specify one or more permissions for the `RequiredScopes` parameter. After you specify this parameter, the required permissions are automatically selected and cannot be revoked when a user grants permissions on the application.
	//
	// If you also specify the `NewPredefinedScopes` parameter, the `NewPredefinedScopes` parameter specifies the permissions that can be granted on the application, and this parameter specifies the required permissions.
	//
	// If you enter multiple permissions, separate them with semicolons (;).
	//
	// The new value of this parameter overwrites the original value, and the required permission specified by the new value takes effect.
	//
	// >  If the permission that you specify for the `RequiredScopes` parameter is not included in value of the `PredefinedScopes` parameter, the permission does not take effect.
	//
	// example:
	//
	// profile;aliuid
	NewRequiredScopes *string `json:"NewRequiredScopes,omitempty" xml:"NewRequiredScopes,omitempty"`
	// Specifies whether a secret is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >
	//
	// 	- For applications of the WebApp and ServerApp types, this parameter is automatically set to true and cannot be changed.
	//
	// 	- For applications of the NativeApp type, this parameter can be set to true or false. If you do not set this parameter, false is used. Applications of the NativeApp type run in untrusted environments and the secrets of these applications are not protected. Therefore, we recommend that you do not set this parameter to true unless otherwise specified. For more information, see [Use an application of the NativeApp type to log on to Alibaba Cloud](https://help.aliyun.com/document_detail/93697.html).
	//
	// example:
	//
	// true
	NewSecretRequired *bool `json:"NewSecretRequired,omitempty" xml:"NewSecretRequired,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) SetAppId(v string) *UpdateApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewAccessTokenValidity(v int32) *UpdateApplicationRequest {
	s.NewAccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewDisplayName(v string) *UpdateApplicationRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewIsMultiTenant(v bool) *UpdateApplicationRequest {
	s.NewIsMultiTenant = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewPredefinedScopes(v string) *UpdateApplicationRequest {
	s.NewPredefinedScopes = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRedirectUris(v string) *UpdateApplicationRequest {
	s.NewRedirectUris = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRefreshTokenValidity(v int32) *UpdateApplicationRequest {
	s.NewRefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRequiredScopes(v string) *UpdateApplicationRequest {
	s.NewRequiredScopes = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewSecretRequired(v bool) *UpdateApplicationRequest {
	s.NewSecretRequired = &v
	return s
}

type UpdateApplicationResponseBody struct {
	// The information about the application.
	Application *UpdateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationResponseBodyApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *UpdateApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// NewApp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant *bool `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	// The callback URLs.
	RedirectUris *UpdateApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *UpdateApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAccountId(v string) *UpdateApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppId(v string) *UpdateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppName(v string) *UpdateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppType(v string) *UpdateApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetCreateDate(v string) *UpdateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDelegatedScope(v *UpdateApplicationResponseBodyApplicationDelegatedScope) *UpdateApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDisplayName(v string) *UpdateApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *UpdateApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetRedirectUris(v *UpdateApplicationResponseBodyApplicationRedirectUris) *UpdateApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *UpdateApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetSecretRequired(v bool) *UpdateApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdateDate(v string) *UpdateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

type UpdateApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *UpdateApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

type UpdateApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s UpdateApplicationResponseBodyApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *UpdateApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type UpdateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) SetHeaders(v map[string]*string) *UpdateApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationResponse) SetStatusCode(v int32) *UpdateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationResponse) SetBody(v *UpdateApplicationResponseBody) *UpdateApplicationResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// Dev-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The new description.
	//
	// The description can be up to 128 characters in length.
	//
	// example:
	//
	// Test-Team
	NewComments *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	// The new display name of the RAM user group.
	//
	// The name can be up to 24 characters in length.
	//
	// example:
	//
	// Test-Team
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new name of the RAM user group.
	//
	// The name can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Test-Team
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewComments(v string) *UpdateGroupRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateGroupRequest) SetNewDisplayName(v string) *UpdateGroupRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

type UpdateGroupResponseBody struct {
	// The information about the RAM user group.
	Group *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CDA656E3-3CE9-4A03-A8A3-B42A0C3C3287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody {
	s.Group = v
	return s
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupResponseBodyGroup struct {
	// The description.
	//
	// example:
	//
	// Test-Team
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-19T16:15:17Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the RAM user group.
	//
	// example:
	//
	// 740317625433843****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the RAM user group.
	//
	// example:
	//
	// Test-Team
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-20T03:44:27Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyGroup) SetComments(v string) *UpdateGroupResponseBodyGroup {
	s.Comments = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetCreateDate(v string) *UpdateGroupResponseBodyGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetDisplayName(v string) *UpdateGroupResponseBodyGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupId(v string) *UpdateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupName(v string) *UpdateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetUpdateDate(v string) *UpdateGroupResponseBodyGroup {
	s.UpdateDate = &v
	return s
}

type UpdateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateLoginProfileRequest struct {
	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	//
	// 	- true: forcefully enables MFA for the RAM user. The RAM user must bind an MFA device upon the next logon.
	//
	// 	- false: does not forcefully enable MFA for the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The new password that is used to log on to the console.
	//
	// The new password must meet the complexity requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user is required to reset the password upon the next logon. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Specifies whether to enable password-based logons to the console. Valid values:
	//
	// 	- Active: enables password-based logons to the console.
	//
	// 	- Inactive: disables password-based logons to the console.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileRequest) SetMFABindRequired(v bool) *UpdateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPassword(v string) *UpdateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPasswordResetRequired(v bool) *UpdateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetStatus(v string) *UpdateLoginProfileRequest {
	s.Status = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetUserPrincipalName(v string) *UpdateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type UpdateLoginProfileResponseBody struct {
	// The console logon configurations.
	LoginProfile *UpdateLoginProfileResponseBodyLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BCDB6A7F-2199-41D9-B577-4FA536A5ADE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLoginProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBody) SetLoginProfile(v *UpdateLoginProfileResponseBodyLoginProfile) *UpdateLoginProfileResponseBody {
	s.LoginProfile = v
	return s
}

func (s *UpdateLoginProfileResponseBody) SetRequestId(v string) *UpdateLoginProfileResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLoginProfileResponseBodyLoginProfile struct {
	// Indicates whether console logon is automatically disabled if a RAM user does not log on to the console in the previous specified number of days. The number of days is specified by MaxIdleDaysForUsers. The default value is true, and you cannot change the value.
	//
	// example:
	//
	// true
	AutoDisableLoginStatus *string `json:"AutoDisableLoginStatus,omitempty" xml:"AutoDisableLoginStatus,omitempty"`
	// Indicates whether MFA must be enabled.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// Indicates whether the RAM user is required to reset the password upon the next logon.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Indicates whether to enable password-based logons to the console.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2020-10-14T07:48:41Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// test@example11.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateLoginProfileResponseBodyLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponseBodyLoginProfile) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetAutoDisableLoginStatus(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.AutoDisableLoginStatus = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetMFABindRequired(v bool) *UpdateLoginProfileResponseBodyLoginProfile {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetPasswordResetRequired(v bool) *UpdateLoginProfileResponseBodyLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetStatus(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.Status = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetUpdateDate(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *UpdateLoginProfileResponseBodyLoginProfile) SetUserPrincipalName(v string) *UpdateLoginProfileResponseBodyLoginProfile {
	s.UserPrincipalName = &v
	return s
}

type UpdateLoginProfileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoginProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponse) SetHeaders(v map[string]*string) *UpdateLoginProfileResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoginProfileResponse) SetStatusCode(v int32) *UpdateLoginProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoginProfileResponse) SetBody(v *UpdateLoginProfileResponseBody) *UpdateLoginProfileResponse {
	s.Body = v
	return s
}

type UpdateOIDCProviderRequest struct {
	// The ID of the client. If you want to specify multiple client IDs, separate the client IDs with commas (,).
	//
	// A client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// A client ID can be up to 128 characters in length.
	//
	// >  If you specify this parameter, all the client IDs of the OIDC IdP are replaced. If you need to only add or remove a client ID, call the AddClientIdToOIDCProvider or RemoveClientIdFromOIDCProvider operation. For more information, see [AddClientIdToOIDCProvider](https://help.aliyun.com/document_detail/332057.html) or [RemoveClientIdFromOIDCProvider](https://help.aliyun.com/document_detail/332058.html).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The description of the OIDC IdP.
	//
	// The description can be up to 256 characters in length.
	//
	// example:
	//
	// This is a new OIDC Provider.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s UpdateOIDCProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderRequest) SetClientIds(v string) *UpdateOIDCProviderRequest {
	s.ClientIds = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetIssuanceLimitTime(v int64) *UpdateOIDCProviderRequest {
	s.IssuanceLimitTime = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetNewDescription(v string) *UpdateOIDCProviderRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateOIDCProviderRequest) SetOIDCProviderName(v string) *UpdateOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

type UpdateOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *UpdateOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4C4D1BD-2558-5BD1-8C26-A5D7FB174A55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOIDCProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponseBody) SetOIDCProvider(v *UpdateOIDCProviderResponseBodyOIDCProvider) *UpdateOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *UpdateOIDCProviderResponseBody) SetRequestId(v string) *UpdateOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateOIDCProviderResponseBodyOIDCProvider) String() string {
	return tea.Prettify(s)
}

func (s UpdateOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

type UpdateOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOIDCProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponse) SetHeaders(v map[string]*string) *UpdateOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateOIDCProviderResponse) SetStatusCode(v int32) *UpdateOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOIDCProviderResponse) SetBody(v *UpdateOIDCProviderResponseBody) *UpdateOIDCProviderResponse {
	s.Body = v
	return s
}

type UpdatePasskeyRequest struct {
	// The ID of the passkey.
	//
	// example:
	//
	// PASSKEY-CuZjEHhWcr7GIQOMGvkS
	PasskeyId *string `json:"PasskeyId,omitempty" xml:"PasskeyId,omitempty"`
	// The name of the passkey.
	//
	// example:
	//
	// device1
	PasskeyName *string `json:"PasskeyName,omitempty" xml:"PasskeyName,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdatePasskeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePasskeyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyRequest) SetPasskeyId(v string) *UpdatePasskeyRequest {
	s.PasskeyId = &v
	return s
}

func (s *UpdatePasskeyRequest) SetPasskeyName(v string) *UpdatePasskeyRequest {
	s.PasskeyName = &v
	return s
}

func (s *UpdatePasskeyRequest) SetUserPrincipalName(v string) *UpdatePasskeyRequest {
	s.UserPrincipalName = &v
	return s
}

type UpdatePasskeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85836703-8D4F-485F-9726-4D1C730F957E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePasskeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePasskeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyResponseBody) SetRequestId(v string) *UpdatePasskeyResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePasskeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePasskeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePasskeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePasskeyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePasskeyResponse) SetHeaders(v map[string]*string) *UpdatePasskeyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePasskeyResponse) SetStatusCode(v int32) *UpdatePasskeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePasskeyResponse) SetBody(v *UpdatePasskeyResponseBody) *UpdatePasskeyResponse {
	s.Body = v
	return s
}

type UpdateSAMLProviderRequest struct {
	// The new description.
	//
	// >  You must specify at least one of the `NewDescription` and `NewEncodedSAMLMetadataDocument` parameters.
	//
	// example:
	//
	// This is a new provider.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new metadata file.
	//
	// >  You must specify at least one of the `NewDescription` and `NewEncodedSAMLMetadataDocument` parameters.
	//
	// example:
	//
	// PD94bWwgdmVy****
	NewEncodedSAMLMetadataDocument *string `json:"NewEncodedSAMLMetadataDocument,omitempty" xml:"NewEncodedSAMLMetadataDocument,omitempty"`
	// The name of the IdP whose information you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s UpdateSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderRequest) SetNewDescription(v string) *UpdateSAMLProviderRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetNewEncodedSAMLMetadataDocument(v string) *UpdateSAMLProviderRequest {
	s.NewEncodedSAMLMetadataDocument = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetSAMLProviderName(v string) *UpdateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type UpdateSAMLProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E5EDDFD2-3654-4F9F-9780-4AE7D81823EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the IdP.
	SAMLProvider *UpdateSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s UpdateSAMLProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponseBody) SetRequestId(v string) *UpdateSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSAMLProviderResponseBody) SetSAMLProvider(v *UpdateSAMLProviderResponseBodySAMLProvider) *UpdateSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

type UpdateSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:saml-provider/test-provider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The point in time at which the IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:37:05Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the IdP.
	//
	// example:
	//
	// This is a new provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the IdP.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The point in time at which the information about the IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-22T02:51:20Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateSAMLProviderResponseBodySAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetArn(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *UpdateSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *UpdateSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

type UpdateSAMLProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSAMLProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponse) SetHeaders(v map[string]*string) *UpdateSAMLProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateSAMLProviderResponse) SetStatusCode(v int32) *UpdateSAMLProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSAMLProviderResponse) SetBody(v *UpdateSAMLProviderResponseBody) *UpdateSAMLProviderResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The new description of the RAM user.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	NewComments *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	// The new display name of the RAM user.
	//
	// The name must be 1 to 24 characters in length.
	//
	// example:
	//
	// new
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new mobile phone number of the RAM user.
	//
	// Format: \\<Country code>-\\<Mobile phone number>.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	NewMobilePhone *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	// The new logon name of the RAM user.
	//
	// The name is in the format of `<username>@<AccountAlias>.onaliyun.com`. `<username>` indicates the name of the RAM user. `<AccountAlias>.onaliyun.com` indicates the default domain name.
	//
	// The value of `NewUserPrincipalName` must be `1 to 128` characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The value of `<username>` must be `1 to 64` characters in length.
	//
	// example:
	//
	// new@example.onaliyun.com
	NewUserPrincipalName *string `json:"NewUserPrincipalName,omitempty" xml:"NewUserPrincipalName,omitempty"`
	// The ID of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetNewComments(v string) *UpdateUserRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateUserRequest) SetNewDisplayName(v string) *UpdateUserRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetNewEmail(v string) *UpdateUserRequest {
	s.NewEmail = &v
	return s
}

func (s *UpdateUserRequest) SetNewMobilePhone(v string) *UpdateUserRequest {
	s.NewMobilePhone = &v
	return s
}

func (s *UpdateUserRequest) SetNewUserPrincipalName(v string) *UpdateUserRequest {
	s.NewUserPrincipalName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUserPrincipalName(v string) *UpdateUserRequest {
	s.UserPrincipalName = &v
	return s
}

type UpdateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B56DD42-6962-4F89-A19C-079EED1F0FE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *UpdateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetUser(v *UpdateUserResponseBodyUser) *UpdateUserResponseBody {
	s.User = v
	return s
}

type UpdateUserResponseBodyUser struct {
	// The description.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// new
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The last time when the RAM user logged on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The source of the RAM user. Valid values:
	//
	// 	- Manual: The RAM user is manually created in the RAM console.
	//
	// 	- SCIM: The RAM user is mapped by using System for Cross-domain Identity Management (SCIM).
	//
	// 	- CloudSSO: The RAM user is mapped from a CloudSSO user.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the RAM user was updated.
	//
	// example:
	//
	// 2020-10-13T09:19:49Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// new@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyUser) SetComments(v string) *UpdateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetCreateDate(v string) *UpdateUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetDisplayName(v string) *UpdateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetEmail(v string) *UpdateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetLastLoginDate(v string) *UpdateUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetMobilePhone(v string) *UpdateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetProvisionType(v string) *UpdateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUpdateDate(v string) *UpdateUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserId(v string) *UpdateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserPrincipalName(v string) *UpdateUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("ims"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a client ID to an OpenID Connect (OIDC) identity provider (IdP).
//
// @param request - AddClientIdToOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddClientIdToOIDCProviderResponse
func (client *Client) AddClientIdToOIDCProviderWithOptions(request *AddClientIdToOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *AddClientIdToOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddClientIdToOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddClientIdToOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a client ID to an OpenID Connect (OIDC) identity provider (IdP).
//
// @param request - AddClientIdToOIDCProviderRequest
//
// @return AddClientIdToOIDCProviderResponse
func (client *Client) AddClientIdToOIDCProvider(request *AddClientIdToOIDCProviderRequest) (_result *AddClientIdToOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddClientIdToOIDCProviderResponse{}
	_body, _err := client.AddClientIdToOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a fingerprint to an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to add the fingerprint `902ef2deeb3c5b13ea4c3d5193629309e231****` to the OIDC IdP named `TestOIDCProvider`.
//
// @param request - AddFingerprintToOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFingerprintToOIDCProviderResponse
func (client *Client) AddFingerprintToOIDCProviderWithOptions(request *AddFingerprintToOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *AddFingerprintToOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fingerprint)) {
		query["Fingerprint"] = request.Fingerprint
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFingerprintToOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFingerprintToOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a fingerprint to an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to add the fingerprint `902ef2deeb3c5b13ea4c3d5193629309e231****` to the OIDC IdP named `TestOIDCProvider`.
//
// @param request - AddFingerprintToOIDCProviderRequest
//
// @return AddFingerprintToOIDCProviderResponse
func (client *Client) AddFingerprintToOIDCProvider(request *AddFingerprintToOIDCProviderRequest) (_result *AddFingerprintToOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddFingerprintToOIDCProviderResponse{}
	_body, _err := client.AddFingerprintToOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Resource Access Management (RAM) user to a RAM user group.
//
// @param request - AddUserToGroupRequest
//
// @return AddUserToGroupResponse
func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (_result *AddUserToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.AddUserToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADeviceWithOptions(request *BindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *BindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode1)) {
		query["AuthenticationCode1"] = request.AuthenticationCode1
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode2)) {
		query["AuthenticationCode2"] = request.AuthenticationCode2
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BindMFADevice"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a multi-factor authentication (MFA) device to a Resource Access Management (RAM) user.
//
// @param request - BindMFADeviceRequest
//
// @return BindMFADeviceResponse
func (client *Client) BindMFADevice(request *BindMFADeviceRequest) (_result *BindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.BindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangePasswordResponse
func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *util.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangePassword"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the password that is used to log on to the console for a Resource Access Management (RAM) user.
//
// Description:
//
// >  This operation is available only for RAM users. Before you call this operation, make sure that `AllowUserToChangePassword` in [SetSecurityPreference](https://help.aliyun.com/document_detail/43765.html) is set to `True`. The value True indicates that RAM users can manage their passwords.
//
// @param request - ChangePasswordRequest
//
// @return ChangePasswordResponse
func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKeyWithOptions(request *CreateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAccessKey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - CreateAccessKeyRequest
//
// @return CreateAccessKeyResponse
func (client *Client) CreateAccessKey(request *CreateAccessKeyRequest) (_result *CreateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CreateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application secret for an application.
//
// @param request - CreateAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSecretResponse
func (client *Client) CreateAppSecretWithOptions(request *CreateAppSecretRequest, runtime *util.RuntimeOptions) (_result *CreateAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSecret"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application secret for an application.
//
// @param request - CreateAppSecretRequest
//
// @return CreateAppSecretResponse
func (client *Client) CreateAppSecret(request *CreateAppSecretRequest) (_result *CreateAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.CreateAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessTokenValidity)) {
		query["AccessTokenValidity"] = request.AccessTokenValidity
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.IsMultiTenant)) {
		query["IsMultiTenant"] = request.IsMultiTenant
	}

	if !tea.BoolValue(util.IsUnset(request.PredefinedScopes)) {
		query["PredefinedScopes"] = request.PredefinedScopes
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUris)) {
		query["RedirectUris"] = request.RedirectUris
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshTokenValidity)) {
		query["RefreshTokenValidity"] = request.RefreshTokenValidity
	}

	if !tea.BoolValue(util.IsUnset(request.RequiredScopes)) {
		query["RequiredScopes"] = request.RequiredScopes
	}

	if !tea.BoolValue(util.IsUnset(request.SecretRequired)) {
		query["SecretRequired"] = request.SecretRequired
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user group.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Resource Access Management (RAM) user group.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfileWithOptions(request *CreateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *CreateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MFABindRequired)) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordResetRequired)) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoginProfile"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - CreateLoginProfileRequest
//
// @return CreateLoginProfileResponse
func (client *Client) CreateLoginProfile(request *CreateLoginProfileRequest) (_result *CreateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CreateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an OpenID Connect (OIDC) identity provider (IdP) to configure a trust relationship between Alibaba Cloud and an external IdP. This topic provides an example on how to create an IdP named TestOIDCProvider to configure a trust relationship between the external IdP Okta and Alibaba Cloud.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the information such as the URL of the issuer, the fingerprints of HTTPS certificate authority (CA) certificates, and the client IDs are obtained from an external IdP, such as Google Workspace or Okta.
//
// ### [](#)Limits
//
//   - You can create a maximum of 100 OIDC IdPs in an Alibaba Cloud account.
//
//   - You can add a maximum of 50 client IDs to an OIDC IdP.
//
//   - You can add a maximum of five fingerprints to an OIDC IdP.
//
// ### [](#)Operation description
//
// This topic provides an example on how to create an IdP named `TestOIDCProvider` to configure a trust relationship between the external IdP and Alibaba Cloud.
//
// @param request - CreateOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOIDCProviderResponse
func (client *Client) CreateOIDCProviderWithOptions(request *CreateOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *CreateOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIds)) {
		query["ClientIds"] = request.ClientIds
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Fingerprints)) {
		query["Fingerprints"] = request.Fingerprints
	}

	if !tea.BoolValue(util.IsUnset(request.IssuanceLimitTime)) {
		query["IssuanceLimitTime"] = request.IssuanceLimitTime
	}

	if !tea.BoolValue(util.IsUnset(request.IssuerUrl)) {
		query["IssuerUrl"] = request.IssuerUrl
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an OpenID Connect (OIDC) identity provider (IdP) to configure a trust relationship between Alibaba Cloud and an external IdP. This topic provides an example on how to create an IdP named TestOIDCProvider to configure a trust relationship between the external IdP Okta and Alibaba Cloud.
//
// Description:
//
// ### [](#)Prerequisites
//
// Before you call this operation, make sure that the information such as the URL of the issuer, the fingerprints of HTTPS certificate authority (CA) certificates, and the client IDs are obtained from an external IdP, such as Google Workspace or Okta.
//
// ### [](#)Limits
//
//   - You can create a maximum of 100 OIDC IdPs in an Alibaba Cloud account.
//
//   - You can add a maximum of 50 client IDs to an OIDC IdP.
//
//   - You can add a maximum of five fingerprints to an OIDC IdP.
//
// ### [](#)Operation description
//
// This topic provides an example on how to create an IdP named `TestOIDCProvider` to configure a trust relationship between the external IdP and Alibaba Cloud.
//
// @param request - CreateOIDCProviderRequest
//
// @return CreateOIDCProviderResponse
func (client *Client) CreateOIDCProvider(request *CreateOIDCProviderRequest) (_result *CreateOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOIDCProviderResponse{}
	_body, _err := client.CreateOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - CreateSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSAMLProviderResponse
func (client *Client) CreateSAMLProviderWithOptions(request *CreateSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *CreateSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncodedSAMLMetadataDocument)) {
		query["EncodedSAMLMetadataDocument"] = request.EncodedSAMLMetadataDocument
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLProviderName)) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSAMLProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - CreateSAMLProviderRequest
//
// @return CreateSAMLProviderResponse
func (client *Client) CreateSAMLProvider(request *CreateSAMLProviderRequest) (_result *CreateSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.CreateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a RAM user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `test`.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comments)) {
		query["Comments"] = request.Comments
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhone)) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a RAM user.
//
// Description:
//
// This topic provides an example on how to create a RAM user named `test`.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADeviceWithOptions(request *CreateVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VirtualMFADeviceName)) {
		query["VirtualMFADeviceName"] = request.VirtualMFADeviceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVirtualMFADevice"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual multi-factor authentication (MFA) device.
//
// @param request - CreateVirtualMFADeviceRequest
//
// @return CreateVirtualMFADeviceResponse
func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (_result *CreateVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CreateVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKeyWithOptions(request *DeleteAccessKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessKey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - DeleteAccessKeyRequest
//
// @return DeleteAccessKeyResponse
func (client *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (_result *DeleteAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.DeleteAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteAccessKeyInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessKeyInRecycleBinResponse
func (client *Client) DeleteAccessKeyInRecycleBinWithOptions(request *DeleteAccessKeyInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessKeyInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAccessKeyInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAccessKeyInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteAccessKeyInRecycleBinRequest
//
// @return DeleteAccessKeyInRecycleBinResponse
func (client *Client) DeleteAccessKeyInRecycleBin(request *DeleteAccessKeyInRecycleBinRequest) (_result *DeleteAccessKeyInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessKeyInRecycleBinResponse{}
	_body, _err := client.DeleteAccessKeyInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the application secret of an application.
//
// @param request - DeleteAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppSecretResponse
func (client *Client) DeleteAppSecretWithOptions(request *DeleteAppSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecretId)) {
		query["AppSecretId"] = request.AppSecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppSecret"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the application secret of an application.
//
// @param request - DeleteAppSecretRequest
//
// @return DeleteAppSecretResponse
func (client *Client) DeleteAppSecret(request *DeleteAppSecretRequest) (_result *DeleteAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.DeleteAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user group.
//
// Description:
//
// Before you delete a RAM user group, make sure that no policies are attached to the group and no RAM users are included in the group.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfileWithOptions(request *DeleteLoginProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLoginProfile"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables logon to the console for a Resource Access Management (RAM) user.
//
// @param request - DeleteLoginProfileRequest
//
// @return DeleteLoginProfileResponse
func (client *Client) DeleteLoginProfile(request *DeleteLoginProfileRequest) (_result *DeleteLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.DeleteLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the OIDC IdP named `TestOIDCProvider`.
//
// @param request - DeleteOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOIDCProviderResponse
func (client *Client) DeleteOIDCProviderWithOptions(request *DeleteOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *DeleteOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the OIDC IdP named `TestOIDCProvider`.
//
// @param request - DeleteOIDCProviderRequest
//
// @return DeleteOIDCProviderResponse
func (client *Client) DeleteOIDCProvider(request *DeleteOIDCProviderRequest) (_result *DeleteOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOIDCProviderResponse{}
	_body, _err := client.DeleteOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a passkey for a Resource Access Management (RAM) user.
//
// @param request - DeletePasskeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePasskeyResponse
func (client *Client) DeletePasskeyWithOptions(request *DeletePasskeyRequest, runtime *util.RuntimeOptions) (_result *DeletePasskeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PasskeyId)) {
		query["PasskeyId"] = request.PasskeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePasskey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePasskeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a passkey for a Resource Access Management (RAM) user.
//
// @param request - DeletePasskeyRequest
//
// @return DeletePasskeyResponse
func (client *Client) DeletePasskey(request *DeletePasskeyRequest) (_result *DeletePasskeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePasskeyResponse{}
	_body, _err := client.DeletePasskeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - DeleteSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSAMLProviderResponse
func (client *Client) DeleteSAMLProviderWithOptions(request *DeleteSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *DeleteSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SAMLProviderName)) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSAMLProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - DeleteSAMLProviderRequest
//
// @return DeleteSAMLProviderResponse
func (client *Client) DeleteSAMLProvider(request *DeleteSAMLProviderRequest) (_result *DeleteSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.DeleteSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) user.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteUserInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserInRecycleBinResponse
func (client *Client) DeleteUserInRecycleBinWithOptions(request *DeleteUserInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *DeleteUserInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - DeleteUserInRecycleBinRequest
//
// @return DeleteUserInRecycleBinResponse
func (client *Client) DeleteUserInRecycleBin(request *DeleteUserInRecycleBinRequest) (_result *DeleteUserInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserInRecycleBinResponse{}
	_body, _err := client.DeleteUserInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualMFADevice"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a multi-factor authentication (MFA) device.
//
// @param request - DeleteVirtualMFADeviceRequest
//
// @return DeleteVirtualMFADeviceResponse
func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls an external application or an internal application of the ServerApp type.
//
// Description:
//
// If you want to call this operation to uninstall an internal application, the type of the internal application must be **ServerApp**. Otherwise, an error occurs when you call this operation.
//
// >  For **internal applications**, only internal applications of the ServerApp type need to be **installed or provisioned**. Therefore, only internal applications of the ServerApp type **can be uninstalled**. Internal applications of the WebApp and NativeApp types **do not need to and cannot be uninstalled**.
//
// @param request - DeprovisionApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionApplicationResponse
func (client *Client) DeprovisionApplicationWithOptions(request *DeprovisionApplicationRequest, runtime *util.RuntimeOptions) (_result *DeprovisionApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeprovisionApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeprovisionApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls an external application or an internal application of the ServerApp type.
//
// Description:
//
// If you want to call this operation to uninstall an internal application, the type of the internal application must be **ServerApp**. Otherwise, an error occurs when you call this operation.
//
// >  For **internal applications**, only internal applications of the ServerApp type need to be **installed or provisioned**. Therefore, only internal applications of the ServerApp type **can be uninstalled**. Internal applications of the WebApp and NativeApp types **do not need to and cannot be uninstalled**.
//
// @param request - DeprovisionApplicationRequest
//
// @return DeprovisionApplicationResponse
func (client *Client) DeprovisionApplication(request *DeprovisionApplicationRequest) (_result *DeprovisionApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeprovisionApplicationResponse{}
	_body, _err := client.DeprovisionApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an installed external application.
//
// @param request - DeprovisionExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeprovisionExternalApplicationResponse
func (client *Client) DeprovisionExternalApplicationWithOptions(request *DeprovisionExternalApplicationRequest, runtime *util.RuntimeOptions) (_result *DeprovisionExternalApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeprovisionExternalApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeprovisionExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an installed external application.
//
// @param request - DeprovisionExternalApplicationRequest
//
// @return DeprovisionExternalApplicationResponse
func (client *Client) DeprovisionExternalApplication(request *DeprovisionExternalApplicationRequest) (_result *DeprovisionExternalApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeprovisionExternalApplicationResponse{}
	_body, _err := client.DeprovisionExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds and deletes a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - DisableVirtualMFARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableVirtualMFAResponse
func (client *Client) DisableVirtualMFAWithOptions(request *DisableVirtualMFARequest, runtime *util.RuntimeOptions) (_result *DisableVirtualMFAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableVirtualMFA"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds and deletes a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - DisableVirtualMFARequest
//
// @return DisableVirtualMFAResponse
func (client *Client) DisableVirtualMFA(request *DisableVirtualMFARequest) (_result *DisableVirtualMFAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.DisableVirtualMFAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the user credential report of an Alibaba Cloud account.
//
// @param request - GenerateCredentialReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCredentialReportResponse
func (client *Client) GenerateCredentialReportWithOptions(runtime *util.RuntimeOptions) (_result *GenerateCredentialReportResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GenerateCredentialReport"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the user credential report of an Alibaba Cloud account.
//
// @return GenerateCredentialReportResponse
func (client *Client) GenerateCredentialReport() (_result *GenerateCredentialReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.GenerateCredentialReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a check report for Cloud Governance.
//
// @param request - GenerateGovernanceReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateGovernanceReportResponse
func (client *Client) GenerateGovernanceReportWithOptions(runtime *util.RuntimeOptions) (_result *GenerateGovernanceReportResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GenerateGovernanceReport"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateGovernanceReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a check report for Cloud Governance.
//
// @return GenerateGovernanceReportResponse
func (client *Client) GenerateGovernanceReport() (_result *GenerateGovernanceReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateGovernanceReportResponse{}
	_body, _err := client.GenerateGovernanceReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specific AccessKey pair of a Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetAccessKeyInfoInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyInfoInRecycleBinResponse
func (client *Client) GetAccessKeyInfoInRecycleBinWithOptions(request *GetAccessKeyInfoInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyInfoInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyInfoInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyInfoInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specific AccessKey pair of a Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetAccessKeyInfoInRecycleBinRequest
//
// @return GetAccessKeyInfoInRecycleBinResponse
func (client *Client) GetAccessKeyInfoInRecycleBin(request *GetAccessKeyInfoInRecycleBinRequest) (_result *GetAccessKeyInfoInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyInfoInRecycleBinResponse{}
	_body, _err := client.GetAccessKeyInfoInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the time when an AccessKey pair was used for the last time.
//
// @param request - GetAccessKeyLastUsedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsedWithOptions(request *GetAccessKeyLastUsedRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAccessKeyLastUsed"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time when an AccessKey pair was used for the last time.
//
// @param request - GetAccessKeyLastUsedRequest
//
// @return GetAccessKeyLastUsedResponse
func (client *Client) GetAccessKeyLastUsed(request *GetAccessKeyLastUsedRequest) (_result *GetAccessKeyLastUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.GetAccessKeyLastUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) devices of an Alibaba Cloud account.
//
// @param request - GetAccountMFAInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountMFAInfoResponse
func (client *Client) GetAccountMFAInfoWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountMFAInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountMFAInfo"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) devices of an Alibaba Cloud account.
//
// @return GetAccountMFAInfoResponse
func (client *Client) GetAccountMFAInfo() (_result *GetAccountMFAInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.GetAccountMFAInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security report of an Alibaba Cloud account.
//
// @param request - GetAccountSecurityPracticeReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountSecurityPracticeReportResponse
func (client *Client) GetAccountSecurityPracticeReportWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountSecurityPracticeReport"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security report of an Alibaba Cloud account.
//
// @return GetAccountSecurityPracticeReportResponse
func (client *Client) GetAccountSecurityPracticeReport() (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.GetAccountSecurityPracticeReportWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the overview information about an Alibaba Cloud account.
//
// @param request - GetAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountSummaryResponse
func (client *Client) GetAccountSummaryWithOptions(runtime *util.RuntimeOptions) (_result *GetAccountSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetAccountSummary"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the overview information about an Alibaba Cloud account.
//
// @return GetAccountSummaryResponse
func (client *Client) GetAccountSummary() (_result *GetAccountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.GetAccountSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an application secret.
//
// @param request - GetAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSecretResponse
func (client *Client) GetAppSecretWithOptions(request *GetAppSecretRequest, runtime *util.RuntimeOptions) (_result *GetAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecretId)) {
		query["AppSecretId"] = request.AppSecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppSecret"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application secret.
//
// @param request - GetAppSecretRequest
//
// @return GetAppSecretResponse
func (client *Client) GetAppSecret(request *GetAppSecretRequest) (_result *GetAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSecretResponse{}
	_body, _err := client.GetAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration information about an application.
//
// Description:
//
// This topic provides an example on how to query the configurations of an application named `472457090344041****`.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about an application.
//
// Description:
//
// This topic provides an example on how to query the configurations of an application named `472457090344041****`.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installation information about a specified installed application.
//
// @param request - GetApplicationProvisionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisionInfoResponse
func (client *Client) GetApplicationProvisionInfoWithOptions(request *GetApplicationProvisionInfoRequest, runtime *util.RuntimeOptions) (_result *GetApplicationProvisionInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationProvisionInfo"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationProvisionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries installation information about a specified installed application.
//
// @param request - GetApplicationProvisionInfoRequest
//
// @return GetApplicationProvisionInfoResponse
func (client *Client) GetApplicationProvisionInfo(request *GetApplicationProvisionInfoRequest) (_result *GetApplicationProvisionInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationProvisionInfoResponse{}
	_body, _err := client.GetApplicationProvisionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the user credential reports of an Alibaba Cloud account.
//
// @param request - GetCredentialReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialReportResponse
func (client *Client) GetCredentialReportWithOptions(request *GetCredentialReportRequest, runtime *util.RuntimeOptions) (_result *GetCredentialReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCredentialReport"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user credential reports of an Alibaba Cloud account.
//
// @param request - GetCredentialReportRequest
//
// @return GetCredentialReportResponse
func (client *Client) GetCredentialReport(request *GetCredentialReportRequest) (_result *GetCredentialReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.GetCredentialReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default domain name of an Alibaba Cloud account.
//
// @param request - GetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultDomainResponse
func (client *Client) GetDefaultDomainWithOptions(runtime *util.RuntimeOptions) (_result *GetDefaultDomainResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetDefaultDomain"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default domain name of an Alibaba Cloud account.
//
// @return GetDefaultDomainResponse
func (client *Client) GetDefaultDomain() (_result *GetDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.GetDefaultDomainWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about an installed external application.
//
// @param request - GetExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExternalApplicationResponse
func (client *Client) GetExternalApplicationWithOptions(request *GetExternalApplicationRequest, runtime *util.RuntimeOptions) (_result *GetExternalApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExternalApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about an installed external application.
//
// @param request - GetExternalApplicationRequest
//
// @return GetExternalApplicationResponse
func (client *Client) GetExternalApplication(request *GetExternalApplicationRequest) (_result *GetExternalApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExternalApplicationResponse{}
	_body, _err := client.GetExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户的单项ram治理报告
//
// @param request - GetGovernanceItemReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceItemReportResponse
func (client *Client) GetGovernanceItemReportWithOptions(request *GetGovernanceItemReportRequest, runtime *util.RuntimeOptions) (_result *GetGovernanceItemReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GovernanceItemType)) {
		query["GovernanceItemType"] = request.GovernanceItemType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGovernanceItemReport"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGovernanceItemReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户的单项ram治理报告
//
// @param request - GetGovernanceItemReportRequest
//
// @return GetGovernanceItemReportResponse
func (client *Client) GetGovernanceItemReport(request *GetGovernanceItemReportRequest) (_result *GetGovernanceItemReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGovernanceItemReportResponse{}
	_body, _err := client.GetGovernanceItemReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询成熟度报告状态
//
// @param request - GetGovernanceReportStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGovernanceReportStatusResponse
func (client *Client) GetGovernanceReportStatusWithOptions(runtime *util.RuntimeOptions) (_result *GetGovernanceReportStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetGovernanceReportStatus"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGovernanceReportStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询成熟度报告状态
//
// @return GetGovernanceReportStatusResponse
func (client *Client) GetGovernanceReportStatus() (_result *GetGovernanceReportStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGovernanceReportStatusResponse{}
	_body, _err := client.GetGovernanceReportStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Resource Access Management (RAM) user group.
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfileWithOptions(request *GetLoginProfileRequest, runtime *util.RuntimeOptions) (_result *GetLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginProfile"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logon configurations of a Resource Access Management (RAM) user.
//
// @param request - GetLoginProfileRequest
//
// @return GetLoginProfileResponse
func (client *Client) GetLoginProfile(request *GetLoginProfileRequest) (_result *GetLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.GetLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an OIDC IdP.
//
// Description:
//
// ###
//
// This topic provides an example on how to query the information about an OpenID Connect (OIDC) identity provider (IdP) named `TestOIDCProvider`.
//
// @param request - GetOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOIDCProviderResponse
func (client *Client) GetOIDCProviderWithOptions(request *GetOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *GetOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an OIDC IdP.
//
// Description:
//
// ###
//
// This topic provides an example on how to query the information about an OpenID Connect (OIDC) identity provider (IdP) named `TestOIDCProvider`.
//
// @param request - GetOIDCProviderRequest
//
// @return GetOIDCProviderResponse
func (client *Client) GetOIDCProvider(request *GetOIDCProviderRequest) (_result *GetOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOIDCProviderResponse{}
	_body, _err := client.GetOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the password policy for RAM users.
//
// @param request - GetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicyWithOptions(runtime *util.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordPolicy"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the password policy for RAM users.
//
// @return GetPasswordPolicyResponse
func (client *Client) GetPasswordPolicy() (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - GetSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSAMLProviderResponse
func (client *Client) GetSAMLProviderWithOptions(request *GetSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *GetSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SAMLProviderName)) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSAMLProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// @param request - GetSAMLProviderRequest
//
// @return GetSAMLProviderResponse
func (client *Client) GetSAMLProvider(request *GetSAMLProviderRequest) (_result *GetSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.GetSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security preferences for RAM users.
//
// @param request - GetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreferenceWithOptions(runtime *util.RuntimeOptions) (_result *GetSecurityPreferenceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetSecurityPreference"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security preferences for RAM users.
//
// @return GetSecurityPreferenceResponse
func (client *Client) GetSecurityPreference() (_result *GetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.GetSecurityPreferenceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a RAM user.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a RAM user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetUserInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInRecycleBinResponse
func (client *Client) GetUserInRecycleBinWithOptions(request *GetUserInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *GetUserInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - GetUserInRecycleBinRequest
//
// @return GetUserInRecycleBinResponse
func (client *Client) GetUserInRecycleBin(request *GetUserInRecycleBinRequest) (_result *GetUserInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserInRecycleBinResponse{}
	_body, _err := client.GetUserInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfoWithOptions(request *GetUserMFAInfoRequest, runtime *util.RuntimeOptions) (_result *GetUserMFAInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserMFAInfo"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the multi-factor authentication (MFA) device that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetUserMFAInfoRequest
//
// @return GetUserMFAInfoResponse
func (client *Client) GetUserMFAInfo(request *GetUserMFAInfoRequest) (_result *GetUserMFAInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.GetUserMFAInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of user-based single sign-on (SSO).
//
// @param request - GetUserSsoSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserSsoSettingsResponse
func (client *Client) GetUserSsoSettingsWithOptions(runtime *util.RuntimeOptions) (_result *GetUserSsoSettingsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetUserSsoSettings"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of user-based single sign-on (SSO).
//
// @return GetUserSsoSettingsResponse
func (client *Client) GetUserSsoSettings() (_result *GetUserSsoSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.GetUserSsoSettingsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the mobile phone or email that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVerificationInfoResponse
func (client *Client) GetVerificationInfoWithOptions(request *GetVerificationInfoRequest, runtime *util.RuntimeOptions) (_result *GetVerificationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVerificationInfo"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the mobile phone or email that is bound to a Resource Access Management (RAM) user.
//
// @param request - GetVerificationInfoRequest
//
// @return GetVerificationInfoResponse
func (client *Client) GetVerificationInfo(request *GetVerificationInfoRequest) (_result *GetVerificationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVerificationInfoResponse{}
	_body, _err := client.GetVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeysWithOptions(request *ListAccessKeysRequest, runtime *util.RuntimeOptions) (_result *ListAccessKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessKeys"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - ListAccessKeysRequest
//
// @return ListAccessKeysResponse
func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (_result *ListAccessKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.ListAccessKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - ListAccessKeysInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessKeysInRecycleBinResponse
func (client *Client) ListAccessKeysInRecycleBinWithOptions(request *ListAccessKeysInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *ListAccessKeysInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAccessKeysInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAccessKeysInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs of a specific Resource Access Management (RAM) user in the recycle bin.
//
// @param request - ListAccessKeysInRecycleBinRequest
//
// @return ListAccessKeysInRecycleBinResponse
func (client *Client) ListAccessKeysInRecycleBin(request *ListAccessKeysInRecycleBinRequest) (_result *ListAccessKeysInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessKeysInRecycleBinResponse{}
	_body, _err := client.ListAccessKeysInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the secret IDs of an application.
//
// @param request - ListAppSecretIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppSecretIdsResponse
func (client *Client) ListAppSecretIdsWithOptions(request *ListAppSecretIdsRequest, runtime *util.RuntimeOptions) (_result *ListAppSecretIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppSecretIds"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the secret IDs of an application.
//
// @param request - ListAppSecretIdsRequest
//
// @return ListAppSecretIdsResponse
func (client *Client) ListAppSecretIds(request *ListAppSecretIdsRequest) (_result *ListAppSecretIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.ListAppSecretIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries installation information about all installed applications.
//
// @param request - ListApplicationProvisionInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationProvisionInfosResponse
func (client *Client) ListApplicationProvisionInfosWithOptions(request *ListApplicationProvisionInfosRequest, runtime *util.RuntimeOptions) (_result *ListApplicationProvisionInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationProvisionInfos"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationProvisionInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries installation information about all installed applications.
//
// @param request - ListApplicationProvisionInfosRequest
//
// @return ListApplicationProvisionInfosResponse
func (client *Client) ListApplicationProvisionInfos(request *ListApplicationProvisionInfosRequest) (_result *ListApplicationProvisionInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationProvisionInfosResponse{}
	_body, _err := client.ListApplicationProvisionInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the created applications.
//
// Description:
//
// This topic provides an example on how to query the applications within the current account. The returned result shows that only one application named `myapp` belongs to the current account.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListApplications"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the created applications.
//
// Description:
//
// This topic provides an example on how to query the applications within the current account. The returned result shows that only one application named `myapp` belongs to the current account.
//
// @return ListApplicationsResponse
func (client *Client) ListApplications() (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all installed external applications.
//
// @param request - ListExternalApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalApplicationsResponse
func (client *Client) ListExternalApplicationsWithOptions(runtime *util.RuntimeOptions) (_result *ListExternalApplicationsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListExternalApplications"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExternalApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all installed external applications.
//
// @return ListExternalApplicationsResponse
func (client *Client) ListExternalApplications() (_result *ListExternalApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExternalApplicationsResponse{}
	_body, _err := client.ListExternalApplicationsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) user groups.
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupsForUser"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) user groups to which a RAM user belongs.
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries OIDC IdPs.
//
// Description:
//
// ###
//
// This topic provides an example on how to query all OpenID Connect (OIDC) identity providers (IdPs) within your Alibaba Cloud account. The response shows that your Alibaba Cloud account has only one OIDC IdP named `TestOIDCProvider`.
//
// @param request - ListOIDCProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOIDCProvidersResponse
func (client *Client) ListOIDCProvidersWithOptions(request *ListOIDCProvidersRequest, runtime *util.RuntimeOptions) (_result *ListOIDCProvidersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOIDCProviders"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOIDCProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries OIDC IdPs.
//
// Description:
//
// ###
//
// This topic provides an example on how to query all OpenID Connect (OIDC) identity providers (IdPs) within your Alibaba Cloud account. The response shows that your Alibaba Cloud account has only one OIDC IdP named `TestOIDCProvider`.
//
// @param request - ListOIDCProvidersRequest
//
// @return ListOIDCProvidersResponse
func (client *Client) ListOIDCProviders(request *ListOIDCProvidersRequest) (_result *ListOIDCProvidersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOIDCProvidersResponse{}
	_body, _err := client.ListOIDCProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the passkeys that are bound to a Resource Access Management (RAM) user.
//
// @param request - ListPasskeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPasskeysResponse
func (client *Client) ListPasskeysWithOptions(request *ListPasskeysRequest, runtime *util.RuntimeOptions) (_result *ListPasskeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPasskeys"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPasskeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the passkeys that are bound to a Resource Access Management (RAM) user.
//
// @param request - ListPasskeysRequest
//
// @return ListPasskeysResponse
func (client *Client) ListPasskeys(request *ListPasskeysRequest) (_result *ListPasskeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPasskeysResponse{}
	_body, _err := client.ListPasskeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries predefined application permissions.
//
// @param request - ListPredefinedScopesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPredefinedScopesResponse
func (client *Client) ListPredefinedScopesWithOptions(request *ListPredefinedScopesRequest, runtime *util.RuntimeOptions) (_result *ListPredefinedScopesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		query["AppType"] = request.AppType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPredefinedScopes"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries predefined application permissions.
//
// @param request - ListPredefinedScopesRequest
//
// @return ListPredefinedScopesResponse
func (client *Client) ListPredefinedScopes(request *ListPredefinedScopesRequest) (_result *ListPredefinedScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.ListPredefinedScopesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all metric values in the most recent governance check.
//
// @param request - ListRecentGovernanceMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentGovernanceMetricsResponse
func (client *Client) ListRecentGovernanceMetricsWithOptions(runtime *util.RuntimeOptions) (_result *ListRecentGovernanceMetricsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRecentGovernanceMetrics"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecentGovernanceMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all metric values in the most recent governance check.
//
// @return ListRecentGovernanceMetricsResponse
func (client *Client) ListRecentGovernanceMetrics() (_result *ListRecentGovernanceMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRecentGovernanceMetricsResponse{}
	_body, _err := client.ListRecentGovernanceMetricsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about identity providers (IdPs) for role-based single sign-on (SSO).
//
// @param request - ListSAMLProvidersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSAMLProvidersResponse
func (client *Client) ListSAMLProvidersWithOptions(request *ListSAMLProvidersRequest, runtime *util.RuntimeOptions) (_result *ListSAMLProvidersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSAMLProviders"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about identity providers (IdPs) for role-based single sign-on (SSO).
//
// @param request - ListSAMLProvidersRequest
//
// @return ListSAMLProvidersResponse
func (client *Client) ListSAMLProviders(request *ListSAMLProvidersRequest) (_result *ListSAMLProvidersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.ListSAMLProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added resources.
//
// Description:
//
// ###
//
// You must specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N.Key`
//
//   - `Tag.N.Key` and `Tag.N.Value`
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePrincipalName)) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added resources.
//
// Description:
//
// ###
//
// You must specify at least one of the following parameters or parameter pairs in a request to determine a query object:
//
//   - `ResourceId.N`
//
//   - `Tag.N.Key`
//
//   - `Tag.N.Key` and `Tag.N.Value`
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users.
//
// Description:
//
// You can call the following API operations to query information about all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUserBasicInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserBasicInfosResponse
func (client *Client) ListUserBasicInfosWithOptions(request *ListUserBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListUserBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserBasicInfos"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users.
//
// Description:
//
// You can call the following API operations to query information about all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUserBasicInfosRequest
//
// @return ListUserBasicInfosResponse
func (client *Client) ListUserBasicInfos(request *ListUserBasicInfosRequest) (_result *ListUserBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.ListUserBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all Resource Access Management (RAM) users.
//
// Description:
//
// ### [](#)
//
// You can call the following API operations to query the details of all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all Resource Access Management (RAM) users.
//
// Description:
//
// ### [](#)
//
// You can call the following API operations to query the details of all RAM users:
//
//   - ListUsers: queries the details of all RAM users.
//
//   - ListUserBasicInfos: queries the basic information about all RAM users. The basic information includes only the logon names (`UserPrincipalName`), display names (`DisplayName`), and user IDs (`UserId`).
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *util.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersForGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Resource Access Management (RAM) users in a RAM user group.
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users in the recycle bin.
//
// @param request - ListUsersInRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersInRecycleBinResponse
func (client *Client) ListUsersInRecycleBinWithOptions(request *ListUsersInRecycleBinRequest, runtime *util.RuntimeOptions) (_result *ListUsersInRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersInRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersInRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about all Resource Access Management (RAM) users in the recycle bin.
//
// @param request - ListUsersInRecycleBinRequest
//
// @return ListUsersInRecycleBinResponse
func (client *Client) ListUsersInRecycleBin(request *ListUsersInRecycleBinRequest) (_result *ListUsersInRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersInRecycleBinResponse{}
	_body, _err := client.ListUsersInRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @param request - ListVirtualMFADevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevicesWithOptions(request *ListVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *ListVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["Marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItems)) {
		query["MaxItems"] = request.MaxItems
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVirtualMFADevices"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries multi-factor authentication (MFA) devices.
//
// @param request - ListVirtualMFADevicesRequest
//
// @return ListVirtualMFADevicesResponse
func (client *Client) ListVirtualMFADevices(request *ListVirtualMFADevicesRequest) (_result *ListVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.ListVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an application.
//
// @param request - ProvisionApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionApplicationResponse
func (client *Client) ProvisionApplicationWithOptions(request *ProvisionApplicationRequest, runtime *util.RuntimeOptions) (_result *ProvisionApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Scopes)) {
		query["Scopes"] = request.Scopes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ProvisionApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProvisionApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an application.
//
// @param request - ProvisionApplicationRequest
//
// @return ProvisionApplicationResponse
func (client *Client) ProvisionApplication(request *ProvisionApplicationRequest) (_result *ProvisionApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProvisionApplicationResponse{}
	_body, _err := client.ProvisionApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs an external application.
//
// @param request - ProvisionExternalApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProvisionExternalApplicationResponse
func (client *Client) ProvisionExternalApplicationWithOptions(request *ProvisionExternalApplicationRequest, runtime *util.RuntimeOptions) (_result *ProvisionExternalApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Scopes)) {
		query["Scopes"] = request.Scopes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ProvisionExternalApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ProvisionExternalApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs an external application.
//
// @param request - ProvisionExternalApplicationRequest
//
// @return ProvisionExternalApplicationResponse
func (client *Client) ProvisionExternalApplication(request *ProvisionExternalApplicationRequest) (_result *ProvisionExternalApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ProvisionExternalApplicationResponse{}
	_body, _err := client.ProvisionExternalApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a client ID from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the client ID `498469743454717****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveClientIdFromOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveClientIdFromOIDCProviderResponse
func (client *Client) RemoveClientIdFromOIDCProviderWithOptions(request *RemoveClientIdFromOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *RemoveClientIdFromOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveClientIdFromOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveClientIdFromOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a client ID from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the client ID `498469743454717****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveClientIdFromOIDCProviderRequest
//
// @return RemoveClientIdFromOIDCProviderResponse
func (client *Client) RemoveClientIdFromOIDCProvider(request *RemoveClientIdFromOIDCProviderRequest) (_result *RemoveClientIdFromOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveClientIdFromOIDCProviderResponse{}
	_body, _err := client.RemoveClientIdFromOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a fingerprint from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the fingerprint `6938fd4d98bab03faadb97b34396831e3780****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveFingerprintFromOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFingerprintFromOIDCProviderResponse
func (client *Client) RemoveFingerprintFromOIDCProviderWithOptions(request *RemoveFingerprintFromOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *RemoveFingerprintFromOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fingerprint)) {
		query["Fingerprint"] = request.Fingerprint
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFingerprintFromOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFingerprintFromOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a fingerprint from an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to remove the fingerprint `6938fd4d98bab03faadb97b34396831e3780****` from the OIDC IdP named `TestOIDCProvider`.
//
// @param request - RemoveFingerprintFromOIDCProviderRequest
//
// @return RemoveFingerprintFromOIDCProviderResponse
func (client *Client) RemoveFingerprintFromOIDCProvider(request *RemoveFingerprintFromOIDCProviderRequest) (_result *RemoveFingerprintFromOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveFingerprintFromOIDCProviderResponse{}
	_body, _err := client.RemoveFingerprintFromOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a Resource Access Management (RAM) user from a RAM user group.
//
// @param request - RemoveUserFromGroupRequest
//
// @return RemoveUserFromGroupResponse
func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreAccessKeyFromRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreAccessKeyFromRecycleBinResponse
func (client *Client) RestoreAccessKeyFromRecycleBinWithOptions(request *RestoreAccessKeyFromRecycleBinRequest, runtime *util.RuntimeOptions) (_result *RestoreAccessKeyFromRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreAccessKeyFromRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreAccessKeyFromRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a specific AccessKey pair that belongs to a Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreAccessKeyFromRecycleBinRequest
//
// @return RestoreAccessKeyFromRecycleBinResponse
func (client *Client) RestoreAccessKeyFromRecycleBin(request *RestoreAccessKeyFromRecycleBinRequest) (_result *RestoreAccessKeyFromRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreAccessKeyFromRecycleBinResponse{}
	_body, _err := client.RestoreAccessKeyFromRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreUserFromRecycleBinRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreUserFromRecycleBinResponse
func (client *Client) RestoreUserFromRecycleBinWithOptions(request *RestoreUserFromRecycleBinRequest, runtime *util.RuntimeOptions) (_result *RestoreUserFromRecycleBinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreUserFromRecycleBin"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreUserFromRecycleBinResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a specific Resource Access Management (RAM) user from the recycle bin.
//
// @param request - RestoreUserFromRecycleBinRequest
//
// @return RestoreUserFromRecycleBinResponse
func (client *Client) RestoreUserFromRecycleBin(request *RestoreUserFromRecycleBinRequest) (_result *RestoreUserFromRecycleBinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreUserFromRecycleBinResponse{}
	_body, _err := client.RestoreUserFromRecycleBinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the default domain name for an Alibaba Cloud account.
//
// @param request - SetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithOptions(request *SetDefaultDomainRequest, runtime *util.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultDomainName)) {
		query["DefaultDomainName"] = request.DefaultDomainName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDefaultDomain"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the default domain name for an Alibaba Cloud account.
//
// @param request - SetDefaultDomainRequest
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomain(request *SetDefaultDomainRequest) (_result *SetDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.SetDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users.
//
// @param request - SetPasswordPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicyWithOptions(request *SetPasswordPolicyRequest, runtime *util.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HardExpire)) {
		query["HardExpire"] = request.HardExpire
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLoginAttemps)) {
		query["MaxLoginAttemps"] = request.MaxLoginAttemps
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPasswordAge)) {
		query["MaxPasswordAge"] = request.MaxPasswordAge
	}

	if !tea.BoolValue(util.IsUnset(request.MinimumPasswordDifferentCharacter)) {
		query["MinimumPasswordDifferentCharacter"] = request.MinimumPasswordDifferentCharacter
	}

	if !tea.BoolValue(util.IsUnset(request.MinimumPasswordLength)) {
		query["MinimumPasswordLength"] = request.MinimumPasswordLength
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordNotContainUserName)) {
		query["PasswordNotContainUserName"] = request.PasswordNotContainUserName
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordReusePrevention)) {
		query["PasswordReusePrevention"] = request.PasswordReusePrevention
	}

	if !tea.BoolValue(util.IsUnset(request.RequireLowercaseCharacters)) {
		query["RequireLowercaseCharacters"] = request.RequireLowercaseCharacters
	}

	if !tea.BoolValue(util.IsUnset(request.RequireNumbers)) {
		query["RequireNumbers"] = request.RequireNumbers
	}

	if !tea.BoolValue(util.IsUnset(request.RequireSymbols)) {
		query["RequireSymbols"] = request.RequireSymbols
	}

	if !tea.BoolValue(util.IsUnset(request.RequireUppercaseCharacters)) {
		query["RequireUppercaseCharacters"] = request.RequireUppercaseCharacters
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordPolicy"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the password policy for Resource Access Management (RAM) users.
//
// @param request - SetPasswordPolicyRequest
//
// @return SetPasswordPolicyResponse
func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (_result *SetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.SetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures security preferences for a RAM user.
//
// Description:
//
// ###
//
// This topic provides an example on how to enable multi-factor authentication (MFA) only for RAM users who initiated unusual logons.
//
// @param tmpReq - SetSecurityPreferenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreferenceWithOptions(tmpReq *SetSecurityPreferenceRequest, runtime *util.RuntimeOptions) (_result *SetSecurityPreferenceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SetSecurityPreferenceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VerificationTypes)) {
		request.VerificationTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VerificationTypes, tea.String("VerificationTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowUserToChangePassword)) {
		query["AllowUserToChangePassword"] = request.AllowUserToChangePassword
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToLoginWithPasskey)) {
		query["AllowUserToLoginWithPasskey"] = request.AllowUserToLoginWithPasskey
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManageAccessKeys)) {
		query["AllowUserToManageAccessKeys"] = request.AllowUserToManageAccessKeys
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManageMFADevices)) {
		query["AllowUserToManageMFADevices"] = request.AllowUserToManageMFADevices
	}

	if !tea.BoolValue(util.IsUnset(request.AllowUserToManagePersonalDingTalk)) {
		query["AllowUserToManagePersonalDingTalk"] = request.AllowUserToManagePersonalDingTalk
	}

	if !tea.BoolValue(util.IsUnset(request.EnableSaveMFATicket)) {
		query["EnableSaveMFATicket"] = request.EnableSaveMFATicket
	}

	if !tea.BoolValue(util.IsUnset(request.LoginNetworkMasks)) {
		query["LoginNetworkMasks"] = request.LoginNetworkMasks
	}

	if !tea.BoolValue(util.IsUnset(request.LoginSessionDuration)) {
		query["LoginSessionDuration"] = request.LoginSessionDuration
	}

	if !tea.BoolValue(util.IsUnset(request.MFAOperationForLogin)) {
		query["MFAOperationForLogin"] = request.MFAOperationForLogin
	}

	if !tea.BoolValue(util.IsUnset(request.OperationForRiskLogin)) {
		query["OperationForRiskLogin"] = request.OperationForRiskLogin
	}

	if !tea.BoolValue(util.IsUnset(request.VerificationTypesShrink)) {
		query["VerificationTypes"] = request.VerificationTypesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSecurityPreference"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures security preferences for a RAM user.
//
// Description:
//
// ###
//
// This topic provides an example on how to enable multi-factor authentication (MFA) only for RAM users who initiated unusual logons.
//
// @param request - SetSecurityPreferenceRequest
//
// @return SetSecurityPreferenceResponse
func (client *Client) SetSecurityPreference(request *SetSecurityPreferenceRequest) (_result *SetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.SetSecurityPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures information about user-based single sign-on (SSO).
//
// @param request - SetUserSsoSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserSsoSettingsResponse
func (client *Client) SetUserSsoSettingsWithOptions(request *SetUserSsoSettingsRequest, runtime *util.RuntimeOptions) (_result *SetUserSsoSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuxiliaryDomain)) {
		query["AuxiliaryDomain"] = request.AuxiliaryDomain
	}

	if !tea.BoolValue(util.IsUnset(request.MetadataDocument)) {
		query["MetadataDocument"] = request.MetadataDocument
	}

	if !tea.BoolValue(util.IsUnset(request.SsoEnabled)) {
		query["SsoEnabled"] = request.SsoEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.SsoLoginWithDomain)) {
		query["SsoLoginWithDomain"] = request.SsoLoginWithDomain
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserSsoSettings"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures information about user-based single sign-on (SSO).
//
// @param request - SetUserSsoSettingsRequest
//
// @return SetUserSsoSettingsResponse
func (client *Client) SetUserSsoSettings(request *SetUserSsoSettingsRequest) (_result *SetUserSsoSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.SetUserSsoSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a mobile phone or email to a Resource Access Management (RAM) user.
//
// @param request - SetVerificationInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVerificationInfoResponse
func (client *Client) SetVerificationInfoWithOptions(request *SetVerificationInfoRequest, runtime *util.RuntimeOptions) (_result *SetVerificationInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhone)) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetVerificationInfo"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetVerificationInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a mobile phone or email to a Resource Access Management (RAM) user.
//
// @param request - SetVerificationInfoRequest
//
// @return SetVerificationInfoResponse
func (client *Client) SetVerificationInfo(request *SetVerificationInfoRequest) (_result *SetVerificationInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVerificationInfoResponse{}
	_body, _err := client.SetVerificationInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePrincipalName)) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADeviceWithOptions(request *UnbindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindMFADevice"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a multi-factor authentication (MFA) device from a Resource Access Management (RAM) user.
//
// @param request - UnbindMFADeviceRequest
//
// @return UnbindMFADeviceResponse
func (client *Client) UnbindMFADevice(request *UnbindMFADeviceRequest) (_result *UnbindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.UnbindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a mobile phone or email from a Resource Access Management (RAM) user.
//
// @param request - UnbindVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindVerificationResponse
func (client *Client) UnbindVerificationWithOptions(request *UnbindVerificationRequest, runtime *util.RuntimeOptions) (_result *UnbindVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.MobilePhone)) {
		query["MobilePhone"] = request.MobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindVerification"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a mobile phone or email from a Resource Access Management (RAM) user.
//
// @param request - UnbindVerificationRequest
//
// @return UnbindVerificationResponse
func (client *Client) UnbindVerification(request *UnbindVerificationRequest) (_result *UnbindVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindVerificationResponse{}
	_body, _err := client.UnbindVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePrincipalName)) {
		query["ResourcePrincipalName"] = request.ResourcePrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKeyWithOptions(request *UpdateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessKeyId)) {
		query["UserAccessKeyId"] = request.UserAccessKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAccessKey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of an AccessKey pair for an Alibaba Cloud account or a Resource Access Management (RAM) user.
//
// @param request - UpdateAccessKeyRequest
//
// @return UpdateAccessKeyResponse
func (client *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (_result *UpdateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.UpdateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a specified application.
//
// @param request - UpdateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithOptions(request *UpdateApplicationRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.NewAccessTokenValidity)) {
		query["NewAccessTokenValidity"] = request.NewAccessTokenValidity
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.NewIsMultiTenant)) {
		query["NewIsMultiTenant"] = request.NewIsMultiTenant
	}

	if !tea.BoolValue(util.IsUnset(request.NewPredefinedScopes)) {
		query["NewPredefinedScopes"] = request.NewPredefinedScopes
	}

	if !tea.BoolValue(util.IsUnset(request.NewRedirectUris)) {
		query["NewRedirectUris"] = request.NewRedirectUris
	}

	if !tea.BoolValue(util.IsUnset(request.NewRefreshTokenValidity)) {
		query["NewRefreshTokenValidity"] = request.NewRefreshTokenValidity
	}

	if !tea.BoolValue(util.IsUnset(request.NewRequiredScopes)) {
		query["NewRequiredScopes"] = request.NewRequiredScopes
	}

	if !tea.BoolValue(util.IsUnset(request.NewSecretRequired)) {
		query["NewSecretRequired"] = request.NewSecretRequired
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplication"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a specified application.
//
// @param request - UpdateApplicationRequest
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NewComments)) {
		query["NewComments"] = request.NewComments
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.NewGroupName)) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a Resource Access Management (RAM) user group.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the console logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfileWithOptions(request *UpdateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *UpdateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MFABindRequired)) {
		query["MFABindRequired"] = request.MFABindRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordResetRequired)) {
		query["PasswordResetRequired"] = request.PasswordResetRequired
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLoginProfile"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the console logon configurations of a Resource Access Management (RAM) user.
//
// @param request - UpdateLoginProfileRequest
//
// @return UpdateLoginProfileResponse
func (client *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (_result *UpdateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.UpdateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description and client IDs of an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to change the description of the OIDC IdP named `TestOIDCProvider` to `This is a new OIDC Provider.`
//
// @param request - UpdateOIDCProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOIDCProviderResponse
func (client *Client) UpdateOIDCProviderWithOptions(request *UpdateOIDCProviderRequest, runtime *util.RuntimeOptions) (_result *UpdateOIDCProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientIds)) {
		query["ClientIds"] = request.ClientIds
	}

	if !tea.BoolValue(util.IsUnset(request.IssuanceLimitTime)) {
		query["IssuanceLimitTime"] = request.IssuanceLimitTime
	}

	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.OIDCProviderName)) {
		query["OIDCProviderName"] = request.OIDCProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOIDCProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOIDCProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description and client IDs of an OpenID Connect (OIDC) identity provider (IdP).
//
// Description:
//
// ###
//
// This topic provides an example on how to change the description of the OIDC IdP named `TestOIDCProvider` to `This is a new OIDC Provider.`
//
// @param request - UpdateOIDCProviderRequest
//
// @return UpdateOIDCProviderResponse
func (client *Client) UpdateOIDCProvider(request *UpdateOIDCProviderRequest) (_result *UpdateOIDCProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOIDCProviderResponse{}
	_body, _err := client.UpdateOIDCProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a passkey.
//
// @param request - UpdatePasskeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePasskeyResponse
func (client *Client) UpdatePasskeyWithOptions(request *UpdatePasskeyRequest, runtime *util.RuntimeOptions) (_result *UpdatePasskeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PasskeyId)) {
		query["PasskeyId"] = request.PasskeyId
	}

	if !tea.BoolValue(util.IsUnset(request.PasskeyName)) {
		query["PasskeyName"] = request.PasskeyName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePasskey"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePasskeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a passkey.
//
// @param request - UpdatePasskeyRequest
//
// @return UpdatePasskeyResponse
func (client *Client) UpdatePasskey(request *UpdatePasskeyRequest) (_result *UpdatePasskeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePasskeyResponse{}
	_body, _err := client.UpdatePasskeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// Description:
//
// This topic provides an example on how to change the description of an IdP named `test-provider` to `This is a new provider.`
//
// @param request - UpdateSAMLProviderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSAMLProviderResponse
func (client *Client) UpdateSAMLProviderWithOptions(request *UpdateSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *UpdateSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewDescription)) {
		query["NewDescription"] = request.NewDescription
	}

	if !tea.BoolValue(util.IsUnset(request.NewEncodedSAMLMetadataDocument)) {
		query["NewEncodedSAMLMetadataDocument"] = request.NewEncodedSAMLMetadataDocument
	}

	if !tea.BoolValue(util.IsUnset(request.SAMLProviderName)) {
		query["SAMLProviderName"] = request.SAMLProviderName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSAMLProvider"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about an identity provider (IdP) for role-based single sign-on (SSO).
//
// Description:
//
// This topic provides an example on how to change the description of an IdP named `test-provider` to `This is a new provider.`
//
// @param request - UpdateSAMLProviderRequest
//
// @return UpdateSAMLProviderResponse
func (client *Client) UpdateSAMLProvider(request *UpdateSAMLProviderRequest) (_result *UpdateSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.UpdateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a RAM user.
//
// Description:
//
// This topic provides an example to show how to modify the name of a RAM user from `test@example.onaliyun.com` to `new@example.onaliyun.com`.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewComments)) {
		query["NewComments"] = request.NewComments
	}

	if !tea.BoolValue(util.IsUnset(request.NewDisplayName)) {
		query["NewDisplayName"] = request.NewDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.NewEmail)) {
		query["NewEmail"] = request.NewEmail
	}

	if !tea.BoolValue(util.IsUnset(request.NewMobilePhone)) {
		query["NewMobilePhone"] = request.NewMobilePhone
	}

	if !tea.BoolValue(util.IsUnset(request.NewUserPrincipalName)) {
		query["NewUserPrincipalName"] = request.NewUserPrincipalName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPrincipalName)) {
		query["UserPrincipalName"] = request.UserPrincipalName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2019-08-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a RAM user.
//
// Description:
//
// This topic provides an example to show how to modify the name of a RAM user from `test@example.onaliyun.com` to `new@example.onaliyun.com`.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
