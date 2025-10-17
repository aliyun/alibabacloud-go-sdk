// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterServerlessConfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterServerlessConfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterServerlessConfResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterServerlessConfResponseBody) *DescribeDBClusterServerlessConfResponse
	GetBody() *DescribeDBClusterServerlessConfResponseBody
}

type DescribeDBClusterServerlessConfResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterServerlessConfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterServerlessConfResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterServerlessConfResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterServerlessConfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterServerlessConfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterServerlessConfResponse) GetBody() *DescribeDBClusterServerlessConfResponseBody {
	return s.Body
}

func (s *DescribeDBClusterServerlessConfResponse) SetHeaders(v map[string]*string) *DescribeDBClusterServerlessConfResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterServerlessConfResponse) SetStatusCode(v int32) *DescribeDBClusterServerlessConfResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterServerlessConfResponse) SetBody(v *DescribeDBClusterServerlessConfResponseBody) *DescribeDBClusterServerlessConfResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterServerlessConfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
