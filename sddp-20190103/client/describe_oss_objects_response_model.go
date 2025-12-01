// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOssObjectsResponseBody) *DescribeOssObjectsResponse
	GetBody() *DescribeOssObjectsResponseBody
}

type DescribeOssObjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssObjectsResponse) GetBody() *DescribeOssObjectsResponseBody {
	return s.Body
}

func (s *DescribeOssObjectsResponse) SetHeaders(v map[string]*string) *DescribeOssObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectsResponse) SetStatusCode(v int32) *DescribeOssObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectsResponse) SetBody(v *DescribeOssObjectsResponseBody) *DescribeOssObjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeOssObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
