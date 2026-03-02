// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibrarySchemaUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *LibrarySchemaUpdateCmd
	GetArtifactId() *string
	SetDescription(v string) *LibrarySchemaUpdateCmd
	GetDescription() *string
	SetGitGroup(v string) *LibrarySchemaUpdateCmd
	GetGitGroup() *string
	SetGroupId(v string) *LibrarySchemaUpdateCmd
	GetGroupId() *string
	SetId(v int64) *LibrarySchemaUpdateCmd
	GetId() *int64
	SetLibraryId(v int64) *LibrarySchemaUpdateCmd
	GetLibraryId() *int64
}

type LibrarySchemaUpdateCmd struct {
	ArtifactId  *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GitGroup    *string `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LibraryId   *int64  `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s LibrarySchemaUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s LibrarySchemaUpdateCmd) GoString() string {
	return s.String()
}

func (s *LibrarySchemaUpdateCmd) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *LibrarySchemaUpdateCmd) GetDescription() *string {
	return s.Description
}

func (s *LibrarySchemaUpdateCmd) GetGitGroup() *string {
	return s.GitGroup
}

func (s *LibrarySchemaUpdateCmd) GetGroupId() *string {
	return s.GroupId
}

func (s *LibrarySchemaUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *LibrarySchemaUpdateCmd) GetLibraryId() *int64 {
	return s.LibraryId
}

func (s *LibrarySchemaUpdateCmd) SetArtifactId(v string) *LibrarySchemaUpdateCmd {
	s.ArtifactId = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) SetDescription(v string) *LibrarySchemaUpdateCmd {
	s.Description = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) SetGitGroup(v string) *LibrarySchemaUpdateCmd {
	s.GitGroup = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) SetGroupId(v string) *LibrarySchemaUpdateCmd {
	s.GroupId = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) SetId(v int64) *LibrarySchemaUpdateCmd {
	s.Id = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) SetLibraryId(v int64) *LibrarySchemaUpdateCmd {
	s.LibraryId = &v
	return s
}

func (s *LibrarySchemaUpdateCmd) Validate() error {
	return dara.Validate(s)
}
