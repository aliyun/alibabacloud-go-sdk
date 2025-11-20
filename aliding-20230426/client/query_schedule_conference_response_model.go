// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryScheduleConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryScheduleConferenceResponse
	GetStatusCode() *int32
	SetBody(v *QueryScheduleConferenceResponseBody) *QueryScheduleConferenceResponse
	GetBody() *QueryScheduleConferenceResponseBody
}

type QueryScheduleConferenceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryScheduleConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryScheduleConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceResponse) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryScheduleConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryScheduleConferenceResponse) GetBody() *QueryScheduleConferenceResponseBody {
	return s.Body
}

func (s *QueryScheduleConferenceResponse) SetHeaders(v map[string]*string) *QueryScheduleConferenceResponse {
	s.Headers = v
	return s
}

func (s *QueryScheduleConferenceResponse) SetStatusCode(v int32) *QueryScheduleConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryScheduleConferenceResponse) SetBody(v *QueryScheduleConferenceResponseBody) *QueryScheduleConferenceResponse {
	s.Body = v
	return s
}

func (s *QueryScheduleConferenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
