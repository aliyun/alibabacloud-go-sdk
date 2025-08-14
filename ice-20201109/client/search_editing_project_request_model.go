// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSource(v string) *SearchEditingProjectRequest
	GetCreateSource() *string
	SetEndTime(v string) *SearchEditingProjectRequest
	GetEndTime() *string
	SetPageNo(v int64) *SearchEditingProjectRequest
	GetPageNo() *int64
	SetPageSize(v int64) *SearchEditingProjectRequest
	GetPageSize() *int64
	SetProjectType(v string) *SearchEditingProjectRequest
	GetProjectType() *string
	SetSortBy(v string) *SearchEditingProjectRequest
	GetSortBy() *string
	SetStartTime(v string) *SearchEditingProjectRequest
	GetStartTime() *string
	SetStatus(v string) *SearchEditingProjectRequest
	GetStatus() *string
	SetTemplateType(v string) *SearchEditingProjectRequest
	GetTemplateType() *string
}

type SearchEditingProjectRequest struct {
	// The source of the project.
	//
	// \\-OpenAPI
	//
	// \\-AliyunConsole
	//
	// \\-WebSDK
	//
	// Valid values:
	//
	// 	- AliyunConsole: The project is created in the Alibaba Cloud console.
	//
	// 	- WebSDK: The project is created by using the SDK for Web.
	//
	// 	- OpenAPI: The project is created by calling API operations.
	//
	// example:
	//
	// WebSDK
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The end of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the editing project. Default value: EditingProject. Valid values:
	//
	// 	- EditingProject: a regular editing project.
	//
	// 	- LiveEditingProject: a live stream editing project.
	//
	// example:
	//
	// EditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The sorting rule of results. Valid values:
	//
	// \\- CreationTime:Desc (default): The results are sorted in reverse chronological order based on the creation time.
	//
	// \\- CreationTime:Asc: The results are sorted in chronological order based on the creation time.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the online editing project. Separate multiple values with commas (,). By default, all online editing projects are queried.
	//
	// Valid values:
	//
	// \\-Draft
	//
	// \\-Producing
	//
	// \\-Produced
	//
	// \\-ProduceFailed
	//
	// example:
	//
	// Producing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type. Valid values:
	//
	// \\-Timeline
	//
	// \\-VETemplate
	//
	// Valid values:
	//
	// 	- Timeline: regular template.
	//
	// 	- VETemplate: advanced template.
	//
	// 	- None: No template is used.
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SearchEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectRequest) GetCreateSource() *string {
	return s.CreateSource
}

func (s *SearchEditingProjectRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *SearchEditingProjectRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *SearchEditingProjectRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SearchEditingProjectRequest) GetProjectType() *string {
	return s.ProjectType
}

func (s *SearchEditingProjectRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *SearchEditingProjectRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *SearchEditingProjectRequest) GetStatus() *string {
	return s.Status
}

func (s *SearchEditingProjectRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *SearchEditingProjectRequest) SetCreateSource(v string) *SearchEditingProjectRequest {
	s.CreateSource = &v
	return s
}

func (s *SearchEditingProjectRequest) SetEndTime(v string) *SearchEditingProjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageNo(v int64) *SearchEditingProjectRequest {
	s.PageNo = &v
	return s
}

func (s *SearchEditingProjectRequest) SetPageSize(v int64) *SearchEditingProjectRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEditingProjectRequest) SetProjectType(v string) *SearchEditingProjectRequest {
	s.ProjectType = &v
	return s
}

func (s *SearchEditingProjectRequest) SetSortBy(v string) *SearchEditingProjectRequest {
	s.SortBy = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStartTime(v string) *SearchEditingProjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStatus(v string) *SearchEditingProjectRequest {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectRequest) SetTemplateType(v string) *SearchEditingProjectRequest {
	s.TemplateType = &v
	return s
}

func (s *SearchEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
