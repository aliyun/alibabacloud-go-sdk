// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpamPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpamPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpamPoolResponseBody) *UpdateIpamPoolResponse
	GetBody() *UpdateIpamPoolResponseBody
}

type UpdateIpamPoolResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpamPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpamPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpamPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpamPoolResponse) GetBody() *UpdateIpamPoolResponseBody {
	return s.Body
}

func (s *UpdateIpamPoolResponse) SetHeaders(v map[string]*string) *UpdateIpamPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpamPoolResponse) SetStatusCode(v int32) *UpdateIpamPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpamPoolResponse) SetBody(v *UpdateIpamPoolResponseBody) *UpdateIpamPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateIpamPoolResponse) Validate() error {
	return dara.Validate(s)
}
