// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetMaliciousFileWhitelistConfigResponseBody) *GetMaliciousFileWhitelistConfigResponse
	GetBody() *GetMaliciousFileWhitelistConfigResponseBody
}

type GetMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *GetMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMaliciousFileWhitelistConfigResponse) GetBody() *GetMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *GetMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *GetMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *GetMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponse) SetBody(v *GetMaliciousFileWhitelistConfigResponseBody) *GetMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *GetMaliciousFileWhitelistConfigResponse) Validate() error {
	return dara.Validate(s)
}
