// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListImagesShrinkRequest
	GetAccessibility() *string
	SetName(v string) *ListImagesShrinkRequest
	GetName() *string
	SetOfficial(v bool) *ListImagesShrinkRequest
	GetOfficial() *bool
	SetPageNumber(v int32) *ListImagesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImagesShrinkRequest
	GetPageSize() *int32
	SetProjectIdsShrink(v string) *ListImagesShrinkRequest
	GetProjectIdsShrink() *string
	SetProviderTypesShrink(v string) *ListImagesShrinkRequest
	GetProviderTypesShrink() *string
	SetSearchAll(v bool) *ListImagesShrinkRequest
	GetSearchAll() *bool
	SetSortBy(v string) *ListImagesShrinkRequest
	GetSortBy() *string
	SetStagesShrink(v string) *ListImagesShrinkRequest
	GetStagesShrink() *string
	SetStatusesShrink(v string) *ListImagesShrinkRequest
	GetStatusesShrink() *string
	SetSupportedModulesShrink(v string) *ListImagesShrinkRequest
	GetSupportedModulesShrink() *string
	SetSupportedTaskTypesShrink(v string) *ListImagesShrinkRequest
	GetSupportedTaskTypesShrink() *string
}

type ListImagesShrinkRequest struct {
	// The accessibility of the image. Valid values:
	//
	// - Public: The image is visible to all users.
	//
	// - Private: The image is visible only to the creator.
	//
	// example:
	//
	// Public
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The image name used for fuzzy match.
	//
	// example:
	//
	// image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the image is an official image.
	Official *bool `json:"Official,omitempty" xml:"Official,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of workspace IDs.
	ProjectIdsShrink *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// The list of image reference data types.
	ProviderTypesShrink *string `json:"ProviderTypes,omitempty" xml:"ProviderTypes,omitempty"`
	// Specifies whether to search for all images.
	SearchAll *bool `json:"SearchAll,omitempty" xml:"SearchAll,omitempty"`
	// The list of sort fields. Fields such as scheduled time and start time are supported. The format is "sort field + sort order (Desc/Asc)". Asc can be omitted. Valid values for sort fields:
	//
	// - CreateTime (Desc/Asc): creation time.
	//
	// - Name (Desc/Asc): image name.
	//
	// Default value: CreateTime Asc.
	//
	// example:
	//
	// CreatedTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The list of image publishing stages to query.
	StagesShrink *string `json:"Stages,omitempty" xml:"Stages,omitempty"`
	// The list of image statuses to query.
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	// The list of supported modules.
	SupportedModulesShrink *string `json:"SupportedModules,omitempty" xml:"SupportedModules,omitempty"`
	// The list of supported task types.
	SupportedTaskTypesShrink *string `json:"SupportedTaskTypes,omitempty" xml:"SupportedTaskTypes,omitempty"`
}

func (s ListImagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImagesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListImagesShrinkRequest) GetOfficial() *bool {
	return s.Official
}

func (s *ListImagesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImagesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesShrinkRequest) GetProjectIdsShrink() *string {
	return s.ProjectIdsShrink
}

func (s *ListImagesShrinkRequest) GetProviderTypesShrink() *string {
	return s.ProviderTypesShrink
}

func (s *ListImagesShrinkRequest) GetSearchAll() *bool {
	return s.SearchAll
}

func (s *ListImagesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListImagesShrinkRequest) GetStagesShrink() *string {
	return s.StagesShrink
}

func (s *ListImagesShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListImagesShrinkRequest) GetSupportedModulesShrink() *string {
	return s.SupportedModulesShrink
}

func (s *ListImagesShrinkRequest) GetSupportedTaskTypesShrink() *string {
	return s.SupportedTaskTypesShrink
}

func (s *ListImagesShrinkRequest) SetAccessibility(v string) *ListImagesShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *ListImagesShrinkRequest) SetName(v string) *ListImagesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListImagesShrinkRequest) SetOfficial(v bool) *ListImagesShrinkRequest {
	s.Official = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageNumber(v int32) *ListImagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageSize(v int32) *ListImagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesShrinkRequest) SetProjectIdsShrink(v string) *ListImagesShrinkRequest {
	s.ProjectIdsShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetProviderTypesShrink(v string) *ListImagesShrinkRequest {
	s.ProviderTypesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetSearchAll(v bool) *ListImagesShrinkRequest {
	s.SearchAll = &v
	return s
}

func (s *ListImagesShrinkRequest) SetSortBy(v string) *ListImagesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesShrinkRequest) SetStagesShrink(v string) *ListImagesShrinkRequest {
	s.StagesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetStatusesShrink(v string) *ListImagesShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetSupportedModulesShrink(v string) *ListImagesShrinkRequest {
	s.SupportedModulesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetSupportedTaskTypesShrink(v string) *ListImagesShrinkRequest {
	s.SupportedTaskTypesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
