// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibrarySchema interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *LibrarySchema
	GetArtifactId() *string
	SetCompanyId(v int64) *LibrarySchema
	GetCompanyId() *int64
	SetDescription(v string) *LibrarySchema
	GetDescription() *string
	SetGitGroup(v string) *LibrarySchema
	GetGitGroup() *string
	SetGroupId(v string) *LibrarySchema
	GetGroupId() *string
	SetId(v int64) *LibrarySchema
	GetId() *int64
	SetIndustry(v string) *LibrarySchema
	GetIndustry() *string
	SetLibraryId(v int64) *LibrarySchema
	GetLibraryId() *int64
	SetMarketId(v int64) *LibrarySchema
	GetMarketId() *int64
	SetName(v string) *LibrarySchema
	GetName() *string
	SetProvider(v *Provider) *LibrarySchema
	GetProvider() *Provider
	SetRequestId(v string) *LibrarySchema
	GetRequestId() *string
}

type LibrarySchema struct {
	ArtifactId  *string   `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	CompanyId   *int64    `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	GitGroup    *string   `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	GroupId     *string   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id          *int64    `json:"id,omitempty" xml:"id,omitempty"`
	Industry    *string   `json:"industry,omitempty" xml:"industry,omitempty"`
	LibraryId   *int64    `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	MarketId    *int64    `json:"marketId,omitempty" xml:"marketId,omitempty"`
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	Provider    *Provider `json:"provider,omitempty" xml:"provider,omitempty"`
	RequestId   *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LibrarySchema) String() string {
	return dara.Prettify(s)
}

func (s LibrarySchema) GoString() string {
	return s.String()
}

func (s *LibrarySchema) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *LibrarySchema) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *LibrarySchema) GetDescription() *string {
	return s.Description
}

func (s *LibrarySchema) GetGitGroup() *string {
	return s.GitGroup
}

func (s *LibrarySchema) GetGroupId() *string {
	return s.GroupId
}

func (s *LibrarySchema) GetId() *int64 {
	return s.Id
}

func (s *LibrarySchema) GetIndustry() *string {
	return s.Industry
}

func (s *LibrarySchema) GetLibraryId() *int64 {
	return s.LibraryId
}

func (s *LibrarySchema) GetMarketId() *int64 {
	return s.MarketId
}

func (s *LibrarySchema) GetName() *string {
	return s.Name
}

func (s *LibrarySchema) GetProvider() *Provider {
	return s.Provider
}

func (s *LibrarySchema) GetRequestId() *string {
	return s.RequestId
}

func (s *LibrarySchema) SetArtifactId(v string) *LibrarySchema {
	s.ArtifactId = &v
	return s
}

func (s *LibrarySchema) SetCompanyId(v int64) *LibrarySchema {
	s.CompanyId = &v
	return s
}

func (s *LibrarySchema) SetDescription(v string) *LibrarySchema {
	s.Description = &v
	return s
}

func (s *LibrarySchema) SetGitGroup(v string) *LibrarySchema {
	s.GitGroup = &v
	return s
}

func (s *LibrarySchema) SetGroupId(v string) *LibrarySchema {
	s.GroupId = &v
	return s
}

func (s *LibrarySchema) SetId(v int64) *LibrarySchema {
	s.Id = &v
	return s
}

func (s *LibrarySchema) SetIndustry(v string) *LibrarySchema {
	s.Industry = &v
	return s
}

func (s *LibrarySchema) SetLibraryId(v int64) *LibrarySchema {
	s.LibraryId = &v
	return s
}

func (s *LibrarySchema) SetMarketId(v int64) *LibrarySchema {
	s.MarketId = &v
	return s
}

func (s *LibrarySchema) SetName(v string) *LibrarySchema {
	s.Name = &v
	return s
}

func (s *LibrarySchema) SetProvider(v *Provider) *LibrarySchema {
	s.Provider = v
	return s
}

func (s *LibrarySchema) SetRequestId(v string) *LibrarySchema {
	s.RequestId = &v
	return s
}

func (s *LibrarySchema) Validate() error {
	if s.Provider != nil {
		if err := s.Provider.Validate(); err != nil {
			return err
		}
	}
	return nil
}
