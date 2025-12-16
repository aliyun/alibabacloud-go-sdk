// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppStatisticsResponseBody) *DescribeAppStatisticsResponse
	GetBody() *DescribeAppStatisticsResponseBody
}

type DescribeAppStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppStatisticsResponse) GetBody() *DescribeAppStatisticsResponseBody {
	return s.Body
}

func (s *DescribeAppStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAppStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppStatisticsResponse) SetStatusCode(v int32) *DescribeAppStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppStatisticsResponse) SetBody(v *DescribeAppStatisticsResponseBody) *DescribeAppStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeAppStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
