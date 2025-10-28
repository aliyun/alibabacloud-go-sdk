// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowLogStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowLogStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowLogStatisticResponseBody) *DescribeSlowLogStatisticResponse
	GetBody() *DescribeSlowLogStatisticResponseBody
}

type DescribeSlowLogStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowLogStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowLogStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowLogStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowLogStatisticResponse) GetBody() *DescribeSlowLogStatisticResponseBody {
	return s.Body
}

func (s *DescribeSlowLogStatisticResponse) SetHeaders(v map[string]*string) *DescribeSlowLogStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogStatisticResponse) SetStatusCode(v int32) *DescribeSlowLogStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowLogStatisticResponse) SetBody(v *DescribeSlowLogStatisticResponseBody) *DescribeSlowLogStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowLogStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
