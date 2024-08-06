// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateFlowRequest struct {
	// The definition of the workflow. The definition must comply with the flow definition language (FDL) syntax. Considering compatibility, the system supports two flow definition specifications.
	//
	// >  In the preceding flow definition example, Name:my_flow_name is the workflow name, which must be consistent with the input parameter Name
	//
	// This parameter is required.
	//
	// example:
	//
	// version:&nbsp;v1.0<br/>type:&nbsp;flow<br/>steps:<br/>&nbsp;-&nbsp;type:&nbsp;pass<br/>&nbsp;name:&nbsp;mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// test flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values: Express and Standard. Considering compatibility, an empty string is equivalent to the Standard execution mode.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The path of the external storage.
	//
	// example:
	//
	// /path
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	// The name of the flow. The name is unique within the same region and cannot be modified after the flow is created. Set this parameter based on the following rules:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram:${region}:${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow. Set this parameter to **FDL**.
	//
	// This parameter is required.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetExecutionMode(v string) *CreateFlowRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateFlowRequest) SetExternalStorageLocation(v string) *CreateFlowRequest {
	s.ExternalStorageLocation = &v
	return s
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetRoleArn(v string) *CreateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowRequest) SetType(v string) *CreateFlowRequest {
	s.Type = &v
	return s
}

type CreateFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Considering compatibility, the system supports two flow definition specifications.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values: Express and Standard. Considering compatibility, an empty string is equivalent to the Standard execution mode.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The unique ID of the flow.
	//
	// example:
	//
	// e589e092-e2c0-4dee-b306-3574ddfdddf5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the flow was last modified.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID. Each time an `HTTP status code` is returned, Serverless Workflow returns a value for the parameter.
	//
	// example:
	//
	// testRequestID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram:${region}:${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow.
	//
	// Valid value:
	//
	// 	- FDL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetCreatedTime(v string) *CreateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetDefinition(v string) *CreateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *CreateFlowResponseBody) SetDescription(v string) *CreateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *CreateFlowResponseBody) SetExecutionMode(v string) *CreateFlowResponseBody {
	s.ExecutionMode = &v
	return s
}

func (s *CreateFlowResponseBody) SetId(v string) *CreateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowResponseBody) SetLastModifiedTime(v string) *CreateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateFlowResponseBody) SetName(v string) *CreateFlowResponseBody {
	s.Name = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowResponseBody) SetRoleArn(v string) *CreateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *CreateFlowResponseBody) SetType(v string) *CreateFlowResponseBody {
	s.Type = &v
	return s
}

type CreateFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateFlowResponse) SetStatusCode(v int32) *CreateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

type CreateScheduleRequest struct {
	// The CRON expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the time-based schedule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The name of the workflow that is associated with the time-based schedule.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The trigger message of the time-based schedule. Specify the value in the JSON format.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The name of the time-based schedule. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- It is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName     *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
	SignatureVersion *string `json:"SignatureVersion,omitempty" xml:"SignatureVersion,omitempty"`
}

func (s CreateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleRequest) SetCronExpression(v string) *CreateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleRequest) SetDescription(v string) *CreateScheduleRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduleRequest) SetEnable(v bool) *CreateScheduleRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduleRequest) SetFlowName(v string) *CreateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *CreateScheduleRequest) SetPayload(v string) *CreateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *CreateScheduleRequest) SetScheduleName(v string) *CreateScheduleRequest {
	s.ScheduleName = &v
	return s
}

func (s *CreateScheduleRequest) SetSignatureVersion(v string) *CreateScheduleRequest {
	s.SignatureVersion = &v
	return s
}

type CreateScheduleResponseBody struct {
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the time-based schedule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last modified.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The trigger message of the time-based schedule.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the time-based schedule.
	//
	// example:
	//
	// testScheduleId
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The name of the time-based schedule.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s CreateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleResponseBody) SetCreatedTime(v string) *CreateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetCronExpression(v string) *CreateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *CreateScheduleResponseBody) SetDescription(v string) *CreateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *CreateScheduleResponseBody) SetEnable(v bool) *CreateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *CreateScheduleResponseBody) SetLastModifiedTime(v string) *CreateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *CreateScheduleResponseBody) SetPayload(v string) *CreateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *CreateScheduleResponseBody) SetRequestId(v string) *CreateScheduleResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateScheduleResponse) SetStatusCode(v int32) *CreateScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleResponse) SetBody(v *CreateScheduleResponseBody) *CreateScheduleResponse {
	s.Body = v
	return s
}

type DeleteFlowRequest struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetName(v string) *DeleteFlowRequest {
	s.Name = &v
	return s
}

type DeleteFlowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// testRequestId
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteFlowResponse) SetStatusCode(v int32) *DeleteFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

type DeleteScheduleRequest struct {
	// This parameter is required.
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DeleteScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduleRequest) GoString() string {
	return s.String()
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
	// The request ID.
	//
	// example:
	//
	// testRequestId
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteScheduleResponse) SetStatusCode(v int32) *DeleteScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduleResponse) SetBody(v *DeleteScheduleResponseBody) *DeleteScheduleResponse {
	s.Body = v
	return s
}

type DescribeExecutionRequest struct {
	// The name of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The maximum period of time for long polling waits. Valid values: 0 to 60. Unit: seconds. Configure this parameter based on the following rules:
	//
	// 	- If the value is 0, the system immediately returns the current execution status.
	//
	// 	- If the value is greater than 0, the long polling request waits until the execution ends or the specified period elapses.
	//
	// example:
	//
	// 20
	WaitTimeSeconds *int32 `json:"WaitTimeSeconds,omitempty" xml:"WaitTimeSeconds,omitempty"`
}

func (s DescribeExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutionRequest) SetExecutionName(v string) *DescribeExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *DescribeExecutionRequest) SetFlowName(v string) *DescribeExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionRequest) SetWaitTimeSeconds(v int32) *DescribeExecutionRequest {
	s.WaitTimeSeconds = &v
	return s
}

type DescribeExecutionResponseBody struct {
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// exec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution result, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The execution status. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s DescribeExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponseBody) SetFlowDefinition(v string) *DescribeExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetFlowName(v string) *DescribeExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetInput(v string) *DescribeExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetName(v string) *DescribeExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetOutput(v string) *DescribeExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetRequestId(v string) *DescribeExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStartedTime(v string) *DescribeExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStatus(v string) *DescribeExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStoppedTime(v string) *DescribeExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

type DescribeExecutionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeExecutionResponse) SetStatusCode(v int32) *DescribeExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExecutionResponse) SetBody(v *DescribeExecutionResponseBody) *DescribeExecutionResponse {
	s.Body = v
	return s
}

type DescribeFlowRequest struct {
	// The name of the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) SetName(v string) *DescribeFlowRequest {
	s.Name = &v
	return s
}

type DescribeFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The flow definition, which follows the flow definition language (FDL) syntax standard. Considering compatibility, the system supports the two flow definition specifications.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode or the enumeration type. Valid values: Express and Standard. A value of Standard indicates an empty string.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The unique ID of the flow.
	//
	// example:
	//
	// e589e092-e2c0-4dee-b306-3574ddfdddf5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the flow was last modified.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, CloudFlow assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the workflow.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) SetCreatedTime(v string) *DescribeFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDefinition(v string) *DescribeFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetExecutionMode(v string) *DescribeFlowResponseBody {
	s.ExecutionMode = &v
	return s
}

func (s *DescribeFlowResponseBody) SetId(v string) *DescribeFlowResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowResponseBody) SetLastModifiedTime(v string) *DescribeFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeFlowResponseBody) SetName(v string) *DescribeFlowResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRoleArn(v string) *DescribeFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *DescribeFlowResponseBody) SetType(v string) *DescribeFlowResponseBody {
	s.Type = &v
	return s
}

type DescribeFlowResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeFlowResponse) SetStatusCode(v int32) *DescribeFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFlowResponse) SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse {
	s.Body = v
	return s
}

type DescribeScheduleRequest struct {
	// The name of the flow that is associated with the time-based schedule. The name must be unique within the region and cannot be modified after the time-based schedule is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The name of the time-based schedule. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DescribeScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleRequest) GoString() string {
	return s.String()
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
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the time-based schedule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last modified.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The trigger message of the time-based schedule.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the time-based schedule.
	//
	// example:
	//
	// testScheduleId
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The name of the time-based schedule.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s DescribeScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleResponseBody) SetCreatedTime(v string) *DescribeScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetCronExpression(v string) *DescribeScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetDescription(v string) *DescribeScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetEnable(v bool) *DescribeScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetLastModifiedTime(v string) *DescribeScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetPayload(v string) *DescribeScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeScheduleResponseBody) SetRequestId(v string) *DescribeScheduleResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeScheduleResponse) SetStatusCode(v int32) *DescribeScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduleResponse) SetBody(v *DescribeScheduleResponseBody) *DescribeScheduleResponse {
	s.Body = v
	return s
}

type GetExecutionHistoryRequest struct {
	// The name of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of workflows that you want to query. Valid values: 1-999. Default value: 60.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event to start the query. You can obtain the value from the response data.
	//
	// example:
	//
	// flow_xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetExecutionHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryRequest) SetExecutionName(v string) *GetExecutionHistoryRequest {
	s.ExecutionName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetFlowName(v string) *GetExecutionHistoryRequest {
	s.FlowName = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetLimit(v int32) *GetExecutionHistoryRequest {
	s.Limit = &v
	return s
}

func (s *GetExecutionHistoryRequest) SetNextToken(v string) *GetExecutionHistoryRequest {
	s.NextToken = &v
	return s
}

type GetExecutionHistoryResponseBody struct {
	// The events.
	Events []*GetExecutionHistoryResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// You do not need to specify this parameter for the first request. The returned value of **ScheduleEventId*	- is used as the token for the next query. No value is returned for the last query.
	//
	// example:
	//
	// 3
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecutionHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBody) SetEvents(v []*GetExecutionHistoryResponseBodyEvents) *GetExecutionHistoryResponseBody {
	s.Events = v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetNextToken(v string) *GetExecutionHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetExecutionHistoryResponseBody) SetRequestId(v string) *GetExecutionHistoryResponseBody {
	s.RequestId = &v
	return s
}

type GetExecutionHistoryResponseBodyEvents struct {
	// The details about the execution step.
	//
	// example:
	//
	// {}
	EventDetail *string `json:"EventDetail,omitempty" xml:"EventDetail,omitempty"`
	// The ID of the execution step.
	//
	// example:
	//
	// 2
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The ID of the scheduling step.
	//
	// example:
	//
	// 1
	ScheduleEventId *int64 `json:"ScheduleEventId,omitempty" xml:"ScheduleEventId,omitempty"`
	// The name of the execution step.
	//
	// example:
	//
	// passStep
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The time when the event was updated.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the execution step. Valid values:
	//
	// 	- **StepEntered**
	//
	// 	- **StepStarted**
	//
	// 	- **StepSucceeded**
	//
	// 	- **StepFailed**
	//
	// 	- **StepExited**
	//
	// 	- **BranchEntered**
	//
	// 	- **BranchExited**
	//
	// 	- **IterationEntered**
	//
	// 	- **IterationExited**
	//
	// 	- **TaskScheduled**
	//
	// 	- **TaskStarted**
	//
	// 	- **TaskSubmitted**
	//
	// 	- **TaskSubmitFailed**
	//
	// 	- **TaskSucceeded**
	//
	// 	- **TaskFailed**
	//
	// 	- **TaskTimedOut**
	//
	// 	- **ExecutionStarted**
	//
	// 	- **ExecutionStopped**
	//
	// 	- **ExecutionSucceeded**
	//
	// 	- **ExecutionFailed**
	//
	// 	- **ExecutionTimedOut**
	//
	// example:
	//
	// TaskSucceeded
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExecutionHistoryResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionHistoryResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventDetail(v string) *GetExecutionHistoryResponseBodyEvents {
	s.EventDetail = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetScheduleEventId(v int64) *GetExecutionHistoryResponseBodyEvents {
	s.ScheduleEventId = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetStepName(v string) *GetExecutionHistoryResponseBodyEvents {
	s.StepName = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetTime(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Time = &v
	return s
}

func (s *GetExecutionHistoryResponseBodyEvents) SetType(v string) *GetExecutionHistoryResponseBodyEvents {
	s.Type = &v
	return s
}

type GetExecutionHistoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecutionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetExecutionHistoryResponse) SetStatusCode(v int32) *GetExecutionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecutionHistoryResponse) SetBody(v *GetExecutionHistoryResponseBody) *GetExecutionHistoryResponse {
	s.Body = v
	return s
}

type ListExecutionsRequest struct {
	// The name prefix of the execution.
	//
	// example:
	//
	// run
	ExecutionNamePrefix *string `json:"ExecutionNamePrefix,omitempty" xml:"ExecutionNamePrefix,omitempty"`
	// The name of the flow. The name must be unique within the region and cannot be modified after the flow is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of executions that you want to query. Valid values: 1-99. Default value: 60.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the execution to start the query. You can obtain the value from the response data. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// flow_xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The beginning of the time range to query executions. Specify the value in the UTC RFC3339 format.
	//
	// example:
	//
	// 2020-12-02T02:39:20.402Z
	StartedTimeBegin *string `json:"StartedTimeBegin,omitempty" xml:"StartedTimeBegin,omitempty"`
	// The end of the time range to query executions. Specify the value in the UTC RFC3339 format.
	//
	// example:
	//
	// 2020-12-02T02:23:54.817Z
	StartedTimeEnd *string `json:"StartedTimeEnd,omitempty" xml:"StartedTimeEnd,omitempty"`
	// The status of the execution that you want to filter. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) SetExecutionNamePrefix(v string) *ListExecutionsRequest {
	s.ExecutionNamePrefix = &v
	return s
}

func (s *ListExecutionsRequest) SetFlowName(v string) *ListExecutionsRequest {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsRequest) SetLimit(v int32) *ListExecutionsRequest {
	s.Limit = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
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

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

type ListExecutionsResponseBody struct {
	// The information about executions.
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The start key for the next query. This parameter is not returned if all results have been returned.
	//
	// example:
	//
	// exec2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// exec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the execution, which is in the JSON format
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the execution.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s ListExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowDefinition(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowDefinition = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowName(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetInput(v string) *ListExecutionsResponseBodyExecutions {
	s.Input = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetName(v string) *ListExecutionsResponseBodyExecutions {
	s.Name = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutput(v string) *ListExecutionsResponseBodyExecutions {
	s.Output = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StartedTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStoppedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StoppedTime = &v
	return s
}

type ListExecutionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListExecutionsResponse) SetStatusCode(v int32) *ListExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionsResponse) SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse {
	s.Body = v
	return s
}

type ListFlowsRequest struct {
	// The number of workflows that you want to query. Valid values: 1 - 999. Default value: 60.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token to start the query.
	//
	// example:
	//
	// flow_nextxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetLimit(v int32) *ListFlowsRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowsRequest) SetNextToken(v string) *ListFlowsRequest {
	s.NextToken = &v
	return s
}

type ListFlowsResponseBody struct {
	// The details of flows.
	Flows []*ListFlowsResponseBodyFlows `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	// The start key for the next query. This parameter is not returned if all results have been returned.
	//
	// example:
	//
	// flow_nextxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *ListFlowsResponseBody) SetNextToken(v string) *ListFlowsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

type ListFlowsResponseBodyFlows struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The definition of the flow. The definition must comply with the Flow Definition Language (FDL) syntax.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test flow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode or the enumeration type. Valid values: Express and Standard. A value of Standard indicates an empty string.
	//
	// example:
	//
	// Standard
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The unique ID of the flow.
	//
	// example:
	//
	// e589e092-e2c0-4dee-b306-3574ddf5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the flow was last modified.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the specified Resource Access Management (RAM) role that Serverless Workflow assumes to invoke resources when the flow is executed.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFlowsResponseBodyFlows) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlows) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlows) SetCreatedTime(v string) *ListFlowsResponseBodyFlows {
	s.CreatedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDefinition(v string) *ListFlowsResponseBodyFlows {
	s.Definition = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetDescription(v string) *ListFlowsResponseBodyFlows {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetExecutionMode(v string) *ListFlowsResponseBodyFlows {
	s.ExecutionMode = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetId(v string) *ListFlowsResponseBodyFlows {
	s.Id = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetLastModifiedTime(v string) *ListFlowsResponseBodyFlows {
	s.LastModifiedTime = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetName(v string) *ListFlowsResponseBodyFlows {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetRoleArn(v string) *ListFlowsResponseBodyFlows {
	s.RoleArn = &v
	return s
}

func (s *ListFlowsResponseBodyFlows) SetType(v string) *ListFlowsResponseBodyFlows {
	s.Type = &v
	return s
}

type ListFlowsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListFlowsResponse) SetStatusCode(v int32) *ListFlowsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowsResponse) SetBody(v *ListFlowsResponseBody) *ListFlowsResponse {
	s.Body = v
	return s
}

type ListSchedulesRequest struct {
	// The name of the flow that is associated with the time-based schedules. The name is unique within the region and cannot be modified after the flow is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The number of schedules to be queried. Valid values: 1 to 1000.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// For the first query, you do not need to specify this parameter. The system uses the value of the **FlowName*	- parameter as the value of the **NextToken*	- parameter. When the query ends, no value is returned for this parameter.
	//
	// example:
	//
	// testNextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSchedulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListSchedulesRequest) SetFlowName(v string) *ListSchedulesRequest {
	s.FlowName = &v
	return s
}

func (s *ListSchedulesRequest) SetLimit(v int32) *ListSchedulesRequest {
	s.Limit = &v
	return s
}

func (s *ListSchedulesRequest) SetNextToken(v string) *ListSchedulesRequest {
	s.NextToken = &v
	return s
}

type ListSchedulesResponseBody struct {
	// The token for the next query.
	//
	// example:
	//
	// testNextToken
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time-based schedules that are queried.
	Schedules []*ListSchedulesResponseBodySchedules `json:"Schedules,omitempty" xml:"Schedules,omitempty" type:"Repeated"`
}

func (s ListSchedulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBody) SetNextToken(v string) *ListSchedulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSchedulesResponseBody) SetRequestId(v string) *ListSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSchedulesResponseBody) SetSchedules(v []*ListSchedulesResponseBodySchedules) *ListSchedulesResponseBody {
	s.Schedules = v
	return s
}

type ListSchedulesResponseBodySchedules struct {
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cron expression of the scheduled task.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the time-based schedule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last modified.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The trigger message of the time-based schedule.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The ID of the time-based schedule.
	//
	// example:
	//
	// testScheduleId
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The name of the time-based schedule.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s ListSchedulesResponseBodySchedules) String() string {
	return tea.Prettify(s)
}

func (s ListSchedulesResponseBodySchedules) GoString() string {
	return s.String()
}

func (s *ListSchedulesResponseBodySchedules) SetCreatedTime(v string) *ListSchedulesResponseBodySchedules {
	s.CreatedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetCronExpression(v string) *ListSchedulesResponseBodySchedules {
	s.CronExpression = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetDescription(v string) *ListSchedulesResponseBodySchedules {
	s.Description = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetEnable(v bool) *ListSchedulesResponseBodySchedules {
	s.Enable = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetLastModifiedTime(v string) *ListSchedulesResponseBodySchedules {
	s.LastModifiedTime = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetPayload(v string) *ListSchedulesResponseBodySchedules {
	s.Payload = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleId(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleId = &v
	return s
}

func (s *ListSchedulesResponseBodySchedules) SetScheduleName(v string) *ListSchedulesResponseBodySchedules {
	s.ScheduleName = &v
	return s
}

type ListSchedulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListSchedulesResponse) SetStatusCode(v int32) *ListSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSchedulesResponse) SetBody(v *ListSchedulesResponseBody) *ListSchedulesResponse {
	s.Body = v
	return s
}

type ReportTaskFailedRequest struct {
	// The cause of the failure. The value must be 1 to 4,096 characters in length.
	//
	// example:
	//
	// emptyString
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The error code for the failed task. The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// nill
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The token of the task whose execution you want to report. The task token is passed to the called service, such as Message Service (MNS) or Function Compute. For MNS, the value of this parameter can be obtained from a message. For Function Compute, the value of this parameter can be obtained from an event. For more information, see [Service integration modes](https://help.aliyun.com/document_detail/2592915.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// emptyString
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
}

func (s ReportTaskFailedRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskFailedRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedRequest) SetCause(v string) *ReportTaskFailedRequest {
	s.Cause = &v
	return s
}

func (s *ReportTaskFailedRequest) SetError(v string) *ReportTaskFailedRequest {
	s.Error = &v
	return s
}

func (s *ReportTaskFailedRequest) SetTaskToken(v string) *ReportTaskFailedRequest {
	s.TaskToken = &v
	return s
}

type ReportTaskFailedResponseBody struct {
	// The ID of the event.
	//
	// example:
	//
	// 1
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportTaskFailedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskFailedResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskFailedResponseBody) SetEventId(v int64) *ReportTaskFailedResponseBody {
	s.EventId = &v
	return s
}

func (s *ReportTaskFailedResponseBody) SetRequestId(v string) *ReportTaskFailedResponseBody {
	s.RequestId = &v
	return s
}

type ReportTaskFailedResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportTaskFailedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ReportTaskFailedResponse) SetStatusCode(v int32) *ReportTaskFailedResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportTaskFailedResponse) SetBody(v *ReportTaskFailedResponseBody) *ReportTaskFailedResponse {
	s.Body = v
	return s
}

type ReportTaskSucceededRequest struct {
	// The output information of the task whose execution success you want to report.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The token of the task whose execution you want to report. The task token is passed to the called service, such as Message Service (MNS) or Function Compute. For MNS, the value of this parameter can be obtained from a message. For Function Compute, the value of this parameter can be obtained from an event. For more information, see [Service integration modes](https://help.aliyun.com/document_detail/2592915.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// emptyString
	TaskToken *string `json:"TaskToken,omitempty" xml:"TaskToken,omitempty"`
}

func (s ReportTaskSucceededRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskSucceededRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededRequest) SetOutput(v string) *ReportTaskSucceededRequest {
	s.Output = &v
	return s
}

func (s *ReportTaskSucceededRequest) SetTaskToken(v string) *ReportTaskSucceededRequest {
	s.TaskToken = &v
	return s
}

type ReportTaskSucceededResponseBody struct {
	// The ID of the event.
	//
	// example:
	//
	// 1
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportTaskSucceededResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskSucceededResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTaskSucceededResponseBody) SetEventId(v int64) *ReportTaskSucceededResponseBody {
	s.EventId = &v
	return s
}

func (s *ReportTaskSucceededResponseBody) SetRequestId(v string) *ReportTaskSucceededResponseBody {
	s.RequestId = &v
	return s
}

type ReportTaskSucceededResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportTaskSucceededResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ReportTaskSucceededResponse) SetStatusCode(v int32) *ReportTaskSucceededResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportTaskSucceededResponse) SetBody(v *ReportTaskSucceededResponseBody) *ReportTaskSucceededResponse {
	s.Body = v
	return s
}

type StartExecutionRequest struct {
	// Specifies that the **TaskToken**-related tasks are called back after the execution in the flow ends.
	//
	// example:
	//
	// 12
	CallbackFnFTaskToken *string `json:"CallbackFnFTaskToken,omitempty" xml:"CallbackFnFTaskToken,omitempty"`
	// The name of the execution. The execution name is unique within a workflow. Configure this parameter based on the following rules:
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s StartExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) SetCallbackFnFTaskToken(v string) *StartExecutionRequest {
	s.CallbackFnFTaskToken = &v
	return s
}

func (s *StartExecutionRequest) SetExecutionName(v string) *StartExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StartExecutionRequest) SetFlowName(v string) *StartExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StartExecutionRequest) SetInput(v string) *StartExecutionRequest {
	s.Input = &v
	return s
}

type StartExecutionResponseBody struct {
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// exec1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution result, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The execution status. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s StartExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) SetFlowDefinition(v string) *StartExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StartExecutionResponseBody) SetFlowName(v string) *StartExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StartExecutionResponseBody) SetInput(v string) *StartExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StartExecutionResponseBody) SetName(v string) *StartExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StartExecutionResponseBody) SetOutput(v string) *StartExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartExecutionResponseBody) SetStartedTime(v string) *StartExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StartExecutionResponseBody) SetStatus(v string) *StartExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBody) SetStoppedTime(v string) *StartExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

type StartExecutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartExecutionResponse) SetStatusCode(v int32) *StartExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartExecutionResponse) SetBody(v *StartExecutionResponseBody) *StartExecutionResponse {
	s.Body = v
	return s
}

type StartSyncExecutionRequest struct {
	// The name of the execution that you want to start. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// Different from the StartExecution operation, in the synchronous execution mode, the execution name is no longer required to be unique within a flow. You can choose to provide an execution name to identify the current execution. In this case, the system adds a UUID to the current execution name. The used format is {ExecutionName}:{UUID}. If you do not specify the execution name, the system automatically generates an execution name.
	//
	// example:
	//
	// my_exec_name
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
}

func (s StartSyncExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSyncExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionRequest) SetExecutionName(v string) *StartSyncExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StartSyncExecutionRequest) SetFlowName(v string) *StartSyncExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StartSyncExecutionRequest) SetInput(v string) *StartSyncExecutionRequest {
	s.Input = &v
	return s
}

type StartSyncExecutionResponseBody struct {
	// The error code that is returned if the execution failed.
	//
	// example:
	//
	// ActionNotSupported
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that indicates the execution timed out.
	//
	// example:
	//
	// Standard execution is not supported
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// my_exec_name:{UUID}
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the execution. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s StartSyncExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSyncExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponseBody) SetErrorCode(v string) *StartSyncExecutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetErrorMessage(v string) *StartSyncExecutionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetFlowName(v string) *StartSyncExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetName(v string) *StartSyncExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetOutput(v string) *StartSyncExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetRequestId(v string) *StartSyncExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStartedTime(v string) *StartSyncExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStatus(v string) *StartSyncExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStoppedTime(v string) *StartSyncExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

type StartSyncExecutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSyncExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSyncExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSyncExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponse) SetHeaders(v map[string]*string) *StartSyncExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartSyncExecutionResponse) SetStatusCode(v int32) *StartSyncExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSyncExecutionResponse) SetBody(v *StartSyncExecutionResponseBody) *StartSyncExecutionResponse {
	s.Body = v
	return s
}

type StopExecutionRequest struct {
	// The reason for stopping the execution. The value must be 1 to 4,096 characters in length.
	//
	// example:
	//
	// for test
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The error for stopping the execution. The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// nill
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The name of the execution to be stopped. You can call the **ListExecutions*	- operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be stopped. You can call the **ListFlows*	- operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s StopExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopExecutionRequest) SetCause(v string) *StopExecutionRequest {
	s.Cause = &v
	return s
}

func (s *StopExecutionRequest) SetError(v string) *StopExecutionRequest {
	s.Error = &v
	return s
}

func (s *StopExecutionRequest) SetExecutionName(v string) *StopExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StopExecutionRequest) SetFlowName(v string) *StopExecutionRequest {
	s.FlowName = &v
	return s
}

type StopExecutionResponseBody struct {
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// exec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution result, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleArn   *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The execution status. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s StopExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopExecutionResponseBody) SetFlowDefinition(v string) *StopExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StopExecutionResponseBody) SetFlowName(v string) *StopExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StopExecutionResponseBody) SetInput(v string) *StopExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StopExecutionResponseBody) SetName(v string) *StopExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StopExecutionResponseBody) SetOutput(v string) *StopExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StopExecutionResponseBody) SetRequestId(v string) *StopExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopExecutionResponseBody) SetRoleArn(v string) *StopExecutionResponseBody {
	s.RoleArn = &v
	return s
}

func (s *StopExecutionResponseBody) SetStartedTime(v string) *StopExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StopExecutionResponseBody) SetStatus(v string) *StopExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StopExecutionResponseBody) SetStoppedTime(v string) *StopExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

type StopExecutionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopExecutionResponse) SetStatusCode(v int32) *StopExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExecutionResponse) SetBody(v *StopExecutionResponseBody) *StopExecutionResponse {
	s.Body = v
	return s
}

type UpdateFlowRequest struct {
	// The definition of the workflow. The definition must comply with the flow definition language (FDL) syntax. Considering compatibility, the system supports the two workflow definition specifications.
	//
	// >  In the preceding workflow definition example, Name:my_flow_name is the workflow name, which must be consistent with the input parameter Name
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test definition
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, the flow execution engine assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow. Valid value: **FDL**.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *UpdateFlowRequest) SetDescription(v string) *UpdateFlowRequest {
	s.Description = &v
	return s
}

func (s *UpdateFlowRequest) SetName(v string) *UpdateFlowRequest {
	s.Name = &v
	return s
}

func (s *UpdateFlowRequest) SetRoleArn(v string) *UpdateFlowRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowRequest) SetType(v string) *UpdateFlowRequest {
	s.Type = &v
	return s
}

type UpdateFlowResponseBody struct {
	// The time when the flow was created.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The flow definition, which follows the FDL syntax standard. Considering compatibility, the system supports the two flow definition specifications.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The description of the flow.
	//
	// example:
	//
	// test definition
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The path of the external storage.
	//
	// example:
	//
	// /path
	ExternalStorageLocation *string `json:"ExternalStorageLocation,omitempty" xml:"ExternalStorageLocation,omitempty"`
	// The unique ID of the flow.
	//
	// example:
	//
	// e589e092-e2c0-4dee-b306-3574ddfdddf5****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the flow was last modified.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the authorized role on which the execution of the flow relies. During the execution of the flow, the flow execution engine assumes the role to call API operations of relevant services.
	//
	// example:
	//
	// acs:ram::${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The type of the flow.
	//
	// example:
	//
	// FDL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFlowResponseBody) SetCreatedTime(v string) *UpdateFlowResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDefinition(v string) *UpdateFlowResponseBody {
	s.Definition = &v
	return s
}

func (s *UpdateFlowResponseBody) SetDescription(v string) *UpdateFlowResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateFlowResponseBody) SetExternalStorageLocation(v string) *UpdateFlowResponseBody {
	s.ExternalStorageLocation = &v
	return s
}

func (s *UpdateFlowResponseBody) SetId(v string) *UpdateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateFlowResponseBody) SetLastModifiedTime(v string) *UpdateFlowResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateFlowResponseBody) SetName(v string) *UpdateFlowResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRequestId(v string) *UpdateFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFlowResponseBody) SetRoleArn(v string) *UpdateFlowResponseBody {
	s.RoleArn = &v
	return s
}

func (s *UpdateFlowResponseBody) SetType(v string) *UpdateFlowResponseBody {
	s.Type = &v
	return s
}

type UpdateFlowResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateFlowResponse) SetStatusCode(v int32) *UpdateFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFlowResponse) SetBody(v *UpdateFlowResponseBody) *UpdateFlowResponse {
	s.Body = v
	return s
}

type UpdateScheduleRequest struct {
	// The CRON expression.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the time-based schedule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The name of the flow that is associated with the time-based schedule. The name must be unique within the region and cannot be modified after the time-based schedule is created. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testFlowName
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The trigger message of the time-based schedule. It must be in the JSON format.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The name of the time-based schedule. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s UpdateScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleRequest) SetCronExpression(v string) *UpdateScheduleRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleRequest) SetDescription(v string) *UpdateScheduleRequest {
	s.Description = &v
	return s
}

func (s *UpdateScheduleRequest) SetEnable(v bool) *UpdateScheduleRequest {
	s.Enable = &v
	return s
}

func (s *UpdateScheduleRequest) SetFlowName(v string) *UpdateScheduleRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateScheduleRequest) SetPayload(v string) *UpdateScheduleRequest {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleRequest) SetScheduleName(v string) *UpdateScheduleRequest {
	s.ScheduleName = &v
	return s
}

type UpdateScheduleResponseBody struct {
	// The time when the time-based schedule was created.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 0 	- 	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The description of the time-based schedule.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the time-based schedule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the time-based schedule was last updated.
	//
	// example:
	//
	// 2020-01-01T01:01:01.001Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The trigger message of the time-based schedule.
	//
	// example:
	//
	// {"key": "value"}
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the time-based schedule.
	//
	// example:
	//
	// testScheduleId
	ScheduleId *string `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// The name of the time-based schedule.
	//
	// example:
	//
	// testScheduleName
	ScheduleName *string `json:"ScheduleName,omitempty" xml:"ScheduleName,omitempty"`
}

func (s UpdateScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleResponseBody) SetCreatedTime(v string) *UpdateScheduleResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetCronExpression(v string) *UpdateScheduleResponseBody {
	s.CronExpression = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetDescription(v string) *UpdateScheduleResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetEnable(v bool) *UpdateScheduleResponseBody {
	s.Enable = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetLastModifiedTime(v string) *UpdateScheduleResponseBody {
	s.LastModifiedTime = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetPayload(v string) *UpdateScheduleResponseBody {
	s.Payload = &v
	return s
}

func (s *UpdateScheduleResponseBody) SetRequestId(v string) *UpdateScheduleResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateScheduleResponse) SetStatusCode(v int32) *UpdateScheduleResponse {
	s.StatusCode = &v
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
	client.SignatureAlgorithm = tea.String("v2")
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

// Summary:
//
// Creates a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- The number of flows that each user can create is restricted by resources. For more information, see [Limits](https://help.aliyun.com/document_detail/122093.html). If you want to create more flows, submit a ticket.
//
// 	- At the user level, flows are distinguished by name. The name of a flow within one account must be unique.
//
// @param request - CreateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		body["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionMode)) {
		body["ExecutionMode"] = request.ExecutionMode
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalStorageLocation)) {
		body["ExternalStorageLocation"] = request.ExternalStorageLocation
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFlow"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- The number of flows that each user can create is restricted by resources. For more information, see [Limits](https://help.aliyun.com/document_detail/122093.html). If you want to create more flows, submit a ticket.
//
// 	- At the user level, flows are distinguished by name. The name of a flow within one account must be unique.
//
// @param request - CreateFlowRequest
//
// @return CreateFlowResponse
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

// Summary:
//
// Creates a time-based schedule.
//
// @param request - CreateScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduleResponse
func (client *Client) CreateScheduleWithOptions(request *CreateScheduleRequest, runtime *util.RuntimeOptions) (_result *CreateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SignatureVersion)) {
		query["SignatureVersion"] = request.SignatureVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CronExpression)) {
		body["CronExpression"] = request.CronExpression
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.Payload)) {
		body["Payload"] = request.Payload
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleName)) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSchedule"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a time-based schedule.
//
// @param request - CreateScheduleRequest
//
// @return CreateScheduleResponse
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

// Summary:
//
// Deletes an existing flow.
//
// Description:
//
// ## [](#)Usage notes
//
// A delete operation is asynchronous. If this operation is successful, the system returns a successful response. If an existing flow is pending to be deleted, a new flow of the same name will not be affected by the existing one. After you delete a flow, you cannot query its historical executions. All executions in progress will stop after their most recent steps are complete.
//
// @param request - DeleteFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFlow"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an existing flow.
//
// Description:
//
// ## [](#)Usage notes
//
// A delete operation is asynchronous. If this operation is successful, the system returns a successful response. If an existing flow is pending to be deleted, a new flow of the same name will not be affected by the existing one. After you delete a flow, you cannot query its historical executions. All executions in progress will stop after their most recent steps are complete.
//
// @param request - DeleteFlowRequest
//
// @return DeleteFlowResponse
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

// Summary:
//
// Deletes a time-based scheduling task.
//
// @param request - DeleteScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduleResponse
func (client *Client) DeleteScheduleWithOptions(request *DeleteScheduleRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleName)) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSchedule"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a time-based scheduling task.
//
// @param request - DeleteScheduleRequest
//
// @return DeleteScheduleResponse
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

// Summary:
//
// Queries an execution in a flow. The long polling mode is supported. The maximum waiting period for long polling depends on the value of the WaitTimeSeconds parameter.
//
// @param request - DescribeExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExecutionResponse
func (client *Client) DescribeExecutionWithOptions(request *DescribeExecutionRequest, runtime *util.RuntimeOptions) (_result *DescribeExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeExecution"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an execution in a flow. The long polling mode is supported. The maximum waiting period for long polling depends on the value of the WaitTimeSeconds parameter.
//
// @param request - DescribeExecutionRequest
//
// @return DescribeExecutionResponse
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

// Summary:
//
// Queries the information about a flow.
//
// @param request - DescribeFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowResponse
func (client *Client) DescribeFlowWithOptions(request *DescribeFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFlow"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a flow.
//
// @param request - DescribeFlowRequest
//
// @return DescribeFlowResponse
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

// Summary:
//
// Queries the detailed information about a time-based schedule.
//
// @param request - DescribeScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduleResponse
func (client *Client) DescribeScheduleWithOptions(request *DescribeScheduleRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSchedule"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a time-based schedule.
//
// @param request - DescribeScheduleRequest
//
// @return DescribeScheduleResponse
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

// Summary:
//
// Queries the details about each step in an execution process.
//
// @param request - GetExecutionHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutionHistoryResponse
func (client *Client) GetExecutionHistoryWithOptions(request *GetExecutionHistoryRequest, runtime *util.RuntimeOptions) (_result *GetExecutionHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExecutionHistory"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExecutionHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about each step in an execution process.
//
// @param request - GetExecutionHistoryRequest
//
// @return GetExecutionHistoryResponse
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

// Summary:
//
// Queries all historical executions of a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// After you delete a flow, you cannot query its historical executions, even if you create a flow of the same name.
//
// @param request - ListExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionsResponse
func (client *Client) ListExecutionsWithOptions(request *ListExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutions"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all historical executions of a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// After you delete a flow, you cannot query its historical executions, even if you create a flow of the same name.
//
// @param request - ListExecutionsRequest
//
// @return ListExecutionsResponse
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

// Summary:
//
// Queries a list of flows.
//
// @param request - ListFlowsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowsResponse
func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *util.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFlows"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of flows.
//
// @param request - ListFlowsRequest
//
// @return ListFlowsResponse
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

// Summary:
//
// Queries time-based schedules in a flow.
//
// @param request - ListSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchedulesResponse
func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, runtime *util.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSchedules"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries time-based schedules in a flow.
//
// @param request - ListSchedulesRequest
//
// @return ListSchedulesResponse
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

// Summary:
//
// Reports a failed task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskFailed is used to call back `pattern: waitForCallback` indicates that the current task fails to be executed.
//
// In the new service (CloudFlow), the task step that ReportTaskFailed is used to call back `TaskMode: WaitForCustomCallback` indicates that the current task fails to be executed.
//
// @param request - ReportTaskFailedRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportTaskFailedResponse
func (client *Client) ReportTaskFailedWithOptions(request *ReportTaskFailedRequest, runtime *util.RuntimeOptions) (_result *ReportTaskFailedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskToken)) {
		query["TaskToken"] = request.TaskToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cause)) {
		body["Cause"] = request.Cause
	}

	if !tea.BoolValue(util.IsUnset(request.Error)) {
		body["Error"] = request.Error
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportTaskFailed"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportTaskFailedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports a failed task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskFailed is used to call back `pattern: waitForCallback` indicates that the current task fails to be executed.
//
// In the new service (CloudFlow), the task step that ReportTaskFailed is used to call back `TaskMode: WaitForCustomCallback` indicates that the current task fails to be executed.
//
// @param request - ReportTaskFailedRequest
//
// @return ReportTaskFailedResponse
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

// Summary:
//
// Reports a successful task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskSucceeded is used to call back pattern: waitForCallback indicates that the current task is successfully executed.
//
// In the new service (CloudFlow), the task step that ReportTaskSucceeded is used to call back TaskMode: WaitForCustomCallback indicates that the current task is successfully executed.
//
// @param request - ReportTaskSucceededRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportTaskSucceededResponse
func (client *Client) ReportTaskSucceededWithOptions(request *ReportTaskSucceededRequest, runtime *util.RuntimeOptions) (_result *ReportTaskSucceededResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskToken)) {
		query["TaskToken"] = request.TaskToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Output)) {
		body["Output"] = request.Output
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportTaskSucceeded"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportTaskSucceededResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reports a successful task.
//
// Description:
//
// ## [](#)Usage notes
//
// In the previous service (Serverless Workflow), the task step that ReportTaskSucceeded is used to call back pattern: waitForCallback indicates that the current task is successfully executed.
//
// In the new service (CloudFlow), the task step that ReportTaskSucceeded is used to call back TaskMode: WaitForCustomCallback indicates that the current task is successfully executed.
//
// @param request - ReportTaskSucceededRequest
//
// @return ReportTaskSucceededResponse
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

// Summary:
//
// Starts the execution of a workflow.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- The flow is created. A flow only in standard mode is supported.
//
// 	- If you do not specify an execution, the system automatically generates an execution and starts the execution.
//
// 	- If an ongoing execution has the same name as that of the execution to be started, the system directly returns the ongoing execution.
//
// 	- If the ongoing execution with the same name has ended (succeeded or failed), `ExecutionAlreadyExists` is returned.
//
// 	- If no execution with the same name exists, the system starts a new execution.
//
// @param request - StartExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartExecutionResponse
func (client *Client) StartExecutionWithOptions(request *StartExecutionRequest, runtime *util.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallbackFnFTaskToken)) {
		body["CallbackFnFTaskToken"] = request.CallbackFnFTaskToken
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionName)) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartExecution"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts the execution of a workflow.
//
// Description:
//
// ## [](#)Usage notes
//
// 	- The flow is created. A flow only in standard mode is supported.
//
// 	- If you do not specify an execution, the system automatically generates an execution and starts the execution.
//
// 	- If an ongoing execution has the same name as that of the execution to be started, the system directly returns the ongoing execution.
//
// 	- If the ongoing execution with the same name has ended (succeeded or failed), `ExecutionAlreadyExists` is returned.
//
// 	- If no execution with the same name exists, the system starts a new execution.
//
// @param request - StartExecutionRequest
//
// @return StartExecutionResponse
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

// Summary:
//
// Synchronously starts an execution in a flow.
//
// Description:
//
//   Only flows of the express execution mode are supported.
//
// @param request - StartSyncExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSyncExecutionResponse
func (client *Client) StartSyncExecutionWithOptions(request *StartSyncExecutionRequest, runtime *util.RuntimeOptions) (_result *StartSyncExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionName)) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSyncExecution"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSyncExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronously starts an execution in a flow.
//
// Description:
//
//   Only flows of the express execution mode are supported.
//
// @param request - StartSyncExecutionRequest
//
// @return StartSyncExecutionResponse
func (client *Client) StartSyncExecution(request *StartSyncExecutionRequest) (_result *StartSyncExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSyncExecutionResponse{}
	_body, _err := client.StartSyncExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an execution that is in progress in a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// The flow must be in progress.
//
// @param request - StopExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopExecutionResponse
func (client *Client) StopExecutionWithOptions(request *StopExecutionRequest, runtime *util.RuntimeOptions) (_result *StopExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cause)) {
		body["Cause"] = request.Cause
	}

	if !tea.BoolValue(util.IsUnset(request.Error)) {
		body["Error"] = request.Error
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionName)) {
		body["ExecutionName"] = request.ExecutionName
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopExecution"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an execution that is in progress in a flow.
//
// Description:
//
// ## [](#)Usage notes
//
// The flow must be in progress.
//
// @param request - StopExecutionRequest
//
// @return StopExecutionResponse
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

// Summary:
//
// Updates a flow.
//
// @param request - UpdateFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlowWithOptions(request *UpdateFlowRequest, runtime *util.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Definition)) {
		body["Definition"] = request.Definition
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		body["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFlow"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a flow.
//
// @param request - UpdateFlowRequest
//
// @return UpdateFlowResponse
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

// Summary:
//
// Updates a time-based schedule.
//
// @param request - UpdateScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScheduleResponse
func (client *Client) UpdateScheduleWithOptions(request *UpdateScheduleRequest, runtime *util.RuntimeOptions) (_result *UpdateScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CronExpression)) {
		body["CronExpression"] = request.CronExpression
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		body["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FlowName)) {
		body["FlowName"] = request.FlowName
	}

	if !tea.BoolValue(util.IsUnset(request.Payload)) {
		body["Payload"] = request.Payload
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleName)) {
		body["ScheduleName"] = request.ScheduleName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSchedule"),
		Version:     tea.String("2019-03-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a time-based schedule.
//
// @param request - UpdateScheduleRequest
//
// @return UpdateScheduleResponse
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
