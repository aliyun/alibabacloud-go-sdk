// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEarlyMediaRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetEarlyMediaRecordingRequest
	GetContactId() *string
	SetInstanceId(v string) *GetEarlyMediaRecordingRequest
	GetInstanceId() *string
}

type GetEarlyMediaRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103689****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEarlyMediaRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEarlyMediaRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetEarlyMediaRecordingRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetEarlyMediaRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEarlyMediaRecordingRequest) SetContactId(v string) *GetEarlyMediaRecordingRequest {
	s.ContactId = &v
	return s
}

func (s *GetEarlyMediaRecordingRequest) SetInstanceId(v string) *GetEarlyMediaRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEarlyMediaRecordingRequest) Validate() error {
	return dara.Validate(s)
}
