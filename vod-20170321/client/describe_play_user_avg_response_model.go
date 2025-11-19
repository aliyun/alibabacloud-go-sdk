// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserAvgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayUserAvgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayUserAvgResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayUserAvgResponseBody) *DescribePlayUserAvgResponse
	GetBody() *DescribePlayUserAvgResponseBody
}

type DescribePlayUserAvgResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayUserAvgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayUserAvgResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserAvgResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayUserAvgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayUserAvgResponse) GetBody() *DescribePlayUserAvgResponseBody {
	return s.Body
}

func (s *DescribePlayUserAvgResponse) SetHeaders(v map[string]*string) *DescribePlayUserAvgResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayUserAvgResponse) SetStatusCode(v int32) *DescribePlayUserAvgResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayUserAvgResponse) SetBody(v *DescribePlayUserAvgResponseBody) *DescribePlayUserAvgResponse {
	s.Body = v
	return s
}

func (s *DescribePlayUserAvgResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
