// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckWarningSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckWarningSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckWarningSummaryResponseBody) *DescribeCheckWarningSummaryResponse
	GetBody() *DescribeCheckWarningSummaryResponseBody
}

type DescribeCheckWarningSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckWarningSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckWarningSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckWarningSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckWarningSummaryResponse) GetBody() *DescribeCheckWarningSummaryResponseBody {
	return s.Body
}

func (s *DescribeCheckWarningSummaryResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningSummaryResponse) SetStatusCode(v int32) *DescribeCheckWarningSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponse) SetBody(v *DescribeCheckWarningSummaryResponseBody) *DescribeCheckWarningSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckWarningSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
