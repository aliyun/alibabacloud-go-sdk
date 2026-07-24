// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateComputeJobRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *UpdateComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *UpdateComputeJobRequest
	GetRegionId() *string
	SetRemark(v string) *UpdateComputeJobRequest
	GetRemark() *string
	SetUpgradeMode(v string) *UpdateComputeJobRequest
	GetUpgradeMode() *string
}

type UpdateComputeJobRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark      *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	UpgradeMode *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
}

func (s UpdateComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *UpdateComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateComputeJobRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateComputeJobRequest) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *UpdateComputeJobRequest) SetClientToken(v string) *UpdateComputeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateComputeJobRequest) SetInstanceId(v string) *UpdateComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateComputeJobRequest) SetJobName(v string) *UpdateComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *UpdateComputeJobRequest) SetRegionId(v string) *UpdateComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateComputeJobRequest) SetRemark(v string) *UpdateComputeJobRequest {
	s.Remark = &v
	return s
}

func (s *UpdateComputeJobRequest) SetUpgradeMode(v string) *UpdateComputeJobRequest {
	s.UpgradeMode = &v
	return s
}

func (s *UpdateComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
