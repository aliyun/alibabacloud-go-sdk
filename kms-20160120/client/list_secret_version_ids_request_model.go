// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretVersionIdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeDeprecated(v string) *ListSecretVersionIdsRequest
	GetIncludeDeprecated() *string
	SetPageNumber(v int32) *ListSecretVersionIdsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretVersionIdsRequest
	GetPageSize() *int32
	SetSecretName(v string) *ListSecretVersionIdsRequest
	GetSecretName() *string
}

type ListSecretVersionIdsRequest struct {
	// Specifies whether to include credential versions that have no version stages in the response.
	//
	// Valid values:
	//
	// - false (default): No
	//
	// - true: Yes
	//
	// example:
	//
	// false
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	// The number of the page to return for a paged query. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name or Alibaba Cloud Resource Name (ARN) of the credential.
	//
	// > When you access a credential that belongs to another Alibaba Cloud account, you must specify the ARN of the credential. The ARN of a credential is in the format of `acs:kms:${region}:${account}:secret/${secret-name}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s ListSecretVersionIdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretVersionIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsRequest) GetIncludeDeprecated() *string {
	return s.IncludeDeprecated
}

func (s *ListSecretVersionIdsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretVersionIdsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretVersionIdsRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *ListSecretVersionIdsRequest) SetIncludeDeprecated(v string) *ListSecretVersionIdsRequest {
	s.IncludeDeprecated = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageNumber(v int32) *ListSecretVersionIdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageSize(v int32) *ListSecretVersionIdsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetSecretName(v string) *ListSecretVersionIdsRequest {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsRequest) Validate() error {
	return dara.Validate(s)
}
