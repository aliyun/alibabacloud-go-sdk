// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetMergeRequestRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *GetMergeRequestRequest
	GetOrganizationId() *string
}

type GetMergeRequestRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *GetMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetMergeRequestRequest) SetAccessToken(v string) *GetMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMergeRequestRequest) SetOrganizationId(v string) *GetMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
