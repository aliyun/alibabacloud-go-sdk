// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventResponseBody) *CreateEventResponse
	GetBody() *CreateEventResponseBody
}

type CreateEventResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventResponse) GoString() string {
	return s.String()
}

func (s *CreateEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventResponse) GetBody() *CreateEventResponseBody {
	return s.Body
}

func (s *CreateEventResponse) SetHeaders(v map[string]*string) *CreateEventResponse {
	s.Headers = v
	return s
}

func (s *CreateEventResponse) SetStatusCode(v int32) *CreateEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventResponse) SetBody(v *CreateEventResponseBody) *CreateEventResponse {
	s.Body = v
	return s
}

func (s *CreateEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
