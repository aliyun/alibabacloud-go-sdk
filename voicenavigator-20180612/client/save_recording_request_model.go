// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *SaveRecordingRequest
	GetConversationId() *string
	SetDuration(v string) *SaveRecordingRequest
	GetDuration() *string
	SetFileName(v string) *SaveRecordingRequest
	GetFileName() *string
	SetFilePath(v string) *SaveRecordingRequest
	GetFilePath() *string
	SetInstanceId(v string) *SaveRecordingRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *SaveRecordingRequest
	GetInstanceOwnerId() *int64
	SetStartTime(v int64) *SaveRecordingRequest
	GetStartTime() *int64
	SetType(v string) *SaveRecordingRequest
	GetType() *string
	SetVoiceSliceRecordingList(v string) *SaveRecordingRequest
	GetVoiceSliceRecordingList() *string
}

type SaveRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e6bef0db439d4048bfcf45322491becf.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://test/record/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1971226538081821
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Source
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VoiceSliceRecordingList *string `json:"VoiceSliceRecordingList,omitempty" xml:"VoiceSliceRecordingList,omitempty"`
}

func (s SaveRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveRecordingRequest) GoString() string {
	return s.String()
}

func (s *SaveRecordingRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *SaveRecordingRequest) GetDuration() *string {
	return s.Duration
}

func (s *SaveRecordingRequest) GetFileName() *string {
	return s.FileName
}

func (s *SaveRecordingRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *SaveRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveRecordingRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *SaveRecordingRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SaveRecordingRequest) GetType() *string {
	return s.Type
}

func (s *SaveRecordingRequest) GetVoiceSliceRecordingList() *string {
	return s.VoiceSliceRecordingList
}

func (s *SaveRecordingRequest) SetConversationId(v string) *SaveRecordingRequest {
	s.ConversationId = &v
	return s
}

func (s *SaveRecordingRequest) SetDuration(v string) *SaveRecordingRequest {
	s.Duration = &v
	return s
}

func (s *SaveRecordingRequest) SetFileName(v string) *SaveRecordingRequest {
	s.FileName = &v
	return s
}

func (s *SaveRecordingRequest) SetFilePath(v string) *SaveRecordingRequest {
	s.FilePath = &v
	return s
}

func (s *SaveRecordingRequest) SetInstanceId(v string) *SaveRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveRecordingRequest) SetInstanceOwnerId(v int64) *SaveRecordingRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *SaveRecordingRequest) SetStartTime(v int64) *SaveRecordingRequest {
	s.StartTime = &v
	return s
}

func (s *SaveRecordingRequest) SetType(v string) *SaveRecordingRequest {
	s.Type = &v
	return s
}

func (s *SaveRecordingRequest) SetVoiceSliceRecordingList(v string) *SaveRecordingRequest {
	s.VoiceSliceRecordingList = &v
	return s
}

func (s *SaveRecordingRequest) Validate() error {
	return dara.Validate(s)
}
