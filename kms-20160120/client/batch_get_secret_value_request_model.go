// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretsList(v []*BatchGetSecretValueRequestSecretsList) *BatchGetSecretValueRequest
	GetSecretsList() []*BatchGetSecretValueRequestSecretsList
}

type BatchGetSecretValueRequest struct {
	// The list of secret information. You can query up to 20 different secrets at a time.
	SecretsList []*BatchGetSecretValueRequestSecretsList `json:"SecretsList,omitempty" xml:"SecretsList,omitempty" type:"Repeated"`
}

func (s BatchGetSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueRequest) GetSecretsList() []*BatchGetSecretValueRequestSecretsList {
	return s.SecretsList
}

func (s *BatchGetSecretValueRequest) SetSecretsList(v []*BatchGetSecretValueRequestSecretsList) *BatchGetSecretValueRequest {
	s.SecretsList = v
	return s
}

func (s *BatchGetSecretValueRequest) Validate() error {
	if s.SecretsList != nil {
		for _, item := range s.SecretsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetSecretValueRequestSecretsList struct {
	// Specifies whether to retrieve the extended configuration of the secret. Valid values:
	//
	// - true: Retrieves the extended configuration.
	//
	// - false: Does not retrieve the extended configuration. This is the default value.
	//
	// > Generic secrets do not support extended configurations. If you specify this parameter for a generic secret, it is ignored.
	//
	// example:
	//
	// true
	FetchExtendedConfig *string `json:"FetchExtendedConfig,omitempty" xml:"FetchExtendedConfig,omitempty"`
	// The name or Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// > To access a secret that belongs to another Alibaba Cloud account, you must specify the ARN of the secret. The ARN is in the `acs:kms:${region}:${account}:secret/${secret-name}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version ID of the secret.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage of the secret version. Default value: ACSCurrent.
	//
	// If you specify this parameter, the secret value of the specified version stage is returned. If you do not specify this parameter, the secret value of the ACSCurrent version stage is returned.
	//
	// > For ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, RAM secrets, and ECS secrets, you can only retrieve the secret values of the versions marked with ACSPrevious and ACSCurrent.
	//
	// example:
	//
	// ACSCurrent
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
}

func (s BatchGetSecretValueRequestSecretsList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueRequestSecretsList) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueRequestSecretsList) GetFetchExtendedConfig() *string {
	return s.FetchExtendedConfig
}

func (s *BatchGetSecretValueRequestSecretsList) GetSecretName() *string {
	return s.SecretName
}

func (s *BatchGetSecretValueRequestSecretsList) GetVersionId() *string {
	return s.VersionId
}

func (s *BatchGetSecretValueRequestSecretsList) GetVersionStage() *string {
	return s.VersionStage
}

func (s *BatchGetSecretValueRequestSecretsList) SetFetchExtendedConfig(v string) *BatchGetSecretValueRequestSecretsList {
	s.FetchExtendedConfig = &v
	return s
}

func (s *BatchGetSecretValueRequestSecretsList) SetSecretName(v string) *BatchGetSecretValueRequestSecretsList {
	s.SecretName = &v
	return s
}

func (s *BatchGetSecretValueRequestSecretsList) SetVersionId(v string) *BatchGetSecretValueRequestSecretsList {
	s.VersionId = &v
	return s
}

func (s *BatchGetSecretValueRequestSecretsList) SetVersionStage(v string) *BatchGetSecretValueRequestSecretsList {
	s.VersionStage = &v
	return s
}

func (s *BatchGetSecretValueRequestSecretsList) Validate() error {
	return dara.Validate(s)
}
