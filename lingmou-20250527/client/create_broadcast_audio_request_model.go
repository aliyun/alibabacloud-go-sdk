// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBroadcastAudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateBroadcastAudioRequest
	GetFileName() *string
	SetOssKey(v string) *CreateBroadcastAudioRequest
	GetOssKey() *string
}

type CreateBroadcastAudioRequest struct {
	// example:
	//
	// audio.mp3
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// material/INPUT_BROADCAST_INFER_AUDIO/Mt.CPRLVQRR27YU2
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s CreateBroadcastAudioRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBroadcastAudioRequest) GoString() string {
	return s.String()
}

func (s *CreateBroadcastAudioRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateBroadcastAudioRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *CreateBroadcastAudioRequest) SetFileName(v string) *CreateBroadcastAudioRequest {
	s.FileName = &v
	return s
}

func (s *CreateBroadcastAudioRequest) SetOssKey(v string) *CreateBroadcastAudioRequest {
	s.OssKey = &v
	return s
}

func (s *CreateBroadcastAudioRequest) Validate() error {
	return dara.Validate(s)
}
