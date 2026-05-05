// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductionStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeductionStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeductionStatisticResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeductionStatisticResponseBody) *DescribeDeductionStatisticResponse
	GetBody() *DescribeDeductionStatisticResponseBody
}

type DescribeDeductionStatisticResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeductionStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeductionStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductionStatisticResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeductionStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeductionStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeductionStatisticResponse) GetBody() *DescribeDeductionStatisticResponseBody {
	return s.Body
}

func (s *DescribeDeductionStatisticResponse) SetHeaders(v map[string]*string) *DescribeDeductionStatisticResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeductionStatisticResponse) SetStatusCode(v int32) *DescribeDeductionStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeductionStatisticResponse) SetBody(v *DescribeDeductionStatisticResponseBody) *DescribeDeductionStatisticResponse {
	s.Body = v
	return s
}

func (s *DescribeDeductionStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
