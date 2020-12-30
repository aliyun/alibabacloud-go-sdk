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

type AddProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Members   *string `json:"Members,omitempty" xml:"Members,omitempty"`
}

func (s AddProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *AddProjectMembersRequest) SetOrgId(v string) *AddProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *AddProjectMembersRequest) SetProjectId(v string) *AddProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *AddProjectMembersRequest) SetMembers(v string) *AddProjectMembersRequest {
	s.Members = &v
	return s
}

type AddProjectMembersResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s AddProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddProjectMembersResponseBody) SetRequestId(v string) *AddProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddProjectMembersResponseBody) SetErrorMsg(v string) *AddProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AddProjectMembersResponseBody) SetObject(v bool) *AddProjectMembersResponseBody {
	s.Object = &v
	return s
}

func (s *AddProjectMembersResponseBody) SetErrorCode(v string) *AddProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddProjectMembersResponseBody) SetSuccessful(v bool) *AddProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type AddProjectMembersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *AddProjectMembersResponse) SetHeaders(v map[string]*string) *AddProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *AddProjectMembersResponse) SetBody(v *AddProjectMembersResponseBody) *AddProjectMembersResponse {
	s.Body = v
	return s
}

type ApplySmallMicroRequest struct {
	OrgId             *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	OrgName           *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	ApplicantName     *string `json:"ApplicantName,omitempty" xml:"ApplicantName,omitempty"`
	ApplicantTel      *string `json:"ApplicantTel,omitempty" xml:"ApplicantTel,omitempty"`
	ApplicantEmail    *string `json:"ApplicantEmail,omitempty" xml:"ApplicantEmail,omitempty"`
	ApplicantPosition *string `json:"ApplicantPosition,omitempty" xml:"ApplicantPosition,omitempty"`
	DevelopScale      *string `json:"DevelopScale,omitempty" xml:"DevelopScale,omitempty"`
	DevelopLanguage   *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	BusinessModel     *string `json:"BusinessModel,omitempty" xml:"BusinessModel,omitempty"`
	Solution          *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	ForHelp           *string `json:"ForHelp,omitempty" xml:"ForHelp,omitempty"`
}

func (s ApplySmallMicroRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplySmallMicroRequest) GoString() string {
	return s.String()
}

func (s *ApplySmallMicroRequest) SetOrgId(v string) *ApplySmallMicroRequest {
	s.OrgId = &v
	return s
}

func (s *ApplySmallMicroRequest) SetType(v string) *ApplySmallMicroRequest {
	s.Type = &v
	return s
}

func (s *ApplySmallMicroRequest) SetOrgName(v string) *ApplySmallMicroRequest {
	s.OrgName = &v
	return s
}

func (s *ApplySmallMicroRequest) SetApplicantName(v string) *ApplySmallMicroRequest {
	s.ApplicantName = &v
	return s
}

func (s *ApplySmallMicroRequest) SetApplicantTel(v string) *ApplySmallMicroRequest {
	s.ApplicantTel = &v
	return s
}

func (s *ApplySmallMicroRequest) SetApplicantEmail(v string) *ApplySmallMicroRequest {
	s.ApplicantEmail = &v
	return s
}

func (s *ApplySmallMicroRequest) SetApplicantPosition(v string) *ApplySmallMicroRequest {
	s.ApplicantPosition = &v
	return s
}

func (s *ApplySmallMicroRequest) SetDevelopScale(v string) *ApplySmallMicroRequest {
	s.DevelopScale = &v
	return s
}

func (s *ApplySmallMicroRequest) SetDevelopLanguage(v string) *ApplySmallMicroRequest {
	s.DevelopLanguage = &v
	return s
}

func (s *ApplySmallMicroRequest) SetBusinessModel(v string) *ApplySmallMicroRequest {
	s.BusinessModel = &v
	return s
}

func (s *ApplySmallMicroRequest) SetSolution(v string) *ApplySmallMicroRequest {
	s.Solution = &v
	return s
}

func (s *ApplySmallMicroRequest) SetForHelp(v string) *ApplySmallMicroRequest {
	s.ForHelp = &v
	return s
}

type ApplySmallMicroResponseBody struct {
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	Code      *int32                 `json:"code,omitempty" xml:"code,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Raw       *bool                  `json:"raw,omitempty" xml:"raw,omitempty"`
	Message   *bool                  `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ApplySmallMicroResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplySmallMicroResponseBody) GoString() string {
	return s.String()
}

func (s *ApplySmallMicroResponseBody) SetResult(v map[string]interface{}) *ApplySmallMicroResponseBody {
	s.Result = v
	return s
}

func (s *ApplySmallMicroResponseBody) SetCode(v int32) *ApplySmallMicroResponseBody {
	s.Code = &v
	return s
}

func (s *ApplySmallMicroResponseBody) SetRequestId(v string) *ApplySmallMicroResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplySmallMicroResponseBody) SetRaw(v bool) *ApplySmallMicroResponseBody {
	s.Raw = &v
	return s
}

func (s *ApplySmallMicroResponseBody) SetMessage(v bool) *ApplySmallMicroResponseBody {
	s.Message = &v
	return s
}

type ApplySmallMicroResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplySmallMicroResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplySmallMicroResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplySmallMicroResponse) GoString() string {
	return s.String()
}

func (s *ApplySmallMicroResponse) SetHeaders(v map[string]*string) *ApplySmallMicroResponse {
	s.Headers = v
	return s
}

func (s *ApplySmallMicroResponse) SetBody(v *ApplySmallMicroResponseBody) *ApplySmallMicroResponse {
	s.Body = v
	return s
}

type BactchInsertMembersRequest struct {
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Members *string `json:"Members,omitempty" xml:"Members,omitempty"`
	RealPk  *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s BactchInsertMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s BactchInsertMembersRequest) GoString() string {
	return s.String()
}

func (s *BactchInsertMembersRequest) SetOrgId(v string) *BactchInsertMembersRequest {
	s.OrgId = &v
	return s
}

func (s *BactchInsertMembersRequest) SetMembers(v string) *BactchInsertMembersRequest {
	s.Members = &v
	return s
}

func (s *BactchInsertMembersRequest) SetRealPk(v string) *BactchInsertMembersRequest {
	s.RealPk = &v
	return s
}

type BactchInsertMembersResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BactchInsertMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BactchInsertMembersResponseBody) GoString() string {
	return s.String()
}

func (s *BactchInsertMembersResponseBody) SetRequestId(v string) *BactchInsertMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *BactchInsertMembersResponseBody) SetObject(v bool) *BactchInsertMembersResponseBody {
	s.Object = &v
	return s
}

func (s *BactchInsertMembersResponseBody) SetErrorCode(v string) *BactchInsertMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BactchInsertMembersResponseBody) SetErrorMessage(v string) *BactchInsertMembersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BactchInsertMembersResponseBody) SetSuccess(v bool) *BactchInsertMembersResponseBody {
	s.Success = &v
	return s
}

type BactchInsertMembersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BactchInsertMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BactchInsertMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s BactchInsertMembersResponse) GoString() string {
	return s.String()
}

func (s *BactchInsertMembersResponse) SetHeaders(v map[string]*string) *BactchInsertMembersResponse {
	s.Headers = v
	return s
}

func (s *BactchInsertMembersResponse) SetBody(v *BactchInsertMembersResponseBody) *BactchInsertMembersResponse {
	s.Body = v
	return s
}

type CheckAliyunUserExistsRequest struct {
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s CheckAliyunUserExistsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunUserExistsRequest) GoString() string {
	return s.String()
}

func (s *CheckAliyunUserExistsRequest) SetUserPk(v string) *CheckAliyunUserExistsRequest {
	s.UserPk = &v
	return s
}

type CheckAliyunUserExistsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CheckAliyunUserExistsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunUserExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAliyunUserExistsResponseBody) SetRequestId(v string) *CheckAliyunUserExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckAliyunUserExistsResponseBody) SetErrorMsg(v string) *CheckAliyunUserExistsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CheckAliyunUserExistsResponseBody) SetObject(v bool) *CheckAliyunUserExistsResponseBody {
	s.Object = &v
	return s
}

func (s *CheckAliyunUserExistsResponseBody) SetErrorCode(v string) *CheckAliyunUserExistsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CheckAliyunUserExistsResponseBody) SetSuccessful(v bool) *CheckAliyunUserExistsResponseBody {
	s.Successful = &v
	return s
}

type CheckAliyunUserExistsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckAliyunUserExistsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAliyunUserExistsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAliyunUserExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckAliyunUserExistsResponse) SetHeaders(v map[string]*string) *CheckAliyunUserExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckAliyunUserExistsResponse) SetBody(v *CheckAliyunUserExistsResponseBody) *CheckAliyunUserExistsResponse {
	s.Body = v
	return s
}

type CreateDevopsOrgRequest struct {
	OrgName            *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	RealPk             *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
	DesiredMemberCount *int32  `json:"DesiredMemberCount,omitempty" xml:"DesiredMemberCount,omitempty"`
}

func (s CreateDevopsOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrgRequest) SetOrgName(v string) *CreateDevopsOrgRequest {
	s.OrgName = &v
	return s
}

func (s *CreateDevopsOrgRequest) SetSource(v string) *CreateDevopsOrgRequest {
	s.Source = &v
	return s
}

func (s *CreateDevopsOrgRequest) SetRealPk(v string) *CreateDevopsOrgRequest {
	s.RealPk = &v
	return s
}

func (s *CreateDevopsOrgRequest) SetDesiredMemberCount(v int32) *CreateDevopsOrgRequest {
	s.DesiredMemberCount = &v
	return s
}

type CreateDevopsOrgResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDevopsOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrgResponseBody) SetRequestId(v string) *CreateDevopsOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDevopsOrgResponseBody) SetObject(v string) *CreateDevopsOrgResponseBody {
	s.Object = &v
	return s
}

func (s *CreateDevopsOrgResponseBody) SetErrorCode(v string) *CreateDevopsOrgResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDevopsOrgResponseBody) SetErrorMessage(v string) *CreateDevopsOrgResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDevopsOrgResponseBody) SetSuccess(v bool) *CreateDevopsOrgResponseBody {
	s.Success = &v
	return s
}

type CreateDevopsOrgResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDevopsOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDevopsOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDevopsOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateDevopsOrgResponse) SetHeaders(v map[string]*string) *CreateDevopsOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateDevopsOrgResponse) SetBody(v *CreateDevopsOrgResponseBody) *CreateDevopsOrgResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetOrgId(v string) *CreateProjectRequest {
	s.OrgId = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

type CreateProjectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateProjectResponseBody) SetObject(v string) *CreateProjectResponseBody {
	s.Object = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrorCode(v string) *CreateProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrorMessage(v string) *CreateProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
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

type CreateProjectSprintRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
}

func (s CreateProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectSprintRequest) SetOrgId(v string) *CreateProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *CreateProjectSprintRequest) SetName(v string) *CreateProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectSprintRequest) SetDescription(v string) *CreateProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectSprintRequest) SetProjectId(v string) *CreateProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectSprintRequest) SetExecutorId(v string) *CreateProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateProjectSprintRequest) SetStartDate(v string) *CreateProjectSprintRequest {
	s.StartDate = &v
	return s
}

func (s *CreateProjectSprintRequest) SetDueDate(v string) *CreateProjectSprintRequest {
	s.DueDate = &v
	return s
}

type CreateProjectSprintResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateProjectSprintResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                  `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectSprintResponseBody) SetRequestId(v string) *CreateProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectSprintResponseBody) SetErrorMsg(v string) *CreateProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProjectSprintResponseBody) SetObject(v *CreateProjectSprintResponseBodyObject) *CreateProjectSprintResponseBody {
	s.Object = v
	return s
}

func (s *CreateProjectSprintResponseBody) SetErrorCode(v string) *CreateProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectSprintResponseBody) SetSuccessful(v bool) *CreateProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type CreateProjectSprintResponseBodyObject struct {
	Status       *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                        `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                        `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                        `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Executor     *string                                        `json:"Executor,omitempty" xml:"Executor,omitempty"`
	Description  *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Accomplished *string                                        `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                          `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                        `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                        `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                        `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *CreateProjectSprintResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                        `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateProjectSprintResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectSprintResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateProjectSprintResponseBodyObject) SetStatus(v string) *CreateProjectSprintResponseBodyObject {
	s.Status = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetProjectId(v string) *CreateProjectSprintResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetStartDate(v string) *CreateProjectSprintResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetCreatorId(v string) *CreateProjectSprintResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetExecutor(v string) *CreateProjectSprintResponseBodyObject {
	s.Executor = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetDescription(v string) *CreateProjectSprintResponseBodyObject {
	s.Description = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetAccomplished(v string) *CreateProjectSprintResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetIsDeleted(v bool) *CreateProjectSprintResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetUpdated(v string) *CreateProjectSprintResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetDueDate(v string) *CreateProjectSprintResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetName(v string) *CreateProjectSprintResponseBodyObject {
	s.Name = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetCreated(v string) *CreateProjectSprintResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetPlanToDo(v *CreateProjectSprintResponseBodyObjectPlanToDo) *CreateProjectSprintResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *CreateProjectSprintResponseBodyObject) SetId(v string) *CreateProjectSprintResponseBodyObject {
	s.Id = &v
	return s
}

type CreateProjectSprintResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s CreateProjectSprintResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectSprintResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *CreateProjectSprintResponseBodyObjectPlanToDo) SetTasks(v int32) *CreateProjectSprintResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *CreateProjectSprintResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *CreateProjectSprintResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *CreateProjectSprintResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type CreateProjectSprintResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectSprintResponse) SetHeaders(v map[string]*string) *CreateProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectSprintResponse) SetBody(v *CreateProjectSprintResponseBody) *CreateProjectSprintResponse {
	s.Body = v
	return s
}

type CreateProjectTaskRequest struct {
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

func (s CreateProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectTaskRequest) SetOrgId(v string) *CreateProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetContent(v string) *CreateProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *CreateProjectTaskRequest) SetProjectId(v string) *CreateProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetExecutorId(v string) *CreateProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetStartDate(v string) *CreateProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *CreateProjectTaskRequest) SetDueDate(v string) *CreateProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *CreateProjectTaskRequest) SetScenarioFieldConfigId(v string) *CreateProjectTaskRequest {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetTaskFlowStatusId(v string) *CreateProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetNote(v string) *CreateProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *CreateProjectTaskRequest) SetPriority(v int32) *CreateProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateProjectTaskRequest) SetVisible(v string) *CreateProjectTaskRequest {
	s.Visible = &v
	return s
}

func (s *CreateProjectTaskRequest) SetParentTaskId(v string) *CreateProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetSprintId(v string) *CreateProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *CreateProjectTaskRequest) SetTaskListId(v string) *CreateProjectTaskRequest {
	s.TaskListId = &v
	return s
}

type CreateProjectTaskResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                              `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *CreateProjectTaskResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s CreateProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectTaskResponseBody) SetRequestId(v string) *CreateProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectTaskResponseBody) SetErrorMsg(v string) *CreateProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProjectTaskResponseBody) SetObject(v *CreateProjectTaskResponseBodyObject) *CreateProjectTaskResponseBody {
	s.Object = v
	return s
}

func (s *CreateProjectTaskResponseBody) SetErrorCode(v string) *CreateProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectTaskResponseBody) SetSuccessful(v bool) *CreateProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type CreateProjectTaskResponseBodyObject struct {
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

func (s CreateProjectTaskResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectTaskResponseBodyObject) GoString() string {
	return s.String()
}

func (s *CreateProjectTaskResponseBodyObject) SetExecutorId(v string) *CreateProjectTaskResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetProjectId(v string) *CreateProjectTaskResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetPriority(v int32) *CreateProjectTaskResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetScenarioFieldConfigId(v string) *CreateProjectTaskResponseBodyObject {
	s.ScenarioFieldConfigId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetAncestorIds(v string) *CreateProjectTaskResponseBodyObject {
	s.AncestorIds = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetTaskType(v string) *CreateProjectTaskResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetTasklistId(v string) *CreateProjectTaskResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetTaskflowstatusId(v string) *CreateProjectTaskResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetNote(v string) *CreateProjectTaskResponseBodyObject {
	s.Note = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetUpdated(v string) *CreateProjectTaskResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetUniqueId(v int32) *CreateProjectTaskResponseBodyObject {
	s.UniqueId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetContent(v string) *CreateProjectTaskResponseBodyObject {
	s.Content = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetRating(v int32) *CreateProjectTaskResponseBodyObject {
	s.Rating = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetPos(v int32) *CreateProjectTaskResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetStartDate(v string) *CreateProjectTaskResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetStoryPoint(v string) *CreateProjectTaskResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetCreatorId(v string) *CreateProjectTaskResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetSource(v string) *CreateProjectTaskResponseBodyObject {
	s.Source = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetOrganizationId(v string) *CreateProjectTaskResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetVisible(v string) *CreateProjectTaskResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetIsDone(v bool) *CreateProjectTaskResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetSprintId(v string) *CreateProjectTaskResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetDueDate(v string) *CreateProjectTaskResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetCreated(v string) *CreateProjectTaskResponseBodyObject {
	s.Created = &v
	return s
}

func (s *CreateProjectTaskResponseBodyObject) SetId(v string) *CreateProjectTaskResponseBodyObject {
	s.Id = &v
	return s
}

type CreateProjectTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectTaskResponse) SetHeaders(v map[string]*string) *CreateProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectTaskResponse) SetBody(v *CreateProjectTaskResponseBody) *CreateProjectTaskResponse {
	s.Body = v
	return s
}

type DeleteMembersForOrgRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RealPk *string `json:"RealPk,omitempty" xml:"RealPk,omitempty"`
}

func (s DeleteMembersForOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersForOrgRequest) GoString() string {
	return s.String()
}

func (s *DeleteMembersForOrgRequest) SetOrgId(v string) *DeleteMembersForOrgRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteMembersForOrgRequest) SetUserId(v string) *DeleteMembersForOrgRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMembersForOrgRequest) SetRealPk(v string) *DeleteMembersForOrgRequest {
	s.RealPk = &v
	return s
}

type DeleteMembersForOrgResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMembersForOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersForOrgResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMembersForOrgResponseBody) SetRequestId(v string) *DeleteMembersForOrgResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMembersForOrgResponseBody) SetObject(v bool) *DeleteMembersForOrgResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteMembersForOrgResponseBody) SetErrorCode(v string) *DeleteMembersForOrgResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteMembersForOrgResponseBody) SetErrorMessage(v string) *DeleteMembersForOrgResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteMembersForOrgResponseBody) SetSuccess(v bool) *DeleteMembersForOrgResponseBody {
	s.Success = &v
	return s
}

type DeleteMembersForOrgResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMembersForOrgResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMembersForOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMembersForOrgResponse) GoString() string {
	return s.String()
}

func (s *DeleteMembersForOrgResponse) SetHeaders(v map[string]*string) *DeleteMembersForOrgResponse {
	s.Headers = v
	return s
}

func (s *DeleteMembersForOrgResponse) SetBody(v *DeleteMembersForOrgResponseBody) *DeleteMembersForOrgResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetOrgId(v string) *DeleteProjectRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteProjectResponseBody) SetObject(v string) *DeleteProjectResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteProjectResponseBody) SetErrorCode(v string) *DeleteProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectResponseBody) SetErrorMessage(v string) *DeleteProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
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

type DeleteProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserIds   *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s DeleteProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersRequest) SetOrgId(v string) *DeleteProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteProjectMembersRequest) SetProjectId(v string) *DeleteProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteProjectMembersRequest) SetUserIds(v string) *DeleteProjectMembersRequest {
	s.UserIds = &v
	return s
}

type DeleteProjectMembersResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersResponseBody) SetRequestId(v string) *DeleteProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectMembersResponseBody) SetErrorMsg(v string) *DeleteProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteProjectMembersResponseBody) SetObject(v bool) *DeleteProjectMembersResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteProjectMembersResponseBody) SetErrorCode(v string) *DeleteProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectMembersResponseBody) SetSuccessful(v bool) *DeleteProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type DeleteProjectMembersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectMembersResponse) SetHeaders(v map[string]*string) *DeleteProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectMembersResponse) SetBody(v *DeleteProjectMembersResponseBody) *DeleteProjectMembersResponse {
	s.Body = v
	return s
}

type DeleteProjectSprintRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	SprintId *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s DeleteProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectSprintRequest) SetOrgId(v string) *DeleteProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteProjectSprintRequest) SetSprintId(v string) *DeleteProjectSprintRequest {
	s.SprintId = &v
	return s
}

type DeleteProjectSprintResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectSprintResponseBody) SetRequestId(v string) *DeleteProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectSprintResponseBody) SetErrorMsg(v string) *DeleteProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteProjectSprintResponseBody) SetObject(v bool) *DeleteProjectSprintResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteProjectSprintResponseBody) SetErrorCode(v string) *DeleteProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectSprintResponseBody) SetSuccessful(v bool) *DeleteProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type DeleteProjectSprintResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectSprintResponse) SetHeaders(v map[string]*string) *DeleteProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectSprintResponse) SetBody(v *DeleteProjectSprintResponseBody) *DeleteProjectSprintResponse {
	s.Body = v
	return s
}

type DeleteProjectTaskRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectTaskRequest) SetOrgId(v string) *DeleteProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *DeleteProjectTaskRequest) SetTaskId(v string) *DeleteProjectTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteProjectTaskResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s DeleteProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectTaskResponseBody) SetRequestId(v string) *DeleteProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectTaskResponseBody) SetErrorMsg(v string) *DeleteProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteProjectTaskResponseBody) SetObject(v bool) *DeleteProjectTaskResponseBody {
	s.Object = &v
	return s
}

func (s *DeleteProjectTaskResponseBody) SetErrorCode(v string) *DeleteProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteProjectTaskResponseBody) SetSuccessful(v bool) *DeleteProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type DeleteProjectTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectTaskResponse) SetHeaders(v map[string]*string) *DeleteProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectTaskResponse) SetBody(v *DeleteProjectTaskResponseBody) *DeleteProjectTaskResponse {
	s.Body = v
	return s
}

type GetOrganizationMembersRequest struct {
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s GetOrganizationMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationMembersRequest) SetOrgId(v string) *GetOrganizationMembersRequest {
	s.OrgId = &v
	return s
}

type GetOrganizationMembersResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                     `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetOrganizationMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                     `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                       `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetOrganizationMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationMembersResponseBody) SetRequestId(v string) *GetOrganizationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrganizationMembersResponseBody) SetErrorMsg(v string) *GetOrganizationMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetOrganizationMembersResponseBody) SetObject(v []*GetOrganizationMembersResponseBodyObject) *GetOrganizationMembersResponseBody {
	s.Object = v
	return s
}

func (s *GetOrganizationMembersResponseBody) SetErrorCode(v string) *GetOrganizationMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrganizationMembersResponseBody) SetSuccessful(v bool) *GetOrganizationMembersResponseBody {
	s.Successful = &v
	return s
}

type GetOrganizationMembersResponseBodyObject struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetOrganizationMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetOrganizationMembersResponseBodyObject) SetEmail(v string) *GetOrganizationMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetAvatarUrl(v string) *GetOrganizationMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetUserId(v string) *GetOrganizationMembersResponseBodyObject {
	s.UserId = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetMemberId(v string) *GetOrganizationMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetRole(v int32) *GetOrganizationMembersResponseBodyObject {
	s.Role = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetName(v string) *GetOrganizationMembersResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetOrganizationMembersResponseBodyObject) SetPhone(v string) *GetOrganizationMembersResponseBodyObject {
	s.Phone = &v
	return s
}

type GetOrganizationMembersResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOrganizationMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOrganizationMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationMembersResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationMembersResponse) SetHeaders(v map[string]*string) *GetOrganizationMembersResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationMembersResponse) SetBody(v *GetOrganizationMembersResponseBody) *GetOrganizationMembersResponse {
	s.Body = v
	return s
}

type GetProjectInfoRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoRequest) GoString() string {
	return s.String()
}

func (s *GetProjectInfoRequest) SetOrgId(v string) *GetProjectInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetProjectInfoRequest) SetProjectId(v string) *GetProjectInfoRequest {
	s.ProjectId = &v
	return s
}

type GetProjectInfoResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                           `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetProjectInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                             `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBody) SetRequestId(v string) *GetProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetErrorMsg(v string) *GetProjectInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetObject(v *GetProjectInfoResponseBodyObject) *GetProjectInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetProjectInfoResponseBody) SetErrorCode(v string) *GetProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetSuccessful(v bool) *GetProjectInfoResponseBody {
	s.Successful = &v
	return s
}

type GetProjectInfoResponseBodyObject struct {
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

func (s GetProjectInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBodyObject) SetSortMethod(v string) *GetProjectInfoResponseBodyObject {
	s.SortMethod = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetUniqueIdPrefix(v string) *GetProjectInfoResponseBodyObject {
	s.UniqueIdPrefix = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetNormalType(v string) *GetProjectInfoResponseBodyObject {
	s.NormalType = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetModifierId(v string) *GetProjectInfoResponseBodyObject {
	s.ModifierId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetSourceType(v string) *GetProjectInfoResponseBodyObject {
	s.SourceType = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetIsTemplate(v bool) *GetProjectInfoResponseBodyObject {
	s.IsTemplate = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetDescription(v string) *GetProjectInfoResponseBodyObject {
	s.Description = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetDefaultRoleId(v string) *GetProjectInfoResponseBodyObject {
	s.DefaultRoleId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetRootCollectionId(v string) *GetProjectInfoResponseBodyObject {
	s.RootCollectionId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetIsDeleted(v bool) *GetProjectInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetUpdated(v string) *GetProjectInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetIsArchived(v bool) *GetProjectInfoResponseBodyObject {
	s.IsArchived = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetName(v string) *GetProjectInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetEndDate(v string) *GetProjectInfoResponseBodyObject {
	s.EndDate = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetLogo(v string) *GetProjectInfoResponseBodyObject {
	s.Logo = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetStartDate(v string) *GetProjectInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetPinyin(v string) *GetProjectInfoResponseBodyObject {
	s.Pinyin = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetCreatorId(v string) *GetProjectInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetSourceId(v string) *GetProjectInfoResponseBodyObject {
	s.SourceId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetOrganizationId(v string) *GetProjectInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetIsSuspended(v bool) *GetProjectInfoResponseBodyObject {
	s.IsSuspended = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetDefaultCollectionId(v string) *GetProjectInfoResponseBodyObject {
	s.DefaultCollectionId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetVisibility(v string) *GetProjectInfoResponseBodyObject {
	s.Visibility = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetPy(v string) *GetProjectInfoResponseBodyObject {
	s.Py = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetCategory(v string) *GetProjectInfoResponseBodyObject {
	s.Category = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetNextTaskUniqueId(v int32) *GetProjectInfoResponseBodyObject {
	s.NextTaskUniqueId = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetCustomfields(v string) *GetProjectInfoResponseBodyObject {
	s.Customfields = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetCreated(v string) *GetProjectInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetProjectInfoResponseBodyObject) SetId(v string) *GetProjectInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetProjectInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponse) SetHeaders(v map[string]*string) *GetProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *GetProjectInfoResponse) SetBody(v *GetProjectInfoResponseBody) *GetProjectInfoResponse {
	s.Body = v
	return s
}

type GetProjectMembersRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMembersRequest) SetOrgId(v string) *GetProjectMembersRequest {
	s.OrgId = &v
	return s
}

func (s *GetProjectMembersRequest) SetProjectId(v string) *GetProjectMembersRequest {
	s.ProjectId = &v
	return s
}

type GetProjectMembersResponseBody struct {
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*GetProjectMembersResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                  `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponseBody) SetRequestId(v string) *GetProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetErrorMsg(v string) *GetProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetObject(v []*GetProjectMembersResponseBodyObject) *GetProjectMembersResponseBody {
	s.Object = v
	return s
}

func (s *GetProjectMembersResponseBody) SetErrorCode(v string) *GetProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetSuccessful(v bool) *GetProjectMembersResponseBody {
	s.Successful = &v
	return s
}

type GetProjectMembersResponseBodyObject struct {
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	MemberId  *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Role      *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetProjectMembersResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponseBodyObject) SetEmail(v string) *GetProjectMembersResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetAvatarUrl(v string) *GetProjectMembersResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetUserId(v string) *GetProjectMembersResponseBodyObject {
	s.UserId = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetMemberId(v string) *GetProjectMembersResponseBodyObject {
	s.MemberId = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetRole(v int32) *GetProjectMembersResponseBodyObject {
	s.Role = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetName(v string) *GetProjectMembersResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetProjectMembersResponseBodyObject) SetPhone(v string) *GetProjectMembersResponseBodyObject {
	s.Phone = &v
	return s
}

type GetProjectMembersResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponse) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponse) SetHeaders(v map[string]*string) *GetProjectMembersResponse {
	s.Headers = v
	return s
}

func (s *GetProjectMembersResponse) SetBody(v *GetProjectMembersResponseBody) *GetProjectMembersResponse {
	s.Body = v
	return s
}

type GetProjectSprintInfoRequest struct {
	OrgId    *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	SprintId *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s GetProjectSprintInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectSprintInfoRequest) GoString() string {
	return s.String()
}

func (s *GetProjectSprintInfoRequest) SetOrgId(v string) *GetProjectSprintInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetProjectSprintInfoRequest) SetSprintId(v string) *GetProjectSprintInfoRequest {
	s.SprintId = &v
	return s
}

type GetProjectSprintInfoResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetProjectSprintInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectSprintInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectSprintInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectSprintInfoResponseBody) SetRequestId(v string) *GetProjectSprintInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectSprintInfoResponseBody) SetErrorMsg(v string) *GetProjectSprintInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectSprintInfoResponseBody) SetObject(v *GetProjectSprintInfoResponseBodyObject) *GetProjectSprintInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetProjectSprintInfoResponseBody) SetErrorCode(v string) *GetProjectSprintInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectSprintInfoResponseBody) SetSuccessful(v bool) *GetProjectSprintInfoResponseBody {
	s.Successful = &v
	return s
}

type GetProjectSprintInfoResponseBodyObject struct {
	Status       *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                         `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                         `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Accomplished *string                                         `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                           `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                         `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                         `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                         `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *GetProjectSprintInfoResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                         `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectSprintInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectSprintInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectSprintInfoResponseBodyObject) SetStatus(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Status = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetProjectId(v string) *GetProjectSprintInfoResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetStartDate(v string) *GetProjectSprintInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetCreatorId(v string) *GetProjectSprintInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetAccomplished(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetIsDeleted(v bool) *GetProjectSprintInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetUpdated(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetDueDate(v string) *GetProjectSprintInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetName(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetCreated(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetPlanToDo(v *GetProjectSprintInfoResponseBodyObjectPlanToDo) *GetProjectSprintInfoResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObject) SetId(v string) *GetProjectSprintInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetProjectSprintInfoResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s GetProjectSprintInfoResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s GetProjectSprintInfoResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *GetProjectSprintInfoResponseBodyObjectPlanToDo) SetTasks(v int32) *GetProjectSprintInfoResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *GetProjectSprintInfoResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *GetProjectSprintInfoResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *GetProjectSprintInfoResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type GetProjectSprintInfoResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectSprintInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectSprintInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectSprintInfoResponse) GoString() string {
	return s.String()
}

func (s *GetProjectSprintInfoResponse) SetHeaders(v map[string]*string) *GetProjectSprintInfoResponse {
	s.Headers = v
	return s
}

func (s *GetProjectSprintInfoResponse) SetBody(v *GetProjectSprintInfoResponseBody) *GetProjectSprintInfoResponse {
	s.Body = v
	return s
}

type GetProjectTaskInfoRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetProjectTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetProjectTaskInfoRequest) SetOrgId(v string) *GetProjectTaskInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetProjectTaskInfoRequest) SetTaskId(v string) *GetProjectTaskInfoRequest {
	s.TaskId = &v
	return s
}

type GetProjectTaskInfoResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetProjectTaskInfoResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetProjectTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectTaskInfoResponseBody) SetRequestId(v string) *GetProjectTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBody) SetErrorMsg(v string) *GetProjectTaskInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetProjectTaskInfoResponseBody) SetObject(v *GetProjectTaskInfoResponseBodyObject) *GetProjectTaskInfoResponseBody {
	s.Object = v
	return s
}

func (s *GetProjectTaskInfoResponseBody) SetErrorCode(v string) *GetProjectTaskInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectTaskInfoResponseBody) SetSuccessful(v bool) *GetProjectTaskInfoResponseBody {
	s.Successful = &v
	return s
}

type GetProjectTaskInfoResponseBodyObject struct {
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

func (s GetProjectTaskInfoResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskInfoResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetProjectTaskInfoResponseBodyObject) SetExecutorId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.ExecutorId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetProjectId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetStoryPoint(v string) *GetProjectTaskInfoResponseBodyObject {
	s.StoryPoint = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetStartDate(v string) *GetProjectTaskInfoResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetIsTopInProject(v bool) *GetProjectTaskInfoResponseBodyObject {
	s.IsTopInProject = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetPriority(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Priority = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetCreatorId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetOrganizationId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.OrganizationId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetTaskType(v string) *GetProjectTaskInfoResponseBodyObject {
	s.TaskType = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetTasklistId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetVisible(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Visible = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetIsDone(v bool) *GetProjectTaskInfoResponseBodyObject {
	s.IsDone = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetIsDeleted(v bool) *GetProjectTaskInfoResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetTaskflowstatusId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.TaskflowstatusId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetNote(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Note = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetSprintId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.SprintId = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetUpdated(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetInvolveMembers(v []*string) *GetProjectTaskInfoResponseBodyObject {
	s.InvolveMembers = v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetDueDate(v string) *GetProjectTaskInfoResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetCreated(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Created = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetContent(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Content = &v
	return s
}

func (s *GetProjectTaskInfoResponseBodyObject) SetId(v string) *GetProjectTaskInfoResponseBodyObject {
	s.Id = &v
	return s
}

type GetProjectTaskInfoResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetProjectTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetProjectTaskInfoResponse) SetHeaders(v map[string]*string) *GetProjectTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetProjectTaskInfoResponse) SetBody(v *GetProjectTaskInfoResponseBody) *GetProjectTaskInfoResponse {
	s.Body = v
	return s
}

type GetUserByUidRequest struct {
	OrgId  *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	UserPk *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
}

func (s GetUserByUidRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUidRequest) GoString() string {
	return s.String()
}

func (s *GetUserByUidRequest) SetOrgId(v string) *GetUserByUidRequest {
	s.OrgId = &v
	return s
}

func (s *GetUserByUidRequest) SetUserPk(v string) *GetUserByUidRequest {
	s.UserPk = &v
	return s
}

type GetUserByUidResponseBody struct {
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                         `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *GetUserByUidResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Struct"`
	ErrorCode  *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                           `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s GetUserByUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUidResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByUidResponseBody) SetRequestId(v string) *GetUserByUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByUidResponseBody) SetErrorMsg(v string) *GetUserByUidResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetUserByUidResponseBody) SetObject(v *GetUserByUidResponseBodyObject) *GetUserByUidResponseBody {
	s.Object = v
	return s
}

func (s *GetUserByUidResponseBody) SetErrorCode(v string) *GetUserByUidResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetUserByUidResponseBody) SetSuccessful(v bool) *GetUserByUidResponseBody {
	s.Successful = &v
	return s
}

type GetUserByUidResponseBodyObject struct {
	AliyunPk  *string `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	Email     *string `json:"Email,omitempty" xml:"Email,omitempty"`
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Phone     *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetUserByUidResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUidResponseBodyObject) GoString() string {
	return s.String()
}

func (s *GetUserByUidResponseBodyObject) SetAliyunPk(v string) *GetUserByUidResponseBodyObject {
	s.AliyunPk = &v
	return s
}

func (s *GetUserByUidResponseBodyObject) SetEmail(v string) *GetUserByUidResponseBodyObject {
	s.Email = &v
	return s
}

func (s *GetUserByUidResponseBodyObject) SetAvatarUrl(v string) *GetUserByUidResponseBodyObject {
	s.AvatarUrl = &v
	return s
}

func (s *GetUserByUidResponseBodyObject) SetName(v string) *GetUserByUidResponseBodyObject {
	s.Name = &v
	return s
}

func (s *GetUserByUidResponseBodyObject) SetId(v string) *GetUserByUidResponseBodyObject {
	s.Id = &v
	return s
}

func (s *GetUserByUidResponseBodyObject) SetPhone(v string) *GetUserByUidResponseBodyObject {
	s.Phone = &v
	return s
}

type GetUserByUidResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserByUidResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByUidResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByUidResponse) GoString() string {
	return s.String()
}

func (s *GetUserByUidResponse) SetHeaders(v map[string]*string) *GetUserByUidResponse {
	s.Headers = v
	return s
}

func (s *GetUserByUidResponse) SetBody(v *GetUserByUidResponseBody) *GetUserByUidResponse {
	s.Body = v
	return s
}

type InsertDevopsMemberRequest struct {
	UserPk   *string `json:"UserPk,omitempty" xml:"UserPk,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Phone    *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
}

func (s InsertDevopsMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsMemberRequest) GoString() string {
	return s.String()
}

func (s *InsertDevopsMemberRequest) SetUserPk(v string) *InsertDevopsMemberRequest {
	s.UserPk = &v
	return s
}

func (s *InsertDevopsMemberRequest) SetUserName(v string) *InsertDevopsMemberRequest {
	s.UserName = &v
	return s
}

func (s *InsertDevopsMemberRequest) SetPhone(v string) *InsertDevopsMemberRequest {
	s.Phone = &v
	return s
}

func (s *InsertDevopsMemberRequest) SetEmail(v string) *InsertDevopsMemberRequest {
	s.Email = &v
	return s
}

type InsertDevopsMemberResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertDevopsMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsMemberResponseBody) GoString() string {
	return s.String()
}

func (s *InsertDevopsMemberResponseBody) SetRequestId(v string) *InsertDevopsMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertDevopsMemberResponseBody) SetObject(v string) *InsertDevopsMemberResponseBody {
	s.Object = &v
	return s
}

func (s *InsertDevopsMemberResponseBody) SetErrorCode(v string) *InsertDevopsMemberResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InsertDevopsMemberResponseBody) SetErrorMessage(v string) *InsertDevopsMemberResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InsertDevopsMemberResponseBody) SetSuccess(v bool) *InsertDevopsMemberResponseBody {
	s.Success = &v
	return s
}

type InsertDevopsMemberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InsertDevopsMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InsertDevopsMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertDevopsMemberResponse) GoString() string {
	return s.String()
}

func (s *InsertDevopsMemberResponse) SetHeaders(v map[string]*string) *InsertDevopsMemberResponse {
	s.Headers = v
	return s
}

func (s *InsertDevopsMemberResponse) SetBody(v *InsertDevopsMemberResponseBody) *InsertDevopsMemberResponse {
	s.Body = v
	return s
}

type ListProjectSprintsRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectSprintsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectSprintsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectSprintsRequest) SetOrgId(v string) *ListProjectSprintsRequest {
	s.OrgId = &v
	return s
}

func (s *ListProjectSprintsRequest) SetProjectId(v string) *ListProjectSprintsRequest {
	s.ProjectId = &v
	return s
}

type ListProjectSprintsResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                 `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectSprintsResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectSprintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectSprintsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectSprintsResponseBody) SetRequestId(v string) *ListProjectSprintsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectSprintsResponseBody) SetErrorMsg(v string) *ListProjectSprintsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectSprintsResponseBody) SetObject(v []*ListProjectSprintsResponseBodyObject) *ListProjectSprintsResponseBody {
	s.Object = v
	return s
}

func (s *ListProjectSprintsResponseBody) SetErrorCode(v string) *ListProjectSprintsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectSprintsResponseBody) SetSuccessful(v bool) *ListProjectSprintsResponseBody {
	s.Successful = &v
	return s
}

type ListProjectSprintsResponseBodyObject struct {
	Status       *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId    *string                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartDate    *string                                       `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	CreatorId    *string                                       `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Accomplished *string                                       `json:"Accomplished,omitempty" xml:"Accomplished,omitempty"`
	IsDeleted    *bool                                         `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	Updated      *string                                       `json:"Updated,omitempty" xml:"Updated,omitempty"`
	DueDate      *string                                       `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	Name         *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	Created      *string                                       `json:"Created,omitempty" xml:"Created,omitempty"`
	PlanToDo     *ListProjectSprintsResponseBodyObjectPlanToDo `json:"PlanToDo,omitempty" xml:"PlanToDo,omitempty" type:"Struct"`
	Id           *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListProjectSprintsResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectSprintsResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectSprintsResponseBodyObject) SetStatus(v string) *ListProjectSprintsResponseBodyObject {
	s.Status = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetProjectId(v string) *ListProjectSprintsResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetStartDate(v string) *ListProjectSprintsResponseBodyObject {
	s.StartDate = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetCreatorId(v string) *ListProjectSprintsResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetAccomplished(v string) *ListProjectSprintsResponseBodyObject {
	s.Accomplished = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetIsDeleted(v bool) *ListProjectSprintsResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetUpdated(v string) *ListProjectSprintsResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetDueDate(v string) *ListProjectSprintsResponseBodyObject {
	s.DueDate = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetName(v string) *ListProjectSprintsResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetCreated(v string) *ListProjectSprintsResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetPlanToDo(v *ListProjectSprintsResponseBodyObjectPlanToDo) *ListProjectSprintsResponseBodyObject {
	s.PlanToDo = v
	return s
}

func (s *ListProjectSprintsResponseBodyObject) SetId(v string) *ListProjectSprintsResponseBodyObject {
	s.Id = &v
	return s
}

type ListProjectSprintsResponseBodyObjectPlanToDo struct {
	Tasks       *int32 `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
	WorkTimes   *int32 `json:"WorkTimes,omitempty" xml:"WorkTimes,omitempty"`
	StoryPoints *int32 `json:"StoryPoints,omitempty" xml:"StoryPoints,omitempty"`
}

func (s ListProjectSprintsResponseBodyObjectPlanToDo) String() string {
	return tea.Prettify(s)
}

func (s ListProjectSprintsResponseBodyObjectPlanToDo) GoString() string {
	return s.String()
}

func (s *ListProjectSprintsResponseBodyObjectPlanToDo) SetTasks(v int32) *ListProjectSprintsResponseBodyObjectPlanToDo {
	s.Tasks = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObjectPlanToDo) SetWorkTimes(v int32) *ListProjectSprintsResponseBodyObjectPlanToDo {
	s.WorkTimes = &v
	return s
}

func (s *ListProjectSprintsResponseBodyObjectPlanToDo) SetStoryPoints(v int32) *ListProjectSprintsResponseBodyObjectPlanToDo {
	s.StoryPoints = &v
	return s
}

type ListProjectSprintsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectSprintsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectSprintsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectSprintsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectSprintsResponse) SetHeaders(v map[string]*string) *ListProjectSprintsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectSprintsResponse) SetBody(v *ListProjectSprintsResponseBody) *ListProjectSprintsResponse {
	s.Body = v
	return s
}

type ListProjectTaskFlowRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectTaskFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowRequest) SetOrgId(v string) *ListProjectTaskFlowRequest {
	s.OrgId = &v
	return s
}

func (s *ListProjectTaskFlowRequest) SetProjectId(v string) *ListProjectTaskFlowRequest {
	s.ProjectId = &v
	return s
}

type ListProjectTaskFlowResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                  `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectTaskFlowResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                    `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectTaskFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowResponseBody) SetRequestId(v string) *ListProjectTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectTaskFlowResponseBody) SetErrorMsg(v string) *ListProjectTaskFlowResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectTaskFlowResponseBody) SetObject(v []*ListProjectTaskFlowResponseBodyObject) *ListProjectTaskFlowResponseBody {
	s.Object = v
	return s
}

func (s *ListProjectTaskFlowResponseBody) SetErrorCode(v string) *ListProjectTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectTaskFlowResponseBody) SetSuccessful(v bool) *ListProjectTaskFlowResponseBody {
	s.Successful = &v
	return s
}

type ListProjectTaskFlowResponseBodyObject struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListProjectTaskFlowResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowResponseBodyObject) SetType(v string) *ListProjectTaskFlowResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListProjectTaskFlowResponseBodyObject) SetId(v string) *ListProjectTaskFlowResponseBodyObject {
	s.Id = &v
	return s
}

type ListProjectTaskFlowResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectTaskFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowResponse) SetHeaders(v map[string]*string) *ListProjectTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *ListProjectTaskFlowResponse) SetBody(v *ListProjectTaskFlowResponseBody) *ListProjectTaskFlowResponse {
	s.Body = v
	return s
}

type ListProjectTaskFlowStatusRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	TaskFlowId *string `json:"TaskFlowId,omitempty" xml:"TaskFlowId,omitempty"`
}

func (s ListProjectTaskFlowStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowStatusRequest) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowStatusRequest) SetOrgId(v string) *ListProjectTaskFlowStatusRequest {
	s.OrgId = &v
	return s
}

func (s *ListProjectTaskFlowStatusRequest) SetTaskFlowId(v string) *ListProjectTaskFlowStatusRequest {
	s.TaskFlowId = &v
	return s
}

type ListProjectTaskFlowStatusResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                        `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectTaskFlowStatusResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                          `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectTaskFlowStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowStatusResponseBody) SetRequestId(v string) *ListProjectTaskFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBody) SetErrorMsg(v string) *ListProjectTaskFlowStatusResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBody) SetObject(v []*ListProjectTaskFlowStatusResponseBodyObject) *ListProjectTaskFlowStatusResponseBody {
	s.Object = v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBody) SetErrorCode(v string) *ListProjectTaskFlowStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBody) SetSuccessful(v bool) *ListProjectTaskFlowStatusResponseBody {
	s.Successful = &v
	return s
}

type ListProjectTaskFlowStatusResponseBodyObject struct {
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

func (s ListProjectTaskFlowStatusResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowStatusResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetTaskflowId(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.TaskflowId = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetKind(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Kind = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetPos(v int32) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Pos = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetIsDeleted(v bool) *ListProjectTaskFlowStatusResponseBodyObject {
	s.IsDeleted = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetUpdated(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetCreatorId(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetName(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetCreated(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetRejectStatusIds(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.RejectStatusIds = &v
	return s
}

func (s *ListProjectTaskFlowStatusResponseBodyObject) SetId(v string) *ListProjectTaskFlowStatusResponseBodyObject {
	s.Id = &v
	return s
}

type ListProjectTaskFlowStatusResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectTaskFlowStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectTaskFlowStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTaskFlowStatusResponse) GoString() string {
	return s.String()
}

func (s *ListProjectTaskFlowStatusResponse) SetHeaders(v map[string]*string) *ListProjectTaskFlowStatusResponse {
	s.Headers = v
	return s
}

func (s *ListProjectTaskFlowStatusResponse) SetBody(v *ListProjectTaskFlowStatusResponseBody) *ListProjectTaskFlowStatusResponse {
	s.Body = v
	return s
}

type ListProjectTasksRequest struct {
	OrgId      *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
}

func (s ListProjectTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTasksRequest) GoString() string {
	return s.String()
}

func (s *ListProjectTasksRequest) SetOrgId(v string) *ListProjectTasksRequest {
	s.OrgId = &v
	return s
}

func (s *ListProjectTasksRequest) SetProjectIds(v string) *ListProjectTasksRequest {
	s.ProjectIds = &v
	return s
}

type ListProjectTasksResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListProjectTasksResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                 `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListProjectTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectTasksResponseBody) SetRequestId(v string) *ListProjectTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectTasksResponseBody) SetErrorMsg(v string) *ListProjectTasksResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectTasksResponseBody) SetObject(v []*ListProjectTasksResponseBodyObject) *ListProjectTasksResponseBody {
	s.Object = v
	return s
}

func (s *ListProjectTasksResponseBody) SetErrorCode(v string) *ListProjectTasksResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectTasksResponseBody) SetSuccessful(v bool) *ListProjectTasksResponseBody {
	s.Successful = &v
	return s
}

type ListProjectTasksResponseBodyObject struct {
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

func (s ListProjectTasksResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTasksResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListProjectTasksResponseBodyObject) SetTaskgroupId(v string) *ListProjectTasksResponseBodyObject {
	s.TaskgroupId = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetTasklistId(v string) *ListProjectTasksResponseBodyObject {
	s.TasklistId = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetProjectId(v string) *ListProjectTasksResponseBodyObject {
	s.ProjectId = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetModifierId(v string) *ListProjectTasksResponseBodyObject {
	s.ModifierId = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetUpdated(v string) *ListProjectTasksResponseBodyObject {
	s.Updated = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetCreatorId(v string) *ListProjectTasksResponseBodyObject {
	s.CreatorId = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetCreated(v string) *ListProjectTasksResponseBodyObject {
	s.Created = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetName(v string) *ListProjectTasksResponseBodyObject {
	s.Name = &v
	return s
}

func (s *ListProjectTasksResponseBodyObject) SetId(v string) *ListProjectTasksResponseBodyObject {
	s.Id = &v
	return s
}

type ListProjectTasksResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListProjectTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectTasksResponse) GoString() string {
	return s.String()
}

func (s *ListProjectTasksResponse) SetHeaders(v map[string]*string) *ListProjectTasksResponse {
	s.Headers = v
	return s
}

func (s *ListProjectTasksResponse) SetBody(v *ListProjectTasksResponseBody) *ListProjectTasksResponse {
	s.Body = v
	return s
}

type ListScenarioFieldConfigRequest struct {
	OrgId     *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListScenarioFieldConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioFieldConfigRequest) GoString() string {
	return s.String()
}

func (s *ListScenarioFieldConfigRequest) SetOrgId(v string) *ListScenarioFieldConfigRequest {
	s.OrgId = &v
	return s
}

func (s *ListScenarioFieldConfigRequest) SetProjectId(v string) *ListScenarioFieldConfigRequest {
	s.ProjectId = &v
	return s
}

type ListScenarioFieldConfigResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string                                      `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     []*ListScenarioFieldConfigResponseBodyObject `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	ErrorCode  *string                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool                                        `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s ListScenarioFieldConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioFieldConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListScenarioFieldConfigResponseBody) SetRequestId(v string) *ListScenarioFieldConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScenarioFieldConfigResponseBody) SetErrorMsg(v string) *ListScenarioFieldConfigResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListScenarioFieldConfigResponseBody) SetObject(v []*ListScenarioFieldConfigResponseBodyObject) *ListScenarioFieldConfigResponseBody {
	s.Object = v
	return s
}

func (s *ListScenarioFieldConfigResponseBody) SetErrorCode(v string) *ListScenarioFieldConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListScenarioFieldConfigResponseBody) SetSuccessful(v bool) *ListScenarioFieldConfigResponseBody {
	s.Successful = &v
	return s
}

type ListScenarioFieldConfigResponseBodyObject struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListScenarioFieldConfigResponseBodyObject) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioFieldConfigResponseBodyObject) GoString() string {
	return s.String()
}

func (s *ListScenarioFieldConfigResponseBodyObject) SetType(v string) *ListScenarioFieldConfigResponseBodyObject {
	s.Type = &v
	return s
}

func (s *ListScenarioFieldConfigResponseBodyObject) SetId(v string) *ListScenarioFieldConfigResponseBodyObject {
	s.Id = &v
	return s
}

type ListScenarioFieldConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScenarioFieldConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScenarioFieldConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScenarioFieldConfigResponse) GoString() string {
	return s.String()
}

func (s *ListScenarioFieldConfigResponse) SetHeaders(v map[string]*string) *ListScenarioFieldConfigResponse {
	s.Headers = v
	return s
}

func (s *ListScenarioFieldConfigResponse) SetBody(v *ListScenarioFieldConfigResponseBody) *ListScenarioFieldConfigResponse {
	s.Body = v
	return s
}

type UpdateProjectRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectRequest) SetOrgId(v string) *UpdateProjectRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateProjectRequest) SetName(v string) *UpdateProjectRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectRequest) SetDescription(v string) *UpdateProjectRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectRequest) SetProjectId(v string) *UpdateProjectRequest {
	s.ProjectId = &v
	return s
}

type UpdateProjectResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Object       *string `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *UpdateProjectResponseBody) SetObject(v string) *UpdateProjectResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateProjectResponseBody) SetErrorCode(v string) *UpdateProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetErrorMessage(v string) *UpdateProjectResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

type UpdateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateProjectResponse) SetBody(v *UpdateProjectResponseBody) *UpdateProjectResponse {
	s.Body = v
	return s
}

type UpdateProjectSprintRequest struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ExecutorId  *string `json:"ExecutorId,omitempty" xml:"ExecutorId,omitempty"`
	StartDate   *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	DueDate     *string `json:"DueDate,omitempty" xml:"DueDate,omitempty"`
	SprintId    *string `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s UpdateProjectSprintRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectSprintRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectSprintRequest) SetOrgId(v string) *UpdateProjectSprintRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetName(v string) *UpdateProjectSprintRequest {
	s.Name = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetDescription(v string) *UpdateProjectSprintRequest {
	s.Description = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetProjectId(v string) *UpdateProjectSprintRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetExecutorId(v string) *UpdateProjectSprintRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetStartDate(v string) *UpdateProjectSprintRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetDueDate(v string) *UpdateProjectSprintRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateProjectSprintRequest) SetSprintId(v string) *UpdateProjectSprintRequest {
	s.SprintId = &v
	return s
}

type UpdateProjectSprintResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateProjectSprintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectSprintResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectSprintResponseBody) SetRequestId(v string) *UpdateProjectSprintResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectSprintResponseBody) SetErrorMsg(v string) *UpdateProjectSprintResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectSprintResponseBody) SetObject(v bool) *UpdateProjectSprintResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateProjectSprintResponseBody) SetErrorCode(v string) *UpdateProjectSprintResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectSprintResponseBody) SetSuccessful(v bool) *UpdateProjectSprintResponseBody {
	s.Successful = &v
	return s
}

type UpdateProjectSprintResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectSprintResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectSprintResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectSprintResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectSprintResponse) SetHeaders(v map[string]*string) *UpdateProjectSprintResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectSprintResponse) SetBody(v *UpdateProjectSprintResponseBody) *UpdateProjectSprintResponse {
	s.Body = v
	return s
}

type UpdateProjectTaskRequest struct {
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

func (s UpdateProjectTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectTaskRequest) SetOrgId(v string) *UpdateProjectTaskRequest {
	s.OrgId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetContent(v string) *UpdateProjectTaskRequest {
	s.Content = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetProjectId(v string) *UpdateProjectTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetExecutorId(v string) *UpdateProjectTaskRequest {
	s.ExecutorId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetStartDate(v string) *UpdateProjectTaskRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetDueDate(v string) *UpdateProjectTaskRequest {
	s.DueDate = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetScenarioFiieldConfigId(v string) *UpdateProjectTaskRequest {
	s.ScenarioFiieldConfigId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetTaskFlowStatusId(v string) *UpdateProjectTaskRequest {
	s.TaskFlowStatusId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetNote(v string) *UpdateProjectTaskRequest {
	s.Note = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetPriority(v int32) *UpdateProjectTaskRequest {
	s.Priority = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetVisible(v string) *UpdateProjectTaskRequest {
	s.Visible = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetParentTaskId(v string) *UpdateProjectTaskRequest {
	s.ParentTaskId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetSprintId(v string) *UpdateProjectTaskRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateProjectTaskRequest) SetTaskId(v string) *UpdateProjectTaskRequest {
	s.TaskId = &v
	return s
}

type UpdateProjectTaskResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	Object     *bool   `json:"Object,omitempty" xml:"Object,omitempty"`
	ErrorCode  *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Successful *bool   `json:"Successful,omitempty" xml:"Successful,omitempty"`
}

func (s UpdateProjectTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectTaskResponseBody) SetRequestId(v string) *UpdateProjectTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectTaskResponseBody) SetErrorMsg(v string) *UpdateProjectTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateProjectTaskResponseBody) SetObject(v bool) *UpdateProjectTaskResponseBody {
	s.Object = &v
	return s
}

func (s *UpdateProjectTaskResponseBody) SetErrorCode(v string) *UpdateProjectTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectTaskResponseBody) SetSuccessful(v bool) *UpdateProjectTaskResponseBody {
	s.Successful = &v
	return s
}

type UpdateProjectTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectTaskResponse) SetHeaders(v map[string]*string) *UpdateProjectTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectTaskResponse) SetBody(v *UpdateProjectTaskResponseBody) *UpdateProjectTaskResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("teambition-aliyun"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddProjectMembersWithOptions(request *AddProjectMembersRequest, runtime *util.RuntimeOptions) (_result *AddProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddProjectMembers"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddProjectMembers(request *AddProjectMembersRequest) (_result *AddProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddProjectMembersResponse{}
	_body, _err := client.AddProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplySmallMicroWithOptions(request *ApplySmallMicroRequest, runtime *util.RuntimeOptions) (_result *ApplySmallMicroResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApplySmallMicroResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplySmallMicro"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplySmallMicro(request *ApplySmallMicroRequest) (_result *ApplySmallMicroResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplySmallMicroResponse{}
	_body, _err := client.ApplySmallMicroWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BactchInsertMembersWithOptions(request *BactchInsertMembersRequest, runtime *util.RuntimeOptions) (_result *BactchInsertMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BactchInsertMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BactchInsertMembers"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BactchInsertMembers(request *BactchInsertMembersRequest) (_result *BactchInsertMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BactchInsertMembersResponse{}
	_body, _err := client.BactchInsertMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAliyunUserExistsWithOptions(request *CheckAliyunUserExistsRequest, runtime *util.RuntimeOptions) (_result *CheckAliyunUserExistsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckAliyunUserExistsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckAliyunUserExists"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAliyunUserExists(request *CheckAliyunUserExistsRequest) (_result *CheckAliyunUserExistsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAliyunUserExistsResponse{}
	_body, _err := client.CheckAliyunUserExistsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDevopsOrgWithOptions(request *CreateDevopsOrgRequest, runtime *util.RuntimeOptions) (_result *CreateDevopsOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDevopsOrgResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDevopsOrg"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDevopsOrg(request *CreateDevopsOrgRequest) (_result *CreateDevopsOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDevopsOrgResponse{}
	_body, _err := client.CreateDevopsOrgWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateProjectSprintWithOptions(request *CreateProjectSprintRequest, runtime *util.RuntimeOptions) (_result *CreateProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProjectSprint"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProjectSprint(request *CreateProjectSprintRequest) (_result *CreateProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectSprintResponse{}
	_body, _err := client.CreateProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectTaskWithOptions(request *CreateProjectTaskRequest, runtime *util.RuntimeOptions) (_result *CreateProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProjectTask"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProjectTask(request *CreateProjectTaskRequest) (_result *CreateProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectTaskResponse{}
	_body, _err := client.CreateProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMembersForOrgWithOptions(request *DeleteMembersForOrgRequest, runtime *util.RuntimeOptions) (_result *DeleteMembersForOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMembersForOrgResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMembersForOrg"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMembersForOrg(request *DeleteMembersForOrgRequest) (_result *DeleteMembersForOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMembersForOrgResponse{}
	_body, _err := client.DeleteMembersForOrgWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteProject"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteProjectMembersWithOptions(request *DeleteProjectMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProjectMembers"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectMembers(request *DeleteProjectMembersRequest) (_result *DeleteProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectMembersResponse{}
	_body, _err := client.DeleteProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectSprintWithOptions(request *DeleteProjectSprintRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProjectSprint"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectSprint(request *DeleteProjectSprintRequest) (_result *DeleteProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectSprintResponse{}
	_body, _err := client.DeleteProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectTaskWithOptions(request *DeleteProjectTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteProjectTask"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProjectTask(request *DeleteProjectTaskRequest) (_result *DeleteProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectTaskResponse{}
	_body, _err := client.DeleteProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOrganizationMembersWithOptions(request *GetOrganizationMembersRequest, runtime *util.RuntimeOptions) (_result *GetOrganizationMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOrganizationMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOrganizationMembers"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOrganizationMembers(request *GetOrganizationMembersRequest) (_result *GetOrganizationMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrganizationMembersResponse{}
	_body, _err := client.GetOrganizationMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectInfoWithOptions(request *GetProjectInfoRequest, runtime *util.RuntimeOptions) (_result *GetProjectInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectInfo"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectInfo(request *GetProjectInfoRequest) (_result *GetProjectInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectInfoResponse{}
	_body, _err := client.GetProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectMembersWithOptions(request *GetProjectMembersRequest, runtime *util.RuntimeOptions) (_result *GetProjectMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectMembers"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectMembers(request *GetProjectMembersRequest) (_result *GetProjectMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectMembersResponse{}
	_body, _err := client.GetProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectSprintInfoWithOptions(request *GetProjectSprintInfoRequest, runtime *util.RuntimeOptions) (_result *GetProjectSprintInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectSprintInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectSprintInfo"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectSprintInfo(request *GetProjectSprintInfoRequest) (_result *GetProjectSprintInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectSprintInfoResponse{}
	_body, _err := client.GetProjectSprintInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectTaskInfoWithOptions(request *GetProjectTaskInfoRequest, runtime *util.RuntimeOptions) (_result *GetProjectTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetProjectTaskInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectTaskInfo"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectTaskInfo(request *GetProjectTaskInfoRequest) (_result *GetProjectTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectTaskInfoResponse{}
	_body, _err := client.GetProjectTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserByUidWithOptions(request *GetUserByUidRequest, runtime *util.RuntimeOptions) (_result *GetUserByUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserByUidResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserByUid"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByUid(request *GetUserByUidRequest) (_result *GetUserByUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserByUidResponse{}
	_body, _err := client.GetUserByUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertDevopsMemberWithOptions(request *InsertDevopsMemberRequest, runtime *util.RuntimeOptions) (_result *InsertDevopsMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InsertDevopsMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InsertDevopsMember"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertDevopsMember(request *InsertDevopsMemberRequest) (_result *InsertDevopsMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertDevopsMemberResponse{}
	_body, _err := client.InsertDevopsMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectSprintsWithOptions(request *ListProjectSprintsRequest, runtime *util.RuntimeOptions) (_result *ListProjectSprintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectSprintsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectSprints"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectSprints(request *ListProjectSprintsRequest) (_result *ListProjectSprintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectSprintsResponse{}
	_body, _err := client.ListProjectSprintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectTaskFlowWithOptions(request *ListProjectTaskFlowRequest, runtime *util.RuntimeOptions) (_result *ListProjectTaskFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectTaskFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectTaskFlow"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectTaskFlow(request *ListProjectTaskFlowRequest) (_result *ListProjectTaskFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectTaskFlowResponse{}
	_body, _err := client.ListProjectTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectTaskFlowStatusWithOptions(request *ListProjectTaskFlowStatusRequest, runtime *util.RuntimeOptions) (_result *ListProjectTaskFlowStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectTaskFlowStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectTaskFlowStatus"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectTaskFlowStatus(request *ListProjectTaskFlowStatusRequest) (_result *ListProjectTaskFlowStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectTaskFlowStatusResponse{}
	_body, _err := client.ListProjectTaskFlowStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectTasksWithOptions(request *ListProjectTasksRequest, runtime *util.RuntimeOptions) (_result *ListProjectTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListProjectTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListProjectTasks"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectTasks(request *ListProjectTasksRequest) (_result *ListProjectTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProjectTasksResponse{}
	_body, _err := client.ListProjectTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScenarioFieldConfigWithOptions(request *ListScenarioFieldConfigRequest, runtime *util.RuntimeOptions) (_result *ListScenarioFieldConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScenarioFieldConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScenarioFieldConfig"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScenarioFieldConfig(request *ListScenarioFieldConfigRequest) (_result *ListScenarioFieldConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScenarioFieldConfigResponse{}
	_body, _err := client.ListScenarioFieldConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectWithOptions(request *UpdateProjectRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProject"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectSprintWithOptions(request *UpdateProjectSprintRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectSprintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectSprintResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProjectSprint"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectSprint(request *UpdateProjectSprintRequest) (_result *UpdateProjectSprintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectSprintResponse{}
	_body, _err := client.UpdateProjectSprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectTaskWithOptions(request *UpdateProjectTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProjectTask"), tea.String("2020-02-26"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectTask(request *UpdateProjectTaskRequest) (_result *UpdateProjectTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectTaskResponse{}
	_body, _err := client.UpdateProjectTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
