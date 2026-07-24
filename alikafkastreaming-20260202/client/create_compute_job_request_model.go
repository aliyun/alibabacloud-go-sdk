// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateComputeJobRequest
	GetClientToken() *string
	SetCuLimit(v float64) *CreateComputeJobRequest
	GetCuLimit() *float64
	SetCuReserved(v float64) *CreateComputeJobRequest
	GetCuReserved() *float64
	SetDraftSql(v string) *CreateComputeJobRequest
	GetDraftSql() *string
	SetInstanceId(v string) *CreateComputeJobRequest
	GetInstanceId() *string
	SetJobConfig(v string) *CreateComputeJobRequest
	GetJobConfig() *string
	SetJobName(v string) *CreateComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *CreateComputeJobRequest
	GetRegionId() *string
	SetRemark(v string) *CreateComputeJobRequest
	GetRemark() *string
	SetUpgradeMode(v string) *CreateComputeJobRequest
	GetUpgradeMode() *string
	SetUserId(v string) *CreateComputeJobRequest
	GetUserId() *string
}

type CreateComputeJobRequest struct {
	ClientToken *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CuLimit     *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	CuReserved  *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	DraftSql    *string  `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobConfig  *string `json:"JobConfig,omitempty" xml:"JobConfig,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateComputeJobRequest) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *CreateComputeJobRequest) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *CreateComputeJobRequest) GetDraftSql() *string {
	return s.DraftSql
}

func (s *CreateComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateComputeJobRequest) GetJobConfig() *string {
	return s.JobConfig
}

func (s *CreateComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateComputeJobRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateComputeJobRequest) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *CreateComputeJobRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateComputeJobRequest) SetClientToken(v string) *CreateComputeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateComputeJobRequest) SetCuLimit(v float64) *CreateComputeJobRequest {
	s.CuLimit = &v
	return s
}

func (s *CreateComputeJobRequest) SetCuReserved(v float64) *CreateComputeJobRequest {
	s.CuReserved = &v
	return s
}

func (s *CreateComputeJobRequest) SetDraftSql(v string) *CreateComputeJobRequest {
	s.DraftSql = &v
	return s
}

func (s *CreateComputeJobRequest) SetInstanceId(v string) *CreateComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateComputeJobRequest) SetJobConfig(v string) *CreateComputeJobRequest {
	s.JobConfig = &v
	return s
}

func (s *CreateComputeJobRequest) SetJobName(v string) *CreateComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *CreateComputeJobRequest) SetRegionId(v string) *CreateComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateComputeJobRequest) SetRemark(v string) *CreateComputeJobRequest {
	s.Remark = &v
	return s
}

func (s *CreateComputeJobRequest) SetUpgradeMode(v string) *CreateComputeJobRequest {
	s.UpgradeMode = &v
	return s
}

func (s *CreateComputeJobRequest) SetUserId(v string) *CreateComputeJobRequest {
	s.UserId = &v
	return s
}

func (s *CreateComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
