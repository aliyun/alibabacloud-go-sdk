// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddEnterpriseMemberRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StaffId    *string `json:"StaffId,omitempty" xml:"StaffId,omitempty"`
	Operator   *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s AddEnterpriseMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseMemberRequest) GoString() string {
	return s.String()
}

func (s *AddEnterpriseMemberRequest) SetInstanceId(v string) *AddEnterpriseMemberRequest {
	s.InstanceId = &v
	return s
}

func (s *AddEnterpriseMemberRequest) SetStaffId(v string) *AddEnterpriseMemberRequest {
	s.StaffId = &v
	return s
}

func (s *AddEnterpriseMemberRequest) SetOperator(v string) *AddEnterpriseMemberRequest {
	s.Operator = &v
	return s
}

type AddEnterpriseMemberResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddEnterpriseMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddEnterpriseMemberResponseBody) SetMessage(v string) *AddEnterpriseMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddEnterpriseMemberResponseBody) SetRequestId(v string) *AddEnterpriseMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEnterpriseMemberResponseBody) SetData(v bool) *AddEnterpriseMemberResponseBody {
	s.Data = &v
	return s
}

func (s *AddEnterpriseMemberResponseBody) SetCode(v int32) *AddEnterpriseMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddEnterpriseMemberResponseBody) SetSuccess(v bool) *AddEnterpriseMemberResponseBody {
	s.Success = &v
	return s
}

type AddEnterpriseMemberResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddEnterpriseMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddEnterpriseMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddEnterpriseMemberResponse) GoString() string {
	return s.String()
}

func (s *AddEnterpriseMemberResponse) SetHeaders(v map[string]*string) *AddEnterpriseMemberResponse {
	s.Headers = v
	return s
}

func (s *AddEnterpriseMemberResponse) SetBody(v *AddEnterpriseMemberResponseBody) *AddEnterpriseMemberResponse {
	s.Body = v
	return s
}

type AddRamMemberRequest struct {
	CorpIdentifier  *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	StaffIdentifier *string `json:"StaffIdentifier,omitempty" xml:"StaffIdentifier,omitempty"`
}

func (s AddRamMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRamMemberRequest) GoString() string {
	return s.String()
}

func (s *AddRamMemberRequest) SetCorpIdentifier(v string) *AddRamMemberRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *AddRamMemberRequest) SetStaffIdentifier(v string) *AddRamMemberRequest {
	s.StaffIdentifier = &v
	return s
}

type AddRamMemberResponseBody struct {
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRamMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRamMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddRamMemberResponseBody) SetMessage(v string) *AddRamMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddRamMemberResponseBody) SetRequestId(v string) *AddRamMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRamMemberResponseBody) SetData(v []*string) *AddRamMemberResponseBody {
	s.Data = v
	return s
}

func (s *AddRamMemberResponseBody) SetCode(v int32) *AddRamMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddRamMemberResponseBody) SetSuccess(v bool) *AddRamMemberResponseBody {
	s.Success = &v
	return s
}

type AddRamMemberResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRamMemberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRamMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRamMemberResponse) GoString() string {
	return s.String()
}

func (s *AddRamMemberResponse) SetHeaders(v map[string]*string) *AddRamMemberResponse {
	s.Headers = v
	return s
}

func (s *AddRamMemberResponse) SetBody(v *AddRamMemberResponseBody) *AddRamMemberResponse {
	s.Body = v
	return s
}

type ApproveJoinCompanyRequest struct {
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	ApplicationIds *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
}

func (s ApproveJoinCompanyRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveJoinCompanyRequest) GoString() string {
	return s.String()
}

func (s *ApproveJoinCompanyRequest) SetCorpIdentifier(v string) *ApproveJoinCompanyRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *ApproveJoinCompanyRequest) SetApplicationIds(v string) *ApproveJoinCompanyRequest {
	s.ApplicationIds = &v
	return s
}

type ApproveJoinCompanyResponseBody struct {
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApproveJoinCompanyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveJoinCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveJoinCompanyResponseBody) SetMessage(v string) *ApproveJoinCompanyResponseBody {
	s.Message = &v
	return s
}

func (s *ApproveJoinCompanyResponseBody) SetRequestId(v string) *ApproveJoinCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveJoinCompanyResponseBody) SetData(v []*string) *ApproveJoinCompanyResponseBody {
	s.Data = v
	return s
}

func (s *ApproveJoinCompanyResponseBody) SetCode(v int32) *ApproveJoinCompanyResponseBody {
	s.Code = &v
	return s
}

func (s *ApproveJoinCompanyResponseBody) SetSuccess(v bool) *ApproveJoinCompanyResponseBody {
	s.Success = &v
	return s
}

type ApproveJoinCompanyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApproveJoinCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApproveJoinCompanyResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveJoinCompanyResponse) GoString() string {
	return s.String()
}

func (s *ApproveJoinCompanyResponse) SetHeaders(v map[string]*string) *ApproveJoinCompanyResponse {
	s.Headers = v
	return s
}

func (s *ApproveJoinCompanyResponse) SetBody(v *ApproveJoinCompanyResponseBody) *ApproveJoinCompanyResponse {
	s.Body = v
	return s
}

type CreateEnterpriseRequest struct {
	Name           *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Emails         map[string]interface{} `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Domain         *string                `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CreatorStaffId *string                `json:"CreatorStaffId,omitempty" xml:"CreatorStaffId,omitempty"`
}

func (s CreateEnterpriseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseRequest) SetName(v string) *CreateEnterpriseRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseRequest) SetDescription(v string) *CreateEnterpriseRequest {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseRequest) SetEmails(v map[string]interface{}) *CreateEnterpriseRequest {
	s.Emails = v
	return s
}

func (s *CreateEnterpriseRequest) SetDomain(v string) *CreateEnterpriseRequest {
	s.Domain = &v
	return s
}

func (s *CreateEnterpriseRequest) SetCreatorStaffId(v string) *CreateEnterpriseRequest {
	s.CreatorStaffId = &v
	return s
}

type CreateEnterpriseShrinkRequest struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EmailsShrink   *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	Domain         *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	CreatorStaffId *string `json:"CreatorStaffId,omitempty" xml:"CreatorStaffId,omitempty"`
}

func (s CreateEnterpriseShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseShrinkRequest) SetName(v string) *CreateEnterpriseShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseShrinkRequest) SetDescription(v string) *CreateEnterpriseShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateEnterpriseShrinkRequest) SetEmailsShrink(v string) *CreateEnterpriseShrinkRequest {
	s.EmailsShrink = &v
	return s
}

func (s *CreateEnterpriseShrinkRequest) SetDomain(v string) *CreateEnterpriseShrinkRequest {
	s.Domain = &v
	return s
}

func (s *CreateEnterpriseShrinkRequest) SetCreatorStaffId(v string) *CreateEnterpriseShrinkRequest {
	s.CreatorStaffId = &v
	return s
}

type CreateEnterpriseResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateEnterpriseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEnterpriseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseResponseBody) SetMessage(v string) *CreateEnterpriseResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnterpriseResponseBody) SetRequestId(v string) *CreateEnterpriseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnterpriseResponseBody) SetData(v *CreateEnterpriseResponseBodyData) *CreateEnterpriseResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnterpriseResponseBody) SetCode(v int32) *CreateEnterpriseResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnterpriseResponseBody) SetSuccess(v bool) *CreateEnterpriseResponseBody {
	s.Success = &v
	return s
}

type CreateEnterpriseResponseBodyData struct {
	Status       *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	ProhibitCode *int32  `json:"ProhibitCode,omitempty" xml:"ProhibitCode,omitempty"`
	Identifier   *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateEnterpriseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseResponseBodyData) SetStatus(v int32) *CreateEnterpriseResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateEnterpriseResponseBodyData) SetType(v int32) *CreateEnterpriseResponseBodyData {
	s.Type = &v
	return s
}

func (s *CreateEnterpriseResponseBodyData) SetProhibitCode(v int32) *CreateEnterpriseResponseBodyData {
	s.ProhibitCode = &v
	return s
}

func (s *CreateEnterpriseResponseBodyData) SetIdentifier(v string) *CreateEnterpriseResponseBodyData {
	s.Identifier = &v
	return s
}

func (s *CreateEnterpriseResponseBodyData) SetName(v string) *CreateEnterpriseResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateEnterpriseResponseBodyData) SetId(v int64) *CreateEnterpriseResponseBodyData {
	s.Id = &v
	return s
}

type CreateEnterpriseResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEnterpriseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEnterpriseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseResponse) SetHeaders(v map[string]*string) *CreateEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseResponse) SetBody(v *CreateEnterpriseResponseBody) *CreateEnterpriseResponse {
	s.Body = v
	return s
}

type CreateWorkitemRequest struct {
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	Author         *string `json:"Author,omitempty" xml:"Author,omitempty"`
	AssignedTo     *string `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	TemplateId     *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Subject        *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Stamp          *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	AKProjectId    *int32  `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	CfList         *string `json:"CfList,omitempty" xml:"CfList,omitempty"`
	Verifier       *string `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	PriorityId     *int32  `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	SeriousLevelId *int32  `json:"SeriousLevelId,omitempty" xml:"SeriousLevelId,omitempty"`
	WatcherUsers   *string `json:"WatcherUsers,omitempty" xml:"WatcherUsers,omitempty"`
	ModuleIds      *string `json:"ModuleIds,omitempty" xml:"ModuleIds,omitempty"`
}

func (s CreateWorkitemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRequest) SetCorpIdentifier(v string) *CreateWorkitemRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *CreateWorkitemRequest) SetAuthor(v string) *CreateWorkitemRequest {
	s.Author = &v
	return s
}

func (s *CreateWorkitemRequest) SetAssignedTo(v string) *CreateWorkitemRequest {
	s.AssignedTo = &v
	return s
}

func (s *CreateWorkitemRequest) SetTemplateId(v int32) *CreateWorkitemRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkitemRequest) SetSubject(v string) *CreateWorkitemRequest {
	s.Subject = &v
	return s
}

func (s *CreateWorkitemRequest) SetDescription(v string) *CreateWorkitemRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRequest) SetStamp(v string) *CreateWorkitemRequest {
	s.Stamp = &v
	return s
}

func (s *CreateWorkitemRequest) SetAKProjectId(v int32) *CreateWorkitemRequest {
	s.AKProjectId = &v
	return s
}

func (s *CreateWorkitemRequest) SetCfList(v string) *CreateWorkitemRequest {
	s.CfList = &v
	return s
}

func (s *CreateWorkitemRequest) SetVerifier(v string) *CreateWorkitemRequest {
	s.Verifier = &v
	return s
}

func (s *CreateWorkitemRequest) SetPriorityId(v int32) *CreateWorkitemRequest {
	s.PriorityId = &v
	return s
}

func (s *CreateWorkitemRequest) SetSeriousLevelId(v int32) *CreateWorkitemRequest {
	s.SeriousLevelId = &v
	return s
}

func (s *CreateWorkitemRequest) SetWatcherUsers(v string) *CreateWorkitemRequest {
	s.WatcherUsers = &v
	return s
}

func (s *CreateWorkitemRequest) SetModuleIds(v string) *CreateWorkitemRequest {
	s.ModuleIds = &v
	return s
}

type CreateWorkitemResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWorkitemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponseBody) SetMessage(v string) *CreateWorkitemResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetRequestId(v string) *CreateWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetData(v int32) *CreateWorkitemResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetCode(v int32) *CreateWorkitemResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkitemResponseBody) SetSuccess(v bool) *CreateWorkitemResponseBody {
	s.Success = &v
	return s
}

type CreateWorkitemResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateWorkitemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkitemResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkitemResponse) SetHeaders(v map[string]*string) *CreateWorkitemResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkitemResponse) SetBody(v *CreateWorkitemResponseBody) *CreateWorkitemResponse {
	s.Body = v
	return s
}

type GetBindedUserByDingIdRequest struct {
	DingId *string `json:"DingId,omitempty" xml:"DingId,omitempty"`
}

func (s GetBindedUserByDingIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdRequest) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdRequest) SetDingId(v string) *GetBindedUserByDingIdRequest {
	s.DingId = &v
	return s
}

type GetBindedUserByDingIdResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetBindedUserByDingIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBindedUserByDingIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponseBody) SetRequestId(v string) *GetBindedUserByDingIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBody) SetData(v *GetBindedUserByDingIdResponseBodyData) *GetBindedUserByDingIdResponseBody {
	s.Data = v
	return s
}

func (s *GetBindedUserByDingIdResponseBody) SetCode(v int32) *GetBindedUserByDingIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBody) SetSuccess(v bool) *GetBindedUserByDingIdResponseBody {
	s.Success = &v
	return s
}

type GetBindedUserByDingIdResponseBodyData struct {
	Uuid            *string                                              `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Guid            *string                                              `json:"Guid,omitempty" xml:"Guid,omitempty"`
	DingtalkUser    *GetBindedUserByDingIdResponseBodyDataDingtalkUser   `json:"DingtalkUser,omitempty" xml:"DingtalkUser,omitempty" type:"Struct"`
	UserProfileDTO  *GetBindedUserByDingIdResponseBodyDataUserProfileDTO `json:"UserProfileDTO,omitempty" xml:"UserProfileDTO,omitempty" type:"Struct"`
	NickName        *string                                              `json:"NickName,omitempty" xml:"NickName,omitempty"`
	MainAccountType *string                                              `json:"MainAccountType,omitempty" xml:"MainAccountType,omitempty"`
	AliyunUser      *GetBindedUserByDingIdResponseBodyDataAliyunUser     `json:"AliyunUser,omitempty" xml:"AliyunUser,omitempty" type:"Struct"`
	IsValid         *bool                                                `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	Id              *int32                                               `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBindedUserByDingIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponseBodyData) SetUuid(v string) *GetBindedUserByDingIdResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetGuid(v string) *GetBindedUserByDingIdResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetDingtalkUser(v *GetBindedUserByDingIdResponseBodyDataDingtalkUser) *GetBindedUserByDingIdResponseBodyData {
	s.DingtalkUser = v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetUserProfileDTO(v *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) *GetBindedUserByDingIdResponseBodyData {
	s.UserProfileDTO = v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetNickName(v string) *GetBindedUserByDingIdResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetMainAccountType(v string) *GetBindedUserByDingIdResponseBodyData {
	s.MainAccountType = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetAliyunUser(v *GetBindedUserByDingIdResponseBodyDataAliyunUser) *GetBindedUserByDingIdResponseBodyData {
	s.AliyunUser = v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetIsValid(v bool) *GetBindedUserByDingIdResponseBodyData {
	s.IsValid = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyData) SetId(v int32) *GetBindedUserByDingIdResponseBodyData {
	s.Id = &v
	return s
}

type GetBindedUserByDingIdResponseBodyDataDingtalkUser struct {
	DingtalkUserId *int32  `json:"DingtalkUserId,omitempty" xml:"DingtalkUserId,omitempty"`
	DingId         *string `json:"DingId,omitempty" xml:"DingId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
	Nick           *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	CodeUserName   *string `json:"CodeUserName,omitempty" xml:"CodeUserName,omitempty"`
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBindedUserByDingIdResponseBodyDataDingtalkUser) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponseBodyDataDingtalkUser) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetDingtalkUserId(v int32) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.DingtalkUserId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetDingId(v string) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.DingId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetUnionId(v string) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.UnionId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetNick(v string) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.Nick = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetCodeUserName(v string) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.CodeUserName = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataDingtalkUser) SetId(v int32) *GetBindedUserByDingIdResponseBodyDataDingtalkUser {
	s.Id = &v
	return s
}

type GetBindedUserByDingIdResponseBodyDataUserProfileDTO struct {
	DataSource  *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Avatar      *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CreatedAt   *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	NickName    *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetBindedUserByDingIdResponseBodyDataUserProfileDTO) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponseBodyDataUserProfileDTO) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetDataSource(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.DataSource = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetAvatar(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.Avatar = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetEmail(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.Email = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetMobile(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.Mobile = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetUserId(v int32) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.UserId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetCreatedAt(v int64) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.CreatedAt = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetEnglishName(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.EnglishName = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetNickName(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.NickName = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataUserProfileDTO) SetName(v string) *GetBindedUserByDingIdResponseBodyDataUserProfileDTO {
	s.Name = &v
	return s
}

type GetBindedUserByDingIdResponseBodyDataAliyunUser struct {
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HavanaId         *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	TaobaoNick       *string `json:"TaobaoNick,omitempty" xml:"TaobaoNick,omitempty"`
	Kp               *string `json:"Kp,omitempty" xml:"Kp,omitempty"`
	Realname         *string `json:"Realname,omitempty" xml:"Realname,omitempty"`
	NickName         *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	AccountStructure *int32  `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	AliyunId         *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetBindedUserByDingIdResponseBodyDataAliyunUser) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponseBodyDataAliyunUser) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetEmail(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.Email = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetHavanaId(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.HavanaId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetTaobaoNick(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.TaobaoNick = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetKp(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.Kp = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetRealname(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.Realname = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetNickName(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.NickName = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetAccountStructure(v int32) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.AccountStructure = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetAliyunId(v string) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.AliyunId = &v
	return s
}

func (s *GetBindedUserByDingIdResponseBodyDataAliyunUser) SetId(v int32) *GetBindedUserByDingIdResponseBodyDataAliyunUser {
	s.Id = &v
	return s
}

type GetBindedUserByDingIdResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBindedUserByDingIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBindedUserByDingIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBindedUserByDingIdResponse) GoString() string {
	return s.String()
}

func (s *GetBindedUserByDingIdResponse) SetHeaders(v map[string]*string) *GetBindedUserByDingIdResponse {
	s.Headers = v
	return s
}

func (s *GetBindedUserByDingIdResponse) SetBody(v *GetBindedUserByDingIdResponseBody) *GetBindedUserByDingIdResponse {
	s.Body = v
	return s
}

type GetChangeLogRequest struct {
	TargetType     *string                `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetIds      map[string]interface{} `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	CorpIdentifier *string                `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetChangeLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeLogRequest) GoString() string {
	return s.String()
}

func (s *GetChangeLogRequest) SetTargetType(v string) *GetChangeLogRequest {
	s.TargetType = &v
	return s
}

func (s *GetChangeLogRequest) SetTargetIds(v map[string]interface{}) *GetChangeLogRequest {
	s.TargetIds = v
	return s
}

func (s *GetChangeLogRequest) SetCorpIdentifier(v string) *GetChangeLogRequest {
	s.CorpIdentifier = &v
	return s
}

type GetChangeLogShrinkRequest struct {
	TargetType      *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	TargetIdsShrink *string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty"`
	CorpIdentifier  *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetChangeLogShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeLogShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetChangeLogShrinkRequest) SetTargetType(v string) *GetChangeLogShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *GetChangeLogShrinkRequest) SetTargetIdsShrink(v string) *GetChangeLogShrinkRequest {
	s.TargetIdsShrink = &v
	return s
}

func (s *GetChangeLogShrinkRequest) SetCorpIdentifier(v string) *GetChangeLogShrinkRequest {
	s.CorpIdentifier = &v
	return s
}

type GetChangeLogResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetChangeLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetChangeLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChangeLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeLogResponseBody) SetRequestId(v string) *GetChangeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeLogResponseBody) SetData(v []*GetChangeLogResponseBodyData) *GetChangeLogResponseBody {
	s.Data = v
	return s
}

func (s *GetChangeLogResponseBody) SetCode(v int32) *GetChangeLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetChangeLogResponseBody) SetSuccess(v bool) *GetChangeLogResponseBody {
	s.Success = &v
	return s
}

type GetChangeLogResponseBodyData struct {
	OldValue     *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	PropertyType *string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	PropertyKey  *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	TargetId     *int32  `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	Other        *string `json:"Other,omitempty" xml:"Other,omitempty"`
	NewValue     *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	TargetType   *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetChangeLogResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetChangeLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChangeLogResponseBodyData) SetOldValue(v string) *GetChangeLogResponseBodyData {
	s.OldValue = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetPropertyType(v string) *GetChangeLogResponseBodyData {
	s.PropertyType = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetPropertyKey(v string) *GetChangeLogResponseBodyData {
	s.PropertyKey = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetTargetId(v int32) *GetChangeLogResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetOther(v string) *GetChangeLogResponseBodyData {
	s.Other = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetNewValue(v string) *GetChangeLogResponseBodyData {
	s.NewValue = &v
	return s
}

func (s *GetChangeLogResponseBodyData) SetTargetType(v string) *GetChangeLogResponseBodyData {
	s.TargetType = &v
	return s
}

type GetChangeLogResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetChangeLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChangeLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChangeLogResponse) GoString() string {
	return s.String()
}

func (s *GetChangeLogResponse) SetHeaders(v map[string]*string) *GetChangeLogResponse {
	s.Headers = v
	return s
}

func (s *GetChangeLogResponse) SetBody(v *GetChangeLogResponseBody) *GetChangeLogResponse {
	s.Body = v
	return s
}

type GetCustomFieldsByTemplateIdRequest struct {
	AKProjectId    *int32  `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	TemplateId     *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetCustomFieldsByTemplateIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldsByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *GetCustomFieldsByTemplateIdRequest) SetAKProjectId(v int32) *GetCustomFieldsByTemplateIdRequest {
	s.AKProjectId = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdRequest) SetTemplateId(v int32) *GetCustomFieldsByTemplateIdRequest {
	s.TemplateId = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdRequest) SetCorpIdentifier(v string) *GetCustomFieldsByTemplateIdRequest {
	s.CorpIdentifier = &v
	return s
}

type GetCustomFieldsByTemplateIdResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetCustomFieldsByTemplateIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCustomFieldsByTemplateIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldsByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomFieldsByTemplateIdResponseBody) SetRequestId(v string) *GetCustomFieldsByTemplateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBody) SetData(v []*GetCustomFieldsByTemplateIdResponseBodyData) *GetCustomFieldsByTemplateIdResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBody) SetCode(v int32) *GetCustomFieldsByTemplateIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBody) SetSuccess(v bool) *GetCustomFieldsByTemplateIdResponseBody {
	s.Success = &v
	return s
}

type GetCustomFieldsByTemplateIdResponseBodyData struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Editable       *bool   `json:"Editable,omitempty" xml:"Editable,omitempty"`
	CreatedAt      *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	IsRemember     *bool   `json:"IsRemember,omitempty" xml:"IsRemember,omitempty"`
	IsDelete       *bool   `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	DefaultValue   *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	MaxLength      *int32  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	FieldFormat    *string `json:"FieldFormat,omitempty" xml:"FieldFormat,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Dynamic        *bool   `json:"Dynamic,omitempty" xml:"Dynamic,omitempty"`
	IsRequired     *bool   `json:"IsRequired,omitempty" xml:"IsRequired,omitempty"`
	NameI18N       *string `json:"NameI18N,omitempty" xml:"NameI18N,omitempty"`
	PossibleValues *string `json:"PossibleValues,omitempty" xml:"PossibleValues,omitempty"`
	MinLength      *int32  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	UpdatedAt      *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Other          *string `json:"Other,omitempty" xml:"Other,omitempty"`
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetCustomFieldsByTemplateIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldsByTemplateIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetType(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetEditable(v bool) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Editable = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetCreatedAt(v int64) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetIsRemember(v bool) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.IsRemember = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetIsDelete(v bool) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.IsDelete = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetDefaultValue(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetMaxLength(v int32) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.MaxLength = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetFieldFormat(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.FieldFormat = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetDescription(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetDynamic(v bool) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Dynamic = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetIsRequired(v bool) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.IsRequired = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetNameI18N(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.NameI18N = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetPossibleValues(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.PossibleValues = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetMinLength(v int32) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.MinLength = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetUpdatedAt(v int64) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetName(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetOther(v string) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Other = &v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponseBodyData) SetId(v int32) *GetCustomFieldsByTemplateIdResponseBodyData {
	s.Id = &v
	return s
}

type GetCustomFieldsByTemplateIdResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCustomFieldsByTemplateIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomFieldsByTemplateIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomFieldsByTemplateIdResponse) GoString() string {
	return s.String()
}

func (s *GetCustomFieldsByTemplateIdResponse) SetHeaders(v map[string]*string) *GetCustomFieldsByTemplateIdResponse {
	s.Headers = v
	return s
}

func (s *GetCustomFieldsByTemplateIdResponse) SetBody(v *GetCustomFieldsByTemplateIdResponseBody) *GetCustomFieldsByTemplateIdResponse {
	s.Body = v
	return s
}

type GetIssueByIdRequest struct {
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetIssueByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIssueByIdRequest) GoString() string {
	return s.String()
}

func (s *GetIssueByIdRequest) SetId(v int32) *GetIssueByIdRequest {
	s.Id = &v
	return s
}

func (s *GetIssueByIdRequest) SetCorpIdentifier(v string) *GetIssueByIdRequest {
	s.CorpIdentifier = &v
	return s
}

type GetIssueByIdResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetIssueByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIssueByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIssueByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetIssueByIdResponseBody) SetRequestId(v string) *GetIssueByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIssueByIdResponseBody) SetData(v *GetIssueByIdResponseBodyData) *GetIssueByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetIssueByIdResponseBody) SetCode(v int32) *GetIssueByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetIssueByIdResponseBody) SetSuccess(v string) *GetIssueByIdResponseBody {
	s.Success = &v
	return s
}

type GetIssueByIdResponseBodyData struct {
	SkipCollab            *bool                                  `json:"SkipCollab,omitempty" xml:"SkipCollab,omitempty"`
	SendWangwang          *bool                                  `json:"SendWangwang,omitempty" xml:"SendWangwang,omitempty"`
	VersionList           *string                                `json:"VersionList,omitempty" xml:"VersionList,omitempty"`
	AkProjectId           *int32                                 `json:"AkProjectId,omitempty" xml:"AkProjectId,omitempty"`
	UserId                *int32                                 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IdPath                *string                                `json:"IdPath,omitempty" xml:"IdPath,omitempty"`
	IgnoreIntegrate       *bool                                  `json:"IgnoreIntegrate,omitempty" xml:"IgnoreIntegrate,omitempty"`
	AttachmentIds         *string                                `json:"AttachmentIds,omitempty" xml:"AttachmentIds,omitempty"`
	CommitDate            *int64                                 `json:"CommitDate,omitempty" xml:"CommitDate,omitempty"`
	UpdateStatusAt        *int64                                 `json:"UpdateStatusAt,omitempty" xml:"UpdateStatusAt,omitempty"`
	Trackers              *string                                `json:"Trackers,omitempty" xml:"Trackers,omitempty"`
	Subject               *string                                `json:"Subject,omitempty" xml:"Subject,omitempty"`
	UserStaffId           *string                                `json:"UserStaffId,omitempty" xml:"UserStaffId,omitempty"`
	AttachmentList        *string                                `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty"`
	VerifierId            *int32                                 `json:"VerifierId,omitempty" xml:"VerifierId,omitempty"`
	IgnoreCheck           *bool                                  `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	ModuleList            *string                                `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Attachmented          *bool                                  `json:"Attachmented,omitempty" xml:"Attachmented,omitempty"`
	SeriousLevelId        *int32                                 `json:"SeriousLevelId,omitempty" xml:"SeriousLevelId,omitempty"`
	TagIdList             *string                                `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	Watched               *bool                                  `json:"Watched,omitempty" xml:"Watched,omitempty"`
	AssignedTo            *string                                `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	AssignedToIds         *string                                `json:"AssignedToIds,omitempty" xml:"AssignedToIds,omitempty"`
	PriorityId            *int32                                 `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	RecordChangeLog       *bool                                  `json:"RecordChangeLog,omitempty" xml:"RecordChangeLog,omitempty"`
	UpdatedAt             *int64                                 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	TemplateId            *int32                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Verifier              *string                                `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	Status                *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	RelatedUserIds        *string                                `json:"RelatedUserIds,omitempty" xml:"RelatedUserIds,omitempty"`
	RelatedAKProjectIds   *string                                `json:"RelatedAKProjectIds,omitempty" xml:"RelatedAKProjectIds,omitempty"`
	RelatedAKProjectGuids *string                                `json:"RelatedAKProjectGuids,omitempty" xml:"RelatedAKProjectGuids,omitempty"`
	CreatedAt             *int64                                 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AssignedToIdList      *string                                `json:"AssignedToIdList,omitempty" xml:"AssignedToIdList,omitempty"`
	CfsList               []*GetIssueByIdResponseBodyDataCfsList `json:"CfsList,omitempty" xml:"CfsList,omitempty" type:"Repeated"`
	Priority              *string                                `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StatusStage           *int32                                 `json:"StatusStage,omitempty" xml:"StatusStage,omitempty"`
	AssignedToStaffId     *string                                `json:"AssignedToStaffId,omitempty" xml:"AssignedToStaffId,omitempty"`
	ParentId              *int32                                 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Description           *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	LogicalStatus         *string                                `json:"LogicalStatus,omitempty" xml:"LogicalStatus,omitempty"`
	VersionIds            *string                                `json:"VersionIds,omitempty" xml:"VersionIds,omitempty"`
	SeriousLevel          *string                                `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	StatusId              *int32                                 `json:"StatusId,omitempty" xml:"StatusId,omitempty"`
	Stamp                 *string                                `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	User                  *string                                `json:"User,omitempty" xml:"User,omitempty"`
	TrackerIds            *string                                `json:"TrackerIds,omitempty" xml:"TrackerIds,omitempty"`
	AssignedToId          *int32                                 `json:"AssignedToId,omitempty" xml:"AssignedToId,omitempty"`
	ChangeLogList         *string                                `json:"ChangeLogList,omitempty" xml:"ChangeLogList,omitempty"`
	IssueTypeId           *int32                                 `json:"IssueTypeId,omitempty" xml:"IssueTypeId,omitempty"`
	RegionId              *int32                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Guid                  *string                                `json:"Guid,omitempty" xml:"Guid,omitempty"`
	AssignedToMaps        *string                                `json:"AssignedToMaps,omitempty" xml:"AssignedToMaps,omitempty"`
	BlackListNotice       *string                                `json:"BlackListNotice,omitempty" xml:"BlackListNotice,omitempty"`
	ModuleUpdated         *bool                                  `json:"ModuleUpdated,omitempty" xml:"ModuleUpdated,omitempty"`
	ProjectIds            *string                                `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	CommentList           *string                                `json:"CommentList,omitempty" xml:"CommentList,omitempty"`
	VerifierStaffId       *string                                `json:"VerifierStaffId,omitempty" xml:"VerifierStaffId,omitempty"`
	ModuleIds             *string                                `json:"ModuleIds,omitempty" xml:"ModuleIds,omitempty"`
	Id                    *int32                                 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetIssueByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIssueByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIssueByIdResponseBodyData) SetSkipCollab(v bool) *GetIssueByIdResponseBodyData {
	s.SkipCollab = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetSendWangwang(v bool) *GetIssueByIdResponseBodyData {
	s.SendWangwang = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetVersionList(v string) *GetIssueByIdResponseBodyData {
	s.VersionList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAkProjectId(v int32) *GetIssueByIdResponseBodyData {
	s.AkProjectId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetUserId(v int32) *GetIssueByIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetIdPath(v string) *GetIssueByIdResponseBodyData {
	s.IdPath = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetIgnoreIntegrate(v bool) *GetIssueByIdResponseBodyData {
	s.IgnoreIntegrate = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAttachmentIds(v string) *GetIssueByIdResponseBodyData {
	s.AttachmentIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetCommitDate(v int64) *GetIssueByIdResponseBodyData {
	s.CommitDate = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetUpdateStatusAt(v int64) *GetIssueByIdResponseBodyData {
	s.UpdateStatusAt = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetTrackers(v string) *GetIssueByIdResponseBodyData {
	s.Trackers = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetSubject(v string) *GetIssueByIdResponseBodyData {
	s.Subject = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetUserStaffId(v string) *GetIssueByIdResponseBodyData {
	s.UserStaffId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAttachmentList(v string) *GetIssueByIdResponseBodyData {
	s.AttachmentList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetVerifierId(v int32) *GetIssueByIdResponseBodyData {
	s.VerifierId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetIgnoreCheck(v bool) *GetIssueByIdResponseBodyData {
	s.IgnoreCheck = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetModuleList(v string) *GetIssueByIdResponseBodyData {
	s.ModuleList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAttachmented(v bool) *GetIssueByIdResponseBodyData {
	s.Attachmented = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetSeriousLevelId(v int32) *GetIssueByIdResponseBodyData {
	s.SeriousLevelId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetTagIdList(v string) *GetIssueByIdResponseBodyData {
	s.TagIdList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetWatched(v bool) *GetIssueByIdResponseBodyData {
	s.Watched = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedTo(v string) *GetIssueByIdResponseBodyData {
	s.AssignedTo = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedToIds(v string) *GetIssueByIdResponseBodyData {
	s.AssignedToIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetPriorityId(v int32) *GetIssueByIdResponseBodyData {
	s.PriorityId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetRecordChangeLog(v bool) *GetIssueByIdResponseBodyData {
	s.RecordChangeLog = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetUpdatedAt(v int64) *GetIssueByIdResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetTemplateId(v int32) *GetIssueByIdResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetVerifier(v string) *GetIssueByIdResponseBodyData {
	s.Verifier = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetStatus(v string) *GetIssueByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetRelatedUserIds(v string) *GetIssueByIdResponseBodyData {
	s.RelatedUserIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetRelatedAKProjectIds(v string) *GetIssueByIdResponseBodyData {
	s.RelatedAKProjectIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetRelatedAKProjectGuids(v string) *GetIssueByIdResponseBodyData {
	s.RelatedAKProjectGuids = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetCreatedAt(v int64) *GetIssueByIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedToIdList(v string) *GetIssueByIdResponseBodyData {
	s.AssignedToIdList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetCfsList(v []*GetIssueByIdResponseBodyDataCfsList) *GetIssueByIdResponseBodyData {
	s.CfsList = v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetPriority(v string) *GetIssueByIdResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetStatusStage(v int32) *GetIssueByIdResponseBodyData {
	s.StatusStage = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedToStaffId(v string) *GetIssueByIdResponseBodyData {
	s.AssignedToStaffId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetParentId(v int32) *GetIssueByIdResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetDescription(v string) *GetIssueByIdResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetLogicalStatus(v string) *GetIssueByIdResponseBodyData {
	s.LogicalStatus = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetVersionIds(v string) *GetIssueByIdResponseBodyData {
	s.VersionIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetSeriousLevel(v string) *GetIssueByIdResponseBodyData {
	s.SeriousLevel = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetStatusId(v int32) *GetIssueByIdResponseBodyData {
	s.StatusId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetStamp(v string) *GetIssueByIdResponseBodyData {
	s.Stamp = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetUser(v string) *GetIssueByIdResponseBodyData {
	s.User = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetTrackerIds(v string) *GetIssueByIdResponseBodyData {
	s.TrackerIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedToId(v int32) *GetIssueByIdResponseBodyData {
	s.AssignedToId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetChangeLogList(v string) *GetIssueByIdResponseBodyData {
	s.ChangeLogList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetIssueTypeId(v int32) *GetIssueByIdResponseBodyData {
	s.IssueTypeId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetRegionId(v int32) *GetIssueByIdResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetGuid(v string) *GetIssueByIdResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetAssignedToMaps(v string) *GetIssueByIdResponseBodyData {
	s.AssignedToMaps = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetBlackListNotice(v string) *GetIssueByIdResponseBodyData {
	s.BlackListNotice = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetModuleUpdated(v bool) *GetIssueByIdResponseBodyData {
	s.ModuleUpdated = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetProjectIds(v string) *GetIssueByIdResponseBodyData {
	s.ProjectIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetCommentList(v string) *GetIssueByIdResponseBodyData {
	s.CommentList = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetVerifierStaffId(v string) *GetIssueByIdResponseBodyData {
	s.VerifierStaffId = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetModuleIds(v string) *GetIssueByIdResponseBodyData {
	s.ModuleIds = &v
	return s
}

func (s *GetIssueByIdResponseBodyData) SetId(v int32) *GetIssueByIdResponseBodyData {
	s.Id = &v
	return s
}

type GetIssueByIdResponseBodyDataCfsList struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetIssueByIdResponseBodyDataCfsList) String() string {
	return tea.Prettify(s)
}

func (s GetIssueByIdResponseBodyDataCfsList) GoString() string {
	return s.String()
}

func (s *GetIssueByIdResponseBodyDataCfsList) SetType(v string) *GetIssueByIdResponseBodyDataCfsList {
	s.Type = &v
	return s
}

func (s *GetIssueByIdResponseBodyDataCfsList) SetValue(v string) *GetIssueByIdResponseBodyDataCfsList {
	s.Value = &v
	return s
}

func (s *GetIssueByIdResponseBodyDataCfsList) SetName(v string) *GetIssueByIdResponseBodyDataCfsList {
	s.Name = &v
	return s
}

func (s *GetIssueByIdResponseBodyDataCfsList) SetId(v string) *GetIssueByIdResponseBodyDataCfsList {
	s.Id = &v
	return s
}

type GetIssueByIdResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIssueByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIssueByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIssueByIdResponse) GoString() string {
	return s.String()
}

func (s *GetIssueByIdResponse) SetHeaders(v map[string]*string) *GetIssueByIdResponse {
	s.Headers = v
	return s
}

func (s *GetIssueByIdResponse) SetBody(v *GetIssueByIdResponseBody) *GetIssueByIdResponse {
	s.Body = v
	return s
}

type GetJoinCodeRequest struct {
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetJoinCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJoinCodeRequest) GoString() string {
	return s.String()
}

func (s *GetJoinCodeRequest) SetCorpIdentifier(v string) *GetJoinCodeRequest {
	s.CorpIdentifier = &v
	return s
}

type GetJoinCodeResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetJoinCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJoinCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJoinCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetJoinCodeResponseBody) SetMessage(v string) *GetJoinCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetJoinCodeResponseBody) SetRequestId(v string) *GetJoinCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJoinCodeResponseBody) SetData(v *GetJoinCodeResponseBodyData) *GetJoinCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetJoinCodeResponseBody) SetCode(v int32) *GetJoinCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetJoinCodeResponseBody) SetSuccess(v bool) *GetJoinCodeResponseBody {
	s.Success = &v
	return s
}

type GetJoinCodeResponseBodyData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetJoinCodeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJoinCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJoinCodeResponseBodyData) SetCode(v string) *GetJoinCodeResponseBodyData {
	s.Code = &v
	return s
}

type GetJoinCodeResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJoinCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJoinCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJoinCodeResponse) GoString() string {
	return s.String()
}

func (s *GetJoinCodeResponse) SetHeaders(v map[string]*string) *GetJoinCodeResponse {
	s.Headers = v
	return s
}

func (s *GetJoinCodeResponse) SetBody(v *GetJoinCodeResponseBody) *GetJoinCodeResponse {
	s.Body = v
	return s
}

type GetProjectMembersRequest struct {
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	ProjectId      *int32  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StaffId        *string `json:"StaffId,omitempty" xml:"StaffId,omitempty"`
}

func (s GetProjectMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *GetProjectMembersRequest) SetCorpIdentifier(v string) *GetProjectMembersRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *GetProjectMembersRequest) SetProjectId(v int32) *GetProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *GetProjectMembersRequest) SetStaffId(v string) *GetProjectMembersRequest {
	s.StaffId = &v
	return s
}

type GetProjectMembersResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetProjectMembersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponseBody) SetMessage(v string) *GetProjectMembersResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetRequestId(v string) *GetProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetData(v []*GetProjectMembersResponseBodyData) *GetProjectMembersResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectMembersResponseBody) SetCode(v int32) *GetProjectMembersResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectMembersResponseBody) SetSuccess(v bool) *GetProjectMembersResponseBody {
	s.Success = &v
	return s
}

type GetProjectMembersResponseBodyData struct {
	Identifier *string                                   `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Users      []*GetProjectMembersResponseBodyDataUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	Name       *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *int32                                    `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectMembersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponseBodyData) SetIdentifier(v string) *GetProjectMembersResponseBodyData {
	s.Identifier = &v
	return s
}

func (s *GetProjectMembersResponseBodyData) SetUsers(v []*GetProjectMembersResponseBodyDataUsers) *GetProjectMembersResponseBodyData {
	s.Users = v
	return s
}

func (s *GetProjectMembersResponseBodyData) SetName(v string) *GetProjectMembersResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProjectMembersResponseBodyData) SetId(v int32) *GetProjectMembersResponseBodyData {
	s.Id = &v
	return s
}

type GetProjectMembersResponseBodyDataUsers struct {
	Avatar   *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	RealName *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	StaffId  *string `json:"StaffId,omitempty" xml:"StaffId,omitempty"`
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Id       *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectMembersResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s GetProjectMembersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *GetProjectMembersResponseBodyDataUsers) SetAvatar(v string) *GetProjectMembersResponseBodyDataUsers {
	s.Avatar = &v
	return s
}

func (s *GetProjectMembersResponseBodyDataUsers) SetEmail(v string) *GetProjectMembersResponseBodyDataUsers {
	s.Email = &v
	return s
}

func (s *GetProjectMembersResponseBodyDataUsers) SetRealName(v string) *GetProjectMembersResponseBodyDataUsers {
	s.RealName = &v
	return s
}

func (s *GetProjectMembersResponseBodyDataUsers) SetStaffId(v string) *GetProjectMembersResponseBodyDataUsers {
	s.StaffId = &v
	return s
}

func (s *GetProjectMembersResponseBodyDataUsers) SetNickName(v string) *GetProjectMembersResponseBodyDataUsers {
	s.NickName = &v
	return s
}

func (s *GetProjectMembersResponseBodyDataUsers) SetId(v int32) *GetProjectMembersResponseBodyDataUsers {
	s.Id = &v
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

type GetUserByAliyunPkRequest struct {
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetUserByAliyunPkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkRequest) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkRequest) SetPk(v string) *GetUserByAliyunPkRequest {
	s.Pk = &v
	return s
}

type GetUserByAliyunPkResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetUserByAliyunPkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserByAliyunPkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponseBody) SetRequestId(v string) *GetUserByAliyunPkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBody) SetData(v *GetUserByAliyunPkResponseBodyData) *GetUserByAliyunPkResponseBody {
	s.Data = v
	return s
}

func (s *GetUserByAliyunPkResponseBody) SetCode(v int32) *GetUserByAliyunPkResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserByAliyunPkResponseBody) SetSuccess(v bool) *GetUserByAliyunPkResponseBody {
	s.Success = &v
	return s
}

type GetUserByAliyunPkResponseBodyData struct {
	Uuid            *string                                          `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Guid            *string                                          `json:"Guid,omitempty" xml:"Guid,omitempty"`
	DingtalkUser    *GetUserByAliyunPkResponseBodyDataDingtalkUser   `json:"DingtalkUser,omitempty" xml:"DingtalkUser,omitempty" type:"Struct"`
	UserProfileDTO  *GetUserByAliyunPkResponseBodyDataUserProfileDTO `json:"UserProfileDTO,omitempty" xml:"UserProfileDTO,omitempty" type:"Struct"`
	NickName        *string                                          `json:"NickName,omitempty" xml:"NickName,omitempty"`
	MainAccountType *string                                          `json:"MainAccountType,omitempty" xml:"MainAccountType,omitempty"`
	AliyunUser      *GetUserByAliyunPkResponseBodyDataAliyunUser     `json:"AliyunUser,omitempty" xml:"AliyunUser,omitempty" type:"Struct"`
	IsValid         *bool                                            `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
	Id              *int32                                           `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserByAliyunPkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponseBodyData) SetUuid(v string) *GetUserByAliyunPkResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetGuid(v string) *GetUserByAliyunPkResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetDingtalkUser(v *GetUserByAliyunPkResponseBodyDataDingtalkUser) *GetUserByAliyunPkResponseBodyData {
	s.DingtalkUser = v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetUserProfileDTO(v *GetUserByAliyunPkResponseBodyDataUserProfileDTO) *GetUserByAliyunPkResponseBodyData {
	s.UserProfileDTO = v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetNickName(v string) *GetUserByAliyunPkResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetMainAccountType(v string) *GetUserByAliyunPkResponseBodyData {
	s.MainAccountType = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetAliyunUser(v *GetUserByAliyunPkResponseBodyDataAliyunUser) *GetUserByAliyunPkResponseBodyData {
	s.AliyunUser = v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetIsValid(v bool) *GetUserByAliyunPkResponseBodyData {
	s.IsValid = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyData) SetId(v int32) *GetUserByAliyunPkResponseBodyData {
	s.Id = &v
	return s
}

type GetUserByAliyunPkResponseBodyDataDingtalkUser struct {
	DingtalkUserId *int32  `json:"DingtalkUserId,omitempty" xml:"DingtalkUserId,omitempty"`
	DingId         *string `json:"DingId,omitempty" xml:"DingId,omitempty"`
	UnionId        *string `json:"UnionId,omitempty" xml:"UnionId,omitempty"`
	Nick           *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	CodeUserName   *string `json:"CodeUserName,omitempty" xml:"CodeUserName,omitempty"`
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserByAliyunPkResponseBodyDataDingtalkUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponseBodyDataDingtalkUser) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetDingtalkUserId(v int32) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.DingtalkUserId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetDingId(v string) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.DingId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetUnionId(v string) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.UnionId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetNick(v string) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.Nick = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetCodeUserName(v string) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.CodeUserName = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataDingtalkUser) SetId(v int32) *GetUserByAliyunPkResponseBodyDataDingtalkUser {
	s.Id = &v
	return s
}

type GetUserByAliyunPkResponseBodyDataUserProfileDTO struct {
	DataSource  *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Avatar      *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile      *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId      *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CreatedAt   *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	NickName    *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserByAliyunPkResponseBodyDataUserProfileDTO) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponseBodyDataUserProfileDTO) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetDataSource(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.DataSource = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetAvatar(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.Avatar = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetEmail(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.Email = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetMobile(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.Mobile = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetUserId(v int32) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.UserId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetCreatedAt(v int64) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.CreatedAt = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetEnglishName(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.EnglishName = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetNickName(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.NickName = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataUserProfileDTO) SetName(v string) *GetUserByAliyunPkResponseBodyDataUserProfileDTO {
	s.Name = &v
	return s
}

type GetUserByAliyunPkResponseBodyDataAliyunUser struct {
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HavanaId         *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	TaobaoNick       *string `json:"TaobaoNick,omitempty" xml:"TaobaoNick,omitempty"`
	Kp               *string `json:"Kp,omitempty" xml:"Kp,omitempty"`
	Realname         *string `json:"Realname,omitempty" xml:"Realname,omitempty"`
	NickName         *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	AccountStructure *int32  `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	AliyunId         *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Id               *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserByAliyunPkResponseBodyDataAliyunUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponseBodyDataAliyunUser) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetEmail(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.Email = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetHavanaId(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.HavanaId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetTaobaoNick(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.TaobaoNick = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetKp(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.Kp = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetRealname(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.Realname = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetNickName(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.NickName = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetAccountStructure(v int32) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.AccountStructure = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetAliyunId(v string) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.AliyunId = &v
	return s
}

func (s *GetUserByAliyunPkResponseBodyDataAliyunUser) SetId(v int32) *GetUserByAliyunPkResponseBodyDataAliyunUser {
	s.Id = &v
	return s
}

type GetUserByAliyunPkResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserByAliyunPkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserByAliyunPkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserByAliyunPkResponse) GoString() string {
	return s.String()
}

func (s *GetUserByAliyunPkResponse) SetHeaders(v map[string]*string) *GetUserByAliyunPkResponse {
	s.Headers = v
	return s
}

func (s *GetUserByAliyunPkResponse) SetBody(v *GetUserByAliyunPkResponseBody) *GetUserByAliyunPkResponse {
	s.Body = v
	return s
}

type GetWorkitemByIdRequest struct {
	Id             *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
}

func (s GetWorkitemByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemByIdRequest) GoString() string {
	return s.String()
}

func (s *GetWorkitemByIdRequest) SetId(v int32) *GetWorkitemByIdRequest {
	s.Id = &v
	return s
}

func (s *GetWorkitemByIdRequest) SetCorpIdentifier(v string) *GetWorkitemByIdRequest {
	s.CorpIdentifier = &v
	return s
}

type GetWorkitemByIdResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetWorkitemByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetWorkitemByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkitemByIdResponseBody) SetRequestId(v string) *GetWorkitemByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkitemByIdResponseBody) SetData(v *GetWorkitemByIdResponseBodyData) *GetWorkitemByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkitemByIdResponseBody) SetCode(v int32) *GetWorkitemByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkitemByIdResponseBody) SetSuccess(v string) *GetWorkitemByIdResponseBody {
	s.Success = &v
	return s
}

type GetWorkitemByIdResponseBodyData struct {
	SkipCollab            *bool                                     `json:"SkipCollab,omitempty" xml:"SkipCollab,omitempty"`
	SendWangwang          *bool                                     `json:"SendWangwang,omitempty" xml:"SendWangwang,omitempty"`
	VersionList           *string                                   `json:"VersionList,omitempty" xml:"VersionList,omitempty"`
	AkProjectId           *int32                                    `json:"AkProjectId,omitempty" xml:"AkProjectId,omitempty"`
	UserId                *int32                                    `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IdPath                *string                                   `json:"IdPath,omitempty" xml:"IdPath,omitempty"`
	IgnoreIntegrate       *bool                                     `json:"IgnoreIntegrate,omitempty" xml:"IgnoreIntegrate,omitempty"`
	AttachmentIds         *string                                   `json:"AttachmentIds,omitempty" xml:"AttachmentIds,omitempty"`
	CommitDate            *int64                                    `json:"CommitDate,omitempty" xml:"CommitDate,omitempty"`
	UpdateStatusAt        *int64                                    `json:"UpdateStatusAt,omitempty" xml:"UpdateStatusAt,omitempty"`
	Trackers              *string                                   `json:"Trackers,omitempty" xml:"Trackers,omitempty"`
	Subject               *string                                   `json:"Subject,omitempty" xml:"Subject,omitempty"`
	UserStaffId           *string                                   `json:"UserStaffId,omitempty" xml:"UserStaffId,omitempty"`
	AttachmentList        *string                                   `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty"`
	VerifierId            *int32                                    `json:"VerifierId,omitempty" xml:"VerifierId,omitempty"`
	IgnoreCheck           *bool                                     `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	ModuleList            *string                                   `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Attachmented          *bool                                     `json:"Attachmented,omitempty" xml:"Attachmented,omitempty"`
	SeriousLevelId        *int32                                    `json:"SeriousLevelId,omitempty" xml:"SeriousLevelId,omitempty"`
	TagIdList             *string                                   `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	Watched               *bool                                     `json:"Watched,omitempty" xml:"Watched,omitempty"`
	AssignedTo            *string                                   `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	AssignedToIds         *string                                   `json:"AssignedToIds,omitempty" xml:"AssignedToIds,omitempty"`
	PriorityId            *int32                                    `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	RecordChangeLog       *bool                                     `json:"RecordChangeLog,omitempty" xml:"RecordChangeLog,omitempty"`
	UpdatedAt             *int64                                    `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	TemplateId            *int32                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Verifier              *string                                   `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	Status                *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	RelatedUserIds        *string                                   `json:"RelatedUserIds,omitempty" xml:"RelatedUserIds,omitempty"`
	RelatedAKProjectIds   *string                                   `json:"RelatedAKProjectIds,omitempty" xml:"RelatedAKProjectIds,omitempty"`
	RelatedAKProjectGuids *string                                   `json:"RelatedAKProjectGuids,omitempty" xml:"RelatedAKProjectGuids,omitempty"`
	CreatedAt             *int64                                    `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AssignedToIdList      *string                                   `json:"AssignedToIdList,omitempty" xml:"AssignedToIdList,omitempty"`
	CfsList               []*GetWorkitemByIdResponseBodyDataCfsList `json:"CfsList,omitempty" xml:"CfsList,omitempty" type:"Repeated"`
	Priority              *string                                   `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StatusStage           *int32                                    `json:"StatusStage,omitempty" xml:"StatusStage,omitempty"`
	AssignedToStaffId     *string                                   `json:"AssignedToStaffId,omitempty" xml:"AssignedToStaffId,omitempty"`
	ParentId              *int32                                    `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Description           *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	LogicalStatus         *string                                   `json:"LogicalStatus,omitempty" xml:"LogicalStatus,omitempty"`
	VersionIds            *string                                   `json:"VersionIds,omitempty" xml:"VersionIds,omitempty"`
	SeriousLevel          *string                                   `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	StatusId              *int32                                    `json:"StatusId,omitempty" xml:"StatusId,omitempty"`
	Stamp                 *string                                   `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	User                  *string                                   `json:"User,omitempty" xml:"User,omitempty"`
	TrackerIds            *string                                   `json:"TrackerIds,omitempty" xml:"TrackerIds,omitempty"`
	AssignedToId          *int32                                    `json:"AssignedToId,omitempty" xml:"AssignedToId,omitempty"`
	ChangeLogList         *string                                   `json:"ChangeLogList,omitempty" xml:"ChangeLogList,omitempty"`
	IssueTypeId           *int32                                    `json:"IssueTypeId,omitempty" xml:"IssueTypeId,omitempty"`
	RegionId              *int32                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Guid                  *string                                   `json:"Guid,omitempty" xml:"Guid,omitempty"`
	AssignedToMaps        *string                                   `json:"AssignedToMaps,omitempty" xml:"AssignedToMaps,omitempty"`
	BlackListNotice       *string                                   `json:"BlackListNotice,omitempty" xml:"BlackListNotice,omitempty"`
	ModuleUpdated         *bool                                     `json:"ModuleUpdated,omitempty" xml:"ModuleUpdated,omitempty"`
	ProjectIds            *string                                   `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	CommentList           *string                                   `json:"CommentList,omitempty" xml:"CommentList,omitempty"`
	VerifierStaffId       *string                                   `json:"VerifierStaffId,omitempty" xml:"VerifierStaffId,omitempty"`
	ModuleIds             *string                                   `json:"ModuleIds,omitempty" xml:"ModuleIds,omitempty"`
	Id                    *int32                                    `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetWorkitemByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkitemByIdResponseBodyData) SetSkipCollab(v bool) *GetWorkitemByIdResponseBodyData {
	s.SkipCollab = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetSendWangwang(v bool) *GetWorkitemByIdResponseBodyData {
	s.SendWangwang = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetVersionList(v string) *GetWorkitemByIdResponseBodyData {
	s.VersionList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAkProjectId(v int32) *GetWorkitemByIdResponseBodyData {
	s.AkProjectId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetUserId(v int32) *GetWorkitemByIdResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetIdPath(v string) *GetWorkitemByIdResponseBodyData {
	s.IdPath = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetIgnoreIntegrate(v bool) *GetWorkitemByIdResponseBodyData {
	s.IgnoreIntegrate = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAttachmentIds(v string) *GetWorkitemByIdResponseBodyData {
	s.AttachmentIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetCommitDate(v int64) *GetWorkitemByIdResponseBodyData {
	s.CommitDate = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetUpdateStatusAt(v int64) *GetWorkitemByIdResponseBodyData {
	s.UpdateStatusAt = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetTrackers(v string) *GetWorkitemByIdResponseBodyData {
	s.Trackers = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetSubject(v string) *GetWorkitemByIdResponseBodyData {
	s.Subject = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetUserStaffId(v string) *GetWorkitemByIdResponseBodyData {
	s.UserStaffId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAttachmentList(v string) *GetWorkitemByIdResponseBodyData {
	s.AttachmentList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetVerifierId(v int32) *GetWorkitemByIdResponseBodyData {
	s.VerifierId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetIgnoreCheck(v bool) *GetWorkitemByIdResponseBodyData {
	s.IgnoreCheck = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetModuleList(v string) *GetWorkitemByIdResponseBodyData {
	s.ModuleList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAttachmented(v bool) *GetWorkitemByIdResponseBodyData {
	s.Attachmented = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetSeriousLevelId(v int32) *GetWorkitemByIdResponseBodyData {
	s.SeriousLevelId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetTagIdList(v string) *GetWorkitemByIdResponseBodyData {
	s.TagIdList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetWatched(v bool) *GetWorkitemByIdResponseBodyData {
	s.Watched = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedTo(v string) *GetWorkitemByIdResponseBodyData {
	s.AssignedTo = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedToIds(v string) *GetWorkitemByIdResponseBodyData {
	s.AssignedToIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetPriorityId(v int32) *GetWorkitemByIdResponseBodyData {
	s.PriorityId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetRecordChangeLog(v bool) *GetWorkitemByIdResponseBodyData {
	s.RecordChangeLog = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetUpdatedAt(v int64) *GetWorkitemByIdResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetTemplateId(v int32) *GetWorkitemByIdResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetVerifier(v string) *GetWorkitemByIdResponseBodyData {
	s.Verifier = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetStatus(v string) *GetWorkitemByIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetRelatedUserIds(v string) *GetWorkitemByIdResponseBodyData {
	s.RelatedUserIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetRelatedAKProjectIds(v string) *GetWorkitemByIdResponseBodyData {
	s.RelatedAKProjectIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetRelatedAKProjectGuids(v string) *GetWorkitemByIdResponseBodyData {
	s.RelatedAKProjectGuids = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetCreatedAt(v int64) *GetWorkitemByIdResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedToIdList(v string) *GetWorkitemByIdResponseBodyData {
	s.AssignedToIdList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetCfsList(v []*GetWorkitemByIdResponseBodyDataCfsList) *GetWorkitemByIdResponseBodyData {
	s.CfsList = v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetPriority(v string) *GetWorkitemByIdResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetStatusStage(v int32) *GetWorkitemByIdResponseBodyData {
	s.StatusStage = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedToStaffId(v string) *GetWorkitemByIdResponseBodyData {
	s.AssignedToStaffId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetParentId(v int32) *GetWorkitemByIdResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetDescription(v string) *GetWorkitemByIdResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetLogicalStatus(v string) *GetWorkitemByIdResponseBodyData {
	s.LogicalStatus = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetVersionIds(v string) *GetWorkitemByIdResponseBodyData {
	s.VersionIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetSeriousLevel(v string) *GetWorkitemByIdResponseBodyData {
	s.SeriousLevel = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetStatusId(v int32) *GetWorkitemByIdResponseBodyData {
	s.StatusId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetStamp(v string) *GetWorkitemByIdResponseBodyData {
	s.Stamp = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetUser(v string) *GetWorkitemByIdResponseBodyData {
	s.User = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetTrackerIds(v string) *GetWorkitemByIdResponseBodyData {
	s.TrackerIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedToId(v int32) *GetWorkitemByIdResponseBodyData {
	s.AssignedToId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetChangeLogList(v string) *GetWorkitemByIdResponseBodyData {
	s.ChangeLogList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetIssueTypeId(v int32) *GetWorkitemByIdResponseBodyData {
	s.IssueTypeId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetRegionId(v int32) *GetWorkitemByIdResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetGuid(v string) *GetWorkitemByIdResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetAssignedToMaps(v string) *GetWorkitemByIdResponseBodyData {
	s.AssignedToMaps = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetBlackListNotice(v string) *GetWorkitemByIdResponseBodyData {
	s.BlackListNotice = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetModuleUpdated(v bool) *GetWorkitemByIdResponseBodyData {
	s.ModuleUpdated = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetProjectIds(v string) *GetWorkitemByIdResponseBodyData {
	s.ProjectIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetCommentList(v string) *GetWorkitemByIdResponseBodyData {
	s.CommentList = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetVerifierStaffId(v string) *GetWorkitemByIdResponseBodyData {
	s.VerifierStaffId = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetModuleIds(v string) *GetWorkitemByIdResponseBodyData {
	s.ModuleIds = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyData) SetId(v int32) *GetWorkitemByIdResponseBodyData {
	s.Id = &v
	return s
}

type GetWorkitemByIdResponseBodyDataCfsList struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetWorkitemByIdResponseBodyDataCfsList) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemByIdResponseBodyDataCfsList) GoString() string {
	return s.String()
}

func (s *GetWorkitemByIdResponseBodyDataCfsList) SetType(v string) *GetWorkitemByIdResponseBodyDataCfsList {
	s.Type = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyDataCfsList) SetValue(v string) *GetWorkitemByIdResponseBodyDataCfsList {
	s.Value = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyDataCfsList) SetName(v string) *GetWorkitemByIdResponseBodyDataCfsList {
	s.Name = &v
	return s
}

func (s *GetWorkitemByIdResponseBodyDataCfsList) SetId(v string) *GetWorkitemByIdResponseBodyDataCfsList {
	s.Id = &v
	return s
}

type GetWorkitemByIdResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkitemByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkitemByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkitemByIdResponse) GoString() string {
	return s.String()
}

func (s *GetWorkitemByIdResponse) SetHeaders(v map[string]*string) *GetWorkitemByIdResponse {
	s.Headers = v
	return s
}

func (s *GetWorkitemByIdResponse) SetBody(v *GetWorkitemByIdResponseBody) *GetWorkitemByIdResponse {
	s.Body = v
	return s
}

type JoinCompanyRequest struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s JoinCompanyRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinCompanyRequest) GoString() string {
	return s.String()
}

func (s *JoinCompanyRequest) SetCode(v string) *JoinCompanyRequest {
	s.Code = &v
	return s
}

type JoinCompanyResponseBody struct {
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *JoinCompanyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinCompanyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinCompanyResponseBody) GoString() string {
	return s.String()
}

func (s *JoinCompanyResponseBody) SetMessage(v string) *JoinCompanyResponseBody {
	s.Message = &v
	return s
}

func (s *JoinCompanyResponseBody) SetRequestId(v string) *JoinCompanyResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinCompanyResponseBody) SetData(v *JoinCompanyResponseBodyData) *JoinCompanyResponseBody {
	s.Data = v
	return s
}

func (s *JoinCompanyResponseBody) SetCode(v int32) *JoinCompanyResponseBody {
	s.Code = &v
	return s
}

func (s *JoinCompanyResponseBody) SetSuccess(v bool) *JoinCompanyResponseBody {
	s.Success = &v
	return s
}

type JoinCompanyResponseBodyData struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s JoinCompanyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s JoinCompanyResponseBodyData) GoString() string {
	return s.String()
}

func (s *JoinCompanyResponseBodyData) SetApplicationId(v string) *JoinCompanyResponseBodyData {
	s.ApplicationId = &v
	return s
}

type JoinCompanyResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinCompanyResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinCompanyResponse) GoString() string {
	return s.String()
}

func (s *JoinCompanyResponse) SetHeaders(v map[string]*string) *JoinCompanyResponse {
	s.Headers = v
	return s
}

func (s *JoinCompanyResponse) SetBody(v *JoinCompanyResponseBody) *JoinCompanyResponse {
	s.Body = v
	return s
}

type SearchProjectsByRegionRequest struct {
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ToPage         *int32  `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s SearchProjectsByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsByRegionRequest) GoString() string {
	return s.String()
}

func (s *SearchProjectsByRegionRequest) SetCorpIdentifier(v string) *SearchProjectsByRegionRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *SearchProjectsByRegionRequest) SetRegion(v string) *SearchProjectsByRegionRequest {
	s.Region = &v
	return s
}

func (s *SearchProjectsByRegionRequest) SetStatus(v string) *SearchProjectsByRegionRequest {
	s.Status = &v
	return s
}

func (s *SearchProjectsByRegionRequest) SetToPage(v int32) *SearchProjectsByRegionRequest {
	s.ToPage = &v
	return s
}

func (s *SearchProjectsByRegionRequest) SetPageSize(v int32) *SearchProjectsByRegionRequest {
	s.PageSize = &v
	return s
}

type SearchProjectsByRegionResponseBody struct {
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*SearchProjectsByRegionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchProjectsByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *SearchProjectsByRegionResponseBody) SetMessage(v string) *SearchProjectsByRegionResponseBody {
	s.Message = &v
	return s
}

func (s *SearchProjectsByRegionResponseBody) SetRequestId(v string) *SearchProjectsByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchProjectsByRegionResponseBody) SetData(v []*SearchProjectsByRegionResponseBodyData) *SearchProjectsByRegionResponseBody {
	s.Data = v
	return s
}

func (s *SearchProjectsByRegionResponseBody) SetCode(v int32) *SearchProjectsByRegionResponseBody {
	s.Code = &v
	return s
}

func (s *SearchProjectsByRegionResponseBody) SetSuccess(v bool) *SearchProjectsByRegionResponseBody {
	s.Success = &v
	return s
}

type SearchProjectsByRegionResponseBodyData struct {
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	AoneType       *string   `json:"AoneType,omitempty" xml:"AoneType,omitempty"`
	Stamp          *string   `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	FullName       *string   `json:"FullName,omitempty" xml:"FullName,omitempty"`
	Icons          []*string `json:"Icons,omitempty" xml:"Icons,omitempty" type:"Repeated"`
	IdPath         *string   `json:"IdPath,omitempty" xml:"IdPath,omitempty"`
	Mode           *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	AoneId         *int32    `json:"AoneId,omitempty" xml:"AoneId,omitempty"`
	ParentId       *int32    `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	CustomFieldMap []*string `json:"CustomFieldMap,omitempty" xml:"CustomFieldMap,omitempty" type:"Repeated"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *int32    `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchProjectsByRegionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsByRegionResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchProjectsByRegionResponseBodyData) SetStatus(v string) *SearchProjectsByRegionResponseBodyData {
	s.Status = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetType(v string) *SearchProjectsByRegionResponseBodyData {
	s.Type = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetAoneType(v string) *SearchProjectsByRegionResponseBodyData {
	s.AoneType = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetStamp(v string) *SearchProjectsByRegionResponseBodyData {
	s.Stamp = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetFullName(v string) *SearchProjectsByRegionResponseBodyData {
	s.FullName = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetIcons(v []*string) *SearchProjectsByRegionResponseBodyData {
	s.Icons = v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetIdPath(v string) *SearchProjectsByRegionResponseBodyData {
	s.IdPath = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetMode(v string) *SearchProjectsByRegionResponseBodyData {
	s.Mode = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetAoneId(v int32) *SearchProjectsByRegionResponseBodyData {
	s.AoneId = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetParentId(v int32) *SearchProjectsByRegionResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetDescription(v string) *SearchProjectsByRegionResponseBodyData {
	s.Description = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetCustomFieldMap(v []*string) *SearchProjectsByRegionResponseBodyData {
	s.CustomFieldMap = v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetName(v string) *SearchProjectsByRegionResponseBodyData {
	s.Name = &v
	return s
}

func (s *SearchProjectsByRegionResponseBodyData) SetId(v int32) *SearchProjectsByRegionResponseBodyData {
	s.Id = &v
	return s
}

type SearchProjectsByRegionResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchProjectsByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchProjectsByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchProjectsByRegionResponse) GoString() string {
	return s.String()
}

func (s *SearchProjectsByRegionResponse) SetHeaders(v map[string]*string) *SearchProjectsByRegionResponse {
	s.Headers = v
	return s
}

func (s *SearchProjectsByRegionResponse) SetBody(v *SearchProjectsByRegionResponseBody) *SearchProjectsByRegionResponse {
	s.Body = v
	return s
}

type SearchTestCaseRequest struct {
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpIdentifier  *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	AkProjectId     *string `json:"AkProjectId,omitempty" xml:"AkProjectId,omitempty"`
	CaseTag         *string `json:"CaseTag,omitempty" xml:"CaseTag,omitempty"`
	PageNum         *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	CreateDateStart *string `json:"CreateDateStart,omitempty" xml:"CreateDateStart,omitempty"`
	CreateDateEnd   *string `json:"CreateDateEnd,omitempty" xml:"CreateDateEnd,omitempty"`
	UpdateDateStart *string `json:"UpdateDateStart,omitempty" xml:"UpdateDateStart,omitempty"`
	UpdateDateEnd   *string `json:"UpdateDateEnd,omitempty" xml:"UpdateDateEnd,omitempty"`
}

func (s SearchTestCaseRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTestCaseRequest) GoString() string {
	return s.String()
}

func (s *SearchTestCaseRequest) SetPageSize(v int32) *SearchTestCaseRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTestCaseRequest) SetCorpIdentifier(v string) *SearchTestCaseRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *SearchTestCaseRequest) SetAkProjectId(v string) *SearchTestCaseRequest {
	s.AkProjectId = &v
	return s
}

func (s *SearchTestCaseRequest) SetCaseTag(v string) *SearchTestCaseRequest {
	s.CaseTag = &v
	return s
}

func (s *SearchTestCaseRequest) SetPageNum(v string) *SearchTestCaseRequest {
	s.PageNum = &v
	return s
}

func (s *SearchTestCaseRequest) SetCreateDateStart(v string) *SearchTestCaseRequest {
	s.CreateDateStart = &v
	return s
}

func (s *SearchTestCaseRequest) SetCreateDateEnd(v string) *SearchTestCaseRequest {
	s.CreateDateEnd = &v
	return s
}

func (s *SearchTestCaseRequest) SetUpdateDateStart(v string) *SearchTestCaseRequest {
	s.UpdateDateStart = &v
	return s
}

func (s *SearchTestCaseRequest) SetUpdateDateEnd(v string) *SearchTestCaseRequest {
	s.UpdateDateEnd = &v
	return s
}

type SearchTestCaseResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchTestCaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTestCaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTestCaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTestCaseResponseBody) SetMessage(v string) *SearchTestCaseResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTestCaseResponseBody) SetRequestId(v string) *SearchTestCaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTestCaseResponseBody) SetData(v *SearchTestCaseResponseBodyData) *SearchTestCaseResponseBody {
	s.Data = v
	return s
}

func (s *SearchTestCaseResponseBody) SetCode(v int32) *SearchTestCaseResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTestCaseResponseBody) SetSuccess(v bool) *SearchTestCaseResponseBody {
	s.Success = &v
	return s
}

type SearchTestCaseResponseBodyData struct {
	PageNum    *string `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	Cases      *string `json:"Cases,omitempty" xml:"Cases,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageTotal  *string `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchTestCaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTestCaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTestCaseResponseBodyData) SetPageNum(v string) *SearchTestCaseResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *SearchTestCaseResponseBodyData) SetCases(v string) *SearchTestCaseResponseBodyData {
	s.Cases = &v
	return s
}

func (s *SearchTestCaseResponseBodyData) SetPageSize(v string) *SearchTestCaseResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SearchTestCaseResponseBodyData) SetPageTotal(v string) *SearchTestCaseResponseBodyData {
	s.PageTotal = &v
	return s
}

func (s *SearchTestCaseResponseBodyData) SetTotalCount(v string) *SearchTestCaseResponseBodyData {
	s.TotalCount = &v
	return s
}

type SearchTestCaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchTestCaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTestCaseResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTestCaseResponse) GoString() string {
	return s.String()
}

func (s *SearchTestCaseResponse) SetHeaders(v map[string]*string) *SearchTestCaseResponse {
	s.Headers = v
	return s
}

func (s *SearchTestCaseResponse) SetBody(v *SearchTestCaseResponseBody) *SearchTestCaseResponse {
	s.Body = v
	return s
}

type SearchWorkitemRequest struct {
	Stamp          *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	AKProjectId    *int32  `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	ToPage         *int32  `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	SprintId       *int32  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	CreatedAtStart *string `json:"CreatedAtStart,omitempty" xml:"CreatedAtStart,omitempty"`
	CreatedAtEnd   *string `json:"CreatedAtEnd,omitempty" xml:"CreatedAtEnd,omitempty"`
}

func (s SearchWorkitemRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemRequest) GoString() string {
	return s.String()
}

func (s *SearchWorkitemRequest) SetStamp(v string) *SearchWorkitemRequest {
	s.Stamp = &v
	return s
}

func (s *SearchWorkitemRequest) SetAKProjectId(v int32) *SearchWorkitemRequest {
	s.AKProjectId = &v
	return s
}

func (s *SearchWorkitemRequest) SetToPage(v int32) *SearchWorkitemRequest {
	s.ToPage = &v
	return s
}

func (s *SearchWorkitemRequest) SetPageSize(v int32) *SearchWorkitemRequest {
	s.PageSize = &v
	return s
}

func (s *SearchWorkitemRequest) SetCorpIdentifier(v string) *SearchWorkitemRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *SearchWorkitemRequest) SetSprintId(v int32) *SearchWorkitemRequest {
	s.SprintId = &v
	return s
}

func (s *SearchWorkitemRequest) SetCreatedAtStart(v string) *SearchWorkitemRequest {
	s.CreatedAtStart = &v
	return s
}

func (s *SearchWorkitemRequest) SetCreatedAtEnd(v string) *SearchWorkitemRequest {
	s.CreatedAtEnd = &v
	return s
}

type SearchWorkitemResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ToPage     *int32                            `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalPages *int32                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Data       []*SearchWorkitemResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchWorkitemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWorkitemResponseBody) SetTotalCount(v int32) *SearchWorkitemResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetToPage(v int32) *SearchWorkitemResponseBody {
	s.ToPage = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetRequestId(v string) *SearchWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetPageSize(v int32) *SearchWorkitemResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetTotalPages(v int32) *SearchWorkitemResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetData(v []*SearchWorkitemResponseBodyData) *SearchWorkitemResponseBody {
	s.Data = v
	return s
}

func (s *SearchWorkitemResponseBody) SetCode(v int32) *SearchWorkitemResponseBody {
	s.Code = &v
	return s
}

func (s *SearchWorkitemResponseBody) SetSuccess(v bool) *SearchWorkitemResponseBody {
	s.Success = &v
	return s
}

type SearchWorkitemResponseBodyData struct {
	FixedUserId           *int32  `json:"FixedUserId,omitempty" xml:"FixedUserId,omitempty"`
	SendWangwang          *bool   `json:"SendWangwang,omitempty" xml:"SendWangwang,omitempty"`
	SkipCollab            *bool   `json:"SkipCollab,omitempty" xml:"SkipCollab,omitempty"`
	VersionList           *string `json:"VersionList,omitempty" xml:"VersionList,omitempty"`
	AkProjectId           *int32  `json:"AkProjectId,omitempty" xml:"AkProjectId,omitempty"`
	ProjectId             *int32  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserId                *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionId             *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	IdPath                *string `json:"IdPath,omitempty" xml:"IdPath,omitempty"`
	IgnoreIntegrate       *bool   `json:"IgnoreIntegrate,omitempty" xml:"IgnoreIntegrate,omitempty"`
	AttachmentIds         *string `json:"AttachmentIds,omitempty" xml:"AttachmentIds,omitempty"`
	CommitDate            *int64  `json:"CommitDate,omitempty" xml:"CommitDate,omitempty"`
	AkVersionIds          *string `json:"AkVersionIds,omitempty" xml:"AkVersionIds,omitempty"`
	Region                *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Trackers              *string `json:"Trackers,omitempty" xml:"Trackers,omitempty"`
	Subject               *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Solution              *int32  `json:"Solution,omitempty" xml:"Solution,omitempty"`
	AttachmentList        *string `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty"`
	UserStaffId           *string `json:"UserStaffId,omitempty" xml:"UserStaffId,omitempty"`
	IgnoreCheck           *bool   `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	AkPaths               *string `json:"AkPaths,omitempty" xml:"AkPaths,omitempty"`
	VerifierId            *int32  `json:"VerifierId,omitempty" xml:"VerifierId,omitempty"`
	ModuleList            *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Attachmented          *bool   `json:"Attachmented,omitempty" xml:"Attachmented,omitempty"`
	SourceId              *int32  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SeriousLevelId        *int32  `json:"SeriousLevelId,omitempty" xml:"SeriousLevelId,omitempty"`
	TagIdList             *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	Watched               *bool   `json:"Watched,omitempty" xml:"Watched,omitempty"`
	AssignedTo            *string `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	AssignedToIds         *string `json:"AssignedToIds,omitempty" xml:"AssignedToIds,omitempty"`
	ClosedDuration        *int32  `json:"ClosedDuration,omitempty" xml:"ClosedDuration,omitempty"`
	PriorityId            *int32  `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	RecordChangeLog       *bool   `json:"RecordChangeLog,omitempty" xml:"RecordChangeLog,omitempty"`
	IssueRelations        *string `json:"IssueRelations,omitempty" xml:"IssueRelations,omitempty"`
	UpdatedAt             *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	TemplateId            *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Verifier              *string `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RelatedUserIds        *string `json:"RelatedUserIds,omitempty" xml:"RelatedUserIds,omitempty"`
	RelatedAKProjectGuids *string `json:"RelatedAKProjectGuids,omitempty" xml:"RelatedAKProjectGuids,omitempty"`
	RelatedAKProjectIds   *string `json:"RelatedAKProjectIds,omitempty" xml:"RelatedAKProjectIds,omitempty"`
	WatcherIdList         *string `json:"WatcherIdList,omitempty" xml:"WatcherIdList,omitempty"`
	CreatedAt             *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AssignedToIdList      *string `json:"AssignedToIdList,omitempty" xml:"AssignedToIdList,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StatusStage           *int32  `json:"StatusStage,omitempty" xml:"StatusStage,omitempty"`
	AssignedToStaffId     *string `json:"AssignedToStaffId,omitempty" xml:"AssignedToStaffId,omitempty"`
	ParentId              *int32  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	LogicalStatus         *string `json:"LogicalStatus,omitempty" xml:"LogicalStatus,omitempty"`
	VersionIds            *string `json:"VersionIds,omitempty" xml:"VersionIds,omitempty"`
	Scope                 *int32  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SeriousLevel          *string `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	StatusId              *int32  `json:"StatusId,omitempty" xml:"StatusId,omitempty"`
	Stamp                 *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	User                  *string `json:"User,omitempty" xml:"User,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	AssignedToId          *int32  `json:"AssignedToId,omitempty" xml:"AssignedToId,omitempty"`
	TrackerIds            *string `json:"TrackerIds,omitempty" xml:"TrackerIds,omitempty"`
	IssueTypeId           *int32  `json:"IssueTypeId,omitempty" xml:"IssueTypeId,omitempty"`
	ChangeLogList         *string `json:"ChangeLogList,omitempty" xml:"ChangeLogList,omitempty"`
	FixedDuration         *int32  `json:"FixedDuration,omitempty" xml:"FixedDuration,omitempty"`
	ScopeUserIds          *string `json:"ScopeUserIds,omitempty" xml:"ScopeUserIds,omitempty"`
	DevDuration           *int32  `json:"DevDuration,omitempty" xml:"DevDuration,omitempty"`
	LinePath              *string `json:"LinePath,omitempty" xml:"LinePath,omitempty"`
	BlackListNotice       *string `json:"BlackListNotice,omitempty" xml:"BlackListNotice,omitempty"`
	ProjectIds            *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	TestsuiteId           *int32  `json:"TestsuiteId,omitempty" xml:"TestsuiteId,omitempty"`
	ModuleUpdated         *bool   `json:"ModuleUpdated,omitempty" xml:"ModuleUpdated,omitempty"`
	CommentList           *string `json:"CommentList,omitempty" xml:"CommentList,omitempty"`
	SprintId              *int32  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	RespondDuration       *int32  `json:"RespondDuration,omitempty" xml:"RespondDuration,omitempty"`
	ModuleIds             *string `json:"ModuleIds,omitempty" xml:"ModuleIds,omitempty"`
	VerifierStaffId       *string `json:"VerifierStaffId,omitempty" xml:"VerifierStaffId,omitempty"`
	Id                    *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchWorkitemResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchWorkitemResponseBodyData) SetFixedUserId(v int32) *SearchWorkitemResponseBodyData {
	s.FixedUserId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSendWangwang(v bool) *SearchWorkitemResponseBodyData {
	s.SendWangwang = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSkipCollab(v bool) *SearchWorkitemResponseBodyData {
	s.SkipCollab = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVersionList(v string) *SearchWorkitemResponseBodyData {
	s.VersionList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAkProjectId(v int32) *SearchWorkitemResponseBodyData {
	s.AkProjectId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetProjectId(v int32) *SearchWorkitemResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetUserId(v int32) *SearchWorkitemResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVersionId(v int32) *SearchWorkitemResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetIdPath(v string) *SearchWorkitemResponseBodyData {
	s.IdPath = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetIgnoreIntegrate(v bool) *SearchWorkitemResponseBodyData {
	s.IgnoreIntegrate = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAttachmentIds(v string) *SearchWorkitemResponseBodyData {
	s.AttachmentIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetCommitDate(v int64) *SearchWorkitemResponseBodyData {
	s.CommitDate = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAkVersionIds(v string) *SearchWorkitemResponseBodyData {
	s.AkVersionIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRegion(v string) *SearchWorkitemResponseBodyData {
	s.Region = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetTrackers(v string) *SearchWorkitemResponseBodyData {
	s.Trackers = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSubject(v string) *SearchWorkitemResponseBodyData {
	s.Subject = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSolution(v int32) *SearchWorkitemResponseBodyData {
	s.Solution = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAttachmentList(v string) *SearchWorkitemResponseBodyData {
	s.AttachmentList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetUserStaffId(v string) *SearchWorkitemResponseBodyData {
	s.UserStaffId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetIgnoreCheck(v bool) *SearchWorkitemResponseBodyData {
	s.IgnoreCheck = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAkPaths(v string) *SearchWorkitemResponseBodyData {
	s.AkPaths = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVerifierId(v int32) *SearchWorkitemResponseBodyData {
	s.VerifierId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetModuleList(v string) *SearchWorkitemResponseBodyData {
	s.ModuleList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAttachmented(v bool) *SearchWorkitemResponseBodyData {
	s.Attachmented = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSourceId(v int32) *SearchWorkitemResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSeriousLevelId(v int32) *SearchWorkitemResponseBodyData {
	s.SeriousLevelId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetTagIdList(v string) *SearchWorkitemResponseBodyData {
	s.TagIdList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetWatched(v bool) *SearchWorkitemResponseBodyData {
	s.Watched = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAssignedTo(v string) *SearchWorkitemResponseBodyData {
	s.AssignedTo = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAssignedToIds(v string) *SearchWorkitemResponseBodyData {
	s.AssignedToIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetClosedDuration(v int32) *SearchWorkitemResponseBodyData {
	s.ClosedDuration = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetPriorityId(v int32) *SearchWorkitemResponseBodyData {
	s.PriorityId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRecordChangeLog(v bool) *SearchWorkitemResponseBodyData {
	s.RecordChangeLog = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetIssueRelations(v string) *SearchWorkitemResponseBodyData {
	s.IssueRelations = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetUpdatedAt(v int64) *SearchWorkitemResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetTemplateId(v int32) *SearchWorkitemResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVerifier(v string) *SearchWorkitemResponseBodyData {
	s.Verifier = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetStatus(v string) *SearchWorkitemResponseBodyData {
	s.Status = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRelatedUserIds(v string) *SearchWorkitemResponseBodyData {
	s.RelatedUserIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRelatedAKProjectGuids(v string) *SearchWorkitemResponseBodyData {
	s.RelatedAKProjectGuids = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRelatedAKProjectIds(v string) *SearchWorkitemResponseBodyData {
	s.RelatedAKProjectIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetWatcherIdList(v string) *SearchWorkitemResponseBodyData {
	s.WatcherIdList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetCreatedAt(v int64) *SearchWorkitemResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAssignedToIdList(v string) *SearchWorkitemResponseBodyData {
	s.AssignedToIdList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetPriority(v string) *SearchWorkitemResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetStatusStage(v int32) *SearchWorkitemResponseBodyData {
	s.StatusStage = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAssignedToStaffId(v string) *SearchWorkitemResponseBodyData {
	s.AssignedToStaffId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetParentId(v int32) *SearchWorkitemResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetLogicalStatus(v string) *SearchWorkitemResponseBodyData {
	s.LogicalStatus = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVersionIds(v string) *SearchWorkitemResponseBodyData {
	s.VersionIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetScope(v int32) *SearchWorkitemResponseBodyData {
	s.Scope = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSeriousLevel(v string) *SearchWorkitemResponseBodyData {
	s.SeriousLevel = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetStatusId(v int32) *SearchWorkitemResponseBodyData {
	s.StatusId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetStamp(v string) *SearchWorkitemResponseBodyData {
	s.Stamp = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetUser(v string) *SearchWorkitemResponseBodyData {
	s.User = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSource(v string) *SearchWorkitemResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetAssignedToId(v int32) *SearchWorkitemResponseBodyData {
	s.AssignedToId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetTrackerIds(v string) *SearchWorkitemResponseBodyData {
	s.TrackerIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetIssueTypeId(v int32) *SearchWorkitemResponseBodyData {
	s.IssueTypeId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetChangeLogList(v string) *SearchWorkitemResponseBodyData {
	s.ChangeLogList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetFixedDuration(v int32) *SearchWorkitemResponseBodyData {
	s.FixedDuration = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetScopeUserIds(v string) *SearchWorkitemResponseBodyData {
	s.ScopeUserIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetDevDuration(v int32) *SearchWorkitemResponseBodyData {
	s.DevDuration = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetLinePath(v string) *SearchWorkitemResponseBodyData {
	s.LinePath = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetBlackListNotice(v string) *SearchWorkitemResponseBodyData {
	s.BlackListNotice = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetProjectIds(v string) *SearchWorkitemResponseBodyData {
	s.ProjectIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetTestsuiteId(v int32) *SearchWorkitemResponseBodyData {
	s.TestsuiteId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetModuleUpdated(v bool) *SearchWorkitemResponseBodyData {
	s.ModuleUpdated = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetCommentList(v string) *SearchWorkitemResponseBodyData {
	s.CommentList = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetSprintId(v int32) *SearchWorkitemResponseBodyData {
	s.SprintId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetRespondDuration(v int32) *SearchWorkitemResponseBodyData {
	s.RespondDuration = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetModuleIds(v string) *SearchWorkitemResponseBodyData {
	s.ModuleIds = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetVerifierStaffId(v string) *SearchWorkitemResponseBodyData {
	s.VerifierStaffId = &v
	return s
}

func (s *SearchWorkitemResponseBodyData) SetId(v int32) *SearchWorkitemResponseBodyData {
	s.Id = &v
	return s
}

type SearchWorkitemResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchWorkitemResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemResponse) GoString() string {
	return s.String()
}

func (s *SearchWorkitemResponse) SetHeaders(v map[string]*string) *SearchWorkitemResponse {
	s.Headers = v
	return s
}

func (s *SearchWorkitemResponse) SetBody(v *SearchWorkitemResponseBody) *SearchWorkitemResponse {
	s.Body = v
	return s
}

type SearchWorkitemWithTotalCountRequest struct {
	Stamp          *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	AKProjectId    *int32  `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	ToPage         *int32  `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	SprintId       *int32  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
}

func (s SearchWorkitemWithTotalCountRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemWithTotalCountRequest) GoString() string {
	return s.String()
}

func (s *SearchWorkitemWithTotalCountRequest) SetStamp(v string) *SearchWorkitemWithTotalCountRequest {
	s.Stamp = &v
	return s
}

func (s *SearchWorkitemWithTotalCountRequest) SetAKProjectId(v int32) *SearchWorkitemWithTotalCountRequest {
	s.AKProjectId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountRequest) SetToPage(v int32) *SearchWorkitemWithTotalCountRequest {
	s.ToPage = &v
	return s
}

func (s *SearchWorkitemWithTotalCountRequest) SetPageSize(v int32) *SearchWorkitemWithTotalCountRequest {
	s.PageSize = &v
	return s
}

func (s *SearchWorkitemWithTotalCountRequest) SetCorpIdentifier(v string) *SearchWorkitemWithTotalCountRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *SearchWorkitemWithTotalCountRequest) SetSprintId(v int32) *SearchWorkitemWithTotalCountRequest {
	s.SprintId = &v
	return s
}

type SearchWorkitemWithTotalCountResponseBody struct {
	TotalCount *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ToPage     *int32                                          `json:"ToPage,omitempty" xml:"ToPage,omitempty"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalPages *int32                                          `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	Data       []*SearchWorkitemWithTotalCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code       *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Success    *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchWorkitemWithTotalCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemWithTotalCountResponseBody) GoString() string {
	return s.String()
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetTotalCount(v int32) *SearchWorkitemWithTotalCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetToPage(v int32) *SearchWorkitemWithTotalCountResponseBody {
	s.ToPage = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetRequestId(v string) *SearchWorkitemWithTotalCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetPageSize(v int32) *SearchWorkitemWithTotalCountResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetTotalPages(v int32) *SearchWorkitemWithTotalCountResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetData(v []*SearchWorkitemWithTotalCountResponseBodyData) *SearchWorkitemWithTotalCountResponseBody {
	s.Data = v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetCode(v int32) *SearchWorkitemWithTotalCountResponseBody {
	s.Code = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBody) SetSuccess(v bool) *SearchWorkitemWithTotalCountResponseBody {
	s.Success = &v
	return s
}

type SearchWorkitemWithTotalCountResponseBodyData struct {
	FixedUserId           *int32  `json:"FixedUserId,omitempty" xml:"FixedUserId,omitempty"`
	SendWangwang          *bool   `json:"SendWangwang,omitempty" xml:"SendWangwang,omitempty"`
	SkipCollab            *bool   `json:"SkipCollab,omitempty" xml:"SkipCollab,omitempty"`
	VersionList           *string `json:"VersionList,omitempty" xml:"VersionList,omitempty"`
	AkProjectId           *int32  `json:"AkProjectId,omitempty" xml:"AkProjectId,omitempty"`
	ProjectId             *int32  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserId                *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VersionId             *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	IdPath                *string `json:"IdPath,omitempty" xml:"IdPath,omitempty"`
	IgnoreIntegrate       *bool   `json:"IgnoreIntegrate,omitempty" xml:"IgnoreIntegrate,omitempty"`
	AttachmentIds         *string `json:"AttachmentIds,omitempty" xml:"AttachmentIds,omitempty"`
	CommitDate            *int64  `json:"CommitDate,omitempty" xml:"CommitDate,omitempty"`
	AkVersionIds          *string `json:"AkVersionIds,omitempty" xml:"AkVersionIds,omitempty"`
	Region                *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Trackers              *string `json:"Trackers,omitempty" xml:"Trackers,omitempty"`
	Subject               *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Solution              *int32  `json:"Solution,omitempty" xml:"Solution,omitempty"`
	AttachmentList        *string `json:"AttachmentList,omitempty" xml:"AttachmentList,omitempty"`
	UserStaffId           *string `json:"UserStaffId,omitempty" xml:"UserStaffId,omitempty"`
	IgnoreCheck           *bool   `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	AkPaths               *string `json:"AkPaths,omitempty" xml:"AkPaths,omitempty"`
	VerifierId            *int32  `json:"VerifierId,omitempty" xml:"VerifierId,omitempty"`
	ModuleList            *string `json:"ModuleList,omitempty" xml:"ModuleList,omitempty"`
	Attachmented          *bool   `json:"Attachmented,omitempty" xml:"Attachmented,omitempty"`
	SourceId              *int32  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SeriousLevelId        *int32  `json:"SeriousLevelId,omitempty" xml:"SeriousLevelId,omitempty"`
	TagIdList             *string `json:"TagIdList,omitempty" xml:"TagIdList,omitempty"`
	Watched               *bool   `json:"Watched,omitempty" xml:"Watched,omitempty"`
	AssignedTo            *string `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	AssignedToIds         *string `json:"AssignedToIds,omitempty" xml:"AssignedToIds,omitempty"`
	ClosedDuration        *int32  `json:"ClosedDuration,omitempty" xml:"ClosedDuration,omitempty"`
	PriorityId            *int32  `json:"PriorityId,omitempty" xml:"PriorityId,omitempty"`
	RecordChangeLog       *bool   `json:"RecordChangeLog,omitempty" xml:"RecordChangeLog,omitempty"`
	IssueRelations        *string `json:"IssueRelations,omitempty" xml:"IssueRelations,omitempty"`
	UpdatedAt             *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	TemplateId            *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Verifier              *string `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RelatedUserIds        *string `json:"RelatedUserIds,omitempty" xml:"RelatedUserIds,omitempty"`
	RelatedAKProjectGuids *string `json:"RelatedAKProjectGuids,omitempty" xml:"RelatedAKProjectGuids,omitempty"`
	RelatedAKProjectIds   *string `json:"RelatedAKProjectIds,omitempty" xml:"RelatedAKProjectIds,omitempty"`
	WatcherIdList         *string `json:"WatcherIdList,omitempty" xml:"WatcherIdList,omitempty"`
	CreatedAt             *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	AssignedToIdList      *string `json:"AssignedToIdList,omitempty" xml:"AssignedToIdList,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StatusStage           *int32  `json:"StatusStage,omitempty" xml:"StatusStage,omitempty"`
	AssignedToStaffId     *string `json:"AssignedToStaffId,omitempty" xml:"AssignedToStaffId,omitempty"`
	ParentId              *int32  `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	LogicalStatus         *string `json:"LogicalStatus,omitempty" xml:"LogicalStatus,omitempty"`
	VersionIds            *string `json:"VersionIds,omitempty" xml:"VersionIds,omitempty"`
	Scope                 *int32  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SeriousLevel          *string `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	StatusId              *int32  `json:"StatusId,omitempty" xml:"StatusId,omitempty"`
	Stamp                 *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	User                  *string `json:"User,omitempty" xml:"User,omitempty"`
	Source                *string `json:"Source,omitempty" xml:"Source,omitempty"`
	AssignedToId          *int32  `json:"AssignedToId,omitempty" xml:"AssignedToId,omitempty"`
	TrackerIds            *string `json:"TrackerIds,omitempty" xml:"TrackerIds,omitempty"`
	IssueTypeId           *int32  `json:"IssueTypeId,omitempty" xml:"IssueTypeId,omitempty"`
	ChangeLogList         *string `json:"ChangeLogList,omitempty" xml:"ChangeLogList,omitempty"`
	FixedDuration         *int32  `json:"FixedDuration,omitempty" xml:"FixedDuration,omitempty"`
	ScopeUserIds          *string `json:"ScopeUserIds,omitempty" xml:"ScopeUserIds,omitempty"`
	DevDuration           *int32  `json:"DevDuration,omitempty" xml:"DevDuration,omitempty"`
	LinePath              *string `json:"LinePath,omitempty" xml:"LinePath,omitempty"`
	BlackListNotice       *string `json:"BlackListNotice,omitempty" xml:"BlackListNotice,omitempty"`
	ProjectIds            *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	TestsuiteId           *int32  `json:"TestsuiteId,omitempty" xml:"TestsuiteId,omitempty"`
	ModuleUpdated         *bool   `json:"ModuleUpdated,omitempty" xml:"ModuleUpdated,omitempty"`
	CommentList           *string `json:"CommentList,omitempty" xml:"CommentList,omitempty"`
	SprintId              *int32  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	RespondDuration       *int32  `json:"RespondDuration,omitempty" xml:"RespondDuration,omitempty"`
	ModuleIds             *string `json:"ModuleIds,omitempty" xml:"ModuleIds,omitempty"`
	VerifierStaffId       *string `json:"VerifierStaffId,omitempty" xml:"VerifierStaffId,omitempty"`
	Id                    *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SearchWorkitemWithTotalCountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemWithTotalCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetFixedUserId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.FixedUserId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSendWangwang(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SendWangwang = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSkipCollab(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SkipCollab = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVersionList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.VersionList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAkProjectId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AkProjectId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetProjectId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetUserId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.UserId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVersionId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetIdPath(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.IdPath = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetIgnoreIntegrate(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.IgnoreIntegrate = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAttachmentIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AttachmentIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetCommitDate(v int64) *SearchWorkitemWithTotalCountResponseBodyData {
	s.CommitDate = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAkVersionIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AkVersionIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRegion(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Region = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetTrackers(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Trackers = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSubject(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Subject = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSolution(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Solution = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAttachmentList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AttachmentList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetUserStaffId(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.UserStaffId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetIgnoreCheck(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.IgnoreCheck = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAkPaths(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AkPaths = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVerifierId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.VerifierId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetModuleList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ModuleList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAttachmented(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Attachmented = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSourceId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSeriousLevelId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SeriousLevelId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetTagIdList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.TagIdList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetWatched(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Watched = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAssignedTo(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AssignedTo = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAssignedToIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AssignedToIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetClosedDuration(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ClosedDuration = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetPriorityId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.PriorityId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRecordChangeLog(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.RecordChangeLog = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetIssueRelations(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.IssueRelations = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetUpdatedAt(v int64) *SearchWorkitemWithTotalCountResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetTemplateId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVerifier(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Verifier = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetStatus(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Status = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRelatedUserIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.RelatedUserIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRelatedAKProjectGuids(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.RelatedAKProjectGuids = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRelatedAKProjectIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.RelatedAKProjectIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetWatcherIdList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.WatcherIdList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetCreatedAt(v int64) *SearchWorkitemWithTotalCountResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAssignedToIdList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AssignedToIdList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetPriority(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetStatusStage(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.StatusStage = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAssignedToStaffId(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AssignedToStaffId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetParentId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetLogicalStatus(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.LogicalStatus = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVersionIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.VersionIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetScope(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Scope = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSeriousLevel(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SeriousLevel = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetStatusId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.StatusId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetStamp(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Stamp = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetUser(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.User = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSource(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetAssignedToId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.AssignedToId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetTrackerIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.TrackerIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetIssueTypeId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.IssueTypeId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetChangeLogList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ChangeLogList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetFixedDuration(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.FixedDuration = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetScopeUserIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ScopeUserIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetDevDuration(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.DevDuration = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetLinePath(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.LinePath = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetBlackListNotice(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.BlackListNotice = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetProjectIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ProjectIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetTestsuiteId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.TestsuiteId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetModuleUpdated(v bool) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ModuleUpdated = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetCommentList(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.CommentList = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetSprintId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.SprintId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetRespondDuration(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.RespondDuration = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetModuleIds(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.ModuleIds = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetVerifierStaffId(v string) *SearchWorkitemWithTotalCountResponseBodyData {
	s.VerifierStaffId = &v
	return s
}

func (s *SearchWorkitemWithTotalCountResponseBodyData) SetId(v int32) *SearchWorkitemWithTotalCountResponseBodyData {
	s.Id = &v
	return s
}

type SearchWorkitemWithTotalCountResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchWorkitemWithTotalCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchWorkitemWithTotalCountResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchWorkitemWithTotalCountResponse) GoString() string {
	return s.String()
}

func (s *SearchWorkitemWithTotalCountResponse) SetHeaders(v map[string]*string) *SearchWorkitemWithTotalCountResponse {
	s.Headers = v
	return s
}

func (s *SearchWorkitemWithTotalCountResponse) SetBody(v *SearchWorkitemWithTotalCountResponseBody) *SearchWorkitemWithTotalCountResponse {
	s.Body = v
	return s
}

type SyncUserToRdcRequest struct {
	LoginTicket *string `json:"LoginTicket,omitempty" xml:"LoginTicket,omitempty"`
}

func (s SyncUserToRdcRequest) String() string {
	return tea.Prettify(s)
}

func (s SyncUserToRdcRequest) GoString() string {
	return s.String()
}

func (s *SyncUserToRdcRequest) SetLoginTicket(v string) *SyncUserToRdcRequest {
	s.LoginTicket = &v
	return s
}

type SyncUserToRdcResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SyncUserToRdcResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncUserToRdcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncUserToRdcResponseBody) GoString() string {
	return s.String()
}

func (s *SyncUserToRdcResponseBody) SetMessage(v string) *SyncUserToRdcResponseBody {
	s.Message = &v
	return s
}

func (s *SyncUserToRdcResponseBody) SetRequestId(v string) *SyncUserToRdcResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncUserToRdcResponseBody) SetData(v *SyncUserToRdcResponseBodyData) *SyncUserToRdcResponseBody {
	s.Data = v
	return s
}

func (s *SyncUserToRdcResponseBody) SetCode(v int32) *SyncUserToRdcResponseBody {
	s.Code = &v
	return s
}

func (s *SyncUserToRdcResponseBody) SetSuccess(v bool) *SyncUserToRdcResponseBody {
	s.Success = &v
	return s
}

type SyncUserToRdcResponseBodyData struct {
	IsValid *bool `json:"IsValid,omitempty" xml:"IsValid,omitempty"`
}

func (s SyncUserToRdcResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SyncUserToRdcResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncUserToRdcResponseBodyData) SetIsValid(v bool) *SyncUserToRdcResponseBodyData {
	s.IsValid = &v
	return s
}

type SyncUserToRdcResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SyncUserToRdcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SyncUserToRdcResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncUserToRdcResponse) GoString() string {
	return s.String()
}

func (s *SyncUserToRdcResponse) SetHeaders(v map[string]*string) *SyncUserToRdcResponse {
	s.Headers = v
	return s
}

func (s *SyncUserToRdcResponse) SetBody(v *SyncUserToRdcResponseBody) *SyncUserToRdcResponse {
	s.Body = v
	return s
}

type UpdateWorkitemRequest struct {
	AKProjectId    *int32                 `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	CorpIdentifier *string                `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	Modifier       *string                `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	AssignedTo     *string                `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	TemplateId     *int32                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Subject        *string                `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Description    *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Stamp          *string                `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	CfList         map[string]interface{} `json:"CfList,omitempty" xml:"CfList,omitempty"`
	IssueId        *int32                 `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Status         *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	Priority       *string                `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SeriousLevel   *string                `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	Verifier       *string                `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	SprintId       *int32                 `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	IgnoreCheck    *bool                  `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	Cfs            map[string]interface{} `json:"Cfs,omitempty" xml:"Cfs,omitempty"`
}

func (s UpdateWorkitemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemRequest) SetAKProjectId(v int32) *UpdateWorkitemRequest {
	s.AKProjectId = &v
	return s
}

func (s *UpdateWorkitemRequest) SetCorpIdentifier(v string) *UpdateWorkitemRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *UpdateWorkitemRequest) SetModifier(v string) *UpdateWorkitemRequest {
	s.Modifier = &v
	return s
}

func (s *UpdateWorkitemRequest) SetAssignedTo(v string) *UpdateWorkitemRequest {
	s.AssignedTo = &v
	return s
}

func (s *UpdateWorkitemRequest) SetTemplateId(v int32) *UpdateWorkitemRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkitemRequest) SetSubject(v string) *UpdateWorkitemRequest {
	s.Subject = &v
	return s
}

func (s *UpdateWorkitemRequest) SetDescription(v string) *UpdateWorkitemRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkitemRequest) SetStamp(v string) *UpdateWorkitemRequest {
	s.Stamp = &v
	return s
}

func (s *UpdateWorkitemRequest) SetCfList(v map[string]interface{}) *UpdateWorkitemRequest {
	s.CfList = v
	return s
}

func (s *UpdateWorkitemRequest) SetIssueId(v int32) *UpdateWorkitemRequest {
	s.IssueId = &v
	return s
}

func (s *UpdateWorkitemRequest) SetStatus(v string) *UpdateWorkitemRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkitemRequest) SetPriority(v string) *UpdateWorkitemRequest {
	s.Priority = &v
	return s
}

func (s *UpdateWorkitemRequest) SetSeriousLevel(v string) *UpdateWorkitemRequest {
	s.SeriousLevel = &v
	return s
}

func (s *UpdateWorkitemRequest) SetVerifier(v string) *UpdateWorkitemRequest {
	s.Verifier = &v
	return s
}

func (s *UpdateWorkitemRequest) SetSprintId(v int32) *UpdateWorkitemRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateWorkitemRequest) SetIgnoreCheck(v bool) *UpdateWorkitemRequest {
	s.IgnoreCheck = &v
	return s
}

func (s *UpdateWorkitemRequest) SetCfs(v map[string]interface{}) *UpdateWorkitemRequest {
	s.Cfs = v
	return s
}

type UpdateWorkitemShrinkRequest struct {
	AKProjectId    *int32  `json:"AKProjectId,omitempty" xml:"AKProjectId,omitempty"`
	CorpIdentifier *string `json:"CorpIdentifier,omitempty" xml:"CorpIdentifier,omitempty"`
	Modifier       *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	AssignedTo     *string `json:"AssignedTo,omitempty" xml:"AssignedTo,omitempty"`
	TemplateId     *int32  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Subject        *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Stamp          *string `json:"Stamp,omitempty" xml:"Stamp,omitempty"`
	CfListShrink   *string `json:"CfList,omitempty" xml:"CfList,omitempty"`
	IssueId        *int32  `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Priority       *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	SeriousLevel   *string `json:"SeriousLevel,omitempty" xml:"SeriousLevel,omitempty"`
	Verifier       *string `json:"Verifier,omitempty" xml:"Verifier,omitempty"`
	SprintId       *int32  `json:"SprintId,omitempty" xml:"SprintId,omitempty"`
	IgnoreCheck    *bool   `json:"IgnoreCheck,omitempty" xml:"IgnoreCheck,omitempty"`
	CfsShrink      *string `json:"Cfs,omitempty" xml:"Cfs,omitempty"`
}

func (s UpdateWorkitemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemShrinkRequest) SetAKProjectId(v int32) *UpdateWorkitemShrinkRequest {
	s.AKProjectId = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetCorpIdentifier(v string) *UpdateWorkitemShrinkRequest {
	s.CorpIdentifier = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetModifier(v string) *UpdateWorkitemShrinkRequest {
	s.Modifier = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetAssignedTo(v string) *UpdateWorkitemShrinkRequest {
	s.AssignedTo = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetTemplateId(v int32) *UpdateWorkitemShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetSubject(v string) *UpdateWorkitemShrinkRequest {
	s.Subject = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetDescription(v string) *UpdateWorkitemShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetStamp(v string) *UpdateWorkitemShrinkRequest {
	s.Stamp = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetCfListShrink(v string) *UpdateWorkitemShrinkRequest {
	s.CfListShrink = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetIssueId(v int32) *UpdateWorkitemShrinkRequest {
	s.IssueId = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetStatus(v string) *UpdateWorkitemShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetPriority(v string) *UpdateWorkitemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetSeriousLevel(v string) *UpdateWorkitemShrinkRequest {
	s.SeriousLevel = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetVerifier(v string) *UpdateWorkitemShrinkRequest {
	s.Verifier = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetSprintId(v int32) *UpdateWorkitemShrinkRequest {
	s.SprintId = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetIgnoreCheck(v bool) *UpdateWorkitemShrinkRequest {
	s.IgnoreCheck = &v
	return s
}

func (s *UpdateWorkitemShrinkRequest) SetCfsShrink(v string) *UpdateWorkitemShrinkRequest {
	s.CfsShrink = &v
	return s
}

type UpdateWorkitemResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkitemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemResponseBody) SetMessage(v string) *UpdateWorkitemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkitemResponseBody) SetRequestId(v string) *UpdateWorkitemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkitemResponseBody) SetData(v bool) *UpdateWorkitemResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateWorkitemResponseBody) SetCode(v int32) *UpdateWorkitemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkitemResponseBody) SetSuccess(v bool) *UpdateWorkitemResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkitemResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateWorkitemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateWorkitemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkitemResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemResponse) SetHeaders(v map[string]*string) *UpdateWorkitemResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkitemResponse) SetBody(v *UpdateWorkitemResponseBody) *UpdateWorkitemResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rdc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddEnterpriseMemberWithOptions(request *AddEnterpriseMemberRequest, runtime *util.RuntimeOptions) (_result *AddEnterpriseMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddEnterpriseMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddEnterpriseMember"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddEnterpriseMember(request *AddEnterpriseMemberRequest) (_result *AddEnterpriseMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddEnterpriseMemberResponse{}
	_body, _err := client.AddEnterpriseMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRamMemberWithOptions(request *AddRamMemberRequest, runtime *util.RuntimeOptions) (_result *AddRamMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddRamMemberResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddRamMember"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRamMember(request *AddRamMemberRequest) (_result *AddRamMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRamMemberResponse{}
	_body, _err := client.AddRamMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveJoinCompanyWithOptions(request *ApproveJoinCompanyRequest, runtime *util.RuntimeOptions) (_result *ApproveJoinCompanyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ApproveJoinCompanyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApproveJoinCompany"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveJoinCompany(request *ApproveJoinCompanyRequest) (_result *ApproveJoinCompanyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveJoinCompanyResponse{}
	_body, _err := client.ApproveJoinCompanyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEnterpriseWithOptions(tmpReq *CreateEnterpriseRequest, runtime *util.RuntimeOptions) (_result *CreateEnterpriseResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEnterpriseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Emails)) {
		request.EmailsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Emails, tea.String("Emails"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEnterpriseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEnterprise"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEnterprise(request *CreateEnterpriseRequest) (_result *CreateEnterpriseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEnterpriseResponse{}
	_body, _err := client.CreateEnterpriseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateWorkitemWithOptions(request *CreateWorkitemRequest, runtime *util.RuntimeOptions) (_result *CreateWorkitemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateWorkitemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateWorkitem"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateWorkitem(request *CreateWorkitemRequest) (_result *CreateWorkitemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkitemResponse{}
	_body, _err := client.CreateWorkitemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBindedUserByDingIdWithOptions(request *GetBindedUserByDingIdRequest, runtime *util.RuntimeOptions) (_result *GetBindedUserByDingIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetBindedUserByDingIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBindedUserByDingId"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBindedUserByDingId(request *GetBindedUserByDingIdRequest) (_result *GetBindedUserByDingIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBindedUserByDingIdResponse{}
	_body, _err := client.GetBindedUserByDingIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetChangeLogWithOptions(tmpReq *GetChangeLogRequest, runtime *util.RuntimeOptions) (_result *GetChangeLogResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetChangeLogShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TargetIds)) {
		request.TargetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TargetIds, tea.String("TargetIds"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetChangeLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetChangeLog"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetChangeLog(request *GetChangeLogRequest) (_result *GetChangeLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChangeLogResponse{}
	_body, _err := client.GetChangeLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomFieldsByTemplateIdWithOptions(request *GetCustomFieldsByTemplateIdRequest, runtime *util.RuntimeOptions) (_result *GetCustomFieldsByTemplateIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetCustomFieldsByTemplateIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCustomFieldsByTemplateId"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomFieldsByTemplateId(request *GetCustomFieldsByTemplateIdRequest) (_result *GetCustomFieldsByTemplateIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomFieldsByTemplateIdResponse{}
	_body, _err := client.GetCustomFieldsByTemplateIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIssueByIdWithOptions(request *GetIssueByIdRequest, runtime *util.RuntimeOptions) (_result *GetIssueByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIssueByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIssueById"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIssueById(request *GetIssueByIdRequest) (_result *GetIssueByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIssueByIdResponse{}
	_body, _err := client.GetIssueByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJoinCodeWithOptions(request *GetJoinCodeRequest, runtime *util.RuntimeOptions) (_result *GetJoinCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJoinCodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJoinCode"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJoinCode(request *GetJoinCodeRequest) (_result *GetJoinCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJoinCodeResponse{}
	_body, _err := client.GetJoinCodeWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetProjectMembersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetProjectMembers"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetUserByAliyunPkWithOptions(request *GetUserByAliyunPkRequest, runtime *util.RuntimeOptions) (_result *GetUserByAliyunPkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetUserByAliyunPkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserByAliyunPk"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserByAliyunPk(request *GetUserByAliyunPkRequest) (_result *GetUserByAliyunPkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserByAliyunPkResponse{}
	_body, _err := client.GetUserByAliyunPkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkitemByIdWithOptions(request *GetWorkitemByIdRequest, runtime *util.RuntimeOptions) (_result *GetWorkitemByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetWorkitemByIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWorkitemById"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkitemById(request *GetWorkitemByIdRequest) (_result *GetWorkitemByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkitemByIdResponse{}
	_body, _err := client.GetWorkitemByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinCompanyWithOptions(request *JoinCompanyRequest, runtime *util.RuntimeOptions) (_result *JoinCompanyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinCompanyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinCompany"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinCompany(request *JoinCompanyRequest) (_result *JoinCompanyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinCompanyResponse{}
	_body, _err := client.JoinCompanyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchProjectsByRegionWithOptions(request *SearchProjectsByRegionRequest, runtime *util.RuntimeOptions) (_result *SearchProjectsByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SearchProjectsByRegionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchProjectsByRegion"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchProjectsByRegion(request *SearchProjectsByRegionRequest) (_result *SearchProjectsByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchProjectsByRegionResponse{}
	_body, _err := client.SearchProjectsByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTestCaseWithOptions(request *SearchTestCaseRequest, runtime *util.RuntimeOptions) (_result *SearchTestCaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchTestCaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchTestCase"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTestCase(request *SearchTestCaseRequest) (_result *SearchTestCaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTestCaseResponse{}
	_body, _err := client.SearchTestCaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchWorkitemWithOptions(request *SearchWorkitemRequest, runtime *util.RuntimeOptions) (_result *SearchWorkitemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchWorkitemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchWorkitem"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchWorkitem(request *SearchWorkitemRequest) (_result *SearchWorkitemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchWorkitemResponse{}
	_body, _err := client.SearchWorkitemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchWorkitemWithTotalCountWithOptions(request *SearchWorkitemWithTotalCountRequest, runtime *util.RuntimeOptions) (_result *SearchWorkitemWithTotalCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchWorkitemWithTotalCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchWorkitemWithTotalCount"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchWorkitemWithTotalCount(request *SearchWorkitemWithTotalCountRequest) (_result *SearchWorkitemWithTotalCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchWorkitemWithTotalCountResponse{}
	_body, _err := client.SearchWorkitemWithTotalCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SyncUserToRdcWithOptions(request *SyncUserToRdcRequest, runtime *util.RuntimeOptions) (_result *SyncUserToRdcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SyncUserToRdcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SyncUserToRdc"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SyncUserToRdc(request *SyncUserToRdcRequest) (_result *SyncUserToRdcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncUserToRdcResponse{}
	_body, _err := client.SyncUserToRdcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateWorkitemWithOptions(tmpReq *UpdateWorkitemRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkitemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateWorkitemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CfList)) {
		request.CfListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CfList, tea.String("CfList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Cfs)) {
		request.CfsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Cfs, tea.String("Cfs"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateWorkitemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateWorkitem"), tea.String("2018-08-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateWorkitem(request *UpdateWorkitemRequest) (_result *UpdateWorkitemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkitemResponse{}
	_body, _err := client.UpdateWorkitemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
