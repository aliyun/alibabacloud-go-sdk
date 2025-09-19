// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateMaliciousFileWhitelistConfigResponseBody) *CreateMaliciousFileWhitelistConfigResponse
	GetBody() *CreateMaliciousFileWhitelistConfigResponseBody
}

type CreateMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMaliciousFileWhitelistConfigResponse) GetBody() *CreateMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *CreateMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *CreateMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *CreateMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponse) SetBody(v *CreateMaliciousFileWhitelistConfigResponseBody) *CreateMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *CreateMaliciousFileWhitelistConfigResponse) Validate() error {
	return dara.Validate(s)
}
