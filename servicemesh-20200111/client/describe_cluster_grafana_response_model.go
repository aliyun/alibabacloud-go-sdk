// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterGrafanaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterGrafanaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterGrafanaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterGrafanaResponseBody) *DescribeClusterGrafanaResponse
	GetBody() *DescribeClusterGrafanaResponseBody
}

type DescribeClusterGrafanaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterGrafanaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterGrafanaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterGrafanaResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterGrafanaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterGrafanaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterGrafanaResponse) GetBody() *DescribeClusterGrafanaResponseBody {
	return s.Body
}

func (s *DescribeClusterGrafanaResponse) SetHeaders(v map[string]*string) *DescribeClusterGrafanaResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterGrafanaResponse) SetStatusCode(v int32) *DescribeClusterGrafanaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterGrafanaResponse) SetBody(v *DescribeClusterGrafanaResponseBody) *DescribeClusterGrafanaResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterGrafanaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
