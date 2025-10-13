// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfigurationPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfigurationPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfigurationPriceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfigurationPriceResponseBody) *DescribeConfigurationPriceResponse
	GetBody() *DescribeConfigurationPriceResponseBody
}

type DescribeConfigurationPriceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfigurationPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfigurationPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfigurationPriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfigurationPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfigurationPriceResponse) GetBody() *DescribeConfigurationPriceResponseBody {
	return s.Body
}

func (s *DescribeConfigurationPriceResponse) SetHeaders(v map[string]*string) *DescribeConfigurationPriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigurationPriceResponse) SetStatusCode(v int32) *DescribeConfigurationPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigurationPriceResponse) SetBody(v *DescribeConfigurationPriceResponseBody) *DescribeConfigurationPriceResponse {
	s.Body = v
	return s
}

func (s *DescribeConfigurationPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
