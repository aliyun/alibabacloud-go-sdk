// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobDraftSqlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateComputeJobDraftSqlRequest
	GetClientToken() *string
	SetDraftSql(v string) *UpdateComputeJobDraftSqlRequest
	GetDraftSql() *string
	SetInstanceId(v string) *UpdateComputeJobDraftSqlRequest
	GetInstanceId() *string
	SetJobName(v string) *UpdateComputeJobDraftSqlRequest
	GetJobName() *string
	SetRegionId(v string) *UpdateComputeJobDraftSqlRequest
	GetRegionId() *string
}

type UpdateComputeJobDraftSqlRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DraftSql *string `json:"DraftSql,omitempty" xml:"DraftSql,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateComputeJobDraftSqlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobDraftSqlRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobDraftSqlRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateComputeJobDraftSqlRequest) GetDraftSql() *string {
	return s.DraftSql
}

func (s *UpdateComputeJobDraftSqlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateComputeJobDraftSqlRequest) GetJobName() *string {
	return s.JobName
}

func (s *UpdateComputeJobDraftSqlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateComputeJobDraftSqlRequest) SetClientToken(v string) *UpdateComputeJobDraftSqlRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateComputeJobDraftSqlRequest) SetDraftSql(v string) *UpdateComputeJobDraftSqlRequest {
	s.DraftSql = &v
	return s
}

func (s *UpdateComputeJobDraftSqlRequest) SetInstanceId(v string) *UpdateComputeJobDraftSqlRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateComputeJobDraftSqlRequest) SetJobName(v string) *UpdateComputeJobDraftSqlRequest {
	s.JobName = &v
	return s
}

func (s *UpdateComputeJobDraftSqlRequest) SetRegionId(v string) *UpdateComputeJobDraftSqlRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateComputeJobDraftSqlRequest) Validate() error {
	return dara.Validate(s)
}
