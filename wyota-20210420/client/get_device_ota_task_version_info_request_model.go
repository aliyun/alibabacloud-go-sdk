// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceOtaTaskVersionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDeviceOtaTaskVersionInfoRequest
	GetTaskId() *string
}

type GetDeviceOtaTaskVersionInfoRequest struct {
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDeviceOtaTaskVersionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceOtaTaskVersionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceOtaTaskVersionInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDeviceOtaTaskVersionInfoRequest) SetTaskId(v string) *GetDeviceOtaTaskVersionInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetDeviceOtaTaskVersionInfoRequest) Validate() error {
	return dara.Validate(s)
}
