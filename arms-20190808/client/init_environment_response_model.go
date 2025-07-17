// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *InitEnvironmentResponseBody) *InitEnvironmentResponse
	GetBody() *InitEnvironmentResponseBody
}

type InitEnvironmentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s InitEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *InitEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitEnvironmentResponse) GetBody() *InitEnvironmentResponseBody {
	return s.Body
}

func (s *InitEnvironmentResponse) SetHeaders(v map[string]*string) *InitEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *InitEnvironmentResponse) SetStatusCode(v int32) *InitEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *InitEnvironmentResponse) SetBody(v *InitEnvironmentResponseBody) *InitEnvironmentResponse {
	s.Body = v
	return s
}

func (s *InitEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
