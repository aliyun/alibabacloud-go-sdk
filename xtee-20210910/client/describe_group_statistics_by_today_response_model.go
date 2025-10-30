// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStatisticsByTodayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupStatisticsByTodayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupStatisticsByTodayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupStatisticsByTodayResponseBody) *DescribeGroupStatisticsByTodayResponse
	GetBody() *DescribeGroupStatisticsByTodayResponseBody
}

type DescribeGroupStatisticsByTodayResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupStatisticsByTodayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupStatisticsByTodayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStatisticsByTodayResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupStatisticsByTodayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupStatisticsByTodayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupStatisticsByTodayResponse) GetBody() *DescribeGroupStatisticsByTodayResponseBody {
	return s.Body
}

func (s *DescribeGroupStatisticsByTodayResponse) SetHeaders(v map[string]*string) *DescribeGroupStatisticsByTodayResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupStatisticsByTodayResponse) SetStatusCode(v int32) *DescribeGroupStatisticsByTodayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupStatisticsByTodayResponse) SetBody(v *DescribeGroupStatisticsByTodayResponseBody) *DescribeGroupStatisticsByTodayResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupStatisticsByTodayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
