// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticQpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticQpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticQpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticQpsResponseBody) *DescribeElasticQpsResponse
	GetBody() *DescribeElasticQpsResponseBody
}

type DescribeElasticQpsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticQpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticQpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticQpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticQpsResponse) GetBody() *DescribeElasticQpsResponseBody {
	return s.Body
}

func (s *DescribeElasticQpsResponse) SetHeaders(v map[string]*string) *DescribeElasticQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticQpsResponse) SetStatusCode(v int32) *DescribeElasticQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticQpsResponse) SetBody(v *DescribeElasticQpsResponseBody) *DescribeElasticQpsResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticQpsResponse) Validate() error {
	return dara.Validate(s)
}
