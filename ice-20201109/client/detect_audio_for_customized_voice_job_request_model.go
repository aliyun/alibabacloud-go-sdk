// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectAudioForCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioRecordId(v int32) *DetectAudioForCustomizedVoiceJobRequest
	GetAudioRecordId() *int32
	SetRecordUrl(v string) *DetectAudioForCustomizedVoiceJobRequest
	GetRecordUrl() *string
	SetVoiceId(v string) *DetectAudioForCustomizedVoiceJobRequest
	GetVoiceId() *string
}

type DetectAudioForCustomizedVoiceJobRequest struct {
	// The sequence number of the recording file.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AudioRecordId *int32 `json:"AudioRecordId,omitempty" xml:"AudioRecordId,omitempty"`
	// The URL of the recording file.
	//
	// > : The URL must be an Object Storage Service (OSS) URL within your Alibaba Cloud account. The OSS bucket must be in the same region in which IMS is activated.
	//
	// > : The audio file must be in the WAV or PCM format and must be a 16-bit mono audio file at 48000 Hz.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://your-bucket.oss-cn-hangzhou.aliyuncs.com/record1.wav
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The voice ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s DetectAudioForCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectAudioForCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *DetectAudioForCustomizedVoiceJobRequest) GetAudioRecordId() *int32 {
	return s.AudioRecordId
}

func (s *DetectAudioForCustomizedVoiceJobRequest) GetRecordUrl() *string {
	return s.RecordUrl
}

func (s *DetectAudioForCustomizedVoiceJobRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *DetectAudioForCustomizedVoiceJobRequest) SetAudioRecordId(v int32) *DetectAudioForCustomizedVoiceJobRequest {
	s.AudioRecordId = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobRequest) SetRecordUrl(v string) *DetectAudioForCustomizedVoiceJobRequest {
	s.RecordUrl = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobRequest) SetVoiceId(v string) *DetectAudioForCustomizedVoiceJobRequest {
	s.VoiceId = &v
	return s
}

func (s *DetectAudioForCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
