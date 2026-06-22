// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskType(v string) *GetVirusScanConfigRequest
	GetTaskType() *string
}

type GetVirusScanConfigRequest struct {
	// The task type. Valid values:
	//
	// - **VIRUS_VUL_SCHEDULE_SCAN**: virus scan.
	//
	// example:
	//
	// VIRUS_VUL_SCHEDULE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVirusScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanConfigRequest) GoString() string {
	return s.String()
}

func (s *GetVirusScanConfigRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVirusScanConfigRequest) SetTaskType(v string) *GetVirusScanConfigRequest {
	s.TaskType = &v
	return s
}

func (s *GetVirusScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
