// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertHistoryListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertHistoryListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertHistoryListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertHistoryListResponseBody) *DescribeAlertHistoryListResponse
	GetBody() *DescribeAlertHistoryListResponseBody
}

type DescribeAlertHistoryListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertHistoryListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertHistoryListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertHistoryListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertHistoryListResponse) GetBody() *DescribeAlertHistoryListResponseBody {
	return s.Body
}

func (s *DescribeAlertHistoryListResponse) SetHeaders(v map[string]*string) *DescribeAlertHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertHistoryListResponse) SetStatusCode(v int32) *DescribeAlertHistoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertHistoryListResponse) SetBody(v *DescribeAlertHistoryListResponseBody) *DescribeAlertHistoryListResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertHistoryListResponse) Validate() error {
	return dara.Validate(s)
}
