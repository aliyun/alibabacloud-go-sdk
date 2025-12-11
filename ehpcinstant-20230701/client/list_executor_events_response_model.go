// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExecutorEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExecutorEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListExecutorEventsResponseBody) *ListExecutorEventsResponse
	GetBody() *ListExecutorEventsResponseBody
}

type ListExecutorEventsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutorEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutorEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorEventsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutorEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExecutorEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExecutorEventsResponse) GetBody() *ListExecutorEventsResponseBody {
	return s.Body
}

func (s *ListExecutorEventsResponse) SetHeaders(v map[string]*string) *ListExecutorEventsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutorEventsResponse) SetStatusCode(v int32) *ListExecutorEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutorEventsResponse) SetBody(v *ListExecutorEventsResponseBody) *ListExecutorEventsResponse {
	s.Body = v
	return s
}

func (s *ListExecutorEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
