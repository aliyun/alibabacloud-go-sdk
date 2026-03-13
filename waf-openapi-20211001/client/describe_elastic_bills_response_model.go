// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticBillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticBillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticBillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticBillsResponseBody) *DescribeElasticBillsResponse
	GetBody() *DescribeElasticBillsResponseBody
}

type DescribeElasticBillsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticBillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticBillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBillsResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticBillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticBillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticBillsResponse) GetBody() *DescribeElasticBillsResponseBody {
	return s.Body
}

func (s *DescribeElasticBillsResponse) SetHeaders(v map[string]*string) *DescribeElasticBillsResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticBillsResponse) SetStatusCode(v int32) *DescribeElasticBillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticBillsResponse) SetBody(v *DescribeElasticBillsResponseBody) *DescribeElasticBillsResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticBillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
