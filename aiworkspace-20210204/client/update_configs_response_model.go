// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigsResponseBody) *UpdateConfigsResponse
	GetBody() *UpdateConfigsResponseBody
}

type UpdateConfigsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigsResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigsResponse) GetBody() *UpdateConfigsResponseBody {
	return s.Body
}

func (s *UpdateConfigsResponse) SetHeaders(v map[string]*string) *UpdateConfigsResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigsResponse) SetStatusCode(v int32) *UpdateConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigsResponse) SetBody(v *UpdateConfigsResponseBody) *UpdateConfigsResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigsResponse) Validate() error {
	return dara.Validate(s)
}
