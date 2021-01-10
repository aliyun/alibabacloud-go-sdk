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

type CheckServiceLinkedRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetRoleName(v string) *CheckServiceLinkedRoleRequest {
	s.RoleName = &v
	return s
}

type CheckServiceLinkedRoleResponseBody struct {
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CheckServiceLinkedRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBody) SetMessage(v string) *CheckServiceLinkedRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetData(v *CheckServiceLinkedRoleResponseBodyData) *CheckServiceLinkedRoleResponseBody {
	s.Data = v
	return s
}

func (s *CheckServiceLinkedRoleResponseBody) SetCode(v string) *CheckServiceLinkedRoleResponseBody {
	s.Code = &v
	return s
}

type CheckServiceLinkedRoleResponseBodyData struct {
	HasRole *bool `json:"HasRole,omitempty" xml:"HasRole,omitempty"`
}

func (s CheckServiceLinkedRoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponseBodyData) SetHasRole(v bool) *CheckServiceLinkedRoleResponseBodyData {
	s.HasRole = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PackageName  *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetName(v string) *CreateAppRequest {
	s.Name = &v
	return s
}

func (s *CreateAppRequest) SetPackageName(v string) *CreateAppRequest {
	s.PackageName = &v
	return s
}

func (s *CreateAppRequest) SetDepartmentId(v string) *CreateAppRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateAppRequest) SetClientToken(v string) *CreateAppRequest {
	s.ClientToken = &v
	return s
}

type CreateAppResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetData(v *CreateAppResponseBodyData) *CreateAppResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppResponseBody) SetCode(v string) *CreateAppResponseBody {
	s.Code = &v
	return s
}

type CreateAppResponseBodyData struct {
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyData) SetCreatorName(v string) *CreateAppResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *CreateAppResponseBodyData) SetDisabled(v bool) *CreateAppResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *CreateAppResponseBodyData) SetCreatedAt(v string) *CreateAppResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateAppResponseBodyData) SetName(v string) *CreateAppResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyData) SetId(v string) *CreateAppResponseBodyData {
	s.Id = &v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateDepartmentRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateDepartmentRequest) SetDescription(v string) *CreateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *CreateDepartmentRequest) SetLabel(v string) *CreateDepartmentRequest {
	s.Label = &v
	return s
}

func (s *CreateDepartmentRequest) SetName(v string) *CreateDepartmentRequest {
	s.Name = &v
	return s
}

func (s *CreateDepartmentRequest) SetClientToken(v string) *CreateDepartmentRequest {
	s.ClientToken = &v
	return s
}

type CreateDepartmentResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBody) SetMessage(v string) *CreateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetRequestId(v string) *CreateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDepartmentResponseBody) SetData(v *CreateDepartmentResponseBodyData) *CreateDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateDepartmentResponseBody) SetCode(v string) *CreateDepartmentResponseBody {
	s.Code = &v
	return s
}

type CreateDepartmentResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponseBodyData) SetDescription(v string) *CreateDepartmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetCreatedAt(v string) *CreateDepartmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetName(v string) *CreateDepartmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateDepartmentResponseBodyData) SetId(v string) *CreateDepartmentResponseBodyData {
	s.Id = &v
	return s
}

type CreateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
	s.Body = v
	return s
}

type CreateDetectProcessRequest struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Draft   *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessRequest) SetName(v string) *CreateDetectProcessRequest {
	s.Name = &v
	return s
}

func (s *CreateDetectProcessRequest) SetDraft(v string) *CreateDetectProcessRequest {
	s.Draft = &v
	return s
}

func (s *CreateDetectProcessRequest) SetContent(v string) *CreateDetectProcessRequest {
	s.Content = &v
	return s
}

type CreateDetectProcessResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponseBody) SetMessage(v string) *CreateDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDetectProcessResponseBody) SetRequestId(v string) *CreateDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDetectProcessResponseBody) SetData(v *CreateDetectProcessResponseBodyData) *CreateDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *CreateDetectProcessResponseBody) SetCode(v string) *CreateDetectProcessResponseBody {
	s.Code = &v
	return s
}

type CreateDetectProcessResponseBodyData struct {
	Draft     *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Disabled  *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Md5       *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponseBodyData) SetDraft(v string) *CreateDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetCreatedAt(v string) *CreateDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetDisabled(v bool) *CreateDetectProcessResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetMd5(v string) *CreateDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetName(v string) *CreateDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetContent(v string) *CreateDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateDetectProcessResponseBodyData) SetId(v string) *CreateDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

type CreateDetectProcessResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *CreateDetectProcessResponse) SetHeaders(v map[string]*string) *CreateDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *CreateDetectProcessResponse) SetBody(v *CreateDetectProcessResponseBody) *CreateDetectProcessResponse {
	s.Body = v
	return s
}

type CreateLiveRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RtcCode     *string `json:"RtcCode,omitempty" xml:"RtcCode,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CreateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRequest) SetName(v string) *CreateLiveRequest {
	s.Name = &v
	return s
}

func (s *CreateLiveRequest) SetRtcCode(v string) *CreateLiveRequest {
	s.RtcCode = &v
	return s
}

func (s *CreateLiveRequest) SetUserId(v string) *CreateLiveRequest {
	s.UserId = &v
	return s
}

func (s *CreateLiveRequest) SetClientToken(v string) *CreateLiveRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLiveRequest) SetAppId(v string) *CreateLiveRequest {
	s.AppId = &v
	return s
}

type CreateLiveResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateLiveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBody) SetMessage(v string) *CreateLiveResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLiveResponseBody) SetRequestId(v string) *CreateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveResponseBody) SetData(v *CreateLiveResponseBodyData) *CreateLiveResponseBody {
	s.Data = v
	return s
}

func (s *CreateLiveResponseBody) SetCode(v string) *CreateLiveResponseBody {
	s.Code = &v
	return s
}

type CreateLiveResponseBodyData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Channel   *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateLiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLiveResponseBodyData) SetCreatedAt(v string) *CreateLiveResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateLiveResponseBodyData) SetChannel(v string) *CreateLiveResponseBodyData {
	s.Channel = &v
	return s
}

func (s *CreateLiveResponseBodyData) SetName(v string) *CreateLiveResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateLiveResponseBodyData) SetId(v string) *CreateLiveResponseBodyData {
	s.Id = &v
	return s
}

type CreateLiveResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveResponse) SetHeaders(v map[string]*string) *CreateLiveResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveResponse) SetBody(v *CreateLiveResponseBody) *CreateLiveResponse {
	s.Body = v
	return s
}

type CreateLiveDetectionRequest struct {
	LiveId      *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	MetaUrl     *string `json:"MetaUrl,omitempty" xml:"MetaUrl,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateLiveDetectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveDetectionRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveDetectionRequest) SetLiveId(v string) *CreateLiveDetectionRequest {
	s.LiveId = &v
	return s
}

func (s *CreateLiveDetectionRequest) SetUserId(v string) *CreateLiveDetectionRequest {
	s.UserId = &v
	return s
}

func (s *CreateLiveDetectionRequest) SetRuleId(v string) *CreateLiveDetectionRequest {
	s.RuleId = &v
	return s
}

func (s *CreateLiveDetectionRequest) SetMetaUrl(v string) *CreateLiveDetectionRequest {
	s.MetaUrl = &v
	return s
}

func (s *CreateLiveDetectionRequest) SetClientToken(v string) *CreateLiveDetectionRequest {
	s.ClientToken = &v
	return s
}

type CreateLiveDetectionResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateLiveDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateLiveDetectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLiveDetectionResponseBody) SetMessage(v string) *CreateLiveDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLiveDetectionResponseBody) SetRequestId(v string) *CreateLiveDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLiveDetectionResponseBody) SetData(v *CreateLiveDetectionResponseBodyData) *CreateLiveDetectionResponseBody {
	s.Data = v
	return s
}

func (s *CreateLiveDetectionResponseBody) SetCode(v string) *CreateLiveDetectionResponseBody {
	s.Code = &v
	return s
}

type CreateLiveDetectionResponseBodyData struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateLiveDetectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLiveDetectionResponseBodyData) SetId(v string) *CreateLiveDetectionResponseBodyData {
	s.Id = &v
	return s
}

type CreateLiveDetectionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLiveDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLiveDetectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLiveDetectionResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveDetectionResponse) SetHeaders(v map[string]*string) *CreateLiveDetectionResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveDetectionResponse) SetBody(v *CreateLiveDetectionResponseBody) *CreateLiveDetectionResponse {
	s.Body = v
	return s
}

type CreateRuleRequest struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRuleRequest) SetName(v string) *CreateRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateRuleRequest) SetContent(v string) *CreateRuleRequest {
	s.Content = &v
	return s
}

func (s *CreateRuleRequest) SetClientToken(v string) *CreateRuleRequest {
	s.ClientToken = &v
	return s
}

type CreateRuleResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) SetMessage(v string) *CreateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetData(v *CreateRuleResponseBodyData) *CreateRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateRuleResponseBody) SetCode(v string) *CreateRuleResponseBody {
	s.Code = &v
	return s
}

type CreateRuleResponseBodyData struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBodyData) SetName(v string) *CreateRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateRuleResponseBodyData) SetContent(v string) *CreateRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *CreateRuleResponseBodyData) SetId(v string) *CreateRuleResponseBodyData {
	s.Id = &v
	return s
}

type CreateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRuleResponse) SetHeaders(v map[string]*string) *CreateRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRuleResponse) SetBody(v *CreateRuleResponseBody) *CreateRuleResponse {
	s.Body = v
	return s
}

type CreateStatisticsRecordRequest struct {
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeviceType  *int32  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	BeginAt     *string `json:"BeginAt,omitempty" xml:"BeginAt,omitempty"`
	EndAt       *string `json:"EndAt,omitempty" xml:"EndAt,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateStatisticsRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateStatisticsRecordRequest) SetDeviceId(v string) *CreateStatisticsRecordRequest {
	s.DeviceId = &v
	return s
}

func (s *CreateStatisticsRecordRequest) SetAppId(v string) *CreateStatisticsRecordRequest {
	s.AppId = &v
	return s
}

func (s *CreateStatisticsRecordRequest) SetDeviceType(v int32) *CreateStatisticsRecordRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateStatisticsRecordRequest) SetBeginAt(v string) *CreateStatisticsRecordRequest {
	s.BeginAt = &v
	return s
}

func (s *CreateStatisticsRecordRequest) SetEndAt(v string) *CreateStatisticsRecordRequest {
	s.EndAt = &v
	return s
}

func (s *CreateStatisticsRecordRequest) SetClientToken(v string) *CreateStatisticsRecordRequest {
	s.ClientToken = &v
	return s
}

type CreateStatisticsRecordResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateStatisticsRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStatisticsRecordResponseBody) SetMessage(v string) *CreateStatisticsRecordResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStatisticsRecordResponseBody) SetRequestId(v string) *CreateStatisticsRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStatisticsRecordResponseBody) SetCode(v string) *CreateStatisticsRecordResponseBody {
	s.Code = &v
	return s
}

type CreateStatisticsRecordResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStatisticsRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStatisticsRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateStatisticsRecordResponse) SetHeaders(v map[string]*string) *CreateStatisticsRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateStatisticsRecordResponse) SetBody(v *CreateStatisticsRecordResponseBody) *CreateStatisticsRecordResponse {
	s.Body = v
	return s
}

type CreateStatisticsTaskRequest struct {
	DateFrom     *string   `json:"DateFrom,omitempty" xml:"DateFrom,omitempty"`
	DateTo       *string   `json:"DateTo,omitempty" xml:"DateTo,omitempty"`
	ClientToken  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
}

func (s CreateStatisticsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateStatisticsTaskRequest) SetDateFrom(v string) *CreateStatisticsTaskRequest {
	s.DateFrom = &v
	return s
}

func (s *CreateStatisticsTaskRequest) SetDateTo(v string) *CreateStatisticsTaskRequest {
	s.DateTo = &v
	return s
}

func (s *CreateStatisticsTaskRequest) SetClientToken(v string) *CreateStatisticsTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStatisticsTaskRequest) SetDepartmentId(v []*string) *CreateStatisticsTaskRequest {
	s.DepartmentId = v
	return s
}

type CreateStatisticsTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateStatisticsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStatisticsTaskResponseBody) SetMessage(v string) *CreateStatisticsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStatisticsTaskResponseBody) SetRequestId(v string) *CreateStatisticsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStatisticsTaskResponseBody) SetCode(v string) *CreateStatisticsTaskResponseBody {
	s.Code = &v
	return s
}

type CreateStatisticsTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateStatisticsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStatisticsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStatisticsTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateStatisticsTaskResponse) SetHeaders(v map[string]*string) *CreateStatisticsTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateStatisticsTaskResponse) SetBody(v *CreateStatisticsTaskResponseBody) *CreateStatisticsTaskResponse {
	s.Body = v
	return s
}

type CreateTaskGroupRequest struct {
	AppId            *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ExpireAt         *string   `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	GroupName        *string   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RuleId           *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RunnableTimeFrom *string   `json:"RunnableTimeFrom,omitempty" xml:"RunnableTimeFrom,omitempty"`
	RunnableTimeTo   *string   `json:"RunnableTimeTo,omitempty" xml:"RunnableTimeTo,omitempty"`
	TriggerPeriod    *string   `json:"TriggerPeriod,omitempty" xml:"TriggerPeriod,omitempty"`
	ClientToken      *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Day              []*int    `json:"Day,omitempty" xml:"Day,omitempty" type:"Repeated"`
	VideoUrl         []*string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty" type:"Repeated"`
}

func (s CreateTaskGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupRequest) SetAppId(v string) *CreateTaskGroupRequest {
	s.AppId = &v
	return s
}

func (s *CreateTaskGroupRequest) SetExpireAt(v string) *CreateTaskGroupRequest {
	s.ExpireAt = &v
	return s
}

func (s *CreateTaskGroupRequest) SetGroupName(v string) *CreateTaskGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRuleId(v string) *CreateTaskGroupRequest {
	s.RuleId = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRunnableTimeFrom(v string) *CreateTaskGroupRequest {
	s.RunnableTimeFrom = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRunnableTimeTo(v string) *CreateTaskGroupRequest {
	s.RunnableTimeTo = &v
	return s
}

func (s *CreateTaskGroupRequest) SetTriggerPeriod(v string) *CreateTaskGroupRequest {
	s.TriggerPeriod = &v
	return s
}

func (s *CreateTaskGroupRequest) SetClientToken(v string) *CreateTaskGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskGroupRequest) SetDay(v []*int) *CreateTaskGroupRequest {
	s.Day = v
	return s
}

func (s *CreateTaskGroupRequest) SetVideoUrl(v []*string) *CreateTaskGroupRequest {
	s.VideoUrl = v
	return s
}

type CreateTaskGroupResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *CreateTaskGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateTaskGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBody) SetMessage(v string) *CreateTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetRequestId(v string) *CreateTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskGroupResponseBody) SetData(v *CreateTaskGroupResponseBodyData) *CreateTaskGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateTaskGroupResponseBody) SetCode(v string) *CreateTaskGroupResponseBody {
	s.Code = &v
	return s
}

type CreateTaskGroupResponseBodyData struct {
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	CompletedTasks *int32    `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	TotalTasks     *int32    `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskIds        []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleId         *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateTaskGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponseBodyData) SetStatus(v string) *CreateTaskGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetCompletedTasks(v int32) *CreateTaskGroupResponseBodyData {
	s.CompletedTasks = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetTotalTasks(v int32) *CreateTaskGroupResponseBodyData {
	s.TotalTasks = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetCreatedAt(v string) *CreateTaskGroupResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetName(v string) *CreateTaskGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetTaskIds(v []*string) *CreateTaskGroupResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetId(v string) *CreateTaskGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetRuleId(v string) *CreateTaskGroupResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *CreateTaskGroupResponseBodyData) SetRuleName(v string) *CreateTaskGroupResponseBodyData {
	s.RuleName = &v
	return s
}

type CreateTaskGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTaskGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupResponse) SetHeaders(v map[string]*string) *CreateTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskGroupResponse) SetBody(v *CreateTaskGroupResponseBody) *CreateTaskGroupResponse {
	s.Body = v
	return s
}

type CreateUserDepartmentsRequest struct {
	UserId       []*string `json:"UserId,omitempty" xml:"UserId,omitempty" type:"Repeated"`
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
}

func (s CreateUserDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsRequest) SetUserId(v []*string) *CreateUserDepartmentsRequest {
	s.UserId = v
	return s
}

func (s *CreateUserDepartmentsRequest) SetDepartmentId(v []*string) *CreateUserDepartmentsRequest {
	s.DepartmentId = v
	return s
}

type CreateUserDepartmentsResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateUserDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsResponseBody) SetMessage(v string) *CreateUserDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetRequestId(v string) *CreateUserDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetData(v map[string]interface{}) *CreateUserDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserDepartmentsResponseBody) SetCode(v string) *CreateUserDepartmentsResponseBody {
	s.Code = &v
	return s
}

type CreateUserDepartmentsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *CreateUserDepartmentsResponse) SetHeaders(v map[string]*string) *CreateUserDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *CreateUserDepartmentsResponse) SetBody(v *CreateUserDepartmentsResponseBody) *CreateUserDepartmentsResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetId(v string) *DeleteAppRequest {
	s.Id = &v
	return s
}

type DeleteAppResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetData(v map[string]interface{}) *DeleteAppResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAppResponseBody) SetCode(v string) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

type DeleteAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteDepartmentRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentRequest) SetId(v string) *DeleteDepartmentRequest {
	s.Id = &v
	return s
}

type DeleteDepartmentResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponseBody) SetMessage(v string) *DeleteDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetRequestId(v string) *DeleteDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDepartmentResponseBody) SetData(v map[string]interface{}) *DeleteDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDepartmentResponseBody) SetCode(v string) *DeleteDepartmentResponseBody {
	s.Code = &v
	return s
}

type DeleteDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponse) SetHeaders(v map[string]*string) *DeleteDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDepartmentResponse) SetBody(v *DeleteDepartmentResponseBody) *DeleteDepartmentResponse {
	s.Body = v
	return s
}

type DeleteDetectProcessRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessRequest) SetId(v string) *DeleteDetectProcessRequest {
	s.Id = &v
	return s
}

type DeleteDetectProcessResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessResponseBody) SetMessage(v string) *DeleteDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetRequestId(v string) *DeleteDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetData(v map[string]interface{}) *DeleteDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDetectProcessResponseBody) SetCode(v string) *DeleteDetectProcessResponseBody {
	s.Code = &v
	return s
}

type DeleteDetectProcessResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *DeleteDetectProcessResponse) SetHeaders(v map[string]*string) *DeleteDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *DeleteDetectProcessResponse) SetBody(v *DeleteDetectProcessResponseBody) *DeleteDetectProcessResponse {
	s.Body = v
	return s
}

type DeleteRuleRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) SetId(v string) *DeleteRuleRequest {
	s.Id = &v
	return s
}

type DeleteRuleResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DeleteRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetData(v *DeleteRuleResponseBodyData) *DeleteRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

type DeleteRuleResponseBodyData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBodyData) SetCreatedAt(v string) *DeleteRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetName(v string) *DeleteRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetContent(v string) *DeleteRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *DeleteRuleResponseBodyData) SetId(v string) *DeleteRuleResponseBodyData {
	s.Id = &v
	return s
}

type DeleteRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponse) SetHeaders(v map[string]*string) *DeleteRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRuleResponse) SetBody(v *DeleteRuleResponseBody) *DeleteRuleResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetId(v string) *DeleteUserRequest {
	s.Id = &v
	return s
}

type DeleteUserResponseBody struct {
	Errors    []*DeleteUserResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{}          `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetErrors(v []*DeleteUserResponseBodyErrors) *DeleteUserResponseBody {
	s.Errors = v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetData(v map[string]interface{}) *DeleteUserResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

type DeleteUserResponseBodyErrors struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
}

func (s DeleteUserResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBodyErrors) SetMessage(v string) *DeleteUserResponseBodyErrors {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBodyErrors) SetField(v string) *DeleteUserResponseBodyErrors {
	s.Field = &v
	return s
}

type DeleteUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserDepartmentsRequest struct {
	UserId       []*string `json:"UserId,omitempty" xml:"UserId,omitempty" type:"Repeated"`
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
}

func (s DeleteUserDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsRequest) SetUserId(v []*string) *DeleteUserDepartmentsRequest {
	s.UserId = v
	return s
}

func (s *DeleteUserDepartmentsRequest) SetDepartmentId(v []*string) *DeleteUserDepartmentsRequest {
	s.DepartmentId = v
	return s
}

type DeleteUserDepartmentsResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteUserDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsResponseBody) SetMessage(v string) *DeleteUserDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetRequestId(v string) *DeleteUserDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetData(v map[string]interface{}) *DeleteUserDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *DeleteUserDepartmentsResponseBody) SetCode(v string) *DeleteUserDepartmentsResponseBody {
	s.Code = &v
	return s
}

type DeleteUserDepartmentsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDepartmentsResponse) SetHeaders(v map[string]*string) *DeleteUserDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDepartmentsResponse) SetBody(v *DeleteUserDepartmentsResponseBody) *DeleteUserDepartmentsResponse {
	s.Body = v
	return s
}

type ExitLiveRequest struct {
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	RtcCode *string `json:"RtcCode,omitempty" xml:"RtcCode,omitempty"`
}

func (s ExitLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s ExitLiveRequest) GoString() string {
	return s.String()
}

func (s *ExitLiveRequest) SetChannel(v string) *ExitLiveRequest {
	s.Channel = &v
	return s
}

func (s *ExitLiveRequest) SetUserId(v string) *ExitLiveRequest {
	s.UserId = &v
	return s
}

func (s *ExitLiveRequest) SetRtcCode(v string) *ExitLiveRequest {
	s.RtcCode = &v
	return s
}

type ExitLiveResponseBody struct {
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ExitLiveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExitLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExitLiveResponseBody) GoString() string {
	return s.String()
}

func (s *ExitLiveResponseBody) SetMessage(v string) *ExitLiveResponseBody {
	s.Message = &v
	return s
}

func (s *ExitLiveResponseBody) SetRequestId(v string) *ExitLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExitLiveResponseBody) SetData(v *ExitLiveResponseBodyData) *ExitLiveResponseBody {
	s.Data = v
	return s
}

func (s *ExitLiveResponseBody) SetCode(v string) *ExitLiveResponseBody {
	s.Code = &v
	return s
}

type ExitLiveResponseBodyData struct {
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExitLiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExitLiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExitLiveResponseBodyData) SetCode(v int32) *ExitLiveResponseBodyData {
	s.Code = &v
	return s
}

type ExitLiveResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExitLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExitLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s ExitLiveResponse) GoString() string {
	return s.String()
}

func (s *ExitLiveResponse) SetHeaders(v map[string]*string) *ExitLiveResponse {
	s.Headers = v
	return s
}

func (s *ExitLiveResponse) SetBody(v *ExitLiveResponseBody) *ExitLiveResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	DeviceId    *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetId(v string) *GetAppRequest {
	s.Id = &v
	return s
}

func (s *GetAppRequest) SetPackageName(v string) *GetAppRequest {
	s.PackageName = &v
	return s
}

func (s *GetAppRequest) SetDeviceId(v string) *GetAppRequest {
	s.DeviceId = &v
	return s
}

type GetAppResponseBody struct {
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetMessage(v string) *GetAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetData(v *GetAppResponseBodyData) *GetAppResponseBody {
	s.Data = v
	return s
}

func (s *GetAppResponseBody) SetCode(v string) *GetAppResponseBody {
	s.Code = &v
	return s
}

type GetAppResponseBodyData struct {
	CreatedAt *int32  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Disabled  *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppResponseBodyData) SetCreatedAt(v int32) *GetAppResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAppResponseBodyData) SetDisabled(v string) *GetAppResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *GetAppResponseBodyData) SetName(v string) *GetAppResponseBodyData {
	s.Name = &v
	return s
}

type GetAppResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetBatchSignedUrlRequest struct {
	FileUrlList []*string `json:"FileUrlList,omitempty" xml:"FileUrlList,omitempty" type:"Repeated"`
}

func (s GetBatchSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBatchSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetBatchSignedUrlRequest) SetFileUrlList(v []*string) *GetBatchSignedUrlRequest {
	s.FileUrlList = v
	return s
}

type GetBatchSignedUrlResponseBody struct {
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetBatchSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBatchSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchSignedUrlResponseBody) SetMessage(v string) *GetBatchSignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchSignedUrlResponseBody) SetRequestId(v string) *GetBatchSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchSignedUrlResponseBody) SetData(v []*string) *GetBatchSignedUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchSignedUrlResponseBody) SetCode(v string) *GetBatchSignedUrlResponseBody {
	s.Code = &v
	return s
}

type GetBatchSignedUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetBatchSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBatchSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBatchSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetBatchSignedUrlResponse) SetHeaders(v map[string]*string) *GetBatchSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetBatchSignedUrlResponse) SetBody(v *GetBatchSignedUrlResponseBody) *GetBatchSignedUrlResponse {
	s.Body = v
	return s
}

type GetDepartmentRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentRequest) GoString() string {
	return s.String()
}

func (s *GetDepartmentRequest) SetId(v string) *GetDepartmentRequest {
	s.Id = &v
	return s
}

type GetDepartmentResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponseBody) SetMessage(v string) *GetDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetDepartmentResponseBody) SetRequestId(v string) *GetDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDepartmentResponseBody) SetData(v *GetDepartmentResponseBodyData) *GetDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *GetDepartmentResponseBody) SetCode(v string) *GetDepartmentResponseBody {
	s.Code = &v
	return s
}

type GetDepartmentResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponseBodyData) SetDescription(v string) *GetDepartmentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetCreatedAt(v string) *GetDepartmentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetUpdatedAt(v string) *GetDepartmentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetName(v string) *GetDepartmentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDepartmentResponseBodyData) SetId(v string) *GetDepartmentResponseBodyData {
	s.Id = &v
	return s
}

type GetDepartmentResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDepartmentResponse) GoString() string {
	return s.String()
}

func (s *GetDepartmentResponse) SetHeaders(v map[string]*string) *GetDepartmentResponse {
	s.Headers = v
	return s
}

func (s *GetDepartmentResponse) SetBody(v *GetDepartmentResponseBody) *GetDepartmentResponse {
	s.Body = v
	return s
}

type GetDetectEvaluationRequest struct {
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetDetectEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectEvaluationRequest) GoString() string {
	return s.String()
}

func (s *GetDetectEvaluationRequest) SetStartTime(v string) *GetDetectEvaluationRequest {
	s.StartTime = &v
	return s
}

func (s *GetDetectEvaluationRequest) SetEndTime(v string) *GetDetectEvaluationRequest {
	s.EndTime = &v
	return s
}

type GetDetectEvaluationResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetDetectEvaluationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDetectEvaluationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectEvaluationResponseBody) SetMessage(v string) *GetDetectEvaluationResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectEvaluationResponseBody) SetRequestId(v string) *GetDetectEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectEvaluationResponseBody) SetData(v []*GetDetectEvaluationResponseBodyData) *GetDetectEvaluationResponseBody {
	s.Data = v
	return s
}

func (s *GetDetectEvaluationResponseBody) SetCode(v string) *GetDetectEvaluationResponseBody {
	s.Code = &v
	return s
}

type GetDetectEvaluationResponseBodyData struct {
	Day                *string                                                  `json:"Day,omitempty" xml:"Day,omitempty"`
	EvaluationItemList []*GetDetectEvaluationResponseBodyDataEvaluationItemList `json:"EvaluationItemList,omitempty" xml:"EvaluationItemList,omitempty" type:"Repeated"`
}

func (s GetDetectEvaluationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetectEvaluationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetectEvaluationResponseBodyData) SetDay(v string) *GetDetectEvaluationResponseBodyData {
	s.Day = &v
	return s
}

func (s *GetDetectEvaluationResponseBodyData) SetEvaluationItemList(v []*GetDetectEvaluationResponseBodyDataEvaluationItemList) *GetDetectEvaluationResponseBodyData {
	s.EvaluationItemList = v
	return s
}

type GetDetectEvaluationResponseBodyDataEvaluationItemList struct {
	SuccessRate  *string `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	HandleCount  *int32  `json:"HandleCount,omitempty" xml:"HandleCount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SuccessCount *int32  `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s GetDetectEvaluationResponseBodyDataEvaluationItemList) String() string {
	return tea.Prettify(s)
}

func (s GetDetectEvaluationResponseBodyDataEvaluationItemList) GoString() string {
	return s.String()
}

func (s *GetDetectEvaluationResponseBodyDataEvaluationItemList) SetSuccessRate(v string) *GetDetectEvaluationResponseBodyDataEvaluationItemList {
	s.SuccessRate = &v
	return s
}

func (s *GetDetectEvaluationResponseBodyDataEvaluationItemList) SetHandleCount(v int32) *GetDetectEvaluationResponseBodyDataEvaluationItemList {
	s.HandleCount = &v
	return s
}

func (s *GetDetectEvaluationResponseBodyDataEvaluationItemList) SetName(v string) *GetDetectEvaluationResponseBodyDataEvaluationItemList {
	s.Name = &v
	return s
}

func (s *GetDetectEvaluationResponseBodyDataEvaluationItemList) SetSuccessCount(v int32) *GetDetectEvaluationResponseBodyDataEvaluationItemList {
	s.SuccessCount = &v
	return s
}

type GetDetectEvaluationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectEvaluationResponse) GoString() string {
	return s.String()
}

func (s *GetDetectEvaluationResponse) SetHeaders(v map[string]*string) *GetDetectEvaluationResponse {
	s.Headers = v
	return s
}

func (s *GetDetectEvaluationResponse) SetBody(v *GetDetectEvaluationResponseBody) *GetDetectEvaluationResponse {
	s.Body = v
	return s
}

type GetDetectionRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetDetectionRequest) SetId(v string) *GetDetectionRequest {
	s.Id = &v
	return s
}

type GetDetectionResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetDetectionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDetectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBody) SetMessage(v string) *GetDetectionResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectionResponseBody) SetRequestId(v string) *GetDetectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectionResponseBody) SetData(v *GetDetectionResponseBodyData) *GetDetectionResponseBody {
	s.Data = v
	return s
}

func (s *GetDetectionResponseBody) SetCode(v string) *GetDetectionResponseBody {
	s.Code = &v
	return s
}

type GetDetectionResponseBodyData struct {
	Status         *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	DepartmentName *string                              `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	Tasks          []*GetDetectionResponseBodyDataTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	RecordingType  *string                              `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	CreatedAt      *string                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentId   *string                              `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Id             *string                              `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleName       *string                              `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId         *string                              `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetDetectionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBodyData) SetStatus(v string) *GetDetectionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetDepartmentName(v string) *GetDetectionResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetTasks(v []*GetDetectionResponseBodyDataTasks) *GetDetectionResponseBodyData {
	s.Tasks = v
	return s
}

func (s *GetDetectionResponseBodyData) SetRecordingType(v string) *GetDetectionResponseBodyData {
	s.RecordingType = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetCreatedAt(v string) *GetDetectionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetDepartmentId(v string) *GetDetectionResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetId(v string) *GetDetectionResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetRuleName(v string) *GetDetectionResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetDetectionResponseBodyData) SetRuleId(v string) *GetDetectionResponseBodyData {
	s.RuleId = &v
	return s
}

type GetDetectionResponseBodyDataTasks struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetDetectionResponseBodyDataTasks) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponseBodyDataTasks) GoString() string {
	return s.String()
}

func (s *GetDetectionResponseBodyDataTasks) SetStatus(v string) *GetDetectionResponseBodyDataTasks {
	s.Status = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetVideoMetaUrl(v string) *GetDetectionResponseBodyDataTasks {
	s.VideoMetaUrl = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetCreatedAt(v string) *GetDetectionResponseBodyDataTasks {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetId(v string) *GetDetectionResponseBodyDataTasks {
	s.Id = &v
	return s
}

func (s *GetDetectionResponseBodyDataTasks) SetVideoUrl(v string) *GetDetectionResponseBodyDataTasks {
	s.VideoUrl = &v
	return s
}

type GetDetectionResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectionResponse) GoString() string {
	return s.String()
}

func (s *GetDetectionResponse) SetHeaders(v map[string]*string) *GetDetectionResponse {
	s.Headers = v
	return s
}

func (s *GetDetectionResponse) SetBody(v *GetDetectionResponseBody) *GetDetectionResponse {
	s.Body = v
	return s
}

type GetDetectProcessRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *GetDetectProcessRequest) SetId(v string) *GetDetectProcessRequest {
	s.Id = &v
	return s
}

type GetDetectProcessResponseBody struct {
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponseBody) SetMessage(v string) *GetDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectProcessResponseBody) SetRequestId(v string) *GetDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectProcessResponseBody) SetData(v *GetDetectProcessResponseBodyData) *GetDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *GetDetectProcessResponseBody) SetCode(v string) *GetDetectProcessResponseBody {
	s.Code = &v
	return s
}

type GetDetectProcessResponseBodyData struct {
	Draft      *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	CreatedAt  *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	UpdatedAt  *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewVersion *bool   `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponseBodyData) SetDraft(v string) *GetDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetCreatedAt(v string) *GetDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetMd5(v string) *GetDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetUpdatedAt(v string) *GetDetectProcessResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetName(v string) *GetDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetNewVersion(v bool) *GetDetectProcessResponseBodyData {
	s.NewVersion = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetContent(v string) *GetDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDetectProcessResponseBodyData) SetId(v string) *GetDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

type GetDetectProcessResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *GetDetectProcessResponse) SetHeaders(v map[string]*string) *GetDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *GetDetectProcessResponse) SetBody(v *GetDetectProcessResponseBody) *GetDetectProcessResponse {
	s.Body = v
	return s
}

type GetDetectProcessJsonFileRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDetectProcessJsonFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileRequest) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileRequest) SetId(v string) *GetDetectProcessJsonFileRequest {
	s.Id = &v
	return s
}

type GetDetectProcessJsonFileResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDetectProcessJsonFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileResponseBody) SetMessage(v string) *GetDetectProcessJsonFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetRequestId(v string) *GetDetectProcessJsonFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetData(v string) *GetDetectProcessJsonFileResponseBody {
	s.Data = &v
	return s
}

func (s *GetDetectProcessJsonFileResponseBody) SetCode(v string) *GetDetectProcessJsonFileResponseBody {
	s.Code = &v
	return s
}

type GetDetectProcessJsonFileResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectProcessJsonFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectProcessJsonFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessJsonFileResponse) GoString() string {
	return s.String()
}

func (s *GetDetectProcessJsonFileResponse) SetHeaders(v map[string]*string) *GetDetectProcessJsonFileResponse {
	s.Headers = v
	return s
}

func (s *GetDetectProcessJsonFileResponse) SetBody(v *GetDetectProcessJsonFileResponseBody) *GetDetectProcessJsonFileResponse {
	s.Body = v
	return s
}

type GetDetectProcessTemplateResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDetectProcessTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetectProcessTemplateResponseBody) SetMessage(v string) *GetDetectProcessTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetectProcessTemplateResponseBody) SetRequestId(v string) *GetDetectProcessTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetectProcessTemplateResponseBody) SetData(v string) *GetDetectProcessTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *GetDetectProcessTemplateResponseBody) SetCode(v string) *GetDetectProcessTemplateResponseBody {
	s.Code = &v
	return s
}

type GetDetectProcessTemplateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDetectProcessTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDetectProcessTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDetectProcessTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetDetectProcessTemplateResponse) SetHeaders(v map[string]*string) *GetDetectProcessTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetDetectProcessTemplateResponse) SetBody(v *GetDetectProcessTemplateResponseBody) *GetDetectProcessTemplateResponse {
	s.Body = v
	return s
}

type GetGlobalConfigRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetGlobalConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalConfigRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalConfigRequest) SetName(v string) *GetGlobalConfigRequest {
	s.Name = &v
	return s
}

type GetGlobalConfigResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetGlobalConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalConfigResponseBody) SetMessage(v string) *GetGlobalConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetGlobalConfigResponseBody) SetRequestId(v string) *GetGlobalConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalConfigResponseBody) SetData(v string) *GetGlobalConfigResponseBody {
	s.Data = &v
	return s
}

func (s *GetGlobalConfigResponseBody) SetCode(v string) *GetGlobalConfigResponseBody {
	s.Code = &v
	return s
}

type GetGlobalConfigResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGlobalConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGlobalConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGlobalConfigResponse) GoString() string {
	return s.String()
}

func (s *GetGlobalConfigResponse) SetHeaders(v map[string]*string) *GetGlobalConfigResponse {
	s.Headers = v
	return s
}

func (s *GetGlobalConfigResponse) SetBody(v *GetGlobalConfigResponseBody) *GetGlobalConfigResponse {
	s.Body = v
	return s
}

type GetModelSignedUrlRequest struct {
	ModelPath []*string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty" type:"Repeated"`
}

func (s GetModelSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetModelSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetModelSignedUrlRequest) SetModelPath(v []*string) *GetModelSignedUrlRequest {
	s.ModelPath = v
	return s
}

type GetModelSignedUrlResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*GetModelSignedUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetModelSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelSignedUrlResponseBody) SetMessage(v string) *GetModelSignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetModelSignedUrlResponseBody) SetRequestId(v string) *GetModelSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelSignedUrlResponseBody) SetData(v []*GetModelSignedUrlResponseBodyData) *GetModelSignedUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetModelSignedUrlResponseBody) SetCode(v string) *GetModelSignedUrlResponseBody {
	s.Code = &v
	return s
}

type GetModelSignedUrlResponseBodyData struct {
	PublicUrl *string `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	Md5       *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	ModelPath *string `json:"ModelPath,omitempty" xml:"ModelPath,omitempty"`
}

func (s GetModelSignedUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetModelSignedUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetModelSignedUrlResponseBodyData) SetPublicUrl(v string) *GetModelSignedUrlResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *GetModelSignedUrlResponseBodyData) SetMd5(v string) *GetModelSignedUrlResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetModelSignedUrlResponseBodyData) SetModelPath(v string) *GetModelSignedUrlResponseBodyData {
	s.ModelPath = &v
	return s
}

type GetModelSignedUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetModelSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetModelSignedUrlResponse) SetHeaders(v map[string]*string) *GetModelSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetModelSignedUrlResponse) SetBody(v *GetModelSignedUrlResponseBody) *GetModelSignedUrlResponse {
	s.Body = v
	return s
}

type GetPreSignedUrlRequest struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s GetPreSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlRequest) SetPrefix(v string) *GetPreSignedUrlRequest {
	s.Prefix = &v
	return s
}

type GetPreSignedUrlResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPreSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlResponseBody) SetMessage(v string) *GetPreSignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetRequestId(v string) *GetPreSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetData(v string) *GetPreSignedUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetPreSignedUrlResponseBody) SetCode(v string) *GetPreSignedUrlResponseBody {
	s.Code = &v
	return s
}

type GetPreSignedUrlResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPreSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPreSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPreSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPreSignedUrlResponse) SetHeaders(v map[string]*string) *GetPreSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPreSignedUrlResponse) SetBody(v *GetPreSignedUrlResponseBody) *GetPreSignedUrlResponse {
	s.Body = v
	return s
}

type GetRuleRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRuleRequest) GoString() string {
	return s.String()
}

func (s *GetRuleRequest) SetId(v string) *GetRuleRequest {
	s.Id = &v
	return s
}

type GetRuleResponseBody struct {
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBody) SetMessage(v string) *GetRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRuleResponseBody) SetRequestId(v string) *GetRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRuleResponseBody) SetData(v *GetRuleResponseBodyData) *GetRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetRuleResponseBody) SetCode(v string) *GetRuleResponseBody {
	s.Code = &v
	return s
}

type GetRuleResponseBodyData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRuleResponseBodyData) SetCreatedAt(v string) *GetRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetRuleResponseBodyData) SetName(v string) *GetRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetRuleResponseBodyData) SetContent(v string) *GetRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetRuleResponseBodyData) SetId(v string) *GetRuleResponseBodyData {
	s.Id = &v
	return s
}

type GetRuleResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

type GetServiceConfigurationRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetServiceConfigurationRequest) SetUserId(v string) *GetServiceConfigurationRequest {
	s.UserId = &v
	return s
}

type GetServiceConfigurationResponseBody struct {
	Message   *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetServiceConfigurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetServiceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConfigurationResponseBody) SetMessage(v string) *GetServiceConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceConfigurationResponseBody) SetRequestId(v string) *GetServiceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceConfigurationResponseBody) SetData(v *GetServiceConfigurationResponseBodyData) *GetServiceConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceConfigurationResponseBody) SetCode(v string) *GetServiceConfigurationResponseBody {
	s.Code = &v
	return s
}

type GetServiceConfigurationResponseBodyData struct {
	LiveRecordAll             *bool   `json:"LiveRecordAll,omitempty" xml:"LiveRecordAll,omitempty"`
	LiveRecordLayout          *int32  `json:"LiveRecordLayout,omitempty" xml:"LiveRecordLayout,omitempty"`
	LiveRecordMaxClient       *int32  `json:"LiveRecordMaxClient,omitempty" xml:"LiveRecordMaxClient,omitempty"`
	LiveRecordEveryOne        *bool   `json:"LiveRecordEveryOne,omitempty" xml:"LiveRecordEveryOne,omitempty"`
	LiveRecordTaskProfile     *string `json:"LiveRecordTaskProfile,omitempty" xml:"LiveRecordTaskProfile,omitempty"`
	LiveRecordVideoResolution *int32  `json:"LiveRecordVideoResolution,omitempty" xml:"LiveRecordVideoResolution,omitempty"`
	TaskItemQueueSize         *int32  `json:"TaskItemQueueSize,omitempty" xml:"TaskItemQueueSize,omitempty"`
	ClientQueueSize           *int32  `json:"ClientQueueSize,omitempty" xml:"ClientQueueSize,omitempty"`
}

func (s GetServiceConfigurationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordAll(v bool) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordAll = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordLayout(v int32) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordLayout = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordMaxClient(v int32) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordMaxClient = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordEveryOne(v bool) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordEveryOne = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordTaskProfile(v string) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordTaskProfile = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetLiveRecordVideoResolution(v int32) *GetServiceConfigurationResponseBodyData {
	s.LiveRecordVideoResolution = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetTaskItemQueueSize(v int32) *GetServiceConfigurationResponseBodyData {
	s.TaskItemQueueSize = &v
	return s
}

func (s *GetServiceConfigurationResponseBodyData) SetClientQueueSize(v int32) *GetServiceConfigurationResponseBodyData {
	s.ClientQueueSize = &v
	return s
}

type GetServiceConfigurationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetServiceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetServiceConfigurationResponse) SetHeaders(v map[string]*string) *GetServiceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetServiceConfigurationResponse) SetBody(v *GetServiceConfigurationResponseBody) *GetServiceConfigurationResponse {
	s.Body = v
	return s
}

type GetSignedUrlRequest struct {
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *GetSignedUrlRequest) SetFileUrl(v string) *GetSignedUrlRequest {
	s.FileUrl = &v
	return s
}

type GetSignedUrlResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetSignedUrlResponseBody) SetMessage(v string) *GetSignedUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetSignedUrlResponseBody) SetRequestId(v string) *GetSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSignedUrlResponseBody) SetData(v string) *GetSignedUrlResponseBody {
	s.Data = &v
	return s
}

func (s *GetSignedUrlResponseBody) SetCode(v string) *GetSignedUrlResponseBody {
	s.Code = &v
	return s
}

type GetSignedUrlResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GetSignedUrlResponse) SetHeaders(v map[string]*string) *GetSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GetSignedUrlResponse) SetBody(v *GetSignedUrlResponseBody) *GetSignedUrlResponse {
	s.Body = v
	return s
}

type GetSlrConfigurationRequest struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetSlrConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSlrConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetSlrConfigurationRequest) SetUserId(v string) *GetSlrConfigurationRequest {
	s.UserId = &v
	return s
}

type GetSlrConfigurationResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetSlrConfigurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetSlrConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSlrConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSlrConfigurationResponseBody) SetMessage(v string) *GetSlrConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetSlrConfigurationResponseBody) SetRequestId(v string) *GetSlrConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSlrConfigurationResponseBody) SetData(v *GetSlrConfigurationResponseBodyData) *GetSlrConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *GetSlrConfigurationResponseBody) SetCode(v string) *GetSlrConfigurationResponseBody {
	s.Code = &v
	return s
}

type GetSlrConfigurationResponseBodyData struct {
	MqGroupId    *string   `json:"MqGroupId,omitempty" xml:"MqGroupId,omitempty"`
	MqTopic      *string   `json:"MqTopic,omitempty" xml:"MqTopic,omitempty"`
	MqInstanceId *string   `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	MqEventList  []*string `json:"MqEventList,omitempty" xml:"MqEventList,omitempty" type:"Repeated"`
	MqEndpoint   *string   `json:"MqEndpoint,omitempty" xml:"MqEndpoint,omitempty"`
	MqSubscribe  *bool     `json:"MqSubscribe,omitempty" xml:"MqSubscribe,omitempty"`
}

func (s GetSlrConfigurationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSlrConfigurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSlrConfigurationResponseBodyData) SetMqGroupId(v string) *GetSlrConfigurationResponseBodyData {
	s.MqGroupId = &v
	return s
}

func (s *GetSlrConfigurationResponseBodyData) SetMqTopic(v string) *GetSlrConfigurationResponseBodyData {
	s.MqTopic = &v
	return s
}

func (s *GetSlrConfigurationResponseBodyData) SetMqInstanceId(v string) *GetSlrConfigurationResponseBodyData {
	s.MqInstanceId = &v
	return s
}

func (s *GetSlrConfigurationResponseBodyData) SetMqEventList(v []*string) *GetSlrConfigurationResponseBodyData {
	s.MqEventList = v
	return s
}

func (s *GetSlrConfigurationResponseBodyData) SetMqEndpoint(v string) *GetSlrConfigurationResponseBodyData {
	s.MqEndpoint = &v
	return s
}

func (s *GetSlrConfigurationResponseBodyData) SetMqSubscribe(v bool) *GetSlrConfigurationResponseBodyData {
	s.MqSubscribe = &v
	return s
}

type GetSlrConfigurationResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSlrConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSlrConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSlrConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetSlrConfigurationResponse) SetHeaders(v map[string]*string) *GetSlrConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetSlrConfigurationResponse) SetBody(v *GetSlrConfigurationResponseBody) *GetSlrConfigurationResponse {
	s.Body = v
	return s
}

type GetStatisticsRequest struct {
	DateFrom     *string   `json:"DateFrom,omitempty" xml:"DateFrom,omitempty"`
	DateTo       *string   `json:"DateTo,omitempty" xml:"DateTo,omitempty"`
	DepartmentId []*string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty" type:"Repeated"`
}

func (s GetStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetStatisticsRequest) SetDateFrom(v string) *GetStatisticsRequest {
	s.DateFrom = &v
	return s
}

func (s *GetStatisticsRequest) SetDateTo(v string) *GetStatisticsRequest {
	s.DateTo = &v
	return s
}

func (s *GetStatisticsRequest) SetDepartmentId(v []*string) *GetStatisticsRequest {
	s.DepartmentId = v
	return s
}

type GetStatisticsResponseBody struct {
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStatisticsResponseBody) SetMessage(v string) *GetStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetStatisticsResponseBody) SetRequestId(v string) *GetStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStatisticsResponseBody) SetData(v *GetStatisticsResponseBodyData) *GetStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *GetStatisticsResponseBody) SetCode(v string) *GetStatisticsResponseBody {
	s.Code = &v
	return s
}

type GetStatisticsResponseBodyData struct {
	Items []*GetStatisticsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetStatisticsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStatisticsResponseBodyData) SetItems(v []*GetStatisticsResponseBodyDataItems) *GetStatisticsResponseBodyData {
	s.Items = v
	return s
}

type GetStatisticsResponseBodyDataItems struct {
	DepartmentName *string                                     `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	CloudCount     *int64                                      `json:"CloudCount,omitempty" xml:"CloudCount,omitempty"`
	Month          *string                                     `json:"Month,omitempty" xml:"Month,omitempty"`
	ClientCount    *int64                                      `json:"ClientCount,omitempty" xml:"ClientCount,omitempty"`
	Detail         []*GetStatisticsResponseBodyDataItemsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s GetStatisticsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetStatisticsResponseBodyDataItems) SetDepartmentName(v string) *GetStatisticsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItems) SetCloudCount(v int64) *GetStatisticsResponseBodyDataItems {
	s.CloudCount = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItems) SetMonth(v string) *GetStatisticsResponseBodyDataItems {
	s.Month = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItems) SetClientCount(v int64) *GetStatisticsResponseBodyDataItems {
	s.ClientCount = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItems) SetDetail(v []*GetStatisticsResponseBodyDataItemsDetail) *GetStatisticsResponseBodyDataItems {
	s.Detail = v
	return s
}

type GetStatisticsResponseBodyDataItemsDetail struct {
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	CloudCount     *int64  `json:"CloudCount,omitempty" xml:"CloudCount,omitempty"`
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Month          *int32  `json:"Month,omitempty" xml:"Month,omitempty"`
	ClientCount    *int64  `json:"ClientCount,omitempty" xml:"ClientCount,omitempty"`
}

func (s GetStatisticsResponseBodyDataItemsDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsResponseBodyDataItemsDetail) GoString() string {
	return s.String()
}

func (s *GetStatisticsResponseBodyDataItemsDetail) SetDepartmentName(v string) *GetStatisticsResponseBodyDataItemsDetail {
	s.DepartmentName = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItemsDetail) SetCloudCount(v int64) *GetStatisticsResponseBodyDataItemsDetail {
	s.CloudCount = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItemsDetail) SetDepartmentId(v string) *GetStatisticsResponseBodyDataItemsDetail {
	s.DepartmentId = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItemsDetail) SetMonth(v int32) *GetStatisticsResponseBodyDataItemsDetail {
	s.Month = &v
	return s
}

func (s *GetStatisticsResponseBodyDataItemsDetail) SetClientCount(v int64) *GetStatisticsResponseBodyDataItemsDetail {
	s.ClientCount = &v
	return s
}

type GetStatisticsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetStatisticsResponse) SetHeaders(v map[string]*string) *GetStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetStatisticsResponse) SetBody(v *GetStatisticsResponseBody) *GetStatisticsResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

type GetTaskResponseBody struct {
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

type GetTaskResponseBodyData struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	VideoUrl  *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s GetTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyData) SetStatus(v string) *GetTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyData) SetCreatedAt(v string) *GetTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTaskResponseBodyData) SetId(v string) *GetTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTaskResponseBodyData) SetVideoUrl(v string) *GetTaskResponseBodyData {
	s.VideoUrl = &v
	return s
}

type GetTaskResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTaskGroupRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetTaskGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *GetTaskGroupRequest) SetId(v string) *GetTaskGroupRequest {
	s.Id = &v
	return s
}

type GetTaskGroupResponseBody struct {
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetTaskGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetTaskGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponseBody) SetMessage(v string) *GetTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskGroupResponseBody) SetRequestId(v string) *GetTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskGroupResponseBody) SetData(v *GetTaskGroupResponseBodyData) *GetTaskGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskGroupResponseBody) SetCode(v string) *GetTaskGroupResponseBody {
	s.Code = &v
	return s
}

type GetTaskGroupResponseBodyData struct {
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	CompletedTasks *int32    `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	TotalTasks     *int32    `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskIds        []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleId         *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetTaskGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponseBodyData) SetStatus(v string) *GetTaskGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetCompletedTasks(v int32) *GetTaskGroupResponseBodyData {
	s.CompletedTasks = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetTotalTasks(v int32) *GetTaskGroupResponseBodyData {
	s.TotalTasks = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetCreatedAt(v string) *GetTaskGroupResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetName(v string) *GetTaskGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetTaskIds(v []*string) *GetTaskGroupResponseBodyData {
	s.TaskIds = v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetId(v string) *GetTaskGroupResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetRuleId(v string) *GetTaskGroupResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetTaskGroupResponseBodyData) SetRuleName(v string) *GetTaskGroupResponseBodyData {
	s.RuleName = &v
	return s
}

type GetTaskGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTaskGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTaskGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskGroupResponse) GoString() string {
	return s.String()
}

func (s *GetTaskGroupResponse) SetHeaders(v map[string]*string) *GetTaskGroupResponse {
	s.Headers = v
	return s
}

func (s *GetTaskGroupResponse) SetBody(v *GetTaskGroupResponseBody) *GetTaskGroupResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetId(v string) *GetUserRequest {
	s.Id = &v
	return s
}

type GetUserResponseBody struct {
	Message   *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

type GetUserResponseBodyData struct {
	Email       *string                               `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNumber *string                               `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Departments []*GetUserResponseBodyDataDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	CreatedAt   *string                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt   *string                               `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Source      *string                               `json:"Source,omitempty" xml:"Source,omitempty"`
	Role        *string                               `json:"Role,omitempty" xml:"Role,omitempty"`
	Name        *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Username    *string                               `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetPhoneNumber(v string) *GetUserResponseBodyData {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBodyData) SetDepartments(v []*GetUserResponseBodyDataDepartments) *GetUserResponseBodyData {
	s.Departments = v
	return s
}

func (s *GetUserResponseBodyData) SetCreatedAt(v string) *GetUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetUpdatedAt(v string) *GetUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetSource(v string) *GetUserResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetUserResponseBodyData) SetRole(v string) *GetUserResponseBodyData {
	s.Role = &v
	return s
}

func (s *GetUserResponseBodyData) SetName(v string) *GetUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyData) SetId(v string) *GetUserResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyData) SetUsername(v string) *GetUserResponseBodyData {
	s.Username = &v
	return s
}

type GetUserResponseBodyDataDepartments struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserResponseBodyDataDepartments) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyDataDepartments) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyDataDepartments) SetDescription(v string) *GetUserResponseBodyDataDepartments {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetGmtCreate(v string) *GetUserResponseBodyDataDepartments {
	s.GmtCreate = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetName(v string) *GetUserResponseBodyDataDepartments {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetGmtModified(v string) *GetUserResponseBodyDataDepartments {
	s.GmtModified = &v
	return s
}

func (s *GetUserResponseBodyDataDepartments) SetId(v string) *GetUserResponseBodyDataDepartments {
	s.Id = &v
	return s
}

type GetUserResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type InitializeServiceLinkedRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s InitializeServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkedRoleRequest) SetRoleName(v string) *InitializeServiceLinkedRoleRequest {
	s.RoleName = &v
	return s
}

type InitializeServiceLinkedRoleResponseBody struct {
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *InitializeServiceLinkedRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s InitializeServiceLinkedRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkedRoleResponseBody) SetMessage(v string) *InitializeServiceLinkedRoleResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeServiceLinkedRoleResponseBody) SetRequestId(v string) *InitializeServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeServiceLinkedRoleResponseBody) SetData(v *InitializeServiceLinkedRoleResponseBodyData) *InitializeServiceLinkedRoleResponseBody {
	s.Data = v
	return s
}

func (s *InitializeServiceLinkedRoleResponseBody) SetCode(v string) *InitializeServiceLinkedRoleResponseBody {
	s.Code = &v
	return s
}

type InitializeServiceLinkedRoleResponseBodyData struct {
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	CreateSuccess *bool   `json:"CreateSuccess,omitempty" xml:"CreateSuccess,omitempty"`
}

func (s InitializeServiceLinkedRoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InitializeServiceLinkedRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkedRoleResponseBodyData) SetErrorMessage(v string) *InitializeServiceLinkedRoleResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *InitializeServiceLinkedRoleResponseBodyData) SetCreateSuccess(v bool) *InitializeServiceLinkedRoleResponseBodyData {
	s.CreateSuccess = &v
	return s
}

type InitializeServiceLinkedRoleResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitializeServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *InitializeServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeServiceLinkedRoleResponse) SetBody(v *InitializeServiceLinkedRoleResponseBody) *InitializeServiceLinkedRoleResponse {
	s.Body = v
	return s
}

type JoinLiveRequest struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	RtcCode *string `json:"RtcCode,omitempty" xml:"RtcCode,omitempty"`
}

func (s JoinLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveRequest) GoString() string {
	return s.String()
}

func (s *JoinLiveRequest) SetUserId(v string) *JoinLiveRequest {
	s.UserId = &v
	return s
}

func (s *JoinLiveRequest) SetChannel(v string) *JoinLiveRequest {
	s.Channel = &v
	return s
}

func (s *JoinLiveRequest) SetRtcCode(v string) *JoinLiveRequest {
	s.RtcCode = &v
	return s
}

type JoinLiveResponseBody struct {
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *JoinLiveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s JoinLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBody) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBody) SetMessage(v string) *JoinLiveResponseBody {
	s.Message = &v
	return s
}

func (s *JoinLiveResponseBody) SetRequestId(v string) *JoinLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinLiveResponseBody) SetData(v *JoinLiveResponseBodyData) *JoinLiveResponseBody {
	s.Data = v
	return s
}

func (s *JoinLiveResponseBody) SetCode(v string) *JoinLiveResponseBody {
	s.Code = &v
	return s
}

type JoinLiveResponseBodyData struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	TokenData *JoinLiveResponseBodyDataTokenData `json:"TokenData,omitempty" xml:"TokenData,omitempty" type:"Struct"`
}

func (s JoinLiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBodyData) SetCode(v int32) *JoinLiveResponseBodyData {
	s.Code = &v
	return s
}

func (s *JoinLiveResponseBodyData) SetTokenData(v *JoinLiveResponseBodyDataTokenData) *JoinLiveResponseBodyData {
	s.TokenData = v
	return s
}

type JoinLiveResponseBodyDataTokenData struct {
	Turn      *JoinLiveResponseBodyDataTokenDataTurn `json:"Turn,omitempty" xml:"Turn,omitempty" type:"Struct"`
	Token     *string                                `json:"Token,omitempty" xml:"Token,omitempty"`
	AppId     *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Nonce     *string                                `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Gslb      []*string                              `json:"Gslb,omitempty" xml:"Gslb,omitempty" type:"Repeated"`
	LiveId    *string                                `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	UserId    *string                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Timestamp *int64                                 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s JoinLiveResponseBodyDataTokenData) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBodyDataTokenData) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBodyDataTokenData) SetTurn(v *JoinLiveResponseBodyDataTokenDataTurn) *JoinLiveResponseBodyDataTokenData {
	s.Turn = v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetToken(v string) *JoinLiveResponseBodyDataTokenData {
	s.Token = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetAppId(v string) *JoinLiveResponseBodyDataTokenData {
	s.AppId = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetNonce(v string) *JoinLiveResponseBodyDataTokenData {
	s.Nonce = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetGslb(v []*string) *JoinLiveResponseBodyDataTokenData {
	s.Gslb = v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetLiveId(v string) *JoinLiveResponseBodyDataTokenData {
	s.LiveId = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetUserId(v string) *JoinLiveResponseBodyDataTokenData {
	s.UserId = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenData) SetTimestamp(v int64) *JoinLiveResponseBodyDataTokenData {
	s.Timestamp = &v
	return s
}

type JoinLiveResponseBodyDataTokenDataTurn struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s JoinLiveResponseBodyDataTokenDataTurn) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponseBodyDataTokenDataTurn) GoString() string {
	return s.String()
}

func (s *JoinLiveResponseBodyDataTokenDataTurn) SetPassword(v string) *JoinLiveResponseBodyDataTokenDataTurn {
	s.Password = &v
	return s
}

func (s *JoinLiveResponseBodyDataTokenDataTurn) SetUsername(v string) *JoinLiveResponseBodyDataTokenDataTurn {
	s.Username = &v
	return s
}

type JoinLiveResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JoinLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinLiveResponse) GoString() string {
	return s.String()
}

func (s *JoinLiveResponse) SetHeaders(v map[string]*string) *JoinLiveResponse {
	s.Headers = v
	return s
}

func (s *JoinLiveResponse) SetBody(v *JoinLiveResponseBody) *JoinLiveResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetPageIndex(v int32) *ListAppsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListAppsRequest) SetPageSize(v int32) *ListAppsRequest {
	s.PageSize = &v
	return s
}

type ListAppsResponseBody struct {
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) SetMessage(v string) *ListAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetData(v *ListAppsResponseBodyData) *ListAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListAppsResponseBody) SetCode(v string) *ListAppsResponseBody {
	s.Code = &v
	return s
}

type ListAppsResponseBodyData struct {
	Items         []*ListAppsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                           `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                           `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyData) SetItems(v []*ListAppsResponseBodyDataItems) *ListAppsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListAppsResponseBodyData) SetTotalPages(v int32) *ListAppsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListAppsResponseBodyData) SetTotalElements(v int64) *ListAppsResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListAppsResponseBodyDataItems struct {
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	PackageName    *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Disabled       *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListAppsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyDataItems) SetDepartmentName(v string) *ListAppsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetPackageName(v string) *ListAppsResponseBodyDataItems {
	s.PackageName = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetCreatedAt(v string) *ListAppsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDepartmentId(v string) *ListAppsResponseBodyDataItems {
	s.DepartmentId = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetDisabled(v bool) *ListAppsResponseBodyDataItems {
	s.Disabled = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetName(v string) *ListAppsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListAppsResponseBodyDataItems) SetId(v string) *ListAppsResponseBodyDataItems {
	s.Id = &v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetBody(v *ListAppsResponseBody) *ListAppsResponse {
	s.Body = v
	return s
}

type ListDepartmentsRequest struct {
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *ListDepartmentsRequest) SetName(v string) *ListDepartmentsRequest {
	s.Name = &v
	return s
}

func (s *ListDepartmentsRequest) SetPageIndex(v int32) *ListDepartmentsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDepartmentsRequest) SetPageSize(v int32) *ListDepartmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDepartmentsRequest) SetUserId(v string) *ListDepartmentsRequest {
	s.UserId = &v
	return s
}

type ListDepartmentsResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListDepartmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBody) SetMessage(v string) *ListDepartmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDepartmentsResponseBody) SetRequestId(v string) *ListDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDepartmentsResponseBody) SetData(v *ListDepartmentsResponseBodyData) *ListDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDepartmentsResponseBody) SetCode(v string) *ListDepartmentsResponseBody {
	s.Code = &v
	return s
}

type ListDepartmentsResponseBodyData struct {
	Items         []*ListDepartmentsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                                  `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                                  `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListDepartmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyData) SetItems(v []*ListDepartmentsResponseBodyDataItems) *ListDepartmentsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDepartmentsResponseBodyData) SetTotalPages(v int32) *ListDepartmentsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDepartmentsResponseBodyData) SetTotalElements(v int64) *ListDepartmentsResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListDepartmentsResponseBodyDataItems struct {
	Description    *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Administrators []*ListDepartmentsResponseBodyDataItemsAdministrators `json:"Administrators,omitempty" xml:"Administrators,omitempty" type:"Repeated"`
	CreatedAt      *string                                               `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt      *string                                               `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name           *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDepartmentsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyDataItems) SetDescription(v string) *ListDepartmentsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetAdministrators(v []*ListDepartmentsResponseBodyDataItemsAdministrators) *ListDepartmentsResponseBodyDataItems {
	s.Administrators = v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetCreatedAt(v string) *ListDepartmentsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetUpdatedAt(v string) *ListDepartmentsResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetName(v string) *ListDepartmentsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItems) SetId(v string) *ListDepartmentsResponseBodyDataItems {
	s.Id = &v
	return s
}

type ListDepartmentsResponseBodyDataItemsAdministrators struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDepartmentsResponseBodyDataItemsAdministrators) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponseBodyDataItemsAdministrators) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponseBodyDataItemsAdministrators) SetName(v string) *ListDepartmentsResponseBodyDataItemsAdministrators {
	s.Name = &v
	return s
}

func (s *ListDepartmentsResponseBodyDataItemsAdministrators) SetId(v string) *ListDepartmentsResponseBodyDataItemsAdministrators {
	s.Id = &v
	return s
}

type ListDepartmentsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *ListDepartmentsResponse) SetHeaders(v map[string]*string) *ListDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *ListDepartmentsResponse) SetBody(v *ListDepartmentsResponseBody) *ListDepartmentsResponse {
	s.Body = v
	return s
}

type ListDetectionsRequest struct {
	CreateDateFrom *string `json:"CreateDateFrom,omitempty" xml:"CreateDateFrom,omitempty"`
	CreateDateTo   *string `json:"CreateDateTo,omitempty" xml:"CreateDateTo,omitempty"`
	DepartmentId   *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	PageIndex      *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecordingType  *string `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	RuleId         *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListDetectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionsRequest) SetCreateDateFrom(v string) *ListDetectionsRequest {
	s.CreateDateFrom = &v
	return s
}

func (s *ListDetectionsRequest) SetCreateDateTo(v string) *ListDetectionsRequest {
	s.CreateDateTo = &v
	return s
}

func (s *ListDetectionsRequest) SetDepartmentId(v string) *ListDetectionsRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListDetectionsRequest) SetPageIndex(v int32) *ListDetectionsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDetectionsRequest) SetPageSize(v int32) *ListDetectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectionsRequest) SetRecordingType(v string) *ListDetectionsRequest {
	s.RecordingType = &v
	return s
}

func (s *ListDetectionsRequest) SetRuleId(v string) *ListDetectionsRequest {
	s.RuleId = &v
	return s
}

type ListDetectionsResponseBody struct {
	Errors    []*ListDetectionsResponseBodyErrors `json:"Errors,omitempty" xml:"Errors,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListDetectionsResponseBodyData     `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDetectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBody) SetErrors(v []*ListDetectionsResponseBodyErrors) *ListDetectionsResponseBody {
	s.Errors = v
	return s
}

func (s *ListDetectionsResponseBody) SetMessage(v string) *ListDetectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDetectionsResponseBody) SetRequestId(v string) *ListDetectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetectionsResponseBody) SetData(v *ListDetectionsResponseBodyData) *ListDetectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListDetectionsResponseBody) SetCode(v string) *ListDetectionsResponseBody {
	s.Code = &v
	return s
}

type ListDetectionsResponseBodyErrors struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Field   *string `json:"Field,omitempty" xml:"Field,omitempty"`
}

func (s ListDetectionsResponseBodyErrors) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyErrors) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyErrors) SetMessage(v string) *ListDetectionsResponseBodyErrors {
	s.Message = &v
	return s
}

func (s *ListDetectionsResponseBodyErrors) SetField(v string) *ListDetectionsResponseBodyErrors {
	s.Field = &v
	return s
}

type ListDetectionsResponseBodyData struct {
	Items         []*ListDetectionsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                                 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                                 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListDetectionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyData) SetItems(v []*ListDetectionsResponseBodyDataItems) *ListDetectionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDetectionsResponseBodyData) SetTotalPages(v int32) *ListDetectionsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDetectionsResponseBodyData) SetTotalElements(v int64) *ListDetectionsResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListDetectionsResponseBodyDataItems struct {
	Status         *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	DepartmentName *string                                     `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	Tasks          []*ListDetectionsResponseBodyDataItemsTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	RecordingType  *string                                     `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	CreatedAt      *string                                     `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DepartmentId   *string                                     `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Id             *string                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleName       *string                                     `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId         *string                                     `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListDetectionsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyDataItems) SetStatus(v string) *ListDetectionsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetDepartmentName(v string) *ListDetectionsResponseBodyDataItems {
	s.DepartmentName = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetTasks(v []*ListDetectionsResponseBodyDataItemsTasks) *ListDetectionsResponseBodyDataItems {
	s.Tasks = v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRecordingType(v string) *ListDetectionsResponseBodyDataItems {
	s.RecordingType = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetCreatedAt(v string) *ListDetectionsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetDepartmentId(v string) *ListDetectionsResponseBodyDataItems {
	s.DepartmentId = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetId(v string) *ListDetectionsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRuleName(v string) *ListDetectionsResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItems) SetRuleId(v string) *ListDetectionsResponseBodyDataItems {
	s.RuleId = &v
	return s
}

type ListDetectionsResponseBodyDataItemsTasks struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListDetectionsResponseBodyDataItemsTasks) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponseBodyDataItemsTasks) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetStatus(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.Status = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetVideoMetaUrl(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.VideoMetaUrl = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetCreatedAt(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetId(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.Id = &v
	return s
}

func (s *ListDetectionsResponseBodyDataItemsTasks) SetVideoUrl(v string) *ListDetectionsResponseBodyDataItemsTasks {
	s.VideoUrl = &v
	return s
}

type ListDetectionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDetectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDetectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDetectionsResponse) SetHeaders(v map[string]*string) *ListDetectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDetectionsResponse) SetBody(v *ListDetectionsResponseBody) *ListDetectionsResponse {
	s.Body = v
	return s
}

type ListDetectProcessesRequest struct {
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishStatus *bool   `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
}

func (s ListDetectProcessesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesRequest) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesRequest) SetPageIndex(v int32) *ListDetectProcessesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDetectProcessesRequest) SetPageSize(v int32) *ListDetectProcessesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectProcessesRequest) SetName(v string) *ListDetectProcessesRequest {
	s.Name = &v
	return s
}

func (s *ListDetectProcessesRequest) SetPublishStatus(v bool) *ListDetectProcessesRequest {
	s.PublishStatus = &v
	return s
}

type ListDetectProcessesResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListDetectProcessesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListDetectProcessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBody) SetMessage(v string) *ListDetectProcessesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDetectProcessesResponseBody) SetRequestId(v string) *ListDetectProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetectProcessesResponseBody) SetData(v *ListDetectProcessesResponseBodyData) *ListDetectProcessesResponseBody {
	s.Data = v
	return s
}

func (s *ListDetectProcessesResponseBody) SetCode(v string) *ListDetectProcessesResponseBody {
	s.Code = &v
	return s
}

type ListDetectProcessesResponseBodyData struct {
	Items         []*ListDetectProcessesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                                      `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                                      `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListDetectProcessesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBodyData) SetItems(v []*ListDetectProcessesResponseBodyDataItems) *ListDetectProcessesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDetectProcessesResponseBodyData) SetTotalPages(v int32) *ListDetectProcessesResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDetectProcessesResponseBodyData) SetTotalElements(v int64) *ListDetectProcessesResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListDetectProcessesResponseBodyDataItems struct {
	Draft      *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	FileUrl    *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	CreatedAt  *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Md5        *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	UpdatedAt  *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	NewVersion *bool   `json:"NewVersion,omitempty" xml:"NewVersion,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListDetectProcessesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponseBodyDataItems) SetDraft(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Draft = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetFileUrl(v string) *ListDetectProcessesResponseBodyDataItems {
	s.FileUrl = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetCreatedAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetMd5(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Md5 = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetUpdatedAt(v string) *ListDetectProcessesResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetNewVersion(v bool) *ListDetectProcessesResponseBodyDataItems {
	s.NewVersion = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetName(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetContent(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *ListDetectProcessesResponseBodyDataItems) SetId(v string) *ListDetectProcessesResponseBodyDataItems {
	s.Id = &v
	return s
}

type ListDetectProcessesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDetectProcessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDetectProcessesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDetectProcessesResponse) GoString() string {
	return s.String()
}

func (s *ListDetectProcessesResponse) SetHeaders(v map[string]*string) *ListDetectProcessesResponse {
	s.Headers = v
	return s
}

func (s *ListDetectProcessesResponse) SetBody(v *ListDetectProcessesResponseBody) *ListDetectProcessesResponse {
	s.Body = v
	return s
}

type ListFilesRequest struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Limit  *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFilesRequest) GoString() string {
	return s.String()
}

func (s *ListFilesRequest) SetPrefix(v string) *ListFilesRequest {
	s.Prefix = &v
	return s
}

func (s *ListFilesRequest) SetLimit(v int32) *ListFilesRequest {
	s.Limit = &v
	return s
}

type ListFilesResponseBody struct {
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFilesResponseBody) SetMessage(v string) *ListFilesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFilesResponseBody) SetRequestId(v string) *ListFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFilesResponseBody) SetData(v []*string) *ListFilesResponseBody {
	s.Data = v
	return s
}

func (s *ListFilesResponseBody) SetCode(v string) *ListFilesResponseBody {
	s.Code = &v
	return s
}

type ListFilesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFilesResponse) GoString() string {
	return s.String()
}

func (s *ListFilesResponse) SetHeaders(v map[string]*string) *ListFilesResponse {
	s.Headers = v
	return s
}

func (s *ListFilesResponse) SetBody(v *ListFilesResponseBody) *ListFilesResponse {
	s.Body = v
	return s
}

type ListLivesRequest struct {
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLivesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLivesRequest) GoString() string {
	return s.String()
}

func (s *ListLivesRequest) SetPageIndex(v int32) *ListLivesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListLivesRequest) SetPageSize(v int32) *ListLivesRequest {
	s.PageSize = &v
	return s
}

type ListLivesResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListLivesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListLivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLivesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivesResponseBody) SetMessage(v string) *ListLivesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLivesResponseBody) SetRequestId(v string) *ListLivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivesResponseBody) SetData(v *ListLivesResponseBodyData) *ListLivesResponseBody {
	s.Data = v
	return s
}

func (s *ListLivesResponseBody) SetCode(v string) *ListLivesResponseBody {
	s.Code = &v
	return s
}

type ListLivesResponseBodyData struct {
	Items         []*ListLivesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                            `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListLivesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLivesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLivesResponseBodyData) SetItems(v []*ListLivesResponseBodyDataItems) *ListLivesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListLivesResponseBodyData) SetTotalPages(v int32) *ListLivesResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListLivesResponseBodyData) SetTotalElements(v int64) *ListLivesResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListLivesResponseBodyDataItems struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Channel   *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	PublicId  *string `json:"PublicId,omitempty" xml:"PublicId,omitempty"`
}

func (s ListLivesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListLivesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListLivesResponseBodyDataItems) SetStatus(v string) *ListLivesResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetUserId(v string) *ListLivesResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetCreatedAt(v string) *ListLivesResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetChannel(v string) *ListLivesResponseBodyDataItems {
	s.Channel = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetName(v string) *ListLivesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetId(v string) *ListLivesResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListLivesResponseBodyDataItems) SetPublicId(v string) *ListLivesResponseBodyDataItems {
	s.PublicId = &v
	return s
}

type ListLivesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLivesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLivesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLivesResponse) GoString() string {
	return s.String()
}

func (s *ListLivesResponse) SetHeaders(v map[string]*string) *ListLivesResponse {
	s.Headers = v
	return s
}

func (s *ListLivesResponse) SetBody(v *ListLivesResponseBody) *ListLivesResponse {
	s.Body = v
	return s
}

type ListRolesResponseBody struct {
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetData(v []*string) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

type ListRolesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListRulesRequest struct {
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRulesRequest) GoString() string {
	return s.String()
}

func (s *ListRulesRequest) SetPageIndex(v int32) *ListRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRulesRequest) SetPageSize(v int32) *ListRulesRequest {
	s.PageSize = &v
	return s
}

type ListRulesResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v *ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

type ListRulesResponseBodyData struct {
	Items         []*ListRulesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                            `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) SetItems(v []*ListRulesResponseBodyDataItems) *ListRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListRulesResponseBodyData) SetTotalPages(v int32) *ListRulesResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListRulesResponseBodyData) SetTotalElements(v int64) *ListRulesResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListRulesResponseBodyDataItems struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListRulesResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyDataItems) SetCreatedAt(v string) *ListRulesResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetName(v string) *ListRulesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetContent(v string) *ListRulesResponseBodyDataItems {
	s.Content = &v
	return s
}

func (s *ListRulesResponseBodyDataItems) SetId(v string) *ListRulesResponseBodyDataItems {
	s.Id = &v
	return s
}

type ListRulesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRulesResponse) SetHeaders(v map[string]*string) *ListRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRulesResponse) SetBody(v *ListRulesResponseBody) *ListRulesResponse {
	s.Body = v
	return s
}

type ListStatisticsTaskRequest struct {
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListStatisticsTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticsTaskRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticsTaskRequest) SetPageIndex(v int32) *ListStatisticsTaskRequest {
	s.PageIndex = &v
	return s
}

func (s *ListStatisticsTaskRequest) SetPageSize(v int32) *ListStatisticsTaskRequest {
	s.PageSize = &v
	return s
}

type ListStatisticsTaskResponseBody struct {
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListStatisticsTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListStatisticsTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticsTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticsTaskResponseBody) SetMessage(v string) *ListStatisticsTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListStatisticsTaskResponseBody) SetRequestId(v string) *ListStatisticsTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticsTaskResponseBody) SetData(v *ListStatisticsTaskResponseBodyData) *ListStatisticsTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListStatisticsTaskResponseBody) SetCode(v string) *ListStatisticsTaskResponseBody {
	s.Code = &v
	return s
}

type ListStatisticsTaskResponseBodyData struct {
	Items         []*ListStatisticsTaskResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                                     `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                                     `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListStatisticsTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticsTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListStatisticsTaskResponseBodyData) SetItems(v []*ListStatisticsTaskResponseBodyDataItems) *ListStatisticsTaskResponseBodyData {
	s.Items = v
	return s
}

func (s *ListStatisticsTaskResponseBodyData) SetTotalPages(v int32) *ListStatisticsTaskResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListStatisticsTaskResponseBodyData) SetTotalElements(v int64) *ListStatisticsTaskResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListStatisticsTaskResponseBodyDataItems struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	FileUrl   *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListStatisticsTaskResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticsTaskResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListStatisticsTaskResponseBodyDataItems) SetStatus(v int32) *ListStatisticsTaskResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListStatisticsTaskResponseBodyDataItems) SetFileUrl(v string) *ListStatisticsTaskResponseBodyDataItems {
	s.FileUrl = &v
	return s
}

func (s *ListStatisticsTaskResponseBodyDataItems) SetCreatedAt(v string) *ListStatisticsTaskResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListStatisticsTaskResponseBodyDataItems) SetName(v string) *ListStatisticsTaskResponseBodyDataItems {
	s.Name = &v
	return s
}

type ListStatisticsTaskResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListStatisticsTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatisticsTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticsTaskResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticsTaskResponse) SetHeaders(v map[string]*string) *ListStatisticsTaskResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticsTaskResponse) SetBody(v *ListStatisticsTaskResponseBody) *ListStatisticsTaskResponse {
	s.Body = v
	return s
}

type ListTaskGroupsRequest struct {
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTaskGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsRequest) SetPageIndex(v int32) *ListTaskGroupsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTaskGroupsRequest) SetPageSize(v int32) *ListTaskGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTaskGroupsRequest) SetStatus(v string) *ListTaskGroupsRequest {
	s.Status = &v
	return s
}

type ListTaskGroupsResponseBody struct {
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTaskGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTaskGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBody) SetMessage(v string) *ListTaskGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskGroupsResponseBody) SetRequestId(v string) *ListTaskGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskGroupsResponseBody) SetData(v *ListTaskGroupsResponseBodyData) *ListTaskGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskGroupsResponseBody) SetCode(v string) *ListTaskGroupsResponseBody {
	s.Code = &v
	return s
}

type ListTaskGroupsResponseBodyData struct {
	Items         []*ListTaskGroupsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                                 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                                 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListTaskGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBodyData) SetItems(v []*ListTaskGroupsResponseBodyDataItems) *ListTaskGroupsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListTaskGroupsResponseBodyData) SetTotalPages(v int32) *ListTaskGroupsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListTaskGroupsResponseBodyData) SetTotalElements(v int64) *ListTaskGroupsResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListTaskGroupsResponseBodyDataItems struct {
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	CompletedTasks *int32    `json:"CompletedTasks,omitempty" xml:"CompletedTasks,omitempty"`
	TotalTasks     *int32    `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
	CreatedAt      *string   `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskIds        []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	Id             *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RuleId         *string   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName       *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListTaskGroupsResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponseBodyDataItems) SetStatus(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetCompletedTasks(v int32) *ListTaskGroupsResponseBodyDataItems {
	s.CompletedTasks = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetTotalTasks(v int32) *ListTaskGroupsResponseBodyDataItems {
	s.TotalTasks = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetCreatedAt(v string) *ListTaskGroupsResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetName(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetTaskIds(v []*string) *ListTaskGroupsResponseBodyDataItems {
	s.TaskIds = v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetId(v string) *ListTaskGroupsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetRuleId(v string) *ListTaskGroupsResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListTaskGroupsResponseBodyDataItems) SetRuleName(v string) *ListTaskGroupsResponseBodyDataItems {
	s.RuleName = &v
	return s
}

type ListTaskGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskGroupsResponse) SetHeaders(v map[string]*string) *ListTaskGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskGroupsResponse) SetBody(v *ListTaskGroupsResponseBody) *ListTaskGroupsResponse {
	s.Body = v
	return s
}

type ListTaskItemsRequest struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListTaskItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskItemsRequest) SetTaskId(v string) *ListTaskItemsRequest {
	s.TaskId = &v
	return s
}

type ListTaskItemsResponseBody struct {
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*ListTaskItemsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTaskItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponseBody) SetMessage(v string) *ListTaskItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskItemsResponseBody) SetRequestId(v string) *ListTaskItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskItemsResponseBody) SetData(v []*ListTaskItemsResponseBodyData) *ListTaskItemsResponseBody {
	s.Data = v
	return s
}

func (s *ListTaskItemsResponseBody) SetCode(v string) *ListTaskItemsResponseBody {
	s.Code = &v
	return s
}

type ListTaskItemsResponseBodyData struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	CreatedAt  *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Output     *string `json:"Output,omitempty" xml:"Output,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SegmentSeq *int64  `json:"SegmentSeq,omitempty" xml:"SegmentSeq,omitempty"`
}

func (s ListTaskItemsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponseBodyData) SetStatus(v string) *ListTaskItemsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetCreatedAt(v string) *ListTaskItemsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetMessage(v string) *ListTaskItemsResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetOutput(v string) *ListTaskItemsResponseBodyData {
	s.Output = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetName(v string) *ListTaskItemsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTaskItemsResponseBodyData) SetSegmentSeq(v int64) *ListTaskItemsResponseBodyData {
	s.SegmentSeq = &v
	return s
}

type ListTaskItemsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTaskItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTaskItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskItemsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskItemsResponse) SetHeaders(v map[string]*string) *ListTaskItemsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskItemsResponse) SetBody(v *ListTaskItemsResponseBody) *ListTaskItemsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	PageIndex   *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskGroupId *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetPageIndex(v int32) *ListTasksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetTaskGroupId(v string) *ListTasksRequest {
	s.TaskGroupId = &v
	return s
}

type ListTasksResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetMessage(v string) *ListTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetData(v *ListTasksResponseBodyData) *ListTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTasksResponseBody) SetCode(v string) *ListTasksResponseBody {
	s.Code = &v
	return s
}

type ListTasksResponseBodyData struct {
	Items         []*ListTasksResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                            `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyData) SetItems(v []*ListTasksResponseBodyDataItems) *ListTasksResponseBodyData {
	s.Items = v
	return s
}

func (s *ListTasksResponseBodyData) SetTotalPages(v int32) *ListTasksResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTotalElements(v int64) *ListTasksResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListTasksResponseBodyDataItems struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VideoMetaUrl *string `json:"VideoMetaUrl,omitempty" xml:"VideoMetaUrl,omitempty"`
	CreatedAt    *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	VideoUrl     *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ListTasksResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyDataItems) SetStatus(v string) *ListTasksResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetVideoMetaUrl(v string) *ListTasksResponseBodyDataItems {
	s.VideoMetaUrl = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetCreatedAt(v string) *ListTasksResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetId(v string) *ListTasksResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListTasksResponseBodyDataItems) SetVideoUrl(v string) *ListTasksResponseBodyDataItems {
	s.VideoUrl = &v
	return s
}

type ListTasksResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	PageIndex    *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetDepartmentId(v string) *ListUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListUsersRequest) SetPageIndex(v int32) *ListUsersRequest {
	s.PageIndex = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetUsername(v string) *ListUsersRequest {
	s.Username = &v
	return s
}

type ListUsersResponseBody struct {
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetMessage(v string) *ListUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetCode(v string) *ListUsersResponseBody {
	s.Code = &v
	return s
}

type ListUsersResponseBodyData struct {
	Items         []*ListUsersResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	TotalPages    *int32                            `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	TotalElements *int64                            `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetItems(v []*ListUsersResponseBodyDataItems) *ListUsersResponseBodyData {
	s.Items = v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalPages(v int32) *ListUsersResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalElements(v int64) *ListUsersResponseBodyData {
	s.TotalElements = &v
	return s
}

type ListUsersResponseBodyDataItems struct {
	RamUsername *string                                      `json:"RamUsername,omitempty" xml:"RamUsername,omitempty"`
	Email       *string                                      `json:"Email,omitempty" xml:"Email,omitempty"`
	PhoneNumber *string                                      `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Departments []*ListUsersResponseBodyDataItemsDepartments `json:"Departments,omitempty" xml:"Departments,omitempty" type:"Repeated"`
	CreatedAt   *string                                      `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt   *string                                      `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Source      *string                                      `json:"Source,omitempty" xml:"Source,omitempty"`
	Role        *string                                      `json:"Role,omitempty" xml:"Role,omitempty"`
	Name        *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Username    *string                                      `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyDataItems) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataItems) SetRamUsername(v string) *ListUsersResponseBodyDataItems {
	s.RamUsername = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetEmail(v string) *ListUsersResponseBodyDataItems {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetPhoneNumber(v string) *ListUsersResponseBodyDataItems {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetDepartments(v []*ListUsersResponseBodyDataItemsDepartments) *ListUsersResponseBodyDataItems {
	s.Departments = v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetCreatedAt(v string) *ListUsersResponseBodyDataItems {
	s.CreatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetUpdatedAt(v string) *ListUsersResponseBodyDataItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetSource(v string) *ListUsersResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetRole(v string) *ListUsersResponseBodyDataItems {
	s.Role = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetName(v string) *ListUsersResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetId(v string) *ListUsersResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyDataItems) SetUsername(v string) *ListUsersResponseBodyDataItems {
	s.Username = &v
	return s
}

type ListUsersResponseBodyDataItemsDepartments struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListUsersResponseBodyDataItemsDepartments) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataItemsDepartments) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetDescription(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetCreatedAt(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.CreatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetUpdatedAt(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.UpdatedAt = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetName(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyDataItemsDepartments) SetId(v string) *ListUsersResponseBodyDataItemsDepartments {
	s.Id = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type RenameDetectProcessRequest struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RenameDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessRequest) SetId(v string) *RenameDetectProcessRequest {
	s.Id = &v
	return s
}

func (s *RenameDetectProcessRequest) SetName(v string) *RenameDetectProcessRequest {
	s.Name = &v
	return s
}

type RenameDetectProcessResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *RenameDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RenameDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponseBody) SetMessage(v string) *RenameDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *RenameDetectProcessResponseBody) SetRequestId(v string) *RenameDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameDetectProcessResponseBody) SetData(v *RenameDetectProcessResponseBodyData) *RenameDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *RenameDetectProcessResponseBody) SetCode(v string) *RenameDetectProcessResponseBody {
	s.Code = &v
	return s
}

type RenameDetectProcessResponseBodyData struct {
	Draft     *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Md5       *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RenameDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponseBodyData) SetDraft(v string) *RenameDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetCreatedAt(v string) *RenameDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetMd5(v string) *RenameDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetName(v string) *RenameDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetContent(v string) *RenameDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *RenameDetectProcessResponseBodyData) SetId(v string) *RenameDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

type RenameDetectProcessResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *RenameDetectProcessResponse) SetHeaders(v map[string]*string) *RenameDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *RenameDetectProcessResponse) SetBody(v *RenameDetectProcessResponseBody) *RenameDetectProcessResponse {
	s.Body = v
	return s
}

type UpdateAppRequest struct {
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Disabled     *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	PackageName  *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
}

func (s UpdateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppRequest) SetId(v string) *UpdateAppRequest {
	s.Id = &v
	return s
}

func (s *UpdateAppRequest) SetName(v string) *UpdateAppRequest {
	s.Name = &v
	return s
}

func (s *UpdateAppRequest) SetDisabled(v bool) *UpdateAppRequest {
	s.Disabled = &v
	return s
}

func (s *UpdateAppRequest) SetPackageName(v string) *UpdateAppRequest {
	s.PackageName = &v
	return s
}

func (s *UpdateAppRequest) SetDepartmentId(v string) *UpdateAppRequest {
	s.DepartmentId = &v
	return s
}

type UpdateAppResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAppResponseBody) SetMessage(v string) *UpdateAppResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAppResponseBody) SetRequestId(v string) *UpdateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAppResponseBody) SetData(v map[string]interface{}) *UpdateAppResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAppResponseBody) SetCode(v string) *UpdateAppResponseBody {
	s.Code = &v
	return s
}

type UpdateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppResponse) SetHeaders(v map[string]*string) *UpdateAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppResponse) SetBody(v *UpdateAppResponseBody) *UpdateAppResponse {
	s.Body = v
	return s
}

type UpdateDepartmentRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentRequest) SetDescription(v string) *UpdateDepartmentRequest {
	s.Description = &v
	return s
}

func (s *UpdateDepartmentRequest) SetLabel(v string) *UpdateDepartmentRequest {
	s.Label = &v
	return s
}

func (s *UpdateDepartmentRequest) SetName(v string) *UpdateDepartmentRequest {
	s.Name = &v
	return s
}

func (s *UpdateDepartmentRequest) SetId(v string) *UpdateDepartmentRequest {
	s.Id = &v
	return s
}

type UpdateDepartmentResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponseBody) SetMessage(v string) *UpdateDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetRequestId(v string) *UpdateDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDepartmentResponseBody) SetData(v map[string]interface{}) *UpdateDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDepartmentResponseBody) SetCode(v string) *UpdateDepartmentResponseBody {
	s.Code = &v
	return s
}

type UpdateDepartmentResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponse) SetHeaders(v map[string]*string) *UpdateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDepartmentResponse) SetBody(v *UpdateDepartmentResponseBody) *UpdateDepartmentResponse {
	s.Body = v
	return s
}

type UpdateDetectProcessRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Draft   *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateDetectProcessRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessRequest) SetId(v string) *UpdateDetectProcessRequest {
	s.Id = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetName(v string) *UpdateDetectProcessRequest {
	s.Name = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetDraft(v string) *UpdateDetectProcessRequest {
	s.Draft = &v
	return s
}

func (s *UpdateDetectProcessRequest) SetContent(v string) *UpdateDetectProcessRequest {
	s.Content = &v
	return s
}

type UpdateDetectProcessResponseBody struct {
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateDetectProcessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateDetectProcessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponseBody) SetMessage(v string) *UpdateDetectProcessResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetRequestId(v string) *UpdateDetectProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetData(v *UpdateDetectProcessResponseBodyData) *UpdateDetectProcessResponseBody {
	s.Data = v
	return s
}

func (s *UpdateDetectProcessResponseBody) SetCode(v string) *UpdateDetectProcessResponseBody {
	s.Code = &v
	return s
}

type UpdateDetectProcessResponseBodyData struct {
	Draft     *string `json:"Draft,omitempty" xml:"Draft,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Md5       *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDetectProcessResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponseBodyData) SetDraft(v string) *UpdateDetectProcessResponseBodyData {
	s.Draft = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetCreatedAt(v string) *UpdateDetectProcessResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetMd5(v string) *UpdateDetectProcessResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetName(v string) *UpdateDetectProcessResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetContent(v string) *UpdateDetectProcessResponseBodyData {
	s.Content = &v
	return s
}

func (s *UpdateDetectProcessResponseBodyData) SetId(v string) *UpdateDetectProcessResponseBodyData {
	s.Id = &v
	return s
}

type UpdateDetectProcessResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDetectProcessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDetectProcessResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDetectProcessResponse) GoString() string {
	return s.String()
}

func (s *UpdateDetectProcessResponse) SetHeaders(v map[string]*string) *UpdateDetectProcessResponse {
	s.Headers = v
	return s
}

func (s *UpdateDetectProcessResponse) SetBody(v *UpdateDetectProcessResponseBody) *UpdateDetectProcessResponse {
	s.Body = v
	return s
}

type UpdateLiveRequest struct {
	LiveId *string `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveRequest) SetLiveId(v string) *UpdateLiveRequest {
	s.LiveId = &v
	return s
}

func (s *UpdateLiveRequest) SetStatus(v string) *UpdateLiveRequest {
	s.Status = &v
	return s
}

func (s *UpdateLiveRequest) SetUserId(v string) *UpdateLiveRequest {
	s.UserId = &v
	return s
}

type UpdateLiveResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateLiveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateLiveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBody) SetMessage(v string) *UpdateLiveResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLiveResponseBody) SetRequestId(v string) *UpdateLiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveResponseBody) SetData(v *UpdateLiveResponseBodyData) *UpdateLiveResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLiveResponseBody) SetCode(v string) *UpdateLiveResponseBody {
	s.Code = &v
	return s
}

type UpdateLiveResponseBodyData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateLiveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponseBodyData) SetCreatedAt(v string) *UpdateLiveResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateLiveResponseBodyData) SetId(v string) *UpdateLiveResponseBodyData {
	s.Id = &v
	return s
}

type UpdateLiveResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLiveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLiveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLiveResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveResponse) SetHeaders(v map[string]*string) *UpdateLiveResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveResponse) SetBody(v *UpdateLiveResponseBody) *UpdateLiveResponse {
	s.Body = v
	return s
}

type UpdateRuleRequest struct {
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRuleRequest) SetId(v string) *UpdateRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateRuleRequest) SetName(v string) *UpdateRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateRuleRequest) SetContent(v string) *UpdateRuleRequest {
	s.Content = &v
	return s
}

type UpdateRuleResponseBody struct {
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UpdateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v *UpdateRuleResponseBodyData) *UpdateRuleResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

type UpdateRuleResponseBodyData struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBodyData) SetCreatedAt(v string) *UpdateRuleResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetName(v string) *UpdateRuleResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetContent(v string) *UpdateRuleResponseBodyData {
	s.Content = &v
	return s
}

func (s *UpdateRuleResponseBodyData) SetId(v string) *UpdateRuleResponseBodyData {
	s.Id = &v
	return s
}

type UpdateRuleResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponse) SetHeaders(v map[string]*string) *UpdateRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRuleResponse) SetBody(v *UpdateRuleResponseBody) *UpdateRuleResponse {
	s.Body = v
	return s
}

type UpdateServiceConfigurationRequest struct {
	TaskItemQueueSize         *int32  `json:"TaskItemQueueSize,omitempty" xml:"TaskItemQueueSize,omitempty"`
	ClientQueueSize           *int32  `json:"ClientQueueSize,omitempty" xml:"ClientQueueSize,omitempty"`
	LiveRecordEveryOne        *bool   `json:"LiveRecordEveryOne,omitempty" xml:"LiveRecordEveryOne,omitempty"`
	LiveRecordAll             *bool   `json:"LiveRecordAll,omitempty" xml:"LiveRecordAll,omitempty"`
	LiveRecordLayout          *int32  `json:"LiveRecordLayout,omitempty" xml:"LiveRecordLayout,omitempty"`
	LiveRecordTaskProfile     *string `json:"LiveRecordTaskProfile,omitempty" xml:"LiveRecordTaskProfile,omitempty"`
	LiveRecordMaxClient       *int32  `json:"LiveRecordMaxClient,omitempty" xml:"LiveRecordMaxClient,omitempty"`
	LiveRecordVideoResolution *int32  `json:"LiveRecordVideoResolution,omitempty" xml:"LiveRecordVideoResolution,omitempty"`
}

func (s UpdateServiceConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigurationRequest) SetTaskItemQueueSize(v int32) *UpdateServiceConfigurationRequest {
	s.TaskItemQueueSize = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetClientQueueSize(v int32) *UpdateServiceConfigurationRequest {
	s.ClientQueueSize = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordEveryOne(v bool) *UpdateServiceConfigurationRequest {
	s.LiveRecordEveryOne = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordAll(v bool) *UpdateServiceConfigurationRequest {
	s.LiveRecordAll = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordLayout(v int32) *UpdateServiceConfigurationRequest {
	s.LiveRecordLayout = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordTaskProfile(v string) *UpdateServiceConfigurationRequest {
	s.LiveRecordTaskProfile = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordMaxClient(v int32) *UpdateServiceConfigurationRequest {
	s.LiveRecordMaxClient = &v
	return s
}

func (s *UpdateServiceConfigurationRequest) SetLiveRecordVideoResolution(v int32) *UpdateServiceConfigurationRequest {
	s.LiveRecordVideoResolution = &v
	return s
}

type UpdateServiceConfigurationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateServiceConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigurationResponseBody) SetMessage(v string) *UpdateServiceConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceConfigurationResponseBody) SetRequestId(v string) *UpdateServiceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceConfigurationResponseBody) SetCode(v string) *UpdateServiceConfigurationResponseBody {
	s.Code = &v
	return s
}

type UpdateServiceConfigurationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateServiceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigurationResponse) SetHeaders(v map[string]*string) *UpdateServiceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceConfigurationResponse) SetBody(v *UpdateServiceConfigurationResponseBody) *UpdateServiceConfigurationResponse {
	s.Body = v
	return s
}

type UpdateSlrConfigurationRequest struct {
	MqSubscribe  *bool     `json:"MqSubscribe,omitempty" xml:"MqSubscribe,omitempty"`
	MqEndpoint   *string   `json:"MqEndpoint,omitempty" xml:"MqEndpoint,omitempty"`
	MqInstanceId *string   `json:"MqInstanceId,omitempty" xml:"MqInstanceId,omitempty"`
	MqTopic      *string   `json:"MqTopic,omitempty" xml:"MqTopic,omitempty"`
	MqGroupId    *string   `json:"MqGroupId,omitempty" xml:"MqGroupId,omitempty"`
	MqEvent      []*string `json:"MqEvent,omitempty" xml:"MqEvent,omitempty" type:"Repeated"`
}

func (s UpdateSlrConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlrConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateSlrConfigurationRequest) SetMqSubscribe(v bool) *UpdateSlrConfigurationRequest {
	s.MqSubscribe = &v
	return s
}

func (s *UpdateSlrConfigurationRequest) SetMqEndpoint(v string) *UpdateSlrConfigurationRequest {
	s.MqEndpoint = &v
	return s
}

func (s *UpdateSlrConfigurationRequest) SetMqInstanceId(v string) *UpdateSlrConfigurationRequest {
	s.MqInstanceId = &v
	return s
}

func (s *UpdateSlrConfigurationRequest) SetMqTopic(v string) *UpdateSlrConfigurationRequest {
	s.MqTopic = &v
	return s
}

func (s *UpdateSlrConfigurationRequest) SetMqGroupId(v string) *UpdateSlrConfigurationRequest {
	s.MqGroupId = &v
	return s
}

func (s *UpdateSlrConfigurationRequest) SetMqEvent(v []*string) *UpdateSlrConfigurationRequest {
	s.MqEvent = v
	return s
}

type UpdateSlrConfigurationResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateSlrConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlrConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSlrConfigurationResponseBody) SetMessage(v string) *UpdateSlrConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSlrConfigurationResponseBody) SetRequestId(v string) *UpdateSlrConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSlrConfigurationResponseBody) SetCode(v string) *UpdateSlrConfigurationResponseBody {
	s.Code = &v
	return s
}

type UpdateSlrConfigurationResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateSlrConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSlrConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSlrConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateSlrConfigurationResponse) SetHeaders(v map[string]*string) *UpdateSlrConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateSlrConfigurationResponse) SetBody(v *UpdateSlrConfigurationResponseBody) *UpdateSlrConfigurationResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetName(v string) *UpdateUserRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumber(v string) *UpdateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

func (s *UpdateUserRequest) SetId(v string) *UpdateUserRequest {
	s.Id = &v
	return s
}

type UpdateUserResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetData(v map[string]interface{}) *UpdateUserResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

type UpdateUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("idrsservice.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("idrsservice.aliyuncs.com"),
		"ap-south-1":                  tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-1":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-2":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-3":              tea.String("idrsservice.aliyuncs.com"),
		"ap-southeast-5":              tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("idrsservice.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-chengdu":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-edge-1":                   tea.String("idrsservice.aliyuncs.com"),
		"cn-fujian":                   tea.String("idrsservice.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("idrsservice.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("idrsservice.aliyuncs.com"),
		"cn-hongkong":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("idrsservice.aliyuncs.com"),
		"cn-huhehaote":                tea.String("idrsservice.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("idrsservice.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("idrsservice.aliyuncs.com"),
		"cn-qingdao":                  tea.String("idrsservice.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("idrsservice.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-wuhan":                    tea.String("idrsservice.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("idrsservice.aliyuncs.com"),
		"cn-yushanfang":               tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("idrsservice.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("idrsservice.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("idrsservice.aliyuncs.com"),
		"eu-central-1":                tea.String("idrsservice.aliyuncs.com"),
		"eu-west-1":                   tea.String("idrsservice.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("idrsservice.aliyuncs.com"),
		"me-east-1":                   tea.String("idrsservice.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("idrsservice.aliyuncs.com"),
		"us-east-1":                   tea.String("idrsservice.aliyuncs.com"),
		"us-west-1":                   tea.String("idrsservice.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("idrsservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckServiceLinkedRole"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateApp"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDepartmentWithOptions(request *CreateDepartmentRequest, runtime *util.RuntimeOptions) (_result *CreateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDepartment"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDepartment(request *CreateDepartmentRequest) (_result *CreateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDepartmentResponse{}
	_body, _err := client.CreateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDetectProcessWithOptions(request *CreateDetectProcessRequest, runtime *util.RuntimeOptions) (_result *CreateDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDetectProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDetectProcess"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDetectProcess(request *CreateDetectProcessRequest) (_result *CreateDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDetectProcessResponse{}
	_body, _err := client.CreateDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveWithOptions(request *CreateLiveRequest, runtime *util.RuntimeOptions) (_result *CreateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLive"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLive(request *CreateLiveRequest) (_result *CreateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveResponse{}
	_body, _err := client.CreateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLiveDetectionWithOptions(request *CreateLiveDetectionRequest, runtime *util.RuntimeOptions) (_result *CreateLiveDetectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLiveDetectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLiveDetection"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLiveDetection(request *CreateLiveDetectionRequest) (_result *CreateLiveDetectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLiveDetectionResponse{}
	_body, _err := client.CreateLiveDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRuleWithOptions(request *CreateRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRule"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRule(request *CreateRuleRequest) (_result *CreateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRuleResponse{}
	_body, _err := client.CreateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStatisticsRecordWithOptions(request *CreateStatisticsRecordRequest, runtime *util.RuntimeOptions) (_result *CreateStatisticsRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStatisticsRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStatisticsRecord"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStatisticsRecord(request *CreateStatisticsRecordRequest) (_result *CreateStatisticsRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStatisticsRecordResponse{}
	_body, _err := client.CreateStatisticsRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStatisticsTaskWithOptions(request *CreateStatisticsTaskRequest, runtime *util.RuntimeOptions) (_result *CreateStatisticsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateStatisticsTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateStatisticsTask"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStatisticsTask(request *CreateStatisticsTaskRequest) (_result *CreateStatisticsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStatisticsTaskResponse{}
	_body, _err := client.CreateStatisticsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTaskGroupWithOptions(request *CreateTaskGroupRequest, runtime *util.RuntimeOptions) (_result *CreateTaskGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTaskGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTaskGroup"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTaskGroup(request *CreateTaskGroupRequest) (_result *CreateTaskGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTaskGroupResponse{}
	_body, _err := client.CreateTaskGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserDepartmentsWithOptions(request *CreateUserDepartmentsRequest, runtime *util.RuntimeOptions) (_result *CreateUserDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserDepartmentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserDepartments"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserDepartments(request *CreateUserDepartmentsRequest) (_result *CreateUserDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserDepartmentsResponse{}
	_body, _err := client.CreateUserDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteApp"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDepartmentWithOptions(request *DeleteDepartmentRequest, runtime *util.RuntimeOptions) (_result *DeleteDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDepartment"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDepartment(request *DeleteDepartmentRequest) (_result *DeleteDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDepartmentResponse{}
	_body, _err := client.DeleteDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDetectProcessWithOptions(request *DeleteDetectProcessRequest, runtime *util.RuntimeOptions) (_result *DeleteDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDetectProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDetectProcess"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDetectProcess(request *DeleteDetectProcessRequest) (_result *DeleteDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDetectProcessResponse{}
	_body, _err := client.DeleteDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRuleWithOptions(request *DeleteRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRule"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRule(request *DeleteRuleRequest) (_result *DeleteRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRuleResponse{}
	_body, _err := client.DeleteRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUser"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserDepartmentsWithOptions(request *DeleteUserDepartmentsRequest, runtime *util.RuntimeOptions) (_result *DeleteUserDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserDepartmentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserDepartments"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserDepartments(request *DeleteUserDepartmentsRequest) (_result *DeleteUserDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserDepartmentsResponse{}
	_body, _err := client.DeleteUserDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExitLiveWithOptions(request *ExitLiveRequest, runtime *util.RuntimeOptions) (_result *ExitLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExitLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExitLive"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExitLive(request *ExitLiveRequest) (_result *ExitLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExitLiveResponse{}
	_body, _err := client.ExitLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetApp"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBatchSignedUrlWithOptions(request *GetBatchSignedUrlRequest, runtime *util.RuntimeOptions) (_result *GetBatchSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetBatchSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetBatchSignedUrl"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBatchSignedUrl(request *GetBatchSignedUrlRequest) (_result *GetBatchSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetBatchSignedUrlResponse{}
	_body, _err := client.GetBatchSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDepartmentWithOptions(request *GetDepartmentRequest, runtime *util.RuntimeOptions) (_result *GetDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDepartment"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDepartment(request *GetDepartmentRequest) (_result *GetDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDepartmentResponse{}
	_body, _err := client.GetDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectEvaluationWithOptions(request *GetDetectEvaluationRequest, runtime *util.RuntimeOptions) (_result *GetDetectEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDetectEvaluationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectEvaluation"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectEvaluation(request *GetDetectEvaluationRequest) (_result *GetDetectEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectEvaluationResponse{}
	_body, _err := client.GetDetectEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectionWithOptions(request *GetDetectionRequest, runtime *util.RuntimeOptions) (_result *GetDetectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDetectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetection"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetection(request *GetDetectionRequest) (_result *GetDetectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectionResponse{}
	_body, _err := client.GetDetectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectProcessWithOptions(request *GetDetectProcessRequest, runtime *util.RuntimeOptions) (_result *GetDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDetectProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectProcess"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectProcess(request *GetDetectProcessRequest) (_result *GetDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectProcessResponse{}
	_body, _err := client.GetDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectProcessJsonFileWithOptions(request *GetDetectProcessJsonFileRequest, runtime *util.RuntimeOptions) (_result *GetDetectProcessJsonFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDetectProcessJsonFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectProcessJsonFile"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectProcessJsonFile(request *GetDetectProcessJsonFileRequest) (_result *GetDetectProcessJsonFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectProcessJsonFileResponse{}
	_body, _err := client.GetDetectProcessJsonFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDetectProcessTemplateWithOptions(runtime *util.RuntimeOptions) (_result *GetDetectProcessTemplateResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &GetDetectProcessTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDetectProcessTemplate"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDetectProcessTemplate() (_result *GetDetectProcessTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDetectProcessTemplateResponse{}
	_body, _err := client.GetDetectProcessTemplateWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGlobalConfigWithOptions(request *GetGlobalConfigRequest, runtime *util.RuntimeOptions) (_result *GetGlobalConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetGlobalConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGlobalConfig"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGlobalConfig(request *GetGlobalConfigRequest) (_result *GetGlobalConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGlobalConfigResponse{}
	_body, _err := client.GetGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelSignedUrlWithOptions(request *GetModelSignedUrlRequest, runtime *util.RuntimeOptions) (_result *GetModelSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetModelSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetModelSignedUrl"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelSignedUrl(request *GetModelSignedUrlRequest) (_result *GetModelSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetModelSignedUrlResponse{}
	_body, _err := client.GetModelSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPreSignedUrlWithOptions(request *GetPreSignedUrlRequest, runtime *util.RuntimeOptions) (_result *GetPreSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPreSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPreSignedUrl"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPreSignedUrl(request *GetPreSignedUrlRequest) (_result *GetPreSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPreSignedUrlResponse{}
	_body, _err := client.GetPreSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRuleWithOptions(request *GetRuleRequest, runtime *util.RuntimeOptions) (_result *GetRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRule"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRule(request *GetRuleRequest) (_result *GetRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRuleResponse{}
	_body, _err := client.GetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceConfigurationWithOptions(request *GetServiceConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetServiceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetServiceConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetServiceConfiguration"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceConfiguration(request *GetServiceConfigurationRequest) (_result *GetServiceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceConfigurationResponse{}
	_body, _err := client.GetServiceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSignedUrlWithOptions(request *GetSignedUrlRequest, runtime *util.RuntimeOptions) (_result *GetSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSignedUrl"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSignedUrl(request *GetSignedUrlRequest) (_result *GetSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSignedUrlResponse{}
	_body, _err := client.GetSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSlrConfigurationWithOptions(request *GetSlrConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetSlrConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSlrConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSlrConfiguration"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSlrConfiguration(request *GetSlrConfigurationRequest) (_result *GetSlrConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSlrConfigurationResponse{}
	_body, _err := client.GetSlrConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetStatisticsWithOptions(request *GetStatisticsRequest, runtime *util.RuntimeOptions) (_result *GetStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetStatisticsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetStatistics"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetStatistics(request *GetStatisticsRequest) (_result *GetStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStatisticsResponse{}
	_body, _err := client.GetStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTask"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskGroupWithOptions(request *GetTaskGroupRequest, runtime *util.RuntimeOptions) (_result *GetTaskGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTaskGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTaskGroup"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTaskGroup(request *GetTaskGroupRequest) (_result *GetTaskGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskGroupResponse{}
	_body, _err := client.GetTaskGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUser"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeServiceLinkedRoleWithOptions(request *InitializeServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *InitializeServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitializeServiceLinkedRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitializeServiceLinkedRole"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeServiceLinkedRole(request *InitializeServiceLinkedRoleRequest) (_result *InitializeServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeServiceLinkedRoleResponse{}
	_body, _err := client.InitializeServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinLiveWithOptions(request *JoinLiveRequest, runtime *util.RuntimeOptions) (_result *JoinLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &JoinLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("JoinLive"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinLive(request *JoinLiveRequest) (_result *JoinLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinLiveResponse{}
	_body, _err := client.JoinLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListApps"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDepartmentsWithOptions(request *ListDepartmentsRequest, runtime *util.RuntimeOptions) (_result *ListDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDepartmentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDepartments"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDepartments(request *ListDepartmentsRequest) (_result *ListDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDepartmentsResponse{}
	_body, _err := client.ListDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectionsWithOptions(request *ListDetectionsRequest, runtime *util.RuntimeOptions) (_result *ListDetectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDetectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDetections"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetections(request *ListDetectionsRequest) (_result *ListDetectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectionsResponse{}
	_body, _err := client.ListDetectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDetectProcessesWithOptions(request *ListDetectProcessesRequest, runtime *util.RuntimeOptions) (_result *ListDetectProcessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDetectProcessesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDetectProcesses"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDetectProcesses(request *ListDetectProcessesRequest) (_result *ListDetectProcessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDetectProcessesResponse{}
	_body, _err := client.ListDetectProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *util.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFiles"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLivesWithOptions(request *ListLivesRequest, runtime *util.RuntimeOptions) (_result *ListLivesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLivesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLives"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLives(request *ListLivesRequest) (_result *ListLivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLivesResponse{}
	_body, _err := client.ListLivesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListRolesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRoles"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles() (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRulesWithOptions(request *ListRulesRequest, runtime *util.RuntimeOptions) (_result *ListRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRules"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRules(request *ListRulesRequest) (_result *ListRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRulesResponse{}
	_body, _err := client.ListRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatisticsTaskWithOptions(request *ListStatisticsTaskRequest, runtime *util.RuntimeOptions) (_result *ListStatisticsTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListStatisticsTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListStatisticsTask"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatisticsTask(request *ListStatisticsTaskRequest) (_result *ListStatisticsTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStatisticsTaskResponse{}
	_body, _err := client.ListStatisticsTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskGroupsWithOptions(request *ListTaskGroupsRequest, runtime *util.RuntimeOptions) (_result *ListTaskGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTaskGroups"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskGroups(request *ListTaskGroupsRequest) (_result *ListTaskGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskGroupsResponse{}
	_body, _err := client.ListTaskGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTaskItemsWithOptions(request *ListTaskItemsRequest, runtime *util.RuntimeOptions) (_result *ListTaskItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTaskItemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTaskItems"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTaskItems(request *ListTaskItemsRequest) (_result *ListTaskItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskItemsResponse{}
	_body, _err := client.ListTaskItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTasks"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameDetectProcessWithOptions(request *RenameDetectProcessRequest, runtime *util.RuntimeOptions) (_result *RenameDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameDetectProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameDetectProcess"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameDetectProcess(request *RenameDetectProcessRequest) (_result *RenameDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameDetectProcessResponse{}
	_body, _err := client.RenameDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAppWithOptions(request *UpdateAppRequest, runtime *util.RuntimeOptions) (_result *UpdateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateApp"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApp(request *UpdateAppRequest) (_result *UpdateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAppResponse{}
	_body, _err := client.UpdateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDepartmentWithOptions(request *UpdateDepartmentRequest, runtime *util.RuntimeOptions) (_result *UpdateDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDepartment"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDepartment(request *UpdateDepartmentRequest) (_result *UpdateDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDepartmentResponse{}
	_body, _err := client.UpdateDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDetectProcessWithOptions(request *UpdateDetectProcessRequest, runtime *util.RuntimeOptions) (_result *UpdateDetectProcessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDetectProcessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDetectProcess"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDetectProcess(request *UpdateDetectProcessRequest) (_result *UpdateDetectProcessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDetectProcessResponse{}
	_body, _err := client.UpdateDetectProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLiveWithOptions(request *UpdateLiveRequest, runtime *util.RuntimeOptions) (_result *UpdateLiveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLiveResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLive"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLive(request *UpdateLiveRequest) (_result *UpdateLiveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLiveResponse{}
	_body, _err := client.UpdateLiveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRuleWithOptions(request *UpdateRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRule"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRule(request *UpdateRuleRequest) (_result *UpdateRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRuleResponse{}
	_body, _err := client.UpdateRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceConfigurationWithOptions(request *UpdateServiceConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateServiceConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateServiceConfiguration"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceConfiguration(request *UpdateServiceConfigurationRequest) (_result *UpdateServiceConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceConfigurationResponse{}
	_body, _err := client.UpdateServiceConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSlrConfigurationWithOptions(request *UpdateSlrConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateSlrConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateSlrConfigurationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSlrConfiguration"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSlrConfiguration(request *UpdateSlrConfigurationRequest) (_result *UpdateSlrConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSlrConfigurationResponse{}
	_body, _err := client.UpdateSlrConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUser"), tea.String("2020-06-30"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
