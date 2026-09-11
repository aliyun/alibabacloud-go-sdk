// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableVscsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKVCacheStoreAvailableVscsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKVCacheStoreAvailableVscsResponse
	GetStatusCode() *int32
	SetBody(v *ListKVCacheStoreAvailableVscsResponseBody) *ListKVCacheStoreAvailableVscsResponse
	GetBody() *ListKVCacheStoreAvailableVscsResponseBody
}

type ListKVCacheStoreAvailableVscsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKVCacheStoreAvailableVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKVCacheStoreAvailableVscsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableVscsResponse) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableVscsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKVCacheStoreAvailableVscsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKVCacheStoreAvailableVscsResponse) GetBody() *ListKVCacheStoreAvailableVscsResponseBody {
	return s.Body
}

func (s *ListKVCacheStoreAvailableVscsResponse) SetHeaders(v map[string]*string) *ListKVCacheStoreAvailableVscsResponse {
	s.Headers = v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponse) SetStatusCode(v int32) *ListKVCacheStoreAvailableVscsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponse) SetBody(v *ListKVCacheStoreAvailableVscsResponseBody) *ListKVCacheStoreAvailableVscsResponse {
	s.Body = v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
