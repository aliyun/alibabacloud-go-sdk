// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonoRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *GetMonoRecordingRequest
	GetContactId() *string
	SetExpireSeconds(v int64) *GetMonoRecordingRequest
	GetExpireSeconds() *int64
	SetInstanceId(v string) *GetMonoRecordingRequest
	GetInstanceId() *string
}

type GetMonoRecordingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103689****
	ContactId     *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ExpireSeconds *int64  `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMonoRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMonoRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetMonoRecordingRequest) GetContactId() *string {
	return s.ContactId
}

func (s *GetMonoRecordingRequest) GetExpireSeconds() *int64 {
	return s.ExpireSeconds
}

func (s *GetMonoRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMonoRecordingRequest) SetContactId(v string) *GetMonoRecordingRequest {
	s.ContactId = &v
	return s
}

func (s *GetMonoRecordingRequest) SetExpireSeconds(v int64) *GetMonoRecordingRequest {
	s.ExpireSeconds = &v
	return s
}

func (s *GetMonoRecordingRequest) SetInstanceId(v string) *GetMonoRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMonoRecordingRequest) Validate() error {
	return dara.Validate(s)
}
