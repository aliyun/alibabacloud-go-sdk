// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessLevel(v int32) *ListOrganizationsRequest
	GetAccessLevel() *int32
	SetAccessToken(v string) *ListOrganizationsRequest
	GetAccessToken() *string
	SetMinAccessLevel(v int32) *ListOrganizationsRequest
	GetMinAccessLevel() *int32
}

type ListOrganizationsRequest struct {
	// example:
	//
	// 5
	AccessLevel *int32  `json:"accessLevel,omitempty" xml:"accessLevel,omitempty"`
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// example:
	//
	// 60
	MinAccessLevel *int32 `json:"minAccessLevel,omitempty" xml:"minAccessLevel,omitempty"`
}

func (s ListOrganizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationsRequest) GetAccessLevel() *int32 {
	return s.AccessLevel
}

func (s *ListOrganizationsRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListOrganizationsRequest) GetMinAccessLevel() *int32 {
	return s.MinAccessLevel
}

func (s *ListOrganizationsRequest) SetAccessLevel(v int32) *ListOrganizationsRequest {
	s.AccessLevel = &v
	return s
}

func (s *ListOrganizationsRequest) SetAccessToken(v string) *ListOrganizationsRequest {
	s.AccessToken = &v
	return s
}

func (s *ListOrganizationsRequest) SetMinAccessLevel(v int32) *ListOrganizationsRequest {
	s.MinAccessLevel = &v
	return s
}

func (s *ListOrganizationsRequest) Validate() error {
	return dara.Validate(s)
}
