// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAttachInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKVCacheStoreAttachInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKVCacheStoreAttachInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListKVCacheStoreAttachInfoResponseBody) *ListKVCacheStoreAttachInfoResponse
	GetBody() *ListKVCacheStoreAttachInfoResponseBody
}

type ListKVCacheStoreAttachInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKVCacheStoreAttachInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKVCacheStoreAttachInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAttachInfoResponse) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAttachInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKVCacheStoreAttachInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKVCacheStoreAttachInfoResponse) GetBody() *ListKVCacheStoreAttachInfoResponseBody {
	return s.Body
}

func (s *ListKVCacheStoreAttachInfoResponse) SetHeaders(v map[string]*string) *ListKVCacheStoreAttachInfoResponse {
	s.Headers = v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponse) SetStatusCode(v int32) *ListKVCacheStoreAttachInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponse) SetBody(v *ListKVCacheStoreAttachInfoResponseBody) *ListKVCacheStoreAttachInfoResponse {
	s.Body = v
	return s
}

func (s *ListKVCacheStoreAttachInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
