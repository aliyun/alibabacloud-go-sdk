// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse
	GetBody() *UpdateEnvironmentResponseBody
}

type UpdateEnvironmentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEnvironmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEnvironmentResponse) GetBody() *UpdateEnvironmentResponseBody {
	return s.Body
}

func (s *UpdateEnvironmentResponse) SetHeaders(v map[string]*string) *UpdateEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateEnvironmentResponse) SetStatusCode(v int32) *UpdateEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEnvironmentResponse) SetBody(v *UpdateEnvironmentResponseBody) *UpdateEnvironmentResponse {
	s.Body = v
	return s
}

func (s *UpdateEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
