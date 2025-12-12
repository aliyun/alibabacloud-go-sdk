// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStatusSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterStatusSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterStatusSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterStatusSetResponseBody) *DescribeDBClusterStatusSetResponse
	GetBody() *DescribeDBClusterStatusSetResponseBody
}

type DescribeDBClusterStatusSetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterStatusSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterStatusSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStatusSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStatusSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterStatusSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterStatusSetResponse) GetBody() *DescribeDBClusterStatusSetResponseBody {
	return s.Body
}

func (s *DescribeDBClusterStatusSetResponse) SetHeaders(v map[string]*string) *DescribeDBClusterStatusSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterStatusSetResponse) SetStatusCode(v int32) *DescribeDBClusterStatusSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterStatusSetResponse) SetBody(v *DescribeDBClusterStatusSetResponseBody) *DescribeDBClusterStatusSetResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterStatusSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
