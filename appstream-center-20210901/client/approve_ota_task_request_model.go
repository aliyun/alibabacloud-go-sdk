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
	// The delivery group ID. You can call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The region ID of the delivery group. You can call [ListRegions](~~ListRegions~~) to query the list of regions supported by Wuying Cloud Application.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The over-the-air upgrade task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Fota
	OtaType *string `json:"OtaType,omitempty" xml:"OtaType,omitempty"`
	// The start time of the over-the-air upgrade task. Specify the time in ISO 8601 format.
	//
	// This parameter is required.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2022-08-04T14:36:00+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The over-the-air upgrade task ID. You can call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the ID.
	//
	// > Each successful call of `ApproveOtaTask` causes the `TaskId` to change. Therefore, before calling this operation again, call `ListAppInstanceGroup` again to obtain the latest `TaskId`.
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
