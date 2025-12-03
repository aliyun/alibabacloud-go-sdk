// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListRepositoryWebhookRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListRepositoryWebhookRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListRepositoryWebhookRequest
	GetPage() *int64
	SetPageSize(v int64) *ListRepositoryWebhookRequest
	GetPageSize() *int64
}

type ListRepositoryWebhookRequest struct {
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRepositoryWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryWebhookRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryWebhookRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListRepositoryWebhookRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListRepositoryWebhookRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListRepositoryWebhookRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRepositoryWebhookRequest) SetAccessToken(v string) *ListRepositoryWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetOrganizationId(v string) *ListRepositoryWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPage(v int64) *ListRepositoryWebhookRequest {
	s.Page = &v
	return s
}

func (s *ListRepositoryWebhookRequest) SetPageSize(v int64) *ListRepositoryWebhookRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryWebhookRequest) Validate() error {
	return dara.Validate(s)
}
