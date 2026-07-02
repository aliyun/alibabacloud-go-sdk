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
	SetVersionStages(v *GetSecretValueResponseBodyVersionStages) *GetSecretValueResponseBody
	GetVersionStages() *GetSecretValueResponseBodyVersionStages
}

type GetSecretValueResponseBody struct {
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// - Enabled: Automatic rotation is enabled.
	//
	// - Disabled: Automatic rotation is disabled.
	//
	// - Invalid: The rotation status is abnormal, and KMS cannot automatically rotate the credential for you.
	//
	// > This parameter is returned only for RDS credentials, PolarDB credentials, Redis/Tair credentials, RAM credentials, or ECS credentials.
	//
	// example:
	//
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// The time when the credential was created.
	//
	// example:
	//
	// 2024-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extended configuration of the credential.
	//
	// > This parameter is returned only for RDS credentials, PolarDB credentials, Redis/Tair credentials, RAM credentials, or ECS credentials when FetchExtendedConfig is set to true.
	//
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The time of the most recent rotation.
	//
	// > This parameter is returned only if the credential has been rotated.
	//
	// example:
	//
	// 2023-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time of the next rotation.
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 2024-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The ID of the current request. Alibaba Cloud generates a unique identifier for each request, which can be used for troubleshooting and issue tracking.
	//
	// example:
	//
	// 6a3e9c36-1150-4881-84d3-eb8672fcafad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The epoch for automatic credential rotation.
	//
	// The format is `integer[unit]`, where `integer` indicates the time duration and `unit` indicates the time unit. Valid value for `unit`: s (seconds). For example, a 7-day rotation epoch is 604800s.
	//
	// > This parameter is returned only when automatic rotation is enabled.
	//
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The value of the credential. KMS decrypts the stored ciphertext and returns this parameter.
	//
	// - For generic secrets, the credential value you specified is returned.
	//
	// - For RDS credentials and Redis/Tair credentials, the credential value is in the format: `{"AccountName":"","AccountPassword":""}`.
	//
	// - For RAM credentials, the credential value is in the format: `{"AccessKeyId":"Adfdsfd","AccessKeySecret":"fdsfdsf","GenerateTimestamp": "2023-03-25T10:42:40Z"}`.
	//
	// - For ECS credentials, the credential value is in one of the following formats:
	//
	//   - Security token type: `{"UserName":"ecs-user","Password":"H5asdasdsads****"}`.
	//
	//   - Public-private key pair type (private key in PEM format): `{"UserName":"ecs-user","PublicKey":"ssh-rsa ****mKwnVix9YTFY9Rs= imported-openssh-key","PrivateKey": "d6bee1cb-2e14-4277-ba6b-73786b21****"}`.
	//
	// - For PolarDB credentials, the credential value is in the format: `{"AccountName":"","AccountPassword":""}`.
	//
	// example:
	//
	// testdata1
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The value type of the credential. Valid values:
	//
	// - text
	//
	// - binary
	//
	// example:
	//
	// binary
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the credential.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the credential. Valid values:
	//
	// - Generic: generic secret.
	//
	// - Rds: RDS credential.
	//
	// - Redis: Redis/Tair credential.
	//
	// - RAMCredentials: RAM credential.
	//
	// - ECS: ECS credential.
	//
	// - PolarDB: PolarDB credential.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The version number of the credential.
	//
	// example:
	//
	// v1
	VersionId     *string                                  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionStages *GetSecretValueResponseBodyVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" type:"Struct"`
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

func (s *GetSecretValueResponseBody) GetVersionStages() *GetSecretValueResponseBodyVersionStages {
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

func (s *GetSecretValueResponseBody) SetVersionStages(v *GetSecretValueResponseBodyVersionStages) *GetSecretValueResponseBody {
	s.VersionStages = v
	return s
}

func (s *GetSecretValueResponseBody) Validate() error {
	if s.VersionStages != nil {
		if err := s.VersionStages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecretValueResponseBodyVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" type:"Repeated"`
}

func (s GetSecretValueResponseBodyVersionStages) String() string {
	return dara.Prettify(s)
}

func (s GetSecretValueResponseBodyVersionStages) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseBodyVersionStages) GetVersionStage() []*string {
	return s.VersionStage
}

func (s *GetSecretValueResponseBodyVersionStages) SetVersionStage(v []*string) *GetSecretValueResponseBodyVersionStages {
	s.VersionStage = v
	return s
}

func (s *GetSecretValueResponseBodyVersionStages) Validate() error {
	return dara.Validate(s)
}
