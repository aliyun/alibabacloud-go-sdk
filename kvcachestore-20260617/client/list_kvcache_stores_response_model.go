// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKVCacheStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKVCacheStoresResponse
	GetStatusCode() *int32
	SetBody(v *ListKVCacheStoresResponseBody) *ListKVCacheStoresResponse
	GetBody() *ListKVCacheStoresResponseBody
}

type ListKVCacheStoresResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKVCacheStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKVCacheStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoresResponse) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKVCacheStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKVCacheStoresResponse) GetBody() *ListKVCacheStoresResponseBody {
	return s.Body
}

func (s *ListKVCacheStoresResponse) SetHeaders(v map[string]*string) *ListKVCacheStoresResponse {
	s.Headers = v
	return s
}

func (s *ListKVCacheStoresResponse) SetStatusCode(v int32) *ListKVCacheStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKVCacheStoresResponse) SetBody(v *ListKVCacheStoresResponseBody) *ListKVCacheStoresResponse {
	s.Body = v
	return s
}

func (s *ListKVCacheStoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
