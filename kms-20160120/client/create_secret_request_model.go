// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDKMSInstanceId(v string) *CreateSecretRequest
	GetDKMSInstanceId() *string
	SetDescription(v string) *CreateSecretRequest
	GetDescription() *string
	SetEnableAutomaticRotation(v bool) *CreateSecretRequest
	GetEnableAutomaticRotation() *bool
	SetEncryptionKeyId(v string) *CreateSecretRequest
	GetEncryptionKeyId() *string
	SetExtendedConfig(v map[string]interface{}) *CreateSecretRequest
	GetExtendedConfig() map[string]interface{}
	SetPolicy(v string) *CreateSecretRequest
	GetPolicy() *string
	SetRotationInterval(v string) *CreateSecretRequest
	GetRotationInterval() *string
	SetSecretData(v string) *CreateSecretRequest
	GetSecretData() *string
	SetSecretDataType(v string) *CreateSecretRequest
	GetSecretDataType() *string
	SetSecretName(v string) *CreateSecretRequest
	GetSecretName() *string
	SetSecretType(v string) *CreateSecretRequest
	GetSecretType() *string
	SetTags(v string) *CreateSecretRequest
	GetTags() *string
	SetVersionId(v string) *CreateSecretRequest
	GetVersionId() *string
}

type CreateSecretRequest struct {
	// The version number of the secret.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// Specifies whether to enable automatic rotation. Valid values:
	//
	// 	- true: specifies to enable automatic rotation.
	//
	// 	- false: specifies to disable automatic rotation. This is the default value.
	//
	// >  This parameter is valid if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// mydbinfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether automatic rotation is enabled. Valid values:
	//
	// 	- Enabled: indicates that automatic rotation is enabled.
	//
	// 	- Disabled: indicates that automatic rotation is disabled.
	//
	// 	- Invalid: indicates that the status of automatic rotation is abnormal. In this case, Secrets Manager cannot automatically rotate the secret.
	//
	// >  This parameter is returned if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// true
	EnableAutomaticRotation *bool `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	// The description of the secret.
	//
	// example:
	//
	// 00aa68af-2c02-4f68-95fe-3435d330****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// {"SecretSubType":"SingleUser", "DBInstanceId":"rm-bp1b3dd3a506e****" ,"CustomData":{}}
	ExtendedConfig map[string]interface{} `json:"ExtendedConfig,omitempty" xml:"ExtendedConfig,omitempty"`
	Policy         *string                `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// 30d
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	// The tags of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// The type of the secret. Valid values:
	//
	// 	- Generic: specifies a generic secret.
	//
	// 	- Rds: specifies a managed ApsaraDB RDS secret.
	//
	// 	- RAMCredentials: specifies a managed RAM secret.
	//
	// 	- ECS: specifies a managed ECS secret.
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The extended configuration of the secret. This parameter specifies the properties of the secret of the specific type. The description can be up to 1,024 characters in length.
	//
	// 	- If you set the SecretType parameter to Generic, you do not need to configure this parameter.
	//
	// 	- If you set the SecretType parameter to Rds, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- SingleUser: Secrets Manager manages the ApsaraDB RDS secret in single-account mode. When the secret is rotated, the password of the specified account is reset to a new random password.
	//
	//         	- DoubleUsers: Secrets Manager manages the ApsaraDB RDS secret in dual-account mode. One account is referenced by the ACSCurrent version, and the other account is referenced by the ACSPrevious version. When the secret is rotated, the password of the account referenced by the ACSPrevious version is reset to a new random password. Then, Secrets Manager switches the referenced accounts between the ACSCurrent and ACSPrevious versions.
	//
	//     	- DBInstanceId: required. The ApsaraDB RDS instance to which the ApsaraDB RDS account belongs.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). Example: `{"Key1": "v1", "fds":"fdsf"}`. The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to RAMCredentials, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Set the value to RamUserAccessKey.
	//
	//     	- UserName: required. The name of the RAM user.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// 	- If you set the SecretType parameter to ECS, configure the following fields for the ExtendedConfig parameter:
	//
	//     	- SecretSubType: required. The subtype of the secret. Valid values:
	//
	//         	- Password: the password that is used to log on to the ECS instance.
	//
	//         	- SSHKey: the SSH public key and private key that are used to log on to the ECS instance.
	//
	//     	- RegionId: required. The ID of the region in which the ECS instance resides.
	//
	//     	- InstanceId: required. The ID of the ECS instance.
	//
	//     	- CustomData: optional. The custom data. The value is a collection of key-value pairs in the JSON format. Up to 10 key-value pairs can be specified. Separate multiple key-value pairs with commas (,). The default value is a pair of empty braces (`{}`).
	//
	// >  This parameter is required if you set the SecretType parameter to Rds, RAMCredentials, or ECS.
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The value of the secret that you want to create. Secrets Manager encrypts the secret value and stores the encrypted value in the initial version.
	//
	// 	- If you set the SecretType parameter to Generic that indicates a generic secret, you can customize the secret value.
	//
	// 	- If you set the SecretType parameter to Rds that indicates a managed ApsaraDB RDS secret, the secret value must be in the format of `{"Accounts":[{"AccountName":"","AccountPassword":""}]}`. In the preceding format, `AccountName` indicates the username of the account that is used to connect to your ApsaraDB RDS instance, and `AccountPassword` specifies the password of the account.
	//
	// 	- If you set the SecretType parameter to RAMCredentials that indicates a managed RAM secret, the secret value must be in the format of `{"AccessKeys":[{"AccessKeyId":"","AccessKeySecret":"",}]}`. In the preceding format, `AccessKeyId` indicates the AccessKey ID of the RAM user and `AccessKeySecret` specifies the AccessKey secret of the RAM user. You must specify all the AccessKey pairs of the RAM user.
	//
	// 	- If you set the SecretType parameter to ECS that indicates a managed ECS secret, the secret value must be in one of the following formats:
	//
	//     	- `{"UserName":"","Password": ""}`: In the format, `UserName` specifies the username that is used to log on to the ECS instance, and `Password` specifies the password that is used to log on to the ECS instance.
	//
	//     	- `{"UserName":"","PublicKey": "", "PrivateKey": ""}`: In the format, `PublicKey` indicates the SSH public key that is used to log on to the ECS instance, and `PrivateKey` specifies the SSH private key that is used to log on to the ECS instance.
	//
	// This parameter is required.
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
	// The interval for automatic rotation. Valid values: 6 hours to 8,760 hours (365 days).
	//
	// The value is in the `integer[unit]` format.
	//
	// The unit can be d (day), h (hour), m (minute), or s (second). For example, both 7d and 604800s indicate a seven-day interval.
	//
	// >  This parameter is required if you set the EnableAutomaticRotation parameter to true. This parameter is ignored if you set the EnableAutomaticRotation parameter to false or if the EnableAutomaticRotation parameter is not configured.
	//
	// example:
	//
	// [{\\"TagKey\\":\\"key1\\",\\"TagValue\\":\\"val1\\"},{\\"TagKey\\":\\"key2\\",\\"TagValue\\":\\"val2\\"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text
	//
	// 	- binary
	//
	// >  If you set the SecretType parameter to Rds, RAMCredentials, or ECS, the SecretDataType parameter must be set to text.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *CreateSecretRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecretRequest) GetEnableAutomaticRotation() *bool {
	return s.EnableAutomaticRotation
}

func (s *CreateSecretRequest) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *CreateSecretRequest) GetExtendedConfig() map[string]interface{} {
	return s.ExtendedConfig
}

func (s *CreateSecretRequest) GetPolicy() *string {
	return s.Policy
}

func (s *CreateSecretRequest) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *CreateSecretRequest) GetSecretData() *string {
	return s.SecretData
}

func (s *CreateSecretRequest) GetSecretDataType() *string {
	return s.SecretDataType
}

func (s *CreateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretRequest) GetSecretType() *string {
	return s.SecretType
}

func (s *CreateSecretRequest) GetTags() *string {
	return s.Tags
}

func (s *CreateSecretRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateSecretRequest) SetDKMSInstanceId(v string) *CreateSecretRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetEnableAutomaticRotation(v bool) *CreateSecretRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateSecretRequest) SetEncryptionKeyId(v string) *CreateSecretRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateSecretRequest) SetExtendedConfig(v map[string]interface{}) *CreateSecretRequest {
	s.ExtendedConfig = v
	return s
}

func (s *CreateSecretRequest) SetPolicy(v string) *CreateSecretRequest {
	s.Policy = &v
	return s
}

func (s *CreateSecretRequest) SetRotationInterval(v string) *CreateSecretRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateSecretRequest) SetSecretData(v string) *CreateSecretRequest {
	s.SecretData = &v
	return s
}

func (s *CreateSecretRequest) SetSecretDataType(v string) *CreateSecretRequest {
	s.SecretDataType = &v
	return s
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetSecretType(v string) *CreateSecretRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretRequest) SetTags(v string) *CreateSecretRequest {
	s.Tags = &v
	return s
}

func (s *CreateSecretRequest) SetVersionId(v string) *CreateSecretRequest {
	s.VersionId = &v
	return s
}

func (s *CreateSecretRequest) Validate() error {
	return dara.Validate(s)
}
