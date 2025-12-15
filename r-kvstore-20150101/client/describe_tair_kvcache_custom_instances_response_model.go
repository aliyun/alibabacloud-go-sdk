// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairKVCacheCustomInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairKVCacheCustomInstancesResponseBody) *DescribeTairKVCacheCustomInstancesResponse
	GetBody() *DescribeTairKVCacheCustomInstancesResponseBody
}

type DescribeTairKVCacheCustomInstancesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairKVCacheCustomInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairKVCacheCustomInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairKVCacheCustomInstancesResponse) GetBody() *DescribeTairKVCacheCustomInstancesResponseBody {
	return s.Body
}

func (s *DescribeTairKVCacheCustomInstancesResponse) SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponse) SetStatusCode(v int32) *DescribeTairKVCacheCustomInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponse) SetBody(v *DescribeTairKVCacheCustomInstancesResponseBody) *DescribeTairKVCacheCustomInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
