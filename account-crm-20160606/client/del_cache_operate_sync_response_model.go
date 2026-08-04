// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelCacheOperateSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DelCacheOperateSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DelCacheOperateSyncResponse
	GetStatusCode() *int32
	SetBody(v *DelCacheOperateSyncResponseBody) *DelCacheOperateSyncResponse
	GetBody() *DelCacheOperateSyncResponseBody
}

type DelCacheOperateSyncResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DelCacheOperateSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DelCacheOperateSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s DelCacheOperateSyncResponse) GoString() string {
	return s.String()
}

func (s *DelCacheOperateSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DelCacheOperateSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DelCacheOperateSyncResponse) GetBody() *DelCacheOperateSyncResponseBody {
	return s.Body
}

func (s *DelCacheOperateSyncResponse) SetHeaders(v map[string]*string) *DelCacheOperateSyncResponse {
	s.Headers = v
	return s
}

func (s *DelCacheOperateSyncResponse) SetStatusCode(v int32) *DelCacheOperateSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *DelCacheOperateSyncResponse) SetBody(v *DelCacheOperateSyncResponseBody) *DelCacheOperateSyncResponse {
	s.Body = v
	return s
}

func (s *DelCacheOperateSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
