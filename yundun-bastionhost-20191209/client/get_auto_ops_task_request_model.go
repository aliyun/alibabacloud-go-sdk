// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoOpsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAutoOpsTaskRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetAutoOpsTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *GetAutoOpsTaskRequest
	GetTaskId() *string
}

type GetAutoOpsTaskRequest struct {
	// The instance ID of the bastion host.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the O&M task that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAutoOpsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoOpsTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAutoOpsTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoOpsTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAutoOpsTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAutoOpsTaskRequest) SetInstanceId(v string) *GetAutoOpsTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoOpsTaskRequest) SetRegionId(v string) *GetAutoOpsTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetAutoOpsTaskRequest) SetTaskId(v string) *GetAutoOpsTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAutoOpsTaskRequest) Validate() error {
	return dara.Validate(s)
}
