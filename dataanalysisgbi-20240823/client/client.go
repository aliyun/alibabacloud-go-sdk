// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelDatasourceAuthorizationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CancelDatasourceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelDatasourceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *CancelDatasourceAuthorizationRequest) SetWorkspaceId(v string) *CancelDatasourceAuthorizationRequest {
	s.WorkspaceId = &v
	return s
}

type CancelDatasourceAuthorizationResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// FB11F719-9AC8-5C99-AB0A-4709D437FCFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CancelDatasourceAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelDatasourceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDatasourceAuthorizationResponseBody) SetCode(v string) *CancelDatasourceAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *CancelDatasourceAuthorizationResponseBody) SetData(v interface{}) *CancelDatasourceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *CancelDatasourceAuthorizationResponseBody) SetErrorMsg(v string) *CancelDatasourceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CancelDatasourceAuthorizationResponseBody) SetRequestId(v string) *CancelDatasourceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDatasourceAuthorizationResponseBody) SetSuccess(v bool) *CancelDatasourceAuthorizationResponseBody {
	s.Success = &v
	return s
}

type CancelDatasourceAuthorizationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelDatasourceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelDatasourceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelDatasourceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *CancelDatasourceAuthorizationResponse) SetHeaders(v map[string]*string) *CancelDatasourceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *CancelDatasourceAuthorizationResponse) SetStatusCode(v int32) *CancelDatasourceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelDatasourceAuthorizationResponse) SetBody(v *CancelDatasourceAuthorizationResponseBody) *CancelDatasourceAuthorizationResponse {
	s.Body = v
	return s
}

type CreateDatasourceAuthorizationRequest struct {
	// example:
	//
	// password
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// jdbc:mysql://rm-2zedvv990c8d8rj8ejo.mysql.rds.aliyuncs.com:3306/gbi_good_case
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// example:
	//
	// root
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
	// example:
	//
	// vdb-E0F693C8-9F72-5830-B81A-696C9D8EBBD1
	VdbId *string `json:"vdbId,omitempty" xml:"vdbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateDatasourceAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasourceAuthorizationRequest) SetPassword(v string) *CreateDatasourceAuthorizationRequest {
	s.Password = &v
	return s
}

func (s *CreateDatasourceAuthorizationRequest) SetType(v int32) *CreateDatasourceAuthorizationRequest {
	s.Type = &v
	return s
}

func (s *CreateDatasourceAuthorizationRequest) SetUrl(v string) *CreateDatasourceAuthorizationRequest {
	s.Url = &v
	return s
}

func (s *CreateDatasourceAuthorizationRequest) SetUserName(v string) *CreateDatasourceAuthorizationRequest {
	s.UserName = &v
	return s
}

func (s *CreateDatasourceAuthorizationRequest) SetVdbId(v string) *CreateDatasourceAuthorizationRequest {
	s.VdbId = &v
	return s
}

func (s *CreateDatasourceAuthorizationRequest) SetWorkspaceId(v string) *CreateDatasourceAuthorizationRequest {
	s.WorkspaceId = &v
	return s
}

type CreateDatasourceAuthorizationResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// E0F693C8-9F72-5830-B81A-696C9D8EBBD1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDatasourceAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasourceAuthorizationResponseBody) SetCode(v string) *CreateDatasourceAuthorizationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasourceAuthorizationResponseBody) SetData(v interface{}) *CreateDatasourceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *CreateDatasourceAuthorizationResponseBody) SetErrorMsg(v string) *CreateDatasourceAuthorizationResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateDatasourceAuthorizationResponseBody) SetRequestId(v string) *CreateDatasourceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasourceAuthorizationResponseBody) SetSuccess(v bool) *CreateDatasourceAuthorizationResponseBody {
	s.Success = &v
	return s
}

type CreateDatasourceAuthorizationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatasourceAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatasourceAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasourceAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasourceAuthorizationResponse) SetHeaders(v map[string]*string) *CreateDatasourceAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasourceAuthorizationResponse) SetStatusCode(v int32) *CreateDatasourceAuthorizationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasourceAuthorizationResponse) SetBody(v *CreateDatasourceAuthorizationResponseBody) *CreateDatasourceAuthorizationResponse {
	s.Body = v
	return s
}

type CreateVirtualDatasourceInstanceRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// virtual-instance-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Type *int32  `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateVirtualDatasourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualDatasourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualDatasourceInstanceRequest) SetDescription(v string) *CreateVirtualDatasourceInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceRequest) SetName(v string) *CreateVirtualDatasourceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceRequest) SetType(v int32) *CreateVirtualDatasourceInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceRequest) SetWorkspaceId(v string) *CreateVirtualDatasourceInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type CreateVirtualDatasourceInstanceResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// D02D895A-5E58-5A9F-963D-D8B027AB7AE2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateVirtualDatasourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualDatasourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVirtualDatasourceInstanceResponseBody) SetCode(v string) *CreateVirtualDatasourceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponseBody) SetData(v interface{}) *CreateVirtualDatasourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponseBody) SetErrorMsg(v string) *CreateVirtualDatasourceInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponseBody) SetRequestId(v string) *CreateVirtualDatasourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponseBody) SetSuccess(v bool) *CreateVirtualDatasourceInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateVirtualDatasourceInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualDatasourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualDatasourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualDatasourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualDatasourceInstanceResponse) SetHeaders(v map[string]*string) *CreateVirtualDatasourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponse) SetStatusCode(v int32) *CreateVirtualDatasourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualDatasourceInstanceResponse) SetBody(v *CreateVirtualDatasourceInstanceResponseBody) *CreateVirtualDatasourceInstanceResponse {
	s.Body = v
	return s
}

type DeleteVirtualDatasourceInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vdb-7D63529B-5D42-5BF0-84E4-F742FFE02E7F
	VdbId *string `json:"vdbId,omitempty" xml:"vdbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteVirtualDatasourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualDatasourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualDatasourceInstanceRequest) SetVdbId(v string) *DeleteVirtualDatasourceInstanceRequest {
	s.VdbId = &v
	return s
}

func (s *DeleteVirtualDatasourceInstanceRequest) SetWorkspaceId(v string) *DeleteVirtualDatasourceInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteVirtualDatasourceInstanceResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 7D63529B-5D42-5BF0-84E4-F742FFE02E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteVirtualDatasourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualDatasourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVirtualDatasourceInstanceResponseBody) SetCode(v string) *DeleteVirtualDatasourceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponseBody) SetData(v interface{}) *DeleteVirtualDatasourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponseBody) SetErrorMsg(v string) *DeleteVirtualDatasourceInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponseBody) SetRequestId(v string) *DeleteVirtualDatasourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponseBody) SetSuccess(v bool) *DeleteVirtualDatasourceInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteVirtualDatasourceInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualDatasourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualDatasourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualDatasourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualDatasourceInstanceResponse) SetHeaders(v map[string]*string) *DeleteVirtualDatasourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponse) SetStatusCode(v int32) *DeleteVirtualDatasourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualDatasourceInstanceResponse) SetBody(v *DeleteVirtualDatasourceInstanceResponseBody) *DeleteVirtualDatasourceInstanceResponse {
	s.Body = v
	return s
}

type ListVirtualDatasourceInstanceRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// e8Z0nRyY51ZQmYljqGNK
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListVirtualDatasourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualDatasourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualDatasourceInstanceRequest) SetMaxResults(v int32) *ListVirtualDatasourceInstanceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVirtualDatasourceInstanceRequest) SetNextToken(v string) *ListVirtualDatasourceInstanceRequest {
	s.NextToken = &v
	return s
}

func (s *ListVirtualDatasourceInstanceRequest) SetWorkspaceId(v string) *ListVirtualDatasourceInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type ListVirtualDatasourceInstanceResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 45390C6D-016D-5030-BF65-031ED1F65003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListVirtualDatasourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualDatasourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirtualDatasourceInstanceResponseBody) SetCode(v string) *ListVirtualDatasourceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListVirtualDatasourceInstanceResponseBody) SetData(v interface{}) *ListVirtualDatasourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListVirtualDatasourceInstanceResponseBody) SetErrorMsg(v string) *ListVirtualDatasourceInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListVirtualDatasourceInstanceResponseBody) SetRequestId(v string) *ListVirtualDatasourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirtualDatasourceInstanceResponseBody) SetSuccess(v bool) *ListVirtualDatasourceInstanceResponseBody {
	s.Success = &v
	return s
}

type ListVirtualDatasourceInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualDatasourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualDatasourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualDatasourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualDatasourceInstanceResponse) SetHeaders(v map[string]*string) *ListVirtualDatasourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualDatasourceInstanceResponse) SetStatusCode(v int32) *ListVirtualDatasourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualDatasourceInstanceResponse) SetBody(v *ListVirtualDatasourceInstanceResponseBody) *ListVirtualDatasourceInstanceResponse {
	s.Body = v
	return s
}

type RunDataAnalysisRequest struct {
	// example:
	//
	// true
	GenerateSqlOnly *bool `json:"generateSqlOnly,omitempty" xml:"generateSqlOnly,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// sessionID
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// STANDARD_MIX
	SpecificationType *string `json:"specificationType,omitempty" xml:"specificationType,omitempty"`
}

func (s RunDataAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisRequest) SetGenerateSqlOnly(v bool) *RunDataAnalysisRequest {
	s.GenerateSqlOnly = &v
	return s
}

func (s *RunDataAnalysisRequest) SetQuery(v string) *RunDataAnalysisRequest {
	s.Query = &v
	return s
}

func (s *RunDataAnalysisRequest) SetSessionId(v string) *RunDataAnalysisRequest {
	s.SessionId = &v
	return s
}

func (s *RunDataAnalysisRequest) SetSpecificationType(v string) *RunDataAnalysisRequest {
	s.SpecificationType = &v
	return s
}

type RunDataAnalysisResponseBody struct {
	Code           *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data           *RunDataAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int64                           `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                          `json:"message,omitempty" xml:"message,omitempty"`
}

func (s RunDataAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBody) SetCode(v string) *RunDataAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *RunDataAnalysisResponseBody) SetData(v *RunDataAnalysisResponseBodyData) *RunDataAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *RunDataAnalysisResponseBody) SetHttpStatusCode(v int64) *RunDataAnalysisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunDataAnalysisResponseBody) SetMessage(v string) *RunDataAnalysisResponseBody {
	s.Message = &v
	return s
}

type RunDataAnalysisResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/tech-scp-chain7.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// rewrite
	Event          *string `json:"event,omitempty" xml:"event,omitempty"`
	Evidence       *string `json:"evidence,omitempty" xml:"evidence,omitempty"`
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// DA2578F7-88A5-5D6E-9305-33E724E97D60
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rewrite   *string   `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
	Selector  []*string `json:"selector,omitempty" xml:"selector,omitempty" type:"Repeated"`
	// example:
	//
	// sessionid1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// select p.product_id, p.product_name, sum(o.quantity) as total_sales from products p join orders o on p.product_id = o.product_id where o.order_date between \\"2022-10-22\\" and \\"2024-10-22\\" group by p.product_id, p.product_name having total_sales > 5
	Sql     *string                                 `json:"sql,omitempty" xml:"sql,omitempty"`
	SqlData *RunDataAnalysisResponseBodyDataSqlData `json:"sqlData,omitempty" xml:"sqlData,omitempty" type:"Struct"`
	// example:
	//
	// Can not issue data manipulation statements with executeQuery()
	SqlError      *string                                       `json:"sqlError,omitempty" xml:"sqlError,omitempty"`
	Visualization *RunDataAnalysisResponseBodyDataVisualization `json:"visualization,omitempty" xml:"visualization,omitempty" type:"Struct"`
}

func (s RunDataAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyData) SetErrorMessage(v string) *RunDataAnalysisResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetEvent(v string) *RunDataAnalysisResponseBodyData {
	s.Event = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetEvidence(v string) *RunDataAnalysisResponseBodyData {
	s.Evidence = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetHttpStatusCode(v int64) *RunDataAnalysisResponseBodyData {
	s.HttpStatusCode = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetRequestId(v string) *RunDataAnalysisResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetRewrite(v string) *RunDataAnalysisResponseBodyData {
	s.Rewrite = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSelector(v []*string) *RunDataAnalysisResponseBodyData {
	s.Selector = v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSessionId(v string) *RunDataAnalysisResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSql(v string) *RunDataAnalysisResponseBodyData {
	s.Sql = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSqlData(v *RunDataAnalysisResponseBodyDataSqlData) *RunDataAnalysisResponseBodyData {
	s.SqlData = v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSqlError(v string) *RunDataAnalysisResponseBodyData {
	s.SqlError = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetVisualization(v *RunDataAnalysisResponseBodyDataVisualization) *RunDataAnalysisResponseBodyData {
	s.Visualization = v
	return s
}

type RunDataAnalysisResponseBodyDataSqlData struct {
	Column []*string                `json:"column,omitempty" xml:"column,omitempty" type:"Repeated"`
	Data   []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s RunDataAnalysisResponseBodyDataSqlData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataSqlData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataSqlData) SetColumn(v []*string) *RunDataAnalysisResponseBodyDataSqlData {
	s.Column = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataSqlData) SetData(v []map[string]interface{}) *RunDataAnalysisResponseBodyDataSqlData {
	s.Data = v
	return s
}

type RunDataAnalysisResponseBodyDataVisualization struct {
	Data *RunDataAnalysisResponseBodyDataVisualizationData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Text *string                                           `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunDataAnalysisResponseBodyDataVisualization) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataVisualization) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataVisualization) SetData(v *RunDataAnalysisResponseBodyDataVisualizationData) *RunDataAnalysisResponseBodyDataVisualization {
	s.Data = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualization) SetText(v string) *RunDataAnalysisResponseBodyDataVisualization {
	s.Text = &v
	return s
}

type RunDataAnalysisResponseBodyDataVisualizationData struct {
	// example:
	//
	// bar
	PlotType *string   `json:"plotType,omitempty" xml:"plotType,omitempty"`
	XAxis    []*string `json:"xAxis,omitempty" xml:"xAxis,omitempty" type:"Repeated"`
	YAxis    []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
}

func (s RunDataAnalysisResponseBodyDataVisualizationData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataVisualizationData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetPlotType(v string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.PlotType = &v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetXAxis(v []*string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.XAxis = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetYAxis(v []*string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.YAxis = v
	return s
}

type RunDataAnalysisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDataAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDataAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponse) SetHeaders(v map[string]*string) *RunDataAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunDataAnalysisResponse) SetStatusCode(v int32) *RunDataAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDataAnalysisResponse) SetBody(v *RunDataAnalysisResponseBody) *RunDataAnalysisResponse {
	s.Body = v
	return s
}

type SaveVirtualDatasourceDdlRequest struct {
	// This parameter is required.
	Ddl *string `json:"ddl,omitempty" xml:"ddl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdb-E0F693C8-9F72-5830-B81A-696C9D8EBBD1
	VdbId *string `json:"vdbId,omitempty" xml:"vdbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SaveVirtualDatasourceDdlRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveVirtualDatasourceDdlRequest) GoString() string {
	return s.String()
}

func (s *SaveVirtualDatasourceDdlRequest) SetDdl(v string) *SaveVirtualDatasourceDdlRequest {
	s.Ddl = &v
	return s
}

func (s *SaveVirtualDatasourceDdlRequest) SetVdbId(v string) *SaveVirtualDatasourceDdlRequest {
	s.VdbId = &v
	return s
}

func (s *SaveVirtualDatasourceDdlRequest) SetWorkspaceId(v string) *SaveVirtualDatasourceDdlRequest {
	s.WorkspaceId = &v
	return s
}

type SaveVirtualDatasourceDdlResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 72ABCA5B-1E93-55D3-8A61-6D8A03CC5C8B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveVirtualDatasourceDdlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveVirtualDatasourceDdlResponseBody) GoString() string {
	return s.String()
}

func (s *SaveVirtualDatasourceDdlResponseBody) SetCode(v string) *SaveVirtualDatasourceDdlResponseBody {
	s.Code = &v
	return s
}

func (s *SaveVirtualDatasourceDdlResponseBody) SetData(v interface{}) *SaveVirtualDatasourceDdlResponseBody {
	s.Data = v
	return s
}

func (s *SaveVirtualDatasourceDdlResponseBody) SetErrorMsg(v string) *SaveVirtualDatasourceDdlResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SaveVirtualDatasourceDdlResponseBody) SetRequestId(v string) *SaveVirtualDatasourceDdlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveVirtualDatasourceDdlResponseBody) SetSuccess(v bool) *SaveVirtualDatasourceDdlResponseBody {
	s.Success = &v
	return s
}

type SaveVirtualDatasourceDdlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveVirtualDatasourceDdlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveVirtualDatasourceDdlResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveVirtualDatasourceDdlResponse) GoString() string {
	return s.String()
}

func (s *SaveVirtualDatasourceDdlResponse) SetHeaders(v map[string]*string) *SaveVirtualDatasourceDdlResponse {
	s.Headers = v
	return s
}

func (s *SaveVirtualDatasourceDdlResponse) SetStatusCode(v int32) *SaveVirtualDatasourceDdlResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveVirtualDatasourceDdlResponse) SetBody(v *SaveVirtualDatasourceDdlResponseBody) *SaveVirtualDatasourceDdlResponse {
	s.Body = v
	return s
}

type SyncRemoteTablesRequest struct {
	KeepTableNames []*string `json:"keepTableNames,omitempty" xml:"keepTableNames,omitempty" type:"Repeated"`
	// example:
	//
	// true
	PullSamples *bool `json:"pullSamples,omitempty" xml:"pullSamples,omitempty"`
	// This parameter is required.
	TableNames []*string `json:"tableNames,omitempty" xml:"tableNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SyncRemoteTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncRemoteTablesRequest) GoString() string {
	return s.String()
}

func (s *SyncRemoteTablesRequest) SetKeepTableNames(v []*string) *SyncRemoteTablesRequest {
	s.KeepTableNames = v
	return s
}

func (s *SyncRemoteTablesRequest) SetPullSamples(v bool) *SyncRemoteTablesRequest {
	s.PullSamples = &v
	return s
}

func (s *SyncRemoteTablesRequest) SetTableNames(v []*string) *SyncRemoteTablesRequest {
	s.TableNames = v
	return s
}

func (s *SyncRemoteTablesRequest) SetWorkspaceId(v string) *SyncRemoteTablesRequest {
	s.WorkspaceId = &v
	return s
}

type SyncRemoteTablesResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// E9563C85-5810-5835-B68C-78580BC3169E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SyncRemoteTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncRemoteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *SyncRemoteTablesResponseBody) SetCode(v string) *SyncRemoteTablesResponseBody {
	s.Code = &v
	return s
}

func (s *SyncRemoteTablesResponseBody) SetData(v interface{}) *SyncRemoteTablesResponseBody {
	s.Data = v
	return s
}

func (s *SyncRemoteTablesResponseBody) SetErrorMsg(v string) *SyncRemoteTablesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SyncRemoteTablesResponseBody) SetRequestId(v string) *SyncRemoteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncRemoteTablesResponseBody) SetSuccess(v bool) *SyncRemoteTablesResponseBody {
	s.Success = &v
	return s
}

type SyncRemoteTablesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncRemoteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncRemoteTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncRemoteTablesResponse) GoString() string {
	return s.String()
}

func (s *SyncRemoteTablesResponse) SetHeaders(v map[string]*string) *SyncRemoteTablesResponse {
	s.Headers = v
	return s
}

func (s *SyncRemoteTablesResponse) SetStatusCode(v int32) *SyncRemoteTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncRemoteTablesResponse) SetBody(v *SyncRemoteTablesResponseBody) *SyncRemoteTablesResponse {
	s.Body = v
	return s
}

type UpdateVirtualDatasourceInstanceRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Type        *int32  `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vdb-E0F693C8-9F72-5830-B81A-696C9D8EBBD1
	VdbId *string `json:"vdbId,omitempty" xml:"vdbId,omitempty"`
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateVirtualDatasourceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualDatasourceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualDatasourceInstanceRequest) SetDescription(v string) *UpdateVirtualDatasourceInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceRequest) SetName(v string) *UpdateVirtualDatasourceInstanceRequest {
	s.Name = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceRequest) SetType(v int32) *UpdateVirtualDatasourceInstanceRequest {
	s.Type = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceRequest) SetVdbId(v string) *UpdateVirtualDatasourceInstanceRequest {
	s.VdbId = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceRequest) SetWorkspaceId(v string) *UpdateVirtualDatasourceInstanceRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateVirtualDatasourceInstanceResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// NoAuth
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// E9563C85-5810-5835-B68C-78580BC3169E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVirtualDatasourceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualDatasourceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVirtualDatasourceInstanceResponseBody) SetCode(v string) *UpdateVirtualDatasourceInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponseBody) SetData(v interface{}) *UpdateVirtualDatasourceInstanceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponseBody) SetErrorMsg(v string) *UpdateVirtualDatasourceInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponseBody) SetRequestId(v string) *UpdateVirtualDatasourceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponseBody) SetSuccess(v bool) *UpdateVirtualDatasourceInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateVirtualDatasourceInstanceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirtualDatasourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirtualDatasourceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateVirtualDatasourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualDatasourceInstanceResponse) SetHeaders(v map[string]*string) *UpdateVirtualDatasourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponse) SetStatusCode(v int32) *UpdateVirtualDatasourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualDatasourceInstanceResponse) SetBody(v *UpdateVirtualDatasourceInstanceResponseBody) *UpdateVirtualDatasourceInstanceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataanalysisgbi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 取消关联的数据源授权
//
// @param request - CancelDatasourceAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDatasourceAuthorizationResponse
func (client *Client) CancelDatasourceAuthorizationWithOptions(request *CancelDatasourceAuthorizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelDatasourceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelDatasourceAuthorization"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/cancel/datasource"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelDatasourceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关联的数据源授权
//
// @param request - CancelDatasourceAuthorizationRequest
//
// @return CancelDatasourceAuthorizationResponse
func (client *Client) CancelDatasourceAuthorization(request *CancelDatasourceAuthorizationRequest) (_result *CancelDatasourceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelDatasourceAuthorizationResponse{}
	_body, _err := client.CancelDatasourceAuthorizationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数据库关联授权
//
// @param request - CreateDatasourceAuthorizationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasourceAuthorizationResponse
func (client *Client) CreateDatasourceAuthorizationWithOptions(request *CreateDatasourceAuthorizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatasourceAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["userName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.VdbId)) {
		body["vdbId"] = request.VdbId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatasourceAuthorization"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/create/datasource"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasourceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据库关联授权
//
// @param request - CreateDatasourceAuthorizationRequest
//
// @return CreateDatasourceAuthorizationResponse
func (client *Client) CreateDatasourceAuthorization(request *CreateDatasourceAuthorizationRequest) (_result *CreateDatasourceAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasourceAuthorizationResponse{}
	_body, _err := client.CreateDatasourceAuthorizationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 在指定的业务空间创建虚拟数据源
//
// @param request - CreateVirtualDatasourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualDatasourceInstanceResponse
func (client *Client) CreateVirtualDatasourceInstanceWithOptions(request *CreateVirtualDatasourceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateVirtualDatasourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
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
		Action:      tea.String("CreateVirtualDatasourceInstance"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/virtualDatasource/createVirtualDatasourceInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVirtualDatasourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 在指定的业务空间创建虚拟数据源
//
// @param request - CreateVirtualDatasourceInstanceRequest
//
// @return CreateVirtualDatasourceInstanceResponse
func (client *Client) CreateVirtualDatasourceInstance(request *CreateVirtualDatasourceInstanceRequest) (_result *CreateVirtualDatasourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVirtualDatasourceInstanceResponse{}
	_body, _err := client.CreateVirtualDatasourceInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定业务空间下面的虚拟数据源实例
//
// @param request - DeleteVirtualDatasourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualDatasourceInstanceResponse
func (client *Client) DeleteVirtualDatasourceInstanceWithOptions(request *DeleteVirtualDatasourceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVirtualDatasourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VdbId)) {
		body["vdbId"] = request.VdbId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVirtualDatasourceInstance"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/virtualDatasource/deleteVirtualDatasourceInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVirtualDatasourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定业务空间下面的虚拟数据源实例
//
// @param request - DeleteVirtualDatasourceInstanceRequest
//
// @return DeleteVirtualDatasourceInstanceResponse
func (client *Client) DeleteVirtualDatasourceInstance(request *DeleteVirtualDatasourceInstanceRequest) (_result *DeleteVirtualDatasourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVirtualDatasourceInstanceResponse{}
	_body, _err := client.DeleteVirtualDatasourceInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前业务空间下的数据源实例列表
//
// @param request - ListVirtualDatasourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirtualDatasourceInstanceResponse
func (client *Client) ListVirtualDatasourceInstanceWithOptions(request *ListVirtualDatasourceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVirtualDatasourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVirtualDatasourceInstance"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/virtualDatasource/listVirtualDatasourceInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVirtualDatasourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取当前业务空间下的数据源实例列表
//
// @param request - ListVirtualDatasourceInstanceRequest
//
// @return ListVirtualDatasourceInstanceResponse
func (client *Client) ListVirtualDatasourceInstance(request *ListVirtualDatasourceInstanceRequest) (_result *ListVirtualDatasourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVirtualDatasourceInstanceResponse{}
	_body, _err := client.ListVirtualDatasourceInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行数据分析
//
// @param request - RunDataAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDataAnalysisResponse
func (client *Client) RunDataAnalysisWithOptions(workspaceId *string, request *RunDataAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunDataAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateSqlOnly)) {
		body["generateSqlOnly"] = request.GenerateSqlOnly
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationType)) {
		body["specificationType"] = request.SpecificationType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunDataAnalysis"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/gbi/runDataAnalysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunDataAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行数据分析
//
// @param request - RunDataAnalysisRequest
//
// @return RunDataAnalysisResponse
func (client *Client) RunDataAnalysis(workspaceId *string, request *RunDataAnalysisRequest) (_result *RunDataAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunDataAnalysisResponse{}
	_body, _err := client.RunDataAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向当前指定的业务空间下的指定虚拟数据源实例添加ddl语句
//
// @param request - SaveVirtualDatasourceDdlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveVirtualDatasourceDdlResponse
func (client *Client) SaveVirtualDatasourceDdlWithOptions(request *SaveVirtualDatasourceDdlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveVirtualDatasourceDdlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ddl)) {
		body["ddl"] = request.Ddl
	}

	if !tea.BoolValue(util.IsUnset(request.VdbId)) {
		body["vdbId"] = request.VdbId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveVirtualDatasourceDdl"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/virtualDatasource/addDdl2VirtualInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveVirtualDatasourceDdlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向当前指定的业务空间下的指定虚拟数据源实例添加ddl语句
//
// @param request - SaveVirtualDatasourceDdlRequest
//
// @return SaveVirtualDatasourceDdlResponse
func (client *Client) SaveVirtualDatasourceDdl(request *SaveVirtualDatasourceDdlRequest) (_result *SaveVirtualDatasourceDdlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveVirtualDatasourceDdlResponse{}
	_body, _err := client.SaveVirtualDatasourceDdlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新当前业务空间所关联的数据表
//
// @param request - SyncRemoteTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncRemoteTablesResponse
func (client *Client) SyncRemoteTablesWithOptions(request *SyncRemoteTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SyncRemoteTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeepTableNames)) {
		body["keepTableNames"] = request.KeepTableNames
	}

	if !tea.BoolValue(util.IsUnset(request.PullSamples)) {
		body["pullSamples"] = request.PullSamples
	}

	if !tea.BoolValue(util.IsUnset(request.TableNames)) {
		body["tableNames"] = request.TableNames
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SyncRemoteTables"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/datasource/tables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncRemoteTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新当前业务空间所关联的数据表
//
// @param request - SyncRemoteTablesRequest
//
// @return SyncRemoteTablesResponse
func (client *Client) SyncRemoteTables(request *SyncRemoteTablesRequest) (_result *SyncRemoteTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncRemoteTablesResponse{}
	_body, _err := client.SyncRemoteTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改指定业务空间下所指定的虚拟数据源的信息
//
// @param request - UpdateVirtualDatasourceInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirtualDatasourceInstanceResponse
func (client *Client) UpdateVirtualDatasourceInstanceWithOptions(request *UpdateVirtualDatasourceInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateVirtualDatasourceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VdbId)) {
		body["vdbId"] = request.VdbId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateVirtualDatasourceInstance"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/virtualDatasource/updateVirtualDatasourceInstance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateVirtualDatasourceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改指定业务空间下所指定的虚拟数据源的信息
//
// @param request - UpdateVirtualDatasourceInstanceRequest
//
// @return UpdateVirtualDatasourceInstanceResponse
func (client *Client) UpdateVirtualDatasourceInstance(request *UpdateVirtualDatasourceInstanceRequest) (_result *UpdateVirtualDatasourceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVirtualDatasourceInstanceResponse{}
	_body, _err := client.UpdateVirtualDatasourceInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
