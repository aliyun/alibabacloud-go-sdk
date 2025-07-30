// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTairKVCacheCustomInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTairKVCacheCustomInstanceAttributeResponseBody) *DescribeTairKVCacheCustomInstanceAttributeResponse
	GetBody() *DescribeTairKVCacheCustomInstanceAttributeResponseBody
}

type DescribeTairKVCacheCustomInstanceAttributeResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTairKVCacheCustomInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) GetBody() *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeTairKVCacheCustomInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) SetStatusCode(v int32) *DescribeTairKVCacheCustomInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) SetBody(v *DescribeTairKVCacheCustomInstanceAttributeResponseBody) *DescribeTairKVCacheCustomInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
