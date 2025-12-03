// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListUserResourcesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListUserResourcesRequest
	GetOrganizationId() *string
	SetPage(v int64) *ListUserResourcesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListUserResourcesRequest
	GetPageSize() *int64
	SetUserIds(v string) *ListUserResourcesRequest
	GetUserIds() *string
}

type ListUserResourcesRequest struct {
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
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1,2,3
	UserIds *string `json:"userIds,omitempty" xml:"userIds,omitempty"`
}

func (s ListUserResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserResourcesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListUserResourcesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListUserResourcesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListUserResourcesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserResourcesRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *ListUserResourcesRequest) SetAccessToken(v string) *ListUserResourcesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListUserResourcesRequest) SetOrganizationId(v string) *ListUserResourcesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListUserResourcesRequest) SetPage(v int64) *ListUserResourcesRequest {
	s.Page = &v
	return s
}

func (s *ListUserResourcesRequest) SetPageSize(v int64) *ListUserResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserResourcesRequest) SetUserIds(v string) *ListUserResourcesRequest {
	s.UserIds = &v
	return s
}

func (s *ListUserResourcesRequest) Validate() error {
	return dara.Validate(s)
}
