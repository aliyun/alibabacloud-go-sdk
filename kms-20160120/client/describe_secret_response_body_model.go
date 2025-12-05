// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *DescribeSecretResponseBody
	GetArn() *string
	SetAutomaticRotation(v string) *DescribeSecretResponseBody
	GetAutomaticRotation() *string
	SetCreateTime(v string) *DescribeSecretResponseBody
	GetCreateTime() *string
	SetDKMSInstanceId(v string) *DescribeSecretResponseBody
	GetDKMSInstanceId() *string
	SetDescription(v string) *DescribeSecretResponseBody
	GetDescription() *string
	SetEncryptionKeyId(v string) *DescribeSecretResponseBody
	GetEncryptionKeyId() *string
	SetExtendedConfig(v string) *DescribeSecretResponseBody
	GetExtendedConfig() *string
	SetLastRotationDate(v string) *DescribeSecretResponseBody
	GetLastRotationDate() *string
	SetNextRotationDate(v string) *DescribeSecretResponseBody
	GetNextRotationDate() *string
	SetOwingService(v string) *DescribeSecretResponseBody
	GetOwingService() *string
	SetPlannedDeleteTime(v string) *DescribeSecretResponseBody
	GetPlannedDeleteTime() *string
	SetRequestId(v string) *DescribeSecretResponseBody
	GetRequestId() *string
	SetRotationInterval(v string) *DescribeSecretResponseBody
	GetRotationInterval() *string
	SetSecretName(v string) *DescribeSecretResponseBody
	GetSecretName() *string
	SetSecretType(v string) *DescribeSecretResponseBody
	GetSecretType() *string
	SetTags(v *DescribeSecretResponseBodyTags) *DescribeSecretResponseBody
	GetTags() *DescribeSecretResponseBodyTags
	SetUpdateTime(v string) *DescribeSecretResponseBody
	GetUpdateTime() *string
}

type DescribeSecretResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/secret001
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
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
	// 2022-02-21T15:39:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the dedicated KMS instance.
	//
	// example:
	//
	// kst-bjj62d8f5e0sgtx8h****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the secret.
	//
	// example:
	//
	// userinfo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the customer master key (CMK) that is used to encrypt the secret value.
	//
	// example:
	//
	// 00aa68af-2c02-4f68-95fe-3435d330****
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	// The extended configuration of the secret.
	//
	// >  This parameter is returned only for a managed ApsaraDB RDS secret, a managed Resource Access Management (RAM) secret, or a managed Elastic Compute Service (ECS) secret.
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
	// 2022-07-05T08:22:03Z
	LastRotationDate *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty"`
	// The time when the next rotation will be performed.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	//
	// example:
	//
	// 2022-07-06T18:22:03Z
	NextRotationDate *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty"`
	OwingService     *string `json:"OwingService,omitempty" xml:"OwingService,omitempty"`
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-03-21T15:45:12Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 93348dfb-3627-4417-8d90-487a76a909c9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The interval for automatic rotation.
	//
	// The value is in the `integer[unit]` format. `integer` indicates the length of time. `unit`: indicates the time unit. The value of `unit` is fixed as s. For example, if the value is 604800s, automatic rotation is performed at a 7-day interval.
	//
	// >  This parameter is returned when automatic rotation is enabled.
	//
	// example:
	//
	// 3153600s
	RotationInterval *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
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
	// Rds
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or you do not specify the FetchTags parameter.
	Tags *DescribeSecretResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret was updated.
	//
	// example:
	//
	// 2022-02-21T15:39:26Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBody) GetArn() *string {
	return s.Arn
}

func (s *DescribeSecretResponseBody) GetAutomaticRotation() *string {
	return s.AutomaticRotation
}

func (s *DescribeSecretResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSecretResponseBody) GetDKMSInstanceId() *string {
	return s.DKMSInstanceId
}

func (s *DescribeSecretResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecretResponseBody) GetEncryptionKeyId() *string {
	return s.EncryptionKeyId
}

func (s *DescribeSecretResponseBody) GetExtendedConfig() *string {
	return s.ExtendedConfig
}

func (s *DescribeSecretResponseBody) GetLastRotationDate() *string {
	return s.LastRotationDate
}

func (s *DescribeSecretResponseBody) GetNextRotationDate() *string {
	return s.NextRotationDate
}

func (s *DescribeSecretResponseBody) GetOwingService() *string {
	return s.OwingService
}

func (s *DescribeSecretResponseBody) GetPlannedDeleteTime() *string {
	return s.PlannedDeleteTime
}

func (s *DescribeSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecretResponseBody) GetRotationInterval() *string {
	return s.RotationInterval
}

func (s *DescribeSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeSecretResponseBody) GetSecretType() *string {
	return s.SecretType
}

func (s *DescribeSecretResponseBody) GetTags() *DescribeSecretResponseBodyTags {
	return s.Tags
}

func (s *DescribeSecretResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSecretResponseBody) SetArn(v string) *DescribeSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *DescribeSecretResponseBody) SetAutomaticRotation(v string) *DescribeSecretResponseBody {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeSecretResponseBody) SetCreateTime(v string) *DescribeSecretResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetDKMSInstanceId(v string) *DescribeSecretResponseBody {
	s.DKMSInstanceId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetDescription(v string) *DescribeSecretResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecretResponseBody) SetEncryptionKeyId(v string) *DescribeSecretResponseBody {
	s.EncryptionKeyId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetExtendedConfig(v string) *DescribeSecretResponseBody {
	s.ExtendedConfig = &v
	return s
}

func (s *DescribeSecretResponseBody) SetLastRotationDate(v string) *DescribeSecretResponseBody {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetNextRotationDate(v string) *DescribeSecretResponseBody {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeSecretResponseBody) SetOwingService(v string) *DescribeSecretResponseBody {
	s.OwingService = &v
	return s
}

func (s *DescribeSecretResponseBody) SetPlannedDeleteTime(v string) *DescribeSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRequestId(v string) *DescribeSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRotationInterval(v string) *DescribeSecretResponseBody {
	s.RotationInterval = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretName(v string) *DescribeSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSecretType(v string) *DescribeSecretResponseBody {
	s.SecretType = &v
	return s
}

func (s *DescribeSecretResponseBody) SetTags(v *DescribeSecretResponseBodyTags) *DescribeSecretResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeSecretResponseBody) SetUpdateTime(v string) *DescribeSecretResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSecretResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecretResponseBodyTags struct {
	Tag []*DescribeSecretResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSecretResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyTags) GetTag() []*DescribeSecretResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeSecretResponseBodyTags) SetTag(v []*DescribeSecretResponseBodyTagsTag) *DescribeSecretResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeSecretResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecretResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// val1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeSecretResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeSecretResponseBodyTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagKey(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSecretResponseBodyTagsTag) SetTagValue(v string) *DescribeSecretResponseBodyTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeSecretResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
