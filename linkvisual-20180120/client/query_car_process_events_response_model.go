// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCarProcessEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCarProcessEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCarProcessEventsResponse
	GetStatusCode() *int32
	SetBody(v *QueryCarProcessEventsResponseBody) *QueryCarProcessEventsResponse
	GetBody() *QueryCarProcessEventsResponseBody
}

type QueryCarProcessEventsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCarProcessEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCarProcessEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCarProcessEventsResponse) GoString() string {
	return s.String()
}

func (s *QueryCarProcessEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCarProcessEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCarProcessEventsResponse) GetBody() *QueryCarProcessEventsResponseBody {
	return s.Body
}

func (s *QueryCarProcessEventsResponse) SetHeaders(v map[string]*string) *QueryCarProcessEventsResponse {
	s.Headers = v
	return s
}

func (s *QueryCarProcessEventsResponse) SetStatusCode(v int32) *QueryCarProcessEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCarProcessEventsResponse) SetBody(v *QueryCarProcessEventsResponseBody) *QueryCarProcessEventsResponse {
	s.Body = v
	return s
}

func (s *QueryCarProcessEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
