// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GlobalHotelQueryAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GlobalHotelQueryAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *GlobalHotelQueryAvailabilityResponseBody) *GlobalHotelQueryAvailabilityResponse
	GetBody() *GlobalHotelQueryAvailabilityResponseBody
}

type GlobalHotelQueryAvailabilityResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalHotelQueryAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalHotelQueryAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GlobalHotelQueryAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GlobalHotelQueryAvailabilityResponse) GetBody() *GlobalHotelQueryAvailabilityResponseBody {
	return s.Body
}

func (s *GlobalHotelQueryAvailabilityResponse) SetHeaders(v map[string]*string) *GlobalHotelQueryAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponse) SetStatusCode(v int32) *GlobalHotelQueryAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponse) SetBody(v *GlobalHotelQueryAvailabilityResponseBody) *GlobalHotelQueryAvailabilityResponse {
	s.Body = v
	return s
}

func (s *GlobalHotelQueryAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
