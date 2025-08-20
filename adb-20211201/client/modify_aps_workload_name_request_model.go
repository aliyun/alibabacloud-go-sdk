// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsWorkloadNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyApsWorkloadNameRequest
	GetDBClusterId() *string
	SetRegionId(v string) *ModifyApsWorkloadNameRequest
	GetRegionId() *string
	SetWorkloadId(v string) *ModifyApsWorkloadNameRequest
	GetWorkloadId() *string
	SetWorkloadName(v string) *ModifyApsWorkloadNameRequest
	GetWorkloadName() *string
}

type ModifyApsWorkloadNameRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps-******
	WorkloadId *string `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// The name of the workload.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-2024***93014
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s ModifyApsWorkloadNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsWorkloadNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsWorkloadNameRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyApsWorkloadNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsWorkloadNameRequest) GetWorkloadId() *string {
	return s.WorkloadId
}

func (s *ModifyApsWorkloadNameRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *ModifyApsWorkloadNameRequest) SetDBClusterId(v string) *ModifyApsWorkloadNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyApsWorkloadNameRequest) SetRegionId(v string) *ModifyApsWorkloadNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsWorkloadNameRequest) SetWorkloadId(v string) *ModifyApsWorkloadNameRequest {
	s.WorkloadId = &v
	return s
}

func (s *ModifyApsWorkloadNameRequest) SetWorkloadName(v string) *ModifyApsWorkloadNameRequest {
	s.WorkloadName = &v
	return s
}

func (s *ModifyApsWorkloadNameRequest) Validate() error {
	return dara.Validate(s)
}
