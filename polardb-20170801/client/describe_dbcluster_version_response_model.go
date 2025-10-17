// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterVersionResponseBody) *DescribeDBClusterVersionResponse
	GetBody() *DescribeDBClusterVersionResponseBody
}

type DescribeDBClusterVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterVersionResponse) GetBody() *DescribeDBClusterVersionResponseBody {
	return s.Body
}

func (s *DescribeDBClusterVersionResponse) SetHeaders(v map[string]*string) *DescribeDBClusterVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterVersionResponse) SetStatusCode(v int32) *DescribeDBClusterVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterVersionResponse) SetBody(v *DescribeDBClusterVersionResponseBody) *DescribeDBClusterVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
