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
	SetSecretName(v string) *GetSecretValueRequest
	GetSecretName() *string
	SetVersionId(v string) *GetSecretValueRequest
	GetVersionId() *string
	SetVersionStage(v string) *GetSecretValueRequest
	GetVersionStage() *string
}

type GetSecretValueRequest struct {
	// Indicates whether to enable DryRun mode.
	//
	// - true: Enabled
	//
	// - false (Default Value): Disabled
	//
	// DryRun mode is used for Testing API Calls to authenticate whether you have the required permissions on the specified resource and whether the Request Parameters are correctly configured. When DryRun mode is enabled, KMS always returns a failed response along with the failure reason. Possible failure reasons include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter were not specified.
	//
	// - ValidationError: One or more parameters in the request are invalid.
	//
	// - AccessDeniedError: You do not have permission to execute this operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Indicates whether to retrieve the extended configuration of the credential. Valid values:
	//
	// - true: Retrieve
	//
	// - false (Default Value): Do not retrieve
	//
	// > Generic secrets do not support extended configuration. If you specify this parameter, it will be ignored.
	//
	// example:
	//
	// true
	FetchExtendedConfig *bool `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// The name or ARN of the credential.
	//
	// > When accessing a credential under another Alibaba Cloud account, you must specify the credential ARN. The ARN format is `acs:kms:${region}:${account}:secret/${secret-name}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// Version number.
	//
	// > The VersionId parameter is not supported for RDS credentials, PolarDB credentials, Redis/Tair credentials, RAM credentials, and ECS credentials. If you specify this parameter, it will be ignored.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version stage. Default value: ACSCurrent.
	//
	// If you specify this parameter, the credential value of the specified version stage is returned. If you do not specify this parameter, the credential value of the ACSCurrent version stage is returned.
	//
	// > For RDS credentials, PolarDB credentials, Redis/Tair credentials, RAM credentials, and ECS credentials, you can retrieve only the credential values corresponding to the ACSPrevious or ACSCurrent version stages.
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
