// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlushExpireKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FlushExpireKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FlushExpireKeysResponse
	GetStatusCode() *int32
	SetBody(v *FlushExpireKeysResponseBody) *FlushExpireKeysResponse
	GetBody() *FlushExpireKeysResponseBody
}

type FlushExpireKeysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FlushExpireKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FlushExpireKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s FlushExpireKeysResponse) GoString() string {
	return s.String()
}

func (s *FlushExpireKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FlushExpireKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FlushExpireKeysResponse) GetBody() *FlushExpireKeysResponseBody {
	return s.Body
}

func (s *FlushExpireKeysResponse) SetHeaders(v map[string]*string) *FlushExpireKeysResponse {
	s.Headers = v
	return s
}

func (s *FlushExpireKeysResponse) SetStatusCode(v int32) *FlushExpireKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *FlushExpireKeysResponse) SetBody(v *FlushExpireKeysResponseBody) *FlushExpireKeysResponse {
	s.Body = v
	return s
}

func (s *FlushExpireKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
