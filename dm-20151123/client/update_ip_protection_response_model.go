// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIpProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIpProtectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIpProtectionResponseBody) *UpdateIpProtectionResponse
	GetBody() *UpdateIpProtectionResponseBody
}

type UpdateIpProtectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIpProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIpProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpProtectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIpProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIpProtectionResponse) GetBody() *UpdateIpProtectionResponseBody {
	return s.Body
}

func (s *UpdateIpProtectionResponse) SetHeaders(v map[string]*string) *UpdateIpProtectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateIpProtectionResponse) SetStatusCode(v int32) *UpdateIpProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIpProtectionResponse) SetBody(v *UpdateIpProtectionResponseBody) *UpdateIpProtectionResponse {
	s.Body = v
	return s
}

func (s *UpdateIpProtectionResponse) Validate() error {
	return dara.Validate(s)
}
