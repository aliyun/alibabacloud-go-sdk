// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcCloudRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRtcCloudRecordingResponseBody
	GetRequestId() *string
	SetTaskId(v string) *UpdateRtcCloudRecordingResponseBody
	GetTaskId() *string
}

type UpdateRtcCloudRecordingResponseBody struct {
	// example:
	//
	// ******58-5876-****-83CA-B56278******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateRtcCloudRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcCloudRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRtcCloudRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRtcCloudRecordingResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateRtcCloudRecordingResponseBody) SetRequestId(v string) *UpdateRtcCloudRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRtcCloudRecordingResponseBody) SetTaskId(v string) *UpdateRtcCloudRecordingResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpdateRtcCloudRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
