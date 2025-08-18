// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOtaTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ApproveOtaTaskRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *ApproveOtaTaskRequest
	GetBizRegionId() *string
	SetOtaType(v string) *ApproveOtaTaskRequest
	GetOtaType() *string
	SetStartTime(v string) *ApproveOtaTaskRequest
	GetStartTime() *string
	SetTaskId(v string) *ApproveOtaTaskRequest
	GetTaskId() *string
}

type ApproveOtaTaskRequest struct {
	// The ID of the delivery group. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the region where the delivery group resides. You can call the [ListRegions](https://help.aliyun.com/document_detail/428500.html) operation to query the list of regions supported by App Streaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The type of the OTA update task.
	//
	// Valid values:
	//
	// 	- Fota: update of the system components of Alibaba Cloud Workspace
	//
	// 	- AppUpdate
	//
	// 	- ImageUpdate
	//
	// This parameter is required.
	//
	// example:
	//
	// Fota
	OtaType *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	// The start time of the OTA update task. The time follows the ISO 8601 standard.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the OTA update task. You can call the [ListAppInstanceGroup](https://help.aliyun.com/document_detail/428506.html) operation to obtain the ID.
	//
	// >  Each successful call to the `ApproveOtaTask` operation causes a value change of this parameter.`` Before you call this operation, call the `ListAppInstanceGroup` operation again to obtain the latest value of this parameter.``
	//
	// This parameter is required.
	//
	// example:
	//
	// ota-be7jzm29wrrz5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApproveOtaTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveOtaTaskRequest) GoString() string {
	return s.String()
}

func (s *ApproveOtaTaskRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ApproveOtaTaskRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ApproveOtaTaskRequest) GetOtaType() *string {
	return s.OtaType
}

func (s *ApproveOtaTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ApproveOtaTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ApproveOtaTaskRequest) SetAppInstanceGroupId(v string) *ApproveOtaTaskRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetBizRegionId(v string) *ApproveOtaTaskRequest {
	s.BizRegionId = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetOtaType(v string) *ApproveOtaTaskRequest {
	s.OtaType = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetStartTime(v string) *ApproveOtaTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ApproveOtaTaskRequest) SetTaskId(v string) *ApproveOtaTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ApproveOtaTaskRequest) Validate() error {
	return dara.Validate(s)
}
