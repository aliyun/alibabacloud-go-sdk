// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnfinishedOnceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTarget(v string) *ListUnfinishedOnceTaskRequest
	GetTarget() *string
	SetTaskType(v string) *ListUnfinishedOnceTaskRequest
	GetTaskType() *string
}

type ListUnfinishedOnceTaskRequest struct {
	// The value of the object on which the task runs. If you set TaskType to IMAGE_SCAN, set this parameter to the UUID of the image that you want to scan. If you set TaskType to ASSETS_COLLECTION, set this parameter to the UUID of the server whose information you want to collect.
	//
	// example:
	//
	// 4fe8e1cd-3c37-4851-b9de-124da32c****
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **ASSETS_COLLECTION**: asset information collection task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListUnfinishedOnceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUnfinishedOnceTaskRequest) GoString() string {
	return s.String()
}

func (s *ListUnfinishedOnceTaskRequest) GetTarget() *string {
	return s.Target
}

func (s *ListUnfinishedOnceTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListUnfinishedOnceTaskRequest) SetTarget(v string) *ListUnfinishedOnceTaskRequest {
	s.Target = &v
	return s
}

func (s *ListUnfinishedOnceTaskRequest) SetTaskType(v string) *ListUnfinishedOnceTaskRequest {
	s.TaskType = &v
	return s
}

func (s *ListUnfinishedOnceTaskRequest) Validate() error {
	return dara.Validate(s)
}
