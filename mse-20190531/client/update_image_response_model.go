// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageResponseBody) *UpdateImageResponse
	GetBody() *UpdateImageResponseBody
}

type UpdateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageResponse) GetBody() *UpdateImageResponseBody {
	return s.Body
}

func (s *UpdateImageResponse) SetHeaders(v map[string]*string) *UpdateImageResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageResponse) SetStatusCode(v int32) *UpdateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageResponse) SetBody(v *UpdateImageResponseBody) *UpdateImageResponse {
	s.Body = v
	return s
}

func (s *UpdateImageResponse) Validate() error {
	return dara.Validate(s)
}
