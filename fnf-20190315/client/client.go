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

type CreateFlowRequest struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetRequestId(v string) *CreateFlowRequest {
	s.RequestId = &v
	return s
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetDefinition(v string) *CreateFlowRequest {
	s.Definition = &v
	return s
}

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetType(v string) *CreateFlowRequest {
	s.Type = &v
	return s
}

func (s *CreateFlowRequest) SetRoleArn(v string) *CreateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowRequest) SetExternalStorageLocation(v string) *CreateFlowRequest {
	s.ExternalStorageLocation = &v
	return s
}

type CreateFlowResponseBody struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime             *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	LastModifiedTime        *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetType(v string) *CreateFlowResponseBody {
	s.Type = &v
	return s
}

func (s *CreateFlowResponseBody) SetDescription(v string) *CreateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *CreateFlowResponseBody) SetCreatedTime(v string) *CreateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponseBody) SetDefinition(v string) *CreateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *CreateFlowResponseBody) SetLastModifiedTime(v string) *CreateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetId(v string) *CreateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowResponseBody) SetExternalStorageLocation(v string) *CreateFlowResponseBody {
	s.ExternalStorageLocation = &v
	return s
}

func (s *CreateFlowResponseBody) SetRoleArn(v string) *CreateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowResponseBody) SetName(v string) *CreateFlowResponseBody {
	s.Name = &v
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

type CreateScheduleRequest struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName       *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ScheduleName   *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Payload        *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Enable         *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s CreateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleRequest) SetRequestId(v string) *CreateScheduleRequest {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleRequest) SetFlowName(v string) *CreateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *CreateScheduleRequest) SetScheduleName(v string) *CreateScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *CreateScheduleRequest) SetDescription(v string) *CreateScheduleRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduleRequest) SetPayload(v string) *CreateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *CreateScheduleRequest) SetCronExpression(v string) *CreateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleRequest) SetEnable(v bool) *CreateScheduleRequest {
	s.Enable = &v
	return s
}

type CreateScheduleResponseBody struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Payload          *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	CronExpression   *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	ScheduleId       *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s CreateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBody) SetDescription(v string) *CreateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *CreateScheduleResponseBody) SetCreatedTime(v string) *CreateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetRequestId(v string) *CreateScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleResponseBody) SetLastModifiedTime(v string) *CreateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetEnable(v bool) *CreateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *CreateScheduleResponseBody) SetPayload(v string) *CreateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *CreateScheduleResponseBody) SetCronExpression(v string) *CreateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleResponseBody) SetScheduleId(v string) *CreateScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *CreateScheduleResponseBody) SetScheduleName(v string) *CreateScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

type CreateScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponse) SetHeaders(v map[string]*string) *CreateScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleResponse) SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse {
	s.Body = v
	return s
}

type DeleteFlowRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetRequestId(v string) *DeleteFlowRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowRequest) SetName(v string) *DeleteFlowRequest {
	s.Name = &v
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

type DeleteScheduleRequest struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName     *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DeleteScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduleRequest) SetRequestId(v string) *DeleteScheduleRequest {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduleRequest) SetFlowName(v string) *DeleteScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *DeleteScheduleRequest) SetScheduleName(v string) *DeleteScheduleRequest {
	s.ScheduleName = &v
	return s
}

type DeleteScheduleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponseBody) SetRequestId(v string) *DeleteScheduleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleResponse) SetHeaders(v map[string]*string) *DeleteScheduleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleResponse) SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse {
	s.Body = v
	return s
}

type DescribeExecutionRequest struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName        *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ExecutionName   *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	WaitTimeSeconds *int32  `json:"WaitTimeSeconds,omitempty" xml:"WaitTimeSeconds,omitempty"`
}

func (s DescribeExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutionRequest) SetRequestId(v string) *DescribeExecutionRequest {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutionRequest) SetFlowName(v string) *DescribeExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionRequest) SetExecutionName(v string) *DescribeExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *DescribeExecutionRequest) SetWaitTimeSeconds(v int32) *DescribeExecutionRequest {
	s.WaitTimeSeconds = &v
	return s
}

type DescribeExecutionResponseBody struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Input             *string `json:"Input,omitempty" xml:"Input,omitempty"`
	StoppedTime       *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName          *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Output            *string `json:"Output,omitempty" xml:"Output,omitempty"`
	ExternalOutputUri *string `json:"ExternalOutputUri,omitempty" xml:"ExternalOutputUri,omitempty"`
	StartedTime       *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	ExternalInputUri  *string `json:"ExternalInputUri,omitempty" xml:"ExternalInputUri,omitempty"`
	FlowDefinition    *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponseBody) SetStatus(v string) *DescribeExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetInput(v string) *DescribeExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStoppedTime(v string) *DescribeExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetRequestId(v string) *DescribeExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetFlowName(v string) *DescribeExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetOutput(v string) *DescribeExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetExternalOutputUri(v string) *DescribeExecutionResponseBody {
	s.ExternalOutputUri = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStartedTime(v string) *DescribeExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetExternalInputUri(v string) *DescribeExecutionResponseBody {
	s.ExternalInputUri = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetFlowDefinition(v string) *DescribeExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetName(v string) *DescribeExecutionResponseBody {
	s.Name = &v
	return s
}

type DescribeExecutionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutionResponse) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponse) SetHeaders(v map[string]*string) *DescribeExecutionResponse {
	s.Headers = v
	return s
}

func (s *DescribeExecutionResponse) SetBody(v *DescribeExecutionResponseBody) *DescribeExecutionResponse {
	s.Body = v
	return s
}

type DescribeFlowRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) SetRequestId(v string) *DescribeFlowRequest {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowRequest) SetName(v string) *DescribeFlowRequest {
	s.Name = &v
	return s
}

type DescribeFlowResponseBody struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime             *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	LastModifiedTime        *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) SetType(v string) *DescribeFlowResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCreatedTime(v string) *DescribeFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDefinition(v string) *DescribeFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *DescribeFlowResponseBody) SetLastModifiedTime(v string) *DescribeFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetId(v string) *DescribeFlowResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowResponseBody) SetExternalStorageLocation(v string) *DescribeFlowResponseBody {
	s.ExternalStorageLocation = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRoleArn(v string) *DescribeFlowResponseBody {
	s.RoleArn = &v
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

type DescribeScheduleRequest struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName     *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DescribeScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleRequest) SetRequestId(v string) *DescribeScheduleRequest {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleRequest) SetFlowName(v string) *DescribeScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeScheduleRequest) SetScheduleName(v string) *DescribeScheduleRequest {
	s.ScheduleName = &v
	return s
}

type DescribeScheduleResponseBody struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Payload          *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	CronExpression   *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	ScheduleId       *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DescribeScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleResponseBody) SetDescription(v string) *DescribeScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetCreatedTime(v string) *DescribeScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetRequestId(v string) *DescribeScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetLastModifiedTime(v string) *DescribeScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetEnable(v bool) *DescribeScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetPayload(v string) *DescribeScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetCronExpression(v string) *DescribeScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetScheduleId(v string) *DescribeScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetScheduleName(v string) *DescribeScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

type DescribeScheduleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleResponse) SetHeaders(v map[string]*string) *DescribeScheduleResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleResponse) SetBody(v *DescribeScheduleResponseBody) *DescribeScheduleResponse {
	s.Body = v
	return s
}

type GetExecutionHistoryRequest struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName      *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Limit         *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s GetExecutionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryRequest) SetRequestId(v string) *GetExecutionHistoryRequest {
	s.RequestId = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetFlowName(v string) *GetExecutionHistoryRequest {
	s.FlowName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetExecutionName(v string) *GetExecutionHistoryRequest {
	s.ExecutionName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetNextToken(v string) *GetExecutionHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetLimit(v int32) *GetExecutionHistoryRequest {
	s.Limit = &v
	return s
}

type GetExecutionHistoryResponseBody struct {
	NextToken *string                                  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*GetExecutionHistoryResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s GetExecutionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBody) SetNextToken(v string) *GetExecutionHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetRequestId(v string) *GetExecutionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetEvents(v []*GetExecutionHistoryResponseBodyEvents) *GetExecutionHistoryResponseBody {
	s.Events = v
	return s
}

type GetExecutionHistoryResponseBodyEvents struct {
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EventId         *int64  `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Time            *string `json:"Time,omitempty" xml:"Time,omitempty"`
	ScheduleEventId *int64  `json:"ScheduleEventId,omitempty" xml:"ScheduleEventId,omitempty"`
	EventDetail     *string `json:"EventDetail,omitempty" xml:"EventDetail,omitempty"`
	StepName        *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
}

func (s GetExecutionHistoryResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBodyEvents) SetType(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Type = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetTime(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetScheduleEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.ScheduleEventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventDetail(v string) *GetExecutionHistoryResponseBodyEvents {
	s.EventDetail = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetStepName(v string) *GetExecutionHistoryResponseBodyEvents {
	s.StepName = &v
	return s
}

type GetExecutionHistoryResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetExecutionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetExecutionHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponse) SetHeaders(v map[string]*string) *GetExecutionHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetExecutionHistoryResponse) SetBody(v *GetExecutionHistoryResponseBody) *GetExecutionHistoryResponse {
	s.Body = v
	return s
}

type ListExecutionsRequest struct {
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName            *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	NextToken           *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Limit               *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartedTimeBegin    *string `json:"StartedTimeBegin,omitempty" xml:"StartedTimeBegin,omitempty"`
	StartedTimeEnd      *string `json:"StartedTimeEnd,omitempty" xml:"StartedTimeEnd,omitempty"`
	ExecutionNamePrefix *string `json:"ExecutionNamePrefix,omitempty" xml:"ExecutionNamePrefix,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) SetRequestId(v string) *ListExecutionsRequest {
	s.RequestId = &v
	return s
}

func (s *ListExecutionsRequest) SetFlowName(v string) *ListExecutionsRequest {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsRequest) SetLimit(v int32) *ListExecutionsRequest {
	s.Limit = &v
	return s
}

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsRequest) SetStartedTimeBegin(v string) *ListExecutionsRequest {
	s.StartedTimeBegin = &v
	return s
}

func (s *ListExecutionsRequest) SetStartedTimeEnd(v string) *ListExecutionsRequest {
	s.StartedTimeEnd = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutionNamePrefix(v string) *ListExecutionsRequest {
	s.ExecutionNamePrefix = &v
	return s
}

type ListExecutionsResponseBody struct {
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBody) SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListExecutionsResponseBody) SetNextToken(v string) *ListExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsResponseBody) SetRequestId(v string) *ListExecutionsResponseBody {
	s.RequestId = &v
	return s
}

type ListExecutionsResponseBodyExecutions struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StoppedTime       *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
	StartedTime       *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	FlowDefinition    *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	ExternalInputUri  *string `json:"ExternalInputUri,omitempty" xml:"ExternalInputUri,omitempty"`
	Output            *string `json:"Output,omitempty" xml:"Output,omitempty"`
	FlowName          *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ExternalOutputUri *string `json:"ExternalOutputUri,omitempty" xml:"ExternalOutputUri,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Input             *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s ListExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStoppedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StoppedTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StartedTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowDefinition(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowDefinition = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExternalInputUri(v string) *ListExecutionsResponseBodyExecutions {
	s.ExternalInputUri = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutput(v string) *ListExecutionsResponseBodyExecutions {
	s.Output = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowName(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExternalOutputUri(v string) *ListExecutionsResponseBodyExecutions {
	s.ExternalOutputUri = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetName(v string) *ListExecutionsResponseBodyExecutions {
	s.Name = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetInput(v string) *ListExecutionsResponseBodyExecutions {
	s.Input = &v
	return s
}

type ListExecutionsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponse) SetHeaders(v map[string]*string) *ListExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionsResponse) SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse {
	s.Body = v
	return s
}

type ListFlowsRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetRequestId(v string) *ListFlowsRequest {
	s.RequestId = &v
	return s
}

func (s *ListFlowsRequest) SetNextToken(v string) *ListFlowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFlowsRequest) SetLimit(v int32) *ListFlowsRequest {
	s.Limit = &v
	return s
}

type ListFlowsResponseBody struct {
	NextToken *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Flows     []*ListFlowsResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
}

func (s ListFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) SetNextToken(v string) *ListFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) SetFlows(v []*ListFlowsResponseBodyFlows) *ListFlowsResponseBody {
	s.Flows = v
	return s
}

type ListFlowsResponseBodyFlows struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	CreatedTime             *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	LastModifiedTime        *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFlowsResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlows) SetType(v string) *ListFlowsResponseBodyFlows {
	s.Type = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDefinition(v string) *ListFlowsResponseBodyFlows {
	s.Definition = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetRoleArn(v string) *ListFlowsResponseBodyFlows {
	s.RoleArn = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDescription(v string) *ListFlowsResponseBodyFlows {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetExternalStorageLocation(v string) *ListFlowsResponseBodyFlows {
	s.ExternalStorageLocation = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetName(v string) *ListFlowsResponseBodyFlows {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetCreatedTime(v string) *ListFlowsResponseBodyFlows {
	s.CreatedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetLastModifiedTime(v string) *ListFlowsResponseBodyFlows {
	s.LastModifiedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetId(v string) *ListFlowsResponseBodyFlows {
	s.Id = &v
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

type ListSchedulesRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName  *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Limit     *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) SetRequestId(v string) *ListSchedulesRequest {
	s.RequestId = &v
	return s
}

func (s *ListSchedulesRequest) SetFlowName(v string) *ListSchedulesRequest {
	s.FlowName = &v
	return s
}

func (s *ListSchedulesRequest) SetNextToken(v string) *ListSchedulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSchedulesRequest) SetLimit(v int32) *ListSchedulesRequest {
	s.Limit = &v
	return s
}

type ListSchedulesResponseBody struct {
	Schedules []*ListSchedulesResponseBodySchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
	NextToken *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) SetSchedules(v []*ListSchedulesResponseBodySchedules) *ListSchedulesResponseBody {
	s.Schedules = v
	return s
}

func (s *ListSchedulesResponseBody) SetNextToken(v string) *ListSchedulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSchedulesResponseBody) SetRequestId(v string) *ListSchedulesResponseBody {
	s.RequestId = &v
	return s
}

type ListSchedulesResponseBodySchedules struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ScheduleId       *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	Payload          *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	CronExpression   *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ListSchedulesResponseBodySchedules) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodySchedules) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodySchedules) SetDescription(v string) *ListSchedulesResponseBodySchedules {
	s.Description = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleId(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleId = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetPayload(v string) *ListSchedulesResponseBodySchedules {
	s.Payload = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleName(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleName = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetCreatedTime(v string) *ListSchedulesResponseBodySchedules {
	s.CreatedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetLastModifiedTime(v string) *ListSchedulesResponseBodySchedules {
	s.LastModifiedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetCronExpression(v string) *ListSchedulesResponseBodySchedules {
	s.CronExpression = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetEnable(v bool) *ListSchedulesResponseBodySchedules {
	s.Enable = &v
	return s
}

type ListSchedulesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSchedulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponse) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponse) SetHeaders(v map[string]*string) *ListSchedulesResponse {
	s.Headers = v
	return s
}

func (s *ListSchedulesResponse) SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse {
	s.Body = v
	return s
}

type ReportTaskFailedRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
	Error     *string `json:"Error,omitempty" xml:"Error,omitempty"`
	Cause     *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
}

func (s ReportTaskFailedRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskFailedRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedRequest) SetRequestId(v string) *ReportTaskFailedRequest {
	s.RequestId = &v
	return s
}

func (s *ReportTaskFailedRequest) SetTaskToken(v string) *ReportTaskFailedRequest {
	s.TaskToken = &v
	return s
}

func (s *ReportTaskFailedRequest) SetError(v string) *ReportTaskFailedRequest {
	s.Error = &v
	return s
}

func (s *ReportTaskFailedRequest) SetCause(v string) *ReportTaskFailedRequest {
	s.Cause = &v
	return s
}

type ReportTaskFailedResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventId   *int64  `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s ReportTaskFailedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskFailedResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedResponseBody) SetRequestId(v string) *ReportTaskFailedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportTaskFailedResponseBody) SetEventId(v int64) *ReportTaskFailedResponseBody {
	s.EventId = &v
	return s
}

type ReportTaskFailedResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportTaskFailedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportTaskFailedResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskFailedResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedResponse) SetHeaders(v map[string]*string) *ReportTaskFailedResponse {
	s.Headers = v
	return s
}

func (s *ReportTaskFailedResponse) SetBody(v *ReportTaskFailedResponseBody) *ReportTaskFailedResponse {
	s.Body = v
	return s
}

type ReportTaskSucceededRequest struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
	Output    *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ReportTaskSucceededRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskSucceededRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededRequest) SetRequestId(v string) *ReportTaskSucceededRequest {
	s.RequestId = &v
	return s
}

func (s *ReportTaskSucceededRequest) SetTaskToken(v string) *ReportTaskSucceededRequest {
	s.TaskToken = &v
	return s
}

func (s *ReportTaskSucceededRequest) SetOutput(v string) *ReportTaskSucceededRequest {
	s.Output = &v
	return s
}

type ReportTaskSucceededResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EventId   *int64  `json:"EventId,omitempty" xml:"EventId,omitempty"`
}

func (s ReportTaskSucceededResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskSucceededResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededResponseBody) SetRequestId(v string) *ReportTaskSucceededResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportTaskSucceededResponseBody) SetEventId(v int64) *ReportTaskSucceededResponseBody {
	s.EventId = &v
	return s
}

type ReportTaskSucceededResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportTaskSucceededResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportTaskSucceededResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskSucceededResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededResponse) SetHeaders(v map[string]*string) *ReportTaskSucceededResponse {
	s.Headers = v
	return s
}

func (s *ReportTaskSucceededResponse) SetBody(v *ReportTaskSucceededResponseBody) *ReportTaskSucceededResponse {
	s.Body = v
	return s
}

type StartExecutionRequest struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName             *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ExecutionName        *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	Input                *string `json:"Input,omitempty" xml:"Input,omitempty"`
	CallbackFnFTaskToken *string `json:"CallbackFnFTaskToken,omitempty" xml:"CallbackFnFTaskToken,omitempty"`
}

func (s StartExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) SetRequestId(v string) *StartExecutionRequest {
	s.RequestId = &v
	return s
}

func (s *StartExecutionRequest) SetFlowName(v string) *StartExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StartExecutionRequest) SetExecutionName(v string) *StartExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StartExecutionRequest) SetInput(v string) *StartExecutionRequest {
	s.Input = &v
	return s
}

func (s *StartExecutionRequest) SetCallbackFnFTaskToken(v string) *StartExecutionRequest {
	s.CallbackFnFTaskToken = &v
	return s
}

type StartExecutionResponseBody struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Input             *string `json:"Input,omitempty" xml:"Input,omitempty"`
	StoppedTime       *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName          *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Output            *string `json:"Output,omitempty" xml:"Output,omitempty"`
	ExternalOutputUri *string `json:"ExternalOutputUri,omitempty" xml:"ExternalOutputUri,omitempty"`
	StartedTime       *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	ExternalInputUri  *string `json:"ExternalInputUri,omitempty" xml:"ExternalInputUri,omitempty"`
	FlowDefinition    *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StartExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) SetStatus(v string) *StartExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBody) SetInput(v string) *StartExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StartExecutionResponseBody) SetStoppedTime(v string) *StartExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartExecutionResponseBody) SetFlowName(v string) *StartExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StartExecutionResponseBody) SetOutput(v string) *StartExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StartExecutionResponseBody) SetExternalOutputUri(v string) *StartExecutionResponseBody {
	s.ExternalOutputUri = &v
	return s
}

func (s *StartExecutionResponseBody) SetStartedTime(v string) *StartExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StartExecutionResponseBody) SetExternalInputUri(v string) *StartExecutionResponseBody {
	s.ExternalInputUri = &v
	return s
}

func (s *StartExecutionResponseBody) SetFlowDefinition(v string) *StartExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StartExecutionResponseBody) SetName(v string) *StartExecutionResponseBody {
	s.Name = &v
	return s
}

type StartExecutionResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartExecutionResponse) SetHeaders(v map[string]*string) *StartExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartExecutionResponse) SetBody(v *StartExecutionResponseBody) *StartExecutionResponse {
	s.Body = v
	return s
}

type StopExecutionRequest struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName      *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	Cause         *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	Error         *string `json:"Error,omitempty" xml:"Error,omitempty"`
}

func (s StopExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopExecutionRequest) SetRequestId(v string) *StopExecutionRequest {
	s.RequestId = &v
	return s
}

func (s *StopExecutionRequest) SetFlowName(v string) *StopExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StopExecutionRequest) SetExecutionName(v string) *StopExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StopExecutionRequest) SetCause(v string) *StopExecutionRequest {
	s.Cause = &v
	return s
}

func (s *StopExecutionRequest) SetError(v string) *StopExecutionRequest {
	s.Error = &v
	return s
}

type StopExecutionResponseBody struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Input             *string `json:"Input,omitempty" xml:"Input,omitempty"`
	StoppedTime       *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName          *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Output            *string `json:"Output,omitempty" xml:"Output,omitempty"`
	ExternalOutputUri *string `json:"ExternalOutputUri,omitempty" xml:"ExternalOutputUri,omitempty"`
	StartedTime       *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	ExternalInputUri  *string `json:"ExternalInputUri,omitempty" xml:"ExternalInputUri,omitempty"`
	FlowDefinition    *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s StopExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopExecutionResponseBody) SetStatus(v string) *StopExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StopExecutionResponseBody) SetInput(v string) *StopExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StopExecutionResponseBody) SetStoppedTime(v string) *StopExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *StopExecutionResponseBody) SetRequestId(v string) *StopExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopExecutionResponseBody) SetFlowName(v string) *StopExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StopExecutionResponseBody) SetOutput(v string) *StopExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StopExecutionResponseBody) SetExternalOutputUri(v string) *StopExecutionResponseBody {
	s.ExternalOutputUri = &v
	return s
}

func (s *StopExecutionResponseBody) SetStartedTime(v string) *StopExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StopExecutionResponseBody) SetExternalInputUri(v string) *StopExecutionResponseBody {
	s.ExternalInputUri = &v
	return s
}

func (s *StopExecutionResponseBody) SetFlowDefinition(v string) *StopExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StopExecutionResponseBody) SetName(v string) *StopExecutionResponseBody {
	s.Name = &v
	return s
}

type StopExecutionResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopExecutionResponse) SetHeaders(v map[string]*string) *StopExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopExecutionResponse) SetBody(v *StopExecutionResponseBody) *StopExecutionResponse {
	s.Body = v
	return s
}

type UpdateFlowRequest struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
}

func (s UpdateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowRequest) GoString() string {
	return s.String()
}

func (s *UpdateFlowRequest) SetRequestId(v string) *UpdateFlowRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowRequest) SetName(v string) *UpdateFlowRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowRequest) SetDefinition(v string) *UpdateFlowRequest {
	s.Definition = &v
	return s
}

func (s *UpdateFlowRequest) SetDescription(v string) *UpdateFlowRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowRequest) SetType(v string) *UpdateFlowRequest {
	s.Type = &v
	return s
}

func (s *UpdateFlowRequest) SetRoleArn(v string) *UpdateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowRequest) SetExternalStorageLocation(v string) *UpdateFlowRequest {
	s.ExternalStorageLocation = &v
	return s
}

type UpdateFlowResponseBody struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime             *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Definition              *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	LastModifiedTime        *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	RoleArn                 *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBody) SetType(v string) *UpdateFlowResponseBody {
	s.Type = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDescription(v string) *UpdateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateFlowResponseBody) SetCreatedTime(v string) *UpdateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRequestId(v string) *UpdateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDefinition(v string) *UpdateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *UpdateFlowResponseBody) SetLastModifiedTime(v string) *UpdateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetId(v string) *UpdateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateFlowResponseBody) SetExternalStorageLocation(v string) *UpdateFlowResponseBody {
	s.ExternalStorageLocation = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRoleArn(v string) *UpdateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowResponseBody) SetName(v string) *UpdateFlowResponseBody {
	s.Name = &v
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

type UpdateScheduleRequest struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	FlowName       *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	ScheduleName   *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Payload        *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Enable         *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleRequest) SetRequestId(v string) *UpdateScheduleRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleRequest) SetFlowName(v string) *UpdateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateScheduleRequest) SetScheduleName(v string) *UpdateScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *UpdateScheduleRequest) SetDescription(v string) *UpdateScheduleRequest {
	s.Description = &v
	return s
}

func (s *UpdateScheduleRequest) SetPayload(v string) *UpdateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleRequest) SetCronExpression(v string) *UpdateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleRequest) SetEnable(v bool) *UpdateScheduleRequest {
	s.Enable = &v
	return s
}

type UpdateScheduleResponseBody struct {
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Enable           *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Payload          *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	CronExpression   *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	ScheduleId       *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s UpdateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleResponseBody) SetDescription(v string) *UpdateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetCreatedTime(v string) *UpdateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetRequestId(v string) *UpdateScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetLastModifiedTime(v string) *UpdateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetEnable(v bool) *UpdateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetPayload(v string) *UpdateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetCronExpression(v string) *UpdateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetScheduleId(v string) *UpdateScheduleResponseBody {
	s.ScheduleId = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetScheduleName(v string) *UpdateScheduleResponseBody {
	s.ScheduleName = &v
	return s
}

type UpdateScheduleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduleResponse) SetHeaders(v map[string]*string) *UpdateScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduleResponse) SetBody(v *UpdateScheduleResponseBody) *UpdateScheduleResponse {
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
		"cn-beijing":  tea.String("cn-beijing.fnf.aliyuncs.com"),
		"cn-hangzhou": tea.String("cn-hangzhou.fnf.aliyuncs.com"),
		"cn-shanghai": tea.String("cn-shanghai.fnf.aliyuncs.com"),
		"cn-shenzhen": tea.String("cn-shenzhen.fnf.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("fnf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlow"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateScheduleWithOptions(request *CreateScheduleRequest, runtime *util.RuntimeOptions) (_result *CreateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSchedule"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSchedule(request *CreateScheduleRequest) (_result *CreateScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduleResponse{}
	_body, _err := client.CreateScheduleWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlow"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteScheduleWithOptions(request *DeleteScheduleRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSchedule"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSchedule(request *DeleteScheduleRequest) (_result *DeleteScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.DeleteScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExecutionWithOptions(request *DescribeExecutionRequest, runtime *util.RuntimeOptions) (_result *DescribeExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExecution"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExecution(request *DescribeExecutionRequest) (_result *DescribeExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExecutionResponse{}
	_body, _err := client.DescribeExecutionWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlow"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeScheduleWithOptions(request *DescribeScheduleRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSchedule"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSchedule(request *DescribeScheduleRequest) (_result *DescribeScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduleResponse{}
	_body, _err := client.DescribeScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExecutionHistoryWithOptions(request *GetExecutionHistoryRequest, runtime *util.RuntimeOptions) (_result *GetExecutionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetExecutionHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetExecutionHistory"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExecutionHistory(request *GetExecutionHistoryRequest) (_result *GetExecutionHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExecutionHistoryResponse{}
	_body, _err := client.GetExecutionHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListExecutionsWithOptions(request *ListExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListExecutions"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListExecutions(request *ListExecutionsRequest) (_result *ListExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionsResponse{}
	_body, _err := client.ListExecutionsWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlows"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, runtime *util.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSchedulesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSchedules"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSchedules(request *ListSchedulesRequest) (_result *ListSchedulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSchedulesResponse{}
	_body, _err := client.ListSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportTaskFailedWithOptions(request *ReportTaskFailedRequest, runtime *util.RuntimeOptions) (_result *ReportTaskFailedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReportTaskFailedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReportTaskFailed"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportTaskFailed(request *ReportTaskFailedRequest) (_result *ReportTaskFailedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportTaskFailedResponse{}
	_body, _err := client.ReportTaskFailedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportTaskSucceededWithOptions(request *ReportTaskSucceededRequest, runtime *util.RuntimeOptions) (_result *ReportTaskSucceededResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReportTaskSucceededResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReportTaskSucceeded"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportTaskSucceeded(request *ReportTaskSucceededRequest) (_result *ReportTaskSucceededResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportTaskSucceededResponse{}
	_body, _err := client.ReportTaskSucceededWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartExecutionWithOptions(request *StartExecutionRequest, runtime *util.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartExecution"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartExecution(request *StartExecutionRequest) (_result *StartExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartExecutionResponse{}
	_body, _err := client.StartExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopExecutionWithOptions(request *StopExecutionRequest, runtime *util.RuntimeOptions) (_result *StopExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopExecutionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopExecution"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopExecution(request *StopExecutionRequest) (_result *StopExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopExecutionResponse{}
	_body, _err := client.StopExecutionWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UpdateFlow"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpdateScheduleWithOptions(request *UpdateScheduleRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateScheduleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateSchedule"), tea.String("2019-03-15"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSchedule(request *UpdateScheduleRequest) (_result *UpdateScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateScheduleResponse{}
	_body, _err := client.UpdateScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
