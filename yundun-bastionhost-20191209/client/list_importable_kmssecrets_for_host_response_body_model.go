// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImportableKMSSecretsForHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListImportableKMSSecretsForHostResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListImportableKMSSecretsForHostResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImportableKMSSecretsForHostResponseBody
	GetRequestId() *string
	SetSecrets(v []*ListImportableKMSSecretsForHostResponseBodySecrets) *ListImportableKMSSecretsForHostResponseBody
	GetSecrets() []*ListImportableKMSSecretsForHostResponseBodySecrets
}

type ListImportableKMSSecretsForHostResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Secrets   []*ListImportableKMSSecretsForHostResponseBodySecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ListImportableKMSSecretsForHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImportableKMSSecretsForHostResponseBody) GoString() string {
	return s.String()
}

func (s *ListImportableKMSSecretsForHostResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListImportableKMSSecretsForHostResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImportableKMSSecretsForHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImportableKMSSecretsForHostResponseBody) GetSecrets() []*ListImportableKMSSecretsForHostResponseBodySecrets {
	return s.Secrets
}

func (s *ListImportableKMSSecretsForHostResponseBody) SetMaxResults(v int32) *ListImportableKMSSecretsForHostResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBody) SetNextToken(v string) *ListImportableKMSSecretsForHostResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBody) SetRequestId(v string) *ListImportableKMSSecretsForHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBody) SetSecrets(v []*ListImportableKMSSecretsForHostResponseBodySecrets) *ListImportableKMSSecretsForHostResponseBody {
	s.Secrets = v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBody) Validate() error {
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

type ListImportableKMSSecretsForHostResponseBodySecrets struct {
	// example:
	//
	// test1
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// example:
	//
	// ECS
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// example:
	//
	// test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListImportableKMSSecretsForHostResponseBodySecrets) String() string {
	return dara.Prettify(s)
}

func (s ListImportableKMSSecretsForHostResponseBodySecrets) GoString() string {
	return s.String()
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) GetSecretName() *string {
	return s.SecretName
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) GetSecretType() *string {
	return s.SecretType
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) GetTags() *string {
	return s.Tags
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) SetSecretName(v string) *ListImportableKMSSecretsForHostResponseBodySecrets {
	s.SecretName = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) SetSecretType(v string) *ListImportableKMSSecretsForHostResponseBodySecrets {
	s.SecretType = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) SetTags(v string) *ListImportableKMSSecretsForHostResponseBodySecrets {
	s.Tags = &v
	return s
}

func (s *ListImportableKMSSecretsForHostResponseBodySecrets) Validate() error {
	return dara.Validate(s)
}
