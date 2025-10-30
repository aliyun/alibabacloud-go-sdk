// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterVulsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterVulsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterVulsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterVulsResponseBody) *DescribeClusterVulsResponse
	GetBody() *DescribeClusterVulsResponseBody
}

type DescribeClusterVulsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterVulsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterVulsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterVulsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterVulsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterVulsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterVulsResponse) GetBody() *DescribeClusterVulsResponseBody {
	return s.Body
}

func (s *DescribeClusterVulsResponse) SetHeaders(v map[string]*string) *DescribeClusterVulsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterVulsResponse) SetStatusCode(v int32) *DescribeClusterVulsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterVulsResponse) SetBody(v *DescribeClusterVulsResponseBody) *DescribeClusterVulsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterVulsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
