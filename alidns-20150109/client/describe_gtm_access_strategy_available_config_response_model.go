// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyAvailableConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyAvailableConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmAccessStrategyAvailableConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmAccessStrategyAvailableConfigResponseBody) *DescribeGtmAccessStrategyAvailableConfigResponse
	GetBody() *DescribeGtmAccessStrategyAvailableConfigResponseBody
}

type DescribeGtmAccessStrategyAvailableConfigResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmAccessStrategyAvailableConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmAccessStrategyAvailableConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyAvailableConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) GetBody() *DescribeGtmAccessStrategyAvailableConfigResponseBody {
	return s.Body
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyAvailableConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) SetStatusCode(v int32) *DescribeGtmAccessStrategyAvailableConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) SetBody(v *DescribeGtmAccessStrategyAvailableConfigResponseBody) *DescribeGtmAccessStrategyAvailableConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmAccessStrategyAvailableConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
