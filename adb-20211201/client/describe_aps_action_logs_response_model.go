// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsActionLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsActionLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsActionLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsActionLogsResponseBody) *DescribeApsActionLogsResponse
	GetBody() *DescribeApsActionLogsResponseBody
}

type DescribeApsActionLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsActionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsActionLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsActionLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsActionLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsActionLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsActionLogsResponse) GetBody() *DescribeApsActionLogsResponseBody {
	return s.Body
}

func (s *DescribeApsActionLogsResponse) SetHeaders(v map[string]*string) *DescribeApsActionLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsActionLogsResponse) SetStatusCode(v int32) *DescribeApsActionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsActionLogsResponse) SetBody(v *DescribeApsActionLogsResponseBody) *DescribeApsActionLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeApsActionLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
