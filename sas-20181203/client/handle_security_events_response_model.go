// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSecurityEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleSecurityEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleSecurityEventsResponse
	GetStatusCode() *int32
	SetBody(v *HandleSecurityEventsResponseBody) *HandleSecurityEventsResponse
	GetBody() *HandleSecurityEventsResponseBody
}

type HandleSecurityEventsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleSecurityEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleSecurityEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleSecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *HandleSecurityEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleSecurityEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleSecurityEventsResponse) GetBody() *HandleSecurityEventsResponseBody {
	return s.Body
}

func (s *HandleSecurityEventsResponse) SetHeaders(v map[string]*string) *HandleSecurityEventsResponse {
	s.Headers = v
	return s
}

func (s *HandleSecurityEventsResponse) SetStatusCode(v int32) *HandleSecurityEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleSecurityEventsResponse) SetBody(v *HandleSecurityEventsResponseBody) *HandleSecurityEventsResponse {
	s.Body = v
	return s
}

func (s *HandleSecurityEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
