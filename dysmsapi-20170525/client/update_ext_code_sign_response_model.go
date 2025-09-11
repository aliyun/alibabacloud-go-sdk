// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtCodeSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExtCodeSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExtCodeSignResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExtCodeSignResponseBody) *UpdateExtCodeSignResponse
	GetBody() *UpdateExtCodeSignResponseBody
}

type UpdateExtCodeSignResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExtCodeSignResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExtCodeSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExtCodeSignResponse) GetBody() *UpdateExtCodeSignResponseBody {
	return s.Body
}

func (s *UpdateExtCodeSignResponse) SetHeaders(v map[string]*string) *UpdateExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *UpdateExtCodeSignResponse) SetStatusCode(v int32) *UpdateExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExtCodeSignResponse) SetBody(v *UpdateExtCodeSignResponseBody) *UpdateExtCodeSignResponse {
	s.Body = v
	return s
}

func (s *UpdateExtCodeSignResponse) Validate() error {
	return dara.Validate(s)
}
