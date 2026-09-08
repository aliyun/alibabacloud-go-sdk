// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutEventsResponse
	GetStatusCode() *int32
	SetBody(v *PutEventsResponseBody) *PutEventsResponse
	GetBody() *PutEventsResponseBody
}

type PutEventsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s PutEventsResponse) GoString() string {
	return s.String()
}

func (s *PutEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutEventsResponse) GetBody() *PutEventsResponseBody {
	return s.Body
}

func (s *PutEventsResponse) SetHeaders(v map[string]*string) *PutEventsResponse {
	s.Headers = v
	return s
}

func (s *PutEventsResponse) SetStatusCode(v int32) *PutEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEventsResponse) SetBody(v *PutEventsResponseBody) *PutEventsResponse {
	s.Body = v
	return s
}

func (s *PutEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
