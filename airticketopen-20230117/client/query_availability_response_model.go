// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvailabilityResponseBody) *QueryAvailabilityResponse
	GetBody() *QueryAvailabilityResponseBody
}

type QueryAvailabilityResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *QueryAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvailabilityResponse) GetBody() *QueryAvailabilityResponseBody {
	return s.Body
}

func (s *QueryAvailabilityResponse) SetHeaders(v map[string]*string) *QueryAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *QueryAvailabilityResponse) SetStatusCode(v int32) *QueryAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvailabilityResponse) SetBody(v *QueryAvailabilityResponseBody) *QueryAvailabilityResponse {
	s.Body = v
	return s
}

func (s *QueryAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
