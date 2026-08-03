// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryCalendarAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelQueryCalendarAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelQueryCalendarAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelQueryCalendarAvailabilityResponseBody) *GlobalHotelQueryCalendarAvailabilityResponse
	GetBody() *GlobalHotelQueryCalendarAvailabilityResponseBody
}

type GlobalHotelQueryCalendarAvailabilityResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelQueryCalendarAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelQueryCalendarAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryCalendarAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) GetBody() *GlobalHotelQueryCalendarAvailabilityResponseBody {
	return s.Body
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) SetHeaders(v map[string]*string) *GlobalHotelQueryCalendarAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) SetStatusCode(v int32) *GlobalHotelQueryCalendarAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) SetBody(v *GlobalHotelQueryCalendarAvailabilityResponseBody) *GlobalHotelQueryCalendarAvailabilityResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelQueryCalendarAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
