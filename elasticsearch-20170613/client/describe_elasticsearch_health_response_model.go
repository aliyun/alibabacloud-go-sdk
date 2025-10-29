// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticsearchHealthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticsearchHealthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticsearchHealthResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticsearchHealthResponseBody) *DescribeElasticsearchHealthResponse
	GetBody() *DescribeElasticsearchHealthResponseBody
}

type DescribeElasticsearchHealthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticsearchHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticsearchHealthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticsearchHealthResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticsearchHealthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticsearchHealthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticsearchHealthResponse) GetBody() *DescribeElasticsearchHealthResponseBody {
	return s.Body
}

func (s *DescribeElasticsearchHealthResponse) SetHeaders(v map[string]*string) *DescribeElasticsearchHealthResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticsearchHealthResponse) SetStatusCode(v int32) *DescribeElasticsearchHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticsearchHealthResponse) SetBody(v *DescribeElasticsearchHealthResponseBody) *DescribeElasticsearchHealthResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticsearchHealthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
