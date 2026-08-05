// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrors(v []*BatchGetSecretValueResponseBodyErrors) *BatchGetSecretValueResponseBody
	GetErrors() []*BatchGetSecretValueResponseBodyErrors
	SetRequestId(v string) *BatchGetSecretValueResponseBody
	GetRequestId() *string
	SetSecretValues(v []*BatchGetSecretValueResponseBodySecretValues) *BatchGetSecretValueResponseBody
	GetSecretValues() []*BatchGetSecretValueResponseBodySecretValues
}

type BatchGetSecretValueResponseBody struct {
	// The list of errors for secrets that failed to be retrieved.
	Errors []*BatchGetSecretValueResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8252db58-2036-408c-a3d5-56e656dc2551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of secrets that are successfully retrieved.
	SecretValues []*BatchGetSecretValueResponseBodySecretValues `json:"SecretValues,omitempty" xml:"SecretValues,omitempty" type:"Repeated"`
}

func (s BatchGetSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueResponseBody) GetErrors() []*BatchGetSecretValueResponseBodyErrors {
	return s.Errors
}

func (s *BatchGetSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetSecretValueResponseBody) GetSecretValues() []*BatchGetSecretValueResponseBodySecretValues {
	return s.SecretValues
}

func (s *BatchGetSecretValueResponseBody) SetErrors(v []*BatchGetSecretValueResponseBodyErrors) *BatchGetSecretValueResponseBody {
	s.Errors = v
	return s
}

func (s *BatchGetSecretValueResponseBody) SetRequestId(v string) *BatchGetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetSecretValueResponseBody) SetSecretValues(v []*BatchGetSecretValueResponseBodySecretValues) *BatchGetSecretValueResponseBody {
	s.SecretValues = v
	return s
}

func (s *BatchGetSecretValueResponseBody) Validate() error {
	if s.Errors != nil {
		for _, item := range s.Errors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecretValues != nil {
		for _, item := range s.SecretValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetSecretValueResponseBodyErrors struct {
	// The service error code.
	//
	// example:
	//
	// Forbidden.ResourceNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned for the error.
	//
	// example:
	//
	// 404
	HttpStatus *int32 `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	// The error message for the failed secret retrieval.
	//
	// example:
	//
	// Resource not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The secret name.
	//
	// example:
	//
	// secret002
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s BatchGetSecretValueResponseBodyErrors) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueResponseBodyErrors) GetCode() *string {
	return s.Code
}

func (s *BatchGetSecretValueResponseBodyErrors) GetHttpStatus() *int32 {
	return s.HttpStatus
}

func (s *BatchGetSecretValueResponseBodyErrors) GetMessage() *string {
	return s.Message
}

func (s *BatchGetSecretValueResponseBodyErrors) GetSecretName() *string {
	return s.SecretName
}

func (s *BatchGetSecretValueResponseBodyErrors) SetCode(v string) *BatchGetSecretValueResponseBodyErrors {
	s.Code = &v
	return s
}

func (s *BatchGetSecretValueResponseBodyErrors) SetHttpStatus(v int32) *BatchGetSecretValueResponseBodyErrors {
	s.HttpStatus = &v
	return s
}

func (s *BatchGetSecretValueResponseBodyErrors) SetMessage(v string) *BatchGetSecretValueResponseBodyErrors {
	s.Message = &v
	return s
}

func (s *BatchGetSecretValueResponseBodyErrors) SetSecretName(v string) *BatchGetSecretValueResponseBodyErrors {
	s.SecretName = &v
	return s
}

func (s *BatchGetSecretValueResponseBodyErrors) Validate() error {
	return dara.Validate(s)
}

type BatchGetSecretValueResponseBodySecretValues struct {
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// - Enabled: Automatic rotation is enabled.
	//
	// - Disabled: Automatic rotation is disabled.
	//
	// - Invalid: The rotation status is abnormal. KMS cannot automatically rotate the secret.
	//
	// > This parameter is returned only for ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, RAM secrets, or ECS secrets.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created (UTC).
	//
	// example:
	//
	// 2024-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended configuration of the secret.
	//
	// > This parameter is returned only for ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, RAM secrets, or ECS secrets when FetchExtendedConfig is set to true.
	//
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time of the most recent rotation (UTC).
	//
	// > This parameter is returned only when the secret has been rotated.
	//
	// example:
	//
	// 2023-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time of the next rotation (UTC).
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 2024-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of this request, which is a unique identifier generated by Alibaba Cloud for troubleshooting purposes.
	//
	// example:
	//
	// 8252db58-2036-408c-a3d5-56e656dc2551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation of the secret.
	//
	// The format is integer[unit], where integer specifies the duration and unit specifies the time unit. The value of unit is s (seconds). For example, a rotation interval of 7 days is 604800s.
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The secret value. KMS decrypts the stored ciphertext secret value and returns this parameter.
	//
	// - A generic secret returns the secret value that you specified.
	//
	// - An ApsaraDB RDS secret or Redis/Tair secret returns the secret value in the following format: `{"AccountName":"","AccountPassword":""}`.
	//
	// - A RAM secret returns the secret value in the following format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2023-03-25T10:42:40Z"}`.
	//
	// - An ECS secret returns the secret value in the following formats:
	//
	//   - Password-type secret: `{"UserName":"ecs-user","Password":"H5asdasdsads****"}`.
	//
	//   - Public-private key-type secret (private key in PEM format): `{"UserName":"ecs-user","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`.
	//
	// - A PolarDB secret returns the secret value in the following format: `{"AccountName":"","AccountPassword":""}`.
	//
	// example:
	//
	// testdata1
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// - text
	//
	// - binary
	//
	// example:
	//
	// binary
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The secret name.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the secret. Valid values:
	//
	// - Generic: generic secret.
	//
	// - Rds: ApsaraDB RDS secret.
	//
	// - Redis: Redis/Tair secret.
	//
	// - RAMCredentials: RAM secret.
	//
	// - ECS: ECS secret.
	//
	// - PolarDB: PolarDB secret.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the secret.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version stages of the secret.
	VersionStages *BatchGetSecretValueResponseBodySecretValuesVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
}

func (s BatchGetSecretValueResponseBodySecretValues) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueResponseBodySecretValues) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetLastRotationDate() *string {
	return s.LastRotationDate
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetSecretData() *string {
	return s.SecretData
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetSecretDataType() *string {
	return s.SecretDataType
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetSecretName() *string {
	return s.SecretName
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetSecretType() *string {
	return s.SecretType
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetVersionId() *string {
	return s.VersionId
}

func (s *BatchGetSecretValueResponseBodySecretValues) GetVersionStages() *BatchGetSecretValueResponseBodySecretValuesVersionStages {
	return s.VersionStages
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetAutomaticRotation(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.AutomaticRotation = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetCreateTime(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.CreateTime = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetExtendedConfig(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.ExtendedConfig = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetLastRotationDate(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.LastRotationDate = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetNextRotationDate(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.NextRotationDate = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetRequestId(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.RequestId = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetRotationInterval(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.RotationInterval = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetSecretData(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.SecretData = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetSecretDataType(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.SecretDataType = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetSecretName(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.SecretName = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetSecretType(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.SecretType = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetVersionId(v string) *BatchGetSecretValueResponseBodySecretValues {
	s.VersionId = &v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) SetVersionStages(v *BatchGetSecretValueResponseBodySecretValuesVersionStages) *BatchGetSecretValueResponseBodySecretValues {
	s.VersionStages = v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValues) Validate() error {
	if s.VersionStages != nil {
		if err := s.VersionStages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetSecretValueResponseBodySecretValuesVersionStages struct {
	// The version stages of the secret.
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s BatchGetSecretValueResponseBodySecretValuesVersionStages) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueResponseBodySecretValuesVersionStages) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueResponseBodySecretValuesVersionStages) GetVersionStage() []*string {
	return s.VersionStage
}

func (s *BatchGetSecretValueResponseBodySecretValuesVersionStages) SetVersionStage(v []*string) *BatchGetSecretValueResponseBodySecretValuesVersionStages {
	s.VersionStage = v
	return s
}

func (s *BatchGetSecretValueResponseBodySecretValuesVersionStages) Validate() error {
	return dara.Validate(s)
}
