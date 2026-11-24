// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceStatisticByRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceStatisticByRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceStatisticByRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceStatisticByRegionResponseBody) *DescribeResourceStatisticByRegionResponse
	GetBody() *DescribeResourceStatisticByRegionResponseBody
}

type DescribeResourceStatisticByRegionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceStatisticByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceStatisticByRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceStatisticByRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticByRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceStatisticByRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceStatisticByRegionResponse) GetBody() *DescribeResourceStatisticByRegionResponseBody {
	return s.Body
}

func (s *DescribeResourceStatisticByRegionResponse) SetHeaders(v map[string]*string) *DescribeResourceStatisticByRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceStatisticByRegionResponse) SetStatusCode(v int32) *DescribeResourceStatisticByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceStatisticByRegionResponse) SetBody(v *DescribeResourceStatisticByRegionResponseBody) *DescribeResourceStatisticByRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceStatisticByRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
