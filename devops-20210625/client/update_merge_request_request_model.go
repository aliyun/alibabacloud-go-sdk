// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *UpdateMergeRequestRequest
	GetAccessToken() *string
	SetDescription(v string) *UpdateMergeRequestRequest
	GetDescription() *string
	SetTitle(v string) *UpdateMergeRequestRequest
	GetTitle() *string
	SetOrganizationId(v string) *UpdateMergeRequestRequest
	GetOrganizationId() *string
}

type UpdateMergeRequestRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s UpdateMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *UpdateMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *UpdateMergeRequestRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMergeRequestRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *UpdateMergeRequestRequest) SetAccessToken(v string) *UpdateMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *UpdateMergeRequestRequest) SetDescription(v string) *UpdateMergeRequestRequest {
	s.Description = &v
	return s
}

func (s *UpdateMergeRequestRequest) SetTitle(v string) *UpdateMergeRequestRequest {
	s.Title = &v
	return s
}

func (s *UpdateMergeRequestRequest) SetOrganizationId(v string) *UpdateMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *UpdateMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
