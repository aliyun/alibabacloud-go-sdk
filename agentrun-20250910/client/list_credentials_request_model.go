// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialAuthType(v string) *ListCredentialsRequest
	GetCredentialAuthType() *string
	SetCredentialName(v string) *ListCredentialsRequest
	GetCredentialName() *string
	SetCredentialSourceType(v string) *ListCredentialsRequest
	GetCredentialSourceType() *string
	SetEnabled(v bool) *ListCredentialsRequest
	GetEnabled() *bool
	SetPageNumber(v int32) *ListCredentialsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCredentialsRequest
	GetPageSize() *int32
	SetProvider(v string) *ListCredentialsRequest
	GetProvider() *string
}

type ListCredentialsRequest struct {
	// credentialAuthType
	//
	// example:
	//
	// credentialAuthType
	CredentialAuthType *string `json:"credentialAuthType,omitempty" xml:"credentialAuthType,omitempty"`
	// credentialName
	//
	// example:
	//
	// credentialName
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// credentialSourceType
	//
	// example:
	//
	// credentialSourceType
	CredentialSourceType *string `json:"credentialSourceType,omitempty" xml:"credentialSourceType,omitempty"`
	// example:
	//
	// False
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Aliyun
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) GetCredentialAuthType() *string {
	return s.CredentialAuthType
}

func (s *ListCredentialsRequest) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ListCredentialsRequest) GetCredentialSourceType() *string {
	return s.CredentialSourceType
}

func (s *ListCredentialsRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListCredentialsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCredentialsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCredentialsRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListCredentialsRequest) SetCredentialAuthType(v string) *ListCredentialsRequest {
	s.CredentialAuthType = &v
	return s
}

func (s *ListCredentialsRequest) SetCredentialName(v string) *ListCredentialsRequest {
	s.CredentialName = &v
	return s
}

func (s *ListCredentialsRequest) SetCredentialSourceType(v string) *ListCredentialsRequest {
	s.CredentialSourceType = &v
	return s
}

func (s *ListCredentialsRequest) SetEnabled(v bool) *ListCredentialsRequest {
	s.Enabled = &v
	return s
}

func (s *ListCredentialsRequest) SetPageNumber(v int32) *ListCredentialsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCredentialsRequest) SetPageSize(v int32) *ListCredentialsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCredentialsRequest) SetProvider(v string) *ListCredentialsRequest {
	s.Provider = &v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
