// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceImageResponseBody) *UpdateInstanceImageResponse
	GetBody() *UpdateInstanceImageResponseBody
}

type UpdateInstanceImageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceImageResponse) GetBody() *UpdateInstanceImageResponseBody {
	return s.Body
}

func (s *UpdateInstanceImageResponse) SetHeaders(v map[string]*string) *UpdateInstanceImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceImageResponse) SetStatusCode(v int32) *UpdateInstanceImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceImageResponse) SetBody(v *UpdateInstanceImageResponseBody) *UpdateInstanceImageResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceImageResponse) Validate() error {
	return dara.Validate(s)
}
