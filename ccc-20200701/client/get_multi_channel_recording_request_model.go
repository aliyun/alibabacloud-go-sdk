// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiChannelRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetMultiChannelRecordingRequest
	GetContactId() *string
	SetInstanceId(v string) *GetMultiChannelRecordingRequest
	GetInstanceId() *string
}

type GetMultiChannelRecordingRequest struct {
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

func (s GetMultiChannelRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiChannelRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetMultiChannelRecordingRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetMultiChannelRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMultiChannelRecordingRequest) SetContactId(v string) *GetMultiChannelRecordingRequest {
	s.ContactId = &v
	return s
}

func (s *GetMultiChannelRecordingRequest) SetInstanceId(v string) *GetMultiChannelRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMultiChannelRecordingRequest) Validate() error {
	return dara.Validate(s)
}
