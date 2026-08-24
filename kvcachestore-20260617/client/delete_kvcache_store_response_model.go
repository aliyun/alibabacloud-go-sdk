// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKVCacheStoreResponseBody) *DeleteKVCacheStoreResponse
	GetBody() *DeleteKVCacheStoreResponseBody
}

type DeleteKVCacheStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *DeleteKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKVCacheStoreResponse) GetBody() *DeleteKVCacheStoreResponseBody {
	return s.Body
}

func (s *DeleteKVCacheStoreResponse) SetHeaders(v map[string]*string) *DeleteKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *DeleteKVCacheStoreResponse) SetStatusCode(v int32) *DeleteKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKVCacheStoreResponse) SetBody(v *DeleteKVCacheStoreResponseBody) *DeleteKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *DeleteKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
