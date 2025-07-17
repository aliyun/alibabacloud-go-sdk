// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse
	GetBody() *GetEnvironmentResponseBody
}

type GetEnvironmentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnvironmentResponse) GetBody() *GetEnvironmentResponseBody {
	return s.Body
}

func (s *GetEnvironmentResponse) SetHeaders(v map[string]*string) *GetEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentResponse) SetStatusCode(v int32) *GetEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentResponse) SetBody(v *GetEnvironmentResponseBody) *GetEnvironmentResponse {
	s.Body = v
	return s
}

func (s *GetEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
