// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFotaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateFotaTaskRequest
	GetRegionId() *string
	SetTaskUid(v string) *UpdateFotaTaskRequest
	GetTaskUid() *string
	SetUserStatus(v string) *UpdateFotaTaskRequest
	GetUserStatus() *string
}

type UpdateFotaTaskRequest struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image update task. You can call the [DescribeFotaTasks](https://help.aliyun.com/document_detail/437001.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// aot-c4khwrp9ocml4****
	TaskUid *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
	// Specifies whether to automatically push the image update task.
	//
	// Valid values:
	//
	// 	- Running: automatically pushes the image update task.
	//
	// 	- Pending: does not automatically push the image update task.
	//
	// example:
	//
	// Pending
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s UpdateFotaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFotaTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateFotaTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateFotaTaskRequest) GetTaskUid() *string {
	return s.TaskUid
}

func (s *UpdateFotaTaskRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *UpdateFotaTaskRequest) SetRegionId(v string) *UpdateFotaTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateFotaTaskRequest) SetTaskUid(v string) *UpdateFotaTaskRequest {
	s.TaskUid = &v
	return s
}

func (s *UpdateFotaTaskRequest) SetUserStatus(v string) *UpdateFotaTaskRequest {
	s.UserStatus = &v
	return s
}

func (s *UpdateFotaTaskRequest) Validate() error {
	return dara.Validate(s)
}
