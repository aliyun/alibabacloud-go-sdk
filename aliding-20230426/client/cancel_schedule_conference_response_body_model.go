// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelScheduleConferenceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelScheduleConferenceResponseBody
	GetSuccess() *bool
}

type CancelScheduleConferenceResponseBody struct {
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

func (s CancelScheduleConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelScheduleConferenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelScheduleConferenceResponseBody) SetRequestId(v string) *CancelScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelScheduleConferenceResponseBody) SetSuccess(v bool) *CancelScheduleConferenceResponseBody {
	s.Success = &v
	return s
}

func (s *CancelScheduleConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
