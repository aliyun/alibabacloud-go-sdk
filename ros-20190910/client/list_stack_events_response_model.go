// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackEventsResponse
	GetStatusCode() *int32
	SetBody(v *ListStackEventsResponseBody) *ListStackEventsResponse
	GetBody() *ListStackEventsResponseBody
}

type ListStackEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackEventsResponse) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackEventsResponse) GetBody() *ListStackEventsResponseBody {
	return s.Body
}

func (s *ListStackEventsResponse) SetHeaders(v map[string]*string) *ListStackEventsResponse {
	s.Headers = v
	return s
}

func (s *ListStackEventsResponse) SetStatusCode(v int32) *ListStackEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackEventsResponse) SetBody(v *ListStackEventsResponseBody) *ListStackEventsResponse {
	s.Body = v
	return s
}

func (s *ListStackEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
