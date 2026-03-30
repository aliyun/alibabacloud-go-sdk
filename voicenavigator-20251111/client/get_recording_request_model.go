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
	// example:
	//
	// 868eef14-7515-4856-8a50-5c9a22abdbcc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 89ab28c2-eb94-4010-a539-f0eee922e371
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
