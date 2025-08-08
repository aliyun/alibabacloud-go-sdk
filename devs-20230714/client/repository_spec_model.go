// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepositorySpec interface {
	dara.Model
	String() string
	GoString() string
	SetCloneUrl(v string) *RepositorySpec
	GetCloneUrl() *string
	SetConnectionName(v string) *RepositorySpec
	GetConnectionName() *string
	SetDisplayName(v string) *RepositorySpec
	GetDisplayName() *string
	SetId(v int64) *RepositorySpec
	GetId() *int64
	SetOwner(v string) *RepositorySpec
	GetOwner() *string
	SetPlatform(v string) *RepositorySpec
	GetPlatform() *string
	SetWebUrl(v string) *RepositorySpec
	GetWebUrl() *string
}

type RepositorySpec struct {
	// This parameter is required.
	//
	// example:
	//
	// https://github.com/DDofDD/start-springboot-lfgy.git
	CloneUrl *string `json:"cloneUrl,omitempty" xml:"cloneUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// awesome-connection
	ConnectionName *string `json:"connectionName,omitempty" xml:"connectionName,omitempty"`
	// example:
	//
	// my-repo-name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 312649
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// my-org-name
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// github
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// https://github.com/my-org-name/my-repo-name
	WebUrl *string `json:"webUrl,omitempty" xml:"webUrl,omitempty"`
}

func (s RepositorySpec) String() string {
	return dara.Prettify(s)
}

func (s RepositorySpec) GoString() string {
	return s.String()
}

func (s *RepositorySpec) GetCloneUrl() *string {
	return s.CloneUrl
}

func (s *RepositorySpec) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *RepositorySpec) GetDisplayName() *string {
	return s.DisplayName
}

func (s *RepositorySpec) GetId() *int64 {
	return s.Id
}

func (s *RepositorySpec) GetOwner() *string {
	return s.Owner
}

func (s *RepositorySpec) GetPlatform() *string {
	return s.Platform
}

func (s *RepositorySpec) GetWebUrl() *string {
	return s.WebUrl
}

func (s *RepositorySpec) SetCloneUrl(v string) *RepositorySpec {
	s.CloneUrl = &v
	return s
}

func (s *RepositorySpec) SetConnectionName(v string) *RepositorySpec {
	s.ConnectionName = &v
	return s
}

func (s *RepositorySpec) SetDisplayName(v string) *RepositorySpec {
	s.DisplayName = &v
	return s
}

func (s *RepositorySpec) SetId(v int64) *RepositorySpec {
	s.Id = &v
	return s
}

func (s *RepositorySpec) SetOwner(v string) *RepositorySpec {
	s.Owner = &v
	return s
}

func (s *RepositorySpec) SetPlatform(v string) *RepositorySpec {
	s.Platform = &v
	return s
}

func (s *RepositorySpec) SetWebUrl(v string) *RepositorySpec {
	s.WebUrl = &v
	return s
}

func (s *RepositorySpec) Validate() error {
	return dara.Validate(s)
}
