// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReportRecipientStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeReportRecipientStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeReportRecipientStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeReportRecipientStatusResponseBody) *DescribeReportRecipientStatusResponse
	GetBody() *DescribeReportRecipientStatusResponseBody
}

type DescribeReportRecipientStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeReportRecipientStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeReportRecipientStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeReportRecipientStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeReportRecipientStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeReportRecipientStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeReportRecipientStatusResponse) GetBody() *DescribeReportRecipientStatusResponseBody {
	return s.Body
}

func (s *DescribeReportRecipientStatusResponse) SetHeaders(v map[string]*string) *DescribeReportRecipientStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeReportRecipientStatusResponse) SetStatusCode(v int32) *DescribeReportRecipientStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeReportRecipientStatusResponse) SetBody(v *DescribeReportRecipientStatusResponseBody) *DescribeReportRecipientStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeReportRecipientStatusResponse) Validate() error {
	return dara.Validate(s)
}
