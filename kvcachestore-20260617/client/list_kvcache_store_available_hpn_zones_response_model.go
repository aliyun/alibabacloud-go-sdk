// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableHpnZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKVCacheStoreAvailableHpnZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKVCacheStoreAvailableHpnZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListKVCacheStoreAvailableHpnZonesResponseBody) *ListKVCacheStoreAvailableHpnZonesResponse
	GetBody() *ListKVCacheStoreAvailableHpnZonesResponseBody
}

type ListKVCacheStoreAvailableHpnZonesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKVCacheStoreAvailableHpnZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKVCacheStoreAvailableHpnZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableHpnZonesResponse) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) GetBody() *ListKVCacheStoreAvailableHpnZonesResponseBody {
	return s.Body
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) SetHeaders(v map[string]*string) *ListKVCacheStoreAvailableHpnZonesResponse {
	s.Headers = v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) SetStatusCode(v int32) *ListKVCacheStoreAvailableHpnZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) SetBody(v *ListKVCacheStoreAvailableHpnZonesResponseBody) *ListKVCacheStoreAvailableHpnZonesResponse {
	s.Body = v
	return s
}

func (s *ListKVCacheStoreAvailableHpnZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
