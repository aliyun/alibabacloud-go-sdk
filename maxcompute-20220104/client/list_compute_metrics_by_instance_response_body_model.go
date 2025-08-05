// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsByInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListComputeMetricsByInstanceResponseBodyData) *ListComputeMetricsByInstanceResponseBody
	GetData() *ListComputeMetricsByInstanceResponseBodyData
	SetErrorCode(v string) *ListComputeMetricsByInstanceResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListComputeMetricsByInstanceResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListComputeMetricsByInstanceResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListComputeMetricsByInstanceResponseBody
	GetRequestId() *string
}

type ListComputeMetricsByInstanceResponseBody struct {
	// The data returned.
	Data *ListComputeMetricsByInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc059b717363029839908920ea631
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBody) GetData() *ListComputeMetricsByInstanceResponseBodyData {
	return s.Data
}

func (s *ListComputeMetricsByInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListComputeMetricsByInstanceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListComputeMetricsByInstanceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListComputeMetricsByInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeMetricsByInstanceResponseBody) SetData(v *ListComputeMetricsByInstanceResponseBodyData) *ListComputeMetricsByInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorCode(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorMsg(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetHttpCode(v int32) *ListComputeMetricsByInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetRequestId(v string) *ListComputeMetricsByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListComputeMetricsByInstanceResponseBodyData struct {
	// List of pay-as-you-go job compute usage.
	InstanceComputeMetrics []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics `json:"instanceComputeMetrics,omitempty" xml:"instanceComputeMetrics,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of results returned.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetInstanceComputeMetrics() []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	return s.InstanceComputeMetrics
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetInstanceComputeMetrics(v []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) *ListComputeMetricsByInstanceResponseBodyData {
	s.InstanceComputeMetrics = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageNumber(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageSize(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetTotalCount(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics struct {
	// The end time of the job execution.
	//
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The job(instance) ID.
	//
	// example:
	//
	// 20240730****ddlr
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The owner of the job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// odps_porject
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// pqrs12345tuv
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// Specifications Type, specifies the resource package that you select when you purchase the MaxCompute service.
	//
	// - OdpsStandard: the pay-as-you-go resource package.
	//
	// - OdpsSpot: the pay-as-you-go spot resource package.
	//
	// example:
	//
	// OdpsStandard
	SpecCode *string `json:"specCode,omitempty" xml:"specCode,omitempty"`
	// The submission time of the job.
	//
	// example:
	//
	// 1610432000000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// Metering types.
	//
	// - ComputationSql: the metering data of SQL jobs that involve internal tables.
	//
	// - ComputationSqlOTS: the metering data of SQL jobs that involve Tablestore external tables.
	//
	// - ComputationSqlOSS: the metering data of SQL jobs that involve OSS external tables.
	//
	// - MapReduce: the metering data of MapReduce jobs.
	//
	// - spark: the metering data of Spark jobs.
	//
	// - mars: the metering data of Mars jobs.
	//
	// example:
	//
	// ComputationSql
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The unit of computing resource usage
	//
	// example:
	//
	// GB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The computing resource usage is calculated based on the following items:
	//
	// - Amount of scanned data in the unit of GB. For the jobs whose metering types are ComputationSql, ComputationSqlOTS, or ComputationSqlOSS, they are billed based on the amount of scanned data. The computing resource usage of such a job is calculated by using the following formula: Amount of scanned data × Complexity. The complexity is fixed at 1 for the jobs whose metering types are ComputationSqlOTS or ComputationSqlOSS.
	//
	// - CU-hours. For the jobs whose metering types are MapReduce, spark, or mars, they are billed based on CU-hours.
	//
	// example:
	//
	// 1024
	Usage *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSignature() *string {
	return s.Signature
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSpecCode() *string {
	return s.SpecCode
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetType() *string {
	return s.Type
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetUnit() *string {
	return s.Unit
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetUsage() *float64 {
	return s.Usage
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetEndTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.EndTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetInstanceId(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetJobOwner(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetProjectName(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.ProjectName = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSignature(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSpecCode(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SpecCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSubmitTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SubmitTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetType(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Type = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUnit(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Unit = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUsage(v float64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Usage = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) Validate() error {
	return dara.Validate(s)
}
