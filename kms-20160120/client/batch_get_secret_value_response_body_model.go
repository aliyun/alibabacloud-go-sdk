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
	// The list of successfully retrieved secrets.
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
	// The error message returned when the secret fails to be retrieved.
	//
	// example:
	//
	// Resource not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the secret.
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
	// - Invalid: The rotation status is abnormal, and KMS cannot automatically rotate the secret.
	//
	// > This parameter is returned only for ApsaraDB RDS secrets, PolarDB secrets, Redis/Tair secrets, RAM secrets, or ECS secrets.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created. The time is displayed in UTC.
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
	// The time when the last rotation was performed. The time is displayed in UTC.
	//
	// > This parameter is returned only when the secret has been rotated.
	//
	// example:
	//
	// 2023-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation is performed. The time is displayed in UTC.
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 2024-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the request. The ID is a unique identifier generated by Alibaba Cloud for the request and can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// 8252db58-2036-408c-a3d5-56e656dc2551
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic secret rotation.
	//
	// The value is in the `integer[unit]` format. `integer` indicates the time length, and `unit` indicates the time unit. Valid values for `unit`: s (seconds). For example, a 7-day rotation interval is 604800s.
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The secret value. KMS decrypts the stored ciphertext secret value and returns this parameter.
	//
	// - For a generic secret, the specified secret value is returned.
	//
	// - For an ApsaraDB RDS secret or a Redis/Tair secret, the secret value is in the following format: `{"AccountName":"","AccountPassword":""}`.
	//
	// - For a RAM secret, the secret value is in the following format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2023-03-25T10:42:40Z"}`.
	//
	// - For an ECS secret, the secret value is in one of the following formats:
	//
	//   - Password secret: `{"UserName":"ecs-user","Password":"H5asdasdsads****"}`.
	//
	//   - Public-private key pair secret (the private key is in PEM format): `{"UserName":"ecs-user","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`.
	//
	// - For a PolarDB secret, the secret value is in the following format: `{"AccountName":"","AccountPassword":""}`.
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
	// The name of the secret.
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
	// The version ID of the secret.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels of the secret version.
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
	// The stage labels of the secret version.
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
