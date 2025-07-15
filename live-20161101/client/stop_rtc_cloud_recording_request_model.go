// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *StopRtcCloudRecordingRequest
	GetTaskId() *string
}

type StopRtcCloudRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRtcCloudRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudRecordingRequest) GoString() string {
	return s.String()
}

func (s *StopRtcCloudRecordingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopRtcCloudRecordingRequest) SetTaskId(v string) *StopRtcCloudRecordingRequest {
	s.TaskId = &v
	return s
}

func (s *StopRtcCloudRecordingRequest) Validate() error {
	return dara.Validate(s)
}
