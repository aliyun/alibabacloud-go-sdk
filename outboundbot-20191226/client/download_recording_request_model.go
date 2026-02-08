// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DownloadRecordingRequest
	GetInstanceId() *string
	SetNeedVoiceSliceRecording(v bool) *DownloadRecordingRequest
	GetNeedVoiceSliceRecording() *bool
	SetTaskId(v string) *DownloadRecordingRequest
	GetTaskId() *string
}

type DownloadRecordingRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether to query segmented recordings
	//
	// example:
	//
	// false
	NeedVoiceSliceRecording *bool `json:"NeedVoiceSliceRecording,omitempty" xml:"NeedVoiceSliceRecording,omitempty"`
	// Call ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 744b27f3-437f-4a8c-a181-f668e492fd24
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DownloadRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadRecordingRequest) GoString() string {
	return s.String()
}

func (s *DownloadRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DownloadRecordingRequest) GetNeedVoiceSliceRecording() *bool {
	return s.NeedVoiceSliceRecording
}

func (s *DownloadRecordingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DownloadRecordingRequest) SetInstanceId(v string) *DownloadRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DownloadRecordingRequest) SetNeedVoiceSliceRecording(v bool) *DownloadRecordingRequest {
	s.NeedVoiceSliceRecording = &v
	return s
}

func (s *DownloadRecordingRequest) SetTaskId(v string) *DownloadRecordingRequest {
	s.TaskId = &v
	return s
}

func (s *DownloadRecordingRequest) Validate() error {
	return dara.Validate(s)
}
