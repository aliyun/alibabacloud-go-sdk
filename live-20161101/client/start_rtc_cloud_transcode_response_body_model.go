// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudTranscodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartRtcCloudTranscodeResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StartRtcCloudTranscodeResponseBody
	GetTaskId() *string
}

type StartRtcCloudTranscodeResponseBody struct {
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartRtcCloudTranscodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudTranscodeResponseBody) GoString() string {
	return s.String()
}

func (s *StartRtcCloudTranscodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartRtcCloudTranscodeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartRtcCloudTranscodeResponseBody) SetRequestId(v string) *StartRtcCloudTranscodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRtcCloudTranscodeResponseBody) SetTaskId(v string) *StartRtcCloudTranscodeResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartRtcCloudTranscodeResponseBody) Validate() error {
	return dara.Validate(s)
}
