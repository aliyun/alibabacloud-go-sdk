// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryCommitDiffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryCommitDiffRequest
	GetAccessToken() *string
	SetContextLine(v int32) *ListRepositoryCommitDiffRequest
	GetContextLine() *int32
	SetOrganizationId(v string) *ListRepositoryCommitDiffRequest
	GetOrganizationId() *string
}

type ListRepositoryCommitDiffRequest struct {
	// accessToken
	//
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 3
	ContextLine *int32 `json:"contextLine,omitempty" xml:"contextLine,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s ListRepositoryCommitDiffRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryCommitDiffRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryCommitDiffRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryCommitDiffRequest) GetContextLine() *int32 {
	return s.ContextLine
}

func (s *ListRepositoryCommitDiffRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryCommitDiffRequest) SetAccessToken(v string) *ListRepositoryCommitDiffRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetContextLine(v int32) *ListRepositoryCommitDiffRequest {
	s.ContextLine = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) SetOrganizationId(v string) *ListRepositoryCommitDiffRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryCommitDiffRequest) Validate() error {
	return dara.Validate(s)
}
