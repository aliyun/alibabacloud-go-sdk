// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceId(v string) *CreateVideoConferenceResponseBody
	GetConferenceId() *string
	SetConferencePassword(v string) *CreateVideoConferenceResponseBody
	GetConferencePassword() *string
	SetExternalLinkUrl(v string) *CreateVideoConferenceResponseBody
	GetExternalLinkUrl() *string
	SetHostPassword(v string) *CreateVideoConferenceResponseBody
	GetHostPassword() *string
	SetPhoneNumbers(v []*string) *CreateVideoConferenceResponseBody
	GetPhoneNumbers() []*string
	SetRequestId(v string) *CreateVideoConferenceResponseBody
	GetRequestId() *string
	SetRoomCode(v string) *CreateVideoConferenceResponseBody
	GetRoomCode() *string
}

type CreateVideoConferenceResponseBody struct {
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 1151302294
	ConferencePassword *string `json:"conferencePassword,omitempty" xml:"conferencePassword,omitempty"`
	// example:
	//
	// https://pre-meeting.dingtalk.com/app?roomCode=68550708396&token=1_59209c43-431c-4e57-a0f8-11bebdb3db7f
	ExternalLinkUrl *string `json:"externalLinkUrl,omitempty" xml:"externalLinkUrl,omitempty"`
	// example:
	//
	// 2142817614
	HostPassword *string   `json:"hostPassword,omitempty" xml:"hostPassword,omitempty"`
	PhoneNumbers []*string `json:"phoneNumbers,omitempty" xml:"phoneNumbers,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 3032809F-8C14-57E2-9B76-7AC2134FE3C8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 123
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
}

func (s CreateVideoConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceResponseBody) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *CreateVideoConferenceResponseBody) GetConferencePassword() *string {
	return s.ConferencePassword
}

func (s *CreateVideoConferenceResponseBody) GetExternalLinkUrl() *string {
	return s.ExternalLinkUrl
}

func (s *CreateVideoConferenceResponseBody) GetHostPassword() *string {
	return s.HostPassword
}

func (s *CreateVideoConferenceResponseBody) GetPhoneNumbers() []*string {
	return s.PhoneNumbers
}

func (s *CreateVideoConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoConferenceResponseBody) GetRoomCode() *string {
	return s.RoomCode
}

func (s *CreateVideoConferenceResponseBody) SetConferenceId(v string) *CreateVideoConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetConferencePassword(v string) *CreateVideoConferenceResponseBody {
	s.ConferencePassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetExternalLinkUrl(v string) *CreateVideoConferenceResponseBody {
	s.ExternalLinkUrl = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetHostPassword(v string) *CreateVideoConferenceResponseBody {
	s.HostPassword = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetPhoneNumbers(v []*string) *CreateVideoConferenceResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetRequestId(v string) *CreateVideoConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) SetRoomCode(v string) *CreateVideoConferenceResponseBody {
	s.RoomCode = &v
	return s
}

func (s *CreateVideoConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
