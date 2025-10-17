// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterVersionZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterVersionZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterVersionZonalResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterVersionZonalResponseBody) *DescribeDBClusterVersionZonalResponse
	GetBody() *DescribeDBClusterVersionZonalResponseBody
}

type DescribeDBClusterVersionZonalResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterVersionZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterVersionZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterVersionZonalResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterVersionZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterVersionZonalResponse) GetBody() *DescribeDBClusterVersionZonalResponseBody {
	return s.Body
}

func (s *DescribeDBClusterVersionZonalResponse) SetHeaders(v map[string]*string) *DescribeDBClusterVersionZonalResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterVersionZonalResponse) SetStatusCode(v int32) *DescribeDBClusterVersionZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterVersionZonalResponse) SetBody(v *DescribeDBClusterVersionZonalResponseBody) *DescribeDBClusterVersionZonalResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterVersionZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
