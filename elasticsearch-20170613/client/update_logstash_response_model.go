// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogstashResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogstashResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogstashResponseBody) *UpdateLogstashResponse
	GetBody() *UpdateLogstashResponseBody
}

type UpdateLogstashResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogstashResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogstashResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogstashResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogstashResponse) GetBody() *UpdateLogstashResponseBody {
	return s.Body
}

func (s *UpdateLogstashResponse) SetHeaders(v map[string]*string) *UpdateLogstashResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashResponse) SetStatusCode(v int32) *UpdateLogstashResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogstashResponse) SetBody(v *UpdateLogstashResponseBody) *UpdateLogstashResponse {
	s.Body = v
	return s
}

func (s *UpdateLogstashResponse) Validate() error {
	return dara.Validate(s)
}
