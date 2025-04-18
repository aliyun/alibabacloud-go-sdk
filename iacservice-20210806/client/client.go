// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type JobStatusDetailValue struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// {}
	JobResult *string `json:"jobResult,omitempty" xml:"jobResult,omitempty"`
	// example:
	//
	// 2022-06-13 17:11:34
	TimeStamps *string `json:"timeStamps,omitempty" xml:"timeStamps,omitempty"`
}

func (s JobStatusDetailValue) String() string {
	return tea.Prettify(s)
}

func (s JobStatusDetailValue) GoString() string {
	return s.String()
}

func (s *JobStatusDetailValue) SetComment(v string) *JobStatusDetailValue {
	s.Comment = &v
	return s
}

func (s *JobStatusDetailValue) SetJobResult(v string) *JobStatusDetailValue {
	s.JobResult = &v
	return s
}

func (s *JobStatusDetailValue) SetTimeStamps(v string) *JobStatusDetailValue {
	s.TimeStamps = &v
	return s
}

type JobsStatusDetailValue struct {
	// example:
	//
	// ""
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// ""
	JobResult *string `json:"jobResult,omitempty" xml:"jobResult,omitempty"`
	// example:
	//
	// 2022-06-13 17:11:34
	TimeStamps *string `json:"timeStamps,omitempty" xml:"timeStamps,omitempty"`
}

func (s JobsStatusDetailValue) String() string {
	return tea.Prettify(s)
}

func (s JobsStatusDetailValue) GoString() string {
	return s.String()
}

func (s *JobsStatusDetailValue) SetComment(v string) *JobsStatusDetailValue {
	s.Comment = &v
	return s
}

func (s *JobsStatusDetailValue) SetJobResult(v string) *JobsStatusDetailValue {
	s.JobResult = &v
	return s
}

func (s *JobsStatusDetailValue) SetTimeStamps(v string) *JobsStatusDetailValue {
	s.TimeStamps = &v
	return s
}

type AssociateGroupRequest struct {
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// p-433aead7560571a87349d054b4
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AssociateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateGroupRequest) GoString() string {
	return s.String()
}

func (s *AssociateGroupRequest) SetClientToken(v string) *AssociateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateGroupRequest) SetProjectId(v string) *AssociateGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociateGroupRequest) SetResourceIds(v []*string) *AssociateGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *AssociateGroupRequest) SetResourceType(v string) *AssociateGroupRequest {
	s.ResourceType = &v
	return s
}

type AssociateGroupResponseBody struct {
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AssociateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateGroupResponseBody) SetRequestId(v string) *AssociateGroupResponseBody {
	s.RequestId = &v
	return s
}

type AssociateGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateGroupResponse) GoString() string {
	return s.String()
}

func (s *AssociateGroupResponse) SetHeaders(v map[string]*string) *AssociateGroupResponse {
	s.Headers = v
	return s
}

func (s *AssociateGroupResponse) SetStatusCode(v int32) *AssociateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateGroupResponse) SetBody(v *AssociateGroupResponseBody) *AssociateGroupResponse {
	s.Body = v
	return s
}

type AssociateParameterSetRequest struct {
	// This parameter is required.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// task-433aead756057ffdf5326bf1e12ed
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AssociateParameterSetRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetRequest) SetParameterSetIds(v []*string) *AssociateParameterSetRequest {
	s.ParameterSetIds = v
	return s
}

func (s *AssociateParameterSetRequest) SetResourceId(v string) *AssociateParameterSetRequest {
	s.ResourceId = &v
	return s
}

func (s *AssociateParameterSetRequest) SetResourceType(v string) *AssociateParameterSetRequest {
	s.ResourceType = &v
	return s
}

type AssociateParameterSetResponseBody struct {
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AssociateParameterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetResponseBody) SetRequestId(v string) *AssociateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

type AssociateParameterSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateParameterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetResponse) SetHeaders(v map[string]*string) *AssociateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *AssociateParameterSetResponse) SetStatusCode(v int32) *AssociateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateParameterSetResponse) SetBody(v *AssociateParameterSetResponseBody) *AssociateParameterSetResponse {
	s.Body = v
	return s
}

type AttachRabbitmqPublisherRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// task-518855d9a058cfffcbeafaebe6c89
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s AttachRabbitmqPublisherRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachRabbitmqPublisherRequest) GoString() string {
	return s.String()
}

func (s *AttachRabbitmqPublisherRequest) SetTaskId(v string) *AttachRabbitmqPublisherRequest {
	s.TaskId = &v
	return s
}

type AttachRabbitmqPublisherResponseBody struct {
	// example:
	//
	// D1DEAA38-D888-5811-A8A6-E1E2FBC0779C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AttachRabbitmqPublisherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachRabbitmqPublisherResponseBody) GoString() string {
	return s.String()
}

func (s *AttachRabbitmqPublisherResponseBody) SetRequestId(v string) *AttachRabbitmqPublisherResponseBody {
	s.RequestId = &v
	return s
}

type AttachRabbitmqPublisherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachRabbitmqPublisherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachRabbitmqPublisherResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachRabbitmqPublisherResponse) GoString() string {
	return s.String()
}

func (s *AttachRabbitmqPublisherResponse) SetHeaders(v map[string]*string) *AttachRabbitmqPublisherResponse {
	s.Headers = v
	return s
}

func (s *AttachRabbitmqPublisherResponse) SetStatusCode(v int32) *AttachRabbitmqPublisherResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachRabbitmqPublisherResponse) SetBody(v *AttachRabbitmqPublisherResponseBody) *AttachRabbitmqPublisherResponse {
	s.Body = v
	return s
}

type CancelProjectBuildResponseBody struct {
	// example:
	//
	// C7070EC3-DF66-58BA-A1DD-A8574FF53143
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelProjectBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelProjectBuildResponseBody) GoString() string {
	return s.String()
}

func (s *CancelProjectBuildResponseBody) SetRequestId(v string) *CancelProjectBuildResponseBody {
	s.RequestId = &v
	return s
}

type CancelProjectBuildResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelProjectBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelProjectBuildResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelProjectBuildResponse) GoString() string {
	return s.String()
}

func (s *CancelProjectBuildResponse) SetHeaders(v map[string]*string) *CancelProjectBuildResponse {
	s.Headers = v
	return s
}

func (s *CancelProjectBuildResponse) SetStatusCode(v int32) *CancelProjectBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelProjectBuildResponse) SetBody(v *CancelProjectBuildResponseBody) *CancelProjectBuildResponse {
	s.Body = v
	return s
}

type CancelRamPolicyExportTaskResponseBody struct {
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelRamPolicyExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRamPolicyExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRamPolicyExportTaskResponseBody) SetRequestId(v string) *CancelRamPolicyExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelRamPolicyExportTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRamPolicyExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRamPolicyExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRamPolicyExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelRamPolicyExportTaskResponse) SetHeaders(v map[string]*string) *CancelRamPolicyExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelRamPolicyExportTaskResponse) SetStatusCode(v int32) *CancelRamPolicyExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRamPolicyExportTaskResponse) SetBody(v *CancelRamPolicyExportTaskResponseBody) *CancelRamPolicyExportTaskResponse {
	s.Body = v
	return s
}

type CancelResourceExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// {}
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
}

func (s CancelResourceExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskRequest) SetClientToken(v string) *CancelResourceExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelResourceExportTaskRequest) SetRamRole(v string) *CancelResourceExportTaskRequest {
	s.RamRole = &v
	return s
}

type CancelResourceExportTaskResponseBody struct {
	// example:
	//
	// ex-3b6cb9fa4751a6e5cdc6460282
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CancelResourceExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskResponseBody) SetExportTaskId(v string) *CancelResourceExportTaskResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *CancelResourceExportTaskResponseBody) SetExportVersion(v string) *CancelResourceExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *CancelResourceExportTaskResponseBody) SetRequestId(v string) *CancelResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelResourceExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelResourceExportTaskResponse) SetHeaders(v map[string]*string) *CancelResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelResourceExportTaskResponse) SetStatusCode(v int32) *CancelResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelResourceExportTaskResponse) SetBody(v *CancelResourceExportTaskResponseBody) *CancelResourceExportTaskResponse {
	s.Body = v
	return s
}

type CheckResourceNameRequest struct {
	// example:
	//
	// mod-433af53ab7cc71afa
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-663a72b75245e3
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CheckResourceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceNameRequest) GoString() string {
	return s.String()
}

func (s *CheckResourceNameRequest) SetId(v string) *CheckResourceNameRequest {
	s.Id = &v
	return s
}

func (s *CheckResourceNameRequest) SetName(v string) *CheckResourceNameRequest {
	s.Name = &v
	return s
}

func (s *CheckResourceNameRequest) SetParentId(v string) *CheckResourceNameRequest {
	s.ParentId = &v
	return s
}

func (s *CheckResourceNameRequest) SetResourceType(v string) *CheckResourceNameRequest {
	s.ResourceType = &v
	return s
}

type CheckResourceNameResponseBody struct {
	// example:
	//
	// 24B9B06B-16D8-5558-9FE1-C29757EB1689
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CheckResourceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceNameResponseBody) SetRequestId(v string) *CheckResourceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceNameResponseBody) SetResult(v bool) *CheckResourceNameResponseBody {
	s.Result = &v
	return s
}

type CheckResourceNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResourceNameResponse) GoString() string {
	return s.String()
}

func (s *CheckResourceNameResponse) SetHeaders(v map[string]*string) *CheckResourceNameResponse {
	s.Headers = v
	return s
}

func (s *CheckResourceNameResponse) SetStatusCode(v int32) *CheckResourceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourceNameResponse) SetBody(v *CheckResourceNameResponseBody) *CheckResourceNameResponse {
	s.Body = v
	return s
}

type CloneGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-433aead7560571a87349d054b4
	ProjectId     *string   `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
}

func (s CloneGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneGroupRequest) GoString() string {
	return s.String()
}

func (s *CloneGroupRequest) SetClientToken(v string) *CloneGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CloneGroupRequest) SetDescription(v string) *CloneGroupRequest {
	s.Description = &v
	return s
}

func (s *CloneGroupRequest) SetName(v string) *CloneGroupRequest {
	s.Name = &v
	return s
}

func (s *CloneGroupRequest) SetProjectId(v string) *CloneGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CloneGroupRequest) SetResourceTypes(v []*string) *CloneGroupRequest {
	s.ResourceTypes = v
	return s
}

type CloneGroupResponseBody struct {
	// example:
	//
	// BF75EF50-955D-5E1F-AB23-A657C2C6D3C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloneGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CloneGroupResponseBody) SetRequestId(v string) *CloneGroupResponseBody {
	s.RequestId = &v
	return s
}

type CloneGroupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneGroupResponse) GoString() string {
	return s.String()
}

func (s *CloneGroupResponse) SetHeaders(v map[string]*string) *CloneGroupResponse {
	s.Headers = v
	return s
}

func (s *CloneGroupResponse) SetStatusCode(v int32) *CloneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneGroupResponse) SetBody(v *CloneGroupResponseBody) *CloneGroupResponse {
	s.Body = v
	return s
}

type CloneModuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ok
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e7853433574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// market
	ModuleSource *string `json:"moduleSource,omitempty" xml:"moduleSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CloneModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneModuleRequest) GoString() string {
	return s.String()
}

func (s *CloneModuleRequest) SetDescription(v string) *CloneModuleRequest {
	s.Description = &v
	return s
}

func (s *CloneModuleRequest) SetModuleId(v string) *CloneModuleRequest {
	s.ModuleId = &v
	return s
}

func (s *CloneModuleRequest) SetModuleSource(v string) *CloneModuleRequest {
	s.ModuleSource = &v
	return s
}

func (s *CloneModuleRequest) SetModuleVersion(v string) *CloneModuleRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CloneModuleRequest) SetName(v string) *CloneModuleRequest {
	s.Name = &v
	return s
}

type CloneModuleResponseBody struct {
	// example:
	//
	// mod-71012916a5a376f265ebc3774
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 78EC2EFB-AED9-5C23-AE98-6DDC6F2F96D6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CloneModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CloneModuleResponseBody) SetModuleId(v string) *CloneModuleResponseBody {
	s.ModuleId = &v
	return s
}

func (s *CloneModuleResponseBody) SetRequestId(v string) *CloneModuleResponseBody {
	s.RequestId = &v
	return s
}

type CloneModuleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneModuleResponse) GoString() string {
	return s.String()
}

func (s *CloneModuleResponse) SetHeaders(v map[string]*string) *CloneModuleResponse {
	s.Headers = v
	return s
}

func (s *CloneModuleResponse) SetStatusCode(v int32) *CloneModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneModuleResponse) SetBody(v *CloneModuleResponseBody) *CloneModuleResponse {
	s.Body = v
	return s
}

type CreateAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	DueTime     *string `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e7853433574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5a220da4-f594-4776-87ed-f37888ec0473
	Uid *int64 `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRequest) SetClientToken(v string) *CreateAuthorizationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAuthorizationRequest) SetDueTime(v string) *CreateAuthorizationRequest {
	s.DueTime = &v
	return s
}

func (s *CreateAuthorizationRequest) SetModuleId(v string) *CreateAuthorizationRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateAuthorizationRequest) SetModuleVersion(v string) *CreateAuthorizationRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateAuthorizationRequest) SetName(v string) *CreateAuthorizationRequest {
	s.Name = &v
	return s
}

func (s *CreateAuthorizationRequest) SetUid(v int64) *CreateAuthorizationRequest {
	s.Uid = &v
	return s
}

type CreateAuthorizationResponseBody struct {
	// example:
	//
	// auth-14e80de4866bf7171264a9b40661
	AuthorizationId *string `json:"authorizationId,omitempty" xml:"authorizationId,omitempty"`
	// example:
	//
	// 5FE84246-BB17-54BF-9F7A-F496270AC6DA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResponseBody) SetAuthorizationId(v string) *CreateAuthorizationResponseBody {
	s.AuthorizationId = &v
	return s
}

func (s *CreateAuthorizationResponseBody) SetRequestId(v string) *CreateAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type CreateAuthorizationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationResponse) SetHeaders(v map[string]*string) *CreateAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *CreateAuthorizationResponse) SetStatusCode(v int32) *CreateAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAuthorizationResponse) SetBody(v *CreateAuthorizationResponseBody) *CreateAuthorizationResponse {
	s.Body = v
	return s
}

type CreateExplorerTaskRequest struct {
	InitModuleState  *bool   `json:"InitModuleState,omitempty" xml:"InitModuleState,omitempty"`
	TerraformVersion *string `json:"TerraformVersion,omitempty" xml:"TerraformVersion,omitempty"`
	// This parameter is required.
	AutoApply   *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// This parameter is required.
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	Name                     *string `json:"name,omitempty" xml:"name,omitempty"`
	SkipPropertyValidation   *bool   `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TriggerValue             *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s CreateExplorerTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExplorerTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateExplorerTaskRequest) SetInitModuleState(v bool) *CreateExplorerTaskRequest {
	s.InitModuleState = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetTerraformVersion(v string) *CreateExplorerTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetAutoApply(v bool) *CreateExplorerTaskRequest {
	s.AutoApply = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetAutoDestroy(v bool) *CreateExplorerTaskRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetClientToken(v string) *CreateExplorerTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetDescription(v string) *CreateExplorerTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetModuleId(v string) *CreateExplorerTaskRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetModuleVersion(v string) *CreateExplorerTaskRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetName(v string) *CreateExplorerTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetSkipPropertyValidation(v bool) *CreateExplorerTaskRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetTerraformProviderVersion(v string) *CreateExplorerTaskRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateExplorerTaskRequest) SetTriggerValue(v string) *CreateExplorerTaskRequest {
	s.TriggerValue = &v
	return s
}

type CreateExplorerTaskResponseBody struct {
	ExplorerTaskId *string `json:"explorerTaskId,omitempty" xml:"explorerTaskId,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateExplorerTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExplorerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExplorerTaskResponseBody) SetExplorerTaskId(v string) *CreateExplorerTaskResponseBody {
	s.ExplorerTaskId = &v
	return s
}

func (s *CreateExplorerTaskResponseBody) SetRequestId(v string) *CreateExplorerTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateExplorerTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExplorerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExplorerTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExplorerTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateExplorerTaskResponse) SetHeaders(v map[string]*string) *CreateExplorerTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateExplorerTaskResponse) SetStatusCode(v int32) *CreateExplorerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExplorerTaskResponse) SetBody(v *CreateExplorerTaskResponseBody) *CreateExplorerTaskResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	ForcedSetting *bool `json:"forcedSetting,omitempty" xml:"forcedSetting,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name                 *string                           `json:"name,omitempty" xml:"name,omitempty"`
	NotifyConfig         []*CreateGroupRequestNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	NotifyOperationTypes []*string                         `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// p-433aead7560571a87349d054b4
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// ramName
	RamRole           *string   `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/test/test
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// example:
	//
	// 1.189.0
	TerraformProviderVersion *string                            `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TriggerConfig            []*CreateGroupRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	TriggerResourceType      []*string                          `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetAutoDestroy(v bool) *CreateGroupRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateGroupRequest) SetAutoTrigger(v bool) *CreateGroupRequest {
	s.AutoTrigger = &v
	return s
}

func (s *CreateGroupRequest) SetClientToken(v string) *CreateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetForcedSetting(v bool) *CreateGroupRequest {
	s.ForcedSetting = &v
	return s
}

func (s *CreateGroupRequest) SetName(v string) *CreateGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateGroupRequest) SetNotifyConfig(v []*CreateGroupRequestNotifyConfig) *CreateGroupRequest {
	s.NotifyConfig = v
	return s
}

func (s *CreateGroupRequest) SetNotifyOperationTypes(v []*string) *CreateGroupRequest {
	s.NotifyOperationTypes = v
	return s
}

func (s *CreateGroupRequest) SetProjectId(v string) *CreateGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateGroupRequest) SetRamRole(v string) *CreateGroupRequest {
	s.RamRole = &v
	return s
}

func (s *CreateGroupRequest) SetReportExportField(v []*string) *CreateGroupRequest {
	s.ReportExportField = v
	return s
}

func (s *CreateGroupRequest) SetReportExportPath(v string) *CreateGroupRequest {
	s.ReportExportPath = &v
	return s
}

func (s *CreateGroupRequest) SetTerraformProviderVersion(v string) *CreateGroupRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateGroupRequest) SetTriggerConfig(v []*CreateGroupRequestTriggerConfig) *CreateGroupRequest {
	s.TriggerConfig = v
	return s
}

func (s *CreateGroupRequest) SetTriggerResourceType(v []*string) *CreateGroupRequest {
	s.TriggerResourceType = v
	return s
}

type CreateGroupRequestNotifyConfig struct {
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s CreateGroupRequestNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequestNotifyConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestNotifyConfig) SetNotifyPath(v string) *CreateGroupRequestNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *CreateGroupRequestNotifyConfig) SetNotifyType(v string) *CreateGroupRequestNotifyConfig {
	s.NotifyType = &v
	return s
}

type CreateGroupRequestTriggerConfig struct {
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// 0 0 19 	- 	- ？
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s CreateGroupRequestTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *CreateGroupRequestTriggerConfig) SetTriggerStrategy(v string) *CreateGroupRequestTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateGroupRequestTriggerConfig) SetTriggerValue(v string) *CreateGroupRequestTriggerConfig {
	s.TriggerValue = &v
	return s
}

type CreateGroupResponseBody struct {
	// example:
	//
	// g-4267dcfbf1b6d128c87adf0e95f
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2daf4227f747cbf11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// null
	SubCommand *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
	TaskType   *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetClientToken(v string) *CreateJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateJobRequest) SetDescription(v string) *CreateJobRequest {
	s.Description = &v
	return s
}

func (s *CreateJobRequest) SetExecuteType(v string) *CreateJobRequest {
	s.ExecuteType = &v
	return s
}

func (s *CreateJobRequest) SetSubCommand(v string) *CreateJobRequest {
	s.SubCommand = &v
	return s
}

func (s *CreateJobRequest) SetTaskType(v string) *CreateJobRequest {
	s.TaskType = &v
	return s
}

type CreateJobResponseBody struct {
	// example:
	//
	// job-518855d9a058c32798c319561f
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateModuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string                       `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *CreateModuleRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// OSS：
	//
	// "oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip"
	//
	// Registry：
	//
	// "alibaba/security-group/alicloud:2.1.0"
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s CreateModuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleRequest) SetClientToken(v string) *CreateModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModuleRequest) SetDescription(v string) *CreateModuleRequest {
	s.Description = &v
	return s
}

func (s *CreateModuleRequest) SetGroupInfo(v *CreateModuleRequestGroupInfo) *CreateModuleRequest {
	s.GroupInfo = v
	return s
}

func (s *CreateModuleRequest) SetName(v string) *CreateModuleRequest {
	s.Name = &v
	return s
}

func (s *CreateModuleRequest) SetSource(v string) *CreateModuleRequest {
	s.Source = &v
	return s
}

func (s *CreateModuleRequest) SetSourcePath(v string) *CreateModuleRequest {
	s.SourcePath = &v
	return s
}

func (s *CreateModuleRequest) SetStatePath(v string) *CreateModuleRequest {
	s.StatePath = &v
	return s
}

func (s *CreateModuleRequest) SetVersionStrategy(v string) *CreateModuleRequest {
	s.VersionStrategy = &v
	return s
}

type CreateModuleRequestGroupInfo struct {
	// example:
	//
	// g-5fd38c9b92d541a7083a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// p-433aead75605713865c386cb9d
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateModuleRequestGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateModuleRequestGroupInfo) SetGroupId(v string) *CreateModuleRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *CreateModuleRequestGroupInfo) SetProjectId(v string) *CreateModuleRequestGroupInfo {
	s.ProjectId = &v
	return s
}

type CreateModuleResponseBody struct {
	// example:
	//
	// mod-518855d9a058cfffcc446d8fe3c99
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-92C7799512E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleResponseBody) SetModuleId(v string) *CreateModuleResponseBody {
	s.ModuleId = &v
	return s
}

func (s *CreateModuleResponseBody) SetRequestId(v string) *CreateModuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleResponse) SetHeaders(v map[string]*string) *CreateModuleResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleResponse) SetStatusCode(v int32) *CreateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModuleResponse) SetBody(v *CreateModuleResponseBody) *CreateModuleResponse {
	s.Body = v
	return s
}

type CreateModuleVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateModuleVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionRequest) SetClientToken(v string) *CreateModuleVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateModuleVersionRequest) SetDescription(v string) *CreateModuleVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateModuleVersionRequest) SetName(v string) *CreateModuleVersionRequest {
	s.Name = &v
	return s
}

type CreateModuleVersionResponseBody struct {
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateModuleVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionResponseBody) SetModuleVersion(v string) *CreateModuleVersionResponseBody {
	s.ModuleVersion = &v
	return s
}

func (s *CreateModuleVersionResponseBody) SetRequestId(v string) *CreateModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateModuleVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionResponse) SetHeaders(v map[string]*string) *CreateModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleVersionResponse) SetStatusCode(v int32) *CreateModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModuleVersionResponse) SetBody(v *CreateModuleVersionResponseBody) *CreateModuleVersionResponse {
	s.Body = v
	return s
}

type CreateParameterSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name       *string                                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters []*CreateParameterSetRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
}

func (s CreateParameterSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterSetRequest) SetClientToken(v string) *CreateParameterSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterSetRequest) SetDescription(v string) *CreateParameterSetRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterSetRequest) SetName(v string) *CreateParameterSetRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterSetRequest) SetParameters(v []*CreateParameterSetRequestParameters) *CreateParameterSetRequest {
	s.Parameters = v
	return s
}

type CreateParameterSetRequestParameters struct {
	// example:
	//
	// test1121
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateParameterSetRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterSetRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateParameterSetRequestParameters) SetName(v string) *CreateParameterSetRequestParameters {
	s.Name = &v
	return s
}

func (s *CreateParameterSetRequestParameters) SetType(v string) *CreateParameterSetRequestParameters {
	s.Type = &v
	return s
}

func (s *CreateParameterSetRequestParameters) SetValue(v string) *CreateParameterSetRequestParameters {
	s.Value = &v
	return s
}

type CreateParameterSetResponseBody struct {
	// example:
	//
	// pts-3b6cb9fa4751afff89a4b73779e0d
	ParameterSetId *string `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateParameterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterSetResponseBody) SetParameterSetId(v string) *CreateParameterSetResponseBody {
	s.ParameterSetId = &v
	return s
}

func (s *CreateParameterSetResponseBody) SetRequestId(v string) *CreateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

type CreateParameterSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParameterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterSetResponse) SetHeaders(v map[string]*string) *CreateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterSetResponse) SetStatusCode(v int32) *CreateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParameterSetResponse) SetBody(v *CreateParameterSetResponseBody) *CreateParameterSetResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetClientToken(v string) *CreateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

type CreateProjectResponseBody struct {
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 3E49127A-BB65-5CCD-AB93-0EC0A43E5446
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetProjectId(v string) *CreateProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateProjectBuildRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// refresh
	ProjectBuildAction *string                                  `json:"projectBuildAction,omitempty" xml:"projectBuildAction,omitempty"`
	TaskIds            []*string                                `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
	TaskPolicies       []*CreateProjectBuildRequestTaskPolicies `json:"taskPolicies,omitempty" xml:"taskPolicies,omitempty" type:"Repeated"`
}

func (s CreateProjectBuildRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectBuildRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectBuildRequest) SetClientToken(v string) *CreateProjectBuildRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectBuildRequest) SetGroupId(v string) *CreateProjectBuildRequest {
	s.GroupId = &v
	return s
}

func (s *CreateProjectBuildRequest) SetProjectBuildAction(v string) *CreateProjectBuildRequest {
	s.ProjectBuildAction = &v
	return s
}

func (s *CreateProjectBuildRequest) SetTaskIds(v []*string) *CreateProjectBuildRequest {
	s.TaskIds = v
	return s
}

func (s *CreateProjectBuildRequest) SetTaskPolicies(v []*CreateProjectBuildRequestTaskPolicies) *CreateProjectBuildRequest {
	s.TaskPolicies = v
	return s
}

type CreateProjectBuildRequestTaskPolicies struct {
	// example:
	//
	// destroyAfterExecution
	DestroyAfterExecution *bool `json:"destroyAfterExecution,omitempty" xml:"destroyAfterExecution,omitempty"`
	// example:
	//
	// 30
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// task-60f24b4eb47f1135b7b14ddbdfd
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateProjectBuildRequestTaskPolicies) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectBuildRequestTaskPolicies) GoString() string {
	return s.String()
}

func (s *CreateProjectBuildRequestTaskPolicies) SetDestroyAfterExecution(v bool) *CreateProjectBuildRequestTaskPolicies {
	s.DestroyAfterExecution = &v
	return s
}

func (s *CreateProjectBuildRequestTaskPolicies) SetPriority(v int64) *CreateProjectBuildRequestTaskPolicies {
	s.Priority = &v
	return s
}

func (s *CreateProjectBuildRequestTaskPolicies) SetTaskId(v string) *CreateProjectBuildRequestTaskPolicies {
	s.TaskId = &v
	return s
}

type CreateProjectBuildResponseBody struct {
	// example:
	//
	// pb-433aead756057193ba8bb410945
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// EE295EEC-EA85-5899-B2D5-5FCA788AC3C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProjectBuildResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectBuildResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectBuildResponseBody) SetId(v string) *CreateProjectBuildResponseBody {
	s.Id = &v
	return s
}

func (s *CreateProjectBuildResponseBody) SetRequestId(v string) *CreateProjectBuildResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectBuildResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectBuildResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectBuildResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectBuildResponse) SetHeaders(v map[string]*string) *CreateProjectBuildResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectBuildResponse) SetStatusCode(v int32) *CreateProjectBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectBuildResponse) SetBody(v *CreateProjectBuildResponseBody) *CreateProjectBuildResponse {
	s.Body = v
	return s
}

type CreateRabbitmqPublisherRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// ok
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ExchangeTest
	ExchangeName *string `json:"exchangeName,omitempty" xml:"exchangeName,omitempty"`
	// example:
	//
	// TOPIC
	ExchangeType *string `json:"exchangeType,omitempty" xml:"exchangeType,omitempty"`
	// This parameter is required.
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MQ
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NDAxREVDQzI2MjA0OT****
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5672
	Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MjoxODgwNzcwODY5MD****
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Test
	VirtualHost *string `json:"virtualHost,omitempty" xml:"virtualHost,omitempty"`
}

func (s CreateRabbitmqPublisherRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRabbitmqPublisherRequest) GoString() string {
	return s.String()
}

func (s *CreateRabbitmqPublisherRequest) SetClientToken(v string) *CreateRabbitmqPublisherRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetDescription(v string) *CreateRabbitmqPublisherRequest {
	s.Description = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetExchangeName(v string) *CreateRabbitmqPublisherRequest {
	s.ExchangeName = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetExchangeType(v string) *CreateRabbitmqPublisherRequest {
	s.ExchangeType = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetHostName(v string) *CreateRabbitmqPublisherRequest {
	s.HostName = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetName(v string) *CreateRabbitmqPublisherRequest {
	s.Name = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetPassword(v string) *CreateRabbitmqPublisherRequest {
	s.Password = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetPort(v int64) *CreateRabbitmqPublisherRequest {
	s.Port = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetUserName(v string) *CreateRabbitmqPublisherRequest {
	s.UserName = &v
	return s
}

func (s *CreateRabbitmqPublisherRequest) SetVirtualHost(v string) *CreateRabbitmqPublisherRequest {
	s.VirtualHost = &v
	return s
}

type CreateRabbitmqPublisherResponseBody struct {
	// example:
	//
	// mqp-3b6cb9fa4751afffb0af06b9ba504
	PublisherId *string `json:"publisherId,omitempty" xml:"publisherId,omitempty"`
	// example:
	//
	// C3DA172D-DDCF-5268-BB0F-060C3A9D2623
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRabbitmqPublisherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRabbitmqPublisherResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRabbitmqPublisherResponseBody) SetPublisherId(v string) *CreateRabbitmqPublisherResponseBody {
	s.PublisherId = &v
	return s
}

func (s *CreateRabbitmqPublisherResponseBody) SetRequestId(v string) *CreateRabbitmqPublisherResponseBody {
	s.RequestId = &v
	return s
}

type CreateRabbitmqPublisherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRabbitmqPublisherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRabbitmqPublisherResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRabbitmqPublisherResponse) GoString() string {
	return s.String()
}

func (s *CreateRabbitmqPublisherResponse) SetHeaders(v map[string]*string) *CreateRabbitmqPublisherResponse {
	s.Headers = v
	return s
}

func (s *CreateRabbitmqPublisherResponse) SetStatusCode(v int32) *CreateRabbitmqPublisherResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRabbitmqPublisherResponse) SetBody(v *CreateRabbitmqPublisherResponseBody) *CreateRabbitmqPublisherResponse {
	s.Body = v
	return s
}

type CreateRamPolicyExportTaskRequest struct {
	AuthorizationAccountIds []*int64  `json:"authorizationAccountIds,omitempty" xml:"authorizationAccountIds,omitempty" type:"Repeated"`
	AuthorizationActions    []*string `json:"authorizationActions,omitempty" xml:"authorizationActions,omitempty" type:"Repeated"`
	AuthorizationRegionIds  []*string `json:"authorizationRegionIds,omitempty" xml:"authorizationRegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2da11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e785fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ramName
	RamRole                  *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s CreateRamPolicyExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRamPolicyExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRamPolicyExportTaskRequest) SetAuthorizationAccountIds(v []*int64) *CreateRamPolicyExportTaskRequest {
	s.AuthorizationAccountIds = v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetAuthorizationActions(v []*string) *CreateRamPolicyExportTaskRequest {
	s.AuthorizationActions = v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetAuthorizationRegionIds(v []*string) *CreateRamPolicyExportTaskRequest {
	s.AuthorizationRegionIds = v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetClientToken(v string) *CreateRamPolicyExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetModuleId(v string) *CreateRamPolicyExportTaskRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetModuleVersion(v string) *CreateRamPolicyExportTaskRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetName(v string) *CreateRamPolicyExportTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetRamRole(v string) *CreateRamPolicyExportTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetTerraformProviderVersion(v string) *CreateRamPolicyExportTaskRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateRamPolicyExportTaskRequest) SetTriggerStrategy(v string) *CreateRamPolicyExportTaskRequest {
	s.TriggerStrategy = &v
	return s
}

type CreateRamPolicyExportTaskResponseBody struct {
	// example:
	//
	// rpe-436057ffe0252e48f9286a
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 9AC66641-C0F1-533B-AB58-982FC7B1C32A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRamPolicyExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRamPolicyExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRamPolicyExportTaskResponseBody) SetId(v string) *CreateRamPolicyExportTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateRamPolicyExportTaskResponseBody) SetRequestId(v string) *CreateRamPolicyExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateRamPolicyExportTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRamPolicyExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRamPolicyExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRamPolicyExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRamPolicyExportTaskResponse) SetHeaders(v map[string]*string) *CreateRamPolicyExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRamPolicyExportTaskResponse) SetStatusCode(v int32) *CreateRamPolicyExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRamPolicyExportTaskResponse) SetBody(v *CreateRamPolicyExportTaskResponseBody) *CreateRamPolicyExportTaskResponse {
	s.Body = v
	return s
}

type CreateResourceExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ConfigPath  *string `json:"configPath,omitempty" xml:"configPath,omitempty"`
	// example:
	//
	// OK
	Description    *string                                        `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeRules   []*CreateResourceExportTaskRequestExcludeRules `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	ExportToModule *CreateResourceExportTaskRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	IncludeRules   []*CreateResourceExportTaskRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {}
	RamRole                  *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Auto
	TriggerStrategy *string                                     `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables       []*CreateResourceExportTaskRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateResourceExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequest) SetClientToken(v string) *CreateResourceExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetConfigPath(v string) *CreateResourceExportTaskRequest {
	s.ConfigPath = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetDescription(v string) *CreateResourceExportTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetExcludeRules(v []*CreateResourceExportTaskRequestExcludeRules) *CreateResourceExportTaskRequest {
	s.ExcludeRules = v
	return s
}

func (s *CreateResourceExportTaskRequest) SetExportToModule(v *CreateResourceExportTaskRequestExportToModule) *CreateResourceExportTaskRequest {
	s.ExportToModule = v
	return s
}

func (s *CreateResourceExportTaskRequest) SetIncludeRules(v []*CreateResourceExportTaskRequestIncludeRules) *CreateResourceExportTaskRequest {
	s.IncludeRules = v
	return s
}

func (s *CreateResourceExportTaskRequest) SetName(v string) *CreateResourceExportTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetRamRole(v string) *CreateResourceExportTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTerraformProviderVersion(v string) *CreateResourceExportTaskRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTerraformVersion(v string) *CreateResourceExportTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetTriggerStrategy(v string) *CreateResourceExportTaskRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateResourceExportTaskRequest) SetVariables(v []*CreateResourceExportTaskRequestVariables) *CreateResourceExportTaskRequest {
	s.Variables = v
	return s
}

type CreateResourceExportTaskRequestExcludeRules struct {
	// example:
	//
	// VPC
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateResourceExportTaskRequestExcludeRules) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskRequestExcludeRules) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestExcludeRules) SetKey(v string) *CreateResourceExportTaskRequestExcludeRules {
	s.Key = &v
	return s
}

func (s *CreateResourceExportTaskRequestExcludeRules) SetValues(v []*string) *CreateResourceExportTaskRequestExcludeRules {
	s.Values = v
	return s
}

type CreateResourceExportTaskRequestExportToModule struct {
	// example:
	//
	// Registry
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// alibaba/security-group/alicloud
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// /
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s CreateResourceExportTaskRequestExportToModule) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskRequestExportToModule) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestExportToModule) SetSource(v string) *CreateResourceExportTaskRequestExportToModule {
	s.Source = &v
	return s
}

func (s *CreateResourceExportTaskRequestExportToModule) SetSourcePath(v string) *CreateResourceExportTaskRequestExportToModule {
	s.SourcePath = &v
	return s
}

func (s *CreateResourceExportTaskRequestExportToModule) SetStatePath(v string) *CreateResourceExportTaskRequestExportToModule {
	s.StatePath = &v
	return s
}

type CreateResourceExportTaskRequestIncludeRules struct {
	// example:
	//
	// ZoneId
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateResourceExportTaskRequestIncludeRules) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskRequestIncludeRules) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestIncludeRules) SetKey(v string) *CreateResourceExportTaskRequestIncludeRules {
	s.Key = &v
	return s
}

func (s *CreateResourceExportTaskRequestIncludeRules) SetValues(v []*string) *CreateResourceExportTaskRequestIncludeRules {
	s.Values = v
	return s
}

type CreateResourceExportTaskRequestVariables struct {
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// AliCloud::VPC::VPC
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateResourceExportTaskRequestVariables) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskRequestVariables) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskRequestVariables) SetProperties(v []*string) *CreateResourceExportTaskRequestVariables {
	s.Properties = v
	return s
}

func (s *CreateResourceExportTaskRequestVariables) SetResourceType(v string) *CreateResourceExportTaskRequestVariables {
	s.ResourceType = &v
	return s
}

type CreateResourceExportTaskResponseBody struct {
	// example:
	//
	// ex-4a1ec8b7003d24528326821be
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// CFD8C2A8-5BE7-56D2-916D-64039B8E06E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateResourceExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskResponseBody) SetExportTaskId(v string) *CreateResourceExportTaskResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *CreateResourceExportTaskResponseBody) SetExportVersion(v string) *CreateResourceExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *CreateResourceExportTaskResponseBody) SetRequestId(v string) *CreateResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceExportTaskResponse) SetHeaders(v map[string]*string) *CreateResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceExportTaskResponse) SetStatusCode(v int32) *CreateResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceExportTaskResponse) SetBody(v *CreateResourceExportTaskResponseBody) *CreateResourceExportTaskResponse {
	s.Body = v
	return s
}

type CreateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// demo
	Description     *string                     `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo       *CreateTaskRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	InitModuleState *bool                       `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e7853433574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name               *string            `json:"name,omitempty" xml:"name,omitempty"`
	Parameters         map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	ProtectionStrategy []*string          `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	RamRole                *string                       `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	SkipPropertyValidation *bool                         `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	TaskBackend            *CreateTaskRequestTaskBackend `json:"taskBackend,omitempty" xml:"taskBackend,omitempty" type:"Struct"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Always
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// ******
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) SetAutoApply(v bool) *CreateTaskRequest {
	s.AutoApply = &v
	return s
}

func (s *CreateTaskRequest) SetAutoDestroy(v bool) *CreateTaskRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateTaskRequest) SetClientToken(v string) *CreateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskRequest) SetDescription(v string) *CreateTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTaskRequest) SetGroupInfo(v *CreateTaskRequestGroupInfo) *CreateTaskRequest {
	s.GroupInfo = v
	return s
}

func (s *CreateTaskRequest) SetInitModuleState(v bool) *CreateTaskRequest {
	s.InitModuleState = &v
	return s
}

func (s *CreateTaskRequest) SetModuleId(v string) *CreateTaskRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateTaskRequest) SetModuleVersion(v string) *CreateTaskRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateTaskRequest) SetName(v string) *CreateTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTaskRequest) SetParameters(v map[string]*string) *CreateTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateTaskRequest) SetProtectionStrategy(v []*string) *CreateTaskRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *CreateTaskRequest) SetRamRole(v string) *CreateTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateTaskRequest) SetSkipPropertyValidation(v bool) *CreateTaskRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *CreateTaskRequest) SetTaskBackend(v *CreateTaskRequestTaskBackend) *CreateTaskRequest {
	s.TaskBackend = v
	return s
}

func (s *CreateTaskRequest) SetTerraformVersion(v string) *CreateTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateTaskRequest) SetTriggerStrategy(v string) *CreateTaskRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateTaskRequest) SetTriggerValue(v string) *CreateTaskRequest {
	s.TriggerValue = &v
	return s
}

type CreateTaskRequestGroupInfo struct {
	// example:
	//
	// g-5fd38c9b92d541a7083a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// project-433aead7560572057e5d9167608
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateTaskRequestGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestGroupInfo) SetGroupId(v string) *CreateTaskRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *CreateTaskRequestGroupInfo) SetProjectId(v string) *CreateTaskRequestGroupInfo {
	s.ProjectId = &v
	return s
}

type CreateTaskRequestTaskBackend struct {
	BucketEndpoint *string `json:"bucketEndpoint,omitempty" xml:"bucketEndpoint,omitempty"`
	BucketName     *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	ObjectPath     *string `json:"objectPath,omitempty" xml:"objectPath,omitempty"`
}

func (s CreateTaskRequestTaskBackend) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestTaskBackend) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTaskBackend) SetBucketEndpoint(v string) *CreateTaskRequestTaskBackend {
	s.BucketEndpoint = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetBucketName(v string) *CreateTaskRequestTaskBackend {
	s.BucketName = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetObjectPath(v string) *CreateTaskRequestTaskBackend {
	s.ObjectPath = &v
	return s
}

type CreateTaskResponseBody struct {
	// example:
	//
	// CD478792-6952-5A1C-9F57-78932BF0FAC6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetTaskId(v string) *CreateTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type DeleteAuthorizationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationResponseBody) SetRequestId(v string) *DeleteAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAuthorizationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationResponse) SetStatusCode(v int32) *DeleteAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationResponse) SetBody(v *DeleteAuthorizationResponseBody) *DeleteAuthorizationResponse {
	s.Body = v
	return s
}

type DeleteGroupResponseBody struct {
	// example:
	//
	// 1E7BA3EB-B0EF-53F5-9999-07CAD6D9F8A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteModuleResponseBody struct {
	// example:
	//
	// 49DA6457-E545-5095-A930-EB8F0BCD4DAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponseBody) SetRequestId(v string) *DeleteModuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteModuleResponse) SetHeaders(v map[string]*string) *DeleteModuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteModuleResponse) SetStatusCode(v int32) *DeleteModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModuleResponse) SetBody(v *DeleteModuleResponseBody) *DeleteModuleResponse {
	s.Body = v
	return s
}

type DeleteParameterSetResponseBody struct {
	// example:
	//
	// 708D670B-67C4-5653-9F88-8F7800429EE1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteParameterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterSetResponseBody) SetRequestId(v string) *DeleteParameterSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParameterSetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterSetResponse) SetHeaders(v map[string]*string) *DeleteParameterSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterSetResponse) SetStatusCode(v int32) *DeleteParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterSetResponse) SetBody(v *DeleteParameterSetResponseBody) *DeleteParameterSetResponse {
	s.Body = v
	return s
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type DeleteRabbitmqPublisherResponseBody struct {
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRabbitmqPublisherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRabbitmqPublisherResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRabbitmqPublisherResponseBody) SetRequestId(v string) *DeleteRabbitmqPublisherResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRabbitmqPublisherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRabbitmqPublisherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRabbitmqPublisherResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRabbitmqPublisherResponse) GoString() string {
	return s.String()
}

func (s *DeleteRabbitmqPublisherResponse) SetHeaders(v map[string]*string) *DeleteRabbitmqPublisherResponse {
	s.Headers = v
	return s
}

func (s *DeleteRabbitmqPublisherResponse) SetStatusCode(v int32) *DeleteRabbitmqPublisherResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRabbitmqPublisherResponse) SetBody(v *DeleteRabbitmqPublisherResponseBody) *DeleteRabbitmqPublisherResponse {
	s.Body = v
	return s
}

type DeleteRamPolicyExportTaskResponseBody struct {
	// example:
	//
	// 3E49127A-BB65-5CCD-AB93-0EC0A43E5446
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRamPolicyExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRamPolicyExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRamPolicyExportTaskResponseBody) SetRequestId(v string) *DeleteRamPolicyExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRamPolicyExportTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRamPolicyExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRamPolicyExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRamPolicyExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteRamPolicyExportTaskResponse) SetHeaders(v map[string]*string) *DeleteRamPolicyExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteRamPolicyExportTaskResponse) SetStatusCode(v int32) *DeleteRamPolicyExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRamPolicyExportTaskResponse) SetBody(v *DeleteRamPolicyExportTaskResponseBody) *DeleteRamPolicyExportTaskResponse {
	s.Body = v
	return s
}

type DeleteRamPolicyExportTaskVersionResponseBody struct {
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRamPolicyExportTaskVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRamPolicyExportTaskVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRamPolicyExportTaskVersionResponseBody) SetRequestId(v string) *DeleteRamPolicyExportTaskVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRamPolicyExportTaskVersionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRamPolicyExportTaskVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRamPolicyExportTaskVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRamPolicyExportTaskVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRamPolicyExportTaskVersionResponse) SetHeaders(v map[string]*string) *DeleteRamPolicyExportTaskVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRamPolicyExportTaskVersionResponse) SetStatusCode(v int32) *DeleteRamPolicyExportTaskVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRamPolicyExportTaskVersionResponse) SetBody(v *DeleteRamPolicyExportTaskVersionResponseBody) *DeleteRamPolicyExportTaskVersionResponse {
	s.Body = v
	return s
}

type DeleteResourceExportTaskResponseBody struct {
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteResourceExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceExportTaskResponseBody) SetRequestId(v string) *DeleteResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceExportTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceExportTaskResponse) SetHeaders(v map[string]*string) *DeleteResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceExportTaskResponse) SetStatusCode(v int32) *DeleteResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceExportTaskResponse) SetBody(v *DeleteResourceExportTaskResponseBody) *DeleteResourceExportTaskResponse {
	s.Body = v
	return s
}

type DeleteSceneTestingTaskResponseBody struct {
	// example:
	//
	// 17793D91-A26F-520D-A948-3452A45D15B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSceneTestingTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneTestingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneTestingTaskResponseBody) SetRequestId(v string) *DeleteSceneTestingTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSceneTestingTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSceneTestingTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSceneTestingTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSceneTestingTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneTestingTaskResponse) SetHeaders(v map[string]*string) *DeleteSceneTestingTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneTestingTaskResponse) SetStatusCode(v int32) *DeleteSceneTestingTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSceneTestingTaskResponse) SetBody(v *DeleteSceneTestingTaskResponseBody) *DeleteSceneTestingTaskResponse {
	s.Body = v
	return s
}

type DeleteTaskResponseBody struct {
	// example:
	//
	// 73B38F77-62BA-5878-8952-530DFE21C93B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponseBody) SetRequestId(v string) *DeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteTaskResponse) SetHeaders(v map[string]*string) *DeleteTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteTaskResponse) SetStatusCode(v int32) *DeleteTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTaskResponse) SetBody(v *DeleteTaskResponseBody) *DeleteTaskResponse {
	s.Body = v
	return s
}

type DetachRabbitmqPublisherRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// task-518855d9a058cf1127d082bec5
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s DetachRabbitmqPublisherRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachRabbitmqPublisherRequest) GoString() string {
	return s.String()
}

func (s *DetachRabbitmqPublisherRequest) SetTaskId(v string) *DetachRabbitmqPublisherRequest {
	s.TaskId = &v
	return s
}

type DetachRabbitmqPublisherResponseBody struct {
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetachRabbitmqPublisherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachRabbitmqPublisherResponseBody) GoString() string {
	return s.String()
}

func (s *DetachRabbitmqPublisherResponseBody) SetRequestId(v string) *DetachRabbitmqPublisherResponseBody {
	s.RequestId = &v
	return s
}

type DetachRabbitmqPublisherResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachRabbitmqPublisherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachRabbitmqPublisherResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachRabbitmqPublisherResponse) GoString() string {
	return s.String()
}

func (s *DetachRabbitmqPublisherResponse) SetHeaders(v map[string]*string) *DetachRabbitmqPublisherResponse {
	s.Headers = v
	return s
}

func (s *DetachRabbitmqPublisherResponse) SetStatusCode(v int32) *DetachRabbitmqPublisherResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachRabbitmqPublisherResponse) SetBody(v *DetachRabbitmqPublisherResponseBody) *DetachRabbitmqPublisherResponse {
	s.Body = v
	return s
}

type DissociateGroupRequest struct {
	// example:
	//
	// 2daf4227f747cbf11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DissociateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateGroupRequest) GoString() string {
	return s.String()
}

func (s *DissociateGroupRequest) SetClientToken(v string) *DissociateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateGroupRequest) SetResourceIds(v []*string) *DissociateGroupRequest {
	s.ResourceIds = v
	return s
}

func (s *DissociateGroupRequest) SetResourceType(v string) *DissociateGroupRequest {
	s.ResourceType = &v
	return s
}

type DissociateGroupResponseBody struct {
	// example:
	//
	// 17793D91-A26F-520D-A948-3452A45D15B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DissociateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateGroupResponseBody) SetRequestId(v string) *DissociateGroupResponseBody {
	s.RequestId = &v
	return s
}

type DissociateGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateGroupResponse) GoString() string {
	return s.String()
}

func (s *DissociateGroupResponse) SetHeaders(v map[string]*string) *DissociateGroupResponse {
	s.Headers = v
	return s
}

func (s *DissociateGroupResponse) SetStatusCode(v int32) *DissociateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateGroupResponse) SetBody(v *DissociateGroupResponseBody) *DissociateGroupResponse {
	s.Body = v
	return s
}

type DissociateParameterSetRequest struct {
	// This parameter is required.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// mod-39cd1e5e58c50e79dd8cd901cd
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DissociateParameterSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DissociateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetRequest) SetParameterSetIds(v []*string) *DissociateParameterSetRequest {
	s.ParameterSetIds = v
	return s
}

func (s *DissociateParameterSetRequest) SetResourceId(v string) *DissociateParameterSetRequest {
	s.ResourceId = &v
	return s
}

func (s *DissociateParameterSetRequest) SetResourceType(v string) *DissociateParameterSetRequest {
	s.ResourceType = &v
	return s
}

type DissociateParameterSetResponseBody struct {
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DissociateParameterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DissociateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetResponseBody) SetRequestId(v string) *DissociateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

type DissociateParameterSetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateParameterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DissociateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetResponse) SetHeaders(v map[string]*string) *DissociateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *DissociateParameterSetResponse) SetStatusCode(v int32) *DissociateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateParameterSetResponse) SetBody(v *DissociateParameterSetResponseBody) *DissociateParameterSetResponse {
	s.Body = v
	return s
}

type ExecuteRamPolicyExportTaskResponseBody struct {
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteRamPolicyExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRamPolicyExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteRamPolicyExportTaskResponseBody) SetExportVersion(v string) *ExecuteRamPolicyExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *ExecuteRamPolicyExportTaskResponseBody) SetRequestId(v string) *ExecuteRamPolicyExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteRamPolicyExportTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteRamPolicyExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteRamPolicyExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRamPolicyExportTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecuteRamPolicyExportTaskResponse) SetHeaders(v map[string]*string) *ExecuteRamPolicyExportTaskResponse {
	s.Headers = v
	return s
}

func (s *ExecuteRamPolicyExportTaskResponse) SetStatusCode(v int32) *ExecuteRamPolicyExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteRamPolicyExportTaskResponse) SetBody(v *ExecuteRamPolicyExportTaskResponseBody) *ExecuteRamPolicyExportTaskResponse {
	s.Body = v
	return s
}

type ExecuteResourceExportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// {}
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
}

func (s ExecuteResourceExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecuteResourceExportTaskRequest) SetClientToken(v string) *ExecuteResourceExportTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteResourceExportTaskRequest) SetRamRole(v string) *ExecuteResourceExportTaskRequest {
	s.RamRole = &v
	return s
}

type ExecuteResourceExportTaskResponseBody struct {
	// example:
	//
	// ex-3b6cb9fa4751a6e645ad8365e6
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// 0B0A7C19-9077-5975-ACBD-DEE718787992
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ExecuteResourceExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteResourceExportTaskResponseBody) SetExportTaskId(v string) *ExecuteResourceExportTaskResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *ExecuteResourceExportTaskResponseBody) SetExportVersion(v string) *ExecuteResourceExportTaskResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *ExecuteResourceExportTaskResponseBody) SetRequestId(v string) *ExecuteResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteResourceExportTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteResourceExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecuteResourceExportTaskResponse) SetHeaders(v map[string]*string) *ExecuteResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *ExecuteResourceExportTaskResponse) SetStatusCode(v int32) *ExecuteResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteResourceExportTaskResponse) SetBody(v *ExecuteResourceExportTaskResponseBody) *ExecuteResourceExportTaskResponse {
	s.Body = v
	return s
}

type GetExplorerTaskResponseBody struct {
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetExplorerTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetExplorerTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExplorerTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetExplorerTaskResponseBody) SetRequestId(v string) *GetExplorerTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExplorerTaskResponseBody) SetTask(v *GetExplorerTaskResponseBodyTask) *GetExplorerTaskResponseBody {
	s.Task = v
	return s
}

type GetExplorerTaskResponseBodyTask struct {
	AutoApply              *bool                                           `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	AutoDestroy            *bool                                           `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	CreateTime             *string                                         `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CurrentJobId           *string                                         `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	Description            *string                                         `json:"description,omitempty" xml:"description,omitempty"`
	ExplorerTaskId         *string                                         `json:"explorerTaskId,omitempty" xml:"explorerTaskId,omitempty"`
	InitModuleState        *bool                                           `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	ModuleId               *string                                         `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleVersion          *string                                         `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	Name                   *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	SkipPropertyValidation *bool                                           `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	Status                 *string                                         `json:"status,omitempty" xml:"status,omitempty"`
	TaskOutputPath         *string                                         `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	TerraformVersion       *string                                         `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	TriggerValue           *string                                         `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
	UploadSignature        *GetExplorerTaskResponseBodyTaskUploadSignature `json:"uploadSignature,omitempty" xml:"uploadSignature,omitempty" type:"Struct"`
}

func (s GetExplorerTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetExplorerTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetExplorerTaskResponseBodyTask) SetAutoApply(v bool) *GetExplorerTaskResponseBodyTask {
	s.AutoApply = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetAutoDestroy(v bool) *GetExplorerTaskResponseBodyTask {
	s.AutoDestroy = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetCreateTime(v string) *GetExplorerTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetCurrentJobId(v string) *GetExplorerTaskResponseBodyTask {
	s.CurrentJobId = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetDescription(v string) *GetExplorerTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetExplorerTaskId(v string) *GetExplorerTaskResponseBodyTask {
	s.ExplorerTaskId = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetInitModuleState(v bool) *GetExplorerTaskResponseBodyTask {
	s.InitModuleState = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetModuleId(v string) *GetExplorerTaskResponseBodyTask {
	s.ModuleId = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetModuleVersion(v string) *GetExplorerTaskResponseBodyTask {
	s.ModuleVersion = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetName(v string) *GetExplorerTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetSkipPropertyValidation(v bool) *GetExplorerTaskResponseBodyTask {
	s.SkipPropertyValidation = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetStatus(v string) *GetExplorerTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetTaskOutputPath(v string) *GetExplorerTaskResponseBodyTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetTerraformVersion(v string) *GetExplorerTaskResponseBodyTask {
	s.TerraformVersion = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetTriggerValue(v string) *GetExplorerTaskResponseBodyTask {
	s.TriggerValue = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTask) SetUploadSignature(v *GetExplorerTaskResponseBodyTaskUploadSignature) *GetExplorerTaskResponseBodyTask {
	s.UploadSignature = v
	return s
}

type GetExplorerTaskResponseBodyTaskUploadSignature struct {
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	BucketName  *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	Key         *string `json:"key,omitempty" xml:"key,omitempty"`
	Policy      *string `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature   *string `json:"signature,omitempty" xml:"signature,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetExplorerTaskResponseBodyTaskUploadSignature) String() string {
	return tea.Prettify(s)
}

func (s GetExplorerTaskResponseBodyTaskUploadSignature) GoString() string {
	return s.String()
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetAccessKeyId(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.AccessKeyId = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetBucketName(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.BucketName = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetKey(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.Key = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetPolicy(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.Policy = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetSignature(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.Signature = &v
	return s
}

func (s *GetExplorerTaskResponseBodyTaskUploadSignature) SetUrl(v string) *GetExplorerTaskResponseBodyTaskUploadSignature {
	s.Url = &v
	return s
}

type GetExplorerTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExplorerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExplorerTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExplorerTaskResponse) GoString() string {
	return s.String()
}

func (s *GetExplorerTaskResponse) SetHeaders(v map[string]*string) *GetExplorerTaskResponse {
	s.Headers = v
	return s
}

func (s *GetExplorerTaskResponse) SetStatusCode(v int32) *GetExplorerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExplorerTaskResponse) SetBody(v *GetExplorerTaskResponseBody) *GetExplorerTaskResponse {
	s.Body = v
	return s
}

type GetGroupResponseBody struct {
	Group *GetGroupResponseBodyGroup `json:"group,omitempty" xml:"group,omitempty" type:"Struct"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupResponseBodyGroup struct {
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// example:
	//
	// 2022-08-21T10:57:11Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	ForcedSetting *bool `json:"forcedSetting,omitempty" xml:"forcedSetting,omitempty"`
	// example:
	//
	// g-14e80de4866bf7ffed0bab6154d738
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// abc
	Name                 *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	NotifyConfig         []*GetGroupResponseBodyGroupNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	NotifyOperationTypes []*string                                `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// example:
	//
	// p-4267dcfbf1b6d126edcadf0e949
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// ramRoleName
	RamRole           *string   `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// example:
	//
	// /
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// example:
	//
	// 3
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
	// example:
	//
	// 1.191.0
	TerraformProviderVersion *string                                   `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TriggerConfig            []*GetGroupResponseBodyGroupTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	TriggerResourceType      []*string                                 `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) SetAutoDestroy(v bool) *GetGroupResponseBodyGroup {
	s.AutoDestroy = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetAutoTrigger(v bool) *GetGroupResponseBodyGroup {
	s.AutoTrigger = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetCreateTime(v string) *GetGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDescription(v string) *GetGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetForcedSetting(v bool) *GetGroupResponseBodyGroup {
	s.ForcedSetting = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetName(v string) *GetGroupResponseBodyGroup {
	s.Name = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetNotifyConfig(v []*GetGroupResponseBodyGroupNotifyConfig) *GetGroupResponseBodyGroup {
	s.NotifyConfig = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetNotifyOperationTypes(v []*string) *GetGroupResponseBodyGroup {
	s.NotifyOperationTypes = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetProjectId(v string) *GetGroupResponseBodyGroup {
	s.ProjectId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetRamRole(v string) *GetGroupResponseBodyGroup {
	s.RamRole = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetReportExportField(v []*string) *GetGroupResponseBodyGroup {
	s.ReportExportField = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetReportExportPath(v string) *GetGroupResponseBodyGroup {
	s.ReportExportPath = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTaskCnt(v int64) *GetGroupResponseBodyGroup {
	s.TaskCnt = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTerraformProviderVersion(v string) *GetGroupResponseBodyGroup {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTriggerConfig(v []*GetGroupResponseBodyGroupTriggerConfig) *GetGroupResponseBodyGroup {
	s.TriggerConfig = v
	return s
}

func (s *GetGroupResponseBodyGroup) SetTriggerResourceType(v []*string) *GetGroupResponseBodyGroup {
	s.TriggerResourceType = v
	return s
}

type GetGroupResponseBodyGroupNotifyConfig struct {
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s GetGroupResponseBodyGroupNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroupNotifyConfig) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroupNotifyConfig) SetNotifyPath(v string) *GetGroupResponseBodyGroupNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *GetGroupResponseBodyGroupNotifyConfig) SetNotifyType(v string) *GetGroupResponseBodyGroupNotifyConfig {
	s.NotifyType = &v
	return s
}

type GetGroupResponseBodyGroupTriggerConfig struct {
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// 0 0 8 	- 	- ?
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s GetGroupResponseBodyGroupTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroupTriggerConfig) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroupTriggerConfig) SetTriggerStrategy(v string) *GetGroupResponseBodyGroupTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *GetGroupResponseBodyGroupTriggerConfig) SetTriggerValue(v string) *GetGroupResponseBodyGroupTriggerConfig {
	s.TriggerValue = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetJobRequest struct {
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetTaskType(v string) *GetJobRequest {
	s.TaskType = &v
	return s
}

type GetJobResponseBody struct {
	Job *GetJobResponseBodyJob `json:"job,omitempty" xml:"job,omitempty" type:"Struct"`
	// example:
	//
	// 1435C78A-AED9-53D6-B7A6-E2661D29B1FA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetJob(v *GetJobResponseBodyJob) *GetJobResponseBody {
	s.Job = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResponseBodyJob struct {
	AssertCheckDetail []*GetJobResponseBodyJobAssertCheckDetail `json:"assertCheckDetail,omitempty" xml:"assertCheckDetail,omitempty" type:"Repeated"`
	Config            *GetJobResponseBodyJobConfig              `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-08-31T03:38:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test1
	Description *string                `json:"description,omitempty" xml:"description,omitempty"`
	DownloadUrl map[string]interface{} `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	ElapsedTime *int64                 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExecuteType *string                `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// example:
	//
	// job-518855d9a058cfff0dc933e6b5767
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// /
	Output      *string            `json:"output,omitempty" xml:"output,omitempty"`
	Parameters  map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	RuntimeType *string            `json:"runtimeType,omitempty" xml:"runtimeType,omitempty"`
	// example:
	//
	// Errored
	Status       *string                          `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail map[string]*JobStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// example:
	//
	// task-3b6cb9fa4751a1b9b5f22cbcf4e
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskType                 *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s GetJobResponseBodyJob) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJob) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJob) SetAssertCheckDetail(v []*GetJobResponseBodyJobAssertCheckDetail) *GetJobResponseBodyJob {
	s.AssertCheckDetail = v
	return s
}

func (s *GetJobResponseBodyJob) SetConfig(v *GetJobResponseBodyJobConfig) *GetJobResponseBodyJob {
	s.Config = v
	return s
}

func (s *GetJobResponseBodyJob) SetCreateTime(v string) *GetJobResponseBodyJob {
	s.CreateTime = &v
	return s
}

func (s *GetJobResponseBodyJob) SetDescription(v string) *GetJobResponseBodyJob {
	s.Description = &v
	return s
}

func (s *GetJobResponseBodyJob) SetDownloadUrl(v map[string]interface{}) *GetJobResponseBodyJob {
	s.DownloadUrl = v
	return s
}

func (s *GetJobResponseBodyJob) SetElapsedTime(v int64) *GetJobResponseBodyJob {
	s.ElapsedTime = &v
	return s
}

func (s *GetJobResponseBodyJob) SetExecuteType(v string) *GetJobResponseBodyJob {
	s.ExecuteType = &v
	return s
}

func (s *GetJobResponseBodyJob) SetIsPassAssertCheck(v bool) *GetJobResponseBodyJob {
	s.IsPassAssertCheck = &v
	return s
}

func (s *GetJobResponseBodyJob) SetJobId(v string) *GetJobResponseBodyJob {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetOutput(v string) *GetJobResponseBodyJob {
	s.Output = &v
	return s
}

func (s *GetJobResponseBodyJob) SetParameters(v map[string]*string) *GetJobResponseBodyJob {
	s.Parameters = v
	return s
}

func (s *GetJobResponseBodyJob) SetRuntimeType(v string) *GetJobResponseBodyJob {
	s.RuntimeType = &v
	return s
}

func (s *GetJobResponseBodyJob) SetStatus(v string) *GetJobResponseBodyJob {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyJob) SetStatusDetail(v map[string]*JobStatusDetailValue) *GetJobResponseBodyJob {
	s.StatusDetail = v
	return s
}

func (s *GetJobResponseBodyJob) SetTaskId(v string) *GetJobResponseBodyJob {
	s.TaskId = &v
	return s
}

func (s *GetJobResponseBodyJob) SetTaskType(v string) *GetJobResponseBodyJob {
	s.TaskType = &v
	return s
}

func (s *GetJobResponseBodyJob) SetTerraformProviderVersion(v string) *GetJobResponseBodyJob {
	s.TerraformProviderVersion = &v
	return s
}

type GetJobResponseBodyJobAssertCheckDetail struct {
	// example:
	//
	// eq
	Comparison    *string `json:"comparison,omitempty" xml:"comparison,omitempty"`
	ExpectedValue *string `json:"expectedValue,omitempty" xml:"expectedValue,omitempty"`
	// example:
	//
	// true
	IsPass *bool `json:"isPass,omitempty" xml:"isPass,omitempty"`
	// example:
	//
	// result
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetJobResponseBodyJobAssertCheckDetail) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobAssertCheckDetail) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetComparison(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.Comparison = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetExpectedValue(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.ExpectedValue = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetIsPass(v bool) *GetJobResponseBodyJobAssertCheckDetail {
	s.IsPass = &v
	return s
}

func (s *GetJobResponseBodyJobAssertCheckDetail) SetType(v string) *GetJobResponseBodyJobAssertCheckDetail {
	s.Type = &v
	return s
}

type GetJobResponseBodyJobConfig struct {
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// fales
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// {}
	ResourcesChanged *string `json:"resourcesChanged,omitempty" xml:"resourcesChanged,omitempty"`
	SubCommand       *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
}

func (s GetJobResponseBodyJobConfig) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyJobConfig) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyJobConfig) SetAutoApply(v bool) *GetJobResponseBodyJobConfig {
	s.AutoApply = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetIsDestroy(v bool) *GetJobResponseBodyJobConfig {
	s.IsDestroy = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetModuleVersion(v string) *GetJobResponseBodyJobConfig {
	s.ModuleVersion = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetResourcesChanged(v string) *GetJobResponseBodyJobConfig {
	s.ResourcesChanged = &v
	return s
}

func (s *GetJobResponseBodyJobConfig) SetSubCommand(v string) *GetJobResponseBodyJobConfig {
	s.SubCommand = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetModuleResponseBody struct {
	Module *GetModuleResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1E7BA3EB-B0EF-53F5-9999-07CAD6D9F8A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetModuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBody) SetModule(v *GetModuleResponseBodyModule) *GetModuleResponseBody {
	s.Module = v
	return s
}

func (s *GetModuleResponseBody) SetRequestId(v string) *GetModuleResponseBody {
	s.RequestId = &v
	return s
}

type GetModuleResponseBodyModule struct {
	// example:
	//
	// 2022-09-06T06:11:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// v1
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// example:
	//
	// mod-4267dcfbf1b6d14625614ddbe15
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// /
	OutputPath *string `json:"outputPath,omitempty" xml:"outputPath,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// OSS：
	//
	// "oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip"
	//
	// Registry：
	//
	// "alibaba/security-group/alicloud:2.1.0"
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s GetModuleResponseBodyModule) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetModuleResponseBodyModule) SetCreateTime(v string) *GetModuleResponseBodyModule {
	s.CreateTime = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetDescription(v string) *GetModuleResponseBodyModule {
	s.Description = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetLatestVersion(v string) *GetModuleResponseBodyModule {
	s.LatestVersion = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetModuleId(v string) *GetModuleResponseBodyModule {
	s.ModuleId = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetName(v string) *GetModuleResponseBodyModule {
	s.Name = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetOutputPath(v string) *GetModuleResponseBodyModule {
	s.OutputPath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetSource(v string) *GetModuleResponseBodyModule {
	s.Source = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetSourcePath(v string) *GetModuleResponseBodyModule {
	s.SourcePath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetStatePath(v string) *GetModuleResponseBodyModule {
	s.StatePath = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetStatus(v string) *GetModuleResponseBodyModule {
	s.Status = &v
	return s
}

func (s *GetModuleResponseBodyModule) SetVersionStrategy(v string) *GetModuleResponseBodyModule {
	s.VersionStrategy = &v
	return s
}

type GetModuleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModuleResponse) GoString() string {
	return s.String()
}

func (s *GetModuleResponse) SetHeaders(v map[string]*string) *GetModuleResponse {
	s.Headers = v
	return s
}

func (s *GetModuleResponse) SetStatusCode(v int32) *GetModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleResponse) SetBody(v *GetModuleResponseBody) *GetModuleResponse {
	s.Body = v
	return s
}

type GetModuleVersionResponseBody struct {
	// example:
	//
	// 0D298375-F92F-5B65-82E4-EA68F02521F1
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Version   *GetModuleVersionResponseBodyVersion `json:"version,omitempty" xml:"version,omitempty" type:"Struct"`
}

func (s GetModuleVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponseBody) SetRequestId(v string) *GetModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModuleVersionResponseBody) SetVersion(v *GetModuleVersionResponseBodyVersion) *GetModuleVersionResponseBody {
	s.Version = v
	return s
}

type GetModuleVersionResponseBodyVersion struct {
	// example:
	//
	// 2022-09-08T18:07:40Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// mod-4267dcfbf1b6dfffbc27e218d1b66
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath        *string                `json:"statePath,omitempty" xml:"statePath,omitempty"`
	TerraformContext map[string]interface{} `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s GetModuleVersionResponseBodyVersion) String() string {
	return tea.Prettify(s)
}

func (s GetModuleVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponseBodyVersion) SetCreateTime(v string) *GetModuleVersionResponseBodyVersion {
	s.CreateTime = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetDescription(v string) *GetModuleVersionResponseBodyVersion {
	s.Description = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetModuleId(v string) *GetModuleVersionResponseBodyVersion {
	s.ModuleId = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetModuleVersion(v string) *GetModuleVersionResponseBodyVersion {
	s.ModuleVersion = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetName(v string) *GetModuleVersionResponseBodyVersion {
	s.Name = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetSource(v string) *GetModuleVersionResponseBodyVersion {
	s.Source = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetSourcePath(v string) *GetModuleVersionResponseBodyVersion {
	s.SourcePath = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetStatePath(v string) *GetModuleVersionResponseBodyVersion {
	s.StatePath = &v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetTerraformContext(v map[string]interface{}) *GetModuleVersionResponseBodyVersion {
	s.TerraformContext = v
	return s
}

func (s *GetModuleVersionResponseBodyVersion) SetVersionStrategy(v string) *GetModuleVersionResponseBodyVersion {
	s.VersionStrategy = &v
	return s
}

type GetModuleVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *GetModuleVersionResponse) SetHeaders(v map[string]*string) *GetModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *GetModuleVersionResponse) SetStatusCode(v int32) *GetModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleVersionResponse) SetBody(v *GetModuleVersionResponseBody) *GetModuleVersionResponse {
	s.Body = v
	return s
}

type GetParameterSetResponseBody struct {
	ParameterSet *GetParameterSetResponseBodyParameterSet `json:"parameterSet,omitempty" xml:"parameterSet,omitempty" type:"Struct"`
	// example:
	//
	// 99905C7C-1320-5E7F-A798-3071482EB08E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetParameterSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBody) SetParameterSet(v *GetParameterSetResponseBodyParameterSet) *GetParameterSetResponseBody {
	s.ParameterSet = v
	return s
}

func (s *GetParameterSetResponseBody) SetRequestId(v string) *GetParameterSetResponseBody {
	s.RequestId = &v
	return s
}

type GetParameterSetResponseBodyParameterSet struct {
	// example:
	//
	// 2022-01-30T02:14:16Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pts-3b6cb9fa4751afff9c5e4e01624b9
	ParameterSetId *string                                                `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	Parameters     []*GetParameterSetResponseBodyParameterSetParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	RelationList   []*GetParameterSetResponseBodyParameterSetRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s GetParameterSetResponseBodyParameterSet) String() string {
	return tea.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSet) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSet) SetCreateTime(v string) *GetParameterSetResponseBodyParameterSet {
	s.CreateTime = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetDescription(v string) *GetParameterSetResponseBodyParameterSet {
	s.Description = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetName(v string) *GetParameterSetResponseBodyParameterSet {
	s.Name = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetParameterSetId(v string) *GetParameterSetResponseBodyParameterSet {
	s.ParameterSetId = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetParameters(v []*GetParameterSetResponseBodyParameterSetParameters) *GetParameterSetResponseBodyParameterSet {
	s.Parameters = v
	return s
}

func (s *GetParameterSetResponseBodyParameterSet) SetRelationList(v []*GetParameterSetResponseBodyParameterSetRelationList) *GetParameterSetResponseBodyParameterSet {
	s.RelationList = v
	return s
}

type GetParameterSetResponseBodyParameterSetParameters struct {
	// example:
	//
	// test1121
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-2ze83xrka9ktxy0pnaxn5
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetParameterSetResponseBodyParameterSetParameters) String() string {
	return tea.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSetParameters) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetName(v string) *GetParameterSetResponseBodyParameterSetParameters {
	s.Name = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetType(v string) *GetParameterSetResponseBodyParameterSetParameters {
	s.Type = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetParameters) SetValue(v interface{}) *GetParameterSetResponseBodyParameterSetParameters {
	s.Value = v
	return s
}

type GetParameterSetResponseBodyParameterSetRelationList struct {
	// example:
	//
	// 2022-04-24T22:58:50Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// mod-433aead756057101546eb5d50c1
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetParameterSetResponseBodyParameterSetRelationList) String() string {
	return tea.Prettify(s)
}

func (s GetParameterSetResponseBodyParameterSetRelationList) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetCreateTime(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.CreateTime = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetResourceId(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.ResourceId = &v
	return s
}

func (s *GetParameterSetResponseBodyParameterSetRelationList) SetResourceType(v string) *GetParameterSetResponseBodyParameterSetRelationList {
	s.ResourceType = &v
	return s
}

type GetParameterSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParameterSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParameterSetResponse) GoString() string {
	return s.String()
}

func (s *GetParameterSetResponse) SetHeaders(v map[string]*string) *GetParameterSetResponse {
	s.Headers = v
	return s
}

func (s *GetParameterSetResponse) SetStatusCode(v int32) *GetParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParameterSetResponse) SetBody(v *GetParameterSetResponseBody) *GetParameterSetResponse {
	s.Body = v
	return s
}

type GetProjectResponseBody struct {
	Project *GetProjectResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponseBodyProject struct {
	// example:
	//
	// 2022-09-06T06:11:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 2
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) SetCreateTime(v string) *GetProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetName(v string) *GetProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetProjectId(v string) *GetProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetTaskCnt(v int64) *GetProjectResponseBodyProject {
	s.TaskCnt = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetProjectBuildContextRequest struct {
	IsPassAssertCheck *bool   `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetProjectBuildContextRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextRequest) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextRequest) SetIsPassAssertCheck(v bool) *GetProjectBuildContextRequest {
	s.IsPassAssertCheck = &v
	return s
}

func (s *GetProjectBuildContextRequest) SetStatus(v string) *GetProjectBuildContextRequest {
	s.Status = &v
	return s
}

type GetProjectBuildContextResponseBody struct {
	ProjectBuild *GetProjectBuildContextResponseBodyProjectBuild `json:"ProjectBuild,omitempty" xml:"ProjectBuild,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C54F3939-9112-5795-A5A7-5322E2FB2257
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectBuildContextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextResponseBody) SetProjectBuild(v *GetProjectBuildContextResponseBodyProjectBuild) *GetProjectBuildContextResponseBody {
	s.ProjectBuild = v
	return s
}

func (s *GetProjectBuildContextResponseBody) SetRequestId(v string) *GetProjectBuildContextResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectBuildContextResponseBodyProjectBuild struct {
	// example:
	//
	// 1646719546155
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2
	FailCnt *int64                                                   `json:"failCnt,omitempty" xml:"failCnt,omitempty"`
	JobList []*GetProjectBuildContextResponseBodyProjectBuildJobList `json:"jobList,omitempty" xml:"jobList,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	JobTotalCnt *int64 `json:"jobTotalCnt,omitempty" xml:"jobTotalCnt,omitempty"`
	// example:
	//
	// pb-4267dcfbf1b6d13c7d2386cba7
	ProjectBuildId *string `json:"projectBuildId,omitempty" xml:"projectBuildId,omitempty"`
	// example:
	//
	// p-4267dcfbf1b6d1e0652bfbbe994
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 0
	ResourceCnt  *int64                                                        `json:"resourceCnt,omitempty" xml:"resourceCnt,omitempty"`
	ResourceList []*GetProjectBuildContextResponseBodyProjectBuildResourceList `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	ResourceTypeCnt *int64 `json:"resourceTypeCnt,omitempty" xml:"resourceTypeCnt,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	SuccessCnt               *int64  `json:"successCnt,omitempty" xml:"successCnt,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TriggerStrategy          *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s GetProjectBuildContextResponseBodyProjectBuild) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextResponseBodyProjectBuild) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetEndTime(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.EndTime = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetFailCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuild {
	s.FailCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetJobList(v []*GetProjectBuildContextResponseBodyProjectBuildJobList) *GetProjectBuildContextResponseBodyProjectBuild {
	s.JobList = v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetJobTotalCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuild {
	s.JobTotalCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetProjectBuildId(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.ProjectBuildId = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetProjectId(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.ProjectId = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetResourceCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuild {
	s.ResourceCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetResourceList(v []*GetProjectBuildContextResponseBodyProjectBuildResourceList) *GetProjectBuildContextResponseBodyProjectBuild {
	s.ResourceList = v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetResourceTypeCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuild {
	s.ResourceTypeCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetStatus(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.Status = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetSuccessCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuild {
	s.SuccessCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetTerraformProviderVersion(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuild) SetTriggerStrategy(v string) *GetProjectBuildContextResponseBodyProjectBuild {
	s.TriggerStrategy = &v
	return s
}

type GetProjectBuildContextResponseBodyProjectBuildJobList struct {
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// 0
	IsDeleted         *int64 `json:"isDeleted,omitempty" xml:"isDeleted,omitempty"`
	IsPassAssertCheck *bool  `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// example:
	//
	// job-518855d9a058cfff262b993646d12
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// mod-3b6cb9fa4751a107afd6c2c3eb9
	ModuleId   *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test7
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// task-433aead756057ffe67aefed4aa75d
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	Type                     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProjectBuildContextResponseBodyProjectBuildJobList) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextResponseBodyProjectBuildJobList) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetElapsedTime(v int64) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.ElapsedTime = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetIsDeleted(v int64) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.IsDeleted = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetIsPassAssertCheck(v bool) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.IsPassAssertCheck = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetJobId(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.JobId = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetModuleId(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.ModuleId = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetModuleName(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.ModuleName = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetModuleVersion(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.ModuleVersion = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetName(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.Name = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetStatus(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.Status = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetTaskId(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.TaskId = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetTerraformProviderVersion(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildJobList) SetType(v string) *GetProjectBuildContextResponseBodyProjectBuildJobList {
	s.Type = &v
	return s
}

type GetProjectBuildContextResponseBodyProjectBuildResourceList struct {
	Info map[string]*string `json:"info,omitempty" xml:"info,omitempty"`
	// example:
	//
	// 1
	ResourceCnt *int64 `json:"resourceCnt,omitempty" xml:"resourceCnt,omitempty"`
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetProjectBuildContextResponseBodyProjectBuildResourceList) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextResponseBodyProjectBuildResourceList) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextResponseBodyProjectBuildResourceList) SetInfo(v map[string]*string) *GetProjectBuildContextResponseBodyProjectBuildResourceList {
	s.Info = v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildResourceList) SetResourceCnt(v int64) *GetProjectBuildContextResponseBodyProjectBuildResourceList {
	s.ResourceCnt = &v
	return s
}

func (s *GetProjectBuildContextResponseBodyProjectBuildResourceList) SetResourceType(v string) *GetProjectBuildContextResponseBodyProjectBuildResourceList {
	s.ResourceType = &v
	return s
}

type GetProjectBuildContextResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProjectBuildContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProjectBuildContextResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectBuildContextResponse) GoString() string {
	return s.String()
}

func (s *GetProjectBuildContextResponse) SetHeaders(v map[string]*string) *GetProjectBuildContextResponse {
	s.Headers = v
	return s
}

func (s *GetProjectBuildContextResponse) SetStatusCode(v int32) *GetProjectBuildContextResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectBuildContextResponse) SetBody(v *GetProjectBuildContextResponseBody) *GetProjectBuildContextResponse {
	s.Body = v
	return s
}

type GetRabbitmqPublisherResponseBody struct {
	Publisher *GetRabbitmqPublisherResponseBodyPublisher `json:"publisher,omitempty" xml:"publisher,omitempty" type:"Struct"`
	// example:
	//
	// 1E7BA3EB-B0EF-53F5-9999-07CAD6D9F8A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRabbitmqPublisherResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRabbitmqPublisherResponseBody) GoString() string {
	return s.String()
}

func (s *GetRabbitmqPublisherResponseBody) SetPublisher(v *GetRabbitmqPublisherResponseBodyPublisher) *GetRabbitmqPublisherResponseBody {
	s.Publisher = v
	return s
}

func (s *GetRabbitmqPublisherResponseBody) SetRequestId(v string) *GetRabbitmqPublisherResponseBody {
	s.RequestId = &v
	return s
}

type GetRabbitmqPublisherResponseBodyPublisher struct {
	// example:
	//
	// 2022-10-05T06:07:48Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ExchangeTest
	ExchangeName *string `json:"exchangeName,omitempty" xml:"exchangeName,omitempty"`
	// example:
	//
	// TOPIC
	ExchangeType *string `json:"exchangeType,omitempty" xml:"exchangeType,omitempty"`
	// example:
	//
	// 1880770****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	HostName *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// example:
	//
	// test123
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5671
	Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// mqp-3b6cb9fa4751a6fa1b69b8dca1
	PublisherId *string   `json:"publisherId,omitempty" xml:"publisherId,omitempty"`
	TaskIds     []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
	// example:
	//
	// MjoxODgwNzcwODY5MD****
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// example:
	//
	// MjoxODgwNzcwODY5MD****
	VirtualHost *string `json:"virtualHost,omitempty" xml:"virtualHost,omitempty"`
}

func (s GetRabbitmqPublisherResponseBodyPublisher) String() string {
	return tea.Prettify(s)
}

func (s GetRabbitmqPublisherResponseBodyPublisher) GoString() string {
	return s.String()
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetCreateTime(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.CreateTime = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetDescription(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.Description = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetExchangeName(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.ExchangeName = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetExchangeType(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.ExchangeType = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetHostName(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.HostName = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetName(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.Name = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetPort(v int64) *GetRabbitmqPublisherResponseBodyPublisher {
	s.Port = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetPublisherId(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.PublisherId = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetTaskIds(v []*string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.TaskIds = v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetUserName(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.UserName = &v
	return s
}

func (s *GetRabbitmqPublisherResponseBodyPublisher) SetVirtualHost(v string) *GetRabbitmqPublisherResponseBodyPublisher {
	s.VirtualHost = &v
	return s
}

type GetRabbitmqPublisherResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRabbitmqPublisherResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRabbitmqPublisherResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRabbitmqPublisherResponse) GoString() string {
	return s.String()
}

func (s *GetRabbitmqPublisherResponse) SetHeaders(v map[string]*string) *GetRabbitmqPublisherResponse {
	s.Headers = v
	return s
}

func (s *GetRabbitmqPublisherResponse) SetStatusCode(v int32) *GetRabbitmqPublisherResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRabbitmqPublisherResponse) SetBody(v *GetRabbitmqPublisherResponseBody) *GetRabbitmqPublisherResponse {
	s.Body = v
	return s
}

type GetRamPolicyExportTaskResponseBody struct {
	RamPolicyExportTask *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask `json:"ramPolicyExportTask,omitempty" xml:"ramPolicyExportTask,omitempty" type:"Struct"`
	// example:
	//
	// 3E49127A-BB65-5CCD-AB93-0EC0A43E5446
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRamPolicyExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskResponseBody) SetRamPolicyExportTask(v *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) *GetRamPolicyExportTaskResponseBody {
	s.RamPolicyExportTask = v
	return s
}

func (s *GetRamPolicyExportTaskResponseBody) SetRequestId(v string) *GetRamPolicyExportTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetRamPolicyExportTaskResponseBodyRamPolicyExportTask struct {
	AuthorizationAccountIds []*int64  `json:"authorizationAccountIds,omitempty" xml:"authorizationAccountIds,omitempty" type:"Repeated"`
	AuthorizationActions    []*string `json:"authorizationActions,omitempty" xml:"authorizationActions,omitempty" type:"Repeated"`
	AuthorizationRegionIds  []*string `json:"authorizationRegionIds,omitempty" xml:"authorizationRegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-06-16T10:03:39Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// mod-66c23a9cc0cacddf
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rpe-40252e48f9286a
	RamPolicyExportTaskId *string `json:"ramPolicyExportTaskId,omitempty" xml:"ramPolicyExportTaskId,omitempty"`
	// example:
	//
	// ramName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 12345678/default/rampolicytask/rpe-4399111870d7e6b0f11a
	TaskOutputPath           *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetAuthorizationAccountIds(v []*int64) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.AuthorizationAccountIds = v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetAuthorizationActions(v []*string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.AuthorizationActions = v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetAuthorizationRegionIds(v []*string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.AuthorizationRegionIds = v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetCreateTime(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.CreateTime = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetModuleId(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.ModuleId = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetModuleVersion(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.ModuleVersion = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetName(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.Name = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetRamPolicyExportTaskId(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.RamPolicyExportTaskId = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetRamRole(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.RamRole = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetStatus(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.Status = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetTaskOutputPath(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetTerraformProviderVersion(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask) SetTriggerStrategy(v string) *GetRamPolicyExportTaskResponseBodyRamPolicyExportTask {
	s.TriggerStrategy = &v
	return s
}

type GetRamPolicyExportTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRamPolicyExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRamPolicyExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskResponse) SetHeaders(v map[string]*string) *GetRamPolicyExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetRamPolicyExportTaskResponse) SetStatusCode(v int32) *GetRamPolicyExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRamPolicyExportTaskResponse) SetBody(v *GetRamPolicyExportTaskResponseBody) *GetRamPolicyExportTaskResponse {
	s.Body = v
	return s
}

type GetRamPolicyExportTaskVersionResponseBody struct {
	RamPolicyExportTaskVersion *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion `json:"ramPolicyExportTaskVersion,omitempty" xml:"ramPolicyExportTaskVersion,omitempty" type:"Struct"`
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRamPolicyExportTaskVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskVersionResponseBody) SetRamPolicyExportTaskVersion(v *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) *GetRamPolicyExportTaskVersionResponseBody {
	s.RamPolicyExportTaskVersion = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBody) SetRequestId(v string) *GetRamPolicyExportTaskVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion struct {
	AuthorizationAccountIds []*int64  `json:"authorizationAccountIds,omitempty" xml:"authorizationAccountIds,omitempty" type:"Repeated"`
	AuthorizationActions    []*string `json:"authorizationActions,omitempty" xml:"authorizationActions,omitempty" type:"Repeated"`
	AuthorizationRegionIds  []*string `json:"authorizationRegionIds,omitempty" xml:"authorizationRegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-18T06:32:06Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ElapsedTime *int64  `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// the ram policy export task has time out, The maximum running time is 1 hour
	FailedReason   *string   `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	MissingActions []*string `json:"missingActions,omitempty" xml:"missingActions,omitempty" type:"Repeated"`
	// example:
	//
	// mod-4357ffeebf6b577c4afa
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion          *string   `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	NoSupportResourceTypes []*string `json:"noSupportResourceTypes,omitempty" xml:"noSupportResourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// "{
	//
	//     "version": "1",
	//
	//     "statement": [
	//
	//       {
	//
	//         "effect": "Allow",
	//
	//         "action": [
	//
	//           "ecs:*"
	//
	//         ],
	//
	//         "resource": [
	//
	//           "*"
	//
	//         ],
	//
	//         "condition": null,
	//
	//         "supportVariables": []
	//
	//       }
	//
	//     ]
	//
	//   }"
	PolicyDocument *string `json:"policyDocument,omitempty" xml:"policyDocument,omitempty"`
	// example:
	//
	// rpe-43ffe0252e48f9286a
	RamPolicyExportTaskId *string `json:"ramPolicyExportTaskId,omitempty" xml:"ramPolicyExportTaskId,omitempty"`
	// example:
	//
	// Success
	Status                   *string   `json:"status,omitempty" xml:"status,omitempty"`
	TerraformProviderVersion *string   `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	UnauthorizedActions      []*string `json:"unauthorizedActions,omitempty" xml:"unauthorizedActions,omitempty" type:"Repeated"`
	WarnMessage              *string   `json:"warnMessage,omitempty" xml:"warnMessage,omitempty"`
}

func (s GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetAuthorizationAccountIds(v []*int64) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.AuthorizationAccountIds = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetAuthorizationActions(v []*string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.AuthorizationActions = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetAuthorizationRegionIds(v []*string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.AuthorizationRegionIds = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetCreateTime(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.CreateTime = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetElapsedTime(v int64) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.ElapsedTime = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetExportVersion(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.ExportVersion = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetFailedReason(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.FailedReason = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetMissingActions(v []*string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.MissingActions = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetModuleId(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.ModuleId = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetModuleVersion(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.ModuleVersion = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetNoSupportResourceTypes(v []*string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.NoSupportResourceTypes = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetPolicyDocument(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.PolicyDocument = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetRamPolicyExportTaskId(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.RamPolicyExportTaskId = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetStatus(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.Status = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetTerraformProviderVersion(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetUnauthorizedActions(v []*string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.UnauthorizedActions = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion) SetWarnMessage(v string) *GetRamPolicyExportTaskVersionResponseBodyRamPolicyExportTaskVersion {
	s.WarnMessage = &v
	return s
}

type GetRamPolicyExportTaskVersionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRamPolicyExportTaskVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRamPolicyExportTaskVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRamPolicyExportTaskVersionResponse) GoString() string {
	return s.String()
}

func (s *GetRamPolicyExportTaskVersionResponse) SetHeaders(v map[string]*string) *GetRamPolicyExportTaskVersionResponse {
	s.Headers = v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponse) SetStatusCode(v int32) *GetRamPolicyExportTaskVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRamPolicyExportTaskVersionResponse) SetBody(v *GetRamPolicyExportTaskVersionResponseBody) *GetRamPolicyExportTaskVersionResponse {
	s.Body = v
	return s
}

type GetResourceExportTaskRequest struct {
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
}

func (s GetResourceExportTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskRequest) SetExportVersion(v string) *GetResourceExportTaskRequest {
	s.ExportVersion = &v
	return s
}

type GetResourceExportTaskResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetResourceExportTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetResourceExportTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBody) SetRequestId(v string) *GetResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceExportTaskResponseBody) SetTask(v *GetResourceExportTaskResponseBodyTask) *GetResourceExportTaskResponseBody {
	s.Task = v
	return s
}

type GetResourceExportTaskResponseBodyTask struct {
	ConfigPath               *string                                              `json:"configPath,omitempty" xml:"configPath,omitempty"`
	CreateTime               *string                                              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description              *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	ElapsedTime              *int64                                               `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExcludeRules             []*GetResourceExportTaskResponseBodyTaskExcludeRules `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	ExportTaskId             *string                                              `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportToModule           *GetResourceExportTaskResponseBodyTaskExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	ExportVersion            *string                                              `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	FailedReason             *string                                              `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	IncludeRules             []*GetResourceExportTaskResponseBodyTaskIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	Modules                  []*GetResourceExportTaskResponseBodyTaskModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	Name                     *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	RamRole                  *string                                              `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	Status                   *string                                              `json:"status,omitempty" xml:"status,omitempty"`
	TaskOutputPath           *string                                              `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	TerraformContext         map[string]interface{}                               `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
	TerraformProviderVersion *string                                              `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TerraformVersion         *string                                              `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	TriggerStrategy          *string                                              `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables                []*GetResourceExportTaskResponseBodyTaskVariables    `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetResourceExportTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTask) SetConfigPath(v string) *GetResourceExportTaskResponseBodyTask {
	s.ConfigPath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetCreateTime(v string) *GetResourceExportTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetDescription(v string) *GetResourceExportTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetElapsedTime(v int64) *GetResourceExportTaskResponseBodyTask {
	s.ElapsedTime = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExcludeRules(v []*GetResourceExportTaskResponseBodyTaskExcludeRules) *GetResourceExportTaskResponseBodyTask {
	s.ExcludeRules = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportTaskId(v string) *GetResourceExportTaskResponseBodyTask {
	s.ExportTaskId = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportToModule(v *GetResourceExportTaskResponseBodyTaskExportToModule) *GetResourceExportTaskResponseBodyTask {
	s.ExportToModule = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.ExportVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetFailedReason(v string) *GetResourceExportTaskResponseBodyTask {
	s.FailedReason = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetIncludeRules(v []*GetResourceExportTaskResponseBodyTaskIncludeRules) *GetResourceExportTaskResponseBodyTask {
	s.IncludeRules = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetModules(v []*GetResourceExportTaskResponseBodyTaskModules) *GetResourceExportTaskResponseBodyTask {
	s.Modules = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetName(v string) *GetResourceExportTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetRamRole(v string) *GetResourceExportTaskResponseBodyTask {
	s.RamRole = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetStatus(v string) *GetResourceExportTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTaskOutputPath(v string) *GetResourceExportTaskResponseBodyTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformContext(v map[string]interface{}) *GetResourceExportTaskResponseBodyTask {
	s.TerraformContext = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformProviderVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.TerraformVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTriggerStrategy(v string) *GetResourceExportTaskResponseBodyTask {
	s.TriggerStrategy = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetVariables(v []*GetResourceExportTaskResponseBodyTaskVariables) *GetResourceExportTaskResponseBodyTask {
	s.Variables = v
	return s
}

type GetResourceExportTaskResponseBodyTaskExcludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetResourceExportTaskResponseBodyTaskExcludeRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskExcludeRules) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskExcludeRules) SetKey(v string) *GetResourceExportTaskResponseBodyTaskExcludeRules {
	s.Key = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExcludeRules) SetValues(v []*string) *GetResourceExportTaskResponseBodyTaskExcludeRules {
	s.Values = v
	return s
}

type GetResourceExportTaskResponseBodyTaskExportToModule struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	StatePath  *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskExportToModule) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskExportToModule) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetSource(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.Source = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetSourcePath(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.SourcePath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetStatePath(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.StatePath = &v
	return s
}

type GetResourceExportTaskResponseBodyTaskIncludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetResourceExportTaskResponseBodyTaskIncludeRules) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskIncludeRules) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) SetKey(v string) *GetResourceExportTaskResponseBodyTaskIncludeRules {
	s.Key = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) SetValues(v []*string) *GetResourceExportTaskResponseBodyTaskIncludeRules {
	s.Values = v
	return s
}

type GetResourceExportTaskResponseBodyTaskModules struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskModules) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskModules) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetSource(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.Source = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetSourcePath(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.SourcePath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetVersion(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.Version = &v
	return s
}

type GetResourceExportTaskResponseBodyTaskVariables struct {
	Properties   []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskVariables) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskVariables) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) SetProperties(v []*string) *GetResourceExportTaskResponseBodyTaskVariables {
	s.Properties = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) SetResourceType(v string) *GetResourceExportTaskResponseBodyTaskVariables {
	s.ResourceType = &v
	return s
}

type GetResourceExportTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceExportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceExportTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceExportTaskResponse) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponse) SetHeaders(v map[string]*string) *GetResourceExportTaskResponse {
	s.Headers = v
	return s
}

func (s *GetResourceExportTaskResponse) SetStatusCode(v int32) *GetResourceExportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceExportTaskResponse) SetBody(v *GetResourceExportTaskResponseBody) *GetResourceExportTaskResponse {
	s.Body = v
	return s
}

type GetTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// C24C498A-09CF-54D3-8972-8DC074CF8614
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

type GetTaskResponseBodyTask struct {
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// false
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// example:
	//
	// 2022-06-15T02:44:37Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// job-absdf
	CurrentJobId *string `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	// example:
	//
	// demo
	Description     *string                           `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo       *GetTaskResponseBodyTaskGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	InitModuleState *bool                             `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// example:
	//
	// mod-4267dcfbf1b6d14625614ddbe15
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v2
	ModuleVersion      *string            `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	Name               *string            `json:"name,omitempty" xml:"name,omitempty"`
	Parameters         map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	ProtectionStrategy []*string          `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	RamRole                *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	SkipPropertyValidation *bool   `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// task-433aead756057154bda7f1c2e98
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// /
	TaskOutputPath *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// ***10*
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) SetAutoApply(v bool) *GetTaskResponseBodyTask {
	s.AutoApply = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetAutoDestroy(v bool) *GetTaskResponseBodyTask {
	s.AutoDestroy = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCreateTime(v string) *GetTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetCurrentJobId(v string) *GetTaskResponseBodyTask {
	s.CurrentJobId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetDescription(v string) *GetTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetGroupInfo(v *GetTaskResponseBodyTaskGroupInfo) *GetTaskResponseBodyTask {
	s.GroupInfo = v
	return s
}

func (s *GetTaskResponseBodyTask) SetInitModuleState(v bool) *GetTaskResponseBodyTask {
	s.InitModuleState = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModuleId(v string) *GetTaskResponseBodyTask {
	s.ModuleId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetModuleVersion(v string) *GetTaskResponseBodyTask {
	s.ModuleVersion = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetName(v string) *GetTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetParameters(v map[string]*string) *GetTaskResponseBodyTask {
	s.Parameters = v
	return s
}

func (s *GetTaskResponseBodyTask) SetProtectionStrategy(v []*string) *GetTaskResponseBodyTask {
	s.ProtectionStrategy = v
	return s
}

func (s *GetTaskResponseBodyTask) SetRamRole(v string) *GetTaskResponseBodyTask {
	s.RamRole = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetSkipPropertyValidation(v bool) *GetTaskResponseBodyTask {
	s.SkipPropertyValidation = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskOutputPath(v string) *GetTaskResponseBodyTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTerraformVersion(v string) *GetTaskResponseBodyTask {
	s.TerraformVersion = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTriggerStrategy(v string) *GetTaskResponseBodyTask {
	s.TriggerStrategy = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTriggerValue(v string) *GetTaskResponseBodyTask {
	s.TriggerValue = &v
	return s
}

type GetTaskResponseBodyTaskGroupInfo struct {
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// p-433aead756057fff47ecbfd94d76
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// abc
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s GetTaskResponseBodyTaskGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskGroupInfo) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetGroupId(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.GroupId = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetGroupName(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.GroupName = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetProjectId(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *GetTaskResponseBodyTaskGroupInfo) SetProjectName(v string) *GetTaskResponseBodyTaskGroupInfo {
	s.ProjectName = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskPolicyRequest struct {
	// example:
	//
	// SceneTestingTask
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTaskPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetTaskPolicyRequest) SetType(v string) *GetTaskPolicyRequest {
	s.Type = &v
	return s
}

type GetTaskPolicyResponseBody struct {
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-92C7799512E3
	RequestId  *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskPolicy *GetTaskPolicyResponseBodyTaskPolicy `json:"taskPolicy,omitempty" xml:"taskPolicy,omitempty" type:"Struct"`
}

func (s GetTaskPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskPolicyResponseBody) SetRequestId(v string) *GetTaskPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskPolicyResponseBody) SetTaskPolicy(v *GetTaskPolicyResponseBodyTaskPolicy) *GetTaskPolicyResponseBody {
	s.TaskPolicy = v
	return s
}

type GetTaskPolicyResponseBodyTaskPolicy struct {
	// example:
	//
	// g-433aead7560571e66e31274ffd3
	GroupId      *string                                            `json:"groupId,omitempty" xml:"groupId,omitempty"`
	TaskPolicies []*GetTaskPolicyResponseBodyTaskPolicyTaskPolicies `json:"taskPolicies,omitempty" xml:"taskPolicies,omitempty" type:"Repeated"`
}

func (s GetTaskPolicyResponseBodyTaskPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetTaskPolicyResponseBodyTaskPolicy) GoString() string {
	return s.String()
}

func (s *GetTaskPolicyResponseBodyTaskPolicy) SetGroupId(v string) *GetTaskPolicyResponseBodyTaskPolicy {
	s.GroupId = &v
	return s
}

func (s *GetTaskPolicyResponseBodyTaskPolicy) SetTaskPolicies(v []*GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) *GetTaskPolicyResponseBodyTaskPolicy {
	s.TaskPolicies = v
	return s
}

type GetTaskPolicyResponseBodyTaskPolicyTaskPolicies struct {
	// example:
	//
	// 5
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// task-433aead7560571d0938bfbbe957
	TaskId   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// example:
	//
	// SceneTestingTask
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) String() string {
	return tea.Prettify(s)
}

func (s GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) GoString() string {
	return s.String()
}

func (s *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) SetPriority(v int64) *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies {
	s.Priority = &v
	return s
}

func (s *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) SetTaskId(v string) *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies {
	s.TaskId = &v
	return s
}

func (s *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) SetTaskName(v string) *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies {
	s.TaskName = &v
	return s
}

func (s *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies) SetType(v string) *GetTaskPolicyResponseBodyTaskPolicyTaskPolicies {
	s.Type = &v
	return s
}

type GetTaskPolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetTaskPolicyResponse) SetHeaders(v map[string]*string) *GetTaskPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetTaskPolicyResponse) SetStatusCode(v int32) *GetTaskPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskPolicyResponse) SetBody(v *GetTaskPolicyResponseBody) *GetTaskPolicyResponse {
	s.Body = v
	return s
}

type ListAuthorizationsRequest struct {
	// example:
	//
	// auth-433aead756057ffec22d5b1ef27ac
	AuthorizationId *string `json:"authorizationId,omitempty" xml:"authorizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// others
	AuthorizationType *string `json:"authorizationType,omitempty" xml:"authorizationType,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAuthorizationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizationsRequest) SetAuthorizationId(v string) *ListAuthorizationsRequest {
	s.AuthorizationId = &v
	return s
}

func (s *ListAuthorizationsRequest) SetAuthorizationType(v string) *ListAuthorizationsRequest {
	s.AuthorizationType = &v
	return s
}

func (s *ListAuthorizationsRequest) SetKeyword(v string) *ListAuthorizationsRequest {
	s.Keyword = &v
	return s
}

func (s *ListAuthorizationsRequest) SetPageNumber(v int32) *ListAuthorizationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizationsRequest) SetPageSize(v int32) *ListAuthorizationsRequest {
	s.PageSize = &v
	return s
}

type ListAuthorizationsResponseBody struct {
	Authorizations []*ListAuthorizationsResponseBodyAuthorizations `json:"authorizations,omitempty" xml:"authorizations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C67A913A-2A0F-53F6-A041-56CC4CA1E593
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 72
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAuthorizationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationsResponseBody) SetAuthorizations(v []*ListAuthorizationsResponseBodyAuthorizations) *ListAuthorizationsResponseBody {
	s.Authorizations = v
	return s
}

func (s *ListAuthorizationsResponseBody) SetPageNumber(v int32) *ListAuthorizationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizationsResponseBody) SetPageSize(v int32) *ListAuthorizationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizationsResponseBody) SetRequestId(v string) *ListAuthorizationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationsResponseBody) SetTotalCount(v int32) *ListAuthorizationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAuthorizationsResponseBodyAuthorizations struct {
	// example:
	//
	// auth-433aead756057ffee37b763564fdd
	AuthorizationId *string `json:"authorizationId,omitempty" xml:"authorizationId,omitempty"`
	// example:
	//
	// 2022-06-16T03:41:34Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DueTime    *string `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// example:
	//
	// mod-395f8626622affff71ccbf5b25885
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// aadfaf
	OwnerUid *int64 `json:"ownerUid,omitempty" xml:"ownerUid,omitempty"`
	// example:
	//
	// asdf123
	Uid *int64 `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListAuthorizationsResponseBodyAuthorizations) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationsResponseBodyAuthorizations) GoString() string {
	return s.String()
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetAuthorizationId(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.AuthorizationId = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetCreateTime(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.CreateTime = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetDueTime(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.DueTime = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetModuleId(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.ModuleId = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetModuleVersion(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.ModuleVersion = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetName(v string) *ListAuthorizationsResponseBodyAuthorizations {
	s.Name = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetOwnerUid(v int64) *ListAuthorizationsResponseBodyAuthorizations {
	s.OwnerUid = &v
	return s
}

func (s *ListAuthorizationsResponseBodyAuthorizations) SetUid(v int64) *ListAuthorizationsResponseBodyAuthorizations {
	s.Uid = &v
	return s
}

type ListAuthorizationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorizationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorizationsResponse) SetHeaders(v map[string]*string) *ListAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorizationsResponse) SetStatusCode(v int32) *ListAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorizationsResponse) SetBody(v *ListAuthorizationsResponseBody) *ListAuthorizationsResponse {
	s.Body = v
	return s
}

type ListAvailableTerraformVersionsRequest struct {
	// example:
	//
	// 1.35
	KeyWord *string `json:"keyWord,omitempty" xml:"keyWord,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAvailableTerraformVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableTerraformVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableTerraformVersionsRequest) SetKeyWord(v string) *ListAvailableTerraformVersionsRequest {
	s.KeyWord = &v
	return s
}

func (s *ListAvailableTerraformVersionsRequest) SetPageNumber(v int64) *ListAvailableTerraformVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAvailableTerraformVersionsRequest) SetPageSize(v int64) *ListAvailableTerraformVersionsRequest {
	s.PageSize = &v
	return s
}

type ListAvailableTerraformVersionsResponseBody struct {
	// example:
	//
	// b19952f6-1e03-43e3-ac9b-1be20c9ac1a6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 58
	TotalCount  *int32    `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	VerisonList []*string `json:"verisonList,omitempty" xml:"verisonList,omitempty" type:"Repeated"`
}

func (s ListAvailableTerraformVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableTerraformVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableTerraformVersionsResponseBody) SetRequestId(v string) *ListAvailableTerraformVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableTerraformVersionsResponseBody) SetTotalCount(v int32) *ListAvailableTerraformVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAvailableTerraformVersionsResponseBody) SetVerisonList(v []*string) *ListAvailableTerraformVersionsResponseBody {
	s.VerisonList = v
	return s
}

type ListAvailableTerraformVersionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableTerraformVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableTerraformVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableTerraformVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableTerraformVersionsResponse) SetHeaders(v map[string]*string) *ListAvailableTerraformVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableTerraformVersionsResponse) SetStatusCode(v int32) *ListAvailableTerraformVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableTerraformVersionsResponse) SetBody(v *ListAvailableTerraformVersionsResponseBody) *ListAvailableTerraformVersionsResponse {
	s.Body = v
	return s
}

type ListExplorerTasksRequest struct {
	Keyword   *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	MaxResult *string `json:"maxResult,omitempty" xml:"maxResult,omitempty"`
	ModuleId  *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExplorerTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExplorerTasksRequest) GoString() string {
	return s.String()
}

func (s *ListExplorerTasksRequest) SetKeyword(v string) *ListExplorerTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListExplorerTasksRequest) SetMaxResult(v string) *ListExplorerTasksRequest {
	s.MaxResult = &v
	return s
}

func (s *ListExplorerTasksRequest) SetModuleId(v string) *ListExplorerTasksRequest {
	s.ModuleId = &v
	return s
}

func (s *ListExplorerTasksRequest) SetNextToken(v string) *ListExplorerTasksRequest {
	s.NextToken = &v
	return s
}

type ListExplorerTasksResponseBody struct {
	MaxResults *int32                                `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId  *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Tasks      []*ListExplorerTasksResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExplorerTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExplorerTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListExplorerTasksResponseBody) SetMaxResults(v int32) *ListExplorerTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerTasksResponseBody) SetNextToken(v string) *ListExplorerTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExplorerTasksResponseBody) SetRequestId(v string) *ListExplorerTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExplorerTasksResponseBody) SetTasks(v []*ListExplorerTasksResponseBodyTasks) *ListExplorerTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListExplorerTasksResponseBody) SetTotalCount(v int32) *ListExplorerTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListExplorerTasksResponseBodyTasks struct {
	AutoApply          *bool   `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CurrentJobId       *string `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	CurrentJobStatus   *string `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	DeletionProtection *bool   `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	ModuleId           *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleName         *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	ModuleVersion      *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskId             *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListExplorerTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListExplorerTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListExplorerTasksResponseBodyTasks) SetAutoApply(v bool) *ListExplorerTasksResponseBodyTasks {
	s.AutoApply = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetCreateTime(v string) *ListExplorerTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetCurrentJobId(v string) *ListExplorerTasksResponseBodyTasks {
	s.CurrentJobId = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetCurrentJobStatus(v string) *ListExplorerTasksResponseBodyTasks {
	s.CurrentJobStatus = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetDeletionProtection(v bool) *ListExplorerTasksResponseBodyTasks {
	s.DeletionProtection = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetModuleId(v string) *ListExplorerTasksResponseBodyTasks {
	s.ModuleId = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetModuleName(v string) *ListExplorerTasksResponseBodyTasks {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetModuleVersion(v string) *ListExplorerTasksResponseBodyTasks {
	s.ModuleVersion = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetName(v string) *ListExplorerTasksResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetStatus(v string) *ListExplorerTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListExplorerTasksResponseBodyTasks) SetTaskId(v string) *ListExplorerTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

type ListExplorerTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExplorerTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExplorerTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExplorerTasksResponse) GoString() string {
	return s.String()
}

func (s *ListExplorerTasksResponse) SetHeaders(v map[string]*string) *ListExplorerTasksResponse {
	s.Headers = v
	return s
}

func (s *ListExplorerTasksResponse) SetStatusCode(v int32) *ListExplorerTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExplorerTasksResponse) SetBody(v *ListExplorerTasksResponseBody) *ListExplorerTasksResponse {
	s.Body = v
	return s
}

type ListGroupRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-14e80de4866bf7ffed0c4072ed9b37
	ProjectId *string                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tag       []*ListGroupRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRequest) SetKeyword(v string) *ListGroupRequest {
	s.Keyword = &v
	return s
}

func (s *ListGroupRequest) SetPageNumber(v string) *ListGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupRequest) SetPageSize(v string) *ListGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupRequest) SetProjectId(v string) *ListGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGroupRequest) SetTag(v []*ListGroupRequestTag) *ListGroupRequest {
	s.Tag = v
	return s
}

type ListGroupRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ListGroupRequestTag) SetKey(v string) *ListGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ListGroupRequestTag) SetValue(v string) *ListGroupRequestTag {
	s.Value = &v
	return s
}

type ListGroupShrinkRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 200
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-14e80de4866bf7ffed0c4072ed9b37
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupShrinkRequest) SetKeyword(v string) *ListGroupShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGroupShrinkRequest) SetPageNumber(v string) *ListGroupShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupShrinkRequest) SetPageSize(v string) *ListGroupShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupShrinkRequest) SetProjectId(v string) *ListGroupShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListGroupShrinkRequest) SetTagShrink(v string) *ListGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListGroupResponseBody struct {
	// example:
	//
	// 3
	Count  *int64                         `json:"count,omitempty" xml:"count,omitempty"`
	Groups []*ListGroupResponseBodyGroups `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBody) SetCount(v int64) *ListGroupResponseBody {
	s.Count = &v
	return s
}

func (s *ListGroupResponseBody) SetGroups(v []*ListGroupResponseBodyGroups) *ListGroupResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupResponseBody) SetPageNumber(v int64) *ListGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGroupResponseBody) SetPageSize(v int64) *ListGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGroupResponseBody) SetRequestId(v string) *ListGroupResponseBody {
	s.RequestId = &v
	return s
}

type ListGroupResponseBodyGroups struct {
	// example:
	//
	// 2022-09-14T07:19:13Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// g-148e7853433574fffe9fec72ed9b73
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// example:
	//
	// 1
	ModuleCnt *int64 `json:"moduleCnt,omitempty" xml:"moduleCnt,omitempty"`
	// example:
	//
	// 1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-4267dcfbf1b6d126edcadf0e949
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 1
	SceneTestingTaskCnt *int64                             `json:"sceneTestingTaskCnt,omitempty" xml:"sceneTestingTaskCnt,omitempty"`
	Tags                []*ListGroupResponseBodyGroupsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s ListGroupResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBodyGroups) SetCreateTime(v string) *ListGroupResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetDescription(v string) *ListGroupResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetGroupId(v string) *ListGroupResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetIsDefault(v bool) *ListGroupResponseBodyGroups {
	s.IsDefault = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetModuleCnt(v int64) *ListGroupResponseBodyGroups {
	s.ModuleCnt = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetName(v string) *ListGroupResponseBodyGroups {
	s.Name = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetProjectId(v string) *ListGroupResponseBodyGroups {
	s.ProjectId = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetSceneTestingTaskCnt(v int64) *ListGroupResponseBodyGroups {
	s.SceneTestingTaskCnt = &v
	return s
}

func (s *ListGroupResponseBodyGroups) SetTags(v []*ListGroupResponseBodyGroupsTags) *ListGroupResponseBodyGroups {
	s.Tags = v
	return s
}

func (s *ListGroupResponseBodyGroups) SetTaskCnt(v int64) *ListGroupResponseBodyGroups {
	s.TaskCnt = &v
	return s
}

type ListGroupResponseBodyGroupsTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGroupResponseBodyGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponseBodyGroupsTags) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBodyGroupsTags) SetKey(v string) *ListGroupResponseBodyGroupsTags {
	s.Key = &v
	return s
}

func (s *ListGroupResponseBodyGroupsTags) SetValue(v string) *ListGroupResponseBodyGroupsTags {
	s.Value = &v
	return s
}

type ListGroupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponse) GoString() string {
	return s.String()
}

func (s *ListGroupResponse) SetHeaders(v map[string]*string) *ListGroupResponse {
	s.Headers = v
	return s
}

func (s *ListGroupResponse) SetStatusCode(v int32) *ListGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupResponse) SetBody(v *ListGroupResponseBody) *ListGroupResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Errored
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetTaskType(v string) *ListJobsRequest {
	s.TaskType = &v
	return s
}

type ListJobsResponseBody struct {
	Jobs []*ListJobsResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 882304F9-6DB1-5593-A719-33473D082B9C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 11
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponseBodyJobs struct {
	Config *ListJobsResponseBodyJobsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 2022-07-05T02:13:43Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	ElapsedTime *int64  `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// true
	IsPassAssertCheck *bool `json:"isPassAssertCheck,omitempty" xml:"isPassAssertCheck,omitempty"`
	// example:
	//
	// job-433aead756057fff9e4dca57b147c
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// example:
	//
	// Errored
	Status       *string                           `json:"status,omitempty" xml:"status,omitempty"`
	StatusDetail map[string]*JobsStatusDetailValue `json:"statusDetail,omitempty" xml:"statusDetail,omitempty"`
	// example:
	//
	// task-518855d9a058c1176866c2c3efb
	TaskId                   *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetConfig(v *ListJobsResponseBodyJobsConfig) *ListJobsResponseBodyJobs {
	s.Config = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetCreateTime(v string) *ListJobsResponseBodyJobs {
	s.CreateTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetDescription(v string) *ListJobsResponseBodyJobs {
	s.Description = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetElapsedTime(v int64) *ListJobsResponseBodyJobs {
	s.ElapsedTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetExecuteType(v string) *ListJobsResponseBodyJobs {
	s.ExecuteType = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetIsPassAssertCheck(v bool) *ListJobsResponseBodyJobs {
	s.IsPassAssertCheck = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetJobId(v string) *ListJobsResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatus(v string) *ListJobsResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStatusDetail(v map[string]*JobsStatusDetailValue) *ListJobsResponseBodyJobs {
	s.StatusDetail = v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTaskId(v string) *ListJobsResponseBodyJobs {
	s.TaskId = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetTerraformProviderVersion(v string) *ListJobsResponseBodyJobs {
	s.TerraformProviderVersion = &v
	return s
}

type ListJobsResponseBodyJobsConfig struct {
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
	// example:
	//
	// v4
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// {}
	ResourcesChanged *string `json:"resourcesChanged,omitempty" xml:"resourcesChanged,omitempty"`
	SubCommand       *string `json:"subCommand,omitempty" xml:"subCommand,omitempty"`
}

func (s ListJobsResponseBodyJobsConfig) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsConfig) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsConfig) SetIsDestroy(v bool) *ListJobsResponseBodyJobsConfig {
	s.IsDestroy = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetModuleVersion(v string) *ListJobsResponseBodyJobsConfig {
	s.ModuleVersion = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetResourcesChanged(v string) *ListJobsResponseBodyJobsConfig {
	s.ResourcesChanged = &v
	return s
}

func (s *ListJobsResponseBodyJobsConfig) SetSubCommand(v string) *ListJobsResponseBodyJobsConfig {
	s.SubCommand = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListModuleVersionRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListModuleVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *ListModuleVersionRequest) SetKeyword(v string) *ListModuleVersionRequest {
	s.Keyword = &v
	return s
}

func (s *ListModuleVersionRequest) SetPageNumber(v int32) *ListModuleVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModuleVersionRequest) SetPageSize(v int32) *ListModuleVersionRequest {
	s.PageSize = &v
	return s
}

type ListModuleVersionResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 792171BB-1A68-5148-8B9B-C7C728E1E98B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 6
	TotalCount *int32                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Versions   []*ListModuleVersionResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListModuleVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponseBody) SetPageNumber(v int32) *ListModuleVersionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetPageSize(v int32) *ListModuleVersionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetRequestId(v string) *ListModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetTotalCount(v int32) *ListModuleVersionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModuleVersionResponseBody) SetVersions(v []*ListModuleVersionResponseBodyVersions) *ListModuleVersionResponseBody {
	s.Versions = v
	return s
}

type ListModuleVersionResponseBodyVersions struct {
	// example:
	//
	// 2022-05-13T02:21:49Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// mod-55f1739d9050fffed3ec3a2c4a5e5
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v3
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
}

func (s ListModuleVersionResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListModuleVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponseBodyVersions) SetCreateTime(v string) *ListModuleVersionResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetDescription(v string) *ListModuleVersionResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetModuleId(v string) *ListModuleVersionResponseBodyVersions {
	s.ModuleId = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetModuleVersion(v string) *ListModuleVersionResponseBodyVersions {
	s.ModuleVersion = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetName(v string) *ListModuleVersionResponseBodyVersions {
	s.Name = &v
	return s
}

func (s *ListModuleVersionResponseBodyVersions) SetSourcePath(v string) *ListModuleVersionResponseBodyVersions {
	s.SourcePath = &v
	return s
}

type ListModuleVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModuleVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModuleVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModuleVersionResponse) GoString() string {
	return s.String()
}

func (s *ListModuleVersionResponse) SetHeaders(v map[string]*string) *ListModuleVersionResponse {
	s.Headers = v
	return s
}

func (s *ListModuleVersionResponse) SetStatusCode(v int32) *ListModuleVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModuleVersionResponse) SetBody(v *ListModuleVersionResponseBody) *ListModuleVersionResponse {
	s.Body = v
	return s
}

type ListModulesRequest struct {
	ExcludeModuleIds []*string `json:"excludeModuleIds,omitempty" xml:"excludeModuleIds,omitempty" type:"Repeated"`
	GroupId          *string   `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize  *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectId *string                  `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tag       []*ListModulesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListModulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModulesRequest) GoString() string {
	return s.String()
}

func (s *ListModulesRequest) SetExcludeModuleIds(v []*string) *ListModulesRequest {
	s.ExcludeModuleIds = v
	return s
}

func (s *ListModulesRequest) SetGroupId(v string) *ListModulesRequest {
	s.GroupId = &v
	return s
}

func (s *ListModulesRequest) SetKeyword(v string) *ListModulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListModulesRequest) SetPageNumber(v int32) *ListModulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesRequest) SetPageSize(v int32) *ListModulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesRequest) SetProjectId(v string) *ListModulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModulesRequest) SetTag(v []*ListModulesRequestTag) *ListModulesRequest {
	s.Tag = v
	return s
}

type ListModulesRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListModulesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListModulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListModulesRequestTag) SetKey(v string) *ListModulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListModulesRequestTag) SetValue(v string) *ListModulesRequestTag {
	s.Value = &v
	return s
}

type ListModulesShrinkRequest struct {
	ExcludeModuleIdsShrink *string `json:"excludeModuleIds,omitempty" xml:"excludeModuleIds,omitempty"`
	GroupId                *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize  *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListModulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModulesShrinkRequest) SetExcludeModuleIdsShrink(v string) *ListModulesShrinkRequest {
	s.ExcludeModuleIdsShrink = &v
	return s
}

func (s *ListModulesShrinkRequest) SetGroupId(v string) *ListModulesShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListModulesShrinkRequest) SetKeyword(v string) *ListModulesShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListModulesShrinkRequest) SetPageNumber(v int32) *ListModulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModulesShrinkRequest) SetPageSize(v int32) *ListModulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModulesShrinkRequest) SetProjectId(v string) *ListModulesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListModulesShrinkRequest) SetTagShrink(v string) *ListModulesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListModulesResponseBody struct {
	Modules []*ListModulesResponseBodyModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// C617E03B-3DD2-5F0C-A6CF-3028B499A2D5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 2790
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBody) SetModules(v []*ListModulesResponseBodyModules) *ListModulesResponseBody {
	s.Modules = v
	return s
}

func (s *ListModulesResponseBody) SetPageNumber(v int32) *ListModulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModulesResponseBody) SetPageSize(v int32) *ListModulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModulesResponseBody) SetRequestId(v string) *ListModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModulesResponseBody) SetTotalCount(v int32) *ListModulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListModulesResponseBodyModules struct {
	// example:
	//
	// 2022-01-30T02:14:16Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// example:
	//
	// description
	Description *string                                  `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *ListModulesResponseBodyModulesGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// v1
	LatestVersion *string                `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	// example:
	//
	// mod-518855d9a058cdbd3fd6951d59
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source       *string                `json:"source,omitempty" xml:"source,omitempty"`
	SourceConfig map[string]interface{} `json:"sourceConfig,omitempty" xml:"sourceConfig,omitempty"`
	// example:
	//
	// Created
	Status *string                               `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListModulesResponseBodyModulesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListModulesResponseBodyModules) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBodyModules) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModules) SetCreateTime(v string) *ListModulesResponseBodyModules {
	s.CreateTime = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetDeletionProtection(v bool) *ListModulesResponseBodyModules {
	s.DeletionProtection = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetDescription(v string) *ListModulesResponseBodyModules {
	s.Description = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetGroupInfo(v *ListModulesResponseBodyModulesGroupInfo) *ListModulesResponseBodyModules {
	s.GroupInfo = v
	return s
}

func (s *ListModulesResponseBodyModules) SetLatestVersion(v string) *ListModulesResponseBodyModules {
	s.LatestVersion = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetMeta(v map[string]interface{}) *ListModulesResponseBodyModules {
	s.Meta = v
	return s
}

func (s *ListModulesResponseBodyModules) SetModuleId(v string) *ListModulesResponseBodyModules {
	s.ModuleId = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetName(v string) *ListModulesResponseBodyModules {
	s.Name = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetSource(v string) *ListModulesResponseBodyModules {
	s.Source = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetSourceConfig(v map[string]interface{}) *ListModulesResponseBodyModules {
	s.SourceConfig = v
	return s
}

func (s *ListModulesResponseBodyModules) SetStatus(v string) *ListModulesResponseBodyModules {
	s.Status = &v
	return s
}

func (s *ListModulesResponseBodyModules) SetTags(v []*ListModulesResponseBodyModulesTags) *ListModulesResponseBodyModules {
	s.Tags = v
	return s
}

type ListModulesResponseBodyModulesGroupInfo struct {
	GroupId     *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	GroupName   *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	ProjectId   *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s ListModulesResponseBodyModulesGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBodyModulesGroupInfo) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetGroupId(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.GroupId = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetGroupName(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.GroupName = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetProjectId(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *ListModulesResponseBodyModulesGroupInfo) SetProjectName(v string) *ListModulesResponseBodyModulesGroupInfo {
	s.ProjectName = &v
	return s
}

type ListModulesResponseBodyModulesTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListModulesResponseBodyModulesTags) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponseBodyModulesTags) GoString() string {
	return s.String()
}

func (s *ListModulesResponseBodyModulesTags) SetKey(v string) *ListModulesResponseBodyModulesTags {
	s.Key = &v
	return s
}

func (s *ListModulesResponseBodyModulesTags) SetValue(v string) *ListModulesResponseBodyModulesTags {
	s.Value = &v
	return s
}

type ListModulesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListModulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListModulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModulesResponse) GoString() string {
	return s.String()
}

func (s *ListModulesResponse) SetHeaders(v map[string]*string) *ListModulesResponse {
	s.Headers = v
	return s
}

func (s *ListModulesResponse) SetStatusCode(v int32) *ListModulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModulesResponse) SetBody(v *ListModulesResponseBody) *ListModulesResponse {
	s.Body = v
	return s
}

type ListParameterSetRelationRequest struct {
	// This parameter is required.
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListParameterSetRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetRelationRequest) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationRequest) SetResourceId(v string) *ListParameterSetRelationRequest {
	s.ResourceId = &v
	return s
}

func (s *ListParameterSetRelationRequest) SetResourceType(v string) *ListParameterSetRelationRequest {
	s.ResourceType = &v
	return s
}

type ListParameterSetRelationResponseBody struct {
	ParameterSets []*ListParameterSetRelationResponseBodyParameterSets `json:"parameterSets,omitempty" xml:"parameterSets,omitempty" type:"Repeated"`
	RequestId     *string                                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount    *int32                                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListParameterSetRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponseBody) SetParameterSets(v []*ListParameterSetRelationResponseBodyParameterSets) *ListParameterSetRelationResponseBody {
	s.ParameterSets = v
	return s
}

func (s *ListParameterSetRelationResponseBody) SetRequestId(v string) *ListParameterSetRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterSetRelationResponseBody) SetTotalCount(v int32) *ListParameterSetRelationResponseBody {
	s.TotalCount = &v
	return s
}

type ListParameterSetRelationResponseBodyParameterSets struct {
	CreateTime     *string            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string            `json:"description,omitempty" xml:"description,omitempty"`
	Name           *string            `json:"name,omitempty" xml:"name,omitempty"`
	ParameterSetId *string            `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	Parameters     map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s ListParameterSetRelationResponseBodyParameterSets) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetRelationResponseBodyParameterSets) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetCreateTime(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetDescription(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Description = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetName(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Name = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetParameterSetId(v string) *ListParameterSetRelationResponseBodyParameterSets {
	s.ParameterSetId = &v
	return s
}

func (s *ListParameterSetRelationResponseBodyParameterSets) SetParameters(v map[string]*string) *ListParameterSetRelationResponseBodyParameterSets {
	s.Parameters = v
	return s
}

type ListParameterSetRelationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterSetRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterSetRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetRelationResponse) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationResponse) SetHeaders(v map[string]*string) *ListParameterSetRelationResponse {
	s.Headers = v
	return s
}

func (s *ListParameterSetRelationResponse) SetStatusCode(v int32) *ListParameterSetRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterSetRelationResponse) SetBody(v *ListParameterSetRelationResponseBody) *ListParameterSetRelationResponse {
	s.Body = v
	return s
}

type ListParameterSetsRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListParameterSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterSetsRequest) SetKeyword(v string) *ListParameterSetsRequest {
	s.Keyword = &v
	return s
}

func (s *ListParameterSetsRequest) SetPageNumber(v int32) *ListParameterSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParameterSetsRequest) SetPageSize(v int32) *ListParameterSetsRequest {
	s.PageSize = &v
	return s
}

type ListParameterSetsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize      *int32                                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ParameterSets []*ListParameterSetsResponseBodyParameterSets `json:"parameterSets,omitempty" xml:"parameterSets,omitempty" type:"Repeated"`
	// example:
	//
	// 4E188A8C-D77A-53F2-9578-E9AD8ABF2FA9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListParameterSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBody) SetPageNumber(v int32) *ListParameterSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetPageSize(v int32) *ListParameterSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetParameterSets(v []*ListParameterSetsResponseBodyParameterSets) *ListParameterSetsResponseBody {
	s.ParameterSets = v
	return s
}

func (s *ListParameterSetsResponseBody) SetRequestId(v string) *ListParameterSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetTotalCount(v int32) *ListParameterSetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListParameterSetsResponseBodyParameterSets struct {
	// example:
	//
	// 2022-05-14T10:05:19Z
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DeletionProtection *bool   `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 12
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pts-433aead756057ea135b21e89c
	ParameterSetId *string                                                   `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	Parameters     []*ListParameterSetsResponseBodyParameterSetsParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	RelationList   []*ListParameterSetsResponseBodyParameterSetsRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s ListParameterSetsResponseBodyParameterSets) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSets) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSets) SetCreateTime(v string) *ListParameterSetsResponseBodyParameterSets {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetDeletionProtection(v bool) *ListParameterSetsResponseBodyParameterSets {
	s.DeletionProtection = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetDescription(v string) *ListParameterSetsResponseBodyParameterSets {
	s.Description = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetName(v string) *ListParameterSetsResponseBodyParameterSets {
	s.Name = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetParameterSetId(v string) *ListParameterSetsResponseBodyParameterSets {
	s.ParameterSetId = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetParameters(v []*ListParameterSetsResponseBodyParameterSetsParameters) *ListParameterSetsResponseBodyParameterSets {
	s.Parameters = v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetRelationList(v []*ListParameterSetsResponseBodyParameterSetsRelationList) *ListParameterSetsResponseBodyParameterSets {
	s.RelationList = v
	return s
}

type ListParameterSetsResponseBodyParameterSetsParameters struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 111
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListParameterSetsResponseBodyParameterSetsParameters) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSetsParameters) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetName(v string) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Name = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetType(v string) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Type = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetValue(v interface{}) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Value = v
	return s
}

type ListParameterSetsResponseBodyParameterSetsRelationList struct {
	// example:
	//
	// 2022-06-09T03:46:18Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// task-433aead756057ffdf5326bf1e12ed
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListParameterSetsResponseBodyParameterSetsRelationList) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSetsRelationList) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetCreateTime(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetResourceId(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.ResourceId = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetResourceType(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.ResourceType = &v
	return s
}

type ListParameterSetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParameterSetsResponse) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponse) SetHeaders(v map[string]*string) *ListParameterSetsResponse {
	s.Headers = v
	return s
}

func (s *ListParameterSetsResponse) SetStatusCode(v int32) *ListParameterSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterSetsResponse) SetBody(v *ListParameterSetsResponseBody) *ListParameterSetsResponse {
	s.Body = v
	return s
}

type ListProjectRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Tag      []*ListProjectRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequest) GoString() string {
	return s.String()
}

func (s *ListProjectRequest) SetKeyword(v string) *ListProjectRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectRequest) SetPageNumber(v string) *ListProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectRequest) SetPageSize(v string) *ListProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectRequest) SetTag(v []*ListProjectRequestTag) *ListProjectRequest {
	s.Tag = v
	return s
}

type ListProjectRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListProjectRequestTag) GoString() string {
	return s.String()
}

func (s *ListProjectRequestTag) SetKey(v string) *ListProjectRequestTag {
	s.Key = &v
	return s
}

func (s *ListProjectRequestTag) SetValue(v string) *ListProjectRequestTag {
	s.Value = &v
	return s
}

type ListProjectShrinkRequest struct {
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListProjectShrinkRequest) SetKeyword(v string) *ListProjectShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectShrinkRequest) SetPageNumber(v string) *ListProjectShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectShrinkRequest) SetPageSize(v string) *ListProjectShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectShrinkRequest) SetTagShrink(v string) *ListProjectShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListProjectResponseBody struct {
	// example:
	//
	// 3
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Projects []*ListProjectResponseBodyProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetPageNumber(v int64) *ListProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectResponseBody) SetPageSize(v int64) *ListProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectResponseBody) SetProjects(v []*ListProjectResponseBodyProjects) *ListProjectResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectResponseBodyProjects struct {
	// example:
	//
	// 2022-05-10T10:08:34Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1234
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-148e7853433574fffe9fec72ed9b72
	ProjectId *string                                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tags      []*ListProjectResponseBodyProjectsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s ListProjectResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyProjects) SetCreateTime(v string) *ListProjectResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetDescription(v string) *ListProjectResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetName(v string) *ListProjectResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetProjectId(v string) *ListProjectResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetTags(v []*ListProjectResponseBodyProjectsTags) *ListProjectResponseBodyProjects {
	s.Tags = v
	return s
}

func (s *ListProjectResponseBodyProjects) SetTaskCnt(v int64) *ListProjectResponseBodyProjects {
	s.TaskCnt = &v
	return s
}

type ListProjectResponseBodyProjectsTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectResponseBodyProjectsTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponseBodyProjectsTags) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyProjectsTags) SetKey(v string) *ListProjectResponseBodyProjectsTags {
	s.Key = &v
	return s
}

func (s *ListProjectResponseBodyProjectsTags) SetValue(v string) *ListProjectResponseBodyProjectsTags {
	s.Value = &v
	return s
}

type ListProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectResponse) GoString() string {
	return s.String()
}

func (s *ListProjectResponse) SetHeaders(v map[string]*string) *ListProjectResponse {
	s.Headers = v
	return s
}

func (s *ListProjectResponse) SetStatusCode(v int32) *ListProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

type ListProjectBuildsRequest struct {
	GroupId            *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	PageNumber         *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize           *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectBuildAction *string `json:"projectBuildAction,omitempty" xml:"projectBuildAction,omitempty"`
}

func (s ListProjectBuildsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectBuildsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectBuildsRequest) SetGroupId(v string) *ListProjectBuildsRequest {
	s.GroupId = &v
	return s
}

func (s *ListProjectBuildsRequest) SetPageNumber(v string) *ListProjectBuildsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectBuildsRequest) SetPageSize(v string) *ListProjectBuildsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectBuildsRequest) SetProjectBuildAction(v string) *ListProjectBuildsRequest {
	s.ProjectBuildAction = &v
	return s
}

type ListProjectBuildsResponseBody struct {
	ProjectBuilds []*ListProjectBuildsResponseBodyProjectBuilds `json:"ProjectBuilds,omitempty" xml:"ProjectBuilds,omitempty" type:"Repeated"`
	PageNumber    *int64                                        `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize      *int64                                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 89154E16-FB0A-573D-8AF5-CF3F2FE28913
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProjectBuildsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectBuildsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectBuildsResponseBody) SetProjectBuilds(v []*ListProjectBuildsResponseBodyProjectBuilds) *ListProjectBuildsResponseBody {
	s.ProjectBuilds = v
	return s
}

func (s *ListProjectBuildsResponseBody) SetPageNumber(v int64) *ListProjectBuildsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectBuildsResponseBody) SetPageSize(v int64) *ListProjectBuildsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectBuildsResponseBody) SetRequestId(v string) *ListProjectBuildsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectBuildsResponseBody) SetTotalCount(v int32) *ListProjectBuildsResponseBody {
	s.TotalCount = &v
	return s
}

type ListProjectBuildsResponseBodyProjectBuilds struct {
	// example:
	//
	// 2022-08-26T02:10:48Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	EndTime    *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// apply
	ProjectBuildAction *string `json:"projectBuildAction,omitempty" xml:"projectBuildAction,omitempty"`
	// example:
	//
	// pb-433aead75605717728b6a27615f
	ProjectBuildId *string `json:"projectBuildId,omitempty" xml:"projectBuildId,omitempty"`
	// example:
	//
	// 1661221432088
	ProjectBuildIndex *int64 `json:"projectBuildIndex,omitempty" xml:"projectBuildIndex,omitempty"`
	// example:
	//
	// p-14e80de4866bf7ffed0bab6154d737
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// Pending
	Status                   *string `json:"status,omitempty" xml:"status,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// 1661221432088
	Timestamp       *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s ListProjectBuildsResponseBodyProjectBuilds) String() string {
	return tea.Prettify(s)
}

func (s ListProjectBuildsResponseBodyProjectBuilds) GoString() string {
	return s.String()
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetCreateTime(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.CreateTime = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetEndTime(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.EndTime = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetProjectBuildAction(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.ProjectBuildAction = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetProjectBuildId(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.ProjectBuildId = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetProjectBuildIndex(v int64) *ListProjectBuildsResponseBodyProjectBuilds {
	s.ProjectBuildIndex = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetProjectId(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.ProjectId = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetStatus(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.Status = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetTerraformProviderVersion(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.TerraformProviderVersion = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetTimestamp(v int64) *ListProjectBuildsResponseBodyProjectBuilds {
	s.Timestamp = &v
	return s
}

func (s *ListProjectBuildsResponseBodyProjectBuilds) SetTriggerStrategy(v string) *ListProjectBuildsResponseBodyProjectBuilds {
	s.TriggerStrategy = &v
	return s
}

type ListProjectBuildsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectBuildsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectBuildsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectBuildsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectBuildsResponse) SetHeaders(v map[string]*string) *ListProjectBuildsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectBuildsResponse) SetStatusCode(v int32) *ListProjectBuildsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectBuildsResponse) SetBody(v *ListProjectBuildsResponseBody) *ListProjectBuildsResponse {
	s.Body = v
	return s
}

type ListRabbitmqPublishersRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRabbitmqPublishersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRabbitmqPublishersRequest) GoString() string {
	return s.String()
}

func (s *ListRabbitmqPublishersRequest) SetKeyword(v string) *ListRabbitmqPublishersRequest {
	s.Keyword = &v
	return s
}

func (s *ListRabbitmqPublishersRequest) SetPageNumber(v int32) *ListRabbitmqPublishersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRabbitmqPublishersRequest) SetPageSize(v int32) *ListRabbitmqPublishersRequest {
	s.PageSize = &v
	return s
}

type ListRabbitmqPublishersResponseBody struct {
	Authorizations []*ListRabbitmqPublishersResponseBodyAuthorizations `json:"authorizations,omitempty" xml:"authorizations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 6599600A-20F6-556D-A15C-55EB9243BB24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 72
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRabbitmqPublishersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRabbitmqPublishersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRabbitmqPublishersResponseBody) SetAuthorizations(v []*ListRabbitmqPublishersResponseBodyAuthorizations) *ListRabbitmqPublishersResponseBody {
	s.Authorizations = v
	return s
}

func (s *ListRabbitmqPublishersResponseBody) SetPageNumber(v int32) *ListRabbitmqPublishersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBody) SetPageSize(v int32) *ListRabbitmqPublishersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBody) SetRequestId(v string) *ListRabbitmqPublishersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBody) SetTotalCount(v int32) *ListRabbitmqPublishersResponseBody {
	s.TotalCount = &v
	return s
}

type ListRabbitmqPublishersResponseBodyAuthorizations struct {
	// example:
	//
	// 2022-06-16T03:41:34Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// exchangeName
	ExchangeName *string `json:"exchangeName,omitempty" xml:"exchangeName,omitempty"`
	// example:
	//
	// TOPIC
	ExchangeType *string `json:"exchangeType,omitempty" xml:"exchangeType,omitempty"`
	HostName     *string `json:"hostName,omitempty" xml:"hostName,omitempty"`
	// example:
	//
	// MQ
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5671
	Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
	// example:
	//
	// mqp-3b6cb9fa4751a6e5cd2682246
	PublisherId *string `json:"publisherId,omitempty" xml:"publisherId,omitempty"`
	// example:
	//
	// MjoxODgwNzcwODY5MD***
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// example:
	//
	// virtualHost
	VirtualHost *string `json:"virtualHost,omitempty" xml:"virtualHost,omitempty"`
}

func (s ListRabbitmqPublishersResponseBodyAuthorizations) String() string {
	return tea.Prettify(s)
}

func (s ListRabbitmqPublishersResponseBodyAuthorizations) GoString() string {
	return s.String()
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetCreateTime(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.CreateTime = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetDescription(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.Description = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetExchangeName(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.ExchangeName = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetExchangeType(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.ExchangeType = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetHostName(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.HostName = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetName(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.Name = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetPort(v int64) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.Port = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetPublisherId(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.PublisherId = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetUserName(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.UserName = &v
	return s
}

func (s *ListRabbitmqPublishersResponseBodyAuthorizations) SetVirtualHost(v string) *ListRabbitmqPublishersResponseBodyAuthorizations {
	s.VirtualHost = &v
	return s
}

type ListRabbitmqPublishersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRabbitmqPublishersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRabbitmqPublishersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRabbitmqPublishersResponse) GoString() string {
	return s.String()
}

func (s *ListRabbitmqPublishersResponse) SetHeaders(v map[string]*string) *ListRabbitmqPublishersResponse {
	s.Headers = v
	return s
}

func (s *ListRabbitmqPublishersResponse) SetStatusCode(v int32) *ListRabbitmqPublishersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRabbitmqPublishersResponse) SetBody(v *ListRabbitmqPublishersResponseBody) *ListRabbitmqPublishersResponse {
	s.Body = v
	return s
}

type ListRamPolicyExportTaskVersionsRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListRamPolicyExportTaskVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTaskVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTaskVersionsRequest) SetPageNumber(v int32) *ListRamPolicyExportTaskVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsRequest) SetPageSize(v int32) *ListRamPolicyExportTaskVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsRequest) SetStatus(v string) *ListRamPolicyExportTaskVersionsRequest {
	s.Status = &v
	return s
}

type ListRamPolicyExportTaskVersionsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize                    *int32                                                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RamPolicyExportTaskVersions []*ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions `json:"ramPolicyExportTaskVersions,omitempty" xml:"ramPolicyExportTaskVersions,omitempty" type:"Repeated"`
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRamPolicyExportTaskVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTaskVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTaskVersionsResponseBody) SetPageNumber(v int32) *ListRamPolicyExportTaskVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBody) SetPageSize(v int32) *ListRamPolicyExportTaskVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBody) SetRamPolicyExportTaskVersions(v []*ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) *ListRamPolicyExportTaskVersionsResponseBody {
	s.RamPolicyExportTaskVersions = v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBody) SetRequestId(v string) *ListRamPolicyExportTaskVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBody) SetTotalCount(v int32) *ListRamPolicyExportTaskVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions struct {
	// example:
	//
	// 2022-07-04T02:11:29Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ElapsedTime *int64  `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// mod-51da8a8a36c166
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// rpe-437ffe0252e48f9286a
	RamPolicyExportTaskId *string `json:"ramPolicyExportTaskId,omitempty" xml:"ramPolicyExportTaskId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetCreateTime(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.CreateTime = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetElapsedTime(v int64) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.ElapsedTime = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetExportVersion(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.ExportVersion = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetModuleId(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.ModuleId = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetModuleVersion(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.ModuleVersion = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetRamPolicyExportTaskId(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.RamPolicyExportTaskId = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions) SetStatus(v string) *ListRamPolicyExportTaskVersionsResponseBodyRamPolicyExportTaskVersions {
	s.Status = &v
	return s
}

type ListRamPolicyExportTaskVersionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRamPolicyExportTaskVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRamPolicyExportTaskVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTaskVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTaskVersionsResponse) SetHeaders(v map[string]*string) *ListRamPolicyExportTaskVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponse) SetStatusCode(v int32) *ListRamPolicyExportTaskVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRamPolicyExportTaskVersionsResponse) SetBody(v *ListRamPolicyExportTaskVersionsResponseBody) *ListRamPolicyExportTaskVersionsResponse {
	s.Body = v
	return s
}

type ListRamPolicyExportTasksRequest struct {
	// example:
	//
	// vpc
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// mod-43143bd9145a5258
	ModuleId      *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListRamPolicyExportTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTasksRequest) SetKeyword(v string) *ListRamPolicyExportTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListRamPolicyExportTasksRequest) SetModuleId(v string) *ListRamPolicyExportTasksRequest {
	s.ModuleId = &v
	return s
}

func (s *ListRamPolicyExportTasksRequest) SetModuleVersion(v string) *ListRamPolicyExportTasksRequest {
	s.ModuleVersion = &v
	return s
}

func (s *ListRamPolicyExportTasksRequest) SetPageNumber(v int32) *ListRamPolicyExportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRamPolicyExportTasksRequest) SetPageSize(v int32) *ListRamPolicyExportTasksRequest {
	s.PageSize = &v
	return s
}

type ListRamPolicyExportTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize             *int32                                                      `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RamPolicyExportTasks []*ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks `json:"ramPolicyExportTasks,omitempty" xml:"ramPolicyExportTasks,omitempty" type:"Repeated"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 43
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListRamPolicyExportTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTasksResponseBody) SetPageNumber(v int32) *ListRamPolicyExportTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBody) SetPageSize(v int32) *ListRamPolicyExportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBody) SetRamPolicyExportTasks(v []*ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) *ListRamPolicyExportTasksResponseBody {
	s.RamPolicyExportTasks = v
	return s
}

func (s *ListRamPolicyExportTasksResponseBody) SetRequestId(v string) *ListRamPolicyExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBody) SetTotalCount(v int32) *ListRamPolicyExportTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks struct {
	// example:
	//
	// 2022-09-16T03:59:27Z
	CreateTime           *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CurrentPolicyStatus  *string `json:"currentPolicyStatus,omitempty" xml:"currentPolicyStatus,omitempty"`
	CurrentPolicyVersion *string `json:"currentPolicyVersion,omitempty" xml:"currentPolicyVersion,omitempty"`
	ElapsedTime          *int64  `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExportTime           *string `json:"exportTime,omitempty" xml:"exportTime,omitempty"`
	// example:
	//
	// mod-51fdfefa8788e82c9
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rpe-457ffe0252e48f9286a
	RamPolicyExportTaskId *string `json:"ramPolicyExportTaskId,omitempty" xml:"ramPolicyExportTaskId,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 12345678/default/rampolicytask/rpe-4399111870d7e6b0f11a
	TaskOutputPath *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
}

func (s ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetCreateTime(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.CreateTime = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetCurrentPolicyStatus(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.CurrentPolicyStatus = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetCurrentPolicyVersion(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.CurrentPolicyVersion = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetElapsedTime(v int64) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.ElapsedTime = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetExportTime(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.ExportTime = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetModuleId(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.ModuleId = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetModuleVersion(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.ModuleVersion = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetName(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.Name = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetRamPolicyExportTaskId(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.RamPolicyExportTaskId = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetStatus(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.Status = &v
	return s
}

func (s *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks) SetTaskOutputPath(v string) *ListRamPolicyExportTasksResponseBodyRamPolicyExportTasks {
	s.TaskOutputPath = &v
	return s
}

type ListRamPolicyExportTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRamPolicyExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRamPolicyExportTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRamPolicyExportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListRamPolicyExportTasksResponse) SetHeaders(v map[string]*string) *ListRamPolicyExportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListRamPolicyExportTasksResponse) SetStatusCode(v int32) *ListRamPolicyExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRamPolicyExportTasksResponse) SetBody(v *ListRamPolicyExportTasksResponseBody) *ListRamPolicyExportTasksResponse {
	s.Body = v
	return s
}

type ListResourceExportTaskVersionsRequest struct {
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	Keyword       *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	PageNumber    *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize      *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListResourceExportTaskVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsRequest) SetExportVersion(v string) *ListResourceExportTaskVersionsRequest {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetKeyword(v string) *ListResourceExportTaskVersionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetPageNumber(v int32) *ListResourceExportTaskVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetPageSize(v int32) *ListResourceExportTaskVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetStatus(v string) *ListResourceExportTaskVersionsRequest {
	s.Status = &v
	return s
}

type ListResourceExportTaskVersionsResponseBody struct {
	ExportTasks []*ListResourceExportTaskVersionsResponseBodyExportTasks `json:"exportTasks,omitempty" xml:"exportTasks,omitempty" type:"Repeated"`
	PageNumber  *int32                                                   `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32                                                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId   *string                                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount  *int32                                                   `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBody) SetExportTasks(v []*ListResourceExportTaskVersionsResponseBodyExportTasks) *ListResourceExportTaskVersionsResponseBody {
	s.ExportTasks = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetPageNumber(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetPageSize(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetRequestId(v string) *ListResourceExportTaskVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetTotalCount(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasks struct {
	CreateTime     *string                                                              `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string                                                              `json:"description,omitempty" xml:"description,omitempty"`
	ElapsedTime    *int64                                                               `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExcludeRules   []*ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	ExportTaskId   *string                                                              `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportToModule *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	ExportVersion  *string                                                              `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	FailedReason   *string                                                              `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	HasDestroy     *bool                                                                `json:"hasDestroy,omitempty" xml:"hasDestroy,omitempty"`
	IncludeRules   []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	Modules        []*ListResourceExportTaskVersionsResponseBodyExportTasksModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	Name           *string                                                              `json:"name,omitempty" xml:"name,omitempty"`
	Status         *string                                                              `json:"status,omitempty" xml:"status,omitempty"`
	Variables      []*ListResourceExportTaskVersionsResponseBodyExportTasksVariables    `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasks) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasks) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetCreateTime(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.CreateTime = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetDescription(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Description = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetElapsedTime(v int64) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ElapsedTime = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExcludeRules(v []*ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExcludeRules = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportTaskId(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportToModule(v *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportToModule = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportVersion(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetFailedReason(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.FailedReason = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetHasDestroy(v bool) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.HasDestroy = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetIncludeRules(v []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.IncludeRules = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetModules(v []*ListResourceExportTaskVersionsResponseBodyExportTasksModules) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Modules = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetName(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Name = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetStatus(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Status = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetVariables(v []*ListResourceExportTaskVersionsResponseBodyExportTasksVariables) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Variables = v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules) SetKey(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules) SetValues(v []*string) *ListResourceExportTaskVersionsResponseBodyExportTasksExcludeRules {
	s.Values = v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	StatePath  *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetSource(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.Source = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetSourcePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetStatePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.StatePath = &v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) SetKey(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) SetValues(v []*string) *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules {
	s.Values = v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasksModules struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksModules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksModules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetSource(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.Source = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetSourcePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetVersion(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.Version = &v
	return s
}

type ListResourceExportTaskVersionsResponseBodyExportTasksVariables struct {
	Properties   []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksVariables) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksVariables) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) SetProperties(v []*string) *ListResourceExportTaskVersionsResponseBodyExportTasksVariables {
	s.Properties = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) SetResourceType(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksVariables {
	s.ResourceType = &v
	return s
}

type ListResourceExportTaskVersionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExportTaskVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExportTaskVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponse) SetHeaders(v map[string]*string) *ListResourceExportTaskVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExportTaskVersionsResponse) SetStatusCode(v int32) *ListResourceExportTaskVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponse) SetBody(v *ListResourceExportTaskVersionsResponseBody) *ListResourceExportTaskVersionsResponse {
	s.Body = v
	return s
}

type ListResourceExportTasksRequest struct {
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	Keyword      *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	PageNumber   *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListResourceExportTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksRequest) SetExportTaskId(v string) *ListResourceExportTasksRequest {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetKeyword(v string) *ListResourceExportTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetPageNumber(v int32) *ListResourceExportTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTasksRequest) SetPageSize(v int32) *ListResourceExportTasksRequest {
	s.PageSize = &v
	return s
}

type ListResourceExportTasksResponseBody struct {
	ExportTasks []*ListResourceExportTasksResponseBodyExportTasks `json:"exportTasks,omitempty" xml:"exportTasks,omitempty" type:"Repeated"`
	PageNumber  *int32                                            `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32                                            `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId   *string                                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount  *int32                                            `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceExportTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBody) SetExportTasks(v []*ListResourceExportTasksResponseBodyExportTasks) *ListResourceExportTasksResponseBody {
	s.ExportTasks = v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetPageNumber(v int32) *ListResourceExportTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetPageSize(v int32) *ListResourceExportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetRequestId(v string) *ListResourceExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetTotalCount(v int32) *ListResourceExportTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceExportTasksResponseBodyExportTasks struct {
	CreateTime     *string                                                       `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description    *string                                                       `json:"description,omitempty" xml:"description,omitempty"`
	ElapsedTime    *int64                                                        `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	ExcludeRules   []*ListResourceExportTasksResponseBodyExportTasksExcludeRules `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	ExportStatus   *string                                                       `json:"exportStatus,omitempty" xml:"exportStatus,omitempty"`
	ExportTaskId   *string                                                       `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportToModule *ListResourceExportTasksResponseBodyExportTasksExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	ExportVersion  *string                                                       `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	HasDestroy     *bool                                                         `json:"hasDestroy,omitempty" xml:"hasDestroy,omitempty"`
	IncludeRules   []*ListResourceExportTasksResponseBodyExportTasksIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	Modules        []*ListResourceExportTasksResponseBodyExportTasksModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	Name           *string                                                       `json:"name,omitempty" xml:"name,omitempty"`
	Status         *string                                                       `json:"status,omitempty" xml:"status,omitempty"`
	Variables      []*ListResourceExportTasksResponseBodyExportTasksVariables    `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListResourceExportTasksResponseBodyExportTasks) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasks) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetCreateTime(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.CreateTime = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetDescription(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Description = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetElapsedTime(v int64) *ListResourceExportTasksResponseBodyExportTasks {
	s.ElapsedTime = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExcludeRules(v []*ListResourceExportTasksResponseBodyExportTasksExcludeRules) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExcludeRules = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportStatus(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportStatus = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportTaskId(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportToModule(v *ListResourceExportTasksResponseBodyExportTasksExportToModule) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportToModule = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportVersion(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetHasDestroy(v bool) *ListResourceExportTasksResponseBodyExportTasks {
	s.HasDestroy = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetIncludeRules(v []*ListResourceExportTasksResponseBodyExportTasksIncludeRules) *ListResourceExportTasksResponseBodyExportTasks {
	s.IncludeRules = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetModules(v []*ListResourceExportTasksResponseBodyExportTasksModules) *ListResourceExportTasksResponseBodyExportTasks {
	s.Modules = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetName(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Name = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetStatus(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Status = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetVariables(v []*ListResourceExportTasksResponseBodyExportTasksVariables) *ListResourceExportTasksResponseBodyExportTasks {
	s.Variables = v
	return s
}

type ListResourceExportTasksResponseBodyExportTasksExcludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTasksResponseBodyExportTasksExcludeRules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksExcludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksExcludeRules) SetKey(v string) *ListResourceExportTasksResponseBodyExportTasksExcludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExcludeRules) SetValues(v []*string) *ListResourceExportTasksResponseBodyExportTasksExcludeRules {
	s.Values = v
	return s
}

type ListResourceExportTasksResponseBodyExportTasksExportToModule struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	StatePath  *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s ListResourceExportTasksResponseBodyExportTasksExportToModule) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksExportToModule) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetSource(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.Source = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetSourcePath(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetStatePath(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.StatePath = &v
	return s
}

type ListResourceExportTasksResponseBodyExportTasksIncludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTasksResponseBodyExportTasksIncludeRules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksIncludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) SetKey(v string) *ListResourceExportTasksResponseBodyExportTasksIncludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) SetValues(v []*string) *ListResourceExportTasksResponseBodyExportTasksIncludeRules {
	s.Values = v
	return s
}

type ListResourceExportTasksResponseBodyExportTasksModules struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListResourceExportTasksResponseBodyExportTasksModules) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksModules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetSource(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.Source = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetSourcePath(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetVersion(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.Version = &v
	return s
}

type ListResourceExportTasksResponseBodyExportTasksVariables struct {
	Properties   []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListResourceExportTasksResponseBodyExportTasksVariables) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksVariables) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) SetProperties(v []*string) *ListResourceExportTasksResponseBodyExportTasksVariables {
	s.Properties = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) SetResourceType(v string) *ListResourceExportTasksResponseBodyExportTasksVariables {
	s.ResourceType = &v
	return s
}

type ListResourceExportTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExportTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponse) SetHeaders(v map[string]*string) *ListResourceExportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExportTasksResponse) SetStatusCode(v int32) *ListResourceExportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExportTasksResponse) SetBody(v *ListResourceExportTasksResponseBody) *ListResourceExportTasksResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// This parameter is required.
	SourceValue *string `json:"sourceValue,omitempty" xml:"sourceValue,omitempty"`
	// This parameter is required.
	SpecType *string `json:"specType,omitempty" xml:"specType,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetSourceType(v string) *ListResourcesRequest {
	s.SourceType = &v
	return s
}

func (s *ListResourcesRequest) SetSourceValue(v string) *ListResourcesRequest {
	s.SourceValue = &v
	return s
}

func (s *ListResourcesRequest) SetSpecType(v string) *ListResourcesRequest {
	s.SpecType = &v
	return s
}

type ListResourcesResponseBody struct {
	Resources  []*ListResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	PageNumber *int32                                `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TotalCount *int32                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetResources(v []*ListResourcesResponseBodyResources) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyResources struct {
	AccountId            *string                                   `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CreateTime           *string                                   `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DependsOnResourceIds []*string                                 `json:"dependsOnResourceIds,omitempty" xml:"dependsOnResourceIds,omitempty" type:"Repeated"`
	ProductCode          *string                                   `json:"productCode,omitempty" xml:"productCode,omitempty"`
	Properties           map[string]interface{}                    `json:"properties,omitempty" xml:"properties,omitempty"`
	PropertyVariables    map[string]interface{}                    `json:"propertyVariables,omitempty" xml:"propertyVariables,omitempty"`
	RegionId             *string                                   `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceArn          *string                                   `json:"resourceArn,omitempty" xml:"resourceArn,omitempty"`
	ResourceGroupId      *string                                   `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ResourceId           *string                                   `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	ResourceName         *string                                   `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	ResourceType         *string                                   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Status               *string                                   `json:"status,omitempty" xml:"status,omitempty"`
	Tags                 []*ListResourcesResponseBodyResourcesTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TerraformArn         *string                                   `json:"terraformArn,omitempty" xml:"terraformArn,omitempty"`
	TerraformCode        *string                                   `json:"terraformCode,omitempty" xml:"terraformCode,omitempty"`
	ZoneId               *string                                   `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ListResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResources) SetAccountId(v string) *ListResourcesResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetCreateTime(v string) *ListResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetDependsOnResourceIds(v []*string) *ListResourcesResponseBodyResources {
	s.DependsOnResourceIds = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProductCode(v string) *ListResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetProperties(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.Properties = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetPropertyVariables(v map[string]interface{}) *ListResourcesResponseBodyResources {
	s.PropertyVariables = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetRegionId(v string) *ListResourcesResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceArn(v string) *ListResourcesResponseBodyResources {
	s.ResourceArn = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceGroupId(v string) *ListResourcesResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceId(v string) *ListResourcesResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceName(v string) *ListResourcesResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetResourceType(v string) *ListResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetStatus(v string) *ListResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTags(v []*ListResourcesResponseBodyResourcesTags) *ListResourcesResponseBodyResources {
	s.Tags = v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTerraformArn(v string) *ListResourcesResponseBodyResources {
	s.TerraformArn = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetTerraformCode(v string) *ListResourcesResponseBodyResources {
	s.TerraformCode = &v
	return s
}

func (s *ListResourcesResponseBodyResources) SetZoneId(v string) *ListResourcesResponseBodyResources {
	s.ZoneId = &v
	return s
}

type ListResourcesResponseBodyResourcesTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListResourcesResponseBodyResourcesTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyResourcesTags) SetTagKey(v string) *ListResourcesResponseBodyResourcesTags {
	s.TagKey = &v
	return s
}

func (s *ListResourcesResponseBodyResourcesTags) SetTagValue(v string) *ListResourcesResponseBodyResourcesTags {
	s.TagValue = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	ExcludeTaskIds []*string `json:"excludeTaskIds,omitempty" xml:"excludeTaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// mod-1525e992f1b62139d1c437d64ae
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Status    *string                `json:"status,omitempty" xml:"status,omitempty"`
	Tag       []*ListTasksRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetExcludeTaskIds(v []*string) *ListTasksRequest {
	s.ExcludeTaskIds = v
	return s
}

func (s *ListTasksRequest) SetGroupId(v string) *ListTasksRequest {
	s.GroupId = &v
	return s
}

func (s *ListTasksRequest) SetKeyword(v string) *ListTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListTasksRequest) SetModuleId(v string) *ListTasksRequest {
	s.ModuleId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v string) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTag(v []*ListTasksRequestTag) *ListTasksRequest {
	s.Tag = v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

type ListTasksRequestTag struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTasksRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequestTag) GoString() string {
	return s.String()
}

func (s *ListTasksRequestTag) SetKey(v string) *ListTasksRequestTag {
	s.Key = &v
	return s
}

func (s *ListTasksRequestTag) SetValue(v string) *ListTasksRequestTag {
	s.Value = &v
	return s
}

type ListTasksShrinkRequest struct {
	ExcludeTaskIdsShrink *string `json:"excludeTaskIds,omitempty" xml:"excludeTaskIds,omitempty"`
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// mod-1525e992f1b62139d1c437d64ae
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) SetExcludeTaskIdsShrink(v string) *ListTasksShrinkRequest {
	s.ExcludeTaskIdsShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetGroupId(v string) *ListTasksShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetKeyword(v string) *ListTasksShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTasksShrinkRequest) SetModuleId(v string) *ListTasksShrinkRequest {
	s.ModuleId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageNumber(v int32) *ListTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageSize(v int32) *ListTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectId(v string) *ListTasksShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTagShrink(v string) *ListTasksShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskId(v string) *ListTasksShrinkRequest {
	s.TaskId = &v
	return s
}

type ListTasksResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 98610149-488B-5E48-B981-8D4CE1AF77CD
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Tasks     []*ListTasksResponseBodyTasks `json:"tasks,omitempty" xml:"tasks,omitempty" type:"Repeated"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// 2022-07-11T15:09:53Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// job-123asd
	CurrentJobId *string `json:"currentJobId,omitempty" xml:"currentJobId,omitempty"`
	// example:
	//
	// Pending
	CurrentJobStatus   *string                              `json:"currentJobStatus,omitempty" xml:"currentJobStatus,omitempty"`
	DeletionProtection *bool                                `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	GroupInfo          *ListTasksResponseBodyTasksGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// mod-518855d9a058c331e9c60bc0ce
	ModuleId   *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Running
	Status *string                           `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*ListTasksResponseBodyTasksTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// task-1525e992f1b621b0ca51647876e
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetAutoApply(v bool) *ListTasksResponseBodyTasks {
	s.AutoApply = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCreateTime(v string) *ListTasksResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentJobId(v string) *ListTasksResponseBodyTasks {
	s.CurrentJobId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentJobStatus(v string) *ListTasksResponseBodyTasks {
	s.CurrentJobStatus = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetDeletionProtection(v bool) *ListTasksResponseBodyTasks {
	s.DeletionProtection = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetGroupInfo(v *ListTasksResponseBodyTasksGroupInfo) *ListTasksResponseBodyTasks {
	s.GroupInfo = v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleId(v string) *ListTasksResponseBodyTasks {
	s.ModuleId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleName(v string) *ListTasksResponseBodyTasks {
	s.ModuleName = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetModuleVersion(v string) *ListTasksResponseBodyTasks {
	s.ModuleVersion = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetName(v string) *ListTasksResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTags(v []*ListTasksResponseBodyTasksTags) *ListTasksResponseBodyTasks {
	s.Tags = v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

type ListTasksResponseBodyTasksGroupInfo struct {
	// example:
	//
	// g-4267dcfbf1b6d1e0652bfbbe995
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// abc
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// example:
	//
	// p-433aead7560571cf1b2bfbbe92b
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// abc
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
}

func (s ListTasksResponseBodyTasksGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasksGroupInfo) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetGroupId(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.GroupId = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetGroupName(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.GroupName = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetProjectId(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyTasksGroupInfo) SetProjectName(v string) *ListTasksResponseBodyTasksGroupInfo {
	s.ProjectName = &v
	return s
}

type ListTasksResponseBodyTasksTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTasksResponseBodyTasksTags) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasksTags) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasksTags) SetKey(v string) *ListTasksResponseBodyTasksTags {
	s.Key = &v
	return s
}

func (s *ListTasksResponseBodyTasksTags) SetValue(v string) *ListTasksResponseBodyTasksTags {
	s.Value = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListTerraformProviderVersionsRequest struct {
	Keyword    *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Usage      *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListTerraformProviderVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTerraformProviderVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsRequest) SetKeyword(v string) *ListTerraformProviderVersionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetMaxResults(v int32) *ListTerraformProviderVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetNextToken(v string) *ListTerraformProviderVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTerraformProviderVersionsRequest) SetUsage(v string) *ListTerraformProviderVersionsRequest {
	s.Usage = &v
	return s
}

type ListTerraformProviderVersionsResponseBody struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 73588ebb-9d40-4660-a59f-764636ae6034
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 50
	TotalCount  *int32                                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	VerisonList []*string                                            `json:"verisonList,omitempty" xml:"verisonList,omitempty" type:"Repeated"`
	Versions    []*ListTerraformProviderVersionsResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListTerraformProviderVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTerraformProviderVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponseBody) SetMaxResults(v int32) *ListTerraformProviderVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetNextToken(v string) *ListTerraformProviderVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetRequestId(v string) *ListTerraformProviderVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetTotalCount(v int32) *ListTerraformProviderVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetVerisonList(v []*string) *ListTerraformProviderVersionsResponseBody {
	s.VerisonList = v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetVersions(v []*ListTerraformProviderVersionsResponseBodyVersions) *ListTerraformProviderVersionsResponseBody {
	s.Versions = v
	return s
}

type ListTerraformProviderVersionsResponseBodyVersions struct {
	PublishedTime *string `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListTerraformProviderVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListTerraformProviderVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetPublishedTime(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.PublishedTime = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetStatus(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.Status = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetVersion(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.Version = &v
	return s
}

type ListTerraformProviderVersionsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTerraformProviderVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTerraformProviderVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTerraformProviderVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponse) SetHeaders(v map[string]*string) *ListTerraformProviderVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTerraformProviderVersionsResponse) SetStatusCode(v int32) *ListTerraformProviderVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTerraformProviderVersionsResponse) SetBody(v *ListTerraformProviderVersionsResponseBody) *ListTerraformProviderVersionsResponse {
	s.Body = v
	return s
}

type OperateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dasd
	Comment  *string `json:"comment,omitempty" xml:"comment,omitempty"`
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s OperateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s OperateJobRequest) GoString() string {
	return s.String()
}

func (s *OperateJobRequest) SetComment(v string) *OperateJobRequest {
	s.Comment = &v
	return s
}

func (s *OperateJobRequest) SetTaskType(v string) *OperateJobRequest {
	s.TaskType = &v
	return s
}

type OperateJobResponseBody struct {
	// example:
	//
	// E602681C-A811-5787-9DC3-48BED7537071
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OperateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OperateJobResponseBody) GoString() string {
	return s.String()
}

func (s *OperateJobResponseBody) SetRequestId(v string) *OperateJobResponseBody {
	s.RequestId = &v
	return s
}

type OperateJobResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s OperateJobResponse) GoString() string {
	return s.String()
}

func (s *OperateJobResponse) SetHeaders(v map[string]*string) *OperateJobResponse {
	s.Headers = v
	return s
}

func (s *OperateJobResponse) SetStatusCode(v int32) *OperateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateJobResponse) SetBody(v *OperateJobResponseBody) *OperateJobResponse {
	s.Body = v
	return s
}

type RemoveResourceExportTaskVersionResponseBody struct {
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveResourceExportTaskVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveResourceExportTaskVersionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveResourceExportTaskVersionResponseBody) SetRequestId(v string) *RemoveResourceExportTaskVersionResponseBody {
	s.RequestId = &v
	return s
}

type RemoveResourceExportTaskVersionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveResourceExportTaskVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveResourceExportTaskVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveResourceExportTaskVersionResponse) GoString() string {
	return s.String()
}

func (s *RemoveResourceExportTaskVersionResponse) SetHeaders(v map[string]*string) *RemoveResourceExportTaskVersionResponse {
	s.Headers = v
	return s
}

func (s *RemoveResourceExportTaskVersionResponse) SetStatusCode(v int32) *RemoveResourceExportTaskVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveResourceExportTaskVersionResponse) SetBody(v *RemoveResourceExportTaskVersionResponseBody) *RemoveResourceExportTaskVersionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// This parameter is required.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// This parameter is required.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// This parameter is required.
	Tags []*TagResourcesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UpdateAuthorizationAttributeRequest struct {
	// This parameter is required.
	DueTime *string `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateAuthorizationAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributeRequest) SetDueTime(v string) *UpdateAuthorizationAttributeRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateAuthorizationAttributeRequest) SetName(v string) *UpdateAuthorizationAttributeRequest {
	s.Name = &v
	return s
}

type UpdateAuthorizationAttributeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 712C87FE-7C24-5298-B3E3-2BCB7AB9ED01
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAuthorizationAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributeResponseBody) SetRequestId(v string) *UpdateAuthorizationAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAuthorizationAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAuthorizationAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributeResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationAttributeResponse) SetStatusCode(v int32) *UpdateAuthorizationAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationAttributeResponse) SetBody(v *UpdateAuthorizationAttributeResponseBody) *UpdateAuthorizationAttributeResponse {
	s.Body = v
	return s
}

type UpdateExplorerTaskAttributeRequest struct {
	AutoApply                *bool   `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	ExplorerTaskName         *string `json:"explorerTaskName,omitempty" xml:"explorerTaskName,omitempty"`
	ModuleId                 *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	ModuleVersion            *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
}

func (s UpdateExplorerTaskAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExplorerTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateExplorerTaskAttributeRequest) SetAutoApply(v bool) *UpdateExplorerTaskAttributeRequest {
	s.AutoApply = &v
	return s
}

func (s *UpdateExplorerTaskAttributeRequest) SetExplorerTaskName(v string) *UpdateExplorerTaskAttributeRequest {
	s.ExplorerTaskName = &v
	return s
}

func (s *UpdateExplorerTaskAttributeRequest) SetModuleId(v string) *UpdateExplorerTaskAttributeRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateExplorerTaskAttributeRequest) SetModuleVersion(v string) *UpdateExplorerTaskAttributeRequest {
	s.ModuleVersion = &v
	return s
}

func (s *UpdateExplorerTaskAttributeRequest) SetTerraformProviderVersion(v string) *UpdateExplorerTaskAttributeRequest {
	s.TerraformProviderVersion = &v
	return s
}

type UpdateExplorerTaskAttributeResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateExplorerTaskAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExplorerTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExplorerTaskAttributeResponseBody) SetRequestId(v string) *UpdateExplorerTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExplorerTaskAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExplorerTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExplorerTaskAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExplorerTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateExplorerTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateExplorerTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateExplorerTaskAttributeResponse) SetStatusCode(v int32) *UpdateExplorerTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExplorerTaskAttributeResponse) SetBody(v *UpdateExplorerTaskAttributeResponseBody) *UpdateExplorerTaskAttributeResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// example:
	//
	// true
	AutoTrigger *bool `json:"autoTrigger,omitempty" xml:"autoTrigger,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	ForcedSetting *bool `json:"forcedSetting,omitempty" xml:"forcedSetting,omitempty"`
	// example:
	//
	// test
	Name                 *string                           `json:"name,omitempty" xml:"name,omitempty"`
	NotifyConfig         []*UpdateGroupRequestNotifyConfig `json:"notifyConfig,omitempty" xml:"notifyConfig,omitempty" type:"Repeated"`
	NotifyOperationTypes []*string                         `json:"notifyOperationTypes,omitempty" xml:"notifyOperationTypes,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	RamRole           *string   `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	ReportExportField []*string `json:"reportExportField,omitempty" xml:"reportExportField,omitempty" type:"Repeated"`
	// example:
	//
	// /
	ReportExportPath *string `json:"reportExportPath,omitempty" xml:"reportExportPath,omitempty"`
	// example:
	//
	// 1.183.0
	TerraformProviderVersion *string                            `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TriggerConfig            []*UpdateGroupRequestTriggerConfig `json:"triggerConfig,omitempty" xml:"triggerConfig,omitempty" type:"Repeated"`
	TriggerResourceType      []*string                          `json:"triggerResourceType,omitempty" xml:"triggerResourceType,omitempty" type:"Repeated"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetAutoDestroy(v bool) *UpdateGroupRequest {
	s.AutoDestroy = &v
	return s
}

func (s *UpdateGroupRequest) SetAutoTrigger(v bool) *UpdateGroupRequest {
	s.AutoTrigger = &v
	return s
}

func (s *UpdateGroupRequest) SetClientToken(v string) *UpdateGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateGroupRequest) SetDescription(v string) *UpdateGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupRequest) SetForcedSetting(v bool) *UpdateGroupRequest {
	s.ForcedSetting = &v
	return s
}

func (s *UpdateGroupRequest) SetName(v string) *UpdateGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateGroupRequest) SetNotifyConfig(v []*UpdateGroupRequestNotifyConfig) *UpdateGroupRequest {
	s.NotifyConfig = v
	return s
}

func (s *UpdateGroupRequest) SetNotifyOperationTypes(v []*string) *UpdateGroupRequest {
	s.NotifyOperationTypes = v
	return s
}

func (s *UpdateGroupRequest) SetRamRole(v string) *UpdateGroupRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateGroupRequest) SetReportExportField(v []*string) *UpdateGroupRequest {
	s.ReportExportField = v
	return s
}

func (s *UpdateGroupRequest) SetReportExportPath(v string) *UpdateGroupRequest {
	s.ReportExportPath = &v
	return s
}

func (s *UpdateGroupRequest) SetTerraformProviderVersion(v string) *UpdateGroupRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *UpdateGroupRequest) SetTriggerConfig(v []*UpdateGroupRequestTriggerConfig) *UpdateGroupRequest {
	s.TriggerConfig = v
	return s
}

func (s *UpdateGroupRequest) SetTriggerResourceType(v []*string) *UpdateGroupRequest {
	s.TriggerResourceType = v
	return s
}

type UpdateGroupRequestNotifyConfig struct {
	// example:
	//
	// /
	NotifyPath *string `json:"notifyPath,omitempty" xml:"notifyPath,omitempty"`
	// example:
	//
	// DingDing
	NotifyType *string `json:"notifyType,omitempty" xml:"notifyType,omitempty"`
}

func (s UpdateGroupRequestNotifyConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequestNotifyConfig) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequestNotifyConfig) SetNotifyPath(v string) *UpdateGroupRequestNotifyConfig {
	s.NotifyPath = &v
	return s
}

func (s *UpdateGroupRequestNotifyConfig) SetNotifyType(v string) *UpdateGroupRequestNotifyConfig {
	s.NotifyType = &v
	return s
}

type UpdateGroupRequestTriggerConfig struct {
	// example:
	//
	// Cron
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// 0 0 	- 	- 	- ？
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s UpdateGroupRequestTriggerConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequestTriggerConfig) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequestTriggerConfig) SetTriggerStrategy(v string) *UpdateGroupRequestTriggerConfig {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateGroupRequestTriggerConfig) SetTriggerValue(v string) *UpdateGroupRequestTriggerConfig {
	s.TriggerValue = &v
	return s
}

type UpdateGroupResponseBody struct {
	// example:
	//
	// 4EF5E823-AC0D-5CB5-8F18-1352455A488D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateModuleAttributeRequest struct {
	// example:
	//
	// test
	Description *string                                `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo   *UpdateModuleAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// OSS：
	//
	// "oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip"
	//
	// Registry：
	//
	// "alibaba/security-group/alicloud:2.1.0"
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/terraform.tfstate
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
	// example:
	//
	// Manual
	VersionStrategy *string `json:"versionStrategy,omitempty" xml:"versionStrategy,omitempty"`
}

func (s UpdateModuleAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeRequest) SetDescription(v string) *UpdateModuleAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetGroupInfo(v *UpdateModuleAttributeRequestGroupInfo) *UpdateModuleAttributeRequest {
	s.GroupInfo = v
	return s
}

func (s *UpdateModuleAttributeRequest) SetName(v string) *UpdateModuleAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetSource(v string) *UpdateModuleAttributeRequest {
	s.Source = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetSourcePath(v string) *UpdateModuleAttributeRequest {
	s.SourcePath = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetStatePath(v string) *UpdateModuleAttributeRequest {
	s.StatePath = &v
	return s
}

func (s *UpdateModuleAttributeRequest) SetVersionStrategy(v string) *UpdateModuleAttributeRequest {
	s.VersionStrategy = &v
	return s
}

type UpdateModuleAttributeRequestGroupInfo struct {
	GroupId   *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s UpdateModuleAttributeRequestGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleAttributeRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeRequestGroupInfo) SetGroupId(v string) *UpdateModuleAttributeRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *UpdateModuleAttributeRequestGroupInfo) SetProjectId(v string) *UpdateModuleAttributeRequestGroupInfo {
	s.ProjectId = &v
	return s
}

type UpdateModuleAttributeResponseBody struct {
	// example:
	//
	// CA05185E-6B90-5941-98D4-7212688AECC8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateModuleAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeResponseBody) SetRequestId(v string) *UpdateModuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateModuleAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateModuleAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateModuleAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateModuleAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateModuleAttributeResponse) SetHeaders(v map[string]*string) *UpdateModuleAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateModuleAttributeResponse) SetStatusCode(v int32) *UpdateModuleAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateModuleAttributeResponse) SetBody(v *UpdateModuleAttributeResponseBody) *UpdateModuleAttributeResponse {
	s.Body = v
	return s
}

type UpdateParameterSetAttributeRequest struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name       *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	Parameters []*UpdateParameterSetAttributeRequestParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
}

func (s UpdateParameterSetAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeRequest) SetDescription(v string) *UpdateParameterSetAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterSetAttributeRequest) SetName(v string) *UpdateParameterSetAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateParameterSetAttributeRequest) SetParameters(v []*UpdateParameterSetAttributeRequestParameters) *UpdateParameterSetAttributeRequest {
	s.Parameters = v
	return s
}

type UpdateParameterSetAttributeRequestParameters struct {
	// example:
	//
	// t
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// vpc-bp1mjm9exduos1bipw9x6
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateParameterSetAttributeRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterSetAttributeRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeRequestParameters) SetName(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Name = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) SetType(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Type = &v
	return s
}

func (s *UpdateParameterSetAttributeRequestParameters) SetValue(v string) *UpdateParameterSetAttributeRequestParameters {
	s.Value = &v
	return s
}

type UpdateParameterSetAttributeResponseBody struct {
	// example:
	//
	// 81CF7E18-D318-5670-9A4D-C08476BC4899
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateParameterSetAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterSetAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeResponseBody) SetRequestId(v string) *UpdateParameterSetAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateParameterSetAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParameterSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParameterSetAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateParameterSetAttributeResponse) SetHeaders(v map[string]*string) *UpdateParameterSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateParameterSetAttributeResponse) SetStatusCode(v int32) *UpdateParameterSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParameterSetAttributeResponse) SetBody(v *UpdateParameterSetAttributeResponseBody) *UpdateParameterSetAttributeResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetClientToken(v string) *UpdateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// C62888F6-254D-5589-BF05-0D9EE698C187
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponse) SetHeaders(v map[string]*string) *UpdateProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectResponse) SetStatusCode(v int32) *UpdateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateRabbitmqPublisherAttributeRequest struct {
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ExchangeTest
	ExchangeName *string `json:"exchangeName,omitempty" xml:"exchangeName,omitempty"`
	// example:
	//
	// TOPIC
	ExchangeType *string `json:"exchangeType,omitempty" xml:"exchangeType,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateRabbitmqPublisherAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRabbitmqPublisherAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRabbitmqPublisherAttributeRequest) SetDescription(v string) *UpdateRabbitmqPublisherAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateRabbitmqPublisherAttributeRequest) SetExchangeName(v string) *UpdateRabbitmqPublisherAttributeRequest {
	s.ExchangeName = &v
	return s
}

func (s *UpdateRabbitmqPublisherAttributeRequest) SetExchangeType(v string) *UpdateRabbitmqPublisherAttributeRequest {
	s.ExchangeType = &v
	return s
}

func (s *UpdateRabbitmqPublisherAttributeRequest) SetName(v string) *UpdateRabbitmqPublisherAttributeRequest {
	s.Name = &v
	return s
}

type UpdateRabbitmqPublisherAttributeResponseBody struct {
	// example:
	//
	// E2D0E863-1651-5E58-823F-B451C8C24615
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRabbitmqPublisherAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRabbitmqPublisherAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRabbitmqPublisherAttributeResponseBody) SetRequestId(v string) *UpdateRabbitmqPublisherAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRabbitmqPublisherAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRabbitmqPublisherAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRabbitmqPublisherAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRabbitmqPublisherAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRabbitmqPublisherAttributeResponse) SetHeaders(v map[string]*string) *UpdateRabbitmqPublisherAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRabbitmqPublisherAttributeResponse) SetStatusCode(v int32) *UpdateRabbitmqPublisherAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRabbitmqPublisherAttributeResponse) SetBody(v *UpdateRabbitmqPublisherAttributeResponseBody) *UpdateRabbitmqPublisherAttributeResponse {
	s.Body = v
	return s
}

type UpdateRamPolicyExportTaskAttributeRequest struct {
	AuthorizationAccountIds []*int64  `json:"authorizationAccountIds,omitempty" xml:"authorizationAccountIds,omitempty" type:"Repeated"`
	AuthorizationActions    []*string `json:"authorizationActions,omitempty" xml:"authorizationActions,omitempty" type:"Repeated"`
	AuthorizationRegionIds  []*string `json:"authorizationRegionIds,omitempty" xml:"authorizationRegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// mod-143574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ramName
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// Auto
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s UpdateRamPolicyExportTaskAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRamPolicyExportTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetAuthorizationAccountIds(v []*int64) *UpdateRamPolicyExportTaskAttributeRequest {
	s.AuthorizationAccountIds = v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetAuthorizationActions(v []*string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.AuthorizationActions = v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetAuthorizationRegionIds(v []*string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.AuthorizationRegionIds = v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetModuleId(v string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetModuleVersion(v string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.ModuleVersion = &v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetName(v string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetRamRole(v string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateRamPolicyExportTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

type UpdateRamPolicyExportTaskAttributeResponseBody struct {
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRamPolicyExportTaskAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRamPolicyExportTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRamPolicyExportTaskAttributeResponseBody) SetRequestId(v string) *UpdateRamPolicyExportTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRamPolicyExportTaskAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRamPolicyExportTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRamPolicyExportTaskAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRamPolicyExportTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRamPolicyExportTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateRamPolicyExportTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeResponse) SetStatusCode(v int32) *UpdateRamPolicyExportTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRamPolicyExportTaskAttributeResponse) SetBody(v *UpdateRamPolicyExportTaskAttributeResponseBody) *UpdateRamPolicyExportTaskAttributeResponse {
	s.Body = v
	return s
}

type UpdateResourceExportTaskAttributeRequest struct {
	// This parameter is required.
	ClientToken    *string                                                 `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	ConfigPath     *string                                                 `json:"configPath,omitempty" xml:"configPath,omitempty"`
	Description    *string                                                 `json:"description,omitempty" xml:"description,omitempty"`
	ExcludeRules   []*UpdateResourceExportTaskAttributeRequestExcludeRules `json:"excludeRules,omitempty" xml:"excludeRules,omitempty" type:"Repeated"`
	ExportToModule *UpdateResourceExportTaskAttributeRequestExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	IncludeRules   []*UpdateResourceExportTaskAttributeRequestIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// This parameter is required.
	Name                     *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	RamRole                  *string                                              `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	TerraformProviderVersion *string                                              `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	TerraformVersion         *string                                              `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	TriggerStrategy          *string                                              `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables                []*UpdateResourceExportTaskAttributeRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s UpdateResourceExportTaskAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequest) SetClientToken(v string) *UpdateResourceExportTaskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetConfigPath(v string) *UpdateResourceExportTaskAttributeRequest {
	s.ConfigPath = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetDescription(v string) *UpdateResourceExportTaskAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetExcludeRules(v []*UpdateResourceExportTaskAttributeRequestExcludeRules) *UpdateResourceExportTaskAttributeRequest {
	s.ExcludeRules = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetExportToModule(v *UpdateResourceExportTaskAttributeRequestExportToModule) *UpdateResourceExportTaskAttributeRequest {
	s.ExportToModule = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetIncludeRules(v []*UpdateResourceExportTaskAttributeRequestIncludeRules) *UpdateResourceExportTaskAttributeRequest {
	s.IncludeRules = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetName(v string) *UpdateResourceExportTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetRamRole(v string) *UpdateResourceExportTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTerraformProviderVersion(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TerraformProviderVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTerraformVersion(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TerraformVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateResourceExportTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequest) SetVariables(v []*UpdateResourceExportTaskAttributeRequestVariables) *UpdateResourceExportTaskAttributeRequest {
	s.Variables = v
	return s
}

type UpdateResourceExportTaskAttributeRequestExcludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateResourceExportTaskAttributeRequestExcludeRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestExcludeRules) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestExcludeRules) SetKey(v string) *UpdateResourceExportTaskAttributeRequestExcludeRules {
	s.Key = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExcludeRules) SetValues(v []*string) *UpdateResourceExportTaskAttributeRequestExcludeRules {
	s.Values = v
	return s
}

type UpdateResourceExportTaskAttributeRequestExportToModule struct {
	Source     *string `json:"source,omitempty" xml:"source,omitempty"`
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	StatePath  *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s UpdateResourceExportTaskAttributeRequestExportToModule) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestExportToModule) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetSource(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.Source = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetSourcePath(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.SourcePath = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestExportToModule) SetStatePath(v string) *UpdateResourceExportTaskAttributeRequestExportToModule {
	s.StatePath = &v
	return s
}

type UpdateResourceExportTaskAttributeRequestIncludeRules struct {
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s UpdateResourceExportTaskAttributeRequestIncludeRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestIncludeRules) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) SetKey(v string) *UpdateResourceExportTaskAttributeRequestIncludeRules {
	s.Key = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestIncludeRules) SetValues(v []*string) *UpdateResourceExportTaskAttributeRequestIncludeRules {
	s.Values = v
	return s
}

type UpdateResourceExportTaskAttributeRequestVariables struct {
	Properties   []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s UpdateResourceExportTaskAttributeRequestVariables) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeRequestVariables) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) SetProperties(v []*string) *UpdateResourceExportTaskAttributeRequestVariables {
	s.Properties = v
	return s
}

func (s *UpdateResourceExportTaskAttributeRequestVariables) SetResourceType(v string) *UpdateResourceExportTaskAttributeRequestVariables {
	s.ResourceType = &v
	return s
}

type UpdateResourceExportTaskAttributeResponseBody struct {
	ExportTaskId  *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateResourceExportTaskAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetExportTaskId(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.ExportTaskId = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetExportVersion(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.ExportVersion = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponseBody) SetRequestId(v string) *UpdateResourceExportTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceExportTaskAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceExportTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceExportTaskAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceExportTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceExportTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateResourceExportTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponse) SetStatusCode(v int32) *UpdateResourceExportTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceExportTaskAttributeResponse) SetBody(v *UpdateResourceExportTaskAttributeResponseBody) *UpdateResourceExportTaskAttributeResponse {
	s.Body = v
	return s
}

type UpdateTaskAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// example:
	//
	// demo
	Description     *string                              `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo       *UpdateTaskAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	InitModuleState *bool                                `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e7853433574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name               *string            `json:"name,omitempty" xml:"name,omitempty"`
	Parameters         map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	ProtectionStrategy []*string          `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {}
	RamRole                *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	SkipPropertyValidation *bool   `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// example:
	//
	// ******
	TriggerValue *string `json:"triggerValue,omitempty" xml:"triggerValue,omitempty"`
}

func (s UpdateTaskAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequest) SetAutoApply(v bool) *UpdateTaskAttributeRequest {
	s.AutoApply = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetAutoDestroy(v bool) *UpdateTaskAttributeRequest {
	s.AutoDestroy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetDescription(v string) *UpdateTaskAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetGroupInfo(v *UpdateTaskAttributeRequestGroupInfo) *UpdateTaskAttributeRequest {
	s.GroupInfo = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetInitModuleState(v bool) *UpdateTaskAttributeRequest {
	s.InitModuleState = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetModuleId(v string) *UpdateTaskAttributeRequest {
	s.ModuleId = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetModuleVersion(v string) *UpdateTaskAttributeRequest {
	s.ModuleVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetName(v string) *UpdateTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetParameters(v map[string]*string) *UpdateTaskAttributeRequest {
	s.Parameters = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetProtectionStrategy(v []*string) *UpdateTaskAttributeRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetRamRole(v string) *UpdateTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetSkipPropertyValidation(v bool) *UpdateTaskAttributeRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTerraformVersion(v string) *UpdateTaskAttributeRequest {
	s.TerraformVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTriggerValue(v string) *UpdateTaskAttributeRequest {
	s.TriggerValue = &v
	return s
}

type UpdateTaskAttributeRequestGroupInfo struct {
	// example:
	//
	// g-433aead7560571e66e31274ffd3
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// p-433aead75605713865c386cb9d
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s UpdateTaskAttributeRequestGroupInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAttributeRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetGroupId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetProjectId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.ProjectId = &v
	return s
}

type UpdateTaskAttributeResponseBody struct {
	// example:
	//
	// 17793D91-A26F-520D-A948-3452A45D15B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTaskAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeResponseBody) SetRequestId(v string) *UpdateTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTaskAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeResponse) SetHeaders(v map[string]*string) *UpdateTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskAttributeResponse) SetStatusCode(v int32) *UpdateTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskAttributeResponse) SetBody(v *UpdateTaskAttributeResponseBody) *UpdateTaskAttributeResponse {
	s.Body = v
	return s
}

type UpdateTaskPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken  *string                                `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	TaskPolicies []*UpdateTaskPolicyRequestTaskPolicies `json:"taskPolicies,omitempty" xml:"taskPolicies,omitempty" type:"Repeated"`
}

func (s UpdateTaskPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskPolicyRequest) SetClientToken(v string) *UpdateTaskPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTaskPolicyRequest) SetTaskPolicies(v []*UpdateTaskPolicyRequestTaskPolicies) *UpdateTaskPolicyRequest {
	s.TaskPolicies = v
	return s
}

type UpdateTaskPolicyRequestTaskPolicies struct {
	// example:
	//
	// 5
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// task-60f24b4eb47f1135b7b14ddbdfd
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SceneTestingTask
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateTaskPolicyRequestTaskPolicies) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPolicyRequestTaskPolicies) GoString() string {
	return s.String()
}

func (s *UpdateTaskPolicyRequestTaskPolicies) SetPriority(v string) *UpdateTaskPolicyRequestTaskPolicies {
	s.Priority = &v
	return s
}

func (s *UpdateTaskPolicyRequestTaskPolicies) SetTaskId(v string) *UpdateTaskPolicyRequestTaskPolicies {
	s.TaskId = &v
	return s
}

func (s *UpdateTaskPolicyRequestTaskPolicies) SetType(v string) *UpdateTaskPolicyRequestTaskPolicies {
	s.Type = &v
	return s
}

type UpdateTaskPolicyResponseBody struct {
	// example:
	//
	// 0D797DC3-FF04-5C21-81EB-92C7799512E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateTaskPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskPolicyResponseBody) SetRequestId(v string) *UpdateTaskPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTaskPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskPolicyResponse) SetHeaders(v map[string]*string) *UpdateTaskPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskPolicyResponse) SetStatusCode(v int32) *UpdateTaskPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskPolicyResponse) SetBody(v *UpdateTaskPolicyResponseBody) *UpdateTaskPolicyResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("iacservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 分组关联
//
// @param request - AssociateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateGroupResponse
func (client *Client) AssociateGroupWithOptions(groupId *string, request *AssociateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AssociateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/associate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分组关联
//
// @param request - AssociateGroupRequest
//
// @return AssociateGroupResponse
func (client *Client) AssociateGroup(groupId *string, request *AssociateGroupRequest) (_result *AssociateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AssociateGroupResponse{}
	_body, _err := client.AssociateGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将参数集关联资源
//
// @param request - AssociateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateParameterSetResponse
func (client *Client) AssociateParameterSetWithOptions(request *AssociateParameterSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AssociateParameterSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParameterSetIds)) {
		body["parameterSetIds"] = request.ParameterSetIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateParameterSet"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/operations/associate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateParameterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将参数集关联资源
//
// @param request - AssociateParameterSetRequest
//
// @return AssociateParameterSetResponse
func (client *Client) AssociateParameterSet(request *AssociateParameterSetRequest) (_result *AssociateParameterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AssociateParameterSetResponse{}
	_body, _err := client.AssociateParameterSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将消息发布者挂载到任务
//
// @param request - AttachRabbitmqPublisherRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachRabbitmqPublisherResponse
func (client *Client) AttachRabbitmqPublisherWithOptions(publisherId *string, request *AttachRabbitmqPublisherRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AttachRabbitmqPublisherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachRabbitmqPublisher"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers/" + tea.StringValue(openapiutil.GetEncodeParam(publisherId)) + "/attach"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachRabbitmqPublisherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将消息发布者挂载到任务
//
// @param request - AttachRabbitmqPublisherRequest
//
// @return AttachRabbitmqPublisherResponse
func (client *Client) AttachRabbitmqPublisher(publisherId *string, request *AttachRabbitmqPublisherRequest) (_result *AttachRabbitmqPublisherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AttachRabbitmqPublisherResponse{}
	_body, _err := client.AttachRabbitmqPublisherWithOptions(publisherId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消执行
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelProjectBuildResponse
func (client *Client) CancelProjectBuildWithOptions(projectId *string, projectBuildId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelProjectBuildResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelProjectBuild"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/build/" + tea.StringValue(openapiutil.GetEncodeParam(projectBuildId)) + "/cancel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelProjectBuildResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消执行
//
// @return CancelProjectBuildResponse
func (client *Client) CancelProjectBuild(projectId *string, projectBuildId *string) (_result *CancelProjectBuildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelProjectBuildResponse{}
	_body, _err := client.CancelProjectBuildWithOptions(projectId, projectBuildId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消RAM策略导出任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRamPolicyExportTaskResponse
func (client *Client) CancelRamPolicyExportTaskWithOptions(ramPolicyExportTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelRamPolicyExportTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRamPolicyExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId)) + "/cancel"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRamPolicyExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消RAM策略导出任务
//
// @return CancelRamPolicyExportTaskResponse
func (client *Client) CancelRamPolicyExportTask(ramPolicyExportTaskId *string) (_result *CancelRamPolicyExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelRamPolicyExportTaskResponse{}
	_body, _err := client.CancelRamPolicyExportTaskWithOptions(ramPolicyExportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消资源导出任务
//
// @param request - CancelResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelResourceExportTaskResponse
func (client *Client) CancelResourceExportTaskWithOptions(exportTaskId *string, request *CancelResourceExportTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelResourceExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelResourceExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/cancel/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelResourceExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消资源导出任务
//
// @param request - CancelResourceExportTaskRequest
//
// @return CancelResourceExportTaskResponse
func (client *Client) CancelResourceExportTask(exportTaskId *string, request *CancelResourceExportTaskRequest) (_result *CancelResourceExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelResourceExportTaskResponse{}
	_body, _err := client.CancelResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 校验资源名称
//
// @param request - CheckResourceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckResourceNameResponse
func (client *Client) CheckResourceNameWithOptions(request *CheckResourceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckResourceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["parentId"] = request.ParentId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResourceName"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/check/name"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckResourceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验资源名称
//
// @param request - CheckResourceNameRequest
//
// @return CheckResourceNameResponse
func (client *Client) CheckResourceName(request *CheckResourceNameRequest) (_result *CheckResourceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckResourceNameResponse{}
	_body, _err := client.CheckResourceNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆分组
//
// @param request - CloneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneGroupResponse
func (client *Client) CloneGroupWithOptions(groupId *string, request *CloneGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		body["resourceTypes"] = request.ResourceTypes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/clone"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆分组
//
// @param request - CloneGroupRequest
//
// @return CloneGroupResponse
func (client *Client) CloneGroup(groupId *string, request *CloneGroupRequest) (_result *CloneGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneGroupResponse{}
	_body, _err := client.CloneGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 克隆模版
//
// @param request - CloneModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloneModuleResponse
func (client *Client) CloneModuleWithOptions(request *CloneModuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloneModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleSource)) {
		body["moduleSource"] = request.ModuleSource
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloneModule"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/operations/clone"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloneModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 克隆模版
//
// @param request - CloneModuleRequest
//
// @return CloneModuleResponse
func (client *Client) CloneModule(request *CloneModuleRequest) (_result *CloneModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloneModuleResponse{}
	_body, _err := client.CloneModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建共享
//
// @param request - CreateAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthorizationResponse
func (client *Client) CreateAuthorizationWithOptions(request *CreateAuthorizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		body["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAuthorization"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/authorizations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建共享
//
// @param request - CreateAuthorizationRequest
//
// @return CreateAuthorizationResponse
func (client *Client) CreateAuthorization(request *CreateAuthorizationRequest) (_result *CreateAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAuthorizationResponse{}
	_body, _err := client.CreateAuthorizationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Explorer任务
//
// @param request - CreateExplorerTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExplorerTaskResponse
func (client *Client) CreateExplorerTaskWithOptions(request *CreateExplorerTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateExplorerTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InitModuleState)) {
		body["InitModuleState"] = request.InitModuleState
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVersion)) {
		body["TerraformVersion"] = request.TerraformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AutoApply)) {
		body["autoApply"] = request.AutoApply
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDestroy)) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InitModuleState)) {
		body["initModuleState"] = request.InitModuleState
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SkipPropertyValidation)) {
		body["skipPropertyValidation"] = request.SkipPropertyValidation
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerValue)) {
		body["triggerValue"] = request.TriggerValue
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExplorerTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/explorerTask"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExplorerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Explorer任务
//
// @param request - CreateExplorerTaskRequest
//
// @return CreateExplorerTaskResponse
func (client *Client) CreateExplorerTask(request *CreateExplorerTaskRequest) (_result *CreateExplorerTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExplorerTaskResponse{}
	_body, _err := client.CreateExplorerTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建分组
//
// @param request - CreateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDestroy)) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoTrigger)) {
		body["autoTrigger"] = request.AutoTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ForcedSetting)) {
		body["forcedSetting"] = request.ForcedSetting
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfig)) {
		body["notifyConfig"] = request.NotifyConfig
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyOperationTypes)) {
		body["notifyOperationTypes"] = request.NotifyOperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.ReportExportField)) {
		body["reportExportField"] = request.ReportExportField
	}

	if !tea.BoolValue(util.IsUnset(request.ReportExportPath)) {
		body["reportExportPath"] = request.ReportExportPath
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerResourceType)) {
		body["triggerResourceType"] = request.TriggerResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建分组
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建作业
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(taskId *string, request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteType)) {
		body["executeType"] = request.ExecuteType
	}

	if !tea.BoolValue(util.IsUnset(request.SubCommand)) {
		body["subCommand"] = request.SubCommand
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建作业
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(taskId *string, request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模版
//
// @param request - CreateModuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleResponse
func (client *Client) CreateModuleWithOptions(request *CreateModuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfo)) {
		body["groupInfo"] = request.GroupInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePath)) {
		body["sourcePath"] = request.SourcePath
	}

	if !tea.BoolValue(util.IsUnset(request.StatePath)) {
		body["statePath"] = request.StatePath
	}

	if !tea.BoolValue(util.IsUnset(request.VersionStrategy)) {
		body["versionStrategy"] = request.VersionStrategy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModule"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模版
//
// @param request - CreateModuleRequest
//
// @return CreateModuleResponse
func (client *Client) CreateModule(request *CreateModuleRequest) (_result *CreateModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModuleResponse{}
	_body, _err := client.CreateModuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建模版版本
//
// @param request - CreateModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModuleVersionResponse
func (client *Client) CreateModuleVersionWithOptions(moduleId *string, request *CreateModuleVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateModuleVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModuleVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId)) + "/versions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建模版版本
//
// @param request - CreateModuleVersionRequest
//
// @return CreateModuleVersionResponse
func (client *Client) CreateModuleVersion(moduleId *string, request *CreateModuleVersionRequest) (_result *CreateModuleVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModuleVersionResponse{}
	_body, _err := client.CreateModuleVersionWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建参数集
//
// @param request - CreateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterSetResponse
func (client *Client) CreateParameterSetWithOptions(request *CreateParameterSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateParameterSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateParameterSet"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateParameterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建参数集
//
// @param request - CreateParameterSetRequest
//
// @return CreateParameterSetResponse
func (client *Client) CreateParameterSet(request *CreateParameterSetRequest) (_result *CreateParameterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateParameterSetResponse{}
	_body, _err := client.CreateParameterSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 构建项目批次
//
// @param request - CreateProjectBuildRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectBuildResponse
func (client *Client) CreateProjectBuildWithOptions(projectId *string, request *CreateProjectBuildRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectBuildResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectBuildAction)) {
		body["projectBuildAction"] = request.ProjectBuildAction
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		body["taskIds"] = request.TaskIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPolicies)) {
		body["taskPolicies"] = request.TaskPolicies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProjectBuild"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/build"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectBuildResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 构建项目批次
//
// @param request - CreateProjectBuildRequest
//
// @return CreateProjectBuildResponse
func (client *Client) CreateProjectBuild(projectId *string, request *CreateProjectBuildRequest) (_result *CreateProjectBuildResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectBuildResponse{}
	_body, _err := client.CreateProjectBuildWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建消息发布者
//
// @param request - CreateRabbitmqPublisherRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRabbitmqPublisherResponse
func (client *Client) CreateRabbitmqPublisherWithOptions(request *CreateRabbitmqPublisherRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRabbitmqPublisherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["exchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeType)) {
		body["exchangeType"] = request.ExchangeType
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		body["hostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VirtualHost)) {
		body["virtualHost"] = request.VirtualHost
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRabbitmqPublisher"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRabbitmqPublisherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消息发布者
//
// @param request - CreateRabbitmqPublisherRequest
//
// @return CreateRabbitmqPublisherResponse
func (client *Client) CreateRabbitmqPublisher(request *CreateRabbitmqPublisherRequest) (_result *CreateRabbitmqPublisherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRabbitmqPublisherResponse{}
	_body, _err := client.CreateRabbitmqPublisherWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建RAM策略导出任务
//
// @param request - CreateRamPolicyExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRamPolicyExportTaskResponse
func (client *Client) CreateRamPolicyExportTaskWithOptions(request *CreateRamPolicyExportTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRamPolicyExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationAccountIds)) {
		body["authorizationAccountIds"] = request.AuthorizationAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationActions)) {
		body["authorizationActions"] = request.AuthorizationActions
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRegionIds)) {
		body["authorizationRegionIds"] = request.AuthorizationRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRamPolicyExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRamPolicyExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建RAM策略导出任务
//
// @param request - CreateRamPolicyExportTaskRequest
//
// @return CreateRamPolicyExportTaskResponse
func (client *Client) CreateRamPolicyExportTask(request *CreateRamPolicyExportTaskRequest) (_result *CreateRamPolicyExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRamPolicyExportTaskResponse{}
	_body, _err := client.CreateRamPolicyExportTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建导出任务
//
// @param request - CreateResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceExportTaskResponse
func (client *Client) CreateResourceExportTaskWithOptions(request *CreateResourceExportTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigPath)) {
		body["configPath"] = request.ConfigPath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeRules)) {
		body["excludeRules"] = request.ExcludeRules
	}

	if !tea.BoolValue(util.IsUnset(request.ExportToModule)) {
		body["exportToModule"] = request.ExportToModule
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeRules)) {
		body["includeRules"] = request.IncludeRules
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVersion)) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建导出任务
//
// @param request - CreateResourceExportTaskRequest
//
// @return CreateResourceExportTaskResponse
func (client *Client) CreateResourceExportTask(request *CreateResourceExportTaskRequest) (_result *CreateResourceExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceExportTaskResponse{}
	_body, _err := client.CreateResourceExportTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoApply)) {
		body["autoApply"] = request.AutoApply
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDestroy)) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfo)) {
		body["groupInfo"] = request.GroupInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InitModuleState)) {
		body["initModuleState"] = request.InitModuleState
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectionStrategy)) {
		body["protectionStrategy"] = request.ProtectionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.SkipPropertyValidation)) {
		body["skipPropertyValidation"] = request.SkipPropertyValidation
	}

	if !tea.BoolValue(util.IsUnset(request.TaskBackend)) {
		body["taskBackend"] = request.TaskBackend
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVersion)) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerValue)) {
		body["triggerValue"] = request.TriggerValue
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除共享
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthorizationResponse
func (client *Client) DeleteAuthorizationWithOptions(authorizationId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAuthorizationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAuthorization"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/authorizations/" + tea.StringValue(openapiutil.GetEncodeParam(authorizationId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除共享
//
// @return DeleteAuthorizationResponse
func (client *Client) DeleteAuthorization(authorizationId *string) (_result *DeleteAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAuthorizationResponse{}
	_body, _err := client.DeleteAuthorizationWithOptions(authorizationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(groupId *string) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除模版
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModuleResponse
func (client *Client) DeleteModuleWithOptions(moduleId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteModuleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModule"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除模版
//
// @return DeleteModuleResponse
func (client *Client) DeleteModule(moduleId *string) (_result *DeleteModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModuleResponse{}
	_body, _err := client.DeleteModuleWithOptions(moduleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除参数集
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterSetResponse
func (client *Client) DeleteParameterSetWithOptions(parameterSetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteParameterSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteParameterSet"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/" + tea.StringValue(openapiutil.GetEncodeParam(parameterSetId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteParameterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除参数集
//
// @return DeleteParameterSetResponse
func (client *Client) DeleteParameterSet(parameterSetId *string) (_result *DeleteParameterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteParameterSetResponse{}
	_body, _err := client.DeleteParameterSetWithOptions(parameterSetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(projectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(projectId *string) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消息发布者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRabbitmqPublisherResponse
func (client *Client) DeleteRabbitmqPublisherWithOptions(publisherId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRabbitmqPublisherResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRabbitmqPublisher"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers/" + tea.StringValue(openapiutil.GetEncodeParam(publisherId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRabbitmqPublisherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消息发布者
//
// @return DeleteRabbitmqPublisherResponse
func (client *Client) DeleteRabbitmqPublisher(publisherId *string) (_result *DeleteRabbitmqPublisherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRabbitmqPublisherResponse{}
	_body, _err := client.DeleteRabbitmqPublisherWithOptions(publisherId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除RAM策略导出任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRamPolicyExportTaskResponse
func (client *Client) DeleteRamPolicyExportTaskWithOptions(ramPolicyExportTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRamPolicyExportTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRamPolicyExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRamPolicyExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除RAM策略导出任务
//
// @return DeleteRamPolicyExportTaskResponse
func (client *Client) DeleteRamPolicyExportTask(ramPolicyExportTaskId *string) (_result *DeleteRamPolicyExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRamPolicyExportTaskResponse{}
	_body, _err := client.DeleteRamPolicyExportTaskWithOptions(ramPolicyExportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除RAM策略导出任务版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRamPolicyExportTaskVersionResponse
func (client *Client) DeleteRamPolicyExportTaskVersionWithOptions(ramPolicyExportTaskId *string, exportVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRamPolicyExportTaskVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRamPolicyExportTaskVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(exportVersion))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRamPolicyExportTaskVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除RAM策略导出任务版本
//
// @return DeleteRamPolicyExportTaskVersionResponse
func (client *Client) DeleteRamPolicyExportTaskVersion(ramPolicyExportTaskId *string, exportVersion *string) (_result *DeleteRamPolicyExportTaskVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRamPolicyExportTaskVersionResponse{}
	_body, _err := client.DeleteRamPolicyExportTaskVersionWithOptions(ramPolicyExportTaskId, exportVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceExportTaskResponse
func (client *Client) DeleteResourceExportTaskWithOptions(exportTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceExportTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @return DeleteResourceExportTaskResponse
func (client *Client) DeleteResourceExportTask(exportTaskId *string) (_result *DeleteResourceExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceExportTaskResponse{}
	_body, _err := client.DeleteResourceExportTaskWithOptions(exportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除场景化测试任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneTestingTaskResponse
func (client *Client) DeleteSceneTestingTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSceneTestingTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSceneTestingTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sceneTestingTask/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSceneTestingTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除场景化测试任务
//
// @return DeleteSceneTestingTaskResponse
func (client *Client) DeleteSceneTestingTask(taskId *string) (_result *DeleteSceneTestingTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSceneTestingTaskResponse{}
	_body, _err := client.DeleteSceneTestingTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除任务
//
// @return DeleteTaskResponse
func (client *Client) DeleteTask(taskId *string) (_result *DeleteTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将消息发布者从任务上卸载
//
// @param request - DetachRabbitmqPublisherRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachRabbitmqPublisherResponse
func (client *Client) DetachRabbitmqPublisherWithOptions(publisherId *string, request *DetachRabbitmqPublisherRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DetachRabbitmqPublisherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachRabbitmqPublisher"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers/" + tea.StringValue(openapiutil.GetEncodeParam(publisherId)) + "/detach"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachRabbitmqPublisherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将消息发布者从任务上卸载
//
// @param request - DetachRabbitmqPublisherRequest
//
// @return DetachRabbitmqPublisherResponse
func (client *Client) DetachRabbitmqPublisher(publisherId *string, request *DetachRabbitmqPublisherRequest) (_result *DetachRabbitmqPublisherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DetachRabbitmqPublisherResponse{}
	_body, _err := client.DetachRabbitmqPublisherWithOptions(publisherId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消关联分组
//
// @param request - DissociateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateGroupResponse
func (client *Client) DissociateGroupWithOptions(projectId *string, groupId *string, request *DissociateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DissociateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/dissociate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关联分组
//
// @param request - DissociateGroupRequest
//
// @return DissociateGroupResponse
func (client *Client) DissociateGroup(projectId *string, groupId *string, request *DissociateGroupRequest) (_result *DissociateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DissociateGroupResponse{}
	_body, _err := client.DissociateGroupWithOptions(projectId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解除参数集关联资源关系
//
// @param request - DissociateParameterSetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateParameterSetResponse
func (client *Client) DissociateParameterSetWithOptions(request *DissociateParameterSetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DissociateParameterSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParameterSetIds)) {
		body["parameterSetIds"] = request.ParameterSetIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DissociateParameterSet"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/operations/dissociate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DissociateParameterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除参数集关联资源关系
//
// @param request - DissociateParameterSetRequest
//
// @return DissociateParameterSetResponse
func (client *Client) DissociateParameterSet(request *DissociateParameterSetRequest) (_result *DissociateParameterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DissociateParameterSetResponse{}
	_body, _err := client.DissociateParameterSetWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行RAM策略导出任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteRamPolicyExportTaskResponse
func (client *Client) ExecuteRamPolicyExportTaskWithOptions(ramPolicyExportTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteRamPolicyExportTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteRamPolicyExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId)) + "/execute"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteRamPolicyExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行RAM策略导出任务
//
// @return ExecuteRamPolicyExportTaskResponse
func (client *Client) ExecuteRamPolicyExportTask(ramPolicyExportTaskId *string) (_result *ExecuteRamPolicyExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteRamPolicyExportTaskResponse{}
	_body, _err := client.ExecuteRamPolicyExportTaskWithOptions(ramPolicyExportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行资源导出任务
//
// @param request - ExecuteResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteResourceExportTaskResponse
func (client *Client) ExecuteResourceExportTaskWithOptions(exportTaskId *string, request *ExecuteResourceExportTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteResourceExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteResourceExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/execute/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteResourceExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行资源导出任务
//
// @param request - ExecuteResourceExportTaskRequest
//
// @return ExecuteResourceExportTaskResponse
func (client *Client) ExecuteResourceExportTask(exportTaskId *string, request *ExecuteResourceExportTaskRequest) (_result *ExecuteResourceExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteResourceExportTaskResponse{}
	_body, _err := client.ExecuteResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Explorer任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExplorerTaskResponse
func (client *Client) GetExplorerTaskWithOptions(explorerTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetExplorerTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetExplorerTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/explorerTask/" + tea.StringValue(openapiutil.GetEncodeParam(explorerTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExplorerTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Explorer任务详情
//
// @return GetExplorerTaskResponse
func (client *Client) GetExplorerTask(explorerTaskId *string) (_result *GetExplorerTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExplorerTaskResponse{}
	_body, _err := client.GetExplorerTaskWithOptions(explorerTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分组
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组
//
// @return GetGroupResponse
func (client *Client) GetGroup(groupId *string) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 作业详情
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(taskId *string, jobId *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业详情
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(taskId *string, jobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(taskId, jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模版详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleResponse
func (client *Client) GetModuleWithOptions(moduleId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModuleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModule"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模版详情
//
// @return GetModuleResponse
func (client *Client) GetModule(moduleId *string) (_result *GetModuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModuleResponse{}
	_body, _err := client.GetModuleWithOptions(moduleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模版版本详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModuleVersionResponse
func (client *Client) GetModuleVersionWithOptions(moduleId *string, moduleVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModuleVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModuleVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(moduleVersion))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模版版本详情
//
// @return GetModuleVersionResponse
func (client *Client) GetModuleVersion(moduleId *string, moduleVersion *string) (_result *GetModuleVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModuleVersionResponse{}
	_body, _err := client.GetModuleVersionWithOptions(moduleId, moduleVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 参数集详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParameterSetResponse
func (client *Client) GetParameterSetWithOptions(parameterSetId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetParameterSetResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetParameterSet"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/" + tea.StringValue(openapiutil.GetEncodeParam(parameterSetId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetParameterSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 参数集详情
//
// @return GetParameterSetResponse
func (client *Client) GetParameterSet(parameterSetId *string) (_result *GetParameterSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetParameterSetResponse{}
	_body, _err := client.GetParameterSetWithOptions(parameterSetId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(projectId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目
//
// @return GetProjectResponse
func (client *Client) GetProject(projectId *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(projectId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目批次概览
//
// @param request - GetProjectBuildContextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectBuildContextResponse
func (client *Client) GetProjectBuildContextWithOptions(projectId *string, projectBuildId *string, request *GetProjectBuildContextRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectBuildContextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsPassAssertCheck)) {
		query["isPassAssertCheck"] = request.IsPassAssertCheck
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectBuildContext"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/build/" + tea.StringValue(openapiutil.GetEncodeParam(projectBuildId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectBuildContextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目批次概览
//
// @param request - GetProjectBuildContextRequest
//
// @return GetProjectBuildContextResponse
func (client *Client) GetProjectBuildContext(projectId *string, projectBuildId *string, request *GetProjectBuildContextRequest) (_result *GetProjectBuildContextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectBuildContextResponse{}
	_body, _err := client.GetProjectBuildContextWithOptions(projectId, projectBuildId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取消息发布者详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRabbitmqPublisherResponse
func (client *Client) GetRabbitmqPublisherWithOptions(publisherId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRabbitmqPublisherResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRabbitmqPublisher"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers/" + tea.StringValue(openapiutil.GetEncodeParam(publisherId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRabbitmqPublisherResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取消息发布者详情
//
// @return GetRabbitmqPublisherResponse
func (client *Client) GetRabbitmqPublisher(publisherId *string) (_result *GetRabbitmqPublisherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRabbitmqPublisherResponse{}
	_body, _err := client.GetRabbitmqPublisherWithOptions(publisherId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRamPolicyExportTaskResponse
func (client *Client) GetRamPolicyExportTaskWithOptions(ramPolicyExportTaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRamPolicyExportTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRamPolicyExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRamPolicyExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务详情
//
// @return GetRamPolicyExportTaskResponse
func (client *Client) GetRamPolicyExportTask(ramPolicyExportTaskId *string) (_result *GetRamPolicyExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRamPolicyExportTaskResponse{}
	_body, _err := client.GetRamPolicyExportTaskWithOptions(ramPolicyExportTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务版本详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRamPolicyExportTaskVersionResponse
func (client *Client) GetRamPolicyExportTaskVersionWithOptions(ramPolicyExportTaskId *string, exportVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRamPolicyExportTaskVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRamPolicyExportTaskVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(exportVersion))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRamPolicyExportTaskVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务版本详情
//
// @return GetRamPolicyExportTaskVersionResponse
func (client *Client) GetRamPolicyExportTaskVersion(ramPolicyExportTaskId *string, exportVersion *string) (_result *GetRamPolicyExportTaskVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRamPolicyExportTaskVersionResponse{}
	_body, _err := client.GetRamPolicyExportTaskVersionWithOptions(ramPolicyExportTaskId, exportVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询导出任务详情
//
// @param request - GetResourceExportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceExportTaskResponse
func (client *Client) GetResourceExportTaskWithOptions(exportTaskId *string, request *GetResourceExportTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourceExportTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportVersion)) {
		query["exportVersion"] = request.ExportVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceExportTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导出任务详情
//
// @param request - GetResourceExportTaskRequest
//
// @return GetResourceExportTaskResponse
func (client *Client) GetResourceExportTask(exportTaskId *string, request *GetResourceExportTaskRequest) (_result *GetResourceExportTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceExportTaskResponse{}
	_body, _err := client.GetResourceExportTaskWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务详情
//
// @return GetTaskResponse
func (client *Client) GetTask(taskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分组优先级详情
//
// @param request - GetTaskPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskPolicyResponse
func (client *Client) GetTaskPolicyWithOptions(groupId *string, request *GetTaskPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskPolicy"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/task/policy/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组优先级详情
//
// @param request - GetTaskPolicyRequest
//
// @return GetTaskPolicyResponse
func (client *Client) GetTaskPolicy(groupId *string, request *GetTaskPolicyRequest) (_result *GetTaskPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskPolicyResponse{}
	_body, _err := client.GetTaskPolicyWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取共享列表
//
// @param request - ListAuthorizationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizationsResponse
func (client *Client) ListAuthorizationsWithOptions(request *ListAuthorizationsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAuthorizationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationId)) {
		query["authorizationId"] = request.AuthorizationId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationType)) {
		query["authorizationType"] = request.AuthorizationType
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAuthorizations"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/authorizations"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取共享列表
//
// @param request - ListAuthorizationsRequest
//
// @return ListAuthorizationsResponse
func (client *Client) ListAuthorizations(request *ListAuthorizationsRequest) (_result *ListAuthorizationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAuthorizationsResponse{}
	_body, _err := client.ListAuthorizationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// terraform版本
//
// @param request - ListAvailableTerraformVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableTerraformVersionsResponse
func (client *Client) ListAvailableTerraformVersionsWithOptions(request *ListAvailableTerraformVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAvailableTerraformVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["keyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableTerraformVersions"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/version/terraform"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableTerraformVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// terraform版本
//
// @param request - ListAvailableTerraformVersionsRequest
//
// @return ListAvailableTerraformVersionsResponse
func (client *Client) ListAvailableTerraformVersions(request *ListAvailableTerraformVersionsRequest) (_result *ListAvailableTerraformVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableTerraformVersionsResponse{}
	_body, _err := client.ListAvailableTerraformVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举Explorer任务
//
// @param request - ListExplorerTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExplorerTasksResponse
func (client *Client) ListExplorerTasksWithOptions(explorerTaskId *string, request *ListExplorerTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListExplorerTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["maxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExplorerTasks"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/explorerTask"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExplorerTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举Explorer任务
//
// @param request - ListExplorerTasksRequest
//
// @return ListExplorerTasksResponse
func (client *Client) ListExplorerTasks(explorerTaskId *string, request *ListExplorerTasksRequest) (_result *ListExplorerTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExplorerTasksResponse{}
	_body, _err := client.ListExplorerTasksWithOptions(explorerTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询分组列表
//
// @param tmpReq - ListGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupResponse
func (client *Client) ListGroupWithOptions(tmpReq *ListGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询分组列表
//
// @param request - ListGroupRequest
//
// @return ListGroupResponse
func (client *Client) ListGroup(request *ListGroupRequest) (_result *ListGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupResponse{}
	_body, _err := client.ListGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 作业列表
//
// @param request - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(taskId *string, request *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 作业列表
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(taskId *string, request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 模版版本列表
//
// @param request - ListModuleVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModuleVersionResponse
func (client *Client) ListModuleVersionWithOptions(moduleId *string, request *ListModuleVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModuleVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModuleVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModuleVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 模版版本列表
//
// @param request - ListModuleVersionRequest
//
// @return ListModuleVersionResponse
func (client *Client) ListModuleVersion(moduleId *string, request *ListModuleVersionRequest) (_result *ListModuleVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModuleVersionResponse{}
	_body, _err := client.ListModuleVersionWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举模版
//
// @param tmpReq - ListModulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModulesResponse
func (client *Client) ListModulesWithOptions(tmpReq *ListModulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListModulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExcludeModuleIds)) {
		request.ExcludeModuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeModuleIds, tea.String("excludeModuleIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeModuleIdsShrink)) {
		query["excludeModuleIds"] = request.ExcludeModuleIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModules"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举模版
//
// @param request - ListModulesRequest
//
// @return ListModulesResponse
func (client *Client) ListModules(request *ListModulesRequest) (_result *ListModulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModulesResponse{}
	_body, _err := client.ListModulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关联到资源的参数集列表
//
// @param request - ListParameterSetRelationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterSetRelationResponse
func (client *Client) ListParameterSetRelationWithOptions(request *ListParameterSetRelationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListParameterSetRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParameterSetRelation"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/operations/relation"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParameterSetRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关联到资源的参数集列表
//
// @param request - ListParameterSetRelationRequest
//
// @return ListParameterSetRelationResponse
func (client *Client) ListParameterSetRelation(request *ListParameterSetRelationRequest) (_result *ListParameterSetRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListParameterSetRelationResponse{}
	_body, _err := client.ListParameterSetRelationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 参数集列表
//
// @param request - ListParameterSetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterSetsResponse
func (client *Client) ListParameterSetsWithOptions(request *ListParameterSetsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListParameterSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParameterSets"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParameterSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 参数集列表
//
// @param request - ListParameterSetsRequest
//
// @return ListParameterSetsResponse
func (client *Client) ListParameterSets(request *ListParameterSetsRequest) (_result *ListParameterSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListParameterSetsResponse{}
	_body, _err := client.ListParameterSetsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目列表
//
// @param tmpReq - ListProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithOptions(tmpReq *ListProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProject"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目列表
//
// @param request - ListProjectRequest
//
// @return ListProjectResponse
func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目批次列表
//
// @param request - ListProjectBuildsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectBuildsResponse
func (client *Client) ListProjectBuildsWithOptions(projectId *string, request *ListProjectBuildsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectBuildsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectBuildAction)) {
		query["projectBuildAction"] = request.ProjectBuildAction
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectBuilds"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId)) + "/build"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectBuildsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目批次列表
//
// @param request - ListProjectBuildsRequest
//
// @return ListProjectBuildsResponse
func (client *Client) ListProjectBuilds(projectId *string, request *ListProjectBuildsRequest) (_result *ListProjectBuildsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectBuildsResponse{}
	_body, _err := client.ListProjectBuildsWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取消息发布者列表
//
// @param request - ListRabbitmqPublishersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRabbitmqPublishersResponse
func (client *Client) ListRabbitmqPublishersWithOptions(request *ListRabbitmqPublishersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRabbitmqPublishersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRabbitmqPublishers"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRabbitmqPublishersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取消息发布者列表
//
// @param request - ListRabbitmqPublishersRequest
//
// @return ListRabbitmqPublishersResponse
func (client *Client) ListRabbitmqPublishers(request *ListRabbitmqPublishersRequest) (_result *ListRabbitmqPublishersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRabbitmqPublishersResponse{}
	_body, _err := client.ListRabbitmqPublishersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务版本列表
//
// @param request - ListRamPolicyExportTaskVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRamPolicyExportTaskVersionsResponse
func (client *Client) ListRamPolicyExportTaskVersionsWithOptions(ramPolicyExportTaskId *string, request *ListRamPolicyExportTaskVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRamPolicyExportTaskVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRamPolicyExportTaskVersions"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId)) + "/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRamPolicyExportTaskVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务版本列表
//
// @param request - ListRamPolicyExportTaskVersionsRequest
//
// @return ListRamPolicyExportTaskVersionsResponse
func (client *Client) ListRamPolicyExportTaskVersions(ramPolicyExportTaskId *string, request *ListRamPolicyExportTaskVersionsRequest) (_result *ListRamPolicyExportTaskVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRamPolicyExportTaskVersionsResponse{}
	_body, _err := client.ListRamPolicyExportTaskVersionsWithOptions(ramPolicyExportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务列表
//
// @param request - ListRamPolicyExportTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRamPolicyExportTasksResponse
func (client *Client) ListRamPolicyExportTasksWithOptions(request *ListRamPolicyExportTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRamPolicyExportTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		query["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRamPolicyExportTasks"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRamPolicyExportTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取RAM策略导出任务列表
//
// @param request - ListRamPolicyExportTasksRequest
//
// @return ListRamPolicyExportTasksResponse
func (client *Client) ListRamPolicyExportTasks(request *ListRamPolicyExportTasksRequest) (_result *ListRamPolicyExportTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRamPolicyExportTasksResponse{}
	_body, _err := client.ListRamPolicyExportTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务版本列表
//
// @param request - ListResourceExportTaskVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExportTaskVersionsResponse
func (client *Client) ListResourceExportTaskVersionsWithOptions(exportTaskId *string, request *ListResourceExportTaskVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceExportTaskVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportVersion)) {
		query["exportVersion"] = request.ExportVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceExportTaskVersions"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId)) + "/exportVersions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceExportTaskVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务版本列表
//
// @param request - ListResourceExportTaskVersionsRequest
//
// @return ListResourceExportTaskVersionsResponse
func (client *Client) ListResourceExportTaskVersions(exportTaskId *string, request *ListResourceExportTaskVersionsRequest) (_result *ListResourceExportTaskVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceExportTaskVersionsResponse{}
	_body, _err := client.ListResourceExportTaskVersionsWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询导出任务列表
//
// @param request - ListResourceExportTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExportTasksResponse
func (client *Client) ListResourceExportTasksWithOptions(request *ListResourceExportTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceExportTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExportTaskId)) {
		query["exportTaskId"] = request.ExportTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceExportTasks"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceExportTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询导出任务列表
//
// @param request - ListResourceExportTasksRequest
//
// @return ListResourceExportTasksResponse
func (client *Client) ListResourceExportTasks(request *ListResourceExportTasksRequest) (_result *ListResourceExportTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceExportTasksResponse{}
	_body, _err := client.ListResourceExportTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源列表
//
// @param request - ListResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["sourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceValue)) {
		query["sourceValue"] = request.SourceValue
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["specType"] = request.SpecType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resources/stateparser"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源列表
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务列表
//
// @param tmpReq - ListTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksResponse
func (client *Client) ListTasksWithOptions(tmpReq *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExcludeTaskIds)) {
		request.ExcludeTaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExcludeTaskIds, tea.String("excludeTaskIds"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeTaskIdsShrink)) {
		query["excludeTaskIds"] = request.ExcludeTaskIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["groupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		query["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["projectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["taskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务列表
//
// @param request - ListTasksRequest
//
// @return ListTasksResponse
func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// terraformProvider版本
//
// @param request - ListTerraformProviderVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerraformProviderVersionsResponse
func (client *Client) ListTerraformProviderVersionsWithOptions(request *ListTerraformProviderVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTerraformProviderVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Usage)) {
		query["usage"] = request.Usage
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTerraformProviderVersions"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/version/terraform/provider"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTerraformProviderVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// terraformProvider版本
//
// @param request - ListTerraformProviderVersionsRequest
//
// @return ListTerraformProviderVersionsResponse
func (client *Client) ListTerraformProviderVersions(request *ListTerraformProviderVersionsRequest) (_result *ListTerraformProviderVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTerraformProviderVersionsResponse{}
	_body, _err := client.ListTerraformProviderVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 控制作业
//
// @param request - OperateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateJobResponse
func (client *Client) OperateJobWithOptions(taskId *string, jobId *string, operationType *string, request *OperateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OperateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["taskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OperateJob"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId)) + "/jobs/" + tea.StringValue(openapiutil.GetEncodeParam(jobId)) + "/operation/" + tea.StringValue(openapiutil.GetEncodeParam(operationType))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OperateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 控制作业
//
// @param request - OperateJobRequest
//
// @return OperateJobResponse
func (client *Client) OperateJob(taskId *string, jobId *string, operationType *string, request *OperateJobRequest) (_result *OperateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OperateJobResponse{}
	_body, _err := client.OperateJobWithOptions(taskId, jobId, operationType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除导出任务版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveResourceExportTaskVersionResponse
func (client *Client) RemoveResourceExportTaskVersionWithOptions(exportTaskId *string, exportVersion *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveResourceExportTaskVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveResourceExportTaskVersion"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId)) + "/" + tea.StringValue(openapiutil.GetEncodeParam(exportVersion))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveResourceExportTaskVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除导出任务版本
//
// @return RemoveResourceExportTaskVersionResponse
func (client *Client) RemoveResourceExportTaskVersion(exportTaskId *string, exportVersion *string) (_result *RemoveResourceExportTaskVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveResourceExportTaskVersionResponse{}
	_body, _err := client.RemoveResourceExportTaskVersionWithOptions(exportTaskId, exportVersion, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIds)) {
		body["resourceIds"] = request.ResourceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新共享
//
// @param request - UpdateAuthorizationAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorizationAttributeResponse
func (client *Client) UpdateAuthorizationAttributeWithOptions(authorizationId *string, request *UpdateAuthorizationAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAuthorizationAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DueTime)) {
		body["dueTime"] = request.DueTime
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAuthorizationAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/authorizations/" + tea.StringValue(openapiutil.GetEncodeParam(authorizationId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAuthorizationAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新共享
//
// @param request - UpdateAuthorizationAttributeRequest
//
// @return UpdateAuthorizationAttributeResponse
func (client *Client) UpdateAuthorizationAttribute(authorizationId *string, request *UpdateAuthorizationAttributeRequest) (_result *UpdateAuthorizationAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAuthorizationAttributeResponse{}
	_body, _err := client.UpdateAuthorizationAttributeWithOptions(authorizationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Explorer任务
//
// @param request - UpdateExplorerTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExplorerTaskAttributeResponse
func (client *Client) UpdateExplorerTaskAttributeWithOptions(explorerTaskId *string, request *UpdateExplorerTaskAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateExplorerTaskAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoApply)) {
		body["autoApply"] = request.AutoApply
	}

	if !tea.BoolValue(util.IsUnset(request.ExplorerTaskName)) {
		body["explorerTaskName"] = request.ExplorerTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExplorerTaskAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/explorerTask/" + tea.StringValue(openapiutil.GetEncodeParam(explorerTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExplorerTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Explorer任务
//
// @param request - UpdateExplorerTaskAttributeRequest
//
// @return UpdateExplorerTaskAttributeResponse
func (client *Client) UpdateExplorerTaskAttribute(explorerTaskId *string, request *UpdateExplorerTaskAttributeRequest) (_result *UpdateExplorerTaskAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateExplorerTaskAttributeResponse{}
	_body, _err := client.UpdateExplorerTaskAttributeWithOptions(explorerTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改分组
//
// @param request - UpdateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(groupId *string, request *UpdateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDestroy)) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !tea.BoolValue(util.IsUnset(request.AutoTrigger)) {
		body["autoTrigger"] = request.AutoTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ForcedSetting)) {
		body["forcedSetting"] = request.ForcedSetting
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyConfig)) {
		body["notifyConfig"] = request.NotifyConfig
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyOperationTypes)) {
		body["notifyOperationTypes"] = request.NotifyOperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.ReportExportField)) {
		body["reportExportField"] = request.ReportExportField
	}

	if !tea.BoolValue(util.IsUnset(request.ReportExportPath)) {
		body["reportExportPath"] = request.ReportExportPath
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerConfig)) {
		body["triggerConfig"] = request.TriggerConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerResourceType)) {
		body["triggerResourceType"] = request.TriggerResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改分组
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(groupId *string, request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新模版
//
// @param request - UpdateModuleAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModuleAttributeResponse
func (client *Client) UpdateModuleAttributeWithOptions(moduleId *string, request *UpdateModuleAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateModuleAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfo)) {
		body["groupInfo"] = request.GroupInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourcePath)) {
		body["sourcePath"] = request.SourcePath
	}

	if !tea.BoolValue(util.IsUnset(request.StatePath)) {
		body["statePath"] = request.StatePath
	}

	if !tea.BoolValue(util.IsUnset(request.VersionStrategy)) {
		body["versionStrategy"] = request.VersionStrategy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateModuleAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/modules/" + tea.StringValue(openapiutil.GetEncodeParam(moduleId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateModuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新模版
//
// @param request - UpdateModuleAttributeRequest
//
// @return UpdateModuleAttributeResponse
func (client *Client) UpdateModuleAttribute(moduleId *string, request *UpdateModuleAttributeRequest) (_result *UpdateModuleAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModuleAttributeResponse{}
	_body, _err := client.UpdateModuleAttributeWithOptions(moduleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新参数集
//
// @param request - UpdateParameterSetAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParameterSetAttributeResponse
func (client *Client) UpdateParameterSetAttributeWithOptions(parameterSetId *string, request *UpdateParameterSetAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateParameterSetAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateParameterSetAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/parameterSets/" + tea.StringValue(openapiutil.GetEncodeParam(parameterSetId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateParameterSetAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新参数集
//
// @param request - UpdateParameterSetAttributeRequest
//
// @return UpdateParameterSetAttributeResponse
func (client *Client) UpdateParameterSetAttribute(parameterSetId *string, request *UpdateParameterSetAttributeRequest) (_result *UpdateParameterSetAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateParameterSetAttributeResponse{}
	_body, _err := client.UpdateParameterSetAttributeWithOptions(parameterSetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改项目
//
// @param request - UpdateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(projectId *string, request *UpdateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProject"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/project/" + tea.StringValue(openapiutil.GetEncodeParam(projectId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改项目
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(projectId *string, request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新消息发布者
//
// @param request - UpdateRabbitmqPublisherAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRabbitmqPublisherAttributeResponse
func (client *Client) UpdateRabbitmqPublisherAttributeWithOptions(publisherId *string, request *UpdateRabbitmqPublisherAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRabbitmqPublisherAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeName)) {
		body["exchangeName"] = request.ExchangeName
	}

	if !tea.BoolValue(util.IsUnset(request.ExchangeType)) {
		body["exchangeType"] = request.ExchangeType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRabbitmqPublisherAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/publishers/" + tea.StringValue(openapiutil.GetEncodeParam(publisherId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRabbitmqPublisherAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消息发布者
//
// @param request - UpdateRabbitmqPublisherAttributeRequest
//
// @return UpdateRabbitmqPublisherAttributeResponse
func (client *Client) UpdateRabbitmqPublisherAttribute(publisherId *string, request *UpdateRabbitmqPublisherAttributeRequest) (_result *UpdateRabbitmqPublisherAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRabbitmqPublisherAttributeResponse{}
	_body, _err := client.UpdateRabbitmqPublisherAttributeWithOptions(publisherId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改RAM策略导出任务
//
// @param request - UpdateRamPolicyExportTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRamPolicyExportTaskAttributeResponse
func (client *Client) UpdateRamPolicyExportTaskAttributeWithOptions(ramPolicyExportTaskId *string, request *UpdateRamPolicyExportTaskAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRamPolicyExportTaskAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizationAccountIds)) {
		body["authorizationAccountIds"] = request.AuthorizationAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationActions)) {
		body["authorizationActions"] = request.AuthorizationActions
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationRegionIds)) {
		body["authorizationRegionIds"] = request.AuthorizationRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRamPolicyExportTaskAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/ramPolicyExportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(ramPolicyExportTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRamPolicyExportTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改RAM策略导出任务
//
// @param request - UpdateRamPolicyExportTaskAttributeRequest
//
// @return UpdateRamPolicyExportTaskAttributeResponse
func (client *Client) UpdateRamPolicyExportTaskAttribute(ramPolicyExportTaskId *string, request *UpdateRamPolicyExportTaskAttributeRequest) (_result *UpdateRamPolicyExportTaskAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRamPolicyExportTaskAttributeResponse{}
	_body, _err := client.UpdateRamPolicyExportTaskAttributeWithOptions(ramPolicyExportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新导出任务
//
// @param request - UpdateResourceExportTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceExportTaskAttributeResponse
func (client *Client) UpdateResourceExportTaskAttributeWithOptions(exportTaskId *string, request *UpdateResourceExportTaskAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceExportTaskAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigPath)) {
		body["configPath"] = request.ConfigPath
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeRules)) {
		body["excludeRules"] = request.ExcludeRules
	}

	if !tea.BoolValue(util.IsUnset(request.ExportToModule)) {
		body["exportToModule"] = request.ExportToModule
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeRules)) {
		body["includeRules"] = request.IncludeRules
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformProviderVersion)) {
		body["terraformProviderVersion"] = request.TerraformProviderVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVersion)) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceExportTaskAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/exportTasks/" + tea.StringValue(openapiutil.GetEncodeParam(exportTaskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceExportTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新导出任务
//
// @param request - UpdateResourceExportTaskAttributeRequest
//
// @return UpdateResourceExportTaskAttributeResponse
func (client *Client) UpdateResourceExportTaskAttribute(exportTaskId *string, request *UpdateResourceExportTaskAttributeRequest) (_result *UpdateResourceExportTaskAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceExportTaskAttributeResponse{}
	_body, _err := client.UpdateResourceExportTaskAttributeWithOptions(exportTaskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改任务
//
// @param request - UpdateTaskAttributeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskAttributeResponse
func (client *Client) UpdateTaskAttributeWithOptions(taskId *string, request *UpdateTaskAttributeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTaskAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoApply)) {
		body["autoApply"] = request.AutoApply
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDestroy)) {
		body["autoDestroy"] = request.AutoDestroy
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfo)) {
		body["groupInfo"] = request.GroupInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InitModuleState)) {
		body["initModuleState"] = request.InitModuleState
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleId)) {
		body["moduleId"] = request.ModuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleVersion)) {
		body["moduleVersion"] = request.ModuleVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.ProtectionStrategy)) {
		body["protectionStrategy"] = request.ProtectionStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		body["ramRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.SkipPropertyValidation)) {
		body["skipPropertyValidation"] = request.SkipPropertyValidation
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVersion)) {
		body["terraformVersion"] = request.TerraformVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerStrategy)) {
		body["triggerStrategy"] = request.TriggerStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerValue)) {
		body["triggerValue"] = request.TriggerValue
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskAttribute"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改任务
//
// @param request - UpdateTaskAttributeRequest
//
// @return UpdateTaskAttributeResponse
func (client *Client) UpdateTaskAttribute(taskId *string, request *UpdateTaskAttributeRequest) (_result *UpdateTaskAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTaskAttributeResponse{}
	_body, _err := client.UpdateTaskAttributeWithOptions(taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改分组优先级配置
//
// @param request - UpdateTaskPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskPolicyResponse
func (client *Client) UpdateTaskPolicyWithOptions(groupId *string, request *UpdateTaskPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTaskPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["clientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskPolicies)) {
		body["taskPolicies"] = request.TaskPolicies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskPolicy"),
		Version:     tea.String("2021-08-06"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/group/task/policy/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改分组优先级配置
//
// @param request - UpdateTaskPolicyRequest
//
// @return UpdateTaskPolicyResponse
func (client *Client) UpdateTaskPolicy(groupId *string, request *UpdateTaskPolicyRequest) (_result *UpdateTaskPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTaskPolicyResponse{}
	_body, _err := client.UpdateTaskPolicyWithOptions(groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
