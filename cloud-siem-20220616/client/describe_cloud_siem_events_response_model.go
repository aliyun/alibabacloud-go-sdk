// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudSiemEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudSiemEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudSiemEventsResponseBody) *DescribeCloudSiemEventsResponse
	GetBody() *DescribeCloudSiemEventsResponseBody
}

type DescribeCloudSiemEventsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudSiemEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudSiemEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudSiemEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudSiemEventsResponse) GetBody() *DescribeCloudSiemEventsResponseBody {
	return s.Body
}

func (s *DescribeCloudSiemEventsResponse) SetHeaders(v map[string]*string) *DescribeCloudSiemEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetStatusCode(v int32) *DescribeCloudSiemEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudSiemEventsResponse) SetBody(v *DescribeCloudSiemEventsResponseBody) *DescribeCloudSiemEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudSiemEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
