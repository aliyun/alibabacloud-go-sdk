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

type CancelSparkStatementRequest struct {
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	StatementId *string `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s CancelSparkStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelSparkStatementRequest) GoString() string {
	return s.String()
}

func (s *CancelSparkStatementRequest) SetJobId(v string) *CancelSparkStatementRequest {
	s.JobId = &v
	return s
}

func (s *CancelSparkStatementRequest) SetStatementId(v string) *CancelSparkStatementRequest {
	s.StatementId = &v
	return s
}

type CancelSparkStatementResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelSparkStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelSparkStatementResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSparkStatementResponseBody) SetData(v string) *CancelSparkStatementResponseBody {
	s.Data = &v
	return s
}

func (s *CancelSparkStatementResponseBody) SetRequestId(v string) *CancelSparkStatementResponseBody {
	s.RequestId = &v
	return s
}

type CancelSparkStatementResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelSparkStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelSparkStatementResponse) GoString() string {
	return s.String()
}

func (s *CancelSparkStatementResponse) SetHeaders(v map[string]*string) *CancelSparkStatementResponse {
	s.Headers = v
	return s
}

func (s *CancelSparkStatementResponse) SetStatusCode(v int32) *CancelSparkStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSparkStatementResponse) SetBody(v *CancelSparkStatementResponseBody) *CancelSparkStatementResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Component    *string `json:"Component,omitempty" xml:"Component,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetComponent(v string) *CreateInstanceRequest {
	s.Component = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

type CreateInstanceResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetErrorCode(v string) *CreateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorInfo(v string) *CreateInstanceResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v string) *CreateInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type ExecuteSparkStatementRequest struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Kind  *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
}

func (s ExecuteSparkStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSparkStatementRequest) GoString() string {
	return s.String()
}

func (s *ExecuteSparkStatementRequest) SetCode(v string) *ExecuteSparkStatementRequest {
	s.Code = &v
	return s
}

func (s *ExecuteSparkStatementRequest) SetJobId(v string) *ExecuteSparkStatementRequest {
	s.JobId = &v
	return s
}

func (s *ExecuteSparkStatementRequest) SetKind(v string) *ExecuteSparkStatementRequest {
	s.Kind = &v
	return s
}

type ExecuteSparkStatementResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteSparkStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSparkStatementResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteSparkStatementResponseBody) SetData(v string) *ExecuteSparkStatementResponseBody {
	s.Data = &v
	return s
}

func (s *ExecuteSparkStatementResponseBody) SetRequestId(v string) *ExecuteSparkStatementResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteSparkStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteSparkStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteSparkStatementResponse) GoString() string {
	return s.String()
}

func (s *ExecuteSparkStatementResponse) SetHeaders(v map[string]*string) *ExecuteSparkStatementResponse {
	s.Headers = v
	return s
}

func (s *ExecuteSparkStatementResponse) SetStatusCode(v int32) *ExecuteSparkStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteSparkStatementResponse) SetBody(v *ExecuteSparkStatementResponseBody) *ExecuteSparkStatementResponse {
	s.Body = v
	return s
}

type GetJobAttemptLogRequest struct {
	JobAttemptId *string `json:"JobAttemptId,omitempty" xml:"JobAttemptId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName       *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobAttemptLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobAttemptLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobAttemptLogRequest) SetJobAttemptId(v string) *GetJobAttemptLogRequest {
	s.JobAttemptId = &v
	return s
}

func (s *GetJobAttemptLogRequest) SetJobId(v string) *GetJobAttemptLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobAttemptLogRequest) SetVcName(v string) *GetJobAttemptLogRequest {
	s.VcName = &v
	return s
}

type GetJobAttemptLogResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobAttemptLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobAttemptLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobAttemptLogResponseBody) SetData(v string) *GetJobAttemptLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetJobAttemptLogResponseBody) SetRequestId(v string) *GetJobAttemptLogResponseBody {
	s.RequestId = &v
	return s
}

type GetJobAttemptLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobAttemptLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobAttemptLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobAttemptLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobAttemptLogResponse) SetHeaders(v map[string]*string) *GetJobAttemptLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobAttemptLogResponse) SetStatusCode(v int32) *GetJobAttemptLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobAttemptLogResponse) SetBody(v *GetJobAttemptLogResponseBody) *GetJobAttemptLogResponse {
	s.Body = v
	return s
}

type GetJobDetailRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailRequest) GoString() string {
	return s.String()
}

func (s *GetJobDetailRequest) SetJobId(v string) *GetJobDetailRequest {
	s.JobId = &v
	return s
}

func (s *GetJobDetailRequest) SetVcName(v string) *GetJobDetailRequest {
	s.VcName = &v
	return s
}

type GetJobDetailResponseBody struct {
	JobDetail *GetJobDetailResponseBodyJobDetail `json:"JobDetail,omitempty" xml:"JobDetail,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBody) SetJobDetail(v *GetJobDetailResponseBodyJobDetail) *GetJobDetailResponseBody {
	s.JobDetail = v
	return s
}

func (s *GetJobDetailResponseBody) SetRequestId(v string) *GetJobDetailResponseBody {
	s.RequestId = &v
	return s
}

type GetJobDetailResponseBodyJobDetail struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	LastJobAttemptId     *string `json:"LastJobAttemptId,omitempty" xml:"LastJobAttemptId,omitempty"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobDetailResponseBodyJobDetail) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponseBodyJobDetail) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponseBodyJobDetail) SetCreateTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.CreateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetCreateTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.CreateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetDetail(v string) *GetJobDetailResponseBodyJobDetail {
	s.Detail = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetDriverResourceSpec(v string) *GetJobDetailResponseBodyJobDetail {
	s.DriverResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetExecutorInstances(v string) *GetJobDetailResponseBodyJobDetail {
	s.ExecutorInstances = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetExecutorResourceSpec(v string) *GetJobDetailResponseBodyJobDetail {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetJobId(v string) *GetJobDetailResponseBodyJobDetail {
	s.JobId = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetJobName(v string) *GetJobDetailResponseBodyJobDetail {
	s.JobName = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetLastJobAttemptId(v string) *GetJobDetailResponseBodyJobDetail {
	s.LastJobAttemptId = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSparkUI(v string) *GetJobDetailResponseBodyJobDetail {
	s.SparkUI = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetStatus(v string) *GetJobDetailResponseBodyJobDetail {
	s.Status = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSubmitTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.SubmitTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetSubmitTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.SubmitTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetUpdateTime(v string) *GetJobDetailResponseBodyJobDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetUpdateTimeValue(v string) *GetJobDetailResponseBodyJobDetail {
	s.UpdateTimeValue = &v
	return s
}

func (s *GetJobDetailResponseBodyJobDetail) SetVcName(v string) *GetJobDetailResponseBodyJobDetail {
	s.VcName = &v
	return s
}

type GetJobDetailResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobDetailResponse) GoString() string {
	return s.String()
}

func (s *GetJobDetailResponse) SetHeaders(v map[string]*string) *GetJobDetailResponse {
	s.Headers = v
	return s
}

func (s *GetJobDetailResponse) SetStatusCode(v int32) *GetJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobDetailResponse) SetBody(v *GetJobDetailResponseBody) *GetJobDetailResponse {
	s.Body = v
	return s
}

type GetJobLogRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogRequest) GoString() string {
	return s.String()
}

func (s *GetJobLogRequest) SetJobId(v string) *GetJobLogRequest {
	s.JobId = &v
	return s
}

func (s *GetJobLogRequest) SetVcName(v string) *GetJobLogRequest {
	s.VcName = &v
	return s
}

type GetJobLogResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobLogResponseBody) SetData(v string) *GetJobLogResponseBody {
	s.Data = &v
	return s
}

func (s *GetJobLogResponseBody) SetRequestId(v string) *GetJobLogResponseBody {
	s.RequestId = &v
	return s
}

type GetJobLogResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobLogResponse) GoString() string {
	return s.String()
}

func (s *GetJobLogResponse) SetHeaders(v map[string]*string) *GetJobLogResponse {
	s.Headers = v
	return s
}

func (s *GetJobLogResponse) SetStatusCode(v int32) *GetJobLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobLogResponse) SetBody(v *GetJobLogResponseBody) *GetJobLogResponse {
	s.Body = v
	return s
}

type GetJobStatusRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s GetJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetJobStatusRequest) SetJobId(v string) *GetJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetJobStatusRequest) SetVcName(v string) *GetJobStatusRequest {
	s.VcName = &v
	return s
}

type GetJobStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponseBody) SetRequestId(v string) *GetJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobStatusResponseBody) SetStatus(v string) *GetJobStatusResponseBody {
	s.Status = &v
	return s
}

type GetJobStatusResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetJobStatusResponse) SetHeaders(v map[string]*string) *GetJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetJobStatusResponse) SetStatusCode(v int32) *GetJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobStatusResponse) SetBody(v *GetJobStatusResponseBody) *GetJobStatusResponse {
	s.Body = v
	return s
}

type GetSparkSessionStateRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetSparkSessionStateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSessionStateRequest) GoString() string {
	return s.String()
}

func (s *GetSparkSessionStateRequest) SetJobId(v string) *GetSparkSessionStateRequest {
	s.JobId = &v
	return s
}

type GetSparkSessionStateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSparkSessionStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSessionStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkSessionStateResponseBody) SetRequestId(v string) *GetSparkSessionStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkSessionStateResponseBody) SetState(v string) *GetSparkSessionStateResponseBody {
	s.State = &v
	return s
}

type GetSparkSessionStateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSparkSessionStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSparkSessionStateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkSessionStateResponse) GoString() string {
	return s.String()
}

func (s *GetSparkSessionStateResponse) SetHeaders(v map[string]*string) *GetSparkSessionStateResponse {
	s.Headers = v
	return s
}

func (s *GetSparkSessionStateResponse) SetStatusCode(v int32) *GetSparkSessionStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkSessionStateResponse) SetBody(v *GetSparkSessionStateResponseBody) *GetSparkSessionStateResponse {
	s.Body = v
	return s
}

type GetSparkStatementRequest struct {
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	StatementId *int32  `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
}

func (s GetSparkStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSparkStatementRequest) GoString() string {
	return s.String()
}

func (s *GetSparkStatementRequest) SetJobId(v string) *GetSparkStatementRequest {
	s.JobId = &v
	return s
}

func (s *GetSparkStatementRequest) SetStatementId(v int32) *GetSparkStatementRequest {
	s.StatementId = &v
	return s
}

type GetSparkStatementResponseBody struct {
	Data      *GetSparkStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkStatementResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSparkStatementResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkStatementResponseBody) SetData(v *GetSparkStatementResponseBodyData) *GetSparkStatementResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkStatementResponseBody) SetRequestId(v string) *GetSparkStatementResponseBody {
	s.RequestId = &v
	return s
}

type GetSparkStatementResponseBodyData struct {
	Code          *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	CompletedTime *int64   `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	Id            *int32   `json:"Id,omitempty" xml:"Id,omitempty"`
	Output        *string  `json:"Output,omitempty" xml:"Output,omitempty"`
	Process       *float32 `json:"Process,omitempty" xml:"Process,omitempty"`
	StartedTime   *int64   `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	State         *string  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetSparkStatementResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSparkStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkStatementResponseBodyData) SetCode(v string) *GetSparkStatementResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetCompletedTime(v int64) *GetSparkStatementResponseBodyData {
	s.CompletedTime = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetId(v int32) *GetSparkStatementResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetOutput(v string) *GetSparkStatementResponseBodyData {
	s.Output = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetProcess(v float32) *GetSparkStatementResponseBodyData {
	s.Process = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetStartedTime(v int64) *GetSparkStatementResponseBodyData {
	s.StartedTime = &v
	return s
}

func (s *GetSparkStatementResponseBodyData) SetState(v string) *GetSparkStatementResponseBodyData {
	s.State = &v
	return s
}

type GetSparkStatementResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSparkStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSparkStatementResponse) GoString() string {
	return s.String()
}

func (s *GetSparkStatementResponse) SetHeaders(v map[string]*string) *GetSparkStatementResponse {
	s.Headers = v
	return s
}

func (s *GetSparkStatementResponse) SetStatusCode(v int32) *GetSparkStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkStatementResponse) SetBody(v *GetSparkStatementResponseBody) *GetSparkStatementResponse {
	s.Body = v
	return s
}

type KillSparkJobRequest struct {
	JobId  *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s KillSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobRequest) GoString() string {
	return s.String()
}

func (s *KillSparkJobRequest) SetJobId(v string) *KillSparkJobRequest {
	s.JobId = &v
	return s
}

func (s *KillSparkJobRequest) SetVcName(v string) *KillSparkJobRequest {
	s.VcName = &v
	return s
}

type KillSparkJobResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkJobResponseBody) SetData(v string) *KillSparkJobResponseBody {
	s.Data = &v
	return s
}

func (s *KillSparkJobResponseBody) SetRequestId(v string) *KillSparkJobResponseBody {
	s.RequestId = &v
	return s
}

type KillSparkJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *KillSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s KillSparkJobResponse) GoString() string {
	return s.String()
}

func (s *KillSparkJobResponse) SetHeaders(v map[string]*string) *KillSparkJobResponse {
	s.Headers = v
	return s
}

func (s *KillSparkJobResponse) SetStatusCode(v int32) *KillSparkJobResponse {
	s.StatusCode = &v
	return s
}

func (s *KillSparkJobResponse) SetBody(v *KillSparkJobResponseBody) *KillSparkJobResponse {
	s.Body = v
	return s
}

type ListSparkJobRequest struct {
	Condition  map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	PageNumber *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VcName     *string                `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobRequest) SetCondition(v map[string]interface{}) *ListSparkJobRequest {
	s.Condition = v
	return s
}

func (s *ListSparkJobRequest) SetPageNumber(v int32) *ListSparkJobRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobRequest) SetPageSize(v int32) *ListSparkJobRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobRequest) SetVcName(v string) *ListSparkJobRequest {
	s.VcName = &v
	return s
}

type ListSparkJobShrinkRequest struct {
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VcName          *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobShrinkRequest) SetConditionShrink(v string) *ListSparkJobShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *ListSparkJobShrinkRequest) SetPageNumber(v int32) *ListSparkJobShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobShrinkRequest) SetPageSize(v int32) *ListSparkJobShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobShrinkRequest) SetVcName(v string) *ListSparkJobShrinkRequest {
	s.VcName = &v
	return s
}

type ListSparkJobResponseBody struct {
	DataResult *ListSparkJobResponseBodyDataResult `json:"DataResult,omitempty" xml:"DataResult,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBody) SetDataResult(v *ListSparkJobResponseBodyDataResult) *ListSparkJobResponseBody {
	s.DataResult = v
	return s
}

func (s *ListSparkJobResponseBody) SetRequestId(v string) *ListSparkJobResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkJobResponseBodyDataResult struct {
	JobList    []*ListSparkJobResponseBodyDataResultJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	PageNumber *string                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *string                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkJobResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBodyDataResult) SetJobList(v []*ListSparkJobResponseBodyDataResultJobList) *ListSparkJobResponseBodyDataResult {
	s.JobList = v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetPageNumber(v string) *ListSparkJobResponseBodyDataResult {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetPageSize(v string) *ListSparkJobResponseBodyDataResult {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResult) SetTotalCount(v string) *ListSparkJobResponseBodyDataResult {
	s.TotalCount = &v
	return s
}

type ListSparkJobResponseBodyDataResultJobList struct {
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeValue      *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	Detail               *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DriverResourceSpec   *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	ExecutorInstances    *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName              *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	SparkUI              *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SubmitTime           *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	SubmitTimeValue      *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	UpdateTime           *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimeValue      *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	VcName               *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobResponseBodyDataResultJobList) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponseBodyDataResultJobList) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetCreateTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.CreateTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetCreateTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.CreateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetDetail(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.Detail = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetDriverResourceSpec(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.DriverResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetExecutorInstances(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.ExecutorInstances = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetExecutorResourceSpec(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.ExecutorResourceSpec = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetJobId(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.JobId = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetJobName(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.JobName = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSparkUI(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SparkUI = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetStatus(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.Status = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSubmitTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SubmitTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetSubmitTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.SubmitTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetUpdateTime(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.UpdateTime = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetUpdateTimeValue(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.UpdateTimeValue = &v
	return s
}

func (s *ListSparkJobResponseBodyDataResultJobList) SetVcName(v string) *ListSparkJobResponseBodyDataResultJobList {
	s.VcName = &v
	return s
}

type ListSparkJobResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobResponse) GoString() string {
	return s.String()
}

func (s *ListSparkJobResponse) SetHeaders(v map[string]*string) *ListSparkJobResponse {
	s.Headers = v
	return s
}

func (s *ListSparkJobResponse) SetStatusCode(v int32) *ListSparkJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkJobResponse) SetBody(v *ListSparkJobResponseBody) *ListSparkJobResponse {
	s.Body = v
	return s
}

type ListSparkJobAttemptRequest struct {
	Condition  map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	JobId      *string                `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VcName     *string                `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobAttemptRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptRequest) SetCondition(v map[string]interface{}) *ListSparkJobAttemptRequest {
	s.Condition = v
	return s
}

func (s *ListSparkJobAttemptRequest) SetJobId(v string) *ListSparkJobAttemptRequest {
	s.JobId = &v
	return s
}

func (s *ListSparkJobAttemptRequest) SetPageNumber(v int32) *ListSparkJobAttemptRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobAttemptRequest) SetPageSize(v int32) *ListSparkJobAttemptRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobAttemptRequest) SetVcName(v string) *ListSparkJobAttemptRequest {
	s.VcName = &v
	return s
}

type ListSparkJobAttemptShrinkRequest struct {
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	JobId           *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	VcName          *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobAttemptShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptShrinkRequest) SetConditionShrink(v string) *ListSparkJobAttemptShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *ListSparkJobAttemptShrinkRequest) SetJobId(v string) *ListSparkJobAttemptShrinkRequest {
	s.JobId = &v
	return s
}

func (s *ListSparkJobAttemptShrinkRequest) SetPageNumber(v int32) *ListSparkJobAttemptShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobAttemptShrinkRequest) SetPageSize(v int32) *ListSparkJobAttemptShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobAttemptShrinkRequest) SetVcName(v string) *ListSparkJobAttemptShrinkRequest {
	s.VcName = &v
	return s
}

type ListSparkJobAttemptResponseBody struct {
	DataResult *ListSparkJobAttemptResponseBodyDataResult `json:"DataResult,omitempty" xml:"DataResult,omitempty" type:"Struct"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSparkJobAttemptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptResponseBody) SetDataResult(v *ListSparkJobAttemptResponseBodyDataResult) *ListSparkJobAttemptResponseBody {
	s.DataResult = v
	return s
}

func (s *ListSparkJobAttemptResponseBody) SetRequestId(v string) *ListSparkJobAttemptResponseBody {
	s.RequestId = &v
	return s
}

type ListSparkJobAttemptResponseBodyDataResult struct {
	JobAttemptList []*ListSparkJobAttemptResponseBodyDataResultJobAttemptList `json:"JobAttemptList,omitempty" xml:"JobAttemptList,omitempty" type:"Repeated"`
	PageNumber     *string                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *string                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *string                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSparkJobAttemptResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptResponseBodyDataResult) SetJobAttemptList(v []*ListSparkJobAttemptResponseBodyDataResultJobAttemptList) *ListSparkJobAttemptResponseBodyDataResult {
	s.JobAttemptList = v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResult) SetPageNumber(v string) *ListSparkJobAttemptResponseBodyDataResult {
	s.PageNumber = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResult) SetPageSize(v string) *ListSparkJobAttemptResponseBodyDataResult {
	s.PageSize = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResult) SetTotalCount(v string) *ListSparkJobAttemptResponseBodyDataResult {
	s.TotalCount = &v
	return s
}

type ListSparkJobAttemptResponseBodyDataResultJobAttemptList struct {
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeValue     *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	Detail              *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	DurationTime        *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	DurationTimeValue   *string `json:"DurationTimeValue,omitempty" xml:"DurationTimeValue,omitempty"`
	JobAttemptId        *string `json:"JobAttemptId,omitempty" xml:"JobAttemptId,omitempty"`
	JobId               *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobName             *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	SparkUI             *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminatedTime      *string `json:"TerminatedTime,omitempty" xml:"TerminatedTime,omitempty"`
	TerminatedTimeValue *string `json:"TerminatedTimeValue,omitempty" xml:"TerminatedTimeValue,omitempty"`
	VcName              *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ListSparkJobAttemptResponseBodyDataResultJobAttemptList) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptResponseBodyDataResultJobAttemptList) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetCreateTime(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.CreateTime = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetCreateTimeValue(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.CreateTimeValue = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetDetail(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.Detail = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetDurationTime(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.DurationTime = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetDurationTimeValue(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.DurationTimeValue = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetJobAttemptId(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.JobAttemptId = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetJobId(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.JobId = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetJobName(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.JobName = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetSparkUI(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.SparkUI = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetStatus(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.Status = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetTerminatedTime(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.TerminatedTime = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetTerminatedTimeValue(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.TerminatedTimeValue = &v
	return s
}

func (s *ListSparkJobAttemptResponseBodyDataResultJobAttemptList) SetVcName(v string) *ListSparkJobAttemptResponseBodyDataResultJobAttemptList {
	s.VcName = &v
	return s
}

type ListSparkJobAttemptResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSparkJobAttemptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSparkJobAttemptResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkJobAttemptResponse) GoString() string {
	return s.String()
}

func (s *ListSparkJobAttemptResponse) SetHeaders(v map[string]*string) *ListSparkJobAttemptResponse {
	s.Headers = v
	return s
}

func (s *ListSparkJobAttemptResponse) SetStatusCode(v int32) *ListSparkJobAttemptResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkJobAttemptResponse) SetBody(v *ListSparkJobAttemptResponseBody) *ListSparkJobAttemptResponse {
	s.Body = v
	return s
}

type ListSparkStatementsRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListSparkStatementsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSparkStatementsRequest) GoString() string {
	return s.String()
}

func (s *ListSparkStatementsRequest) SetJobId(v string) *ListSparkStatementsRequest {
	s.JobId = &v
	return s
}

type ListSparkStatementsResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statements []*ListSparkStatementsResponseBodyStatements `json:"Statements,omitempty" xml:"Statements,omitempty" type:"Repeated"`
}

func (s ListSparkStatementsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSparkStatementsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSparkStatementsResponseBody) SetRequestId(v string) *ListSparkStatementsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSparkStatementsResponseBody) SetStatements(v []*ListSparkStatementsResponseBodyStatements) *ListSparkStatementsResponseBody {
	s.Statements = v
	return s
}

type ListSparkStatementsResponseBodyStatements struct {
	Code          *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	CompletedTime *int64   `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	Id            *int32   `json:"Id,omitempty" xml:"Id,omitempty"`
	Output        *string  `json:"Output,omitempty" xml:"Output,omitempty"`
	Progress      *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	StartedTime   *int64   `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	State         *string  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListSparkStatementsResponseBodyStatements) String() string {
	return tea.Prettify(s)
}

func (s ListSparkStatementsResponseBodyStatements) GoString() string {
	return s.String()
}

func (s *ListSparkStatementsResponseBodyStatements) SetCode(v string) *ListSparkStatementsResponseBodyStatements {
	s.Code = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetCompletedTime(v int64) *ListSparkStatementsResponseBodyStatements {
	s.CompletedTime = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetId(v int32) *ListSparkStatementsResponseBodyStatements {
	s.Id = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetOutput(v string) *ListSparkStatementsResponseBodyStatements {
	s.Output = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetProgress(v float32) *ListSparkStatementsResponseBodyStatements {
	s.Progress = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetStartedTime(v int64) *ListSparkStatementsResponseBodyStatements {
	s.StartedTime = &v
	return s
}

func (s *ListSparkStatementsResponseBodyStatements) SetState(v string) *ListSparkStatementsResponseBodyStatements {
	s.State = &v
	return s
}

type ListSparkStatementsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSparkStatementsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSparkStatementsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSparkStatementsResponse) GoString() string {
	return s.String()
}

func (s *ListSparkStatementsResponse) SetHeaders(v map[string]*string) *ListSparkStatementsResponse {
	s.Headers = v
	return s
}

func (s *ListSparkStatementsResponse) SetStatusCode(v int32) *ListSparkStatementsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSparkStatementsResponse) SetBody(v *ListSparkStatementsResponseBody) *ListSparkStatementsResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetRegionId(v string) *ReleaseInstanceRequest {
	s.RegionId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetErrorCode(v string) *ReleaseInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetErrorInfo(v string) *ReleaseInstanceResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetResult(v string) *ReleaseInstanceResponseBody {
	s.Result = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetStatusCode(v int32) *ReleaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type SubmitSparkJobRequest struct {
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	VcName     *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s SubmitSparkJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobRequest) SetConfigJson(v string) *SubmitSparkJobRequest {
	s.ConfigJson = &v
	return s
}

func (s *SubmitSparkJobRequest) SetVcName(v string) *SubmitSparkJobRequest {
	s.VcName = &v
	return s
}

type SubmitSparkJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobResponseBody) SetJobId(v string) *SubmitSparkJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSparkJobResponseBody) SetRequestId(v string) *SubmitSparkJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSparkJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSparkJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkJobResponse) SetHeaders(v map[string]*string) *SubmitSparkJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkJobResponse) SetStatusCode(v int32) *SubmitSparkJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkJobResponse) SetBody(v *SubmitSparkJobResponseBody) *SubmitSparkJobResponse {
	s.Body = v
	return s
}

type SubmitSparkSQLRequest struct {
	Sql    *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s SubmitSparkSQLRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkSQLRequest) GoString() string {
	return s.String()
}

func (s *SubmitSparkSQLRequest) SetSql(v string) *SubmitSparkSQLRequest {
	s.Sql = &v
	return s
}

func (s *SubmitSparkSQLRequest) SetVcName(v string) *SubmitSparkSQLRequest {
	s.VcName = &v
	return s
}

type SubmitSparkSQLResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkSQLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkSQLResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkSQLResponseBody) SetJobId(v string) *SubmitSparkSQLResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitSparkSQLResponseBody) SetRequestId(v string) *SubmitSparkSQLResponseBody {
	s.RequestId = &v
	return s
}

type SubmitSparkSQLResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitSparkSQLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitSparkSQLResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitSparkSQLResponse) GoString() string {
	return s.String()
}

func (s *SubmitSparkSQLResponse) SetHeaders(v map[string]*string) *SubmitSparkSQLResponse {
	s.Headers = v
	return s
}

func (s *SubmitSparkSQLResponse) SetStatusCode(v int32) *SubmitSparkSQLResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSparkSQLResponse) SetBody(v *SubmitSparkSQLResponseBody) *SubmitSparkSQLResponse {
	s.Body = v
	return s
}

type ValidateVirtualClusterNameRequest struct {
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
}

func (s ValidateVirtualClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameRequest) SetVcName(v string) *ValidateVirtualClusterNameRequest {
	s.VcName = &v
	return s
}

type ValidateVirtualClusterNameResponseBody struct {
	Data      *ValidateVirtualClusterNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateVirtualClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponseBody) SetData(v *ValidateVirtualClusterNameResponseBodyData) *ValidateVirtualClusterNameResponseBody {
	s.Data = v
	return s
}

func (s *ValidateVirtualClusterNameResponseBody) SetRequestId(v string) *ValidateVirtualClusterNameResponseBody {
	s.RequestId = &v
	return s
}

type ValidateVirtualClusterNameResponseBodyData struct {
	Legal   *string `json:"Legal,omitempty" xml:"Legal,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ValidateVirtualClusterNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponseBodyData) SetLegal(v string) *ValidateVirtualClusterNameResponseBodyData {
	s.Legal = &v
	return s
}

func (s *ValidateVirtualClusterNameResponseBodyData) SetMessage(v string) *ValidateVirtualClusterNameResponseBodyData {
	s.Message = &v
	return s
}

type ValidateVirtualClusterNameResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateVirtualClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateVirtualClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateVirtualClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ValidateVirtualClusterNameResponse) SetHeaders(v map[string]*string) *ValidateVirtualClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ValidateVirtualClusterNameResponse) SetStatusCode(v int32) *ValidateVirtualClusterNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateVirtualClusterNameResponse) SetBody(v *ValidateVirtualClusterNameResponseBody) *ValidateVirtualClusterNameResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-beijing":                  tea.String("openanalytics.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("openanalytics.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("openanalytics.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":                 tea.String("openanalytics.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("openanalytics.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":                 tea.String("openanalytics.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":              tea.String("openanalytics.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("datalakeanalytics.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("openanalytics.ap-southeast-3.aliyuncs.com"),
		"ap-northeast-1":              tea.String("datalakeanalytics.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("openanalytics.eu-west-1.aliyuncs.com"),
		"us-west-1":                   tea.String("openanalytics.us-west-1.aliyuncs.com"),
		"us-east-1":                   tea.String("datalakeanalytics.us-east-1.aliyuncs.com"),
		"eu-central-1":                tea.String("datalakeanalytics.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("openanalytics.ap-south-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              tea.String("openanalytics.ap-southeast-5.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-chengdu":                  tea.String("openanalytics.aliyuncs.com"),
		"cn-edge-1":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-fujian":                   tea.String("openanalytics.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("openanalytics.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("openanalytics.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("openanalytics.aliyuncs.com"),
		"cn-huhehaote":                tea.String("openanalytics.cn-huhehaote.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("openanalytics.aliyuncs.com"),
		"cn-qingdao":                  tea.String("openanalytics.cn-qingdao.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("openanalytics.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-wuhan":                    tea.String("openanalytics.aliyuncs.com"),
		"cn-yushanfang":               tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("openanalytics.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("openanalytics.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("openanalytics.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   tea.String("openanalytics.me-east-1.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("openanalytics.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("openanalytics-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelSparkStatementWithOptions(request *CancelSparkStatementRequest, runtime *util.RuntimeOptions) (_result *CancelSparkStatementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.StatementId)) {
		body["StatementId"] = request.StatementId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelSparkStatement"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelSparkStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelSparkStatement(request *CancelSparkStatementRequest) (_result *CancelSparkStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelSparkStatementResponse{}
	_body, _err := client.CancelSparkStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Component)) {
		body["Component"] = request.Component
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteSparkStatementWithOptions(request *ExecuteSparkStatementRequest, runtime *util.RuntimeOptions) (_result *ExecuteSparkStatementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Kind)) {
		body["Kind"] = request.Kind
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteSparkStatement"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteSparkStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteSparkStatement(request *ExecuteSparkStatementRequest) (_result *ExecuteSparkStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteSparkStatementResponse{}
	_body, _err := client.ExecuteSparkStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobAttemptLogWithOptions(request *GetJobAttemptLogRequest, runtime *util.RuntimeOptions) (_result *GetJobAttemptLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobAttemptId)) {
		body["JobAttemptId"] = request.JobAttemptId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobAttemptLog"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobAttemptLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobAttemptLog(request *GetJobAttemptLogRequest) (_result *GetJobAttemptLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobAttemptLogResponse{}
	_body, _err := client.GetJobAttemptLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobDetailWithOptions(request *GetJobDetailRequest, runtime *util.RuntimeOptions) (_result *GetJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobDetail"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobDetail(request *GetJobDetailRequest) (_result *GetJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobDetailResponse{}
	_body, _err := client.GetJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobLogWithOptions(request *GetJobLogRequest, runtime *util.RuntimeOptions) (_result *GetJobLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobLog"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobLog(request *GetJobLogRequest) (_result *GetJobLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobLogResponse{}
	_body, _err := client.GetJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobStatusWithOptions(request *GetJobStatusRequest, runtime *util.RuntimeOptions) (_result *GetJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobStatus"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobStatus(request *GetJobStatusRequest) (_result *GetJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJobStatusResponse{}
	_body, _err := client.GetJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkSessionStateWithOptions(request *GetSparkSessionStateRequest, runtime *util.RuntimeOptions) (_result *GetSparkSessionStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkSessionState"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkSessionStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkSessionState(request *GetSparkSessionStateRequest) (_result *GetSparkSessionStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkSessionStateResponse{}
	_body, _err := client.GetSparkSessionStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSparkStatementWithOptions(request *GetSparkStatementRequest, runtime *util.RuntimeOptions) (_result *GetSparkStatementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.StatementId)) {
		body["StatementId"] = request.StatementId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSparkStatement"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSparkStatementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSparkStatement(request *GetSparkStatementRequest) (_result *GetSparkStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSparkStatementResponse{}
	_body, _err := client.GetSparkStatementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillSparkJobWithOptions(request *KillSparkJobRequest, runtime *util.RuntimeOptions) (_result *KillSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("KillSparkJob"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &KillSparkJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillSparkJob(request *KillSparkJobRequest) (_result *KillSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillSparkJobResponse{}
	_body, _err := client.KillSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkJobWithOptions(tmpReq *ListSparkJobRequest, runtime *util.RuntimeOptions) (_result *ListSparkJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSparkJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Condition)) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, tea.String("Condition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionShrink)) {
		query["Condition"] = request.ConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		query["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkJob"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkJob(request *ListSparkJobRequest) (_result *ListSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkJobResponse{}
	_body, _err := client.ListSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkJobAttemptWithOptions(tmpReq *ListSparkJobAttemptRequest, runtime *util.RuntimeOptions) (_result *ListSparkJobAttemptResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSparkJobAttemptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Condition)) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, tea.String("Condition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConditionShrink)) {
		query["Condition"] = request.ConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		query["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkJobAttempt"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkJobAttemptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkJobAttempt(request *ListSparkJobAttemptRequest) (_result *ListSparkJobAttemptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkJobAttemptResponse{}
	_body, _err := client.ListSparkJobAttemptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSparkStatementsWithOptions(request *ListSparkStatementsRequest, runtime *util.RuntimeOptions) (_result *ListSparkStatementsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSparkStatements"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSparkStatementsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSparkStatements(request *ListSparkStatementsRequest) (_result *ListSparkStatementsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSparkStatementsResponse{}
	_body, _err := client.ListSparkStatementsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstance"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkJobWithOptions(request *SubmitSparkJobRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigJson)) {
		body["ConfigJson"] = request.ConfigJson
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSparkJob"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkJob(request *SubmitSparkJobRequest) (_result *SubmitSparkJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkJobResponse{}
	_body, _err := client.SubmitSparkJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitSparkSQLWithOptions(request *SubmitSparkSQLRequest, runtime *util.RuntimeOptions) (_result *SubmitSparkSQLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["Sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitSparkSQL"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitSparkSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitSparkSQL(request *SubmitSparkSQLRequest) (_result *SubmitSparkSQLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitSparkSQLResponse{}
	_body, _err := client.SubmitSparkSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateVirtualClusterNameWithOptions(request *ValidateVirtualClusterNameRequest, runtime *util.RuntimeOptions) (_result *ValidateVirtualClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VcName)) {
		body["VcName"] = request.VcName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateVirtualClusterName"),
		Version:     tea.String("2018-06-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateVirtualClusterName(request *ValidateVirtualClusterNameRequest) (_result *ValidateVirtualClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateVirtualClusterNameResponse{}
	_body, _err := client.ValidateVirtualClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
