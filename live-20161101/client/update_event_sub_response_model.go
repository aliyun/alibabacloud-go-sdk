// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventSubResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventSubResponseBody) *UpdateEventSubResponse
	GetBody() *UpdateEventSubResponseBody
}

type UpdateEventSubResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventSubResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSubResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventSubResponse) GetBody() *UpdateEventSubResponseBody {
	return s.Body
}

func (s *UpdateEventSubResponse) SetHeaders(v map[string]*string) *UpdateEventSubResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventSubResponse) SetStatusCode(v int32) *UpdateEventSubResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventSubResponse) SetBody(v *UpdateEventSubResponseBody) *UpdateEventSubResponse {
	s.Body = v
	return s
}

func (s *UpdateEventSubResponse) Validate() error {
	return dara.Validate(s)
}
