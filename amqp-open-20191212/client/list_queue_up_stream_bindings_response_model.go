// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueUpStreamBindingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueueUpStreamBindingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueueUpStreamBindingsResponse
	GetStatusCode() *int32
	SetBody(v *ListQueueUpStreamBindingsResponseBody) *ListQueueUpStreamBindingsResponse
	GetBody() *ListQueueUpStreamBindingsResponseBody
}

type ListQueueUpStreamBindingsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueUpStreamBindingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueUpStreamBindingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueueUpStreamBindingsResponse) GoString() string {
	return s.String()
}

func (s *ListQueueUpStreamBindingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueueUpStreamBindingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueueUpStreamBindingsResponse) GetBody() *ListQueueUpStreamBindingsResponseBody {
	return s.Body
}

func (s *ListQueueUpStreamBindingsResponse) SetHeaders(v map[string]*string) *ListQueueUpStreamBindingsResponse {
	s.Headers = v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetStatusCode(v int32) *ListQueueUpStreamBindingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) SetBody(v *ListQueueUpStreamBindingsResponseBody) *ListQueueUpStreamBindingsResponse {
	s.Body = v
	return s
}

func (s *ListQueueUpStreamBindingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
