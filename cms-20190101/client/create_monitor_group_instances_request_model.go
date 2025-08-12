// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateMonitorGroupInstancesRequest
	GetGroupId() *string
	SetInstances(v []*CreateMonitorGroupInstancesRequestInstances) *CreateMonitorGroupInstancesRequest
	GetInstances() []*CreateMonitorGroupInstancesRequestInstances
	SetRegionId(v string) *CreateMonitorGroupInstancesRequest
	GetRegionId() *string
}

type CreateMonitorGroupInstancesRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instances that you want to add to the application group.
	//
	// This parameter is required.
	Instances []*CreateMonitorGroupInstancesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RegionId  *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMonitorGroupInstancesRequest) GetInstances() []*CreateMonitorGroupInstancesRequestInstances {
	return s.Instances
}

func (s *CreateMonitorGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupInstancesRequest) SetGroupId(v string) *CreateMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetInstances(v []*CreateMonitorGroupInstancesRequestInstances) *CreateMonitorGroupInstancesRequest {
	s.Instances = v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) SetRegionId(v string) *CreateMonitorGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateMonitorGroupInstancesRequestInstances struct {
	// The abbreviation of the Alibaba Cloud service name.
	//
	// To obtain the abbreviation of an Alibaba Cloud service name, call the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation. The `metricCategory` tag in the `Labels` response parameter indicates the abbreviation of the Alibaba Cloud service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2ze26xj5wwy12****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-instance-ecs
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMonitorGroupInstancesRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupInstancesRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetCategory() *string {
	return s.Category
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateMonitorGroupInstancesRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetCategory(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.Category = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetInstanceName(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.InstanceName = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) SetRegionId(v string) *CreateMonitorGroupInstancesRequestInstances {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupInstancesRequestInstances) Validate() error {
	return dara.Validate(s)
}
