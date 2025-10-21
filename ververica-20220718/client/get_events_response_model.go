// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventsResponse
	GetStatusCode() *int32
	SetBody(v *GetEventsResponseBody) *GetEventsResponse
	GetBody() *GetEventsResponseBody
}

type GetEventsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventsResponse) GoString() string {
	return s.String()
}

func (s *GetEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventsResponse) GetBody() *GetEventsResponseBody {
	return s.Body
}

func (s *GetEventsResponse) SetHeaders(v map[string]*string) *GetEventsResponse {
	s.Headers = v
	return s
}

func (s *GetEventsResponse) SetStatusCode(v int32) *GetEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventsResponse) SetBody(v *GetEventsResponseBody) *GetEventsResponse {
	s.Body = v
	return s
}

func (s *GetEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
