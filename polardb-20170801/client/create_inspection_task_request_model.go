// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInspectionTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCronExpression(v string) *CreateInspectionTaskRequest
	GetCronExpression() *string
	SetDBClusterList(v []*CreateInspectionTaskRequestDBClusterList) *CreateInspectionTaskRequest
	GetDBClusterList() []*CreateInspectionTaskRequestDBClusterList
	SetRegionId(v string) *CreateInspectionTaskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInspectionTaskRequest
	GetResourceGroupId() *string
	SetSpaceId(v string) *CreateInspectionTaskRequest
	GetSpaceId() *string
	SetTaskDescription(v string) *CreateInspectionTaskRequest
	GetTaskDescription() *string
}

type CreateInspectionTaskRequest struct {
	// The cron expression for the inspection task, which specifies the periodic execution schedule.
	//
	// example:
	//
	// 0 0 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The list of clusters associated with the inspection task.
	DBClusterList []*CreateInspectionTaskRequestDBClusterList `json:"DBClusterList,omitempty" xml:"DBClusterList,omitempty" type:"Repeated"`
	// The ID of the region where the task is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-123
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the inspection workspace.
	//
	// example:
	//
	// space-123
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The description of the inspection task.
	//
	// example:
	//
	// Periodic inspection task
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
}

func (s CreateInspectionTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateInspectionTaskRequest) GetDBClusterList() []*CreateInspectionTaskRequestDBClusterList {
	return s.DBClusterList
}

func (s *CreateInspectionTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInspectionTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInspectionTaskRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *CreateInspectionTaskRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *CreateInspectionTaskRequest) SetCronExpression(v string) *CreateInspectionTaskRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetDBClusterList(v []*CreateInspectionTaskRequestDBClusterList) *CreateInspectionTaskRequest {
	s.DBClusterList = v
	return s
}

func (s *CreateInspectionTaskRequest) SetRegionId(v string) *CreateInspectionTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetResourceGroupId(v string) *CreateInspectionTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetSpaceId(v string) *CreateInspectionTaskRequest {
	s.SpaceId = &v
	return s
}

func (s *CreateInspectionTaskRequest) SetTaskDescription(v string) *CreateInspectionTaskRequest {
	s.TaskDescription = &v
	return s
}

func (s *CreateInspectionTaskRequest) Validate() error {
	if s.DBClusterList != nil {
		for _, item := range s.DBClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInspectionTaskRequestDBClusterList struct {
	// The ID of the PolarDB cluster to be inspected.
	//
	// example:
	//
	// pc-test123456
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the region where the cluster to be inspected is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInspectionTaskRequestDBClusterList) String() string {
	return dara.Prettify(s)
}

func (s CreateInspectionTaskRequestDBClusterList) GoString() string {
	return s.String()
}

func (s *CreateInspectionTaskRequestDBClusterList) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateInspectionTaskRequestDBClusterList) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInspectionTaskRequestDBClusterList) SetDBClusterId(v string) *CreateInspectionTaskRequestDBClusterList {
	s.DBClusterId = &v
	return s
}

func (s *CreateInspectionTaskRequestDBClusterList) SetRegionId(v string) *CreateInspectionTaskRequestDBClusterList {
	s.RegionId = &v
	return s
}

func (s *CreateInspectionTaskRequestDBClusterList) Validate() error {
	return dara.Validate(s)
}
