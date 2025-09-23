// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExtensionsResponseBody) *UpdateExtensionsResponse
	GetBody() *UpdateExtensionsResponseBody
}

type UpdateExtensionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtensionsResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExtensionsResponse) GetBody() *UpdateExtensionsResponseBody {
	return s.Body
}

func (s *UpdateExtensionsResponse) SetHeaders(v map[string]*string) *UpdateExtensionsResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtensionsResponse) SetStatusCode(v int32) *UpdateExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExtensionsResponse) SetBody(v *UpdateExtensionsResponseBody) *UpdateExtensionsResponse {
	s.Body = v
	return s
}

func (s *UpdateExtensionsResponse) Validate() error {
	return dara.Validate(s)
}
