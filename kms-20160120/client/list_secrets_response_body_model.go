// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSecretsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSecretsResponseBody
	GetRequestId() *string
	SetSecretList(v *ListSecretsResponseBodySecretList) *ListSecretsResponseBody
	GetSecretList() *ListSecretsResponseBodySecretList
	SetTotalCount(v int32) *ListSecretsResponseBody
	GetTotalCount() *int32
}

type ListSecretsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned secrets.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of secrets.
	//
	// example:
	//
	// 6a6287a0-ff34-4780-a790-fdfca900557f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the secret was updated.
	SecretList *ListSecretsResponseBodySecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" type:"Struct"`
	// The secret name.
	//
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretsResponseBody) GetSecretList() *ListSecretsResponseBodySecretList {
	return s.SecretList
}

func (s *ListSecretsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSecretsResponseBody) SetPageNumber(v int32) *ListSecretsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsResponseBody) SetPageSize(v int32) *ListSecretsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSecretsResponseBody) SetRequestId(v string) *ListSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponseBody) SetSecretList(v *ListSecretsResponseBodySecretList) *ListSecretsResponseBody {
	s.SecretList = v
	return s
}

func (s *ListSecretsResponseBody) SetTotalCount(v int32) *ListSecretsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretsResponseBody) Validate() error {
	if s.SecretList != nil {
		if err := s.SecretList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretsResponseBodySecretList struct {
	Secret []*ListSecretsResponseBodySecretListSecret `json:"Secret,omitempty" xml:"Secret,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodySecretList) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecretList) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretList) GetSecret() []*ListSecretsResponseBodySecretListSecret {
	return s.Secret
}

func (s *ListSecretsResponseBodySecretList) SetSecret(v []*ListSecretsResponseBodySecretListSecret) *ListSecretsResponseBodySecretList {
	s.Secret = v
	return s
}

func (s *ListSecretsResponseBodySecretList) Validate() error {
	if s.Secret != nil {
		for _, item := range s.Secret {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecretsResponseBodySecretListSecret struct {
	// The tag value.
	//
	// example:
	//
	// 2022-07-17T07:59:05Z
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	OwingService *string `json:"OwingService,omitempty" xml:"OwingService,omitempty"`
	// The resource tags of the secret.
	//
	// This parameter is not returned if you set the FetchTags parameter to false or do not specify the FetchTags parameter.
	//
	// example:
	//
	// 2022-08-17T07:59:05Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The type of the secret. Valid values:
	//
	// 	- Generic: indicates a generic secret.
	//
	// 	- Rds: indicates a managed ApsaraDB RDS secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The time when the secret was created.
	//
	// example:
	//
	// Generic
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The tag key.
	Tags *ListSecretsResponseBodySecretListSecretTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-07-17T07:59:05Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecret) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecret) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecret) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSecretsResponseBodySecretListSecret) GetOwingService() *string {
	return s.OwingService
}

func (s *ListSecretsResponseBodySecretListSecret) GetPlannedDeleteTime() *string {
	return s.PlannedDeleteTime
}

func (s *ListSecretsResponseBodySecretListSecret) GetSecretName() *string {
	return s.SecretName
}

func (s *ListSecretsResponseBodySecretListSecret) GetSecretType() *string {
	return s.SecretType
}

func (s *ListSecretsResponseBodySecretListSecret) GetTags() *ListSecretsResponseBodySecretListSecretTags {
	return s.Tags
}

func (s *ListSecretsResponseBodySecretListSecret) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSecretsResponseBodySecretListSecret) SetCreateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.CreateTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetOwingService(v string) *ListSecretsResponseBodySecretListSecret {
	s.OwingService = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetPlannedDeleteTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.PlannedDeleteTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretName(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetSecretType(v string) *ListSecretsResponseBodySecretListSecret {
	s.SecretType = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetTags(v *ListSecretsResponseBodySecretListSecretTags) *ListSecretsResponseBodySecretListSecret {
	s.Tags = v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) SetUpdateTime(v string) *ListSecretsResponseBodySecretListSecret {
	s.UpdateTime = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecret) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretsResponseBodySecretListSecretTags struct {
	Tag []*ListSecretsResponseBodySecretListSecretTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodySecretListSecretTags) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecretTags) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecretTags) GetTag() []*ListSecretsResponseBodySecretListSecretTagsTag {
	return s.Tag
}

func (s *ListSecretsResponseBodySecretListSecretTags) SetTag(v []*ListSecretsResponseBodySecretListSecretTagsTag) *ListSecretsResponseBodySecretListSecretTags {
	s.Tag = v
	return s
}

func (s *ListSecretsResponseBodySecretListSecretTags) Validate() error {
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

type ListSecretsResponseBodySecretListSecretTagsTag struct {
	// example:
	//
	// key1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// val1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecretListSecretTagsTag) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagKey(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) SetTagValue(v string) *ListSecretsResponseBodySecretListSecretTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListSecretsResponseBodySecretListSecretTagsTag) Validate() error {
	return dara.Validate(s)
}
