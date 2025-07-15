// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopRtcCloudRecordingResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopRtcCloudRecordingResponseBody
	GetTaskId() *string
}

type StopRtcCloudRecordingResponseBody struct {
	// example:
	//
	// ******58-5876-****-83CA-B56278******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRtcCloudRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *StopRtcCloudRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRtcCloudRecordingResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopRtcCloudRecordingResponseBody) SetRequestId(v string) *StopRtcCloudRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRtcCloudRecordingResponseBody) SetTaskId(v string) *StopRtcCloudRecordingResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopRtcCloudRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
