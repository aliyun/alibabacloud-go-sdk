// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCheckReportUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataCheckReportUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataCheckReportUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataCheckReportUrlResponseBody) *DescribeDataCheckReportUrlResponse
	GetBody() *DescribeDataCheckReportUrlResponseBody
}

type DescribeDataCheckReportUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCheckReportUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataCheckReportUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCheckReportUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCheckReportUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataCheckReportUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataCheckReportUrlResponse) GetBody() *DescribeDataCheckReportUrlResponseBody {
	return s.Body
}

func (s *DescribeDataCheckReportUrlResponse) SetHeaders(v map[string]*string) *DescribeDataCheckReportUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCheckReportUrlResponse) SetStatusCode(v int32) *DescribeDataCheckReportUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCheckReportUrlResponse) SetBody(v *DescribeDataCheckReportUrlResponseBody) *DescribeDataCheckReportUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeDataCheckReportUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
