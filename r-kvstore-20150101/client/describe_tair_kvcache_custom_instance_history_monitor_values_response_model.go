// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse
	GetBody() *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody
}

type DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse struct {
	Headers    map[string]*string                                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) GetBody() *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody {
	return s.Body
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) SetStatusCode(v int32) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) SetBody(v *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponseBody) *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse {
	s.Body = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceHistoryMonitorValuesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
