// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteJobs(v bool) *DeleteAppGroupRequest
	GetDeleteJobs() *bool
	SetGroupId(v string) *DeleteAppGroupRequest
	GetGroupId() *string
	SetNamespace(v string) *DeleteAppGroupRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteAppGroupRequest
	GetRegionId() *string
}

type DeleteAppGroupRequest struct {
	// Specifies whether to delete the jobs in the application group. Valid values:
	//
	// true: deletes the jobs in the application group.
	//
	// false: does not delete the jobs in the application group.
	//
	// example:
	//
	// true
	DeleteJobs *bool `json:"DeleteJobs,omitempty" xml:"DeleteJobs,omitempty"`
	// The ID of the application. You can obtain the ID of the application on the Application Management page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the namespace. You can obtain the ID of the namespace on the Namespace page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupRequest) GetDeleteJobs() *bool {
	return s.DeleteJobs
}

func (s *DeleteAppGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteAppGroupRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteAppGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAppGroupRequest) SetDeleteJobs(v bool) *DeleteAppGroupRequest {
	s.DeleteJobs = &v
	return s
}

func (s *DeleteAppGroupRequest) SetGroupId(v string) *DeleteAppGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteAppGroupRequest) SetNamespace(v string) *DeleteAppGroupRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteAppGroupRequest) SetRegionId(v string) *DeleteAppGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
