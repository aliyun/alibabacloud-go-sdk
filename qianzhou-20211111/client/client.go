// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AICreateSessionMessageRequest struct {
	Context   *AICreateSessionMessageRequestContext `json:"context,omitempty" xml:"context,omitempty" type:"Struct"`
	Language  *string                               `json:"language,omitempty" xml:"language,omitempty"`
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	SessionId *string                               `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// QA / K8S_DIAGNOSE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	EmployeeId *string `json:"employee_id,omitempty" xml:"employee_id,omitempty"`
}

func (s AICreateSessionMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s AICreateSessionMessageRequest) GoString() string {
	return s.String()
}

func (s *AICreateSessionMessageRequest) SetContext(v *AICreateSessionMessageRequestContext) *AICreateSessionMessageRequest {
	s.Context = v
	return s
}

func (s *AICreateSessionMessageRequest) SetLanguage(v string) *AICreateSessionMessageRequest {
	s.Language = &v
	return s
}

func (s *AICreateSessionMessageRequest) SetMessage(v string) *AICreateSessionMessageRequest {
	s.Message = &v
	return s
}

func (s *AICreateSessionMessageRequest) SetSessionId(v string) *AICreateSessionMessageRequest {
	s.SessionId = &v
	return s
}

func (s *AICreateSessionMessageRequest) SetType(v string) *AICreateSessionMessageRequest {
	s.Type = &v
	return s
}

func (s *AICreateSessionMessageRequest) SetEmployeeId(v string) *AICreateSessionMessageRequest {
	s.EmployeeId = &v
	return s
}

type AICreateSessionMessageRequestContext struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	Error     *string `json:"error,omitempty" xml:"error,omitempty"`
	Kind      *string `json:"kind,omitempty" xml:"kind,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Uuid      *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s AICreateSessionMessageRequestContext) String() string {
	return tea.Prettify(s)
}

func (s AICreateSessionMessageRequestContext) GoString() string {
	return s.String()
}

func (s *AICreateSessionMessageRequestContext) SetClusterId(v string) *AICreateSessionMessageRequestContext {
	s.ClusterId = &v
	return s
}

func (s *AICreateSessionMessageRequestContext) SetError(v string) *AICreateSessionMessageRequestContext {
	s.Error = &v
	return s
}

func (s *AICreateSessionMessageRequestContext) SetKind(v string) *AICreateSessionMessageRequestContext {
	s.Kind = &v
	return s
}

func (s *AICreateSessionMessageRequestContext) SetName(v string) *AICreateSessionMessageRequestContext {
	s.Name = &v
	return s
}

func (s *AICreateSessionMessageRequestContext) SetNamespace(v string) *AICreateSessionMessageRequestContext {
	s.Namespace = &v
	return s
}

func (s *AICreateSessionMessageRequestContext) SetUuid(v string) *AICreateSessionMessageRequestContext {
	s.Uuid = &v
	return s
}

type AICreateSessionMessageResponseBody struct {
	Answer    *string                                        `json:"answer,omitempty" xml:"answer,omitempty"`
	Code      *int64                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string                                        `json:"data,omitempty" xml:"data,omitempty"`
	Msg       *string                                        `json:"msg,omitempty" xml:"msg,omitempty"`
	Reference []*AICreateSessionMessageResponseBodyReference `json:"reference,omitempty" xml:"reference,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SessionId *string                                        `json:"session_id,omitempty" xml:"session_id,omitempty"`
}

func (s AICreateSessionMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AICreateSessionMessageResponseBody) GoString() string {
	return s.String()
}

func (s *AICreateSessionMessageResponseBody) SetAnswer(v string) *AICreateSessionMessageResponseBody {
	s.Answer = &v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetCode(v int64) *AICreateSessionMessageResponseBody {
	s.Code = &v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetData(v string) *AICreateSessionMessageResponseBody {
	s.Data = &v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetMsg(v string) *AICreateSessionMessageResponseBody {
	s.Msg = &v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetReference(v []*AICreateSessionMessageResponseBodyReference) *AICreateSessionMessageResponseBody {
	s.Reference = v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetRequestId(v string) *AICreateSessionMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AICreateSessionMessageResponseBody) SetSessionId(v string) *AICreateSessionMessageResponseBody {
	s.SessionId = &v
	return s
}

type AICreateSessionMessageResponseBodyReference struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url   *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s AICreateSessionMessageResponseBodyReference) String() string {
	return tea.Prettify(s)
}

func (s AICreateSessionMessageResponseBodyReference) GoString() string {
	return s.String()
}

func (s *AICreateSessionMessageResponseBodyReference) SetTitle(v string) *AICreateSessionMessageResponseBodyReference {
	s.Title = &v
	return s
}

func (s *AICreateSessionMessageResponseBodyReference) SetUrl(v string) *AICreateSessionMessageResponseBodyReference {
	s.Url = &v
	return s
}

type AICreateSessionMessageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AICreateSessionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AICreateSessionMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s AICreateSessionMessageResponse) GoString() string {
	return s.String()
}

func (s *AICreateSessionMessageResponse) SetHeaders(v map[string]*string) *AICreateSessionMessageResponse {
	s.Headers = v
	return s
}

func (s *AICreateSessionMessageResponse) SetStatusCode(v int32) *AICreateSessionMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *AICreateSessionMessageResponse) SetBody(v *AICreateSessionMessageResponseBody) *AICreateSessionMessageResponse {
	s.Body = v
	return s
}

type CreateDiagnosisRequest struct {
	Body          *string `json:"body,omitempty" xml:"body,omitempty"`
	ClusterID     *string `json:"clusterID,omitempty" xml:"clusterID,omitempty"`
	DiagnosisType *string `json:"diagnosisType,omitempty" xml:"diagnosisType,omitempty"`
}

func (s CreateDiagnosisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisRequest) SetBody(v string) *CreateDiagnosisRequest {
	s.Body = &v
	return s
}

func (s *CreateDiagnosisRequest) SetClusterID(v string) *CreateDiagnosisRequest {
	s.ClusterID = &v
	return s
}

func (s *CreateDiagnosisRequest) SetDiagnosisType(v string) *CreateDiagnosisRequest {
	s.DiagnosisType = &v
	return s
}

type CreateDiagnosisResponseBody struct {
	Code      *int64      `json:"code,omitempty" xml:"code,omitempty"`
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateDiagnosisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponseBody) SetCode(v int64) *CreateDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosisResponseBody) SetData(v interface{}) *CreateDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *CreateDiagnosisResponseBody) SetRequestId(v string) *CreateDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiagnosisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosisResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosisResponse) SetHeaders(v map[string]*string) *CreateDiagnosisResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosisResponse) SetStatusCode(v int32) *CreateDiagnosisResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosisResponse) SetBody(v *CreateDiagnosisResponseBody) *CreateDiagnosisResponse {
	s.Body = v
	return s
}

type GetClusterRequest struct {
	ClusterID *string `json:"clusterID,omitempty" xml:"clusterID,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetClusterID(v string) *GetClusterRequest {
	s.ClusterID = &v
	return s
}

type GetClusterResponseBody struct {
	Data      *GetClusterResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) SetData(v *GetClusterResponseBodyData) *GetClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterResponseBodyData struct {
	ClusterID       *string `json:"clusterID,omitempty" xml:"clusterID,omitempty"`
	ClusterType     *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	CurrentVersion  *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Profile         *string `json:"profile,omitempty" xml:"profile,omitempty"`
	RegionID        *string `json:"regionID,omitempty" xml:"regionID,omitempty"`
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
	State           *string `json:"state,omitempty" xml:"state,omitempty"`
	UserID          *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s GetClusterResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyData) SetClusterID(v string) *GetClusterResponseBodyData {
	s.ClusterID = &v
	return s
}

func (s *GetClusterResponseBodyData) SetClusterType(v string) *GetClusterResponseBodyData {
	s.ClusterType = &v
	return s
}

func (s *GetClusterResponseBodyData) SetCurrentVersion(v string) *GetClusterResponseBodyData {
	s.CurrentVersion = &v
	return s
}

func (s *GetClusterResponseBodyData) SetName(v string) *GetClusterResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetClusterResponseBodyData) SetProfile(v string) *GetClusterResponseBodyData {
	s.Profile = &v
	return s
}

func (s *GetClusterResponseBodyData) SetRegionID(v string) *GetClusterResponseBodyData {
	s.RegionID = &v
	return s
}

func (s *GetClusterResponseBodyData) SetSecurityGroupId(v string) *GetClusterResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *GetClusterResponseBodyData) SetSpec(v string) *GetClusterResponseBodyData {
	s.Spec = &v
	return s
}

func (s *GetClusterResponseBodyData) SetState(v string) *GetClusterResponseBodyData {
	s.State = &v
	return s
}

func (s *GetClusterResponseBodyData) SetUserID(v string) *GetClusterResponseBodyData {
	s.UserID = &v
	return s
}

type GetClusterResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

func (s *GetClusterResponse) SetStatusCode(v int32) *GetClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterResponse) SetBody(v *GetClusterResponseBody) *GetClusterResponse {
	s.Body = v
	return s
}

type GetClusterWarningRequest struct {
	ClusterID *string `json:"clusterID,omitempty" xml:"clusterID,omitempty"`
}

func (s GetClusterWarningRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterWarningRequest) GoString() string {
	return s.String()
}

func (s *GetClusterWarningRequest) SetClusterID(v string) *GetClusterWarningRequest {
	s.ClusterID = &v
	return s
}

type GetClusterWarningResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetClusterWarningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterWarningResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterWarningResponseBody) SetRequestId(v string) *GetClusterWarningResponseBody {
	s.RequestId = &v
	return s
}

type GetClusterWarningResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterWarningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterWarningResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterWarningResponse) GoString() string {
	return s.String()
}

func (s *GetClusterWarningResponse) SetHeaders(v map[string]*string) *GetClusterWarningResponse {
	s.Headers = v
	return s
}

func (s *GetClusterWarningResponse) SetStatusCode(v int32) *GetClusterWarningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterWarningResponse) SetBody(v *GetClusterWarningResponseBody) *GetClusterWarningResponse {
	s.Body = v
	return s
}

type GetDiagnosisResultRequest struct {
	DiagnosisId *string `json:"diagnosisId,omitempty" xml:"diagnosisId,omitempty"`
	OwnerUid    *string `json:"ownerUid,omitempty" xml:"ownerUid,omitempty"`
}

func (s GetDiagnosisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosisResultRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosisResultRequest) SetDiagnosisId(v string) *GetDiagnosisResultRequest {
	s.DiagnosisId = &v
	return s
}

func (s *GetDiagnosisResultRequest) SetOwnerUid(v string) *GetDiagnosisResultRequest {
	s.OwnerUid = &v
	return s
}

type GetDiagnosisResultResponseBody struct {
	Code      *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data      interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Msg       *string     `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
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

func (s *GetDiagnosisResultResponseBody) SetData(v interface{}) *GetDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetMsg(v string) *GetDiagnosisResultResponseBody {
	s.Msg = &v
	return s
}

func (s *GetDiagnosisResultResponseBody) SetRequestId(v string) *GetDiagnosisResultResponseBody {
	s.RequestId = &v
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

type GetUserClusterWarningRequest struct {
	UserID *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s GetUserClusterWarningRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserClusterWarningRequest) GoString() string {
	return s.String()
}

func (s *GetUserClusterWarningRequest) SetUserID(v string) *GetUserClusterWarningRequest {
	s.UserID = &v
	return s
}

type GetUserClusterWarningResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetUserClusterWarningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserClusterWarningResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserClusterWarningResponseBody) SetRequestId(v string) *GetUserClusterWarningResponseBody {
	s.RequestId = &v
	return s
}

type GetUserClusterWarningResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserClusterWarningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserClusterWarningResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserClusterWarningResponse) GoString() string {
	return s.String()
}

func (s *GetUserClusterWarningResponse) SetHeaders(v map[string]*string) *GetUserClusterWarningResponse {
	s.Headers = v
	return s
}

func (s *GetUserClusterWarningResponse) SetStatusCode(v int32) *GetUserClusterWarningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserClusterWarningResponse) SetBody(v *GetUserClusterWarningResponseBody) *GetUserClusterWarningResponse {
	s.Body = v
	return s
}

type GetWebshellTokenResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetWebshellTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWebshellTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebshellTokenResponseBody) SetRequestId(v string) *GetWebshellTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetWebshellTokenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWebshellTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWebshellTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWebshellTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWebshellTokenResponse) SetHeaders(v map[string]*string) *GetWebshellTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWebshellTokenResponse) SetStatusCode(v int32) *GetWebshellTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWebshellTokenResponse) SetBody(v *GetWebshellTokenResponseBody) *GetWebshellTokenResponse {
	s.Body = v
	return s
}

type HelloRequest struct {
	// example:
	//
	// binzhi
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s HelloRequest) String() string {
	return tea.Prettify(s)
}

func (s HelloRequest) GoString() string {
	return s.String()
}

func (s *HelloRequest) SetUser(v string) *HelloRequest {
	s.User = &v
	return s
}

type HelloResponseBody struct {
	// example:
	//
	// 200
	Code *string                `json:"code,omitempty" xml:"code,omitempty"`
	Data *HelloResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	Msg       *string `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s HelloResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HelloResponseBody) GoString() string {
	return s.String()
}

func (s *HelloResponseBody) SetCode(v string) *HelloResponseBody {
	s.Code = &v
	return s
}

func (s *HelloResponseBody) SetData(v *HelloResponseBodyData) *HelloResponseBody {
	s.Data = v
	return s
}

func (s *HelloResponseBody) SetMsg(v string) *HelloResponseBody {
	s.Msg = &v
	return s
}

func (s *HelloResponseBody) SetRequestId(v string) *HelloResponseBody {
	s.RequestId = &v
	return s
}

type HelloResponseBodyData struct {
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// ack1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HelloResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s HelloResponseBodyData) GoString() string {
	return s.String()
}

func (s *HelloResponseBodyData) SetDate(v string) *HelloResponseBodyData {
	s.Date = &v
	return s
}

func (s *HelloResponseBodyData) SetName(v string) *HelloResponseBodyData {
	s.Name = &v
	return s
}

func (s *HelloResponseBodyData) SetSuccess(v bool) *HelloResponseBodyData {
	s.Success = &v
	return s
}

type HelloResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HelloResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HelloResponse) String() string {
	return tea.Prettify(s)
}

func (s HelloResponse) GoString() string {
	return s.String()
}

func (s *HelloResponse) SetHeaders(v map[string]*string) *HelloResponse {
	s.Headers = v
	return s
}

func (s *HelloResponse) SetStatusCode(v int32) *HelloResponse {
	s.StatusCode = &v
	return s
}

func (s *HelloResponse) SetBody(v *HelloResponseBody) *HelloResponse {
	s.Body = v
	return s
}

type ListClusterDeprecatedAPIRequest struct {
	ClusterId     *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	PageNo        *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize      *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TargetVersion *string `json:"target_version,omitempty" xml:"target_version,omitempty"`
}

func (s ListClusterDeprecatedAPIRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterDeprecatedAPIRequest) GoString() string {
	return s.String()
}

func (s *ListClusterDeprecatedAPIRequest) SetClusterId(v string) *ListClusterDeprecatedAPIRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterDeprecatedAPIRequest) SetPageNo(v int32) *ListClusterDeprecatedAPIRequest {
	s.PageNo = &v
	return s
}

func (s *ListClusterDeprecatedAPIRequest) SetPageSize(v int32) *ListClusterDeprecatedAPIRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterDeprecatedAPIRequest) SetTargetVersion(v string) *ListClusterDeprecatedAPIRequest {
	s.TargetVersion = &v
	return s
}

type ListClusterDeprecatedAPIResponseBody struct {
	Code      *int32                                     `json:"code,omitempty" xml:"code,omitempty"`
	Datas     *ListClusterDeprecatedAPIResponseBodyDatas `json:"datas,omitempty" xml:"datas,omitempty" type:"Struct"`
	Msg       *string                                    `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListClusterDeprecatedAPIResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterDeprecatedAPIResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterDeprecatedAPIResponseBody) SetCode(v int32) *ListClusterDeprecatedAPIResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBody) SetDatas(v *ListClusterDeprecatedAPIResponseBodyDatas) *ListClusterDeprecatedAPIResponseBody {
	s.Datas = v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBody) SetMsg(v string) *ListClusterDeprecatedAPIResponseBody {
	s.Msg = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBody) SetRequestId(v string) *ListClusterDeprecatedAPIResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterDeprecatedAPIResponseBodyDatas struct {
	Current  *int32                                           `json:"current,omitempty" xml:"current,omitempty"`
	Data     []*ListClusterDeprecatedAPIResponseBodyDatasData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageSize *int32                                           `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                           `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClusterDeprecatedAPIResponseBodyDatas) String() string {
	return tea.Prettify(s)
}

func (s ListClusterDeprecatedAPIResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *ListClusterDeprecatedAPIResponseBodyDatas) SetCurrent(v int32) *ListClusterDeprecatedAPIResponseBodyDatas {
	s.Current = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatas) SetData(v []*ListClusterDeprecatedAPIResponseBodyDatasData) *ListClusterDeprecatedAPIResponseBodyDatas {
	s.Data = v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatas) SetPageSize(v int32) *ListClusterDeprecatedAPIResponseBodyDatas {
	s.PageSize = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatas) SetTotal(v int32) *ListClusterDeprecatedAPIResponseBodyDatas {
	s.Total = &v
	return s
}

type ListClusterDeprecatedAPIResponseBodyDatasData struct {
	ClusterId            *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	Content              *string `json:"content,omitempty" xml:"content,omitempty"`
	Count                *int32  `json:"count,omitempty" xml:"count,omitempty"`
	CurrentVersion       *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	DeprecatedK8sVersion *string `json:"deprecated_k8s_version,omitempty" xml:"deprecated_k8s_version,omitempty"`
	Ds                   *string `json:"ds,omitempty" xml:"ds,omitempty"`
	Label                *string `json:"label,omitempty" xml:"label,omitempty"`
	LastTime             *string `json:"last_time,omitempty" xml:"last_time,omitempty"`
	RequestUri           *string `json:"request_uri,omitempty" xml:"request_uri,omitempty"`
	SourceIps            *string `json:"source_ips,omitempty" xml:"source_ips,omitempty"`
	TargetVersion        *string `json:"target_version,omitempty" xml:"target_version,omitempty"`
	UserAgent            *string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
}

func (s ListClusterDeprecatedAPIResponseBodyDatasData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterDeprecatedAPIResponseBodyDatasData) GoString() string {
	return s.String()
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetClusterId(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetContent(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.Content = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetCount(v int32) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.Count = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetCurrentVersion(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.CurrentVersion = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetDeprecatedK8sVersion(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.DeprecatedK8sVersion = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetDs(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.Ds = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetLabel(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.Label = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetLastTime(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.LastTime = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetRequestUri(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.RequestUri = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetSourceIps(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.SourceIps = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetTargetVersion(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.TargetVersion = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponseBodyDatasData) SetUserAgent(v string) *ListClusterDeprecatedAPIResponseBodyDatasData {
	s.UserAgent = &v
	return s
}

type ListClusterDeprecatedAPIResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterDeprecatedAPIResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterDeprecatedAPIResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterDeprecatedAPIResponse) GoString() string {
	return s.String()
}

func (s *ListClusterDeprecatedAPIResponse) SetHeaders(v map[string]*string) *ListClusterDeprecatedAPIResponse {
	s.Headers = v
	return s
}

func (s *ListClusterDeprecatedAPIResponse) SetStatusCode(v int32) *ListClusterDeprecatedAPIResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterDeprecatedAPIResponse) SetBody(v *ListClusterDeprecatedAPIResponseBody) *ListClusterDeprecatedAPIResponse {
	s.Body = v
	return s
}

type ListClusterImagesRequest struct {
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	ImageHash *string `json:"image_hash,omitempty" xml:"image_hash,omitempty"`
	ImageName *string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	PageNo    *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize  *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	Uid       *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ListClusterImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterImagesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterImagesRequest) SetClusterId(v string) *ListClusterImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterImagesRequest) SetImageHash(v string) *ListClusterImagesRequest {
	s.ImageHash = &v
	return s
}

func (s *ListClusterImagesRequest) SetImageName(v string) *ListClusterImagesRequest {
	s.ImageName = &v
	return s
}

func (s *ListClusterImagesRequest) SetPageNo(v int32) *ListClusterImagesRequest {
	s.PageNo = &v
	return s
}

func (s *ListClusterImagesRequest) SetPageSize(v int32) *ListClusterImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterImagesRequest) SetUid(v string) *ListClusterImagesRequest {
	s.Uid = &v
	return s
}

type ListClusterImagesResponseBody struct {
	Code      *int32                              `json:"code,omitempty" xml:"code,omitempty"`
	Datas     *ListClusterImagesResponseBodyDatas `json:"datas,omitempty" xml:"datas,omitempty" type:"Struct"`
	Msg       *string                             `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListClusterImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterImagesResponseBody) SetCode(v int32) *ListClusterImagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterImagesResponseBody) SetDatas(v *ListClusterImagesResponseBodyDatas) *ListClusterImagesResponseBody {
	s.Datas = v
	return s
}

func (s *ListClusterImagesResponseBody) SetMsg(v string) *ListClusterImagesResponseBody {
	s.Msg = &v
	return s
}

func (s *ListClusterImagesResponseBody) SetRequestId(v string) *ListClusterImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterImagesResponseBodyDatas struct {
	Current  *int32                                    `json:"current,omitempty" xml:"current,omitempty"`
	Data     []*ListClusterImagesResponseBodyDatasData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	PageSize *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListClusterImagesResponseBodyDatas) String() string {
	return tea.Prettify(s)
}

func (s ListClusterImagesResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *ListClusterImagesResponseBodyDatas) SetCurrent(v int32) *ListClusterImagesResponseBodyDatas {
	s.Current = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatas) SetData(v []*ListClusterImagesResponseBodyDatasData) *ListClusterImagesResponseBodyDatas {
	s.Data = v
	return s
}

func (s *ListClusterImagesResponseBodyDatas) SetPageSize(v int32) *ListClusterImagesResponseBodyDatas {
	s.PageSize = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatas) SetTotal(v int32) *ListClusterImagesResponseBodyDatas {
	s.Total = &v
	return s
}

type ListClusterImagesResponseBodyDatasData struct {
	ClusterId      *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	ControllerName *string `json:"controller_name,omitempty" xml:"controller_name,omitempty"`
	ControllerType *string `json:"controller_type,omitempty" xml:"controller_type,omitempty"`
	Created        *string `json:"created,omitempty" xml:"created,omitempty"`
	ImageHash      *string `json:"image_hash,omitempty" xml:"image_hash,omitempty"`
	ImageName      *string `json:"image_name,omitempty" xml:"image_name,omitempty"`
	Namespace      *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Uid            *string `json:"uid,omitempty" xml:"uid,omitempty"`
	Updated        *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListClusterImagesResponseBodyDatasData) String() string {
	return tea.Prettify(s)
}

func (s ListClusterImagesResponseBodyDatasData) GoString() string {
	return s.String()
}

func (s *ListClusterImagesResponseBodyDatasData) SetClusterId(v string) *ListClusterImagesResponseBodyDatasData {
	s.ClusterId = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetControllerName(v string) *ListClusterImagesResponseBodyDatasData {
	s.ControllerName = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetControllerType(v string) *ListClusterImagesResponseBodyDatasData {
	s.ControllerType = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetCreated(v string) *ListClusterImagesResponseBodyDatasData {
	s.Created = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetImageHash(v string) *ListClusterImagesResponseBodyDatasData {
	s.ImageHash = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetImageName(v string) *ListClusterImagesResponseBodyDatasData {
	s.ImageName = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetNamespace(v string) *ListClusterImagesResponseBodyDatasData {
	s.Namespace = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetUid(v string) *ListClusterImagesResponseBodyDatasData {
	s.Uid = &v
	return s
}

func (s *ListClusterImagesResponseBodyDatasData) SetUpdated(v string) *ListClusterImagesResponseBodyDatasData {
	s.Updated = &v
	return s
}

type ListClusterImagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterImagesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterImagesResponse) SetHeaders(v map[string]*string) *ListClusterImagesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterImagesResponse) SetStatusCode(v int32) *ListClusterImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterImagesResponse) SetBody(v *ListClusterImagesResponseBody) *ListClusterImagesResponse {
	s.Body = v
	return s
}

type ListUserClustersRequest struct {
	Current  *string `json:"current,omitempty" xml:"current,omitempty"`
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	UserID   *string `json:"userID,omitempty" xml:"userID,omitempty"`
}

func (s ListUserClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserClustersRequest) GoString() string {
	return s.String()
}

func (s *ListUserClustersRequest) SetCurrent(v string) *ListUserClustersRequest {
	s.Current = &v
	return s
}

func (s *ListUserClustersRequest) SetPageSize(v string) *ListUserClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserClustersRequest) SetUserID(v string) *ListUserClustersRequest {
	s.UserID = &v
	return s
}

type ListUserClustersResponseBody struct {
	Code      *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	Msg       *string `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUserClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClustersResponseBody) SetCode(v int64) *ListUserClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserClustersResponseBody) SetData(v string) *ListUserClustersResponseBody {
	s.Data = &v
	return s
}

func (s *ListUserClustersResponseBody) SetMsg(v string) *ListUserClustersResponseBody {
	s.Msg = &v
	return s
}

func (s *ListUserClustersResponseBody) SetRequestId(v string) *ListUserClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListUserClustersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserClustersResponse) GoString() string {
	return s.String()
}

func (s *ListUserClustersResponse) SetHeaders(v map[string]*string) *ListUserClustersResponse {
	s.Headers = v
	return s
}

func (s *ListUserClustersResponse) SetStatusCode(v int32) *ListUserClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserClustersResponse) SetBody(v *ListUserClustersResponseBody) *ListUserClustersResponse {
	s.Body = v
	return s
}

type QueryByInstanceInfoRequest struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryByInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryByInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryByInstanceInfoRequest) SetInstanceId(v []*string) *QueryByInstanceInfoRequest {
	s.InstanceId = v
	return s
}

func (s *QueryByInstanceInfoRequest) SetRegionId(v string) *QueryByInstanceInfoRequest {
	s.RegionId = &v
	return s
}

type QueryByInstanceInfoResponseBody struct {
	Code      *int32                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data      []*QueryByInstanceInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg       *string                                `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryByInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryByInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryByInstanceInfoResponseBody) SetCode(v int32) *QueryByInstanceInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryByInstanceInfoResponseBody) SetData(v []*QueryByInstanceInfoResponseBodyData) *QueryByInstanceInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryByInstanceInfoResponseBody) SetMsg(v string) *QueryByInstanceInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *QueryByInstanceInfoResponseBody) SetRequestId(v string) *QueryByInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryByInstanceInfoResponseBodyData struct {
	ClusterId   *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RegionId    *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	Uid         *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s QueryByInstanceInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryByInstanceInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryByInstanceInfoResponseBodyData) SetClusterId(v string) *QueryByInstanceInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *QueryByInstanceInfoResponseBodyData) SetClusterName(v string) *QueryByInstanceInfoResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *QueryByInstanceInfoResponseBodyData) SetInstanceId(v string) *QueryByInstanceInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryByInstanceInfoResponseBodyData) SetRegionId(v string) *QueryByInstanceInfoResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *QueryByInstanceInfoResponseBodyData) SetUid(v string) *QueryByInstanceInfoResponseBodyData {
	s.Uid = &v
	return s
}

type QueryByInstanceInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryByInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryByInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryByInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryByInstanceInfoResponse) SetHeaders(v map[string]*string) *QueryByInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryByInstanceInfoResponse) SetStatusCode(v int32) *QueryByInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryByInstanceInfoResponse) SetBody(v *QueryByInstanceInfoResponseBody) *QueryByInstanceInfoResponse {
	s.Body = v
	return s
}

type QueryNodeInfoRequest struct {
	// This parameter is required.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s QueryNodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryNodeInfoRequest) SetInstanceId(v string) *QueryNodeInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryNodeInfoRequest) SetRegionId(v string) *QueryNodeInfoRequest {
	s.RegionId = &v
	return s
}

type QueryNodeInfoResponseBody struct {
	Code      *int32                           `json:"code,omitempty" xml:"code,omitempty"`
	Data      []*QueryNodeInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	Msg       *string                          `json:"msg,omitempty" xml:"msg,omitempty"`
	RequestId *string                          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryNodeInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryNodeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryNodeInfoResponseBody) SetCode(v int32) *QueryNodeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryNodeInfoResponseBody) SetData(v []*QueryNodeInfoResponseBodyData) *QueryNodeInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryNodeInfoResponseBody) SetMsg(v string) *QueryNodeInfoResponseBody {
	s.Msg = &v
	return s
}

func (s *QueryNodeInfoResponseBody) SetRequestId(v string) *QueryNodeInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryNodeInfoResponseBodyData struct {
	CluserName *string `json:"cluserName,omitempty" xml:"cluserName,omitempty"`
	ClusterId  *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region     *string `json:"region,omitempty" xml:"region,omitempty"`
	Uid        *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s QueryNodeInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryNodeInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryNodeInfoResponseBodyData) SetCluserName(v string) *QueryNodeInfoResponseBodyData {
	s.CluserName = &v
	return s
}

func (s *QueryNodeInfoResponseBodyData) SetClusterId(v string) *QueryNodeInfoResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *QueryNodeInfoResponseBodyData) SetInstanceId(v string) *QueryNodeInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryNodeInfoResponseBodyData) SetRegion(v string) *QueryNodeInfoResponseBodyData {
	s.Region = &v
	return s
}

func (s *QueryNodeInfoResponseBodyData) SetUid(v string) *QueryNodeInfoResponseBodyData {
	s.Uid = &v
	return s
}

type QueryNodeInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryNodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryNodeInfoResponse) SetHeaders(v map[string]*string) *QueryNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryNodeInfoResponse) SetStatusCode(v int32) *QueryNodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryNodeInfoResponse) SetBody(v *QueryNodeInfoResponseBody) *QueryNodeInfoResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("qianzhou"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// ACK AI助手千舟InnerAPI
//
// @param request - AICreateSessionMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AICreateSessionMessageResponse
func (client *Client) AICreateSessionMessageWithOptions(request *AICreateSessionMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AICreateSessionMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EmployeeId)) {
		query["employee_id"] = request.EmployeeId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Context)) {
		body["context"] = request.Context
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		body["language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["session_id"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AICreateSessionMessage"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/AICreateSessionMessage"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AICreateSessionMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ACK AI助手千舟InnerAPI
//
// @param request - AICreateSessionMessageRequest
//
// @return AICreateSessionMessageResponse
func (client *Client) AICreateSessionMessage(request *AICreateSessionMessageRequest) (_result *AICreateSessionMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AICreateSessionMessageResponse{}
	_body, _err := client.AICreateSessionMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// CreateDiagnosis
//
// @param request - CreateDiagnosisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosisResponse
func (client *Client) CreateDiagnosisWithOptions(request *CreateDiagnosisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDiagnosisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterID)) {
		query["clusterID"] = request.ClusterID
	}

	if !tea.BoolValue(util.IsUnset(request.DiagnosisType)) {
		query["diagnosisType"] = request.DiagnosisType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnosis"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/createDiagnosis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CreateDiagnosis
//
// @param request - CreateDiagnosisRequest
//
// @return CreateDiagnosisResponse
func (client *Client) CreateDiagnosis(request *CreateDiagnosisRequest) (_result *CreateDiagnosisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDiagnosisResponse{}
	_body, _err := client.CreateDiagnosisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群信息
//
// @param request - GetClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterID)) {
		query["clusterID"] = request.ClusterID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCluster"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/getCluster"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群信息
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 单个集群的高危风险项
//
// @param request - GetClusterWarningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterWarningResponse
func (client *Client) GetClusterWarningWithOptions(request *GetClusterWarningRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterWarningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterID)) {
		query["clusterID"] = request.ClusterID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClusterWarning"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/getKeyClusterWarningList"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClusterWarningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单个集群的高危风险项
//
// @param request - GetClusterWarningRequest
//
// @return GetClusterWarningResponse
func (client *Client) GetClusterWarning(request *GetClusterWarningRequest) (_result *GetClusterWarningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterWarningResponse{}
	_body, _err := client.GetClusterWarningWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// GetDiagnosisResult
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
	if !tea.BoolValue(util.IsUnset(request.DiagnosisId)) {
		query["diagnosisId"] = request.DiagnosisId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerUid)) {
		query["ownerUid"] = request.OwnerUid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnosisResult"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/GetDiagnosisResult"),
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
// GetDiagnosisResult
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
// 单个用户的高危集群
//
// @param request - GetUserClusterWarningRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserClusterWarningResponse
func (client *Client) GetUserClusterWarningWithOptions(request *GetUserClusterWarningRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserClusterWarningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserID)) {
		query["userID"] = request.UserID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserClusterWarning"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/listUserKeyClusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserClusterWarningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单个用户的高危集群
//
// @param request - GetUserClusterWarningRequest
//
// @return GetUserClusterWarningResponse
func (client *Client) GetUserClusterWarning(request *GetUserClusterWarningRequest) (_result *GetUserClusterWarningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserClusterWarningResponse{}
	_body, _err := client.GetUserClusterWarningWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetWebshellToken is deprecated
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebshellTokenResponse
// Deprecated
func (client *Client) GetWebshellTokenWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWebshellTokenResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWebshellToken"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/getChorusToken"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWebshellTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetWebshellToken is deprecated
//
// @return GetWebshellTokenResponse
// Deprecated
func (client *Client) GetWebshellToken() (_result *GetWebshellTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWebshellTokenResponse{}
	_body, _err := client.GetWebshellTokenWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 测试接口
//
// @param request - HelloRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HelloResponse
func (client *Client) HelloWithOptions(request *HelloRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HelloResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.User)) {
		query["user"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Hello"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/hello"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &HelloResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 测试接口
//
// @param request - HelloRequest
//
// @return HelloResponse
func (client *Client) Hello(request *HelloRequest) (_result *HelloResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HelloResponse{}
	_body, _err := client.HelloWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某集群废弃api的统计情况
//
// @param request - ListClusterDeprecatedAPIRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterDeprecatedAPIResponse
func (client *Client) ListClusterDeprecatedAPIWithOptions(request *ListClusterDeprecatedAPIRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterDeprecatedAPIResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		query["target_version"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterDeprecatedAPI"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/listDeprecatedK8sAPI"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterDeprecatedAPIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某集群废弃api的统计情况
//
// @param request - ListClusterDeprecatedAPIRequest
//
// @return ListClusterDeprecatedAPIResponse
func (client *Client) ListClusterDeprecatedAPI(request *ListClusterDeprecatedAPIRequest) (_result *ListClusterDeprecatedAPIResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterDeprecatedAPIResponse{}
	_body, _err := client.ListClusterDeprecatedAPIWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某集群中的Image列表
//
// @param request - ListClusterImagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterImagesResponse
func (client *Client) ListClusterImagesWithOptions(request *ListClusterImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["cluster_id"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageHash)) {
		query["image_hash"] = request.ImageHash
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["image_name"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["page_no"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["page_size"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterImages"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/listClusterPodImages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某集群中的Image列表
//
// @param request - ListClusterImagesRequest
//
// @return ListClusterImagesResponse
func (client *Client) ListClusterImages(request *ListClusterImagesRequest) (_result *ListClusterImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterImagesResponse{}
	_body, _err := client.ListClusterImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户下的集群列表
//
// @param request - ListUserClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserClustersResponse
func (client *Client) ListUserClustersWithOptions(request *ListUserClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserClustersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserID)) {
		query["userID"] = request.UserID
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserClusters"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/listUserClusters"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户下的集群列表
//
// @param request - ListUserClustersRequest
//
// @return ListUserClustersResponse
func (client *Client) ListUserClusters(request *ListUserClustersRequest) (_result *ListUserClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserClustersResponse{}
	_body, _err := client.ListUserClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询节点相关集群、用户信息
//
// @param request - QueryByInstanceInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryByInstanceInfoResponse
func (client *Client) QueryByInstanceInfoWithOptions(request *QueryByInstanceInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryByInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryByInstanceInfo"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/queryByInstanceInfo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryByInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询节点相关集群、用户信息
//
// @param request - QueryByInstanceInfoRequest
//
// @return QueryByInstanceInfoResponse
func (client *Client) QueryByInstanceInfo(request *QueryByInstanceInfoRequest) (_result *QueryByInstanceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryByInstanceInfoResponse{}
	_body, _err := client.QueryByInstanceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据节点instanceId查询集群/用户信息
//
// @param request - QueryNodeInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryNodeInfoResponse
func (client *Client) QueryNodeInfoWithOptions(request *QueryNodeInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryNodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryNodeInfo"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/popapi/queryByInstanceId"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryNodeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据节点instanceId查询集群/用户信息
//
// @param request - QueryNodeInfoRequest
//
// @return QueryNodeInfoResponse
func (client *Client) QueryNodeInfo(request *QueryNodeInfoRequest) (_result *QueryNodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryNodeInfoResponse{}
	_body, _err := client.QueryNodeInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
