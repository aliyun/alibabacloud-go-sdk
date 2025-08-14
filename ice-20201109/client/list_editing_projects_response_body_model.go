// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEditingProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEditingProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEditingProjectsResponseBody
	GetNextToken() *string
	SetProjectList(v []*ListEditingProjectsResponseBodyProjectList) *ListEditingProjectsResponseBody
	GetProjectList() []*ListEditingProjectsResponseBodyProjectList
	SetRequestId(v string) *ListEditingProjectsResponseBody
	GetRequestId() *string
}

type ListEditingProjectsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// This parameter is required.
	//
	// example:
	//
	// Nzv3rcKla9wHUGua9YXHNA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried online editing projects.
	ProjectList []*ListEditingProjectsResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEditingProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEditingProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEditingProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEditingProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEditingProjectsResponseBody) GetProjectList() []*ListEditingProjectsResponseBodyProjectList {
	return s.ProjectList
}

func (s *ListEditingProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEditingProjectsResponseBody) SetMaxResults(v int32) *ListEditingProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEditingProjectsResponseBody) SetNextToken(v string) *ListEditingProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEditingProjectsResponseBody) SetProjectList(v []*ListEditingProjectsResponseBodyProjectList) *ListEditingProjectsResponseBody {
	s.ProjectList = v
	return s
}

func (s *ListEditingProjectsResponseBody) SetRequestId(v string) *ListEditingProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEditingProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEditingProjectsResponseBodyProjectList struct {
	// The business configuration of the project. This parameter can be ignored for general editing projects.
	//
	// example:
	//
	// {}
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The business status of the project. This parameter can be ignored for general editing projects.
	//
	// example:
	//
	// {}
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://xxx.com/cover/xxx.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The method for editing the online editing project. Valid values:
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
	// The time when the online editing project was created.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the online editing project.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error code returned if the production of the online editing project failed.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the production of the online editing project failed.
	//
	// example:
	//
	// The specified parameter \\"LiveStreamConfig\\" is not valid. specified parameter example is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The method for modifying the online editing project last time.
	//
	// example:
	//
	// OpenAPI
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// The time when the online editing project was last modified.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ****fddd7748b58bf1d47e95****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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
	// The status of the online editing project. Valid values:
	//
	// \\- Draft
	//
	// \\- Editing
	//
	// \\- Producing
	//
	// \\- Produced
	//
	// \\- ProduceFailed
	//
	// example:
	//
	// Produced
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type. Valid values:
	//
	// 	- Timeline: a regular template.
	//
	// 	- VETemplate: an advanced template.
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The title of the online editing project.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListEditingProjectsResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s ListEditingProjectsResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *ListEditingProjectsResponseBodyProjectList) GetBusinessConfig() *string {
	return s.BusinessConfig
}

func (s *ListEditingProjectsResponseBodyProjectList) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *ListEditingProjectsResponseBodyProjectList) GetCoverURL() *string {
	return s.CoverURL
}

func (s *ListEditingProjectsResponseBodyProjectList) GetCreateSource() *string {
	return s.CreateSource
}

func (s *ListEditingProjectsResponseBodyProjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEditingProjectsResponseBodyProjectList) GetDescription() *string {
	return s.Description
}

func (s *ListEditingProjectsResponseBodyProjectList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEditingProjectsResponseBodyProjectList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListEditingProjectsResponseBodyProjectList) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *ListEditingProjectsResponseBodyProjectList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListEditingProjectsResponseBodyProjectList) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListEditingProjectsResponseBodyProjectList) GetProjectType() *string {
	return s.ProjectType
}

func (s *ListEditingProjectsResponseBodyProjectList) GetStatus() *string {
	return s.Status
}

func (s *ListEditingProjectsResponseBodyProjectList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListEditingProjectsResponseBodyProjectList) GetTitle() *string {
	return s.Title
}

func (s *ListEditingProjectsResponseBodyProjectList) SetBusinessConfig(v string) *ListEditingProjectsResponseBodyProjectList {
	s.BusinessConfig = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetBusinessStatus(v string) *ListEditingProjectsResponseBodyProjectList {
	s.BusinessStatus = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetCoverURL(v string) *ListEditingProjectsResponseBodyProjectList {
	s.CoverURL = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetCreateSource(v string) *ListEditingProjectsResponseBodyProjectList {
	s.CreateSource = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetCreateTime(v string) *ListEditingProjectsResponseBodyProjectList {
	s.CreateTime = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetDescription(v string) *ListEditingProjectsResponseBodyProjectList {
	s.Description = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetErrorCode(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ErrorCode = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetErrorMessage(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ErrorMessage = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetModifiedSource(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ModifiedSource = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetModifiedTime(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ModifiedTime = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetProjectId(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetProjectType(v string) *ListEditingProjectsResponseBodyProjectList {
	s.ProjectType = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetStatus(v string) *ListEditingProjectsResponseBodyProjectList {
	s.Status = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetTemplateType(v string) *ListEditingProjectsResponseBodyProjectList {
	s.TemplateType = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) SetTitle(v string) *ListEditingProjectsResponseBodyProjectList {
	s.Title = &v
	return s
}

func (s *ListEditingProjectsResponseBodyProjectList) Validate() error {
	return dara.Validate(s)
}
