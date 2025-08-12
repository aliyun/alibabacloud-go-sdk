// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DeleteMonitorGroupInstancesRequest
	GetCategory() *string
	SetGroupId(v int64) *DeleteMonitorGroupInstancesRequest
	GetGroupId() *int64
	SetInstanceIdList(v string) *DeleteMonitorGroupInstancesRequest
	GetInstanceIdList() *string
	SetRegionId(v string) *DeleteMonitorGroupInstancesRequest
	GetRegionId() *string
}

type DeleteMonitorGroupInstancesRequest struct {
	// The abbreviation of the cloud service name.
	//
	// >  For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instances to be removed from the application group. Separate multiple instances with commas (,). You can remove a maximum of 20 instances from an application group at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-a2d5q7pm3f912****,i-a2d5q7pm3f222****
	InstanceIdList *string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitorGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupInstancesRequest) GetCategory() *string {
	return s.Category
}

func (s *DeleteMonitorGroupInstancesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DeleteMonitorGroupInstancesRequest) GetInstanceIdList() *string {
	return s.InstanceIdList
}

func (s *DeleteMonitorGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitorGroupInstancesRequest) SetCategory(v string) *DeleteMonitorGroupInstancesRequest {
	s.Category = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetGroupId(v int64) *DeleteMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetInstanceIdList(v string) *DeleteMonitorGroupInstancesRequest {
	s.InstanceIdList = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetRegionId(v string) *DeleteMonitorGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) Validate() error {
	return dara.Validate(s)
}
