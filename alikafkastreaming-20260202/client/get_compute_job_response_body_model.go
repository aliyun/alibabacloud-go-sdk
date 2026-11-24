// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetComputeJobResponseBody
	GetCode() *int64
	SetData(v *GetComputeJobResponseBodyData) *GetComputeJobResponseBody
	GetData() *GetComputeJobResponseBodyData
	SetRequestId(v string) *GetComputeJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetComputeJobResponseBody
	GetSuccess() *bool
}

type GetComputeJobResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetComputeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 062D8E8B-8D47-5DCC-BB12-5A1D93C3A66B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetComputeJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetComputeJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetComputeJobResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetComputeJobResponseBody) GetData() *GetComputeJobResponseBodyData {
	return s.Data
}

func (s *GetComputeJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetComputeJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetComputeJobResponseBody) SetCode(v int64) *GetComputeJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetComputeJobResponseBody) SetData(v *GetComputeJobResponseBodyData) *GetComputeJobResponseBody {
	s.Data = v
	return s
}

func (s *GetComputeJobResponseBody) SetRequestId(v string) *GetComputeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetComputeJobResponseBody) SetSuccess(v bool) *GetComputeJobResponseBody {
	s.Success = &v
	return s
}

func (s *GetComputeJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetComputeJobResponseBodyData struct {
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-09-02T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The CU limit for the job.
	//
	// example:
	//
	// 2.0
	CuLimit *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	// The reserved CUs for the job.
	//
	// example:
	//
	// 1.0
	CuReserved *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	// The used CUs.
	//
	// example:
	//
	// 1.5
	CuUsed *float64 `json:"CuUsed,omitempty" xml:"CuUsed,omitempty"`
	// Indicates whether the job is a debug job. Valid values:
	//
	// - 0: No.
	//
	// - 1: Yes.
	//
	// example:
	//
	// 0
	DebugMode *int32 `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	// The currently deployed SQL.
	//
	// example:
	//
	// INSERT INTO sink_table SELECT 	- FROM source_table;
	DeployedSql *string `json:"DeployedSql,omitempty" xml:"DeployedSql,omitempty"`
	// The Flink SQL draft.
	//
	// example:
	//
	// INSERT INTO sink_table SELECT 	- FROM source_table;
	DraftSql *string `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	// The error message.
	//
	// example:
	//
	// SQL validation or compilation failed: Column \\"xxx\\" not found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The expiration time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-09-11T04:33:03Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The compute instance ID.
	//
	// example:
	//
	// alikafka_streaming-cn-hangzhou-a1b2c3d4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the compute job.
	//
	// example:
	//
	// order_enrichment
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The owner of the resource.
	//
	// example:
	//
	// 1234567890123456
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remark for the job.
	//
	// example:
	//
	// Real-time order stream cleansing
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the job.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The startup mode, such as savepoint or stateless.
	//
	// example:
	//
	// savepoint
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
}

func (s GetComputeJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetComputeJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetComputeJobResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetComputeJobResponseBodyData) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *GetComputeJobResponseBodyData) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *GetComputeJobResponseBodyData) GetCuUsed() *float64 {
	return s.CuUsed
}

func (s *GetComputeJobResponseBodyData) GetDebugMode() *int32 {
	return s.DebugMode
}

func (s *GetComputeJobResponseBodyData) GetDeployedSql() *string {
	return s.DeployedSql
}

func (s *GetComputeJobResponseBodyData) GetDraftSql() *string {
	return s.DraftSql
}

func (s *GetComputeJobResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetComputeJobResponseBodyData) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *GetComputeJobResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeJobResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *GetComputeJobResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetComputeJobResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeJobResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetComputeJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetComputeJobResponseBodyData) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *GetComputeJobResponseBodyData) SetCreateTime(v string) *GetComputeJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetCuLimit(v float64) *GetComputeJobResponseBodyData {
	s.CuLimit = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetCuReserved(v float64) *GetComputeJobResponseBodyData {
	s.CuReserved = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetCuUsed(v float64) *GetComputeJobResponseBodyData {
	s.CuUsed = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetDebugMode(v int32) *GetComputeJobResponseBodyData {
	s.DebugMode = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetDeployedSql(v string) *GetComputeJobResponseBodyData {
	s.DeployedSql = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetDraftSql(v string) *GetComputeJobResponseBodyData {
	s.DraftSql = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetErrorMsg(v string) *GetComputeJobResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetExpirationTime(v string) *GetComputeJobResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetInstanceId(v string) *GetComputeJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetJobName(v string) *GetComputeJobResponseBodyData {
	s.JobName = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetOwner(v string) *GetComputeJobResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetRegionId(v string) *GetComputeJobResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetRemark(v string) *GetComputeJobResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetStatus(v string) *GetComputeJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetUpgradeMode(v string) *GetComputeJobResponseBodyData {
	s.UpgradeMode = &v
	return s
}

func (s *GetComputeJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
