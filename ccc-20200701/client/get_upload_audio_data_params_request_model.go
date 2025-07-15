// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadAudioDataParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetUploadAudioDataParamsRequest
	GetContactId() *string
	SetInstanceId(v string) *GetUploadAudioDataParamsRequest
	GetInstanceId() *string
}

type GetUploadAudioDataParamsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-38860977107324****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetUploadAudioDataParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadAudioDataParamsRequest) GoString() string {
	return s.String()
}

func (s *GetUploadAudioDataParamsRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetUploadAudioDataParamsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUploadAudioDataParamsRequest) SetContactId(v string) *GetUploadAudioDataParamsRequest {
	s.ContactId = &v
	return s
}

func (s *GetUploadAudioDataParamsRequest) SetInstanceId(v string) *GetUploadAudioDataParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUploadAudioDataParamsRequest) Validate() error {
	return dara.Validate(s)
}
