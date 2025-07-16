// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateScheduleConferenceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScheduleConferenceResponseBody
	GetSuccess() *bool
}

type UpdateScheduleConferenceResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScheduleConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduleConferenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScheduleConferenceResponseBody) SetRequestId(v string) *UpdateScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleConferenceResponseBody) SetSuccess(v bool) *UpdateScheduleConferenceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScheduleConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
