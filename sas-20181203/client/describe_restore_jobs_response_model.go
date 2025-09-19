// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreJobsResponseBody) *DescribeRestoreJobsResponse
	GetBody() *DescribeRestoreJobsResponseBody
}

type DescribeRestoreJobsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreJobsResponse) GetBody() *DescribeRestoreJobsResponseBody {
	return s.Body
}

func (s *DescribeRestoreJobsResponse) SetHeaders(v map[string]*string) *DescribeRestoreJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreJobsResponse) SetStatusCode(v int32) *DescribeRestoreJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreJobsResponse) SetBody(v *DescribeRestoreJobsResponseBody) *DescribeRestoreJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreJobsResponse) Validate() error {
	return dara.Validate(s)
}
