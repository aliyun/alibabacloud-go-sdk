// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairKVCacheInferInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairKVCacheInferInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairKVCacheInferInstanceAttributeResponseBody) *DescribeTairKVCacheInferInstanceAttributeResponse
	GetBody() *DescribeTairKVCacheInferInstanceAttributeResponseBody
}

type DescribeTairKVCacheInferInstanceAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairKVCacheInferInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) GetBody() *DescribeTairKVCacheInferInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeTairKVCacheInferInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) SetStatusCode(v int32) *DescribeTairKVCacheInferInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) SetBody(v *DescribeTairKVCacheInferInstanceAttributeResponseBody) *DescribeTairKVCacheInferInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
