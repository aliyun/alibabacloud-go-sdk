// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateKVCacheStoreResponseBody) *CreateKVCacheStoreResponse
	GetBody() *CreateKVCacheStoreResponseBody
}

type CreateKVCacheStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKVCacheStoreResponse) GetBody() *CreateKVCacheStoreResponseBody {
	return s.Body
}

func (s *CreateKVCacheStoreResponse) SetHeaders(v map[string]*string) *CreateKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateKVCacheStoreResponse) SetStatusCode(v int32) *CreateKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKVCacheStoreResponse) SetBody(v *CreateKVCacheStoreResponseBody) *CreateKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *CreateKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
