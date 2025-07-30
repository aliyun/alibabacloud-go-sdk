// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairKVCacheInferInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairKVCacheInferInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairKVCacheInferInstancesResponseBody) *DescribeTairKVCacheInferInstancesResponse
	GetBody() *DescribeTairKVCacheInferInstancesResponseBody
}

type DescribeTairKVCacheInferInstancesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairKVCacheInferInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairKVCacheInferInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairKVCacheInferInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairKVCacheInferInstancesResponse) GetBody() *DescribeTairKVCacheInferInstancesResponseBody {
	return s.Body
}

func (s *DescribeTairKVCacheInferInstancesResponse) SetHeaders(v map[string]*string) *DescribeTairKVCacheInferInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponse) SetStatusCode(v int32) *DescribeTairKVCacheInferInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponse) SetBody(v *DescribeTairKVCacheInferInstancesResponseBody) *DescribeTairKVCacheInferInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeTairKVCacheInferInstancesResponse) Validate() error {
	return dara.Validate(s)
}
