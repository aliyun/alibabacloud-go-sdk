// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileTypeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebLockFileTypeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebLockFileTypeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebLockFileTypeSummaryResponseBody) *DescribeWebLockFileTypeSummaryResponse
	GetBody() *DescribeWebLockFileTypeSummaryResponseBody
}

type DescribeWebLockFileTypeSummaryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebLockFileTypeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebLockFileTypeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileTypeSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileTypeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebLockFileTypeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebLockFileTypeSummaryResponse) GetBody() *DescribeWebLockFileTypeSummaryResponseBody {
	return s.Body
}

func (s *DescribeWebLockFileTypeSummaryResponse) SetHeaders(v map[string]*string) *DescribeWebLockFileTypeSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponse) SetStatusCode(v int32) *DescribeWebLockFileTypeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponse) SetBody(v *DescribeWebLockFileTypeSummaryResponseBody) *DescribeWebLockFileTypeSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeWebLockFileTypeSummaryResponse) Validate() error {
	return dara.Validate(s)
}
