// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEtlJobLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEtlJobLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEtlJobLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEtlJobLogsResponseBody) *DescribeEtlJobLogsResponse
	GetBody() *DescribeEtlJobLogsResponseBody
}

type DescribeEtlJobLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEtlJobLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEtlJobLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEtlJobLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEtlJobLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEtlJobLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEtlJobLogsResponse) GetBody() *DescribeEtlJobLogsResponseBody {
	return s.Body
}

func (s *DescribeEtlJobLogsResponse) SetHeaders(v map[string]*string) *DescribeEtlJobLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEtlJobLogsResponse) SetStatusCode(v int32) *DescribeEtlJobLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEtlJobLogsResponse) SetBody(v *DescribeEtlJobLogsResponseBody) *DescribeEtlJobLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeEtlJobLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
