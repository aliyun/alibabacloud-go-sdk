// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKVCacheStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKVCacheStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKVCacheStoreResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKVCacheStoreResponseBody) *UpdateKVCacheStoreResponse
	GetBody() *UpdateKVCacheStoreResponseBody
}

type UpdateKVCacheStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKVCacheStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKVCacheStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKVCacheStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateKVCacheStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKVCacheStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKVCacheStoreResponse) GetBody() *UpdateKVCacheStoreResponseBody {
	return s.Body
}

func (s *UpdateKVCacheStoreResponse) SetHeaders(v map[string]*string) *UpdateKVCacheStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateKVCacheStoreResponse) SetStatusCode(v int32) *UpdateKVCacheStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKVCacheStoreResponse) SetBody(v *UpdateKVCacheStoreResponseBody) *UpdateKVCacheStoreResponse {
	s.Body = v
	return s
}

func (s *UpdateKVCacheStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
