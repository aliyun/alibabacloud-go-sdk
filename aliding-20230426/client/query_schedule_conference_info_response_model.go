// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryScheduleConferenceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryScheduleConferenceInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryScheduleConferenceInfoResponseBody) *QueryScheduleConferenceInfoResponse
	GetBody() *QueryScheduleConferenceInfoResponseBody
}

type QueryScheduleConferenceInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConferenceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConferenceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryScheduleConferenceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryScheduleConferenceInfoResponse) GetBody() *QueryScheduleConferenceInfoResponseBody {
	return s.Body
}

func (s *QueryScheduleConferenceInfoResponse) SetHeaders(v map[string]*string) *QueryScheduleConferenceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConferenceInfoResponse) SetStatusCode(v int32) *QueryScheduleConferenceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConferenceInfoResponse) SetBody(v *QueryScheduleConferenceInfoResponseBody) *QueryScheduleConferenceInfoResponse {
	s.Body = v
	return s
}

func (s *QueryScheduleConferenceInfoResponse) Validate() error {
	return dara.Validate(s)
}
