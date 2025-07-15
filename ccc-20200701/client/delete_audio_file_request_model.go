// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAudioFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioResourceId(v string) *DeleteAudioFileRequest
	GetAudioResourceId() *string
	SetInstanceId(v string) *DeleteAudioFileRequest
	GetInstanceId() *string
}

type DeleteAudioFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// d5cd7a94-3b6a-47d2-b7fd-0b1cd839bf77
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteAudioFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAudioFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteAudioFileRequest) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *DeleteAudioFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteAudioFileRequest) SetAudioResourceId(v string) *DeleteAudioFileRequest {
	s.AudioResourceId = &v
	return s
}

func (s *DeleteAudioFileRequest) SetInstanceId(v string) *DeleteAudioFileRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAudioFileRequest) Validate() error {
	return dara.Validate(s)
}
