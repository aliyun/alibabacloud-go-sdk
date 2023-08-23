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

type CampaignDetail struct {
	ActualEndTime      *int64  `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	ActualStartTime    *int64  `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	CasesAborted       *int64  `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	CasesConnected     *int64  `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	CasesUncompleted   *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletedRate      *int64  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	CreateTime         *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxAttemptCount    *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanedEndTime      *int64  `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	PlanedStartTime    *int64  `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	QueueName          *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	TotalCases         *int64  `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
	UpdateTime         *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CampaignDetail) String() string {
	return tea.Prettify(s)
}

func (s CampaignDetail) GoString() string {
	return s.String()
}

func (s *CampaignDetail) SetActualEndTime(v int64) *CampaignDetail {
	s.ActualEndTime = &v
	return s
}

func (s *CampaignDetail) SetActualStartTime(v int64) *CampaignDetail {
	s.ActualStartTime = &v
	return s
}

func (s *CampaignDetail) SetCasesAborted(v int64) *CampaignDetail {
	s.CasesAborted = &v
	return s
}

func (s *CampaignDetail) SetCasesConnected(v int64) *CampaignDetail {
	s.CasesConnected = &v
	return s
}

func (s *CampaignDetail) SetCasesUncompleted(v int64) *CampaignDetail {
	s.CasesUncompleted = &v
	return s
}

func (s *CampaignDetail) SetCompletedRate(v int64) *CampaignDetail {
	s.CompletedRate = &v
	return s
}

func (s *CampaignDetail) SetCreateTime(v int64) *CampaignDetail {
	s.CreateTime = &v
	return s
}

func (s *CampaignDetail) SetId(v string) *CampaignDetail {
	s.Id = &v
	return s
}

func (s *CampaignDetail) SetMaxAttemptCount(v int64) *CampaignDetail {
	s.MaxAttemptCount = &v
	return s
}

func (s *CampaignDetail) SetMinAttemptInterval(v int64) *CampaignDetail {
	s.MinAttemptInterval = &v
	return s
}

func (s *CampaignDetail) SetName(v string) *CampaignDetail {
	s.Name = &v
	return s
}

func (s *CampaignDetail) SetPlanedEndTime(v int64) *CampaignDetail {
	s.PlanedEndTime = &v
	return s
}

func (s *CampaignDetail) SetPlanedStartTime(v int64) *CampaignDetail {
	s.PlanedStartTime = &v
	return s
}

func (s *CampaignDetail) SetQueueName(v string) *CampaignDetail {
	s.QueueName = &v
	return s
}

func (s *CampaignDetail) SetState(v string) *CampaignDetail {
	s.State = &v
	return s
}

func (s *CampaignDetail) SetTotalCases(v int64) *CampaignDetail {
	s.TotalCases = &v
	return s
}

func (s *CampaignDetail) SetUpdateTime(v int64) *CampaignDetail {
	s.UpdateTime = &v
	return s
}

type AbortCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AbortCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortCampaignRequest) GoString() string {
	return s.String()
}

func (s *AbortCampaignRequest) SetCampaignId(v string) *AbortCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCampaignRequest) SetInstanceId(v string) *AbortCampaignRequest {
	s.InstanceId = &v
	return s
}

type AbortCampaignResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCampaignResponseBody) SetCode(v string) *AbortCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCampaignResponseBody) SetHttpStatusCode(v string) *AbortCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCampaignResponseBody) SetMessage(v string) *AbortCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCampaignResponseBody) SetRequestId(v string) *AbortCampaignResponseBody {
	s.RequestId = &v
	return s
}

type AbortCampaignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbortCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortCampaignResponse) GoString() string {
	return s.String()
}

func (s *AbortCampaignResponse) SetHeaders(v map[string]*string) *AbortCampaignResponse {
	s.Headers = v
	return s
}

func (s *AbortCampaignResponse) SetStatusCode(v int32) *AbortCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortCampaignResponse) SetBody(v *AbortCampaignResponseBody) *AbortCampaignResponse {
	s.Body = v
	return s
}

type AbortCasesRequest struct {
	CampaignId      *string   `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId      *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNumberList []*string `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty" type:"Repeated"`
}

func (s AbortCasesRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortCasesRequest) GoString() string {
	return s.String()
}

func (s *AbortCasesRequest) SetCampaignId(v string) *AbortCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesRequest) SetInstanceId(v string) *AbortCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesRequest) SetPhoneNumberList(v []*string) *AbortCasesRequest {
	s.PhoneNumberList = v
	return s
}

type AbortCasesShrinkRequest struct {
	CampaignId            *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNumberListShrink *string `json:"PhoneNumberList,omitempty" xml:"PhoneNumberList,omitempty"`
}

func (s AbortCasesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortCasesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AbortCasesShrinkRequest) SetCampaignId(v string) *AbortCasesShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetInstanceId(v string) *AbortCasesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AbortCasesShrinkRequest) SetPhoneNumberListShrink(v string) *AbortCasesShrinkRequest {
	s.PhoneNumberListShrink = &v
	return s
}

type AbortCasesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbortCasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortCasesResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCasesResponseBody) SetCode(v string) *AbortCasesResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCasesResponseBody) SetHttpStatusCode(v string) *AbortCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCasesResponseBody) SetMessage(v string) *AbortCasesResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCasesResponseBody) SetRequestId(v string) *AbortCasesResponseBody {
	s.RequestId = &v
	return s
}

type AbortCasesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbortCasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortCasesResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortCasesResponse) GoString() string {
	return s.String()
}

func (s *AbortCasesResponse) SetHeaders(v map[string]*string) *AbortCasesResponse {
	s.Headers = v
	return s
}

func (s *AbortCasesResponse) SetStatusCode(v int32) *AbortCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortCasesResponse) SetBody(v *AbortCasesResponseBody) *AbortCasesResponse {
	s.Body = v
	return s
}

type CheckDemoInstanceExistsResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool     `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDemoInstanceExistsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDemoInstanceExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDemoInstanceExistsResponseBody) SetCode(v string) *CheckDemoInstanceExistsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetData(v bool) *CheckDemoInstanceExistsResponseBody {
	s.Data = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetHttpStatusCode(v int32) *CheckDemoInstanceExistsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetMessage(v string) *CheckDemoInstanceExistsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetParams(v []*string) *CheckDemoInstanceExistsResponseBody {
	s.Params = v
	return s
}

func (s *CheckDemoInstanceExistsResponseBody) SetRequestId(v string) *CheckDemoInstanceExistsResponseBody {
	s.RequestId = &v
	return s
}

type CheckDemoInstanceExistsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckDemoInstanceExistsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDemoInstanceExistsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDemoInstanceExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckDemoInstanceExistsResponse) SetHeaders(v map[string]*string) *CheckDemoInstanceExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckDemoInstanceExistsResponse) SetStatusCode(v int32) *CheckDemoInstanceExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDemoInstanceExistsResponse) SetBody(v *CheckDemoInstanceExistsResponseBody) *CheckDemoInstanceExistsResponse {
	s.Body = v
	return s
}

type CheckMQRoleAssumptionAuthorityResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckMQRoleAssumptionAuthorityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckMQRoleAssumptionAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetCode(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Code = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetHttpStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetMessage(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Message = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetParams(v []*string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.Params = v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponseBody) SetRequestId(v string) *CheckMQRoleAssumptionAuthorityResponseBody {
	s.RequestId = &v
	return s
}

type CheckMQRoleAssumptionAuthorityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckMQRoleAssumptionAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckMQRoleAssumptionAuthorityResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckMQRoleAssumptionAuthorityResponse) GoString() string {
	return s.String()
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetHeaders(v map[string]*string) *CheckMQRoleAssumptionAuthorityResponse {
	s.Headers = v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetStatusCode(v int32) *CheckMQRoleAssumptionAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMQRoleAssumptionAuthorityResponse) SetBody(v *CheckMQRoleAssumptionAuthorityResponseBody) *CheckMQRoleAssumptionAuthorityResponse {
	s.Body = v
	return s
}

type CreateCampaignRequest struct {
	CallableTime          *string                          `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	CaseFileKey           *string                          `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	CaseList              []*CreateCampaignRequestCaseList `json:"CaseList,omitempty" xml:"CaseList,omitempty" type:"Repeated"`
	ContactFlowId         *string                          `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	EndTime               *string                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutingUntilTimeout *bool                            `json:"ExecutingUntilTimeout,omitempty" xml:"ExecutingUntilTimeout,omitempty"`
	InstanceId            *string                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxAttemptCount       *int64                           `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval    *int64                           `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name                  *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	QueueId               *string                          `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	Simulation            *bool                            `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters  *string                          `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	StartTime             *string                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyParameters    *string                          `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType          *string                          `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequest) SetCallableTime(v string) *CreateCampaignRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignRequest) SetCaseFileKey(v string) *CreateCampaignRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignRequest) SetCaseList(v []*CreateCampaignRequestCaseList) *CreateCampaignRequest {
	s.CaseList = v
	return s
}

func (s *CreateCampaignRequest) SetContactFlowId(v string) *CreateCampaignRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CreateCampaignRequest) SetEndTime(v string) *CreateCampaignRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignRequest) SetExecutingUntilTimeout(v bool) *CreateCampaignRequest {
	s.ExecutingUntilTimeout = &v
	return s
}

func (s *CreateCampaignRequest) SetInstanceId(v string) *CreateCampaignRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignRequest) SetMaxAttemptCount(v int64) *CreateCampaignRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignRequest) SetMinAttemptInterval(v int64) *CreateCampaignRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignRequest) SetName(v string) *CreateCampaignRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignRequest) SetQueueId(v string) *CreateCampaignRequest {
	s.QueueId = &v
	return s
}

func (s *CreateCampaignRequest) SetSimulation(v bool) *CreateCampaignRequest {
	s.Simulation = &v
	return s
}

func (s *CreateCampaignRequest) SetSimulationParameters(v string) *CreateCampaignRequest {
	s.SimulationParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetStartTime(v string) *CreateCampaignRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignRequest) SetStrategyParameters(v string) *CreateCampaignRequest {
	s.StrategyParameters = &v
	return s
}

func (s *CreateCampaignRequest) SetStrategyType(v string) *CreateCampaignRequest {
	s.StrategyType = &v
	return s
}

type CreateCampaignRequestCaseList struct {
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	PhoneNumber     *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ReferenceId     *string `json:"ReferenceId,omitempty" xml:"ReferenceId,omitempty"`
}

func (s CreateCampaignRequestCaseList) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignRequestCaseList) GoString() string {
	return s.String()
}

func (s *CreateCampaignRequestCaseList) SetCustomVariables(v string) *CreateCampaignRequestCaseList {
	s.CustomVariables = &v
	return s
}

func (s *CreateCampaignRequestCaseList) SetPhoneNumber(v string) *CreateCampaignRequestCaseList {
	s.PhoneNumber = &v
	return s
}

func (s *CreateCampaignRequestCaseList) SetReferenceId(v string) *CreateCampaignRequestCaseList {
	s.ReferenceId = &v
	return s
}

type CreateCampaignShrinkRequest struct {
	CallableTime          *string `json:"CallableTime,omitempty" xml:"CallableTime,omitempty"`
	CaseFileKey           *string `json:"CaseFileKey,omitempty" xml:"CaseFileKey,omitempty"`
	CaseListShrink        *string `json:"CaseList,omitempty" xml:"CaseList,omitempty"`
	ContactFlowId         *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExecutingUntilTimeout *bool   `json:"ExecutingUntilTimeout,omitempty" xml:"ExecutingUntilTimeout,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxAttemptCount       *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval    *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name                  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	QueueId               *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	Simulation            *bool   `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters  *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyParameters    *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType          *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
}

func (s CreateCampaignShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCampaignShrinkRequest) SetCallableTime(v string) *CreateCampaignShrinkRequest {
	s.CallableTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCaseFileKey(v string) *CreateCampaignShrinkRequest {
	s.CaseFileKey = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetCaseListShrink(v string) *CreateCampaignShrinkRequest {
	s.CaseListShrink = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetContactFlowId(v string) *CreateCampaignShrinkRequest {
	s.ContactFlowId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetEndTime(v string) *CreateCampaignShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetExecutingUntilTimeout(v bool) *CreateCampaignShrinkRequest {
	s.ExecutingUntilTimeout = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetInstanceId(v string) *CreateCampaignShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMaxAttemptCount(v int64) *CreateCampaignShrinkRequest {
	s.MaxAttemptCount = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetMinAttemptInterval(v int64) *CreateCampaignShrinkRequest {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetName(v string) *CreateCampaignShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetQueueId(v string) *CreateCampaignShrinkRequest {
	s.QueueId = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetSimulation(v bool) *CreateCampaignShrinkRequest {
	s.Simulation = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetSimulationParameters(v string) *CreateCampaignShrinkRequest {
	s.SimulationParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStartTime(v string) *CreateCampaignShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStrategyParameters(v string) *CreateCampaignShrinkRequest {
	s.StrategyParameters = &v
	return s
}

func (s *CreateCampaignShrinkRequest) SetStrategyType(v string) *CreateCampaignShrinkRequest {
	s.StrategyType = &v
	return s
}

type CreateCampaignResponseBody struct {
	CampaignId     *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBody) SetCampaignId(v string) *CreateCampaignResponseBody {
	s.CampaignId = &v
	return s
}

func (s *CreateCampaignResponseBody) SetCode(v string) *CreateCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCampaignResponseBody) SetHttpStatusCode(v int64) *CreateCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCampaignResponseBody) SetMessage(v string) *CreateCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCampaignResponseBody) SetRequestId(v string) *CreateCampaignResponseBody {
	s.RequestId = &v
	return s
}

type CreateCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCampaignResponse) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponse) SetHeaders(v map[string]*string) *CreateCampaignResponse {
	s.Headers = v
	return s
}

func (s *CreateCampaignResponse) SetStatusCode(v int32) *CreateCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCampaignResponse) SetBody(v *CreateCampaignResponseBody) *CreateCampaignResponse {
	s.Body = v
	return s
}

type CreateCorpNumberGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCorpNumberGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpNumberGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupRequest) SetDescription(v string) *CreateCorpNumberGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateCorpNumberGroupRequest) SetName(v string) *CreateCorpNumberGroupRequest {
	s.Name = &v
	return s
}

type CreateCorpNumberGroupResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateCorpNumberGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCorpNumberGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpNumberGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponseBody) SetCode(v string) *CreateCorpNumberGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetData(v *CreateCorpNumberGroupResponseBodyData) *CreateCorpNumberGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetHttpStatusCode(v int32) *CreateCorpNumberGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetMessage(v string) *CreateCorpNumberGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBody) SetRequestId(v string) *CreateCorpNumberGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateCorpNumberGroupResponseBodyData struct {
	AliyunUid       *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NumberCount     *string `json:"NumberCount,omitempty" xml:"NumberCount,omitempty"`
	NumberGroupId   *string `json:"NumberGroupId,omitempty" xml:"NumberGroupId,omitempty"`
	NumberGroupName *string `json:"NumberGroupName,omitempty" xml:"NumberGroupName,omitempty"`
}

func (s CreateCorpNumberGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpNumberGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponseBodyData) SetAliyunUid(v string) *CreateCorpNumberGroupResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetDescription(v string) *CreateCorpNumberGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberCount(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberCount = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberGroupId(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberGroupId = &v
	return s
}

func (s *CreateCorpNumberGroupResponseBodyData) SetNumberGroupName(v string) *CreateCorpNumberGroupResponseBodyData {
	s.NumberGroupName = &v
	return s
}

type CreateCorpNumberGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCorpNumberGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCorpNumberGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCorpNumberGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateCorpNumberGroupResponse) SetHeaders(v map[string]*string) *CreateCorpNumberGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateCorpNumberGroupResponse) SetStatusCode(v int32) *CreateCorpNumberGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCorpNumberGroupResponse) SetBody(v *CreateCorpNumberGroupResponseBody) *CreateCorpNumberGroupResponse {
	s.Body = v
	return s
}

type CreateDemoInstanceRequest struct {
	OutboundCallWhitelist *string `json:"OutboundCallWhitelist,omitempty" xml:"OutboundCallWhitelist,omitempty"`
}

func (s CreateDemoInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDemoInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceRequest) SetOutboundCallWhitelist(v string) *CreateDemoInstanceRequest {
	s.OutboundCallWhitelist = &v
	return s
}

type CreateDemoInstanceResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDemoInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDemoInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceResponseBody) SetCode(v string) *CreateDemoInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetData(v string) *CreateDemoInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetHttpStatusCode(v int32) *CreateDemoInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetMessage(v string) *CreateDemoInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetParams(v []*string) *CreateDemoInstanceResponseBody {
	s.Params = v
	return s
}

func (s *CreateDemoInstanceResponseBody) SetRequestId(v string) *CreateDemoInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDemoInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDemoInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDemoInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDemoInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceResponse) SetHeaders(v map[string]*string) *CreateDemoInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDemoInstanceResponse) SetStatusCode(v int32) *CreateDemoInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDemoInstanceResponse) SetBody(v *CreateDemoInstanceResponseBody) *CreateDemoInstanceResponse {
	s.Body = v
	return s
}

type GetCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignRequest) GoString() string {
	return s.String()
}

func (s *GetCampaignRequest) SetCampaignId(v string) *GetCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *GetCampaignRequest) SetInstanceId(v string) *GetCampaignRequest {
	s.InstanceId = &v
	return s
}

type GetCampaignResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBody) SetCode(v string) *GetCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *GetCampaignResponseBody) SetData(v *GetCampaignResponseBodyData) *GetCampaignResponseBody {
	s.Data = v
	return s
}

func (s *GetCampaignResponseBody) SetHttpStatusCode(v int64) *GetCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCampaignResponseBody) SetRequestId(v string) *GetCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCampaignResponseBody) SetSuccess(v bool) *GetCampaignResponseBody {
	s.Success = &v
	return s
}

type GetCampaignResponseBodyData struct {
	ActualEndTime        *int64  `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	ActualStartTime      *int64  `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	CampaignId           *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CasesAborted         *int64  `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	CasesConnected       *int64  `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	CasesUncompleted     *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletedRate        *int64  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	MaxAttemptCount      *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval   *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanedEndTime        *int64  `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	PlanedStartTime      *int64  `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	QueueId              *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName            *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Simulation           *bool   `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	SimulationParameters *string `json:"SimulationParameters,omitempty" xml:"SimulationParameters,omitempty"`
	State                *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyParameters   *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType         *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	TotalCases           *int64  `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
}

func (s GetCampaignResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCampaignResponseBodyData) SetActualEndTime(v int64) *GetCampaignResponseBodyData {
	s.ActualEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetActualStartTime(v int64) *GetCampaignResponseBodyData {
	s.ActualStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCampaignId(v string) *GetCampaignResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesAborted(v int64) *GetCampaignResponseBodyData {
	s.CasesAborted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesConnected(v int64) *GetCampaignResponseBodyData {
	s.CasesConnected = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCasesUncompleted(v int64) *GetCampaignResponseBodyData {
	s.CasesUncompleted = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetCompletedRate(v int64) *GetCampaignResponseBodyData {
	s.CompletedRate = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMaxAttemptCount(v int64) *GetCampaignResponseBodyData {
	s.MaxAttemptCount = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetMinAttemptInterval(v int64) *GetCampaignResponseBodyData {
	s.MinAttemptInterval = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetName(v string) *GetCampaignResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlanedEndTime(v int64) *GetCampaignResponseBodyData {
	s.PlanedEndTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetPlanedStartTime(v int64) *GetCampaignResponseBodyData {
	s.PlanedStartTime = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetQueueId(v string) *GetCampaignResponseBodyData {
	s.QueueId = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetQueueName(v string) *GetCampaignResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetSimulation(v bool) *GetCampaignResponseBodyData {
	s.Simulation = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetSimulationParameters(v string) *GetCampaignResponseBodyData {
	s.SimulationParameters = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetState(v string) *GetCampaignResponseBodyData {
	s.State = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetStrategyParameters(v string) *GetCampaignResponseBodyData {
	s.StrategyParameters = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetStrategyType(v string) *GetCampaignResponseBodyData {
	s.StrategyType = &v
	return s
}

func (s *GetCampaignResponseBodyData) SetTotalCases(v int64) *GetCampaignResponseBodyData {
	s.TotalCases = &v
	return s
}

type GetCampaignResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCampaignResponse) GoString() string {
	return s.String()
}

func (s *GetCampaignResponse) SetHeaders(v map[string]*string) *GetCampaignResponse {
	s.Headers = v
	return s
}

func (s *GetCampaignResponse) SetStatusCode(v int32) *GetCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCampaignResponse) SetBody(v *GetCampaignResponseBody) *GetCampaignResponse {
	s.Body = v
	return s
}

type GetHistoricalCampaignReportRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHistoricalCampaignReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCampaignReportRequest) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportRequest) SetCampaignId(v string) *GetHistoricalCampaignReportRequest {
	s.CampaignId = &v
	return s
}

func (s *GetHistoricalCampaignReportRequest) SetInstanceId(v string) *GetHistoricalCampaignReportRequest {
	s.InstanceId = &v
	return s
}

type GetHistoricalCampaignReportResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetHistoricalCampaignReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHistoricalCampaignReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCampaignReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponseBody) SetCode(v string) *GetHistoricalCampaignReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetData(v *GetHistoricalCampaignReportResponseBodyData) *GetHistoricalCampaignReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetHttpStatusCode(v int32) *GetHistoricalCampaignReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetMessage(v string) *GetHistoricalCampaignReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBody) SetRequestId(v string) *GetHistoricalCampaignReportResponseBody {
	s.RequestId = &v
	return s
}

type GetHistoricalCampaignReportResponseBodyData struct {
	AbandonedRate  *float32 `json:"AbandonedRate,omitempty" xml:"AbandonedRate,omitempty"`
	CallsAbandoned *int64   `json:"CallsAbandoned,omitempty" xml:"CallsAbandoned,omitempty"`
	CallsConnected *int64   `json:"CallsConnected,omitempty" xml:"CallsConnected,omitempty"`
	CallsDialed    *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	ConnectedRate  *float32 `json:"ConnectedRate,omitempty" xml:"ConnectedRate,omitempty"`
	OccupancyRate  *float32 `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
}

func (s GetHistoricalCampaignReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCampaignReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetAbandonedRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.AbandonedRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsAbandoned(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsAbandoned = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsConnected(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsConnected = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetCallsDialed(v int64) *GetHistoricalCampaignReportResponseBodyData {
	s.CallsDialed = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetConnectedRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.ConnectedRate = &v
	return s
}

func (s *GetHistoricalCampaignReportResponseBodyData) SetOccupancyRate(v float32) *GetHistoricalCampaignReportResponseBodyData {
	s.OccupancyRate = &v
	return s
}

type GetHistoricalCampaignReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHistoricalCampaignReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHistoricalCampaignReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHistoricalCampaignReportResponse) GoString() string {
	return s.String()
}

func (s *GetHistoricalCampaignReportResponse) SetHeaders(v map[string]*string) *GetHistoricalCampaignReportResponse {
	s.Headers = v
	return s
}

func (s *GetHistoricalCampaignReportResponse) SetStatusCode(v int32) *GetHistoricalCampaignReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoricalCampaignReportResponse) SetBody(v *GetHistoricalCampaignReportResponseBody) *GetHistoricalCampaignReportResponse {
	s.Body = v
	return s
}

type GetInstanceIdsByAliyunUidV2Request struct {
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
}

func (s GetInstanceIdsByAliyunUidV2Request) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2Request) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2Request) SetAliyunUid(v string) *GetInstanceIdsByAliyunUidV2Request {
	s.AliyunUid = &v
	return s
}

type GetInstanceIdsByAliyunUidV2ResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceIds    []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceIdsByAliyunUidV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2ResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetCode(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetHttpStatusCode(v int64) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetInstanceIds(v []*string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.InstanceIds = v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetMessage(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetRequestId(v string) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2ResponseBody) SetSuccess(v bool) *GetInstanceIdsByAliyunUidV2ResponseBody {
	s.Success = &v
	return s
}

type GetInstanceIdsByAliyunUidV2Response struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceIdsByAliyunUidV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceIdsByAliyunUidV2Response) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIdsByAliyunUidV2Response) GoString() string {
	return s.String()
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetHeaders(v map[string]*string) *GetInstanceIdsByAliyunUidV2Response {
	s.Headers = v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetStatusCode(v int32) *GetInstanceIdsByAliyunUidV2Response {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIdsByAliyunUidV2Response) SetBody(v *GetInstanceIdsByAliyunUidV2ResponseBody) *GetInstanceIdsByAliyunUidV2Response {
	s.Body = v
	return s
}

type GetRealtimeCampaignStatsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueueId    *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
}

func (s GetRealtimeCampaignStatsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeCampaignStatsRequest) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsRequest) SetInstanceId(v string) *GetRealtimeCampaignStatsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealtimeCampaignStatsRequest) SetQueueId(v string) *GetRealtimeCampaignStatsRequest {
	s.QueueId = &v
	return s
}

type GetRealtimeCampaignStatsResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetRealtimeCampaignStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRealtimeCampaignStatsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponseBody) SetCode(v string) *GetRealtimeCampaignStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetData(v *GetRealtimeCampaignStatsResponseBodyData) *GetRealtimeCampaignStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetHttpStatusCode(v int32) *GetRealtimeCampaignStatsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetMessage(v string) *GetRealtimeCampaignStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBody) SetRequestId(v string) *GetRealtimeCampaignStatsResponseBody {
	s.RequestId = &v
	return s
}

type GetRealtimeCampaignStatsResponseBodyData struct {
	BreakingAgents                 *int64 `json:"BreakingAgents,omitempty" xml:"BreakingAgents,omitempty"`
	Caps                           *int64 `json:"Caps,omitempty" xml:"Caps,omitempty"`
	LoggedInAgents                 *int64 `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	OutboundScenarioBreakingAgents *int64 `json:"OutboundScenarioBreakingAgents,omitempty" xml:"OutboundScenarioBreakingAgents,omitempty"`
	OutboundScenarioReadyAgents    *int64 `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundScenarioTalkingAgents  *int64 `json:"OutboundScenarioTalkingAgents,omitempty" xml:"OutboundScenarioTalkingAgents,omitempty"`
	OutboundScenarioWorkingAgents  *int64 `json:"OutboundScenarioWorkingAgents,omitempty" xml:"OutboundScenarioWorkingAgents,omitempty"`
	ReadyAgents                    *int64 `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	TalkingAgents                  *int64 `json:"TalkingAgents,omitempty" xml:"TalkingAgents,omitempty"`
	TotalAgents                    *int64 `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	WorkingAgents                  *int64 `json:"WorkingAgents,omitempty" xml:"WorkingAgents,omitempty"`
}

func (s GetRealtimeCampaignStatsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetBreakingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.BreakingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetCaps(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.Caps = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetLoggedInAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioBreakingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioBreakingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioReadyAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioTalkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioTalkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetOutboundScenarioWorkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.OutboundScenarioWorkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetReadyAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetTalkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.TalkingAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetTotalAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.TotalAgents = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponseBodyData) SetWorkingAgents(v int64) *GetRealtimeCampaignStatsResponseBodyData {
	s.WorkingAgents = &v
	return s
}

type GetRealtimeCampaignStatsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRealtimeCampaignStatsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRealtimeCampaignStatsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRealtimeCampaignStatsResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeCampaignStatsResponse) SetHeaders(v map[string]*string) *GetRealtimeCampaignStatsResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeCampaignStatsResponse) SetStatusCode(v int32) *GetRealtimeCampaignStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeCampaignStatsResponse) SetBody(v *GetRealtimeCampaignStatsResponseBody) *GetRealtimeCampaignStatsResponse {
	s.Body = v
	return s
}

type ImportAdminsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RamIdList  *string `json:"RamIdList,omitempty" xml:"RamIdList,omitempty"`
}

func (s ImportAdminsRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportAdminsRequest) GoString() string {
	return s.String()
}

func (s *ImportAdminsRequest) SetInstanceId(v string) *ImportAdminsRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportAdminsRequest) SetRamIdList(v string) *ImportAdminsRequest {
	s.RamIdList = &v
	return s
}

type ImportAdminsResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ImportAdminsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAdminsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportAdminsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponseBody) SetCode(v string) *ImportAdminsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportAdminsResponseBody) SetData(v []*ImportAdminsResponseBodyData) *ImportAdminsResponseBody {
	s.Data = v
	return s
}

func (s *ImportAdminsResponseBody) SetHttpStatusCode(v int32) *ImportAdminsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportAdminsResponseBody) SetMessage(v string) *ImportAdminsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportAdminsResponseBody) SetRequestId(v string) *ImportAdminsResponseBody {
	s.RequestId = &v
	return s
}

type ImportAdminsResponseBodyData struct {
	Extension  *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RamId      *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	RoleId     *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ImportAdminsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportAdminsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponseBodyData) SetExtension(v string) *ImportAdminsResponseBodyData {
	s.Extension = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetInstanceId(v string) *ImportAdminsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetRamId(v string) *ImportAdminsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetRoleId(v string) *ImportAdminsResponseBodyData {
	s.RoleId = &v
	return s
}

func (s *ImportAdminsResponseBodyData) SetUserId(v string) *ImportAdminsResponseBodyData {
	s.UserId = &v
	return s
}

type ImportAdminsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportAdminsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportAdminsResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportAdminsResponse) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponse) SetHeaders(v map[string]*string) *ImportAdminsResponse {
	s.Headers = v
	return s
}

func (s *ImportAdminsResponse) SetStatusCode(v int32) *ImportAdminsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAdminsResponse) SetBody(v *ImportAdminsResponseBody) *ImportAdminsResponse {
	s.Body = v
	return s
}

type IssueSoftphoneCommandRequest struct {
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s IssueSoftphoneCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s IssueSoftphoneCommandRequest) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandRequest) SetData(v string) *IssueSoftphoneCommandRequest {
	s.Data = &v
	return s
}

type IssueSoftphoneCommandResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IssueSoftphoneCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IssueSoftphoneCommandResponseBody) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandResponseBody) SetCode(v string) *IssueSoftphoneCommandResponseBody {
	s.Code = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetData(v string) *IssueSoftphoneCommandResponseBody {
	s.Data = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetHttpStatusCode(v int32) *IssueSoftphoneCommandResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetMessage(v string) *IssueSoftphoneCommandResponseBody {
	s.Message = &v
	return s
}

func (s *IssueSoftphoneCommandResponseBody) SetRequestId(v string) *IssueSoftphoneCommandResponseBody {
	s.RequestId = &v
	return s
}

type IssueSoftphoneCommandResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IssueSoftphoneCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IssueSoftphoneCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s IssueSoftphoneCommandResponse) GoString() string {
	return s.String()
}

func (s *IssueSoftphoneCommandResponse) SetHeaders(v map[string]*string) *IssueSoftphoneCommandResponse {
	s.Headers = v
	return s
}

func (s *IssueSoftphoneCommandResponse) SetStatusCode(v int32) *IssueSoftphoneCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *IssueSoftphoneCommandResponse) SetBody(v *IssueSoftphoneCommandResponseBody) *IssueSoftphoneCommandResponse {
	s.Body = v
	return s
}

type ListAttemptsRequest struct {
	AgentId    *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AttemptId  *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	Callee     *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller     *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CaseId     *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueueId    *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAttemptsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAttemptsRequest) GoString() string {
	return s.String()
}

func (s *ListAttemptsRequest) SetAgentId(v string) *ListAttemptsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAttemptsRequest) SetAttemptId(v string) *ListAttemptsRequest {
	s.AttemptId = &v
	return s
}

func (s *ListAttemptsRequest) SetCallee(v string) *ListAttemptsRequest {
	s.Callee = &v
	return s
}

func (s *ListAttemptsRequest) SetCaller(v string) *ListAttemptsRequest {
	s.Caller = &v
	return s
}

func (s *ListAttemptsRequest) SetCampaignId(v string) *ListAttemptsRequest {
	s.CampaignId = &v
	return s
}

func (s *ListAttemptsRequest) SetCaseId(v string) *ListAttemptsRequest {
	s.CaseId = &v
	return s
}

func (s *ListAttemptsRequest) SetContactId(v string) *ListAttemptsRequest {
	s.ContactId = &v
	return s
}

func (s *ListAttemptsRequest) SetEndTime(v int64) *ListAttemptsRequest {
	s.EndTime = &v
	return s
}

func (s *ListAttemptsRequest) SetInstanceId(v string) *ListAttemptsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAttemptsRequest) SetPageNumber(v int32) *ListAttemptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAttemptsRequest) SetPageSize(v int32) *ListAttemptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAttemptsRequest) SetQueueId(v string) *ListAttemptsRequest {
	s.QueueId = &v
	return s
}

func (s *ListAttemptsRequest) SetStartTime(v int64) *ListAttemptsRequest {
	s.StartTime = &v
	return s
}

type ListAttemptsResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListAttemptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAttemptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAttemptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBody) SetCode(v string) *ListAttemptsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAttemptsResponseBody) SetData(v *ListAttemptsResponseBodyData) *ListAttemptsResponseBody {
	s.Data = v
	return s
}

func (s *ListAttemptsResponseBody) SetHttpStatusCode(v int32) *ListAttemptsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAttemptsResponseBody) SetMessage(v string) *ListAttemptsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAttemptsResponseBody) SetRequestId(v string) *ListAttemptsResponseBody {
	s.RequestId = &v
	return s
}

type ListAttemptsResponseBodyData struct {
	List       []*ListAttemptsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAttemptsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAttemptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBodyData) SetList(v []*ListAttemptsResponseBodyDataList) *ListAttemptsResponseBodyData {
	s.List = v
	return s
}

func (s *ListAttemptsResponseBodyData) SetPageNumber(v int32) *ListAttemptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAttemptsResponseBodyData) SetPageSize(v int32) *ListAttemptsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAttemptsResponseBodyData) SetTotalCount(v int32) *ListAttemptsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListAttemptsResponseBodyDataList struct {
	AgentEstablishedTime    *int64  `json:"AgentEstablishedTime,omitempty" xml:"AgentEstablishedTime,omitempty"`
	AgentId                 *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentRingDuration       *int64  `json:"AgentRingDuration,omitempty" xml:"AgentRingDuration,omitempty"`
	AssignAgentTime         *int64  `json:"AssignAgentTime,omitempty" xml:"AssignAgentTime,omitempty"`
	AttemptId               *string `json:"AttemptId,omitempty" xml:"AttemptId,omitempty"`
	Callee                  *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller                  *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CampaignId              *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CaseId                  *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	ContactId               *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	CustomerEstablishedTime *int64  `json:"CustomerEstablishedTime,omitempty" xml:"CustomerEstablishedTime,omitempty"`
	CustomerReleasedTime    *int64  `json:"CustomerReleasedTime,omitempty" xml:"CustomerReleasedTime,omitempty"`
	DialDuration            *int64  `json:"DialDuration,omitempty" xml:"DialDuration,omitempty"`
	DialTime                *int64  `json:"DialTime,omitempty" xml:"DialTime,omitempty"`
	EnqueueTime             *int64  `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	EnterIvrTime            *int64  `json:"EnterIvrTime,omitempty" xml:"EnterIvrTime,omitempty"`
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IvrDuration             *int64  `json:"IvrDuration,omitempty" xml:"IvrDuration,omitempty"`
	QueueDuration           *int64  `json:"QueueDuration,omitempty" xml:"QueueDuration,omitempty"`
	QueueId                 *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
}

func (s ListAttemptsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListAttemptsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponseBodyDataList) SetAgentEstablishedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.AgentEstablishedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAgentId(v string) *ListAttemptsResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAgentRingDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.AgentRingDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAssignAgentTime(v int64) *ListAttemptsResponseBodyDataList {
	s.AssignAgentTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetAttemptId(v string) *ListAttemptsResponseBodyDataList {
	s.AttemptId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCallee(v string) *ListAttemptsResponseBodyDataList {
	s.Callee = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCaller(v string) *ListAttemptsResponseBodyDataList {
	s.Caller = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCampaignId(v string) *ListAttemptsResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCaseId(v string) *ListAttemptsResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetContactId(v string) *ListAttemptsResponseBodyDataList {
	s.ContactId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCustomerEstablishedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.CustomerEstablishedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetCustomerReleasedTime(v int64) *ListAttemptsResponseBodyDataList {
	s.CustomerReleasedTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetDialDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.DialDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetDialTime(v int64) *ListAttemptsResponseBodyDataList {
	s.DialTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetEnqueueTime(v int64) *ListAttemptsResponseBodyDataList {
	s.EnqueueTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetEnterIvrTime(v int64) *ListAttemptsResponseBodyDataList {
	s.EnterIvrTime = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetInstanceId(v string) *ListAttemptsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetIvrDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.IvrDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetQueueDuration(v int64) *ListAttemptsResponseBodyDataList {
	s.QueueDuration = &v
	return s
}

func (s *ListAttemptsResponseBodyDataList) SetQueueId(v string) *ListAttemptsResponseBodyDataList {
	s.QueueId = &v
	return s
}

type ListAttemptsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAttemptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAttemptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAttemptsResponse) GoString() string {
	return s.String()
}

func (s *ListAttemptsResponse) SetHeaders(v map[string]*string) *ListAttemptsResponse {
	s.Headers = v
	return s
}

func (s *ListAttemptsResponse) SetStatusCode(v int32) *ListAttemptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAttemptsResponse) SetBody(v *ListAttemptsResponseBody) *ListAttemptsResponse {
	s.Body = v
	return s
}

type ListCampaignTrendingReportRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	EndTime    *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StartTime  *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCampaignTrendingReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportRequest) SetCampaignId(v string) *ListCampaignTrendingReportRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetEndTime(v int64) *ListCampaignTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetInstanceId(v string) *ListCampaignTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCampaignTrendingReportRequest) SetStartTime(v int64) *ListCampaignTrendingReportRequest {
	s.StartTime = &v
	return s
}

type ListCampaignTrendingReportResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListCampaignTrendingReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCampaignTrendingReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignTrendingReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponseBody) SetCode(v string) *ListCampaignTrendingReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetData(v []*ListCampaignTrendingReportResponseBodyData) *ListCampaignTrendingReportResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetHttpStatusCode(v int32) *ListCampaignTrendingReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetMessage(v string) *ListCampaignTrendingReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBody) SetRequestId(v string) *ListCampaignTrendingReportResponseBody {
	s.RequestId = &v
	return s
}

type ListCampaignTrendingReportResponseBodyData struct {
	BreakAgents                    *int64  `json:"BreakAgents,omitempty" xml:"BreakAgents,omitempty"`
	Concurrency                    *int64  `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	Datetime                       *int64  `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	LoggedInAgents                 *int64  `json:"LoggedInAgents,omitempty" xml:"LoggedInAgents,omitempty"`
	OutboundScenarioBreakingAgents *string `json:"OutboundScenarioBreakingAgents,omitempty" xml:"OutboundScenarioBreakingAgents,omitempty"`
	OutboundScenarioReadyAgents    *string `json:"OutboundScenarioReadyAgents,omitempty" xml:"OutboundScenarioReadyAgents,omitempty"`
	OutboundScenarioTalkingAgents  *string `json:"OutboundScenarioTalkingAgents,omitempty" xml:"OutboundScenarioTalkingAgents,omitempty"`
	OutboundScenarioWorkingAgents  *string `json:"OutboundScenarioWorkingAgents,omitempty" xml:"OutboundScenarioWorkingAgents,omitempty"`
	ReadyAgents                    *int64  `json:"ReadyAgents,omitempty" xml:"ReadyAgents,omitempty"`
	TalkAgents                     *int64  `json:"TalkAgents,omitempty" xml:"TalkAgents,omitempty"`
	WorkAgents                     *int64  `json:"WorkAgents,omitempty" xml:"WorkAgents,omitempty"`
}

func (s ListCampaignTrendingReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignTrendingReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponseBodyData) SetBreakAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.BreakAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetConcurrency(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetDatetime(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.Datetime = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetLoggedInAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.LoggedInAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioBreakingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioBreakingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioReadyAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioTalkingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioTalkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetOutboundScenarioWorkingAgents(v string) *ListCampaignTrendingReportResponseBodyData {
	s.OutboundScenarioWorkingAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetReadyAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.ReadyAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetTalkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.TalkAgents = &v
	return s
}

func (s *ListCampaignTrendingReportResponseBodyData) SetWorkAgents(v int64) *ListCampaignTrendingReportResponseBodyData {
	s.WorkAgents = &v
	return s
}

type ListCampaignTrendingReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCampaignTrendingReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCampaignTrendingReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignTrendingReportResponse) GoString() string {
	return s.String()
}

func (s *ListCampaignTrendingReportResponse) SetHeaders(v map[string]*string) *ListCampaignTrendingReportResponse {
	s.Headers = v
	return s
}

func (s *ListCampaignTrendingReportResponse) SetStatusCode(v int32) *ListCampaignTrendingReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCampaignTrendingReportResponse) SetBody(v *ListCampaignTrendingReportResponseBody) *ListCampaignTrendingReportResponse {
	s.Body = v
	return s
}

type ListCampaignsRequest struct {
	ActualStartTimeFrom *string `json:"ActualStartTimeFrom,omitempty" xml:"ActualStartTimeFrom,omitempty"`
	ActualStartTimeTo   *string `json:"ActualStartTimeTo,omitempty" xml:"ActualStartTimeTo,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber          *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlanedStartTimeFrom *string `json:"PlanedStartTimeFrom,omitempty" xml:"PlanedStartTimeFrom,omitempty"`
	PlanedStartTimeTo   *string `json:"PlanedStartTimeTo,omitempty" xml:"PlanedStartTimeTo,omitempty"`
	QueueId             *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	State               *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCampaignsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsRequest) GoString() string {
	return s.String()
}

func (s *ListCampaignsRequest) SetActualStartTimeFrom(v string) *ListCampaignsRequest {
	s.ActualStartTimeFrom = &v
	return s
}

func (s *ListCampaignsRequest) SetActualStartTimeTo(v string) *ListCampaignsRequest {
	s.ActualStartTimeTo = &v
	return s
}

func (s *ListCampaignsRequest) SetInstanceId(v string) *ListCampaignsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCampaignsRequest) SetName(v string) *ListCampaignsRequest {
	s.Name = &v
	return s
}

func (s *ListCampaignsRequest) SetPageNumber(v int64) *ListCampaignsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsRequest) SetPageSize(v int64) *ListCampaignsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsRequest) SetPlanedStartTimeFrom(v string) *ListCampaignsRequest {
	s.PlanedStartTimeFrom = &v
	return s
}

func (s *ListCampaignsRequest) SetPlanedStartTimeTo(v string) *ListCampaignsRequest {
	s.PlanedStartTimeTo = &v
	return s
}

func (s *ListCampaignsRequest) SetQueueId(v string) *ListCampaignsRequest {
	s.QueueId = &v
	return s
}

func (s *ListCampaignsRequest) SetState(v string) *ListCampaignsRequest {
	s.State = &v
	return s
}

type ListCampaignsResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListCampaignsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCampaignsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBody) SetCode(v string) *ListCampaignsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCampaignsResponseBody) SetData(v *ListCampaignsResponseBodyData) *ListCampaignsResponseBody {
	s.Data = v
	return s
}

func (s *ListCampaignsResponseBody) SetHttpStatusCode(v int64) *ListCampaignsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCampaignsResponseBody) SetMessage(v string) *ListCampaignsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCampaignsResponseBody) SetRequestId(v string) *ListCampaignsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCampaignsResponseBody) SetSuccess(v bool) *ListCampaignsResponseBody {
	s.Success = &v
	return s
}

type ListCampaignsResponseBodyData struct {
	List       []*ListCampaignsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int64                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCampaignsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyData) SetList(v []*ListCampaignsResponseBodyDataList) *ListCampaignsResponseBodyData {
	s.List = v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageNumber(v int64) *ListCampaignsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetPageSize(v int64) *ListCampaignsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCampaignsResponseBodyData) SetTotalCount(v int64) *ListCampaignsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCampaignsResponseBodyDataList struct {
	ActualEndTime      *int64  `json:"ActualEndTime,omitempty" xml:"ActualEndTime,omitempty"`
	ActualStartTime    *int64  `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	CampaignId         *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	CasesAborted       *int64  `json:"CasesAborted,omitempty" xml:"CasesAborted,omitempty"`
	CasesConnected     *int64  `json:"CasesConnected,omitempty" xml:"CasesConnected,omitempty"`
	CasesUncompleted   *int64  `json:"CasesUncompleted,omitempty" xml:"CasesUncompleted,omitempty"`
	CompletedRate      *int64  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	MaxAttemptCount    *int64  `json:"MaxAttemptCount,omitempty" xml:"MaxAttemptCount,omitempty"`
	MinAttemptInterval *int64  `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PlanedEndTime      *int64  `json:"PlanedEndTime,omitempty" xml:"PlanedEndTime,omitempty"`
	PlanedStartTime    *int64  `json:"PlanedStartTime,omitempty" xml:"PlanedStartTime,omitempty"`
	QueueId            *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	QueueName          *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Simulation         *bool   `json:"Simulation,omitempty" xml:"Simulation,omitempty"`
	State              *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyParameters *string `json:"StrategyParameters,omitempty" xml:"StrategyParameters,omitempty"`
	StrategyType       *string `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	TotalCases         *int64  `json:"TotalCases,omitempty" xml:"TotalCases,omitempty"`
}

func (s ListCampaignsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponseBodyDataList) SetActualEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetActualStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.ActualStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCampaignId(v string) *ListCampaignsResponseBodyDataList {
	s.CampaignId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesAborted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesAborted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesConnected(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesConnected = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCasesUncompleted(v int64) *ListCampaignsResponseBodyDataList {
	s.CasesUncompleted = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetCompletedRate(v int64) *ListCampaignsResponseBodyDataList {
	s.CompletedRate = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMaxAttemptCount(v int64) *ListCampaignsResponseBodyDataList {
	s.MaxAttemptCount = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetMinAttemptInterval(v int64) *ListCampaignsResponseBodyDataList {
	s.MinAttemptInterval = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetName(v string) *ListCampaignsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedEndTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedEndTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetPlanedStartTime(v int64) *ListCampaignsResponseBodyDataList {
	s.PlanedStartTime = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueId(v string) *ListCampaignsResponseBodyDataList {
	s.QueueId = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetQueueName(v string) *ListCampaignsResponseBodyDataList {
	s.QueueName = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetSimulation(v bool) *ListCampaignsResponseBodyDataList {
	s.Simulation = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetState(v string) *ListCampaignsResponseBodyDataList {
	s.State = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyParameters(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyParameters = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetStrategyType(v string) *ListCampaignsResponseBodyDataList {
	s.StrategyType = &v
	return s
}

func (s *ListCampaignsResponseBodyDataList) SetTotalCases(v int64) *ListCampaignsResponseBodyDataList {
	s.TotalCases = &v
	return s
}

type ListCampaignsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCampaignsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCampaignsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCampaignsResponse) GoString() string {
	return s.String()
}

func (s *ListCampaignsResponse) SetHeaders(v map[string]*string) *ListCampaignsResponse {
	s.Headers = v
	return s
}

func (s *ListCampaignsResponse) SetStatusCode(v int32) *ListCampaignsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCampaignsResponse) SetBody(v *ListCampaignsResponseBody) *ListCampaignsResponse {
	s.Body = v
	return s
}

type ListCasesRequest struct {
	CampaignId  *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber  *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s ListCasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCasesRequest) GoString() string {
	return s.String()
}

func (s *ListCasesRequest) SetCampaignId(v string) *ListCasesRequest {
	s.CampaignId = &v
	return s
}

func (s *ListCasesRequest) SetInstanceId(v string) *ListCasesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCasesRequest) SetPageNumber(v int64) *ListCasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCasesRequest) SetPageSize(v int64) *ListCasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCasesRequest) SetPhoneNumber(v string) *ListCasesRequest {
	s.PhoneNumber = &v
	return s
}

type ListCasesResponseBody struct {
	Code           *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListCasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBody) SetCode(v string) *ListCasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCasesResponseBody) SetData(v *ListCasesResponseBodyData) *ListCasesResponseBody {
	s.Data = v
	return s
}

func (s *ListCasesResponseBody) SetHttpStatusCode(v int64) *ListCasesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCasesResponseBody) SetMessage(v string) *ListCasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCasesResponseBody) SetRequestId(v string) *ListCasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCasesResponseBody) SetSuccess(v bool) *ListCasesResponseBody {
	s.Success = &v
	return s
}

type ListCasesResponseBodyData struct {
	List       []*ListCasesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int64                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyData) SetList(v []*ListCasesResponseBodyDataList) *ListCasesResponseBodyData {
	s.List = v
	return s
}

func (s *ListCasesResponseBodyData) SetPageNumber(v int64) *ListCasesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCasesResponseBodyData) SetPageSize(v int64) *ListCasesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCasesResponseBodyData) SetTotalCount(v int64) *ListCasesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCasesResponseBodyDataList struct {
	AbandonType     *string `json:"AbandonType,omitempty" xml:"AbandonType,omitempty"`
	AttemptCount    *int64  `json:"AttemptCount,omitempty" xml:"AttemptCount,omitempty"`
	CaseId          *string `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CustomVariables *string `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	ExpandInfo      *string `json:"ExpandInfo,omitempty" xml:"ExpandInfo,omitempty"`
	FailureReason   *string `json:"FailureReason,omitempty" xml:"FailureReason,omitempty"`
	PhoneNumber     *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListCasesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListCasesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListCasesResponseBodyDataList) SetAbandonType(v string) *ListCasesResponseBodyDataList {
	s.AbandonType = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetAttemptCount(v int64) *ListCasesResponseBodyDataList {
	s.AttemptCount = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCaseId(v string) *ListCasesResponseBodyDataList {
	s.CaseId = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetCustomVariables(v string) *ListCasesResponseBodyDataList {
	s.CustomVariables = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetExpandInfo(v string) *ListCasesResponseBodyDataList {
	s.ExpandInfo = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetFailureReason(v string) *ListCasesResponseBodyDataList {
	s.FailureReason = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetPhoneNumber(v string) *ListCasesResponseBodyDataList {
	s.PhoneNumber = &v
	return s
}

func (s *ListCasesResponseBodyDataList) SetState(v string) *ListCasesResponseBodyDataList {
	s.State = &v
	return s
}

type ListCasesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCasesResponse) GoString() string {
	return s.String()
}

func (s *ListCasesResponse) SetHeaders(v map[string]*string) *ListCasesResponse {
	s.Headers = v
	return s
}

func (s *ListCasesResponse) SetStatusCode(v int32) *ListCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCasesResponse) SetBody(v *ListCasesResponseBody) *ListCasesResponse {
	s.Body = v
	return s
}

type ListHistoricalAgentSkillGroupReportRequest struct {
	AgentIdList      *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType        *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetAgentIdList(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetEndTime(v int64) *ListHistoricalAgentSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetInstanceId(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetMediaType(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetPageNumber(v int32) *ListHistoricalAgentSkillGroupReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetPageSize(v int32) *ListHistoricalAgentSkillGroupReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetSkillGroupIdList(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetStartTime(v int64) *ListHistoricalAgentSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListHistoricalAgentSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetCode(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetData(v *ListHistoricalAgentSkillGroupReportResponseBodyData) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetMessage(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBody) SetRequestId(v string) *ListHistoricalAgentSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyData struct {
	List       []*ListHistoricalAgentSkillGroupReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNumber *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetList(v []*ListHistoricalAgentSkillGroupReportResponseBodyDataList) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.List = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetPageNumber(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetPageSize(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyData) SetTotalCount(v int32) *ListHistoricalAgentSkillGroupReportResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataList struct {
	AgentId        *string                                                           `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName      *string                                                           `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	Back2Back      *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	DisplayId      *string                                                           `json:"DisplayId,omitempty" xml:"DisplayId,omitempty"`
	Inbound        *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound   `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal       *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal  `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound       *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound  `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall        *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall   `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	SkillGroupId   *string                                                           `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string                                                           `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetAgentId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetAgentName(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetBack2Back(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Back2Back = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetDisplayId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.DisplayId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetInbound(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Inbound = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetInternal(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Internal = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetOutbound(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Outbound = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetOverall(v *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.Overall = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetSkillGroupId(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataList) SetSkillGroupName(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataList {
	s.SkillGroupName = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back struct {
	AgentAnswerRate         *float32 `json:"AgentAnswerRate,omitempty" xml:"AgentAnswerRate,omitempty"`
	AnswerRate              *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	AverageRingTime         *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime         *int64   `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered           *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerHandled    *int64   `json:"CallsCustomerHandled,omitempty" xml:"CallsCustomerHandled,omitempty"`
	CallsDialed             *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerHandleRate      *float32 `json:"CustomerHandleRate,omitempty" xml:"CustomerHandleRate,omitempty"`
	MaxCustomerRingTime     *int64   `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	MaxRingTime             *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime             *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCustomerRingTime   *int64   `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	TotalRingTime           *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime           *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAgentAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AgentAnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageCustomerRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetAverageTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsCustomerHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsCustomerHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetCustomerHandleRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.CustomerHandleRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxCustomerRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalCustomerRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListBack2Back {
	s.TotalTalkTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound struct {
	AverageFirstResponseTime     *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageResponseTime          *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalMessagesSent            *int64   `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	TotalMessagesSentByAgent     *int64   `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	TotalMessagesSentByCustomer  *int64   `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageFirstResponseTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageResponseTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsHold(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetCallsRinged(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetHandleRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.HandleRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSent(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByAgent(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalMessagesSentByCustomer(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInbound {
	s.TotalWorkTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal struct {
	AverageTalkTime *int64 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered   *int64 `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsDialed     *int64 `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHandled    *int64 `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsOffered    *int64 `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsTalk       *int64 `json:"CallsTalk,omitempty" xml:"CallsTalk,omitempty"`
	MaxTalkTime     *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalTalkTime   *int64 `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetAverageTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsHandled(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetCallsTalk(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.CallsTalk = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListInternal {
	s.TotalTalkTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound struct {
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAnswerRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageDialingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageRingTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAnswered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsAttendedTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferIn(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsBlindTransferOut(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsDialed(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsHold(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetCallsRinged(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxDialingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalDialingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalRingTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOutbound {
	s.TotalWorkTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall struct {
	AverageBreakTime               *float32                                                                             `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	AverageHoldTime                *float32                                                                             `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageReadyTime               *float32                                                                             `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageTalkTime                *float32                                                                             `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime                *float32                                                                             `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList            []*ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	FirstCheckInTime               *int64                                                                               `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	LastCheckOutTime               *int64                                                                               `json:"LastCheckOutTime,omitempty" xml:"LastCheckOutTime,omitempty"`
	MaxBreakTime                   *int64                                                                               `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	MaxHoldTime                    *int64                                                                               `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxReadyTime                   *int64                                                                               `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	MaxTalkTime                    *int64                                                                               `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                    *int64                                                                               `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	OccupancyRate                  *float32                                                                             `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	SatisfactionIndex              *float32                                                                             `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate               *float32                                                                             `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered     *int64                                                                               `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded   *int64                                                                               `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalBreakTime                 *int64                                                                               `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	TotalCalls                     *int64                                                                               `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalHoldTime                  *int64                                                                               `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalLoggedInTime              *int64                                                                               `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	TotalOffSiteOnlineTime         *int64                                                                               `json:"TotalOffSiteOnlineTime,omitempty" xml:"TotalOffSiteOnlineTime,omitempty"`
	TotalOfficePhoneOnlineTime     *int64                                                                               `json:"TotalOfficePhoneOnlineTime,omitempty" xml:"TotalOfficePhoneOnlineTime,omitempty"`
	TotalOnSiteOnlineTime          *int64                                                                               `json:"TotalOnSiteOnlineTime,omitempty" xml:"TotalOnSiteOnlineTime,omitempty"`
	TotalOutboundScenarioReadyTime *int64                                                                               `json:"TotalOutboundScenarioReadyTime,omitempty" xml:"TotalOutboundScenarioReadyTime,omitempty"`
	TotalOutboundScenarioTime      *int64                                                                               `json:"TotalOutboundScenarioTime,omitempty" xml:"TotalOutboundScenarioTime,omitempty"`
	TotalReadyTime                 *int64                                                                               `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalTalkTime                  *int64                                                                               `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                  *int64                                                                               `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageBreakTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageHoldTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageReadyTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageTalkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetAverageWorkTime(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetBreakCodeDetailList(v []*ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetFirstCheckInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetLastCheckOutTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.LastCheckOutTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxBreakTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetMaxWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetOccupancyRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionIndex(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionRate(v float32) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysOffered(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetSatisfactionSurveysResponded(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalBreakTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalCalls(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalHoldTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalLoggedInTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOffSiteOnlineTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOffSiteOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOfficePhoneOnlineTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOfficePhoneOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOnSiteOnlineTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOnSiteOnlineTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalOutboundScenarioTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalReadyTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalTalkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall) SetTotalWorkTime(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverall {
	s.TotalWorkTime = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetBreakCode(v string) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetCount(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList) SetDuration(v int64) *ListHistoricalAgentSkillGroupReportResponseBodyDataListOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

type ListHistoricalAgentSkillGroupReportResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHistoricalAgentSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHistoricalAgentSkillGroupReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListHistoricalAgentSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetBody(v *ListHistoricalAgentSkillGroupReportResponseBody) *ListHistoricalAgentSkillGroupReportResponse {
	s.Body = v
	return s
}

type ListIntervalAgentSkillGroupReportRequest struct {
	AgentId      *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Interval     *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	MediaType    *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetAgentId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.AgentId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetEndTime(v int64) *ListIntervalAgentSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetInstanceId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetInterval(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetMediaType(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetSkillGroupId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetStartTime(v int64) *ListIntervalAgentSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBody struct {
	Code           *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListIntervalAgentSkillGroupReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetCode(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetData(v []*ListIntervalAgentSkillGroupReportResponseBodyData) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetHttpStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetMessage(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBody) SetRequestId(v string) *ListIntervalAgentSkillGroupReportResponseBody {
	s.RequestId = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyData struct {
	Back2Back *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back `json:"Back2Back,omitempty" xml:"Back2Back,omitempty" type:"Struct"`
	Inbound   *ListIntervalAgentSkillGroupReportResponseBodyDataInbound   `json:"Inbound,omitempty" xml:"Inbound,omitempty" type:"Struct"`
	Internal  *ListIntervalAgentSkillGroupReportResponseBodyDataInternal  `json:"Internal,omitempty" xml:"Internal,omitempty" type:"Struct"`
	Outbound  *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound  `json:"Outbound,omitempty" xml:"Outbound,omitempty" type:"Struct"`
	Overall   *ListIntervalAgentSkillGroupReportResponseBodyDataOverall   `json:"Overall,omitempty" xml:"Overall,omitempty" type:"Struct"`
	StatsTime *int64                                                      `json:"StatsTime,omitempty" xml:"StatsTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetBack2Back(v *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Back2Back = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetInbound(v *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Inbound = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetInternal(v *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Internal = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetOutbound(v *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Outbound = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetOverall(v *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.Overall = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyData) SetStatsTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyData {
	s.StatsTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back struct {
	AgentAnswerRate         *float32 `json:"AgentAnswerRate,omitempty" xml:"AgentAnswerRate,omitempty"`
	AnswerRate              *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageCustomerRingTime *float32 `json:"AverageCustomerRingTime,omitempty" xml:"AverageCustomerRingTime,omitempty"`
	AverageRingTime         *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime         *int64   `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered           *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsCustomerHandled    *int64   `json:"CallsCustomerHandled,omitempty" xml:"CallsCustomerHandled,omitempty"`
	CallsDialed             *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CustomerHandleRate      *float32 `json:"CustomerHandleRate,omitempty" xml:"CustomerHandleRate,omitempty"`
	MaxCustomerRingTime     *int64   `json:"MaxCustomerRingTime,omitempty" xml:"MaxCustomerRingTime,omitempty"`
	MaxRingTime             *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime             *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalCustomerRingTime   *int64   `json:"TotalCustomerRingTime,omitempty" xml:"TotalCustomerRingTime,omitempty"`
	TotalRingTime           *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime           *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAgentAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AgentAnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageCustomerRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetAverageTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsCustomerHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsCustomerHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetCustomerHandleRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.CustomerHandleRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxCustomerRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalCustomerRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalCustomerRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataBack2Back {
	s.TotalTalkTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataInbound struct {
	AverageFirstResponseTime     *float32 `json:"AverageFirstResponseTime,omitempty" xml:"AverageFirstResponseTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageResponseTime          *float32 `json:"AverageResponseTime,omitempty" xml:"AverageResponseTime,omitempty"`
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsHandled                 *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	CallsOffered                 *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	HandleRate                   *float32 `json:"HandleRate,omitempty" xml:"HandleRate,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalMessagesSent            *int64   `json:"TotalMessagesSent,omitempty" xml:"TotalMessagesSent,omitempty"`
	TotalMessagesSentByAgent     *int64   `json:"TotalMessagesSentByAgent,omitempty" xml:"TotalMessagesSentByAgent,omitempty"`
	TotalMessagesSentByCustomer  *int64   `json:"TotalMessagesSentByCustomer,omitempty" xml:"TotalMessagesSentByCustomer,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageFirstResponseTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageFirstResponseTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageResponseTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageResponseTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsHold(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetCallsRinged(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetHandleRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.HandleRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSent(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSent = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByAgent(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByAgent = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalMessagesSentByCustomer(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalMessagesSentByCustomer = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInbound) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInbound {
	s.TotalWorkTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataInternal struct {
	AverageTalkTime *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	CallsAnswered   *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsDialed     *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHandled    *int64   `json:"CallsHandled,omitempty" xml:"CallsHandled,omitempty"`
	CallsOffered    *int64   `json:"CallsOffered,omitempty" xml:"CallsOffered,omitempty"`
	CallsTalk       *int64   `json:"CallsTalk,omitempty" xml:"CallsTalk,omitempty"`
	MaxTalkTime     *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	TotalTalkTime   *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInternal) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataInternal) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsHandled(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsHandled = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetCallsTalk(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.CallsTalk = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataInternal) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataInternal {
	s.TotalTalkTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOutbound struct {
	AnswerRate                   *float32 `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	AverageDialingTime           *float32 `json:"AverageDialingTime,omitempty" xml:"AverageDialingTime,omitempty"`
	AverageHoldTime              *float32 `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageRingTime              *float32 `json:"AverageRingTime,omitempty" xml:"AverageRingTime,omitempty"`
	AverageTalkTime              *float32 `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime              *float32 `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	CallsAnswered                *int64   `json:"CallsAnswered,omitempty" xml:"CallsAnswered,omitempty"`
	CallsAttendedTransferIn      *int64   `json:"CallsAttendedTransferIn,omitempty" xml:"CallsAttendedTransferIn,omitempty"`
	CallsAttendedTransferOut     *int64   `json:"CallsAttendedTransferOut,omitempty" xml:"CallsAttendedTransferOut,omitempty"`
	CallsBlindTransferIn         *int64   `json:"CallsBlindTransferIn,omitempty" xml:"CallsBlindTransferIn,omitempty"`
	CallsBlindTransferOut        *int64   `json:"CallsBlindTransferOut,omitempty" xml:"CallsBlindTransferOut,omitempty"`
	CallsDialed                  *int64   `json:"CallsDialed,omitempty" xml:"CallsDialed,omitempty"`
	CallsHold                    *int64   `json:"CallsHold,omitempty" xml:"CallsHold,omitempty"`
	CallsRinged                  *int64   `json:"CallsRinged,omitempty" xml:"CallsRinged,omitempty"`
	MaxDialingTime               *int64   `json:"MaxDialingTime,omitempty" xml:"MaxDialingTime,omitempty"`
	MaxHoldTime                  *int64   `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxRingTime                  *int64   `json:"MaxRingTime,omitempty" xml:"MaxRingTime,omitempty"`
	MaxTalkTime                  *int64   `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                  *int64   `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	SatisfactionIndex            *float32 `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate             *float32 `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered   *int64   `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded *int64   `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalDialingTime             *int64   `json:"TotalDialingTime,omitempty" xml:"TotalDialingTime,omitempty"`
	TotalHoldTime                *int64   `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalRingTime                *int64   `json:"TotalRingTime,omitempty" xml:"TotalRingTime,omitempty"`
	TotalTalkTime                *int64   `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                *int64   `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAnswerRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AnswerRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageDialingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageRingTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAnswered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAnswered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsAttendedTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsAttendedTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferIn(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferIn = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsBlindTransferOut(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsBlindTransferOut = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsDialed(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsDialed = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsHold(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsHold = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetCallsRinged(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.CallsRinged = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxDialingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalDialingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalDialingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalRingTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalRingTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOutbound {
	s.TotalWorkTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOverall struct {
	AverageBreakTime               *float32                                                                       `json:"AverageBreakTime,omitempty" xml:"AverageBreakTime,omitempty"`
	AverageHoldTime                *float32                                                                       `json:"AverageHoldTime,omitempty" xml:"AverageHoldTime,omitempty"`
	AverageReadyTime               *float32                                                                       `json:"AverageReadyTime,omitempty" xml:"AverageReadyTime,omitempty"`
	AverageTalkTime                *float32                                                                       `json:"AverageTalkTime,omitempty" xml:"AverageTalkTime,omitempty"`
	AverageWorkTime                *float32                                                                       `json:"AverageWorkTime,omitempty" xml:"AverageWorkTime,omitempty"`
	BreakCodeDetailList            []*ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList `json:"BreakCodeDetailList,omitempty" xml:"BreakCodeDetailList,omitempty" type:"Repeated"`
	FirstCheckInTime               *int64                                                                         `json:"FirstCheckInTime,omitempty" xml:"FirstCheckInTime,omitempty"`
	LastCheckoutTime               *int64                                                                         `json:"LastCheckoutTime,omitempty" xml:"LastCheckoutTime,omitempty"`
	MaxBreakTime                   *int64                                                                         `json:"MaxBreakTime,omitempty" xml:"MaxBreakTime,omitempty"`
	MaxHoldTime                    *int64                                                                         `json:"MaxHoldTime,omitempty" xml:"MaxHoldTime,omitempty"`
	MaxReadyTime                   *int64                                                                         `json:"MaxReadyTime,omitempty" xml:"MaxReadyTime,omitempty"`
	MaxTalkTime                    *int64                                                                         `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
	MaxWorkTime                    *int64                                                                         `json:"MaxWorkTime,omitempty" xml:"MaxWorkTime,omitempty"`
	OccupancyRate                  *float32                                                                       `json:"OccupancyRate,omitempty" xml:"OccupancyRate,omitempty"`
	SatisfactionIndex              *float32                                                                       `json:"SatisfactionIndex,omitempty" xml:"SatisfactionIndex,omitempty"`
	SatisfactionRate               *float32                                                                       `json:"SatisfactionRate,omitempty" xml:"SatisfactionRate,omitempty"`
	SatisfactionSurveysOffered     *int64                                                                         `json:"SatisfactionSurveysOffered,omitempty" xml:"SatisfactionSurveysOffered,omitempty"`
	SatisfactionSurveysResponded   *int64                                                                         `json:"SatisfactionSurveysResponded,omitempty" xml:"SatisfactionSurveysResponded,omitempty"`
	TotalBreakTime                 *int64                                                                         `json:"TotalBreakTime,omitempty" xml:"TotalBreakTime,omitempty"`
	TotalCalls                     *int64                                                                         `json:"TotalCalls,omitempty" xml:"TotalCalls,omitempty"`
	TotalHoldTime                  *int64                                                                         `json:"TotalHoldTime,omitempty" xml:"TotalHoldTime,omitempty"`
	TotalLoggedInTime              *int64                                                                         `json:"TotalLoggedInTime,omitempty" xml:"TotalLoggedInTime,omitempty"`
	TotalOffSiteOnlineTime         *int64                                                                         `json:"TotalOffSiteOnlineTime,omitempty" xml:"TotalOffSiteOnlineTime,omitempty"`
	TotalOfficePhoneOnlineTime     *int64                                                                         `json:"TotalOfficePhoneOnlineTime,omitempty" xml:"TotalOfficePhoneOnlineTime,omitempty"`
	TotalOnSiteOnlineTime          *int64                                                                         `json:"TotalOnSiteOnlineTime,omitempty" xml:"TotalOnSiteOnlineTime,omitempty"`
	TotalOutboundScenarioReadyTime *int64                                                                         `json:"TotalOutboundScenarioReadyTime,omitempty" xml:"TotalOutboundScenarioReadyTime,omitempty"`
	TotalOutboundScenarioTime      *int64                                                                         `json:"TotalOutboundScenarioTime,omitempty" xml:"TotalOutboundScenarioTime,omitempty"`
	TotalReadyTime                 *int64                                                                         `json:"TotalReadyTime,omitempty" xml:"TotalReadyTime,omitempty"`
	TotalTalkTime                  *int64                                                                         `json:"TotalTalkTime,omitempty" xml:"TotalTalkTime,omitempty"`
	TotalWorkTime                  *int64                                                                         `json:"TotalWorkTime,omitempty" xml:"TotalWorkTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverall) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverall) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageBreakTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageHoldTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageReadyTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageTalkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetAverageWorkTime(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.AverageWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetBreakCodeDetailList(v []*ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.BreakCodeDetailList = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetFirstCheckInTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.FirstCheckInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetLastCheckoutTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.LastCheckoutTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxBreakTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetMaxWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.MaxWorkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetOccupancyRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.OccupancyRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionIndex(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionIndex = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionRate(v float32) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionRate = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysOffered(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysOffered = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetSatisfactionSurveysResponded(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.SatisfactionSurveysResponded = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalBreakTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalBreakTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalCalls(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalCalls = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalHoldTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalHoldTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalLoggedInTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalLoggedInTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOffSiteOnlineTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOffSiteOnlineTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOfficePhoneOnlineTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOfficePhoneOnlineTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOnSiteOnlineTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOnSiteOnlineTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOutboundScenarioReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOutboundScenarioReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalOutboundScenarioTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalOutboundScenarioTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalReadyTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalReadyTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalTalkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalTalkTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverall) SetTotalWorkTime(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverall {
	s.TotalWorkTime = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList struct {
	BreakCode *string `json:"BreakCode,omitempty" xml:"BreakCode,omitempty"`
	Count     *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetBreakCode(v string) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.BreakCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetCount(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Count = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList) SetDuration(v int64) *ListIntervalAgentSkillGroupReportResponseBodyDataOverallBreakCodeDetailList {
	s.Duration = &v
	return s
}

type ListIntervalAgentSkillGroupReportResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIntervalAgentSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIntervalAgentSkillGroupReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListIntervalAgentSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetBody(v *ListIntervalAgentSkillGroupReportResponseBody) *ListIntervalAgentSkillGroupReportResponse {
	s.Body = v
	return s
}

type ListMonoRecordingsRequest struct {
	ContactId  *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListMonoRecordingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMonoRecordingsRequest) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsRequest) SetContactId(v string) *ListMonoRecordingsRequest {
	s.ContactId = &v
	return s
}

func (s *ListMonoRecordingsRequest) SetInstanceId(v string) *ListMonoRecordingsRequest {
	s.InstanceId = &v
	return s
}

type ListMonoRecordingsResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListMonoRecordingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMonoRecordingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMonoRecordingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponseBody) SetCode(v string) *ListMonoRecordingsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetData(v []*ListMonoRecordingsResponseBodyData) *ListMonoRecordingsResponseBody {
	s.Data = v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetHttpStatusCode(v int32) *ListMonoRecordingsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetMessage(v string) *ListMonoRecordingsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMonoRecordingsResponseBody) SetRequestId(v string) *ListMonoRecordingsResponseBody {
	s.RequestId = &v
	return s
}

type ListMonoRecordingsResponseBodyData struct {
	AgentId      *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentName    *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	ContactId    *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	RamId        *string `json:"RamId,omitempty" xml:"RamId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListMonoRecordingsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMonoRecordingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponseBodyData) SetAgentId(v string) *ListMonoRecordingsResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetAgentName(v string) *ListMonoRecordingsResponseBodyData {
	s.AgentName = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetContactId(v string) *ListMonoRecordingsResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetDuration(v string) *ListMonoRecordingsResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetFileName(v string) *ListMonoRecordingsResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetFileUrl(v string) *ListMonoRecordingsResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetRamId(v string) *ListMonoRecordingsResponseBodyData {
	s.RamId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetSkillGroupId(v string) *ListMonoRecordingsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *ListMonoRecordingsResponseBodyData) SetStartTime(v string) *ListMonoRecordingsResponseBodyData {
	s.StartTime = &v
	return s
}

type ListMonoRecordingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMonoRecordingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMonoRecordingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMonoRecordingsResponse) GoString() string {
	return s.String()
}

func (s *ListMonoRecordingsResponse) SetHeaders(v map[string]*string) *ListMonoRecordingsResponse {
	s.Headers = v
	return s
}

func (s *ListMonoRecordingsResponse) SetStatusCode(v int32) *ListMonoRecordingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMonoRecordingsResponse) SetBody(v *ListMonoRecordingsResponseBody) *ListMonoRecordingsResponse {
	s.Body = v
	return s
}

type PauseCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PauseCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseCampaignRequest) GoString() string {
	return s.String()
}

func (s *PauseCampaignRequest) SetCampaignId(v string) *PauseCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *PauseCampaignRequest) SetInstanceId(v string) *PauseCampaignRequest {
	s.InstanceId = &v
	return s
}

type PauseCampaignResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *PauseCampaignResponseBody) SetCode(v string) *PauseCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *PauseCampaignResponseBody) SetHttpStatusCode(v string) *PauseCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PauseCampaignResponseBody) SetMessage(v string) *PauseCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *PauseCampaignResponseBody) SetRequestId(v string) *PauseCampaignResponseBody {
	s.RequestId = &v
	return s
}

type PauseCampaignResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PauseCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PauseCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseCampaignResponse) GoString() string {
	return s.String()
}

func (s *PauseCampaignResponse) SetHeaders(v map[string]*string) *PauseCampaignResponse {
	s.Headers = v
	return s
}

func (s *PauseCampaignResponse) SetStatusCode(v int32) *PauseCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseCampaignResponse) SetBody(v *PauseCampaignResponseBody) *PauseCampaignResponse {
	s.Body = v
	return s
}

type ReplaceMigrationAvailableNumbersRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OperatorUid  *int64  `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	V1Numbers    *string `json:"V1Numbers,omitempty" xml:"V1Numbers,omitempty"`
	V2Numbers    *string `json:"V2Numbers,omitempty" xml:"V2Numbers,omitempty"`
}

func (s ReplaceMigrationAvailableNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersRequest) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetInstanceId(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetOperatorName(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.OperatorName = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetOperatorUid(v int64) *ReplaceMigrationAvailableNumbersRequest {
	s.OperatorUid = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetV1Numbers(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.V1Numbers = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersRequest) SetV2Numbers(v string) *ReplaceMigrationAvailableNumbersRequest {
	s.V2Numbers = &v
	return s
}

type ReplaceMigrationAvailableNumbersResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceMigrationAvailableNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetCode(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetHttpStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetMessage(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponseBody) SetRequestId(v string) *ReplaceMigrationAvailableNumbersResponseBody {
	s.RequestId = &v
	return s
}

type ReplaceMigrationAvailableNumbersResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplaceMigrationAvailableNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceMigrationAvailableNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceMigrationAvailableNumbersResponse) GoString() string {
	return s.String()
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetHeaders(v map[string]*string) *ReplaceMigrationAvailableNumbersResponse {
	s.Headers = v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetStatusCode(v int32) *ReplaceMigrationAvailableNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceMigrationAvailableNumbersResponse) SetBody(v *ReplaceMigrationAvailableNumbersResponseBody) *ReplaceMigrationAvailableNumbersResponse {
	s.Body = v
	return s
}

type ResumeCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ResumeCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeCampaignRequest) GoString() string {
	return s.String()
}

func (s *ResumeCampaignRequest) SetCampaignId(v string) *ResumeCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *ResumeCampaignRequest) SetInstanceId(v string) *ResumeCampaignRequest {
	s.InstanceId = &v
	return s
}

type ResumeCampaignResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeCampaignResponseBody) SetCode(v string) *ResumeCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetHttpStatusCode(v string) *ResumeCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetMessage(v string) *ResumeCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetRequestId(v string) *ResumeCampaignResponseBody {
	s.RequestId = &v
	return s
}

type ResumeCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeCampaignResponse) GoString() string {
	return s.String()
}

func (s *ResumeCampaignResponse) SetHeaders(v map[string]*string) *ResumeCampaignResponse {
	s.Headers = v
	return s
}

func (s *ResumeCampaignResponse) SetStatusCode(v int32) *ResumeCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeCampaignResponse) SetBody(v *ResumeCampaignResponseBody) *ResumeCampaignResponse {
	s.Body = v
	return s
}

type SaveRTCStatsV2Request struct {
	CallId         *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	GeneralInfo    *string `json:"GeneralInfo,omitempty" xml:"GeneralInfo,omitempty"`
	GoogAddress    *string `json:"GoogAddress,omitempty" xml:"GoogAddress,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ReceiverReport *string `json:"ReceiverReport,omitempty" xml:"ReceiverReport,omitempty"`
	SenderReport   *string `json:"SenderReport,omitempty" xml:"SenderReport,omitempty"`
}

func (s SaveRTCStatsV2Request) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2Request) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Request) SetCallId(v string) *SaveRTCStatsV2Request {
	s.CallId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGeneralInfo(v string) *SaveRTCStatsV2Request {
	s.GeneralInfo = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetGoogAddress(v string) *SaveRTCStatsV2Request {
	s.GoogAddress = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetInstanceId(v string) *SaveRTCStatsV2Request {
	s.InstanceId = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetReceiverReport(v string) *SaveRTCStatsV2Request {
	s.ReceiverReport = &v
	return s
}

func (s *SaveRTCStatsV2Request) SetSenderReport(v string) *SaveRTCStatsV2Request {
	s.SenderReport = &v
	return s
}

type SaveRTCStatsV2ResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RowCount       *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveRTCStatsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2ResponseBody) SetCode(v string) *SaveRTCStatsV2ResponseBody {
	s.Code = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetHttpStatusCode(v int64) *SaveRTCStatsV2ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetMessage(v string) *SaveRTCStatsV2ResponseBody {
	s.Message = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRequestId(v string) *SaveRTCStatsV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetRowCount(v int64) *SaveRTCStatsV2ResponseBody {
	s.RowCount = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetSuccess(v bool) *SaveRTCStatsV2ResponseBody {
	s.Success = &v
	return s
}

func (s *SaveRTCStatsV2ResponseBody) SetTimeStamp(v int64) *SaveRTCStatsV2ResponseBody {
	s.TimeStamp = &v
	return s
}

type SaveRTCStatsV2Response struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveRTCStatsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveRTCStatsV2Response) String() string {
	return tea.Prettify(s)
}

func (s SaveRTCStatsV2Response) GoString() string {
	return s.String()
}

func (s *SaveRTCStatsV2Response) SetHeaders(v map[string]*string) *SaveRTCStatsV2Response {
	s.Headers = v
	return s
}

func (s *SaveRTCStatsV2Response) SetStatusCode(v int32) *SaveRTCStatsV2Response {
	s.StatusCode = &v
	return s
}

func (s *SaveRTCStatsV2Response) SetBody(v *SaveRTCStatsV2ResponseBody) *SaveRTCStatsV2Response {
	s.Body = v
	return s
}

type SaveTerminalLogRequest struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CallId          *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Content         *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DataType        *int32  `json:"DataType,omitempty" xml:"DataType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MethodName      *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UniqueRequestId *string `json:"UniqueRequestId,omitempty" xml:"UniqueRequestId,omitempty"`
}

func (s SaveTerminalLogRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogRequest) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogRequest) SetAppName(v string) *SaveTerminalLogRequest {
	s.AppName = &v
	return s
}

func (s *SaveTerminalLogRequest) SetCallId(v string) *SaveTerminalLogRequest {
	s.CallId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetContent(v string) *SaveTerminalLogRequest {
	s.Content = &v
	return s
}

func (s *SaveTerminalLogRequest) SetDataType(v int32) *SaveTerminalLogRequest {
	s.DataType = &v
	return s
}

func (s *SaveTerminalLogRequest) SetInstanceId(v string) *SaveTerminalLogRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetJobId(v string) *SaveTerminalLogRequest {
	s.JobId = &v
	return s
}

func (s *SaveTerminalLogRequest) SetMethodName(v string) *SaveTerminalLogRequest {
	s.MethodName = &v
	return s
}

func (s *SaveTerminalLogRequest) SetStatus(v string) *SaveTerminalLogRequest {
	s.Status = &v
	return s
}

func (s *SaveTerminalLogRequest) SetUniqueRequestId(v string) *SaveTerminalLogRequest {
	s.UniqueRequestId = &v
	return s
}

type SaveTerminalLogResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveTerminalLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponseBody) SetCode(v string) *SaveTerminalLogResponseBody {
	s.Code = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetHttpStatusCode(v int64) *SaveTerminalLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetMessage(v string) *SaveTerminalLogResponseBody {
	s.Message = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetRequestId(v string) *SaveTerminalLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetSuccess(v bool) *SaveTerminalLogResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTerminalLogResponseBody) SetTimeStamp(v int64) *SaveTerminalLogResponseBody {
	s.TimeStamp = &v
	return s
}

type SaveTerminalLogResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveTerminalLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveTerminalLogResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveTerminalLogResponse) GoString() string {
	return s.String()
}

func (s *SaveTerminalLogResponse) SetHeaders(v map[string]*string) *SaveTerminalLogResponse {
	s.Headers = v
	return s
}

func (s *SaveTerminalLogResponse) SetStatusCode(v int32) *SaveTerminalLogResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveTerminalLogResponse) SetBody(v *SaveTerminalLogResponseBody) *SaveTerminalLogResponse {
	s.Body = v
	return s
}

type SaveWebRtcInfoRequest struct {
	CallId      *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SaveWebRtcInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoRequest) SetCallId(v string) *SaveWebRtcInfoRequest {
	s.CallId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContent(v string) *SaveWebRtcInfoRequest {
	s.Content = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetContentType(v string) *SaveWebRtcInfoRequest {
	s.ContentType = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetInstanceId(v string) *SaveWebRtcInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveWebRtcInfoRequest) SetJobId(v string) *SaveWebRtcInfoRequest {
	s.JobId = &v
	return s
}

type SaveWebRtcInfoResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RowCount       *int64  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TimeStamp      *int64  `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s SaveWebRtcInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponseBody) SetCode(v string) *SaveWebRtcInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetHttpStatusCode(v int64) *SaveWebRtcInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetMessage(v string) *SaveWebRtcInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRequestId(v string) *SaveWebRtcInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetRowCount(v int64) *SaveWebRtcInfoResponseBody {
	s.RowCount = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetSuccess(v bool) *SaveWebRtcInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SaveWebRtcInfoResponseBody) SetTimeStamp(v int64) *SaveWebRtcInfoResponseBody {
	s.TimeStamp = &v
	return s
}

type SaveWebRtcInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveWebRtcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveWebRtcInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWebRtcInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponse) SetHeaders(v map[string]*string) *SaveWebRtcInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveWebRtcInfoResponse) SetStatusCode(v int32) *SaveWebRtcInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWebRtcInfoResponse) SetBody(v *SaveWebRtcInfoResponseBody) *SaveWebRtcInfoResponse {
	s.Body = v
	return s
}

type SubmitCampaignRequest struct {
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SubmitCampaignRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitCampaignRequest) GoString() string {
	return s.String()
}

func (s *SubmitCampaignRequest) SetCampaignId(v string) *SubmitCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *SubmitCampaignRequest) SetInstanceId(v string) *SubmitCampaignRequest {
	s.InstanceId = &v
	return s
}

type SubmitCampaignResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitCampaignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCampaignResponseBody) SetCode(v string) *SubmitCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetHttpStatusCode(v string) *SubmitCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetMessage(v string) *SubmitCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetRequestId(v string) *SubmitCampaignResponseBody {
	s.RequestId = &v
	return s
}

type SubmitCampaignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitCampaignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitCampaignResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitCampaignResponse) GoString() string {
	return s.String()
}

func (s *SubmitCampaignResponse) SetHeaders(v map[string]*string) *SubmitCampaignResponse {
	s.Headers = v
	return s
}

func (s *SubmitCampaignResponse) SetStatusCode(v int32) *SubmitCampaignResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitCampaignResponse) SetBody(v *SubmitCampaignResponseBody) *SubmitCampaignResponse {
	s.Body = v
	return s
}

type UnregisterDeviceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnregisterDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnregisterDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceRequest) SetInstanceId(v string) *UnregisterDeviceRequest {
	s.InstanceId = &v
	return s
}

func (s *UnregisterDeviceRequest) SetUserId(v string) *UnregisterDeviceRequest {
	s.UserId = &v
	return s
}

type UnregisterDeviceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnregisterDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnregisterDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceResponseBody) SetCode(v string) *UnregisterDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetHttpStatusCode(v int32) *UnregisterDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetMessage(v string) *UnregisterDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *UnregisterDeviceResponseBody) SetRequestId(v string) *UnregisterDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnregisterDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnregisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnregisterDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnregisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceResponse) SetHeaders(v map[string]*string) *UnregisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnregisterDeviceResponse) SetStatusCode(v int32) *UnregisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnregisterDeviceResponse) SetBody(v *UnregisterDeviceResponseBody) *UnregisterDeviceResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudcallcenter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AbortCampaignWithOptions(request *AbortCampaignRequest, runtime *util.RuntimeOptions) (_result *AbortCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortCampaign(request *AbortCampaignRequest) (_result *AbortCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortCampaignResponse{}
	_body, _err := client.AbortCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortCasesWithOptions(tmpReq *AbortCasesRequest, runtime *util.RuntimeOptions) (_result *AbortCasesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AbortCasesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PhoneNumberList)) {
		request.PhoneNumberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PhoneNumberList, tea.String("PhoneNumberList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberListShrink)) {
		query["PhoneNumberList"] = request.PhoneNumberListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortCases"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortCases(request *AbortCasesRequest) (_result *AbortCasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbortCasesResponse{}
	_body, _err := client.AbortCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDemoInstanceExistsWithOptions(runtime *util.RuntimeOptions) (_result *CheckDemoInstanceExistsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckDemoInstanceExists"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckDemoInstanceExistsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDemoInstanceExists() (_result *CheckDemoInstanceExistsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDemoInstanceExistsResponse{}
	_body, _err := client.CheckDemoInstanceExistsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckMQRoleAssumptionAuthorityWithOptions(runtime *util.RuntimeOptions) (_result *CheckMQRoleAssumptionAuthorityResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("CheckMQRoleAssumptionAuthority"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckMQRoleAssumptionAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckMQRoleAssumptionAuthority() (_result *CheckMQRoleAssumptionAuthorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckMQRoleAssumptionAuthorityResponse{}
	_body, _err := client.CheckMQRoleAssumptionAuthorityWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCampaignWithOptions(tmpReq *CreateCampaignRequest, runtime *util.RuntimeOptions) (_result *CreateCampaignResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCampaignShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CaseList)) {
		request.CaseListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CaseList, tea.String("CaseList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallableTime)) {
		query["CallableTime"] = request.CallableTime
	}

	if !tea.BoolValue(util.IsUnset(request.CaseFileKey)) {
		query["CaseFileKey"] = request.CaseFileKey
	}

	if !tea.BoolValue(util.IsUnset(request.CaseListShrink)) {
		query["CaseList"] = request.CaseListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ContactFlowId)) {
		query["ContactFlowId"] = request.ContactFlowId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutingUntilTimeout)) {
		query["ExecutingUntilTimeout"] = request.ExecutingUntilTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxAttemptCount)) {
		query["MaxAttemptCount"] = request.MaxAttemptCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinAttemptInterval)) {
		query["MinAttemptInterval"] = request.MinAttemptInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.QueueId)) {
		query["QueueId"] = request.QueueId
	}

	if !tea.BoolValue(util.IsUnset(request.Simulation)) {
		query["Simulation"] = request.Simulation
	}

	if !tea.BoolValue(util.IsUnset(request.SimulationParameters)) {
		query["SimulationParameters"] = request.SimulationParameters
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StrategyParameters)) {
		query["StrategyParameters"] = request.StrategyParameters
	}

	if !tea.BoolValue(util.IsUnset(request.StrategyType)) {
		query["StrategyType"] = request.StrategyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCampaign(request *CreateCampaignRequest) (_result *CreateCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCampaignResponse{}
	_body, _err := client.CreateCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCorpNumberGroupWithOptions(request *CreateCorpNumberGroupRequest, runtime *util.RuntimeOptions) (_result *CreateCorpNumberGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCorpNumberGroup"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCorpNumberGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCorpNumberGroup(request *CreateCorpNumberGroupRequest) (_result *CreateCorpNumberGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCorpNumberGroupResponse{}
	_body, _err := client.CreateCorpNumberGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDemoInstanceWithOptions(request *CreateDemoInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDemoInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OutboundCallWhitelist)) {
		query["OutboundCallWhitelist"] = request.OutboundCallWhitelist
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDemoInstance"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDemoInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDemoInstance(request *CreateDemoInstanceRequest) (_result *CreateDemoInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDemoInstanceResponse{}
	_body, _err := client.CreateDemoInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCampaignWithOptions(request *GetCampaignRequest, runtime *util.RuntimeOptions) (_result *GetCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCampaign(request *GetCampaignRequest) (_result *GetCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCampaignResponse{}
	_body, _err := client.GetCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHistoricalCampaignReportWithOptions(request *GetHistoricalCampaignReportRequest, runtime *util.RuntimeOptions) (_result *GetHistoricalCampaignReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHistoricalCampaignReport"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHistoricalCampaignReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHistoricalCampaignReport(request *GetHistoricalCampaignReportRequest) (_result *GetHistoricalCampaignReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHistoricalCampaignReportResponse{}
	_body, _err := client.GetHistoricalCampaignReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceIdsByAliyunUidV2WithOptions(request *GetInstanceIdsByAliyunUidV2Request, runtime *util.RuntimeOptions) (_result *GetInstanceIdsByAliyunUidV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliyunUid)) {
		query["AliyunUid"] = request.AliyunUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceIdsByAliyunUidV2"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceIdsByAliyunUidV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceIdsByAliyunUidV2(request *GetInstanceIdsByAliyunUidV2Request) (_result *GetInstanceIdsByAliyunUidV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceIdsByAliyunUidV2Response{}
	_body, _err := client.GetInstanceIdsByAliyunUidV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRealtimeCampaignStatsWithOptions(request *GetRealtimeCampaignStatsRequest, runtime *util.RuntimeOptions) (_result *GetRealtimeCampaignStatsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRealtimeCampaignStats"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRealtimeCampaignStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRealtimeCampaignStats(request *GetRealtimeCampaignStatsRequest) (_result *GetRealtimeCampaignStatsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRealtimeCampaignStatsResponse{}
	_body, _err := client.GetRealtimeCampaignStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportAdminsWithOptions(request *ImportAdminsRequest, runtime *util.RuntimeOptions) (_result *ImportAdminsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RamIdList)) {
		query["RamIdList"] = request.RamIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportAdmins"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportAdminsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportAdmins(request *ImportAdminsRequest) (_result *ImportAdminsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportAdminsResponse{}
	_body, _err := client.ImportAdminsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IssueSoftphoneCommandWithOptions(request *IssueSoftphoneCommandRequest, runtime *util.RuntimeOptions) (_result *IssueSoftphoneCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IssueSoftphoneCommand"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IssueSoftphoneCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IssueSoftphoneCommand(request *IssueSoftphoneCommandRequest) (_result *IssueSoftphoneCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IssueSoftphoneCommandResponse{}
	_body, _err := client.IssueSoftphoneCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAttemptsWithOptions(request *ListAttemptsRequest, runtime *util.RuntimeOptions) (_result *ListAttemptsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAttempts"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAttemptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAttempts(request *ListAttemptsRequest) (_result *ListAttemptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAttemptsResponse{}
	_body, _err := client.ListAttemptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCampaignTrendingReportWithOptions(request *ListCampaignTrendingReportRequest, runtime *util.RuntimeOptions) (_result *ListCampaignTrendingReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCampaignTrendingReport"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCampaignTrendingReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCampaignTrendingReport(request *ListCampaignTrendingReportRequest) (_result *ListCampaignTrendingReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCampaignTrendingReportResponse{}
	_body, _err := client.ListCampaignTrendingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCampaignsWithOptions(request *ListCampaignsRequest, runtime *util.RuntimeOptions) (_result *ListCampaignsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualStartTimeFrom)) {
		query["ActualStartTimeFrom"] = request.ActualStartTimeFrom
	}

	if !tea.BoolValue(util.IsUnset(request.ActualStartTimeTo)) {
		query["ActualStartTimeTo"] = request.ActualStartTimeTo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlanedStartTimeFrom)) {
		query["PlanedStartTimeFrom"] = request.PlanedStartTimeFrom
	}

	if !tea.BoolValue(util.IsUnset(request.PlanedStartTimeTo)) {
		query["PlanedStartTimeTo"] = request.PlanedStartTimeTo
	}

	if !tea.BoolValue(util.IsUnset(request.QueueId)) {
		query["QueueId"] = request.QueueId
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["State"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCampaigns"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCampaignsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCampaigns(request *ListCampaignsRequest) (_result *ListCampaignsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCampaignsResponse{}
	_body, _err := client.ListCampaignsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCasesWithOptions(request *ListCasesRequest, runtime *util.RuntimeOptions) (_result *ListCasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCases"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCases(request *ListCasesRequest) (_result *ListCasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCasesResponse{}
	_body, _err := client.ListCasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHistoricalAgentSkillGroupReportWithOptions(request *ListHistoricalAgentSkillGroupReportRequest, runtime *util.RuntimeOptions) (_result *ListHistoricalAgentSkillGroupReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupIdList)) {
		query["SkillGroupIdList"] = request.SkillGroupIdList
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentIdList)) {
		body["AgentIdList"] = request.AgentIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHistoricalAgentSkillGroupReport"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHistoricalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHistoricalAgentSkillGroupReport(request *ListHistoricalAgentSkillGroupReportRequest) (_result *ListHistoricalAgentSkillGroupReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHistoricalAgentSkillGroupReportResponse{}
	_body, _err := client.ListHistoricalAgentSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntervalAgentSkillGroupReportWithOptions(request *ListIntervalAgentSkillGroupReportRequest, runtime *util.RuntimeOptions) (_result *ListIntervalAgentSkillGroupReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntervalAgentSkillGroupReport"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntervalAgentSkillGroupReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntervalAgentSkillGroupReport(request *ListIntervalAgentSkillGroupReportRequest) (_result *ListIntervalAgentSkillGroupReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervalAgentSkillGroupReportResponse{}
	_body, _err := client.ListIntervalAgentSkillGroupReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMonoRecordingsWithOptions(request *ListMonoRecordingsRequest, runtime *util.RuntimeOptions) (_result *ListMonoRecordingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContactId)) {
		query["ContactId"] = request.ContactId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMonoRecordings"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMonoRecordingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMonoRecordings(request *ListMonoRecordingsRequest) (_result *ListMonoRecordingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMonoRecordingsResponse{}
	_body, _err := client.ListMonoRecordingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PauseCampaignWithOptions(request *PauseCampaignRequest, runtime *util.RuntimeOptions) (_result *PauseCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseCampaign(request *PauseCampaignRequest) (_result *PauseCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseCampaignResponse{}
	_body, _err := client.PauseCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceMigrationAvailableNumbersWithOptions(request *ReplaceMigrationAvailableNumbersRequest, runtime *util.RuntimeOptions) (_result *ReplaceMigrationAvailableNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorName)) {
		query["OperatorName"] = request.OperatorName
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorUid)) {
		query["OperatorUid"] = request.OperatorUid
	}

	if !tea.BoolValue(util.IsUnset(request.V1Numbers)) {
		query["V1Numbers"] = request.V1Numbers
	}

	if !tea.BoolValue(util.IsUnset(request.V2Numbers)) {
		query["V2Numbers"] = request.V2Numbers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceMigrationAvailableNumbers"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceMigrationAvailableNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceMigrationAvailableNumbers(request *ReplaceMigrationAvailableNumbersRequest) (_result *ReplaceMigrationAvailableNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReplaceMigrationAvailableNumbersResponse{}
	_body, _err := client.ReplaceMigrationAvailableNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeCampaignWithOptions(request *ResumeCampaignRequest, runtime *util.RuntimeOptions) (_result *ResumeCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeCampaign(request *ResumeCampaignRequest) (_result *ResumeCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeCampaignResponse{}
	_body, _err := client.ResumeCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveRTCStatsV2WithOptions(request *SaveRTCStatsV2Request, runtime *util.RuntimeOptions) (_result *SaveRTCStatsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.GeneralInfo)) {
		query["GeneralInfo"] = request.GeneralInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GoogAddress)) {
		query["GoogAddress"] = request.GoogAddress
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ReceiverReport)) {
		query["ReceiverReport"] = request.ReceiverReport
	}

	if !tea.BoolValue(util.IsUnset(request.SenderReport)) {
		query["SenderReport"] = request.SenderReport
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveRTCStatsV2"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveRTCStatsV2(request *SaveRTCStatsV2Request) (_result *SaveRTCStatsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveRTCStatsV2Response{}
	_body, _err := client.SaveRTCStatsV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveTerminalLogWithOptions(request *SaveTerminalLogRequest, runtime *util.RuntimeOptions) (_result *SaveTerminalLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MethodName)) {
		query["MethodName"] = request.MethodName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueRequestId)) {
		query["UniqueRequestId"] = request.UniqueRequestId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveTerminalLog"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveTerminalLog(request *SaveTerminalLogRequest) (_result *SaveTerminalLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveTerminalLogResponse{}
	_body, _err := client.SaveTerminalLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveWebRtcInfoWithOptions(request *SaveWebRtcInfoRequest, runtime *util.RuntimeOptions) (_result *SaveWebRtcInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		query["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveWebRtcInfo"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveWebRtcInfo(request *SaveWebRtcInfoRequest) (_result *SaveWebRtcInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveWebRtcInfoResponse{}
	_body, _err := client.SaveWebRtcInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitCampaignWithOptions(request *SubmitCampaignRequest, runtime *util.RuntimeOptions) (_result *SubmitCampaignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CampaignId)) {
		query["CampaignId"] = request.CampaignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitCampaign"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitCampaign(request *SubmitCampaignRequest) (_result *SubmitCampaignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitCampaignResponse{}
	_body, _err := client.SubmitCampaignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnregisterDeviceWithOptions(request *UnregisterDeviceRequest, runtime *util.RuntimeOptions) (_result *UnregisterDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnregisterDevice"),
		Version:     tea.String("2020-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnregisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnregisterDevice(request *UnregisterDeviceRequest) (_result *UnregisterDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnregisterDeviceResponse{}
	_body, _err := client.UnregisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
