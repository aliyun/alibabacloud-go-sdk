// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingBusinessOptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventStreamingBusinessOptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventStreamingBusinessOptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventStreamingBusinessOptionResponseBody) *UpdateEventStreamingBusinessOptionResponse
	GetBody() *UpdateEventStreamingBusinessOptionResponseBody
}

type UpdateEventStreamingBusinessOptionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventStreamingBusinessOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventStreamingBusinessOptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingBusinessOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingBusinessOptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventStreamingBusinessOptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventStreamingBusinessOptionResponse) GetBody() *UpdateEventStreamingBusinessOptionResponseBody {
	return s.Body
}

func (s *UpdateEventStreamingBusinessOptionResponse) SetHeaders(v map[string]*string) *UpdateEventStreamingBusinessOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponse) SetStatusCode(v int32) *UpdateEventStreamingBusinessOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponse) SetBody(v *UpdateEventStreamingBusinessOptionResponseBody) *UpdateEventStreamingBusinessOptionResponse {
	s.Body = v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
