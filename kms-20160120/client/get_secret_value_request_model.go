// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v string) *GetSecretValueRequest
	GetDryRun() *string
	SetFetchExtendedConfig(v bool) *GetSecretValueRequest
	GetFetchExtendedConfig() *bool
	SetRecipient(v string) *GetSecretValueRequest
	GetRecipient() *string
	SetSecretName(v string) *GetSecretValueRequest
	GetSecretName() *string
	SetVersionId(v string) *GetSecretValueRequest
	GetVersionId() *string
	SetVersionStage(v string) *GetSecretValueRequest
	GetVersionStage() *string
}

type GetSecretValueRequest struct {
	// Specifies whether to enable DryRun mode. Valid values:
	//
	// - true: Enables DryRun mode.
	//
	// - false (default): Disables DryRun mode.
	//
	// DryRun mode is used to test API calls, verify whether you have the required permissions on the specified resources, and check whether the request parameters are configured correctly. When DryRun mode is enabled, KMS always returns a failure with the reason for the failure. Possible failure reasons include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter were not configured.
	//
	// - ValidationError: A parameter specified in the request is invalid.
	//
	// - AccessDeniedError: You do not have permission to perform this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to retrieve the extended configuration of the secret. Valid values:
	//
	// - true: Retrieves the extended configuration.
	//
	// - false (default): Does not retrieve the extended configuration.
	//
	// > Generic secrets do not support extended configuration. If you set this parameter, it is ignored.
	//
	// example:
	//
	// true
	FetchExtendedConfig *bool `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// example:
	//
	// { "AttestationDocument":"base64-encoded-attestion-document",  "KeyEncryptionAlgorithm":"RSAES_OAEP_SHA_256" }
	Recipient *string `json:"Recipient,omitempty" xml:"Recipient,omitempty"`
	// The secret name or secret Alibaba Cloud Resource Name (ARN).
	//
	// >When accessing a secret in another Alibaba Cloud account, you must enter the secret ARN. The format of the secret ARN is `acs:kms:${region}:${account}:secret/${secret-name}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number.
	//
	// > Specifying VersionId is not supported for ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, dynamic RAM secrets, or dynamic ECS secrets. If you set this parameter, it is ignored.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version status. Default value: ACSCurrent.
	//
	// If you specify this parameter, the secret value of the specified version status is returned. If you do not specify this parameter, the secret value of the ACSCurrent version status is returned.
	//
	// > For ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, dynamic RAM secrets, and dynamic ECS secrets, only the secret values corresponding to ACSPrevious and ACSCurrent can be retrieved.
	//
	// example:
	//
	// ACSCurrent
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
}

func (s GetSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *GetSecretValueRequest) GetFetchExtendedConfig() *bool {
	return s.FetchExtendedConfig
}

func (s *GetSecretValueRequest) GetRecipient() *string {
	return s.Recipient
}

func (s *GetSecretValueRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *GetSecretValueRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetSecretValueRequest) GetVersionStage() *string {
	return s.VersionStage
}

func (s *GetSecretValueRequest) SetDryRun(v string) *GetSecretValueRequest {
	s.DryRun = &v
	return s
}

func (s *GetSecretValueRequest) SetFetchExtendedConfig(v bool) *GetSecretValueRequest {
	s.FetchExtendedConfig = &v
	return s
}

func (s *GetSecretValueRequest) SetRecipient(v string) *GetSecretValueRequest {
	s.Recipient = &v
	return s
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionId(v string) *GetSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionStage(v string) *GetSecretValueRequest {
	s.VersionStage = &v
	return s
}

func (s *GetSecretValueRequest) Validate() error {
	return dara.Validate(s)
}
