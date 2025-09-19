// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMaliciousFileWhitelistConfigResponseBody) *UpdateMaliciousFileWhitelistConfigResponse
	GetBody() *UpdateMaliciousFileWhitelistConfigResponseBody
}

type UpdateMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) GetBody() *UpdateMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *UpdateMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *UpdateMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) SetBody(v *UpdateMaliciousFileWhitelistConfigResponseBody) *UpdateMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateMaliciousFileWhitelistConfigResponse) Validate() error {
	return dara.Validate(s)
}
