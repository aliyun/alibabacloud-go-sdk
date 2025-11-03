// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventSourceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventSourceResponseBody) *UpdateEventSourceResponse
	GetBody() *UpdateEventSourceResponseBody
}

type UpdateEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventSourceResponse) GetBody() *UpdateEventSourceResponseBody {
	return s.Body
}

func (s *UpdateEventSourceResponse) SetHeaders(v map[string]*string) *UpdateEventSourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventSourceResponse) SetStatusCode(v int32) *UpdateEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventSourceResponse) SetBody(v *UpdateEventSourceResponseBody) *UpdateEventSourceResponse {
	s.Body = v
	return s
}

func (s *UpdateEventSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
