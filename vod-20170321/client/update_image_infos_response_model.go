// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageInfosResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageInfosResponseBody) *UpdateImageInfosResponse
	GetBody() *UpdateImageInfosResponseBody
}

type UpdateImageInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageInfosResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageInfosResponse) GetBody() *UpdateImageInfosResponseBody {
	return s.Body
}

func (s *UpdateImageInfosResponse) SetHeaders(v map[string]*string) *UpdateImageInfosResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageInfosResponse) SetStatusCode(v int32) *UpdateImageInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageInfosResponse) SetBody(v *UpdateImageInfosResponseBody) *UpdateImageInfosResponse {
	s.Body = v
	return s
}

func (s *UpdateImageInfosResponse) Validate() error {
	return dara.Validate(s)
}
