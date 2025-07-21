// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDeviceOtaTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationStatus(v int32) *SetDeviceOtaTaskStatusRequest
	GetOperationStatus() *int32
	SetTaskId(v int32) *SetDeviceOtaTaskStatusRequest
	GetTaskId() *int32
}

type SetDeviceOtaTaskStatusRequest struct {
	// This parameter is required.
	OperationStatus *int32 `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// This parameter is required.
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetDeviceOtaTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDeviceOtaTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDeviceOtaTaskStatusRequest) GetOperationStatus() *int32 {
	return s.OperationStatus
}

func (s *SetDeviceOtaTaskStatusRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *SetDeviceOtaTaskStatusRequest) SetOperationStatus(v int32) *SetDeviceOtaTaskStatusRequest {
	s.OperationStatus = &v
	return s
}

func (s *SetDeviceOtaTaskStatusRequest) SetTaskId(v int32) *SetDeviceOtaTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *SetDeviceOtaTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
