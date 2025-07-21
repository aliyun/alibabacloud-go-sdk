// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse
	GetBody() *UpdateInstanceResponseBody
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceResponse) GetBody() *UpdateInstanceResponseBody {
	return s.Body
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceResponse) Validate() error {
	return dara.Validate(s)
}
