// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartConferenceRequest
	GetInstanceId() *string
	SetParticipantListJson(v string) *StartConferenceRequest
	GetParticipantListJson() *string
	SetTags(v string) *StartConferenceRequest
	GetTags() *string
	SetTimeoutSeconds(v int32) *StartConferenceRequest
	GetTimeoutSeconds() *int32
	SetUserId(v string) *StartConferenceRequest
	GetUserId() *string
}

type StartConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["8045****","1317511****"]
	ParticipantListJson *string `json:"ParticipantListJson,omitempty" xml:"ParticipantListJson,omitempty"`
	Tags                *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// example:
	//
	// 30
	TimeoutSeconds *int32 `json:"TimeoutSeconds,omitempty" xml:"TimeoutSeconds,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartConferenceRequest) GoString() string {
	return s.String()
}

func (s *StartConferenceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartConferenceRequest) GetParticipantListJson() *string {
	return s.ParticipantListJson
}

func (s *StartConferenceRequest) GetTags() *string {
	return s.Tags
}

func (s *StartConferenceRequest) GetTimeoutSeconds() *int32 {
	return s.TimeoutSeconds
}

func (s *StartConferenceRequest) GetUserId() *string {
	return s.UserId
}

func (s *StartConferenceRequest) SetInstanceId(v string) *StartConferenceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartConferenceRequest) SetParticipantListJson(v string) *StartConferenceRequest {
	s.ParticipantListJson = &v
	return s
}

func (s *StartConferenceRequest) SetTags(v string) *StartConferenceRequest {
	s.Tags = &v
	return s
}

func (s *StartConferenceRequest) SetTimeoutSeconds(v int32) *StartConferenceRequest {
	s.TimeoutSeconds = &v
	return s
}

func (s *StartConferenceRequest) SetUserId(v string) *StartConferenceRequest {
	s.UserId = &v
	return s
}

func (s *StartConferenceRequest) Validate() error {
	return dara.Validate(s)
}
