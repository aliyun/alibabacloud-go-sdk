// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlanMaintenanceWindowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanWindowId(v string) *DeletePlanMaintenanceWindowRequest
	GetPlanWindowId() *string
	SetRegionId(v string) *DeletePlanMaintenanceWindowRequest
	GetRegionId() *string
}

type DeletePlanMaintenanceWindowRequest struct {
	// The ID of the maintenance window to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// pw-bp12kkvnebe7hksqnx9w
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// The ID of the region. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to retrieve the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePlanMaintenanceWindowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlanMaintenanceWindowRequest) GoString() string {
	return s.String()
}

func (s *DeletePlanMaintenanceWindowRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *DeletePlanMaintenanceWindowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePlanMaintenanceWindowRequest) SetPlanWindowId(v string) *DeletePlanMaintenanceWindowRequest {
	s.PlanWindowId = &v
	return s
}

func (s *DeletePlanMaintenanceWindowRequest) SetRegionId(v string) *DeletePlanMaintenanceWindowRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePlanMaintenanceWindowRequest) Validate() error {
	return dara.Validate(s)
}
