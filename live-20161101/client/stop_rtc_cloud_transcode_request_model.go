// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcCloudTranscodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StopRtcCloudTranscodeRequest
	GetAppId() *string
	SetTaskId(v string) *StopRtcCloudTranscodeRequest
	GetTaskId() *string
}

type StopRtcCloudTranscodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ******73-8501-****-8ac1-72295a******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRtcCloudTranscodeRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRtcCloudTranscodeRequest) GoString() string {
	return s.String()
}

func (s *StopRtcCloudTranscodeRequest) GetAppId() *string {
	return s.AppId
}

func (s *StopRtcCloudTranscodeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopRtcCloudTranscodeRequest) SetAppId(v string) *StopRtcCloudTranscodeRequest {
	s.AppId = &v
	return s
}

func (s *StopRtcCloudTranscodeRequest) SetTaskId(v string) *StopRtcCloudTranscodeRequest {
	s.TaskId = &v
	return s
}

func (s *StopRtcCloudTranscodeRequest) Validate() error {
	return dara.Validate(s)
}
