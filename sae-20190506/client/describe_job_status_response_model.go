// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobStatusResponseBody) *DescribeJobStatusResponse
	GetBody() *DescribeJobStatusResponseBody
}

type DescribeJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobStatusResponse) GetBody() *DescribeJobStatusResponseBody {
	return s.Body
}

func (s *DescribeJobStatusResponse) SetHeaders(v map[string]*string) *DescribeJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobStatusResponse) SetStatusCode(v int32) *DescribeJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobStatusResponse) SetBody(v *DescribeJobStatusResponseBody) *DescribeJobStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeJobStatusResponse) Validate() error {
	return dara.Validate(s)
}
