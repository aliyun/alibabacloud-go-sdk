// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudRecordingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRtcCloudRecordingResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartRtcCloudRecordingResponseBody
	GetTaskId() *string
}

type StartRtcCloudRecordingResponseBody struct {
	// example:
	//
	// ******58-5876-****-83CA-B56278******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartRtcCloudRecordingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingResponseBody) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRtcCloudRecordingResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartRtcCloudRecordingResponseBody) SetRequestId(v string) *StartRtcCloudRecordingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRtcCloudRecordingResponseBody) SetTaskId(v string) *StartRtcCloudRecordingResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartRtcCloudRecordingResponseBody) Validate() error {
	return dara.Validate(s)
}
