// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventsResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventsResponseBody) *CreateEventsResponse
	GetBody() *CreateEventsResponseBody
}

type CreateEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsResponse) GoString() string {
	return s.String()
}

func (s *CreateEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventsResponse) GetBody() *CreateEventsResponseBody {
	return s.Body
}

func (s *CreateEventsResponse) SetHeaders(v map[string]*string) *CreateEventsResponse {
	s.Headers = v
	return s
}

func (s *CreateEventsResponse) SetStatusCode(v int32) *CreateEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventsResponse) SetBody(v *CreateEventsResponseBody) *CreateEventsResponse {
	s.Body = v
	return s
}

func (s *CreateEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
