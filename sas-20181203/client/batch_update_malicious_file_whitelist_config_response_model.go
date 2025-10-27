// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateMaliciousFileWhitelistConfigResponseBody) *BatchUpdateMaliciousFileWhitelistConfigResponse
	GetBody() *BatchUpdateMaliciousFileWhitelistConfigResponseBody
}

type BatchUpdateMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) GetBody() *BatchUpdateMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *BatchUpdateMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *BatchUpdateMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) SetBody(v *BatchUpdateMaliciousFileWhitelistConfigResponseBody) *BatchUpdateMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateMaliciousFileWhitelistConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
