// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibrary interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *Library
	GetArtifactId() *string
	SetAssetType(v string) *Library
	GetAssetType() *string
	SetCompany(v string) *Library
	GetCompany() *string
	SetCompanyId(v int64) *Library
	GetCompanyId() *int64
	SetDependCount(v int32) *Library
	GetDependCount() *int32
	SetDescription(v string) *Library
	GetDescription() *string
	SetGroupId(v string) *Library
	GetGroupId() *string
	SetId(v int64) *Library
	GetId() *int64
	SetIndustry(v string) *Library
	GetIndustry() *string
	SetIsWatched(v bool) *Library
	GetIsWatched() *bool
	SetName(v string) *Library
	GetName() *string
	SetProvider(v string) *Library
	GetProvider() *string
	SetProviderName(v string) *Library
	GetProviderName() *string
	SetRepoUrl(v string) *Library
	GetRepoUrl() *string
	SetRequestId(v string) *Library
	GetRequestId() *string
	SetReviewId(v int64) *Library
	GetReviewId() *int64
	SetStatus(v string) *Library
	GetStatus() *string
	SetStepStatus(v string) *Library
	GetStepStatus() *string
	SetWatchCount(v int32) *Library
	GetWatchCount() *int32
}

type Library struct {
	ArtifactId   *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	AssetType    *string `json:"assetType,omitempty" xml:"assetType,omitempty"`
	Company      *string `json:"company,omitempty" xml:"company,omitempty"`
	CompanyId    *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	DependCount  *int32  `json:"dependCount,omitempty" xml:"dependCount,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupId      *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Industry     *string `json:"industry,omitempty" xml:"industry,omitempty"`
	IsWatched    *bool   `json:"isWatched,omitempty" xml:"isWatched,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Provider     *string `json:"provider,omitempty" xml:"provider,omitempty"`
	ProviderName *string `json:"providerName,omitempty" xml:"providerName,omitempty"`
	RepoUrl      *string `json:"repoUrl,omitempty" xml:"repoUrl,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ReviewId     *int64  `json:"reviewId,omitempty" xml:"reviewId,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	StepStatus   *string `json:"stepStatus,omitempty" xml:"stepStatus,omitempty"`
	WatchCount   *int32  `json:"watchCount,omitempty" xml:"watchCount,omitempty"`
}

func (s Library) String() string {
	return dara.Prettify(s)
}

func (s Library) GoString() string {
	return s.String()
}

func (s *Library) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *Library) GetAssetType() *string {
	return s.AssetType
}

func (s *Library) GetCompany() *string {
	return s.Company
}

func (s *Library) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *Library) GetDependCount() *int32 {
	return s.DependCount
}

func (s *Library) GetDescription() *string {
	return s.Description
}

func (s *Library) GetGroupId() *string {
	return s.GroupId
}

func (s *Library) GetId() *int64 {
	return s.Id
}

func (s *Library) GetIndustry() *string {
	return s.Industry
}

func (s *Library) GetIsWatched() *bool {
	return s.IsWatched
}

func (s *Library) GetName() *string {
	return s.Name
}

func (s *Library) GetProvider() *string {
	return s.Provider
}

func (s *Library) GetProviderName() *string {
	return s.ProviderName
}

func (s *Library) GetRepoUrl() *string {
	return s.RepoUrl
}

func (s *Library) GetRequestId() *string {
	return s.RequestId
}

func (s *Library) GetReviewId() *int64 {
	return s.ReviewId
}

func (s *Library) GetStatus() *string {
	return s.Status
}

func (s *Library) GetStepStatus() *string {
	return s.StepStatus
}

func (s *Library) GetWatchCount() *int32 {
	return s.WatchCount
}

func (s *Library) SetArtifactId(v string) *Library {
	s.ArtifactId = &v
	return s
}

func (s *Library) SetAssetType(v string) *Library {
	s.AssetType = &v
	return s
}

func (s *Library) SetCompany(v string) *Library {
	s.Company = &v
	return s
}

func (s *Library) SetCompanyId(v int64) *Library {
	s.CompanyId = &v
	return s
}

func (s *Library) SetDependCount(v int32) *Library {
	s.DependCount = &v
	return s
}

func (s *Library) SetDescription(v string) *Library {
	s.Description = &v
	return s
}

func (s *Library) SetGroupId(v string) *Library {
	s.GroupId = &v
	return s
}

func (s *Library) SetId(v int64) *Library {
	s.Id = &v
	return s
}

func (s *Library) SetIndustry(v string) *Library {
	s.Industry = &v
	return s
}

func (s *Library) SetIsWatched(v bool) *Library {
	s.IsWatched = &v
	return s
}

func (s *Library) SetName(v string) *Library {
	s.Name = &v
	return s
}

func (s *Library) SetProvider(v string) *Library {
	s.Provider = &v
	return s
}

func (s *Library) SetProviderName(v string) *Library {
	s.ProviderName = &v
	return s
}

func (s *Library) SetRepoUrl(v string) *Library {
	s.RepoUrl = &v
	return s
}

func (s *Library) SetRequestId(v string) *Library {
	s.RequestId = &v
	return s
}

func (s *Library) SetReviewId(v int64) *Library {
	s.ReviewId = &v
	return s
}

func (s *Library) SetStatus(v string) *Library {
	s.Status = &v
	return s
}

func (s *Library) SetStepStatus(v string) *Library {
	s.StepStatus = &v
	return s
}

func (s *Library) SetWatchCount(v int32) *Library {
	s.WatchCount = &v
	return s
}

func (s *Library) Validate() error {
	return dara.Validate(s)
}
