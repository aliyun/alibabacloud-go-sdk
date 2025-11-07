// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListGitRepositoriesRequest
	GetClientToken() *string
	SetOrgId(v string) *ListGitRepositoriesRequest
	GetOrgId() *string
	SetOrgName(v string) *ListGitRepositoriesRequest
	GetOrgName() *string
	SetOwner(v string) *ListGitRepositoriesRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListGitRepositoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListGitRepositoriesRequest
	GetPageSize() *int32
	SetPlatform(v string) *ListGitRepositoriesRequest
	GetPlatform() *string
	SetRegionId(v string) *ListGitRepositoriesRequest
	GetRegionId() *string
}

type ListGitRepositoriesRequest struct {
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 6fekfmnfll135123kdf33
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// aliyun-computenest
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dhuai
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s ListGitRepositoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListGitRepositoriesRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *ListGitRepositoriesRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *ListGitRepositoriesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListGitRepositoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGitRepositoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGitRepositoriesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListGitRepositoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListGitRepositoriesRequest) SetClientToken(v string) *ListGitRepositoriesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetOrgId(v string) *ListGitRepositoriesRequest {
	s.OrgId = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetOrgName(v string) *ListGitRepositoriesRequest {
	s.OrgName = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetOwner(v string) *ListGitRepositoriesRequest {
	s.Owner = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPageNumber(v int32) *ListGitRepositoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPageSize(v int32) *ListGitRepositoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPlatform(v string) *ListGitRepositoriesRequest {
	s.Platform = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetRegionId(v string) *ListGitRepositoriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListGitRepositoriesRequest) Validate() error {
	return dara.Validate(s)
}
