// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferRepositoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *TransferRepositoryRequest
	GetAccessToken() *string
	SetGroupId(v string) *TransferRepositoryRequest
	GetGroupId() *string
	SetOrganizationId(v string) *TransferRepositoryRequest
	GetOrganizationId() *string
	SetRepositoryId(v string) *TransferRepositoryRequest
	GetRepositoryId() *string
}

type TransferRepositoryRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	RepositoryId *string `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
}

func (s TransferRepositoryRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferRepositoryRequest) GoString() string {
	return s.String()
}

func (s *TransferRepositoryRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *TransferRepositoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *TransferRepositoryRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *TransferRepositoryRequest) GetRepositoryId() *string {
	return s.RepositoryId
}

func (s *TransferRepositoryRequest) SetAccessToken(v string) *TransferRepositoryRequest {
	s.AccessToken = &v
	return s
}

func (s *TransferRepositoryRequest) SetGroupId(v string) *TransferRepositoryRequest {
	s.GroupId = &v
	return s
}

func (s *TransferRepositoryRequest) SetOrganizationId(v string) *TransferRepositoryRequest {
	s.OrganizationId = &v
	return s
}

func (s *TransferRepositoryRequest) SetRepositoryId(v string) *TransferRepositoryRequest {
	s.RepositoryId = &v
	return s
}

func (s *TransferRepositoryRequest) Validate() error {
	return dara.Validate(s)
}
