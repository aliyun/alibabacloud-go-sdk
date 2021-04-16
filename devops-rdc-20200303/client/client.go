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

type BatchInsertMembersRequest struct {
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Members *string `json:"Members,omitempty" xml:"Members,omitempty"`
	RealPk  *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s BatchInsertMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertMembersRequest) GoString() string {
	return s.String()
}

func (s *BatchInsertMembersRequest) SetOrgId(v string) *BatchInsertMembersRequest {
	s.OrgId = &v
	return s
}

func (s *BatchInsertMembersRequest) SetMembers(v string) *BatchInsertMembersRequest {
	s.Members = &v
	return s
}

func (s *BatchInsertMembersRequest) SetRealPk(v string) *BatchInsertMembersRequest {
	s.RealPk = &v
	return s
}

type BatchInsertMembersResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchInsertMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchInsertMembersResponseBody) GoString() string {
	return s.String()
}

func (s *BatchInsertMembersResponseBody) SetRequestId(v string) *BatchInsertMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetObject(v bool) *BatchInsertMembersResponseBody {
	s.Object = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetErrorCode(v string) *BatchInsertMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetErrorMessage(v string) *BatchInsertMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BatchInsertMembersResponseBody) SetSuccess(v bool) *BatchInsertMembersResponseBody {
	s.Success = &v
	return s
}

type BatchInsertMembersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchInsertMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *BatchInsertMembersResponse) SetBody(v *BatchInsertMembersResponseBody) *BatchInsertMembersResponse {
	s.Body = v
	return s
}

type CancelPipelineRequest struct {
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CancelPipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineRequest) GoString() string {
	return s.String()
}

func (s *CancelPipelineRequest) SetOrgId(v string) *CancelPipelineRequest {
	s.OrgId = &v
	return s
}

func (s *CancelPipelineRequest) SetPipelineId(v int64) *CancelPipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *CancelPipelineRequest) SetFlowInstanceId(v int64) *CancelPipelineRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *CancelPipelineRequest) SetUserPk(v string) *CancelPipelineRequest {
	s.UserPk = &v
	return s
}

type CancelPipelineResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelPipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelPipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CancelPipelineResponseBody) SetRequestId(v string) *CancelPipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelPipelineResponseBody) SetObject(v bool) *CancelPipelineResponseBody {
	s.Object = &v
	return s
}

func (s *CancelPipelineResponseBody) SetErrorCode(v string) *CancelPipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CancelPipelineResponseBody) SetErrorMessage(v string) *CancelPipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CancelPipelineResponseBody) SetSuccess(v bool) *CancelPipelineResponseBody {
	s.Success = &v
	return s
}

type CancelPipelineResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CancelPipelineResponse) SetBody(v *CancelPipelineResponseBody) *CancelPipelineResponse {
	s.Body = v
	return s
}

type CheckAliyunAccountExistsRequest struct {
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CheckAliyunAccountExistsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunAccountExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckAliyunAccountExistsRequest) SetUserPk(v string) *CheckAliyunAccountExistsRequest {
	s.UserPk = &v
	return s
}

type CheckAliyunAccountExistsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CheckAliyunAccountExistsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunAccountExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAliyunAccountExistsResponseBody) SetRequestId(v string) *CheckAliyunAccountExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAliyunAccountExistsResponseBody) SetErrorMsg(v string) *CheckAliyunAccountExistsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckAliyunAccountExistsResponseBody) SetObject(v bool) *CheckAliyunAccountExistsResponseBody {
	s.Object = &v
	return s
}

func (s *CheckAliyunAccountExistsResponseBody) SetErrorCode(v string) *CheckAliyunAccountExistsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckAliyunAccountExistsResponseBody) SetSuccessful(v bool) *CheckAliyunAccountExistsResponseBody {
	s.Successful = &v
	return s
}

type CheckAliyunAccountExistsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckAliyunAccountExistsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAliyunAccountExistsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunAccountExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckAliyunAccountExistsResponse) SetHeaders(v map[string]*string) *CheckAliyunAccountExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckAliyunAccountExistsResponse) SetBody(v *CheckAliyunAccountExistsResponseBody) *CheckAliyunAccountExistsResponse {
	s.Body = v
	return s
}

type CreateCommonGroupRequest struct {
	OrgId        *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SmartGroupId *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupRequest) SetOrgId(v string) *CreateCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *CreateCommonGroupRequest) SetProjectId(v string) *CreateCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCommonGroupRequest) SetDescription(v string) *CreateCommonGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateCommonGroupRequest) SetSmartGroupId(v string) *CreateCommonGroupRequest {
	s.SmartGroupId = &v
	return s
}

func (s *CreateCommonGroupRequest) SetName(v string) *CreateCommonGroupRequest {
	s.Name = &v
	return s
}

type CreateCommonGroupResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommonGroupResponseBody) SetRequestId(v string) *CreateCommonGroupResponseBody {
	s.RequestId = &v
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

func (s *CreateCommonGroupResponseBody) SetErrorCode(v string) *CreateCommonGroupResponseBody {
	s.ErrorCode = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCommonGroupResponse) SetBody(v *CreateCommonGroupResponseBody) *CreateCommonGroupResponse {
	s.Body = v
	return s
}

type CreateCredentialRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserPk   *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) SetOrgId(v string) *CreateCredentialRequest {
	s.OrgId = &v
	return s
}

func (s *CreateCredentialRequest) SetName(v string) *CreateCredentialRequest {
	s.Name = &v
	return s
}

func (s *CreateCredentialRequest) SetUserName(v string) *CreateCredentialRequest {
	s.UserName = &v
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

func (s *CreateCredentialRequest) SetUserPk(v string) *CreateCredentialRequest {
	s.UserPk = &v
	return s
}

type CreateCredentialResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetObject(v int64) *CreateCredentialResponseBody {
	s.Object = &v
	return s
}

func (s *CreateCredentialResponseBody) SetErrorCode(v string) *CreateCredentialResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCredentialResponseBody) SetErrorMessage(v string) *CreateCredentialResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCredentialResponseBody) SetSuccess(v bool) *CreateCredentialResponseBody {
	s.Success = &v
	return s
}

type CreateCredentialResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCredentialResponse) SetBody(v *CreateCredentialResponseBody) *CreateCredentialResponse {
	s.Body = v
	return s
}

type CreateDevopsOrganizationRequest struct {
	OrgName            *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	RealPk             *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
	DesiredMemberCount *int32  `json:"DesiredMemberCount,omitempty" xml:"DesiredMemberCount,omitempty"`
}

func (s CreateDevopsOrganizationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrganizationRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrganizationRequest) SetOrgName(v string) *CreateDevopsOrganizationRequest {
	s.OrgName = &v
	return s
}

func (s *CreateDevopsOrganizationRequest) SetSource(v string) *CreateDevopsOrganizationRequest {
	s.Source = &v
	return s
}

func (s *CreateDevopsOrganizationRequest) SetRealPk(v string) *CreateDevopsOrganizationRequest {
	s.RealPk = &v
	return s
}

func (s *CreateDevopsOrganizationRequest) SetDesiredMemberCount(v int32) *CreateDevopsOrganizationRequest {
	s.DesiredMemberCount = &v
	return s
}

type CreateDevopsOrganizationResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevopsOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrganizationResponseBody) SetRequestId(v string) *CreateDevopsOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsOrganizationResponseBody) SetObject(v string) *CreateDevopsOrganizationResponseBody {
	s.Object = &v
	return s
}

func (s *CreateDevopsOrganizationResponseBody) SetErrorCode(v string) *CreateDevopsOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsOrganizationResponseBody) SetErrorMessage(v string) *CreateDevopsOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevopsOrganizationResponseBody) SetSuccess(v bool) *CreateDevopsOrganizationResponseBody {
	s.Success = &v
	return s
}

type CreateDevopsOrganizationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevopsOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevopsOrganizationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrganizationResponse) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrganizationResponse) SetHeaders(v map[string]*string) *CreateDevopsOrganizationResponse {
	s.Headers = v
	return s
}

func (s *CreateDevopsOrganizationResponse) SetBody(v *CreateDevopsOrganizationResponseBody) *CreateDevopsOrganizationResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateDevopsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectRequest) SetOrgId(v string) *CreateDevopsProjectRequest {
	s.OrgId = &v
	return s
}

func (s *CreateDevopsProjectRequest) SetName(v string) *CreateDevopsProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectRequest) SetDescription(v string) *CreateDevopsProjectRequest {
	s.Description = &v
	return s
}

type CreateDevopsProjectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectResponseBody) SetRequestId(v string) *CreateDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetObject(v string) *CreateDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetErrorCode(v string) *CreateDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetErrorMessage(v string) *CreateDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevopsProjectResponseBody) SetSuccess(v bool) *CreateDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type CreateDevopsProjectResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDevopsProjectResponse) SetBody(v *CreateDevopsProjectResponseBody) *CreateDevopsProjectResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectSprintRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
}

func (s CreateDevopsProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintRequest) SetOrgId(v string) *CreateDevopsProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetName(v string) *CreateDevopsProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetDescription(v string) *CreateDevopsProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetProjectId(v string) *CreateDevopsProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetExecutorId(v string) *CreateDevopsProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetStartDate(v string) *CreateDevopsProjectSprintRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDevopsProjectSprintRequest) SetDueDate(v string) *CreateDevopsProjectSprintRequest {
	s.DueDate = &v
	return s
}

type CreateDevopsProjectSprintResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateDevopsProjectSprintResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBody) SetRequestId(v string) *CreateDevopsProjectSprintResponseBody {
	s.RequestId = &v
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

func (s *CreateDevopsProjectSprintResponseBody) SetErrorCode(v string) *CreateDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBody) SetSuccessful(v bool) *CreateDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type CreateDevopsProjectSprintResponseBodyObject struct {
	Status       *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                              `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                              `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                              `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Executor     *string                                              `json:"Executor,omitempty" xml:"Executor,omitempty"`
	Description  *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Accomplished *string                                              `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                                `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                              `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                              `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                              `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *CreateDevopsProjectSprintResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetStatus(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Status = &v
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

func (s *CreateDevopsProjectSprintResponseBodyObject) SetCreatorId(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetExecutor(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Executor = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetDescription(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Description = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetAccomplished(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetIsDeleted(v bool) *CreateDevopsProjectSprintResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetUpdated(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetDueDate(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetName(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Name = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetCreated(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetPlanToDo(v *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) *CreateDevopsProjectSprintResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObject) SetId(v string) *CreateDevopsProjectSprintResponseBodyObject {
	s.Id = &v
	return s
}

type CreateDevopsProjectSprintResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s CreateDevopsProjectSprintResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectSprintResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetTasks(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *CreateDevopsProjectSprintResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *CreateDevopsProjectSprintResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type CreateDevopsProjectSprintResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDevopsProjectSprintResponse) SetBody(v *CreateDevopsProjectSprintResponseBody) *CreateDevopsProjectSprintResponse {
	s.Body = v
	return s
}

type CreateDevopsProjectTaskRequest struct {
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Content               *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate             *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate               *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	TaskFlowStatusId      *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	Note                  *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Priority              *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Visible               *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
	ParentTaskId          *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	TaskListId            *string `json:"TaskListId,omitempty" xml:"TaskListId,omitempty"`
}

func (s CreateDevopsProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskRequest) SetOrgId(v string) *CreateDevopsProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetContent(v string) *CreateDevopsProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetProjectId(v string) *CreateDevopsProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetExecutorId(v string) *CreateDevopsProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetStartDate(v string) *CreateDevopsProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetDueDate(v string) *CreateDevopsProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetScenarioFieldConfigId(v string) *CreateDevopsProjectTaskRequest {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetTaskFlowStatusId(v string) *CreateDevopsProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetNote(v string) *CreateDevopsProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetPriority(v int32) *CreateDevopsProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetVisible(v string) *CreateDevopsProjectTaskRequest {
	s.Visible = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetParentTaskId(v string) *CreateDevopsProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetSprintId(v string) *CreateDevopsProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *CreateDevopsProjectTaskRequest) SetTaskListId(v string) *CreateDevopsProjectTaskRequest {
	s.TaskListId = &v
	return s
}

type CreateDevopsProjectTaskResponseBody struct {
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateDevopsProjectTaskResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                      `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskResponseBody) SetRequestId(v string) *CreateDevopsProjectTaskResponseBody {
	s.RequestId = &v
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

func (s *CreateDevopsProjectTaskResponseBody) SetErrorCode(v string) *CreateDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBody) SetSuccessful(v bool) *CreateDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type CreateDevopsProjectTaskResponseBodyObject struct {
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Priority              *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	AncestorIds           *string `json:"AncestorIds,omitempty" xml:"AncestorIds,omitempty"`
	TaskType              *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TasklistId            *string `json:"TasklistId,omitempty" xml:"TasklistId,omitempty"`
	TaskflowstatusId      *string `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	Note                  *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Updated               *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	UniqueId              *int32  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	Content               *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Rating                *int32  `json:"Rating,omitempty" xml:"Rating,omitempty"`
	Pos                   *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	StartDate             *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	StoryPoint            *string `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	CreatorId             *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	OrganizationId        *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	Visible               *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
	IsDone                *bool   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	DueDate               *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Created               *string `json:"Created,omitempty" xml:"Created,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDevopsProjectTaskResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsProjectTaskResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetExecutorId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetProjectId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetPriority(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetScenarioFieldConfigId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetAncestorIds(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.AncestorIds = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTaskType(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTasklistId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetTaskflowstatusId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetNote(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Note = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetUpdated(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetUniqueId(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.UniqueId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetContent(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Content = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetRating(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Rating = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetPos(v int32) *CreateDevopsProjectTaskResponseBodyObject {
	s.Pos = &v
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

func (s *CreateDevopsProjectTaskResponseBodyObject) SetCreatorId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetSource(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Source = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetOrganizationId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetVisible(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetIsDone(v bool) *CreateDevopsProjectTaskResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetSprintId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetDueDate(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetCreated(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateDevopsProjectTaskResponseBodyObject) SetId(v string) *CreateDevopsProjectTaskResponseBodyObject {
	s.Id = &v
	return s
}

type CreateDevopsProjectTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDevopsProjectTaskResponse) SetBody(v *CreateDevopsProjectTaskResponseBody) *CreateDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type CreatePipelineRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Pipeline *string `json:"Pipeline,omitempty" xml:"Pipeline,omitempty"`
	UserPk   *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CreatePipelineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineRequest) GoString() string {
	return s.String()
}

func (s *CreatePipelineRequest) SetOrgId(v string) *CreatePipelineRequest {
	s.OrgId = &v
	return s
}

func (s *CreatePipelineRequest) SetPipeline(v string) *CreatePipelineRequest {
	s.Pipeline = &v
	return s
}

func (s *CreatePipelineRequest) SetUserPk(v string) *CreatePipelineRequest {
	s.UserPk = &v
	return s
}

type CreatePipelineResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineResponseBody) SetObject(v int64) *CreatePipelineResponseBody {
	s.Object = &v
	return s
}

func (s *CreatePipelineResponseBody) SetErrorCode(v string) *CreatePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreatePipelineResponseBody) SetErrorMessage(v string) *CreatePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePipelineResponseBody) SetSuccess(v bool) *CreatePipelineResponseBody {
	s.Success = &v
	return s
}

type CreatePipelineResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePipelineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePipelineResponse) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponse) SetHeaders(v map[string]*string) *CreatePipelineResponse {
	s.Headers = v
	return s
}

func (s *CreatePipelineResponse) SetBody(v *CreatePipelineResponseBody) *CreatePipelineResponse {
	s.Body = v
	return s
}

type CreateServiceConnectionRequest struct {
	ServiceConnectionType *string `json:"ServiceConnectionType,omitempty" xml:"ServiceConnectionType,omitempty"`
	UserPk                *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s CreateServiceConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionRequest) SetServiceConnectionType(v string) *CreateServiceConnectionRequest {
	s.ServiceConnectionType = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetUserPk(v string) *CreateServiceConnectionRequest {
	s.UserPk = &v
	return s
}

func (s *CreateServiceConnectionRequest) SetOrgId(v string) *CreateServiceConnectionRequest {
	s.OrgId = &v
	return s
}

type CreateServiceConnectionResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceConnectionResponseBody) SetRequestId(v string) *CreateServiceConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetObject(v int64) *CreateServiceConnectionResponseBody {
	s.Object = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetErrorCode(v string) *CreateServiceConnectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetErrorMessage(v string) *CreateServiceConnectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateServiceConnectionResponseBody) SetSuccess(v bool) *CreateServiceConnectionResponseBody {
	s.Success = &v
	return s
}

type CreateServiceConnectionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateServiceConnectionResponse) SetBody(v *CreateServiceConnectionResponseBody) *CreateServiceConnectionResponse {
	s.Body = v
	return s
}

type DeleteCommonGroupRequest struct {
	OrgId         *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	CommonGroupId *string `json:"CommonGroupId,omitempty" xml:"CommonGroupId,omitempty"`
}

func (s DeleteCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupRequest) SetOrgId(v string) *DeleteCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteCommonGroupRequest) SetProjectId(v string) *DeleteCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteCommonGroupRequest) SetCommonGroupId(v string) *DeleteCommonGroupRequest {
	s.CommonGroupId = &v
	return s
}

type DeleteCommonGroupResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *DeleteCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommonGroupResponseBody) SetRequestId(v string) *DeleteCommonGroupResponseBody {
	s.RequestId = &v
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

func (s *DeleteCommonGroupResponseBody) SetErrorCode(v string) *DeleteCommonGroupResponseBody {
	s.ErrorCode = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteCommonGroupResponse) SetBody(v *DeleteCommonGroupResponseBody) *DeleteCommonGroupResponse {
	s.Body = v
	return s
}

type DeleteDevopsOrganizationMembersRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RealPk *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
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

func (s *DeleteDevopsOrganizationMembersRequest) SetUserId(v string) *DeleteDevopsOrganizationMembersRequest {
	s.UserId = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersRequest) SetRealPk(v string) *DeleteDevopsOrganizationMembersRequest {
	s.RealPk = &v
	return s
}

type DeleteDevopsOrganizationMembersResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDevopsOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetRequestId(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetObject(v bool) *DeleteDevopsOrganizationMembersResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetErrorCode(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetErrorMessage(v string) *DeleteDevopsOrganizationMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDevopsOrganizationMembersResponseBody) SetSuccess(v bool) *DeleteDevopsOrganizationMembersResponseBody {
	s.Success = &v
	return s
}

type DeleteDevopsOrganizationMembersResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevopsOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectResponseBody) SetRequestId(v string) *DeleteDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetObject(v string) *DeleteDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetErrorCode(v string) *DeleteDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetErrorMessage(v string) *DeleteDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDevopsProjectResponseBody) SetSuccess(v bool) *DeleteDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteDevopsProjectResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectMembersResponseBody) SetRequestId(v string) *DeleteDevopsProjectMembersResponseBody {
	s.RequestId = &v
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

func (s *DeleteDevopsProjectMembersResponseBody) SetErrorCode(v string) *DeleteDevopsProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectMembersResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectMembersResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevopsProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectSprintResponseBody) SetRequestId(v string) *DeleteDevopsProjectSprintResponseBody {
	s.RequestId = &v
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

func (s *DeleteDevopsProjectSprintResponseBody) SetErrorCode(v string) *DeleteDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectSprintResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectSprintResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDevopsProjectTaskResponseBody) SetRequestId(v string) *DeleteDevopsProjectTaskResponseBody {
	s.RequestId = &v
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

func (s *DeleteDevopsProjectTaskResponseBody) SetErrorCode(v string) *DeleteDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDevopsProjectTaskResponseBody) SetSuccessful(v bool) *DeleteDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type DeleteDevopsProjectTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDevopsProjectTaskResponse) SetBody(v *DeleteDevopsProjectTaskResponseBody) *DeleteDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type DeletePipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *DeletePipelineMemberRequest) SetUserPk(v string) *DeletePipelineMemberRequest {
	s.UserPk = &v
	return s
}

func (s *DeletePipelineMemberRequest) SetUserId(v string) *DeletePipelineMemberRequest {
	s.UserId = &v
	return s
}

type DeletePipelineMemberResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePipelineMemberResponseBody) SetRequestId(v string) *DeletePipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetObject(v bool) *DeletePipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetErrorCode(v string) *DeletePipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetErrorMessage(v string) *DeletePipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeletePipelineMemberResponseBody) SetSuccess(v bool) *DeletePipelineMemberResponseBody {
	s.Success = &v
	return s
}

type DeletePipelineMemberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeletePipelineMemberResponse) SetBody(v *DeletePipelineMemberResponseBody) *DeletePipelineMemberResponse {
	s.Body = v
	return s
}

type ExecutePipelineRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s *ExecutePipelineRequest) SetPipelineId(v int64) *ExecutePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *ExecutePipelineRequest) SetParameters(v string) *ExecutePipelineRequest {
	s.Parameters = &v
	return s
}

func (s *ExecutePipelineRequest) SetUserPk(v string) *ExecutePipelineRequest {
	s.UserPk = &v
	return s
}

type ExecutePipelineResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *int64  `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecutePipelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecutePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ExecutePipelineResponseBody) SetRequestId(v string) *ExecutePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetObject(v int64) *ExecutePipelineResponseBody {
	s.Object = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetErrorCode(v string) *ExecutePipelineResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetErrorMessage(v string) *ExecutePipelineResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ExecutePipelineResponseBody) SetSuccess(v bool) *ExecutePipelineResponseBody {
	s.Success = &v
	return s
}

type ExecutePipelineResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecutePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetDevopsOrganizationMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                             `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersResponseBody) SetRequestId(v string) *GetDevopsOrganizationMembersResponseBody {
	s.RequestId = &v
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

func (s *GetDevopsOrganizationMembersResponseBody) SetErrorCode(v string) *GetDevopsOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBody) SetSuccessful(v bool) *GetDevopsOrganizationMembersResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsOrganizationMembersResponseBodyObject struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetDevopsOrganizationMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsOrganizationMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetEmail(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetAvatarUrl(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetUserId(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.UserId = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetMemberId(v string) *GetDevopsOrganizationMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetDevopsOrganizationMembersResponseBodyObject) SetRole(v int32) *GetDevopsOrganizationMembersResponseBodyObject {
	s.Role = &v
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

type GetDevopsOrganizationMembersResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevopsOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoResponseBody) SetRequestId(v string) *GetDevopsProjectInfoResponseBody {
	s.RequestId = &v
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

func (s *GetDevopsProjectInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectInfoResponseBodyObject struct {
	SortMethod          *string `json:"SortMethod,omitempty" xml:"SortMethod,omitempty"`
	UniqueIdPrefix      *string `json:"UniqueIdPrefix,omitempty" xml:"UniqueIdPrefix,omitempty"`
	NormalType          *string `json:"NormalType,omitempty" xml:"NormalType,omitempty"`
	ModifierId          *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	SourceType          *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	IsTemplate          *bool   `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DefaultRoleId       *string `json:"DefaultRoleId,omitempty" xml:"DefaultRoleId,omitempty"`
	RootCollectionId    *string `json:"RootCollectionId,omitempty" xml:"RootCollectionId,omitempty"`
	IsDeleted           *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated             *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	IsArchived          *bool   `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	EndDate             *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Logo                *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	StartDate           *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Pinyin              *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	CreatorId           *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	SourceId            *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	OrganizationId      *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	IsSuspended         *bool   `json:"IsSuspended,omitempty" xml:"IsSuspended,omitempty"`
	DefaultCollectionId *string `json:"DefaultCollectionId,omitempty" xml:"DefaultCollectionId,omitempty"`
	Visibility          *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	Py                  *string `json:"Py,omitempty" xml:"Py,omitempty"`
	Category            *string `json:"Category,omitempty" xml:"Category,omitempty"`
	NextTaskUniqueId    *int32  `json:"NextTaskUniqueId,omitempty" xml:"NextTaskUniqueId,omitempty"`
	Customfields        *string `json:"Customfields,omitempty" xml:"Customfields,omitempty"`
	Created             *string `json:"Created,omitempty" xml:"Created,omitempty"`
	Id                  *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDevopsProjectInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSortMethod(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SortMethod = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetUniqueIdPrefix(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.UniqueIdPrefix = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetNormalType(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.NormalType = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetModifierId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.ModifierId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSourceType(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SourceType = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsTemplate(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsTemplate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDescription(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Description = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDefaultRoleId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.DefaultRoleId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetRootCollectionId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.RootCollectionId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsArchived(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsArchived = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetName(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetEndDate(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.EndDate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetLogo(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Logo = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetStartDate(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetPinyin(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Pinyin = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetSourceId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.SourceId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetOrganizationId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetIsSuspended(v bool) *GetDevopsProjectInfoResponseBodyObject {
	s.IsSuspended = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetDefaultCollectionId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.DefaultCollectionId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetVisibility(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Visibility = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetPy(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Py = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCategory(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Category = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetNextTaskUniqueId(v int32) *GetDevopsProjectInfoResponseBodyObject {
	s.NextTaskUniqueId = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCustomfields(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Customfields = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectInfoResponseBodyObject) SetId(v string) *GetDevopsProjectInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetDevopsProjectInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevopsProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDevopsProjectInfoResponse) SetBody(v *GetDevopsProjectInfoResponseBody) *GetDevopsProjectInfoResponse {
	s.Body = v
	return s
}

type GetDevopsProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
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

func (s *GetDevopsProjectMembersRequest) SetProjectId(v string) *GetDevopsProjectMembersRequest {
	s.ProjectId = &v
	return s
}

type GetDevopsProjectMembersResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetDevopsProjectMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersResponseBody) SetRequestId(v string) *GetDevopsProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetErrorMsg(v string) *GetDevopsProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetObject(v []*GetDevopsProjectMembersResponseBodyObject) *GetDevopsProjectMembersResponseBody {
	s.Object = v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetErrorCode(v string) *GetDevopsProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBody) SetSuccessful(v bool) *GetDevopsProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectMembersResponseBodyObject struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetDevopsProjectMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetEmail(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetAvatarUrl(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetUserId(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.UserId = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetMemberId(v string) *GetDevopsProjectMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetDevopsProjectMembersResponseBodyObject) SetRole(v int32) *GetDevopsProjectMembersResponseBodyObject {
	s.Role = &v
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

type GetDevopsProjectMembersResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevopsProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectSprintInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                         `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetRequestId(v string) *GetDevopsProjectSprintInfoResponseBody {
	s.RequestId = &v
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

func (s *GetDevopsProjectSprintInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectSprintInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectSprintInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectSprintInfoResponseBodyObject struct {
	Status       *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                               `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                               `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                               `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Accomplished *string                                               `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                                 `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                               `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                               `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                               `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetStatus(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Status = &v
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

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetAccomplished(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetDueDate(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetName(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetPlanToDo(v *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObject) SetId(v string) *GetDevopsProjectSprintInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetTasks(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *GetDevopsProjectSprintInfoResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type GetDevopsProjectSprintInfoResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevopsProjectSprintInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetDevopsProjectTaskInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                       `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetDevopsProjectTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetRequestId(v string) *GetDevopsProjectTaskInfoResponseBody {
	s.RequestId = &v
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

func (s *GetDevopsProjectTaskInfoResponseBody) SetErrorCode(v string) *GetDevopsProjectTaskInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBody) SetSuccessful(v bool) *GetDevopsProjectTaskInfoResponseBody {
	s.Successful = &v
	return s
}

type GetDevopsProjectTaskInfoResponseBodyObject struct {
	ExecutorId       *string   `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	ProjectId        *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StoryPoint       *string   `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	StartDate        *string   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	IsTopInProject   *bool     `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Priority         *string   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	CreatorId        *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	OrganizationId   *string   `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	TaskType         *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TasklistId       *string   `json:"TasklistId,omitempty" xml:"TasklistId,omitempty"`
	Visible          *string   `json:"Visible,omitempty" xml:"Visible,omitempty"`
	IsDone           *bool     `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	IsDeleted        *bool     `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	TaskflowstatusId *string   `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	Note             *string   `json:"Note,omitempty" xml:"Note,omitempty"`
	SprintId         *string   `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	Updated          *string   `json:"Updated,omitempty" xml:"Updated,omitempty"`
	InvolveMembers   []*string `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	DueDate          *string   `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Created          *string   `json:"Created,omitempty" xml:"Created,omitempty"`
	Content          *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Id               *string   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDevopsProjectTaskInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetDevopsProjectTaskInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetExecutorId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetProjectId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetStoryPoint(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetStartDate(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsTopInProject(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsTopInProject = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetPriority(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetCreatorId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetOrganizationId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTaskType(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTasklistId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetVisible(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsDone(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetIsDeleted(v bool) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetTaskflowstatusId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetNote(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Note = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetSprintId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetUpdated(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetInvolveMembers(v []*string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.InvolveMembers = v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetDueDate(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetCreated(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetContent(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Content = &v
	return s
}

func (s *GetDevopsProjectTaskInfoResponseBodyObject) SetId(v string) *GetDevopsProjectTaskInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetDevopsProjectTaskInfoResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDevopsProjectTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDevopsProjectTaskInfoResponse) SetBody(v *GetDevopsProjectTaskInfoResponseBody) *GetDevopsProjectTaskInfoResponse {
	s.Body = v
	return s
}

type GetLastWorkspaceRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	RealPk *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s GetLastWorkspaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLastWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *GetLastWorkspaceRequest) SetOrgId(v string) *GetLastWorkspaceRequest {
	s.OrgId = &v
	return s
}

func (s *GetLastWorkspaceRequest) SetRealPk(v string) *GetLastWorkspaceRequest {
	s.RealPk = &v
	return s
}

type GetLastWorkspaceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLastWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLastWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLastWorkspaceResponseBody) SetRequestId(v string) *GetLastWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLastWorkspaceResponseBody) SetObject(v string) *GetLastWorkspaceResponseBody {
	s.Object = &v
	return s
}

func (s *GetLastWorkspaceResponseBody) SetErrorCode(v string) *GetLastWorkspaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLastWorkspaceResponseBody) SetErrorMessage(v string) *GetLastWorkspaceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLastWorkspaceResponseBody) SetSuccess(v bool) *GetLastWorkspaceResponseBody {
	s.Success = &v
	return s
}

type GetLastWorkspaceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLastWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLastWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLastWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetLastWorkspaceResponse) SetHeaders(v map[string]*string) *GetLastWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetLastWorkspaceResponse) SetBody(v *GetLastWorkspaceResponseBody) *GetLastWorkspaceResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceBuildNumberStatusRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	BuildNum   *int64  `json:"BuildNum,omitempty" xml:"BuildNum,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusRequest) GoString() string {
	return s.String()
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

func (s *GetPipelineInstanceBuildNumberStatusRequest) SetBuildNum(v int64) *GetPipelineInstanceBuildNumberStatusRequest {
	s.BuildNum = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBody struct {
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *GetPipelineInstanceBuildNumberStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode    *string                                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetObject(v *GetPipelineInstanceBuildNumberStatusResponseBodyObject) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceBuildNumberStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObject struct {
	Status *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Groups []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObject) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObject) SetGroups(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) *GetPipelineInstanceBuildNumberStatusResponseBodyObject {
	s.Groups = v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups struct {
	Status *string                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	Stages []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Name   *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetStages(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups) SetName(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages struct {
	Status     *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Sign       *string                                                                         `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Components []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	JobId  *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v int64) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipelineInstanceBuildNumberStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

type GetPipelineInstanceBuildNumberStatusResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineInstanceBuildNumberStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineInstanceBuildNumberStatusResponse) SetBody(v *GetPipelineInstanceBuildNumberStatusResponseBody) *GetPipelineInstanceBuildNumberStatusResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceGroupStatusRequest struct {
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
}

func (s GetPipelineInstanceGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusRequest) GoString() string {
	return s.String()
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

func (s *GetPipelineInstanceGroupStatusRequest) SetFlowInstanceId(v int64) *GetPipelineInstanceGroupStatusRequest {
	s.FlowInstanceId = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBody struct {
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *GetPipelineInstanceGroupStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode    *string                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetObject(v *GetPipelineInstanceGroupStatusResponseBodyObject) *GetPipelineInstanceGroupStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceGroupStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceGroupStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObject struct {
	Status *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Groups []*GetPipelineInstanceGroupStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObject) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObject) SetGroups(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroups) *GetPipelineInstanceGroupStatusResponseBodyObject {
	s.Groups = v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroups struct {
	Status *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Stages []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Name   *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetStages(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroups) SetName(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages struct {
	Status     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Sign       *string                                                                   `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Components []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

type GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipelineInstanceGroupStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

type GetPipelineInstanceGroupStatusResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineInstanceGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineInstanceGroupStatusResponse) SetBody(v *GetPipelineInstanceGroupStatusResponseBody) *GetPipelineInstanceGroupStatusResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceInfoRequest struct {
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoRequest) SetOrgId(v string) *GetPipelineInstanceInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetPipelineId(v int64) *GetPipelineInstanceInfoRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetFlowInstanceId(v string) *GetPipelineInstanceInfoRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *GetPipelineInstanceInfoRequest) SetUserPk(v string) *GetPipelineInstanceInfoRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceInfoResponseBody struct {
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *GetPipelineInstanceInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode    *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoResponseBody) SetRequestId(v string) *GetPipelineInstanceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetObject(v *GetPipelineInstanceInfoResponseBodyObject) *GetPipelineInstanceInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetErrorCode(v string) *GetPipelineInstanceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetErrorMessage(v string) *GetPipelineInstanceInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBody) SetSuccess(v bool) *GetPipelineInstanceInfoResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceInfoResponseBodyObject struct {
	EndTime             *int64    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status              *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime           *int64    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PackageDownloadUrls []*string `json:"PackageDownloadUrls,omitempty" xml:"PackageDownloadUrls,omitempty" type:"Repeated"`
	EmployeeId          *string   `json:"EmployeeId,omitempty" xml:"EmployeeId,omitempty"`
	DockerImages        []*string `json:"DockerImages,omitempty" xml:"DockerImages,omitempty" type:"Repeated"`
	Sources             *string   `json:"Sources,omitempty" xml:"Sources,omitempty"`
}

func (s GetPipelineInstanceInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetEndTime(v int64) *GetPipelineInstanceInfoResponseBodyObject {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetStatus(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetStartTime(v int64) *GetPipelineInstanceInfoResponseBodyObject {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetPackageDownloadUrls(v []*string) *GetPipelineInstanceInfoResponseBodyObject {
	s.PackageDownloadUrls = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetEmployeeId(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.EmployeeId = &v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetDockerImages(v []*string) *GetPipelineInstanceInfoResponseBodyObject {
	s.DockerImages = v
	return s
}

func (s *GetPipelineInstanceInfoResponseBodyObject) SetSources(v string) *GetPipelineInstanceInfoResponseBodyObject {
	s.Sources = &v
	return s
}

type GetPipelineInstanceInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineInstanceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineInstanceInfoResponse) SetBody(v *GetPipelineInstanceInfoResponseBody) *GetPipelineInstanceInfoResponse {
	s.Body = v
	return s
}

type GetPipelineInstanceStatusRequest struct {
	OrgId          *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId     *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	FlowInstanceId *int64  `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	UserPk         *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetPipelineInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceStatusRequest) SetOrgId(v string) *GetPipelineInstanceStatusRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetPipelineId(v int64) *GetPipelineInstanceStatusRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetFlowInstanceId(v int64) *GetPipelineInstanceStatusRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *GetPipelineInstanceStatusRequest) SetUserPk(v string) *GetPipelineInstanceStatusRequest {
	s.UserPk = &v
	return s
}

type GetPipelineInstanceStatusResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstanceStatusResponseBody) SetRequestId(v string) *GetPipelineInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetObject(v string) *GetPipelineInstanceStatusResponseBody {
	s.Object = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetErrorCode(v string) *GetPipelineInstanceStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetErrorMessage(v string) *GetPipelineInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineInstanceStatusResponseBody) SetSuccess(v bool) *GetPipelineInstanceStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipelineInstanceStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineInstanceStatusResponse) SetBody(v *GetPipelineInstanceStatusResponseBody) *GetPipelineInstanceStatusResponse {
	s.Body = v
	return s
}

type GetPipelineInstHistoryRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageStart  *int64  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetPipelineInstHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryRequest) SetOrgId(v string) *GetPipelineInstHistoryRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPipelineId(v int64) *GetPipelineInstHistoryRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetUserPk(v string) *GetPipelineInstHistoryRequest {
	s.UserPk = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetStartTime(v string) *GetPipelineInstHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetEndTime(v string) *GetPipelineInstHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPageStart(v int64) *GetPipelineInstHistoryRequest {
	s.PageStart = &v
	return s
}

func (s *GetPipelineInstHistoryRequest) SetPageSize(v int64) *GetPipelineInstHistoryRequest {
	s.PageSize = &v
	return s
}

type GetPipelineInstHistoryResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data         *GetPipelineInstHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode    *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineInstHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBody) SetRequestId(v string) *GetPipelineInstHistoryResponseBody {
	s.RequestId = &v
	return s
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
	Status           *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StatusName       *string                                                     `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	Deletion         *string                                                     `json:"Deletion,omitempty" xml:"Deletion,omitempty"`
	PipelineConfigId *int32                                                      `json:"PipelineConfigId,omitempty" xml:"PipelineConfigId,omitempty"`
	TriggerMode      *int32                                                      `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
	Creator          *string                                                     `json:"Creator,omitempty" xml:"Creator,omitempty"`
	InstNumber       *int32                                                      `json:"InstNumber,omitempty" xml:"InstNumber,omitempty"`
	Modifier         *string                                                     `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	FlowInstance     *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance `json:"FlowInstance,omitempty" xml:"FlowInstance,omitempty" type:"Struct"`
	Packages         *string                                                     `json:"Packages,omitempty" xml:"Packages,omitempty"`
	FlowInstId       *int32                                                      `json:"FlowInstId,omitempty" xml:"FlowInstId,omitempty"`
	PipelineId       *int32                                                      `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	Id               *int32                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime       *int64                                                      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetStatusName(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.StatusName = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetDeletion(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Deletion = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPipelineConfigId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.PipelineConfigId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetTriggerMode(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.TriggerMode = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Creator = &v
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

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetFlowInstance(v *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.FlowInstance = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPackages(v string) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Packages = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetFlowInstId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.FlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetPipelineId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataList) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataList {
	s.ModifyTime = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstance struct {
	Status           *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Stages           map[string]interface{}                                              `json:"Stages,omitempty" xml:"Stages,omitempty"`
	Result           *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult   `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	CreateTime       *int64                                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	StatusName       *string                                                             `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	RunningStatus    *string                                                             `json:"RunningStatus,omitempty" xml:"RunningStatus,omitempty"`
	StageTopo        *string                                                             `json:"StageTopo,omitempty" xml:"StageTopo,omitempty"`
	Creator          *string                                                             `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Modifier         *string                                                             `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	Groups           []*GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	AutoDrivenStatus *bool                                                               `json:"AutoDrivenStatus,omitempty" xml:"AutoDrivenStatus,omitempty"`
	ResultStatus     *string                                                             `json:"ResultStatus,omitempty" xml:"ResultStatus,omitempty"`
	Id               *int32                                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	SystemCode       *string                                                             `json:"SystemCode,omitempty" xml:"SystemCode,omitempty"`
	ModifyTime       *int64                                                              `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	SystemId         *string                                                             `json:"SystemId,omitempty" xml:"SystemId,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Status = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStages(v map[string]interface{}) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Stages = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetResult(v *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Result = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetStatusName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.StatusName = &v
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

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Creator = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetModifier(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Modifier = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetGroups(v []*GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Groups = v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetAutoDrivenStatus(v bool) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.AutoDrivenStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetResultStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.ResultStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetSystemCode(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.SystemCode = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.ModifyTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance) SetSystemId(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstance {
	s.SystemId = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult struct {
	EnginePipelineNumber *int32  `json:"EnginePipelineNumber,omitempty" xml:"EnginePipelineNumber,omitempty"`
	MixFlowInstId        *string `json:"MixFlowInstId,omitempty" xml:"MixFlowInstId,omitempty"`
	EnginePipelineName   *string `json:"EnginePipelineName,omitempty" xml:"EnginePipelineName,omitempty"`
	EnginePipelineId     *int32  `json:"EnginePipelineId,omitempty" xml:"EnginePipelineId,omitempty"`
	EnginePipelineInstId *int32  `json:"EnginePipelineInstId,omitempty" xml:"EnginePipelineInstId,omitempty"`
	TimeStamp            *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	TriggerMode          *string `json:"TriggerMode,omitempty" xml:"TriggerMode,omitempty"`
	Sources              *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	Caches               *string `json:"Caches,omitempty" xml:"Caches,omitempty"`
	DateTime             *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineNumber(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineNumber = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetMixFlowInstId(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.MixFlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetEnginePipelineName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.EnginePipelineName = &v
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

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetTimeStamp(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.TimeStamp = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetTriggerMode(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.TriggerMode = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetSources(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.Sources = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetCaches(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.Caches = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult) SetDateTime(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceResult {
	s.DateTime = &v
	return s
}

type GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeleteStatus *string `json:"DeleteStatus,omitempty" xml:"DeleteStatus,omitempty"`
	IdExtent     *int32  `json:"IdExtent,omitempty" xml:"IdExtent,omitempty"`
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Modifier     *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	ResultStatus *string `json:"ResultStatus,omitempty" xml:"ResultStatus,omitempty"`
	FlowInstId   *int32  `json:"FlowInstId,omitempty" xml:"FlowInstId,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime   *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) GoString() string {
	return s.String()
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Status = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetCreateTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.CreateTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetDeleteStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.DeleteStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetIdExtent(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.IdExtent = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetCreator(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Creator = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetEndTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.EndTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetStartTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.StartTime = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetModifier(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Modifier = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetResultStatus(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.ResultStatus = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetFlowInstId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.FlowInstId = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetName(v string) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Name = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetId(v int32) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.Id = &v
	return s
}

func (s *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups) SetModifyTime(v int64) *GetPipelineInstHistoryResponseBodyDataDataListFlowInstanceGroups {
	s.ModifyTime = &v
	return s
}

type GetPipelineInstHistoryResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineInstHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineInstHistoryResponse) SetBody(v *GetPipelineInstHistoryResponseBody) *GetPipelineInstHistoryResponse {
	s.Body = v
	return s
}

type GetPipelineLogRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	JobId      *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetPipelineLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogRequest) GoString() string {
	return s.String()
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

func (s *GetPipelineLogRequest) SetJobId(v int64) *GetPipelineLogRequest {
	s.JobId = &v
	return s
}

type GetPipelineLogResponseBody struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       []*GetPipelineLogResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode    *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBody) SetRequestId(v string) *GetPipelineLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetObject(v []*GetPipelineLogResponseBodyObject) *GetPipelineLogResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineLogResponseBody) SetErrorCode(v string) *GetPipelineLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetErrorMessage(v string) *GetPipelineLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineLogResponseBody) SetSuccess(v bool) *GetPipelineLogResponseBody {
	s.Success = &v
	return s
}

type GetPipelineLogResponseBodyObject struct {
	StartTime         *string                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	JobId             *int64                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ActionName        *string                                              `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	BuildProcessNodes []*GetPipelineLogResponseBodyObjectBuildProcessNodes `json:"BuildProcessNodes,omitempty" xml:"BuildProcessNodes,omitempty" type:"Repeated"`
}

func (s GetPipelineLogResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBodyObject) SetStartTime(v string) *GetPipelineLogResponseBodyObject {
	s.StartTime = &v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetJobId(v int64) *GetPipelineLogResponseBodyObject {
	s.JobId = &v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetActionName(v string) *GetPipelineLogResponseBodyObject {
	s.ActionName = &v
	return s
}

func (s *GetPipelineLogResponseBodyObject) SetBuildProcessNodes(v []*GetPipelineLogResponseBodyObjectBuildProcessNodes) *GetPipelineLogResponseBodyObject {
	s.BuildProcessNodes = v
	return s
}

type GetPipelineLogResponseBodyObjectBuildProcessNodes struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	NodeName  *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeIndex *int32  `json:"NodeIndex,omitempty" xml:"NodeIndex,omitempty"`
}

func (s GetPipelineLogResponseBodyObjectBuildProcessNodes) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineLogResponseBodyObjectBuildProcessNodes) GoString() string {
	return s.String()
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetStatus(v string) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.Status = &v
	return s
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetNodeName(v string) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.NodeName = &v
	return s
}

func (s *GetPipelineLogResponseBodyObjectBuildProcessNodes) SetNodeIndex(v int32) *GetPipelineLogResponseBodyObjectBuildProcessNodes {
	s.NodeIndex = &v
	return s
}

type GetPipelineLogResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipelineLogResponse) SetBody(v *GetPipelineLogResponseBody) *GetPipelineLogResponse {
	s.Body = v
	return s
}

type GetPipelineStepLogRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	JobId      *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	StepIndex  *string `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	Offset     *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Limit      *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s GetPipelineStepLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogRequest) SetOrgId(v string) *GetPipelineStepLogRequest {
	s.OrgId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetPipelineId(v int64) *GetPipelineStepLogRequest {
	s.PipelineId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetUserPk(v string) *GetPipelineStepLogRequest {
	s.UserPk = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetJobId(v int64) *GetPipelineStepLogRequest {
	s.JobId = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetStepIndex(v string) *GetPipelineStepLogRequest {
	s.StepIndex = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetOffset(v int64) *GetPipelineStepLogRequest {
	s.Offset = &v
	return s
}

func (s *GetPipelineStepLogRequest) SetLimit(v int64) *GetPipelineStepLogRequest {
	s.Limit = &v
	return s
}

type GetPipelineStepLogResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *GetPipelineStepLogResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode    *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipelineStepLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogResponseBody) SetRequestId(v string) *GetPipelineStepLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetObject(v *GetPipelineStepLogResponseBodyObject) *GetPipelineStepLogResponseBody {
	s.Object = v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetErrorCode(v string) *GetPipelineStepLogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetErrorMessage(v string) *GetPipelineStepLogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipelineStepLogResponseBody) SetSuccess(v bool) *GetPipelineStepLogResponseBody {
	s.Success = &v
	return s
}

type GetPipelineStepLogResponseBodyObject struct {
	Logs *string `json:"Logs,omitempty" xml:"Logs,omitempty"`
	Last *int32  `json:"Last,omitempty" xml:"Last,omitempty"`
	More *bool   `json:"More,omitempty" xml:"More,omitempty"`
}

func (s GetPipelineStepLogResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipelineStepLogResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipelineStepLogResponseBodyObject) SetLogs(v string) *GetPipelineStepLogResponseBodyObject {
	s.Logs = &v
	return s
}

func (s *GetPipelineStepLogResponseBodyObject) SetLast(v int32) *GetPipelineStepLogResponseBodyObject {
	s.Last = &v
	return s
}

func (s *GetPipelineStepLogResponseBodyObject) SetMore(v bool) *GetPipelineStepLogResponseBodyObject {
	s.More = &v
	return s
}

type GetPipelineStepLogResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipelineStepLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *GetPipleineLatestInstanceStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode    *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetRequestId(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetObject(v *GetPipleineLatestInstanceStatusResponseBodyObject) *GetPipleineLatestInstanceStatusResponseBody {
	s.Object = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetErrorCode(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetErrorMessage(v string) *GetPipleineLatestInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBody) SetSuccess(v bool) *GetPipleineLatestInstanceStatusResponseBody {
	s.Success = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObject struct {
	Status *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Groups []*GetPipleineLatestInstanceStatusResponseBodyObjectGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObject) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObject) SetGroups(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroups) *GetPipleineLatestInstanceStatusResponseBodyObject {
	s.Groups = v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroups struct {
	Status *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Stages []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	Name   *string                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroups) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroups) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Status = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetStages(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Stages = v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroups) SetName(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroups {
	s.Name = &v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages struct {
	Status     *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Sign       *string                                                                    `json:"Sign,omitempty" xml:"Sign,omitempty"`
	Components []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Status = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetSign(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Sign = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages) SetComponents(v []*GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStages {
	s.Components = v
	return s
}

type GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	JobId  *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) String() string {
	return tea.Prettify(s)
}

func (s GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) GoString() string {
	return s.String()
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetStatus(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.Status = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetJobId(v int64) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.JobId = &v
	return s
}

func (s *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents) SetName(v string) *GetPipleineLatestInstanceStatusResponseBodyObjectGroupsStagesComponents {
	s.Name = &v
	return s
}

type GetPipleineLatestInstanceStatusResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPipleineLatestInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPipleineLatestInstanceStatusResponse) SetBody(v *GetPipleineLatestInstanceStatusResponseBody) *GetPipleineLatestInstanceStatusResponse {
	s.Body = v
	return s
}

type GetProjectOptionRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Query     *string `json:"Query,omitempty" xml:"Query,omitempty"`
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

func (s *GetProjectOptionRequest) SetType(v string) *GetProjectOptionRequest {
	s.Type = &v
	return s
}

func (s *GetProjectOptionRequest) SetQuery(v string) *GetProjectOptionRequest {
	s.Query = &v
	return s
}

type GetProjectOptionResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetProjectOptionResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectOptionResponseBody) SetRequestId(v string) *GetProjectOptionResponseBody {
	s.RequestId = &v
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

func (s *GetProjectOptionResponseBody) SetErrorCode(v string) *GetProjectOptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectOptionResponseBody) SetSuccessful(v bool) *GetProjectOptionResponseBody {
	s.Successful = &v
	return s
}

type GetProjectOptionResponseBodyObject struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProjectOptionResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectOptionResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectOptionResponseBodyObject) SetValue(v string) *GetProjectOptionResponseBodyObject {
	s.Value = &v
	return s
}

func (s *GetProjectOptionResponseBodyObject) SetName(v string) *GetProjectOptionResponseBodyObject {
	s.Name = &v
	return s
}

type GetProjectOptionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectOptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId      *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ErrorMsg       *string                                    `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object         []*GetTaskDetailActivityResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode      *string                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful     *bool                                      `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskDetailActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailActivityResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskDetailActivityResponseBody) SetRequestId(v string) *GetTaskDetailActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetHttpStatusCode(v int32) *GetTaskDetailActivityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetErrorMsg(v string) *GetTaskDetailActivityResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetObject(v []*GetTaskDetailActivityResponseBodyObject) *GetTaskDetailActivityResponseBody {
	s.Object = v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetErrorCode(v string) *GetTaskDetailActivityResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskDetailActivityResponseBody) SetSuccessful(v bool) *GetTaskDetailActivityResponseBody {
	s.Successful = &v
	return s
}

type GetTaskDetailActivityResponseBodyObject struct {
	Action  *string                `json:"Action,omitempty" xml:"Action,omitempty"`
	Updated *string                `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Title   *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	Created *string                `json:"Created,omitempty" xml:"Created,omitempty"`
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
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

func (s *GetTaskDetailActivityResponseBodyObject) SetUpdated(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetTitle(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Title = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetCreated(v string) *GetTaskDetailActivityResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetTaskDetailActivityResponseBodyObject) SetContent(v map[string]interface{}) *GetTaskDetailActivityResponseBodyObject {
	s.Content = v
	return s
}

type GetTaskDetailActivityResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskDetailActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetTaskDetailBaseResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskDetailBaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBody) SetRequestId(v string) *GetTaskDetailBaseResponseBody {
	s.RequestId = &v
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

func (s *GetTaskDetailBaseResponseBody) SetErrorCode(v string) *GetTaskDetailBaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskDetailBaseResponseBody) SetSuccessful(v bool) *GetTaskDetailBaseResponseBody {
	s.Successful = &v
	return s
}

type GetTaskDetailBaseResponseBodyObject struct {
	Organization          *string                                                 `json:"Organization,omitempty" xml:"Organization,omitempty"`
	ScenariofieldconfigId *string                                                 `json:"ScenariofieldconfigId,omitempty" xml:"ScenariofieldconfigId,omitempty"`
	ProjectId             *string                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsTopInProject        *bool                                                   `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Tasklist              *GetTaskDetailBaseResponseBodyObjectTasklist            `json:"Tasklist,omitempty" xml:"Tasklist,omitempty" type:"Struct"`
	Badges                *GetTaskDetailBaseResponseBodyObjectBadges              `json:"Badges,omitempty" xml:"Badges,omitempty" type:"Struct"`
	AncestorIds           []*string                                               `json:"AncestorIds,omitempty" xml:"AncestorIds,omitempty" type:"Repeated"`
	ShareStatus           *int32                                                  `json:"ShareStatus,omitempty" xml:"ShareStatus,omitempty"`
	Reminder              *GetTaskDetailBaseResponseBodyObjectReminder            `json:"Reminder,omitempty" xml:"Reminder,omitempty" type:"Struct"`
	Ancestors             []*string                                               `json:"Ancestors,omitempty" xml:"Ancestors,omitempty" type:"Repeated"`
	TaskflowstatusId      *string                                                 `json:"TaskflowstatusId,omitempty" xml:"TaskflowstatusId,omitempty"`
	Updated               *string                                                 `json:"Updated,omitempty" xml:"Updated,omitempty"`
	Note                  *string                                                 `json:"Note,omitempty" xml:"Note,omitempty"`
	IsArchived            *bool                                                   `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	Content               *string                                                 `json:"Content,omitempty" xml:"Content,omitempty"`
	Rating                *int32                                                  `json:"Rating,omitempty" xml:"Rating,omitempty"`
	Progress              *int32                                                  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Stage                 *GetTaskDetailBaseResponseBodyObjectStage               `json:"Stage,omitempty" xml:"Stage,omitempty" type:"Struct"`
	Labels                []*string                                               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	StartDate             *string                                                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Sprint                *string                                                 `json:"Sprint,omitempty" xml:"Sprint,omitempty"`
	CreatorId             *string                                                 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	SourceId              *string                                                 `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	OrganizationId        *string                                                 `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SourceDate            *string                                                 `json:"SourceDate,omitempty" xml:"SourceDate,omitempty"`
	IsFavorite            *bool                                                   `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	ExecutorId            *string                                                 `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	Scenariofieldconfig   *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig `json:"Scenariofieldconfig,omitempty" xml:"Scenariofieldconfig,omitempty" type:"Struct"`
	WorkTime              *GetTaskDetailBaseResponseBodyObjectWorkTime            `json:"WorkTime,omitempty" xml:"WorkTime,omitempty" type:"Struct"`
	TagIds                []*string                                               `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	Priority              *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Creator               *GetTaskDetailBaseResponseBodyObjectCreator             `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Executor              *GetTaskDetailBaseResponseBodyObjectExecutor            `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	Accomplished          *string                                                 `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	InvolveMembers        []*string                                               `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	UniqueId              *int32                                                  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	CommentsCount         *int32                                                  `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	Recurrence            *string                                                 `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	ObjectType            *string                                                 `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SubtaskCount          *GetTaskDetailBaseResponseBodyObjectSubtaskCount        `json:"SubtaskCount,omitempty" xml:"SubtaskCount,omitempty" type:"Struct"`
	UntilDate             *string                                                 `json:"UntilDate,omitempty" xml:"UntilDate,omitempty"`
	StoryPoint            *string                                                 `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	ObjectlinksCount      *int32                                                  `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	Source                *string                                                 `json:"Source,omitempty" xml:"Source,omitempty"`
	LikesCount            *int32                                                  `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	StageId               *string                                                 `json:"StageId,omitempty" xml:"StageId,omitempty"`
	Divisions             []*string                                               `json:"Divisions,omitempty" xml:"Divisions,omitempty" type:"Repeated"`
	Visible               *string                                                 `json:"Visible,omitempty" xml:"Visible,omitempty"`
	IsDone                *bool                                                   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	Involvers             []*GetTaskDetailBaseResponseBodyObjectInvolvers         `json:"Involvers,omitempty" xml:"Involvers,omitempty" type:"Repeated"`
	Parent                *string                                                 `json:"Parent,omitempty" xml:"Parent,omitempty"`
	SprintId              *string                                                 `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	DueDate               *string                                                 `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	AttachmentsCount      *int32                                                  `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	Subtasks              []*GetTaskDetailBaseResponseBodyObjectSubtasks          `json:"Subtasks,omitempty" xml:"Subtasks,omitempty" type:"Repeated"`
	Customfields          []*GetTaskDetailBaseResponseBodyObjectCustomfields      `json:"Customfields,omitempty" xml:"Customfields,omitempty" type:"Repeated"`
	Created               *string                                                 `json:"Created,omitempty" xml:"Created,omitempty"`
	TaskId                *string                                                 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Taskflowstatus        *GetTaskDetailBaseResponseBodyObjectTaskflowstatus      `json:"Taskflowstatus,omitempty" xml:"Taskflowstatus,omitempty" type:"Struct"`
	Id                    *string                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObject) SetOrganization(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Organization = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetScenariofieldconfigId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ScenariofieldconfigId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetProjectId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsTopInProject(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsTopInProject = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTasklist(v *GetTaskDetailBaseResponseBodyObjectTasklist) *GetTaskDetailBaseResponseBodyObject {
	s.Tasklist = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetBadges(v *GetTaskDetailBaseResponseBodyObjectBadges) *GetTaskDetailBaseResponseBodyObject {
	s.Badges = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAncestorIds(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.AncestorIds = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetShareStatus(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.ShareStatus = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetReminder(v *GetTaskDetailBaseResponseBodyObjectReminder) *GetTaskDetailBaseResponseBodyObject {
	s.Reminder = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAncestors(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Ancestors = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTaskflowstatusId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUpdated(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetNote(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Note = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsArchived(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsArchived = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetContent(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Content = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetRating(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Rating = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetProgress(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Progress = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStage(v *GetTaskDetailBaseResponseBodyObjectStage) *GetTaskDetailBaseResponseBodyObject {
	s.Stage = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetLabels(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Labels = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStartDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSprint(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Sprint = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreatorId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSourceId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SourceId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetOrganizationId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSourceDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SourceDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsFavorite(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsFavorite = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetExecutorId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetScenariofieldconfig(v *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) *GetTaskDetailBaseResponseBodyObject {
	s.Scenariofieldconfig = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetWorkTime(v *GetTaskDetailBaseResponseBodyObjectWorkTime) *GetTaskDetailBaseResponseBodyObject {
	s.WorkTime = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetTagIds(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.TagIds = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetPriority(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreator(v *GetTaskDetailBaseResponseBodyObjectCreator) *GetTaskDetailBaseResponseBodyObject {
	s.Creator = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetExecutor(v *GetTaskDetailBaseResponseBodyObjectExecutor) *GetTaskDetailBaseResponseBodyObject {
	s.Executor = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAccomplished(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetInvolveMembers(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.InvolveMembers = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUniqueId(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.UniqueId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCommentsCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetRecurrence(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Recurrence = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetObjectType(v string) *GetTaskDetailBaseResponseBodyObject {
	s.ObjectType = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSubtaskCount(v *GetTaskDetailBaseResponseBodyObjectSubtaskCount) *GetTaskDetailBaseResponseBodyObject {
	s.SubtaskCount = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetUntilDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.UntilDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStoryPoint(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetObjectlinksCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSource(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Source = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetLikesCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.LikesCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetStageId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.StageId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetDivisions(v []*string) *GetTaskDetailBaseResponseBodyObject {
	s.Divisions = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetVisible(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetIsDone(v bool) *GetTaskDetailBaseResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetInvolvers(v []*GetTaskDetailBaseResponseBodyObjectInvolvers) *GetTaskDetailBaseResponseBodyObject {
	s.Involvers = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetParent(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Parent = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSprintId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetDueDate(v string) *GetTaskDetailBaseResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetAttachmentsCount(v int32) *GetTaskDetailBaseResponseBodyObject {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetSubtasks(v []*GetTaskDetailBaseResponseBodyObjectSubtasks) *GetTaskDetailBaseResponseBodyObject {
	s.Subtasks = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCustomfields(v []*GetTaskDetailBaseResponseBodyObjectCustomfields) *GetTaskDetailBaseResponseBodyObject {
	s.Customfields = v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObject) SetCreated(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Created = &v
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

func (s *GetTaskDetailBaseResponseBodyObject) SetId(v string) *GetTaskDetailBaseResponseBodyObject {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectTasklist struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectTasklist) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectTasklist) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectTasklist) SetTitle(v string) *GetTaskDetailBaseResponseBodyObjectTasklist {
	s.Title = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectTasklist) SetId(v string) *GetTaskDetailBaseResponseBodyObjectTasklist {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectBadges struct {
	LikesCount       *int32 `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	ObjectlinksCount *int32 `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	AttachmentsCount *int32 `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	CommentsCount    *int32 `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectBadges) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectBadges) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetLikesCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.LikesCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetObjectlinksCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetAttachmentsCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectBadges) SetCommentsCount(v int32) *GetTaskDetailBaseResponseBodyObjectBadges {
	s.CommentsCount = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectReminder struct {
	Type        *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Members     []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Date        *string   `json:"Date,omitempty" xml:"Date,omitempty"`
	MemberRoles []*string `json:"MemberRoles,omitempty" xml:"MemberRoles,omitempty" type:"Repeated"`
	Method      *string   `json:"Method,omitempty" xml:"Method,omitempty"`
	CreatorId   *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Rules       []*string `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetTaskDetailBaseResponseBodyObjectReminder) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectReminder) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetType(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Type = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetMembers(v []*string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Members = v
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

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetMethod(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Method = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetCreatorId(v string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.CreatorId = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectReminder) SetRules(v []*string) *GetTaskDetailBaseResponseBodyObjectReminder {
	s.Rules = v
	return s
}

type GetTaskDetailBaseResponseBodyObjectStage struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectStage) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectStage) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectStage) SetName(v string) *GetTaskDetailBaseResponseBodyObjectStage {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectStage) SetId(v string) *GetTaskDetailBaseResponseBodyObjectStage {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectScenariofieldconfig struct {
	Icon                  *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	ProTemplateConfigType *string `json:"ProTemplateConfigType,omitempty" xml:"ProTemplateConfigType,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetProTemplateConfigType(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.ProTemplateConfigType = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetName(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig) SetId(v string) *GetTaskDetailBaseResponseBodyObjectScenariofieldconfig {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectWorkTime struct {
	UsedTime  *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	TotalTime *int32  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Unit      *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectWorkTime) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectWorkTime) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetUsedTime(v int32) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.UsedTime = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetTotalTime(v int32) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.TotalTime = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectWorkTime) SetUnit(v string) *GetTaskDetailBaseResponseBodyObjectWorkTime {
	s.Unit = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectCreator struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectCreator) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCreator) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectCreator) SetName(v string) *GetTaskDetailBaseResponseBodyObjectCreator {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCreator) SetId(v string) *GetTaskDetailBaseResponseBodyObjectCreator {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectExecutor struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetTaskDetailBaseResponseBodyObjectExecutor) SetName(v string) *GetTaskDetailBaseResponseBodyObjectExecutor {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectExecutor) SetId(v string) *GetTaskDetailBaseResponseBodyObjectExecutor {
	s.Id = &v
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

type GetTaskDetailBaseResponseBodyObjectInvolvers struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectInvolvers) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectInvolvers) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectInvolvers) SetName(v string) *GetTaskDetailBaseResponseBodyObjectInvolvers {
	s.Name = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectInvolvers) SetId(v string) *GetTaskDetailBaseResponseBodyObjectInvolvers {
	s.Id = &v
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

type GetTaskDetailBaseResponseBodyObjectCustomfields struct {
	Type          *string                                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         []*GetTaskDetailBaseResponseBodyObjectCustomfieldsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Values        []*string                                               `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	CustomfieldId *string                                                 `json:"CustomfieldId,omitempty" xml:"CustomfieldId,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfields) GoString() string {
	return s.String()
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

func (s *GetTaskDetailBaseResponseBodyObjectCustomfields) SetCustomfieldId(v string) *GetTaskDetailBaseResponseBodyObjectCustomfields {
	s.CustomfieldId = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectCustomfieldsValue struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) SetTitle(v string) *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue {
	s.Title = &v
	return s
}

func (s *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue) SetId(v string) *GetTaskDetailBaseResponseBodyObjectCustomfieldsValue {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponseBodyObjectTaskflowstatus struct {
	TaskflowId *string `json:"TaskflowId,omitempty" xml:"TaskflowId,omitempty"`
	Kind       *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskDetailBaseResponseBodyObjectTaskflowstatus) String() string {
	return tea.Prettify(s)
}

func (s GetTaskDetailBaseResponseBodyObjectTaskflowstatus) GoString() string {
	return s.String()
}

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetTaskflowId(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.TaskflowId = &v
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

func (s *GetTaskDetailBaseResponseBodyObjectTaskflowstatus) SetId(v string) *GetTaskDetailBaseResponseBodyObjectTaskflowstatus {
	s.Id = &v
	return s
}

type GetTaskDetailBaseResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskDetailBaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskDetailBaseResponse) SetBody(v *GetTaskDetailBaseResponseBody) *GetTaskDetailBaseResponse {
	s.Body = v
	return s
}

type GetTaskListFilterRequest struct {
	OrgId                 *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId             *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ScenarioFieldConfigId *string `json:"ScenarioFieldConfigId,omitempty" xml:"ScenarioFieldConfigId,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderCondition        *string `json:"OrderCondition,omitempty" xml:"OrderCondition,omitempty"`
	Order                 *string `json:"Order,omitempty" xml:"Order,omitempty"`
	ExecutorId            *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	TagId                 *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	DueDateStart          *string `json:"DueDateStart,omitempty" xml:"DueDateStart,omitempty"`
	DueDateEnd            *string `json:"DueDateEnd,omitempty" xml:"DueDateEnd,omitempty"`
	CreatorId             *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	InvolveMembers        *string `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty"`
	IsDone                *bool   `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageToken             *string `json:"PageToken,omitempty" xml:"PageToken,omitempty"`
	ObjectType            *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	TaskFlowStatusId      *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	SprintId              *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	Extra                 *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
}

func (s GetTaskListFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterRequest) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterRequest) SetOrgId(v string) *GetTaskListFilterRequest {
	s.OrgId = &v
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

func (s *GetTaskListFilterRequest) SetName(v string) *GetTaskListFilterRequest {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterRequest) SetOrderCondition(v string) *GetTaskListFilterRequest {
	s.OrderCondition = &v
	return s
}

func (s *GetTaskListFilterRequest) SetOrder(v string) *GetTaskListFilterRequest {
	s.Order = &v
	return s
}

func (s *GetTaskListFilterRequest) SetExecutorId(v string) *GetTaskListFilterRequest {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetTagId(v string) *GetTaskListFilterRequest {
	s.TagId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetDueDateStart(v string) *GetTaskListFilterRequest {
	s.DueDateStart = &v
	return s
}

func (s *GetTaskListFilterRequest) SetDueDateEnd(v string) *GetTaskListFilterRequest {
	s.DueDateEnd = &v
	return s
}

func (s *GetTaskListFilterRequest) SetCreatorId(v string) *GetTaskListFilterRequest {
	s.CreatorId = &v
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

func (s *GetTaskListFilterRequest) SetPriority(v string) *GetTaskListFilterRequest {
	s.Priority = &v
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

func (s *GetTaskListFilterRequest) SetObjectType(v string) *GetTaskListFilterRequest {
	s.ObjectType = &v
	return s
}

func (s *GetTaskListFilterRequest) SetTaskFlowStatusId(v string) *GetTaskListFilterRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetSprintId(v string) *GetTaskListFilterRequest {
	s.SprintId = &v
	return s
}

func (s *GetTaskListFilterRequest) SetExtra(v string) *GetTaskListFilterRequest {
	s.Extra = &v
	return s
}

type GetTaskListFilterResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetTaskListFilterResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *string                              `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetTaskListFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBody) SetRequestId(v string) *GetTaskListFilterResponseBody {
	s.RequestId = &v
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

func (s *GetTaskListFilterResponseBody) SetErrorCode(v string) *GetTaskListFilterResponseBody {
	s.ErrorCode = &v
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
	ProjectId             *string                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsTopInProject        *bool                                                    `json:"IsTopInProject,omitempty" xml:"IsTopInProject,omitempty"`
	Badges                *GetTaskListFilterResponseBodyObjectResultBadges         `json:"Badges,omitempty" xml:"Badges,omitempty" type:"Struct"`
	AncestorIds           []*string                                                `json:"AncestorIds,omitempty" xml:"AncestorIds,omitempty" type:"Repeated"`
	ShareStatus           *int32                                                   `json:"ShareStatus,omitempty" xml:"ShareStatus,omitempty"`
	Reminder              *GetTaskListFilterResponseBodyObjectResultReminder       `json:"Reminder,omitempty" xml:"Reminder,omitempty" type:"Struct"`
	Note                  *string                                                  `json:"Note,omitempty" xml:"Note,omitempty"`
	Updated               *string                                                  `json:"Updated,omitempty" xml:"Updated,omitempty"`
	IsArchived            *bool                                                    `json:"IsArchived,omitempty" xml:"IsArchived,omitempty"`
	Content               *string                                                  `json:"Content,omitempty" xml:"Content,omitempty"`
	Rating                *int32                                                   `json:"Rating,omitempty" xml:"Rating,omitempty"`
	TaskFlowStatusId      *string                                                  `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	Progress              *int32                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Stage                 *GetTaskListFilterResponseBodyObjectResultStage          `json:"Stage,omitempty" xml:"Stage,omitempty" type:"Struct"`
	Labels                []*string                                                `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Pos                   *int32                                                   `json:"Pos,omitempty" xml:"Pos,omitempty"`
	StartDate             *string                                                  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	Sprint                *string                                                  `json:"Sprint,omitempty" xml:"Sprint,omitempty"`
	CreatorId             *string                                                  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	SourceId              *string                                                  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	OrganizationId        *string                                                  `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	SourceDate            *string                                                  `json:"SourceDate,omitempty" xml:"SourceDate,omitempty"`
	IsFavorite            *bool                                                    `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	ExecutorId            *string                                                  `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	WorkTime              *GetTaskListFilterResponseBodyObjectResultWorkTime       `json:"WorkTime,omitempty" xml:"WorkTime,omitempty" type:"Struct"`
	TagIds                []*string                                                `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	Priority              *int32                                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ScenariofFeldConfigId *string                                                  `json:"ScenariofFeldConfigId,omitempty" xml:"ScenariofFeldConfigId,omitempty"`
	Creator               *GetTaskListFilterResponseBodyObjectResultCreator        `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Executor              *GetTaskListFilterResponseBodyObjectResultExecutor       `json:"Executor,omitempty" xml:"Executor,omitempty" type:"Struct"`
	Accomplished          *string                                                  `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	TaskListId            *string                                                  `json:"TaskListId,omitempty" xml:"TaskListId,omitempty"`
	InvolveMembers        []*string                                                `json:"InvolveMembers,omitempty" xml:"InvolveMembers,omitempty" type:"Repeated"`
	UniqueId              *int32                                                   `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	TaskFlowStatus        *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus `json:"TaskFlowStatus,omitempty" xml:"TaskFlowStatus,omitempty" type:"Struct"`
	CommentsCount         *int32                                                   `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
	Recurrence            *string                                                  `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	ObjectType            *string                                                  `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SubtaskCount          *GetTaskListFilterResponseBodyObjectResultSubtaskCount   `json:"SubtaskCount,omitempty" xml:"SubtaskCount,omitempty" type:"Struct"`
	UntilDate             *string                                                  `json:"UntilDate,omitempty" xml:"UntilDate,omitempty"`
	StoryPoint            *string                                                  `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	ObjectlinksCount      *int32                                                   `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	Source                *string                                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	LikesCount            *int32                                                   `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	StageId               *string                                                  `json:"StageId,omitempty" xml:"StageId,omitempty"`
	Divisions             []*string                                                `json:"Divisions,omitempty" xml:"Divisions,omitempty" type:"Repeated"`
	Visible               *string                                                  `json:"Visible,omitempty" xml:"Visible,omitempty"`
	IsDone                *bool                                                    `json:"IsDone,omitempty" xml:"IsDone,omitempty"`
	Parent                *string                                                  `json:"Parent,omitempty" xml:"Parent,omitempty"`
	SprintId              *string                                                  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	DueDate               *string                                                  `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	AttachmentsCount      *int32                                                   `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	Customfields          []*GetTaskListFilterResponseBodyObjectResultCustomfields `json:"Customfields,omitempty" xml:"Customfields,omitempty" type:"Repeated"`
	Created               *string                                                  `json:"Created,omitempty" xml:"Created,omitempty"`
	TaskUniqueId          *string                                                  `json:"TaskUniqueId,omitempty" xml:"TaskUniqueId,omitempty"`
	TaskId                *string                                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Id                    *string                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResult) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetProjectId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ProjectId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsTopInProject(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsTopInProject = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetBadges(v *GetTaskListFilterResponseBodyObjectResultBadges) *GetTaskListFilterResponseBodyObjectResult {
	s.Badges = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAncestorIds(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.AncestorIds = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetShareStatus(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.ShareStatus = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetReminder(v *GetTaskListFilterResponseBodyObjectResultReminder) *GetTaskListFilterResponseBodyObjectResult {
	s.Reminder = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetNote(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Note = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUpdated(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Updated = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsArchived(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsArchived = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetContent(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Content = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetRating(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Rating = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskFlowStatusId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskFlowStatusId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetProgress(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Progress = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStage(v *GetTaskListFilterResponseBodyObjectResultStage) *GetTaskListFilterResponseBodyObjectResult {
	s.Stage = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetLabels(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.Labels = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetPos(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Pos = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStartDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StartDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSprint(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Sprint = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreatorId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.CreatorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSourceId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SourceId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetOrganizationId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.OrganizationId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSourceDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SourceDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsFavorite(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsFavorite = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetExecutorId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ExecutorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetWorkTime(v *GetTaskListFilterResponseBodyObjectResultWorkTime) *GetTaskListFilterResponseBodyObjectResult {
	s.WorkTime = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTagIds(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.TagIds = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetPriority(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.Priority = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetScenariofFeldConfigId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ScenariofFeldConfigId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreator(v *GetTaskListFilterResponseBodyObjectResultCreator) *GetTaskListFilterResponseBodyObjectResult {
	s.Creator = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetExecutor(v *GetTaskListFilterResponseBodyObjectResultExecutor) *GetTaskListFilterResponseBodyObjectResult {
	s.Executor = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAccomplished(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Accomplished = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskListId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskListId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetInvolveMembers(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.InvolveMembers = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUniqueId(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.UniqueId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskFlowStatus(v *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskFlowStatus = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCommentsCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.CommentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetRecurrence(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Recurrence = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetObjectType(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.ObjectType = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSubtaskCount(v *GetTaskListFilterResponseBodyObjectResultSubtaskCount) *GetTaskListFilterResponseBodyObjectResult {
	s.SubtaskCount = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetUntilDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.UntilDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStoryPoint(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StoryPoint = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetObjectlinksCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSource(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Source = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetLikesCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.LikesCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetStageId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.StageId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetDivisions(v []*string) *GetTaskListFilterResponseBodyObjectResult {
	s.Divisions = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetVisible(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Visible = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetIsDone(v bool) *GetTaskListFilterResponseBodyObjectResult {
	s.IsDone = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetParent(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Parent = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetSprintId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.SprintId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetDueDate(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.DueDate = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetAttachmentsCount(v int32) *GetTaskListFilterResponseBodyObjectResult {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCustomfields(v []*GetTaskListFilterResponseBodyObjectResultCustomfields) *GetTaskListFilterResponseBodyObjectResult {
	s.Customfields = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetCreated(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Created = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskUniqueId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskUniqueId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetTaskId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.TaskId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResult) SetId(v string) *GetTaskListFilterResponseBodyObjectResult {
	s.Id = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultBadges struct {
	LikesCount       *int32 `json:"LikesCount,omitempty" xml:"LikesCount,omitempty"`
	ObjectlinksCount *int32 `json:"ObjectlinksCount,omitempty" xml:"ObjectlinksCount,omitempty"`
	AttachmentsCount *int32 `json:"AttachmentsCount,omitempty" xml:"AttachmentsCount,omitempty"`
	CommentsCount    *int32 `json:"CommentsCount,omitempty" xml:"CommentsCount,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultBadges) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultBadges) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetLikesCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.LikesCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetObjectlinksCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.ObjectlinksCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetAttachmentsCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.AttachmentsCount = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultBadges) SetCommentsCount(v int32) *GetTaskListFilterResponseBodyObjectResultBadges {
	s.CommentsCount = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultReminder struct {
	Type      *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Members   []*string `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Date      *string   `json:"Date,omitempty" xml:"Date,omitempty"`
	CreatorId *string   `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Rules     []*string `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetTaskListFilterResponseBodyObjectResultReminder) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultReminder) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetType(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Type = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetMembers(v []*string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Members = v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetDate(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Date = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetCreatorId(v string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.CreatorId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultReminder) SetRules(v []*string) *GetTaskListFilterResponseBodyObjectResultReminder {
	s.Rules = v
	return s
}

type GetTaskListFilterResponseBodyObjectResultStage struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultStage) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultStage) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultStage) SetName(v string) *GetTaskListFilterResponseBodyObjectResultStage {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultStage) SetId(v string) *GetTaskListFilterResponseBodyObjectResultStage {
	s.Id = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultWorkTime struct {
	UsedTime  *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	TotalTime *int32  `json:"TotalTime,omitempty" xml:"TotalTime,omitempty"`
	Unit      *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultWorkTime) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultWorkTime) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetUsedTime(v int32) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.UsedTime = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetTotalTime(v int32) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.TotalTime = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultWorkTime) SetUnit(v string) *GetTaskListFilterResponseBodyObjectResultWorkTime {
	s.Unit = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultCreator struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetTaskListFilterResponseBodyObjectResultCreator) SetName(v string) *GetTaskListFilterResponseBodyObjectResultCreator {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCreator) SetId(v string) *GetTaskListFilterResponseBodyObjectResultCreator {
	s.Id = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultExecutor struct {
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *GetTaskListFilterResponseBodyObjectResultExecutor) SetName(v string) *GetTaskListFilterResponseBodyObjectResultExecutor {
	s.Name = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultExecutor) SetId(v string) *GetTaskListFilterResponseBodyObjectResultExecutor {
	s.Id = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultTaskFlowStatus struct {
	TaskFlowId *string `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
	Pos        *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	Kind       *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetTaskFlowId(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.TaskFlowId = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetPos(v int32) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Pos = &v
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

func (s *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus) SetId(v string) *GetTaskListFilterResponseBodyObjectResultTaskFlowStatus {
	s.Id = &v
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

type GetTaskListFilterResponseBodyObjectResultCustomfields struct {
	Type          *string                                                       `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         []*GetTaskListFilterResponseBodyObjectResultCustomfieldsValue `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
	Values        *string                                                       `json:"Values,omitempty" xml:"Values,omitempty"`
	CustomfieldId *string                                                       `json:"CustomfieldId,omitempty" xml:"CustomfieldId,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfields) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfields) GoString() string {
	return s.String()
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

func (s *GetTaskListFilterResponseBodyObjectResultCustomfields) SetCustomfieldId(v string) *GetTaskListFilterResponseBodyObjectResultCustomfields {
	s.CustomfieldId = &v
	return s
}

type GetTaskListFilterResponseBodyObjectResultCustomfieldsValue struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) String() string {
	return tea.Prettify(s)
}

func (s GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) GoString() string {
	return s.String()
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) SetTitle(v string) *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue {
	s.Title = &v
	return s
}

func (s *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue) SetId(v string) *GetTaskListFilterResponseBodyObjectResultCustomfieldsValue {
	s.Id = &v
	return s
}

type GetTaskListFilterResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskListFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskListFilterResponse) SetBody(v *GetTaskListFilterResponseBody) *GetTaskListFilterResponse {
	s.Body = v
	return s
}

type GetUserByAliyunUidRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetUserByAliyunUidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunUidRequest) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunUidRequest) SetOrgId(v string) *GetUserByAliyunUidRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserByAliyunUidRequest) SetUserPk(v string) *GetUserByAliyunUidRequest {
	s.UserPk = &v
	return s
}

type GetUserByAliyunUidResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetUserByAliyunUidResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetUserByAliyunUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunUidResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunUidResponseBody) SetRequestId(v string) *GetUserByAliyunUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByAliyunUidResponseBody) SetErrorMsg(v string) *GetUserByAliyunUidResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserByAliyunUidResponseBody) SetObject(v *GetUserByAliyunUidResponseBodyObject) *GetUserByAliyunUidResponseBody {
	s.Object = v
	return s
}

func (s *GetUserByAliyunUidResponseBody) SetErrorCode(v string) *GetUserByAliyunUidResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserByAliyunUidResponseBody) SetSuccessful(v bool) *GetUserByAliyunUidResponseBody {
	s.Successful = &v
	return s
}

type GetUserByAliyunUidResponseBodyObject struct {
	AliyunPk  *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetUserByAliyunUidResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunUidResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunUidResponseBodyObject) SetAliyunPk(v string) *GetUserByAliyunUidResponseBodyObject {
	s.AliyunPk = &v
	return s
}

func (s *GetUserByAliyunUidResponseBodyObject) SetEmail(v string) *GetUserByAliyunUidResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetUserByAliyunUidResponseBodyObject) SetAvatarUrl(v string) *GetUserByAliyunUidResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserByAliyunUidResponseBodyObject) SetName(v string) *GetUserByAliyunUidResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetUserByAliyunUidResponseBodyObject) SetId(v string) *GetUserByAliyunUidResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetUserByAliyunUidResponseBodyObject) SetPhone(v string) *GetUserByAliyunUidResponseBodyObject {
	s.Phone = &v
	return s
}

type GetUserByAliyunUidResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserByAliyunUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByAliyunUidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunUidResponse) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunUidResponse) SetHeaders(v map[string]*string) *GetUserByAliyunUidResponse {
	s.Headers = v
	return s
}

func (s *GetUserByAliyunUidResponse) SetBody(v *GetUserByAliyunUidResponseBody) *GetUserByAliyunUidResponse {
	s.Body = v
	return s
}

type GetUserNameRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserNameRequest) GoString() string {
	return s.String()
}

func (s *GetUserNameRequest) SetOrgId(v string) *GetUserNameRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserNameRequest) SetUserId(v string) *GetUserNameRequest {
	s.UserId = &v
	return s
}

type GetUserNameResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetUserNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserNameResponseBody) SetRequestId(v string) *GetUserNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserNameResponseBody) SetErrorMsg(v string) *GetUserNameResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserNameResponseBody) SetObject(v string) *GetUserNameResponseBody {
	s.Object = &v
	return s
}

func (s *GetUserNameResponseBody) SetErrorCode(v string) *GetUserNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserNameResponseBody) SetSuccessful(v bool) *GetUserNameResponseBody {
	s.Successful = &v
	return s
}

type GetUserNameResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserNameResponse) GoString() string {
	return s.String()
}

func (s *GetUserNameResponse) SetHeaders(v map[string]*string) *GetUserNameResponse {
	s.Headers = v
	return s
}

func (s *GetUserNameResponse) SetBody(v *GetUserNameResponseBody) *GetUserNameResponse {
	s.Body = v
	return s
}

type InsertDevopsUserRequest struct {
	UserPk   *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Phone    *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s InsertDevopsUserRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsUserRequest) GoString() string {
	return s.String()
}

func (s *InsertDevopsUserRequest) SetUserPk(v string) *InsertDevopsUserRequest {
	s.UserPk = &v
	return s
}

func (s *InsertDevopsUserRequest) SetUserName(v string) *InsertDevopsUserRequest {
	s.UserName = &v
	return s
}

func (s *InsertDevopsUserRequest) SetPhone(v string) *InsertDevopsUserRequest {
	s.Phone = &v
	return s
}

func (s *InsertDevopsUserRequest) SetEmail(v string) *InsertDevopsUserRequest {
	s.Email = &v
	return s
}

type InsertDevopsUserResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertDevopsUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsUserResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDevopsUserResponseBody) SetRequestId(v string) *InsertDevopsUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertDevopsUserResponseBody) SetObject(v string) *InsertDevopsUserResponseBody {
	s.Object = &v
	return s
}

func (s *InsertDevopsUserResponseBody) SetErrorCode(v string) *InsertDevopsUserResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertDevopsUserResponseBody) SetErrorMessage(v string) *InsertDevopsUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InsertDevopsUserResponseBody) SetSuccess(v bool) *InsertDevopsUserResponseBody {
	s.Success = &v
	return s
}

type InsertDevopsUserResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertDevopsUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertDevopsUserResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsUserResponse) GoString() string {
	return s.String()
}

func (s *InsertDevopsUserResponse) SetHeaders(v map[string]*string) *InsertDevopsUserResponse {
	s.Headers = v
	return s
}

func (s *InsertDevopsUserResponse) SetBody(v *InsertDevopsUserResponseBody) *InsertDevopsUserResponse {
	s.Body = v
	return s
}

type InsertPipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

func (s *InsertPipelineMemberRequest) SetUserPk(v string) *InsertPipelineMemberRequest {
	s.UserPk = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetUserId(v string) *InsertPipelineMemberRequest {
	s.UserId = &v
	return s
}

func (s *InsertPipelineMemberRequest) SetRoleName(v string) *InsertPipelineMemberRequest {
	s.RoleName = &v
	return s
}

type InsertPipelineMemberResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertPipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertPipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *InsertPipelineMemberResponseBody) SetRequestId(v string) *InsertPipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetObject(v bool) *InsertPipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetErrorCode(v string) *InsertPipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetErrorMessage(v string) *InsertPipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InsertPipelineMemberResponseBody) SetSuccess(v bool) *InsertPipelineMemberResponseBody {
	s.Success = &v
	return s
}

type InsertPipelineMemberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertPipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *InsertPipelineMemberResponse) SetBody(v *InsertPipelineMemberResponseBody) *InsertPipelineMemberResponse {
	s.Body = v
	return s
}

type InsertProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Members   *string `json:"Members,omitempty" xml:"Members,omitempty"`
}

func (s InsertProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *InsertProjectMembersRequest) SetOrgId(v string) *InsertProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *InsertProjectMembersRequest) SetProjectId(v string) *InsertProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *InsertProjectMembersRequest) SetMembers(v string) *InsertProjectMembersRequest {
	s.Members = &v
	return s
}

type InsertProjectMembersResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s InsertProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *InsertProjectMembersResponseBody) SetRequestId(v string) *InsertProjectMembersResponseBody {
	s.RequestId = &v
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

func (s *InsertProjectMembersResponseBody) SetErrorCode(v string) *InsertProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertProjectMembersResponseBody) SetSuccessful(v bool) *InsertProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type InsertProjectMembersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *InsertProjectMembersResponse) SetBody(v *InsertProjectMembersResponseBody) *InsertProjectMembersResponse {
	s.Body = v
	return s
}

type ListCommonGroupRequest struct {
	OrgId        *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SmartGroupId *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
}

func (s ListCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupRequest) GoString() string {
	return s.String()
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

func (s *ListCommonGroupRequest) SetAll(v bool) *ListCommonGroupRequest {
	s.All = &v
	return s
}

type ListCommonGroupResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonGroupResponseBody) SetRequestId(v string) *ListCommonGroupResponseBody {
	s.RequestId = &v
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

func (s *ListCommonGroupResponseBody) SetErrorCode(v string) *ListCommonGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCommonGroupResponseBody) SetSuccessful(v bool) *ListCommonGroupResponseBody {
	s.Successful = &v
	return s
}

type ListCommonGroupResponseBodyObject struct {
	ResourceCount *int32  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	SmartGroupId  *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
	Pos           *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IsRoot        *bool   `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	Pinyin        *string `json:"Pinyin,omitempty" xml:"Pinyin,omitempty"`
	CreatorId     *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id            *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListCommonGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListCommonGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListCommonGroupResponseBodyObject) SetResourceCount(v int32) *ListCommonGroupResponseBodyObject {
	s.ResourceCount = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetSmartGroupId(v string) *ListCommonGroupResponseBodyObject {
	s.SmartGroupId = &v
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

func (s *ListCommonGroupResponseBodyObject) SetIsRoot(v bool) *ListCommonGroupResponseBodyObject {
	s.IsRoot = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetPinyin(v string) *ListCommonGroupResponseBodyObject {
	s.Pinyin = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetCreatorId(v string) *ListCommonGroupResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetName(v string) *ListCommonGroupResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListCommonGroupResponseBodyObject) SetId(v string) *ListCommonGroupResponseBodyObject {
	s.Id = &v
	return s
}

type ListCommonGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       []map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode    *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetObject(v []map[string]interface{}) *ListCredentialsResponseBody {
	s.Object = v
	return s
}

func (s *ListCredentialsResponseBody) SetErrorCode(v string) *ListCredentialsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCredentialsResponseBody) SetErrorMessage(v string) *ListCredentialsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCredentialsResponseBody) SetSuccess(v bool) *ListCredentialsResponseBody {
	s.Success = &v
	return s
}

type ListCredentialsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListCredentialsResponse) SetBody(v *ListCredentialsResponseBody) *ListCredentialsResponse {
	s.Body = v
	return s
}

type ListDevopsProjectSprintsRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
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

func (s *ListDevopsProjectSprintsRequest) SetProjectId(v string) *ListDevopsProjectSprintsRequest {
	s.ProjectId = &v
	return s
}

type ListDevopsProjectSprintsResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                       `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectSprintsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                       `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                         `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBody) SetRequestId(v string) *ListDevopsProjectSprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetErrorMsg(v string) *ListDevopsProjectSprintsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetObject(v []*ListDevopsProjectSprintsResponseBodyObject) *ListDevopsProjectSprintsResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetErrorCode(v string) *ListDevopsProjectSprintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBody) SetSuccessful(v bool) *ListDevopsProjectSprintsResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectSprintsResponseBodyObject struct {
	Status       *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                             `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                             `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Accomplished *string                                             `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                               `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                             `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                             `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                             `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *ListDevopsProjectSprintsResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetStatus(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Status = &v
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

func (s *ListDevopsProjectSprintsResponseBodyObject) SetCreatorId(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetAccomplished(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetIsDeleted(v bool) *ListDevopsProjectSprintsResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetUpdated(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetDueDate(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetName(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetCreated(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetPlanToDo(v *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) *ListDevopsProjectSprintsResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObject) SetId(v string) *ListDevopsProjectSprintsResponseBodyObject {
	s.Id = &v
	return s
}

type ListDevopsProjectSprintsResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s ListDevopsProjectSprintsResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectSprintsResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetTasks(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *ListDevopsProjectSprintsResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *ListDevopsProjectSprintsResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type ListDevopsProjectSprintsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsProjectSprintsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                        `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectTaskFlowResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                          `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetRequestId(v string) *ListDevopsProjectTaskFlowResponseBody {
	s.RequestId = &v
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

func (s *ListDevopsProjectTaskFlowResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBody) SetSuccessful(v bool) *ListDevopsProjectTaskFlowResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTaskFlowResponseBodyObject struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsProjectTaskFlowResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetType(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetName(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectTaskFlowResponseBodyObject) SetId(v string) *ListDevopsProjectTaskFlowResponseBodyObject {
	s.Id = &v
	return s
}

type ListDevopsProjectTaskFlowResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsProjectTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectTaskFlowStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskFlowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetRequestId(v string) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.RequestId = &v
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

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBody) SetSuccessful(v bool) *ListDevopsProjectTaskFlowStatusResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTaskFlowStatusResponseBodyObject struct {
	TaskflowId      *string `json:"TaskflowId,omitempty" xml:"TaskflowId,omitempty"`
	Kind            *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	Pos             *int32  `json:"Pos,omitempty" xml:"Pos,omitempty"`
	IsDeleted       *bool   `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated         *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	CreatorId       *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Created         *string `json:"Created,omitempty" xml:"Created,omitempty"`
	RejectStatusIds *string `json:"RejectStatusIds,omitempty" xml:"RejectStatusIds,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsProjectTaskFlowStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskFlowStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetTaskflowId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.TaskflowId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetKind(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Kind = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetPos(v int32) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetIsDeleted(v bool) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetUpdated(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetCreatorId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetName(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetCreated(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetRejectStatusIds(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.RejectStatusIds = &v
	return s
}

func (s *ListDevopsProjectTaskFlowStatusResponseBodyObject) SetId(v string) *ListDevopsProjectTaskFlowStatusResponseBodyObject {
	s.Id = &v
	return s
}

type ListDevopsProjectTaskFlowStatusResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsProjectTaskFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *ListDevopsProjectTaskListResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTaskListResponseBody) SetRequestId(v string) *ListDevopsProjectTaskListResponseBody {
	s.RequestId = &v
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

func (s *ListDevopsProjectTaskListResponseBody) SetErrorCode(v string) *ListDevopsProjectTaskListResponseBody {
	s.ErrorCode = &v
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsProjectTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDevopsProjectTaskListResponse) SetBody(v *ListDevopsProjectTaskListResponseBody) *ListDevopsProjectTaskListResponse {
	s.Body = v
	return s
}

type ListDevopsProjectTasksRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s ListDevopsProjectTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTasksRequest) SetOrgId(v string) *ListDevopsProjectTasksRequest {
	s.OrgId = &v
	return s
}

func (s *ListDevopsProjectTasksRequest) SetProjectIds(v string) *ListDevopsProjectTasksRequest {
	s.ProjectIds = &v
	return s
}

type ListDevopsProjectTasksResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsProjectTasksResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                       `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsProjectTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTasksResponseBody) SetRequestId(v string) *ListDevopsProjectTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBody) SetErrorMsg(v string) *ListDevopsProjectTasksResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBody) SetObject(v []*ListDevopsProjectTasksResponseBodyObject) *ListDevopsProjectTasksResponseBody {
	s.Object = v
	return s
}

func (s *ListDevopsProjectTasksResponseBody) SetErrorCode(v string) *ListDevopsProjectTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBody) SetSuccessful(v bool) *ListDevopsProjectTasksResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsProjectTasksResponseBodyObject struct {
	TaskgroupId *string `json:"TaskgroupId,omitempty" xml:"TaskgroupId,omitempty"`
	TasklistId  *string `json:"TasklistId,omitempty" xml:"TasklistId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ModifierId  *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	Updated     *string `json:"Updated,omitempty" xml:"Updated,omitempty"`
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Created     *string `json:"Created,omitempty" xml:"Created,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsProjectTasksResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTasksResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetTaskgroupId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.TaskgroupId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetTasklistId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetProjectId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetModifierId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.ModifierId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetUpdated(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetCreatorId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetCreated(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetName(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListDevopsProjectTasksResponseBodyObject) SetId(v string) *ListDevopsProjectTasksResponseBodyObject {
	s.Id = &v
	return s
}

type ListDevopsProjectTasksResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsProjectTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDevopsProjectTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsProjectTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDevopsProjectTasksResponse) SetHeaders(v map[string]*string) *ListDevopsProjectTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDevopsProjectTasksResponse) SetBody(v *ListDevopsProjectTasksResponseBody) *ListDevopsProjectTasksResponse {
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
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                            `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListDevopsScenarioFieldConfigResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                              `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListDevopsScenarioFieldConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetRequestId(v string) *ListDevopsScenarioFieldConfigResponseBody {
	s.RequestId = &v
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

func (s *ListDevopsScenarioFieldConfigResponseBody) SetErrorCode(v string) *ListDevopsScenarioFieldConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBody) SetSuccessful(v bool) *ListDevopsScenarioFieldConfigResponseBody {
	s.Successful = &v
	return s
}

type ListDevopsScenarioFieldConfigResponseBodyObject struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDevopsScenarioFieldConfigResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListDevopsScenarioFieldConfigResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListDevopsScenarioFieldConfigResponseBodyObject) SetType(v string) *ListDevopsScenarioFieldConfigResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListDevopsScenarioFieldConfigResponseBodyObject) SetId(v string) *ListDevopsScenarioFieldConfigResponseBodyObject {
	s.Id = &v
	return s
}

type ListDevopsScenarioFieldConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDevopsScenarioFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDevopsScenarioFieldConfigResponse) SetBody(v *ListDevopsScenarioFieldConfigResponseBody) *ListDevopsScenarioFieldConfigResponse {
	s.Body = v
	return s
}

type ListPipelinesRequest struct {
	OrgId            *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineName     *string `json:"PipelineName,omitempty" xml:"PipelineName,omitempty"`
	Creators         *string `json:"Creators,omitempty" xml:"Creators,omitempty"`
	Operators        *string `json:"Operators,omitempty" xml:"Operators,omitempty"`
	ResultStatusList *string `json:"ResultStatusList,omitempty" xml:"ResultStatusList,omitempty"`
	CreateStartTime  *string `json:"CreateStartTime,omitempty" xml:"CreateStartTime,omitempty"`
	CreateEndTime    *string `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	ExecuteStartTime *string `json:"ExecuteStartTime,omitempty" xml:"ExecuteStartTime,omitempty"`
	ExecuteEndTime   *string `json:"ExecuteEndTime,omitempty" xml:"ExecuteEndTime,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageStart        *int32  `json:"PageStart,omitempty" xml:"PageStart,omitempty"`
	UserPk           *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) SetOrgId(v string) *ListPipelinesRequest {
	s.OrgId = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineName(v string) *ListPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesRequest) SetCreators(v string) *ListPipelinesRequest {
	s.Creators = &v
	return s
}

func (s *ListPipelinesRequest) SetOperators(v string) *ListPipelinesRequest {
	s.Operators = &v
	return s
}

func (s *ListPipelinesRequest) SetResultStatusList(v string) *ListPipelinesRequest {
	s.ResultStatusList = &v
	return s
}

func (s *ListPipelinesRequest) SetCreateStartTime(v string) *ListPipelinesRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetCreateEndTime(v string) *ListPipelinesRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteStartTime(v string) *ListPipelinesRequest {
	s.ExecuteStartTime = &v
	return s
}

func (s *ListPipelinesRequest) SetExecuteEndTime(v string) *ListPipelinesRequest {
	s.ExecuteEndTime = &v
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

func (s *ListPipelinesRequest) SetUserPk(v string) *ListPipelinesRequest {
	s.UserPk = &v
	return s
}

type ListPipelinesResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetObject(v map[string]interface{}) *ListPipelinesResponseBody {
	s.Object = v
	return s
}

func (s *ListPipelinesResponseBody) SetErrorCode(v string) *ListPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelinesResponseBody) SetErrorMessage(v string) *ListPipelinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelinesResponseBody) SetSuccess(v bool) *ListPipelinesResponseBody {
	s.Success = &v
	return s
}

type ListPipelinesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPipelinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectCustomFieldsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectCustomFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBody) SetRequestId(v string) *ListProjectCustomFieldsResponseBody {
	s.RequestId = &v
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

func (s *ListProjectCustomFieldsResponseBody) SetErrorCode(v string) *ListProjectCustomFieldsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBody) SetSuccessful(v bool) *ListProjectCustomFieldsResponseBody {
	s.Successful = &v
	return s
}

type ListProjectCustomFieldsResponseBodyObject struct {
	Type          *string                                            `json:"Type,omitempty" xml:"Type,omitempty"`
	Subtype       *string                                            `json:"Subtype,omitempty" xml:"Subtype,omitempty"`
	Values        []*ListProjectCustomFieldsResponseBodyObjectValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
	CustomFieldId *string                                            `json:"CustomFieldId,omitempty" xml:"CustomFieldId,omitempty"`
	Name          *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListProjectCustomFieldsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetType(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetSubtype(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Subtype = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetValues(v []*ListProjectCustomFieldsResponseBodyObjectValues) *ListProjectCustomFieldsResponseBodyObject {
	s.Values = v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetCustomFieldId(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.CustomFieldId = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObject) SetName(v string) *ListProjectCustomFieldsResponseBodyObject {
	s.Name = &v
	return s
}

type ListProjectCustomFieldsResponseBodyObjectValues struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListProjectCustomFieldsResponseBodyObjectValues) String() string {
	return tea.Prettify(s)
}

func (s ListProjectCustomFieldsResponseBodyObjectValues) GoString() string {
	return s.String()
}

func (s *ListProjectCustomFieldsResponseBodyObjectValues) SetValue(v string) *ListProjectCustomFieldsResponseBodyObjectValues {
	s.Value = &v
	return s
}

func (s *ListProjectCustomFieldsResponseBodyObjectValues) SetId(v string) *ListProjectCustomFieldsResponseBodyObjectValues {
	s.Id = &v
	return s
}

type ListProjectCustomFieldsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectCustomFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       []map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode    *string                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListServiceConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceConnectionsResponseBody) SetRequestId(v string) *ListServiceConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetObject(v []map[string]interface{}) *ListServiceConnectionsResponseBody {
	s.Object = v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetErrorCode(v string) *ListServiceConnectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetErrorMessage(v string) *ListServiceConnectionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListServiceConnectionsResponseBody) SetSuccess(v bool) *ListServiceConnectionsResponseBody {
	s.Success = &v
	return s
}

type ListServiceConnectionsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListServiceConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                             `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListSmartGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                               `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListSmartGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmartGroupResponseBody) SetRequestId(v string) *ListSmartGroupResponseBody {
	s.RequestId = &v
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

func (s *ListSmartGroupResponseBody) SetErrorCode(v string) *ListSmartGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSmartGroupResponseBody) SetSuccessful(v bool) *ListSmartGroupResponseBody {
	s.Successful = &v
	return s
}

type ListSmartGroupResponseBodyObject struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListSmartGroupResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListSmartGroupResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListSmartGroupResponseBodyObject) SetType(v string) *ListSmartGroupResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListSmartGroupResponseBodyObject) SetId(v string) *ListSmartGroupResponseBodyObject {
	s.Id = &v
	return s
}

type ListSmartGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSmartGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       []*ListUserOrganizationResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode    *string                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserOrganizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationResponseBody) SetRequestId(v string) *ListUserOrganizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetObject(v []*ListUserOrganizationResponseBodyObject) *ListUserOrganizationResponseBody {
	s.Object = v
	return s
}

func (s *ListUserOrganizationResponseBody) SetErrorCode(v string) *ListUserOrganizationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetErrorMessage(v string) *ListUserOrganizationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListUserOrganizationResponseBody) SetSuccess(v bool) *ListUserOrganizationResponseBody {
	s.Success = &v
	return s
}

type ListUserOrganizationResponseBodyObject struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUserOrganizationResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListUserOrganizationResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListUserOrganizationResponseBodyObject) SetName(v string) *ListUserOrganizationResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListUserOrganizationResponseBodyObject) SetId(v string) *ListUserOrganizationResponseBodyObject {
	s.Id = &v
	return s
}

type ListUserOrganizationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserOrganizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUserOrganizationResponse) SetBody(v *ListUserOrganizationResponseBody) *ListUserOrganizationResponse {
	s.Body = v
	return s
}

type TransferPipelineOwnerRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	NewOwnerId *string `json:"NewOwnerId,omitempty" xml:"NewOwnerId,omitempty"`
}

func (s TransferPipelineOwnerRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferPipelineOwnerRequest) GoString() string {
	return s.String()
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

func (s *TransferPipelineOwnerRequest) SetNewOwnerId(v string) *TransferPipelineOwnerRequest {
	s.NewOwnerId = &v
	return s
}

type TransferPipelineOwnerResponseBody struct {
	RequestId    *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       map[string]interface{} `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferPipelineOwnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferPipelineOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPipelineOwnerResponseBody) SetRequestId(v string) *TransferPipelineOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetObject(v map[string]interface{}) *TransferPipelineOwnerResponseBody {
	s.Object = v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetErrorCode(v string) *TransferPipelineOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetErrorMessage(v string) *TransferPipelineOwnerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TransferPipelineOwnerResponseBody) SetSuccess(v bool) *TransferPipelineOwnerResponseBody {
	s.Success = &v
	return s
}

type TransferPipelineOwnerResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferPipelineOwnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TransferPipelineOwnerResponse) SetBody(v *TransferPipelineOwnerResponseBody) *TransferPipelineOwnerResponse {
	s.Body = v
	return s
}

type UpdateCommonGroupRequest struct {
	OrgId         *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SmartGroupId  *string `json:"SmartGroupId,omitempty" xml:"SmartGroupId,omitempty"`
	CommonGroupId *string `json:"CommonGroupId,omitempty" xml:"CommonGroupId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCommonGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupRequest) SetOrgId(v string) *UpdateCommonGroupRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetProjectId(v string) *UpdateCommonGroupRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetDescription(v string) *UpdateCommonGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetSmartGroupId(v string) *UpdateCommonGroupRequest {
	s.SmartGroupId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetCommonGroupId(v string) *UpdateCommonGroupRequest {
	s.CommonGroupId = &v
	return s
}

func (s *UpdateCommonGroupRequest) SetName(v string) *UpdateCommonGroupRequest {
	s.Name = &v
	return s
}

type UpdateCommonGroupResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *UpdateCommonGroupResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateCommonGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommonGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommonGroupResponseBody) SetRequestId(v string) *UpdateCommonGroupResponseBody {
	s.RequestId = &v
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

func (s *UpdateCommonGroupResponseBody) SetErrorCode(v string) *UpdateCommonGroupResponseBody {
	s.ErrorCode = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateCommonGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateCommonGroupResponse) SetBody(v *UpdateCommonGroupResponseBody) *UpdateCommonGroupResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDevopsProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectRequest) SetOrgId(v string) *UpdateDevopsProjectRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetName(v string) *UpdateDevopsProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetDescription(v string) *UpdateDevopsProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateDevopsProjectRequest) SetProjectId(v string) *UpdateDevopsProjectRequest {
	s.ProjectId = &v
	return s
}

type UpdateDevopsProjectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDevopsProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectResponseBody) SetRequestId(v string) *UpdateDevopsProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetObject(v string) *UpdateDevopsProjectResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetErrorCode(v string) *UpdateDevopsProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetErrorMessage(v string) *UpdateDevopsProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDevopsProjectResponseBody) SetSuccess(v bool) *UpdateDevopsProjectResponseBody {
	s.Success = &v
	return s
}

type UpdateDevopsProjectResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevopsProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDevopsProjectResponse) SetBody(v *UpdateDevopsProjectResponseBody) *UpdateDevopsProjectResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectSprintRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	SprintId    *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s UpdateDevopsProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectSprintRequest) SetOrgId(v string) *UpdateDevopsProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetName(v string) *UpdateDevopsProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetDescription(v string) *UpdateDevopsProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetProjectId(v string) *UpdateDevopsProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetExecutorId(v string) *UpdateDevopsProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetStartDate(v string) *UpdateDevopsProjectSprintRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetDueDate(v string) *UpdateDevopsProjectSprintRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateDevopsProjectSprintRequest) SetSprintId(v string) *UpdateDevopsProjectSprintRequest {
	s.SprintId = &v
	return s
}

type UpdateDevopsProjectSprintResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateDevopsProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectSprintResponseBody) SetRequestId(v string) *UpdateDevopsProjectSprintResponseBody {
	s.RequestId = &v
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

func (s *UpdateDevopsProjectSprintResponseBody) SetErrorCode(v string) *UpdateDevopsProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectSprintResponseBody) SetSuccessful(v bool) *UpdateDevopsProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type UpdateDevopsProjectSprintResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevopsProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDevopsProjectSprintResponse) SetBody(v *UpdateDevopsProjectSprintResponseBody) *UpdateDevopsProjectSprintResponse {
	s.Body = v
	return s
}

type UpdateDevopsProjectTaskRequest struct {
	OrgId                  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Content                *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProjectId              *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId             *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate              *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate                *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	ScenarioFiieldConfigId *string `json:"ScenarioFiieldConfigId,omitempty" xml:"ScenarioFiieldConfigId,omitempty"`
	TaskFlowStatusId       *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	Note                   *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Priority               *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Visible                *string `json:"Visible,omitempty" xml:"Visible,omitempty"`
	ParentTaskId           *string `json:"ParentTaskId,omitempty" xml:"ParentTaskId,omitempty"`
	SprintId               *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateDevopsProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectTaskRequest) SetOrgId(v string) *UpdateDevopsProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetContent(v string) *UpdateDevopsProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetProjectId(v string) *UpdateDevopsProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetExecutorId(v string) *UpdateDevopsProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetStartDate(v string) *UpdateDevopsProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetDueDate(v string) *UpdateDevopsProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetScenarioFiieldConfigId(v string) *UpdateDevopsProjectTaskRequest {
	s.ScenarioFiieldConfigId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetTaskFlowStatusId(v string) *UpdateDevopsProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetNote(v string) *UpdateDevopsProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetPriority(v int32) *UpdateDevopsProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetVisible(v string) *UpdateDevopsProjectTaskRequest {
	s.Visible = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetParentTaskId(v string) *UpdateDevopsProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetSprintId(v string) *UpdateDevopsProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateDevopsProjectTaskRequest) SetTaskId(v string) *UpdateDevopsProjectTaskRequest {
	s.TaskId = &v
	return s
}

type UpdateDevopsProjectTaskResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateDevopsProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDevopsProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDevopsProjectTaskResponseBody) SetRequestId(v string) *UpdateDevopsProjectTaskResponseBody {
	s.RequestId = &v
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

func (s *UpdateDevopsProjectTaskResponseBody) SetErrorCode(v string) *UpdateDevopsProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDevopsProjectTaskResponseBody) SetSuccessful(v bool) *UpdateDevopsProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type UpdateDevopsProjectTaskResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDevopsProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateDevopsProjectTaskResponse) SetBody(v *UpdateDevopsProjectTaskResponseBody) *UpdateDevopsProjectTaskResponse {
	s.Body = v
	return s
}

type UpdatePipelineMemberRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	PipelineId *int64  `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	UserPk     *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RoleName   *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
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

func (s *UpdatePipelineMemberRequest) SetUserPk(v string) *UpdatePipelineMemberRequest {
	s.UserPk = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetUserId(v string) *UpdatePipelineMemberRequest {
	s.UserId = &v
	return s
}

func (s *UpdatePipelineMemberRequest) SetRoleName(v string) *UpdatePipelineMemberRequest {
	s.RoleName = &v
	return s
}

type UpdatePipelineMemberResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePipelineMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePipelineMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePipelineMemberResponseBody) SetRequestId(v string) *UpdatePipelineMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetObject(v bool) *UpdatePipelineMemberResponseBody {
	s.Object = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetErrorCode(v string) *UpdatePipelineMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetErrorMessage(v string) *UpdatePipelineMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePipelineMemberResponseBody) SetSuccess(v bool) *UpdatePipelineMemberResponseBody {
	s.Success = &v
	return s
}

type UpdatePipelineMemberResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePipelineMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdatePipelineMemberResponse) SetBody(v *UpdatePipelineMemberResponseBody) *UpdatePipelineMemberResponse {
	s.Body = v
	return s
}

type UpdateTaskDetailRequest struct {
	OrgId             *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId        *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate         *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate           *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	TaskFlowStatusId  *string `json:"TaskFlowStatusId,omitempty" xml:"TaskFlowStatusId,omitempty"`
	Note              *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Priority          *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SprintId          *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	TaskId            *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WorkTimes         *int64  `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	CustomFieldId     *string `json:"CustomFieldId,omitempty" xml:"CustomFieldId,omitempty"`
	CustomFieldValues *string `json:"CustomFieldValues,omitempty" xml:"CustomFieldValues,omitempty"`
	StoryPoint        *string `json:"StoryPoint,omitempty" xml:"StoryPoint,omitempty"`
	TagIds            *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	DelInvolvers      *string `json:"DelInvolvers,omitempty" xml:"DelInvolvers,omitempty"`
	AddInvolvers      *string `json:"AddInvolvers,omitempty" xml:"AddInvolvers,omitempty"`
}

func (s UpdateTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskDetailRequest) SetOrgId(v string) *UpdateTaskDetailRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetContent(v string) *UpdateTaskDetailRequest {
	s.Content = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetProjectId(v string) *UpdateTaskDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetExecutorId(v string) *UpdateTaskDetailRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetStartDate(v string) *UpdateTaskDetailRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetDueDate(v string) *UpdateTaskDetailRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetTaskFlowStatusId(v string) *UpdateTaskDetailRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetNote(v string) *UpdateTaskDetailRequest {
	s.Note = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetPriority(v int64) *UpdateTaskDetailRequest {
	s.Priority = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetSprintId(v string) *UpdateTaskDetailRequest {
	s.SprintId = &v
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

func (s *UpdateTaskDetailRequest) SetCustomFieldId(v string) *UpdateTaskDetailRequest {
	s.CustomFieldId = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetCustomFieldValues(v string) *UpdateTaskDetailRequest {
	s.CustomFieldValues = &v
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

func (s *UpdateTaskDetailRequest) SetDelInvolvers(v string) *UpdateTaskDetailRequest {
	s.DelInvolvers = &v
	return s
}

func (s *UpdateTaskDetailRequest) SetAddInvolvers(v string) *UpdateTaskDetailRequest {
	s.AddInvolvers = &v
	return s
}

type UpdateTaskDetailResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTaskDetailResponseBody) SetRequestId(v string) *UpdateTaskDetailResponseBody {
	s.RequestId = &v
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

func (s *UpdateTaskDetailResponseBody) SetErrorCode(v string) *UpdateTaskDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTaskDetailResponseBody) SetSuccessful(v bool) *UpdateTaskDetailResponseBody {
	s.Successful = &v
	return s
}

type UpdateTaskDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (client *Client) BatchInsertMembersWithOptions(request *BatchInsertMembersRequest, runtime *util.RuntimeOptions) (_result *BatchInsertMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchInsertMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchInsertMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelPipelineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelPipeline"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CheckAliyunAccountExistsWithOptions(request *CheckAliyunAccountExistsRequest, runtime *util.RuntimeOptions) (_result *CheckAliyunAccountExistsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckAliyunAccountExistsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckAliyunAccountExists"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAliyunAccountExists(request *CheckAliyunAccountExistsRequest) (_result *CheckAliyunAccountExistsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAliyunAccountExistsResponse{}
	_body, _err := client.CheckAliyunAccountExistsWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCommonGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCommonGroup"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCredential"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateDevopsOrganizationWithOptions(request *CreateDevopsOrganizationRequest, runtime *util.RuntimeOptions) (_result *CreateDevopsOrganizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevopsOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevopsOrganization"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevopsOrganization(request *CreateDevopsOrganizationRequest) (_result *CreateDevopsOrganizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevopsOrganizationResponse{}
	_body, _err := client.CreateDevopsOrganizationWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevopsProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevopsProject"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevopsProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevopsProjectSprint"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevopsProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevopsProjectTask"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreatePipelineWithOptions(request *CreatePipelineRequest, runtime *util.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePipeline"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePipeline(request *CreatePipelineRequest) (_result *CreatePipelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CreatePipelineWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceConnection"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCommonGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCommonGroup"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevopsOrganizationMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevopsOrganizationMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevopsProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevopsProject"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevopsProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevopsProjectMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevopsProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevopsProjectSprint"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDevopsProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDevopsProjectTask"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePipelineMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePipelineMember"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecutePipelineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecutePipeline"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevopsOrganizationMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevopsOrganizationMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevopsProjectInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevopsProjectInfo"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevopsProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevopsProjectMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevopsProjectSprintInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevopsProjectSprintInfo"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDevopsProjectTaskInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDevopsProjectTaskInfo"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetLastWorkspaceWithOptions(request *GetLastWorkspaceRequest, runtime *util.RuntimeOptions) (_result *GetLastWorkspaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLastWorkspaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLastWorkspace"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLastWorkspace(request *GetLastWorkspaceRequest) (_result *GetLastWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLastWorkspaceResponse{}
	_body, _err := client.GetLastWorkspaceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineInstanceBuildNumberStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineInstanceBuildNumberStatus"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineInstanceGroupStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineInstanceGroupStatus"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineInstanceInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineInstanceInfo"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineInstanceStatus"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetPipelineInstHistoryWithOptions(request *GetPipelineInstHistoryRequest, runtime *util.RuntimeOptions) (_result *GetPipelineInstHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineInstHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineInstHistory"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetPipelineLogWithOptions(request *GetPipelineLogRequest, runtime *util.RuntimeOptions) (_result *GetPipelineLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineLog"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipelineStepLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipelineStepLog"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPipleineLatestInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPipleineLatestInstanceStatus"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectOptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectOption"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskDetailActivityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskDetailActivity"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskDetailBaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskDetailBase"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskListFilterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskListFilter"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetUserByAliyunUidWithOptions(request *GetUserByAliyunUidRequest, runtime *util.RuntimeOptions) (_result *GetUserByAliyunUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserByAliyunUidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserByAliyunUid"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByAliyunUid(request *GetUserByAliyunUidRequest) (_result *GetUserByAliyunUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserByAliyunUidResponse{}
	_body, _err := client.GetUserByAliyunUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserNameWithOptions(request *GetUserNameRequest, runtime *util.RuntimeOptions) (_result *GetUserNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserName"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserName(request *GetUserNameRequest) (_result *GetUserNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserNameResponse{}
	_body, _err := client.GetUserNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertDevopsUserWithOptions(request *InsertDevopsUserRequest, runtime *util.RuntimeOptions) (_result *InsertDevopsUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertDevopsUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertDevopsUser"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertDevopsUser(request *InsertDevopsUserRequest) (_result *InsertDevopsUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertDevopsUserResponse{}
	_body, _err := client.InsertDevopsUserWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertPipelineMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertPipelineMember"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertProjectMembers"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCommonGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCommonGroup"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCredentials"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsProjectSprintsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsProjectSprints"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsProjectTaskFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsProjectTaskFlow"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsProjectTaskFlowStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsProjectTaskFlowStatus"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsProjectTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsProjectTaskList"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListDevopsProjectTasksWithOptions(request *ListDevopsProjectTasksRequest, runtime *util.RuntimeOptions) (_result *ListDevopsProjectTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsProjectTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsProjectTasks"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDevopsProjectTasks(request *ListDevopsProjectTasksRequest) (_result *ListDevopsProjectTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDevopsProjectTasksResponse{}
	_body, _err := client.ListDevopsProjectTasksWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDevopsScenarioFieldConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDevopsScenarioFieldConfig"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListPipelinesWithOptions(request *ListPipelinesRequest, runtime *util.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPipelines"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectCustomFieldsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectCustomFields"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListServiceConnectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListServiceConnections"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListSmartGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSmartGroup"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserOrganizationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserOrganization"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferPipelineOwnerResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferPipelineOwner"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateCommonGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateCommonGroup"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevopsProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevopsProject"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevopsProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevopsProjectSprint"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDevopsProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDevopsProjectTask"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdatePipelineMemberWithOptions(request *UpdatePipelineMemberRequest, runtime *util.RuntimeOptions) (_result *UpdatePipelineMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePipelineMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePipelineMember"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateTaskDetail"), tea.String("2020-03-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
