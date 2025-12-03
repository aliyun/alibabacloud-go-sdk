// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMergeRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *CloseMergeRequestRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *CloseMergeRequestRequest
	GetOrganizationId() *string
}

type CloseMergeRequestRequest struct {
	// example:
	//
	// 0cf2c8458ac44d9481aab2dd6ec10596v3
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s CloseMergeRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseMergeRequestRequest) GoString() string {
	return s.String()
}

func (s *CloseMergeRequestRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *CloseMergeRequestRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *CloseMergeRequestRequest) SetAccessToken(v string) *CloseMergeRequestRequest {
	s.AccessToken = &v
	return s
}

func (s *CloseMergeRequestRequest) SetOrganizationId(v string) *CloseMergeRequestRequest {
	s.OrganizationId = &v
	return s
}

func (s *CloseMergeRequestRequest) Validate() error {
	return dara.Validate(s)
}
