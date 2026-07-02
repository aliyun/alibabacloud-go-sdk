// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccount(v *GetCloudAccountResponseBodyCloudAccount) *GetCloudAccountResponseBody
	GetCloudAccount() *GetCloudAccountResponseBodyCloudAccount
	SetRequestId(v string) *GetCloudAccountResponseBody
	GetRequestId() *string
}

type GetCloudAccountResponseBody struct {
	// The cloud account details.
	CloudAccount *GetCloudAccountResponseBodyCloudAccount `json:"CloudAccount,omitempty" xml:"CloudAccount,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponseBody) GetCloudAccount() *GetCloudAccountResponseBodyCloudAccount {
	return s.CloudAccount
}

func (s *GetCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAccountResponseBody) SetCloudAccount(v *GetCloudAccountResponseBodyCloudAccount) *GetCloudAccountResponseBody {
	s.CloudAccount = v
	return s
}

func (s *GetCloudAccountResponseBody) SetRequestId(v string) *GetCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAccountResponseBody) Validate() error {
	if s.CloudAccount != nil {
		if err := s.CloudAccount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountResponseBodyCloudAccount struct {
	// The external unique identifier of the cloud account.
	//
	// example:
	//
	// 1234567
	CloudAccountExternalId *string `json:"CloudAccountExternalId,omitempty" xml:"CloudAccountExternalId,omitempty"`
	// The health status of the cloud account. Valid values:
	//
	// - healthy: Healthy.
	//
	// - unhealthy: Unhealthy.
	//
	// - unknown: Unknown.
	//
	// example:
	//
	// healthy
	CloudAccountHealth *string `json:"CloudAccountHealth,omitempty" xml:"CloudAccountHealth,omitempty"`
	// The health check result of the cloud account.
	CloudAccountHealthCheckResult *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult `json:"CloudAccountHealthCheckResult,omitempty" xml:"CloudAccountHealthCheckResult,omitempty" type:"Struct"`
	// The cloud account ID.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The cloud account name.
	//
	// example:
	//
	// cloud_accout_xxxx
	CloudAccountName *string `json:"CloudAccountName,omitempty" xml:"CloudAccountName,omitempty"`
	// The identity provider configuration.
	CloudAccountProviderConfig *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig `json:"CloudAccountProviderConfig,omitempty" xml:"CloudAccountProviderConfig,omitempty" type:"Struct"`
	// The identity provider name.
	//
	// example:
	//
	// idaas-eiam-oidc-provider
	CloudAccountProviderName *string `json:"CloudAccountProviderName,omitempty" xml:"CloudAccountProviderName,omitempty"`
	CloudAccountSite         *string `json:"CloudAccountSite,omitempty" xml:"CloudAccountSite,omitempty"`
	// The cloud account type. Valid values:
	//
	// - alibaba_cloud: Alibaba Cloud.
	//
	// example:
	//
	// alibaba_cloud
	CloudAccountVendorType *string `json:"CloudAccountVendorType,omitempty" xml:"CloudAccountVendorType,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cloud account description.
	//
	// example:
	//
	// cloud_accout_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The most recent update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCloudAccountResponseBodyCloudAccount) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponseBodyCloudAccount) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountExternalId() *string {
	return s.CloudAccountExternalId
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountHealth() *string {
	return s.CloudAccountHealth
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountHealthCheckResult() *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult {
	return s.CloudAccountHealthCheckResult
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountName() *string {
	return s.CloudAccountName
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountProviderConfig() *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig {
	return s.CloudAccountProviderConfig
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountProviderName() *string {
	return s.CloudAccountProviderName
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountSite() *string {
	return s.CloudAccountSite
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCloudAccountVendorType() *string {
	return s.CloudAccountVendorType
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetDescription() *string {
	return s.Description
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAccountResponseBodyCloudAccount) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountExternalId(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountExternalId = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountHealth(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountHealth = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountHealthCheckResult(v *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountHealthCheckResult = v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountId(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountName(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountName = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountProviderConfig(v *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountProviderConfig = v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountProviderName(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountProviderName = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountSite(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountSite = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCloudAccountVendorType(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.CloudAccountVendorType = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetCreateTime(v int64) *GetCloudAccountResponseBodyCloudAccount {
	s.CreateTime = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetDescription(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.Description = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetInstanceId(v string) *GetCloudAccountResponseBodyCloudAccount {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) SetUpdateTime(v int64) *GetCloudAccountResponseBodyCloudAccount {
	s.UpdateTime = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccount) Validate() error {
	if s.CloudAccountHealthCheckResult != nil {
		if err := s.CloudAccountHealthCheckResult.Validate(); err != nil {
			return err
		}
	}
	if s.CloudAccountProviderConfig != nil {
		if err := s.CloudAccountProviderConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult struct {
	// The error reason. This field is returned when the health check status is unhealthy.
	ErrorReason *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// The time of the last health check. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The health check result of the cloud account. Valid values:
	//
	// - success: Succeeded.
	//
	// - failed: Failed.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) GetErrorReason() *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason {
	return s.ErrorReason
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) GetResult() *string {
	return s.Result
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) SetErrorReason(v *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult {
	s.ErrorReason = v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) SetLastCheckTime(v int64) *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult {
	s.LastCheckTime = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) SetResult(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult {
	s.Result = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResult) Validate() error {
	if s.ErrorReason != nil {
		if err := s.ErrorReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason struct {
	// The error code.
	//
	// example:
	//
	// AuthenticationFail.NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// There is no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) SetErrorCode(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason {
	s.ErrorCode = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) SetErrorMessage(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason {
	s.ErrorMessage = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountHealthCheckResultErrorReason) Validate() error {
	return dara.Validate(s)
}

type GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig struct {
	// The audience identifier.
	//
	// example:
	//
	// urn:cloud:idaas:sts:xxx:xxx
	Audience *string `json:"Audience,omitempty" xml:"Audience,omitempty"`
	// The authorization server ID.
	//
	// example:
	//
	// iauths_system
	AuthorizationServerId *string `json:"AuthorizationServerId,omitempty" xml:"AuthorizationServerId,omitempty"`
	// Issuer。
	//
	// example:
	//
	// https://xxxxx.aliyunidaas.com/api/v2/iauths_system/oauth2
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The signature verification public key endpoint.
	//
	// example:
	//
	// https://xxxxx.aliyunidaas.com/api/v2/iauths_system/oauth2/jwks
	OidcJwksEndpoint *string `json:"OidcJwksEndpoint,omitempty" xml:"OidcJwksEndpoint,omitempty"`
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) GetAudience() *string {
	return s.Audience
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) GetAuthorizationServerId() *string {
	return s.AuthorizationServerId
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) GetOidcJwksEndpoint() *string {
	return s.OidcJwksEndpoint
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) SetAudience(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig {
	s.Audience = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) SetAuthorizationServerId(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig {
	s.AuthorizationServerId = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) SetIssuer(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig {
	s.Issuer = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) SetOidcJwksEndpoint(v string) *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig {
	s.OidcJwksEndpoint = &v
	return s
}

func (s *GetCloudAccountResponseBodyCloudAccountCloudAccountProviderConfig) Validate() error {
	return dara.Validate(s)
}
