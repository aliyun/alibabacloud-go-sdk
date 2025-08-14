// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *SearchEditingProjectResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *SearchEditingProjectResponseBody
	GetNextToken() *string
	SetProjectList(v []*SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody
	GetProjectList() []*SearchEditingProjectResponseBodyProjectList
	SetRequestId(v string) *SearchEditingProjectResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchEditingProjectResponseBody
	GetTotalCount() *int64
}

type SearchEditingProjectResponseBody struct {
	// The maximum number of entries returned on a single page. The value is set to the maximum number of entries returned on each page except for the last page.
	//
	// Examples:
	//
	// Valid example: 10,10,5. Invalid example: 10,5,10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// null
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried online editing projects.
	ProjectList []*SearchEditingProjectResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ****9262E3DA-07FA-4862-FCBB6BC61D08*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Optional. The total number of entries returned. By default, this parameter is not returned.
	//
	// example:
	//
	// 110
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *SearchEditingProjectResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchEditingProjectResponseBody) GetProjectList() []*SearchEditingProjectResponseBodyProjectList {
	return s.ProjectList
}

func (s *SearchEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchEditingProjectResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchEditingProjectResponseBody) SetMaxResults(v int64) *SearchEditingProjectResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetNextToken(v string) *SearchEditingProjectResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetProjectList(v []*SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody {
	s.ProjectList = v
	return s
}

func (s *SearchEditingProjectResponseBody) SetRequestId(v string) *SearchEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetTotalCount(v int64) *SearchEditingProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchEditingProjectResponseBodyProjectList struct {
	// The business configuration of the project. This parameter can be ignored for general editing projects.
	//
	// example:
	//
	// { "OutputMediaConfig" : { "StorageLocation": "test-bucket.oss-cn-shanghai.aliyuncs.com", "Path": "test-path" }, "OutputMediaTarget": "oss-object", "ReservationTime": "2021-06-21T08:05:00Z" }
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The business status of the project. This parameter can be ignored for general editing projects. Valid values:
	//
	// Valid values:
	//
	// 	- BroadCasting:
	//
	// 	- ReservationCanceled
	//
	// 	- LiveFinished
	//
	// 	- LoadingFailed
	//
	// 	- Reserving
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example-cover.jpg
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The method for editing the online editing project.
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
	// OpenAPI
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The time when the online editing project was created.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// sample description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total length of the online editing project. Unit: seconds.
	//
	// example:
	//
	// 30.100000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
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
	// "EventTime":"2021-08-12T10:04:15Z","ErrorCode":"InvalidParameter","ErrorMessage":"The specified parameter \\"LiveStreamConfig\\" is not valid. specified parameter example is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The method used when the online editing project was last modified.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
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
	// The type of the editing project.
	//
	// Valid values:
	//
	// 	- LiveEditingProject: a live stream editing project.
	//
	// 	- EditingProject: a regular editing project.
	//
	// example:
	//
	// EditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The status of the online editing project. Valid values:
	//
	// \\-Draft
	//
	// \\-Editing
	//
	// \\-Producing
	//
	// \\-Produced
	//
	// \\-ProduceFailed
	//
	// Valid values:
	//
	// 	- Draft
	//
	// 	- Produced
	//
	// 	- Editing
	//
	// 	- Producing
	//
	// 	- ProduceFailed
	//
	// example:
	//
	// PRODUCE_FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The timeline of the online editing project.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchEditingProjectResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectList) GetBusinessConfig() *string {
	return s.BusinessConfig
}

func (s *SearchEditingProjectResponseBodyProjectList) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *SearchEditingProjectResponseBodyProjectList) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchEditingProjectResponseBodyProjectList) GetCreateSource() *string {
	return s.CreateSource
}

func (s *SearchEditingProjectResponseBodyProjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchEditingProjectResponseBodyProjectList) GetDescription() *string {
	return s.Description
}

func (s *SearchEditingProjectResponseBodyProjectList) GetDuration() *int64 {
	return s.Duration
}

func (s *SearchEditingProjectResponseBodyProjectList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchEditingProjectResponseBodyProjectList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchEditingProjectResponseBodyProjectList) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *SearchEditingProjectResponseBodyProjectList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchEditingProjectResponseBodyProjectList) GetProjectId() *string {
	return s.ProjectId
}

func (s *SearchEditingProjectResponseBodyProjectList) GetProjectType() *string {
	return s.ProjectType
}

func (s *SearchEditingProjectResponseBodyProjectList) GetStatus() *string {
	return s.Status
}

func (s *SearchEditingProjectResponseBodyProjectList) GetTemplateType() *string {
	return s.TemplateType
}

func (s *SearchEditingProjectResponseBodyProjectList) GetTimeline() *string {
	return s.Timeline
}

func (s *SearchEditingProjectResponseBodyProjectList) GetTitle() *string {
	return s.Title
}

func (s *SearchEditingProjectResponseBodyProjectList) SetBusinessConfig(v string) *SearchEditingProjectResponseBodyProjectList {
	s.BusinessConfig = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetBusinessStatus(v string) *SearchEditingProjectResponseBodyProjectList {
	s.BusinessStatus = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCoverURL(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CoverURL = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCreateSource(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CreateSource = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCreateTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CreateTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDescription(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Description = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDuration(v int64) *SearchEditingProjectResponseBodyProjectList {
	s.Duration = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetErrorCode(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ErrorCode = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetErrorMessage(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ErrorMessage = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetModifiedSource(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ModifiedSource = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetModifiedTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ModifiedTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProjectId(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProjectType(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ProjectType = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetStatus(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTemplateType(v string) *SearchEditingProjectResponseBodyProjectList {
	s.TemplateType = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTimeline(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Timeline = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTitle(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Title = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) Validate() error {
	return dara.Validate(s)
}
