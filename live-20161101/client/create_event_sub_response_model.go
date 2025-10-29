// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventSubResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventSubResponseBody) *CreateEventSubResponse
	GetBody() *CreateEventSubResponseBody
}

type CreateEventSubResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventSubResponse) GoString() string {
	return s.String()
}

func (s *CreateEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventSubResponse) GetBody() *CreateEventSubResponseBody {
	return s.Body
}

func (s *CreateEventSubResponse) SetHeaders(v map[string]*string) *CreateEventSubResponse {
	s.Headers = v
	return s
}

func (s *CreateEventSubResponse) SetStatusCode(v int32) *CreateEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventSubResponse) SetBody(v *CreateEventSubResponseBody) *CreateEventSubResponse {
	s.Body = v
	return s
}

func (s *CreateEventSubResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
