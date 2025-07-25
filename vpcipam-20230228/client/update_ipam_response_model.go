// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpamResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpamResponseBody) *UpdateIpamResponse
	GetBody() *UpdateIpamResponseBody
}

type UpdateIpamResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpamResponse) GetBody() *UpdateIpamResponseBody {
	return s.Body
}

func (s *UpdateIpamResponse) SetHeaders(v map[string]*string) *UpdateIpamResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamResponse) SetStatusCode(v int32) *UpdateIpamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamResponse) SetBody(v *UpdateIpamResponseBody) *UpdateIpamResponse {
	s.Body = v
	return s
}

func (s *UpdateIpamResponse) Validate() error {
	return dara.Validate(s)
}
