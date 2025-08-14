// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditingProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSource(v string) *ListEditingProjectsRequest
	GetCreateSource() *string
	SetEndTime(v string) *ListEditingProjectsRequest
	GetEndTime() *string
	SetKeyword(v string) *ListEditingProjectsRequest
	GetKeyword() *string
	SetMaxResults(v string) *ListEditingProjectsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListEditingProjectsRequest
	GetNextToken() *string
	SetProjectType(v string) *ListEditingProjectsRequest
	GetProjectType() *string
	SetSortBy(v string) *ListEditingProjectsRequest
	GetSortBy() *string
	SetStartTime(v string) *ListEditingProjectsRequest
	GetStartTime() *string
	SetStatus(v string) *ListEditingProjectsRequest
	GetStatus() *string
	SetTemplateType(v string) *ListEditingProjectsRequest
	GetTemplateType() *string
}

type ListEditingProjectsRequest struct {
	// The method for creating the online editing project. Valid values:
	//
	// \\- OpenAPI
	//
	// \\- AliyunConsole
	//
	// \\- WebSDK
	//
	// example:
	//
	// OpenAPI
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-02-02T23:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search keyword. You can search by job ID.
	//
	// example:
	//
	// ******6f36bc45d09a9d5cde49******
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per page. A maximum of 100 entries can be returned on each page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 8EqYpQbZ6Eh7+Zz8DxVYoQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the editing project. Valid values:
	//
	// 	- EditingProject: a regular editing project.
	//
	// 	- LiveEditingProject: a live stream editing project.
	//
	// example:
	//
	// EditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The order of sorting of the results. Valid values:
	//
	// 	- CreationTime:Desc (default): sorts the results in reverse chronological order.
	//
	// 	- CreationTime:Asc: sorts the results in chronological order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-12-21T08:00:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the online editing project. By default, online editing projects in all states are queried.
	//
	// example:
	//
	// Produced
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type. This parameter is required if you create a template-based online editing project. Default value: Timeline.
	//
	// *
	//
	// *
	//
	// Valid values:
	//
	// 	- Timeline: a regular template.
	//
	// 	- VETemplate: an advanced template.
	//
	// 	- None: general editing.
	//
	// example:
	//
	// None
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListEditingProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEditingProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListEditingProjectsRequest) GetCreateSource() *string {
	return s.CreateSource
}

func (s *ListEditingProjectsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEditingProjectsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEditingProjectsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListEditingProjectsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEditingProjectsRequest) GetProjectType() *string {
	return s.ProjectType
}

func (s *ListEditingProjectsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEditingProjectsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEditingProjectsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEditingProjectsRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListEditingProjectsRequest) SetCreateSource(v string) *ListEditingProjectsRequest {
	s.CreateSource = &v
	return s
}

func (s *ListEditingProjectsRequest) SetEndTime(v string) *ListEditingProjectsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEditingProjectsRequest) SetKeyword(v string) *ListEditingProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListEditingProjectsRequest) SetMaxResults(v string) *ListEditingProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEditingProjectsRequest) SetNextToken(v string) *ListEditingProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEditingProjectsRequest) SetProjectType(v string) *ListEditingProjectsRequest {
	s.ProjectType = &v
	return s
}

func (s *ListEditingProjectsRequest) SetSortBy(v string) *ListEditingProjectsRequest {
	s.SortBy = &v
	return s
}

func (s *ListEditingProjectsRequest) SetStartTime(v string) *ListEditingProjectsRequest {
	s.StartTime = &v
	return s
}

func (s *ListEditingProjectsRequest) SetStatus(v string) *ListEditingProjectsRequest {
	s.Status = &v
	return s
}

func (s *ListEditingProjectsRequest) SetTemplateType(v string) *ListEditingProjectsRequest {
	s.TemplateType = &v
	return s
}

func (s *ListEditingProjectsRequest) Validate() error {
	return dara.Validate(s)
}
