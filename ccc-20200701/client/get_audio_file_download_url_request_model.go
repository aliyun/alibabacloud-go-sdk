// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioResourceId(v string) *GetAudioFileDownloadUrlRequest
	GetAudioResourceId() *string
	SetInstanceId(v string) *GetAudioFileDownloadUrlRequest
	GetInstanceId() *string
}

type GetAudioFileDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acc300c4-75c9-41ba-ba5e-2a365c96c248
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAudioFileDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetAudioFileDownloadUrlRequest) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *GetAudioFileDownloadUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAudioFileDownloadUrlRequest) SetAudioResourceId(v string) *GetAudioFileDownloadUrlRequest {
	s.AudioResourceId = &v
	return s
}

func (s *GetAudioFileDownloadUrlRequest) SetInstanceId(v string) *GetAudioFileDownloadUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAudioFileDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
