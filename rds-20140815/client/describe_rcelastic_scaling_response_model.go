// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCElasticScalingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCElasticScalingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCElasticScalingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCElasticScalingResponseBody) *DescribeRCElasticScalingResponse
	GetBody() *DescribeRCElasticScalingResponseBody
}

type DescribeRCElasticScalingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCElasticScalingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCElasticScalingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCElasticScalingResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCElasticScalingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCElasticScalingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCElasticScalingResponse) GetBody() *DescribeRCElasticScalingResponseBody {
	return s.Body
}

func (s *DescribeRCElasticScalingResponse) SetHeaders(v map[string]*string) *DescribeRCElasticScalingResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCElasticScalingResponse) SetStatusCode(v int32) *DescribeRCElasticScalingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCElasticScalingResponse) SetBody(v *DescribeRCElasticScalingResponseBody) *DescribeRCElasticScalingResponse {
	s.Body = v
	return s
}

func (s *DescribeRCElasticScalingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
