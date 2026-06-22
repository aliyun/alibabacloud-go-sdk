// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParam(v string) *GenerateOnceTaskRequest
	GetParam() *string
	SetSource(v string) *GenerateOnceTaskRequest
	GetSource() *string
	SetTaskName(v string) *GenerateOnceTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *GenerateOnceTaskRequest
	GetTaskType() *string
}

type GenerateOnceTaskRequest struct {
	// A JSON string that contains additional parameters for the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"mode":1,"problemType":"offline","uuids":"inet-795dcad1-360f-49d2-b01e-b7da7f1c****"}
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The source that initiated the task.
	//
	// example:
	//
	// Manual
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The name of the scan task.
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// -
	//
	// This parameter is required.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the scan task. Valid values:
	//
	// - **CLIENT_PROBLEM_CHECK**: a client troubleshooting task
	//
	// - **CLIENT_DEV_OPS**: a cloud DevOps task
	//
	// - **ASSET_SECURITY_CHECK**: an asset collection task
	//
	// - **ASSETS_COLLECTION**: an asset fingerprinting task
	//
	// - **IMAGE_SCAN**: a container image scan task
	//
	// - **AI_SECURITY_CHECK**: an AI asset synchronization task
	//
	// - **IDC_PROBE_SCAN**: an IDC probe scan task
	//
	// - **ATTACK_SURFACE_SCAN**: an attack surface scan task
	//
	// - **ASSET_EXPOSURE_SCAN**: an asset exposure scan task
	//
	// - **VUL_CHECK_TASK**: a vulnerability scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// CLIENT_PROBLEM_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GenerateOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *GenerateOnceTaskRequest) GetParam() *string {
	return s.Param
}

func (s *GenerateOnceTaskRequest) GetSource() *string {
	return s.Source
}

func (s *GenerateOnceTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GenerateOnceTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GenerateOnceTaskRequest) SetParam(v string) *GenerateOnceTaskRequest {
	s.Param = &v
	return s
}

func (s *GenerateOnceTaskRequest) SetSource(v string) *GenerateOnceTaskRequest {
	s.Source = &v
	return s
}

func (s *GenerateOnceTaskRequest) SetTaskName(v string) *GenerateOnceTaskRequest {
	s.TaskName = &v
	return s
}

func (s *GenerateOnceTaskRequest) SetTaskType(v string) *GenerateOnceTaskRequest {
	s.TaskType = &v
	return s
}

func (s *GenerateOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
