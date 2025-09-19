// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsSecurityEventSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetsSecurityEventSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetsSecurityEventSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetsSecurityEventSummaryResponseBody) *DescribeAssetsSecurityEventSummaryResponse
	GetBody() *DescribeAssetsSecurityEventSummaryResponseBody
}

type DescribeAssetsSecurityEventSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetsSecurityEventSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetsSecurityEventSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsSecurityEventSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetsSecurityEventSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetsSecurityEventSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetsSecurityEventSummaryResponse) GetBody() *DescribeAssetsSecurityEventSummaryResponseBody {
	return s.Body
}

func (s *DescribeAssetsSecurityEventSummaryResponse) SetHeaders(v map[string]*string) *DescribeAssetsSecurityEventSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponse) SetStatusCode(v int32) *DescribeAssetsSecurityEventSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponse) SetBody(v *DescribeAssetsSecurityEventSummaryResponseBody) *DescribeAssetsSecurityEventSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponse) Validate() error {
	return dara.Validate(s)
}
