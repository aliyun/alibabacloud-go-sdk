// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceConfigResponseBody) *UpdateServiceConfigResponse
	GetBody() *UpdateServiceConfigResponseBody
}

type UpdateServiceConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceConfigResponse) GetBody() *UpdateServiceConfigResponseBody {
	return s.Body
}

func (s *UpdateServiceConfigResponse) SetHeaders(v map[string]*string) *UpdateServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceConfigResponse) SetStatusCode(v int32) *UpdateServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceConfigResponse) SetBody(v *UpdateServiceConfigResponseBody) *UpdateServiceConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceConfigResponse) Validate() error {
	return dara.Validate(s)
}
