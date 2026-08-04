// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCacheOperateSyncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCacheOperateSyncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCacheOperateSyncResponse
	GetStatusCode() *int32
	SetBody(v *SetCacheOperateSyncResponseBody) *SetCacheOperateSyncResponse
	GetBody() *SetCacheOperateSyncResponseBody
}

type SetCacheOperateSyncResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCacheOperateSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCacheOperateSyncResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCacheOperateSyncResponse) GoString() string {
	return s.String()
}

func (s *SetCacheOperateSyncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCacheOperateSyncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCacheOperateSyncResponse) GetBody() *SetCacheOperateSyncResponseBody {
	return s.Body
}

func (s *SetCacheOperateSyncResponse) SetHeaders(v map[string]*string) *SetCacheOperateSyncResponse {
	s.Headers = v
	return s
}

func (s *SetCacheOperateSyncResponse) SetStatusCode(v int32) *SetCacheOperateSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCacheOperateSyncResponse) SetBody(v *SetCacheOperateSyncResponseBody) *SetCacheOperateSyncResponse {
	s.Body = v
	return s
}

func (s *SetCacheOperateSyncResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
