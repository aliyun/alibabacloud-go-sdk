// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEventsViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEventsViewResponse
	GetStatusCode() *int32
	SetBody(v *ListEventsViewResponseBody) *ListEventsViewResponse
	GetBody() *ListEventsViewResponseBody
}

type ListEventsViewResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventsViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventsViewResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponse) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEventsViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEventsViewResponse) GetBody() *ListEventsViewResponseBody {
	return s.Body
}

func (s *ListEventsViewResponse) SetHeaders(v map[string]*string) *ListEventsViewResponse {
	s.Headers = v
	return s
}

func (s *ListEventsViewResponse) SetStatusCode(v int32) *ListEventsViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventsViewResponse) SetBody(v *ListEventsViewResponseBody) *ListEventsViewResponse {
	s.Body = v
	return s
}

func (s *ListEventsViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
