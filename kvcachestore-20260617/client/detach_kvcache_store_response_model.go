// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *DetachKVCacheStoreResponseBody) *DetachKVCacheStoreResponse
	GetBody() *DetachKVCacheStoreResponseBody
}

type DetachKVCacheStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *DetachKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachKVCacheStoreResponse) GetBody() *DetachKVCacheStoreResponseBody {
	return s.Body
}

func (s *DetachKVCacheStoreResponse) SetHeaders(v map[string]*string) *DetachKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *DetachKVCacheStoreResponse) SetStatusCode(v int32) *DetachKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachKVCacheStoreResponse) SetBody(v *DetachKVCacheStoreResponseBody) *DetachKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *DetachKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
