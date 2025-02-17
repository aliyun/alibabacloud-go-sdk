// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchDeleteSynonymsRequest struct {
	// This parameter is required.
	SynonymIdKeys []*string `json:"synonymIdKeys,omitempty" xml:"synonymIdKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s BatchDeleteSynonymsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteSynonymsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteSynonymsRequest) SetSynonymIdKeys(v []*string) *BatchDeleteSynonymsRequest {
	s.SynonymIdKeys = v
	return s
}

func (s *BatchDeleteSynonymsRequest) SetWorkspaceId(v string) *BatchDeleteSynonymsRequest {
	s.WorkspaceId = &v
	return s
}

type BatchDeleteSynonymsResponseBody struct {
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
	// 45390C6D-016D-5030-BF65-031ED1F65003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s BatchDeleteSynonymsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteSynonymsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteSynonymsResponseBody) SetCode(v string) *BatchDeleteSynonymsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteSynonymsResponseBody) SetData(v interface{}) *BatchDeleteSynonymsResponseBody {
	s.Data = v
	return s
}

func (s *BatchDeleteSynonymsResponseBody) SetErrorMsg(v string) *BatchDeleteSynonymsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BatchDeleteSynonymsResponseBody) SetRequestId(v string) *BatchDeleteSynonymsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteSynonymsResponseBody) SetSuccess(v bool) *BatchDeleteSynonymsResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteSynonymsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteSynonymsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteSynonymsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteSynonymsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteSynonymsResponse) SetHeaders(v map[string]*string) *BatchDeleteSynonymsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteSynonymsResponse) SetStatusCode(v int32) *BatchDeleteSynonymsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteSynonymsResponse) SetBody(v *BatchDeleteSynonymsResponseBody) *BatchDeleteSynonymsResponse {
	s.Body = v
	return s
}

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

type CreateBusinessLogicRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateBusinessLogicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessLogicRequest) GoString() string {
	return s.String()
}

func (s *CreateBusinessLogicRequest) SetDescription(v string) *CreateBusinessLogicRequest {
	s.Description = &v
	return s
}

func (s *CreateBusinessLogicRequest) SetType(v int32) *CreateBusinessLogicRequest {
	s.Type = &v
	return s
}

func (s *CreateBusinessLogicRequest) SetWorkspaceId(v string) *CreateBusinessLogicRequest {
	s.WorkspaceId = &v
	return s
}

type CreateBusinessLogicResponseBody struct {
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

func (s CreateBusinessLogicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessLogicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBusinessLogicResponseBody) SetCode(v string) *CreateBusinessLogicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBusinessLogicResponseBody) SetData(v interface{}) *CreateBusinessLogicResponseBody {
	s.Data = v
	return s
}

func (s *CreateBusinessLogicResponseBody) SetErrorMsg(v string) *CreateBusinessLogicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateBusinessLogicResponseBody) SetRequestId(v string) *CreateBusinessLogicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBusinessLogicResponseBody) SetSuccess(v bool) *CreateBusinessLogicResponseBody {
	s.Success = &v
	return s
}

type CreateBusinessLogicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBusinessLogicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBusinessLogicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBusinessLogicResponse) GoString() string {
	return s.String()
}

func (s *CreateBusinessLogicResponse) SetHeaders(v map[string]*string) *CreateBusinessLogicResponse {
	s.Headers = v
	return s
}

func (s *CreateBusinessLogicResponse) SetStatusCode(v int32) *CreateBusinessLogicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBusinessLogicResponse) SetBody(v *CreateBusinessLogicResponseBody) *CreateBusinessLogicResponse {
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

type CreateSynonymsRequest struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// This parameter is required.
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
	// This parameter is required.
	WordSynonyms []*string `json:"wordSynonyms,omitempty" xml:"wordSynonyms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateSynonymsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSynonymsRequest) GoString() string {
	return s.String()
}

func (s *CreateSynonymsRequest) SetColumns(v []*string) *CreateSynonymsRequest {
	s.Columns = v
	return s
}

func (s *CreateSynonymsRequest) SetWord(v string) *CreateSynonymsRequest {
	s.Word = &v
	return s
}

func (s *CreateSynonymsRequest) SetWordSynonyms(v []*string) *CreateSynonymsRequest {
	s.WordSynonyms = v
	return s
}

func (s *CreateSynonymsRequest) SetWorkspaceId(v string) *CreateSynonymsRequest {
	s.WorkspaceId = &v
	return s
}

type CreateSynonymsResponseBody struct {
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

func (s CreateSynonymsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSynonymsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynonymsResponseBody) SetCode(v string) *CreateSynonymsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSynonymsResponseBody) SetData(v interface{}) *CreateSynonymsResponseBody {
	s.Data = v
	return s
}

func (s *CreateSynonymsResponseBody) SetErrorMsg(v string) *CreateSynonymsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateSynonymsResponseBody) SetRequestId(v string) *CreateSynonymsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSynonymsResponseBody) SetSuccess(v bool) *CreateSynonymsResponseBody {
	s.Success = &v
	return s
}

type CreateSynonymsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSynonymsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSynonymsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSynonymsResponse) GoString() string {
	return s.String()
}

func (s *CreateSynonymsResponse) SetHeaders(v map[string]*string) *CreateSynonymsResponse {
	s.Headers = v
	return s
}

func (s *CreateSynonymsResponse) SetStatusCode(v int32) *CreateSynonymsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSynonymsResponse) SetBody(v *CreateSynonymsResponseBody) *CreateSynonymsResponse {
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

type DeleteBusinessLogicRequest struct {
	// This parameter is required.
	BusinessLogicIdKeys []*string `json:"businessLogicIdKeys,omitempty" xml:"businessLogicIdKeys,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteBusinessLogicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessLogicRequest) GoString() string {
	return s.String()
}

func (s *DeleteBusinessLogicRequest) SetBusinessLogicIdKeys(v []*string) *DeleteBusinessLogicRequest {
	s.BusinessLogicIdKeys = v
	return s
}

func (s *DeleteBusinessLogicRequest) SetWorkspaceId(v string) *DeleteBusinessLogicRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteBusinessLogicResponseBody struct {
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

func (s DeleteBusinessLogicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessLogicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBusinessLogicResponseBody) SetCode(v string) *DeleteBusinessLogicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBusinessLogicResponseBody) SetData(v interface{}) *DeleteBusinessLogicResponseBody {
	s.Data = v
	return s
}

func (s *DeleteBusinessLogicResponseBody) SetErrorMsg(v string) *DeleteBusinessLogicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteBusinessLogicResponseBody) SetRequestId(v string) *DeleteBusinessLogicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBusinessLogicResponseBody) SetSuccess(v bool) *DeleteBusinessLogicResponseBody {
	s.Success = &v
	return s
}

type DeleteBusinessLogicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBusinessLogicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBusinessLogicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBusinessLogicResponse) GoString() string {
	return s.String()
}

func (s *DeleteBusinessLogicResponse) SetHeaders(v map[string]*string) *DeleteBusinessLogicResponse {
	s.Headers = v
	return s
}

func (s *DeleteBusinessLogicResponse) SetStatusCode(v int32) *DeleteBusinessLogicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBusinessLogicResponse) SetBody(v *DeleteBusinessLogicResponseBody) *DeleteBusinessLogicResponse {
	s.Body = v
	return s
}

type DeleteColumnRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// column-AAAAAAAAh6Q9ERazKYFvkA
	ColumnIdKey *string `json:"columnIdKey,omitempty" xml:"columnIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnRequest) GoString() string {
	return s.String()
}

func (s *DeleteColumnRequest) SetColumnIdKey(v string) *DeleteColumnRequest {
	s.ColumnIdKey = &v
	return s
}

func (s *DeleteColumnRequest) SetWorkspaceId(v string) *DeleteColumnRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteColumnResponseBody struct {
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

func (s DeleteColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteColumnResponseBody) SetCode(v string) *DeleteColumnResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteColumnResponseBody) SetData(v interface{}) *DeleteColumnResponseBody {
	s.Data = v
	return s
}

func (s *DeleteColumnResponseBody) SetErrorMsg(v string) *DeleteColumnResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteColumnResponseBody) SetRequestId(v string) *DeleteColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteColumnResponseBody) SetSuccess(v bool) *DeleteColumnResponseBody {
	s.Success = &v
	return s
}

type DeleteColumnResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteColumnResponse) GoString() string {
	return s.String()
}

func (s *DeleteColumnResponse) SetHeaders(v map[string]*string) *DeleteColumnResponse {
	s.Headers = v
	return s
}

func (s *DeleteColumnResponse) SetStatusCode(v int32) *DeleteColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteColumnResponse) SetBody(v *DeleteColumnResponseBody) *DeleteColumnResponse {
	s.Body = v
	return s
}

type DeleteSelectedTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s DeleteSelectedTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSelectedTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteSelectedTableRequest) SetTableIdKey(v string) *DeleteSelectedTableRequest {
	s.TableIdKey = &v
	return s
}

func (s *DeleteSelectedTableRequest) SetWorkspaceId(v string) *DeleteSelectedTableRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteSelectedTableResponseBody struct {
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
	// 45390C6D-016D-5030-BF65-031ED1F65003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteSelectedTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSelectedTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSelectedTableResponseBody) SetCode(v string) *DeleteSelectedTableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSelectedTableResponseBody) SetData(v interface{}) *DeleteSelectedTableResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSelectedTableResponseBody) SetErrorMsg(v string) *DeleteSelectedTableResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteSelectedTableResponseBody) SetRequestId(v string) *DeleteSelectedTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSelectedTableResponseBody) SetSuccess(v bool) *DeleteSelectedTableResponseBody {
	s.Success = &v
	return s
}

type DeleteSelectedTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSelectedTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSelectedTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSelectedTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteSelectedTableResponse) SetHeaders(v map[string]*string) *DeleteSelectedTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteSelectedTableResponse) SetStatusCode(v int32) *DeleteSelectedTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSelectedTableResponse) SetBody(v *DeleteSelectedTableResponseBody) *DeleteSelectedTableResponse {
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

type ListBusinessLogicRequest struct {
	// example:
	//
	// 10
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

func (s ListBusinessLogicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessLogicRequest) GoString() string {
	return s.String()
}

func (s *ListBusinessLogicRequest) SetMaxResults(v int32) *ListBusinessLogicRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBusinessLogicRequest) SetNextToken(v string) *ListBusinessLogicRequest {
	s.NextToken = &v
	return s
}

func (s *ListBusinessLogicRequest) SetWorkspaceId(v string) *ListBusinessLogicRequest {
	s.WorkspaceId = &v
	return s
}

type ListBusinessLogicResponseBody struct {
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
	// FB11F719-9AC8-5C99-AB0A-4709D437FCFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListBusinessLogicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessLogicResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessLogicResponseBody) SetCode(v string) *ListBusinessLogicResponseBody {
	s.Code = &v
	return s
}

func (s *ListBusinessLogicResponseBody) SetData(v interface{}) *ListBusinessLogicResponseBody {
	s.Data = v
	return s
}

func (s *ListBusinessLogicResponseBody) SetErrorMsg(v string) *ListBusinessLogicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListBusinessLogicResponseBody) SetRequestId(v string) *ListBusinessLogicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessLogicResponseBody) SetSuccess(v bool) *ListBusinessLogicResponseBody {
	s.Success = &v
	return s
}

type ListBusinessLogicResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBusinessLogicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBusinessLogicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBusinessLogicResponse) GoString() string {
	return s.String()
}

func (s *ListBusinessLogicResponse) SetHeaders(v map[string]*string) *ListBusinessLogicResponse {
	s.Headers = v
	return s
}

func (s *ListBusinessLogicResponse) SetStatusCode(v int32) *ListBusinessLogicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBusinessLogicResponse) SetBody(v *ListBusinessLogicResponseBody) *ListBusinessLogicResponse {
	s.Body = v
	return s
}

type ListColumnRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// e8Z0nRyY51ZQmYljqGNK
	NextToken *int32 `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// This parameter is required.
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s ListColumnRequest) GoString() string {
	return s.String()
}

func (s *ListColumnRequest) SetMaxResults(v int32) *ListColumnRequest {
	s.MaxResults = &v
	return s
}

func (s *ListColumnRequest) SetNextToken(v int32) *ListColumnRequest {
	s.NextToken = &v
	return s
}

func (s *ListColumnRequest) SetTableIdKey(v string) *ListColumnRequest {
	s.TableIdKey = &v
	return s
}

func (s *ListColumnRequest) SetWorkspaceId(v string) *ListColumnRequest {
	s.WorkspaceId = &v
	return s
}

type ListColumnResponseBody struct {
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
	// FB11F719-9AC8-5C99-AB0A-4709D437FCFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListColumnResponseBody) GoString() string {
	return s.String()
}

func (s *ListColumnResponseBody) SetCode(v string) *ListColumnResponseBody {
	s.Code = &v
	return s
}

func (s *ListColumnResponseBody) SetData(v interface{}) *ListColumnResponseBody {
	s.Data = v
	return s
}

func (s *ListColumnResponseBody) SetErrorMsg(v string) *ListColumnResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListColumnResponseBody) SetRequestId(v string) *ListColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListColumnResponseBody) SetSuccess(v bool) *ListColumnResponseBody {
	s.Success = &v
	return s
}

type ListColumnResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s ListColumnResponse) GoString() string {
	return s.String()
}

func (s *ListColumnResponse) SetHeaders(v map[string]*string) *ListColumnResponse {
	s.Headers = v
	return s
}

func (s *ListColumnResponse) SetStatusCode(v int32) *ListColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *ListColumnResponse) SetBody(v *ListColumnResponseBody) *ListColumnResponse {
	s.Body = v
	return s
}

type ListEnumMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// column-AAAAAAAAh6cWOUPagYstkg
	ColumnIdKey *string `json:"columnIdKey,omitempty" xml:"columnIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListEnumMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnumMappingRequest) GoString() string {
	return s.String()
}

func (s *ListEnumMappingRequest) SetColumnIdKey(v string) *ListEnumMappingRequest {
	s.ColumnIdKey = &v
	return s
}

func (s *ListEnumMappingRequest) SetTableIdKey(v string) *ListEnumMappingRequest {
	s.TableIdKey = &v
	return s
}

func (s *ListEnumMappingRequest) SetWorkspaceId(v string) *ListEnumMappingRequest {
	s.WorkspaceId = &v
	return s
}

type ListEnumMappingResponseBody struct {
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
	// FB11F719-9AC8-5C99-AB0A-4709D437FCFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListEnumMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnumMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnumMappingResponseBody) SetCode(v string) *ListEnumMappingResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnumMappingResponseBody) SetData(v interface{}) *ListEnumMappingResponseBody {
	s.Data = v
	return s
}

func (s *ListEnumMappingResponseBody) SetErrorMsg(v string) *ListEnumMappingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListEnumMappingResponseBody) SetRequestId(v string) *ListEnumMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnumMappingResponseBody) SetSuccess(v bool) *ListEnumMappingResponseBody {
	s.Success = &v
	return s
}

type ListEnumMappingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnumMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnumMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnumMappingResponse) GoString() string {
	return s.String()
}

func (s *ListEnumMappingResponse) SetHeaders(v map[string]*string) *ListEnumMappingResponse {
	s.Headers = v
	return s
}

func (s *ListEnumMappingResponse) SetStatusCode(v int32) *ListEnumMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnumMappingResponse) SetBody(v *ListEnumMappingResponseBody) *ListEnumMappingResponse {
	s.Body = v
	return s
}

type ListSelectedTablesRequest struct {
	// example:
	//
	// 10
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

func (s ListSelectedTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSelectedTablesRequest) GoString() string {
	return s.String()
}

func (s *ListSelectedTablesRequest) SetMaxResults(v int32) *ListSelectedTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSelectedTablesRequest) SetNextToken(v string) *ListSelectedTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSelectedTablesRequest) SetWorkspaceId(v string) *ListSelectedTablesRequest {
	s.WorkspaceId = &v
	return s
}

type ListSelectedTablesResponseBody struct {
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
	// FB11F719-9AC8-5C99-AB0A-4709D437FCFC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListSelectedTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSelectedTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSelectedTablesResponseBody) SetCode(v string) *ListSelectedTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSelectedTablesResponseBody) SetData(v interface{}) *ListSelectedTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListSelectedTablesResponseBody) SetErrorMsg(v string) *ListSelectedTablesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSelectedTablesResponseBody) SetRequestId(v string) *ListSelectedTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSelectedTablesResponseBody) SetSuccess(v bool) *ListSelectedTablesResponseBody {
	s.Success = &v
	return s
}

type ListSelectedTablesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSelectedTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectedTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSelectedTablesResponse) GoString() string {
	return s.String()
}

func (s *ListSelectedTablesResponse) SetHeaders(v map[string]*string) *ListSelectedTablesResponse {
	s.Headers = v
	return s
}

func (s *ListSelectedTablesResponse) SetStatusCode(v int32) *ListSelectedTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectedTablesResponse) SetBody(v *ListSelectedTablesResponseBody) *ListSelectedTablesResponse {
	s.Body = v
	return s
}

type ListSynonymsRequest struct {
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
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

func (s ListSynonymsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSynonymsRequest) GoString() string {
	return s.String()
}

func (s *ListSynonymsRequest) SetMaxResults(v int64) *ListSynonymsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSynonymsRequest) SetNextToken(v string) *ListSynonymsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSynonymsRequest) SetWorkspaceId(v string) *ListSynonymsRequest {
	s.WorkspaceId = &v
	return s
}

type ListSynonymsResponseBody struct {
	// example:
	//
	// NoAuth
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {
	//
	//   "data": {
	//
	//     "data": [
	//
	//       {
	//
	//         "gmtModified": 1734401404000,
	//
	//         "columns": [
	//
	//           "test.id",
	//
	//           "user_info.createdt"
	//
	//         ],
	//
	//         "synonymIdKey": "synonyms-AAAAAAAAAVLaD8z63NnFhA",
	//
	//         "wordSynonyms": [
	//
	//           "1"
	//
	//         ],
	//
	//         "workSpaceId": "10024809",
	//
	//         "gmtCreate": 1734401404000,
	//
	//         "word": "1",
	//
	//         "status": 1
	//
	//       }
	//
	//     ],
	//
	//     "nextToken": "k1BLjEN114wyfrhDHoJlbg==",
	//
	//     "totalCount": 0
	//
	//   }
	//
	// }
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

func (s ListSynonymsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSynonymsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSynonymsResponseBody) SetCode(v string) *ListSynonymsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSynonymsResponseBody) SetData(v interface{}) *ListSynonymsResponseBody {
	s.Data = v
	return s
}

func (s *ListSynonymsResponseBody) SetErrorMsg(v string) *ListSynonymsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSynonymsResponseBody) SetRequestId(v string) *ListSynonymsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSynonymsResponseBody) SetSuccess(v bool) *ListSynonymsResponseBody {
	s.Success = &v
	return s
}

type ListSynonymsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSynonymsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSynonymsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSynonymsResponse) GoString() string {
	return s.String()
}

func (s *ListSynonymsResponse) SetHeaders(v map[string]*string) *ListSynonymsResponse {
	s.Headers = v
	return s
}

func (s *ListSynonymsResponse) SetStatusCode(v int32) *ListSynonymsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSynonymsResponse) SetBody(v *ListSynonymsResponseBody) *ListSynonymsResponse {
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

type RecoverColumnRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// column-AAAAAAAAh6cWOUPagYstkg
	ColumnIdKey *string `json:"columnIdKey,omitempty" xml:"columnIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RecoverColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverColumnRequest) GoString() string {
	return s.String()
}

func (s *RecoverColumnRequest) SetColumnIdKey(v string) *RecoverColumnRequest {
	s.ColumnIdKey = &v
	return s
}

func (s *RecoverColumnRequest) SetTableIdKey(v string) *RecoverColumnRequest {
	s.TableIdKey = &v
	return s
}

func (s *RecoverColumnRequest) SetWorkspaceId(v string) *RecoverColumnRequest {
	s.WorkspaceId = &v
	return s
}

type RecoverColumnResponseBody struct {
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

func (s RecoverColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverColumnResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverColumnResponseBody) SetCode(v string) *RecoverColumnResponseBody {
	s.Code = &v
	return s
}

func (s *RecoverColumnResponseBody) SetData(v interface{}) *RecoverColumnResponseBody {
	s.Data = v
	return s
}

func (s *RecoverColumnResponseBody) SetErrorMsg(v string) *RecoverColumnResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RecoverColumnResponseBody) SetRequestId(v string) *RecoverColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverColumnResponseBody) SetSuccess(v bool) *RecoverColumnResponseBody {
	s.Success = &v
	return s
}

type RecoverColumnResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverColumnResponse) GoString() string {
	return s.String()
}

func (s *RecoverColumnResponse) SetHeaders(v map[string]*string) *RecoverColumnResponse {
	s.Headers = v
	return s
}

func (s *RecoverColumnResponse) SetStatusCode(v int32) *RecoverColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverColumnResponse) SetBody(v *RecoverColumnResponseBody) *RecoverColumnResponse {
	s.Body = v
	return s
}

type ResyncTableRequest struct {
	Keep *bool `json:"keep,omitempty" xml:"keep,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ResyncTableRequest) String() string {
	return tea.Prettify(s)
}

func (s ResyncTableRequest) GoString() string {
	return s.String()
}

func (s *ResyncTableRequest) SetKeep(v bool) *ResyncTableRequest {
	s.Keep = &v
	return s
}

func (s *ResyncTableRequest) SetTableIdKey(v string) *ResyncTableRequest {
	s.TableIdKey = &v
	return s
}

func (s *ResyncTableRequest) SetWorkspaceId(v string) *ResyncTableRequest {
	s.WorkspaceId = &v
	return s
}

type ResyncTableResponseBody struct {
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

func (s ResyncTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResyncTableResponseBody) GoString() string {
	return s.String()
}

func (s *ResyncTableResponseBody) SetCode(v string) *ResyncTableResponseBody {
	s.Code = &v
	return s
}

func (s *ResyncTableResponseBody) SetData(v interface{}) *ResyncTableResponseBody {
	s.Data = v
	return s
}

func (s *ResyncTableResponseBody) SetErrorMsg(v string) *ResyncTableResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ResyncTableResponseBody) SetRequestId(v string) *ResyncTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResyncTableResponseBody) SetSuccess(v bool) *ResyncTableResponseBody {
	s.Success = &v
	return s
}

type ResyncTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResyncTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResyncTableResponse) String() string {
	return tea.Prettify(s)
}

func (s ResyncTableResponse) GoString() string {
	return s.String()
}

func (s *ResyncTableResponse) SetHeaders(v map[string]*string) *ResyncTableResponse {
	s.Headers = v
	return s
}

func (s *ResyncTableResponse) SetStatusCode(v int32) *ResyncTableResponse {
	s.StatusCode = &v
	return s
}

func (s *ResyncTableResponse) SetBody(v *ResyncTableResponseBody) *ResyncTableResponse {
	s.Body = v
	return s
}

type RunDataAnalysisRequest struct {
	DataRole []*string `json:"dataRole,omitempty" xml:"dataRole,omitempty" type:"Repeated"`
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
	SpecificationType *string     `json:"specificationType,omitempty" xml:"specificationType,omitempty"`
	UserParams        interface{} `json:"userParams,omitempty" xml:"userParams,omitempty"`
}

func (s RunDataAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisRequest) SetDataRole(v []*string) *RunDataAnalysisRequest {
	s.DataRole = v
	return s
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

func (s *RunDataAnalysisRequest) SetUserParams(v interface{}) *RunDataAnalysisRequest {
	s.UserParams = v
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

type RunDataResultAnalysisRequest struct {
	// example:
	//
	// all
	AnalysisMode *string `json:"analysisMode,omitempty" xml:"analysisMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FF76AD3F-8B32-567E-819B-0D3738917006
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SqlData   *RunDataResultAnalysisRequestSqlData `json:"sqlData,omitempty" xml:"sqlData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RunDataResultAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisRequest) SetAnalysisMode(v string) *RunDataResultAnalysisRequest {
	s.AnalysisMode = &v
	return s
}

func (s *RunDataResultAnalysisRequest) SetRequestId(v string) *RunDataResultAnalysisRequest {
	s.RequestId = &v
	return s
}

func (s *RunDataResultAnalysisRequest) SetSqlData(v *RunDataResultAnalysisRequestSqlData) *RunDataResultAnalysisRequest {
	s.SqlData = v
	return s
}

func (s *RunDataResultAnalysisRequest) SetWorkspaceId(v string) *RunDataResultAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

type RunDataResultAnalysisRequestSqlData struct {
	Column []*string            `json:"column,omitempty" xml:"column,omitempty" type:"Repeated"`
	Data   []map[string]*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s RunDataResultAnalysisRequestSqlData) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisRequestSqlData) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisRequestSqlData) SetColumn(v []*string) *RunDataResultAnalysisRequestSqlData {
	s.Column = v
	return s
}

func (s *RunDataResultAnalysisRequestSqlData) SetData(v []map[string]*string) *RunDataResultAnalysisRequestSqlData {
	s.Data = v
	return s
}

type RunDataResultAnalysisResponseBody struct {
	Data *RunDataResultAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s RunDataResultAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisResponseBody) SetData(v *RunDataResultAnalysisResponseBodyData) *RunDataResultAnalysisResponseBody {
	s.Data = v
	return s
}

type RunDataResultAnalysisResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/tech-scp-chain7.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// DA2578F7-88A5-5D6E-9305-33E724E97D60
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rewrite   *string `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
	// example:
	//
	// select p.product_id, p.product_name, sum(o.quantity) as total_sales from products p join orders o on p.product_id = o.product_id where o.order_date between \\"2022-10-22\\" and \\"2024-10-22\\" group by p.product_id, p.product_name having total_sales > 5
	Sql           *string                                             `json:"sql,omitempty" xml:"sql,omitempty"`
	Visualization *RunDataResultAnalysisResponseBodyDataVisualization `json:"visualization,omitempty" xml:"visualization,omitempty" type:"Struct"`
}

func (s RunDataResultAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisResponseBodyData) SetErrorMessage(v string) *RunDataResultAnalysisResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyData) SetEvent(v string) *RunDataResultAnalysisResponseBodyData {
	s.Event = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyData) SetRequestId(v string) *RunDataResultAnalysisResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyData) SetRewrite(v string) *RunDataResultAnalysisResponseBodyData {
	s.Rewrite = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyData) SetSql(v string) *RunDataResultAnalysisResponseBodyData {
	s.Sql = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyData) SetVisualization(v *RunDataResultAnalysisResponseBodyDataVisualization) *RunDataResultAnalysisResponseBodyData {
	s.Visualization = v
	return s
}

type RunDataResultAnalysisResponseBodyDataVisualization struct {
	Data *RunDataResultAnalysisResponseBodyDataVisualizationData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Text *string                                                 `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunDataResultAnalysisResponseBodyDataVisualization) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisResponseBodyDataVisualization) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisResponseBodyDataVisualization) SetData(v *RunDataResultAnalysisResponseBodyDataVisualizationData) *RunDataResultAnalysisResponseBodyDataVisualization {
	s.Data = v
	return s
}

func (s *RunDataResultAnalysisResponseBodyDataVisualization) SetText(v string) *RunDataResultAnalysisResponseBodyDataVisualization {
	s.Text = &v
	return s
}

type RunDataResultAnalysisResponseBodyDataVisualizationData struct {
	// example:
	//
	// bar
	PlotType *string   `json:"plotType,omitempty" xml:"plotType,omitempty"`
	XAxis    []*string `json:"xAxis,omitempty" xml:"xAxis,omitempty" type:"Repeated"`
	YAxis    []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
}

func (s RunDataResultAnalysisResponseBodyDataVisualizationData) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisResponseBodyDataVisualizationData) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisResponseBodyDataVisualizationData) SetPlotType(v string) *RunDataResultAnalysisResponseBodyDataVisualizationData {
	s.PlotType = &v
	return s
}

func (s *RunDataResultAnalysisResponseBodyDataVisualizationData) SetXAxis(v []*string) *RunDataResultAnalysisResponseBodyDataVisualizationData {
	s.XAxis = v
	return s
}

func (s *RunDataResultAnalysisResponseBodyDataVisualizationData) SetYAxis(v []*string) *RunDataResultAnalysisResponseBodyDataVisualizationData {
	s.YAxis = v
	return s
}

type RunDataResultAnalysisResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDataResultAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDataResultAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDataResultAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunDataResultAnalysisResponse) SetHeaders(v map[string]*string) *RunDataResultAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunDataResultAnalysisResponse) SetStatusCode(v int32) *RunDataResultAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDataResultAnalysisResponse) SetBody(v *RunDataResultAnalysisResponseBody) *RunDataResultAnalysisResponse {
	s.Body = v
	return s
}

type RunSqlGenerationRequest struct {
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// d5eced84-fd25-43ee-a245-adb4e4a8c3be
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// STANDARD_MIX
	SpecificationType *string `json:"specificationType,omitempty" xml:"specificationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s RunSqlGenerationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSqlGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunSqlGenerationRequest) SetQuery(v string) *RunSqlGenerationRequest {
	s.Query = &v
	return s
}

func (s *RunSqlGenerationRequest) SetSessionId(v string) *RunSqlGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunSqlGenerationRequest) SetSpecificationType(v string) *RunSqlGenerationRequest {
	s.SpecificationType = &v
	return s
}

func (s *RunSqlGenerationRequest) SetWorkspaceId(v string) *RunSqlGenerationRequest {
	s.WorkspaceId = &v
	return s
}

type RunSqlGenerationResponseBody struct {
	Data *RunSqlGenerationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s RunSqlGenerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSqlGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunSqlGenerationResponseBody) SetData(v *RunSqlGenerationResponseBodyData) *RunSqlGenerationResponseBody {
	s.Data = v
	return s
}

type RunSqlGenerationResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/tech-scp-chain7.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// sql
	Event    *string `json:"event,omitempty" xml:"event,omitempty"`
	Evidence *string `json:"evidence,omitempty" xml:"evidence,omitempty"`
	// example:
	//
	// DA2578F7-88A5-5D6E-9305-33E724E97D60
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rewrite   *string   `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
	Selector  []*string `json:"selector,omitempty" xml:"selector,omitempty" type:"Repeated"`
	// example:
	//
	// f64c38dd-a235-4bb4-ae6c-79eaedcba699
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// select p.product_id, p.product_name, sum(o.quantity) as total_sales from products p join orders o on p.product_id = o.product_id where o.order_date between \\"2022-10-22\\" and \\"2024-10-22\\" group by p.product_id, p.product_name having total_sales > 5
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
	// example:
	//
	// Can not issue data manipulation statements with executeQuery()
	SqlError *string `json:"sqlError,omitempty" xml:"sqlError,omitempty"`
}

func (s RunSqlGenerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunSqlGenerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSqlGenerationResponseBodyData) SetErrorMessage(v string) *RunSqlGenerationResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetEvent(v string) *RunSqlGenerationResponseBodyData {
	s.Event = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetEvidence(v string) *RunSqlGenerationResponseBodyData {
	s.Evidence = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetRequestId(v string) *RunSqlGenerationResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetRewrite(v string) *RunSqlGenerationResponseBodyData {
	s.Rewrite = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetSelector(v []*string) *RunSqlGenerationResponseBodyData {
	s.Selector = v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetSessionId(v string) *RunSqlGenerationResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetSql(v string) *RunSqlGenerationResponseBodyData {
	s.Sql = &v
	return s
}

func (s *RunSqlGenerationResponseBodyData) SetSqlError(v string) *RunSqlGenerationResponseBodyData {
	s.SqlError = &v
	return s
}

type RunSqlGenerationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSqlGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSqlGenerationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSqlGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunSqlGenerationResponse) SetHeaders(v map[string]*string) *RunSqlGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunSqlGenerationResponse) SetStatusCode(v int32) *RunSqlGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSqlGenerationResponse) SetBody(v *RunSqlGenerationResponseBody) *RunSqlGenerationResponse {
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
	KeepTableNames       []*string `json:"keepTableNames,omitempty" xml:"keepTableNames,omitempty" type:"Repeated"`
	NoModifiedTableNames []*string `json:"noModifiedTableNames,omitempty" xml:"noModifiedTableNames,omitempty" type:"Repeated"`
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

func (s *SyncRemoteTablesRequest) SetNoModifiedTableNames(v []*string) *SyncRemoteTablesRequest {
	s.NoModifiedTableNames = v
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

type UpdateBusinessLogicRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// businessLogic-AAAAAAAAAGAIyJoP7LbKuA
	BusinessLogicIdKey *string `json:"businessLogicIdKey,omitempty" xml:"businessLogicIdKey,omitempty"`
	// This parameter is required.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateBusinessLogicRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBusinessLogicRequest) GoString() string {
	return s.String()
}

func (s *UpdateBusinessLogicRequest) SetBusinessLogicIdKey(v string) *UpdateBusinessLogicRequest {
	s.BusinessLogicIdKey = &v
	return s
}

func (s *UpdateBusinessLogicRequest) SetDescription(v string) *UpdateBusinessLogicRequest {
	s.Description = &v
	return s
}

func (s *UpdateBusinessLogicRequest) SetType(v int64) *UpdateBusinessLogicRequest {
	s.Type = &v
	return s
}

func (s *UpdateBusinessLogicRequest) SetWorkspaceId(v string) *UpdateBusinessLogicRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateBusinessLogicResponseBody struct {
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

func (s UpdateBusinessLogicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBusinessLogicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBusinessLogicResponseBody) SetCode(v string) *UpdateBusinessLogicResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBusinessLogicResponseBody) SetData(v interface{}) *UpdateBusinessLogicResponseBody {
	s.Data = v
	return s
}

func (s *UpdateBusinessLogicResponseBody) SetErrorMsg(v string) *UpdateBusinessLogicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateBusinessLogicResponseBody) SetRequestId(v string) *UpdateBusinessLogicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBusinessLogicResponseBody) SetSuccess(v bool) *UpdateBusinessLogicResponseBody {
	s.Success = &v
	return s
}

type UpdateBusinessLogicResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBusinessLogicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBusinessLogicResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBusinessLogicResponse) GoString() string {
	return s.String()
}

func (s *UpdateBusinessLogicResponse) SetHeaders(v map[string]*string) *UpdateBusinessLogicResponse {
	s.Headers = v
	return s
}

func (s *UpdateBusinessLogicResponse) SetStatusCode(v int32) *UpdateBusinessLogicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBusinessLogicResponse) SetBody(v *UpdateBusinessLogicResponseBody) *UpdateBusinessLogicResponse {
	s.Body = v
	return s
}

type UpdateColumnRequest struct {
	ChineseName *string `json:"chineseName,omitempty" xml:"chineseName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// column-AAAAAAAAh6cWOUPagYstkg
	ColumnIdKey *string `json:"columnIdKey,omitempty" xml:"columnIdKey,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EnumType   *int32    `json:"enumType,omitempty" xml:"enumType,omitempty"`
	EnumValues []*string `json:"enumValues,omitempty" xml:"enumValues,omitempty" type:"Repeated"`
	// example:
	//
	// 2000
	RangeMax *int64 `json:"rangeMax,omitempty" xml:"rangeMax,omitempty"`
	// example:
	//
	// 0
	RangeMin *int64    `json:"rangeMin,omitempty" xml:"rangeMin,omitempty"`
	Samples  []*string `json:"samples,omitempty" xml:"samples,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateColumnRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateColumnRequest) GoString() string {
	return s.String()
}

func (s *UpdateColumnRequest) SetChineseName(v string) *UpdateColumnRequest {
	s.ChineseName = &v
	return s
}

func (s *UpdateColumnRequest) SetColumnIdKey(v string) *UpdateColumnRequest {
	s.ColumnIdKey = &v
	return s
}

func (s *UpdateColumnRequest) SetDescription(v string) *UpdateColumnRequest {
	s.Description = &v
	return s
}

func (s *UpdateColumnRequest) SetEnumType(v int32) *UpdateColumnRequest {
	s.EnumType = &v
	return s
}

func (s *UpdateColumnRequest) SetEnumValues(v []*string) *UpdateColumnRequest {
	s.EnumValues = v
	return s
}

func (s *UpdateColumnRequest) SetRangeMax(v int64) *UpdateColumnRequest {
	s.RangeMax = &v
	return s
}

func (s *UpdateColumnRequest) SetRangeMin(v int64) *UpdateColumnRequest {
	s.RangeMin = &v
	return s
}

func (s *UpdateColumnRequest) SetSamples(v []*string) *UpdateColumnRequest {
	s.Samples = v
	return s
}

func (s *UpdateColumnRequest) SetTableIdKey(v string) *UpdateColumnRequest {
	s.TableIdKey = &v
	return s
}

func (s *UpdateColumnRequest) SetWorkspaceId(v string) *UpdateColumnRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateColumnResponseBody struct {
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
	// 45390C6D-016D-5030-BF65-031ED1F65003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateColumnResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateColumnResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateColumnResponseBody) SetCode(v string) *UpdateColumnResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateColumnResponseBody) SetData(v interface{}) *UpdateColumnResponseBody {
	s.Data = v
	return s
}

func (s *UpdateColumnResponseBody) SetErrorMsg(v string) *UpdateColumnResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateColumnResponseBody) SetRequestId(v string) *UpdateColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateColumnResponseBody) SetSuccess(v bool) *UpdateColumnResponseBody {
	s.Success = &v
	return s
}

type UpdateColumnResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateColumnResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateColumnResponse) GoString() string {
	return s.String()
}

func (s *UpdateColumnResponse) SetHeaders(v map[string]*string) *UpdateColumnResponse {
	s.Headers = v
	return s
}

func (s *UpdateColumnResponse) SetStatusCode(v int32) *UpdateColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateColumnResponse) SetBody(v *UpdateColumnResponseBody) *UpdateColumnResponse {
	s.Body = v
	return s
}

type UpdateEnumMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// column-AAAAAAAAh6cWOUPagYstkg
	ColumnIdKey *string              `json:"columnIdKey,omitempty" xml:"columnIdKey,omitempty"`
	EnumMapping map[string][]*string `json:"enumMapping,omitempty" xml:"enumMapping,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateEnumMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnumMappingRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnumMappingRequest) SetColumnIdKey(v string) *UpdateEnumMappingRequest {
	s.ColumnIdKey = &v
	return s
}

func (s *UpdateEnumMappingRequest) SetEnumMapping(v map[string][]*string) *UpdateEnumMappingRequest {
	s.EnumMapping = v
	return s
}

func (s *UpdateEnumMappingRequest) SetTableIdKey(v string) *UpdateEnumMappingRequest {
	s.TableIdKey = &v
	return s
}

func (s *UpdateEnumMappingRequest) SetWorkspaceId(v string) *UpdateEnumMappingRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateEnumMappingResponseBody struct {
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

func (s UpdateEnumMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnumMappingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnumMappingResponseBody) SetCode(v string) *UpdateEnumMappingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnumMappingResponseBody) SetData(v interface{}) *UpdateEnumMappingResponseBody {
	s.Data = v
	return s
}

func (s *UpdateEnumMappingResponseBody) SetErrorMsg(v string) *UpdateEnumMappingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateEnumMappingResponseBody) SetRequestId(v string) *UpdateEnumMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnumMappingResponseBody) SetSuccess(v bool) *UpdateEnumMappingResponseBody {
	s.Success = &v
	return s
}

type UpdateEnumMappingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnumMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnumMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEnumMappingResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnumMappingResponse) SetHeaders(v map[string]*string) *UpdateEnumMappingResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnumMappingResponse) SetStatusCode(v int32) *UpdateEnumMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnumMappingResponse) SetBody(v *UpdateEnumMappingResponseBody) *UpdateEnumMappingResponse {
	s.Body = v
	return s
}

type UpdateSynonymsRequest struct {
	Columns []*string `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// synonyms-AAAAAAAAAUpwTTVrwTFJwQ
	SynonymIdKey *string `json:"synonymIdKey,omitempty" xml:"synonymIdKey,omitempty"`
	// This parameter is required.
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
	// This parameter is required.
	WordSynonyms []*string `json:"wordSynonyms,omitempty" xml:"wordSynonyms,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateSynonymsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsRequest) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsRequest) SetColumns(v []*string) *UpdateSynonymsRequest {
	s.Columns = v
	return s
}

func (s *UpdateSynonymsRequest) SetSynonymIdKey(v string) *UpdateSynonymsRequest {
	s.SynonymIdKey = &v
	return s
}

func (s *UpdateSynonymsRequest) SetWord(v string) *UpdateSynonymsRequest {
	s.Word = &v
	return s
}

func (s *UpdateSynonymsRequest) SetWordSynonyms(v []*string) *UpdateSynonymsRequest {
	s.WordSynonyms = v
	return s
}

func (s *UpdateSynonymsRequest) SetWorkspaceId(v string) *UpdateSynonymsRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateSynonymsResponseBody struct {
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

func (s UpdateSynonymsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsResponseBody) SetCode(v string) *UpdateSynonymsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSynonymsResponseBody) SetData(v interface{}) *UpdateSynonymsResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSynonymsResponseBody) SetErrorMsg(v string) *UpdateSynonymsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateSynonymsResponseBody) SetRequestId(v string) *UpdateSynonymsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSynonymsResponseBody) SetSuccess(v bool) *UpdateSynonymsResponseBody {
	s.Success = &v
	return s
}

type UpdateSynonymsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSynonymsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSynonymsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSynonymsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSynonymsResponse) SetHeaders(v map[string]*string) *UpdateSynonymsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSynonymsResponse) SetStatusCode(v int32) *UpdateSynonymsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSynonymsResponse) SetBody(v *UpdateSynonymsResponseBody) *UpdateSynonymsResponse {
	s.Body = v
	return s
}

type UpdateTableInfoRequest struct {
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	ForeignKeys []*string `json:"foreignKeys,omitempty" xml:"foreignKeys,omitempty" type:"Repeated"`
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table-AAAAAAAAFQBwSLJkUj4CYg
	TableIdKey *string `json:"tableIdKey,omitempty" xml:"tableIdKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2v3934xtp49esw64
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s UpdateTableInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableInfoRequest) SetDescription(v string) *UpdateTableInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateTableInfoRequest) SetForeignKeys(v []*string) *UpdateTableInfoRequest {
	s.ForeignKeys = v
	return s
}

func (s *UpdateTableInfoRequest) SetPrimaryKey(v string) *UpdateTableInfoRequest {
	s.PrimaryKey = &v
	return s
}

func (s *UpdateTableInfoRequest) SetTableIdKey(v string) *UpdateTableInfoRequest {
	s.TableIdKey = &v
	return s
}

func (s *UpdateTableInfoRequest) SetWorkspaceId(v string) *UpdateTableInfoRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateTableInfoResponseBody struct {
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

func (s UpdateTableInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableInfoResponseBody) SetCode(v string) *UpdateTableInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTableInfoResponseBody) SetData(v interface{}) *UpdateTableInfoResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTableInfoResponseBody) SetErrorMsg(v string) *UpdateTableInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTableInfoResponseBody) SetRequestId(v string) *UpdateTableInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableInfoResponseBody) SetSuccess(v bool) *UpdateTableInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateTableInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTableInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTableInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableInfoResponse) SetHeaders(v map[string]*string) *UpdateTableInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableInfoResponse) SetStatusCode(v int32) *UpdateTableInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableInfoResponse) SetBody(v *UpdateTableInfoResponseBody) *UpdateTableInfoResponse {
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
// 批量删除当前指定业务空间下的同义词
//
// @param request - BatchDeleteSynonymsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteSynonymsResponse
func (client *Client) BatchDeleteSynonymsWithOptions(request *BatchDeleteSynonymsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchDeleteSynonymsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynonymIdKeys)) {
		body["synonymIdKeys"] = request.SynonymIdKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteSynonyms"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/batchDelete/synonyms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BatchDeleteSynonymsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BatchDeleteSynonymsResponse{}
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
// 批量删除当前指定业务空间下的同义词
//
// @param request - BatchDeleteSynonymsRequest
//
// @return BatchDeleteSynonymsResponse
func (client *Client) BatchDeleteSynonyms(request *BatchDeleteSynonymsRequest) (_result *BatchDeleteSynonymsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteSynonymsResponse{}
	_body, _err := client.BatchDeleteSynonymsWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CancelDatasourceAuthorizationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CancelDatasourceAuthorizationResponse{}
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
// 在指定的业务空间下创建新的业务逻辑解释
//
// @param request - CreateBusinessLogicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessLogicResponse
func (client *Client) CreateBusinessLogicWithOptions(request *CreateBusinessLogicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBusinessLogicResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBusinessLogic"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/create/logic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateBusinessLogicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateBusinessLogicResponse{}
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
// 在指定的业务空间下创建新的业务逻辑解释
//
// @param request - CreateBusinessLogicRequest
//
// @return CreateBusinessLogicResponse
func (client *Client) CreateBusinessLogic(request *CreateBusinessLogicRequest) (_result *CreateBusinessLogicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBusinessLogicResponse{}
	_body, _err := client.CreateBusinessLogicWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateDatasourceAuthorizationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateDatasourceAuthorizationResponse{}
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
// 在当前指定的业务空间下面，新建同义词
//
// @param request - CreateSynonymsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSynonymsResponse
func (client *Client) CreateSynonymsWithOptions(request *CreateSynonymsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSynonymsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		body["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		body["word"] = request.Word
	}

	if !tea.BoolValue(util.IsUnset(request.WordSynonyms)) {
		body["wordSynonyms"] = request.WordSynonyms
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSynonyms"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/create/synonyms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSynonymsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSynonymsResponse{}
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
// 在当前指定的业务空间下面，新建同义词
//
// @param request - CreateSynonymsRequest
//
// @return CreateSynonymsResponse
func (client *Client) CreateSynonyms(request *CreateSynonymsRequest) (_result *CreateSynonymsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSynonymsResponse{}
	_body, _err := client.CreateSynonymsWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateVirtualDatasourceInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateVirtualDatasourceInstanceResponse{}
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
// 删除指定业务空间下所指定的业务逻辑解释
//
// @param request - DeleteBusinessLogicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessLogicResponse
func (client *Client) DeleteBusinessLogicWithOptions(request *DeleteBusinessLogicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBusinessLogicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessLogicIdKeys)) {
		body["businessLogicIdKeys"] = request.BusinessLogicIdKeys
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBusinessLogic"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/delete/logic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteBusinessLogicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteBusinessLogicResponse{}
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
// 删除指定业务空间下所指定的业务逻辑解释
//
// @param request - DeleteBusinessLogicRequest
//
// @return DeleteBusinessLogicResponse
func (client *Client) DeleteBusinessLogic(request *DeleteBusinessLogicRequest) (_result *DeleteBusinessLogicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBusinessLogicResponse{}
	_body, _err := client.DeleteBusinessLogicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从当前所指定的业务空间中，删除所指定的列
//
// @param request - DeleteColumnRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteColumnResponse
func (client *Client) DeleteColumnWithOptions(request *DeleteColumnRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteColumnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnIdKey)) {
		body["columnIdKey"] = request.ColumnIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteColumn"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/delete/column"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteColumnResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteColumnResponse{}
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
// 从当前所指定的业务空间中，删除所指定的列
//
// @param request - DeleteColumnRequest
//
// @return DeleteColumnResponse
func (client *Client) DeleteColumn(request *DeleteColumnRequest) (_result *DeleteColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteColumnResponse{}
	_body, _err := client.DeleteColumnWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将当前指定数据表从指定业务空间管控中删除
//
// @param request - DeleteSelectedTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSelectedTableResponse
func (client *Client) DeleteSelectedTableWithOptions(request *DeleteSelectedTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSelectedTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSelectedTable"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/delete/table"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSelectedTableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSelectedTableResponse{}
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
// 将当前指定数据表从指定业务空间管控中删除
//
// @param request - DeleteSelectedTableRequest
//
// @return DeleteSelectedTableResponse
func (client *Client) DeleteSelectedTable(request *DeleteSelectedTableRequest) (_result *DeleteSelectedTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSelectedTableResponse{}
	_body, _err := client.DeleteSelectedTableWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteVirtualDatasourceInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteVirtualDatasourceInstanceResponse{}
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
// 获取当前指定业务空间下的企业知识名词解释列表
//
// @param request - ListBusinessLogicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBusinessLogicResponse
func (client *Client) ListBusinessLogicWithOptions(request *ListBusinessLogicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBusinessLogicResponse, _err error) {
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
		Action:      tea.String("ListBusinessLogic"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/list/logic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListBusinessLogicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListBusinessLogicResponse{}
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
// 获取当前指定业务空间下的企业知识名词解释列表
//
// @param request - ListBusinessLogicRequest
//
// @return ListBusinessLogicResponse
func (client *Client) ListBusinessLogic(request *ListBusinessLogicRequest) (_result *ListBusinessLogicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBusinessLogicResponse{}
	_body, _err := client.ListBusinessLogicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前指定业务空间，指定表下面的列信息
//
// @param request - ListColumnRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListColumnResponse
func (client *Client) ListColumnWithOptions(request *ListColumnRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListColumnResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListColumn"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/list/column"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListColumnResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListColumnResponse{}
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
// 获取当前指定业务空间，指定表下面的列信息
//
// @param request - ListColumnRequest
//
// @return ListColumnResponse
func (client *Client) ListColumn(request *ListColumnRequest) (_result *ListColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListColumnResponse{}
	_body, _err := client.ListColumnWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前业务空间，指定表、列下的枚举值
//
// @param request - ListEnumMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnumMappingResponse
func (client *Client) ListEnumMappingWithOptions(request *ListEnumMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEnumMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnIdKey)) {
		body["columnIdKey"] = request.ColumnIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEnumMapping"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/list/mapping"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListEnumMappingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListEnumMappingResponse{}
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
// 获取当前业务空间，指定表、列下的枚举值
//
// @param request - ListEnumMappingRequest
//
// @return ListEnumMappingResponse
func (client *Client) ListEnumMapping(request *ListEnumMappingRequest) (_result *ListEnumMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnumMappingResponse{}
	_body, _err := client.ListEnumMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前业务空间处于以关联状态的数据表
//
// @param request - ListSelectedTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSelectedTablesResponse
func (client *Client) ListSelectedTablesWithOptions(request *ListSelectedTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSelectedTablesResponse, _err error) {
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
		Action:      tea.String("ListSelectedTables"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/list/datasource/table"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSelectedTablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSelectedTablesResponse{}
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
// 获取当前业务空间处于以关联状态的数据表
//
// @param request - ListSelectedTablesRequest
//
// @return ListSelectedTablesResponse
func (client *Client) ListSelectedTables(request *ListSelectedTablesRequest) (_result *ListSelectedTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSelectedTablesResponse{}
	_body, _err := client.ListSelectedTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取当前指定业务空间下的同义词列表
//
// @param request - ListSynonymsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSynonymsResponse
func (client *Client) ListSynonymsWithOptions(request *ListSynonymsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSynonymsResponse, _err error) {
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
		Action:      tea.String("ListSynonyms"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/list/synonyms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListSynonymsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListSynonymsResponse{}
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
// 获取当前指定业务空间下的同义词列表
//
// @param request - ListSynonymsRequest
//
// @return ListSynonymsResponse
func (client *Client) ListSynonyms(request *ListSynonymsRequest) (_result *ListSynonymsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSynonymsResponse{}
	_body, _err := client.ListSynonymsWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListVirtualDatasourceInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListVirtualDatasourceInstanceResponse{}
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
// 将指定数据表的数据列恢复到初始话状态
//
// @param request - RecoverColumnRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverColumnResponse
func (client *Client) RecoverColumnWithOptions(request *RecoverColumnRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecoverColumnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnIdKey)) {
		body["columnIdKey"] = request.ColumnIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverColumn"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/recover/column"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RecoverColumnResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RecoverColumnResponse{}
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
// 将指定数据表的数据列恢复到初始话状态
//
// @param request - RecoverColumnRequest
//
// @return RecoverColumnResponse
func (client *Client) RecoverColumn(request *RecoverColumnRequest) (_result *RecoverColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecoverColumnResponse{}
	_body, _err := client.RecoverColumnWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从远程数据库刷新当前所关联的数据表信息
//
// @param request - ResyncTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResyncTableResponse
func (client *Client) ResyncTableWithOptions(request *ResyncTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResyncTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keep)) {
		body["keep"] = request.Keep
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResyncTable"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/refresh/table"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResyncTableResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResyncTableResponse{}
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
// 从远程数据库刷新当前所关联的数据表信息
//
// @param request - ResyncTableRequest
//
// @return ResyncTableResponse
func (client *Client) ResyncTable(request *ResyncTableRequest) (_result *ResyncTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResyncTableResponse{}
	_body, _err := client.ResyncTableWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DataRole)) {
		body["dataRole"] = request.DataRole
	}

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

	if !tea.BoolValue(util.IsUnset(request.UserParams)) {
		body["userParams"] = request.UserParams
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunDataAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunDataAnalysisResponse{}
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
// 对结构化结果进行分析、可视化信息生成
//
// @param request - RunDataResultAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDataResultAnalysisResponse
func (client *Client) RunDataResultAnalysisWithOptions(request *RunDataResultAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunDataResultAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalysisMode)) {
		body["analysisMode"] = request.AnalysisMode
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		body["requestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.SqlData)) {
		body["sqlData"] = request.SqlData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunDataResultAnalysis"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/runDataResultAnalysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunDataResultAnalysisResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunDataResultAnalysisResponse{}
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
// 对结构化结果进行分析、可视化信息生成
//
// @param request - RunDataResultAnalysisRequest
//
// @return RunDataResultAnalysisResponse
func (client *Client) RunDataResultAnalysis(request *RunDataResultAnalysisRequest) (_result *RunDataResultAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunDataResultAnalysisResponse{}
	_body, _err := client.RunDataResultAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行sql生成
//
// @param request - RunSqlGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSqlGenerationResponse
func (client *Client) RunSqlGenerationWithOptions(request *RunSqlGenerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunSqlGenerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
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
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunSqlGeneration"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/runSqlGeneration"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RunSqlGenerationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RunSqlGenerationResponse{}
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
// 运行sql生成
//
// @param request - RunSqlGenerationRequest
//
// @return RunSqlGenerationResponse
func (client *Client) RunSqlGeneration(request *RunSqlGenerationRequest) (_result *RunSqlGenerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunSqlGenerationResponse{}
	_body, _err := client.RunSqlGenerationWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SaveVirtualDatasourceDdlResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SaveVirtualDatasourceDdlResponse{}
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

	if !tea.BoolValue(util.IsUnset(request.NoModifiedTableNames)) {
		body["noModifiedTableNames"] = request.NoModifiedTableNames
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SyncRemoteTablesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SyncRemoteTablesResponse{}
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
// 修改当前指定业务空间下所指定的业务逻辑解释
//
// @param request - UpdateBusinessLogicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBusinessLogicResponse
func (client *Client) UpdateBusinessLogicWithOptions(request *UpdateBusinessLogicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBusinessLogicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessLogicIdKey)) {
		body["businessLogicIdKey"] = request.BusinessLogicIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
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
		Action:      tea.String("UpdateBusinessLogic"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/logic"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateBusinessLogicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateBusinessLogicResponse{}
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
// 修改当前指定业务空间下所指定的业务逻辑解释
//
// @param request - UpdateBusinessLogicRequest
//
// @return UpdateBusinessLogicResponse
func (client *Client) UpdateBusinessLogic(request *UpdateBusinessLogicRequest) (_result *UpdateBusinessLogicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBusinessLogicResponse{}
	_body, _err := client.UpdateBusinessLogicWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改当前指定业务空间中，指定列的信息
//
// @param request - UpdateColumnRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateColumnResponse
func (client *Client) UpdateColumnWithOptions(request *UpdateColumnRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateColumnResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChineseName)) {
		body["chineseName"] = request.ChineseName
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnIdKey)) {
		body["columnIdKey"] = request.ColumnIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnumType)) {
		body["enumType"] = request.EnumType
	}

	if !tea.BoolValue(util.IsUnset(request.EnumValues)) {
		body["enumValues"] = request.EnumValues
	}

	if !tea.BoolValue(util.IsUnset(request.RangeMax)) {
		body["rangeMax"] = request.RangeMax
	}

	if !tea.BoolValue(util.IsUnset(request.RangeMin)) {
		body["rangeMin"] = request.RangeMin
	}

	if !tea.BoolValue(util.IsUnset(request.Samples)) {
		body["samples"] = request.Samples
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateColumn"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/column"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateColumnResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateColumnResponse{}
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
// 修改当前指定业务空间中，指定列的信息
//
// @param request - UpdateColumnRequest
//
// @return UpdateColumnResponse
func (client *Client) UpdateColumn(request *UpdateColumnRequest) (_result *UpdateColumnResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateColumnResponse{}
	_body, _err := client.UpdateColumnWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改当前指定业务空间指定列下的枚举值信息
//
// @param request - UpdateEnumMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnumMappingResponse
func (client *Client) UpdateEnumMappingWithOptions(request *UpdateEnumMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateEnumMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColumnIdKey)) {
		body["columnIdKey"] = request.ColumnIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.EnumMapping)) {
		body["enumMapping"] = request.EnumMapping
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEnumMapping"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/mapping"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateEnumMappingResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateEnumMappingResponse{}
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
// 修改当前指定业务空间指定列下的枚举值信息
//
// @param request - UpdateEnumMappingRequest
//
// @return UpdateEnumMappingResponse
func (client *Client) UpdateEnumMapping(request *UpdateEnumMappingRequest) (_result *UpdateEnumMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnumMappingResponse{}
	_body, _err := client.UpdateEnumMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改当前业务空间指定的同义词信息
//
// @param request - UpdateSynonymsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSynonymsResponse
func (client *Client) UpdateSynonymsWithOptions(request *UpdateSynonymsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSynonymsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		body["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.SynonymIdKey)) {
		body["synonymIdKey"] = request.SynonymIdKey
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		body["word"] = request.Word
	}

	if !tea.BoolValue(util.IsUnset(request.WordSynonyms)) {
		body["wordSynonyms"] = request.WordSynonyms
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSynonyms"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/synonyms"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateSynonymsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateSynonymsResponse{}
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
// 修改当前业务空间指定的同义词信息
//
// @param request - UpdateSynonymsRequest
//
// @return UpdateSynonymsResponse
func (client *Client) UpdateSynonyms(request *UpdateSynonymsRequest) (_result *UpdateSynonymsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSynonymsResponse{}
	_body, _err := client.UpdateSynonymsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改当前所指定的数据表的信息
//
// @param request - UpdateTableInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableInfoResponse
func (client *Client) UpdateTableInfoWithOptions(request *UpdateTableInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTableInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ForeignKeys)) {
		body["foreignKeys"] = request.ForeignKeys
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryKey)) {
		body["primaryKey"] = request.PrimaryKey
	}

	if !tea.BoolValue(util.IsUnset(request.TableIdKey)) {
		body["tableIdKey"] = request.TableIdKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTableInfo"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/gbi/update/table"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTableInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTableInfoResponse{}
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
// 修改当前所指定的数据表的信息
//
// @param request - UpdateTableInfoRequest
//
// @return UpdateTableInfoResponse
func (client *Client) UpdateTableInfo(request *UpdateTableInfoRequest) (_result *UpdateTableInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableInfoResponse{}
	_body, _err := client.UpdateTableInfoWithOptions(request, headers, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateVirtualDatasourceInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateVirtualDatasourceInstanceResponse{}
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
