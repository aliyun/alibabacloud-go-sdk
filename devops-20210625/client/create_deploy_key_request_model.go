// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CreateDeployKeyRequest
	GetAccessToken() *string
	SetKey(v string) *CreateDeployKeyRequest
	GetKey() *string
	SetTitle(v string) *CreateDeployKeyRequest
	GetTitle() *string
	SetOrganizationId(v string) *CreateDeployKeyRequest
	GetOrganizationId() *string
}

type CreateDeployKeyRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ""
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60de7a6852743a5162b5f957
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CreateDeployKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateDeployKeyRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CreateDeployKeyRequest) GetKey() *string {
	return s.Key
}

func (s *CreateDeployKeyRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateDeployKeyRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CreateDeployKeyRequest) SetAccessToken(v string) *CreateDeployKeyRequest {
	s.AccessToken = &v
	return s
}

func (s *CreateDeployKeyRequest) SetKey(v string) *CreateDeployKeyRequest {
	s.Key = &v
	return s
}

func (s *CreateDeployKeyRequest) SetTitle(v string) *CreateDeployKeyRequest {
	s.Title = &v
	return s
}

func (s *CreateDeployKeyRequest) SetOrganizationId(v string) *CreateDeployKeyRequest {
	s.OrganizationId = &v
	return s
}

func (s *CreateDeployKeyRequest) Validate() error {
	return dara.Validate(s)
}
