// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreSummaryResponseBody) *DescribeRestoreSummaryResponse
	GetBody() *DescribeRestoreSummaryResponseBody
}

type DescribeRestoreSummaryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreSummaryResponse) GetBody() *DescribeRestoreSummaryResponseBody {
	return s.Body
}

func (s *DescribeRestoreSummaryResponse) SetHeaders(v map[string]*string) *DescribeRestoreSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreSummaryResponse) SetStatusCode(v int32) *DescribeRestoreSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreSummaryResponse) SetBody(v *DescribeRestoreSummaryResponseBody) *DescribeRestoreSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
