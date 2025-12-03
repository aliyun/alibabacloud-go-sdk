// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetGroupDetailRequest
	GetAccessToken() *string
	SetGroupId(v int64) *GetGroupDetailRequest
	GetGroupId() *int64
	SetOrganizationId(v string) *GetGroupDetailRequest
	GetOrganizationId() *string
}

type GetGroupDetailRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36612
	GroupId *int64 `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6218766746d4d2ca636d0497
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetGroupDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetGroupDetailRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetGroupDetailRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *GetGroupDetailRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetGroupDetailRequest) SetAccessToken(v string) *GetGroupDetailRequest {
	s.AccessToken = &v
	return s
}

func (s *GetGroupDetailRequest) SetGroupId(v int64) *GetGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupDetailRequest) SetOrganizationId(v string) *GetGroupDetailRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetGroupDetailRequest) Validate() error {
	return dara.Validate(s)
}
