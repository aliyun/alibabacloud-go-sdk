// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticalDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStatisticalDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStatisticalDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStatisticalDataResponseBody) *DescribeStatisticalDataResponse
	GetBody() *DescribeStatisticalDataResponseBody
}

type DescribeStatisticalDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStatisticalDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStatisticalDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticalDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStatisticalDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStatisticalDataResponse) GetBody() *DescribeStatisticalDataResponseBody {
	return s.Body
}

func (s *DescribeStatisticalDataResponse) SetHeaders(v map[string]*string) *DescribeStatisticalDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeStatisticalDataResponse) SetStatusCode(v int32) *DescribeStatisticalDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStatisticalDataResponse) SetBody(v *DescribeStatisticalDataResponseBody) *DescribeStatisticalDataResponse {
	s.Body = v
	return s
}

func (s *DescribeStatisticalDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
