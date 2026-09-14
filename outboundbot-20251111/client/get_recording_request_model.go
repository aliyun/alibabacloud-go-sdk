// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRecordingRequest
	GetInstanceId() *string
	SetSessionId(v string) *GetRecordingRequest
	GetSessionId() *string
}

type GetRecordingRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 1dcb09c5-d5db-4397-bf65-db854463beea
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// job-b4c0d17b-e871-498f-b6c7-5219bc22b87a
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordingRequest) GoString() string {
	return s.String()
}

func (s *GetRecordingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecordingRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetRecordingRequest) SetInstanceId(v string) *GetRecordingRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecordingRequest) SetSessionId(v string) *GetRecordingRequest {
	s.SessionId = &v
	return s
}

func (s *GetRecordingRequest) Validate() error {
	return dara.Validate(s)
}
