// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCalendarAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCalendarAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCalendarAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *QueryCalendarAvailabilityResponseBody) *QueryCalendarAvailabilityResponse
	GetBody() *QueryCalendarAvailabilityResponseBody
}

type QueryCalendarAvailabilityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCalendarAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCalendarAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCalendarAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *QueryCalendarAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCalendarAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCalendarAvailabilityResponse) GetBody() *QueryCalendarAvailabilityResponseBody {
	return s.Body
}

func (s *QueryCalendarAvailabilityResponse) SetHeaders(v map[string]*string) *QueryCalendarAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *QueryCalendarAvailabilityResponse) SetStatusCode(v int32) *QueryCalendarAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCalendarAvailabilityResponse) SetBody(v *QueryCalendarAvailabilityResponseBody) *QueryCalendarAvailabilityResponse {
	s.Body = v
	return s
}

func (s *QueryCalendarAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
