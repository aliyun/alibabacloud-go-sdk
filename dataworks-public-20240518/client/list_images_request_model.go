// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *ListImagesRequest
	GetAccessibility() *string
	SetName(v string) *ListImagesRequest
	GetName() *string
	SetOfficial(v bool) *ListImagesRequest
	GetOfficial() *bool
	SetPageNumber(v int32) *ListImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListImagesRequest
	GetPageSize() *int32
	SetProjectIds(v []*int64) *ListImagesRequest
	GetProjectIds() []*int64
	SetProviderTypes(v []*string) *ListImagesRequest
	GetProviderTypes() []*string
	SetSearchAll(v bool) *ListImagesRequest
	GetSearchAll() *bool
	SetSortBy(v string) *ListImagesRequest
	GetSortBy() *string
	SetStages(v []*string) *ListImagesRequest
	GetStages() []*string
	SetStatuses(v []*string) *ListImagesRequest
	GetStatuses() []*string
	SetSupportedModules(v []*string) *ListImagesRequest
	GetSupportedModules() []*string
	SetSupportedTaskTypes(v []*string) *ListImagesRequest
	GetSupportedTaskTypes() []*string
}

type ListImagesRequest struct {
	// example:
	//
	// Public
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// image
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Official *bool   `json:"Official,omitempty" xml:"Official,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectIds    []*int64  `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	ProviderTypes []*string `json:"ProviderTypes,omitempty" xml:"ProviderTypes,omitempty" type:"Repeated"`
	SearchAll     *bool     `json:"SearchAll,omitempty" xml:"SearchAll,omitempty"`
	// example:
	//
	// CreatedTime Desc
	SortBy             *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Stages             []*string `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Statuses           []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	SupportedModules   []*string `json:"SupportedModules,omitempty" xml:"SupportedModules,omitempty" type:"Repeated"`
	SupportedTaskTypes []*string `json:"SupportedTaskTypes,omitempty" xml:"SupportedTaskTypes,omitempty" type:"Repeated"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImagesRequest) GetName() *string {
	return s.Name
}

func (s *ListImagesRequest) GetOfficial() *bool {
	return s.Official
}

func (s *ListImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesRequest) GetProjectIds() []*int64 {
	return s.ProjectIds
}

func (s *ListImagesRequest) GetProviderTypes() []*string {
	return s.ProviderTypes
}

func (s *ListImagesRequest) GetSearchAll() *bool {
	return s.SearchAll
}

func (s *ListImagesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListImagesRequest) GetStages() []*string {
	return s.Stages
}

func (s *ListImagesRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListImagesRequest) GetSupportedModules() []*string {
	return s.SupportedModules
}

func (s *ListImagesRequest) GetSupportedTaskTypes() []*string {
	return s.SupportedTaskTypes
}

func (s *ListImagesRequest) SetAccessibility(v string) *ListImagesRequest {
	s.Accessibility = &v
	return s
}

func (s *ListImagesRequest) SetName(v string) *ListImagesRequest {
	s.Name = &v
	return s
}

func (s *ListImagesRequest) SetOfficial(v bool) *ListImagesRequest {
	s.Official = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int32) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int32) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) SetProjectIds(v []*int64) *ListImagesRequest {
	s.ProjectIds = v
	return s
}

func (s *ListImagesRequest) SetProviderTypes(v []*string) *ListImagesRequest {
	s.ProviderTypes = v
	return s
}

func (s *ListImagesRequest) SetSearchAll(v bool) *ListImagesRequest {
	s.SearchAll = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetStages(v []*string) *ListImagesRequest {
	s.Stages = v
	return s
}

func (s *ListImagesRequest) SetStatuses(v []*string) *ListImagesRequest {
	s.Statuses = v
	return s
}

func (s *ListImagesRequest) SetSupportedModules(v []*string) *ListImagesRequest {
	s.SupportedModules = v
	return s
}

func (s *ListImagesRequest) SetSupportedTaskTypes(v []*string) *ListImagesRequest {
	s.SupportedTaskTypes = v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
