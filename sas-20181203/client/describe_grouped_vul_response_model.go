// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedVulResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupedVulResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupedVulResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupedVulResponseBody) *DescribeGroupedVulResponse
	GetBody() *DescribeGroupedVulResponseBody
}

type DescribeGroupedVulResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupedVulResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupedVulResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedVulResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupedVulResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupedVulResponse) GetBody() *DescribeGroupedVulResponseBody {
	return s.Body
}

func (s *DescribeGroupedVulResponse) SetHeaders(v map[string]*string) *DescribeGroupedVulResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupedVulResponse) SetStatusCode(v int32) *DescribeGroupedVulResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupedVulResponse) SetBody(v *DescribeGroupedVulResponseBody) *DescribeGroupedVulResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupedVulResponse) Validate() error {
	return dara.Validate(s)
}
