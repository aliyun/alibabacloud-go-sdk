// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioResourceId(v string) *GetAudioFileRequest
	GetAudioResourceId() *string
	SetInstanceId(v string) *GetAudioFileRequest
	GetInstanceId() *string
}

type GetAudioFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c1a06b46-302a-4c6e-928b-a43c0df485cf
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAudioFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileRequest) GoString() string {
	return s.String()
}

func (s *GetAudioFileRequest) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *GetAudioFileRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAudioFileRequest) SetAudioResourceId(v string) *GetAudioFileRequest {
	s.AudioResourceId = &v
	return s
}

func (s *GetAudioFileRequest) SetInstanceId(v string) *GetAudioFileRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAudioFileRequest) Validate() error {
	return dara.Validate(s)
}
