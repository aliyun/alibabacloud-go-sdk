// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnceTaskResultInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetOnceTaskResultInfoRequest
	GetTaskId() *string
	SetTaskName(v string) *GetOnceTaskResultInfoRequest
	GetTaskName() *string
	SetTaskType(v string) *GetOnceTaskResultInfoRequest
	GetTaskType() *string
}

type GetOnceTaskResultInfoRequest struct {
	// The ID of the one-time task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9dfa3a7eb9547781632785b49003****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task. Valid values:
	//
	// - **CLIENT_PROBLEM_CHECK**: client troubleshooting task
	//
	// - **CLIENT_DEV_OPS**: cloud O\\&M task
	//
	// - **ASSET_SECURITY_CHECK**: asset collection task
	//
	// - **ASSETS_COLLECTION**: asset fingerprint collection task
	//
	// - **IMAGE_SCAN**: container image scan task
	//
	// - **AI_SECURITY_CHECK**: AI asset synchronization task
	//
	// - **IDC_PROBE_SCAN**: IDC probe scan task
	//
	// - **ATTACK_SURFACE_SCAN**: attack surface and boundary asset scan task
	//
	// - **ASSET_EXPOSURE_SCAN**: asset exposure scan task
	//
	// - **VUL_CHECK_TASK**: vulnerability scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// - **CLIENT_PROBLEM_CHECK**: client troubleshooting task
	//
	// - **CLIENT_DEV_OPS**: cloud O\\&M task
	//
	// - **ASSET_SECURITY_CHECK**: asset collection task
	//
	// - **ASSETS_COLLECTION**: asset fingerprint collection task
	//
	// - **IMAGE_SCAN**: container image scan task
	//
	// - **AI_SECURITY_CHECK**: AI asset synchronization task
	//
	// - **IDC_PROBE_SCAN**: IDC probe scan task
	//
	// - **ATTACK_SURFACE_SCAN**: attack surface and boundary asset scan task
	//
	// - **ASSET_EXPOSURE_SCAN**: asset exposure scan task
	//
	// - **VUL_CHECK_TASK**: vulnerability scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// ASSETS_COLLECTION
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetOnceTaskResultInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOnceTaskResultInfoRequest) GoString() string {
	return s.String()
}

func (s *GetOnceTaskResultInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetOnceTaskResultInfoRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetOnceTaskResultInfoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetOnceTaskResultInfoRequest) SetTaskId(v string) *GetOnceTaskResultInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) SetTaskName(v string) *GetOnceTaskResultInfoRequest {
	s.TaskName = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) SetTaskType(v string) *GetOnceTaskResultInfoRequest {
	s.TaskType = &v
	return s
}

func (s *GetOnceTaskResultInfoRequest) Validate() error {
	return dara.Validate(s)
}
