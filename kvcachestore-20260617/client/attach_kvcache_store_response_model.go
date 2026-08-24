// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *AttachKVCacheStoreResponseBody) *AttachKVCacheStoreResponse
	GetBody() *AttachKVCacheStoreResponseBody
}

type AttachKVCacheStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *AttachKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachKVCacheStoreResponse) GetBody() *AttachKVCacheStoreResponseBody {
	return s.Body
}

func (s *AttachKVCacheStoreResponse) SetHeaders(v map[string]*string) *AttachKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *AttachKVCacheStoreResponse) SetStatusCode(v int32) *AttachKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachKVCacheStoreResponse) SetBody(v *AttachKVCacheStoreResponseBody) *AttachKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *AttachKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
