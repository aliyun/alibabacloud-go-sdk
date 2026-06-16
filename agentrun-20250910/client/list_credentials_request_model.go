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
	SetWorkspaceId(v string) *ListCredentialsRequest
	GetWorkspaceId() *string
	SetWorkspaceIds(v string) *ListCredentialsRequest
	GetWorkspaceIds() *string
}

type ListCredentialsRequest struct {
	// Filters the results by credential type.
	//
	// example:
	//
	// credentialAuthType
	CredentialAuthType *string `json:"credentialAuthType,omitempty" xml:"credentialAuthType,omitempty"`
	// Filters the results by credential name.
	//
	// example:
	//
	// credentialName
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// Filters the results by credential source type.
	//
	// example:
	//
	// credentialSourceType
	CredentialSourceType *string `json:"credentialSourceType,omitempty" xml:"credentialSourceType,omitempty"`
	// Filters the results based on the credential\\"s enabled status.
	//
	// example:
	//
	// False
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The page number of the results to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of credentials to return per page.
	//
	// example:
	//
	// 0
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filters the results by provider.
	//
	// example:
	//
	// Aliyun
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// Unique identifier of the workspace
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// Multiple workspace identifiers
	//
	// example:
	//
	// ws-1,ws-2
	WorkspaceIds *string `json:"workspaceIds,omitempty" xml:"workspaceIds,omitempty"`
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

func (s *ListCredentialsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCredentialsRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
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

func (s *ListCredentialsRequest) SetWorkspaceId(v string) *ListCredentialsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCredentialsRequest) SetWorkspaceIds(v string) *ListCredentialsRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
