// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticBandwidthSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeElasticBandwidthSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeElasticBandwidthSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeElasticBandwidthSpecResponseBody) *DescribeElasticBandwidthSpecResponse
	GetBody() *DescribeElasticBandwidthSpecResponseBody
}

type DescribeElasticBandwidthSpecResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeElasticBandwidthSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeElasticBandwidthSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticBandwidthSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeElasticBandwidthSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeElasticBandwidthSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeElasticBandwidthSpecResponse) GetBody() *DescribeElasticBandwidthSpecResponseBody {
	return s.Body
}

func (s *DescribeElasticBandwidthSpecResponse) SetHeaders(v map[string]*string) *DescribeElasticBandwidthSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetStatusCode(v int32) *DescribeElasticBandwidthSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) SetBody(v *DescribeElasticBandwidthSpecResponseBody) *DescribeElasticBandwidthSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeElasticBandwidthSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
