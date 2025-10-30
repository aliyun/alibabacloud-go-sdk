// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterNodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterNodeResponseBody) *DescribeDBClusterNodeResponse
	GetBody() *DescribeDBClusterNodeResponseBody
}

type DescribeDBClusterNodeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterNodeResponse) GetBody() *DescribeDBClusterNodeResponseBody {
	return s.Body
}

func (s *DescribeDBClusterNodeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetStatusCode(v int32) *DescribeDBClusterNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNodeResponse) SetBody(v *DescribeDBClusterNodeResponseBody) *DescribeDBClusterNodeResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
