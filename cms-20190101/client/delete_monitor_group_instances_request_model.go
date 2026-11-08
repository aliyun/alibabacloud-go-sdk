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
	SetClientToken(v string) *DeleteMonitorGroupInstancesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteMonitorGroupInstancesRequest
	GetDryRun() *bool
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
	// > For more information about how to obtain the abbreviation of a cloud service name, see the `metricCategory` field in the `Labels` response parameter of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs a dry run and does not delete the VPN connection. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the request passes the dry run, the VPN connection is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instances to delete. Separate multiple instance IDs with commas (,). You can delete up to 20 instances at a time.
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

func (s *DeleteMonitorGroupInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteMonitorGroupInstancesRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *DeleteMonitorGroupInstancesRequest) SetClientToken(v string) *DeleteMonitorGroupInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteMonitorGroupInstancesRequest) SetDryRun(v bool) *DeleteMonitorGroupInstancesRequest {
	s.DryRun = &v
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
