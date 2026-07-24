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
	Code      *int64                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetComputeJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CuLimit      *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	CuReserved   *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	CuUsed       *float64 `json:"CuUsed,omitempty" xml:"CuUsed,omitempty"`
	DebugMode    *int32   `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	DeployedSql  *string  `json:"DeployedSql,omitempty" xml:"DeployedSql,omitempty"`
	DraftSql     *string  `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	HistoryInfos *string  `json:"HistoryInfos,omitempty" xml:"HistoryInfos,omitempty"`
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobConfig    *string  `json:"JobConfig,omitempty" xml:"JobConfig,omitempty"`
	JobName      *string  `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Owner        *string  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RegionId     *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark       *string  `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpgradeMode  *string  `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
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

func (s *GetComputeJobResponseBodyData) GetHistoryInfos() *string {
	return s.HistoryInfos
}

func (s *GetComputeJobResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeJobResponseBodyData) GetJobConfig() *string {
	return s.JobConfig
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

func (s *GetComputeJobResponseBodyData) SetHistoryInfos(v string) *GetComputeJobResponseBodyData {
	s.HistoryInfos = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetInstanceId(v string) *GetComputeJobResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetComputeJobResponseBodyData) SetJobConfig(v string) *GetComputeJobResponseBodyData {
	s.JobConfig = &v
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
