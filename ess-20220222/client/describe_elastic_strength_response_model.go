// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticStrengthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticStrengthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticStrengthResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticStrengthResponseBody) *DescribeElasticStrengthResponse
	GetBody() *DescribeElasticStrengthResponseBody
}

type DescribeElasticStrengthResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticStrengthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticStrengthResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticStrengthResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticStrengthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticStrengthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticStrengthResponse) GetBody() *DescribeElasticStrengthResponseBody {
	return s.Body
}

func (s *DescribeElasticStrengthResponse) SetHeaders(v map[string]*string) *DescribeElasticStrengthResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticStrengthResponse) SetStatusCode(v int32) *DescribeElasticStrengthResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticStrengthResponse) SetBody(v *DescribeElasticStrengthResponseBody) *DescribeElasticStrengthResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticStrengthResponse) Validate() error {
	return dara.Validate(s)
}
