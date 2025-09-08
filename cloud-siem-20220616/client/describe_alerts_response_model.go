// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertsResponseBody) *DescribeAlertsResponse
	GetBody() *DescribeAlertsResponseBody
}

type DescribeAlertsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertsResponse) GetBody() *DescribeAlertsResponseBody {
	return s.Body
}

func (s *DescribeAlertsResponse) SetHeaders(v map[string]*string) *DescribeAlertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsResponse) SetStatusCode(v int32) *DescribeAlertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsResponse) SetBody(v *DescribeAlertsResponseBody) *DescribeAlertsResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertsResponse) Validate() error {
	return dara.Validate(s)
}
