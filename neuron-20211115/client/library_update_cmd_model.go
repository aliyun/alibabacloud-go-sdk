// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *LibraryUpdateCmd
	GetArtifactId() *string
	SetDescription(v string) *LibraryUpdateCmd
	GetDescription() *string
	SetGroupId(v string) *LibraryUpdateCmd
	GetGroupId() *string
	SetId(v int64) *LibraryUpdateCmd
	GetId() *int64
	SetName(v string) *LibraryUpdateCmd
	GetName() *string
	SetRepoUrl(v string) *LibraryUpdateCmd
	GetRepoUrl() *string
}

type LibraryUpdateCmd struct {
	ArtifactId  *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	RepoUrl     *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
}

func (s LibraryUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s LibraryUpdateCmd) GoString() string {
	return s.String()
}

func (s *LibraryUpdateCmd) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *LibraryUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *LibraryUpdateCmd) GetGroupId() *string {
	return s.GroupId
}

func (s *LibraryUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *LibraryUpdateCmd) GetName() *string {
	return s.Name
}

func (s *LibraryUpdateCmd) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *LibraryUpdateCmd) SetArtifactId(v string) *LibraryUpdateCmd {
	s.ArtifactId = &v
	return s
}

func (s *LibraryUpdateCmd) SetDescription(v string) *LibraryUpdateCmd {
	s.Description = &v
	return s
}

func (s *LibraryUpdateCmd) SetGroupId(v string) *LibraryUpdateCmd {
	s.GroupId = &v
	return s
}

func (s *LibraryUpdateCmd) SetId(v int64) *LibraryUpdateCmd {
	s.Id = &v
	return s
}

func (s *LibraryUpdateCmd) SetName(v string) *LibraryUpdateCmd {
	s.Name = &v
	return s
}

func (s *LibraryUpdateCmd) SetRepoUrl(v string) *LibraryUpdateCmd {
	s.RepoUrl = &v
	return s
}

func (s *LibraryUpdateCmd) Validate() error {
	return dara.Validate(s)
}
