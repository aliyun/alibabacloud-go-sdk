// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCodeupSourceToPipelineRequest struct {
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	CodePath   *string `json:"CodePath,omitempty" xml:"CodePath,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s AddCodeupSourceToPipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCodeupSourceToPipelineRequest) GoString() string {
	return s.String()
}

func (s *AddCodeupSourceToPipelineRequest) SetCodeBranch(v string) *AddCodeupSourceToPipelineRequest {
	s.CodeBranch = &v
	return s
}

func (s *AddCodeupSourceToPipelineRequest) SetCodePath(v string) *AddCodeupSourceToPipelineRequest {
	s.CodePath = &v
	return s
}

func (s *AddCodeupSourceToPipelineRequest) SetOrgId(v string) *AddCodeupSourceToPipelineRequest {
	s.OrgId = &v
	return s
}

func (s *AddCodeupSourceToPipelineRequest) SetPipelineId(v int64) *AddCodeupSourceToPipelineRequest {
	s.PipelineId = &v
	return s
}

type AddCodeupSourceToPipelineResponseBody struct {
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCodeupSourceToPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCodeupSourceToPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *AddCodeupSourceToPipelineResponseBody) SetPipelineId(v int64) *AddCodeupSourceToPipelineResponseBody {
	s.PipelineId = &v
	return s
}

func (s *AddCodeupSourceToPipelineResponseBody) SetRequestId(v string) *AddCodeupSourceToPipelineResponseBody {
	s.RequestId = &v
	return s
}

type AddCodeupSourceToPipelineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddCodeupSourceToPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddCodeupSourceToPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCodeupSourceToPipelineResponse) GoString() string {
	return s.String()
}

func (s *AddCodeupSourceToPipelineResponse) SetHeaders(v map[string]*string) *AddCodeupSourceToPipelineResponse {
	s.Headers = v
	return s
}

func (s *AddCodeupSourceToPipelineResponse) SetStatusCode(v int32) *AddCodeupSourceToPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCodeupSourceToPipelineResponse) SetBody(v *AddCodeupSourceToPipelineResponseBody) *AddCodeupSourceToPipelineResponse {
	s.Body = v
	return s
}

type BatchInsertMembersRequest struct {
	Members *string `json:"Members,omitempty" xml:"Members,omitempty"`
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	RealPk  *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s BatchInsertMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertMembersRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertMembersRequest) SetMembers(v string) *BatchInsertMembersRequest {
	s.Members = &v
	return s
}

func (s *BatchInsertMembersRequest) SetOrgId(v string) *BatchInsertMembersRequest {
	s.OrgId = &v
	return s
}

func (s *BatchInsertMembersRequest) SetRealPk(v string) *BatchInsertMembersRequest {
	s.RealPk = &v
	return s
}

type BatchInsertMembersResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertMembersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertMembersResponseBody) SetErrorCode(v string) *BatchInsertMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetErrorMessage(v string) *BatchInsertMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetObject(v bool) *BatchInsertMembersResponseBody {
	s.Object = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetRequestId(v string) *BatchInsertMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetSuccess(v bool) *BatchInsertMembersResponseBody {
	s.Success = &v
	return s
}

type BatchInsertMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchInsertMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchInsertMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertMembersResponse) GoString() string {
	return s.String()
}

func (s *BatchInsertMembersResponse) SetHeaders(v map[string]*string) *BatchInsertMembersResponse {
	s.Headers = v
	return s
}

func (s *BatchInsertMembersResponse) SetStatusCode(v int32) *BatchInsertMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchInsertMembersResponse) SetBody(v *BatchInsertMembersResponseBody) *BatchInsertMembersResponse {
	s.Body = v
	return s
}

type CancelPipelineRequest struct {
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CancelPipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineRequest) GoString() string {
	return s.String()
}

func (s *CancelPipelineRequest) SetFlowInstanceId(v int64) *CancelPipelineRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *CancelPipelineRequest) SetOrgId(v string) *CancelPipelineRequest {
	s.OrgId = &v
	return s
}

func (s *CancelPipelineRequest) SetPipelineId(v int64) *CancelPipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *CancelPipelineRequest) SetUserPk(v string) *CancelPipelineRequest {
	s.UserPk = &v
	return s
}

type CancelPipelineResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPipelineResponseBody) SetErrorCode(v string) *CancelPipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelPipelineResponseBody) SetErrorMessage(v string) *CancelPipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CancelPipelineResponseBody) SetObject(v bool) *CancelPipelineResponseBody {
	s.Object = &v
	return s
}

func (s *CancelPipelineResponseBody) SetRequestId(v string) *CancelPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPipelineResponseBody) SetSuccess(v bool) *CancelPipelineResponseBody {
	s.Success = &v
	return s
}

type CancelPipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelPipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineResponse) GoString() string {
	return s.String()
}

func (s *CancelPipelineResponse) SetHeaders(v map[string]*string) *CancelPipelineResponse {
	s.Headers = v
	return s
}

func (s *CancelPipelineResponse) SetStatusCode(v int32) *CancelPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelPipelineResponse) SetBody(v *CancelPipelineResponseBody) *CancelPipelineResponse {
	s.Body = v
	return s
}

type CreateCommonGroupRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId        *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SmartGroupId *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
}

func (s CreateCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupRequest) SetDescription(v string) *CreateCommonGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateCommonGroupRequest) SetName(v string) *CreateCommonGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateCommonGroupRequest) SetOrgId(v string) *CreateCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *CreateCommonGroupRequest) SetProjectId(v string) *CreateCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCommonGroupRequest) SetSmartGroupId(v string) *CreateCommonGroupRequest {
	s.SmartGroupId = &v
	return s
}

type CreateCommonGroupResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupResponseBody) SetErrorCode(v string) *CreateCommonGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCommonGroupResponseBody) SetErrorMsg(v string) *CreateCommonGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateCommonGroupResponseBody) SetObject(v *CreateCommonGroupResponseBodyObject) *CreateCommonGroupResponseBody {
	s.Object = v
	return s
}

func (s *CreateCommonGroupResponseBody) SetRequestId(v string) *CreateCommonGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCommonGroupResponseBody) SetSuccessful(v bool) *CreateCommonGroupResponseBody {
	s.Successful = &v
	return s
}

type CreateCommonGroupResponseBodyObject struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateCommonGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupResponseBodyObject) SetId(v string) *CreateCommonGroupResponseBodyObject {
	s.Id = &v
	return s
}

type CreateCommonGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCommonGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupResponse) SetHeaders(v map[string]*string) *CreateCommonGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCommonGroupResponse) SetStatusCode(v int32) *CreateCommonGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommonGroupResponse) SetBody(v *CreateCommonGroupResponseBody) *CreateCommonGroupResponse {
	s.Body = v
	return s
}

type CreateCredentialRequest struct {
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPk   *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) SetName(v string) *CreateCredentialRequest {
	s.Name = &v
	return s
}

func (s *CreateCredentialRequest) SetOrgId(v string) *CreateCredentialRequest {
	s.OrgId = &v
	return s
}

func (s *CreateCredentialRequest) SetPassword(v string) *CreateCredentialRequest {
	s.Password = &v
	return s
}

func (s *CreateCredentialRequest) SetType(v string) *CreateCredentialRequest {
	s.Type = &v
	return s
}

func (s *CreateCredentialRequest) SetUserName(v string) *CreateCredentialRequest {
	s.UserName = &v
	return s
}

func (s *CreateCredentialRequest) SetUserPk(v string) *CreateCredentialRequest {
	s.UserPk = &v
	return s
}

type CreateCredentialResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) SetErrorCode(v string) *CreateCredentialResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCredentialResponseBody) SetErrorMessage(v string) *CreateCredentialResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCredentialResponseBody) SetObject(v int64) *CreateCredentialResponseBody {
	s.Object = &v
	return s
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetSuccess(v bool) *CreateCredentialResponseBody {
	s.Success = &v
	return s
}

type CreateCredentialResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponse) SetHeaders(v map[string]*string) *CreateCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateCredentialResponse) SetStatusCode(v int32) *CreateCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCredentialResponse) SetBody(v *CreateCredentialResponseBody) *CreateCredentialResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s CreateDevopsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectRequest) SetDescription(v string) *CreateDevopsProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateDevopsProjectRequest) SetName(v string) *CreateDevopsProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectRequest) SetOrgId(v string) *CreateDevopsProjectRequest {
	s.OrgId = &v
	return s
}

type CreateDevopsProjectResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectResponseBody) SetErrorCode(v string) *CreateDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetErrorMessage(v string) *CreateDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetObject(v string) *CreateDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetRequestId(v string) *CreateDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetSuccess(v bool) *CreateDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type CreateDevopsProjectResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevopsProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectResponse) SetHeaders(v map[string]*string) *CreateDevopsProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateDevopsProjectResponse) SetStatusCode(v int32) *CreateDevopsProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDevopsProjectResponse) SetBody(v *CreateDevopsProjectResponseBody) *CreateDevopsProjectResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectSprintRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s CreateDevopsProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintRequest) SetDescription(v string) *CreateDevopsProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetDueDate(v string) *CreateDevopsProjectSprintRequest {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetExecutorId(v string) *CreateDevopsProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetName(v string) *CreateDevopsProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetOrgId(v string) *CreateDevopsProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetProjectId(v string) *CreateDevopsProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetStartDate(v string) *CreateDevopsProjectSprintRequest {
	s.StartDate = &v
	return s
}

type CreateDevopsProjectSprintResponseBody struct {
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateDevopsProjectSprintResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBody) SetErrorCode(v string) *CreateDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBody) SetErrorMsg(v string) *CreateDevopsProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBody) SetObject(v *CreateDevopsProjectSprintResponseBodyObject) *CreateDevopsProjectSprintResponseBody {
	s.Object = v
	return s
}

func (s *CreateDevopsProjectSprintResponseBody) SetRequestId(v string) *CreateDevopsProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBody) SetSuccessful(v bool) *CreateDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type CreateDevopsProjectSprintResponseBodyObject struct {
	Accomplished *string                                              `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	Created      *string                                              `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId    *string                                              `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description  *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	DueDate      *string                                              `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Executor     *string                                              `json:"Executor,omitempty" xml:"Executor,omitempty"`
	Id           *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDeleted    *bool                                                `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Name         *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanToDo     *CreateDevopsProjectSprintResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	ProjectId    *string                                              `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                              `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status       *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Updated      *string                                              `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetAccomplished(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetCreated(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetCreatorId(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetDescription(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Description = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetDueDate(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetExecutor(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Executor = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetId(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Id = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetIsDeleted(v bool) *CreateDevopsProjectSprintResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetName(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetPlanToDo(v *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) *CreateDevopsProjectSprintResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetProjectId(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetStartDate(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetStatus(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Status = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetUpdated(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Updated = &v
	return s
}

type CreateDevopsProjectSprintResponseBodyObjectPlanToDo struct {
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetTasks(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

type CreateDevopsProjectSprintResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevopsProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponse) SetHeaders(v map[string]*string) *CreateDevopsProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *CreateDevopsProjectSprintResponse) SetStatusCode(v int32) *CreateDevopsProjectSprintResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDevopsProjectSprintResponse) SetBody(v *CreateDevopsProjectSprintResponseBody) *CreateDevopsProjectSprintResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectTaskRequest struct {
	Content               *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DueDate               *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Note                  *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ParentTaskId          *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	Priority              *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate             *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaskFlowStatusId      *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	TaskListId            *string `json:"TaskListId,omitempty" xml:"TaskListId,omitempty"`
	Visible               *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateDevopsProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskRequest) SetContent(v string) *CreateDevopsProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetDueDate(v string) *CreateDevopsProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetExecutorId(v string) *CreateDevopsProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetNote(v string) *CreateDevopsProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetOrgId(v string) *CreateDevopsProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetParentTaskId(v string) *CreateDevopsProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetPriority(v int32) *CreateDevopsProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetProjectId(v string) *CreateDevopsProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetScenarioFieldConfigId(v string) *CreateDevopsProjectTaskRequest {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetSprintId(v string) *CreateDevopsProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetStartDate(v string) *CreateDevopsProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetTaskFlowStatusId(v string) *CreateDevopsProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetTaskListId(v string) *CreateDevopsProjectTaskRequest {
	s.TaskListId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetVisible(v string) *CreateDevopsProjectTaskRequest {
	s.Visible = &v
	return s
}

type CreateDevopsProjectTaskResponseBody struct {
	ErrorCode  *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateDevopsProjectTaskResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                      `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskResponseBody) SetErrorCode(v string) *CreateDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBody) SetErrorMsg(v string) *CreateDevopsProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBody) SetObject(v *CreateDevopsProjectTaskResponseBodyObject) *CreateDevopsProjectTaskResponseBody {
	s.Object = v
	return s
}

func (s *CreateDevopsProjectTaskResponseBody) SetRequestId(v string) *CreateDevopsProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBody) SetSuccessful(v bool) *CreateDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type CreateDevopsProjectTaskResponseBodyObject struct {
	Content               *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Created               *string `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId             *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DueDate               *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDone                *bool   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	Note                  *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OrganizationId        *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Pos                   *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Priority              *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Rating                *int32  `json:"Rating,omitempty" xml:"Rating,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate             *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint            *string `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	TaskType              *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskflowstatusId      *string `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	TasklistId            *string `json:"TasklistId,omitempty" xml:"TasklistId,omitempty"`
	UniqueId              *int32  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	Updated               *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visible               *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateDevopsProjectTaskResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetContent(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Content = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetCreated(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetCreatorId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetDueDate(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetExecutorId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Id = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetIsDone(v bool) *CreateDevopsProjectTaskResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetNote(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Note = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetOrganizationId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetPos(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetPriority(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetProjectId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetRating(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Rating = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetScenarioFieldConfigId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetSource(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Source = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetSprintId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetStartDate(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetStoryPoint(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTaskType(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTaskflowstatusId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTasklistId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetUniqueId(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.UniqueId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetUpdated(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetVisible(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Visible = &v
	return s
}

type CreateDevopsProjectTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevopsProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskResponse) SetHeaders(v map[string]*string) *CreateDevopsProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDevopsProjectTaskResponse) SetStatusCode(v int32) *CreateDevopsProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDevopsProjectTaskResponse) SetBody(v *CreateDevopsProjectTaskResponseBody) *CreateDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type CreatePipelineFromTemplateRequest struct {
	OrgId              *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineName       *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	PipelineTemplateId *int64  `json:"PipelineTemplateId,omitempty" xml:"PipelineTemplateId,omitempty"`
}

func (s CreatePipelineFromTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineFromTemplateRequest) SetOrgId(v string) *CreatePipelineFromTemplateRequest {
	s.OrgId = &v
	return s
}

func (s *CreatePipelineFromTemplateRequest) SetPipelineName(v string) *CreatePipelineFromTemplateRequest {
	s.PipelineName = &v
	return s
}

func (s *CreatePipelineFromTemplateRequest) SetPipelineTemplateId(v int64) *CreatePipelineFromTemplateRequest {
	s.PipelineTemplateId = &v
	return s
}

type CreatePipelineFromTemplateResponseBody struct {
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePipelineFromTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineFromTemplateResponseBody) SetPipelineId(v int64) *CreatePipelineFromTemplateResponseBody {
	s.PipelineId = &v
	return s
}

func (s *CreatePipelineFromTemplateResponseBody) SetRequestId(v string) *CreatePipelineFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreatePipelineFromTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePipelineFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelineFromTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineFromTemplateResponse) SetHeaders(v map[string]*string) *CreatePipelineFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineFromTemplateResponse) SetStatusCode(v int32) *CreatePipelineFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePipelineFromTemplateResponse) SetBody(v *CreatePipelineFromTemplateResponseBody) *CreatePipelineFromTemplateResponse {
	s.Body = v
	return s
}

type CreateServiceConnectionRequest struct {
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ServiceConnectionType *string `json:"ServiceConnectionType,omitempty" xml:"ServiceConnectionType,omitempty"`
	UserPk                *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CreateServiceConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionRequest) SetOrgId(v string) *CreateServiceConnectionRequest {
	s.OrgId = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetServiceConnectionType(v string) *CreateServiceConnectionRequest {
	s.ServiceConnectionType = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetUserPk(v string) *CreateServiceConnectionRequest {
	s.UserPk = &v
	return s
}

type CreateServiceConnectionResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionResponseBody) SetErrorCode(v string) *CreateServiceConnectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetErrorMessage(v string) *CreateServiceConnectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetObject(v int64) *CreateServiceConnectionResponseBody {
	s.Object = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetRequestId(v string) *CreateServiceConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetSuccess(v bool) *CreateServiceConnectionResponseBody {
	s.Success = &v
	return s
}

type CreateServiceConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionResponse) SetHeaders(v map[string]*string) *CreateServiceConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceConnectionResponse) SetStatusCode(v int32) *CreateServiceConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceConnectionResponse) SetBody(v *CreateServiceConnectionResponseBody) *CreateServiceConnectionResponse {
	s.Body = v
	return s
}

type DeleteCommonGroupRequest struct {
	CommonGroupId *string `json:"CommonGroupId,omitempty" xml:"CommonGroupId,omitempty"`
	OrgId         *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupRequest) SetCommonGroupId(v string) *DeleteCommonGroupRequest {
	s.CommonGroupId = &v
	return s
}

func (s *DeleteCommonGroupRequest) SetOrgId(v string) *DeleteCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteCommonGroupRequest) SetProjectId(v string) *DeleteCommonGroupRequest {
	s.ProjectId = &v
	return s
}

type DeleteCommonGroupResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *DeleteCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupResponseBody) SetErrorCode(v string) *DeleteCommonGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCommonGroupResponseBody) SetErrorMsg(v string) *DeleteCommonGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteCommonGroupResponseBody) SetObject(v *DeleteCommonGroupResponseBodyObject) *DeleteCommonGroupResponseBody {
	s.Object = v
	return s
}

func (s *DeleteCommonGroupResponseBody) SetRequestId(v string) *DeleteCommonGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommonGroupResponseBody) SetSuccessful(v bool) *DeleteCommonGroupResponseBody {
	s.Successful = &v
	return s
}

type DeleteCommonGroupResponseBodyObject struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteCommonGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupResponseBodyObject) SetId(v string) *DeleteCommonGroupResponseBodyObject {
	s.Id = &v
	return s
}

type DeleteCommonGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCommonGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupResponse) SetHeaders(v map[string]*string) *DeleteCommonGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommonGroupResponse) SetStatusCode(v int32) *DeleteCommonGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommonGroupResponse) SetBody(v *DeleteCommonGroupResponseBody) *DeleteCommonGroupResponse {
	s.Body = v
	return s
}

type DeleteDevopsOrganizationMembersRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	RealPk *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteDevopsOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevopsOrganizationMembersRequest) SetOrgId(v string) *DeleteDevopsOrganizationMembersRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersRequest) SetRealPk(v string) *DeleteDevopsOrganizationMembersRequest {
	s.RealPk = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersRequest) SetUserId(v string) *DeleteDevopsOrganizationMembersRequest {
	s.UserId = &v
	return s
}

type DeleteDevopsOrganizationMembersResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDevopsOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetErrorCode(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetErrorMessage(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetObject(v bool) *DeleteDevopsOrganizationMembersResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetRequestId(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetSuccess(v bool) *DeleteDevopsOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

type DeleteDevopsOrganizationMembersResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDevopsOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevopsOrganizationMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevopsOrganizationMembersResponse) SetHeaders(v map[string]*string) *DeleteDevopsOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponse) SetStatusCode(v int32) *DeleteDevopsOrganizationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponse) SetBody(v *DeleteDevopsOrganizationMembersResponseBody) *DeleteDevopsOrganizationMembersResponse {
	s.Body = v
	return s
}

type DeleteDevopsProjectRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteDevopsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectRequest) SetOrgId(v string) *DeleteDevopsProjectRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteDevopsProjectRequest) SetProjectId(v string) *DeleteDevopsProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteDevopsProjectResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectResponseBody) SetErrorCode(v string) *DeleteDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetErrorMessage(v string) *DeleteDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetObject(v string) *DeleteDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetRequestId(v string) *DeleteDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetSuccess(v bool) *DeleteDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteDevopsProjectResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevopsProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectResponse) SetHeaders(v map[string]*string) *DeleteDevopsProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevopsProjectResponse) SetStatusCode(v int32) *DeleteDevopsProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevopsProjectResponse) SetBody(v *DeleteDevopsProjectResponseBody) *DeleteDevopsProjectResponse {
	s.Body = v
	return s
}

type DeleteDevopsProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserIds   *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s DeleteDevopsProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectMembersRequest) SetOrgId(v string) *DeleteDevopsProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteDevopsProjectMembersRequest) SetProjectId(v string) *DeleteDevopsProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteDevopsProjectMembersRequest) SetUserIds(v string) *DeleteDevopsProjectMembersRequest {
	s.UserIds = &v
	return s
}

type DeleteDevopsProjectMembersResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectMembersResponseBody) SetErrorCode(v string) *DeleteDevopsProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponseBody) SetErrorMsg(v string) *DeleteDevopsProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponseBody) SetObject(v bool) *DeleteDevopsProjectMembersResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponseBody) SetRequestId(v string) *DeleteDevopsProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectMembersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDevopsProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevopsProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectMembersResponse) SetHeaders(v map[string]*string) *DeleteDevopsProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevopsProjectMembersResponse) SetStatusCode(v int32) *DeleteDevopsProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponse) SetBody(v *DeleteDevopsProjectMembersResponseBody) *DeleteDevopsProjectMembersResponse {
	s.Body = v
	return s
}

type DeleteDevopsProjectSprintRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	SprintId *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s DeleteDevopsProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectSprintRequest) SetOrgId(v string) *DeleteDevopsProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteDevopsProjectSprintRequest) SetSprintId(v string) *DeleteDevopsProjectSprintRequest {
	s.SprintId = &v
	return s
}

type DeleteDevopsProjectSprintResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectSprintResponseBody) SetErrorCode(v string) *DeleteDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponseBody) SetErrorMsg(v string) *DeleteDevopsProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponseBody) SetObject(v bool) *DeleteDevopsProjectSprintResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponseBody) SetRequestId(v string) *DeleteDevopsProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectSprintResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevopsProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectSprintResponse) SetHeaders(v map[string]*string) *DeleteDevopsProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevopsProjectSprintResponse) SetStatusCode(v int32) *DeleteDevopsProjectSprintResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponse) SetBody(v *DeleteDevopsProjectSprintResponseBody) *DeleteDevopsProjectSprintResponse {
	s.Body = v
	return s
}

type DeleteDevopsProjectTaskRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDevopsProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectTaskRequest) SetOrgId(v string) *DeleteDevopsProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteDevopsProjectTaskRequest) SetTaskId(v string) *DeleteDevopsProjectTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteDevopsProjectTaskResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectTaskResponseBody) SetErrorCode(v string) *DeleteDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponseBody) SetErrorMsg(v string) *DeleteDevopsProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponseBody) SetObject(v bool) *DeleteDevopsProjectTaskResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponseBody) SetRequestId(v string) *DeleteDevopsProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDevopsProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectTaskResponse) SetHeaders(v map[string]*string) *DeleteDevopsProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDevopsProjectTaskResponse) SetStatusCode(v int32) *DeleteDevopsProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponse) SetBody(v *DeleteDevopsProjectTaskResponseBody) *DeleteDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type DeletePipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s DeletePipelineMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineMemberRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelineMemberRequest) SetOrgId(v string) *DeletePipelineMemberRequest {
	s.OrgId = &v
	return s
}

func (s *DeletePipelineMemberRequest) SetPipelineId(v int64) *DeletePipelineMemberRequest {
	s.PipelineId = &v
	return s
}

func (s *DeletePipelineMemberRequest) SetUserId(v string) *DeletePipelineMemberRequest {
	s.UserId = &v
	return s
}

func (s *DeletePipelineMemberRequest) SetUserPk(v string) *DeletePipelineMemberRequest {
	s.UserPk = &v
	return s
}

type DeletePipelineMemberResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineMemberResponseBody) SetErrorCode(v string) *DeletePipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetErrorMessage(v string) *DeletePipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetObject(v bool) *DeletePipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetRequestId(v string) *DeletePipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetSuccess(v bool) *DeletePipelineMemberResponseBody {
	s.Success = &v
	return s
}

type DeletePipelineMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePipelineMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineMemberResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineMemberResponse) SetHeaders(v map[string]*string) *DeletePipelineMemberResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineMemberResponse) SetStatusCode(v int32) *DeletePipelineMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineMemberResponse) SetBody(v *DeletePipelineMemberResponseBody) *DeletePipelineMemberResponse {
	s.Body = v
	return s
}

type ExecutePipelineRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s ExecutePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecutePipelineRequest) GoString() string {
	return s.String()
}

func (s *ExecutePipelineRequest) SetOrgId(v string) *ExecutePipelineRequest {
	s.OrgId = &v
	return s
}

func (s *ExecutePipelineRequest) SetParameters(v string) *ExecutePipelineRequest {
	s.Parameters = &v
	return s
}

func (s *ExecutePipelineRequest) SetPipelineId(v int64) *ExecutePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *ExecutePipelineRequest) SetUserPk(v string) *ExecutePipelineRequest {
	s.UserPk = &v
	return s
}

type ExecutePipelineResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecutePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecutePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ExecutePipelineResponseBody) SetErrorCode(v string) *ExecutePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetErrorMessage(v string) *ExecutePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetObject(v int64) *ExecutePipelineResponseBody {
	s.Object = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetRequestId(v string) *ExecutePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetSuccess(v bool) *ExecutePipelineResponseBody {
	s.Success = &v
	return s
}

type ExecutePipelineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecutePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecutePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecutePipelineResponse) GoString() string {
	return s.String()
}

func (s *ExecutePipelineResponse) SetHeaders(v map[string]*string) *ExecutePipelineResponse {
	s.Headers = v
	return s
}

func (s *ExecutePipelineResponse) SetStatusCode(v int32) *ExecutePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecutePipelineResponse) SetBody(v *ExecutePipelineResponseBody) *ExecutePipelineResponse {
	s.Body = v
	return s
}

type GetDevopsOrganizationMembersRequest struct {
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s GetDevopsOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersRequest) SetOrgId(v string) *GetDevopsOrganizationMembersRequest {
	s.OrgId = &v
	return s
}

type GetDevopsOrganizationMembersResponseBody struct {
	ErrorCode  *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetDevopsOrganizationMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                             `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersResponseBody) SetErrorCode(v string) *GetDevopsOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBody) SetErrorMsg(v string) *GetDevopsOrganizationMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBody) SetObject(v []*GetDevopsOrganizationMembersResponseBodyObject) *GetDevopsOrganizationMembersResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBody) SetRequestId(v string) *GetDevopsOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBody) SetSuccessful(v bool) *GetDevopsOrganizationMembersResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsOrganizationMembersResponseBodyObject struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDevopsOrganizationMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetAvatarUrl(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetEmail(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetMemberId(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetName(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetPhone(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Phone = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetRole(v int32) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Role = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetUserId(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.UserId = &v
	return s
}

type GetDevopsOrganizationMembersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDevopsOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevopsOrganizationMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersResponse) SetHeaders(v map[string]*string) *GetDevopsOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *GetDevopsOrganizationMembersResponse) SetStatusCode(v int32) *GetDevopsOrganizationMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponse) SetBody(v *GetDevopsOrganizationMembersResponseBody) *GetDevopsOrganizationMembersResponse {
	s.Body = v
	return s
}

type GetDevopsProjectInfoRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDevopsProjectInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoRequest) SetOrgId(v string) *GetDevopsProjectInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetDevopsProjectInfoRequest) SetProjectId(v string) *GetDevopsProjectInfoRequest {
	s.ProjectId = &v
	return s
}

type GetDevopsProjectInfoResponseBody struct {
	ErrorCode  *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBody) SetErrorMsg(v string) *GetDevopsProjectInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBody) SetObject(v *GetDevopsProjectInfoResponseBodyObject) *GetDevopsProjectInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsProjectInfoResponseBody) SetRequestId(v string) *GetDevopsProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectInfoResponseBodyObject struct {
	Category            *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Created             *string `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId           *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Customfields        *string `json:"Customfields,omitempty" xml:"Customfields,omitempty"`
	DefaultCollectionId *string `json:"DefaultCollectionId,omitempty" xml:"DefaultCollectionId,omitempty"`
	DefaultRoleId       *string `json:"DefaultRoleId,omitempty" xml:"DefaultRoleId,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndDate             *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Id                  *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArchived          *bool   `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	IsDeleted           *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	IsSuspended         *bool   `json:"IsSuspended,omitempty" xml:"IsSuspended,omitempty"`
	IsTemplate          *bool   `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	Logo                *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	ModifierId          *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NextTaskUniqueId    *int32  `json:"NextTaskUniqueId,omitempty" xml:"NextTaskUniqueId,omitempty"`
	NormalType          *string `json:"NormalType,omitempty" xml:"NormalType,omitempty"`
	OrganizationId      *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Pinyin              *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Py                  *string `json:"Py,omitempty" xml:"Py,omitempty"`
	RootCollectionId    *string `json:"RootCollectionId,omitempty" xml:"RootCollectionId,omitempty"`
	SortMethod          *string `json:"SortMethod,omitempty" xml:"SortMethod,omitempty"`
	SourceId            *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StartDate           *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	UniqueIdPrefix      *string `json:"UniqueIdPrefix,omitempty" xml:"UniqueIdPrefix,omitempty"`
	Updated             *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visibility          *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetDevopsProjectInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCategory(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Category = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCustomfields(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Customfields = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDefaultCollectionId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.DefaultCollectionId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDefaultRoleId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.DefaultRoleId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDescription(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Description = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetEndDate(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.EndDate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsArchived(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsArchived = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsSuspended(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsSuspended = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsTemplate(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsTemplate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetLogo(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Logo = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetModifierId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.ModifierId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetName(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetNextTaskUniqueId(v int32) *GetDevopsProjectInfoResponseBodyObject {
	s.NextTaskUniqueId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetNormalType(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.NormalType = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetOrganizationId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetPinyin(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Pinyin = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetPy(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Py = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetRootCollectionId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.RootCollectionId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSortMethod(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SortMethod = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSourceId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SourceId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSourceType(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SourceType = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetStartDate(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetUniqueIdPrefix(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.UniqueIdPrefix = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetVisibility(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Visibility = &v
	return s
}

type GetDevopsProjectInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDevopsProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevopsProjectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoResponse) SetHeaders(v map[string]*string) *GetDevopsProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDevopsProjectInfoResponse) SetStatusCode(v int32) *GetDevopsProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevopsProjectInfoResponse) SetBody(v *GetDevopsProjectInfoResponseBody) *GetDevopsProjectInfoResponse {
	s.Body = v
	return s
}

type GetDevopsProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageToken *string `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDevopsProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersRequest) SetOrgId(v string) *GetDevopsProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *GetDevopsProjectMembersRequest) SetPageSize(v int32) *GetDevopsProjectMembersRequest {
	s.PageSize = &v
	return s
}

func (s *GetDevopsProjectMembersRequest) SetPageToken(v string) *GetDevopsProjectMembersRequest {
	s.PageToken = &v
	return s
}

func (s *GetDevopsProjectMembersRequest) SetProjectId(v string) *GetDevopsProjectMembersRequest {
	s.ProjectId = &v
	return s
}

type GetDevopsProjectMembersResponseBody struct {
	ErrorCode     *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg      *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	NextPageToken *string                                      `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	Object        []*GetDevopsProjectMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful    *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
	Total         *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetDevopsProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersResponseBody) SetErrorCode(v string) *GetDevopsProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetErrorMsg(v string) *GetDevopsProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetNextPageToken(v string) *GetDevopsProjectMembersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetObject(v []*GetDevopsProjectMembersResponseBodyObject) *GetDevopsProjectMembersResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetRequestId(v string) *GetDevopsProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetSuccessful(v bool) *GetDevopsProjectMembersResponseBody {
	s.Successful = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetTotal(v int32) *GetDevopsProjectMembersResponseBody {
	s.Total = &v
	return s
}

type GetDevopsProjectMembersResponseBodyObject struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDevopsProjectMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetAvatarUrl(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetEmail(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetMemberId(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetName(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetPhone(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.Phone = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetRole(v int32) *GetDevopsProjectMembersResponseBodyObject {
	s.Role = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetUserId(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.UserId = &v
	return s
}

type GetDevopsProjectMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDevopsProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevopsProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersResponse) SetHeaders(v map[string]*string) *GetDevopsProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *GetDevopsProjectMembersResponse) SetStatusCode(v int32) *GetDevopsProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevopsProjectMembersResponse) SetBody(v *GetDevopsProjectMembersResponseBody) *GetDevopsProjectMembersResponse {
	s.Body = v
	return s
}

type GetDevopsProjectSprintInfoRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	SprintId *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s GetDevopsProjectSprintInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoRequest) SetOrgId(v string) *GetDevopsProjectSprintInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetDevopsProjectSprintInfoRequest) SetSprintId(v string) *GetDevopsProjectSprintInfoRequest {
	s.SprintId = &v
	return s
}

type GetDevopsProjectSprintInfoResponseBody struct {
	ErrorCode  *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectSprintInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                         `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectSprintInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetErrorMsg(v string) *GetDevopsProjectSprintInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetObject(v *GetDevopsProjectSprintInfoResponseBodyObject) *GetDevopsProjectSprintInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetRequestId(v string) *GetDevopsProjectSprintInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectSprintInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectSprintInfoResponseBodyObject struct {
	Accomplished *string                                               `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	Created      *string                                               `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId    *string                                               `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DueDate      *string                                               `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Id           *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDeleted    *bool                                                 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Name         *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanToDo     *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	ProjectId    *string                                               `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                               `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status       *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Updated      *string                                               `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetAccomplished(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetDueDate(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetId(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetName(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetPlanToDo(v *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetProjectId(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetStartDate(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetStatus(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Updated = &v
	return s
}

type GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo struct {
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetTasks(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

type GetDevopsProjectSprintInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDevopsProjectSprintInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevopsProjectSprintInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponse) SetHeaders(v map[string]*string) *GetDevopsProjectSprintInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDevopsProjectSprintInfoResponse) SetStatusCode(v int32) *GetDevopsProjectSprintInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponse) SetBody(v *GetDevopsProjectSprintInfoResponseBody) *GetDevopsProjectSprintInfoResponse {
	s.Body = v
	return s
}

type GetDevopsProjectTaskInfoRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDevopsProjectTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoRequest) SetOrgId(v string) *GetDevopsProjectTaskInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoRequest) SetTaskId(v string) *GetDevopsProjectTaskInfoRequest {
	s.TaskId = &v
	return s
}

type GetDevopsProjectTaskInfoResponseBody struct {
	ErrorCode  *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectTaskInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                       `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectTaskInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetErrorMsg(v string) *GetDevopsProjectTaskInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetObject(v *GetDevopsProjectTaskInfoResponseBodyObject) *GetDevopsProjectTaskInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetRequestId(v string) *GetDevopsProjectTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectTaskInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectTaskInfoResponseBodyObject struct {
	Content          *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Created          *string   `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId        *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DueDate          *string   `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId       *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Id               *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	InvolveMembers   []*string `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	IsDeleted        *bool     `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	IsDone           *bool     `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	IsTopInProject   *bool     `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Note             *string   `json:"Note,omitempty" xml:"Note,omitempty"`
	OrganizationId   *string   `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Priority         *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId        *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SprintId         *string   `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate        *string   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint       *string   `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	TaskType         *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskflowstatusId *string   `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	TasklistId       *string   `json:"TasklistId,omitempty" xml:"TasklistId,omitempty"`
	Updated          *string   `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visible          *string   `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetDevopsProjectTaskInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetContent(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Content = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetDueDate(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetExecutorId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetInvolveMembers(v []*string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.InvolveMembers = v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsDone(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsTopInProject(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsTopInProject = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetNote(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Note = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetOrganizationId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetPriority(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetProjectId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetSprintId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetStartDate(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetStoryPoint(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTaskType(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTaskflowstatusId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTasklistId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetVisible(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Visible = &v
	return s
}

type GetDevopsProjectTaskInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDevopsProjectTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDevopsProjectTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoResponse) SetHeaders(v map[string]*string) *GetDevopsProjectTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDevopsProjectTaskInfoResponse) SetStatusCode(v int32) *GetDevopsProjectTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponse) SetBody(v *GetDevopsProjectTaskInfoResponseBody) *GetDevopsProjectTaskInfoResponse {
	s.Body = v
	return s
}

type GetPipelineInstHistoryRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart  *int64  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryRequest) SetEndTime(v string) *GetPipelineInstHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetOrgId(v string) *GetPipelineInstHistoryRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPageSize(v int64) *GetPipelineInstHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPageStart(v int64) *GetPipelineInstHistoryRequest {
	s.PageStart = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPipelineId(v int64) *GetPipelineInstHistoryRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetStartTime(v string) *GetPipelineInstHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetUserPk(v string) *GetPipelineInstHistoryRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstHistoryResponseBody struct {
	Data         *GetPipelineInstHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBody) SetData(v *GetPipelineInstHistoryResponseBodyData) *GetPipelineInstHistoryResponseBody {
	s.Data = v
	return s
}

func (s *GetPipelineInstHistoryResponseBody) SetErrorCode(v string) *GetPipelineInstHistoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBody) SetErrorMessage(v string) *GetPipelineInstHistoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBody) SetRequestId(v string) *GetPipelineInstHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBody) SetSuccess(v bool) *GetPipelineInstHistoryResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstHistoryResponseBodyData struct {
	DataList []*GetPipelineInstHistoryResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Total    *int32                                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyData) SetDataList(v []*GetPipelineInstHistoryResponseBodyDataDataList) *GetPipelineInstHistoryResponseBodyData {
	s.DataList = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyData) SetTotal(v int32) *GetPipelineInstHistoryResponseBodyData {
	s.Total = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataList struct {
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator          *string                                                     `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Deletion         *string                                                     `json:"Deletion,omitempty" xml:"Deletion,omitempty"`
	FlowInstId       *int32                                                      `json:"FlowInstId,omitempty" xml:"FlowInstId,omitempty"`
	FlowInstance     *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance `json:"FlowInstance,omitempty" xml:"FlowInstance,omitempty" type:"Struct"`
	Id               *int32                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	InstNumber       *int32                                                      `json:"InstNumber,omitempty" xml:"InstNumber,omitempty"`
	Modifier         *string                                                     `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyTime       *int64                                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Packages         *string                                                     `json:"Packages,omitempty" xml:"Packages,omitempty"`
	PipelineConfigId *int32                                                      `json:"PipelineConfigId,omitempty" xml:"PipelineConfigId,omitempty"`
	PipelineId       *int32                                                      `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Status           *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName       *string                                                     `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	TriggerMode      *int32                                                      `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Creator = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetDeletion(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Deletion = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetFlowInstId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.FlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetFlowInstance(v *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.FlowInstance = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetInstNumber(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.InstNumber = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetModifier(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Modifier = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.ModifyTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPackages(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Packages = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPipelineConfigId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.PipelineConfigId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPipelineId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetStatusName(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.StatusName = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetTriggerMode(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.TriggerMode = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstance struct {
	AutoDrivenStatus *bool                                                               `json:"AutoDrivenStatus,omitempty" xml:"AutoDrivenStatus,omitempty"`
	CreateTime       *int64                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator          *string                                                             `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Groups           []*GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Id               *int32                                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Modifier         *string                                                             `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyTime       *int64                                                              `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Result           *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult   `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	ResultStatus     *string                                                             `json:"ResultStatus,omitempty" xml:"ResultStatus,omitempty"`
	RunningStatus    *string                                                             `json:"RunningStatus,omitempty" xml:"RunningStatus,omitempty"`
	StageTopo        *string                                                             `json:"StageTopo,omitempty" xml:"StageTopo,omitempty"`
	Stages           map[string]interface{}                                              `json:"Stages,omitempty" xml:"Stages,omitempty"`
	Status           *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName       *string                                                             `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	SystemCode       *string                                                             `json:"SystemCode,omitempty" xml:"SystemCode,omitempty"`
	SystemId         *string                                                             `json:"SystemId,omitempty" xml:"SystemId,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetAutoDrivenStatus(v bool) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.AutoDrivenStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Creator = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetGroups(v []*GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Groups = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetModifier(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Modifier = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.ModifyTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetResult(v *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Result = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetResultStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.ResultStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetRunningStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.RunningStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStageTopo(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.StageTopo = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStages(v map[string]interface{}) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Stages = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Status = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStatusName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.StatusName = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetSystemCode(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.SystemCode = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetSystemId(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.SystemId = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups struct {
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeleteStatus *string `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FlowInstId   *int32  `json:"FlowInstId,omitempty" xml:"FlowInstId,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdExtent     *int32  `json:"IdExtent,omitempty" xml:"IdExtent,omitempty"`
	Modifier     *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ModifyTime   *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResultStatus *string `json:"ResultStatus,omitempty" xml:"ResultStatus,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Creator = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetDeleteStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.DeleteStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetEndTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetFlowInstId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.FlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetIdExtent(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.IdExtent = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetModifier(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Modifier = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.ModifyTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Name = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetResultStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.ResultStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetStartTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Status = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult struct {
	Caches               *string `json:"Caches,omitempty" xml:"Caches,omitempty"`
	DateTime             *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	EnginePipelineId     *int32  `json:"EnginePipelineId,omitempty" xml:"EnginePipelineId,omitempty"`
	EnginePipelineInstId *int32  `json:"EnginePipelineInstId,omitempty" xml:"EnginePipelineInstId,omitempty"`
	EnginePipelineName   *string `json:"EnginePipelineName,omitempty" xml:"EnginePipelineName,omitempty"`
	EnginePipelineNumber *int32  `json:"EnginePipelineNumber,omitempty" xml:"EnginePipelineNumber,omitempty"`
	MixFlowInstId        *string `json:"MixFlowInstId,omitempty" xml:"MixFlowInstId,omitempty"`
	Sources              *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TriggerMode          *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetCaches(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.Caches = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetDateTime(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.DateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineInstId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineName = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineNumber(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineNumber = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetMixFlowInstId(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.MixFlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetSources(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.Sources = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetTimeStamp(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.TimeStamp = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetTriggerMode(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.TriggerMode = &v
	return s
}

type GetPipelineInstHistoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineInstHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineInstHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponse) SetHeaders(v map[string]*string) *GetPipelineInstHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineInstHistoryResponse) SetStatusCode(v int32) *GetPipelineInstHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineInstHistoryResponse) SetBody(v *GetPipelineInstHistoryResponseBody) *GetPipelineInstHistoryResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceBuildNumberStatusRequest struct {
	BuildNum   *int64  `json:"BuildNum,omitempty" xml:"BuildNum,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusRequest) SetBuildNum(v int64) *GetPipelineInstanceBuildNumberStatusRequest {
	s.BuildNum = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusRequest) SetOrgId(v string) *GetPipelineInstanceBuildNumberStatusRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusRequest) SetPipelineId(v int64) *GetPipelineInstanceBuildNumberStatusRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusRequest) SetUserPk(v string) *GetPipelineInstanceBuildNumberStatusRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBody struct {
	ErrorCode    *string                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *GetPipelineInstanceBuildNumberStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetObject(v *GetPipelineInstanceBuildNumberStatusResponseBodyObject) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObject struct {
	Groups []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Status *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObject) SetGroups(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) *GetPipelineInstanceBuildNumberStatusResponseBodyObject {
	s.Groups = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObject) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObject {
	s.Status = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups struct {
	Name   *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Stages []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Status *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetName(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetStages(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages struct {
	Components []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Sign       *string                                                                         `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Status     *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents struct {
	JobId  *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v int64) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineInstanceBuildNumberStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineInstanceBuildNumberStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponse) SetHeaders(v map[string]*string) *GetPipelineInstanceBuildNumberStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponse) SetStatusCode(v int32) *GetPipelineInstanceBuildNumberStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponse) SetBody(v *GetPipelineInstanceBuildNumberStatusResponseBody) *GetPipelineInstanceBuildNumberStatusResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceGroupStatusRequest struct {
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusRequest) SetFlowInstanceId(v int64) *GetPipelineInstanceGroupStatusRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusRequest) SetOrgId(v string) *GetPipelineInstanceGroupStatusRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusRequest) SetPipelineId(v int64) *GetPipelineInstanceGroupStatusRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusRequest) SetUserPk(v string) *GetPipelineInstanceGroupStatusRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBody struct {
	ErrorCode    *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *GetPipelineInstanceGroupStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetObject(v *GetPipelineInstanceGroupStatusResponseBodyObject) *GetPipelineInstanceGroupStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceGroupStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObject struct {
	Groups []*GetPipelineInstanceGroupStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Status *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObject) SetGroups(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroups) *GetPipelineInstanceGroupStatusResponseBodyObject {
	s.Groups = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObject) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObject {
	s.Status = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroups struct {
	Name   *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Stages []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Status *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetName(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetStages(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages struct {
	Components []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Sign       *string                                                                   `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Status     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

type GetPipelineInstanceGroupStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineInstanceGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineInstanceGroupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponse) SetHeaders(v map[string]*string) *GetPipelineInstanceGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponse) SetStatusCode(v int32) *GetPipelineInstanceGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponse) SetBody(v *GetPipelineInstanceGroupStatusResponseBody) *GetPipelineInstanceGroupStatusResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceInfoRequest struct {
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoRequest) SetFlowInstanceId(v string) *GetPipelineInstanceInfoRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetOrgId(v string) *GetPipelineInstanceInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetPipelineId(v int64) *GetPipelineInstanceInfoRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetUserPk(v string) *GetPipelineInstanceInfoRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceInfoResponseBody struct {
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *GetPipelineInstanceInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoResponseBody) SetErrorCode(v string) *GetPipelineInstanceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetErrorMessage(v string) *GetPipelineInstanceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetObject(v *GetPipelineInstanceInfoResponseBodyObject) *GetPipelineInstanceInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetRequestId(v string) *GetPipelineInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetSuccess(v bool) *GetPipelineInstanceInfoResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceInfoResponseBodyObject struct {
	DockerImages        []*string `json:"DockerImages,omitempty" xml:"DockerImages,omitempty" type:"Repeated"`
	EmployeeId          *string   `json:"EmployeeId,omitempty" xml:"EmployeeId,omitempty"`
	EndTime             *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PackageDownloadUrls []*string `json:"PackageDownloadUrls,omitempty" xml:"PackageDownloadUrls,omitempty" type:"Repeated"`
	Sources             *string   `json:"Sources,omitempty" xml:"Sources,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineInstanceInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetDockerImages(v []*string) *GetPipelineInstanceInfoResponseBodyObject {
	s.DockerImages = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetEmployeeId(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.EmployeeId = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetEndTime(v int64) *GetPipelineInstanceInfoResponseBodyObject {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetPackageDownloadUrls(v []*string) *GetPipelineInstanceInfoResponseBodyObject {
	s.PackageDownloadUrls = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetSources(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.Sources = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetStartTime(v int64) *GetPipelineInstanceInfoResponseBodyObject {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetStatus(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.Status = &v
	return s
}

type GetPipelineInstanceInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineInstanceInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoResponse) SetHeaders(v map[string]*string) *GetPipelineInstanceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineInstanceInfoResponse) SetStatusCode(v int32) *GetPipelineInstanceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineInstanceInfoResponse) SetBody(v *GetPipelineInstanceInfoResponseBody) *GetPipelineInstanceInfoResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceStatusRequest struct {
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceStatusRequest) SetFlowInstanceId(v int64) *GetPipelineInstanceStatusRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetOrgId(v string) *GetPipelineInstanceStatusRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetPipelineId(v int64) *GetPipelineInstanceStatusRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetUserPk(v string) *GetPipelineInstanceStatusRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceStatusResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetObject(v string) *GetPipelineInstanceStatusResponseBody {
	s.Object = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceStatusResponse) SetHeaders(v map[string]*string) *GetPipelineInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineInstanceStatusResponse) SetStatusCode(v int32) *GetPipelineInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineInstanceStatusResponse) SetBody(v *GetPipelineInstanceStatusResponseBody) *GetPipelineInstanceStatusResponse {
	s.Body = v
	return s
}

type GetPipelineLogRequest struct {
	JobId      *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineLogRequest) SetJobId(v int64) *GetPipelineLogRequest {
	s.JobId = &v
	return s
}

func (s *GetPipelineLogRequest) SetOrgId(v string) *GetPipelineLogRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineLogRequest) SetPipelineId(v int64) *GetPipelineLogRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineLogRequest) SetUserPk(v string) *GetPipelineLogRequest {
	s.UserPk = &v
	return s
}

type GetPipelineLogResponseBody struct {
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       []*GetPipelineLogResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBody) SetErrorCode(v string) *GetPipelineLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetErrorMessage(v string) *GetPipelineLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetObject(v []*GetPipelineLogResponseBodyObject) *GetPipelineLogResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineLogResponseBody) SetRequestId(v string) *GetPipelineLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetSuccess(v bool) *GetPipelineLogResponseBody {
	s.Success = &v
	return s
}

type GetPipelineLogResponseBodyObject struct {
	ActionName        *string                                              `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	BuildProcessNodes []*GetPipelineLogResponseBodyObjectBuildProcessNodes `json:"BuildProcessNodes,omitempty" xml:"BuildProcessNodes,omitempty" type:"Repeated"`
	JobId             *int64                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	StartTime         *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPipelineLogResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBodyObject) SetActionName(v string) *GetPipelineLogResponseBodyObject {
	s.ActionName = &v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetBuildProcessNodes(v []*GetPipelineLogResponseBodyObjectBuildProcessNodes) *GetPipelineLogResponseBodyObject {
	s.BuildProcessNodes = v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetJobId(v int64) *GetPipelineLogResponseBodyObject {
	s.JobId = &v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetStartTime(v string) *GetPipelineLogResponseBodyObject {
	s.StartTime = &v
	return s
}

type GetPipelineLogResponseBodyObjectBuildProcessNodes struct {
	NodeIndex *int32  `json:"NodeIndex,omitempty" xml:"NodeIndex,omitempty"`
	NodeName  *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipelineLogResponseBodyObjectBuildProcessNodes) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBodyObjectBuildProcessNodes) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetNodeIndex(v int32) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.NodeIndex = &v
	return s
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetNodeName(v string) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.NodeName = &v
	return s
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetStatus(v string) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.Status = &v
	return s
}

type GetPipelineLogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponse) SetHeaders(v map[string]*string) *GetPipelineLogResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineLogResponse) SetStatusCode(v int32) *GetPipelineLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineLogResponse) SetBody(v *GetPipelineLogResponseBody) *GetPipelineLogResponse {
	s.Body = v
	return s
}

type GetPipelineStepLogRequest struct {
	JobId      *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Limit      *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Offset     *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	StepIndex  *string `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineStepLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogRequest) SetJobId(v int64) *GetPipelineStepLogRequest {
	s.JobId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetLimit(v int64) *GetPipelineStepLogRequest {
	s.Limit = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetOffset(v int64) *GetPipelineStepLogRequest {
	s.Offset = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetOrgId(v string) *GetPipelineStepLogRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetPipelineId(v int64) *GetPipelineStepLogRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetStepIndex(v string) *GetPipelineStepLogRequest {
	s.StepIndex = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetUserPk(v string) *GetPipelineStepLogRequest {
	s.UserPk = &v
	return s
}

type GetPipelineStepLogResponseBody struct {
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *GetPipelineStepLogResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineStepLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogResponseBody) SetErrorCode(v string) *GetPipelineStepLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetErrorMessage(v string) *GetPipelineStepLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetObject(v *GetPipelineStepLogResponseBodyObject) *GetPipelineStepLogResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetRequestId(v string) *GetPipelineStepLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetSuccess(v bool) *GetPipelineStepLogResponseBody {
	s.Success = &v
	return s
}

type GetPipelineStepLogResponseBodyObject struct {
	Last *int32  `json:"Last,omitempty" xml:"Last,omitempty"`
	Logs *string `json:"Logs,omitempty" xml:"Logs,omitempty"`
	More *bool   `json:"More,omitempty" xml:"More,omitempty"`
}

func (s GetPipelineStepLogResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogResponseBodyObject) SetLast(v int32) *GetPipelineStepLogResponseBodyObject {
	s.Last = &v
	return s
}

func (s *GetPipelineStepLogResponseBodyObject) SetLogs(v string) *GetPipelineStepLogResponseBodyObject {
	s.Logs = &v
	return s
}

func (s *GetPipelineStepLogResponseBodyObject) SetMore(v bool) *GetPipelineStepLogResponseBodyObject {
	s.More = &v
	return s
}

type GetPipelineStepLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipelineStepLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipelineStepLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogResponse) SetHeaders(v map[string]*string) *GetPipelineStepLogResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineStepLogResponse) SetStatusCode(v int32) *GetPipelineStepLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineStepLogResponse) SetBody(v *GetPipelineStepLogResponseBody) *GetPipelineStepLogResponse {
	s.Body = v
	return s
}

type GetPipleineLatestInstanceStatusRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipleineLatestInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusRequest) SetOrgId(v string) *GetPipleineLatestInstanceStatusRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusRequest) SetPipelineId(v int64) *GetPipleineLatestInstanceStatusRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusRequest) SetUserPk(v string) *GetPipleineLatestInstanceStatusRequest {
	s.UserPk = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBody struct {
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *GetPipleineLatestInstanceStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetErrorCode(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetErrorMessage(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetObject(v *GetPipleineLatestInstanceStatusResponseBodyObject) *GetPipleineLatestInstanceStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetRequestId(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetSuccess(v bool) *GetPipleineLatestInstanceStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObject struct {
	Groups []*GetPipleineLatestInstanceStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	Status *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObject) SetGroups(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroups) *GetPipleineLatestInstanceStatusResponseBodyObject {
	s.Groups = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObject) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObject {
	s.Status = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroups struct {
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Stages []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Status *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetName(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetStages(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages struct {
	Components []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	Sign       *string                                                                    `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Status     *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents struct {
	JobId  *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v int64) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

type GetPipleineLatestInstanceStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPipleineLatestInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPipleineLatestInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponse) SetHeaders(v map[string]*string) *GetPipleineLatestInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponse) SetStatusCode(v int32) *GetPipleineLatestInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponse) SetBody(v *GetPipleineLatestInstanceStatusResponseBody) *GetPipleineLatestInstanceStatusResponse {
	s.Body = v
	return s
}

type GetProjectOptionRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProjectOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionRequest) GoString() string {
	return s.String()
}

func (s *GetProjectOptionRequest) SetOrgId(v string) *GetProjectOptionRequest {
	s.OrgId = &v
	return s
}

func (s *GetProjectOptionRequest) SetProjectId(v string) *GetProjectOptionRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectOptionRequest) SetQuery(v string) *GetProjectOptionRequest {
	s.Query = &v
	return s
}

func (s *GetProjectOptionRequest) SetType(v string) *GetProjectOptionRequest {
	s.Type = &v
	return s
}

type GetProjectOptionResponseBody struct {
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetProjectOptionResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectOptionResponseBody) SetErrorCode(v string) *GetProjectOptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectOptionResponseBody) SetErrorMsg(v string) *GetProjectOptionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectOptionResponseBody) SetObject(v []*GetProjectOptionResponseBodyObject) *GetProjectOptionResponseBody {
	s.Object = v
	return s
}

func (s *GetProjectOptionResponseBody) SetRequestId(v string) *GetProjectOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectOptionResponseBody) SetSuccessful(v bool) *GetProjectOptionResponseBody {
	s.Successful = &v
	return s
}

type GetProjectOptionResponseBodyObject struct {
	Kind      *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScopeName *string `json:"ScopeName,omitempty" xml:"ScopeName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProjectOptionResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectOptionResponseBodyObject) SetKind(v string) *GetProjectOptionResponseBodyObject {
	s.Kind = &v
	return s
}

func (s *GetProjectOptionResponseBodyObject) SetName(v string) *GetProjectOptionResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetProjectOptionResponseBodyObject) SetScopeName(v string) *GetProjectOptionResponseBodyObject {
	s.ScopeName = &v
	return s
}

func (s *GetProjectOptionResponseBodyObject) SetValue(v string) *GetProjectOptionResponseBodyObject {
	s.Value = &v
	return s
}

type GetProjectOptionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionResponse) GoString() string {
	return s.String()
}

func (s *GetProjectOptionResponse) SetHeaders(v map[string]*string) *GetProjectOptionResponse {
	s.Headers = v
	return s
}

func (s *GetProjectOptionResponse) SetStatusCode(v int32) *GetProjectOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectOptionResponse) SetBody(v *GetProjectOptionResponseBody) *GetProjectOptionResponse {
	s.Body = v
	return s
}

type GetTaskDetailActivityRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskDetailActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailActivityRequest) GoString() string {
	return s.String()
}

func (s *GetTaskDetailActivityRequest) SetOrgId(v string) *GetTaskDetailActivityRequest {
	s.OrgId = &v
	return s
}

func (s *GetTaskDetailActivityRequest) SetProjectId(v string) *GetTaskDetailActivityRequest {
	s.ProjectId = &v
	return s
}

func (s *GetTaskDetailActivityRequest) SetTaskId(v string) *GetTaskDetailActivityRequest {
	s.TaskId = &v
	return s
}

type GetTaskDetailActivityResponseBody struct {
	ErrorCode      *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg       *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Object         []*GetTaskDetailActivityResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful     *bool                                      `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskDetailActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailActivityResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskDetailActivityResponseBody) SetErrorCode(v string) *GetTaskDetailActivityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetErrorMsg(v string) *GetTaskDetailActivityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetHttpStatusCode(v int32) *GetTaskDetailActivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetObject(v []*GetTaskDetailActivityResponseBodyObject) *GetTaskDetailActivityResponseBody {
	s.Object = v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetRequestId(v string) *GetTaskDetailActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetSuccessful(v bool) *GetTaskDetailActivityResponseBody {
	s.Successful = &v
	return s
}

type GetTaskDetailActivityResponseBodyObject struct {
	Action  *string                `json:"Action,omitempty" xml:"Action,omitempty"`
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Created *string                `json:"Created,omitempty" xml:"Created,omitempty"`
	Title   *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	Updated *string                `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s GetTaskDetailActivityResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailActivityResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetTaskDetailActivityResponseBodyObject) SetAction(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Action = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetContent(v map[string]interface{}) *GetTaskDetailActivityResponseBodyObject {
	s.Content = v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetCreated(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetTitle(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Title = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetUpdated(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Updated = &v
	return s
}

type GetTaskDetailActivityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskDetailActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskDetailActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailActivityResponse) GoString() string {
	return s.String()
}

func (s *GetTaskDetailActivityResponse) SetHeaders(v map[string]*string) *GetTaskDetailActivityResponse {
	s.Headers = v
	return s
}

func (s *GetTaskDetailActivityResponse) SetStatusCode(v int32) *GetTaskDetailActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskDetailActivityResponse) SetBody(v *GetTaskDetailActivityResponseBody) *GetTaskDetailActivityResponse {
	s.Body = v
	return s
}

type GetTaskDetailBaseRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskDetailBaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseRequest) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseRequest) SetOrgId(v string) *GetTaskDetailBaseRequest {
	s.OrgId = &v
	return s
}

func (s *GetTaskDetailBaseRequest) SetProjectId(v string) *GetTaskDetailBaseRequest {
	s.ProjectId = &v
	return s
}

func (s *GetTaskDetailBaseRequest) SetTaskId(v string) *GetTaskDetailBaseRequest {
	s.TaskId = &v
	return s
}

type GetTaskDetailBaseResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetTaskDetailBaseResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskDetailBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBody) SetErrorCode(v string) *GetTaskDetailBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskDetailBaseResponseBody) SetErrorMsg(v string) *GetTaskDetailBaseResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTaskDetailBaseResponseBody) SetObject(v *GetTaskDetailBaseResponseBodyObject) *GetTaskDetailBaseResponseBody {
	s.Object = v
	return s
}

func (s *GetTaskDetailBaseResponseBody) SetRequestId(v string) *GetTaskDetailBaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBody) SetSuccessful(v bool) *GetTaskDetailBaseResponseBody {
	s.Successful = &v
	return s
}

type GetTaskDetailBaseResponseBodyObject struct {
	Accomplished          *string                                                 `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	AncestorIds           []*string                                               `json:"AncestorIds,omitempty" xml:"AncestorIds,omitempty" type:"Repeated"`
	Ancestors             []*string                                               `json:"Ancestors,omitempty" xml:"Ancestors,omitempty" type:"Repeated"`
	AttachmentsCount      *int32                                                  `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	Badges                *GetTaskDetailBaseResponseBodyObjectBadges              `json:"Badges,omitempty" xml:"Badges,omitempty" type:"Struct"`
	CommentsCount         *int32                                                  `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	Content               *string                                                 `json:"Content,omitempty" xml:"Content,omitempty"`
	Created               *string                                                 `json:"Created,omitempty" xml:"Created,omitempty"`
	Creator               *GetTaskDetailBaseResponseBodyObjectCreator             `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	CreatorId             *string                                                 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Customfields          []*GetTaskDetailBaseResponseBodyObjectCustomfields      `json:"Customfields,omitempty" xml:"Customfields,omitempty" type:"Repeated"`
	Divisions             []*string                                               `json:"Divisions,omitempty" xml:"Divisions,omitempty" type:"Repeated"`
	DueDate               *string                                                 `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Executor              *GetTaskDetailBaseResponseBodyObjectExecutor            `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	ExecutorId            *string                                                 `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Id                    *string                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	InvolveMembers        []*string                                               `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	Involvers             []*GetTaskDetailBaseResponseBodyObjectInvolvers         `json:"Involvers,omitempty" xml:"Involvers,omitempty" type:"Repeated"`
	IsArchived            *bool                                                   `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	IsDone                *bool                                                   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	IsFavorite            *bool                                                   `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	IsTopInProject        *bool                                                   `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Labels                []*string                                               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LikesCount            *int32                                                  `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	Note                  *string                                                 `json:"Note,omitempty" xml:"Note,omitempty"`
	ObjectType            *string                                                 `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ObjectlinksCount      *int32                                                  `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	Organization          *string                                                 `json:"Organization,omitempty" xml:"Organization,omitempty"`
	OrganizationId        *string                                                 `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Parent                *string                                                 `json:"Parent,omitempty" xml:"Parent,omitempty"`
	Priority              *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Progress              *int32                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProjectId             *string                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Rating                *int32                                                  `json:"Rating,omitempty" xml:"Rating,omitempty"`
	Recurrence            *string                                                 `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	Reminder              *GetTaskDetailBaseResponseBodyObjectReminder            `json:"Reminder,omitempty" xml:"Reminder,omitempty" type:"Struct"`
	Scenariofieldconfig   *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig `json:"Scenariofieldconfig,omitempty" xml:"Scenariofieldconfig,omitempty" type:"Struct"`
	ScenariofieldconfigId *string                                                 `json:"ScenariofieldconfigId,omitempty" xml:"ScenariofieldconfigId,omitempty"`
	ShareStatus           *int32                                                  `json:"ShareStatus,omitempty" xml:"ShareStatus,omitempty"`
	Source                *string                                                 `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceDate            *string                                                 `json:"SourceDate,omitempty" xml:"SourceDate,omitempty"`
	SourceId              *string                                                 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Sprint                *string                                                 `json:"Sprint,omitempty" xml:"Sprint,omitempty"`
	SprintId              *string                                                 `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	Stage                 *GetTaskDetailBaseResponseBodyObjectStage               `json:"Stage,omitempty" xml:"Stage,omitempty" type:"Struct"`
	StageId               *string                                                 `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StartDate             *string                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint            *string                                                 `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	SubtaskCount          *GetTaskDetailBaseResponseBodyObjectSubtaskCount        `json:"SubtaskCount,omitempty" xml:"SubtaskCount,omitempty" type:"Struct"`
	Subtasks              []*GetTaskDetailBaseResponseBodyObjectSubtasks          `json:"Subtasks,omitempty" xml:"Subtasks,omitempty" type:"Repeated"`
	TagIds                []*string                                               `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	TaskId                *string                                                 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Taskflowstatus        *GetTaskDetailBaseResponseBodyObjectTaskflowstatus      `json:"Taskflowstatus,omitempty" xml:"Taskflowstatus,omitempty" type:"Struct"`
	TaskflowstatusId      *string                                                 `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	Tasklist              *GetTaskDetailBaseResponseBodyObjectTasklist            `json:"Tasklist,omitempty" xml:"Tasklist,omitempty" type:"Struct"`
	UniqueId              *int32                                                  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	UntilDate             *string                                                 `json:"UntilDate,omitempty" xml:"UntilDate,omitempty"`
	Updated               *string                                                 `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visible               *string                                                 `json:"Visible,omitempty" xml:"Visible,omitempty"`
	WorkTime              *GetTaskDetailBaseResponseBodyObjectWorkTime            `json:"WorkTime,omitempty" xml:"WorkTime,omitempty" type:"Struct"`
}

func (s GetTaskDetailBaseResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAccomplished(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAncestorIds(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.AncestorIds = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAncestors(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Ancestors = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAttachmentsCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetBadges(v *GetTaskDetailBaseResponseBodyObjectBadges) *GetTaskDetailBaseResponseBodyObject {
	s.Badges = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCommentsCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetContent(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Content = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreated(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreator(v *GetTaskDetailBaseResponseBodyObjectCreator) *GetTaskDetailBaseResponseBodyObject {
	s.Creator = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreatorId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCustomfields(v []*GetTaskDetailBaseResponseBodyObjectCustomfields) *GetTaskDetailBaseResponseBodyObject {
	s.Customfields = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetDivisions(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Divisions = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetDueDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetExecutor(v *GetTaskDetailBaseResponseBodyObjectExecutor) *GetTaskDetailBaseResponseBodyObject {
	s.Executor = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetExecutorId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetInvolveMembers(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.InvolveMembers = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetInvolvers(v []*GetTaskDetailBaseResponseBodyObjectInvolvers) *GetTaskDetailBaseResponseBodyObject {
	s.Involvers = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsArchived(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsArchived = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsDone(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsFavorite(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsFavorite = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsTopInProject(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsTopInProject = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetLabels(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Labels = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetLikesCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.LikesCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetNote(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Note = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetObjectType(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ObjectType = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetObjectlinksCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetOrganization(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Organization = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetOrganizationId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetParent(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Parent = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetPriority(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetProgress(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Progress = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetProjectId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetRating(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Rating = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetRecurrence(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Recurrence = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetReminder(v *GetTaskDetailBaseResponseBodyObjectReminder) *GetTaskDetailBaseResponseBodyObject {
	s.Reminder = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetScenariofieldconfig(v *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) *GetTaskDetailBaseResponseBodyObject {
	s.Scenariofieldconfig = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetScenariofieldconfigId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetShareStatus(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.ShareStatus = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSource(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Source = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSourceDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SourceDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSourceId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SourceId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSprint(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Sprint = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSprintId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStage(v *GetTaskDetailBaseResponseBodyObjectStage) *GetTaskDetailBaseResponseBodyObject {
	s.Stage = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStageId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StageId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStartDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStoryPoint(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSubtaskCount(v *GetTaskDetailBaseResponseBodyObjectSubtaskCount) *GetTaskDetailBaseResponseBodyObject {
	s.SubtaskCount = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSubtasks(v []*GetTaskDetailBaseResponseBodyObjectSubtasks) *GetTaskDetailBaseResponseBodyObject {
	s.Subtasks = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTagIds(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.TagIds = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTaskId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.TaskId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTaskflowstatus(v *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) *GetTaskDetailBaseResponseBodyObject {
	s.Taskflowstatus = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTaskflowstatusId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTasklist(v *GetTaskDetailBaseResponseBodyObjectTasklist) *GetTaskDetailBaseResponseBodyObject {
	s.Tasklist = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUniqueId(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.UniqueId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUntilDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.UntilDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUpdated(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetVisible(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetWorkTime(v *GetTaskDetailBaseResponseBodyObjectWorkTime) *GetTaskDetailBaseResponseBodyObject {
	s.WorkTime = v
	return s
}

type GetTaskDetailBaseResponseBodyObjectBadges struct {
	AttachmentsCount *int32 `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	CommentsCount    *int32 `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	LikesCount       *int32 `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	ObjectlinksCount *int32 `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectBadges) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectBadges) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetAttachmentsCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetCommentsCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetLikesCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.LikesCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetObjectlinksCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.ObjectlinksCount = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectCreator struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectCreator) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCreator) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectCreator) SetId(v string) *GetTaskDetailBaseResponseBodyObjectCreator {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCreator) SetName(v string) *GetTaskDetailBaseResponseBodyObjectCreator {
	s.Name = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectCustomfields struct {
	CustomfieldId *string                                                 `json:"CustomfieldId,omitempty" xml:"CustomfieldId,omitempty"`
	Type          *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         []*GetTaskDetailBaseResponseBodyObjectCustomfieldsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Values        []*string                                               `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfields) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfields) SetCustomfieldId(v string) *GetTaskDetailBaseResponseBodyObjectCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfields) SetType(v string) *GetTaskDetailBaseResponseBodyObjectCustomfields {
	s.Type = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfields) SetValue(v []*GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) *GetTaskDetailBaseResponseBodyObjectCustomfields {
	s.Value = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfields) SetValues(v []*string) *GetTaskDetailBaseResponseBodyObjectCustomfields {
	s.Values = v
	return s
}

type GetTaskDetailBaseResponseBodyObjectCustomfieldsValue struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) SetId(v string) *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) SetTitle(v string) *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue {
	s.Title = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectExecutor struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectExecutor) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectExecutor) SetAvatarUrl(v string) *GetTaskDetailBaseResponseBodyObjectExecutor {
	s.AvatarUrl = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectExecutor) SetId(v string) *GetTaskDetailBaseResponseBodyObjectExecutor {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectExecutor) SetName(v string) *GetTaskDetailBaseResponseBodyObjectExecutor {
	s.Name = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectInvolvers struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectInvolvers) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectInvolvers) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectInvolvers) SetId(v string) *GetTaskDetailBaseResponseBodyObjectInvolvers {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectInvolvers) SetName(v string) *GetTaskDetailBaseResponseBodyObjectInvolvers {
	s.Name = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectReminder struct {
	CreatorId   *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Date        *string   `json:"Date,omitempty" xml:"Date,omitempty"`
	MemberRoles []*string `json:"MemberRoles,omitempty" xml:"MemberRoles,omitempty" type:"Repeated"`
	Members     []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Method      *string   `json:"Method,omitempty" xml:"Method,omitempty"`
	Rules       []*string `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Type        *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectReminder) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectReminder) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetCreatorId(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.CreatorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetDate(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Date = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetMemberRoles(v []*string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.MemberRoles = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetMembers(v []*string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Members = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetMethod(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Method = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetRules(v []*string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Rules = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetType(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Type = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectScenariofieldconfig struct {
	Icon                  *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ProTemplateConfigType *string `json:"ProTemplateConfigType,omitempty" xml:"ProTemplateConfigType,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetIcon(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.Icon = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetId(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetName(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetProTemplateConfigType(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.ProTemplateConfigType = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectStage struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectStage) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectStage) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectStage) SetId(v string) *GetTaskDetailBaseResponseBodyObjectStage {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectStage) SetName(v string) *GetTaskDetailBaseResponseBodyObjectStage {
	s.Name = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectSubtaskCount struct {
	Done  *int32 `json:"Done,omitempty" xml:"Done,omitempty"`
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectSubtaskCount) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectSubtaskCount) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectSubtaskCount) SetDone(v int32) *GetTaskDetailBaseResponseBodyObjectSubtaskCount {
	s.Done = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectSubtaskCount) SetTotal(v int32) *GetTaskDetailBaseResponseBodyObjectSubtaskCount {
	s.Total = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectSubtasks struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectSubtasks) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectSubtasks) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectSubtasks) SetContent(v string) *GetTaskDetailBaseResponseBodyObjectSubtasks {
	s.Content = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectSubtasks) SetId(v string) *GetTaskDetailBaseResponseBodyObjectSubtasks {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectTaskflowstatus struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Kind       *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskflowId *string `json:"TaskflowId,omitempty" xml:"TaskflowId,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectTaskflowstatus) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectTaskflowstatus) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetId(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetKind(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.Kind = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetName(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetTaskflowId(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.TaskflowId = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectTasklist struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectTasklist) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectTasklist) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectTasklist) SetId(v string) *GetTaskDetailBaseResponseBodyObjectTasklist {
	s.Id = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectTasklist) SetTitle(v string) *GetTaskDetailBaseResponseBodyObjectTasklist {
	s.Title = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectWorkTime struct {
	TotalTime *int32  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Unit      *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UsedTime  *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectWorkTime) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectWorkTime) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetTotalTime(v int32) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.TotalTime = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetUnit(v string) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.Unit = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetUsedTime(v int32) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.UsedTime = &v
	return s
}

type GetTaskDetailBaseResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskDetailBaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskDetailBaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponse) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponse) SetHeaders(v map[string]*string) *GetTaskDetailBaseResponse {
	s.Headers = v
	return s
}

func (s *GetTaskDetailBaseResponse) SetStatusCode(v int32) *GetTaskDetailBaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskDetailBaseResponse) SetBody(v *GetTaskDetailBaseResponseBody) *GetTaskDetailBaseResponse {
	s.Body = v
	return s
}

type GetTaskListFilterRequest struct {
	CreatorId             *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DueDateEnd            *string `json:"DueDateEnd,omitempty" xml:"DueDateEnd,omitempty"`
	DueDateStart          *string `json:"DueDateStart,omitempty" xml:"DueDateStart,omitempty"`
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InvolveMembers        *string `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty"`
	IsDone                *bool   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectType            *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Order                 *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OrderCondition        *string `json:"OrderCondition,omitempty" xml:"OrderCondition,omitempty"`
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageToken             *string `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	TagId                 *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TaskFlowStatusId      *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
}

func (s GetTaskListFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterRequest) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterRequest) SetCreatorId(v string) *GetTaskListFilterRequest {
	s.CreatorId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetDueDateEnd(v string) *GetTaskListFilterRequest {
	s.DueDateEnd = &v
	return s
}

func (s *GetTaskListFilterRequest) SetDueDateStart(v string) *GetTaskListFilterRequest {
	s.DueDateStart = &v
	return s
}

func (s *GetTaskListFilterRequest) SetExecutorId(v string) *GetTaskListFilterRequest {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetExtra(v string) *GetTaskListFilterRequest {
	s.Extra = &v
	return s
}

func (s *GetTaskListFilterRequest) SetInvolveMembers(v string) *GetTaskListFilterRequest {
	s.InvolveMembers = &v
	return s
}

func (s *GetTaskListFilterRequest) SetIsDone(v bool) *GetTaskListFilterRequest {
	s.IsDone = &v
	return s
}

func (s *GetTaskListFilterRequest) SetName(v string) *GetTaskListFilterRequest {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterRequest) SetObjectType(v string) *GetTaskListFilterRequest {
	s.ObjectType = &v
	return s
}

func (s *GetTaskListFilterRequest) SetOrder(v string) *GetTaskListFilterRequest {
	s.Order = &v
	return s
}

func (s *GetTaskListFilterRequest) SetOrderCondition(v string) *GetTaskListFilterRequest {
	s.OrderCondition = &v
	return s
}

func (s *GetTaskListFilterRequest) SetOrgId(v string) *GetTaskListFilterRequest {
	s.OrgId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetPageSize(v int32) *GetTaskListFilterRequest {
	s.PageSize = &v
	return s
}

func (s *GetTaskListFilterRequest) SetPageToken(v string) *GetTaskListFilterRequest {
	s.PageToken = &v
	return s
}

func (s *GetTaskListFilterRequest) SetPriority(v string) *GetTaskListFilterRequest {
	s.Priority = &v
	return s
}

func (s *GetTaskListFilterRequest) SetProjectId(v string) *GetTaskListFilterRequest {
	s.ProjectId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetScenarioFieldConfigId(v string) *GetTaskListFilterRequest {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetSprintId(v string) *GetTaskListFilterRequest {
	s.SprintId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetTagId(v string) *GetTaskListFilterRequest {
	s.TagId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetTaskFlowStatusId(v string) *GetTaskListFilterRequest {
	s.TaskFlowStatusId = &v
	return s
}

type GetTaskListFilterResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetTaskListFilterResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *string                              `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskListFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBody) SetErrorCode(v string) *GetTaskListFilterResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskListFilterResponseBody) SetErrorMsg(v string) *GetTaskListFilterResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTaskListFilterResponseBody) SetObject(v *GetTaskListFilterResponseBodyObject) *GetTaskListFilterResponseBody {
	s.Object = v
	return s
}

func (s *GetTaskListFilterResponseBody) SetRequestId(v string) *GetTaskListFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskListFilterResponseBody) SetSuccessful(v string) *GetTaskListFilterResponseBody {
	s.Successful = &v
	return s
}

type GetTaskListFilterResponseBodyObject struct {
	NextPageToken *string                                      `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	Result        []*GetTaskListFilterResponseBodyObjectResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	TotalSize     *int32                                       `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetTaskListFilterResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObject) SetNextPageToken(v string) *GetTaskListFilterResponseBodyObject {
	s.NextPageToken = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObject) SetResult(v []*GetTaskListFilterResponseBodyObjectResult) *GetTaskListFilterResponseBodyObject {
	s.Result = v
	return s
}

func (s *GetTaskListFilterResponseBodyObject) SetTotalSize(v int32) *GetTaskListFilterResponseBodyObject {
	s.TotalSize = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResult struct {
	Accomplished          *string                                                  `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	AncestorIds           []*string                                                `json:"AncestorIds,omitempty" xml:"AncestorIds,omitempty" type:"Repeated"`
	AttachmentsCount      *int32                                                   `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	Badges                *GetTaskListFilterResponseBodyObjectResultBadges         `json:"Badges,omitempty" xml:"Badges,omitempty" type:"Struct"`
	CommentsCount         *int32                                                   `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	Content               *string                                                  `json:"Content,omitempty" xml:"Content,omitempty"`
	Created               *string                                                  `json:"Created,omitempty" xml:"Created,omitempty"`
	Creator               *GetTaskListFilterResponseBodyObjectResultCreator        `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	CreatorId             *string                                                  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Customfields          []*GetTaskListFilterResponseBodyObjectResultCustomfields `json:"Customfields,omitempty" xml:"Customfields,omitempty" type:"Repeated"`
	Divisions             []*string                                                `json:"Divisions,omitempty" xml:"Divisions,omitempty" type:"Repeated"`
	DueDate               *string                                                  `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Executor              *GetTaskListFilterResponseBodyObjectResultExecutor       `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	ExecutorId            *string                                                  `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Id                    *string                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	InvolveMembers        []*string                                                `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	IsArchived            *bool                                                    `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	IsDone                *bool                                                    `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	IsFavorite            *bool                                                    `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	IsTopInProject        *bool                                                    `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Labels                []*string                                                `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LikesCount            *int32                                                   `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	Note                  *string                                                  `json:"Note,omitempty" xml:"Note,omitempty"`
	ObjectType            *string                                                  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ObjectlinksCount      *int32                                                   `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	OrganizationId        *string                                                  `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Parent                *string                                                  `json:"Parent,omitempty" xml:"Parent,omitempty"`
	Priority              *int32                                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Progress              *int32                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ProjectId             *string                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Rating                *int32                                                   `json:"Rating,omitempty" xml:"Rating,omitempty"`
	Recurrence            *string                                                  `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	Reminder              *GetTaskListFilterResponseBodyObjectResultReminder       `json:"Reminder,omitempty" xml:"Reminder,omitempty" type:"Struct"`
	ScenariofFeldConfigId *string                                                  `json:"ScenariofFeldConfigId,omitempty" xml:"ScenariofFeldConfigId,omitempty"`
	ShareStatus           *int32                                                   `json:"ShareStatus,omitempty" xml:"ShareStatus,omitempty"`
	Source                *string                                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceDate            *string                                                  `json:"SourceDate,omitempty" xml:"SourceDate,omitempty"`
	SourceId              *string                                                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	Sprint                *string                                                  `json:"Sprint,omitempty" xml:"Sprint,omitempty"`
	SprintId              *string                                                  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	Stage                 *GetTaskListFilterResponseBodyObjectResultStage          `json:"Stage,omitempty" xml:"Stage,omitempty" type:"Struct"`
	StageId               *string                                                  `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StartDate             *string                                                  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint            *string                                                  `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	SubtaskCount          *GetTaskListFilterResponseBodyObjectResultSubtaskCount   `json:"SubtaskCount,omitempty" xml:"SubtaskCount,omitempty" type:"Struct"`
	TagIds                []*string                                                `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	TaskFlowStatus        *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus `json:"TaskFlowStatus,omitempty" xml:"TaskFlowStatus,omitempty" type:"Struct"`
	TaskFlowStatusId      *string                                                  `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	TaskId                *string                                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskListId            *string                                                  `json:"TaskListId,omitempty" xml:"TaskListId,omitempty"`
	TaskUniqueId          *string                                                  `json:"TaskUniqueId,omitempty" xml:"TaskUniqueId,omitempty"`
	UniqueId              *int32                                                   `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	UntilDate             *string                                                  `json:"UntilDate,omitempty" xml:"UntilDate,omitempty"`
	Updated               *string                                                  `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visible               *string                                                  `json:"Visible,omitempty" xml:"Visible,omitempty"`
	WorkTime              *GetTaskListFilterResponseBodyObjectResultWorkTime       `json:"WorkTime,omitempty" xml:"WorkTime,omitempty" type:"Struct"`
}

func (s GetTaskListFilterResponseBodyObjectResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResult) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAccomplished(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Accomplished = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAncestorIds(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.AncestorIds = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAttachmentsCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetBadges(v *GetTaskListFilterResponseBodyObjectResultBadges) *GetTaskListFilterResponseBodyObjectResult {
	s.Badges = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCommentsCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetContent(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Content = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreated(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Created = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreator(v *GetTaskListFilterResponseBodyObjectResultCreator) *GetTaskListFilterResponseBodyObjectResult {
	s.Creator = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreatorId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.CreatorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCustomfields(v []*GetTaskListFilterResponseBodyObjectResultCustomfields) *GetTaskListFilterResponseBodyObjectResult {
	s.Customfields = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetDivisions(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.Divisions = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetDueDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.DueDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetExecutor(v *GetTaskListFilterResponseBodyObjectResultExecutor) *GetTaskListFilterResponseBodyObjectResult {
	s.Executor = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetExecutorId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetInvolveMembers(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.InvolveMembers = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsArchived(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsArchived = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsDone(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsDone = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsFavorite(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsFavorite = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsTopInProject(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsTopInProject = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetLabels(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.Labels = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetLikesCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.LikesCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetNote(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Note = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetObjectType(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ObjectType = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetObjectlinksCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetOrganizationId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.OrganizationId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetParent(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Parent = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetPriority(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Priority = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetProgress(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Progress = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetProjectId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ProjectId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetRating(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Rating = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetRecurrence(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Recurrence = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetReminder(v *GetTaskListFilterResponseBodyObjectResultReminder) *GetTaskListFilterResponseBodyObjectResult {
	s.Reminder = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetScenariofFeldConfigId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ScenariofFeldConfigId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetShareStatus(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.ShareStatus = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSource(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Source = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSourceDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SourceDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSourceId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SourceId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSprint(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Sprint = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSprintId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SprintId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStage(v *GetTaskListFilterResponseBodyObjectResultStage) *GetTaskListFilterResponseBodyObjectResult {
	s.Stage = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStageId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StageId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStartDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StartDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStoryPoint(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StoryPoint = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSubtaskCount(v *GetTaskListFilterResponseBodyObjectResultSubtaskCount) *GetTaskListFilterResponseBodyObjectResult {
	s.SubtaskCount = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTagIds(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.TagIds = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskFlowStatus(v *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskFlowStatus = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskFlowStatusId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskFlowStatusId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskListId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskListId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskUniqueId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskUniqueId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUniqueId(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.UniqueId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUntilDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.UntilDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUpdated(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Updated = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetVisible(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Visible = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetWorkTime(v *GetTaskListFilterResponseBodyObjectResultWorkTime) *GetTaskListFilterResponseBodyObjectResult {
	s.WorkTime = v
	return s
}

type GetTaskListFilterResponseBodyObjectResultBadges struct {
	AttachmentsCount *int32 `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	CommentsCount    *int32 `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	LikesCount       *int32 `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	ObjectlinksCount *int32 `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultBadges) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultBadges) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetAttachmentsCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetCommentsCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetLikesCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.LikesCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetObjectlinksCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.ObjectlinksCount = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultCreator struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultCreator) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultCreator) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultCreator) SetAvatarUrl(v string) *GetTaskListFilterResponseBodyObjectResultCreator {
	s.AvatarUrl = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCreator) SetId(v string) *GetTaskListFilterResponseBodyObjectResultCreator {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCreator) SetName(v string) *GetTaskListFilterResponseBodyObjectResultCreator {
	s.Name = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultCustomfields struct {
	CustomfieldId *string                                                       `json:"CustomfieldId,omitempty" xml:"CustomfieldId,omitempty"`
	Type          *string                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         []*GetTaskListFilterResponseBodyObjectResultCustomfieldsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Values        *string                                                       `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfields) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfields) SetCustomfieldId(v string) *GetTaskListFilterResponseBodyObjectResultCustomfields {
	s.CustomfieldId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfields) SetType(v string) *GetTaskListFilterResponseBodyObjectResultCustomfields {
	s.Type = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfields) SetValue(v []*GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) *GetTaskListFilterResponseBodyObjectResultCustomfields {
	s.Value = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfields) SetValues(v string) *GetTaskListFilterResponseBodyObjectResultCustomfields {
	s.Values = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultCustomfieldsValue struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) SetId(v string) *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) SetTitle(v string) *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue {
	s.Title = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultExecutor struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultExecutor) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultExecutor) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultExecutor) SetAvatarUrl(v string) *GetTaskListFilterResponseBodyObjectResultExecutor {
	s.AvatarUrl = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultExecutor) SetId(v string) *GetTaskListFilterResponseBodyObjectResultExecutor {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultExecutor) SetName(v string) *GetTaskListFilterResponseBodyObjectResultExecutor {
	s.Name = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultReminder struct {
	CreatorId *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Date      *string   `json:"Date,omitempty" xml:"Date,omitempty"`
	Members   []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Rules     []*string `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	Type      *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultReminder) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultReminder) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetCreatorId(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.CreatorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetDate(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Date = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetMembers(v []*string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Members = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetRules(v []*string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Rules = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetType(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Type = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultStage struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultStage) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultStage) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultStage) SetId(v string) *GetTaskListFilterResponseBodyObjectResultStage {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultStage) SetName(v string) *GetTaskListFilterResponseBodyObjectResultStage {
	s.Name = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultSubtaskCount struct {
	Done  *int32 `json:"Done,omitempty" xml:"Done,omitempty"`
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultSubtaskCount) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultSubtaskCount) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultSubtaskCount) SetDone(v int32) *GetTaskListFilterResponseBodyObjectResultSubtaskCount {
	s.Done = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultSubtaskCount) SetTotal(v int32) *GetTaskListFilterResponseBodyObjectResultSubtaskCount {
	s.Total = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultTaskFlowStatus struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Kind       *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pos        *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	TaskFlowId *string `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetId(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Id = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetKind(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Kind = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetName(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetPos(v int32) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Pos = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetTaskFlowId(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.TaskFlowId = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultWorkTime struct {
	TotalTime *int32  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Unit      *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	UsedTime  *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultWorkTime) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultWorkTime) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetTotalTime(v int32) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.TotalTime = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetUnit(v string) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.Unit = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetUsedTime(v int32) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.UsedTime = &v
	return s
}

type GetTaskListFilterResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTaskListFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskListFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponse) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponse) SetHeaders(v map[string]*string) *GetTaskListFilterResponse {
	s.Headers = v
	return s
}

func (s *GetTaskListFilterResponse) SetStatusCode(v int32) *GetTaskListFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskListFilterResponse) SetBody(v *GetTaskListFilterResponseBody) *GetTaskListFilterResponse {
	s.Body = v
	return s
}

type InsertPipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s InsertPipelineMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertPipelineMemberRequest) GoString() string {
	return s.String()
}

func (s *InsertPipelineMemberRequest) SetOrgId(v string) *InsertPipelineMemberRequest {
	s.OrgId = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetPipelineId(v int64) *InsertPipelineMemberRequest {
	s.PipelineId = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetRoleName(v string) *InsertPipelineMemberRequest {
	s.RoleName = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetUserId(v string) *InsertPipelineMemberRequest {
	s.UserId = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetUserPk(v string) *InsertPipelineMemberRequest {
	s.UserPk = &v
	return s
}

type InsertPipelineMemberResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertPipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertPipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *InsertPipelineMemberResponseBody) SetErrorCode(v string) *InsertPipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetErrorMessage(v string) *InsertPipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetObject(v bool) *InsertPipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetRequestId(v string) *InsertPipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetSuccess(v bool) *InsertPipelineMemberResponseBody {
	s.Success = &v
	return s
}

type InsertPipelineMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertPipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertPipelineMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertPipelineMemberResponse) GoString() string {
	return s.String()
}

func (s *InsertPipelineMemberResponse) SetHeaders(v map[string]*string) *InsertPipelineMemberResponse {
	s.Headers = v
	return s
}

func (s *InsertPipelineMemberResponse) SetStatusCode(v int32) *InsertPipelineMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertPipelineMemberResponse) SetBody(v *InsertPipelineMemberResponseBody) *InsertPipelineMemberResponse {
	s.Body = v
	return s
}

type InsertProjectMembersRequest struct {
	Members   *string `json:"Members,omitempty" xml:"Members,omitempty"`
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s InsertProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *InsertProjectMembersRequest) SetMembers(v string) *InsertProjectMembersRequest {
	s.Members = &v
	return s
}

func (s *InsertProjectMembersRequest) SetOrgId(v string) *InsertProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *InsertProjectMembersRequest) SetProjectId(v string) *InsertProjectMembersRequest {
	s.ProjectId = &v
	return s
}

type InsertProjectMembersResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s InsertProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *InsertProjectMembersResponseBody) SetErrorCode(v string) *InsertProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertProjectMembersResponseBody) SetErrorMsg(v string) *InsertProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *InsertProjectMembersResponseBody) SetObject(v bool) *InsertProjectMembersResponseBody {
	s.Object = &v
	return s
}

func (s *InsertProjectMembersResponseBody) SetRequestId(v string) *InsertProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertProjectMembersResponseBody) SetSuccessful(v bool) *InsertProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type InsertProjectMembersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InsertProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *InsertProjectMembersResponse) SetHeaders(v map[string]*string) *InsertProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *InsertProjectMembersResponse) SetStatusCode(v int32) *InsertProjectMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertProjectMembersResponse) SetBody(v *InsertProjectMembersResponseBody) *InsertProjectMembersResponse {
	s.Body = v
	return s
}

type ListCommonGroupRequest struct {
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OrgId        *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SmartGroupId *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
}

func (s ListCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *ListCommonGroupRequest) SetAll(v bool) *ListCommonGroupRequest {
	s.All = &v
	return s
}

func (s *ListCommonGroupRequest) SetOrgId(v string) *ListCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *ListCommonGroupRequest) SetProjectId(v string) *ListCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCommonGroupRequest) SetSmartGroupId(v string) *ListCommonGroupRequest {
	s.SmartGroupId = &v
	return s
}

type ListCommonGroupResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonGroupResponseBody) SetErrorCode(v string) *ListCommonGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCommonGroupResponseBody) SetErrorMsg(v string) *ListCommonGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListCommonGroupResponseBody) SetObject(v []*ListCommonGroupResponseBodyObject) *ListCommonGroupResponseBody {
	s.Object = v
	return s
}

func (s *ListCommonGroupResponseBody) SetRequestId(v string) *ListCommonGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonGroupResponseBody) SetSuccessful(v bool) *ListCommonGroupResponseBody {
	s.Successful = &v
	return s
}

type ListCommonGroupResponseBodyObject struct {
	CreatorId     *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	IsRoot        *bool   `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pinyin        *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	Pos           *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResourceCount *int32  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	SmartGroupId  *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListCommonGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListCommonGroupResponseBodyObject) SetCreatorId(v string) *ListCommonGroupResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetIsRoot(v bool) *ListCommonGroupResponseBodyObject {
	s.IsRoot = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetName(v string) *ListCommonGroupResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetPinyin(v string) *ListCommonGroupResponseBodyObject {
	s.Pinyin = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetPos(v int32) *ListCommonGroupResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetProjectId(v string) *ListCommonGroupResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetResourceCount(v int32) *ListCommonGroupResponseBodyObject {
	s.ResourceCount = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetSmartGroupId(v string) *ListCommonGroupResponseBodyObject {
	s.SmartGroupId = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetId(v string) *ListCommonGroupResponseBodyObject {
	s.Id = &v
	return s
}

type ListCommonGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommonGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupResponse) GoString() string {
	return s.String()
}

func (s *ListCommonGroupResponse) SetHeaders(v map[string]*string) *ListCommonGroupResponse {
	s.Headers = v
	return s
}

func (s *ListCommonGroupResponse) SetStatusCode(v int32) *ListCommonGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCommonGroupResponse) SetBody(v *ListCommonGroupResponseBody) *ListCommonGroupResponse {
	s.Body = v
	return s
}

type ListCredentialsRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s ListCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCredentialsRequest) GoString() string {
	return s.String()
}

func (s *ListCredentialsRequest) SetOrgId(v string) *ListCredentialsRequest {
	s.OrgId = &v
	return s
}

func (s *ListCredentialsRequest) SetUserPk(v string) *ListCredentialsRequest {
	s.UserPk = &v
	return s
}

type ListCredentialsResponseBody struct {
	ErrorCode    *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       []map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) SetErrorCode(v string) *ListCredentialsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCredentialsResponseBody) SetErrorMessage(v string) *ListCredentialsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCredentialsResponseBody) SetObject(v []map[string]interface{}) *ListCredentialsResponseBody {
	s.Object = v
	return s
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetSuccess(v bool) *ListCredentialsResponseBody {
	s.Success = &v
	return s
}

type ListCredentialsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCredentialsResponse) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponse) SetHeaders(v map[string]*string) *ListCredentialsResponse {
	s.Headers = v
	return s
}

func (s *ListCredentialsResponse) SetStatusCode(v int32) *ListCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCredentialsResponse) SetBody(v *ListCredentialsResponseBody) *ListCredentialsResponse {
	s.Body = v
	return s
}

type ListDevopsProjectSprintsRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageToken *string `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDevopsProjectSprintsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsRequest) SetOrgId(v string) *ListDevopsProjectSprintsRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectSprintsRequest) SetPageSize(v int64) *ListDevopsProjectSprintsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevopsProjectSprintsRequest) SetPageToken(v string) *ListDevopsProjectSprintsRequest {
	s.PageToken = &v
	return s
}

func (s *ListDevopsProjectSprintsRequest) SetProjectId(v string) *ListDevopsProjectSprintsRequest {
	s.ProjectId = &v
	return s
}

type ListDevopsProjectSprintsResponseBody struct {
	ErrorCode     *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg      *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	NextPageToken *string                                       `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	Object        []*ListDevopsProjectSprintsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful    *bool                                         `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBody) SetErrorCode(v string) *ListDevopsProjectSprintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetErrorMsg(v string) *ListDevopsProjectSprintsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetNextPageToken(v string) *ListDevopsProjectSprintsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetObject(v []*ListDevopsProjectSprintsResponseBodyObject) *ListDevopsProjectSprintsResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetRequestId(v string) *ListDevopsProjectSprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetSuccessful(v bool) *ListDevopsProjectSprintsResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectSprintsResponseBodyObject struct {
	Accomplished *string                                             `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	Created      *string                                             `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId    *string                                             `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DueDate      *string                                             `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Id           *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDeleted    *bool                                               `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Name         *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanToDo     *ListDevopsProjectSprintsResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	ProjectId    *string                                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                             `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Status       *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Updated      *string                                             `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetAccomplished(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetCreated(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetCreatorId(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetDueDate(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetId(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetIsDeleted(v bool) *ListDevopsProjectSprintsResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetName(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetPlanToDo(v *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) *ListDevopsProjectSprintsResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetProjectId(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetStartDate(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetStatus(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Status = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetUpdated(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Updated = &v
	return s
}

type ListDevopsProjectSprintsResponseBodyObjectPlanToDo struct {
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetTasks(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

type ListDevopsProjectSprintsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsProjectSprintsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectSprintsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponse) SetHeaders(v map[string]*string) *ListDevopsProjectSprintsResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectSprintsResponse) SetStatusCode(v int32) *ListDevopsProjectSprintsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsProjectSprintsResponse) SetBody(v *ListDevopsProjectSprintsResponseBody) *ListDevopsProjectSprintsResponse {
	s.Body = v
	return s
}

type ListDevopsProjectTaskFlowRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDevopsProjectTaskFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowRequest) SetOrgId(v string) *ListDevopsProjectTaskFlowRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowRequest) SetProjectId(v string) *ListDevopsProjectTaskFlowRequest {
	s.ProjectId = &v
	return s
}

type ListDevopsProjectTaskFlowResponseBody struct {
	ErrorCode  *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                        `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectTaskFlowResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                          `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetErrorMsg(v string) *ListDevopsProjectTaskFlowResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetObject(v []*ListDevopsProjectTaskFlowResponseBodyObject) *ListDevopsProjectTaskFlowResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetRequestId(v string) *ListDevopsProjectTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetSuccessful(v bool) *ListDevopsProjectTaskFlowResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTaskFlowResponseBodyObject struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDevopsProjectTaskFlowResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetId(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetName(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetType(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Type = &v
	return s
}

type ListDevopsProjectTaskFlowResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsProjectTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectTaskFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowResponse) SetHeaders(v map[string]*string) *ListDevopsProjectTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectTaskFlowResponse) SetStatusCode(v int32) *ListDevopsProjectTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponse) SetBody(v *ListDevopsProjectTaskFlowResponseBody) *ListDevopsProjectTaskFlowResponse {
	s.Body = v
	return s
}

type ListDevopsProjectTaskFlowStatusRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskFlowId *string `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
}

func (s ListDevopsProjectTaskFlowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusRequest) SetOrgId(v string) *ListDevopsProjectTaskFlowStatusRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusRequest) SetTaskFlowId(v string) *ListDevopsProjectTaskFlowStatusRequest {
	s.TaskFlowId = &v
	return s
}

type ListDevopsProjectTaskFlowStatusResponseBody struct {
	ErrorCode  *string                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectTaskFlowStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskFlowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetErrorMsg(v string) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetObject(v []*ListDevopsProjectTaskFlowStatusResponseBodyObject) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetRequestId(v string) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetSuccessful(v bool) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTaskFlowStatusResponseBodyObject struct {
	Created         *string `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId       *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsDeleted       *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Kind            *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pos             *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	RejectStatusIds *string `json:"RejectStatusIds,omitempty" xml:"RejectStatusIds,omitempty"`
	TaskflowId      *string `json:"TaskflowId,omitempty" xml:"TaskflowId,omitempty"`
	Updated         *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
}

func (s ListDevopsProjectTaskFlowStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetCreated(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetCreatorId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetIsDeleted(v bool) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetKind(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Kind = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetName(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetPos(v int32) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetRejectStatusIds(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.RejectStatusIds = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetTaskflowId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.TaskflowId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetUpdated(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Updated = &v
	return s
}

type ListDevopsProjectTaskFlowStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsProjectTaskFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectTaskFlowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusResponse) SetHeaders(v map[string]*string) *ListDevopsProjectTaskFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponse) SetStatusCode(v int32) *ListDevopsProjectTaskFlowStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponse) SetBody(v *ListDevopsProjectTaskFlowStatusResponseBody) *ListDevopsProjectTaskFlowStatusResponse {
	s.Body = v
	return s
}

type ListDevopsProjectTaskListRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDevopsProjectTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListRequest) SetOrgId(v string) *ListDevopsProjectTaskListRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectTaskListRequest) SetProjectId(v string) *ListDevopsProjectTaskListRequest {
	s.ProjectId = &v
	return s
}

type ListDevopsProjectTaskListResponseBody struct {
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *ListDevopsProjectTaskListResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTaskListResponseBody) SetErrorMsg(v string) *ListDevopsProjectTaskListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectTaskListResponseBody) SetObject(v *ListDevopsProjectTaskListResponseBodyObject) *ListDevopsProjectTaskListResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectTaskListResponseBody) SetRequestId(v string) *ListDevopsProjectTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectTaskListResponseBody) SetSuccessful(v bool) *ListDevopsProjectTaskListResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTaskListResponseBodyObject struct {
	Result []*ListDevopsProjectTaskListResponseBodyObjectResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDevopsProjectTaskListResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListResponseBodyObject) SetResult(v []*ListDevopsProjectTaskListResponseBodyObjectResult) *ListDevopsProjectTaskListResponseBodyObject {
	s.Result = v
	return s
}

type ListDevopsProjectTaskListResponseBodyObjectResult struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsProjectTaskListResponseBodyObjectResult) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListResponseBodyObjectResult) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListResponseBodyObjectResult) SetId(v string) *ListDevopsProjectTaskListResponseBodyObjectResult {
	s.Id = &v
	return s
}

type ListDevopsProjectTaskListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsProjectTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListResponse) SetHeaders(v map[string]*string) *ListDevopsProjectTaskListResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectTaskListResponse) SetStatusCode(v int32) *ListDevopsProjectTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsProjectTaskListResponse) SetBody(v *ListDevopsProjectTaskListResponseBody) *ListDevopsProjectTaskListResponse {
	s.Body = v
	return s
}

type ListDevopsProjectsRequest struct {
	OrderBy   *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageToken *string `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	SelectBy  *string `json:"SelectBy,omitempty" xml:"SelectBy,omitempty"`
}

func (s ListDevopsProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectsRequest) SetOrderBy(v string) *ListDevopsProjectsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDevopsProjectsRequest) SetOrgId(v string) *ListDevopsProjectsRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectsRequest) SetPageSize(v int32) *ListDevopsProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDevopsProjectsRequest) SetPageToken(v string) *ListDevopsProjectsRequest {
	s.PageToken = &v
	return s
}

func (s *ListDevopsProjectsRequest) SetSelectBy(v string) *ListDevopsProjectsRequest {
	s.SelectBy = &v
	return s
}

type ListDevopsProjectsResponseBody struct {
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *ListDevopsProjectsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectsResponseBody) SetErrorCode(v string) *ListDevopsProjectsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectsResponseBody) SetErrorMsg(v string) *ListDevopsProjectsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectsResponseBody) SetObject(v *ListDevopsProjectsResponseBodyObject) *ListDevopsProjectsResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectsResponseBody) SetRequestId(v string) *ListDevopsProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectsResponseBody) SetSuccessful(v bool) *ListDevopsProjectsResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectsResponseBodyObject struct {
	NextPageToken *string                                       `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	Result        []*ListDevopsProjectsResponseBodyObjectResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListDevopsProjectsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectsResponseBodyObject) SetNextPageToken(v string) *ListDevopsProjectsResponseBodyObject {
	s.NextPageToken = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObject) SetResult(v []*ListDevopsProjectsResponseBodyObjectResult) *ListDevopsProjectsResponseBodyObject {
	s.Result = v
	return s
}

type ListDevopsProjectsResponseBodyObjectResult struct {
	Created        *string `json:"Created,omitempty" xml:"Created,omitempty"`
	CreatorId      *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	IsArchived     *bool   `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	IsPublic       *bool   `json:"IsPublic,omitempty" xml:"IsPublic,omitempty"`
	IsStar         *bool   `json:"IsStar,omitempty" xml:"IsStar,omitempty"`
	IsTemplate     *bool   `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	Logo           *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	MembersCount   *int32  `json:"MembersCount,omitempty" xml:"MembersCount,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	RoleId         *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TasksCount     *int32  `json:"TasksCount,omitempty" xml:"TasksCount,omitempty"`
	Updated        *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Visibility     *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListDevopsProjectsResponseBodyObjectResult) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectsResponseBodyObjectResult) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetCreated(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetCreatorId(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetDescription(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Description = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetId(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Id = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetIsArchived(v bool) *ListDevopsProjectsResponseBodyObjectResult {
	s.IsArchived = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetIsPublic(v bool) *ListDevopsProjectsResponseBodyObjectResult {
	s.IsPublic = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetIsStar(v bool) *ListDevopsProjectsResponseBodyObjectResult {
	s.IsStar = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetIsTemplate(v bool) *ListDevopsProjectsResponseBodyObjectResult {
	s.IsTemplate = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetLogo(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Logo = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetMembersCount(v int32) *ListDevopsProjectsResponseBodyObjectResult {
	s.MembersCount = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetName(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetOrganizationId(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.OrganizationId = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetRoleId(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.RoleId = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetTasksCount(v int32) *ListDevopsProjectsResponseBodyObjectResult {
	s.TasksCount = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetUpdated(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Updated = &v
	return s
}

func (s *ListDevopsProjectsResponseBodyObjectResult) SetVisibility(v string) *ListDevopsProjectsResponseBodyObjectResult {
	s.Visibility = &v
	return s
}

type ListDevopsProjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectsResponse) SetHeaders(v map[string]*string) *ListDevopsProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectsResponse) SetStatusCode(v int32) *ListDevopsProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsProjectsResponse) SetBody(v *ListDevopsProjectsResponseBody) *ListDevopsProjectsResponse {
	s.Body = v
	return s
}

type ListDevopsScenarioFieldConfigRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDevopsScenarioFieldConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigRequest) SetOrgId(v string) *ListDevopsScenarioFieldConfigRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigRequest) SetProjectId(v string) *ListDevopsScenarioFieldConfigRequest {
	s.ProjectId = &v
	return s
}

type ListDevopsScenarioFieldConfigResponseBody struct {
	ErrorCode  *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                            `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsScenarioFieldConfigResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                              `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsScenarioFieldConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetErrorCode(v string) *ListDevopsScenarioFieldConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetErrorMsg(v string) *ListDevopsScenarioFieldConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetObject(v []*ListDevopsScenarioFieldConfigResponseBodyObject) *ListDevopsScenarioFieldConfigResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetRequestId(v string) *ListDevopsScenarioFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetSuccessful(v bool) *ListDevopsScenarioFieldConfigResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsScenarioFieldConfigResponseBodyObject struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDevopsScenarioFieldConfigResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigResponseBodyObject) SetId(v string) *ListDevopsScenarioFieldConfigResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBodyObject) SetName(v string) *ListDevopsScenarioFieldConfigResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBodyObject) SetType(v string) *ListDevopsScenarioFieldConfigResponseBodyObject {
	s.Type = &v
	return s
}

type ListDevopsScenarioFieldConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDevopsScenarioFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsScenarioFieldConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigResponse) SetHeaders(v map[string]*string) *ListDevopsScenarioFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponse) SetStatusCode(v int32) *ListDevopsScenarioFieldConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponse) SetBody(v *ListDevopsScenarioFieldConfigResponseBody) *ListDevopsScenarioFieldConfigResponse {
	s.Body = v
	return s
}

type ListPipelineTemplatesRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageStart *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
}

func (s ListPipelineTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesRequest) SetOrgId(v string) *ListPipelineTemplatesRequest {
	s.OrgId = &v
	return s
}

func (s *ListPipelineTemplatesRequest) SetPageNum(v int32) *ListPipelineTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *ListPipelineTemplatesRequest) SetPageStart(v int32) *ListPipelineTemplatesRequest {
	s.PageStart = &v
	return s
}

type ListPipelineTemplatesResponseBody struct {
	Data      *ListPipelineTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPipelineTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesResponseBody) SetData(v *ListPipelineTemplatesResponseBodyData) *ListPipelineTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListPipelineTemplatesResponseBody) SetRequestId(v string) *ListPipelineTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListPipelineTemplatesResponseBodyData struct {
	DataList []*ListPipelineTemplatesResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	Total    *float32                                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPipelineTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesResponseBodyData) SetDataList(v []*ListPipelineTemplatesResponseBodyDataDataList) *ListPipelineTemplatesResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListPipelineTemplatesResponseBodyData) SetTotal(v float32) *ListPipelineTemplatesResponseBodyData {
	s.Total = &v
	return s
}

type ListPipelineTemplatesResponseBodyDataDataList struct {
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListPipelineTemplatesResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesResponseBodyDataDataList) SetId(v int64) *ListPipelineTemplatesResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListPipelineTemplatesResponseBodyDataDataList) SetTemplateName(v string) *ListPipelineTemplatesResponseBodyDataDataList {
	s.TemplateName = &v
	return s
}

type ListPipelineTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelineTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelineTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelineTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineTemplatesResponse) SetHeaders(v map[string]*string) *ListPipelineTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineTemplatesResponse) SetStatusCode(v int32) *ListPipelineTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineTemplatesResponse) SetBody(v *ListPipelineTemplatesResponseBody) *ListPipelineTemplatesResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	CreateEndTime    *string `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	CreateStartTime  *string `json:"CreateStartTime,omitempty" xml:"CreateStartTime,omitempty"`
	Creators         *string `json:"Creators,omitempty" xml:"Creators,omitempty"`
	ExecuteEndTime   *string `json:"ExecuteEndTime,omitempty" xml:"ExecuteEndTime,omitempty"`
	ExecuteStartTime *string `json:"ExecuteStartTime,omitempty" xml:"ExecuteStartTime,omitempty"`
	Operators        *string `json:"Operators,omitempty" xml:"Operators,omitempty"`
	OrgId            *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart        *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	PipelineName     *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	ResultStatusList *string `json:"ResultStatusList,omitempty" xml:"ResultStatusList,omitempty"`
	UserPk           *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) SetCreateEndTime(v string) *ListPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreateStartTime(v string) *ListPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreators(v string) *ListPipelinesRequest {
	s.Creators = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteEndTime(v string) *ListPipelinesRequest {
	s.ExecuteEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteStartTime(v string) *ListPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetOperators(v string) *ListPipelinesRequest {
	s.Operators = &v
	return s
}

func (s *ListPipelinesRequest) SetOrgId(v string) *ListPipelinesRequest {
	s.OrgId = &v
	return s
}

func (s *ListPipelinesRequest) SetPageSize(v int32) *ListPipelinesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPipelinesRequest) SetPageStart(v int32) *ListPipelinesRequest {
	s.PageStart = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineName(v string) *ListPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesRequest) SetResultStatusList(v string) *ListPipelinesRequest {
	s.ResultStatusList = &v
	return s
}

func (s *ListPipelinesRequest) SetUserPk(v string) *ListPipelinesRequest {
	s.UserPk = &v
	return s
}

type ListPipelinesResponseBody struct {
	ErrorCode    *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) SetErrorCode(v string) *ListPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelinesResponseBody) SetErrorMessage(v string) *ListPipelinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelinesResponseBody) SetObject(v map[string]interface{}) *ListPipelinesResponseBody {
	s.Object = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetSuccess(v bool) *ListPipelinesResponseBody {
	s.Success = &v
	return s
}

type ListPipelinesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPipelinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v *ListPipelinesResponseBody) *ListPipelinesResponse {
	s.Body = v
	return s
}

type ListProjectCustomFieldsRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectCustomFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsRequest) SetOrgId(v string) *ListProjectCustomFieldsRequest {
	s.OrgId = &v
	return s
}

func (s *ListProjectCustomFieldsRequest) SetProjectId(v string) *ListProjectCustomFieldsRequest {
	s.ProjectId = &v
	return s
}

type ListProjectCustomFieldsResponseBody struct {
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectCustomFieldsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectCustomFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBody) SetErrorCode(v string) *ListProjectCustomFieldsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBody) SetErrorMsg(v string) *ListProjectCustomFieldsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBody) SetObject(v []*ListProjectCustomFieldsResponseBodyObject) *ListProjectCustomFieldsResponseBody {
	s.Object = v
	return s
}

func (s *ListProjectCustomFieldsResponseBody) SetRequestId(v string) *ListProjectCustomFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBody) SetSuccessful(v bool) *ListProjectCustomFieldsResponseBody {
	s.Successful = &v
	return s
}

type ListProjectCustomFieldsResponseBodyObject struct {
	CustomFieldId *string                                            `json:"CustomFieldId,omitempty" xml:"CustomFieldId,omitempty"`
	Name          *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Subtype       *string                                            `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	Type          *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Values        []*ListProjectCustomFieldsResponseBodyObjectValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListProjectCustomFieldsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetCustomFieldId(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.CustomFieldId = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetName(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetSubtype(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Subtype = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetType(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetValues(v []*ListProjectCustomFieldsResponseBodyObjectValues) *ListProjectCustomFieldsResponseBodyObject {
	s.Values = v
	return s
}

type ListProjectCustomFieldsResponseBodyObjectValues struct {
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProjectCustomFieldsResponseBodyObjectValues) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBodyObjectValues) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBodyObjectValues) SetId(v string) *ListProjectCustomFieldsResponseBodyObjectValues {
	s.Id = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObjectValues) SetValue(v string) *ListProjectCustomFieldsResponseBodyObjectValues {
	s.Value = &v
	return s
}

type ListProjectCustomFieldsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectCustomFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectCustomFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponse) SetHeaders(v map[string]*string) *ListProjectCustomFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectCustomFieldsResponse) SetStatusCode(v int32) *ListProjectCustomFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectCustomFieldsResponse) SetBody(v *ListProjectCustomFieldsResponseBody) *ListProjectCustomFieldsResponse {
	s.Body = v
	return s
}

type ListServiceConnectionsRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ScType *string `json:"ScType,omitempty" xml:"ScType,omitempty"`
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s ListServiceConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsRequest) SetOrgId(v string) *ListServiceConnectionsRequest {
	s.OrgId = &v
	return s
}

func (s *ListServiceConnectionsRequest) SetScType(v string) *ListServiceConnectionsRequest {
	s.ScType = &v
	return s
}

func (s *ListServiceConnectionsRequest) SetUserPk(v string) *ListServiceConnectionsRequest {
	s.UserPk = &v
	return s
}

type ListServiceConnectionsResponseBody struct {
	ErrorCode    *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       []map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServiceConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBody) SetErrorCode(v string) *ListServiceConnectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetErrorMessage(v string) *ListServiceConnectionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetObject(v []map[string]interface{}) *ListServiceConnectionsResponseBody {
	s.Object = v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetRequestId(v string) *ListServiceConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetSuccess(v bool) *ListServiceConnectionsResponseBody {
	s.Success = &v
	return s
}

type ListServiceConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponse) SetHeaders(v map[string]*string) *ListServiceConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceConnectionsResponse) SetStatusCode(v int32) *ListServiceConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceConnectionsResponse) SetBody(v *ListServiceConnectionsResponseBody) *ListServiceConnectionsResponse {
	s.Body = v
	return s
}

type ListSmartGroupRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListSmartGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSmartGroupRequest) SetOrgId(v string) *ListSmartGroupRequest {
	s.OrgId = &v
	return s
}

func (s *ListSmartGroupRequest) SetProjectId(v string) *ListSmartGroupRequest {
	s.ProjectId = &v
	return s
}

type ListSmartGroupResponseBody struct {
	ErrorCode  *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListSmartGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                               `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListSmartGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartGroupResponseBody) SetErrorCode(v string) *ListSmartGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSmartGroupResponseBody) SetErrorMsg(v string) *ListSmartGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListSmartGroupResponseBody) SetObject(v []*ListSmartGroupResponseBodyObject) *ListSmartGroupResponseBody {
	s.Object = v
	return s
}

func (s *ListSmartGroupResponseBody) SetRequestId(v string) *ListSmartGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmartGroupResponseBody) SetSuccessful(v bool) *ListSmartGroupResponseBody {
	s.Successful = &v
	return s
}

type ListSmartGroupResponseBodyObject struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSmartGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListSmartGroupResponseBodyObject) SetId(v string) *ListSmartGroupResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListSmartGroupResponseBodyObject) SetType(v string) *ListSmartGroupResponseBodyObject {
	s.Type = &v
	return s
}

type ListSmartGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSmartGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSmartGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSmartGroupResponse) SetHeaders(v map[string]*string) *ListSmartGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSmartGroupResponse) SetStatusCode(v int32) *ListSmartGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartGroupResponse) SetBody(v *ListSmartGroupResponseBody) *ListSmartGroupResponse {
	s.Body = v
	return s
}

type ListUserOrganizationRequest struct {
	RealPk *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s ListUserOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationRequest) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationRequest) SetRealPk(v string) *ListUserOrganizationRequest {
	s.RealPk = &v
	return s
}

type ListUserOrganizationResponseBody struct {
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       []*ListUserOrganizationResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationResponseBody) SetErrorCode(v string) *ListUserOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetErrorMessage(v string) *ListUserOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetObject(v []*ListUserOrganizationResponseBodyObject) *ListUserOrganizationResponseBody {
	s.Object = v
	return s
}

func (s *ListUserOrganizationResponseBody) SetRequestId(v string) *ListUserOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetSuccess(v bool) *ListUserOrganizationResponseBody {
	s.Success = &v
	return s
}

type ListUserOrganizationResponseBodyObject struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListUserOrganizationResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationResponseBodyObject) SetId(v string) *ListUserOrganizationResponseBodyObject {
	s.Id = &v
	return s
}

func (s *ListUserOrganizationResponseBodyObject) SetName(v string) *ListUserOrganizationResponseBodyObject {
	s.Name = &v
	return s
}

type ListUserOrganizationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationResponse) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationResponse) SetHeaders(v map[string]*string) *ListUserOrganizationResponse {
	s.Headers = v
	return s
}

func (s *ListUserOrganizationResponse) SetStatusCode(v int32) *ListUserOrganizationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserOrganizationResponse) SetBody(v *ListUserOrganizationResponseBody) *ListUserOrganizationResponse {
	s.Body = v
	return s
}

type TransferPipelineOwnerRequest struct {
	NewOwnerId *string `json:"NewOwnerId,omitempty" xml:"NewOwnerId,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s TransferPipelineOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferPipelineOwnerRequest) GoString() string {
	return s.String()
}

func (s *TransferPipelineOwnerRequest) SetNewOwnerId(v string) *TransferPipelineOwnerRequest {
	s.NewOwnerId = &v
	return s
}

func (s *TransferPipelineOwnerRequest) SetOrgId(v string) *TransferPipelineOwnerRequest {
	s.OrgId = &v
	return s
}

func (s *TransferPipelineOwnerRequest) SetPipelineId(v int64) *TransferPipelineOwnerRequest {
	s.PipelineId = &v
	return s
}

func (s *TransferPipelineOwnerRequest) SetUserPk(v string) *TransferPipelineOwnerRequest {
	s.UserPk = &v
	return s
}

type TransferPipelineOwnerResponseBody struct {
	ErrorCode    *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferPipelineOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferPipelineOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPipelineOwnerResponseBody) SetErrorCode(v string) *TransferPipelineOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetErrorMessage(v string) *TransferPipelineOwnerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetObject(v map[string]interface{}) *TransferPipelineOwnerResponseBody {
	s.Object = v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetRequestId(v string) *TransferPipelineOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetSuccess(v bool) *TransferPipelineOwnerResponseBody {
	s.Success = &v
	return s
}

type TransferPipelineOwnerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferPipelineOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferPipelineOwnerResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferPipelineOwnerResponse) GoString() string {
	return s.String()
}

func (s *TransferPipelineOwnerResponse) SetHeaders(v map[string]*string) *TransferPipelineOwnerResponse {
	s.Headers = v
	return s
}

func (s *TransferPipelineOwnerResponse) SetStatusCode(v int32) *TransferPipelineOwnerResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferPipelineOwnerResponse) SetBody(v *TransferPipelineOwnerResponseBody) *TransferPipelineOwnerResponse {
	s.Body = v
	return s
}

type UpdateCommonGroupRequest struct {
	CommonGroupId *string `json:"CommonGroupId,omitempty" xml:"CommonGroupId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId         *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SmartGroupId  *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
}

func (s UpdateCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupRequest) SetCommonGroupId(v string) *UpdateCommonGroupRequest {
	s.CommonGroupId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetDescription(v string) *UpdateCommonGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetName(v string) *UpdateCommonGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetOrgId(v string) *UpdateCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetProjectId(v string) *UpdateCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetSmartGroupId(v string) *UpdateCommonGroupRequest {
	s.SmartGroupId = &v
	return s
}

type UpdateCommonGroupResponseBody struct {
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *UpdateCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupResponseBody) SetErrorCode(v string) *UpdateCommonGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateCommonGroupResponseBody) SetErrorMsg(v string) *UpdateCommonGroupResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateCommonGroupResponseBody) SetObject(v *UpdateCommonGroupResponseBodyObject) *UpdateCommonGroupResponseBody {
	s.Object = v
	return s
}

func (s *UpdateCommonGroupResponseBody) SetRequestId(v string) *UpdateCommonGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCommonGroupResponseBody) SetSuccessful(v bool) *UpdateCommonGroupResponseBody {
	s.Successful = &v
	return s
}

type UpdateCommonGroupResponseBodyObject struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateCommonGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupResponseBodyObject) SetId(v string) *UpdateCommonGroupResponseBodyObject {
	s.Id = &v
	return s
}

type UpdateCommonGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCommonGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupResponse) SetHeaders(v map[string]*string) *UpdateCommonGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommonGroupResponse) SetStatusCode(v int32) *UpdateCommonGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommonGroupResponse) SetBody(v *UpdateCommonGroupResponseBody) *UpdateCommonGroupResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDevopsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectRequest) SetDescription(v string) *UpdateDevopsProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetName(v string) *UpdateDevopsProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetOrgId(v string) *UpdateDevopsProjectRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetProjectId(v string) *UpdateDevopsProjectRequest {
	s.ProjectId = &v
	return s
}

type UpdateDevopsProjectResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectResponseBody) SetErrorCode(v string) *UpdateDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetErrorMessage(v string) *UpdateDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetObject(v string) *UpdateDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetRequestId(v string) *UpdateDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetSuccess(v bool) *UpdateDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type UpdateDevopsProjectResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevopsProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectResponse) SetHeaders(v map[string]*string) *UpdateDevopsProjectResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevopsProjectResponse) SetStatusCode(v int32) *UpdateDevopsProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDevopsProjectResponse) SetBody(v *UpdateDevopsProjectResponseBody) *UpdateDevopsProjectResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectSprintRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SprintId    *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s UpdateDevopsProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectSprintRequest) SetDescription(v string) *UpdateDevopsProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetDueDate(v string) *UpdateDevopsProjectSprintRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetExecutorId(v string) *UpdateDevopsProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetName(v string) *UpdateDevopsProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetOrgId(v string) *UpdateDevopsProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetProjectId(v string) *UpdateDevopsProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetSprintId(v string) *UpdateDevopsProjectSprintRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetStartDate(v string) *UpdateDevopsProjectSprintRequest {
	s.StartDate = &v
	return s
}

type UpdateDevopsProjectSprintResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectSprintResponseBody) SetErrorCode(v string) *UpdateDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponseBody) SetErrorMsg(v string) *UpdateDevopsProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponseBody) SetObject(v bool) *UpdateDevopsProjectSprintResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponseBody) SetRequestId(v string) *UpdateDevopsProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponseBody) SetSuccessful(v bool) *UpdateDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type UpdateDevopsProjectSprintResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevopsProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectSprintResponse) SetHeaders(v map[string]*string) *UpdateDevopsProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevopsProjectSprintResponse) SetStatusCode(v int32) *UpdateDevopsProjectSprintResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponse) SetBody(v *UpdateDevopsProjectSprintResponseBody) *UpdateDevopsProjectSprintResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectTaskRequest struct {
	Content                *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DueDate                *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId             *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Note                   *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OrgId                  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ParentTaskId           *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	Priority               *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId              *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScenarioFiieldConfigId *string `json:"ScenarioFiieldConfigId,omitempty" xml:"ScenarioFiieldConfigId,omitempty"`
	SprintId               *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate              *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TaskFlowStatusId       *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Visible                *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s UpdateDevopsProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectTaskRequest) SetContent(v string) *UpdateDevopsProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetDueDate(v string) *UpdateDevopsProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetExecutorId(v string) *UpdateDevopsProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetNote(v string) *UpdateDevopsProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetOrgId(v string) *UpdateDevopsProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetParentTaskId(v string) *UpdateDevopsProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetPriority(v int32) *UpdateDevopsProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetProjectId(v string) *UpdateDevopsProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetScenarioFiieldConfigId(v string) *UpdateDevopsProjectTaskRequest {
	s.ScenarioFiieldConfigId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetSprintId(v string) *UpdateDevopsProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetStartDate(v string) *UpdateDevopsProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetTaskFlowStatusId(v string) *UpdateDevopsProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetTaskId(v string) *UpdateDevopsProjectTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetVisible(v string) *UpdateDevopsProjectTaskRequest {
	s.Visible = &v
	return s
}

type UpdateDevopsProjectTaskResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectTaskResponseBody) SetErrorCode(v string) *UpdateDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponseBody) SetErrorMsg(v string) *UpdateDevopsProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponseBody) SetObject(v bool) *UpdateDevopsProjectTaskResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponseBody) SetRequestId(v string) *UpdateDevopsProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponseBody) SetSuccessful(v bool) *UpdateDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type UpdateDevopsProjectTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDevopsProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectTaskResponse) SetHeaders(v map[string]*string) *UpdateDevopsProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateDevopsProjectTaskResponse) SetStatusCode(v int32) *UpdateDevopsProjectTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponse) SetBody(v *UpdateDevopsProjectTaskResponseBody) *UpdateDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type UpdatePipelineEnvVarsRequest struct {
	EnvVars    *string `json:"EnvVars,omitempty" xml:"EnvVars,omitempty"`
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s UpdatePipelineEnvVarsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineEnvVarsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineEnvVarsRequest) SetEnvVars(v string) *UpdatePipelineEnvVarsRequest {
	s.EnvVars = &v
	return s
}

func (s *UpdatePipelineEnvVarsRequest) SetOrgId(v string) *UpdatePipelineEnvVarsRequest {
	s.OrgId = &v
	return s
}

func (s *UpdatePipelineEnvVarsRequest) SetPipelineId(v int64) *UpdatePipelineEnvVarsRequest {
	s.PipelineId = &v
	return s
}

type UpdatePipelineEnvVarsResponseBody struct {
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePipelineEnvVarsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineEnvVarsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineEnvVarsResponseBody) SetPipelineId(v int64) *UpdatePipelineEnvVarsResponseBody {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineEnvVarsResponseBody) SetRequestId(v string) *UpdatePipelineEnvVarsResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePipelineEnvVarsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineEnvVarsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineEnvVarsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineEnvVarsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineEnvVarsResponse) SetHeaders(v map[string]*string) *UpdatePipelineEnvVarsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineEnvVarsResponse) SetStatusCode(v int32) *UpdatePipelineEnvVarsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineEnvVarsResponse) SetBody(v *UpdatePipelineEnvVarsResponseBody) *UpdatePipelineEnvVarsResponse {
	s.Body = v
	return s
}

type UpdatePipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s UpdatePipelineMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineMemberRequest) SetOrgId(v string) *UpdatePipelineMemberRequest {
	s.OrgId = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetPipelineId(v int64) *UpdatePipelineMemberRequest {
	s.PipelineId = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetRoleName(v string) *UpdatePipelineMemberRequest {
	s.RoleName = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetUserId(v string) *UpdatePipelineMemberRequest {
	s.UserId = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetUserPk(v string) *UpdatePipelineMemberRequest {
	s.UserPk = &v
	return s
}

type UpdatePipelineMemberResponseBody struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineMemberResponseBody) SetErrorCode(v string) *UpdatePipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetErrorMessage(v string) *UpdatePipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetObject(v bool) *UpdatePipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetRequestId(v string) *UpdatePipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetSuccess(v bool) *UpdatePipelineMemberResponseBody {
	s.Success = &v
	return s
}

type UpdatePipelineMemberResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePipelineMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdatePipelineMemberResponse) SetHeaders(v map[string]*string) *UpdatePipelineMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdatePipelineMemberResponse) SetStatusCode(v int32) *UpdatePipelineMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePipelineMemberResponse) SetBody(v *UpdatePipelineMemberResponseBody) *UpdatePipelineMemberResponse {
	s.Body = v
	return s
}

type UpdateTaskDetailRequest struct {
	AddInvolvers      *string `json:"AddInvolvers,omitempty" xml:"AddInvolvers,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CustomFieldId     *string `json:"CustomFieldId,omitempty" xml:"CustomFieldId,omitempty"`
	CustomFieldValues *string `json:"CustomFieldValues,omitempty" xml:"CustomFieldValues,omitempty"`
	DelInvolvers      *string `json:"DelInvolvers,omitempty" xml:"DelInvolvers,omitempty"`
	DueDate           *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ExecutorId        *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Note              *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OrgId             *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Priority          *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SprintId          *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint        *string `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	TagIds            *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	TaskFlowStatusId  *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WorkTimes         *int64  `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
}

func (s UpdateTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskDetailRequest) SetAddInvolvers(v string) *UpdateTaskDetailRequest {
	s.AddInvolvers = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetContent(v string) *UpdateTaskDetailRequest {
	s.Content = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetCustomFieldId(v string) *UpdateTaskDetailRequest {
	s.CustomFieldId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetCustomFieldValues(v string) *UpdateTaskDetailRequest {
	s.CustomFieldValues = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetDelInvolvers(v string) *UpdateTaskDetailRequest {
	s.DelInvolvers = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetDueDate(v string) *UpdateTaskDetailRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetExecutorId(v string) *UpdateTaskDetailRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetNote(v string) *UpdateTaskDetailRequest {
	s.Note = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetOrgId(v string) *UpdateTaskDetailRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetPriority(v int64) *UpdateTaskDetailRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetProjectId(v string) *UpdateTaskDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetSprintId(v string) *UpdateTaskDetailRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetStartDate(v string) *UpdateTaskDetailRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetStoryPoint(v string) *UpdateTaskDetailRequest {
	s.StoryPoint = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetTagIds(v string) *UpdateTaskDetailRequest {
	s.TagIds = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetTaskFlowStatusId(v string) *UpdateTaskDetailRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetTaskId(v string) *UpdateTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetWorkTimes(v int64) *UpdateTaskDetailRequest {
	s.WorkTimes = &v
	return s
}

type UpdateTaskDetailResponseBody struct {
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskDetailResponseBody) SetErrorCode(v string) *UpdateTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskDetailResponseBody) SetErrorMsg(v string) *UpdateTaskDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateTaskDetailResponseBody) SetObject(v bool) *UpdateTaskDetailResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateTaskDetailResponseBody) SetRequestId(v string) *UpdateTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTaskDetailResponseBody) SetSuccessful(v bool) *UpdateTaskDetailResponseBody {
	s.Successful = &v
	return s
}

type UpdateTaskDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskDetailResponse) SetHeaders(v map[string]*string) *UpdateTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskDetailResponse) SetStatusCode(v int32) *UpdateTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskDetailResponse) SetBody(v *UpdateTaskDetailResponseBody) *UpdateTaskDetailResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("devops-rdc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddCodeupSourceToPipelineWithOptions(request *AddCodeupSourceToPipelineRequest, runtime *util.RuntimeOptions) (_result *AddCodeupSourceToPipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CodeBranch)) {
		query["CodeBranch"] = request.CodeBranch
	}

	if !tea.BoolValue(util.IsUnset(request.CodePath)) {
		query["CodePath"] = request.CodePath
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCodeupSourceToPipeline"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCodeupSourceToPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddCodeupSourceToPipeline(request *AddCodeupSourceToPipelineRequest) (_result *AddCodeupSourceToPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCodeupSourceToPipelineResponse{}
	_body, _err := client.AddCodeupSourceToPipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchInsertMembersWithOptions(request *BatchInsertMembersRequest, runtime *util.RuntimeOptions) (_result *BatchInsertMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["Members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.RealPk)) {
		body["RealPk"] = request.RealPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchInsertMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchInsertMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchInsertMembers(request *BatchInsertMembersRequest) (_result *BatchInsertMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchInsertMembersResponse{}
	_body, _err := client.BatchInsertMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelPipelineWithOptions(request *CancelPipelineRequest, runtime *util.RuntimeOptions) (_result *CancelPipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		body["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelPipeline"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelPipeline(request *CancelPipelineRequest) (_result *CancelPipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelPipelineResponse{}
	_body, _err := client.CancelPipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCommonGroupWithOptions(request *CreateCommonGroupRequest, runtime *util.RuntimeOptions) (_result *CreateCommonGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SmartGroupId)) {
		body["SmartGroupId"] = request.SmartGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCommonGroup"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCommonGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCommonGroup(request *CreateCommonGroupRequest) (_result *CreateCommonGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCommonGroupResponse{}
	_body, _err := client.CreateCommonGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCredentialWithOptions(request *CreateCredentialRequest, runtime *util.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCredential"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCredential(request *CreateCredentialRequest) (_result *CreateCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CreateCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevopsProjectWithOptions(request *CreateDevopsProjectRequest, runtime *util.RuntimeOptions) (_result *CreateDevopsProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevopsProject"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDevopsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevopsProject(request *CreateDevopsProjectRequest) (_result *CreateDevopsProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevopsProjectResponse{}
	_body, _err := client.CreateDevopsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevopsProjectSprintWithOptions(request *CreateDevopsProjectSprintRequest, runtime *util.RuntimeOptions) (_result *CreateDevopsProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["DueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevopsProjectSprint"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDevopsProjectSprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevopsProjectSprint(request *CreateDevopsProjectSprintRequest) (_result *CreateDevopsProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevopsProjectSprintResponse{}
	_body, _err := client.CreateDevopsProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevopsProjectTaskWithOptions(request *CreateDevopsProjectTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDevopsProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["DueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioFieldConfigId)) {
		body["ScenarioFieldConfigId"] = request.ScenarioFieldConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlowStatusId)) {
		body["TaskFlowStatusId"] = request.TaskFlowStatusId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskListId)) {
		body["TaskListId"] = request.TaskListId
	}

	if !tea.BoolValue(util.IsUnset(request.Visible)) {
		body["Visible"] = request.Visible
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDevopsProjectTask"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDevopsProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevopsProjectTask(request *CreateDevopsProjectTaskRequest) (_result *CreateDevopsProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevopsProjectTaskResponse{}
	_body, _err := client.CreateDevopsProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePipelineFromTemplateWithOptions(request *CreatePipelineFromTemplateRequest, runtime *util.RuntimeOptions) (_result *CreatePipelineFromTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineName)) {
		query["PipelineName"] = request.PipelineName
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineTemplateId)) {
		query["PipelineTemplateId"] = request.PipelineTemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePipelineFromTemplate"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePipelineFromTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipelineFromTemplate(request *CreatePipelineFromTemplateRequest) (_result *CreatePipelineFromTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePipelineFromTemplateResponse{}
	_body, _err := client.CreatePipelineFromTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceConnectionWithOptions(request *CreateServiceConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateServiceConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConnectionType)) {
		body["ServiceConnectionType"] = request.ServiceConnectionType
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceConnection"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceConnection(request *CreateServiceConnectionRequest) (_result *CreateServiceConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceConnectionResponse{}
	_body, _err := client.CreateServiceConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCommonGroupWithOptions(request *DeleteCommonGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteCommonGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommonGroupId)) {
		body["CommonGroupId"] = request.CommonGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCommonGroup"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommonGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCommonGroup(request *DeleteCommonGroupRequest) (_result *DeleteCommonGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommonGroupResponse{}
	_body, _err := client.DeleteCommonGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevopsOrganizationMembersWithOptions(request *DeleteDevopsOrganizationMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteDevopsOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.RealPk)) {
		body["RealPk"] = request.RealPk
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevopsOrganizationMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevopsOrganizationMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevopsOrganizationMembers(request *DeleteDevopsOrganizationMembersRequest) (_result *DeleteDevopsOrganizationMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevopsOrganizationMembersResponse{}
	_body, _err := client.DeleteDevopsOrganizationMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevopsProjectWithOptions(request *DeleteDevopsProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteDevopsProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevopsProject"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevopsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevopsProject(request *DeleteDevopsProjectRequest) (_result *DeleteDevopsProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevopsProjectResponse{}
	_body, _err := client.DeleteDevopsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevopsProjectMembersWithOptions(request *DeleteDevopsProjectMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteDevopsProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		body["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevopsProjectMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevopsProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevopsProjectMembers(request *DeleteDevopsProjectMembersRequest) (_result *DeleteDevopsProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevopsProjectMembersResponse{}
	_body, _err := client.DeleteDevopsProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevopsProjectSprintWithOptions(request *DeleteDevopsProjectSprintRequest, runtime *util.RuntimeOptions) (_result *DeleteDevopsProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevopsProjectSprint"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevopsProjectSprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevopsProjectSprint(request *DeleteDevopsProjectSprintRequest) (_result *DeleteDevopsProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevopsProjectSprintResponse{}
	_body, _err := client.DeleteDevopsProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDevopsProjectTaskWithOptions(request *DeleteDevopsProjectTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteDevopsProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDevopsProjectTask"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDevopsProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDevopsProjectTask(request *DeleteDevopsProjectTaskRequest) (_result *DeleteDevopsProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDevopsProjectTaskResponse{}
	_body, _err := client.DeleteDevopsProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePipelineMemberWithOptions(request *DeletePipelineMemberRequest, runtime *util.RuntimeOptions) (_result *DeletePipelineMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePipelineMember"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePipelineMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePipelineMember(request *DeletePipelineMemberRequest) (_result *DeletePipelineMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePipelineMemberResponse{}
	_body, _err := client.DeletePipelineMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecutePipelineWithOptions(request *ExecutePipelineRequest, runtime *util.RuntimeOptions) (_result *ExecutePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecutePipeline"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecutePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecutePipeline(request *ExecutePipelineRequest) (_result *ExecutePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecutePipelineResponse{}
	_body, _err := client.ExecutePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevopsOrganizationMembersWithOptions(request *GetDevopsOrganizationMembersRequest, runtime *util.RuntimeOptions) (_result *GetDevopsOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevopsOrganizationMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevopsOrganizationMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevopsOrganizationMembers(request *GetDevopsOrganizationMembersRequest) (_result *GetDevopsOrganizationMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevopsOrganizationMembersResponse{}
	_body, _err := client.GetDevopsOrganizationMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevopsProjectInfoWithOptions(request *GetDevopsProjectInfoRequest, runtime *util.RuntimeOptions) (_result *GetDevopsProjectInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevopsProjectInfo"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevopsProjectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevopsProjectInfo(request *GetDevopsProjectInfoRequest) (_result *GetDevopsProjectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevopsProjectInfoResponse{}
	_body, _err := client.GetDevopsProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevopsProjectMembersWithOptions(request *GetDevopsProjectMembersRequest, runtime *util.RuntimeOptions) (_result *GetDevopsProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		body["PageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevopsProjectMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevopsProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevopsProjectMembers(request *GetDevopsProjectMembersRequest) (_result *GetDevopsProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevopsProjectMembersResponse{}
	_body, _err := client.GetDevopsProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevopsProjectSprintInfoWithOptions(request *GetDevopsProjectSprintInfoRequest, runtime *util.RuntimeOptions) (_result *GetDevopsProjectSprintInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevopsProjectSprintInfo"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevopsProjectSprintInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevopsProjectSprintInfo(request *GetDevopsProjectSprintInfoRequest) (_result *GetDevopsProjectSprintInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevopsProjectSprintInfoResponse{}
	_body, _err := client.GetDevopsProjectSprintInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDevopsProjectTaskInfoWithOptions(request *GetDevopsProjectTaskInfoRequest, runtime *util.RuntimeOptions) (_result *GetDevopsProjectTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDevopsProjectTaskInfo"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDevopsProjectTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDevopsProjectTaskInfo(request *GetDevopsProjectTaskInfoRequest) (_result *GetDevopsProjectTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDevopsProjectTaskInfoResponse{}
	_body, _err := client.GetDevopsProjectTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineInstHistoryWithOptions(request *GetPipelineInstHistoryRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		body["PageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineInstHistory"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineInstHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineInstHistory(request *GetPipelineInstHistoryRequest) (_result *GetPipelineInstHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineInstHistoryResponse{}
	_body, _err := client.GetPipelineInstHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineInstanceBuildNumberStatusWithOptions(request *GetPipelineInstanceBuildNumberStatusRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstanceBuildNumberStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildNum)) {
		body["BuildNum"] = request.BuildNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineInstanceBuildNumberStatus"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineInstanceBuildNumberStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineInstanceBuildNumberStatus(request *GetPipelineInstanceBuildNumberStatusRequest) (_result *GetPipelineInstanceBuildNumberStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineInstanceBuildNumberStatusResponse{}
	_body, _err := client.GetPipelineInstanceBuildNumberStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineInstanceGroupStatusWithOptions(request *GetPipelineInstanceGroupStatusRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstanceGroupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		body["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineInstanceGroupStatus"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineInstanceGroupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineInstanceGroupStatus(request *GetPipelineInstanceGroupStatusRequest) (_result *GetPipelineInstanceGroupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineInstanceGroupStatusResponse{}
	_body, _err := client.GetPipelineInstanceGroupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineInstanceInfoWithOptions(request *GetPipelineInstanceInfoRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstanceInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		body["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineInstanceInfo"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineInstanceInfo(request *GetPipelineInstanceInfoRequest) (_result *GetPipelineInstanceInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineInstanceInfoResponse{}
	_body, _err := client.GetPipelineInstanceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineInstanceStatusWithOptions(request *GetPipelineInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowInstanceId)) {
		query["FlowInstanceId"] = request.FlowInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineInstanceStatus"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineInstanceStatus(request *GetPipelineInstanceStatusRequest) (_result *GetPipelineInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineInstanceStatusResponse{}
	_body, _err := client.GetPipelineInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineLogWithOptions(request *GetPipelineLogRequest, runtime *util.RuntimeOptions) (_result *GetPipelineLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineLog"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineLog(request *GetPipelineLogRequest) (_result *GetPipelineLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineLogResponse{}
	_body, _err := client.GetPipelineLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipelineStepLogWithOptions(request *GetPipelineStepLogRequest, runtime *util.RuntimeOptions) (_result *GetPipelineStepLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		body["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		body["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.StepIndex)) {
		body["StepIndex"] = request.StepIndex
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipelineStepLog"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipelineStepLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipelineStepLog(request *GetPipelineStepLogRequest) (_result *GetPipelineStepLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipelineStepLogResponse{}
	_body, _err := client.GetPipelineStepLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPipleineLatestInstanceStatusWithOptions(request *GetPipleineLatestInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *GetPipleineLatestInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPipleineLatestInstanceStatus"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPipleineLatestInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPipleineLatestInstanceStatus(request *GetPipleineLatestInstanceStatusRequest) (_result *GetPipleineLatestInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPipleineLatestInstanceStatusResponse{}
	_body, _err := client.GetPipleineLatestInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectOptionWithOptions(request *GetProjectOptionRequest, runtime *util.RuntimeOptions) (_result *GetProjectOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProjectOption"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectOption(request *GetProjectOptionRequest) (_result *GetProjectOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectOptionResponse{}
	_body, _err := client.GetProjectOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskDetailActivityWithOptions(request *GetTaskDetailActivityRequest, runtime *util.RuntimeOptions) (_result *GetTaskDetailActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskDetailActivity"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskDetailActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskDetailActivity(request *GetTaskDetailActivityRequest) (_result *GetTaskDetailActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskDetailActivityResponse{}
	_body, _err := client.GetTaskDetailActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskDetailBaseWithOptions(request *GetTaskDetailBaseRequest, runtime *util.RuntimeOptions) (_result *GetTaskDetailBaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskDetailBase"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskDetailBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskDetailBase(request *GetTaskDetailBaseRequest) (_result *GetTaskDetailBaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskDetailBaseResponse{}
	_body, _err := client.GetTaskDetailBaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskListFilterWithOptions(request *GetTaskListFilterRequest, runtime *util.RuntimeOptions) (_result *GetTaskListFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.DueDateEnd)) {
		body["DueDateEnd"] = request.DueDateEnd
	}

	if !tea.BoolValue(util.IsUnset(request.DueDateStart)) {
		body["DueDateStart"] = request.DueDateStart
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.InvolveMembers)) {
		body["InvolveMembers"] = request.InvolveMembers
	}

	if !tea.BoolValue(util.IsUnset(request.IsDone)) {
		body["IsDone"] = request.IsDone
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		body["ObjectType"] = request.ObjectType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OrderCondition)) {
		body["OrderCondition"] = request.OrderCondition
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		body["PageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioFieldConfigId)) {
		body["ScenarioFieldConfigId"] = request.ScenarioFieldConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		body["TagId"] = request.TagId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlowStatusId)) {
		body["TaskFlowStatusId"] = request.TaskFlowStatusId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskListFilter"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskListFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskListFilter(request *GetTaskListFilterRequest) (_result *GetTaskListFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskListFilterResponse{}
	_body, _err := client.GetTaskListFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertPipelineMemberWithOptions(request *InsertPipelineMemberRequest, runtime *util.RuntimeOptions) (_result *InsertPipelineMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertPipelineMember"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertPipelineMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertPipelineMember(request *InsertPipelineMemberRequest) (_result *InsertPipelineMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertPipelineMemberResponse{}
	_body, _err := client.InsertPipelineMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertProjectMembersWithOptions(request *InsertProjectMembersRequest, runtime *util.RuntimeOptions) (_result *InsertProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		body["Members"] = request.Members
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertProjectMembers"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertProjectMembers(request *InsertProjectMembersRequest) (_result *InsertProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertProjectMembersResponse{}
	_body, _err := client.InsertProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommonGroupWithOptions(request *ListCommonGroupRequest, runtime *util.RuntimeOptions) (_result *ListCommonGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SmartGroupId)) {
		body["SmartGroupId"] = request.SmartGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCommonGroup"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCommonGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommonGroup(request *ListCommonGroupRequest) (_result *ListCommonGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommonGroupResponse{}
	_body, _err := client.ListCommonGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCredentialsWithOptions(request *ListCredentialsRequest, runtime *util.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCredentials"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCredentials(request *ListCredentialsRequest) (_result *ListCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCredentialsResponse{}
	_body, _err := client.ListCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsProjectSprintsWithOptions(request *ListDevopsProjectSprintsRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectSprintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		body["PageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsProjectSprints"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsProjectSprintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjectSprints(request *ListDevopsProjectSprintsRequest) (_result *ListDevopsProjectSprintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectSprintsResponse{}
	_body, _err := client.ListDevopsProjectSprintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskFlowWithOptions(request *ListDevopsProjectTaskFlowRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectTaskFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsProjectTaskFlow"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsProjectTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskFlow(request *ListDevopsProjectTaskFlowRequest) (_result *ListDevopsProjectTaskFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectTaskFlowResponse{}
	_body, _err := client.ListDevopsProjectTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskFlowStatusWithOptions(request *ListDevopsProjectTaskFlowStatusRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectTaskFlowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlowId)) {
		body["TaskFlowId"] = request.TaskFlowId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsProjectTaskFlowStatus"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsProjectTaskFlowStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskFlowStatus(request *ListDevopsProjectTaskFlowStatusRequest) (_result *ListDevopsProjectTaskFlowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectTaskFlowStatusResponse{}
	_body, _err := client.ListDevopsProjectTaskFlowStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskListWithOptions(request *ListDevopsProjectTaskListRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsProjectTaskList"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsProjectTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjectTaskList(request *ListDevopsProjectTaskListRequest) (_result *ListDevopsProjectTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectTaskListResponse{}
	_body, _err := client.ListDevopsProjectTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsProjectsWithOptions(request *ListDevopsProjectsRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageToken)) {
		body["PageToken"] = request.PageToken
	}

	if !tea.BoolValue(util.IsUnset(request.SelectBy)) {
		body["SelectBy"] = request.SelectBy
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsProjects"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjects(request *ListDevopsProjectsRequest) (_result *ListDevopsProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectsResponse{}
	_body, _err := client.ListDevopsProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDevopsScenarioFieldConfigWithOptions(request *ListDevopsScenarioFieldConfigRequest, runtime *util.RuntimeOptions) (_result *ListDevopsScenarioFieldConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDevopsScenarioFieldConfig"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDevopsScenarioFieldConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsScenarioFieldConfig(request *ListDevopsScenarioFieldConfigRequest) (_result *ListDevopsScenarioFieldConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsScenarioFieldConfigResponse{}
	_body, _err := client.ListDevopsScenarioFieldConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelineTemplatesWithOptions(request *ListPipelineTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListPipelineTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelineTemplates"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelineTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelineTemplates(request *ListPipelineTemplatesRequest) (_result *ListPipelineTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPipelineTemplatesResponse{}
	_body, _err := client.ListPipelineTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPipelinesWithOptions(request *ListPipelinesRequest, runtime *util.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateEndTime)) {
		body["CreateEndTime"] = request.CreateEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.CreateStartTime)) {
		body["CreateStartTime"] = request.CreateStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Creators)) {
		body["Creators"] = request.Creators
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteEndTime)) {
		body["ExecuteEndTime"] = request.ExecuteEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecuteStartTime)) {
		body["ExecuteStartTime"] = request.ExecuteStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Operators)) {
		body["Operators"] = request.Operators
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		body["PageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineName)) {
		body["PipelineName"] = request.PipelineName
	}

	if !tea.BoolValue(util.IsUnset(request.ResultStatusList)) {
		body["ResultStatusList"] = request.ResultStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPipelines"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPipelines(request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectCustomFieldsWithOptions(request *ListProjectCustomFieldsRequest, runtime *util.RuntimeOptions) (_result *ListProjectCustomFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectCustomFields"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectCustomFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectCustomFields(request *ListProjectCustomFieldsRequest) (_result *ListProjectCustomFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectCustomFieldsResponse{}
	_body, _err := client.ListProjectCustomFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceConnectionsWithOptions(request *ListServiceConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListServiceConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ScType)) {
		body["ScType"] = request.ScType
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceConnections"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceConnections(request *ListServiceConnectionsRequest) (_result *ListServiceConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.ListServiceConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSmartGroupWithOptions(request *ListSmartGroupRequest, runtime *util.RuntimeOptions) (_result *ListSmartGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSmartGroup"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSmartGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSmartGroup(request *ListSmartGroupRequest) (_result *ListSmartGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSmartGroupResponse{}
	_body, _err := client.ListSmartGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserOrganizationWithOptions(request *ListUserOrganizationRequest, runtime *util.RuntimeOptions) (_result *ListUserOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RealPk)) {
		body["RealPk"] = request.RealPk
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserOrganization"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserOrganizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserOrganization(request *ListUserOrganizationRequest) (_result *ListUserOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserOrganizationResponse{}
	_body, _err := client.ListUserOrganizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferPipelineOwnerWithOptions(request *TransferPipelineOwnerRequest, runtime *util.RuntimeOptions) (_result *TransferPipelineOwnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewOwnerId)) {
		body["NewOwnerId"] = request.NewOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferPipelineOwner"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferPipelineOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferPipelineOwner(request *TransferPipelineOwnerRequest) (_result *TransferPipelineOwnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferPipelineOwnerResponse{}
	_body, _err := client.TransferPipelineOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCommonGroupWithOptions(request *UpdateCommonGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateCommonGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommonGroupId)) {
		body["CommonGroupId"] = request.CommonGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SmartGroupId)) {
		body["SmartGroupId"] = request.SmartGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCommonGroup"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCommonGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCommonGroup(request *UpdateCommonGroupRequest) (_result *UpdateCommonGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCommonGroupResponse{}
	_body, _err := client.UpdateCommonGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevopsProjectWithOptions(request *UpdateDevopsProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateDevopsProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDevopsProject"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDevopsProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevopsProject(request *UpdateDevopsProjectRequest) (_result *UpdateDevopsProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevopsProjectResponse{}
	_body, _err := client.UpdateDevopsProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevopsProjectSprintWithOptions(request *UpdateDevopsProjectSprintRequest, runtime *util.RuntimeOptions) (_result *UpdateDevopsProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["DueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDevopsProjectSprint"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDevopsProjectSprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevopsProjectSprint(request *UpdateDevopsProjectSprintRequest) (_result *UpdateDevopsProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevopsProjectSprintResponse{}
	_body, _err := client.UpdateDevopsProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDevopsProjectTaskWithOptions(request *UpdateDevopsProjectTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateDevopsProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["DueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskId)) {
		body["ParentTaskId"] = request.ParentTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioFiieldConfigId)) {
		body["ScenarioFiieldConfigId"] = request.ScenarioFiieldConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlowStatusId)) {
		body["TaskFlowStatusId"] = request.TaskFlowStatusId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Visible)) {
		body["Visible"] = request.Visible
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDevopsProjectTask"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDevopsProjectTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDevopsProjectTask(request *UpdateDevopsProjectTaskRequest) (_result *UpdateDevopsProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDevopsProjectTaskResponse{}
	_body, _err := client.UpdateDevopsProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineEnvVarsWithOptions(request *UpdatePipelineEnvVarsRequest, runtime *util.RuntimeOptions) (_result *UpdatePipelineEnvVarsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnvVars)) {
		query["EnvVars"] = request.EnvVars
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineEnvVars"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineEnvVarsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineEnvVars(request *UpdatePipelineEnvVarsRequest) (_result *UpdatePipelineEnvVarsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePipelineEnvVarsResponse{}
	_body, _err := client.UpdatePipelineEnvVarsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePipelineMemberWithOptions(request *UpdatePipelineMemberRequest, runtime *util.RuntimeOptions) (_result *UpdatePipelineMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserPk)) {
		body["UserPk"] = request.UserPk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePipelineMember"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePipelineMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePipelineMember(request *UpdatePipelineMemberRequest) (_result *UpdatePipelineMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePipelineMemberResponse{}
	_body, _err := client.UpdatePipelineMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTaskDetailWithOptions(request *UpdateTaskDetailRequest, runtime *util.RuntimeOptions) (_result *UpdateTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddInvolvers)) {
		body["AddInvolvers"] = request.AddInvolvers
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFieldId)) {
		body["CustomFieldId"] = request.CustomFieldId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomFieldValues)) {
		body["CustomFieldValues"] = request.CustomFieldValues
	}

	if !tea.BoolValue(util.IsUnset(request.DelInvolvers)) {
		body["DelInvolvers"] = request.DelInvolvers
	}

	if !tea.BoolValue(util.IsUnset(request.DueDate)) {
		body["DueDate"] = request.DueDate
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutorId)) {
		body["ExecutorId"] = request.ExecutorId
	}

	if !tea.BoolValue(util.IsUnset(request.Note)) {
		body["Note"] = request.Note
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SprintId)) {
		body["SprintId"] = request.SprintId
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StoryPoint)) {
		body["StoryPoint"] = request.StoryPoint
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		body["TagIds"] = request.TagIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskFlowStatusId)) {
		body["TaskFlowStatusId"] = request.TaskFlowStatusId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkTimes)) {
		body["WorkTimes"] = request.WorkTimes
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTaskDetail"),
		Version:     tea.String("2020-03-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTaskDetail(request *UpdateTaskDetailRequest) (_result *UpdateTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTaskDetailResponse{}
	_body, _err := client.UpdateTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
