// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMaliciousFileWhitelistConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteMaliciousFileWhitelistConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteMaliciousFileWhitelistConfigResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteMaliciousFileWhitelistConfigResponseBody) *BatchDeleteMaliciousFileWhitelistConfigResponse
	GetBody() *BatchDeleteMaliciousFileWhitelistConfigResponseBody
}

type BatchDeleteMaliciousFileWhitelistConfigResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteMaliciousFileWhitelistConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteMaliciousFileWhitelistConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMaliciousFileWhitelistConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) GetBody() *BatchDeleteMaliciousFileWhitelistConfigResponseBody {
	return s.Body
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) SetHeaders(v map[string]*string) *BatchDeleteMaliciousFileWhitelistConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) SetStatusCode(v int32) *BatchDeleteMaliciousFileWhitelistConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) SetBody(v *BatchDeleteMaliciousFileWhitelistConfigResponseBody) *BatchDeleteMaliciousFileWhitelistConfigResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteMaliciousFileWhitelistConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
