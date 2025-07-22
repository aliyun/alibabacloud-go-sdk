// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSpnSummaryInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserSpnSummaryInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserSpnSummaryInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserSpnSummaryInfoResponseBody) *DescribeUserSpnSummaryInfoResponse
	GetBody() *DescribeUserSpnSummaryInfoResponseBody
}

type DescribeUserSpnSummaryInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserSpnSummaryInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserSpnSummaryInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSpnSummaryInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserSpnSummaryInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserSpnSummaryInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserSpnSummaryInfoResponse) GetBody() *DescribeUserSpnSummaryInfoResponseBody {
	return s.Body
}

func (s *DescribeUserSpnSummaryInfoResponse) SetHeaders(v map[string]*string) *DescribeUserSpnSummaryInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponse) SetStatusCode(v int32) *DescribeUserSpnSummaryInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponse) SetBody(v *DescribeUserSpnSummaryInfoResponseBody) *DescribeUserSpnSummaryInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeUserSpnSummaryInfoResponse) Validate() error {
	return dara.Validate(s)
}
