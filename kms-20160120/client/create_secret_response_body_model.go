// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *CreateSecretResponseBody
	GetArn() *string
	SetAutomaticRotation(v string) *CreateSecretResponseBody
	GetAutomaticRotation() *string
	SetDKMSInstanceId(v string) *CreateSecretResponseBody
	GetDKMSInstanceId() *string
	SetExtendedConfig(v string) *CreateSecretResponseBody
	GetExtendedConfig() *string
	SetNextRotationDate(v string) *CreateSecretResponseBody
	GetNextRotationDate() *string
	SetRequestId(v string) *CreateSecretResponseBody
	GetRequestId() *string
	SetRotationInterval(v string) *CreateSecretResponseBody
	GetRotationInterval() *string
	SetSecretName(v string) *CreateSecretResponseBody
	GetSecretName() *string
	SetSecretType(v string) *CreateSecretResponseBody
	GetSecretType() *string
	SetVersionId(v string) *CreateSecretResponseBody
	GetVersionId() *string
}

type CreateSecretResponseBody struct {
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/mydbconninfo
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
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
	// Enabled
	AutomaticRotation *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty"`
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// example:
	//
	// {\\"SecretSubType\\":\\"SingleUser\\", \\"DBInstanceId\\":\\"rm-uf667446pc955****\\",  \\"CustomData\\":{} }
	ExtendedConfig *string `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// 2022-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// 3bf02f7a-015b-4f93-be0f-cc043fda2dd3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 604800s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. The value of the `unit` field is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned if automatic rotation is enabled.
	//
	// example:
	//
	// mydbconninfo
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretResponseBody) GetArn() *string {
	return s.Arn
}

func (s *CreateSecretResponseBody) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *CreateSecretResponseBody) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateSecretResponseBody) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *CreateSecretResponseBody) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *CreateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecretResponseBody) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *CreateSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretResponseBody) GetSecretType() *string {
	return s.SecretType
}

func (s *CreateSecretResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateSecretResponseBody) SetArn(v string) *CreateSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *CreateSecretResponseBody) SetAutomaticRotation(v string) *CreateSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateSecretResponseBody) SetDKMSInstanceId(v string) *CreateSecretResponseBody {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretResponseBody) SetExtendedConfig(v string) *CreateSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *CreateSecretResponseBody) SetNextRotationDate(v string) *CreateSecretResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *CreateSecretResponseBody) SetRequestId(v string) *CreateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponseBody) SetRotationInterval(v string) *CreateSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretName(v string) *CreateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *CreateSecretResponseBody) SetSecretType(v string) *CreateSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *CreateSecretResponseBody) SetVersionId(v string) *CreateSecretResponseBody {
	s.VersionId = &v
	return s
}

func (s *CreateSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
