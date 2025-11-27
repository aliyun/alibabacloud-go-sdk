// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableCrossRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableCrossRegionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableCrossRegionResponseBody) *DescribeAvailableCrossRegionResponse
	GetBody() *DescribeAvailableCrossRegionResponseBody
}

type DescribeAvailableCrossRegionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableCrossRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableCrossRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableCrossRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableCrossRegionResponse) GetBody() *DescribeAvailableCrossRegionResponseBody {
	return s.Body
}

func (s *DescribeAvailableCrossRegionResponse) SetHeaders(v map[string]*string) *DescribeAvailableCrossRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableCrossRegionResponse) SetStatusCode(v int32) *DescribeAvailableCrossRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableCrossRegionResponse) SetBody(v *DescribeAvailableCrossRegionResponseBody) *DescribeAvailableCrossRegionResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableCrossRegionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
