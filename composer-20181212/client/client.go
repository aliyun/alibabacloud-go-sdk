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

type CloneFlowRequest struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CloneFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowRequest) SetFlowId(v string) *CloneFlowRequest {
	s.FlowId = &v
	return s
}

func (s *CloneFlowRequest) SetVersionId(v string) *CloneFlowRequest {
	s.VersionId = &v
	return s
}

type CloneFlowResponseBody struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFlowResponseBody) SetFlowId(v string) *CloneFlowResponseBody {
	s.FlowId = &v
	return s
}

func (s *CloneFlowResponseBody) SetRequestId(v string) *CloneFlowResponseBody {
	s.RequestId = &v
	return s
}

type CloneFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowResponse) SetHeaders(v map[string]*string) *CloneFlowResponse {
	s.Headers = v
	return s
}

func (s *CloneFlowResponse) SetBody(v *CloneFlowResponseBody) *CloneFlowResponse {
	s.Body = v
	return s
}

type CreateFlowRequest struct {
	Definition      *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowSource      *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetDefinition(v string) *CreateFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateFlowRequest) SetFlowDescription(v string) *CreateFlowRequest {
	s.FlowDescription = &v
	return s
}

func (s *CreateFlowRequest) SetFlowName(v string) *CreateFlowRequest {
	s.FlowName = &v
	return s
}

func (s *CreateFlowRequest) SetFlowSource(v string) *CreateFlowRequest {
	s.FlowSource = &v
	return s
}

func (s *CreateFlowRequest) SetTemplateId(v string) *CreateFlowRequest {
	s.TemplateId = &v
	return s
}

type CreateFlowResponseBody struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetFlowId(v string) *CreateFlowResponseBody {
	s.FlowId = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) SetHeaders(v map[string]*string) *CreateFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

type DeleteFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetFlowId(v string) *DeleteFlowRequest {
	s.FlowId = &v
	return s
}

type DeleteFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowResponseBody) SetSuccess(v bool) *DeleteFlowResponseBody {
	s.Success = &v
	return s
}

type DeleteFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) SetHeaders(v map[string]*string) *DeleteFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

type DisableFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DisableFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowRequest) GoString() string {
	return s.String()
}

func (s *DisableFlowRequest) SetFlowId(v string) *DisableFlowRequest {
	s.FlowId = &v
	return s
}

type DisableFlowResponseBody struct {
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFlowResponseBody) SetFlowStatus(v string) *DisableFlowResponseBody {
	s.FlowStatus = &v
	return s
}

func (s *DisableFlowResponseBody) SetRequestId(v string) *DisableFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableFlowResponseBody) SetSuccess(v bool) *DisableFlowResponseBody {
	s.Success = &v
	return s
}

type DisableFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowResponse) GoString() string {
	return s.String()
}

func (s *DisableFlowResponse) SetHeaders(v map[string]*string) *DisableFlowResponse {
	s.Headers = v
	return s
}

func (s *DisableFlowResponse) SetBody(v *DisableFlowResponseBody) *DisableFlowResponse {
	s.Body = v
	return s
}

type EnableFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s EnableFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowRequest) GoString() string {
	return s.String()
}

func (s *EnableFlowRequest) SetFlowId(v string) *EnableFlowRequest {
	s.FlowId = &v
	return s
}

type EnableFlowResponseBody struct {
	FlowStatus *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowResponseBody) GoString() string {
	return s.String()
}

func (s *EnableFlowResponseBody) SetFlowStatus(v string) *EnableFlowResponseBody {
	s.FlowStatus = &v
	return s
}

func (s *EnableFlowResponseBody) SetRequestId(v string) *EnableFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableFlowResponseBody) SetSuccess(v bool) *EnableFlowResponseBody {
	s.Success = &v
	return s
}

type EnableFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowResponse) GoString() string {
	return s.String()
}

func (s *EnableFlowResponse) SetHeaders(v map[string]*string) *EnableFlowResponse {
	s.Headers = v
	return s
}

func (s *EnableFlowResponse) SetBody(v *EnableFlowResponseBody) *EnableFlowResponse {
	s.Body = v
	return s
}

type GetFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s GetFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFlowRequest) GoString() string {
	return s.String()
}

func (s *GetFlowRequest) SetFlowId(v string) *GetFlowRequest {
	s.FlowId = &v
	return s
}

type GetFlowResponseBody struct {
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentVersionId *int32  `json:"CurrentVersionId,omitempty" xml:"CurrentVersionId,omitempty"`
	Definition       *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FlowDescription  *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	FlowEditMode     *string `json:"FlowEditMode,omitempty" xml:"FlowEditMode,omitempty"`
	FlowId           *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName         *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowSource       *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty"`
	FlowStatus       *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId       *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UpdateTime       *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetFlowResponseBody) SetCreateTime(v string) *GetFlowResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetFlowResponseBody) SetCurrentVersionId(v int32) *GetFlowResponseBody {
	s.CurrentVersionId = &v
	return s
}

func (s *GetFlowResponseBody) SetDefinition(v string) *GetFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowDescription(v string) *GetFlowResponseBody {
	s.FlowDescription = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowEditMode(v string) *GetFlowResponseBody {
	s.FlowEditMode = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowId(v string) *GetFlowResponseBody {
	s.FlowId = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowName(v string) *GetFlowResponseBody {
	s.FlowName = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowSource(v string) *GetFlowResponseBody {
	s.FlowSource = &v
	return s
}

func (s *GetFlowResponseBody) SetFlowStatus(v string) *GetFlowResponseBody {
	s.FlowStatus = &v
	return s
}

func (s *GetFlowResponseBody) SetRegionId(v string) *GetFlowResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetFlowResponseBody) SetRequestId(v string) *GetFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFlowResponseBody) SetTemplateId(v string) *GetFlowResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetFlowResponseBody) SetUpdateTime(v string) *GetFlowResponseBody {
	s.UpdateTime = &v
	return s
}

type GetFlowResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFlowResponse) GoString() string {
	return s.String()
}

func (s *GetFlowResponse) SetHeaders(v map[string]*string) *GetFlowResponse {
	s.Headers = v
	return s
}

func (s *GetFlowResponse) SetBody(v *GetFlowResponseBody) *GetFlowResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetTemplateResponseBody struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Definition          *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateConnector   *string `json:"TemplateConnector,omitempty" xml:"TemplateConnector,omitempty"`
	TemplateCreator     *string `json:"TemplateCreator,omitempty" xml:"TemplateCreator,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateLocale      *string `json:"TemplateLocale,omitempty" xml:"TemplateLocale,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateOverview    *string `json:"TemplateOverview,omitempty" xml:"TemplateOverview,omitempty"`
	TemplateSummary     *string `json:"TemplateSummary,omitempty" xml:"TemplateSummary,omitempty"`
	TemplateSummaryEn   *string `json:"TemplateSummaryEn,omitempty" xml:"TemplateSummaryEn,omitempty"`
	TemplateTag         *string `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty"`
	TemplateVersion     *int32  `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetDefinition(v string) *GetTemplateResponseBody {
	s.Definition = &v
	return s
}

func (s *GetTemplateResponseBody) SetRegionId(v string) *GetTemplateResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateConnector(v string) *GetTemplateResponseBody {
	s.TemplateConnector = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateCreator(v string) *GetTemplateResponseBody {
	s.TemplateCreator = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateDescription(v string) *GetTemplateResponseBody {
	s.TemplateDescription = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateLocale(v string) *GetTemplateResponseBody {
	s.TemplateLocale = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateOverview(v string) *GetTemplateResponseBody {
	s.TemplateOverview = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateSummary(v string) *GetTemplateResponseBody {
	s.TemplateSummary = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateSummaryEn(v string) *GetTemplateResponseBody {
	s.TemplateSummaryEn = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateTag(v string) *GetTemplateResponseBody {
	s.TemplateTag = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateVersion(v int32) *GetTemplateResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBody) SetUpdateTime(v string) *GetTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

type GetTemplateResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetVersionRequest struct {
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVersionRequest) GoString() string {
	return s.String()
}

func (s *GetVersionRequest) SetFlowId(v string) *GetVersionRequest {
	s.FlowId = &v
	return s
}

func (s *GetVersionRequest) SetVersionId(v string) *GetVersionRequest {
	s.VersionId = &v
	return s
}

type GetVersionResponseBody struct {
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Definition         *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FlowId             *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpdateTime         *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	VersionId          *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionName        *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VersionStatus      *string `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty"`
}

func (s GetVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionResponseBody) SetCreateTime(v string) *GetVersionResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVersionResponseBody) SetDefinition(v string) *GetVersionResponseBody {
	s.Definition = &v
	return s
}

func (s *GetVersionResponseBody) SetFlowId(v string) *GetVersionResponseBody {
	s.FlowId = &v
	return s
}

func (s *GetVersionResponseBody) SetRegionId(v string) *GetVersionResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVersionResponseBody) SetRequestId(v string) *GetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVersionResponseBody) SetUpdateTime(v string) *GetVersionResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetVersionResponseBody) SetVersionDescription(v string) *GetVersionResponseBody {
	s.VersionDescription = &v
	return s
}

func (s *GetVersionResponseBody) SetVersionId(v string) *GetVersionResponseBody {
	s.VersionId = &v
	return s
}

func (s *GetVersionResponseBody) SetVersionName(v string) *GetVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetVersionResponseBody) SetVersionStatus(v string) *GetVersionResponseBody {
	s.VersionStatus = &v
	return s
}

type GetVersionResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVersionResponse) GoString() string {
	return s.String()
}

func (s *GetVersionResponse) SetHeaders(v map[string]*string) *GetVersionResponse {
	s.Headers = v
	return s
}

func (s *GetVersionResponse) SetBody(v *GetVersionResponseBody) *GetVersionResponse {
	s.Body = v
	return s
}

type GroupInvokeFlowRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	GroupKey    *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GroupInvokeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupInvokeFlowRequest) GoString() string {
	return s.String()
}

func (s *GroupInvokeFlowRequest) SetClientToken(v string) *GroupInvokeFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetData(v string) *GroupInvokeFlowRequest {
	s.Data = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetFlowId(v string) *GroupInvokeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetGroupKey(v string) *GroupInvokeFlowRequest {
	s.GroupKey = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetTags(v string) *GroupInvokeFlowRequest {
	s.Tags = &v
	return s
}

func (s *GroupInvokeFlowRequest) SetTotalCount(v int32) *GroupInvokeFlowRequest {
	s.TotalCount = &v
	return s
}

type GroupInvokeFlowResponseBody struct {
	CurrentCount      *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	GroupInvocationId *string `json:"GroupInvocationId,omitempty" xml:"GroupInvocationId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success           *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GroupInvokeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GroupInvokeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GroupInvokeFlowResponseBody) SetCurrentCount(v int32) *GroupInvokeFlowResponseBody {
	s.CurrentCount = &v
	return s
}

func (s *GroupInvokeFlowResponseBody) SetGroupInvocationId(v string) *GroupInvokeFlowResponseBody {
	s.GroupInvocationId = &v
	return s
}

func (s *GroupInvokeFlowResponseBody) SetRequestId(v string) *GroupInvokeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GroupInvokeFlowResponseBody) SetStatus(v string) *GroupInvokeFlowResponseBody {
	s.Status = &v
	return s
}

func (s *GroupInvokeFlowResponseBody) SetSuccess(v bool) *GroupInvokeFlowResponseBody {
	s.Success = &v
	return s
}

type GroupInvokeFlowResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GroupInvokeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GroupInvokeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupInvokeFlowResponse) GoString() string {
	return s.String()
}

func (s *GroupInvokeFlowResponse) SetHeaders(v map[string]*string) *GroupInvokeFlowResponse {
	s.Headers = v
	return s
}

func (s *GroupInvokeFlowResponse) SetBody(v *GroupInvokeFlowResponseBody) *GroupInvokeFlowResponse {
	s.Body = v
	return s
}

type InvokeFlowRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Parameters  *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s InvokeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowRequest) GoString() string {
	return s.String()
}

func (s *InvokeFlowRequest) SetClientToken(v string) *InvokeFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *InvokeFlowRequest) SetData(v string) *InvokeFlowRequest {
	s.Data = &v
	return s
}

func (s *InvokeFlowRequest) SetFlowId(v string) *InvokeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *InvokeFlowRequest) SetParameters(v string) *InvokeFlowRequest {
	s.Parameters = &v
	return s
}

type InvokeFlowResponseBody struct {
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvokeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeFlowResponseBody) SetInvocationId(v string) *InvokeFlowResponseBody {
	s.InvocationId = &v
	return s
}

func (s *InvokeFlowResponseBody) SetRequestId(v string) *InvokeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeFlowResponseBody) SetSuccess(v bool) *InvokeFlowResponseBody {
	s.Success = &v
	return s
}

type InvokeFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InvokeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowResponse) GoString() string {
	return s.String()
}

func (s *InvokeFlowResponse) SetHeaders(v map[string]*string) *InvokeFlowResponse {
	s.Headers = v
	return s
}

func (s *InvokeFlowResponse) SetBody(v *InvokeFlowResponseBody) *InvokeFlowResponse {
	s.Body = v
	return s
}

type ListFlowsRequest struct {
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	FlowName   *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetFilter(v string) *ListFlowsRequest {
	s.Filter = &v
	return s
}

func (s *ListFlowsRequest) SetFlowName(v string) *ListFlowsRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowsRequest) SetPageNumber(v int32) *ListFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsRequest) SetPageSize(v int32) *ListFlowsRequest {
	s.PageSize = &v
	return s
}

type ListFlowsResponseBody struct {
	Flows      []*ListFlowsResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody {
	s.Flows = v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) SetTotalCount(v int32) *ListFlowsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFlowsResponseBodyFlows struct {
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	FlowEditMode    *string `json:"FlowEditMode,omitempty" xml:"FlowEditMode,omitempty"`
	FlowId          *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowSource      *string `json:"FlowSource,omitempty" xml:"FlowSource,omitempty"`
	FlowStatus      *string `json:"FlowStatus,omitempty" xml:"FlowStatus,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UpdateTime      *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionId       *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListFlowsResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlows) SetCreateTime(v string) *ListFlowsResponseBodyFlows {
	s.CreateTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowDescription(v string) *ListFlowsResponseBodyFlows {
	s.FlowDescription = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowEditMode(v string) *ListFlowsResponseBodyFlows {
	s.FlowEditMode = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowId(v string) *ListFlowsResponseBodyFlows {
	s.FlowId = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowName(v string) *ListFlowsResponseBodyFlows {
	s.FlowName = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowSource(v string) *ListFlowsResponseBodyFlows {
	s.FlowSource = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetFlowStatus(v string) *ListFlowsResponseBodyFlows {
	s.FlowStatus = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetRegionId(v string) *ListFlowsResponseBodyFlows {
	s.RegionId = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetTemplateId(v string) *ListFlowsResponseBodyFlows {
	s.TemplateId = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetUpdateTime(v string) *ListFlowsResponseBodyFlows {
	s.UpdateTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetVersionId(v int32) *ListFlowsResponseBodyFlows {
	s.VersionId = &v
	return s
}

type ListFlowsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowsResponse) SetHeaders(v map[string]*string) *ListFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowsResponse) SetBody(v *ListFlowsResponseBody) *ListFlowsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	MaxResults   *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetMaxResults(v int32) *ListTagResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetLang(v string) *ListTemplatesRequest {
	s.Lang = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetTag(v string) *ListTemplatesRequest {
	s.Tag = &v
	return s
}

type ListTemplatesResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates  []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TemplateConnector   *string `json:"TemplateConnector,omitempty" xml:"TemplateConnector,omitempty"`
	TemplateCreator     *string `json:"TemplateCreator,omitempty" xml:"TemplateCreator,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	TemplateId          *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateLocale      *string `json:"TemplateLocale,omitempty" xml:"TemplateLocale,omitempty"`
	TemplateName        *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateOverview    *string `json:"TemplateOverview,omitempty" xml:"TemplateOverview,omitempty"`
	TemplateSummary     *string `json:"TemplateSummary,omitempty" xml:"TemplateSummary,omitempty"`
	TemplateSummaryEn   *string `json:"TemplateSummaryEn,omitempty" xml:"TemplateSummaryEn,omitempty"`
	TemplateTag         *string `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty"`
	TemplateVersion     *int32  `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	UpdateTime          *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateConnector(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateConnector = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateCreator(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateCreator = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateDescription = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateLocale(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateLocale = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateOverview(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateOverview = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateSummary(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateSummary = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateSummaryEn(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateSummaryEn = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateTag(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateTag = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v int32) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdateTime = &v
	return s
}

type ListTemplatesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type ListVersionsRequest struct {
	FlowId     *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) SetFlowId(v string) *ListVersionsRequest {
	s.FlowId = &v
	return s
}

func (s *ListVersionsRequest) SetPageNumber(v int32) *ListVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVersionsRequest) SetPageSize(v int32) *ListVersionsRequest {
	s.PageSize = &v
	return s
}

type ListVersionsResponseBody struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Versions   []*ListVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBody) SetRequestId(v string) *ListVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponseBody) SetTotalCount(v int32) *ListVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVersionsResponseBody) SetVersions(v []*ListVersionsResponseBodyVersions) *ListVersionsResponseBody {
	s.Versions = v
	return s
}

type ListVersionsResponseBodyVersions struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FlowId        *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VersionId     *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionName   *int32  `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VersionStatus *int32  `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty"`
}

func (s ListVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBodyVersions) SetCreateTime(v string) *ListVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListVersionsResponseBodyVersions) SetFlowId(v string) *ListVersionsResponseBodyVersions {
	s.FlowId = &v
	return s
}

func (s *ListVersionsResponseBodyVersions) SetUpdateTime(v string) *ListVersionsResponseBodyVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListVersionsResponseBodyVersions) SetVersionId(v string) *ListVersionsResponseBodyVersions {
	s.VersionId = &v
	return s
}

func (s *ListVersionsResponseBodyVersions) SetVersionName(v int32) *ListVersionsResponseBodyVersions {
	s.VersionName = &v
	return s
}

func (s *ListVersionsResponseBodyVersions) SetVersionStatus(v int32) *ListVersionsResponseBodyVersions {
	s.VersionStatus = &v
	return s
}

type ListVersionsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionsResponse) SetHeaders(v map[string]*string) *ListVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListVersionsResponse) SetBody(v *ListVersionsResponseBody) *ListVersionsResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateFlowRequest struct {
	Definition      *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FlowDescription *string `json:"FlowDescription,omitempty" xml:"FlowDescription,omitempty"`
	FlowId          *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s UpdateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) SetDefinition(v string) *UpdateFlowRequest {
	s.Definition = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowDescription(v string) *UpdateFlowRequest {
	s.FlowDescription = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowId(v string) *UpdateFlowRequest {
	s.FlowId = &v
	return s
}

func (s *UpdateFlowRequest) SetFlowName(v string) *UpdateFlowRequest {
	s.FlowName = &v
	return s
}

type UpdateFlowResponseBody struct {
	CurrentVersionId *int32  `json:"CurrentVersionId,omitempty" xml:"CurrentVersionId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBody) SetCurrentVersionId(v int32) *UpdateFlowResponseBody {
	s.CurrentVersionId = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRequestId(v string) *UpdateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponseBody) SetSuccess(v bool) *UpdateFlowResponseBody {
	s.Success = &v
	return s
}

type UpdateFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowResponse) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponse) SetHeaders(v map[string]*string) *UpdateFlowResponse {
	s.Headers = v
	return s
}

func (s *UpdateFlowResponse) SetBody(v *UpdateFlowResponseBody) *UpdateFlowResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("composer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CloneFlowWithOptions(request *CloneFlowRequest, runtime *util.RuntimeOptions) (_result *CloneFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloneFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloneFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlow(request *CloneFlowRequest) (_result *CloneFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowResponse{}
	_body, _err := client.CloneFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableFlowWithOptions(request *DisableFlowRequest, runtime *util.RuntimeOptions) (_result *DisableFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableFlow(request *DisableFlowRequest) (_result *DisableFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableFlowResponse{}
	_body, _err := client.DisableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableFlowWithOptions(request *EnableFlowRequest, runtime *util.RuntimeOptions) (_result *EnableFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableFlow(request *EnableFlowRequest) (_result *EnableFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFlowResponse{}
	_body, _err := client.EnableFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFlowWithOptions(request *GetFlowRequest, runtime *util.RuntimeOptions) (_result *GetFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFlow(request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetTemplate"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVersionWithOptions(request *GetVersionRequest, runtime *util.RuntimeOptions) (_result *GetVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVersion"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVersion(request *GetVersionRequest) (_result *GetVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVersionResponse{}
	_body, _err := client.GetVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GroupInvokeFlowWithOptions(request *GroupInvokeFlowRequest, runtime *util.RuntimeOptions) (_result *GroupInvokeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GroupInvokeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GroupInvokeFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GroupInvokeFlow(request *GroupInvokeFlowRequest) (_result *GroupInvokeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GroupInvokeFlowResponse{}
	_body, _err := client.GroupInvokeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeFlowWithOptions(request *InvokeFlowRequest, runtime *util.RuntimeOptions) (_result *InvokeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InvokeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InvokeFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeFlow(request *InvokeFlowRequest) (_result *InvokeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeFlowResponse{}
	_body, _err := client.InvokeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *util.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlows"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTemplates"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVersionsWithOptions(request *ListVersionsRequest, runtime *util.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVersionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVersions"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVersions(request *ListVersionsRequest) (_result *ListVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVersionsResponse{}
	_body, _err := client.ListVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFlowWithOptions(request *UpdateFlowRequest, runtime *util.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFlow(request *UpdateFlowRequest) (_result *UpdateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFlowResponse{}
	_body, _err := client.UpdateFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
