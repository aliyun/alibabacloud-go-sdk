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

type AbolishFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s AbolishFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishFlowRequest) GoString() string {
	return s.String()
}

func (s *AbolishFlowRequest) SetFlowId(v string) *AbolishFlowRequest {
	s.FlowId = &v
	return s
}

type AbolishFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbolishFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishFlowResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishFlowResponseBody) SetRequestId(v string) *AbolishFlowResponseBody {
	s.RequestId = &v
	return s
}

type AbolishFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbolishFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbolishFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishFlowResponse) GoString() string {
	return s.String()
}

func (s *AbolishFlowResponse) SetHeaders(v map[string]*string) *AbolishFlowResponse {
	s.Headers = v
	return s
}

func (s *AbolishFlowResponse) SetBody(v *AbolishFlowResponseBody) *AbolishFlowResponse {
	s.Body = v
	return s
}

type CloneFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowVersion *int32  `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
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

func (s *CloneFlowRequest) SetFlowVersion(v int32) *CloneFlowRequest {
	s.FlowVersion = &v
	return s
}

type CloneFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponseBody) GoString() string {
	return s.String()
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
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetDefinition(v string) *CreateFlowRequest {
	s.Definition = &v
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
	FlowId  *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowIds *string `json:"FlowIds,omitempty" xml:"FlowIds,omitempty"`
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

func (s *DeleteFlowRequest) SetFlowIds(v string) *DeleteFlowRequest {
	s.FlowIds = &v
	return s
}

type DeleteFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeployFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DeployFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployFlowRequest) GoString() string {
	return s.String()
}

func (s *DeployFlowRequest) SetFlowId(v string) *DeployFlowRequest {
	s.FlowId = &v
	return s
}

type DeployFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeployFlowResponseBody) SetRequestId(v string) *DeployFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeployFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeployFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployFlowResponse) GoString() string {
	return s.String()
}

func (s *DeployFlowResponse) SetHeaders(v map[string]*string) *DeployFlowResponse {
	s.Headers = v
	return s
}

func (s *DeployFlowResponse) SetBody(v *DeployFlowResponseBody) *DeployFlowResponse {
	s.Body = v
	return s
}

type DescribeAccountSummaryResponseBody struct {
	InstanceCount             *int64  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	OnlineInstanceCount       *int64  `json:"OnlineInstanceCount,omitempty" xml:"OnlineInstanceCount,omitempty"`
	RequestId                 *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvocationCount           *int64  `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	DailyInvocationErrorCount *int64  `json:"DailyInvocationErrorCount,omitempty" xml:"DailyInvocationErrorCount,omitempty"`
}

func (s DescribeAccountSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountSummaryResponseBody) SetInstanceCount(v int64) *DescribeAccountSummaryResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeAccountSummaryResponseBody) SetOnlineInstanceCount(v int64) *DescribeAccountSummaryResponseBody {
	s.OnlineInstanceCount = &v
	return s
}

func (s *DescribeAccountSummaryResponseBody) SetRequestId(v string) *DescribeAccountSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountSummaryResponseBody) SetInvocationCount(v int64) *DescribeAccountSummaryResponseBody {
	s.InvocationCount = &v
	return s
}

func (s *DescribeAccountSummaryResponseBody) SetDailyInvocationErrorCount(v int64) *DescribeAccountSummaryResponseBody {
	s.DailyInvocationErrorCount = &v
	return s
}

type DescribeAccountSummaryResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountSummaryResponse) SetHeaders(v map[string]*string) *DescribeAccountSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountSummaryResponse) SetBody(v *DescribeAccountSummaryResponseBody) *DescribeAccountSummaryResponse {
	s.Body = v
	return s
}

type DescribeConnectorAttributeRequest struct {
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StepInfo      *string `json:"StepInfo,omitempty" xml:"StepInfo,omitempty"`
}

func (s DescribeConnectorAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeRequest) SetConnectorName(v string) *DescribeConnectorAttributeRequest {
	s.ConnectorName = &v
	return s
}

func (s *DescribeConnectorAttributeRequest) SetLang(v string) *DescribeConnectorAttributeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConnectorAttributeRequest) SetStepInfo(v string) *DescribeConnectorAttributeRequest {
	s.StepInfo = &v
	return s
}

type DescribeConnectorAttributeResponseBody struct {
	Description  *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Category     *string                                             `json:"Category,omitempty" xml:"Category,omitempty"`
	RequestId    *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Capabilities *DescribeConnectorAttributeResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
	FullName     *string                                             `json:"FullName,omitempty" xml:"FullName,omitempty"`
	DisplayName  *string                                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	RegionId     *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Icon         *DescribeConnectorAttributeResponseBodyIcon         `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	BrandColor   *string                                             `json:"BrandColor,omitempty" xml:"BrandColor,omitempty"`
	StepResult   *DescribeConnectorAttributeResponseBodyStepResult   `json:"StepResult,omitempty" xml:"StepResult,omitempty" type:"Struct"`
	Name         *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	ConnectorId  *string                                             `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DescribeConnectorAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBody) SetDescription(v string) *DescribeConnectorAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetCategory(v string) *DescribeConnectorAttributeResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetRequestId(v string) *DescribeConnectorAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetCapabilities(v *DescribeConnectorAttributeResponseBodyCapabilities) *DescribeConnectorAttributeResponseBody {
	s.Capabilities = v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetFullName(v string) *DescribeConnectorAttributeResponseBody {
	s.FullName = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetDisplayName(v string) *DescribeConnectorAttributeResponseBody {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetRegionId(v string) *DescribeConnectorAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetIcon(v *DescribeConnectorAttributeResponseBodyIcon) *DescribeConnectorAttributeResponseBody {
	s.Icon = v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetBrandColor(v string) *DescribeConnectorAttributeResponseBody {
	s.BrandColor = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetStepResult(v *DescribeConnectorAttributeResponseBodyStepResult) *DescribeConnectorAttributeResponseBody {
	s.StepResult = v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetName(v string) *DescribeConnectorAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBody) SetConnectorId(v string) *DescribeConnectorAttributeResponseBody {
	s.ConnectorId = &v
	return s
}

type DescribeConnectorAttributeResponseBodyCapabilities struct {
	Actions  []*DescribeConnectorAttributeResponseBodyCapabilitiesActions  `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	Triggers []*DescribeConnectorAttributeResponseBodyCapabilitiesTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s DescribeConnectorAttributeResponseBodyCapabilities) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBodyCapabilities) SetActions(v []*DescribeConnectorAttributeResponseBodyCapabilitiesActions) *DescribeConnectorAttributeResponseBodyCapabilities {
	s.Actions = v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilities) SetTriggers(v []*DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) *DescribeConnectorAttributeResponseBodyCapabilities {
	s.Triggers = v
	return s
}

type DescribeConnectorAttributeResponseBodyCapabilitiesActions struct {
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Visibility        *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentUrl       *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	System            *bool   `json:"System,omitempty" xml:"System,omitempty"`
	DefaultActionName *string `json:"DefaultActionName,omitempty" xml:"DefaultActionName,omitempty"`
}

func (s DescribeConnectorAttributeResponseBodyCapabilitiesActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBodyCapabilitiesActions) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetDisplayName(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetType(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.Type = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetVisibility(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.Visibility = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetDescription(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.Description = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetDocumentUrl(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.DocumentUrl = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetName(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.Name = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetSystem(v bool) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.System = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesActions) SetDefaultActionName(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesActions {
	s.DefaultActionName = &v
	return s
}

type DescribeConnectorAttributeResponseBodyCapabilitiesTriggers struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Visibility  *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	System      *bool   `json:"System,omitempty" xml:"System,omitempty"`
}

func (s DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetDisplayName(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetType(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.Type = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetVisibility(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.Visibility = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetDescription(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.Description = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetDocumentUrl(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.DocumentUrl = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetName(v string) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.Name = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers) SetSystem(v bool) *DescribeConnectorAttributeResponseBodyCapabilitiesTriggers {
	s.System = &v
	return s
}

type DescribeConnectorAttributeResponseBodyIcon struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeConnectorAttributeResponseBodyIcon) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBodyIcon) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBodyIcon) SetType(v string) *DescribeConnectorAttributeResponseBodyIcon {
	s.Type = &v
	return s
}

func (s *DescribeConnectorAttributeResponseBodyIcon) SetValue(v string) *DescribeConnectorAttributeResponseBodyIcon {
	s.Value = &v
	return s
}

type DescribeConnectorAttributeResponseBodyStepResult struct {
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
}

func (s DescribeConnectorAttributeResponseBodyStepResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponseBodyStepResult) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponseBodyStepResult) SetHasNext(v bool) *DescribeConnectorAttributeResponseBodyStepResult {
	s.HasNext = &v
	return s
}

type DescribeConnectorAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConnectorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectorAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectorAttributeResponse) SetHeaders(v map[string]*string) *DescribeConnectorAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectorAttributeResponse) SetBody(v *DescribeConnectorAttributeResponseBody) *DescribeConnectorAttributeResponse {
	s.Body = v
	return s
}

type DescribeConnectorCapabilityRequest struct {
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Lang   *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Preset *string `json:"Preset,omitempty" xml:"Preset,omitempty"`
}

func (s DescribeConnectorCapabilityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityRequest) SetType(v string) *DescribeConnectorCapabilityRequest {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityRequest) SetLang(v string) *DescribeConnectorCapabilityRequest {
	s.Lang = &v
	return s
}

func (s *DescribeConnectorCapabilityRequest) SetPreset(v string) *DescribeConnectorCapabilityRequest {
	s.Preset = &v
	return s
}

type DescribeConnectorCapabilityResponseBody struct {
	Connections   []*DescribeConnectorCapabilityResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	Connector     *DescribeConnectorCapabilityResponseBodyConnector     `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	Type          *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	Parameters    []*DescribeConnectorCapabilityResponseBodyParameters  `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId     *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DisplayName   *string                                               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DocumentUrl   *string                                               `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	Visibility    *string                                               `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	DefaultInputs *string                                               `json:"DefaultInputs,omitempty" xml:"DefaultInputs,omitempty"`
	System        *bool                                                 `json:"System,omitempty" xml:"System,omitempty"`
	Name          *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Responses     []*DescribeConnectorCapabilityResponseBodyResponses   `json:"Responses,omitempty" xml:"Responses,omitempty" type:"Repeated"`
}

func (s DescribeConnectorCapabilityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBody) SetConnections(v []*DescribeConnectorCapabilityResponseBodyConnections) *DescribeConnectorCapabilityResponseBody {
	s.Connections = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetConnector(v *DescribeConnectorCapabilityResponseBodyConnector) *DescribeConnectorCapabilityResponseBody {
	s.Connector = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetType(v string) *DescribeConnectorCapabilityResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetParameters(v []*DescribeConnectorCapabilityResponseBodyParameters) *DescribeConnectorCapabilityResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetRequestId(v string) *DescribeConnectorCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetDisplayName(v string) *DescribeConnectorCapabilityResponseBody {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetDocumentUrl(v string) *DescribeConnectorCapabilityResponseBody {
	s.DocumentUrl = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetVisibility(v string) *DescribeConnectorCapabilityResponseBody {
	s.Visibility = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetDefaultInputs(v string) *DescribeConnectorCapabilityResponseBody {
	s.DefaultInputs = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetSystem(v bool) *DescribeConnectorCapabilityResponseBody {
	s.System = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetName(v string) *DescribeConnectorCapabilityResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBody) SetResponses(v []*DescribeConnectorCapabilityResponseBodyResponses) *DescribeConnectorCapabilityResponseBody {
	s.Responses = v
	return s
}

type DescribeConnectorCapabilityResponseBodyConnections struct {
	Definition     *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyConnections) SetDefinition(v string) *DescribeConnectorCapabilityResponseBodyConnections {
	s.Definition = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnections) SetConnectionName(v string) *DescribeConnectorCapabilityResponseBodyConnections {
	s.ConnectionName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnections) SetContent(v string) *DescribeConnectorCapabilityResponseBodyConnections {
	s.Content = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyConnector struct {
	DisplayName          *string                                                               `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DefaultPolicy        []*string                                                             `json:"DefaultPolicy,omitempty" xml:"DefaultPolicy,omitempty" type:"Repeated"`
	FullName             *string                                                               `json:"FullName,omitempty" xml:"FullName,omitempty"`
	Icon                 *DescribeConnectorCapabilityResponseBodyConnectorIcon                 `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	ConnectorId          *string                                                               `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	RegionId             *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RamMap               *string                                                               `json:"RamMap,omitempty" xml:"RamMap,omitempty"`
	Description          *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	BrandColor           *string                                                               `json:"BrandColor,omitempty" xml:"BrandColor,omitempty"`
	Category             *string                                                               `json:"Category,omitempty" xml:"Category,omitempty"`
	ConnectionParameters *DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters `json:"ConnectionParameters,omitempty" xml:"ConnectionParameters,omitempty" type:"Struct"`
	Name                 *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyConnector) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyConnector) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetDisplayName(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetDefaultPolicy(v []*string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.DefaultPolicy = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetFullName(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.FullName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetIcon(v *DescribeConnectorCapabilityResponseBodyConnectorIcon) *DescribeConnectorCapabilityResponseBodyConnector {
	s.Icon = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetConnectorId(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.ConnectorId = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetRegionId(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.RegionId = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetRamMap(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.RamMap = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetDescription(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.Description = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetBrandColor(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.BrandColor = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetCategory(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.Category = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetConnectionParameters(v *DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters) *DescribeConnectorCapabilityResponseBodyConnector {
	s.ConnectionParameters = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnector) SetName(v string) *DescribeConnectorCapabilityResponseBodyConnector {
	s.Name = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyConnectorIcon struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyConnectorIcon) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyConnectorIcon) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyConnectorIcon) SetType(v string) *DescribeConnectorCapabilityResponseBodyConnectorIcon {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyConnectorIcon) SetValue(v string) *DescribeConnectorCapabilityResponseBodyConnectorIcon {
	s.Value = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters) SetType(v string) *DescribeConnectorCapabilityResponseBodyConnectorConnectionParameters {
	s.Type = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyParameters struct {
	DisplayName     *string                                                           `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Type            *string                                                           `json:"Type,omitempty" xml:"Type,omitempty"`
	PlaceHolder     *string                                                           `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	ReadOnly        *bool                                                             `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Example         *string                                                           `json:"Example,omitempty" xml:"Example,omitempty"`
	DefaultValue    *string                                                           `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	SubType         *string                                                           `json:"SubType,omitempty" xml:"SubType,omitempty"`
	EnumDisplayName []*string                                                         `json:"EnumDisplayName,omitempty" xml:"EnumDisplayName,omitempty" type:"Repeated"`
	Required        *bool                                                             `json:"Required,omitempty" xml:"Required,omitempty"`
	Description     *string                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Position        *string                                                           `json:"Position,omitempty" xml:"Position,omitempty"`
	Enum            []*string                                                         `json:"Enum,omitempty" xml:"Enum,omitempty" type:"Repeated"`
	SubParameters   []*DescribeConnectorCapabilityResponseBodyParametersSubParameters `json:"SubParameters,omitempty" xml:"SubParameters,omitempty" type:"Repeated"`
	Name            *string                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetDisplayName(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetType(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetPlaceHolder(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.PlaceHolder = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetReadOnly(v bool) *DescribeConnectorCapabilityResponseBodyParameters {
	s.ReadOnly = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetExample(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Example = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetDefaultValue(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.DefaultValue = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetSubType(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.SubType = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetEnumDisplayName(v []*string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.EnumDisplayName = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetRequired(v bool) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Required = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetDescription(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetPosition(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Position = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetEnum(v []*string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Enum = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetSubParameters(v []*DescribeConnectorCapabilityResponseBodyParametersSubParameters) *DescribeConnectorCapabilityResponseBodyParameters {
	s.SubParameters = v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParameters) SetName(v string) *DescribeConnectorCapabilityResponseBodyParameters {
	s.Name = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyParametersSubParameters struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Required    *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyParametersSubParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyParametersSubParameters) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyParametersSubParameters) SetDisplayName(v string) *DescribeConnectorCapabilityResponseBodyParametersSubParameters {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParametersSubParameters) SetType(v string) *DescribeConnectorCapabilityResponseBodyParametersSubParameters {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParametersSubParameters) SetRequired(v bool) *DescribeConnectorCapabilityResponseBodyParametersSubParameters {
	s.Required = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyParametersSubParameters) SetName(v string) *DescribeConnectorCapabilityResponseBodyParametersSubParameters {
	s.Name = &v
	return s
}

type DescribeConnectorCapabilityResponseBodyResponses struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Expression  *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	Example     *string `json:"Example,omitempty" xml:"Example,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeConnectorCapabilityResponseBodyResponses) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetType(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.Type = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetDisplayName(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.DisplayName = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetDescription(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.Description = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetExpression(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.Expression = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetExample(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.Example = &v
	return s
}

func (s *DescribeConnectorCapabilityResponseBodyResponses) SetName(v string) *DescribeConnectorCapabilityResponseBodyResponses {
	s.Name = &v
	return s
}

type DescribeConnectorCapabilityResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConnectorCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectorCapabilityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectorCapabilityResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectorCapabilityResponse) SetHeaders(v map[string]*string) *DescribeConnectorCapabilityResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectorCapabilityResponse) SetBody(v *DescribeConnectorCapabilityResponseBody) *DescribeConnectorCapabilityResponse {
	s.Body = v
	return s
}

type DescribeFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowVersion *int32  `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) SetFlowId(v string) *DescribeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowRequest) SetFlowVersion(v int32) *DescribeFlowRequest {
	s.FlowVersion = &v
	return s
}

type DescribeFlowResponseBody struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Version     *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	EditMode    *string `json:"EditMode,omitempty" xml:"EditMode,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) SetStatus(v string) *DescribeFlowResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowResponseBody) SetFlowId(v string) *DescribeFlowResponseBody {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetVersion(v int32) *DescribeFlowResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCreatedAt(v string) *DescribeFlowResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDefinition(v string) *DescribeFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEditMode(v string) *DescribeFlowResponseBody {
	s.EditMode = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRegionId(v string) *DescribeFlowResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetUpdatedAt(v string) *DescribeFlowResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeFlowResponseBody) SetSource(v string) *DescribeFlowResponseBody {
	s.Source = &v
	return s
}

func (s *DescribeFlowResponseBody) SetName(v string) *DescribeFlowResponseBody {
	s.Name = &v
	return s
}

type DescribeFlowResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponse) SetHeaders(v map[string]*string) *DescribeFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowResponse) SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse {
	s.Body = v
	return s
}

type DescribeFlowMetricRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DescribeFlowMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricRequest) SetFlowId(v string) *DescribeFlowMetricRequest {
	s.FlowId = &v
	return s
}

type DescribeFlowMetricResponseBody struct {
	RequestId            *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvocationCount      *int64   `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	InvocationErrorCount *int64   `json:"InvocationErrorCount,omitempty" xml:"InvocationErrorCount,omitempty"`
	ErrorRate            *float32 `json:"ErrorRate,omitempty" xml:"ErrorRate,omitempty"`
	CostAverage          *int64   `json:"CostAverage,omitempty" xml:"CostAverage,omitempty"`
}

func (s DescribeFlowMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponseBody) SetRequestId(v string) *DescribeFlowMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetInvocationCount(v int64) *DescribeFlowMetricResponseBody {
	s.InvocationCount = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetInvocationErrorCount(v int64) *DescribeFlowMetricResponseBody {
	s.InvocationErrorCount = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetErrorRate(v float32) *DescribeFlowMetricResponseBody {
	s.ErrorRate = &v
	return s
}

func (s *DescribeFlowMetricResponseBody) SetCostAverage(v int64) *DescribeFlowMetricResponseBody {
	s.CostAverage = &v
	return s
}

type DescribeFlowMetricResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowMetricResponse) SetHeaders(v map[string]*string) *DescribeFlowMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowMetricResponse) SetBody(v *DescribeFlowMetricResponseBody) *DescribeFlowMetricResponse {
	s.Body = v
	return s
}

type DescribeFlowTemplateRequest struct {
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeFlowTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowTemplateRequest) SetTemplateId(v string) *DescribeFlowTemplateRequest {
	s.TemplateId = &v
	return s
}

type DescribeFlowTemplateResponseBody struct {
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Connector   *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	SummaryEn   *string `json:"SummaryEn,omitempty" xml:"SummaryEn,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Overview    *string `json:"Overview,omitempty" xml:"Overview,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeFlowTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowTemplateResponseBody) SetLocale(v string) *DescribeFlowTemplateResponseBody {
	s.Locale = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetConnector(v string) *DescribeFlowTemplateResponseBody {
	s.Connector = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetSummaryEn(v string) *DescribeFlowTemplateResponseBody {
	s.SummaryEn = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetDescription(v string) *DescribeFlowTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetRequestId(v string) *DescribeFlowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetCreatedAt(v string) *DescribeFlowTemplateResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetDefinition(v string) *DescribeFlowTemplateResponseBody {
	s.Definition = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetOverview(v string) *DescribeFlowTemplateResponseBody {
	s.Overview = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetUpdatedAt(v string) *DescribeFlowTemplateResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetName(v string) *DescribeFlowTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetSummary(v string) *DescribeFlowTemplateResponseBody {
	s.Summary = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetTag(v string) *DescribeFlowTemplateResponseBody {
	s.Tag = &v
	return s
}

func (s *DescribeFlowTemplateResponseBody) SetTemplateId(v string) *DescribeFlowTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type DescribeFlowTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowTemplateResponse) SetHeaders(v map[string]*string) *DescribeFlowTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowTemplateResponse) SetBody(v *DescribeFlowTemplateResponseBody) *DescribeFlowTemplateResponse {
	s.Body = v
	return s
}

type DescribeInvocationLogRequest struct {
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	Pages        *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
}

func (s DescribeInvocationLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogRequest) SetInvocationId(v string) *DescribeInvocationLogRequest {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationLogRequest) SetPages(v string) *DescribeInvocationLogRequest {
	s.Pages = &v
	return s
}

type DescribeInvocationLogResponseBody struct {
	Status          *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Parameters      map[string]interface{}                            `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ReturnCode      *string                                           `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	EndTime         *int64                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Actions         []*DescribeInvocationLogResponseBodyActions       `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	TimeoutTime     *int64                                            `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
	StartTime       *int64                                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Workflow        *DescribeInvocationLogResponseBodyWorkflow        `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
	InvocationError *DescribeInvocationLogResponseBodyInvocationError `json:"InvocationError,omitempty" xml:"InvocationError,omitempty" type:"Struct"`
	Trigger         *DescribeInvocationLogResponseBodyTrigger         `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	ResponseResult  *DescribeInvocationLogResponseBodyResponseResult  `json:"ResponseResult,omitempty" xml:"ResponseResult,omitempty" type:"Struct"`
	InvocationId    *string                                           `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	OutputsResult   *DescribeInvocationLogResponseBodyOutputsResult   `json:"OutputsResult,omitempty" xml:"OutputsResult,omitempty" type:"Struct"`
}

func (s DescribeInvocationLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBody) SetStatus(v string) *DescribeInvocationLogResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetParameters(v map[string]interface{}) *DescribeInvocationLogResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetReturnCode(v string) *DescribeInvocationLogResponseBody {
	s.ReturnCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetEndTime(v int64) *DescribeInvocationLogResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetRequestId(v string) *DescribeInvocationLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetActions(v []*DescribeInvocationLogResponseBodyActions) *DescribeInvocationLogResponseBody {
	s.Actions = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetTimeoutTime(v int64) *DescribeInvocationLogResponseBody {
	s.TimeoutTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetStartTime(v int64) *DescribeInvocationLogResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetWorkflow(v *DescribeInvocationLogResponseBodyWorkflow) *DescribeInvocationLogResponseBody {
	s.Workflow = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetInvocationError(v *DescribeInvocationLogResponseBodyInvocationError) *DescribeInvocationLogResponseBody {
	s.InvocationError = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetTrigger(v *DescribeInvocationLogResponseBodyTrigger) *DescribeInvocationLogResponseBody {
	s.Trigger = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetResponseResult(v *DescribeInvocationLogResponseBodyResponseResult) *DescribeInvocationLogResponseBody {
	s.ResponseResult = v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetInvocationId(v string) *DescribeInvocationLogResponseBody {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationLogResponseBody) SetOutputsResult(v *DescribeInvocationLogResponseBodyOutputsResult) *DescribeInvocationLogResponseBody {
	s.OutputsResult = v
	return s
}

type DescribeInvocationLogResponseBodyActions struct {
	LoopCount     *int32                                                 `json:"LoopCount,omitempty" xml:"LoopCount,omitempty"`
	EndTime       *int64                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime     *int64                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InvocationId  *string                                                `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	ReturnCode    *string                                                `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	Error         *DescribeInvocationLogResponseBodyActionsError         `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	OutputsResult *DescribeInvocationLogResponseBodyActionsOutputsResult `json:"OutputsResult,omitempty" xml:"OutputsResult,omitempty" type:"Struct"`
	Name          *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	InputsResult  *DescribeInvocationLogResponseBodyActionsInputsResult  `json:"InputsResult,omitempty" xml:"InputsResult,omitempty" type:"Struct"`
}

func (s DescribeInvocationLogResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyActions) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyActions) SetLoopCount(v int32) *DescribeInvocationLogResponseBodyActions {
	s.LoopCount = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetEndTime(v int64) *DescribeInvocationLogResponseBodyActions {
	s.EndTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetStatus(v string) *DescribeInvocationLogResponseBodyActions {
	s.Status = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetStartTime(v int64) *DescribeInvocationLogResponseBodyActions {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetInvocationId(v string) *DescribeInvocationLogResponseBodyActions {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetReturnCode(v string) *DescribeInvocationLogResponseBodyActions {
	s.ReturnCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetError(v *DescribeInvocationLogResponseBodyActionsError) *DescribeInvocationLogResponseBodyActions {
	s.Error = v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetOutputsResult(v *DescribeInvocationLogResponseBodyActionsOutputsResult) *DescribeInvocationLogResponseBodyActions {
	s.OutputsResult = v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetName(v string) *DescribeInvocationLogResponseBodyActions {
	s.Name = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActions) SetInputsResult(v *DescribeInvocationLogResponseBodyActionsInputsResult) *DescribeInvocationLogResponseBodyActions {
	s.InputsResult = v
	return s
}

type DescribeInvocationLogResponseBodyActionsError struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeInvocationLogResponseBodyActionsError) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyActionsError) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyActionsError) SetErrorCode(v string) *DescribeInvocationLogResponseBodyActionsError {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActionsError) SetMessage(v string) *DescribeInvocationLogResponseBodyActionsError {
	s.Message = &v
	return s
}

type DescribeInvocationLogResponseBodyActionsOutputsResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyActionsOutputsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyActionsOutputsResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyActionsOutputsResult) SetContentType(v string) *DescribeInvocationLogResponseBodyActionsOutputsResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActionsOutputsResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyActionsOutputsResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActionsOutputsResult) SetUrl(v string) *DescribeInvocationLogResponseBodyActionsOutputsResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponseBodyActionsInputsResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyActionsInputsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyActionsInputsResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyActionsInputsResult) SetContentType(v string) *DescribeInvocationLogResponseBodyActionsInputsResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActionsInputsResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyActionsInputsResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyActionsInputsResult) SetUrl(v string) *DescribeInvocationLogResponseBodyActionsInputsResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponseBodyWorkflow struct {
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Version    *string `json:"Version,omitempty" xml:"Version,omitempty"`
	FlowId     *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DescribeInvocationLogResponseBodyWorkflow) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyWorkflow) SetDefinition(v string) *DescribeInvocationLogResponseBodyWorkflow {
	s.Definition = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyWorkflow) SetVersion(v string) *DescribeInvocationLogResponseBodyWorkflow {
	s.Version = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyWorkflow) SetFlowId(v string) *DescribeInvocationLogResponseBodyWorkflow {
	s.FlowId = &v
	return s
}

type DescribeInvocationLogResponseBodyInvocationError struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeInvocationLogResponseBodyInvocationError) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyInvocationError) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyInvocationError) SetErrorCode(v string) *DescribeInvocationLogResponseBodyInvocationError {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyInvocationError) SetMessage(v string) *DescribeInvocationLogResponseBodyInvocationError {
	s.Message = &v
	return s
}

type DescribeInvocationLogResponseBodyTrigger struct {
	EndTime       *int64                                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status        *string                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime     *int64                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InvocationId  *string                                                `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	ReturnCode    *string                                                `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	Error         *DescribeInvocationLogResponseBodyTriggerError         `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	OutputsResult *DescribeInvocationLogResponseBodyTriggerOutputsResult `json:"OutputsResult,omitempty" xml:"OutputsResult,omitempty" type:"Struct"`
	Name          *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	InputsResult  *DescribeInvocationLogResponseBodyTriggerInputsResult  `json:"InputsResult,omitempty" xml:"InputsResult,omitempty" type:"Struct"`
}

func (s DescribeInvocationLogResponseBodyTrigger) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyTrigger) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetEndTime(v int64) *DescribeInvocationLogResponseBodyTrigger {
	s.EndTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetStatus(v string) *DescribeInvocationLogResponseBodyTrigger {
	s.Status = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetStartTime(v int64) *DescribeInvocationLogResponseBodyTrigger {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetInvocationId(v string) *DescribeInvocationLogResponseBodyTrigger {
	s.InvocationId = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetReturnCode(v string) *DescribeInvocationLogResponseBodyTrigger {
	s.ReturnCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetError(v *DescribeInvocationLogResponseBodyTriggerError) *DescribeInvocationLogResponseBodyTrigger {
	s.Error = v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetOutputsResult(v *DescribeInvocationLogResponseBodyTriggerOutputsResult) *DescribeInvocationLogResponseBodyTrigger {
	s.OutputsResult = v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetName(v string) *DescribeInvocationLogResponseBodyTrigger {
	s.Name = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTrigger) SetInputsResult(v *DescribeInvocationLogResponseBodyTriggerInputsResult) *DescribeInvocationLogResponseBodyTrigger {
	s.InputsResult = v
	return s
}

type DescribeInvocationLogResponseBodyTriggerError struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeInvocationLogResponseBodyTriggerError) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyTriggerError) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyTriggerError) SetErrorCode(v string) *DescribeInvocationLogResponseBodyTriggerError {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTriggerError) SetMessage(v string) *DescribeInvocationLogResponseBodyTriggerError {
	s.Message = &v
	return s
}

type DescribeInvocationLogResponseBodyTriggerOutputsResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyTriggerOutputsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyTriggerOutputsResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyTriggerOutputsResult) SetContentType(v string) *DescribeInvocationLogResponseBodyTriggerOutputsResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTriggerOutputsResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyTriggerOutputsResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTriggerOutputsResult) SetUrl(v string) *DescribeInvocationLogResponseBodyTriggerOutputsResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponseBodyTriggerInputsResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyTriggerInputsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyTriggerInputsResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyTriggerInputsResult) SetContentType(v string) *DescribeInvocationLogResponseBodyTriggerInputsResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTriggerInputsResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyTriggerInputsResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyTriggerInputsResult) SetUrl(v string) *DescribeInvocationLogResponseBodyTriggerInputsResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponseBodyResponseResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyResponseResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyResponseResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyResponseResult) SetContentType(v string) *DescribeInvocationLogResponseBodyResponseResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyResponseResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyResponseResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyResponseResult) SetUrl(v string) *DescribeInvocationLogResponseBodyResponseResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponseBodyOutputsResult struct {
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	ContentSize *string `json:"ContentSize,omitempty" xml:"ContentSize,omitempty"`
	Url         *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeInvocationLogResponseBodyOutputsResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponseBodyOutputsResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponseBodyOutputsResult) SetContentType(v string) *DescribeInvocationLogResponseBodyOutputsResult {
	s.ContentType = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyOutputsResult) SetContentSize(v string) *DescribeInvocationLogResponseBodyOutputsResult {
	s.ContentSize = &v
	return s
}

func (s *DescribeInvocationLogResponseBodyOutputsResult) SetUrl(v string) *DescribeInvocationLogResponseBodyOutputsResult {
	s.Url = &v
	return s
}

type DescribeInvocationLogResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInvocationLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInvocationLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationLogResponse) SetHeaders(v map[string]*string) *DescribeInvocationLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationLogResponse) SetBody(v *DescribeInvocationLogResponseBody) *DescribeInvocationLogResponse {
	s.Body = v
	return s
}

type DescribeMetricDetailRequest struct {
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeMetricDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricDetailRequest) SetMetricName(v string) *DescribeMetricDetailRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricDetailRequest) SetPeriod(v string) *DescribeMetricDetailRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricDetailRequest) SetStartTime(v string) *DescribeMetricDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricDetailRequest) SetEndTime(v string) *DescribeMetricDetailRequest {
	s.EndTime = &v
	return s
}

type DescribeMetricDetailResponseBody struct {
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Datapoints []*DescribeMetricDetailResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Repeated"`
}

func (s DescribeMetricDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricDetailResponseBody) SetRequestId(v string) *DescribeMetricDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricDetailResponseBody) SetDatapoints(v []*DescribeMetricDetailResponseBodyDatapoints) *DescribeMetricDetailResponseBody {
	s.Datapoints = v
	return s
}

type DescribeMetricDetailResponseBodyDatapoints struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricDetailResponseBodyDatapoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDetailResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricDetailResponseBodyDatapoints) SetValue(v string) *DescribeMetricDetailResponseBodyDatapoints {
	s.Value = &v
	return s
}

func (s *DescribeMetricDetailResponseBodyDatapoints) SetTimestamp(v int64) *DescribeMetricDetailResponseBodyDatapoints {
	s.Timestamp = &v
	return s
}

type DescribeMetricDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetricDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetricDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetricDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetricDetailResponse) SetHeaders(v map[string]*string) *DescribeMetricDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetricDetailResponse) SetBody(v *DescribeMetricDetailResponseBody) *DescribeMetricDetailResponse {
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableFlowResponseBody) GoString() string {
	return s.String()
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFlowResponseBody) GoString() string {
	return s.String()
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

type InvokeFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Parameters  *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Data        *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
}

func (s InvokeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowRequest) GoString() string {
	return s.String()
}

func (s *InvokeFlowRequest) SetFlowId(v string) *InvokeFlowRequest {
	s.FlowId = &v
	return s
}

func (s *InvokeFlowRequest) SetParameters(v string) *InvokeFlowRequest {
	s.Parameters = &v
	return s
}

func (s *InvokeFlowRequest) SetData(v string) *InvokeFlowRequest {
	s.Data = &v
	return s
}

func (s *InvokeFlowRequest) SetClientToken(v string) *InvokeFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *InvokeFlowRequest) SetDefinition(v string) *InvokeFlowRequest {
	s.Definition = &v
	return s
}

type InvokeFlowResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvokeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeFlowResponseBody) SetRequestId(v string) *InvokeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeFlowResponseBody) SetInvocationId(v string) *InvokeFlowResponseBody {
	s.InvocationId = &v
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

type ListConnectorsRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListConnectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorsRequest) SetCategory(v string) *ListConnectorsRequest {
	s.Category = &v
	return s
}

func (s *ListConnectorsRequest) SetLang(v string) *ListConnectorsRequest {
	s.Lang = &v
	return s
}

type ListConnectorsResponseBody struct {
	Connectors []*ListConnectorsResponseBodyConnectors `json:"Connectors,omitempty" xml:"Connectors,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBody) SetConnectors(v []*ListConnectorsResponseBodyConnectors) *ListConnectorsResponseBody {
	s.Connectors = v
	return s
}

func (s *ListConnectorsResponseBody) SetRequestId(v string) *ListConnectorsResponseBody {
	s.RequestId = &v
	return s
}

type ListConnectorsResponseBodyConnectors struct {
	DisplayName          *string                                                   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Capabilities         []*string                                                 `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	ParentConnector      *string                                                   `json:"ParentConnector,omitempty" xml:"ParentConnector,omitempty"`
	FullName             *string                                                   `json:"FullName,omitempty" xml:"FullName,omitempty"`
	Icon                 *ListConnectorsResponseBodyConnectorsIcon                 `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	ConnectorId          *string                                                   `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	RegionId             *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Description          *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	BrandColor           *string                                                   `json:"BrandColor,omitempty" xml:"BrandColor,omitempty"`
	Category             *string                                                   `json:"Category,omitempty" xml:"Category,omitempty"`
	ConnectionParameters *ListConnectorsResponseBodyConnectorsConnectionParameters `json:"ConnectionParameters,omitempty" xml:"ConnectionParameters,omitempty" type:"Struct"`
	Name                 *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListConnectorsResponseBodyConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectors) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectors) SetDisplayName(v string) *ListConnectorsResponseBodyConnectors {
	s.DisplayName = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetCapabilities(v []*string) *ListConnectorsResponseBodyConnectors {
	s.Capabilities = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetParentConnector(v string) *ListConnectorsResponseBodyConnectors {
	s.ParentConnector = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetFullName(v string) *ListConnectorsResponseBodyConnectors {
	s.FullName = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetIcon(v *ListConnectorsResponseBodyConnectorsIcon) *ListConnectorsResponseBodyConnectors {
	s.Icon = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectorId(v string) *ListConnectorsResponseBodyConnectors {
	s.ConnectorId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetRegionId(v string) *ListConnectorsResponseBodyConnectors {
	s.RegionId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetDescription(v string) *ListConnectorsResponseBodyConnectors {
	s.Description = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetBrandColor(v string) *ListConnectorsResponseBodyConnectors {
	s.BrandColor = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetCategory(v string) *ListConnectorsResponseBodyConnectors {
	s.Category = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectionParameters(v *ListConnectorsResponseBodyConnectorsConnectionParameters) *ListConnectorsResponseBodyConnectors {
	s.ConnectionParameters = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetName(v string) *ListConnectorsResponseBodyConnectors {
	s.Name = &v
	return s
}

type ListConnectorsResponseBodyConnectorsIcon struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsIcon) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsIcon) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsIcon) SetType(v string) *ListConnectorsResponseBodyConnectorsIcon {
	s.Type = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsIcon) SetValue(v string) *ListConnectorsResponseBodyConnectorsIcon {
	s.Value = &v
	return s
}

type ListConnectorsResponseBodyConnectorsConnectionParameters struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsConnectionParameters) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsConnectionParameters) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsConnectionParameters) SetType(v string) *ListConnectorsResponseBodyConnectorsConnectionParameters {
	s.Type = &v
	return s
}

type ListConnectorsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponse) SetHeaders(v map[string]*string) *ListConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorsResponse) SetBody(v *ListConnectorsResponseBody) *ListConnectorsResponse {
	s.Body = v
	return s
}

type ListConnectorTriggersRequest struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListConnectorTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersRequest) SetCategory(v string) *ListConnectorTriggersRequest {
	s.Category = &v
	return s
}

func (s *ListConnectorTriggersRequest) SetLang(v string) *ListConnectorTriggersRequest {
	s.Lang = &v
	return s
}

type ListConnectorTriggersResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Triggers  []*ListConnectorTriggersResponseBodyTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListConnectorTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersResponseBody) SetRequestId(v string) *ListConnectorTriggersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorTriggersResponseBody) SetTriggers(v []*ListConnectorTriggersResponseBodyTriggers) *ListConnectorTriggersResponseBody {
	s.Triggers = v
	return s
}

type ListConnectorTriggersResponseBodyTriggers struct {
	DisplayName *string                                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Type        *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Visibility  *string                                             `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	Description *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentUrl *string                                             `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	Connector   *ListConnectorTriggersResponseBodyTriggersConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	Name        *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	System      *bool                                               `json:"System,omitempty" xml:"System,omitempty"`
}

func (s ListConnectorTriggersResponseBodyTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetDisplayName(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.DisplayName = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetType(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.Type = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetVisibility(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.Visibility = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetDescription(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.Description = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetDocumentUrl(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.DocumentUrl = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetConnector(v *ListConnectorTriggersResponseBodyTriggersConnector) *ListConnectorTriggersResponseBodyTriggers {
	s.Connector = v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetName(v string) *ListConnectorTriggersResponseBodyTriggers {
	s.Name = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggers) SetSystem(v bool) *ListConnectorTriggersResponseBodyTriggers {
	s.System = &v
	return s
}

type ListConnectorTriggersResponseBodyTriggersConnector struct {
	DisplayName *string                                                 `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	FullName    *string                                                 `json:"FullName,omitempty" xml:"FullName,omitempty"`
	BrandColor  *string                                                 `json:"BrandColor,omitempty" xml:"BrandColor,omitempty"`
	Icon        *ListConnectorTriggersResponseBodyTriggersConnectorIcon `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	Category    *string                                                 `json:"Category,omitempty" xml:"Category,omitempty"`
	Name        *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListConnectorTriggersResponseBodyTriggersConnector) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersResponseBodyTriggersConnector) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetDisplayName(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.DisplayName = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetDescription(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.Description = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetFullName(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.FullName = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetBrandColor(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.BrandColor = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetIcon(v *ListConnectorTriggersResponseBodyTriggersConnectorIcon) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.Icon = v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetCategory(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.Category = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetName(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.Name = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnector) SetRegionId(v string) *ListConnectorTriggersResponseBodyTriggersConnector {
	s.RegionId = &v
	return s
}

type ListConnectorTriggersResponseBodyTriggersConnectorIcon struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConnectorTriggersResponseBodyTriggersConnectorIcon) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersResponseBodyTriggersConnectorIcon) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersResponseBodyTriggersConnectorIcon) SetType(v string) *ListConnectorTriggersResponseBodyTriggersConnectorIcon {
	s.Type = &v
	return s
}

func (s *ListConnectorTriggersResponseBodyTriggersConnectorIcon) SetValue(v string) *ListConnectorTriggersResponseBodyTriggersConnectorIcon {
	s.Value = &v
	return s
}

type ListConnectorTriggersResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListConnectorTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConnectorTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorTriggersResponse) SetHeaders(v map[string]*string) *ListConnectorTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorTriggersResponse) SetBody(v *ListConnectorTriggersResponseBody) *ListConnectorTriggersResponse {
	s.Body = v
	return s
}

type ListFlowRequest struct {
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRequest) SetPageSize(v int32) *ListFlowRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowRequest) SetPageNumber(v int32) *ListFlowRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowRequest) SetName(v string) *ListFlowRequest {
	s.Name = &v
	return s
}

type ListFlowResponseBody struct {
	TotalCount *int32                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                       `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Flows      []*ListFlowResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
}

func (s ListFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBody) SetTotalCount(v int32) *ListFlowResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFlowResponseBody) SetTotalPage(v int32) *ListFlowResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListFlowResponseBody) SetRequestId(v string) *ListFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowResponseBody) SetFlows(v []*ListFlowResponseBodyFlows) *ListFlowResponseBody {
	s.Flows = v
	return s
}

type ListFlowResponseBodyFlows struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ProdVersion    *int32  `json:"ProdVersion,omitempty" xml:"ProdVersion,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedAt      *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CurrentVersion *int32  `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	EditMode       *string `json:"EditMode,omitempty" xml:"EditMode,omitempty"`
	UpdatedAt      *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Source         *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	FlowId         *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ListFlowResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyFlows) SetStatus(v string) *ListFlowResponseBodyFlows {
	s.Status = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetProdVersion(v int32) *ListFlowResponseBodyFlows {
	s.ProdVersion = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetDescription(v string) *ListFlowResponseBodyFlows {
	s.Description = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetCreatedAt(v string) *ListFlowResponseBodyFlows {
	s.CreatedAt = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetCurrentVersion(v int32) *ListFlowResponseBodyFlows {
	s.CurrentVersion = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetEditMode(v string) *ListFlowResponseBodyFlows {
	s.EditMode = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetUpdatedAt(v string) *ListFlowResponseBodyFlows {
	s.UpdatedAt = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetSource(v string) *ListFlowResponseBodyFlows {
	s.Source = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetName(v string) *ListFlowResponseBodyFlows {
	s.Name = &v
	return s
}

func (s *ListFlowResponseBodyFlows) SetFlowId(v string) *ListFlowResponseBodyFlows {
	s.FlowId = &v
	return s
}

type ListFlowResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponse) GoString() string {
	return s.String()
}

func (s *ListFlowResponse) SetHeaders(v map[string]*string) *ListFlowResponse {
	s.Headers = v
	return s
}

func (s *ListFlowResponse) SetBody(v *ListFlowResponseBody) *ListFlowResponse {
	s.Body = v
	return s
}

type ListFlowConnectionsRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ListFlowConnectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsRequest) SetFlowId(v string) *ListFlowConnectionsRequest {
	s.FlowId = &v
	return s
}

type ListFlowConnectionsResponseBody struct {
	Connections []*ListFlowConnectionsResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBody) SetConnections(v []*ListFlowConnectionsResponseBodyConnections) *ListFlowConnectionsResponseBody {
	s.Connections = v
	return s
}

func (s *ListFlowConnectionsResponseBody) SetRequestId(v string) *ListFlowConnectionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowConnectionsResponseBodyConnections struct {
	Definition     *string                                              `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Actions        []*ListFlowConnectionsResponseBodyConnectionsActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	ConnectionName *string                                              `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	Connector      *ListFlowConnectionsResponseBodyConnectionsConnector `json:"Connector,omitempty" xml:"Connector,omitempty" type:"Struct"`
	Content        *string                                              `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ListFlowConnectionsResponseBodyConnections) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBodyConnections) SetDefinition(v string) *ListFlowConnectionsResponseBodyConnections {
	s.Definition = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnections) SetActions(v []*ListFlowConnectionsResponseBodyConnectionsActions) *ListFlowConnectionsResponseBodyConnections {
	s.Actions = v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnections) SetConnectionName(v string) *ListFlowConnectionsResponseBodyConnections {
	s.ConnectionName = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnections) SetConnector(v *ListFlowConnectionsResponseBodyConnectionsConnector) *ListFlowConnectionsResponseBodyConnections {
	s.Connector = v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnections) SetContent(v string) *ListFlowConnectionsResponseBodyConnections {
	s.Content = &v
	return s
}

type ListFlowConnectionsResponseBodyConnectionsActions struct {
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ActionName *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
}

func (s ListFlowConnectionsResponseBodyConnectionsActions) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBodyConnectionsActions) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBodyConnectionsActions) SetType(v string) *ListFlowConnectionsResponseBodyConnectionsActions {
	s.Type = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsActions) SetActionName(v string) *ListFlowConnectionsResponseBodyConnectionsActions {
	s.ActionName = &v
	return s
}

type ListFlowConnectionsResponseBodyConnectionsConnector struct {
	DisplayName          *string                                                                  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	BrandColor           *string                                                                  `json:"BrandColor,omitempty" xml:"BrandColor,omitempty"`
	Icon                 *ListFlowConnectionsResponseBodyConnectionsConnectorIcon                 `json:"Icon,omitempty" xml:"Icon,omitempty" type:"Struct"`
	Name                 *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	ConnectionParameters *ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters `json:"ConnectionParameters,omitempty" xml:"ConnectionParameters,omitempty" type:"Struct"`
}

func (s ListFlowConnectionsResponseBodyConnectionsConnector) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBodyConnectionsConnector) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnector) SetDisplayName(v string) *ListFlowConnectionsResponseBodyConnectionsConnector {
	s.DisplayName = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnector) SetBrandColor(v string) *ListFlowConnectionsResponseBodyConnectionsConnector {
	s.BrandColor = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnector) SetIcon(v *ListFlowConnectionsResponseBodyConnectionsConnectorIcon) *ListFlowConnectionsResponseBodyConnectionsConnector {
	s.Icon = v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnector) SetName(v string) *ListFlowConnectionsResponseBodyConnectionsConnector {
	s.Name = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnector) SetConnectionParameters(v *ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters) *ListFlowConnectionsResponseBodyConnectionsConnector {
	s.ConnectionParameters = v
	return s
}

type ListFlowConnectionsResponseBodyConnectionsConnectorIcon struct {
	Type  *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFlowConnectionsResponseBodyConnectionsConnectorIcon) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBodyConnectionsConnectorIcon) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnectorIcon) SetType(v string) *ListFlowConnectionsResponseBodyConnectionsConnectorIcon {
	s.Type = &v
	return s
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnectorIcon) SetValue(v string) *ListFlowConnectionsResponseBodyConnectionsConnectorIcon {
	s.Value = &v
	return s
}

type ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters) SetType(v string) *ListFlowConnectionsResponseBodyConnectionsConnectorConnectionParameters {
	s.Type = &v
	return s
}

type ListFlowConnectionsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowConnectionsResponse) SetHeaders(v map[string]*string) *ListFlowConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowConnectionsResponse) SetBody(v *ListFlowConnectionsResponseBody) *ListFlowConnectionsResponse {
	s.Body = v
	return s
}

type ListFlowTemplateRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListFlowTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListFlowTemplateRequest) SetPageNumber(v int32) *ListFlowTemplateRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowTemplateRequest) SetPageSize(v int32) *ListFlowTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowTemplateRequest) SetName(v string) *ListFlowTemplateRequest {
	s.Name = &v
	return s
}

func (s *ListFlowTemplateRequest) SetTag(v string) *ListFlowTemplateRequest {
	s.Tag = &v
	return s
}

func (s *ListFlowTemplateRequest) SetLang(v string) *ListFlowTemplateRequest {
	s.Lang = &v
	return s
}

type ListFlowTemplateResponseBody struct {
	TotalPage     *int32                                       `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowTemplates []*ListFlowTemplateResponseBodyFlowTemplates `json:"FlowTemplates,omitempty" xml:"FlowTemplates,omitempty" type:"Repeated"`
}

func (s ListFlowTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowTemplateResponseBody) SetTotalPage(v int32) *ListFlowTemplateResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListFlowTemplateResponseBody) SetRequestId(v string) *ListFlowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowTemplateResponseBody) SetFlowTemplates(v []*ListFlowTemplateResponseBodyFlowTemplates) *ListFlowTemplateResponseBody {
	s.FlowTemplates = v
	return s
}

type ListFlowTemplateResponseBodyFlowTemplates struct {
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Locale      *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	SummaryEn   *string `json:"SummaryEn,omitempty" xml:"SummaryEn,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Overview    *string `json:"Overview,omitempty" xml:"Overview,omitempty"`
	Connector   *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Version     *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	UpdatedAt   *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId  *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListFlowTemplateResponseBodyFlowTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTemplateResponseBodyFlowTemplates) GoString() string {
	return s.String()
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetSummary(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Summary = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetLocale(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Locale = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetSummaryEn(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.SummaryEn = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetCreatedAt(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.CreatedAt = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetOverview(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Overview = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetConnector(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Connector = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetTag(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Tag = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetCreator(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Creator = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetDescription(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Description = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetVersion(v int32) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Version = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetUpdatedAt(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.UpdatedAt = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetName(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.Name = &v
	return s
}

func (s *ListFlowTemplateResponseBodyFlowTemplates) SetTemplateId(v string) *ListFlowTemplateResponseBodyFlowTemplates {
	s.TemplateId = &v
	return s
}

type ListFlowTemplateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListFlowTemplateResponse) SetHeaders(v map[string]*string) *ListFlowTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListFlowTemplateResponse) SetBody(v *ListFlowTemplateResponseBody) *ListFlowTemplateResponse {
	s.Body = v
	return s
}

type ListFlowTriggersRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ListFlowTriggersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListFlowTriggersRequest) SetFlowId(v string) *ListFlowTriggersRequest {
	s.FlowId = &v
	return s
}

type ListFlowTriggersResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Triggers  []*ListFlowTriggersResponseBodyTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListFlowTriggersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTriggersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowTriggersResponseBody) SetRequestId(v string) *ListFlowTriggersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowTriggersResponseBody) SetTriggers(v []*ListFlowTriggersResponseBodyTriggers) *ListFlowTriggersResponseBody {
	s.Triggers = v
	return s
}

type ListFlowTriggersResponseBodyTriggers struct {
	TriggerName              *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	Endpoint                 *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ActionsCount             *int32  `json:"ActionsCount,omitempty" xml:"ActionsCount,omitempty"`
	TriggerActionName        *string `json:"TriggerActionName,omitempty" xml:"TriggerActionName,omitempty"`
	TriggerDescription       *string `json:"TriggerDescription,omitempty" xml:"TriggerDescription,omitempty"`
	TriggerActionDescription *string `json:"TriggerActionDescription,omitempty" xml:"TriggerActionDescription,omitempty"`
	TriggerType              *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListFlowTriggersResponseBodyTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTriggersResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *ListFlowTriggersResponseBodyTriggers) SetTriggerName(v string) *ListFlowTriggersResponseBodyTriggers {
	s.TriggerName = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetEndpoint(v string) *ListFlowTriggersResponseBodyTriggers {
	s.Endpoint = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetActionsCount(v int32) *ListFlowTriggersResponseBodyTriggers {
	s.ActionsCount = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetTriggerActionName(v string) *ListFlowTriggersResponseBodyTriggers {
	s.TriggerActionName = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetTriggerDescription(v string) *ListFlowTriggersResponseBodyTriggers {
	s.TriggerDescription = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetTriggerActionDescription(v string) *ListFlowTriggersResponseBodyTriggers {
	s.TriggerActionDescription = &v
	return s
}

func (s *ListFlowTriggersResponseBodyTriggers) SetTriggerType(v string) *ListFlowTriggersResponseBodyTriggers {
	s.TriggerType = &v
	return s
}

type ListFlowTriggersResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowTriggersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowTriggersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowTriggersResponse) GoString() string {
	return s.String()
}

func (s *ListFlowTriggersResponse) SetHeaders(v map[string]*string) *ListFlowTriggersResponse {
	s.Headers = v
	return s
}

func (s *ListFlowTriggersResponse) SetBody(v *ListFlowTriggersResponseBody) *ListFlowTriggersResponse {
	s.Body = v
	return s
}

type ListFlowVersionRequest struct {
	FlowId     *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *ListFlowVersionRequest) SetFlowId(v string) *ListFlowVersionRequest {
	s.FlowId = &v
	return s
}

func (s *ListFlowVersionRequest) SetPageNumber(v int32) *ListFlowVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowVersionRequest) SetPageSize(v int32) *ListFlowVersionRequest {
	s.PageSize = &v
	return s
}

type ListFlowVersionResponseBody struct {
	Versions  []*ListFlowVersionResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	TotalPage *int32                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlowVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowVersionResponseBody) SetVersions(v []*ListFlowVersionResponseBodyVersions) *ListFlowVersionResponseBody {
	s.Versions = v
	return s
}

func (s *ListFlowVersionResponseBody) SetTotalPage(v int32) *ListFlowVersionResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListFlowVersionResponseBody) SetRequestId(v string) *ListFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowVersionResponseBodyVersions struct {
	Version   *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
	State     *int32  `json:"State,omitempty" xml:"State,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ListFlowVersionResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListFlowVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListFlowVersionResponseBodyVersions) SetVersion(v int32) *ListFlowVersionResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *ListFlowVersionResponseBodyVersions) SetState(v int32) *ListFlowVersionResponseBodyVersions {
	s.State = &v
	return s
}

func (s *ListFlowVersionResponseBodyVersions) SetVersionId(v string) *ListFlowVersionResponseBodyVersions {
	s.VersionId = &v
	return s
}

func (s *ListFlowVersionResponseBodyVersions) SetCreatedAt(v string) *ListFlowVersionResponseBodyVersions {
	s.CreatedAt = &v
	return s
}

func (s *ListFlowVersionResponseBodyVersions) SetUpdatedAt(v string) *ListFlowVersionResponseBodyVersions {
	s.UpdatedAt = &v
	return s
}

func (s *ListFlowVersionResponseBodyVersions) SetFlowId(v string) *ListFlowVersionResponseBodyVersions {
	s.FlowId = &v
	return s
}

type ListFlowVersionResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowVersionResponse) GoString() string {
	return s.String()
}

func (s *ListFlowVersionResponse) SetHeaders(v map[string]*string) *ListFlowVersionResponse {
	s.Headers = v
	return s
}

func (s *ListFlowVersionResponse) SetBody(v *ListFlowVersionResponseBody) *ListFlowVersionResponse {
	s.Body = v
	return s
}

type ListInvocationLogsRequest struct {
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回数量
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 执行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 工作流版本
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// 执行开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 结束执行时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 标签
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListInvocationLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsRequest) SetFlowId(v string) *ListInvocationLogsRequest {
	s.FlowId = &v
	return s
}

func (s *ListInvocationLogsRequest) SetPageNumber(v int32) *ListInvocationLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationLogsRequest) SetPageSize(v int32) *ListInvocationLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInvocationLogsRequest) SetStatus(v string) *ListInvocationLogsRequest {
	s.Status = &v
	return s
}

func (s *ListInvocationLogsRequest) SetFlowVersion(v string) *ListInvocationLogsRequest {
	s.FlowVersion = &v
	return s
}

func (s *ListInvocationLogsRequest) SetStartTime(v string) *ListInvocationLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListInvocationLogsRequest) SetEndTime(v string) *ListInvocationLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListInvocationLogsRequest) SetTags(v string) *ListInvocationLogsRequest {
	s.Tags = &v
	return s
}

type ListInvocationLogsResponseBody struct {
	// 总页数
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// 请求 ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 日志列表
	Logs []*ListInvocationLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// 总数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInvocationLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponseBody) SetTotalPage(v int64) *ListInvocationLogsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListInvocationLogsResponseBody) SetRequestId(v string) *ListInvocationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvocationLogsResponseBody) SetLogs(v []*ListInvocationLogsResponseBodyLogs) *ListInvocationLogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListInvocationLogsResponseBody) SetTotalCount(v int64) *ListInvocationLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListInvocationLogsResponseBodyLogs struct {
	// 执行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 执行完成时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 开始执行时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 执行唯一标识符
	InvocationId *string `json:"InvocationId,omitempty" xml:"InvocationId,omitempty"`
	// 返回码
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// 错误信息
	InvocationError *ListInvocationLogsResponseBodyLogsInvocationError `json:"InvocationError,omitempty" xml:"InvocationError,omitempty" type:"Struct"`
	// 工作流详情
	Workflow *ListInvocationLogsResponseBodyLogsWorkflow `json:"Workflow,omitempty" xml:"Workflow,omitempty" type:"Struct"`
	// 本次执行的标签
	Tags []*ListInvocationLogsResponseBodyLogsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListInvocationLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponseBodyLogs) SetStatus(v string) *ListInvocationLogsResponseBodyLogs {
	s.Status = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetEndTime(v string) *ListInvocationLogsResponseBodyLogs {
	s.EndTime = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetStartTime(v string) *ListInvocationLogsResponseBodyLogs {
	s.StartTime = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetInvocationId(v string) *ListInvocationLogsResponseBodyLogs {
	s.InvocationId = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetReturnCode(v string) *ListInvocationLogsResponseBodyLogs {
	s.ReturnCode = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetInvocationError(v *ListInvocationLogsResponseBodyLogsInvocationError) *ListInvocationLogsResponseBodyLogs {
	s.InvocationError = v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetWorkflow(v *ListInvocationLogsResponseBodyLogsWorkflow) *ListInvocationLogsResponseBodyLogs {
	s.Workflow = v
	return s
}

func (s *ListInvocationLogsResponseBodyLogs) SetTags(v []*ListInvocationLogsResponseBodyLogsTags) *ListInvocationLogsResponseBodyLogs {
	s.Tags = v
	return s
}

type ListInvocationLogsResponseBodyLogsInvocationError struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListInvocationLogsResponseBodyLogsInvocationError) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponseBodyLogsInvocationError) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponseBodyLogsInvocationError) SetErrorCode(v string) *ListInvocationLogsResponseBodyLogsInvocationError {
	s.ErrorCode = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogsInvocationError) SetMessage(v string) *ListInvocationLogsResponseBodyLogsInvocationError {
	s.Message = &v
	return s
}

type ListInvocationLogsResponseBodyLogsWorkflow struct {
	// 工作流定义
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// 工作流版本
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// 工作流 ID
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s ListInvocationLogsResponseBodyLogsWorkflow) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponseBodyLogsWorkflow) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponseBodyLogsWorkflow) SetDefinition(v string) *ListInvocationLogsResponseBodyLogsWorkflow {
	s.Definition = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogsWorkflow) SetVersion(v string) *ListInvocationLogsResponseBodyLogsWorkflow {
	s.Version = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogsWorkflow) SetFlowId(v string) *ListInvocationLogsResponseBodyLogsWorkflow {
	s.FlowId = &v
	return s
}

type ListInvocationLogsResponseBodyLogsTags struct {
	// 标签名
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInvocationLogsResponseBodyLogsTags) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponseBodyLogsTags) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponseBodyLogsTags) SetKey(v string) *ListInvocationLogsResponseBodyLogsTags {
	s.Key = &v
	return s
}

func (s *ListInvocationLogsResponseBodyLogsTags) SetValue(v string) *ListInvocationLogsResponseBodyLogsTags {
	s.Value = &v
	return s
}

type ListInvocationLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInvocationLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationLogsResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationLogsResponse) SetHeaders(v map[string]*string) *ListInvocationLogsResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationLogsResponse) SetBody(v *ListInvocationLogsResponseBody) *ListInvocationLogsResponse {
	s.Body = v
	return s
}

type ListTagResponseBody struct {
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResponseBody) SetRequestId(v string) *ListTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResponseBody) SetTags(v []*ListTagResponseBodyTags) *ListTagResponseBody {
	s.Tags = v
	return s
}

type ListTagResponseBodyTags struct {
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id        *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListTagResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagResponseBodyTags) SetCreatedAt(v string) *ListTagResponseBodyTags {
	s.CreatedAt = &v
	return s
}

func (s *ListTagResponseBodyTags) SetName(v string) *ListTagResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListTagResponseBodyTags) SetId(v int32) *ListTagResponseBodyTags {
	s.Id = &v
	return s
}

func (s *ListTagResponseBodyTags) SetCount(v int32) *ListTagResponseBodyTags {
	s.Count = &v
	return s
}

type ListTagResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResponse) GoString() string {
	return s.String()
}

func (s *ListTagResponse) SetHeaders(v map[string]*string) *ListTagResponse {
	s.Headers = v
	return s
}

func (s *ListTagResponse) SetBody(v *ListTagResponseBody) *ListTagResponse {
	s.Body = v
	return s
}

type ModifyFlowRequest struct {
	FlowId      *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	FlowRole    *string `json:"FlowRole,omitempty" xml:"FlowRole,omitempty"`
}

func (s ModifyFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRequest) SetFlowId(v string) *ModifyFlowRequest {
	s.FlowId = &v
	return s
}

func (s *ModifyFlowRequest) SetName(v string) *ModifyFlowRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowRequest) SetDescription(v string) *ModifyFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowRequest) SetDefinition(v string) *ModifyFlowRequest {
	s.Definition = &v
	return s
}

func (s *ModifyFlowRequest) SetFlowRole(v string) *ModifyFlowRequest {
	s.FlowRole = &v
	return s
}

type ModifyFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBody) SetRequestId(v string) *ModifyFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponse) SetHeaders(v map[string]*string) *ModifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowResponse) SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse {
	s.Body = v
	return s
}

type RollBackFlowRequest struct {
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s RollBackFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s RollBackFlowRequest) GoString() string {
	return s.String()
}

func (s *RollBackFlowRequest) SetFlowId(v string) *RollBackFlowRequest {
	s.FlowId = &v
	return s
}

type RollBackFlowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollBackFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollBackFlowResponseBody) GoString() string {
	return s.String()
}

func (s *RollBackFlowResponseBody) SetRequestId(v string) *RollBackFlowResponseBody {
	s.RequestId = &v
	return s
}

type RollBackFlowResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RollBackFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RollBackFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s RollBackFlowResponse) GoString() string {
	return s.String()
}

func (s *RollBackFlowResponse) SetHeaders(v map[string]*string) *RollBackFlowResponse {
	s.Headers = v
	return s
}

func (s *RollBackFlowResponse) SetBody(v *RollBackFlowResponseBody) *RollBackFlowResponse {
	s.Body = v
	return s
}

type UpdateConnectionRequest struct {
	Connector         *string `json:"Connector,omitempty" xml:"Connector,omitempty"`
	ConnectionName    *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	ConnectionContent *string `json:"ConnectionContent,omitempty" xml:"ConnectionContent,omitempty"`
	ActionType        *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	ConnectionType    *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
}

func (s UpdateConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionRequest) SetConnector(v string) *UpdateConnectionRequest {
	s.Connector = &v
	return s
}

func (s *UpdateConnectionRequest) SetConnectionName(v string) *UpdateConnectionRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionRequest) SetConnectionContent(v string) *UpdateConnectionRequest {
	s.ConnectionContent = &v
	return s
}

func (s *UpdateConnectionRequest) SetActionType(v string) *UpdateConnectionRequest {
	s.ActionType = &v
	return s
}

func (s *UpdateConnectionRequest) SetConnectionType(v string) *UpdateConnectionRequest {
	s.ConnectionType = &v
	return s
}

type UpdateConnectionResponseBody struct {
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	ConnectionId   *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Definition     *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
}

func (s UpdateConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponseBody) SetConnectionName(v string) *UpdateConnectionResponseBody {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetConnectionId(v string) *UpdateConnectionResponseBody {
	s.ConnectionId = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetRequestId(v string) *UpdateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetContent(v string) *UpdateConnectionResponseBody {
	s.Content = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetDefinition(v string) *UpdateConnectionResponseBody {
	s.Definition = &v
	return s
}

type UpdateConnectionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponse) SetHeaders(v map[string]*string) *UpdateConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectionResponse) SetBody(v *UpdateConnectionResponseBody) *UpdateConnectionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("logiccomposer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AbolishFlowWithOptions(request *AbolishFlowRequest, runtime *util.RuntimeOptions) (_result *AbolishFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AbolishFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AbolishFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbolishFlow(request *AbolishFlowRequest) (_result *AbolishFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishFlowResponse{}
	_body, _err := client.AbolishFlowWithOptions(request, runtime)
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

func (client *Client) DeployFlowWithOptions(request *DeployFlowRequest, runtime *util.RuntimeOptions) (_result *DeployFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeployFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeployFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployFlow(request *DeployFlowRequest) (_result *DeployFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployFlowResponse{}
	_body, _err := client.DeployFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountSummaryWithOptions(runtime *util.RuntimeOptions) (_result *DescribeAccountSummaryResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &DescribeAccountSummaryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccountSummary"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccountSummary() (_result *DescribeAccountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountSummaryResponse{}
	_body, _err := client.DescribeAccountSummaryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConnectorAttributeWithOptions(request *DescribeConnectorAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeConnectorAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConnectorAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConnectorAttribute"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectorAttribute(request *DescribeConnectorAttributeRequest) (_result *DescribeConnectorAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConnectorAttributeResponse{}
	_body, _err := client.DescribeConnectorAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConnectorCapabilityWithOptions(request *DescribeConnectorCapabilityRequest, runtime *util.RuntimeOptions) (_result *DescribeConnectorCapabilityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConnectorCapabilityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConnectorCapability"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectorCapability(request *DescribeConnectorCapabilityRequest) (_result *DescribeConnectorCapabilityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConnectorCapabilityResponse{}
	_body, _err := client.DescribeConnectorCapabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowWithOptions(request *DescribeFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlow(request *DescribeFlowRequest) (_result *DescribeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DescribeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowMetricWithOptions(request *DescribeFlowMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowMetric"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowMetric(request *DescribeFlowMetricRequest) (_result *DescribeFlowMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowMetricResponse{}
	_body, _err := client.DescribeFlowMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowTemplateWithOptions(request *DescribeFlowTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowTemplate"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowTemplate(request *DescribeFlowTemplateRequest) (_result *DescribeFlowTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowTemplateResponse{}
	_body, _err := client.DescribeFlowTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvocationLogWithOptions(request *DescribeInvocationLogRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInvocationLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInvocationLog"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvocationLog(request *DescribeInvocationLogRequest) (_result *DescribeInvocationLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationLogResponse{}
	_body, _err := client.DescribeInvocationLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetricDetailWithOptions(request *DescribeMetricDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeMetricDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetricDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetricDetail"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetricDetail(request *DescribeMetricDetailRequest) (_result *DescribeMetricDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetricDetailResponse{}
	_body, _err := client.DescribeMetricDetailWithOptions(request, runtime)
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

func (client *Client) ListConnectorsWithOptions(request *ListConnectorsRequest, runtime *util.RuntimeOptions) (_result *ListConnectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectorsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectors"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectors(request *ListConnectorsRequest) (_result *ListConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectorsResponse{}
	_body, _err := client.ListConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConnectorTriggersWithOptions(request *ListConnectorTriggersRequest, runtime *util.RuntimeOptions) (_result *ListConnectorTriggersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListConnectorTriggersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListConnectorTriggers"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConnectorTriggers(request *ListConnectorTriggersRequest) (_result *ListConnectorTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectorTriggersResponse{}
	_body, _err := client.ListConnectorTriggersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowWithOptions(request *ListFlowRequest, runtime *util.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlow(request *ListFlowRequest) (_result *ListFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowResponse{}
	_body, _err := client.ListFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowConnectionsWithOptions(request *ListFlowConnectionsRequest, runtime *util.RuntimeOptions) (_result *ListFlowConnectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowConnectionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowConnections"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowConnections(request *ListFlowConnectionsRequest) (_result *ListFlowConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowConnectionsResponse{}
	_body, _err := client.ListFlowConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowTemplateWithOptions(request *ListFlowTemplateRequest, runtime *util.RuntimeOptions) (_result *ListFlowTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowTemplate"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowTemplate(request *ListFlowTemplateRequest) (_result *ListFlowTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowTemplateResponse{}
	_body, _err := client.ListFlowTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowTriggersWithOptions(request *ListFlowTriggersRequest, runtime *util.RuntimeOptions) (_result *ListFlowTriggersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowTriggersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowTriggers"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowTriggers(request *ListFlowTriggersRequest) (_result *ListFlowTriggersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowTriggersResponse{}
	_body, _err := client.ListFlowTriggersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowVersionWithOptions(request *ListFlowVersionRequest, runtime *util.RuntimeOptions) (_result *ListFlowVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowVersion"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowVersion(request *ListFlowVersionRequest) (_result *ListFlowVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowVersionResponse{}
	_body, _err := client.ListFlowVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationLogsWithOptions(request *ListInvocationLogsRequest, runtime *util.RuntimeOptions) (_result *ListInvocationLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListInvocationLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInvocationLogs"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationLogs(request *ListInvocationLogsRequest) (_result *ListInvocationLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationLogsResponse{}
	_body, _err := client.ListInvocationLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagWithOptions(runtime *util.RuntimeOptions) (_result *ListTagResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTag"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTag() (_result *ListTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResponse{}
	_body, _err := client.ListTagWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowWithOptions(request *ModifyFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlow(request *ModifyFlowRequest) (_result *ModifyFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowResponse{}
	_body, _err := client.ModifyFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RollBackFlowWithOptions(request *RollBackFlowRequest, runtime *util.RuntimeOptions) (_result *RollBackFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RollBackFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RollBackFlow"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RollBackFlow(request *RollBackFlowRequest) (_result *RollBackFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollBackFlowResponse{}
	_body, _err := client.RollBackFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConnectionWithOptions(request *UpdateConnectionRequest, runtime *util.RuntimeOptions) (_result *UpdateConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateConnection"), tea.String("2018-12-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConnection(request *UpdateConnectionRequest) (_result *UpdateConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.UpdateConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
