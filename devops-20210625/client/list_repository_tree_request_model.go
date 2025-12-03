// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryTreeRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListRepositoryTreeRequest
	GetOrganizationId() *string
	SetPath(v string) *ListRepositoryTreeRequest
	GetPath() *string
	SetRefName(v string) *ListRepositoryTreeRequest
	GetRefName() *string
	SetType(v string) *ListRepositoryTreeRequest
	GetType() *string
}

type ListRepositoryTreeRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 624666bd54d036291ae13a36
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// module
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// master / tag1.0 / sjjfssa
	RefName *string `json:"refName,omitempty" xml:"refName,omitempty"`
	// example:
	//
	// RECURSIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRepositoryTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryTreeRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryTreeRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryTreeRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryTreeRequest) GetPath() *string {
	return s.Path
}

func (s *ListRepositoryTreeRequest) GetRefName() *string {
	return s.RefName
}

func (s *ListRepositoryTreeRequest) GetType() *string {
	return s.Type
}

func (s *ListRepositoryTreeRequest) SetAccessToken(v string) *ListRepositoryTreeRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetOrganizationId(v string) *ListRepositoryTreeRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetPath(v string) *ListRepositoryTreeRequest {
	s.Path = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetRefName(v string) *ListRepositoryTreeRequest {
	s.RefName = &v
	return s
}

func (s *ListRepositoryTreeRequest) SetType(v string) *ListRepositoryTreeRequest {
	s.Type = &v
	return s
}

func (s *ListRepositoryTreeRequest) Validate() error {
	return dara.Validate(s)
}
