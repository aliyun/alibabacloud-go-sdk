// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHostGroupElasticStrategyParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHostGroupElasticStrategyParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHostGroupElasticStrategyParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHostGroupElasticStrategyParametersResponseBody) *DescribeHostGroupElasticStrategyParametersResponse
	GetBody() *DescribeHostGroupElasticStrategyParametersResponseBody
}

type DescribeHostGroupElasticStrategyParametersResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHostGroupElasticStrategyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHostGroupElasticStrategyParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHostGroupElasticStrategyParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) GetBody() *DescribeHostGroupElasticStrategyParametersResponseBody {
	return s.Body
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) SetHeaders(v map[string]*string) *DescribeHostGroupElasticStrategyParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) SetStatusCode(v int32) *DescribeHostGroupElasticStrategyParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) SetBody(v *DescribeHostGroupElasticStrategyParametersResponseBody) *DescribeHostGroupElasticStrategyParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeHostGroupElasticStrategyParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
