// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReopenMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ReopenMergeRequestRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ReopenMergeRequestRequest
	GetOrganizationId() *string
}

type ReopenMergeRequestRequest struct {
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

func (s ReopenMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s ReopenMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *ReopenMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ReopenMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ReopenMergeRequestRequest) SetAccessToken(v string) *ReopenMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *ReopenMergeRequestRequest) SetOrganizationId(v string) *ReopenMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *ReopenMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
