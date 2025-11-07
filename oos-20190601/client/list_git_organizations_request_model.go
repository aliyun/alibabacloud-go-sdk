// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitOrganizationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListGitOrganizationsRequest
	GetClientToken() *string
	SetOwner(v string) *ListGitOrganizationsRequest
	GetOwner() *string
	SetPlatform(v string) *ListGitOrganizationsRequest
	GetPlatform() *string
	SetRegionId(v string) *ListGitOrganizationsRequest
	GetRegionId() *string
}

type ListGitOrganizationsRequest struct {
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// geegenw-j-imwinmtuej
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// github
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGitOrganizationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGitOrganizationsRequest) GoString() string {
	return s.String()
}

func (s *ListGitOrganizationsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListGitOrganizationsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListGitOrganizationsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListGitOrganizationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGitOrganizationsRequest) SetClientToken(v string) *ListGitOrganizationsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitOrganizationsRequest) SetOwner(v string) *ListGitOrganizationsRequest {
	s.Owner = &v
	return s
}

func (s *ListGitOrganizationsRequest) SetPlatform(v string) *ListGitOrganizationsRequest {
	s.Platform = &v
	return s
}

func (s *ListGitOrganizationsRequest) SetRegionId(v string) *ListGitOrganizationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListGitOrganizationsRequest) Validate() error {
	return dara.Validate(s)
}
