// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAvailabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHostAvailabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHostAvailabilityResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHostAvailabilityResponseBody) *DeleteHostAvailabilityResponse
	GetBody() *DeleteHostAvailabilityResponseBody
}

type DeleteHostAvailabilityResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostAvailabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHostAvailabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAvailabilityResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostAvailabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHostAvailabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHostAvailabilityResponse) GetBody() *DeleteHostAvailabilityResponseBody {
	return s.Body
}

func (s *DeleteHostAvailabilityResponse) SetHeaders(v map[string]*string) *DeleteHostAvailabilityResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostAvailabilityResponse) SetStatusCode(v int32) *DeleteHostAvailabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostAvailabilityResponse) SetBody(v *DeleteHostAvailabilityResponseBody) *DeleteHostAvailabilityResponse {
	s.Body = v
	return s
}

func (s *DeleteHostAvailabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
