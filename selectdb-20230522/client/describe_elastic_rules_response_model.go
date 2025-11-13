// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticRulesResponseBody) *DescribeElasticRulesResponse
	GetBody() *DescribeElasticRulesResponseBody
}

type DescribeElasticRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticRulesResponse) GetBody() *DescribeElasticRulesResponseBody {
	return s.Body
}

func (s *DescribeElasticRulesResponse) SetHeaders(v map[string]*string) *DescribeElasticRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticRulesResponse) SetStatusCode(v int32) *DescribeElasticRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticRulesResponse) SetBody(v *DescribeElasticRulesResponseBody) *DescribeElasticRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
