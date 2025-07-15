// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVoicemailRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetVoicemailRecordingRequest
	GetContactId() *string
	SetInstanceId(v string) *GetVoicemailRecordingRequest
	GetInstanceId() *string
}

type GetVoicemailRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-12515239414412****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetVoicemailRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVoicemailRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetVoicemailRecordingRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetVoicemailRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVoicemailRecordingRequest) SetContactId(v string) *GetVoicemailRecordingRequest {
	s.ContactId = &v
	return s
}

func (s *GetVoicemailRecordingRequest) SetInstanceId(v string) *GetVoicemailRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVoicemailRecordingRequest) Validate() error {
	return dara.Validate(s)
}
