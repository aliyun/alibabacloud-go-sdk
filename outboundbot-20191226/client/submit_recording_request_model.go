// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SubmitRecordingRequest
	GetInstanceId() *string
	SetMergedRecording(v string) *SubmitRecordingRequest
	GetMergedRecording() *string
	SetResourceRecording(v string) *SubmitRecordingRequest
	GetResourceRecording() *string
	SetTaskId(v string) *SubmitRecordingRequest
	GetTaskId() *string
}

type SubmitRecordingRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script recording data
	//
	// This parameter is required.
	//
	// example:
	//
	// {"contactId":"3d35c0487cc041abb7ad0ce61752601f","duration":27,"fileName":"ce2659e5-a20b-4f8e-91b5-5cd909c6b96e_3d35c0487cc041abb7ad0ce61752601f.wav","filePath":"oss://ForCompatibility/waveforms/","startTime":1579057583670,"type":"Merged"}
	MergedRecording *string `json:"MergedRecording,omitempty" xml:"MergedRecording,omitempty"`
	// Script recording data
	//
	// example:
	//
	// {"contactId":"3d35c0487cc041abb7ad0ce61752601f","duration":27,"fileName":"ce2659e5-a20b-4f8e-91b5-5cd909c6b96e_3d35c0487cc041abb7ad0ce61752601f.wav","filePath":"oss://ForCompatibility/waveforms/","startTime":1579057583670,"type":"Source"}
	ResourceRecording *string `json:"ResourceRecording,omitempty" xml:"ResourceRecording,omitempty"`
	// Job ID
	//
	// This parameter is required.
	//
	// example:
	//
	// e4e2a770-b97b-465a-80d8-06dca008c503
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitRecordingRequest) GoString() string {
	return s.String()
}

func (s *SubmitRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SubmitRecordingRequest) GetMergedRecording() *string {
	return s.MergedRecording
}

func (s *SubmitRecordingRequest) GetResourceRecording() *string {
	return s.ResourceRecording
}

func (s *SubmitRecordingRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitRecordingRequest) SetInstanceId(v string) *SubmitRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *SubmitRecordingRequest) SetMergedRecording(v string) *SubmitRecordingRequest {
	s.MergedRecording = &v
	return s
}

func (s *SubmitRecordingRequest) SetResourceRecording(v string) *SubmitRecordingRequest {
	s.ResourceRecording = &v
	return s
}

func (s *SubmitRecordingRequest) SetTaskId(v string) *SubmitRecordingRequest {
	s.TaskId = &v
	return s
}

func (s *SubmitRecordingRequest) Validate() error {
	return dara.Validate(s)
}
