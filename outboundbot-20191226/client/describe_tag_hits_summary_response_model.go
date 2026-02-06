// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagHitsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagHitsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagHitsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagHitsSummaryResponseBody) *DescribeTagHitsSummaryResponse
	GetBody() *DescribeTagHitsSummaryResponseBody
}

type DescribeTagHitsSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagHitsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagHitsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagHitsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagHitsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagHitsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagHitsSummaryResponse) GetBody() *DescribeTagHitsSummaryResponseBody {
	return s.Body
}

func (s *DescribeTagHitsSummaryResponse) SetHeaders(v map[string]*string) *DescribeTagHitsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagHitsSummaryResponse) SetStatusCode(v int32) *DescribeTagHitsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagHitsSummaryResponse) SetBody(v *DescribeTagHitsSummaryResponseBody) *DescribeTagHitsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeTagHitsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
