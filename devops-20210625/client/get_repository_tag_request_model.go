// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRepositoryTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetRepositoryTagRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *GetRepositoryTagRequest
	GetOrganizationId() *string
	SetTagName(v string) *GetRepositoryTagRequest
	GetTagName() *string
}

type GetRepositoryTagRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag v1.0
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s GetRepositoryTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRepositoryTagRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryTagRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetRepositoryTagRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetRepositoryTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *GetRepositoryTagRequest) SetAccessToken(v string) *GetRepositoryTagRequest {
	s.AccessToken = &v
	return s
}

func (s *GetRepositoryTagRequest) SetOrganizationId(v string) *GetRepositoryTagRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetRepositoryTagRequest) SetTagName(v string) *GetRepositoryTagRequest {
	s.TagName = &v
	return s
}

func (s *GetRepositoryTagRequest) Validate() error {
	return dara.Validate(s)
}
