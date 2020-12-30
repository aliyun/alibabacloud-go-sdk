// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateProjectRequest struct {
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ModelId     *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetModelId(v string) *CreateProjectRequest {
	s.ModelId = &v
	return s
}

func (s *CreateProjectRequest) SetProjectType(v string) *CreateProjectRequest {
	s.ProjectType = &v
	return s
}

type CreateProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeployServiceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeployServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceRequest) SetProjectId(v string) *DeployServiceRequest {
	s.ProjectId = &v
	return s
}

type DeployServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceResponseBody) SetRequestId(v string) *DeployServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeployServiceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceResponse) SetHeaders(v map[string]*string) *DeployServiceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceResponse) SetBody(v *DeployServiceResponseBody) *DeployServiceResponse {
	s.Body = v
	return s
}

type DescribeProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectRequest) SetProjectId(v string) *DescribeProjectRequest {
	s.ProjectId = &v
	return s
}

type DescribeProjectResponseBody struct {
	QuestionCount       *int32  `json:"QuestionCount,omitempty" xml:"QuestionCount,omitempty"`
	DeployAvailable     *string `json:"DeployAvailable,omitempty" xml:"DeployAvailable,omitempty"`
	ModelName           *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectName         *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectId           *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OnlineServiceStatus *string `json:"OnlineServiceStatus,omitempty" xml:"OnlineServiceStatus,omitempty"`
	DeployTime          *int64  `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ProjectType         *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	DataStatus          *string `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
	ModelId             *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	TestServiceStatus   *string `json:"TestServiceStatus,omitempty" xml:"TestServiceStatus,omitempty"`
}

func (s DescribeProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponseBody) SetQuestionCount(v int32) *DescribeProjectResponseBody {
	s.QuestionCount = &v
	return s
}

func (s *DescribeProjectResponseBody) SetDeployAvailable(v string) *DescribeProjectResponseBody {
	s.DeployAvailable = &v
	return s
}

func (s *DescribeProjectResponseBody) SetModelName(v string) *DescribeProjectResponseBody {
	s.ModelName = &v
	return s
}

func (s *DescribeProjectResponseBody) SetRequestId(v string) *DescribeProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectResponseBody) SetProjectName(v string) *DescribeProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeProjectResponseBody) SetCreateTime(v int64) *DescribeProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeProjectResponseBody) SetProjectId(v string) *DescribeProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectResponseBody) SetOnlineServiceStatus(v string) *DescribeProjectResponseBody {
	s.OnlineServiceStatus = &v
	return s
}

func (s *DescribeProjectResponseBody) SetDeployTime(v int64) *DescribeProjectResponseBody {
	s.DeployTime = &v
	return s
}

func (s *DescribeProjectResponseBody) SetProjectType(v string) *DescribeProjectResponseBody {
	s.ProjectType = &v
	return s
}

func (s *DescribeProjectResponseBody) SetDataStatus(v string) *DescribeProjectResponseBody {
	s.DataStatus = &v
	return s
}

func (s *DescribeProjectResponseBody) SetModelId(v string) *DescribeProjectResponseBody {
	s.ModelId = &v
	return s
}

func (s *DescribeProjectResponseBody) SetTestServiceStatus(v string) *DescribeProjectResponseBody {
	s.TestServiceStatus = &v
	return s
}

type DescribeProjectResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectResponse) SetHeaders(v map[string]*string) *DescribeProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectResponse) SetBody(v *DescribeProjectResponseBody) *DescribeProjectResponse {
	s.Body = v
	return s
}

type GetPredictResultRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Question  *string `json:"Question,omitempty" xml:"Question,omitempty"`
	TopK      *int32  `json:"TopK,omitempty" xml:"TopK,omitempty"`
	TraceTag  *string `json:"TraceTag,omitempty" xml:"TraceTag,omitempty"`
}

func (s GetPredictResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultRequest) GoString() string {
	return s.String()
}

func (s *GetPredictResultRequest) SetProjectId(v string) *GetPredictResultRequest {
	s.ProjectId = &v
	return s
}

func (s *GetPredictResultRequest) SetQuestion(v string) *GetPredictResultRequest {
	s.Question = &v
	return s
}

func (s *GetPredictResultRequest) SetTopK(v int32) *GetPredictResultRequest {
	s.TopK = &v
	return s
}

func (s *GetPredictResultRequest) SetTraceTag(v string) *GetPredictResultRequest {
	s.TraceTag = &v
	return s
}

type GetPredictResultResponseBody struct {
	Trace          *string                                       `json:"Trace,omitempty" xml:"Trace,omitempty"`
	CostTime       *int64                                        `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	TopK           *int32                                        `json:"TopK,omitempty" xml:"TopK,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceTag       *string                                       `json:"TraceTag,omitempty" xml:"TraceTag,omitempty"`
	ProjectId      *string                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Question       *string                                       `json:"Question,omitempty" xml:"Question,omitempty"`
	PredictResults []*GetPredictResultResponseBodyPredictResults `json:"PredictResults,omitempty" xml:"PredictResults,omitempty" type:"Repeated"`
}

func (s GetPredictResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponseBody) SetTrace(v string) *GetPredictResultResponseBody {
	s.Trace = &v
	return s
}

func (s *GetPredictResultResponseBody) SetCostTime(v int64) *GetPredictResultResponseBody {
	s.CostTime = &v
	return s
}

func (s *GetPredictResultResponseBody) SetTopK(v int32) *GetPredictResultResponseBody {
	s.TopK = &v
	return s
}

func (s *GetPredictResultResponseBody) SetRequestId(v string) *GetPredictResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPredictResultResponseBody) SetTraceTag(v string) *GetPredictResultResponseBody {
	s.TraceTag = &v
	return s
}

func (s *GetPredictResultResponseBody) SetProjectId(v string) *GetPredictResultResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetPredictResultResponseBody) SetQuestion(v string) *GetPredictResultResponseBody {
	s.Question = &v
	return s
}

func (s *GetPredictResultResponseBody) SetPredictResults(v []*GetPredictResultResponseBodyPredictResults) *GetPredictResultResponseBody {
	s.PredictResults = v
	return s
}

type GetPredictResultResponseBodyPredictResults struct {
	Answer     *string  `json:"Answer,omitempty" xml:"Answer,omitempty"`
	QuestionId *string  `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	Rank       *int32   `json:"Rank,omitempty" xml:"Rank,omitempty"`
	Score      *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Question   *string  `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s GetPredictResultResponseBodyPredictResults) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponseBodyPredictResults) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponseBodyPredictResults) SetAnswer(v string) *GetPredictResultResponseBodyPredictResults {
	s.Answer = &v
	return s
}

func (s *GetPredictResultResponseBodyPredictResults) SetQuestionId(v string) *GetPredictResultResponseBodyPredictResults {
	s.QuestionId = &v
	return s
}

func (s *GetPredictResultResponseBodyPredictResults) SetRank(v int32) *GetPredictResultResponseBodyPredictResults {
	s.Rank = &v
	return s
}

func (s *GetPredictResultResponseBodyPredictResults) SetScore(v float32) *GetPredictResultResponseBodyPredictResults {
	s.Score = &v
	return s
}

func (s *GetPredictResultResponseBodyPredictResults) SetQuestion(v string) *GetPredictResultResponseBodyPredictResults {
	s.Question = &v
	return s
}

type GetPredictResultResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPredictResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPredictResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPredictResultResponse) GoString() string {
	return s.String()
}

func (s *GetPredictResultResponse) SetHeaders(v map[string]*string) *GetPredictResultResponse {
	s.Headers = v
	return s
}

func (s *GetPredictResultResponse) SetBody(v *GetPredictResultResponseBody) *GetPredictResultResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	FilterParam *string `json:"FilterParam,omitempty" xml:"FilterParam,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetFilterParam(v string) *ListProjectsRequest {
	s.FilterParam = &v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetProjectType(v string) *ListProjectsRequest {
	s.ProjectType = &v
	return s
}

type ListProjectsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Projects   []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetTotalCount(v int32) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageSize(v int32) *ListProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageNumber(v int32) *ListProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

type ListProjectsResponseBodyProjects struct {
	DeployAvailable     *string `json:"DeployAvailable,omitempty" xml:"DeployAvailable,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectName         *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectId           *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	QuestionCount       *int32  `json:"QuestionCount,omitempty" xml:"QuestionCount,omitempty"`
	DeployTime          *int64  `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	ProjectType         *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	OnlineServiceStatus *string `json:"OnlineServiceStatus,omitempty" xml:"OnlineServiceStatus,omitempty"`
	TestServiceStatus   *string `json:"TestServiceStatus,omitempty" xml:"TestServiceStatus,omitempty"`
	ModelName           *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	DataStatus          *string `json:"DataStatus,omitempty" xml:"DataStatus,omitempty"`
	ModelId             *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) SetDeployAvailable(v string) *ListProjectsResponseBodyProjects {
	s.DeployAvailable = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetCreateTime(v int64) *ListProjectsResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectName(v string) *ListProjectsResponseBodyProjects {
	s.ProjectName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectId(v string) *ListProjectsResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetQuestionCount(v int32) *ListProjectsResponseBodyProjects {
	s.QuestionCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDeployTime(v int64) *ListProjectsResponseBodyProjects {
	s.DeployTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectType(v string) *ListProjectsResponseBodyProjects {
	s.ProjectType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineServiceStatus(v string) *ListProjectsResponseBodyProjects {
	s.OnlineServiceStatus = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetTestServiceStatus(v string) *ListProjectsResponseBodyProjects {
	s.TestServiceStatus = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModelName(v string) *ListProjectsResponseBodyProjects {
	s.ModelName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetDataStatus(v string) *ListProjectsResponseBodyProjects {
	s.DataStatus = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModelId(v string) *ListProjectsResponseBodyProjects {
	s.ModelId = &v
	return s
}

type ListProjectsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ModifiyProjectRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModelId     *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ModifiyProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifiyProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifiyProjectRequest) SetProjectId(v string) *ModifiyProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifiyProjectRequest) SetModelId(v string) *ModifiyProjectRequest {
	s.ModelId = &v
	return s
}

func (s *ModifiyProjectRequest) SetProjectName(v string) *ModifiyProjectRequest {
	s.ProjectName = &v
	return s
}

type ModifiyProjectResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ModifiyProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifiyProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifiyProjectResponseBody) SetRequestId(v string) *ModifiyProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifiyProjectResponseBody) SetProjectId(v string) *ModifiyProjectResponseBody {
	s.ProjectId = &v
	return s
}

type ModifiyProjectResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifiyProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifiyProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifiyProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifiyProjectResponse) SetHeaders(v map[string]*string) *ModifiyProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifiyProjectResponse) SetBody(v *ModifiyProjectResponseBody) *ModifiyProjectResponse {
	s.Body = v
	return s
}

type UploadDictionaryRequest struct {
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DictionaryFileUrl *string `json:"DictionaryFileUrl,omitempty" xml:"DictionaryFileUrl,omitempty"`
	DictionaryData    *string `json:"DictionaryData,omitempty" xml:"DictionaryData,omitempty"`
}

func (s UploadDictionaryRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDictionaryRequest) GoString() string {
	return s.String()
}

func (s *UploadDictionaryRequest) SetProjectId(v string) *UploadDictionaryRequest {
	s.ProjectId = &v
	return s
}

func (s *UploadDictionaryRequest) SetDictionaryFileUrl(v string) *UploadDictionaryRequest {
	s.DictionaryFileUrl = &v
	return s
}

func (s *UploadDictionaryRequest) SetDictionaryData(v string) *UploadDictionaryRequest {
	s.DictionaryData = &v
	return s
}

type UploadDictionaryResponseBody struct {
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileDataCount *int32  `json:"FileDataCount,omitempty" xml:"FileDataCount,omitempty"`
	JsonDataCount *int32  `json:"JsonDataCount,omitempty" xml:"JsonDataCount,omitempty"`
}

func (s UploadDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDictionaryResponseBody) SetTotalCount(v int32) *UploadDictionaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *UploadDictionaryResponseBody) SetRequestId(v string) *UploadDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDictionaryResponseBody) SetProjectId(v string) *UploadDictionaryResponseBody {
	s.ProjectId = &v
	return s
}

func (s *UploadDictionaryResponseBody) SetFileDataCount(v int32) *UploadDictionaryResponseBody {
	s.FileDataCount = &v
	return s
}

func (s *UploadDictionaryResponseBody) SetJsonDataCount(v int32) *UploadDictionaryResponseBody {
	s.JsonDataCount = &v
	return s
}

type UploadDictionaryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDictionaryResponse) GoString() string {
	return s.String()
}

func (s *UploadDictionaryResponse) SetHeaders(v map[string]*string) *UploadDictionaryResponse {
	s.Headers = v
	return s
}

func (s *UploadDictionaryResponse) SetBody(v *UploadDictionaryResponseBody) *UploadDictionaryResponse {
	s.Body = v
	return s
}

type UploadDocumentRequest struct {
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DocumentFileUrl *string `json:"DocumentFileUrl,omitempty" xml:"DocumentFileUrl,omitempty"`
	DocumentData    *string `json:"DocumentData,omitempty" xml:"DocumentData,omitempty"`
}

func (s UploadDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentRequest) SetProjectId(v string) *UploadDocumentRequest {
	s.ProjectId = &v
	return s
}

func (s *UploadDocumentRequest) SetDocumentFileUrl(v string) *UploadDocumentRequest {
	s.DocumentFileUrl = &v
	return s
}

func (s *UploadDocumentRequest) SetDocumentData(v string) *UploadDocumentRequest {
	s.DocumentData = &v
	return s
}

type UploadDocumentResponseBody struct {
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FileDataCount *int32  `json:"FileDataCount,omitempty" xml:"FileDataCount,omitempty"`
	JsonDataCount *int32  `json:"JsonDataCount,omitempty" xml:"JsonDataCount,omitempty"`
}

func (s UploadDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponseBody) SetTotalCount(v int32) *UploadDocumentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *UploadDocumentResponseBody) SetRequestId(v string) *UploadDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetProjectId(v string) *UploadDocumentResponseBody {
	s.ProjectId = &v
	return s
}

func (s *UploadDocumentResponseBody) SetFileDataCount(v int32) *UploadDocumentResponseBody {
	s.FileDataCount = &v
	return s
}

func (s *UploadDocumentResponseBody) SetJsonDataCount(v int32) *UploadDocumentResponseBody {
	s.JsonDataCount = &v
	return s
}

type UploadDocumentResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadDocumentResponse) GoString() string {
	return s.String()
}

func (s *UploadDocumentResponse) SetHeaders(v map[string]*string) *UploadDocumentResponse {
	s.Headers = v
	return s
}

func (s *UploadDocumentResponse) SetBody(v *UploadDocumentResponseBody) *UploadDocumentResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("iqa"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployServiceWithOptions(request *DeployServiceRequest, runtime *util.RuntimeOptions) (_result *DeployServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployService"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployService(request *DeployServiceRequest) (_result *DeployServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployServiceResponse{}
	_body, _err := client.DeployServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectWithOptions(request *DescribeProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeProject"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProject(request *DescribeProjectRequest) (_result *DescribeProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectResponse{}
	_body, _err := client.DescribeProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPredictResultWithOptions(request *GetPredictResultRequest, runtime *util.RuntimeOptions) (_result *GetPredictResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPredictResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPredictResult"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPredictResult(request *GetPredictResultRequest) (_result *GetPredictResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPredictResultResponse{}
	_body, _err := client.GetPredictResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjects"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifiyProjectWithOptions(request *ModifiyProjectRequest, runtime *util.RuntimeOptions) (_result *ModifiyProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifiyProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifiyProject"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifiyProject(request *ModifiyProjectRequest) (_result *ModifiyProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifiyProjectResponse{}
	_body, _err := client.ModifiyProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDictionaryWithOptions(request *UploadDictionaryRequest, runtime *util.RuntimeOptions) (_result *UploadDictionaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadDictionaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadDictionary"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadDictionary(request *UploadDictionaryRequest) (_result *UploadDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDictionaryResponse{}
	_body, _err := client.UploadDictionaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadDocumentWithOptions(request *UploadDocumentRequest, runtime *util.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadDocumentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadDocument"), tea.String("2019-08-13"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadDocument(request *UploadDocumentRequest) (_result *UploadDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadDocumentResponse{}
	_body, _err := client.UploadDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
