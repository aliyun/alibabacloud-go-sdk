// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutomaticRotation(v string) *GetSecretValueResponseBody
	GetAutomaticRotation() *string
	SetCreateTime(v string) *GetSecretValueResponseBody
	GetCreateTime() *string
	SetExtendedConfig(v string) *GetSecretValueResponseBody
	GetExtendedConfig() *string
	SetLastRotationDate(v string) *GetSecretValueResponseBody
	GetLastRotationDate() *string
	SetNextRotationDate(v string) *GetSecretValueResponseBody
	GetNextRotationDate() *string
	SetRequestId(v string) *GetSecretValueResponseBody
	GetRequestId() *string
	SetRotationInterval(v string) *GetSecretValueResponseBody
	GetRotationInterval() *string
	SetSecretData(v string) *GetSecretValueResponseBody
	GetSecretData() *string
	SetSecretDataType(v string) *GetSecretValueResponseBody
	GetSecretDataType() *string
	SetSecretName(v string) *GetSecretValueResponseBody
	GetSecretName() *string
	SetSecretType(v string) *GetSecretValueResponseBody
	GetSecretType() *string
	SetVersionId(v string) *GetSecretValueResponseBody
	GetVersionId() *string
	SetVersionStages(v []*string) *GetSecretValueResponseBody
	GetVersionStages() []*string
}

type GetSecretValueResponseBody struct {
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the secret was created.
	//
	// example:
	//
	// 2020-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the FetchExtendedConfig parameter to true. This parameter is returned only for a managed ApsaraDB RDS secret, a managed RAM secret, or a managed ECS secret.
	//
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time when the last rotation was performed.
	//
	// >  This parameter is returned if the secret was rotated.
	//
	// example:
	//
	// 2020-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 2020-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6a3e9c36-1150-4881-84d3-eb8672fcafad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The `unit` field has a fixed value of s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The secret value. Secrets Manager decrypts the ciphertext of the secret value and returns the plaintext of the secret value in this parameter.
	//
	// 	- For a generic secret, the secret value of the specified version is returned.
	//
	// 	- For a managed ApsaraDB RDS secret, the value is returned in the following format:`{"AccountName":"","AccountPassword":""}` .
	//
	// 	- For a managed RAM secret, the secret value is returned in the following format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2016-03-25T10:42:40Z"}`.
	//
	// 	- For a managed ECS secret, the secret value is returned in one of the following formats:
	//
	//     	- `{"UserName":"root","Password":"H5asdasdsads****"}`: The secret value is returned in this format if the ECS secret is a password.
	//
	//     	- `{"UserName":"root","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`: The secret value is returned in this format is the ECS secret is a pair of SSH keys. The private key is in the Privacy Enhanced Mail (PEM) format.
	//
	// example:
	//
	// testdata1
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text
	//
	// 	- binary
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
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: indicates a managed RAM secret.
	//
	// 	- ECS: indicates a managed ECS secret.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the secret value.
	//
	// example:
	//
	// 00000000000000000000000000000001
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that mark the secret versions.
	VersionStages []*string `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Repeated"`
}

func (s GetSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBody) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *GetSecretValueResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSecretValueResponseBody) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *GetSecretValueResponseBody) GetLastRotationDate() *string {
	return s.LastRotationDate
}

func (s *GetSecretValueResponseBody) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *GetSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecretValueResponseBody) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *GetSecretValueResponseBody) GetSecretData() *string {
	return s.SecretData
}

func (s *GetSecretValueResponseBody) GetSecretDataType() *string {
	return s.SecretDataType
}

func (s *GetSecretValueResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *GetSecretValueResponseBody) GetSecretType() *string {
	return s.SecretType
}

func (s *GetSecretValueResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *GetSecretValueResponseBody) GetVersionStages() []*string {
	return s.VersionStages
}

func (s *GetSecretValueResponseBody) SetAutomaticRotation(v string) *GetSecretValueResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *GetSecretValueResponseBody) SetCreateTime(v string) *GetSecretValueResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetSecretValueResponseBody) SetExtendedConfig(v string) *GetSecretValueResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *GetSecretValueResponseBody) SetLastRotationDate(v string) *GetSecretValueResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetNextRotationDate(v string) *GetSecretValueResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRequestId(v string) *GetSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetRotationInterval(v string) *GetSecretValueResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretData(v string) *GetSecretValueResponseBody {
	s.SecretData = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretDataType(v string) *GetSecretValueResponseBody {
	s.SecretDataType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretName(v string) *GetSecretValueResponseBody {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponseBody) SetSecretType(v string) *GetSecretValueResponseBody {
	s.SecretType = &v
	return s
}

func (s *GetSecretValueResponseBody) SetVersionId(v string) *GetSecretValueResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueResponseBody) SetVersionStages(v []*string) *GetSecretValueResponseBody {
	s.VersionStages = v
	return s
}

func (s *GetSecretValueResponseBody) Validate() error {
	return dara.Validate(s)
}
