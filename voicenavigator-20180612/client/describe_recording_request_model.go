// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DescribeRecordingRequest
	GetConversationId() *string
	SetInstanceId(v string) *DescribeRecordingRequest
	GetInstanceId() *string
	SetNeedVoiceSliceRecording(v bool) *DescribeRecordingRequest
	GetNeedVoiceSliceRecording() *bool
}

type DescribeRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	InstanceId              *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedVoiceSliceRecording *bool   `json:"NeedVoiceSliceRecording,omitempty" xml:"NeedVoiceSliceRecording,omitempty"`
}

func (s DescribeRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordingRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DescribeRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRecordingRequest) GetNeedVoiceSliceRecording() *bool {
	return s.NeedVoiceSliceRecording
}

func (s *DescribeRecordingRequest) SetConversationId(v string) *DescribeRecordingRequest {
	s.ConversationId = &v
	return s
}

func (s *DescribeRecordingRequest) SetInstanceId(v string) *DescribeRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRecordingRequest) SetNeedVoiceSliceRecording(v bool) *DescribeRecordingRequest {
	s.NeedVoiceSliceRecording = &v
	return s
}

func (s *DescribeRecordingRequest) Validate() error {
	return dara.Validate(s)
}
