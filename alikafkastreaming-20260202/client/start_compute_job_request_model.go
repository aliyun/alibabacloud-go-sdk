// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCuLimit(v float64) *StartComputeJobRequest
	GetCuLimit() *float64
	SetCuReserved(v float64) *StartComputeJobRequest
	GetCuReserved() *float64
	SetDraftSql(v string) *StartComputeJobRequest
	GetDraftSql() *string
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
	// example:
	//
	// 2.0
	CuLimit *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	// example:
	//
	// 1.0
	CuReserved *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	// example:
	//
	// CREATE TEMPORARY TABLE src (id BIGINT) WITH (\\"connector\\" = \\"datagen\\"); CREATE TEMPORARY TABLE sink (id BIGINT) WITH (\\"connector\\" = \\"print\\"); INSERT INTO sink SELECT id FROM src;
	DraftSql *string `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order_enrichment
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// savepoint
	RecoveryMode *string `json:"RecoveryMode,omitempty" xml:"RecoveryMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartComputeJobRequest) GoString() string {
	return s.String()
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
