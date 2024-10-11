// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelSparkStatementRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// j202106071620hangzhou****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026****
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
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 341ADFA1-4A68-4707-9CA7-FA86523A5F14
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{\\"operable\\":\\"false\\",\\"max_cu\\":\\"8\\",\\"instance_name\\":\\"api-test-9\\",\\"min_cu\\":\\"0\\",\\"engine\\":\\"spark\\",\\"region\\":\\"cn-hangzhou\\",\\"app_name\\":\\"xx\\",\\"app_version_name\\":\\"xx\\"}"
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CU_PRESTO
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 600
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// An internal error occurred
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// example:
	//
	// 312F48C7-AFA0-46AA-B8F6-1B16B8808045
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// print(2+2)\\n
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// j202106071620hangzhou****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// sql
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
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
	// example:
	//
	// 2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0DC5B6F080E8-0DC5B6F080E8-0DC5B
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202105272322hangzhou5d64f1560000128-0001
	JobAttemptId *string `json:"JobAttemptId,omitempty" xml:"JobAttemptId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// j202105272322hangzhou5d64f1560000128
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// release-test
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// local:///opt/spark/jars/offline-sql.jar, main_file\\n+ exec /usr/bin/tini -s -- /jdk/jdk8/bin/java -cp \\"/opt/tools/exec-wrapper.jar:.:::/opt/spark/jars/*\\" com.aliyun.dla.spark.SparkJobWrapper /opt/spark/bin/spark-submit --conf spark.driver.host=172.16.6.205 --conf spark.ui.port=4040 --conf \\"spark.driver.extraJavaOptions=-Dlog4j.configuration=file:///opt/spark/log-conf/log4j.properties -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/dump.hprof -XX:OnOutOfMemoryError=\\"bash /opt/tools/oss-cp.sh /tmp/dump.hprof oss://dla-test-cn-hangzhou/spark-logs/release-test/j202105272322hangzhou5d64f1560000128-0001/driver/dump.hprof; bash /opt/tools/job-stop.sh\\" \\"
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 9CE8F271-F918-43B6-8F58-F9F1C2DCFDB8
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobAttemptLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202010271622hangzhouf742a4330000923
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCluster
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
	// example:
	//
	// 5F10AB6E-8984-4E32-B821-4E1512711B8C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2020-10-27 16:23:16
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1603786996000
	CreateTimeValue *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	// example:
	//
	// {\\"args\\":[\\"100\\"],\\"name\\":\\"SparkPi\\",\\"className\\":\\"org.apache.spark.examples.SparkPi\\",\\"conf\\":{\\"spark.driver.resourceSpec\\":\\"medium\\",\\"spark.executor.instances\\":5,\\"spark.executor.resourceSpec\\":\\"medium\\"},\\"file\\":\\"local:///tmp/spark-examples.jar\\"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// small
	DriverResourceSpec *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	// example:
	//
	// 1
	ExecutorInstances *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	// example:
	//
	// small
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	// example:
	//
	// j202010271622hangzhouf742a4330000923
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// SparkPi
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 202105251618hzslot9906b0b40000005-0001
	LastJobAttemptId *string `json:"LastJobAttemptId,omitempty" xml:"LastJobAttemptId,omitempty"`
	// example:
	//
	// https://dlaui-cn-hangzhou.aliyuncs.com/?token=xxx
	SparkUI *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2020-10-27 16:23:16
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// 1603786996000
	SubmitTimeValue *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	// example:
	//
	// 2020-10-27 16:23:16
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1603786996000
	UpdateTimeValue *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	// example:
	//
	// MyCluster
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202010271622hangzhouf742a4330000923
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCluster
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
	// example:
	//
	// INFO SparkContext: Running Spark version 2.x
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// EC031B4B-3E07-40D2-9BC5-xxxx
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202011031935hangzhouf742a4330003667
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCluster
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
	// example:
	//
	// 5C3F9ADD-2350-4C21-A3A1-A906F8DF8016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202104171040zjkb4cc449d0000000
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
	// example:
	//
	// 9AD49396-8F80-5511-9268-A6360B205136
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// killed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkSessionStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202106071620hangzhou00000000001
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StatementId *int32 `json:"StatementId,omitempty" xml:"StatementId,omitempty"`
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
	Data *GetSparkStatementResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C5DB60E0-6EB3-493C-8956-DA0D3A2EE656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// print(2+2)
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1623033200754
	CompletedTime *int64 `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// "{\\"status\\":\\"ok\\",\\"execution_count\\":1,\\"data\\":{\\"text/plain\\":\\"4\\"}}"
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 0.5
	Process *float32 `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 1623033200750
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202011031935hangzhouf742a4330003667
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCluster
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
	// example:
	//
	// deleted
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 3970ACB5-3E30-4A3A-B63A-1790FCC261F4
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *KillSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {   "createTime":"2020-12-28 09:00:00",   "jobName":"SparkPi" }
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySparkCluster
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// {   "createTime":"2020-12-28 09:00:00",   "jobName":"SparkPi" }
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySparkCluster
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// E18B0634-F003-486A-90B9-4AA6932DA3AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	JobList []*ListSparkJobResponseBodyDataResultJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4998
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 2020-10-27 17:51:05
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1603792267000
	CreateTimeValue *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	// example:
	//
	// {"args":["100"],"name":...}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// small
	DriverResourceSpec *string `json:"DriverResourceSpec,omitempty" xml:"DriverResourceSpec,omitempty"`
	// example:
	//
	// 100
	ExecutorInstances *string `json:"ExecutorInstances,omitempty" xml:"ExecutorInstances,omitempty"`
	// example:
	//
	// small
	ExecutorResourceSpec *string `json:"ExecutorResourceSpec,omitempty" xml:"ExecutorResourceSpec,omitempty"`
	// example:
	//
	// j202010271750hangzhouf742a4330000950
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// SparkPi
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// https://dlaui-cn-hangzhou.aliyuncs.com/?token=xxx
	SparkUI *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2020-10-27 17:51:05
	SubmitTime *string `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	// example:
	//
	// 1603792267000
	SubmitTimeValue *string `json:"SubmitTimeValue,omitempty" xml:"SubmitTimeValue,omitempty"`
	// example:
	//
	// 2020-10-27 17:51:05
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1603792267000
	UpdateTimeValue *string `json:"UpdateTimeValue,omitempty" xml:"UpdateTimeValue,omitempty"`
	// example:
	//
	// MyCluster
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// {"status":"success", "createTime":"2021-05-27 11:00:00", "timeZone":"Asia/Shanghai"}
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// j202105272322hangzhou5d64f1560000128
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// release-test
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// {"status":"success", "createTime":"2021-05-27 11:00:00", "timeZone":"Asia/Shanghai"}
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// j202105272322hangzhou5d64f1560000128
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// release-test
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// 4A839055-70C7-423D-A6F2-E714EA4502DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 2021-05-27 23:22:08
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1622128928000
	CreateTimeValue *string `json:"CreateTimeValue,omitempty" xml:"CreateTimeValue,omitempty"`
	// example:
	//
	// {"xxx":"xxx"}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 00:01:12
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// example:
	//
	// 72000
	DurationTimeValue *string `json:"DurationTimeValue,omitempty" xml:"DurationTimeValue,omitempty"`
	// example:
	//
	// j202105272322hangzhou5d64f1560000128-0001
	JobAttemptId *string `json:"JobAttemptId,omitempty" xml:"JobAttemptId,omitempty"`
	// example:
	//
	// j202105272322hangzhou5d64f1560000128
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// xxx
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// xxx
	SparkUI *string `json:"SparkUI,omitempty" xml:"SparkUI,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2021-05-27 23:23:20
	TerminatedTime *string `json:"TerminatedTime,omitempty" xml:"TerminatedTime,omitempty"`
	// example:
	//
	// 1622129000000
	TerminatedTimeValue *string `json:"TerminatedTimeValue,omitempty" xml:"TerminatedTimeValue,omitempty"`
	// example:
	//
	// release-test
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkJobAttemptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// j202106071620hangzhou00000000001
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
	// example:
	//
	// C2AB7692-B9EF-4A50-BCFF-0DC5B6F080E8
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
	// example:
	//
	// print(2 + 2)
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1623033200754
	CompletedTime *int64 `json:"CompletedTime,omitempty" xml:"CompletedTime,omitempty"`
	// example:
	//
	// 2
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {\\"status\\":\\"ok\\",\\"execution_count\\":1,\\"data\\":{\\"text/plain\\":\\"4\\"}}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 1
	Progress *float32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 1623033200750
	StartedTime *int64 `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSparkStatementsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// dla-3qntmdtbqklsbsmceu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// example:
	//
	// 600
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// An internal error occurred
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// example:
	//
	// 312F48C7-AFA0-46AA-B8F6-1B16B8808045
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// {     "name": "SparkPi",     "file": "local:///tmp/spark-examples.jar",     "className": "org.apache.spark.examples.SparkPi",     "args": [         "100"     ],     "conf": {         "spark.driver.resourceSpec": "medium",         "spark.executor.instances": 5,         "spark.executor.resourceSpec": "medium"     } }
	ConfigJson *string `json:"ConfigJson,omitempty" xml:"ConfigJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MyCluster
	VcName *string `json:"VcName,omitempty" xml:"VcName,omitempty"`
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
	// example:
	//
	// j202011032013hangzhouf742a4330003679
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// C039FD4D-2F3C-4556-AF09-864D3A6485B2
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// set spark.driver.resourceSpec=medium;set spark.executor.instances=5;  set spark.executor.resourceSpec=medium;  set spark.app.name=sparksqltest;  show databases;
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MySparkCluster
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
	// example:
	//
	// j202104261729hangzhou224ee5230000930
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// B570310A-8A49-461B-B81D-AE0B58D63DA4
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSparkSQLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// abc123
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
	Data *ValidateVirtualClusterNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C039FD4D-2F3C-4556-AF09-864D3A6485B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// true
	Legal *string `json:"Legal,omitempty" xml:"Legal,omitempty"`
	// example:
	//
	// null
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateVirtualClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// @param request - CancelSparkStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelSparkStatementResponse
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

// @param request - CancelSparkStatementRequest
//
// @return CancelSparkStatementResponse
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

// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
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

// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
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

// @param request - ExecuteSparkStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteSparkStatementResponse
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

// @param request - ExecuteSparkStatementRequest
//
// @return ExecuteSparkStatementResponse
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

// @param request - GetJobAttemptLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobAttemptLogResponse
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

// @param request - GetJobAttemptLogRequest
//
// @return GetJobAttemptLogResponse
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

// @param request - GetJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobDetailResponse
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

// @param request - GetJobDetailRequest
//
// @return GetJobDetailResponse
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

// @param request - GetJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobLogResponse
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

// @param request - GetJobLogRequest
//
// @return GetJobLogResponse
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

// @param request - GetJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobStatusResponse
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

// @param request - GetJobStatusRequest
//
// @return GetJobStatusResponse
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

// @param request - GetSparkSessionStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkSessionStateResponse
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

// @param request - GetSparkSessionStateRequest
//
// @return GetSparkSessionStateResponse
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

// @param request - GetSparkStatementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSparkStatementResponse
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

// @param request - GetSparkStatementRequest
//
// @return GetSparkStatementResponse
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

// @param request - KillSparkJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KillSparkJobResponse
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

// @param request - KillSparkJobRequest
//
// @return KillSparkJobResponse
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

// @param tmpReq - ListSparkJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkJobResponse
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

// @param request - ListSparkJobRequest
//
// @return ListSparkJobResponse
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

// @param tmpReq - ListSparkJobAttemptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkJobAttemptResponse
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

// @param request - ListSparkJobAttemptRequest
//
// @return ListSparkJobAttemptResponse
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

// @param request - ListSparkStatementsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSparkStatementsResponse
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

// @param request - ListSparkStatementsRequest
//
// @return ListSparkStatementsResponse
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

// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
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

// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
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

// @param request - SubmitSparkJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkJobResponse
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

// @param request - SubmitSparkJobRequest
//
// @return SubmitSparkJobResponse
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

// @param request - SubmitSparkSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSparkSQLResponse
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

// @param request - SubmitSparkSQLRequest
//
// @return SubmitSparkSQLResponse
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

// @param request - ValidateVirtualClusterNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateVirtualClusterNameResponse
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

// @param request - ValidateVirtualClusterNameRequest
//
// @return ValidateVirtualClusterNameResponse
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
