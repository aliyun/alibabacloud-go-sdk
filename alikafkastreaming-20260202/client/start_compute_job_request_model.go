// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartComputeJobRequest
	GetClientToken() *string
	SetCuLimit(v float64) *StartComputeJobRequest
	GetCuLimit() *float64
	SetCuReserved(v float64) *StartComputeJobRequest
	GetCuReserved() *float64
	SetDraftSql(v string) *StartComputeJobRequest
	GetDraftSql() *string
	SetDraftSqlStart(v bool) *StartComputeJobRequest
	GetDraftSqlStart() *bool
	SetInstanceId(v string) *StartComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *StartComputeJobRequest
	GetJobName() *string
	SetRecoveryMode(v string) *StartComputeJobRequest
	GetRecoveryMode() *string
	SetRegionId(v string) *StartComputeJobRequest
	GetRegionId() *string
}

type StartComputeJobRequest struct {
	ClientToken   *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CuLimit       *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	CuReserved    *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	DraftSql      *string  `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	DraftSqlStart *bool    `json:"DraftSqlStart,omitempty" xml:"DraftSqlStart,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName      *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	RecoveryMode *string `json:"RecoveryMode,omitempty" xml:"RecoveryMode,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartComputeJobRequest) GoString() string {
	return s.String()
}

func (s *StartComputeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartComputeJobRequest) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *StartComputeJobRequest) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *StartComputeJobRequest) GetDraftSql() *string {
	return s.DraftSql
}

func (s *StartComputeJobRequest) GetDraftSqlStart() *bool {
	return s.DraftSqlStart
}

func (s *StartComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *StartComputeJobRequest) GetRecoveryMode() *string {
	return s.RecoveryMode
}

func (s *StartComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartComputeJobRequest) SetClientToken(v string) *StartComputeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StartComputeJobRequest) SetCuLimit(v float64) *StartComputeJobRequest {
	s.CuLimit = &v
	return s
}

func (s *StartComputeJobRequest) SetCuReserved(v float64) *StartComputeJobRequest {
	s.CuReserved = &v
	return s
}

func (s *StartComputeJobRequest) SetDraftSql(v string) *StartComputeJobRequest {
	s.DraftSql = &v
	return s
}

func (s *StartComputeJobRequest) SetDraftSqlStart(v bool) *StartComputeJobRequest {
	s.DraftSqlStart = &v
	return s
}

func (s *StartComputeJobRequest) SetInstanceId(v string) *StartComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *StartComputeJobRequest) SetJobName(v string) *StartComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *StartComputeJobRequest) SetRecoveryMode(v string) *StartComputeJobRequest {
	s.RecoveryMode = &v
	return s
}

func (s *StartComputeJobRequest) SetRegionId(v string) *StartComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
