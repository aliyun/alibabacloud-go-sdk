// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListGitAccountsRequest
	GetClientToken() *string
	SetPlatform(v string) *ListGitAccountsRequest
	GetPlatform() *string
	SetRegionId(v string) *ListGitAccountsRequest
	GetRegionId() *string
	SetRoleName(v string) *ListGitAccountsRequest
	GetRoleName() *string
}

type ListGitAccountsRequest struct {
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// example:
	//
	// roleName
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListGitAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGitAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListGitAccountsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListGitAccountsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListGitAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGitAccountsRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ListGitAccountsRequest) SetClientToken(v string) *ListGitAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitAccountsRequest) SetPlatform(v string) *ListGitAccountsRequest {
	s.Platform = &v
	return s
}

func (s *ListGitAccountsRequest) SetRegionId(v string) *ListGitAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListGitAccountsRequest) SetRoleName(v string) *ListGitAccountsRequest {
	s.RoleName = &v
	return s
}

func (s *ListGitAccountsRequest) Validate() error {
	return dara.Validate(s)
}
