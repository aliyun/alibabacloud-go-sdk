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
	// Specifies whether to return deprecated secret versions.
	//
	// Valid values:
	//
	// 	- false: no
	//
	// 	- true: yes
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the secret.
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
