// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterRecoverTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterRecoverTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterRecoverTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterRecoverTimeResponseBody) *DescribeClusterRecoverTimeResponse
	GetBody() *DescribeClusterRecoverTimeResponseBody
}

type DescribeClusterRecoverTimeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterRecoverTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterRecoverTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterRecoverTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterRecoverTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterRecoverTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterRecoverTimeResponse) GetBody() *DescribeClusterRecoverTimeResponseBody {
	return s.Body
}

func (s *DescribeClusterRecoverTimeResponse) SetHeaders(v map[string]*string) *DescribeClusterRecoverTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterRecoverTimeResponse) SetStatusCode(v int32) *DescribeClusterRecoverTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterRecoverTimeResponse) SetBody(v *DescribeClusterRecoverTimeResponseBody) *DescribeClusterRecoverTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterRecoverTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
