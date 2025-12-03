// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepositoryWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeleteRepositoryWebhookRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeleteRepositoryWebhookRequest
	GetOrganizationId() *string
}

type DeleteRepositoryWebhookRequest struct {
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
}

func (s DeleteRepositoryWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepositoryWebhookRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryWebhookRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeleteRepositoryWebhookRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeleteRepositoryWebhookRequest) SetAccessToken(v string) *DeleteRepositoryWebhookRequest {
	s.AccessToken = &v
	return s
}

func (s *DeleteRepositoryWebhookRequest) SetOrganizationId(v string) *DeleteRepositoryWebhookRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeleteRepositoryWebhookRequest) Validate() error {
	return dara.Validate(s)
}
