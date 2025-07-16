// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceId(v string) *QueryConferenceInfoRequest
	GetConferenceId() *string
}

type QueryConferenceInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryConferenceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryConferenceInfoRequest) SetConferenceId(v string) *QueryConferenceInfoRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryConferenceInfoRequest) Validate() error {
	return dara.Validate(s)
}
