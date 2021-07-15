// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListSmartJobsRequest struct {
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	PageNo     *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	JobType    *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	JobState   *string `json:"JobState,omitempty" xml:"JobState,omitempty"`
}

func (s ListSmartJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSmartJobsRequest) SetStatus(v int64) *ListSmartJobsRequest {
	s.Status = &v
	return s
}

func (s *ListSmartJobsRequest) SetNextToken(v string) *ListSmartJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSmartJobsRequest) SetMaxResults(v int64) *ListSmartJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSmartJobsRequest) SetPageNo(v int64) *ListSmartJobsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSmartJobsRequest) SetPageSize(v int64) *ListSmartJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSmartJobsRequest) SetJobType(v string) *ListSmartJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListSmartJobsRequest) SetSortBy(v string) *ListSmartJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSmartJobsRequest) SetJobState(v string) *ListSmartJobsRequest {
	s.JobState = &v
	return s
}

type ListSmartJobsResponseBody struct {
	// Id of the request
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SmartJobList []*ListSmartJobsResponseBodySmartJobList `json:"SmartJobList,omitempty" xml:"SmartJobList,omitempty" type:"Repeated"`
	NextToken    *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults   *string                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount   *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSmartJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBody) SetRequestId(v string) *ListSmartJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetSmartJobList(v []*ListSmartJobsResponseBodySmartJobList) *ListSmartJobsResponseBody {
	s.SmartJobList = v
	return s
}

func (s *ListSmartJobsResponseBody) SetNextToken(v string) *ListSmartJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetMaxResults(v string) *ListSmartJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSmartJobsResponseBody) SetTotalCount(v string) *ListSmartJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSmartJobsResponseBodySmartJobList struct {
	JobId         *string                                            `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Title         *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
	Description   *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId        *int64                                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	JobType       *string                                            `json:"JobType,omitempty" xml:"JobType,omitempty"`
	EditingConfig *string                                            `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	UserData      *string                                            `json:"UserData,omitempty" xml:"UserData,omitempty"`
	JobState      *string                                            `json:"JobState,omitempty" xml:"JobState,omitempty"`
	CreateTime    *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime  *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	InputConfig   *ListSmartJobsResponseBodySmartJobListInputConfig  `json:"InputConfig,omitempty" xml:"InputConfig,omitempty" type:"Struct"`
	OutputConfig  *ListSmartJobsResponseBodySmartJobListOutputConfig `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty" type:"Struct"`
}

func (s ListSmartJobsResponseBodySmartJobList) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobList) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobId(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobId = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetTitle(v string) *ListSmartJobsResponseBodySmartJobList {
	s.Title = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetDescription(v string) *ListSmartJobsResponseBodySmartJobList {
	s.Description = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetUserId(v int64) *ListSmartJobsResponseBodySmartJobList {
	s.UserId = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobType(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobType = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetEditingConfig(v string) *ListSmartJobsResponseBodySmartJobList {
	s.EditingConfig = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetUserData(v string) *ListSmartJobsResponseBodySmartJobList {
	s.UserData = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetJobState(v string) *ListSmartJobsResponseBodySmartJobList {
	s.JobState = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetCreateTime(v string) *ListSmartJobsResponseBodySmartJobList {
	s.CreateTime = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetModifiedTime(v string) *ListSmartJobsResponseBodySmartJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetInputConfig(v *ListSmartJobsResponseBodySmartJobListInputConfig) *ListSmartJobsResponseBodySmartJobList {
	s.InputConfig = v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobList) SetOutputConfig(v *ListSmartJobsResponseBodySmartJobListOutputConfig) *ListSmartJobsResponseBodySmartJobList {
	s.OutputConfig = v
	return s
}

type ListSmartJobsResponseBodySmartJobListInputConfig struct {
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	Keyword   *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListSmartJobsResponseBodySmartJobListInputConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobListInputConfig) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) SetInputFile(v string) *ListSmartJobsResponseBodySmartJobListInputConfig {
	s.InputFile = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListInputConfig) SetKeyword(v string) *ListSmartJobsResponseBodySmartJobListInputConfig {
	s.Keyword = &v
	return s
}

type ListSmartJobsResponseBodySmartJobListOutputConfig struct {
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
}

func (s ListSmartJobsResponseBodySmartJobListOutputConfig) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsResponseBodySmartJobListOutputConfig) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) SetBucket(v string) *ListSmartJobsResponseBodySmartJobListOutputConfig {
	s.Bucket = &v
	return s
}

func (s *ListSmartJobsResponseBodySmartJobListOutputConfig) SetObject(v string) *ListSmartJobsResponseBodySmartJobListOutputConfig {
	s.Object = &v
	return s
}

type ListSmartJobsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSmartJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSmartJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSmartJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSmartJobsResponse) SetHeaders(v map[string]*string) *ListSmartJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSmartJobsResponse) SetBody(v *ListSmartJobsResponseBody) *ListSmartJobsResponse {
	s.Body = v
	return s
}

type DescribeRelatedAuthorizationStatusResponseBody struct {
	// Id of the request
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OSSAuthorized *bool   `json:"OSSAuthorized,omitempty" xml:"OSSAuthorized,omitempty"`
	MTSAuthorized *bool   `json:"MTSAuthorized,omitempty" xml:"MTSAuthorized,omitempty"`
	MNSAuthorized *bool   `json:"MNSAuthorized,omitempty" xml:"MNSAuthorized,omitempty"`
	Authorized    *bool   `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
}

func (s DescribeRelatedAuthorizationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelatedAuthorizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetRequestId(v string) *DescribeRelatedAuthorizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetOSSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.OSSAuthorized = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetMTSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.MTSAuthorized = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetMNSAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.MNSAuthorized = &v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponseBody) SetAuthorized(v bool) *DescribeRelatedAuthorizationStatusResponseBody {
	s.Authorized = &v
	return s
}

type DescribeRelatedAuthorizationStatusResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRelatedAuthorizationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRelatedAuthorizationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRelatedAuthorizationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRelatedAuthorizationStatusResponse) SetHeaders(v map[string]*string) *DescribeRelatedAuthorizationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRelatedAuthorizationStatusResponse) SetBody(v *DescribeRelatedAuthorizationStatusResponseBody) *DescribeRelatedAuthorizationStatusResponse {
	s.Body = v
	return s
}

type DeleteSmartJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteSmartJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmartJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobRequest) SetJobId(v string) *DeleteSmartJobRequest {
	s.JobId = &v
	return s
}

type DeleteSmartJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteSmartJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmartJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobResponseBody) SetRequestId(v string) *DeleteSmartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmartJobResponseBody) SetState(v string) *DeleteSmartJobResponseBody {
	s.State = &v
	return s
}

func (s *DeleteSmartJobResponseBody) SetJobId(v string) *DeleteSmartJobResponseBody {
	s.JobId = &v
	return s
}

type DeleteSmartJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSmartJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSmartJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSmartJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobResponse) SetHeaders(v map[string]*string) *DeleteSmartJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmartJobResponse) SetBody(v *DeleteSmartJobResponseBody) *DeleteSmartJobResponse {
	s.Body = v
	return s
}

type AddTemplateRequest struct {
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板类型，取值范围：Timeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 参见Timeline模板Config文档
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 模板封面
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 预览视频媒资id
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板创建来源，默认OpenAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 模板相关素材，模板编辑器使用
	RelatedMediaids *string `json:"RelatedMediaids,omitempty" xml:"RelatedMediaids,omitempty"`
}

func (s AddTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddTemplateRequest) SetName(v string) *AddTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddTemplateRequest) SetType(v string) *AddTemplateRequest {
	s.Type = &v
	return s
}

func (s *AddTemplateRequest) SetConfig(v string) *AddTemplateRequest {
	s.Config = &v
	return s
}

func (s *AddTemplateRequest) SetCoverUrl(v string) *AddTemplateRequest {
	s.CoverUrl = &v
	return s
}

func (s *AddTemplateRequest) SetPreviewMedia(v string) *AddTemplateRequest {
	s.PreviewMedia = &v
	return s
}

func (s *AddTemplateRequest) SetStatus(v string) *AddTemplateRequest {
	s.Status = &v
	return s
}

func (s *AddTemplateRequest) SetSource(v string) *AddTemplateRequest {
	s.Source = &v
	return s
}

func (s *AddTemplateRequest) SetRelatedMediaids(v string) *AddTemplateRequest {
	s.RelatedMediaids = &v
	return s
}

type AddTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 模板信息
	Template *AddTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s AddTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBody) SetRequestId(v string) *AddTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTemplateResponseBody) SetTemplate(v *AddTemplateResponseBodyTemplate) *AddTemplateResponseBody {
	s.Template = v
	return s
}

type AddTemplateResponseBodyTemplate struct {
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 参见Timeline模板Config文档
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 模板封面
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 预览视频媒资id
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 模板创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 模板修改来源
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
}

func (s AddTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s AddTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *AddTemplateResponseBodyTemplate) SetTemplateId(v string) *AddTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetName(v string) *AddTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetType(v string) *AddTemplateResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetConfig(v string) *AddTemplateResponseBodyTemplate {
	s.Config = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetCoverUrl(v string) *AddTemplateResponseBodyTemplate {
	s.CoverUrl = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetPreviewMedia(v string) *AddTemplateResponseBodyTemplate {
	s.PreviewMedia = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetStatus(v string) *AddTemplateResponseBodyTemplate {
	s.Status = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetCreateSource(v string) *AddTemplateResponseBodyTemplate {
	s.CreateSource = &v
	return s
}

func (s *AddTemplateResponseBodyTemplate) SetModifiedSource(v string) *AddTemplateResponseBodyTemplate {
	s.ModifiedSource = &v
	return s
}

type AddTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddTemplateResponse) SetHeaders(v map[string]*string) *AddTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddTemplateResponse) SetBody(v *AddTemplateResponseBody) *AddTemplateResponse {
	s.Body = v
	return s
}

type UpdateEditingProjectRequest struct {
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程时间线，Json格式
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectRequest) SetTitle(v string) *UpdateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetDescription(v string) *UpdateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetTimeline(v string) *UpdateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetCoverURL(v string) *UpdateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateEditingProjectRequest) SetProjectId(v string) *UpdateEditingProjectRequest {
	s.ProjectId = &v
	return s
}

type UpdateEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponseBody) SetRequestId(v string) *UpdateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateEditingProjectResponse) SetHeaders(v map[string]*string) *UpdateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateEditingProjectResponse) SetBody(v *UpdateEditingProjectResponseBody) *UpdateEditingProjectResponse {
	s.Body = v
	return s
}

type ListMediaProducingJobsRequest struct {
	// 查询以下状态的合成任务，支持多值，以英文逗号分隔
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMediaProducingJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsRequest) SetStatus(v string) *ListMediaProducingJobsRequest {
	s.Status = &v
	return s
}

type ListMediaProducingJobsResponseBody struct {
	// Id of the request
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaProducingJobList []*ListMediaProducingJobsResponseBodyMediaProducingJobList `json:"MediaProducingJobList,omitempty" xml:"MediaProducingJobList,omitempty" type:"Repeated"`
}

func (s ListMediaProducingJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBody) SetRequestId(v string) *ListMediaProducingJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBody) SetMediaProducingJobList(v []*ListMediaProducingJobsResponseBodyMediaProducingJobList) *ListMediaProducingJobsResponseBody {
	s.MediaProducingJobList = v
	return s
}

type ListMediaProducingJobsResponseBodyMediaProducingJobList struct {
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProjectId    *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MediaId      *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaURL     *string  `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	TemplateId   *string  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam   *string  `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	Duration     *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CompleteTime *string  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	ModifiedTime *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Code         *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponseBodyMediaProducingJobList) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetJobId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.JobId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetProjectId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ProjectId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMediaURL(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.MediaURL = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetTemplateId(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.TemplateId = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetClipsParam(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ClipsParam = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetDuration(v float32) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Duration = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCreateTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CreateTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCompleteTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.CompleteTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetModifiedTime(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetStatus(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Status = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetCode(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Code = &v
	return s
}

func (s *ListMediaProducingJobsResponseBodyMediaProducingJobList) SetMessage(v string) *ListMediaProducingJobsResponseBodyMediaProducingJobList {
	s.Message = &v
	return s
}

type ListMediaProducingJobsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMediaProducingJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMediaProducingJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMediaProducingJobsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaProducingJobsResponse) SetHeaders(v map[string]*string) *ListMediaProducingJobsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaProducingJobsResponse) SetBody(v *ListMediaProducingJobsResponseBody) *ListMediaProducingJobsResponse {
	s.Body = v
	return s
}

type GetEditingProjectMaterialsRequest struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetEditingProjectMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsRequest) SetProjectId(v string) *GetEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

type GetEditingProjectMaterialsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*GetEditingProjectMaterialsResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBody) SetRequestId(v string) *GetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetProjectId(v string) *GetEditingProjectMaterialsResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetMediaInfos(v []*GetEditingProjectMaterialsResponseBodyMediaInfos) *GetEditingProjectMaterialsResponseBody {
	s.MediaInfos = v
	return s
}

type GetEditingProjectMaterialsResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetMediaId(v string) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetMediaBasicInfo(v *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfos) SetFileInfoList(v []*GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) *GetEditingProjectMaterialsResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// 雪碧图
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

type GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *GetEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type GetEditingProjectMaterialsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEditingProjectMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *GetEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectMaterialsResponse) SetBody(v *GetEditingProjectMaterialsResponseBody) *GetEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

type GetDefaultStorageLocationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 存储类型
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// oss bucket 名称
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// 路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDefaultStorageLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultStorageLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultStorageLocationResponseBody) SetRequestId(v string) *GetDefaultStorageLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetStorageType(v string) *GetDefaultStorageLocationResponseBody {
	s.StorageType = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetBucket(v string) *GetDefaultStorageLocationResponseBody {
	s.Bucket = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetPath(v string) *GetDefaultStorageLocationResponseBody {
	s.Path = &v
	return s
}

func (s *GetDefaultStorageLocationResponseBody) SetStatus(v string) *GetDefaultStorageLocationResponseBody {
	s.Status = &v
	return s
}

type GetDefaultStorageLocationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDefaultStorageLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDefaultStorageLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultStorageLocationResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultStorageLocationResponse) SetHeaders(v map[string]*string) *GetDefaultStorageLocationResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultStorageLocationResponse) SetBody(v *GetDefaultStorageLocationResponseBody) *GetDefaultStorageLocationResponse {
	s.Body = v
	return s
}

type DeleteMediaInfosRequest struct {
	// ICE 媒资ID
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURLs *string `json:"InputURLs,omitempty" xml:"InputURLs,omitempty"`
}

func (s DeleteMediaInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosRequest) SetMediaIds(v string) *DeleteMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *DeleteMediaInfosRequest) SetInputURLs(v string) *DeleteMediaInfosRequest {
	s.InputURLs = &v
	return s
}

type DeleteMediaInfosResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 出现获取错误的ID或inputUr
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
}

func (s DeleteMediaInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponseBody) SetRequestId(v string) *DeleteMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaInfosResponseBody) SetIgnoredList(v []*string) *DeleteMediaInfosResponseBody {
	s.IgnoredList = v
	return s
}

type DeleteMediaInfosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMediaInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponse) SetHeaders(v map[string]*string) *DeleteMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaInfosResponse) SetBody(v *DeleteMediaInfosResponseBody) *DeleteMediaInfosResponse {
	s.Body = v
	return s
}

type SetEventCallbackRequest struct {
	CallbackQueueName *string `json:"CallbackQueueName,omitempty" xml:"CallbackQueueName,omitempty"`
}

func (s SetEventCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SetEventCallbackRequest) GoString() string {
	return s.String()
}

func (s *SetEventCallbackRequest) SetCallbackQueueName(v string) *SetEventCallbackRequest {
	s.CallbackQueueName = &v
	return s
}

type SetEventCallbackResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否设置成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetEventCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetEventCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SetEventCallbackResponseBody) SetRequestId(v string) *SetEventCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetEventCallbackResponseBody) SetSuccess(v bool) *SetEventCallbackResponseBody {
	s.Success = &v
	return s
}

type SetEventCallbackResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetEventCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetEventCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SetEventCallbackResponse) GoString() string {
	return s.String()
}

func (s *SetEventCallbackResponse) SetHeaders(v map[string]*string) *SetEventCallbackResponse {
	s.Headers = v
	return s
}

func (s *SetEventCallbackResponse) SetBody(v *SetEventCallbackResponseBody) *SetEventCallbackResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// 模板Id
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetTemplateResponseBody struct {
	// Id of the request
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Template  *GetTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody {
	s.Template = v
	return s
}

type GetTemplateResponseBodyTemplate struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 模板配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 预览素材
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 修改来源
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// 预览素材状态
	PreviewMediaStatus *string `json:"PreviewMediaStatus,omitempty" xml:"PreviewMediaStatus,omitempty"`
	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 封面URL
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
}

func (s GetTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateId(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetName(v string) *GetTemplateResponseBodyTemplate {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetType(v string) *GetTemplateResponseBodyTemplate {
	s.Type = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetConfig(v string) *GetTemplateResponseBodyTemplate {
	s.Config = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetPreviewMedia(v string) *GetTemplateResponseBodyTemplate {
	s.PreviewMedia = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetStatus(v string) *GetTemplateResponseBodyTemplate {
	s.Status = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreateSource(v string) *GetTemplateResponseBodyTemplate {
	s.CreateSource = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetModifiedSource(v string) *GetTemplateResponseBodyTemplate {
	s.ModifiedSource = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetPreviewMediaStatus(v string) *GetTemplateResponseBodyTemplate {
	s.PreviewMediaStatus = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreationTime(v string) *GetTemplateResponseBodyTemplate {
	s.CreationTime = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetModifiedTime(v string) *GetTemplateResponseBodyTemplate {
	s.ModifiedTime = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCoverURL(v string) *GetTemplateResponseBodyTemplate {
	s.CoverURL = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type RegisterMediaInfoRequest struct {
	// 媒资媒体url
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签,如果有多个标签用逗号隔开
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面图，仅视频媒资有效
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户自定义元数据
	DynamicMetaDataList *string `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty"`
	// 用户数据，最大1024字节
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 是否覆盖已有媒资
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// 客户端token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 注册媒资的配置
	RegisterConfig *string `json:"RegisterConfig,omitempty" xml:"RegisterConfig,omitempty"`
}

func (s RegisterMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoRequest) SetInputURL(v string) *RegisterMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaType(v string) *RegisterMediaInfoRequest {
	s.MediaType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetBusinessType(v string) *RegisterMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetTitle(v string) *RegisterMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetDescription(v string) *RegisterMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCategory(v string) *RegisterMediaInfoRequest {
	s.Category = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetMediaTags(v string) *RegisterMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetCoverURL(v string) *RegisterMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetDynamicMetaDataList(v string) *RegisterMediaInfoRequest {
	s.DynamicMetaDataList = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetUserData(v string) *RegisterMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetOverwrite(v bool) *RegisterMediaInfoRequest {
	s.Overwrite = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetClientToken(v string) *RegisterMediaInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterMediaInfoRequest) SetRegisterConfig(v string) *RegisterMediaInfoRequest {
	s.RegisterConfig = &v
	return s
}

type RegisterMediaInfoResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ICE媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s RegisterMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponseBody) SetRequestId(v string) *RegisterMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterMediaInfoResponseBody) SetMediaId(v string) *RegisterMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

type RegisterMediaInfoResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *RegisterMediaInfoResponse) SetHeaders(v map[string]*string) *RegisterMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *RegisterMediaInfoResponse) SetBody(v *RegisterMediaInfoResponseBody) *RegisterMediaInfoResponse {
	s.Body = v
	return s
}

type CreateEditingProjectRequest struct {
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程时间线，Json格式
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// FEExtend
	FEExtend *string `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
}

func (s CreateEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectRequest) SetTitle(v string) *CreateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectRequest) SetDescription(v string) *CreateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTimeline(v string) *CreateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectRequest) SetCoverURL(v string) *CreateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectRequest) SetFEExtend(v string) *CreateEditingProjectRequest {
	s.FEExtend = &v
	return s
}

type CreateEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Project   *CreateEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
}

func (s CreateEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBody) SetRequestId(v string) *CreateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEditingProjectResponseBody) SetProject(v *CreateEditingProjectResponseBodyProject) *CreateEditingProjectResponseBody {
	s.Project = v
	return s
}

type CreateEditingProjectResponseBodyProject struct {
	ProjectId      *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Title          *string  `json:"Title,omitempty" xml:"Title,omitempty"`
	Description    *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Timeline       *string  `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	CoverURL       *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	Status         *int64   `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName     *string  `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	CreateTime     *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime   *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Duration       *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CreateSource   *string  `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	ModifiedSource *string  `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	TemplateType   *string  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBodyProject) SetProjectId(v string) *CreateEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTitle(v string) *CreateEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDescription(v string) *CreateEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTimeline(v string) *CreateEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCoverURL(v string) *CreateEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetStatus(v int64) *CreateEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetStatusName(v string) *CreateEditingProjectResponseBodyProject {
	s.StatusName = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCreateTime(v string) *CreateEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetModifiedTime(v string) *CreateEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDuration(v float32) *CreateEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCreateSource(v string) *CreateEditingProjectResponseBodyProject {
	s.CreateSource = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetModifiedSource(v string) *CreateEditingProjectResponseBodyProject {
	s.ModifiedSource = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTemplateType(v string) *CreateEditingProjectResponseBodyProject {
	s.TemplateType = &v
	return s
}

type CreateEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponse) SetHeaders(v map[string]*string) *CreateEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateEditingProjectResponse) SetBody(v *CreateEditingProjectResponseBody) *CreateEditingProjectResponse {
	s.Body = v
	return s
}

type BatchGetMediaInfosRequest struct {
	MediaIds     *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
}

func (s BatchGetMediaInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosRequest) SetMediaIds(v string) *BatchGetMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetMediaInfosRequest) SetAdditionType(v string) *BatchGetMediaInfosRequest {
	s.AdditionType = &v
	return s
}

type BatchGetMediaInfosResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*BatchGetMediaInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
}

func (s BatchGetMediaInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBody) SetRequestId(v string) *BatchGetMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBody) SetMediaInfos(v []*BatchGetMediaInfosResponseBodyMediaInfos) *BatchGetMediaInfosResponseBody {
	s.MediaInfos = v
	return s
}

type BatchGetMediaInfosResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfos) SetFileInfoList(v []*BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) *BatchGetMediaInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// 雪碧图
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *BatchGetMediaInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *BatchGetMediaInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type BatchGetMediaInfosResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetMediaInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetMediaInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetMediaInfosResponse) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosResponse) SetHeaders(v map[string]*string) *BatchGetMediaInfosResponse {
	s.Headers = v
	return s
}

func (s *BatchGetMediaInfosResponse) SetBody(v *BatchGetMediaInfosResponseBody) *BatchGetMediaInfosResponse {
	s.Body = v
	return s
}

type SetDefaultStorageLocationRequest struct {
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Bucket      *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Path        *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s SetDefaultStorageLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultStorageLocationRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationRequest) SetStorageType(v string) *SetDefaultStorageLocationRequest {
	s.StorageType = &v
	return s
}

func (s *SetDefaultStorageLocationRequest) SetBucket(v string) *SetDefaultStorageLocationRequest {
	s.Bucket = &v
	return s
}

func (s *SetDefaultStorageLocationRequest) SetPath(v string) *SetDefaultStorageLocationRequest {
	s.Path = &v
	return s
}

type SetDefaultStorageLocationResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetDefaultStorageLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultStorageLocationResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationResponseBody) SetRequestId(v string) *SetDefaultStorageLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultStorageLocationResponseBody) SetSuccess(v bool) *SetDefaultStorageLocationResponseBody {
	s.Success = &v
	return s
}

type SetDefaultStorageLocationResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDefaultStorageLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDefaultStorageLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultStorageLocationResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultStorageLocationResponse) SetHeaders(v map[string]*string) *SetDefaultStorageLocationResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultStorageLocationResponse) SetBody(v *SetDefaultStorageLocationResponseBody) *SetDefaultStorageLocationResponse {
	s.Body = v
	return s
}

type UpdateMediaInfoRequest struct {
	// 媒资Id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 媒资媒体类型
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签,如果有多个标签用逗号隔开
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面图，仅视频媒资有效
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户自定义元数据
	DynamicMetaDataList *string `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty"`
	// 用户数据，最大1024字节
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 是否以append的形式更新Tags字段
	AppendTags *bool `json:"AppendTags,omitempty" xml:"AppendTags,omitempty"`
	// 是否以append的形式更新DynamicMetaDataList字段
	AppendDynamicMeta *bool `json:"AppendDynamicMeta,omitempty" xml:"AppendDynamicMeta,omitempty"`
}

func (s UpdateMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoRequest) SetMediaId(v string) *UpdateMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetInputURL(v string) *UpdateMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetBusinessType(v string) *UpdateMediaInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetTitle(v string) *UpdateMediaInfoRequest {
	s.Title = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetDescription(v string) *UpdateMediaInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCategory(v string) *UpdateMediaInfoRequest {
	s.Category = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetMediaTags(v string) *UpdateMediaInfoRequest {
	s.MediaTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetCoverURL(v string) *UpdateMediaInfoRequest {
	s.CoverURL = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetDynamicMetaDataList(v string) *UpdateMediaInfoRequest {
	s.DynamicMetaDataList = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetUserData(v string) *UpdateMediaInfoRequest {
	s.UserData = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetAppendTags(v bool) *UpdateMediaInfoRequest {
	s.AppendTags = &v
	return s
}

func (s *UpdateMediaInfoRequest) SetAppendDynamicMeta(v bool) *UpdateMediaInfoRequest {
	s.AppendDynamicMeta = &v
	return s
}

type UpdateMediaInfoResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ICE媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s UpdateMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponseBody) SetRequestId(v string) *UpdateMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaInfoResponseBody) SetMediaId(v string) *UpdateMediaInfoResponseBody {
	s.MediaId = &v
	return s
}

type UpdateMediaInfoResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaInfoResponse) SetHeaders(v map[string]*string) *UpdateMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaInfoResponse) SetBody(v *UpdateMediaInfoResponseBody) *UpdateMediaInfoResponse {
	s.Body = v
	return s
}

type GetMediaProducingJobRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaProducingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobRequest) SetJobId(v string) *GetMediaProducingJobRequest {
	s.JobId = &v
	return s
}

type GetMediaProducingJobResponseBody struct {
	// Id of the request
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaProducingJob *GetMediaProducingJobResponseBodyMediaProducingJob `json:"MediaProducingJob,omitempty" xml:"MediaProducingJob,omitempty" type:"Struct"`
}

func (s GetMediaProducingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBody) SetRequestId(v string) *GetMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaProducingJobResponseBody) SetMediaProducingJob(v *GetMediaProducingJobResponseBodyMediaProducingJob) *GetMediaProducingJobResponseBody {
	s.MediaProducingJob = v
	return s
}

type GetMediaProducingJobResponseBodyMediaProducingJob struct {
	JobId        *string  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProjectId    *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	MediaId      *string  `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	MediaURL     *string  `json:"MediaURL,omitempty" xml:"MediaURL,omitempty"`
	Timeline     *string  `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	TemplateId   *string  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam   *string  `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	Duration     *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CompleteTime *string  `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	ModifiedTime *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	Code         *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string  `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponseBodyMediaProducingJob) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetJobId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.JobId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetProjectId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ProjectId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMediaURL(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.MediaURL = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTimeline(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Timeline = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetTemplateId(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.TemplateId = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetClipsParam(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ClipsParam = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetDuration(v float32) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Duration = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCreateTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CreateTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCompleteTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.CompleteTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetModifiedTime(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetStatus(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Status = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetCode(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Code = &v
	return s
}

func (s *GetMediaProducingJobResponseBodyMediaProducingJob) SetMessage(v string) *GetMediaProducingJobResponseBodyMediaProducingJob {
	s.Message = &v
	return s
}

type GetMediaProducingJobResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaProducingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaProducingJobResponse) SetHeaders(v map[string]*string) *GetMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaProducingJobResponse) SetBody(v *GetMediaProducingJobResponseBody) *GetMediaProducingJobResponse {
	s.Body = v
	return s
}

type DescribeIceProductStatusResponseBody struct {
	// Id of the request
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ICEServiceAvaliable *bool   `json:"ICEServiceAvaliable,omitempty" xml:"ICEServiceAvaliable,omitempty"`
}

func (s DescribeIceProductStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIceProductStatusResponseBody) SetRequestId(v string) *DescribeIceProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIceProductStatusResponseBody) SetICEServiceAvaliable(v bool) *DescribeIceProductStatusResponseBody {
	s.ICEServiceAvaliable = &v
	return s
}

type DescribeIceProductStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIceProductStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIceProductStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIceProductStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIceProductStatusResponse) SetHeaders(v map[string]*string) *DescribeIceProductStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIceProductStatusResponse) SetBody(v *DescribeIceProductStatusResponseBody) *DescribeIceProductStatusResponse {
	s.Body = v
	return s
}

type ListMediaBasicInfosRequest struct {
	// 创建时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 页号
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页大小
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 返回值中是否包含文件基础信息
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
}

func (s ListMediaBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosRequest) SetStartTime(v string) *ListMediaBasicInfosRequest {
	s.StartTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetEndTime(v string) *ListMediaBasicInfosRequest {
	s.EndTime = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMediaType(v string) *ListMediaBasicInfosRequest {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetBusinessType(v string) *ListMediaBasicInfosRequest {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSource(v string) *ListMediaBasicInfosRequest {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetCategory(v string) *ListMediaBasicInfosRequest {
	s.Category = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetStatus(v string) *ListMediaBasicInfosRequest {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetNextToken(v string) *ListMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetMaxResults(v int32) *ListMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetSortBy(v string) *ListMediaBasicInfosRequest {
	s.SortBy = &v
	return s
}

func (s *ListMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

type ListMediaBasicInfosResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的媒资总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*ListMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	NextToken  *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListMediaBasicInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBody) SetRequestId(v string) *ListMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetMediaInfos(v []*ListMediaBasicInfosResponseBodyMediaInfos) *ListMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetNextToken(v string) *ListMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// 雪碧图
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *ListMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type ListMediaBasicInfosResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMediaBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListMediaBasicInfosResponse) SetBody(v *ListMediaBasicInfosResponseBody) *ListMediaBasicInfosResponse {
	s.Body = v
	return s
}

type SubmitSubtitleProduceJobRequest struct {
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	OutputConfig  *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	InputConfig   *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	IsAsync       *int64  `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserData      *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSubtitleProduceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSubtitleProduceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSubtitleProduceJobRequest) SetEditingConfig(v string) *SubmitSubtitleProduceJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetType(v string) *SubmitSubtitleProduceJobRequest {
	s.Type = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetOutputConfig(v string) *SubmitSubtitleProduceJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetInputConfig(v string) *SubmitSubtitleProduceJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetIsAsync(v int64) *SubmitSubtitleProduceJobRequest {
	s.IsAsync = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetTitle(v string) *SubmitSubtitleProduceJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetDescription(v string) *SubmitSubtitleProduceJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitSubtitleProduceJobRequest) SetUserData(v string) *SubmitSubtitleProduceJobRequest {
	s.UserData = &v
	return s
}

type SubmitSubtitleProduceJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitSubtitleProduceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSubtitleProduceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSubtitleProduceJobResponseBody) SetRequestId(v string) *SubmitSubtitleProduceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSubtitleProduceJobResponseBody) SetJobId(v string) *SubmitSubtitleProduceJobResponseBody {
	s.JobId = &v
	return s
}

type SubmitSubtitleProduceJobResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSubtitleProduceJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSubtitleProduceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSubtitleProduceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSubtitleProduceJobResponse) SetHeaders(v map[string]*string) *SubmitSubtitleProduceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSubtitleProduceJobResponse) SetBody(v *SubmitSubtitleProduceJobResponseBody) *SubmitSubtitleProduceJobResponse {
	s.Body = v
	return s
}

type SubmitKeyWordCutJobRequest struct {
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s SubmitKeyWordCutJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitKeyWordCutJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitKeyWordCutJobRequest) SetKeyword(v string) *SubmitKeyWordCutJobRequest {
	s.Keyword = &v
	return s
}

func (s *SubmitKeyWordCutJobRequest) SetInputFile(v string) *SubmitKeyWordCutJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitKeyWordCutJobRequest) SetUserData(v string) *SubmitKeyWordCutJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitKeyWordCutJobRequest) SetTitle(v string) *SubmitKeyWordCutJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitKeyWordCutJobRequest) SetDescription(v string) *SubmitKeyWordCutJobRequest {
	s.Description = &v
	return s
}

type SubmitKeyWordCutJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitKeyWordCutJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitKeyWordCutJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitKeyWordCutJobResponseBody) SetRequestId(v string) *SubmitKeyWordCutJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitKeyWordCutJobResponseBody) SetJobId(v string) *SubmitKeyWordCutJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitKeyWordCutJobResponseBody) SetOutput(v string) *SubmitKeyWordCutJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitKeyWordCutJobResponseBody) SetState(v string) *SubmitKeyWordCutJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitKeyWordCutJobResponseBody) SetUserData(v string) *SubmitKeyWordCutJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitKeyWordCutJobResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitKeyWordCutJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitKeyWordCutJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitKeyWordCutJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitKeyWordCutJobResponse) SetHeaders(v map[string]*string) *SubmitKeyWordCutJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitKeyWordCutJobResponse) SetBody(v *SubmitKeyWordCutJobResponseBody) *SubmitKeyWordCutJobResponse {
	s.Body = v
	return s
}

type AddEditingProjectMaterialsRequest struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 素材ID
	MaterialMaps *string `json:"MaterialMaps,omitempty" xml:"MaterialMaps,omitempty"`
}

func (s AddEditingProjectMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsRequest) SetProjectId(v string) *AddEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectMaterialsRequest) SetMaterialMaps(v string) *AddEditingProjectMaterialsRequest {
	s.MaterialMaps = &v
	return s
}

type AddEditingProjectMaterialsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*AddEditingProjectMaterialsResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
}

func (s AddEditingProjectMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBody) SetRequestId(v string) *AddEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetProjectId(v string) *AddEditingProjectMaterialsResponseBody {
	s.ProjectId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetMediaInfos(v []*AddEditingProjectMaterialsResponseBodyMediaInfos) *AddEditingProjectMaterialsResponseBody {
	s.MediaInfos = v
	return s
}

type AddEditingProjectMaterialsResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetMediaId(v string) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetMediaBasicInfo(v *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfos) SetFileInfoList(v []*AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) *AddEditingProjectMaterialsResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// 雪碧图
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo) SetSpriteImages(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

type AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *AddEditingProjectMaterialsResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type AddEditingProjectMaterialsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEditingProjectMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *AddEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *AddEditingProjectMaterialsResponse) SetBody(v *AddEditingProjectMaterialsResponseBody) *AddEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

type SubmitASRJobRequest struct {
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 持续时间
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s SubmitASRJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitASRJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitASRJobRequest) SetInputFile(v string) *SubmitASRJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitASRJobRequest) SetUserData(v string) *SubmitASRJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitASRJobRequest) SetTitle(v string) *SubmitASRJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitASRJobRequest) SetDescription(v string) *SubmitASRJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitASRJobRequest) SetStartTime(v string) *SubmitASRJobRequest {
	s.StartTime = &v
	return s
}

func (s *SubmitASRJobRequest) SetDuration(v string) *SubmitASRJobRequest {
	s.Duration = &v
	return s
}

type SubmitASRJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitASRJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitASRJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitASRJobResponseBody) SetRequestId(v string) *SubmitASRJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetJobId(v string) *SubmitASRJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetOutput(v string) *SubmitASRJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetState(v string) *SubmitASRJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitASRJobResponseBody) SetUserData(v string) *SubmitASRJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitASRJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitASRJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitASRJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitASRJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitASRJobResponse) SetHeaders(v map[string]*string) *SubmitASRJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitASRJobResponse) SetBody(v *SubmitASRJobResponseBody) *SubmitASRJobResponse {
	s.Body = v
	return s
}

type GetEditingProjectRequest struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *GetEditingProjectRequest) SetProjectId(v string) *GetEditingProjectRequest {
	s.ProjectId = &v
	return s
}

type GetEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Project   *GetEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
}

func (s GetEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBody) SetRequestId(v string) *GetEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectResponseBody) SetProject(v *GetEditingProjectResponseBodyProject) *GetEditingProjectResponseBody {
	s.Project = v
	return s
}

type GetEditingProjectResponseBodyProject struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程时间线
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 云剪辑工程最新修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 云剪辑工程总时长
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 云剪辑工程状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 云剪辑工程创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 云剪辑工程模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 云剪辑工程修改来源
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
}

func (s GetEditingProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBodyProject) SetProjectId(v string) *GetEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTitle(v string) *GetEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimeline(v string) *GetEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDescription(v string) *GetEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCoverURL(v string) *GetEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreateTime(v string) *GetEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedTime(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDuration(v int64) *GetEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStatus(v string) *GetEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreateSource(v string) *GetEditingProjectResponseBodyProject {
	s.CreateSource = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTemplateType(v string) *GetEditingProjectResponseBodyProject {
	s.TemplateType = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedSource(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedSource = &v
	return s
}

type GetEditingProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponse) SetHeaders(v map[string]*string) *GetEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *GetEditingProjectResponse) SetBody(v *GetEditingProjectResponseBody) *GetEditingProjectResponse {
	s.Body = v
	return s
}

type ListSysTemplatesRequest struct {
	// 标记当前开始读取的位置，置空表示从头开始
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次读取的最大数据记录数量
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSysTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSysTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListSysTemplatesRequest) SetNextToken(v string) *ListSysTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSysTemplatesRequest) SetMaxResults(v int32) *ListSysTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSysTemplatesRequest) SetType(v string) *ListSysTemplatesRequest {
	s.Type = &v
	return s
}

type ListSysTemplatesResponseBody struct {
	// TotalCount本次请求条件下的数据总量，此参数为可选参数，默认可不返回
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 表示当前调用返回读取到的位置，空代表数据已经读取完毕
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// MaxResults本次请求所返回的最大记录条数
	MaxResults *int32                                   `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Templates  []*ListSysTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListSysTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSysTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSysTemplatesResponseBody) SetTotalCount(v int32) *ListSysTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSysTemplatesResponseBody) SetRequestId(v string) *ListSysTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSysTemplatesResponseBody) SetNextToken(v string) *ListSysTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSysTemplatesResponseBody) SetMaxResults(v int32) *ListSysTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSysTemplatesResponseBody) SetTemplates(v []*ListSysTemplatesResponseBodyTemplates) *ListSysTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListSysTemplatesResponseBodyTemplates struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Config     *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s ListSysTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListSysTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListSysTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListSysTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListSysTemplatesResponseBodyTemplates) SetName(v string) *ListSysTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListSysTemplatesResponseBodyTemplates) SetType(v string) *ListSysTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *ListSysTemplatesResponseBodyTemplates) SetConfig(v string) *ListSysTemplatesResponseBodyTemplates {
	s.Config = &v
	return s
}

type ListSysTemplatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSysTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSysTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSysTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListSysTemplatesResponse) SetHeaders(v map[string]*string) *ListSysTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListSysTemplatesResponse) SetBody(v *ListSysTemplatesResponseBody) *ListSysTemplatesResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	// 模板id，多个id用英文逗号隔开
	TemplateIds *string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetTemplateIds(v string) *DeleteTemplateRequest {
	s.TemplateIds = &v
	return s
}

type DeleteTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type SubmitIRJobRequest struct {
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s SubmitIRJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIRJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIRJobRequest) SetInputFile(v string) *SubmitIRJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitIRJobRequest) SetUserData(v string) *SubmitIRJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitIRJobRequest) SetTitle(v string) *SubmitIRJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitIRJobRequest) SetDescription(v string) *SubmitIRJobRequest {
	s.Description = &v
	return s
}

type SubmitIRJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitIRJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIRJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIRJobResponseBody) SetRequestId(v string) *SubmitIRJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIRJobResponseBody) SetJobId(v string) *SubmitIRJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitIRJobResponseBody) SetOutput(v string) *SubmitIRJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitIRJobResponseBody) SetState(v string) *SubmitIRJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitIRJobResponseBody) SetUserData(v string) *SubmitIRJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitIRJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitIRJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitIRJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIRJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIRJobResponse) SetHeaders(v map[string]*string) *SubmitIRJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIRJobResponse) SetBody(v *SubmitIRJobResponseBody) *SubmitIRJobResponse {
	s.Body = v
	return s
}

type DeleteEditingProjectMaterialsRequest struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 素材ID
	MaterialIds *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	// 素材类型
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
}

func (s DeleteEditingProjectMaterialsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectMaterialsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsRequest) SetProjectId(v string) *DeleteEditingProjectMaterialsRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialIds(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialIds = &v
	return s
}

func (s *DeleteEditingProjectMaterialsRequest) SetMaterialType(v string) *DeleteEditingProjectMaterialsRequest {
	s.MaterialType = &v
	return s
}

type DeleteEditingProjectMaterialsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEditingProjectMaterialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsResponseBody) SetRequestId(v string) *DeleteEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEditingProjectMaterialsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEditingProjectMaterialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEditingProjectMaterialsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectMaterialsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectMaterialsResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectMaterialsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectMaterialsResponse) SetBody(v *DeleteEditingProjectMaterialsResponseBody) *DeleteEditingProjectMaterialsResponse {
	s.Body = v
	return s
}

type SearchEditingProjectRequest struct {
	// CreateTime（创建时间）的开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// CreationTime（创建时间）的结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 云剪辑工程状态。多个用逗号分隔
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 结果排序方式
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 分页参数
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页参数
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SearchEditingProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectRequest) SetStartTime(v string) *SearchEditingProjectRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetEndTime(v string) *SearchEditingProjectRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEditingProjectRequest) SetStatus(v string) *SearchEditingProjectRequest {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectRequest) SetSortBy(v string) *SearchEditingProjectRequest {
	s.SortBy = &v
	return s
}

func (s *SearchEditingProjectRequest) SetNextToken(v string) *SearchEditingProjectRequest {
	s.NextToken = &v
	return s
}

func (s *SearchEditingProjectRequest) SetMaxResults(v int64) *SearchEditingProjectRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchEditingProjectRequest) SetCreateSource(v string) *SearchEditingProjectRequest {
	s.CreateSource = &v
	return s
}

func (s *SearchEditingProjectRequest) SetTemplateType(v string) *SearchEditingProjectRequest {
	s.TemplateType = &v
	return s
}

type SearchEditingProjectResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 云剪辑工程列表
	ProjectList []*SearchEditingProjectResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// 云剪辑工程总数
	MaxResults *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s SearchEditingProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBody) SetRequestId(v string) *SearchEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetProjectList(v []*SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody {
	s.ProjectList = v
	return s
}

func (s *SearchEditingProjectResponseBody) SetMaxResults(v int64) *SearchEditingProjectResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetTotalCount(v int64) *SearchEditingProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetNextToken(v string) *SearchEditingProjectResponseBody {
	s.NextToken = &v
	return s
}

type SearchEditingProjectResponseBodyProjectList struct {
	// 云剪辑工程ID
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 云剪辑工程标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 云剪辑工程时间线
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// 云剪辑工程描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 云剪辑工程封面
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 云剪辑工程创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 云剪辑工程最新修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 云剪辑工程总时长
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 云剪辑工程状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 云剪辑工程合成失败的错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 云剪辑工程合成失败的消息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 最后一次修改来源
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// 模板类型
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SearchEditingProjectResponseBodyProjectList) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProjectId(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTitle(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Title = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTimeline(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Timeline = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDescription(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Description = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCoverURL(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CoverURL = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetCreateTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CreateTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetModifiedTime(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ModifiedTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetDuration(v int64) *SearchEditingProjectResponseBodyProjectList {
	s.Duration = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetStatus(v string) *SearchEditingProjectResponseBodyProjectList {
	s.Status = &v
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

func (s *SearchEditingProjectResponseBodyProjectList) SetCreateSource(v string) *SearchEditingProjectResponseBodyProjectList {
	s.CreateSource = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetModifiedSource(v string) *SearchEditingProjectResponseBodyProjectList {
	s.ModifiedSource = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) SetTemplateType(v string) *SearchEditingProjectResponseBodyProjectList {
	s.TemplateType = &v
	return s
}

type SearchEditingProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchEditingProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponse) SetHeaders(v map[string]*string) *SearchEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *SearchEditingProjectResponse) SetBody(v *SearchEditingProjectResponseBody) *SearchEditingProjectResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// 模板类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 搜索关键词，可以根据模板id和title搜索
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 排序参数，默认根据创建时间倒序
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetType(v string) *ListTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListTemplatesRequest) SetStatus(v string) *ListTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplatesRequest) SetCreateSource(v string) *ListTemplatesRequest {
	s.CreateSource = &v
	return s
}

func (s *ListTemplatesRequest) SetKeyword(v string) *ListTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListTemplatesRequest) SetSortType(v string) *ListTemplatesRequest {
	s.SortType = &v
	return s
}

type ListTemplatesResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 本次请求条件下的数据总量。
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Templates  []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 模板类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 模板配置
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 预览素材
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 创建来源
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// 修改来源
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// 预览素材状态
	PreviewMediaStatus *string `json:"PreviewMediaStatus,omitempty" xml:"PreviewMediaStatus,omitempty"`
	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 封面URL
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetName(v string) *ListTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetType(v string) *ListTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetConfig(v string) *ListTemplatesResponseBodyTemplates {
	s.Config = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPreviewMedia(v string) *ListTemplatesResponseBodyTemplates {
	s.PreviewMedia = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetStatus(v string) *ListTemplatesResponseBodyTemplates {
	s.Status = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateSource(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateSource = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetModifiedSource(v string) *ListTemplatesResponseBodyTemplates {
	s.ModifiedSource = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPreviewMediaStatus(v string) *ListTemplatesResponseBodyTemplates {
	s.PreviewMediaStatus = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreationTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreationTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetModifiedTime(v string) *ListTemplatesResponseBodyTemplates {
	s.ModifiedTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCoverURL(v string) *ListTemplatesResponseBodyTemplates {
	s.CoverURL = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type DeleteEditingProjectsRequest struct {
	// 云剪辑工程ID。支持多个云剪辑工程，以逗号分隔。
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s DeleteEditingProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsRequest) SetProjectIds(v string) *DeleteEditingProjectsRequest {
	s.ProjectIds = &v
	return s
}

type DeleteEditingProjectsResponseBody struct {
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IgnoredList *string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty"`
}

func (s DeleteEditingProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponseBody) SetRequestId(v string) *DeleteEditingProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEditingProjectsResponseBody) SetIgnoredList(v string) *DeleteEditingProjectsResponseBody {
	s.IgnoredList = &v
	return s
}

type DeleteEditingProjectsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEditingProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEditingProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEditingProjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectsResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectsResponse) SetBody(v *DeleteEditingProjectsResponseBody) *DeleteEditingProjectsResponse {
	s.Body = v
	return s
}

type GetMediaInfoRequest struct {
	MediaId    *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	InputURL   *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	OutputType *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
}

func (s GetMediaInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMediaInfoRequest) SetMediaId(v string) *GetMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoRequest) SetInputURL(v string) *GetMediaInfoRequest {
	s.InputURL = &v
	return s
}

func (s *GetMediaInfoRequest) SetOutputType(v string) *GetMediaInfoRequest {
	s.OutputType = &v
	return s
}

type GetMediaInfoResponseBody struct {
	// RequestId
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MediaInfo *GetMediaInfoResponseBodyMediaInfo `json:"MediaInfo,omitempty" xml:"MediaInfo,omitempty" type:"Struct"`
}

func (s GetMediaInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBody) SetRequestId(v string) *GetMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaInfoResponseBody) SetMediaInfo(v *GetMediaInfoResponseBodyMediaInfo) *GetMediaInfoResponseBody {
	s.MediaInfo = v
	return s
}

type GetMediaInfoResponseBodyMediaInfo struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// 其他元数据
	DynamicMetaDataList []*GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList `json:"DynamicMetaDataList,omitempty" xml:"DynamicMetaDataList,omitempty" type:"Repeated"`
	// AIMetadata
	AiRoughDataList []*GetMediaInfoResponseBodyMediaInfoAiRoughDataList `json:"AiRoughDataList,omitempty" xml:"AiRoughDataList,omitempty" type:"Repeated"`
	// FileInfos
	FileInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoResponseBodyMediaInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetMediaBasicInfo(v *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) *GetMediaInfoResponseBodyMediaInfo {
	s.MediaBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetDynamicMetaDataList(v []*GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) *GetMediaInfoResponseBodyMediaInfo {
	s.DynamicMetaDataList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetAiRoughDataList(v []*GetMediaInfoResponseBodyMediaInfoAiRoughDataList) *GetMediaInfoResponseBodyMediaInfo {
	s.AiRoughDataList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfo) SetFileInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoList) *GetMediaInfoResponseBodyMediaInfo {
	s.FileInfoList = v
	return s
}

type GetMediaInfoResponseBodyMediaInfoMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	// 雪碧图
	SpriteImages *string `json:"SpriteImages,omitempty" xml:"SpriteImages,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaId(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetInputURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetBusinessType(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSource(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetTitle(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDescription(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCategory(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetMediaTags(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCoverURL(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetUserData(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetStatus(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetCreateTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetModifiedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetDeletedTime(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo) SetSpriteImages(v string) *GetMediaInfoResponseBodyMediaInfoMediaBasicInfo {
	s.SpriteImages = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList struct {
	// 开始时间
	In *float32 `json:"In,omitempty" xml:"In,omitempty"`
	// 结束时间
	Out *float32 `json:"Out,omitempty" xml:"Out,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 元数据json string
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetIn(v float32) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.In = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetOut(v float32) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Out = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetType(v string) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Type = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList) SetData(v string) *GetMediaInfoResponseBodyMediaInfoDynamicMetaDataList {
	s.Data = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoAiRoughDataList struct {
	// AI类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// AI原始结果
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoAiRoughDataList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataList) SetType(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataList {
	s.Type = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoAiRoughDataList) SetResult(v string) *GetMediaInfoResponseBodyMediaInfoAiRoughDataList {
	s.Result = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
	// 音频流信息，一个媒资可能有多条音频流
	AudioStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList `json:"AudioStreamInfoList,omitempty" xml:"AudioStreamInfoList,omitempty" type:"Repeated"`
	// 视频流信息，一个媒资可能有多条视频流
	VideoStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList `json:"VideoStreamInfoList,omitempty" xml:"VideoStreamInfoList,omitempty" type:"Repeated"`
	// 字幕流信息，一个媒资可能有多条字幕流
	SubtitleStreamInfoList []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList `json:"SubtitleStreamInfoList,omitempty" xml:"SubtitleStreamInfoList,omitempty" type:"Repeated"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetFileBasicInfo(v *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.FileBasicInfo = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetAudioStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.AudioStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetVideoStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.VideoStreamInfoList = v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoList) SetSubtitleStreamInfoList(v []*GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) *GetMediaInfoResponseBodyMediaInfoFileInfoList {
	s.SubtitleStreamInfoList = v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileStatus(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileType(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileSize(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFileUrl(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetRegion(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetFormatName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList struct {
	// 音频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 编码预置
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// 采样格式
	SampleFmt *string `json:"SampleFmt,omitempty" xml:"SampleFmt,omitempty"`
	// 采样率
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// 声道数
	Channels *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
	// 声道输出样式
	ChannelLayout *string `json:"ChannelLayout,omitempty" xml:"ChannelLayout,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 音频帧率
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// 总帧数
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetSampleRate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.SampleRate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannels(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Channels = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetChannelLayout(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.ChannelLayout = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListAudioStreamInfoList {
	s.Lang = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList struct {
	// 视频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码预置
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// 是否有B帧
	HasBFrames *string `json:"HasBFrames,omitempty" xml:"HasBFrames,omitempty"`
	// 编码信号分辨率比
	Sar *string `json:"Sar,omitempty" xml:"Sar,omitempty"`
	// 编码显示分辨率比
	Dar *string `json:"Dar,omitempty" xml:"Dar,omitempty"`
	// 像素格式
	PixFmt *string `json:"PixFmt,omitempty" xml:"PixFmt,omitempty"`
	// 编码等级
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// 视频帧率
	Fps *string `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// 平均帧率
	AvgFPS *string `json:"AvgFPS,omitempty" xml:"AvgFPS,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 总帧数
	NumFrames *string `json:"NumFrames,omitempty" xml:"NumFrames,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 旋转
	Rotate *string `json:"Rotate,omitempty" xml:"Rotate,omitempty"`
	// 总帧数
	NbFrames *string `json:"Nb_frames,omitempty" xml:"Nb_frames,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetProfile(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Profile = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetWidth(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Width = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHeight(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Height = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetHasBFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.HasBFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetSar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Sar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDar(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Dar = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetPixFmt(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.PixFmt = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLevel(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Level = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetFps(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Fps = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetAvgFPS(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.AvgFPS = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetBitrate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Bitrate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNumFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NumFrames = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Lang = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetRotate(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.Rotate = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList) SetNbFrames(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListVideoStreamInfoList {
	s.NbFrames = &v
	return s
}

type GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList struct {
	// 音频流序号
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// 编码格式简述名
	CodecName *string `json:"CodecName,omitempty" xml:"CodecName,omitempty"`
	// 编码格式长述名
	CodecLongName *string `json:"CodecLongName,omitempty" xml:"CodecLongName,omitempty"`
	// 编码时基
	CodecTimeBase *string `json:"CodecTimeBase,omitempty" xml:"CodecTimeBase,omitempty"`
	// 编码格式标记文本
	CodecTagString *string `json:"CodecTagString,omitempty" xml:"CodecTagString,omitempty"`
	// 编码格式标记
	CodecTag *string `json:"CodecTag,omitempty" xml:"CodecTag,omitempty"`
	// 时基
	Timebase *string `json:"Timebase,omitempty" xml:"Timebase,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetIndex(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Index = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecLongName(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecLongName = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTimeBase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTimeBase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTagString(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTagString = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetCodecTag(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.CodecTag = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetTimebase(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Timebase = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetStartTime(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.StartTime = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetDuration(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Duration = &v
	return s
}

func (s *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList) SetLang(v string) *GetMediaInfoResponseBodyMediaInfoFileInfoListSubtitleStreamInfoList {
	s.Lang = &v
	return s
}

type GetMediaInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMediaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMediaInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMediaInfoResponse) GoString() string {
	return s.String()
}

func (s *GetMediaInfoResponse) SetHeaders(v map[string]*string) *GetMediaInfoResponse {
	s.Headers = v
	return s
}

func (s *GetMediaInfoResponse) SetBody(v *GetMediaInfoResponseBody) *GetMediaInfoResponse {
	s.Body = v
	return s
}

type SubmitSmartJobRequest struct {
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	OutputConfig  *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	InputConfig   *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserData      *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	JobType       *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
}

func (s SubmitSmartJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmartJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartJobRequest) SetEditingConfig(v string) *SubmitSmartJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitSmartJobRequest) SetOutputConfig(v string) *SubmitSmartJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSmartJobRequest) SetInputConfig(v string) *SubmitSmartJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSmartJobRequest) SetTitle(v string) *SubmitSmartJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitSmartJobRequest) SetDescription(v string) *SubmitSmartJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitSmartJobRequest) SetUserData(v string) *SubmitSmartJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSmartJobRequest) SetJobType(v string) *SubmitSmartJobRequest {
	s.JobType = &v
	return s
}

type SubmitSmartJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitSmartJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmartJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSmartJobResponseBody) SetRequestId(v string) *SubmitSmartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSmartJobResponseBody) SetJobId(v string) *SubmitSmartJobResponseBody {
	s.JobId = &v
	return s
}

type SubmitSmartJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitSmartJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSmartJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSmartJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSmartJobResponse) SetHeaders(v map[string]*string) *SubmitSmartJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSmartJobResponse) SetBody(v *SubmitSmartJobResponseBody) *SubmitSmartJobResponse {
	s.Body = v
	return s
}

type SubmitDelogoJobRequest struct {
	// 输入文件
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输出bucket
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// 输入文件类型
	InputType *string `json:"InputType,omitempty" xml:"InputType,omitempty"`
	// 是否强制覆盖现有OSS文件
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// 输出类型
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
}

func (s SubmitDelogoJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitDelogoJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDelogoJobRequest) SetInputFile(v string) *SubmitDelogoJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetUserData(v string) *SubmitDelogoJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetTitle(v string) *SubmitDelogoJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetDescription(v string) *SubmitDelogoJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetOutputConfig(v string) *SubmitDelogoJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetInputType(v string) *SubmitDelogoJobRequest {
	s.InputType = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetOverwrite(v bool) *SubmitDelogoJobRequest {
	s.Overwrite = &v
	return s
}

func (s *SubmitDelogoJobRequest) SetOutputMediaTarget(v string) *SubmitDelogoJobRequest {
	s.OutputMediaTarget = &v
	return s
}

type SubmitDelogoJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitDelogoJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitDelogoJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDelogoJobResponseBody) SetRequestId(v string) *SubmitDelogoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDelogoJobResponseBody) SetJobId(v string) *SubmitDelogoJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitDelogoJobResponseBody) SetOutput(v string) *SubmitDelogoJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitDelogoJobResponseBody) SetState(v string) *SubmitDelogoJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitDelogoJobResponseBody) SetUserData(v string) *SubmitDelogoJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitDelogoJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitDelogoJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitDelogoJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitDelogoJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDelogoJobResponse) SetHeaders(v map[string]*string) *SubmitDelogoJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDelogoJobResponse) SetBody(v *SubmitDelogoJobResponseBody) *SubmitDelogoJobResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	// 模板ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// 模板名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 参见Timeline模板Config文档
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 模板封面
	CoverUrl *string `json:"CoverUrl,omitempty" xml:"CoverUrl,omitempty"`
	// 预览视频媒资id
	PreviewMedia *string `json:"PreviewMedia,omitempty" xml:"PreviewMedia,omitempty"`
	// 模板状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 修改来源，默认OpenAPI
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetName(v string) *UpdateTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateTemplateRequest) SetConfig(v string) *UpdateTemplateRequest {
	s.Config = &v
	return s
}

func (s *UpdateTemplateRequest) SetCoverUrl(v string) *UpdateTemplateRequest {
	s.CoverUrl = &v
	return s
}

func (s *UpdateTemplateRequest) SetPreviewMedia(v string) *UpdateTemplateRequest {
	s.PreviewMedia = &v
	return s
}

func (s *UpdateTemplateRequest) SetStatus(v string) *UpdateTemplateRequest {
	s.Status = &v
	return s
}

func (s *UpdateTemplateRequest) SetSource(v string) *UpdateTemplateRequest {
	s.Source = &v
	return s
}

type UpdateTemplateResponseBody struct {
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type SubmitAudioProduceJobRequest struct {
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	OutputConfig  *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	InputConfig   *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserData      *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Overwrite     *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
}

func (s SubmitAudioProduceJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioProduceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobRequest) SetEditingConfig(v string) *SubmitAudioProduceJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetOutputConfig(v string) *SubmitAudioProduceJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetInputConfig(v string) *SubmitAudioProduceJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetTitle(v string) *SubmitAudioProduceJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetDescription(v string) *SubmitAudioProduceJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetUserData(v string) *SubmitAudioProduceJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAudioProduceJobRequest) SetOverwrite(v bool) *SubmitAudioProduceJobRequest {
	s.Overwrite = &v
	return s
}

type SubmitAudioProduceJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 任务ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 任务状态
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitAudioProduceJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioProduceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobResponseBody) SetRequestId(v string) *SubmitAudioProduceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) SetJobId(v string) *SubmitAudioProduceJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) SetState(v string) *SubmitAudioProduceJobResponseBody {
	s.State = &v
	return s
}

type SubmitAudioProduceJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitAudioProduceJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAudioProduceJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioProduceJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobResponse) SetHeaders(v map[string]*string) *SubmitAudioProduceJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAudioProduceJobResponse) SetBody(v *SubmitAudioProduceJobResponseBody) *SubmitAudioProduceJobResponse {
	s.Body = v
	return s
}

type SubmitMediaProducingJobRequest struct {
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Timeline          *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	TemplateId        *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	ClipsParam        *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	ProjectMetadata   *string `json:"ProjectMetadata,omitempty" xml:"ProjectMetadata,omitempty"`
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	UserData          *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SubmitMediaProducingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobRequest) SetProjectId(v string) *SubmitMediaProducingJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTimeline(v string) *SubmitMediaProducingJobRequest {
	s.Timeline = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetTemplateId(v string) *SubmitMediaProducingJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetClipsParam(v string) *SubmitMediaProducingJobRequest {
	s.ClipsParam = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetProjectMetadata(v string) *SubmitMediaProducingJobRequest {
	s.ProjectMetadata = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaTarget(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaTarget = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetOutputMediaConfig(v string) *SubmitMediaProducingJobRequest {
	s.OutputMediaConfig = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetUserData(v string) *SubmitMediaProducingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetClientToken(v string) *SubmitMediaProducingJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SubmitMediaProducingJobRequest) SetSource(v string) *SubmitMediaProducingJobRequest {
	s.Source = &v
	return s
}

type SubmitMediaProducingJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 剪辑工程Id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 合成作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 合成媒资Id
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s SubmitMediaProducingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponseBody) SetRequestId(v string) *SubmitMediaProducingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetProjectId(v string) *SubmitMediaProducingJobResponseBody {
	s.ProjectId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetJobId(v string) *SubmitMediaProducingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMediaProducingJobResponseBody) SetMediaId(v string) *SubmitMediaProducingJobResponseBody {
	s.MediaId = &v
	return s
}

type SubmitMediaProducingJobResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitMediaProducingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitMediaProducingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMediaProducingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMediaProducingJobResponse) SetHeaders(v map[string]*string) *SubmitMediaProducingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMediaProducingJobResponse) SetBody(v *SubmitMediaProducingJobResponseBody) *SubmitMediaProducingJobResponse {
	s.Body = v
	return s
}

type UpdateSmartJobRequest struct {
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	FEExtend *string `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
}

func (s UpdateSmartJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmartJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartJobRequest) SetJobId(v string) *UpdateSmartJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateSmartJobRequest) SetFEExtend(v string) *UpdateSmartJobRequest {
	s.FEExtend = &v
	return s
}

type UpdateSmartJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	FEExtend  *string `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
}

func (s UpdateSmartJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmartJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSmartJobResponseBody) SetRequestId(v string) *UpdateSmartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSmartJobResponseBody) SetJobId(v string) *UpdateSmartJobResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateSmartJobResponseBody) SetFEExtend(v string) *UpdateSmartJobResponseBody {
	s.FEExtend = &v
	return s
}

type UpdateSmartJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSmartJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSmartJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSmartJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateSmartJobResponse) SetHeaders(v map[string]*string) *UpdateSmartJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateSmartJobResponse) SetBody(v *UpdateSmartJobResponseBody) *UpdateSmartJobResponse {
	s.Body = v
	return s
}

type ListAllPublicMediaTagsRequest struct {
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
}

func (s ListAllPublicMediaTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsRequest) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsRequest) SetBusinessType(v string) *ListAllPublicMediaTagsRequest {
	s.BusinessType = &v
	return s
}

type ListAllPublicMediaTagsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 公共素材库标签列表
	MediaTagList []*ListAllPublicMediaTagsResponseBodyMediaTagList `json:"MediaTagList,omitempty" xml:"MediaTagList,omitempty" type:"Repeated"`
}

func (s ListAllPublicMediaTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBody) SetRequestId(v string) *ListAllPublicMediaTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBody) SetMediaTagList(v []*ListAllPublicMediaTagsResponseBodyMediaTagList) *ListAllPublicMediaTagsResponseBody {
	s.MediaTagList = v
	return s
}

type ListAllPublicMediaTagsResponseBodyMediaTagList struct {
	// 素材标签id
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
	// 素材标签中文名
	MediaTagNameChinese *string `json:"MediaTagNameChinese,omitempty" xml:"MediaTagNameChinese,omitempty"`
	// 素材标签英文名
	MediaTagNameEnglish *string `json:"MediaTagNameEnglish,omitempty" xml:"MediaTagNameEnglish,omitempty"`
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponseBodyMediaTagList) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagId(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagId = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagNameChinese(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagNameChinese = &v
	return s
}

func (s *ListAllPublicMediaTagsResponseBodyMediaTagList) SetMediaTagNameEnglish(v string) *ListAllPublicMediaTagsResponseBodyMediaTagList {
	s.MediaTagNameEnglish = &v
	return s
}

type ListAllPublicMediaTagsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAllPublicMediaTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAllPublicMediaTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllPublicMediaTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAllPublicMediaTagsResponse) SetHeaders(v map[string]*string) *ListAllPublicMediaTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAllPublicMediaTagsResponse) SetBody(v *ListAllPublicMediaTagsResponseBody) *ListAllPublicMediaTagsResponse {
	s.Body = v
	return s
}

type SubmitMattingJobRequest struct {
	// 输入文件
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输出bucket
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// 输入文件类型
	InputType *string `json:"InputType,omitempty" xml:"InputType,omitempty"`
	// 是否强制覆盖现有OSS文件
	Overwrite *string `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// 输出类型
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
}

func (s SubmitMattingJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitMattingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMattingJobRequest) SetInputFile(v string) *SubmitMattingJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitMattingJobRequest) SetUserData(v string) *SubmitMattingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMattingJobRequest) SetTitle(v string) *SubmitMattingJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitMattingJobRequest) SetDescription(v string) *SubmitMattingJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitMattingJobRequest) SetOutputConfig(v string) *SubmitMattingJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitMattingJobRequest) SetInputType(v string) *SubmitMattingJobRequest {
	s.InputType = &v
	return s
}

func (s *SubmitMattingJobRequest) SetOverwrite(v string) *SubmitMattingJobRequest {
	s.Overwrite = &v
	return s
}

func (s *SubmitMattingJobRequest) SetOutputMediaTarget(v string) *SubmitMattingJobRequest {
	s.OutputMediaTarget = &v
	return s
}

type SubmitMattingJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMattingJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitMattingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitMattingJobResponseBody) SetRequestId(v string) *SubmitMattingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitMattingJobResponseBody) SetJobId(v string) *SubmitMattingJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitMattingJobResponseBody) SetOutput(v string) *SubmitMattingJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitMattingJobResponseBody) SetState(v string) *SubmitMattingJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitMattingJobResponseBody) SetUserData(v string) *SubmitMattingJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitMattingJobResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitMattingJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitMattingJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitMattingJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitMattingJobResponse) SetHeaders(v map[string]*string) *SubmitMattingJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitMattingJobResponse) SetBody(v *SubmitMattingJobResponseBody) *SubmitMattingJobResponse {
	s.Body = v
	return s
}

type GetEventCallbackResponseBody struct {
	// Id of the request
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallbackQueueName *string `json:"CallbackQueueName,omitempty" xml:"CallbackQueueName,omitempty"`
}

func (s GetEventCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEventCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventCallbackResponseBody) SetRequestId(v string) *GetEventCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventCallbackResponseBody) SetCallbackQueueName(v string) *GetEventCallbackResponseBody {
	s.CallbackQueueName = &v
	return s
}

type GetEventCallbackResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetEventCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEventCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEventCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetEventCallbackResponse) SetHeaders(v map[string]*string) *GetEventCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetEventCallbackResponse) SetBody(v *GetEventCallbackResponseBody) *GetEventCallbackResponse {
	s.Body = v
	return s
}

type ListPublicMediaBasicInfosRequest struct {
	// 标签
	MediaTagId *string `json:"MediaTagId,omitempty" xml:"MediaTagId,omitempty"`
	// 下一次读取的位置
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 分页大小
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 返回值中是否包含文件基础信息
	IncludeFileBasicInfo *bool `json:"IncludeFileBasicInfo,omitempty" xml:"IncludeFileBasicInfo,omitempty"`
}

func (s ListPublicMediaBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosRequest) SetMediaTagId(v string) *ListPublicMediaBasicInfosRequest {
	s.MediaTagId = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetNextToken(v string) *ListPublicMediaBasicInfosRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetMaxResults(v int32) *ListPublicMediaBasicInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicMediaBasicInfosRequest) SetIncludeFileBasicInfo(v bool) *ListPublicMediaBasicInfosRequest {
	s.IncludeFileBasicInfo = &v
	return s
}

type ListPublicMediaBasicInfosResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合要求的媒资总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 符合要求的媒资集合
	MediaInfos []*ListPublicMediaBasicInfosResponseBodyMediaInfos `json:"MediaInfos,omitempty" xml:"MediaInfos,omitempty" type:"Repeated"`
	NextToken  *string                                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	MaxResults *int32                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBody) SetRequestId(v string) *ListPublicMediaBasicInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetTotalCount(v int64) *ListPublicMediaBasicInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMediaInfos(v []*ListPublicMediaBasicInfosResponseBodyMediaInfos) *ListPublicMediaBasicInfosResponseBody {
	s.MediaInfos = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetNextToken(v string) *ListPublicMediaBasicInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBody) SetMaxResults(v int32) *ListPublicMediaBasicInfosResponseBody {
	s.MaxResults = &v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfos struct {
	// 媒资ID
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// BasicInfo
	MediaBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo `json:"MediaBasicInfo,omitempty" xml:"MediaBasicInfo,omitempty" type:"Struct"`
	// FileInfos
	FileInfoList []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfos) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetMediaBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.MediaBasicInfo = v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfos) SetFileInfoList(v []*ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) *ListPublicMediaBasicInfosResponseBodyMediaInfos {
	s.FileInfoList = v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo struct {
	// MediaId
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// 待注册的媒资在相应系统中的地址
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// 媒资媒体类型
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// 媒资业务类型
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// 来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// 内容描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分类
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// 标签
	MediaTags *string `json:"MediaTags,omitempty" xml:"MediaTags,omitempty"`
	// 封面地址
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// 用户数据
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// 截图
	Snapshots *string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty"`
	// 资源状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 转码状态
	TranscodeStatus *string `json:"TranscodeStatus,omitempty" xml:"TranscodeStatus,omitempty"`
	// 媒资创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 媒资修改时间
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// 媒资删除时间
	DeletedTime *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaId(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaId = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetInputURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.InputURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetBusinessType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.BusinessType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSource(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Source = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTitle(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Title = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDescription(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Description = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCategory(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Category = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetMediaTags(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.MediaTags = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCoverURL(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CoverURL = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetUserData(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.UserData = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetSnapshots(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Snapshots = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.Status = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetTranscodeStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.TranscodeStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetCreateTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.CreateTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetModifiedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.ModifiedTime = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo) SetDeletedTime(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosMediaBasicInfo {
	s.DeletedTime = &v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList struct {
	// 文件基础信息，包含时长，大小等
	FileBasicInfo *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo `json:"FileBasicInfo,omitempty" xml:"FileBasicInfo,omitempty" type:"Struct"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList) SetFileBasicInfo(v *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoList {
	s.FileBasicInfo = v
	return s
}

type ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件状态
	FileStatus *string `json:"FileStatus,omitempty" xml:"FileStatus,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件大小（字节）
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件oss地址
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// 文件存储区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 封装格式
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// 时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 码率
	Bitrate *string `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// 宽
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// 高
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileStatus(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileStatus = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileType(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileType = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileSize(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileSize = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFileUrl(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FileUrl = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetRegion(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Region = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetFormatName(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.FormatName = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetDuration(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Duration = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetBitrate(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Bitrate = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetWidth(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Width = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo) SetHeight(v string) *ListPublicMediaBasicInfosResponseBodyMediaInfosFileInfoListFileBasicInfo {
	s.Height = &v
	return s
}

type ListPublicMediaBasicInfosResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPublicMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublicMediaBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListPublicMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListPublicMediaBasicInfosResponse) SetBody(v *ListPublicMediaBasicInfosResponseBody) *ListPublicMediaBasicInfosResponse {
	s.Body = v
	return s
}

type SubmitCoverJobRequest struct {
	// 输入文件
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输出bucket
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
}

func (s SubmitCoverJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCoverJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCoverJobRequest) SetInputFile(v string) *SubmitCoverJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitCoverJobRequest) SetUserData(v string) *SubmitCoverJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitCoverJobRequest) SetTitle(v string) *SubmitCoverJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitCoverJobRequest) SetDescription(v string) *SubmitCoverJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitCoverJobRequest) SetOutputConfig(v string) *SubmitCoverJobRequest {
	s.OutputConfig = &v
	return s
}

type SubmitCoverJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitCoverJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCoverJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCoverJobResponseBody) SetRequestId(v string) *SubmitCoverJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCoverJobResponseBody) SetJobId(v string) *SubmitCoverJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitCoverJobResponseBody) SetOutput(v string) *SubmitCoverJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitCoverJobResponseBody) SetState(v string) *SubmitCoverJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitCoverJobResponseBody) SetUserData(v string) *SubmitCoverJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitCoverJobResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitCoverJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCoverJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCoverJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitCoverJobResponse) SetHeaders(v map[string]*string) *SubmitCoverJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitCoverJobResponse) SetBody(v *SubmitCoverJobResponseBody) *SubmitCoverJobResponse {
	s.Body = v
	return s
}

type GetSmartHandleJobRequest struct {
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	WithAiResult *string `json:"WithAiResult,omitempty" xml:"WithAiResult,omitempty"`
}

func (s GetSmartHandleJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSmartHandleJobRequest) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobRequest) SetJobId(v string) *GetSmartHandleJobRequest {
	s.JobId = &v
	return s
}

func (s *GetSmartHandleJobRequest) SetWithAiResult(v string) *GetSmartHandleJobRequest {
	s.WithAiResult = &v
	return s
}

type GetSmartHandleJobResponseBody struct {
	// Id of the request
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId        *string                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output       *string                                    `json:"Output,omitempty" xml:"Output,omitempty"`
	State        *string                                    `json:"State,omitempty" xml:"State,omitempty"`
	UserData     *string                                    `json:"UserData,omitempty" xml:"UserData,omitempty"`
	FEExtend     *string                                    `json:"FEExtend,omitempty" xml:"FEExtend,omitempty"`
	SmartJobInfo *GetSmartHandleJobResponseBodySmartJobInfo `json:"SmartJobInfo,omitempty" xml:"SmartJobInfo,omitempty" type:"Struct"`
}

func (s GetSmartHandleJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSmartHandleJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBody) SetRequestId(v string) *GetSmartHandleJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetJobId(v string) *GetSmartHandleJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetOutput(v string) *GetSmartHandleJobResponseBody {
	s.Output = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetState(v string) *GetSmartHandleJobResponseBody {
	s.State = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetUserData(v string) *GetSmartHandleJobResponseBody {
	s.UserData = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetFEExtend(v string) *GetSmartHandleJobResponseBody {
	s.FEExtend = &v
	return s
}

func (s *GetSmartHandleJobResponseBody) SetSmartJobInfo(v *GetSmartHandleJobResponseBodySmartJobInfo) *GetSmartHandleJobResponseBody {
	s.SmartJobInfo = v
	return s
}

type GetSmartHandleJobResponseBodySmartJobInfo struct {
	Title         *string                                               `json:"Title,omitempty" xml:"Title,omitempty"`
	Description   *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId        *string                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	EditingConfig *string                                               `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	InputConfig   *GetSmartHandleJobResponseBodySmartJobInfoInputConfig `json:"InputConfig,omitempty" xml:"InputConfig,omitempty" type:"Struct"`
	OutputConfig  *string                                               `json:"outputConfig,omitempty" xml:"outputConfig,omitempty"`
	CreateTime    *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime  *string                                               `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	JobType       *string                                               `json:"JobType,omitempty" xml:"JobType,omitempty"`
}

func (s GetSmartHandleJobResponseBodySmartJobInfo) String() string {
	return tea.Prettify(s)
}

func (s GetSmartHandleJobResponseBodySmartJobInfo) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetTitle(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.Title = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetDescription(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.Description = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetUserId(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.UserId = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetEditingConfig(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.EditingConfig = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetInputConfig(v *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.InputConfig = v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetOutputConfig(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.OutputConfig = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetCreateTime(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetModifiedTime(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.ModifiedTime = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfo) SetJobType(v string) *GetSmartHandleJobResponseBodySmartJobInfo {
	s.JobType = &v
	return s
}

type GetSmartHandleJobResponseBodySmartJobInfoInputConfig struct {
	InputFile     *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	JobParameters *string `json:"JobParameters,omitempty" xml:"JobParameters,omitempty"`
}

func (s GetSmartHandleJobResponseBodySmartJobInfoInputConfig) String() string {
	return tea.Prettify(s)
}

func (s GetSmartHandleJobResponseBodySmartJobInfoInputConfig) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) SetInputFile(v string) *GetSmartHandleJobResponseBodySmartJobInfoInputConfig {
	s.InputFile = &v
	return s
}

func (s *GetSmartHandleJobResponseBodySmartJobInfoInputConfig) SetJobParameters(v string) *GetSmartHandleJobResponseBodySmartJobInfoInputConfig {
	s.JobParameters = &v
	return s
}

type GetSmartHandleJobResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSmartHandleJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSmartHandleJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSmartHandleJobResponse) GoString() string {
	return s.String()
}

func (s *GetSmartHandleJobResponse) SetHeaders(v map[string]*string) *GetSmartHandleJobResponse {
	s.Headers = v
	return s
}

func (s *GetSmartHandleJobResponse) SetBody(v *GetSmartHandleJobResponseBody) *GetSmartHandleJobResponse {
	s.Body = v
	return s
}

type SubmitH2VJobRequest struct {
	// 输入文件
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 输出bucket
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// 输入文件类型
	InputType *string `json:"InputType,omitempty" xml:"InputType,omitempty"`
	// 是否强制覆盖现有OSS文件
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// 输出类型
	OutputMediaTarget *string `json:"OutputMediaTarget,omitempty" xml:"OutputMediaTarget,omitempty"`
}

func (s SubmitH2VJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitH2VJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitH2VJobRequest) SetInputFile(v string) *SubmitH2VJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitH2VJobRequest) SetUserData(v string) *SubmitH2VJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitH2VJobRequest) SetTitle(v string) *SubmitH2VJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitH2VJobRequest) SetDescription(v string) *SubmitH2VJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitH2VJobRequest) SetOutputConfig(v string) *SubmitH2VJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitH2VJobRequest) SetInputType(v string) *SubmitH2VJobRequest {
	s.InputType = &v
	return s
}

func (s *SubmitH2VJobRequest) SetOverwrite(v bool) *SubmitH2VJobRequest {
	s.Overwrite = &v
	return s
}

func (s *SubmitH2VJobRequest) SetOutputMediaTarget(v string) *SubmitH2VJobRequest {
	s.OutputMediaTarget = &v
	return s
}

type SubmitH2VJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitH2VJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitH2VJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitH2VJobResponseBody) SetRequestId(v string) *SubmitH2VJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitH2VJobResponseBody) SetJobId(v string) *SubmitH2VJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitH2VJobResponseBody) SetOutput(v string) *SubmitH2VJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitH2VJobResponseBody) SetState(v string) *SubmitH2VJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitH2VJobResponseBody) SetUserData(v string) *SubmitH2VJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitH2VJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitH2VJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitH2VJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitH2VJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitH2VJobResponse) SetHeaders(v map[string]*string) *SubmitH2VJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitH2VJobResponse) SetBody(v *SubmitH2VJobResponseBody) *SubmitH2VJobResponse {
	s.Body = v
	return s
}

type SubmitPPTCutJobRequest struct {
	InputFile   *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s SubmitPPTCutJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitPPTCutJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitPPTCutJobRequest) SetInputFile(v string) *SubmitPPTCutJobRequest {
	s.InputFile = &v
	return s
}

func (s *SubmitPPTCutJobRequest) SetUserData(v string) *SubmitPPTCutJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitPPTCutJobRequest) SetTitle(v string) *SubmitPPTCutJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitPPTCutJobRequest) SetDescription(v string) *SubmitPPTCutJobRequest {
	s.Description = &v
	return s
}

type SubmitPPTCutJobResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitPPTCutJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitPPTCutJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPPTCutJobResponseBody) SetRequestId(v string) *SubmitPPTCutJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPPTCutJobResponseBody) SetJobId(v string) *SubmitPPTCutJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitPPTCutJobResponseBody) SetOutput(v string) *SubmitPPTCutJobResponseBody {
	s.Output = &v
	return s
}

func (s *SubmitPPTCutJobResponseBody) SetState(v string) *SubmitPPTCutJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitPPTCutJobResponseBody) SetUserData(v string) *SubmitPPTCutJobResponseBody {
	s.UserData = &v
	return s
}

type SubmitPPTCutJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitPPTCutJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitPPTCutJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitPPTCutJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitPPTCutJobResponse) SetHeaders(v map[string]*string) *SubmitPPTCutJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitPPTCutJobResponse) SetBody(v *SubmitPPTCutJobResponseBody) *SubmitPPTCutJobResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("ice.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ice.aliyuncs.com"),
		"ap-south-1":                  tea.String("ice.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-2":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-3":              tea.String("ice.aliyuncs.com"),
		"ap-southeast-5":              tea.String("ice.aliyuncs.com"),
		"cn-beijing":                  tea.String("ice.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("ice.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ice.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ice.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ice.aliyuncs.com"),
		"cn-chengdu":                  tea.String("ice.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ice.aliyuncs.com"),
		"cn-fujian":                   tea.String("ice.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ice.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ice.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ice.aliyuncs.com"),
		"cn-hongkong":                 tea.String("ice.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ice.aliyuncs.com"),
		"cn-huhehaote":                tea.String("ice.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ice.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ice.aliyuncs.com"),
		"cn-qingdao":                  tea.String("ice.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ice.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ice.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ice.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ice.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ice.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ice.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ice.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ice.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ice.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("ice.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ice.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ice.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ice.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("ice.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ice.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ice.aliyuncs.com"),
		"eu-central-1":                tea.String("ice.aliyuncs.com"),
		"eu-west-1":                   tea.String("ice.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ice.aliyuncs.com"),
		"me-east-1":                   tea.String("ice.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ice.aliyuncs.com"),
		"us-east-1":                   tea.String("ice.aliyuncs.com"),
		"us-west-1":                   tea.String("ice.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSmartJobsWithOptions(request *ListSmartJobsRequest, runtime *util.RuntimeOptions) (_result *ListSmartJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSmartJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSmartJobs"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSmartJobs(request *ListSmartJobsRequest) (_result *ListSmartJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSmartJobsResponse{}
	_body, _err := client.ListSmartJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRelatedAuthorizationStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeRelatedAuthorizationStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeRelatedAuthorizationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRelatedAuthorizationStatus"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRelatedAuthorizationStatus() (_result *DescribeRelatedAuthorizationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRelatedAuthorizationStatusResponse{}
	_body, _err := client.DescribeRelatedAuthorizationStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSmartJobWithOptions(request *DeleteSmartJobRequest, runtime *util.RuntimeOptions) (_result *DeleteSmartJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSmartJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSmartJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSmartJob(request *DeleteSmartJobRequest) (_result *DeleteSmartJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSmartJobResponse{}
	_body, _err := client.DeleteSmartJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTemplateWithOptions(request *AddTemplateRequest, runtime *util.RuntimeOptions) (_result *AddTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddTemplate"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTemplate(request *AddTemplateRequest) (_result *AddTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTemplateResponse{}
	_body, _err := client.AddTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEditingProjectWithOptions(request *UpdateEditingProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEditingProject(request *UpdateEditingProjectRequest) (_result *UpdateEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.UpdateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMediaProducingJobsWithOptions(request *ListMediaProducingJobsRequest, runtime *util.RuntimeOptions) (_result *ListMediaProducingJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMediaProducingJobs"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMediaProducingJobs(request *ListMediaProducingJobsRequest) (_result *ListMediaProducingJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.ListMediaProducingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEditingProjectMaterialsWithOptions(request *GetEditingProjectMaterialsRequest, runtime *util.RuntimeOptions) (_result *GetEditingProjectMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEditingProjectMaterials"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEditingProjectMaterials(request *GetEditingProjectMaterialsRequest) (_result *GetEditingProjectMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.GetEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultStorageLocationWithOptions(runtime *util.RuntimeOptions) (_result *GetDefaultStorageLocationResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetDefaultStorageLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDefaultStorageLocation"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultStorageLocation() (_result *GetDefaultStorageLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultStorageLocationResponse{}
	_body, _err := client.GetDefaultStorageLocationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMediaInfosWithOptions(request *DeleteMediaInfosRequest, runtime *util.RuntimeOptions) (_result *DeleteMediaInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMediaInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMediaInfos(request *DeleteMediaInfosRequest) (_result *DeleteMediaInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.DeleteMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetEventCallbackWithOptions(request *SetEventCallbackRequest, runtime *util.RuntimeOptions) (_result *SetEventCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetEventCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetEventCallback"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetEventCallback(request *SetEventCallbackRequest) (_result *SetEventCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetEventCallbackResponse{}
	_body, _err := client.SetEventCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterMediaInfoWithOptions(request *RegisterMediaInfoRequest, runtime *util.RuntimeOptions) (_result *RegisterMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterMediaInfo(request *RegisterMediaInfoRequest) (_result *RegisterMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.RegisterMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEditingProjectWithOptions(request *CreateEditingProjectRequest, runtime *util.RuntimeOptions) (_result *CreateEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEditingProject(request *CreateEditingProjectRequest) (_result *CreateEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.CreateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetMediaInfosWithOptions(request *BatchGetMediaInfosRequest, runtime *util.RuntimeOptions) (_result *BatchGetMediaInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchGetMediaInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchGetMediaInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetMediaInfos(request *BatchGetMediaInfosRequest) (_result *BatchGetMediaInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetMediaInfosResponse{}
	_body, _err := client.BatchGetMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultStorageLocationWithOptions(request *SetDefaultStorageLocationRequest, runtime *util.RuntimeOptions) (_result *SetDefaultStorageLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetDefaultStorageLocationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetDefaultStorageLocation"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultStorageLocation(request *SetDefaultStorageLocationRequest) (_result *SetDefaultStorageLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultStorageLocationResponse{}
	_body, _err := client.SetDefaultStorageLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMediaInfoWithOptions(request *UpdateMediaInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMediaInfo(request *UpdateMediaInfoRequest) (_result *UpdateMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.UpdateMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaProducingJobWithOptions(request *GetMediaProducingJobRequest, runtime *util.RuntimeOptions) (_result *GetMediaProducingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaProducingJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaProducingJob(request *GetMediaProducingJobRequest) (_result *GetMediaProducingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.GetMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIceProductStatusWithOptions(runtime *util.RuntimeOptions) (_result *DescribeIceProductStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeIceProductStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIceProductStatus"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIceProductStatus() (_result *DescribeIceProductStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIceProductStatusResponse{}
	_body, _err := client.DescribeIceProductStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMediaBasicInfosWithOptions(request *ListMediaBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListMediaBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMediaBasicInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMediaBasicInfos(request *ListMediaBasicInfosRequest) (_result *ListMediaBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.ListMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSubtitleProduceJobWithOptions(request *SubmitSubtitleProduceJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSubtitleProduceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSubtitleProduceJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSubtitleProduceJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSubtitleProduceJob(request *SubmitSubtitleProduceJobRequest) (_result *SubmitSubtitleProduceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSubtitleProduceJobResponse{}
	_body, _err := client.SubmitSubtitleProduceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitKeyWordCutJobWithOptions(request *SubmitKeyWordCutJobRequest, runtime *util.RuntimeOptions) (_result *SubmitKeyWordCutJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SubmitKeyWordCutJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitKeyWordCutJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitKeyWordCutJob(request *SubmitKeyWordCutJobRequest) (_result *SubmitKeyWordCutJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitKeyWordCutJobResponse{}
	_body, _err := client.SubmitKeyWordCutJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddEditingProjectMaterialsWithOptions(request *AddEditingProjectMaterialsRequest, runtime *util.RuntimeOptions) (_result *AddEditingProjectMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEditingProjectMaterialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEditingProjectMaterials"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEditingProjectMaterials(request *AddEditingProjectMaterialsRequest) (_result *AddEditingProjectMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEditingProjectMaterialsResponse{}
	_body, _err := client.AddEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitASRJobWithOptions(request *SubmitASRJobRequest, runtime *util.RuntimeOptions) (_result *SubmitASRJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitASRJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitASRJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitASRJob(request *SubmitASRJobRequest) (_result *SubmitASRJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitASRJobResponse{}
	_body, _err := client.SubmitASRJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEditingProjectWithOptions(request *GetEditingProjectRequest, runtime *util.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEditingProject(request *GetEditingProjectRequest) (_result *GetEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.GetEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSysTemplatesWithOptions(request *ListSysTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListSysTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSysTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSysTemplates"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSysTemplates(request *ListSysTemplatesRequest) (_result *ListSysTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSysTemplatesResponse{}
	_body, _err := client.ListSysTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteTemplate"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitIRJobWithOptions(request *SubmitIRJobRequest, runtime *util.RuntimeOptions) (_result *SubmitIRJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitIRJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitIRJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitIRJob(request *SubmitIRJobRequest) (_result *SubmitIRJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitIRJobResponse{}
	_body, _err := client.SubmitIRJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEditingProjectMaterialsWithOptions(request *DeleteEditingProjectMaterialsRequest, runtime *util.RuntimeOptions) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteEditingProjectMaterialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEditingProjectMaterials"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEditingProjectMaterials(request *DeleteEditingProjectMaterialsRequest) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEditingProjectMaterialsResponse{}
	_body, _err := client.DeleteEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchEditingProjectWithOptions(request *SearchEditingProjectRequest, runtime *util.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchEditingProject"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchEditingProject(request *SearchEditingProjectRequest) (_result *SearchEditingProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.SearchEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEditingProjectsWithOptions(request *DeleteEditingProjectsRequest, runtime *util.RuntimeOptions) (_result *DeleteEditingProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEditingProjects"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEditingProjects(request *DeleteEditingProjectsRequest) (_result *DeleteEditingProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.DeleteEditingProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMediaInfoWithOptions(request *GetMediaInfoRequest, runtime *util.RuntimeOptions) (_result *GetMediaInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMediaInfo"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMediaInfo(request *GetMediaInfoRequest) (_result *GetMediaInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.GetMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSmartJobWithOptions(request *SubmitSmartJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSmartJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitSmartJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitSmartJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSmartJob(request *SubmitSmartJobRequest) (_result *SubmitSmartJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSmartJobResponse{}
	_body, _err := client.SubmitSmartJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDelogoJobWithOptions(request *SubmitDelogoJobRequest, runtime *util.RuntimeOptions) (_result *SubmitDelogoJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitDelogoJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitDelogoJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitDelogoJob(request *SubmitDelogoJobRequest) (_result *SubmitDelogoJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitDelogoJobResponse{}
	_body, _err := client.SubmitDelogoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTemplate"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAudioProduceJobWithOptions(request *SubmitAudioProduceJobRequest, runtime *util.RuntimeOptions) (_result *SubmitAudioProduceJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitAudioProduceJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitAudioProduceJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAudioProduceJob(request *SubmitAudioProduceJobRequest) (_result *SubmitAudioProduceJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAudioProduceJobResponse{}
	_body, _err := client.SubmitAudioProduceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitMediaProducingJobWithOptions(request *SubmitMediaProducingJobRequest, runtime *util.RuntimeOptions) (_result *SubmitMediaProducingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitMediaProducingJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitMediaProducingJob(request *SubmitMediaProducingJobRequest) (_result *SubmitMediaProducingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.SubmitMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSmartJobWithOptions(request *UpdateSmartJobRequest, runtime *util.RuntimeOptions) (_result *UpdateSmartJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSmartJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSmartJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSmartJob(request *UpdateSmartJobRequest) (_result *UpdateSmartJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSmartJobResponse{}
	_body, _err := client.UpdateSmartJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAllPublicMediaTagsWithOptions(request *ListAllPublicMediaTagsRequest, runtime *util.RuntimeOptions) (_result *ListAllPublicMediaTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAllPublicMediaTags"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAllPublicMediaTags(request *ListAllPublicMediaTagsRequest) (_result *ListAllPublicMediaTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.ListAllPublicMediaTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitMattingJobWithOptions(request *SubmitMattingJobRequest, runtime *util.RuntimeOptions) (_result *SubmitMattingJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitMattingJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitMattingJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitMattingJob(request *SubmitMattingJobRequest) (_result *SubmitMattingJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitMattingJobResponse{}
	_body, _err := client.SubmitMattingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEventCallbackWithOptions(runtime *util.RuntimeOptions) (_result *GetEventCallbackResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetEventCallbackResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetEventCallback"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEventCallback() (_result *GetEventCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEventCallbackResponse{}
	_body, _err := client.GetEventCallbackWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublicMediaBasicInfosWithOptions(request *ListPublicMediaBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPublicMediaBasicInfos"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublicMediaBasicInfos(request *ListPublicMediaBasicInfosRequest) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.ListPublicMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCoverJobWithOptions(request *SubmitCoverJobRequest, runtime *util.RuntimeOptions) (_result *SubmitCoverJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitCoverJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitCoverJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCoverJob(request *SubmitCoverJobRequest) (_result *SubmitCoverJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCoverJobResponse{}
	_body, _err := client.SubmitCoverJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSmartHandleJobWithOptions(request *GetSmartHandleJobRequest, runtime *util.RuntimeOptions) (_result *GetSmartHandleJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetSmartHandleJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSmartHandleJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSmartHandleJob(request *GetSmartHandleJobRequest) (_result *GetSmartHandleJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSmartHandleJobResponse{}
	_body, _err := client.GetSmartHandleJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitH2VJobWithOptions(request *SubmitH2VJobRequest, runtime *util.RuntimeOptions) (_result *SubmitH2VJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitH2VJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitH2VJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitH2VJob(request *SubmitH2VJobRequest) (_result *SubmitH2VJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitH2VJobResponse{}
	_body, _err := client.SubmitH2VJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitPPTCutJobWithOptions(request *SubmitPPTCutJobRequest, runtime *util.RuntimeOptions) (_result *SubmitPPTCutJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SubmitPPTCutJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitPPTCutJob"), tea.String("2020-11-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitPPTCutJob(request *SubmitPPTCutJobRequest) (_result *SubmitPPTCutJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitPPTCutJobResponse{}
	_body, _err := client.SubmitPPTCutJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
