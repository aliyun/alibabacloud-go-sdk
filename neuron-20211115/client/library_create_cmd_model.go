// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *LibraryCreateCmd
	GetArtifactId() *string
	SetCompanyId(v int64) *LibraryCreateCmd
	GetCompanyId() *int64
	SetDescription(v string) *LibraryCreateCmd
	GetDescription() *string
	SetGroupId(v string) *LibraryCreateCmd
	GetGroupId() *string
	SetMarketId(v int64) *LibraryCreateCmd
	GetMarketId() *int64
	SetName(v string) *LibraryCreateCmd
	GetName() *string
	SetRepoUrl(v string) *LibraryCreateCmd
	GetRepoUrl() *string
}

type LibraryCreateCmd struct {
	ArtifactId  *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	MarketId    *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	RepoUrl     *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
}

func (s LibraryCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s LibraryCreateCmd) GoString() string {
	return s.String()
}

func (s *LibraryCreateCmd) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *LibraryCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *LibraryCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *LibraryCreateCmd) GetGroupId() *string {
	return s.GroupId
}

func (s *LibraryCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *LibraryCreateCmd) GetName() *string {
	return s.Name
}

func (s *LibraryCreateCmd) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *LibraryCreateCmd) SetArtifactId(v string) *LibraryCreateCmd {
	s.ArtifactId = &v
	return s
}

func (s *LibraryCreateCmd) SetCompanyId(v int64) *LibraryCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *LibraryCreateCmd) SetDescription(v string) *LibraryCreateCmd {
	s.Description = &v
	return s
}

func (s *LibraryCreateCmd) SetGroupId(v string) *LibraryCreateCmd {
	s.GroupId = &v
	return s
}

func (s *LibraryCreateCmd) SetMarketId(v int64) *LibraryCreateCmd {
	s.MarketId = &v
	return s
}

func (s *LibraryCreateCmd) SetName(v string) *LibraryCreateCmd {
	s.Name = &v
	return s
}

func (s *LibraryCreateCmd) SetRepoUrl(v string) *LibraryCreateCmd {
	s.RepoUrl = &v
	return s
}

func (s *LibraryCreateCmd) Validate() error {
	return dara.Validate(s)
}
