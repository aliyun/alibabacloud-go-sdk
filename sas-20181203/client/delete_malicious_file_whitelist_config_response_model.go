// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMaliciousFileWhitelistConfigResponseBody) *DeleteMaliciousFileWhitelistConfigResponse
	GetBody() *DeleteMaliciousFileWhitelistConfigResponseBody
}

type DeleteMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) GetBody() *DeleteMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *DeleteMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *DeleteMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) SetBody(v *DeleteMaliciousFileWhitelistConfigResponseBody) *DeleteMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteMaliciousFileWhitelistConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
