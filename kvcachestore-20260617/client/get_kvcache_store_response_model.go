// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *GetKVCacheStoreResponseBody) *GetKVCacheStoreResponse
	GetBody() *GetKVCacheStoreResponseBody
}

type GetKVCacheStoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *GetKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKVCacheStoreResponse) GetBody() *GetKVCacheStoreResponseBody {
	return s.Body
}

func (s *GetKVCacheStoreResponse) SetHeaders(v map[string]*string) *GetKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *GetKVCacheStoreResponse) SetStatusCode(v int32) *GetKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKVCacheStoreResponse) SetBody(v *GetKVCacheStoreResponseBody) *GetKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *GetKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
