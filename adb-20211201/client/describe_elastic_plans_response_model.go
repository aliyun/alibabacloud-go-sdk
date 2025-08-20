// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlansResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticPlansResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticPlansResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticPlansResponseBody) *DescribeElasticPlansResponse
	GetBody() *DescribeElasticPlansResponseBody
}

type DescribeElasticPlansResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticPlansResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticPlansResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticPlansResponse) GetBody() *DescribeElasticPlansResponseBody {
	return s.Body
}

func (s *DescribeElasticPlansResponse) SetHeaders(v map[string]*string) *DescribeElasticPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticPlansResponse) SetStatusCode(v int32) *DescribeElasticPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticPlansResponse) SetBody(v *DescribeElasticPlansResponseBody) *DescribeElasticPlansResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticPlansResponse) Validate() error {
	return dara.Validate(s)
}
