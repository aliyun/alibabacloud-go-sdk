// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuthDiagnosisRequest struct {
	AutoCreateRole   *bool                            `json:"autoCreateRole,omitempty" xml:"autoCreateRole,omitempty"`
	AutoInstallAgent *bool                            `json:"autoInstallAgent,omitempty" xml:"autoInstallAgent,omitempty"`
	Instances        []*AuthDiagnosisRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s AuthDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequest) SetAutoCreateRole(v bool) *AuthDiagnosisRequest {
	s.AutoCreateRole = &v
	return s
}

func (s *AuthDiagnosisRequest) SetAutoInstallAgent(v bool) *AuthDiagnosisRequest {
	s.AutoInstallAgent = &v
	return s
}

func (s *AuthDiagnosisRequest) SetInstances(v []*AuthDiagnosisRequestInstances) *AuthDiagnosisRequest {
	s.Instances = v
	return s
}

type AuthDiagnosisRequestInstances struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s AuthDiagnosisRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisRequestInstances) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisRequestInstances) SetInstance(v string) *AuthDiagnosisRequestInstances {
	s.Instance = &v
	return s
}

func (s *AuthDiagnosisRequestInstances) SetRegion(v string) *AuthDiagnosisRequestInstances {
	s.Region = &v
	return s
}

type AuthDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s AuthDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponseBody) SetCode(v string) *AuthDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetData(v interface{}) *AuthDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *AuthDiagnosisResponseBody) SetMessage(v string) *AuthDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *AuthDiagnosisResponseBody) SetRequestId(v string) *AuthDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type AuthDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *AuthDiagnosisResponse) SetHeaders(v map[string]*string) *AuthDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *AuthDiagnosisResponse) SetStatusCode(v int32) *AuthDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthDiagnosisResponse) SetBody(v *AuthDiagnosisResponseBody) *AuthDiagnosisResponse {
	s.Body = v
	return s
}

type GenerateCopilotResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseRequest) SetLlmParamString(v string) *GenerateCopilotResponseRequest {
	s.LlmParamString = &v
	return s
}

type GenerateCopilotResponseResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Requests for llm service failed
	Massage *string `json:"massage,omitempty" xml:"massage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponseBody) SetCode(v string) *GenerateCopilotResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetData(v string) *GenerateCopilotResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetMassage(v string) *GenerateCopilotResponseResponseBody {
	s.Massage = &v
	return s
}

func (s *GenerateCopilotResponseResponseBody) SetRequestId(v string) *GenerateCopilotResponseResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCopilotResponseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotResponseResponse) SetStatusCode(v int32) *GenerateCopilotResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotResponseResponse) SetBody(v *GenerateCopilotResponseResponseBody) *GenerateCopilotResponseResponse {
	s.Body = v
	return s
}

type GetAIQueryResultRequest struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
}

func (s GetAIQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAIQueryResultRequest) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultRequest) SetAnalysisId(v string) *GetAIQueryResultRequest {
	s.AnalysisId = &v
	return s
}

type GetAIQueryResultResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {\\"task_id\\": \\"y4ba8uRV\\"}
	Data    *string `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAIQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAIQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultResponseBody) SetCode(v string) *GetAIQueryResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetData(v string) *GetAIQueryResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetMessage(v string) *GetAIQueryResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetAIQueryResultResponseBody) SetRequestId(v string) *GetAIQueryResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAIQueryResultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAIQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAIQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAIQueryResultResponse) GoString() string {
	return s.String()
}

func (s *GetAIQueryResultResponse) SetHeaders(v map[string]*string) *GetAIQueryResultResponse {
	s.Headers = v
	return s
}

func (s *GetAIQueryResultResponse) SetStatusCode(v int32) *GetAIQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAIQueryResultResponse) SetBody(v *GetAIQueryResultResponseBody) *GetAIQueryResultResponse {
	s.Body = v
	return s
}

type GetAbnormalEventsCountRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod     *string `json:"pod,omitempty" xml:"pod,omitempty"`
	ShowPod *int32  `json:"showPod,omitempty" xml:"showPod,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetAbnormalEventsCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountRequest) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountRequest) SetCluster(v string) *GetAbnormalEventsCountRequest {
	s.Cluster = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetEnd(v float32) *GetAbnormalEventsCountRequest {
	s.End = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetInstance(v string) *GetAbnormalEventsCountRequest {
	s.Instance = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetNamespace(v string) *GetAbnormalEventsCountRequest {
	s.Namespace = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetPod(v string) *GetAbnormalEventsCountRequest {
	s.Pod = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetShowPod(v int32) *GetAbnormalEventsCountRequest {
	s.ShowPod = &v
	return s
}

func (s *GetAbnormalEventsCountRequest) SetStart(v float32) *GetAbnormalEventsCountRequest {
	s.Start = &v
	return s
}

type GetAbnormalEventsCountResponseBody struct {
	// example:
	//
	// Success
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetAbnormalEventsCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAbnormalEventsCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBody) SetCode(v string) *GetAbnormalEventsCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetData(v []*GetAbnormalEventsCountResponseBodyData) *GetAbnormalEventsCountResponseBody {
	s.Data = v
	return s
}

func (s *GetAbnormalEventsCountResponseBody) SetMessage(v string) *GetAbnormalEventsCountResponseBody {
	s.Message = &v
	return s
}

type GetAbnormalEventsCountResponseBodyData struct {
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAbnormalEventsCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponseBodyData) SetType(v string) *GetAbnormalEventsCountResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAbnormalEventsCountResponseBodyData) SetValue(v int64) *GetAbnormalEventsCountResponseBodyData {
	s.Value = &v
	return s
}

type GetAbnormalEventsCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAbnormalEventsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAbnormalEventsCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAbnormalEventsCountResponse) GoString() string {
	return s.String()
}

func (s *GetAbnormalEventsCountResponse) SetHeaders(v map[string]*string) *GetAbnormalEventsCountResponse {
	s.Headers = v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetStatusCode(v int32) *GetAbnormalEventsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAbnormalEventsCountResponse) SetBody(v *GetAbnormalEventsCountResponseBody) *GetAbnormalEventsCountResponse {
	s.Body = v
	return s
}

type GetAgentRequest struct {
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

func (s GetAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) SetAgentId(v string) *GetAgentRequest {
	s.AgentId = &v
	return s
}

type GetAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

type GetAgentResponseBodyData struct {
	CreatedAt   *string                             `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string                             `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string                             `json:"name,omitempty" xml:"name,omitempty"`
	SupportArch *string                             `json:"support_arch,omitempty" xml:"support_arch,omitempty"`
	Type        *string                             `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt   *string                             `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Versions    []*GetAgentResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) SetCreatedAt(v string) *GetAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDescription(v string) *GetAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyData) SetId(v string) *GetAgentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAgentResponseBodyData) SetName(v string) *GetAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyData) SetSupportArch(v string) *GetAgentResponseBodyData {
	s.SupportArch = &v
	return s
}

func (s *GetAgentResponseBodyData) SetType(v string) *GetAgentResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAgentResponseBodyData) SetUpdatedAt(v string) *GetAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentResponseBodyData) SetVersions(v []*GetAgentResponseBodyDataVersions) *GetAgentResponseBodyData {
	s.Versions = v
	return s
}

type GetAgentResponseBodyDataVersions struct {
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	InstallScript   *string `json:"install_script,omitempty" xml:"install_script,omitempty"`
	UninstallScript *string `json:"uninstall_script,omitempty" xml:"uninstall_script,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UpgradeScript   *string `json:"upgrade_script,omitempty" xml:"upgrade_script,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAgentResponseBodyDataVersions) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataVersions) SetCreatedAt(v string) *GetAgentResponseBodyDataVersions {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetInstallScript(v string) *GetAgentResponseBodyDataVersions {
	s.InstallScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUninstallScript(v string) *GetAgentResponseBodyDataVersions {
	s.UninstallScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUpdatedAt(v string) *GetAgentResponseBodyDataVersions {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUpgradeScript(v string) *GetAgentResponseBodyDataVersions {
	s.UpgradeScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetVersion(v string) *GetAgentResponseBodyDataVersions {
	s.Version = &v
	return s
}

type GetAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetStatusCode(v int32) *GetAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

type GetAgentTaskRequest struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAgentTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAgentTaskRequest) SetTaskId(v string) *GetAgentTaskRequest {
	s.TaskId = &v
	return s
}

type GetAgentTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAgentTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAgentTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBody) SetRequestId(v string) *GetAgentTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentTaskResponseBody) SetCode(v string) *GetAgentTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentTaskResponseBody) SetData(v *GetAgentTaskResponseBodyData) *GetAgentTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentTaskResponseBody) SetMessage(v string) *GetAgentTaskResponseBody {
	s.Message = &v
	return s
}

type GetAgentTaskResponseBodyData struct {
	Jobs   []*GetAgentTaskResponseBodyDataJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	TaskId *string                             `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetAgentTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBodyData) SetJobs(v []*GetAgentTaskResponseBodyDataJobs) *GetAgentTaskResponseBodyData {
	s.Jobs = v
	return s
}

func (s *GetAgentTaskResponseBodyData) SetTaskId(v string) *GetAgentTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type GetAgentTaskResponseBodyDataJobs struct {
	Error    *string     `json:"error,omitempty" xml:"error,omitempty"`
	Instance *string     `json:"instance,omitempty" xml:"instance,omitempty"`
	Params   interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Region   *string     `json:"region,omitempty" xml:"region,omitempty"`
	Result   *string     `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetAgentTaskResponseBodyDataJobs) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTaskResponseBodyDataJobs) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponseBodyDataJobs) SetError(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Error = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetInstance(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Instance = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetParams(v interface{}) *GetAgentTaskResponseBodyDataJobs {
	s.Params = v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetRegion(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Region = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetResult(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Result = &v
	return s
}

func (s *GetAgentTaskResponseBodyDataJobs) SetStatus(v string) *GetAgentTaskResponseBodyDataJobs {
	s.Status = &v
	return s
}

type GetAgentTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAgentTaskResponse) SetHeaders(v map[string]*string) *GetAgentTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAgentTaskResponse) SetStatusCode(v int32) *GetAgentTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentTaskResponse) SetBody(v *GetAgentTaskResponseBody) *GetAgentTaskResponse {
	s.Body = v
	return s
}

type GetDiagnosisResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// quzuYl23
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) SetTaskId(v string) *GetDiagnosisResultRequest {
	s.TaskId = &v
	return s
}

type GetDiagnosisResultResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetDiagnosisResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9515E5A0-8905-59B0-9BBF-5F0BE568C3A0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetDiagnosisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBody) SetCode(v string) *GetDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetData(v *GetDiagnosisResultResponseBodyData) *GetDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetMessage(v string) *GetDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetRequestId(v string) *GetDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

type GetDiagnosisResultResponseBodyData struct {
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {
	//
	//     "jobs":[
	//
	//         {
	//
	//             "cmd":"mkdir -p /var/log/sysak && sysak podmem -r 100  -a -j /var/log/sysak/podmem.json > /dev/null 2>&1 && cat /var/log/sysak/podmem.json",
	//
	//             "instance":"172.20.12.174",
	//
	//             "fetch_file_list":[
	//
	//             ]
	//
	//         }
	//
	//     ],
	//
	//     "in_order":true,
	//
	//     "offline_mode":false,
	//
	//     "offline_results":[
	//
	//     ]
	//
	// }
	Command   interface{} `json:"command,omitempty" xml:"command,omitempty"`
	CreatedAt *string     `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// Diagnosis failed
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// example:
	//
	// {
	//
	//     "type":"all",
	//
	//     "value":"",
	//
	//     "channel":"ssh",
	//
	//     "instance":"172.1.2.174",
	//
	//     "service_name":"filecache"
	//
	// }
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// {
	//
	//     "summary":"  memory cgroup leak",
	//
	//     "dataMemEvent":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"Util",
	//
	//                 "value":20
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemLeak",
	//
	//                 "value":"OK"
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemcgLeak",
	//
	//                 "value":"NG"
	//
	//             },
	//
	//             {
	//
	//                 "key":"MemFrag",
	//
	//                 "value":"OK"
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataMemOverView":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"app",
	//
	//                 "value":10937332
	//
	//             },
	//
	//             {
	//
	//                 "key":"free",
	//
	//                 "value":806800
	//
	//             },
	//
	//             {
	//
	//                 "key":"kernel",
	//
	//                 "value":4527660
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataKerMem":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"SReclaimable",
	//
	//                 "value":3411292
	//
	//             },
	//
	//             {
	//
	//                 "key":"VmallocUsed",
	//
	//                 "value":30980
	//
	//             },
	//
	//             {
	//
	//                 "key":"allocPage",
	//
	//                 "value":177732
	//
	//             },
	//
	//             {
	//
	//                 "key":"KernelStack",
	//
	//                 "value":9280
	//
	//             },
	//
	//             {
	//
	//                 "key":"PageTables",
	//
	//                 "value":38056
	//
	//             },
	//
	//             {
	//
	//                 "key":"SUnreclaim",
	//
	//                 "value":170248
	//
	//             },
	//
	//             {
	//
	//                 "key":"reserved",
	//
	//                 "value":690072
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataUserMem":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":"filecache",
	//
	//                 "value":8010008
	//
	//             },
	//
	//             {
	//
	//                 "key":"anon",
	//
	//                 "value":2468608
	//
	//             },
	//
	//             {
	//
	//                 "key":"mlock",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"huge1G",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"huge2M",
	//
	//                 "value":0
	//
	//             },
	//
	//             {
	//
	//                 "key":"buffers",
	//
	//                 "value":458608
	//
	//             },
	//
	//             {
	//
	//                 "key":"shmem",
	//
	//                 "value":2284
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataCacheList":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":0,
	//
	//                 "Name":"/var/lib/mysql/sysom/sys_handler_log.ibd",
	//
	//                 "cached":576764,
	//
	//                 "Task":"mysqld_78575 "
	//
	//             },
	//
	//             {
	//
	//                 "key":1,
	//
	//                 "Name":"/var/log/sysom/sysom-migration-access.log",
	//
	//                 "cached":276688,
	//
	//                 "Task":"gunicorn_33647 ,gunicorn_460836 ,gunicorn_559934 ,gunicorn_731758 ,gunicorn_2362682 "
	//
	//             },
	//
	//             {
	//
	//                 "key":2,
	//
	//                 "Name":"/var/log/sysom/sysom-rtdemo-access.log",
	//
	//                 "cached":229404,
	//
	//                 "Task":"gunicorn_60718 ,gunicorn_720734 ,gunicorn_722168 "
	//
	//             },
	//
	//             {
	//
	//                 "key":3,
	//
	//                 "Name":"/var/log/sysom/sysom-monitor-server-access.log",
	//
	//                 "cached":197368,
	//
	//                 "Task":"gunicorn_33682 ,gunicorn_671155 ,gunicorn_714998 "
	//
	//             },
	//
	//             {
	//
	//                 "key":4,
	//
	//                 "Name":"/var/log/sysom/sysom-channel-access.log",
	//
	//                 "cached":180276,
	//
	//                 "Task":"gunicorn_33233 ,gunicorn_499735 ,gunicorn_725497 "
	//
	//             },
	//
	//             {
	//
	//                 "key":5,
	//
	//                 "Name":"total cached of close file",
	//
	//                 "cached":3729668,
	//
	//                 "Task":""
	//
	//             }
	//
	//         ]
	//
	//     },
	//
	//     "dataProcMemList":{
	//
	//         "data":[
	//
	//             {
	//
	//                 "key":0,
	//
	//                 "task":"mysqld",
	//
	//                 "MemTotal":240856,
	//
	//                 "RssAnon":218248,
	//
	//                 "RssFile":22608
	//
	//             },
	//
	//             {
	//
	//                 "key":1,
	//
	//                 "task":"systemd-journal",
	//
	//                 "MemTotal":150248,
	//
	//                 "RssAnon":74300,
	//
	//                 "RssFile":75944
	//
	//             },
	//
	//             {
	//
	//                 "key":2,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":144640,
	//
	//                 "RssAnon":114200,
	//
	//                 "RssFile":30440
	//
	//             },
	//
	//             {
	//
	//                 "key":3,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":141480,
	//
	//                 "RssAnon":111040,
	//
	//                 "RssFile":30440
	//
	//             },
	//
	//             {
	//
	//                 "key":4,
	//
	//                 "task":"grafana-server",
	//
	//                 "MemTotal":103660,
	//
	//                 "RssAnon":42732,
	//
	//                 "RssFile":60928
	//
	//             },
	//
	//             {
	//
	//                 "key":5,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":97444,
	//
	//                 "RssAnon":76256,
	//
	//                 "RssFile":21188
	//
	//             },
	//
	//             {
	//
	//                 "key":6,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":97260,
	//
	//                 "RssAnon":76072,
	//
	//                 "RssFile":21188
	//
	//             },
	//
	//             {
	//
	//                 "key":7,
	//
	//                 "task":"prometheus",
	//
	//                 "MemTotal":95356,
	//
	//                 "RssAnon":45376,
	//
	//                 "RssFile":49980
	//
	//             },
	//
	//             {
	//
	//                 "key":8,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":90144,
	//
	//                 "RssAnon":76456,
	//
	//                 "RssFile":13688
	//
	//             },
	//
	//             {
	//
	//                 "key":9,
	//
	//                 "task":"gunicorn",
	//
	//                 "MemTotal":89796,
	//
	//                 "RssAnon":76108,
	//
	//                 "RssFile":13688
	//
	//             }
	//
	//         ]
	//
	//     }
	//
	// }
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// grcuU21a
	TaskId    *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// example:
	//
	// /diagnose/detail/qe3Z34sa
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDiagnosisResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponseBodyData) SetCode(v int32) *GetDiagnosisResultResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetCommand(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Command = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetCreatedAt(v string) *GetDiagnosisResultResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetErrMsg(v string) *GetDiagnosisResultResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetParams(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Params = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetResult(v interface{}) *GetDiagnosisResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetServiceName(v string) *GetDiagnosisResultResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetStatus(v string) *GetDiagnosisResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetTaskId(v string) *GetDiagnosisResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetUpdatedAt(v string) *GetDiagnosisResultResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDiagnosisResultResponseBodyData) SetUrl(v string) *GetDiagnosisResultResponseBodyData {
	s.Url = &v
	return s
}

type GetDiagnosisResultResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiagnosisResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDiagnosisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultResponse) SetHeaders(v map[string]*string) *GetDiagnosisResultResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosisResultResponse) SetStatusCode(v int32) *GetDiagnosisResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosisResultResponse) SetBody(v *GetDiagnosisResultResponseBody) *GetDiagnosisResultResponse {
	s.Body = v
	return s
}

type GetHealthPercentageRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetHealthPercentageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageRequest) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageRequest) SetCluster(v string) *GetHealthPercentageRequest {
	s.Cluster = &v
	return s
}

func (s *GetHealthPercentageRequest) SetEnd(v float32) *GetHealthPercentageRequest {
	s.End = &v
	return s
}

func (s *GetHealthPercentageRequest) SetInstance(v string) *GetHealthPercentageRequest {
	s.Instance = &v
	return s
}

func (s *GetHealthPercentageRequest) SetStart(v float32) *GetHealthPercentageRequest {
	s.Start = &v
	return s
}

type GetHealthPercentageResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetHealthPercentageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetHealthPercentageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBody) SetCode(v string) *GetHealthPercentageResponseBody {
	s.Code = &v
	return s
}

func (s *GetHealthPercentageResponseBody) SetData(v []*GetHealthPercentageResponseBodyData) *GetHealthPercentageResponseBody {
	s.Data = v
	return s
}

func (s *GetHealthPercentageResponseBody) SetMessage(v string) *GetHealthPercentageResponseBody {
	s.Message = &v
	return s
}

type GetHealthPercentageResponseBodyData struct {
	// example:
	//
	// health
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetHealthPercentageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponseBodyData) SetType(v string) *GetHealthPercentageResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetHealthPercentageResponseBodyData) SetValue(v int64) *GetHealthPercentageResponseBodyData {
	s.Value = &v
	return s
}

type GetHealthPercentageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHealthPercentageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHealthPercentageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthPercentageResponse) GoString() string {
	return s.String()
}

func (s *GetHealthPercentageResponse) SetHeaders(v map[string]*string) *GetHealthPercentageResponse {
	s.Headers = v
	return s
}

func (s *GetHealthPercentageResponse) SetStatusCode(v int32) *GetHealthPercentageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHealthPercentageResponse) SetBody(v *GetHealthPercentageResponseBody) *GetHealthPercentageResponse {
	s.Body = v
	return s
}

type GetHostCountRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetHostCountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostCountRequest) GoString() string {
	return s.String()
}

func (s *GetHostCountRequest) SetCluster(v string) *GetHostCountRequest {
	s.Cluster = &v
	return s
}

func (s *GetHostCountRequest) SetEnd(v float32) *GetHostCountRequest {
	s.End = &v
	return s
}

func (s *GetHostCountRequest) SetInstance(v string) *GetHostCountRequest {
	s.Instance = &v
	return s
}

func (s *GetHostCountRequest) SetStart(v float32) *GetHostCountRequest {
	s.Start = &v
	return s
}

type GetHostCountResponseBody struct {
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetHostCountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// “”
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 43A910E9-A739-525E-855D-A32C257F1826
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetHostCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostCountResponseBody) SetCode(v string) *GetHostCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetHostCountResponseBody) SetData(v []*GetHostCountResponseBodyData) *GetHostCountResponseBody {
	s.Data = v
	return s
}

func (s *GetHostCountResponseBody) SetMessage(v string) *GetHostCountResponseBody {
	s.Message = &v
	return s
}

func (s *GetHostCountResponseBody) SetRequestId(v string) *GetHostCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostCountResponseBody) SetTotal(v int64) *GetHostCountResponseBody {
	s.Total = &v
	return s
}

type GetHostCountResponseBodyData struct {
	// example:
	//
	// 1725797727754
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 5
	Value *int32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetHostCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHostCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHostCountResponseBodyData) SetTime(v int64) *GetHostCountResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetHostCountResponseBodyData) SetValue(v int32) *GetHostCountResponseBodyData {
	s.Value = &v
	return s
}

type GetHostCountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostCountResponse) GoString() string {
	return s.String()
}

func (s *GetHostCountResponse) SetHeaders(v map[string]*string) *GetHostCountResponse {
	s.Headers = v
	return s
}

func (s *GetHostCountResponse) SetStatusCode(v int32) *GetHostCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostCountResponse) SetBody(v *GetHostCountResponseBody) *GetHostCountResponse {
	s.Body = v
	return s
}

type GetHotspotAnalysisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// GR
	AppType *string `json:"appType,omitempty" xml:"appType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725413948000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725410348000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1657494
	Pid *int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotAnalysisRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisRequest) SetAppType(v string) *GetHotspotAnalysisRequest {
	s.AppType = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetBegEnd(v int64) *GetHotspotAnalysisRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetBegStart(v int64) *GetHotspotAnalysisRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetInstance(v string) *GetHotspotAnalysisRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetPid(v int64) *GetHotspotAnalysisRequest {
	s.Pid = &v
	return s
}

func (s *GetHotspotAnalysisRequest) SetTable(v string) *GetHotspotAnalysisRequest {
	s.Table = &v
	return s
}

type GetHotspotAnalysisResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotspotAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisResponseBody) SetCode(v string) *GetHotspotAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetData(v string) *GetHotspotAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetMessage(v string) *GetHotspotAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetRequestId(v string) *GetHotspotAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetSuccess(v bool) *GetHotspotAnalysisResponseBody {
	s.Success = &v
	return s
}

type GetHotspotAnalysisResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotAnalysisResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisResponse) SetHeaders(v map[string]*string) *GetHotspotAnalysisResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotAnalysisResponse) SetStatusCode(v int32) *GetHotspotAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotAnalysisResponse) SetBody(v *GetHotspotAnalysisResponseBody) *GetHotspotAnalysisResponse {
	s.Body = v
	return s
}

type GetHotspotInstanceListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1725413947000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725413647000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListRequest) SetBegEnd(v int64) *GetHotspotInstanceListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotInstanceListRequest) SetBegStart(v int64) *GetHotspotInstanceListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotInstanceListRequest) SetTable(v string) *GetHotspotInstanceListRequest {
	s.Table = &v
	return s
}

type GetHotspotInstanceListResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotInstanceListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponseBody) SetCode(v string) *GetHotspotInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetData(v *GetHotspotInstanceListResponseBodyData) *GetHotspotInstanceListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetMessage(v string) *GetHotspotInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotInstanceListResponseBody) SetRequestId(v string) *GetHotspotInstanceListResponseBody {
	s.RequestId = &v
	return s
}

type GetHotspotInstanceListResponseBodyData struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotInstanceListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotInstanceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponseBodyData) SetColumns(v []*string) *GetHotspotInstanceListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotspotInstanceListResponseBodyData) SetValues(v []*string) *GetHotspotInstanceListResponseBodyData {
	s.Values = v
	return s
}

type GetHotspotInstanceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotInstanceListResponse) SetHeaders(v map[string]*string) *GetHotspotInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotInstanceListResponse) SetStatusCode(v int32) *GetHotspotInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotInstanceListResponse) SetBody(v *GetHotspotInstanceListResponseBody) *GetHotspotInstanceListResponse {
	s.Body = v
	return s
}

type GetHotspotPidListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1725413947000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725413647000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-2ze5ru5rjurix7f71sxv
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotPidListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotPidListRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListRequest) SetBegEnd(v int64) *GetHotspotPidListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotPidListRequest) SetBegStart(v int64) *GetHotspotPidListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotPidListRequest) SetInstance(v string) *GetHotspotPidListRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotPidListRequest) SetTable(v string) *GetHotspotPidListRequest {
	s.Table = &v
	return s
}

type GetHotspotPidListResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotPidListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotPidListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotPidListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponseBody) SetCode(v string) *GetHotspotPidListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotPidListResponseBody) SetData(v *GetHotspotPidListResponseBodyData) *GetHotspotPidListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotPidListResponseBody) SetMessage(v string) *GetHotspotPidListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotPidListResponseBody) SetRequestId(v string) *GetHotspotPidListResponseBody {
	s.RequestId = &v
	return s
}

type GetHotspotPidListResponseBodyData struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotPidListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotPidListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponseBodyData) SetColumns(v []*string) *GetHotspotPidListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotspotPidListResponseBodyData) SetValues(v [][]*string) *GetHotspotPidListResponseBodyData {
	s.Values = v
	return s
}

type GetHotspotPidListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotPidListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotPidListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotPidListResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotPidListResponse) SetHeaders(v map[string]*string) *GetHotspotPidListResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotPidListResponse) SetStatusCode(v int32) *GetHotspotPidListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotPidListResponse) SetBody(v *GetHotspotPidListResponseBody) *GetHotspotPidListResponse {
	s.Body = v
	return s
}

type GetListRecordRequest struct {
	// example:
	//
	// 5
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetListRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetListRecordRequest) GoString() string {
	return s.String()
}

func (s *GetListRecordRequest) SetCurrent(v int64) *GetListRecordRequest {
	s.Current = &v
	return s
}

func (s *GetListRecordRequest) SetPageSize(v int64) *GetListRecordRequest {
	s.PageSize = &v
	return s
}

type GetListRecordResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetListRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 19
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetListRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetListRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetListRecordResponseBody) SetCode(v string) *GetListRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetListRecordResponseBody) SetData(v []*GetListRecordResponseBodyData) *GetListRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetListRecordResponseBody) SetMessage(v string) *GetListRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetListRecordResponseBody) SetRequestId(v string) *GetListRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListRecordResponseBody) SetTotal(v int64) *GetListRecordResponseBody {
	s.Total = &v
	return s
}

type GetListRecordResponseBodyData struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysisId,omitempty" xml:"analysisId,omitempty"`
	// example:
	//
	// 2024-12-24 12:02:05
	AnalysisTime *string `json:"analysisTime,omitempty" xml:"analysisTime,omitempty"`
	// example:
	//
	// timeout=2000 ms
	Arguments *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
	FailedLog *string `json:"failedLog,omitempty" xml:"failedLog,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetListRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetListRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetListRecordResponseBodyData) SetAnalysisId(v string) *GetListRecordResponseBodyData {
	s.AnalysisId = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetAnalysisTime(v string) *GetListRecordResponseBodyData {
	s.AnalysisTime = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetArguments(v string) *GetListRecordResponseBodyData {
	s.Arguments = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetFailedLog(v string) *GetListRecordResponseBodyData {
	s.FailedLog = &v
	return s
}

func (s *GetListRecordResponseBodyData) SetStatus(v string) *GetListRecordResponseBodyData {
	s.Status = &v
	return s
}

type GetListRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetListRecordResponse) GoString() string {
	return s.String()
}

func (s *GetListRecordResponse) SetHeaders(v map[string]*string) *GetListRecordResponse {
	s.Headers = v
	return s
}

func (s *GetListRecordResponse) SetStatusCode(v int32) *GetListRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListRecordResponse) SetBody(v *GetListRecordResponseBody) *GetListRecordResponse {
	s.Body = v
	return s
}

type GetProblemPercentageRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetProblemPercentageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPercentageRequest) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageRequest) SetCluster(v string) *GetProblemPercentageRequest {
	s.Cluster = &v
	return s
}

func (s *GetProblemPercentageRequest) SetEnd(v float32) *GetProblemPercentageRequest {
	s.End = &v
	return s
}

func (s *GetProblemPercentageRequest) SetInstance(v string) *GetProblemPercentageRequest {
	s.Instance = &v
	return s
}

func (s *GetProblemPercentageRequest) SetStart(v float32) *GetProblemPercentageRequest {
	s.Start = &v
	return s
}

type GetProblemPercentageResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetProblemPercentageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 19
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetProblemPercentageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPercentageResponseBody) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponseBody) SetCode(v string) *GetProblemPercentageResponseBody {
	s.Code = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetData(v []*GetProblemPercentageResponseBodyData) *GetProblemPercentageResponseBody {
	s.Data = v
	return s
}

func (s *GetProblemPercentageResponseBody) SetMessage(v string) *GetProblemPercentageResponseBody {
	s.Message = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetRequestId(v string) *GetProblemPercentageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProblemPercentageResponseBody) SetTotal(v int64) *GetProblemPercentageResponseBody {
	s.Total = &v
	return s
}

type GetProblemPercentageResponseBodyData struct {
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 5
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProblemPercentageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPercentageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponseBodyData) SetType(v string) *GetProblemPercentageResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetProblemPercentageResponseBodyData) SetValue(v int64) *GetProblemPercentageResponseBodyData {
	s.Value = &v
	return s
}

type GetProblemPercentageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProblemPercentageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProblemPercentageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProblemPercentageResponse) GoString() string {
	return s.String()
}

func (s *GetProblemPercentageResponse) SetHeaders(v map[string]*string) *GetProblemPercentageResponse {
	s.Headers = v
	return s
}

func (s *GetProblemPercentageResponse) SetStatusCode(v int32) *GetProblemPercentageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProblemPercentageResponse) SetBody(v *GetProblemPercentageResponseBody) *GetProblemPercentageResponse {
	s.Body = v
	return s
}

type GetRangeScoreRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s GetRangeScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRangeScoreRequest) GoString() string {
	return s.String()
}

func (s *GetRangeScoreRequest) SetCluster(v string) *GetRangeScoreRequest {
	s.Cluster = &v
	return s
}

func (s *GetRangeScoreRequest) SetEnd(v float32) *GetRangeScoreRequest {
	s.End = &v
	return s
}

func (s *GetRangeScoreRequest) SetInstance(v string) *GetRangeScoreRequest {
	s.Instance = &v
	return s
}

func (s *GetRangeScoreRequest) SetStart(v float32) *GetRangeScoreRequest {
	s.Start = &v
	return s
}

type GetRangeScoreResponseBody struct {
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetRangeScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 2
	Total *float32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRangeScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRangeScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponseBody) SetCode(v string) *GetRangeScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetData(v []*GetRangeScoreResponseBodyData) *GetRangeScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetRangeScoreResponseBody) SetMessage(v string) *GetRangeScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetRequestId(v string) *GetRangeScoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRangeScoreResponseBody) SetTotal(v float32) *GetRangeScoreResponseBody {
	s.Total = &v
	return s
}

type GetRangeScoreResponseBodyData struct {
	// example:
	//
	// 1725797727754
	Time *float32 `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 100
	Value *float32 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRangeScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRangeScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponseBodyData) SetTime(v float32) *GetRangeScoreResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetRangeScoreResponseBodyData) SetType(v string) *GetRangeScoreResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetRangeScoreResponseBodyData) SetValue(v float32) *GetRangeScoreResponseBodyData {
	s.Value = &v
	return s
}

type GetRangeScoreResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRangeScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRangeScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRangeScoreResponse) GoString() string {
	return s.String()
}

func (s *GetRangeScoreResponse) SetHeaders(v map[string]*string) *GetRangeScoreResponse {
	s.Headers = v
	return s
}

func (s *GetRangeScoreResponse) SetStatusCode(v int32) *GetRangeScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRangeScoreResponse) SetBody(v *GetRangeScoreResponseBody) *GetRangeScoreResponse {
	s.Body = v
	return s
}

type GetResourcesRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// mem
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) SetCluster(v string) *GetResourcesRequest {
	s.Cluster = &v
	return s
}

func (s *GetResourcesRequest) SetInstance(v string) *GetResourcesRequest {
	s.Instance = &v
	return s
}

func (s *GetResourcesRequest) SetType(v string) *GetResourcesRequest {
	s.Type = &v
	return s
}

type GetResourcesResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s GetResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBody) SetCode(v string) *GetResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcesResponseBody) SetData(v *GetResourcesResponseBodyData) *GetResourcesResponseBody {
	s.Data = v
	return s
}

func (s *GetResourcesResponseBody) SetMessage(v string) *GetResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcesResponseBody) SetRequestId(v string) *GetResourcesResponseBody {
	s.RequestId = &v
	return s
}

type GetResourcesResponseBodyData struct {
	// example:
	//
	// 2354
	Total *float32 `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// Kbytes
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 100
	Usage *float32 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s GetResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourcesResponseBodyData) SetTotal(v float32) *GetResourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetResourcesResponseBodyData) SetUnit(v string) *GetResourcesResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetResourcesResponseBodyData) SetUsage(v float32) *GetResourcesResponseBodyData {
	s.Usage = &v
	return s
}

type GetResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcesResponse) GoString() string {
	return s.String()
}

func (s *GetResourcesResponse) SetHeaders(v map[string]*string) *GetResourcesResponse {
	s.Headers = v
	return s
}

func (s *GetResourcesResponse) SetStatusCode(v int32) *GetResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcesResponse) SetBody(v *GetResourcesResponseBody) *GetResourcesResponse {
	s.Body = v
	return s
}

type InstallAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// InstallAndUpgrade
	InstallType *string `json:"install_type,omitempty" xml:"install_type,omitempty"`
	// This parameter is required.
	Instances []*InstallAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s InstallAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentRequest) SetAgentId(v string) *InstallAgentRequest {
	s.AgentId = &v
	return s
}

func (s *InstallAgentRequest) SetAgentVersion(v string) *InstallAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *InstallAgentRequest) SetInstallType(v string) *InstallAgentRequest {
	s.InstallType = &v
	return s
}

func (s *InstallAgentRequest) SetInstances(v []*InstallAgentRequestInstances) *InstallAgentRequest {
	s.Instances = v
	return s
}

type InstallAgentRequestInstances struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s InstallAgentRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *InstallAgentRequestInstances) SetInstance(v string) *InstallAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *InstallAgentRequestInstances) SetRegion(v string) *InstallAgentRequestInstances {
	s.Region = &v
	return s
}

type InstallAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *InstallAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InstallAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBody) SetRequestId(v string) *InstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentResponseBody) SetCode(v string) *InstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentResponseBody) SetData(v *InstallAgentResponseBodyData) *InstallAgentResponseBody {
	s.Data = v
	return s
}

func (s *InstallAgentResponseBody) SetMessage(v string) *InstallAgentResponseBody {
	s.Message = &v
	return s
}

type InstallAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InstallAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAgentResponseBodyData) SetTaskId(v string) *InstallAgentResponseBodyData {
	s.TaskId = &v
	return s
}

type InstallAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentResponse) SetHeaders(v map[string]*string) *InstallAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentResponse) SetStatusCode(v int32) *InstallAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAgentResponse) SetBody(v *InstallAgentResponseBody) *InstallAgentResponse {
	s.Body = v
	return s
}

type InvokeAnomalyDiagnosisRequest struct {
	// example:
	//
	// 8047d763-5465-4a8c-b1cd-23f5a8ba2594
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s InvokeAnomalyDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeAnomalyDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisRequest) SetUuid(v string) *InvokeAnomalyDiagnosisRequest {
	s.Uuid = &v
	return s
}

type InvokeAnomalyDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InvokeAnomalyDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeAnomalyDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetCode(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetMessage(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponseBody) SetRequestId(v string) *InvokeAnomalyDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type InvokeAnomalyDiagnosisResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeAnomalyDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeAnomalyDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeAnomalyDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeAnomalyDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeAnomalyDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeAnomalyDiagnosisResponse) SetStatusCode(v int32) *InvokeAnomalyDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeAnomalyDiagnosisResponse) SetBody(v *InvokeAnomalyDiagnosisResponseBody) *InvokeAnomalyDiagnosisResponse {
	s.Body = v
	return s
}

type InvokeDiagnosisRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cloud_assist
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "instance": "i-wz9gdv7qmdhusamc4dl01",
	//
	//     "uid": "xxxxxxxxxxxxxx",
	//
	//     "region": "cn-shenzhen"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// memgraph
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s InvokeDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisRequest) SetChannel(v string) *InvokeDiagnosisRequest {
	s.Channel = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetParams(v string) *InvokeDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetServiceName(v string) *InvokeDiagnosisRequest {
	s.ServiceName = &v
	return s
}

type InvokeDiagnosisResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *InvokeDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s InvokeDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBody) SetCode(v string) *InvokeDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetData(v *InvokeDiagnosisResponseBodyData) *InvokeDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetMessage(v string) *InvokeDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeDiagnosisResponseBody) SetRequestId(v string) *InvokeDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type InvokeDiagnosisResponseBodyData struct {
	// example:
	//
	// ihqhAcrt
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InvokeDiagnosisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponseBodyData) SetTaskId(v string) *InvokeDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

type InvokeDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisResponse) SetHeaders(v map[string]*string) *InvokeDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *InvokeDiagnosisResponse) SetStatusCode(v int32) *InvokeDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeDiagnosisResponse) SetBody(v *InvokeDiagnosisResponseBody) *InvokeDiagnosisResponse {
	s.Body = v
	return s
}

type ListAbnormalyEventsRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// potential
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// example:
	//
	// 1
	ShowPod *int32 `json:"showPod,omitempty" xml:"showPod,omitempty"`
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListAbnormalyEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAbnormalyEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsRequest) SetCluster(v string) *ListAbnormalyEventsRequest {
	s.Cluster = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetCurrent(v int32) *ListAbnormalyEventsRequest {
	s.Current = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetEnd(v float32) *ListAbnormalyEventsRequest {
	s.End = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetInstance(v string) *ListAbnormalyEventsRequest {
	s.Instance = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetLevel(v string) *ListAbnormalyEventsRequest {
	s.Level = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetNamespace(v string) *ListAbnormalyEventsRequest {
	s.Namespace = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetPageSize(v int32) *ListAbnormalyEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetPod(v string) *ListAbnormalyEventsRequest {
	s.Pod = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetShowPod(v int32) *ListAbnormalyEventsRequest {
	s.ShowPod = &v
	return s
}

func (s *ListAbnormalyEventsRequest) SetStart(v float32) *ListAbnormalyEventsRequest {
	s.Start = &v
	return s
}

type ListAbnormalyEventsResponseBody struct {
	// example:
	//
	// Success
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data    []*ListAbnormalyEventsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListAbnormalyEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAbnormalyEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBody) SetCode(v string) *ListAbnormalyEventsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAbnormalyEventsResponseBody) SetData(v []*ListAbnormalyEventsResponseBodyData) *ListAbnormalyEventsResponseBody {
	s.Data = v
	return s
}

func (s *ListAbnormalyEventsResponseBody) SetMessage(v string) *ListAbnormalyEventsResponseBody {
	s.Message = &v
	return s
}

type ListAbnormalyEventsResponseBodyData struct {
	// example:
	//
	// 1725801090000
	CreatedAt   *float32 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string  `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string                                  `json:"instance,omitempty" xml:"instance,omitempty"`
	Item     *string                                  `json:"item,omitempty" xml:"item,omitempty"`
	Opts     *ListAbnormalyEventsResponseBodyDataOpts `json:"opts,omitempty" xml:"opts,omitempty" type:"Struct"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// example:
	//
	// saturation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyData) SetCreatedAt(v float32) *ListAbnormalyEventsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetDescription(v string) *ListAbnormalyEventsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetId(v string) *ListAbnormalyEventsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetInstance(v string) *ListAbnormalyEventsResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetItem(v string) *ListAbnormalyEventsResponseBodyData {
	s.Item = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetOpts(v *ListAbnormalyEventsResponseBodyDataOpts) *ListAbnormalyEventsResponseBodyData {
	s.Opts = v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetRegionId(v string) *ListAbnormalyEventsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyData) SetType(v string) *ListAbnormalyEventsResponseBodyData {
	s.Type = &v
	return s
}

type ListAbnormalyEventsResponseBodyDataOpts struct {
	// example:
	//
	// diagnose
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// "{\\"service_name\\": \\"oomcheck\\", \\"params\\": {\\"auto_initial\\": true, \\"instance\\": \\"i-wz9d00ut2ska3mlyhn6i\\", \\"region\\": \\"cn-shenzhen\\", \\"uuid\\": \\"24576d0c-a19d-49dd-8a64-3867440fd7a6\\", \\"is_history\\": 1}}"
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// example:
	//
	// realtime
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAbnormalyEventsResponseBodyDataOpts) String() string {
	return tea.Prettify(s)
}

func (s ListAbnormalyEventsResponseBodyDataOpts) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetLabel(v string) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Label = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetParams(v string) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Params = &v
	return s
}

func (s *ListAbnormalyEventsResponseBodyDataOpts) SetType(v string) *ListAbnormalyEventsResponseBodyDataOpts {
	s.Type = &v
	return s
}

type ListAbnormalyEventsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAbnormalyEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAbnormalyEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAbnormalyEventsResponse) GoString() string {
	return s.String()
}

func (s *ListAbnormalyEventsResponse) SetHeaders(v map[string]*string) *ListAbnormalyEventsResponse {
	s.Headers = v
	return s
}

func (s *ListAbnormalyEventsResponse) SetStatusCode(v int32) *ListAbnormalyEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAbnormalyEventsResponse) SetBody(v *ListAbnormalyEventsResponseBody) *ListAbnormalyEventsResponse {
	s.Body = v
	return s
}

type ListAgentInstallRecordsRequest struct {
	Current       *int64  `json:"current,omitempty" xml:"current,omitempty"`
	InstanceId    *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	PageSize      *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PluginId      *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentInstallRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentInstallRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsRequest) SetCurrent(v int64) *ListAgentInstallRecordsRequest {
	s.Current = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetInstanceId(v string) *ListAgentInstallRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPageSize(v int64) *ListAgentInstallRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPluginId(v string) *ListAgentInstallRecordsRequest {
	s.PluginId = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetPluginVersion(v string) *ListAgentInstallRecordsRequest {
	s.PluginVersion = &v
	return s
}

func (s *ListAgentInstallRecordsRequest) SetStatus(v string) *ListAgentInstallRecordsRequest {
	s.Status = &v
	return s
}

type ListAgentInstallRecordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAgentInstallRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 64
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentInstallRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentInstallRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponseBody) SetRequestId(v string) *ListAgentInstallRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetCode(v string) *ListAgentInstallRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetData(v []*ListAgentInstallRecordsResponseBodyData) *ListAgentInstallRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetMessage(v string) *ListAgentInstallRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBody) SetTotal(v int64) *ListAgentInstallRecordsResponseBody {
	s.Total = &v
	return s
}

type ListAgentInstallRecordsResponseBodyData struct {
	CreatedAt     *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	InstanceId    *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	PluginId      *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt     *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListAgentInstallRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentInstallRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponseBodyData) SetCreatedAt(v string) *ListAgentInstallRecordsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetInstanceId(v string) *ListAgentInstallRecordsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetPluginId(v string) *ListAgentInstallRecordsResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetPluginVersion(v string) *ListAgentInstallRecordsResponseBodyData {
	s.PluginVersion = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetStatus(v string) *ListAgentInstallRecordsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentInstallRecordsResponseBodyData) SetUpdatedAt(v string) *ListAgentInstallRecordsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type ListAgentInstallRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentInstallRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentInstallRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentInstallRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentInstallRecordsResponse) SetHeaders(v map[string]*string) *ListAgentInstallRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentInstallRecordsResponse) SetStatusCode(v int32) *ListAgentInstallRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentInstallRecordsResponse) SetBody(v *ListAgentInstallRecordsResponseBody) *ListAgentInstallRecordsResponse {
	s.Body = v
	return s
}

type ListAgentsRequest struct {
	Current  *int64  `json:"current,omitempty" xml:"current,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	PageSize *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) SetCurrent(v int64) *ListAgentsRequest {
	s.Current = &v
	return s
}

func (s *ListAgentsRequest) SetName(v string) *ListAgentsRequest {
	s.Name = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int64) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetType(v string) *ListAgentsRequest {
	s.Type = &v
	return s
}

type ListAgentsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAgentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Total   *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetCode(v string) *ListAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentsResponseBody) SetData(v []*ListAgentsResponseBodyData) *ListAgentsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentsResponseBody) SetMessage(v string) *ListAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentsResponseBody) SetTotal(v int64) *ListAgentsResponseBody {
	s.Total = &v
	return s
}

type ListAgentsResponseBodyData struct {
	CreatedAt   *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// SysOM Agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// x86
	SupportArch *string `json:"support_arch,omitempty" xml:"support_arch,omitempty"`
	// example:
	//
	// Control
	Type      *string                               `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt *string                               `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Versions  []*ListAgentsResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListAgentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyData) SetCreatedAt(v string) *ListAgentsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetDescription(v string) *ListAgentsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetId(v string) *ListAgentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetName(v string) *ListAgentsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetSupportArch(v string) *ListAgentsResponseBodyData {
	s.SupportArch = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetType(v string) *ListAgentsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetUpdatedAt(v string) *ListAgentsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetVersions(v []*ListAgentsResponseBodyDataVersions) *ListAgentsResponseBodyData {
	s.Versions = v
	return s
}

type ListAgentsResponseBodyDataVersions struct {
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	InstallScript   *string `json:"install_script,omitempty" xml:"install_script,omitempty"`
	UninstallScript *string `json:"uninstall_script,omitempty" xml:"uninstall_script,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UpgradeScript   *string `json:"upgrade_script,omitempty" xml:"upgrade_script,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAgentsResponseBodyDataVersions) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyDataVersions) SetCreatedAt(v string) *ListAgentsResponseBodyDataVersions {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetInstallScript(v string) *ListAgentsResponseBodyDataVersions {
	s.InstallScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUninstallScript(v string) *ListAgentsResponseBodyDataVersions {
	s.UninstallScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUpdatedAt(v string) *ListAgentsResponseBodyDataVersions {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUpgradeScript(v string) *ListAgentsResponseBodyDataVersions {
	s.UpgradeScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetVersion(v string) *ListAgentsResponseBodyDataVersions {
	s.Version = &v
	return s
}

type ListAgentsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentsResponse) SetHeaders(v map[string]*string) *ListAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentsResponse) SetStatusCode(v int32) *ListAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentsResponse) SetBody(v *ListAgentsResponseBody) *ListAgentsResponse {
	s.Body = v
	return s
}

type ListInstanceHealthRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725801327754
	End *float32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725797727754
	Start *float32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListInstanceHealthRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceHealthRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthRequest) SetCluster(v string) *ListInstanceHealthRequest {
	s.Cluster = &v
	return s
}

func (s *ListInstanceHealthRequest) SetCurrent(v int32) *ListInstanceHealthRequest {
	s.Current = &v
	return s
}

func (s *ListInstanceHealthRequest) SetEnd(v float32) *ListInstanceHealthRequest {
	s.End = &v
	return s
}

func (s *ListInstanceHealthRequest) SetInstance(v string) *ListInstanceHealthRequest {
	s.Instance = &v
	return s
}

func (s *ListInstanceHealthRequest) SetPageSize(v int32) *ListInstanceHealthRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceHealthRequest) SetStart(v float32) *ListInstanceHealthRequest {
	s.Start = &v
	return s
}

type ListInstanceHealthResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListInstanceHealthResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Query no data
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 35F91AAB-5FDF-5A22-B211-C7C6B00817D0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 42
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceHealthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponseBody) SetCode(v string) *ListInstanceHealthResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetData(v *ListInstanceHealthResponseBodyData) *ListInstanceHealthResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceHealthResponseBody) SetMessage(v string) *ListInstanceHealthResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetRequestId(v string) *ListInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceHealthResponseBody) SetTotal(v int32) *ListInstanceHealthResponseBody {
	s.Total = &v
	return s
}

type ListInstanceHealthResponseBodyData struct {
	Images []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// example:
	//
	// 100
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceHealthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceHealthResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponseBodyData) SetImages(v []*string) *ListInstanceHealthResponseBodyData {
	s.Images = v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetInstance(v string) *ListInstanceHealthResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetNamespace(v string) *ListInstanceHealthResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetPod(v string) *ListInstanceHealthResponseBodyData {
	s.Pod = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetRegionId(v string) *ListInstanceHealthResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetScore(v float32) *ListInstanceHealthResponseBodyData {
	s.Score = &v
	return s
}

func (s *ListInstanceHealthResponseBodyData) SetStatus(v string) *ListInstanceHealthResponseBodyData {
	s.Status = &v
	return s
}

type ListInstanceHealthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceHealthResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceHealthResponse) SetHeaders(v map[string]*string) *ListInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceHealthResponse) SetStatusCode(v int32) *ListInstanceHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceHealthResponse) SetBody(v *ListInstanceHealthResponseBody) *ListInstanceHealthResponse {
	s.Body = v
	return s
}

type StartAIAnalysisRequest struct {
	// example:
	//
	// ecs_sysom
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// python_test
	Comms *string `json:"comms,omitempty" xml:"comms,omitempty"`
	// example:
	//
	// i-wz9dej066kii4goqxxxx
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 2421,36547,10043
	Pids *string `json:"pids,omitempty" xml:"pids,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 2000
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s StartAIAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAIAnalysisRequest) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisRequest) SetChannel(v string) *StartAIAnalysisRequest {
	s.Channel = &v
	return s
}

func (s *StartAIAnalysisRequest) SetComms(v string) *StartAIAnalysisRequest {
	s.Comms = &v
	return s
}

func (s *StartAIAnalysisRequest) SetInstance(v string) *StartAIAnalysisRequest {
	s.Instance = &v
	return s
}

func (s *StartAIAnalysisRequest) SetPids(v string) *StartAIAnalysisRequest {
	s.Pids = &v
	return s
}

func (s *StartAIAnalysisRequest) SetRegion(v string) *StartAIAnalysisRequest {
	s.Region = &v
	return s
}

func (s *StartAIAnalysisRequest) SetTimeout(v int32) *StartAIAnalysisRequest {
	s.Timeout = &v
	return s
}

type StartAIAnalysisResponseBody struct {
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *StartAIAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ""
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartAIAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAIAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponseBody) SetCode(v string) *StartAIAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *StartAIAnalysisResponseBody) SetData(v *StartAIAnalysisResponseBodyData) *StartAIAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *StartAIAnalysisResponseBody) SetMessage(v string) *StartAIAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *StartAIAnalysisResponseBody) SetRequestId(v string) *StartAIAnalysisResponseBody {
	s.RequestId = &v
	return s
}

type StartAIAnalysisResponseBodyData struct {
	// example:
	//
	// 16896fa8-37f6-4c70-bb32-67fa9817d426
	AnalysisId *string `json:"analysis_id,omitempty" xml:"analysis_id,omitempty"`
}

func (s StartAIAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartAIAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponseBodyData) SetAnalysisId(v string) *StartAIAnalysisResponseBodyData {
	s.AnalysisId = &v
	return s
}

type StartAIAnalysisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAIAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAIAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAIAnalysisResponse) GoString() string {
	return s.String()
}

func (s *StartAIAnalysisResponse) SetHeaders(v map[string]*string) *StartAIAnalysisResponse {
	s.Headers = v
	return s
}

func (s *StartAIAnalysisResponse) SetStatusCode(v int32) *StartAIAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAIAnalysisResponse) SetBody(v *StartAIAnalysisResponseBody) *StartAIAnalysisResponse {
	s.Body = v
	return s
}

type UninstallAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	Instances []*UninstallAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UninstallAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentRequest) SetAgentId(v string) *UninstallAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentRequest) SetAgentVersion(v string) *UninstallAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentRequest) SetInstances(v []*UninstallAgentRequestInstances) *UninstallAgentRequest {
	s.Instances = v
	return s
}

type UninstallAgentRequestInstances struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s UninstallAgentRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *UninstallAgentRequestInstances) SetInstance(v string) *UninstallAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *UninstallAgentRequestInstances) SetRegion(v string) *UninstallAgentRequestInstances {
	s.Region = &v
	return s
}

type UninstallAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *UninstallAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UninstallAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponseBody) SetRequestId(v string) *UninstallAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAgentResponseBody) SetCode(v string) *UninstallAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallAgentResponseBody) SetData(v *UninstallAgentResponseBodyData) *UninstallAgentResponseBody {
	s.Data = v
	return s
}

func (s *UninstallAgentResponseBody) SetMessage(v string) *UninstallAgentResponseBody {
	s.Message = &v
	return s
}

type UninstallAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UninstallAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponseBodyData) SetTaskId(v string) *UninstallAgentResponseBodyData {
	s.TaskId = &v
	return s
}

type UninstallAgentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentResponse) GoString() string {
	return s.String()
}

func (s *UninstallAgentResponse) SetHeaders(v map[string]*string) *UninstallAgentResponse {
	s.Headers = v
	return s
}

func (s *UninstallAgentResponse) SetStatusCode(v int32) *UninstallAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAgentResponse) SetBody(v *UninstallAgentResponseBody) *UninstallAgentResponse {
	s.Body = v
	return s
}

type UpgradeAgentRequest struct {
	// This parameter is required.
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// This parameter is required.
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// This parameter is required.
	Instances []*UpgradeAgentRequestInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
}

func (s UpgradeAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentRequest) SetAgentId(v string) *UpgradeAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentRequest) SetAgentVersion(v string) *UpgradeAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentRequest) SetInstances(v []*UpgradeAgentRequestInstances) *UpgradeAgentRequest {
	s.Instances = v
	return s
}

type UpgradeAgentRequestInstances struct {
	// This parameter is required.
	//
	// example:
	//
	// i-wz9b9vucz1iubsz8sjqo
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s UpgradeAgentRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentRequestInstances) GoString() string {
	return s.String()
}

func (s *UpgradeAgentRequestInstances) SetInstance(v string) *UpgradeAgentRequestInstances {
	s.Instance = &v
	return s
}

func (s *UpgradeAgentRequestInstances) SetRegion(v string) *UpgradeAgentRequestInstances {
	s.Region = &v
	return s
}

type UpgradeAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpgradeAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpgradeAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponseBody) SetRequestId(v string) *UpgradeAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAgentResponseBody) SetCode(v string) *UpgradeAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAgentResponseBody) SetData(v *UpgradeAgentResponseBodyData) *UpgradeAgentResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeAgentResponseBody) SetMessage(v string) *UpgradeAgentResponseBody {
	s.Message = &v
	return s
}

type UpgradeAgentResponseBodyData struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponseBodyData) SetTaskId(v string) *UpgradeAgentResponseBodyData {
	s.TaskId = &v
	return s
}

type UpgradeAgentResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAgentResponse) SetHeaders(v map[string]*string) *UpgradeAgentResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAgentResponse) SetStatusCode(v int32) *UpgradeAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAgentResponse) SetBody(v *UpgradeAgentResponseBody) *UpgradeAgentResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("sysom"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 授权 SysOM 对某个机器进行诊断
//
// @param request - AuthDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosisWithOptions(request *AuthDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AuthDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRole)) {
		body["autoCreateRole"] = request.AutoCreateRole
	}

	if !tea.BoolValue(util.IsUnset(request.AutoInstallAgent)) {
		body["autoInstallAgent"] = request.AutoInstallAgent
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/auth"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权 SysOM 对某个机器进行诊断
//
// @param request - AuthDiagnosisRequest
//
// @return AuthDiagnosisResponse
func (client *Client) AuthDiagnosis(request *AuthDiagnosisRequest) (_result *AuthDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthDiagnosisResponse{}
	_body, _err := client.AuthDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取copilot服务的返回结果
//
// @param request - GenerateCopilotResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponseWithOptions(request *GenerateCopilotResponseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateCopilotResponseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LlmParamString)) {
		body["llmParamString"] = request.LlmParamString
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCopilotResponse"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/copilot/generate_copilot_response"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取copilot服务的返回结果
//
// @param request - GenerateCopilotResponseRequest
//
// @return GenerateCopilotResponseResponse
func (client *Client) GenerateCopilotResponse(request *GenerateCopilotResponseRequest) (_result *GenerateCopilotResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotResponseResponse{}
	_body, _err := client.GenerateCopilotResponseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看AI Infra分析结果
//
// @param request - GetAIQueryResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAIQueryResultResponse
func (client *Client) GetAIQueryResultWithOptions(request *GetAIQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAIQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalysisId)) {
		body["analysisId"] = request.AnalysisId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAIQueryResult"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/app_observ/aiAnalysis/query_result"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAIQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看AI Infra分析结果
//
// @param request - GetAIQueryResultRequest
//
// @return GetAIQueryResultResponse
func (client *Client) GetAIQueryResult(request *GetAIQueryResultRequest) (_result *GetAIQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAIQueryResultResponse{}
	_body, _err := client.GetAIQueryResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点/Pod不同等级异常事件的数量
//
// @param request - GetAbnormalEventsCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCountWithOptions(request *GetAbnormalEventsCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAbnormalEventsCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Pod)) {
		query["pod"] = request.Pod
	}

	if !tea.BoolValue(util.IsUnset(request.ShowPod)) {
		query["showPod"] = request.ShowPod
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAbnormalEventsCount"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/abnormaly_events_count"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点/Pod不同等级异常事件的数量
//
// @param request - GetAbnormalEventsCountRequest
//
// @return GetAbnormalEventsCountResponse
func (client *Client) GetAbnormalEventsCount(request *GetAbnormalEventsCountRequest) (_result *GetAbnormalEventsCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAbnormalEventsCountResponse{}
	_body, _err := client.GetAbnormalEventsCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某个组件的详情
//
// @param request - GetAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithOptions(request *GetAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["agent_id"] = request.AgentId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgent"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/get_agent"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个组件的详情
//
// @param request - GetAgentRequest
//
// @return GetAgentResponse
func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Agent安装任务执行状态
//
// @param request - GetAgentTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentTaskResponse
func (client *Client) GetAgentTaskWithOptions(request *GetAgentTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAgentTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentTask"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/get_agent_task"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agent安装任务执行状态
//
// @param request - GetAgentTaskRequest
//
// @return GetAgentTaskResponse
func (client *Client) GetAgentTask(request *GetAgentTaskRequest) (_result *GetAgentTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentTaskResponse{}
	_body, _err := client.GetAgentTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取诊断结果
//
// @param request - GetDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResultWithOptions(request *GetDiagnosisResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDiagnosisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisResult"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/get_diagnosis_results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取诊断结果
//
// @param request - GetDiagnosisResultRequest
//
// @return GetDiagnosisResultResponse
func (client *Client) GetDiagnosisResult(request *GetDiagnosisResultRequest) (_result *GetDiagnosisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDiagnosisResultResponse{}
	_body, _err := client.GetDiagnosisResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一段时间的节点/pod健康度比例
//
// @param request - GetHealthPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentageWithOptions(request *GetHealthPercentageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHealthPercentageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHealthPercentage"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/health_percentage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一段时间的节点/pod健康度比例
//
// @param request - GetHealthPercentageRequest
//
// @return GetHealthPercentageResponse
func (client *Client) GetHealthPercentage(request *GetHealthPercentageRequest) (_result *GetHealthPercentageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHealthPercentageResponse{}
	_body, _err := client.GetHealthPercentageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群节点数量
//
// @param request - GetHostCountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHostCountResponse
func (client *Client) GetHostCountWithOptions(request *GetHostCountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHostCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostCount"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/host_count"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群节点数量
//
// @param request - GetHostCountRequest
//
// @return GetHostCountResponse
func (client *Client) GetHostCount(request *GetHostCountRequest) (_result *GetHostCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHostCountResponse{}
	_body, _err := client.GetHostCountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热定分析结果
//
// @param request - GetHotspotAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotAnalysisResponse
func (client *Client) GetHotspotAnalysisWithOptions(request *GetHotspotAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotspotAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["appType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.BegEnd)) {
		body["beg_end"] = request.BegEnd
	}

	if !tea.BoolValue(util.IsUnset(request.BegStart)) {
		body["beg_start"] = request.BegStart
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		body["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		body["pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["table"] = request.Table
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotAnalysis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/livetrace_hotspot_analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热定分析结果
//
// @param request - GetHotspotAnalysisRequest
//
// @return GetHotspotAnalysisResponse
func (client *Client) GetHotspotAnalysis(request *GetHotspotAnalysisRequest) (_result *GetHotspotAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotAnalysisResponse{}
	_body, _err := client.GetHotspotAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热点实例列表
//
// @param request - GetHotspotInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotInstanceListResponse
func (client *Client) GetHotspotInstanceListWithOptions(request *GetHotspotInstanceListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotspotInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BegEnd)) {
		body["beg_end"] = request.BegEnd
	}

	if !tea.BoolValue(util.IsUnset(request.BegStart)) {
		body["beg_start"] = request.BegStart
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["table"] = request.Table
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotInstanceList"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/livetrace_hotspot_instance_list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热点实例列表
//
// @param request - GetHotspotInstanceListRequest
//
// @return GetHotspotInstanceListResponse
func (client *Client) GetHotspotInstanceList(request *GetHotspotInstanceListRequest) (_result *GetHotspotInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotInstanceListResponse{}
	_body, _err := client.GetHotspotInstanceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某个实例的pid列表
//
// @param request - GetHotspotPidListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotPidListResponse
func (client *Client) GetHotspotPidListWithOptions(request *GetHotspotPidListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotspotPidListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BegEnd)) {
		body["beg_end"] = request.BegEnd
	}

	if !tea.BoolValue(util.IsUnset(request.BegStart)) {
		body["beg_start"] = request.BegStart
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		body["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["table"] = request.Table
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotPidList"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/livetrace_hotspot_pid_list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotspotPidListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个实例的pid列表
//
// @param request - GetHotspotPidListRequest
//
// @return GetHotspotPidListResponse
func (client *Client) GetHotspotPidList(request *GetHotspotPidListRequest) (_result *GetHotspotPidListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotPidListResponse{}
	_body, _err := client.GetHotspotPidListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// AI Infra获取分析记录列表
//
// @param request - GetListRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetListRecordResponse
func (client *Client) GetListRecordWithOptions(request *GetListRecordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetListRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetListRecord"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/app_observ/aiAnalysis/list_record"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetListRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI Infra获取分析记录列表
//
// @param request - GetListRecordRequest
//
// @return GetListRecordResponse
func (client *Client) GetListRecord(request *GetListRecordRequest) (_result *GetListRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetListRecordResponse{}
	_body, _err := client.GetListRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一定时间内集群中节点/节点中pod异常问题占比
//
// @param request - GetProblemPercentageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProblemPercentageResponse
func (client *Client) GetProblemPercentageWithOptions(request *GetProblemPercentageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProblemPercentageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProblemPercentage"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/problem_percentage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProblemPercentageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一定时间内集群中节点/节点中pod异常问题占比
//
// @param request - GetProblemPercentageRequest
//
// @return GetProblemPercentageResponse
func (client *Client) GetProblemPercentage(request *GetProblemPercentageRequest) (_result *GetProblemPercentageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProblemPercentageResponse{}
	_body, _err := client.GetProblemPercentageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取健康分趋势
//
// @param request - GetRangeScoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRangeScoreResponse
func (client *Client) GetRangeScoreWithOptions(request *GetRangeScoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRangeScoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRangeScore"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/score"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRangeScoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取健康分趋势
//
// @param request - GetRangeScoreRequest
//
// @return GetRangeScoreResponse
func (client *Client) GetRangeScore(request *GetRangeScoreRequest) (_result *GetRangeScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRangeScoreResponse{}
	_body, _err := client.GetRangeScoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群/节点资源实时使用情况
//
// @param request - GetResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithOptions(request *GetResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResources"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/instant/resource"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群/节点资源实时使用情况
//
// @param request - GetResourcesRequest
//
// @return GetResourcesResponse
func (client *Client) GetResources(request *GetResourcesRequest) (_result *GetResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourcesResponse{}
	_body, _err := client.GetResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 在指定的实例上安装 Agent
//
// @param request - InstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentResponse
func (client *Client) InstallAgentWithOptions(request *InstallAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agent_id"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentVersion)) {
		body["agent_version"] = request.AgentVersion
	}

	if !tea.BoolValue(util.IsUnset(request.InstallType)) {
		body["install_type"] = request.InstallType
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallAgent"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/install_agent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在指定的实例上安装 Agent
//
// @param request - InstallAgentRequest
//
// @return InstallAgentResponse
func (client *Client) InstallAgent(request *InstallAgentRequest) (_result *InstallAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentResponse{}
	_body, _err := client.InstallAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异常项诊断跳转
//
// @param request - InvokeAnomalyDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeAnomalyDiagnosisResponse
func (client *Client) InvokeAnomalyDiagnosisWithOptions(request *InvokeAnomalyDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokeAnomalyDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeAnomalyDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/diagnosis/invoke_anomaly_diagnose"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeAnomalyDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 异常项诊断跳转
//
// @param request - InvokeAnomalyDiagnosisRequest
//
// @return InvokeAnomalyDiagnosisResponse
func (client *Client) InvokeAnomalyDiagnosis(request *InvokeAnomalyDiagnosisRequest) (_result *InvokeAnomalyDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeAnomalyDiagnosisResponse{}
	_body, _err := client.InvokeAnomalyDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发起诊断
//
// @param request - InvokeDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosisWithOptions(request *InvokeDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvokeDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		body["service_name"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/invoke_diagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发起诊断
//
// @param request - InvokeDiagnosisRequest
//
// @return InvokeDiagnosisResponse
func (client *Client) InvokeDiagnosis(request *InvokeDiagnosisRequest) (_result *InvokeDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvokeDiagnosisResponse{}
	_body, _err := client.InvokeDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一定时间段内的异常事件
//
// @param request - ListAbnormalyEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAbnormalyEventsResponse
func (client *Client) ListAbnormalyEventsWithOptions(request *ListAbnormalyEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAbnormalyEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Level)) {
		query["level"] = request.Level
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Pod)) {
		query["pod"] = request.Pod
	}

	if !tea.BoolValue(util.IsUnset(request.ShowPod)) {
		query["showPod"] = request.ShowPod
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAbnormalyEvents"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/abnormaly_events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAbnormalyEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一定时间段内的异常事件
//
// @param request - ListAbnormalyEventsRequest
//
// @return ListAbnormalyEventsResponse
func (client *Client) ListAbnormalyEvents(request *ListAbnormalyEventsRequest) (_result *ListAbnormalyEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAbnormalyEventsResponse{}
	_body, _err := client.ListAbnormalyEventsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出 Agent 的安装记录
//
// @param request - ListAgentInstallRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentInstallRecordsResponse
func (client *Client) ListAgentInstallRecordsWithOptions(request *ListAgentInstallRecordsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAgentInstallRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instance_id"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PluginId)) {
		query["plugin_id"] = request.PluginId
	}

	if !tea.BoolValue(util.IsUnset(request.PluginVersion)) {
		query["plugin_version"] = request.PluginVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgentInstallRecords"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/list_agent_install_list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAgentInstallRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出 Agent 的安装记录
//
// @param request - ListAgentInstallRecordsRequest
//
// @return ListAgentInstallRecordsResponse
func (client *Client) ListAgentInstallRecords(request *ListAgentInstallRecordsRequest) (_result *ListAgentInstallRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentInstallRecordsResponse{}
	_body, _err := client.ListAgentInstallRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 Agent 列表
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithOptions(request *ListAgentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgents"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/list_agents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 Agent 列表
//
// @param request - ListAgentsRequest
//
// @return ListAgentsResponse
func (client *Client) ListAgents(request *ListAgentsRequest) (_result *ListAgentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentsResponse{}
	_body, _err := client.ListAgentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取一定时间内集群节点/Pod健康度列表
//
// @param request - ListInstanceHealthRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHealthResponse
func (client *Client) ListInstanceHealthWithOptions(request *ListInstanceHealthRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceHealthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		query["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.End)) {
		query["end"] = request.End
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Start)) {
		query["start"] = request.Start
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceHealth"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/range/instance_health_list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceHealthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取一定时间内集群节点/Pod健康度列表
//
// @param request - ListInstanceHealthRequest
//
// @return ListInstanceHealthResponse
func (client *Client) ListInstanceHealth(request *ListInstanceHealthRequest) (_result *ListInstanceHealthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceHealthResponse{}
	_body, _err := client.ListInstanceHealthWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动AI作业分析
//
// @param request - StartAIAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIAnalysisResponse
func (client *Client) StartAIAnalysisWithOptions(request *StartAIAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartAIAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		body["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.Comms)) {
		body["comms"] = request.Comms
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		body["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.Pids)) {
		body["pids"] = request.Pids
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAIAnalysis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/start_ai_analysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAIAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动AI作业分析
//
// @param request - StartAIAnalysisRequest
//
// @return StartAIAnalysisResponse
func (client *Client) StartAIAnalysis(request *StartAIAnalysisRequest) (_result *StartAIAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAIAnalysisResponse{}
	_body, _err := client.StartAIAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载 SysOM Agent
//
// @param request - UninstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentResponse
func (client *Client) UninstallAgentWithOptions(request *UninstallAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UninstallAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agent_id"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentVersion)) {
		body["agent_version"] = request.AgentVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallAgent"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/uninstall_agent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载 SysOM Agent
//
// @param request - UninstallAgentRequest
//
// @return UninstallAgentResponse
func (client *Client) UninstallAgent(request *UninstallAgentRequest) (_result *UninstallAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallAgentResponse{}
	_body, _err := client.UninstallAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 SysOM Agent
//
// @param request - UpgradeAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentResponse
func (client *Client) UpgradeAgentWithOptions(request *UpgradeAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		body["agent_id"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.AgentVersion)) {
		body["agent_version"] = request.AgentVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		body["instances"] = request.Instances
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeAgent"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/upgrade_agent"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 SysOM Agent
//
// @param request - UpgradeAgentRequest
//
// @return UpgradeAgentResponse
func (client *Client) UpgradeAgent(request *UpgradeAgentRequest) (_result *UpgradeAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeAgentResponse{}
	_body, _err := client.UpgradeAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
