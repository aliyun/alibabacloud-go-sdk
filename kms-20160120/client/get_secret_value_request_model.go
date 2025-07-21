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
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to obtain the extended configuration of the secret. Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// >  This parameter is ignored for a generic secret.
	//
	// example:
	//
	// true
	FetchExtendedConfig *bool `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret value. If you specify this parameter, Secrets Manager returns the secret value of the specified version.
	//
	// >  This parameter is ignored for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// 00000000000000000000000000000001
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage label that marks the secret version. If you specify this parameter, Secrets Manager returns the secret value of the version that is marked with the specified stage label.
	//
	// Default value: ACSCurrent.
	//
	// >  For a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret, Secrets Manager can return only the secret value of the version marked with ACSPrevious or ACSCurrent.
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
