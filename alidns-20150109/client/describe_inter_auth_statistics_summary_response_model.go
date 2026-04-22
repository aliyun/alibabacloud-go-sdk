// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInterAuthStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInterAuthStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInterAuthStatisticsSummaryResponseBody) *DescribeInterAuthStatisticsSummaryResponse
	GetBody() *DescribeInterAuthStatisticsSummaryResponseBody
}

type DescribeInterAuthStatisticsSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterAuthStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInterAuthStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInterAuthStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterAuthStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInterAuthStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInterAuthStatisticsSummaryResponse) GetBody() *DescribeInterAuthStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeInterAuthStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeInterAuthStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeInterAuthStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponse) SetBody(v *DescribeInterAuthStatisticsSummaryResponseBody) *DescribeInterAuthStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeInterAuthStatisticsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
