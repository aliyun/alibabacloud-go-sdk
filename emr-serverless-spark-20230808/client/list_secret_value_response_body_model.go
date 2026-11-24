// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSecretValueResponseBody
	GetRequestId() *string
	SetSecrets(v []*ListSecretValueResponseBodySecrets) *ListSecretValueResponseBody
	GetSecrets() []*ListSecretValueResponseBodySecrets
}

type ListSecretValueResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 79555076-7087-59EF-A936-8C31A48AA1E2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of secrets.
	Secrets []*ListSecretValueResponseBodySecrets `json:"secrets,omitempty" xml:"secrets,omitempty" type:"Repeated"`
}

func (s ListSecretValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretValueResponseBody) GetSecrets() []*ListSecretValueResponseBodySecrets {
	return s.Secrets
}

func (s *ListSecretValueResponseBody) SetRequestId(v string) *ListSecretValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretValueResponseBody) SetSecrets(v []*ListSecretValueResponseBodySecrets) *ListSecretValueResponseBody {
	s.Secrets = v
	return s
}

func (s *ListSecretValueResponseBody) Validate() error {
	if s.Secrets != nil {
		for _, item := range s.Secrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSecretValueResponseBodySecrets struct {
	// The comment.
	//
	// example:
	//
	// My secret。
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The secret content.
	//
	// example:
	//
	// password
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787220243000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The UID of the creator.
	//
	// example:
	//
	// 12550000006396
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The nickname of the creator.
	//
	// example:
	//
	// admin
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// The secret name.
	//
	// example:
	//
	// mysecret
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The secret ID.
	//
	// example:
	//
	// st-xxxxxxxxxxxx
	SecretId *string `json:"secretId,omitempty" xml:"secretId,omitempty"`
	// The secret type. The value is fixed to secret.
	//
	// example:
	//
	// secret
	SecretType *string `json:"secretType,omitempty" xml:"secretType,omitempty"`
}

func (s ListSecretValueResponseBodySecrets) String() string {
	return dara.Prettify(s)
}

func (s ListSecretValueResponseBodySecrets) GoString() string {
	return s.String()
}

func (s *ListSecretValueResponseBodySecrets) GetComment() *string {
	return s.Comment
}

func (s *ListSecretValueResponseBodySecrets) GetContent() *string {
	return s.Content
}

func (s *ListSecretValueResponseBodySecrets) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSecretValueResponseBodySecrets) GetCreator() *string {
	return s.Creator
}

func (s *ListSecretValueResponseBodySecrets) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListSecretValueResponseBodySecrets) GetName() *string {
	return s.Name
}

func (s *ListSecretValueResponseBodySecrets) GetSecretId() *string {
	return s.SecretId
}

func (s *ListSecretValueResponseBodySecrets) GetSecretType() *string {
	return s.SecretType
}

func (s *ListSecretValueResponseBodySecrets) SetComment(v string) *ListSecretValueResponseBodySecrets {
	s.Comment = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetContent(v string) *ListSecretValueResponseBodySecrets {
	s.Content = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetCreateTime(v int64) *ListSecretValueResponseBodySecrets {
	s.CreateTime = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetCreator(v string) *ListSecretValueResponseBodySecrets {
	s.Creator = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetCreatorName(v string) *ListSecretValueResponseBodySecrets {
	s.CreatorName = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetName(v string) *ListSecretValueResponseBodySecrets {
	s.Name = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetSecretId(v string) *ListSecretValueResponseBodySecrets {
	s.SecretId = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) SetSecretType(v string) *ListSecretValueResponseBodySecrets {
	s.SecretType = &v
	return s
}

func (s *ListSecretValueResponseBodySecrets) Validate() error {
	return dara.Validate(s)
}
