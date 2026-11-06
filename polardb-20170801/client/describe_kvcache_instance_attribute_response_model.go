// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKVCacheInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKVCacheInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKVCacheInstanceAttributeResponseBody) *DescribeKVCacheInstanceAttributeResponse
	GetBody() *DescribeKVCacheInstanceAttributeResponseBody
}

type DescribeKVCacheInstanceAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKVCacheInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKVCacheInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKVCacheInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKVCacheInstanceAttributeResponse) GetBody() *DescribeKVCacheInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeKVCacheInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeKVCacheInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponse) SetStatusCode(v int32) *DescribeKVCacheInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponse) SetBody(v *DescribeKVCacheInstanceAttributeResponseBody) *DescribeKVCacheInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeKVCacheInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
