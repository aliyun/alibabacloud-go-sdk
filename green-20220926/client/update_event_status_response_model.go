// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventStatusResponseBody) *UpdateEventStatusResponse
	GetBody() *UpdateEventStatusResponseBody
}

type UpdateEventStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventStatusResponse) GetBody() *UpdateEventStatusResponseBody {
	return s.Body
}

func (s *UpdateEventStatusResponse) SetHeaders(v map[string]*string) *UpdateEventStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventStatusResponse) SetStatusCode(v int32) *UpdateEventStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventStatusResponse) SetBody(v *UpdateEventStatusResponseBody) *UpdateEventStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateEventStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
