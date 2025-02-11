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

type GenerateCopilotStreamResponseRequest struct {
	LlmParamString *string `json:"llmParamString,omitempty" xml:"llmParamString,omitempty"`
}

func (s GenerateCopilotStreamResponseRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotStreamResponseRequest) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseRequest) SetLlmParamString(v string) *GenerateCopilotStreamResponseRequest {
	s.LlmParamString = &v
	return s
}

type GenerateCopilotStreamResponseResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GenerateCopilotStreamResponseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotStreamResponseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseResponseBody) SetCode(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetData(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetMessage(v string) *GenerateCopilotStreamResponseResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponseBody) SetRequestId(v string) *GenerateCopilotStreamResponseResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCopilotStreamResponseResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCopilotStreamResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCopilotStreamResponseResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCopilotStreamResponseResponse) GoString() string {
	return s.String()
}

func (s *GenerateCopilotStreamResponseResponse) SetHeaders(v map[string]*string) *GenerateCopilotStreamResponseResponse {
	s.Headers = v
	return s
}

func (s *GenerateCopilotStreamResponseResponse) SetStatusCode(v int32) *GenerateCopilotStreamResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCopilotStreamResponseResponse) SetBody(v *GenerateCopilotStreamResponseResponseBody) *GenerateCopilotStreamResponseResponse {
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

type GetCopilotHistoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
}

func (s GetCopilotHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCopilotHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryRequest) SetCount(v int64) *GetCopilotHistoryRequest {
	s.Count = &v
	return s
}

type GetCopilotHistoryResponseBody struct {
	// example:
	//
	// SysomOpenAPI.InvalidParameter
	Code *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Data []*GetCopilotHistoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCopilotHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCopilotHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponseBody) SetCode(v string) *GetCopilotHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetData(v []*GetCopilotHistoryResponseBodyData) *GetCopilotHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetMessage(v string) *GetCopilotHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetCopilotHistoryResponseBody) SetRequestId(v string) *GetCopilotHistoryResponseBody {
	s.RequestId = &v
	return s
}

type GetCopilotHistoryResponseBodyData struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-02 10:02:39
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// user
	//
	// copilot
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetCopilotHistoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCopilotHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponseBodyData) SetContent(v string) *GetCopilotHistoryResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetCopilotHistoryResponseBodyData) SetTime(v string) *GetCopilotHistoryResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetCopilotHistoryResponseBodyData) SetUser(v string) *GetCopilotHistoryResponseBodyData {
	s.User = &v
	return s
}

type GetCopilotHistoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCopilotHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponse) SetHeaders(v map[string]*string) *GetCopilotHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotHistoryResponse) SetStatusCode(v int32) *GetCopilotHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotHistoryResponse) SetBody(v *GetCopilotHistoryResponseBody) *GetCopilotHistoryResponse {
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

type GetHotSpotUniqListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1735010351000
	BegEnd *int64 `json:"beg_end,omitempty" xml:"beg_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735008551000
	BegStart *int64 `json:"beg_start,omitempty" xml:"beg_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g2i0siirefgwnnnvy
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 12345
	Pid *int64 `json:"pid,omitempty" xml:"pid,omitempty"`
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// This parameter is required.
	Uniq *string `json:"uniq,omitempty" xml:"uniq,omitempty"`
}

func (s GetHotSpotUniqListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotSpotUniqListRequest) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListRequest) SetBegEnd(v int64) *GetHotSpotUniqListRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetBegStart(v int64) *GetHotSpotUniqListRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetInstance(v string) *GetHotSpotUniqListRequest {
	s.Instance = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetPid(v int64) *GetHotSpotUniqListRequest {
	s.Pid = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetTable(v string) *GetHotSpotUniqListRequest {
	s.Table = &v
	return s
}

func (s *GetHotSpotUniqListRequest) SetUniq(v string) *GetHotSpotUniqListRequest {
	s.Uniq = &v
	return s
}

type GetHotSpotUniqListResponseBody struct {
	// example:
	//
	// Success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotSpotUniqListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotSpotUniqListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotSpotUniqListResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponseBody) SetCode(v string) *GetHotSpotUniqListResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetData(v *GetHotSpotUniqListResponseBodyData) *GetHotSpotUniqListResponseBody {
	s.Data = v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetMessage(v string) *GetHotSpotUniqListResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetRequestId(v string) *GetHotSpotUniqListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotSpotUniqListResponseBody) SetSuccess(v string) *GetHotSpotUniqListResponseBody {
	s.Success = &v
	return s
}

type GetHotSpotUniqListResponseBodyData struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotSpotUniqListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotSpotUniqListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponseBodyData) SetColumns(v []*string) *GetHotSpotUniqListResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotSpotUniqListResponseBodyData) SetValues(v []*string) *GetHotSpotUniqListResponseBodyData {
	s.Values = v
	return s
}

type GetHotSpotUniqListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotSpotUniqListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotSpotUniqListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotSpotUniqListResponse) GoString() string {
	return s.String()
}

func (s *GetHotSpotUniqListResponse) SetHeaders(v map[string]*string) *GetHotSpotUniqListResponse {
	s.Headers = v
	return s
}

func (s *GetHotSpotUniqListResponse) SetStatusCode(v int32) *GetHotSpotUniqListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotSpotUniqListResponse) SetBody(v *GetHotSpotUniqListResponseBody) *GetHotSpotUniqListResponse {
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

type GetHotspotCompareRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1725415774000
	Beg1End *int64 `json:"beg1_end,omitempty" xml:"beg1_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725415474000
	Beg1Start *int64 `json:"beg1_start,omitempty" xml:"beg1_start,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725415774000
	Beg2End *int64 `json:"beg2_end,omitempty" xml:"beg2_end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1725415474000
	Beg2Start *int64  `json:"beg2_start,omitempty" xml:"beg2_start,omitempty"`
	HotType   *string `json:"hot_type,omitempty" xml:"hot_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-2zei55fwj8nnu31h3z46"
	Instance1 *string `json:"instance1,omitempty" xml:"instance1,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Instance2 *string `json:"instance2,omitempty" xml:"instance2,omitempty"`
	// example:
	//
	// 0
	Pid1 *int64 `json:"pid1,omitempty" xml:"pid1,omitempty"`
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Pid2 *int64 `json:"pid2,omitempty" xml:"pid2,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareRequest) SetBeg1End(v int64) *GetHotspotCompareRequest {
	s.Beg1End = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg1Start(v int64) *GetHotspotCompareRequest {
	s.Beg1Start = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg2End(v int64) *GetHotspotCompareRequest {
	s.Beg2End = &v
	return s
}

func (s *GetHotspotCompareRequest) SetBeg2Start(v int64) *GetHotspotCompareRequest {
	s.Beg2Start = &v
	return s
}

func (s *GetHotspotCompareRequest) SetHotType(v string) *GetHotspotCompareRequest {
	s.HotType = &v
	return s
}

func (s *GetHotspotCompareRequest) SetInstance1(v string) *GetHotspotCompareRequest {
	s.Instance1 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetInstance2(v string) *GetHotspotCompareRequest {
	s.Instance2 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetPid1(v int64) *GetHotspotCompareRequest {
	s.Pid1 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetPid2(v int64) *GetHotspotCompareRequest {
	s.Pid2 = &v
	return s
}

func (s *GetHotspotCompareRequest) SetTable(v string) *GetHotspotCompareRequest {
	s.Table = &v
	return s
}

type GetHotspotCompareResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotCompareResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotspotCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBody) SetCode(v string) *GetHotspotCompareResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetData(v *GetHotspotCompareResponseBodyData) *GetHotspotCompareResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotCompareResponseBody) SetMessage(v string) *GetHotspotCompareResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetRequestId(v string) *GetHotspotCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotCompareResponseBody) SetSuccess(v bool) *GetHotspotCompareResponseBody {
	s.Success = &v
	return s
}

type GetHotspotCompareResponseBodyData struct {
	Flame           *GetHotspotCompareResponseBodyDataFlame           `json:"flame,omitempty" xml:"flame,omitempty" type:"Struct"`
	SeriesInstance1 *GetHotspotCompareResponseBodyDataSeriesInstance1 `json:"series_instance1,omitempty" xml:"series_instance1,omitempty" type:"Struct"`
	SeriesInstance2 *GetHotspotCompareResponseBodyDataSeriesInstance2 `json:"series_instance2,omitempty" xml:"series_instance2,omitempty" type:"Struct"`
}

func (s GetHotspotCompareResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyData) SetFlame(v *GetHotspotCompareResponseBodyDataFlame) *GetHotspotCompareResponseBodyData {
	s.Flame = v
	return s
}

func (s *GetHotspotCompareResponseBodyData) SetSeriesInstance1(v *GetHotspotCompareResponseBodyDataSeriesInstance1) *GetHotspotCompareResponseBodyData {
	s.SeriesInstance1 = v
	return s
}

func (s *GetHotspotCompareResponseBodyData) SetSeriesInstance2(v *GetHotspotCompareResponseBodyDataSeriesInstance2) *GetHotspotCompareResponseBodyData {
	s.SeriesInstance2 = v
	return s
}

type GetHotspotCompareResponseBodyDataFlame struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataFlame) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataFlame) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataFlame) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataFlame {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataFlame) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataFlame {
	s.Values = v
	return s
}

type GetHotspotCompareResponseBodyDataSeriesInstance1 struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance1) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance1) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataSeriesInstance1 {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance1) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataSeriesInstance1 {
	s.Values = v
	return s
}

type GetHotspotCompareResponseBodyDataSeriesInstance2 struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance2) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponseBodyDataSeriesInstance2) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) SetColumns(v []*string) *GetHotspotCompareResponseBodyDataSeriesInstance2 {
	s.Columns = v
	return s
}

func (s *GetHotspotCompareResponseBodyDataSeriesInstance2) SetValues(v [][]*string) *GetHotspotCompareResponseBodyDataSeriesInstance2 {
	s.Values = v
	return s
}

type GetHotspotCompareResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotCompareResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotCompareResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotCompareResponse) SetHeaders(v map[string]*string) *GetHotspotCompareResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotCompareResponse) SetStatusCode(v int32) *GetHotspotCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotCompareResponse) SetBody(v *GetHotspotCompareResponseBody) *GetHotspotCompareResponse {
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

type GetHotspotTrackingRequest struct {
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
	HotType *string `json:"hot_type,omitempty" xml:"hot_type,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// prof_on
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetHotspotTrackingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingRequest) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingRequest) SetBegEnd(v int64) *GetHotspotTrackingRequest {
	s.BegEnd = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetBegStart(v int64) *GetHotspotTrackingRequest {
	s.BegStart = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetHotType(v string) *GetHotspotTrackingRequest {
	s.HotType = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetInstance(v string) *GetHotspotTrackingRequest {
	s.Instance = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetPid(v int64) *GetHotspotTrackingRequest {
	s.Pid = &v
	return s
}

func (s *GetHotspotTrackingRequest) SetTable(v string) *GetHotspotTrackingRequest {
	s.Table = &v
	return s
}

type GetHotspotTrackingResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetHotspotTrackingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetHotspotTrackingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBody) SetCode(v string) *GetHotspotTrackingResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetData(v *GetHotspotTrackingResponseBodyData) *GetHotspotTrackingResponseBody {
	s.Data = v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetMessage(v string) *GetHotspotTrackingResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotTrackingResponseBody) SetRequestId(v string) *GetHotspotTrackingResponseBody {
	s.RequestId = &v
	return s
}

type GetHotspotTrackingResponseBodyData struct {
	Flame  *GetHotspotTrackingResponseBodyDataFlame  `json:"flame,omitempty" xml:"flame,omitempty" type:"Struct"`
	Series *GetHotspotTrackingResponseBodyDataSeries `json:"series,omitempty" xml:"series,omitempty" type:"Struct"`
}

func (s GetHotspotTrackingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyData) SetFlame(v *GetHotspotTrackingResponseBodyDataFlame) *GetHotspotTrackingResponseBodyData {
	s.Flame = v
	return s
}

func (s *GetHotspotTrackingResponseBodyData) SetSeries(v *GetHotspotTrackingResponseBodyDataSeries) *GetHotspotTrackingResponseBodyData {
	s.Series = v
	return s
}

type GetHotspotTrackingResponseBodyDataFlame struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotTrackingResponseBodyDataFlame) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyDataFlame) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyDataFlame) SetColumns(v []*string) *GetHotspotTrackingResponseBodyDataFlame {
	s.Columns = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataFlame) SetValues(v [][]*string) *GetHotspotTrackingResponseBodyDataFlame {
	s.Values = v
	return s
}

type GetHotspotTrackingResponseBodyDataSeries struct {
	Columns []*string   `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	Values  [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetHotspotTrackingResponseBodyDataSeries) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingResponseBodyDataSeries) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponseBodyDataSeries) SetColumns(v []*string) *GetHotspotTrackingResponseBodyDataSeries {
	s.Columns = v
	return s
}

func (s *GetHotspotTrackingResponseBodyDataSeries) SetValues(v [][]*string) *GetHotspotTrackingResponseBodyDataSeries {
	s.Values = v
	return s
}

type GetHotspotTrackingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotspotTrackingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotspotTrackingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotspotTrackingResponse) GoString() string {
	return s.String()
}

func (s *GetHotspotTrackingResponse) SetHeaders(v map[string]*string) *GetHotspotTrackingResponse {
	s.Headers = v
	return s
}

func (s *GetHotspotTrackingResponse) SetStatusCode(v int32) *GetHotspotTrackingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotspotTrackingResponse) SetBody(v *GetHotspotTrackingResponseBody) *GetHotspotTrackingResponse {
	s.Body = v
	return s
}

type GetInstantScoreRequest struct {
	// example:
	//
	// 1808078950770264
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
}

func (s GetInstantScoreRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstantScoreRequest) GoString() string {
	return s.String()
}

func (s *GetInstantScoreRequest) SetCluster(v string) *GetInstantScoreRequest {
	s.Cluster = &v
	return s
}

func (s *GetInstantScoreRequest) SetInstance(v string) *GetInstantScoreRequest {
	s.Instance = &v
	return s
}

type GetInstantScoreResponseBody struct {
	// 集群ID
	//
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetInstantScoreResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Query no data
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInstantScoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstantScoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponseBody) SetCode(v string) *GetInstantScoreResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstantScoreResponseBody) SetData(v *GetInstantScoreResponseBodyData) *GetInstantScoreResponseBody {
	s.Data = v
	return s
}

func (s *GetInstantScoreResponseBody) SetMessage(v string) *GetInstantScoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstantScoreResponseBody) SetRequestId(v string) *GetInstantScoreResponseBody {
	s.RequestId = &v
	return s
}

type GetInstantScoreResponseBodyData struct {
	// example:
	//
	// 100
	Error *float32 `json:"error,omitempty" xml:"error,omitempty"`
	// example:
	//
	// 100
	Latency *float32 `json:"latency,omitempty" xml:"latency,omitempty"`
	// example:
	//
	// 100
	Load *float32 `json:"load,omitempty" xml:"load,omitempty"`
	// example:
	//
	// 100
	Saturation *float32 `json:"saturation,omitempty" xml:"saturation,omitempty"`
	Total      *float32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetInstantScoreResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstantScoreResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponseBodyData) SetError(v float32) *GetInstantScoreResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetLatency(v float32) *GetInstantScoreResponseBodyData {
	s.Latency = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetLoad(v float32) *GetInstantScoreResponseBodyData {
	s.Load = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetSaturation(v float32) *GetInstantScoreResponseBodyData {
	s.Saturation = &v
	return s
}

func (s *GetInstantScoreResponseBodyData) SetTotal(v float32) *GetInstantScoreResponseBodyData {
	s.Total = &v
	return s
}

type GetInstantScoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstantScoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstantScoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstantScoreResponse) GoString() string {
	return s.String()
}

func (s *GetInstantScoreResponse) SetHeaders(v map[string]*string) *GetInstantScoreResponse {
	s.Headers = v
	return s
}

func (s *GetInstantScoreResponse) SetStatusCode(v int32) *GetInstantScoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstantScoreResponse) SetBody(v *GetInstantScoreResponseBody) *GetInstantScoreResponse {
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

type GetServiceFuncStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	Params *GetServiceFuncStatusRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s GetServiceFuncStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusRequest) SetChannel(v string) *GetServiceFuncStatusRequest {
	s.Channel = &v
	return s
}

func (s *GetServiceFuncStatusRequest) SetParams(v *GetServiceFuncStatusRequestParams) *GetServiceFuncStatusRequest {
	s.Params = v
	return s
}

func (s *GetServiceFuncStatusRequest) SetServiceName(v string) *GetServiceFuncStatusRequest {
	s.ServiceName = &v
	return s
}

type GetServiceFuncStatusRequestParams struct {
	// This parameter is required.
	//
	// example:
	//
	// mullprof
	FunctionName *string `json:"function_name,omitempty" xml:"function_name,omitempty"`
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 1338904783509062
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GetServiceFuncStatusRequestParams) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusRequestParams) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusRequestParams) SetFunctionName(v string) *GetServiceFuncStatusRequestParams {
	s.FunctionName = &v
	return s
}

func (s *GetServiceFuncStatusRequestParams) SetInstance(v string) *GetServiceFuncStatusRequestParams {
	s.Instance = &v
	return s
}

func (s *GetServiceFuncStatusRequestParams) SetUid(v string) *GetServiceFuncStatusRequestParams {
	s.Uid = &v
	return s
}

type GetServiceFuncStatusShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s GetServiceFuncStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusShrinkRequest) SetChannel(v string) *GetServiceFuncStatusShrinkRequest {
	s.Channel = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) SetParamsShrink(v string) *GetServiceFuncStatusShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) SetServiceName(v string) *GetServiceFuncStatusShrinkRequest {
	s.ServiceName = &v
	return s
}

type GetServiceFuncStatusResponseBody struct {
	// example:
	//
	// Success
	Code    *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data    *GetServiceFuncStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message *string                               `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceFuncStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBody) SetCode(v string) *GetServiceFuncStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetData(v *GetServiceFuncStatusResponseBodyData) *GetServiceFuncStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetMessage(v string) *GetServiceFuncStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceFuncStatusResponseBody) SetRequestId(v string) *GetServiceFuncStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceFuncStatusResponseBodyData struct {
	Args *GetServiceFuncStatusResponseBodyDataArgs `json:"args,omitempty" xml:"args,omitempty" type:"Struct"`
}

func (s GetServiceFuncStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBodyData) SetArgs(v *GetServiceFuncStatusResponseBodyDataArgs) *GetServiceFuncStatusResponseBodyData {
	s.Args = v
	return s
}

type GetServiceFuncStatusResponseBodyDataArgs struct {
	// example:
	//
	// java
	AddCmd *string `json:"add_cmd,omitempty" xml:"add_cmd,omitempty"`
	// example:
	//
	// true
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// /tmp/sysom/java-profiler
	JavaStorePath *string `json:"java_store_path,omitempty" xml:"java_store_path,omitempty"`
	// example:
	//
	// true
	Locks *string `json:"locks,omitempty" xml:"locks,omitempty"`
	// example:
	//
	// -1
	Loop *int32 `json:"loop,omitempty" xml:"loop,omitempty"`
	// example:
	//
	// true
	Mem *string `json:"mem,omitempty" xml:"mem,omitempty"`
	// example:
	//
	// true
	SystemProfiling *string `json:"system_profiling,omitempty" xml:"system_profiling,omitempty"`
}

func (s GetServiceFuncStatusResponseBodyDataArgs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusResponseBodyDataArgs) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetAddCmd(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.AddCmd = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetCpu(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Cpu = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetJavaStorePath(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.JavaStorePath = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetLocks(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Locks = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetLoop(v int32) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Loop = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetMem(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.Mem = &v
	return s
}

func (s *GetServiceFuncStatusResponseBodyDataArgs) SetSystemProfiling(v string) *GetServiceFuncStatusResponseBodyDataArgs {
	s.SystemProfiling = &v
	return s
}

type GetServiceFuncStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceFuncStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceFuncStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceFuncStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusResponse) SetHeaders(v map[string]*string) *GetServiceFuncStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceFuncStatusResponse) SetStatusCode(v int32) *GetServiceFuncStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceFuncStatusResponse) SetBody(v *GetServiceFuncStatusResponseBody) *GetServiceFuncStatusResponse {
	s.Body = v
	return s
}

type InitialSysomRequest struct {
	CheckOnly *bool `json:"check_only,omitempty" xml:"check_only,omitempty"`
}

func (s InitialSysomRequest) String() string {
	return tea.Prettify(s)
}

func (s InitialSysomRequest) GoString() string {
	return s.String()
}

func (s *InitialSysomRequest) SetCheckOnly(v bool) *InitialSysomRequest {
	s.CheckOnly = &v
	return s
}

type InitialSysomResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data      *InitialSysomResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                       `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InitialSysomResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitialSysomResponseBody) GoString() string {
	return s.String()
}

func (s *InitialSysomResponseBody) SetRequestId(v string) *InitialSysomResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitialSysomResponseBody) SetCode(v string) *InitialSysomResponseBody {
	s.Code = &v
	return s
}

func (s *InitialSysomResponseBody) SetData(v *InitialSysomResponseBodyData) *InitialSysomResponseBody {
	s.Data = v
	return s
}

func (s *InitialSysomResponseBody) SetMessage(v string) *InitialSysomResponseBody {
	s.Message = &v
	return s
}

type InitialSysomResponseBodyData struct {
	RoleExist *bool `json:"role_exist,omitempty" xml:"role_exist,omitempty"`
}

func (s InitialSysomResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitialSysomResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitialSysomResponseBodyData) SetRoleExist(v bool) *InitialSysomResponseBodyData {
	s.RoleExist = &v
	return s
}

type InitialSysomResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitialSysomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitialSysomResponse) String() string {
	return tea.Prettify(s)
}

func (s InitialSysomResponse) GoString() string {
	return s.String()
}

func (s *InitialSysomResponse) SetHeaders(v map[string]*string) *InitialSysomResponse {
	s.Headers = v
	return s
}

func (s *InitialSysomResponse) SetStatusCode(v int32) *InitialSysomResponse {
	s.StatusCode = &v
	return s
}

func (s *InitialSysomResponse) SetBody(v *InitialSysomResponseBody) *InitialSysomResponse {
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

type InstallAgentForClusterRequest struct {
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// example:
	//
	// c9d7f3fc3d42942afbcb65c1100ffb19d
	ClusterId       *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	GrayscaleConfig *string `json:"grayscale_config,omitempty" xml:"grayscale_config,omitempty"`
}

func (s InstallAgentForClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterRequest) SetAgentId(v string) *InstallAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetAgentVersion(v string) *InstallAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetClusterId(v string) *InstallAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallAgentForClusterRequest) SetGrayscaleConfig(v string) *InstallAgentForClusterRequest {
	s.GrayscaleConfig = &v
	return s
}

type InstallAgentForClusterResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *InstallAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s InstallAgentForClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponseBody) SetRequestId(v string) *InstallAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetCode(v string) *InstallAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetData(v *InstallAgentForClusterResponseBodyData) *InstallAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *InstallAgentForClusterResponseBody) SetMessage(v string) *InstallAgentForClusterResponseBody {
	s.Message = &v
	return s
}

type InstallAgentForClusterResponseBodyData struct {
	// example:
	//
	// 049ea0609515414b9e19c3389d7ba638
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s InstallAgentForClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponseBodyData) SetTaskId(v string) *InstallAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

type InstallAgentForClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallAgentForClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *InstallAgentForClusterResponse) SetHeaders(v map[string]*string) *InstallAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *InstallAgentForClusterResponse) SetStatusCode(v int32) *InstallAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallAgentForClusterResponse) SetBody(v *InstallAgentForClusterResponseBody) *InstallAgentForClusterResponse {
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

type ListClusterAgentInstallRecordsRequest struct {
	// example:
	//
	// cbd80af02b9d6454ebdc579c5e022d0c8
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
}

func (s ListClusterAgentInstallRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterAgentInstallRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsRequest) SetClusterId(v string) *ListClusterAgentInstallRecordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetCurrent(v int64) *ListClusterAgentInstallRecordsRequest {
	s.Current = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPageSize(v int64) *ListClusterAgentInstallRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPluginId(v string) *ListClusterAgentInstallRecordsRequest {
	s.PluginId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsRequest) SetPluginVersion(v string) *ListClusterAgentInstallRecordsRequest {
	s.PluginVersion = &v
	return s
}

type ListClusterAgentInstallRecordsResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                           `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListClusterAgentInstallRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetRequestId(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetCode(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetData(v []*ListClusterAgentInstallRecordsResponseBodyData) *ListClusterAgentInstallRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetMessage(v string) *ListClusterAgentInstallRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBody) SetTotal(v int64) *ListClusterAgentInstallRecordsResponseBody {
	s.Total = &v
	return s
}

type ListClusterAgentInstallRecordsResponseBodyData struct {
	// example:
	//
	// cbf7a37bc905d4682a3338b3744810269
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	GrayscaleConfig *string `json:"grayscale_config,omitempty" xml:"grayscale_config,omitempty"`
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	PluginId *string `json:"plugin_id,omitempty" xml:"plugin_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	PluginVersion *string `json:"plugin_version,omitempty" xml:"plugin_version,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetClusterId(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetCreatedAt(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetGrayscaleConfig(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.GrayscaleConfig = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetPluginId(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.PluginId = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetPluginVersion(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.PluginVersion = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponseBodyData) SetUpdatedAt(v string) *ListClusterAgentInstallRecordsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type ListClusterAgentInstallRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterAgentInstallRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterAgentInstallRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterAgentInstallRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterAgentInstallRecordsResponse) SetHeaders(v map[string]*string) *ListClusterAgentInstallRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterAgentInstallRecordsResponse) SetStatusCode(v int32) *ListClusterAgentInstallRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterAgentInstallRecordsResponse) SetBody(v *ListClusterAgentInstallRecordsResponseBody) *ListClusterAgentInstallRecordsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// example:
	//
	// cb7d4cc26c8f845fb8a8255ffd394820e
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// Running
	ClusterStatus *string `json:"cluster_status,omitempty" xml:"cluster_status,omitempty"`
	// example:
	//
	// ACK
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// cb7d4cc26c8f845fb8a8255ffd394820e
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// proxy-next-upstream
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetClusterId(v string) *ListClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClustersRequest) SetClusterStatus(v string) *ListClustersRequest {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersRequest) SetClusterType(v string) *ListClustersRequest {
	s.ClusterType = &v
	return s
}

func (s *ListClustersRequest) SetCurrent(v int64) *ListClustersRequest {
	s.Current = &v
	return s
}

func (s *ListClustersRequest) SetId(v string) *ListClustersRequest {
	s.Id = &v
	return s
}

func (s *ListClustersRequest) SetName(v string) *ListClustersRequest {
	s.Name = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int64) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                         `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListClustersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 64
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetCode(v string) *ListClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClustersResponseBody) SetData(v []*ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetTotal(v int64) *ListClustersResponseBody {
	s.Total = &v
	return s
}

type ListClustersResponseBodyData struct {
	// example:
	//
	// c666d4774f0e2440b979bf917bf100e40
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// Running
	ClusterStatus *string `json:"cluster_status,omitempty" xml:"cluster_status,omitempty"`
	// example:
	//
	// ACK
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// 5389fba5-92a1-4ff4-9b26-773b97828144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// auto-name-sbvCT
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 2024-12-25T15:08:19
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) SetClusterId(v string) *ListClustersResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterStatus(v string) *ListClustersResponseBodyData {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersResponseBodyData) SetClusterType(v string) *ListClustersResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyData) SetCreatedAt(v string) *ListClustersResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListClustersResponseBodyData) SetId(v string) *ListClustersResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyData) SetName(v string) *ListClustersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyData) SetRegion(v string) *ListClustersResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListClustersResponseBodyData) SetUpdatedAt(v string) *ListClustersResponseBodyData {
	s.UpdatedAt = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListDiagnosisRequest struct {
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Params      *string `json:"params,omitempty" xml:"params,omitempty"`
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosisRequest) SetCurrent(v int64) *ListDiagnosisRequest {
	s.Current = &v
	return s
}

func (s *ListDiagnosisRequest) SetPageSize(v int64) *ListDiagnosisRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiagnosisRequest) SetParams(v string) *ListDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *ListDiagnosisRequest) SetServiceName(v string) *ListDiagnosisRequest {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisRequest) SetStatus(v string) *ListDiagnosisRequest {
	s.Status = &v
	return s
}

type ListDiagnosisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListDiagnosisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Total   *int64  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBody) SetRequestId(v string) *ListDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetCode(v string) *ListDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetData(v []*ListDiagnosisResponseBodyData) *ListDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *ListDiagnosisResponseBody) SetMessage(v string) *ListDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiagnosisResponseBody) SetTotal(v int64) *ListDiagnosisResponseBody {
	s.Total = &v
	return s
}

type ListDiagnosisResponseBodyData struct {
	Code        *int32      `json:"code,omitempty" xml:"code,omitempty"`
	Command     interface{} `json:"command,omitempty" xml:"command,omitempty"`
	CreatedAt   *string     `json:"created_at,omitempty" xml:"created_at,omitempty"`
	ErrMsg      *string     `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	Params      interface{} `json:"params,omitempty" xml:"params,omitempty"`
	Result      interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ServiceName *string     `json:"service_name,omitempty" xml:"service_name,omitempty"`
	Status      *string     `json:"status,omitempty" xml:"status,omitempty"`
	TaskId      *string     `json:"task_id,omitempty" xml:"task_id,omitempty"`
	UpdatedAt   *string     `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Url         *string     `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListDiagnosisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponseBodyData) SetCode(v int32) *ListDiagnosisResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCommand(v interface{}) *ListDiagnosisResponseBodyData {
	s.Command = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetCreatedAt(v string) *ListDiagnosisResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetErrMsg(v string) *ListDiagnosisResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetParams(v interface{}) *ListDiagnosisResponseBodyData {
	s.Params = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetResult(v interface{}) *ListDiagnosisResponseBodyData {
	s.Result = v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetServiceName(v string) *ListDiagnosisResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetStatus(v string) *ListDiagnosisResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetTaskId(v string) *ListDiagnosisResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUpdatedAt(v string) *ListDiagnosisResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListDiagnosisResponseBodyData) SetUrl(v string) *ListDiagnosisResponseBodyData {
	s.Url = &v
	return s
}

type ListDiagnosisResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosisResponse) SetHeaders(v map[string]*string) *ListDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosisResponse) SetStatusCode(v int32) *ListDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosisResponse) SetBody(v *ListDiagnosisResponseBody) *ListDiagnosisResponse {
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

type ListInstanceStatusRequest struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// i-wz9b9vucz1iubsz355rh
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusRequest) SetCurrent(v int64) *ListInstanceStatusRequest {
	s.Current = &v
	return s
}

func (s *ListInstanceStatusRequest) SetInstance(v string) *ListInstanceStatusRequest {
	s.Instance = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageSize(v int64) *ListInstanceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusRequest) SetRegion(v string) *ListInstanceStatusRequest {
	s.Region = &v
	return s
}

func (s *ListInstanceStatusRequest) SetStatus(v string) *ListInstanceStatusRequest {
	s.Status = &v
	return s
}

type ListInstanceStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstanceStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 218
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBody) SetRequestId(v string) *ListInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetCode(v string) *ListInstanceStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetData(v []*ListInstanceStatusResponseBodyData) *ListInstanceStatusResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceStatusResponseBody) SetMessage(v string) *ListInstanceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetTotal(v int64) *ListInstanceStatusResponseBody {
	s.Total = &v
	return s
}

type ListInstanceStatusResponseBodyData struct {
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstanceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBodyData) SetInstance(v string) *ListInstanceStatusResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstanceStatusResponseBodyData) SetRegion(v string) *ListInstanceStatusResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListInstanceStatusResponseBodyData) SetStatus(v string) *ListInstanceStatusResponseBodyData {
	s.Status = &v
	return s
}

type ListInstanceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponse) SetHeaders(v map[string]*string) *ListInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStatusResponse) SetStatusCode(v int32) *ListInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceStatusResponse) SetBody(v *ListInstanceStatusResponseBody) *ListInstanceStatusResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// xxxxx
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 1
	Current  *int64  `json:"current,omitempty" xml:"current,omitempty"`
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetClusterId(v string) *ListInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesRequest) SetCurrent(v int64) *ListInstancesRequest {
	s.Current = &v
	return s
}

func (s *ListInstancesRequest) SetInstance(v string) *ListInstancesRequest {
	s.Instance = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetRegion(v string) *ListInstancesRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

type ListInstancesResponseBody struct {
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// Requests for llm service failed
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 9515E5A0-8905-59B0-9BBF-5F0BE568C3A0
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 623
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v []*ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotal(v int64) *ListInstancesResponseBody {
	s.Total = &v
	return s
}

type ListInstancesResponseBodyData struct {
	ClusterId     *string     `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	Instance      *string     `json:"instance,omitempty" xml:"instance,omitempty"`
	KernelVersion *string     `json:"kernel_version,omitempty" xml:"kernel_version,omitempty"`
	Meta          interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	OsArch        *string     `json:"os_arch,omitempty" xml:"os_arch,omitempty"`
	OsHealthScore *string     `json:"os_health_score,omitempty" xml:"os_health_score,omitempty"`
	OsName        *string     `json:"os_name,omitempty" xml:"os_name,omitempty"`
	OsNameId      *string     `json:"os_name_id,omitempty" xml:"os_name_id,omitempty"`
	OsVersion     *string     `json:"os_version,omitempty" xml:"os_version,omitempty"`
	OsVersionId   *string     `json:"os_version_id,omitempty" xml:"os_version_id,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetClusterId(v string) *ListInstancesResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetInstance(v string) *ListInstancesResponseBodyData {
	s.Instance = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetKernelVersion(v string) *ListInstancesResponseBodyData {
	s.KernelVersion = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetMeta(v interface{}) *ListInstancesResponseBodyData {
	s.Meta = v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsArch(v string) *ListInstancesResponseBodyData {
	s.OsArch = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsHealthScore(v string) *ListInstancesResponseBodyData {
	s.OsHealthScore = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsName(v string) *ListInstancesResponseBodyData {
	s.OsName = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsNameId(v string) *ListInstancesResponseBodyData {
	s.OsNameId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsVersion(v string) *ListInstancesResponseBodyData {
	s.OsVersion = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetOsVersionId(v string) *ListInstancesResponseBodyData {
	s.OsVersionId = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetRegion(v string) *ListInstancesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetStatus(v string) *ListInstancesResponseBodyData {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListPodsOfInstanceRequest struct {
	// example:
	//
	// c96e34d74eb6748f3b2a46552d5d653f6
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// i-wz9d00ut2ska3mlyhn6j
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPodsOfInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPodsOfInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceRequest) SetClusterId(v string) *ListPodsOfInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetCurrent(v int64) *ListPodsOfInstanceRequest {
	s.Current = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetInstance(v string) *ListPodsOfInstanceRequest {
	s.Instance = &v
	return s
}

func (s *ListPodsOfInstanceRequest) SetPageSize(v int64) *ListPodsOfInstanceRequest {
	s.PageSize = &v
	return s
}

type ListPodsOfInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListPodsOfInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// instance not exists
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 42
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPodsOfInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPodsOfInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponseBody) SetRequestId(v string) *ListPodsOfInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetCode(v string) *ListPodsOfInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetData(v []*ListPodsOfInstanceResponseBodyData) *ListPodsOfInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetMessage(v string) *ListPodsOfInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListPodsOfInstanceResponseBody) SetTotal(v int64) *ListPodsOfInstanceResponseBody {
	s.Total = &v
	return s
}

type ListPodsOfInstanceResponseBodyData struct {
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// test-pod
	Pod *string `json:"pod,omitempty" xml:"pod,omitempty"`
}

func (s ListPodsOfInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPodsOfInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponseBodyData) SetNamespace(v string) *ListPodsOfInstanceResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *ListPodsOfInstanceResponseBodyData) SetPod(v string) *ListPodsOfInstanceResponseBodyData {
	s.Pod = &v
	return s
}

type ListPodsOfInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPodsOfInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPodsOfInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPodsOfInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListPodsOfInstanceResponse) SetHeaders(v map[string]*string) *ListPodsOfInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListPodsOfInstanceResponse) SetStatusCode(v int32) *ListPodsOfInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPodsOfInstanceResponse) SetBody(v *ListPodsOfInstanceResponseBody) *ListPodsOfInstanceResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ["cn-hangzhou", "cn-shengzhen"]
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetData(v []*string) *ListRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
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

type UninstallAgentForClusterRequest struct {
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// example:
	//
	// c822f83bb45994ddbac9326b4c2f04f35
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
}

func (s UninstallAgentForClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterRequest) SetAgentId(v string) *UninstallAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetAgentVersion(v string) *UninstallAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UninstallAgentForClusterRequest) SetClusterId(v string) *UninstallAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

type UninstallAgentForClusterResponseBody struct {
	// example:
	//
	// 44841312-7227-55C9-AE03-D59729BFAE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *UninstallAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.NotAuthorizedInstance Instance 21 is not authorized
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UninstallAgentForClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponseBody) SetRequestId(v string) *UninstallAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetCode(v string) *UninstallAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetData(v *UninstallAgentForClusterResponseBodyData) *UninstallAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *UninstallAgentForClusterResponseBody) SetMessage(v string) *UninstallAgentForClusterResponseBody {
	s.Message = &v
	return s
}

type UninstallAgentForClusterResponseBodyData struct {
	// example:
	//
	// 049ea0609515414b9e19c3389d7ba638
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UninstallAgentForClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponseBodyData) SetTaskId(v string) *UninstallAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

type UninstallAgentForClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UninstallAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UninstallAgentForClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *UninstallAgentForClusterResponse) SetHeaders(v map[string]*string) *UninstallAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *UninstallAgentForClusterResponse) SetStatusCode(v int32) *UninstallAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallAgentForClusterResponse) SetBody(v *UninstallAgentForClusterResponseBody) *UninstallAgentForClusterResponse {
	s.Body = v
	return s
}

type UpdateEventsAttentionRequest struct {
	Body *UpdateEventsAttentionRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UpdateEventsAttentionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionRequest) SetBody(v *UpdateEventsAttentionRequestBody) *UpdateEventsAttentionRequest {
	s.Body = v
	return s
}

type UpdateEventsAttentionRequestBody struct {
	// example:
	//
	// 0
	Mode *int32 `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// cluster
	Range *string `json:"range,omitempty" xml:"range,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 03de78af-f49f-433d-b5b1-0f6a70c493ba
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s UpdateEventsAttentionRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionRequestBody) SetMode(v int32) *UpdateEventsAttentionRequestBody {
	s.Mode = &v
	return s
}

func (s *UpdateEventsAttentionRequestBody) SetRange(v string) *UpdateEventsAttentionRequestBody {
	s.Range = &v
	return s
}

func (s *UpdateEventsAttentionRequestBody) SetUuid(v string) *UpdateEventsAttentionRequestBody {
	s.Uuid = &v
	return s
}

type UpdateEventsAttentionShrinkRequest struct {
	BodyShrink *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventsAttentionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionShrinkRequest) SetBodyShrink(v string) *UpdateEventsAttentionShrinkRequest {
	s.BodyShrink = &v
	return s
}

type UpdateEventsAttentionResponseBody struct {
	// example:
	//
	// 44841312-7227-55C9-AE03-D59729BFAE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateEventsAttentionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// Instance not belong to current user
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateEventsAttentionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponseBody) SetRequestId(v string) *UpdateEventsAttentionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetCode(v string) *UpdateEventsAttentionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetData(v *UpdateEventsAttentionResponseBodyData) *UpdateEventsAttentionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEventsAttentionResponseBody) SetMessage(v string) *UpdateEventsAttentionResponseBody {
	s.Message = &v
	return s
}

type UpdateEventsAttentionResponseBodyData struct {
	// example:
	//
	// 1
	Mode *int32 `json:"mode,omitempty" xml:"mode,omitempty"`
}

func (s UpdateEventsAttentionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponseBodyData) SetMode(v int32) *UpdateEventsAttentionResponseBodyData {
	s.Mode = &v
	return s
}

type UpdateEventsAttentionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventsAttentionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventsAttentionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventsAttentionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventsAttentionResponse) SetHeaders(v map[string]*string) *UpdateEventsAttentionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventsAttentionResponse) SetStatusCode(v int32) *UpdateEventsAttentionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventsAttentionResponse) SetBody(v *UpdateEventsAttentionResponseBody) *UpdateEventsAttentionResponse {
	s.Body = v
	return s
}

type UpdateFuncSwitchRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	Params *UpdateFuncSwitchRecordRequestParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s UpdateFuncSwitchRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequest) SetChannel(v string) *UpdateFuncSwitchRecordRequest {
	s.Channel = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequest) SetParams(v *UpdateFuncSwitchRecordRequestParams) *UpdateFuncSwitchRecordRequest {
	s.Params = v
	return s
}

func (s *UpdateFuncSwitchRecordRequest) SetServiceName(v string) *UpdateFuncSwitchRecordRequest {
	s.ServiceName = &v
	return s
}

type UpdateFuncSwitchRecordRequestParams struct {
	Args *UpdateFuncSwitchRecordRequestParamsArgs `json:"args,omitempty" xml:"args,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// mullprof
	FunctionName *string `json:"function_name,omitempty" xml:"function_name,omitempty"`
	// example:
	//
	// i-2zei55fwj8nnu31h3z46
	Instance *string `json:"instance,omitempty" xml:"instance,omitempty"`
	// example:
	//
	// restart
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// example:
	//
	// 1664516888213680
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s UpdateFuncSwitchRecordRequestParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequestParams) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequestParams) SetArgs(v *UpdateFuncSwitchRecordRequestParamsArgs) *UpdateFuncSwitchRecordRequestParams {
	s.Args = v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetFunctionName(v string) *UpdateFuncSwitchRecordRequestParams {
	s.FunctionName = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetInstance(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Instance = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetOp(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Op = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParams) SetUid(v string) *UpdateFuncSwitchRecordRequestParams {
	s.Uid = &v
	return s
}

type UpdateFuncSwitchRecordRequestParamsArgs struct {
	// example:
	//
	// java
	AddCmd *string `json:"add_cmd,omitempty" xml:"add_cmd,omitempty"`
	// example:
	//
	// true
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// example:
	//
	// /tmp/sysom/java-profiler
	JavaStorePath *string `json:"java_store_path,omitempty" xml:"java_store_path,omitempty"`
	// example:
	//
	// true
	Locks *string `json:"locks,omitempty" xml:"locks,omitempty"`
	// example:
	//
	// -1
	Loop *int32 `json:"loop,omitempty" xml:"loop,omitempty"`
	// example:
	//
	// true
	Mem *string `json:"mem,omitempty" xml:"mem,omitempty"`
	// example:
	//
	// true
	SystemProfiling *string `json:"system_profiling,omitempty" xml:"system_profiling,omitempty"`
}

func (s UpdateFuncSwitchRecordRequestParamsArgs) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordRequestParamsArgs) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetAddCmd(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.AddCmd = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetCpu(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Cpu = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetJavaStorePath(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.JavaStorePath = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetLocks(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Locks = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetLoop(v int32) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Loop = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetMem(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.Mem = &v
	return s
}

func (s *UpdateFuncSwitchRecordRequestParamsArgs) SetSystemProfiling(v string) *UpdateFuncSwitchRecordRequestParamsArgs {
	s.SystemProfiling = &v
	return s
}

type UpdateFuncSwitchRecordShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s UpdateFuncSwitchRecordShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetChannel(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.Channel = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetParamsShrink(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetServiceName(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.ServiceName = &v
	return s
}

type UpdateFuncSwitchRecordResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateFuncSwitchRecordResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// result: code=1 msg=(Request failed, status_code != 200)
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateFuncSwitchRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponseBody) SetCode(v string) *UpdateFuncSwitchRecordResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetData(v *UpdateFuncSwitchRecordResponseBodyData) *UpdateFuncSwitchRecordResponseBody {
	s.Data = v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetMessage(v string) *UpdateFuncSwitchRecordResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponseBody) SetRequestId(v string) *UpdateFuncSwitchRecordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateFuncSwitchRecordResponseBodyData struct {
	// example:
	//
	// 63fc5acb99e642d793f42912612e8001
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpdateFuncSwitchRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponseBodyData) SetTaskId(v string) *UpdateFuncSwitchRecordResponseBodyData {
	s.TaskId = &v
	return s
}

type UpdateFuncSwitchRecordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFuncSwitchRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFuncSwitchRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFuncSwitchRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordResponse) SetHeaders(v map[string]*string) *UpdateFuncSwitchRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateFuncSwitchRecordResponse) SetStatusCode(v int32) *UpdateFuncSwitchRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFuncSwitchRecordResponse) SetBody(v *UpdateFuncSwitchRecordResponseBody) *UpdateFuncSwitchRecordResponse {
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

type UpgradeAgentForClusterRequest struct {
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	AgentId *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// example:
	//
	// 3.4.0-1
	AgentVersion *string `json:"agent_version,omitempty" xml:"agent_version,omitempty"`
	// example:
	//
	// c1c187fd513cb41a19876bac0e6b05212
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
}

func (s UpgradeAgentForClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentForClusterRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterRequest) SetAgentId(v string) *UpgradeAgentForClusterRequest {
	s.AgentId = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetAgentVersion(v string) *UpgradeAgentForClusterRequest {
	s.AgentVersion = &v
	return s
}

func (s *UpgradeAgentForClusterRequest) SetClusterId(v string) *UpgradeAgentForClusterRequest {
	s.ClusterId = &v
	return s
}

type UpgradeAgentForClusterResponseBody struct {
	// example:
	//
	// B149FD9C-ED5C-5765-B3AD-05AA4A4D64D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpgradeAgentForClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpgradeAgentForClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentForClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponseBody) SetRequestId(v string) *UpgradeAgentForClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetCode(v string) *UpgradeAgentForClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetData(v *UpgradeAgentForClusterResponseBodyData) *UpgradeAgentForClusterResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeAgentForClusterResponseBody) SetMessage(v string) *UpgradeAgentForClusterResponseBody {
	s.Message = &v
	return s
}

type UpgradeAgentForClusterResponseBodyData struct {
	// example:
	//
	// 7523e9e0ddc74d99a5236f4f4d5056e6
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s UpgradeAgentForClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentForClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponseBodyData) SetTaskId(v string) *UpgradeAgentForClusterResponseBodyData {
	s.TaskId = &v
	return s
}

type UpgradeAgentForClusterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeAgentForClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeAgentForClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeAgentForClusterResponse) GoString() string {
	return s.String()
}

func (s *UpgradeAgentForClusterResponse) SetHeaders(v map[string]*string) *UpgradeAgentForClusterResponse {
	s.Headers = v
	return s
}

func (s *UpgradeAgentForClusterResponse) SetStatusCode(v int32) *UpgradeAgentForClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeAgentForClusterResponse) SetBody(v *UpgradeAgentForClusterResponseBody) *UpgradeAgentForClusterResponse {
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &AuthDiagnosisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &AuthDiagnosisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateCopilotResponseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateCopilotResponseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 流式copilot服务接口
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponseWithOptions(request *GenerateCopilotStreamResponseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateCopilotStreamResponseResponse, _err error) {
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
		Action:      tea.String("GenerateCopilotStreamResponse"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/copilot/generate_copilot_stream_response"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GenerateCopilotStreamResponseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GenerateCopilotStreamResponseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 流式copilot服务接口
//
// @param request - GenerateCopilotStreamResponseRequest
//
// @return GenerateCopilotStreamResponseResponse
func (client *Client) GenerateCopilotStreamResponse(request *GenerateCopilotStreamResponseRequest) (_result *GenerateCopilotStreamResponseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateCopilotStreamResponseResponse{}
	_body, _err := client.GenerateCopilotStreamResponseWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAIQueryResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAIQueryResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAbnormalEventsCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAbnormalEventsCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAgentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAgentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAgentTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAgentTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取copilot历史聊天记录
//
// @param request - GetCopilotHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCopilotHistoryResponse
func (client *Client) GetCopilotHistoryWithOptions(request *GetCopilotHistoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCopilotHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Count)) {
		body["count"] = request.Count
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCopilotHistory"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/copilot/get_copilot_history"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetCopilotHistoryResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetCopilotHistoryResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取copilot历史聊天记录
//
// @param request - GetCopilotHistoryRequest
//
// @return GetCopilotHistoryResponse
func (client *Client) GetCopilotHistory(request *GetCopilotHistoryRequest) (_result *GetCopilotHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCopilotHistoryResponse{}
	_body, _err := client.GetCopilotHistoryWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDiagnosisResultResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDiagnosisResultResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHealthPercentageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHealthPercentageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHostCountResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHostCountResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取实例下的某个字段列表
//
// @param request - GetHotSpotUniqListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotSpotUniqListResponse
func (client *Client) GetHotSpotUniqListWithOptions(request *GetHotSpotUniqListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotSpotUniqListResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Pid)) {
		body["pid"] = request.Pid
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["table"] = request.Table
	}

	if !tea.BoolValue(util.IsUnset(request.Uniq)) {
		body["uniq"] = request.Uniq
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotSpotUniqList"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/livetrace_proxy/hotspot_uniq_list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotSpotUniqListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotSpotUniqListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例下的某个字段列表
//
// @param request - GetHotSpotUniqListRequest
//
// @return GetHotSpotUniqListResponse
func (client *Client) GetHotSpotUniqList(request *GetHotSpotUniqListRequest) (_result *GetHotSpotUniqListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotSpotUniqListResponse{}
	_body, _err := client.GetHotSpotUniqListWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotspotAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotspotAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 热点对比
//
// @param request - GetHotspotCompareRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotCompareResponse
func (client *Client) GetHotspotCompareWithOptions(request *GetHotspotCompareRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotspotCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Beg1End)) {
		body["beg1_end"] = request.Beg1End
	}

	if !tea.BoolValue(util.IsUnset(request.Beg1Start)) {
		body["beg1_start"] = request.Beg1Start
	}

	if !tea.BoolValue(util.IsUnset(request.Beg2End)) {
		body["beg2_end"] = request.Beg2End
	}

	if !tea.BoolValue(util.IsUnset(request.Beg2Start)) {
		body["beg2_start"] = request.Beg2Start
	}

	if !tea.BoolValue(util.IsUnset(request.HotType)) {
		body["hot_type"] = request.HotType
	}

	if !tea.BoolValue(util.IsUnset(request.Instance1)) {
		body["instance1"] = request.Instance1
	}

	if !tea.BoolValue(util.IsUnset(request.Instance2)) {
		body["instance2"] = request.Instance2
	}

	if !tea.BoolValue(util.IsUnset(request.Pid1)) {
		body["pid1"] = request.Pid1
	}

	if !tea.BoolValue(util.IsUnset(request.Pid2)) {
		body["pid2"] = request.Pid2
	}

	if !tea.BoolValue(util.IsUnset(request.Table)) {
		body["table"] = request.Table
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotspotCompare"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/livetrace_hotspot_compare"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotspotCompareResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotspotCompareResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 热点对比
//
// @param request - GetHotspotCompareRequest
//
// @return GetHotspotCompareResponse
func (client *Client) GetHotspotCompare(request *GetHotspotCompareRequest) (_result *GetHotspotCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotCompareResponse{}
	_body, _err := client.GetHotspotCompareWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotspotInstanceListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotspotInstanceListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotspotPidListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotspotPidListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 发起热点追踪
//
// @param request - GetHotspotTrackingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotspotTrackingResponse
func (client *Client) GetHotspotTrackingWithOptions(request *GetHotspotTrackingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetHotspotTrackingResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.HotType)) {
		body["hot_type"] = request.HotType
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
		Action:      tea.String("GetHotspotTracking"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/livetrace_hotspot_tracking"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetHotspotTrackingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetHotspotTrackingResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 发起热点追踪
//
// @param request - GetHotspotTrackingRequest
//
// @return GetHotspotTrackingResponse
func (client *Client) GetHotspotTracking(request *GetHotspotTrackingRequest) (_result *GetHotspotTrackingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHotspotTrackingResponse{}
	_body, _err := client.GetHotspotTrackingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实时集群/节点健康度分数
//
// @param request - GetInstantScoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstantScoreResponse
func (client *Client) GetInstantScoreWithOptions(request *GetInstantScoreRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstantScoreResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstantScore"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/cluster_health/instant/score"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstantScoreResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstantScoreResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实时集群/节点健康度分数
//
// @param request - GetInstantScoreRequest
//
// @return GetInstantScoreResponse
func (client *Client) GetInstantScore(request *GetInstantScoreRequest) (_result *GetInstantScoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstantScoreResponse{}
	_body, _err := client.GetInstantScoreWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetListRecordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetListRecordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetProblemPercentageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetProblemPercentageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRangeScoreResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRangeScoreResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取功能模块配置
//
// @param tmpReq - GetServiceFuncStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceFuncStatusResponse
func (client *Client) GetServiceFuncStatusWithOptions(tmpReq *GetServiceFuncStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceFuncStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceFuncStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Params)) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, tea.String("params"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ParamsShrink)) {
		query["params"] = request.ParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["service_name"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceFuncStatus"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/func-switch/get-service-func-status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetServiceFuncStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetServiceFuncStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取功能模块配置
//
// @param request - GetServiceFuncStatusRequest
//
// @return GetServiceFuncStatusResponse
func (client *Client) GetServiceFuncStatus(request *GetServiceFuncStatusRequest) (_result *GetServiceFuncStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceFuncStatusResponse{}
	_body, _err := client.GetServiceFuncStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化SysOM，确保角色存在
//
// @param request - InitialSysomRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitialSysomResponse
func (client *Client) InitialSysomWithOptions(request *InitialSysomRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitialSysomResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckOnly)) {
		body["check_only"] = request.CheckOnly
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitialSysom"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/initial"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InitialSysomResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InitialSysomResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 初始化SysOM，确保角色存在
//
// @param request - InitialSysomRequest
//
// @return InitialSysomResponse
func (client *Client) InitialSysom(request *InitialSysomRequest) (_result *InitialSysomResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitialSysomResponse{}
	_body, _err := client.InitialSysomWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InstallAgentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InstallAgentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 给集群安装组件
//
// @param request - InstallAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentForClusterResponse
func (client *Client) InstallAgentForClusterWithOptions(request *InstallAgentForClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InstallAgentForClusterResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.GrayscaleConfig)) {
		body["grayscale_config"] = request.GrayscaleConfig
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallAgentForCluster"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/install_agent_by_cluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InstallAgentForClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InstallAgentForClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 给集群安装组件
//
// @param request - InstallAgentForClusterRequest
//
// @return InstallAgentForClusterResponse
func (client *Client) InstallAgentForCluster(request *InstallAgentForClusterRequest) (_result *InstallAgentForClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentForClusterResponse{}
	_body, _err := client.InstallAgentForClusterWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InvokeAnomalyDiagnosisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InvokeAnomalyDiagnosisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &InvokeDiagnosisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &InvokeDiagnosisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAbnormalyEventsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAbnormalyEventsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAgentInstallRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAgentInstallRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAgentsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAgentsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取集群组件安装记录
//
// @param request - ListClusterAgentInstallRecordsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterAgentInstallRecordsResponse
func (client *Client) ListClusterAgentInstallRecordsWithOptions(request *ListClusterAgentInstallRecordsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterAgentInstallRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterAgentInstallRecords"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/list_cluster_agent_install_list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListClusterAgentInstallRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListClusterAgentInstallRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取集群组件安装记录
//
// @param request - ListClusterAgentInstallRecordsRequest
//
// @return ListClusterAgentInstallRecordsResponse
func (client *Client) ListClusterAgentInstallRecords(request *ListClusterAgentInstallRecordsRequest) (_result *ListClusterAgentInstallRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterAgentInstallRecordsResponse{}
	_body, _err := client.ListClusterAgentInstallRecordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前用户的所有集群
//
// @param request - ListClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterStatus)) {
		query["cluster_status"] = request.ClusterStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["cluster_type"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/cluster/list_clusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListClustersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListClustersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取当前用户的所有集群
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取诊断历史记录列表
//
// @param request - ListDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosisResponse
func (client *Client) ListDiagnosisWithOptions(request *ListDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDiagnosisResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["service_name"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnosis"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/diagnosis/list_diagnosis"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListDiagnosisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListDiagnosisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取诊断历史记录列表
//
// @param request - ListDiagnosisRequest
//
// @return ListDiagnosisResponse
func (client *Client) ListDiagnosis(request *ListDiagnosisRequest) (_result *ListDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDiagnosisResponse{}
	_body, _err := client.ListDiagnosisWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstanceHealthResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstanceHealthResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取实例状态
//
// @param request - ListInstanceStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatusWithOptions(request *ListInstanceStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceStatus"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/instance/list_instance_status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstanceStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstanceStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例状态
//
// @param request - ListInstanceStatusRequest
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatus(request *ListInstanceStatusRequest) (_result *ListInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceStatusResponse{}
	_body, _err := client.ListInstanceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例列表
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/instance/list_instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例列表
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例中的pod列表
//
// @param request - ListPodsOfInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPodsOfInstanceResponse
func (client *Client) ListPodsOfInstanceWithOptions(request *ListPodsOfInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPodsOfInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		query["current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Instance)) {
		query["instance"] = request.Instance
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPodsOfInstance"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/instance/list_pod_of_instance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListPodsOfInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListPodsOfInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取实例中的pod列表
//
// @param request - ListPodsOfInstanceRequest
//
// @return ListPodsOfInstanceResponse
func (client *Client) ListPodsOfInstance(request *ListPodsOfInstanceRequest) (_result *ListPodsOfInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPodsOfInstanceResponse{}
	_body, _err := client.ListPodsOfInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出所有纳管了机器的区域
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/instance/list_regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 列出所有纳管了机器的区域
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartAIAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartAIAnalysisResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UninstallAgentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UninstallAgentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 给集群卸载组件
//
// @param request - UninstallAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallAgentForClusterResponse
func (client *Client) UninstallAgentForClusterWithOptions(request *UninstallAgentForClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UninstallAgentForClusterResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["cluster_id"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallAgentForCluster"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/uninstall_agent_by_cluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UninstallAgentForClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UninstallAgentForClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 给集群卸载组件
//
// @param request - UninstallAgentForClusterRequest
//
// @return UninstallAgentForClusterResponse
func (client *Client) UninstallAgentForCluster(request *UninstallAgentForClusterRequest) (_result *UninstallAgentForClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallAgentForClusterResponse{}
	_body, _err := client.UninstallAgentForClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 异常项关注度更新
//
// @param tmpReq - UpdateEventsAttentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventsAttentionResponse
func (client *Client) UpdateEventsAttentionWithOptions(tmpReq *UpdateEventsAttentionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEventsAttentionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEventsAttentionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Body)) {
		request.BodyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Body, tea.String("body"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BodyShrink)) {
		query["body"] = request.BodyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventsAttention"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/openapi/proxy/post/cluster_update_events_attention"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateEventsAttentionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateEventsAttentionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 异常项关注度更新
//
// @param request - UpdateEventsAttentionRequest
//
// @return UpdateEventsAttentionResponse
func (client *Client) UpdateEventsAttention(request *UpdateEventsAttentionRequest) (_result *UpdateEventsAttentionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEventsAttentionResponse{}
	_body, _err := client.UpdateEventsAttentionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取功能模块配置
//
// @param tmpReq - UpdateFuncSwitchRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFuncSwitchRecordResponse
func (client *Client) UpdateFuncSwitchRecordWithOptions(tmpReq *UpdateFuncSwitchRecordRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFuncSwitchRecordResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateFuncSwitchRecordShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Params)) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, tea.String("params"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Channel)) {
		query["channel"] = request.Channel
	}

	if !tea.BoolValue(util.IsUnset(request.ParamsShrink)) {
		query["params"] = request.ParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["service_name"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFuncSwitchRecord"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/func-switch/update-service-func-switch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateFuncSwitchRecordResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateFuncSwitchRecordResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取功能模块配置
//
// @param request - UpdateFuncSwitchRecordRequest
//
// @return UpdateFuncSwitchRecordResponse
func (client *Client) UpdateFuncSwitchRecord(request *UpdateFuncSwitchRecordRequest) (_result *UpdateFuncSwitchRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFuncSwitchRecordResponse{}
	_body, _err := client.UpdateFuncSwitchRecordWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeAgentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeAgentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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

// Summary:
//
// 给集群更新组件
//
// @param request - UpgradeAgentForClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeAgentForClusterResponse
func (client *Client) UpgradeAgentForClusterWithOptions(request *UpgradeAgentForClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpgradeAgentForClusterResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["cluster_id"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeAgentForCluster"),
		Version:     tea.String("2023-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/am/agent/upgrade_agent_by_cluster"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeAgentForClusterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeAgentForClusterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 给集群更新组件
//
// @param request - UpgradeAgentForClusterRequest
//
// @return UpgradeAgentForClusterResponse
func (client *Client) UpgradeAgentForCluster(request *UpgradeAgentForClusterRequest) (_result *UpgradeAgentForClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeAgentForClusterResponse{}
	_body, _err := client.UpgradeAgentForClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
