// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateCacheResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateCacheResponseBody) *DescribeTemplateCacheResponse
	GetBody() *DescribeTemplateCacheResponseBody
}

type DescribeTemplateCacheResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateCacheResponse) GetBody() *DescribeTemplateCacheResponseBody {
	return s.Body
}

func (s *DescribeTemplateCacheResponse) SetHeaders(v map[string]*string) *DescribeTemplateCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateCacheResponse) SetStatusCode(v int32) *DescribeTemplateCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateCacheResponse) SetBody(v *DescribeTemplateCacheResponseBody) *DescribeTemplateCacheResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
