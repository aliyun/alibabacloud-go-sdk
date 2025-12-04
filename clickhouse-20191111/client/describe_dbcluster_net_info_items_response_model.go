// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterNetInfoItemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterNetInfoItemsResponseBody) *DescribeDBClusterNetInfoItemsResponse
	GetBody() *DescribeDBClusterNetInfoItemsResponseBody
}

type DescribeDBClusterNetInfoItemsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterNetInfoItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterNetInfoItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterNetInfoItemsResponse) GetBody() *DescribeDBClusterNetInfoItemsResponseBody {
	return s.Body
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterNetInfoItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetStatusCode(v int32) *DescribeDBClusterNetInfoItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponse) SetBody(v *DescribeDBClusterNetInfoItemsResponseBody) *DescribeDBClusterNetInfoItemsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
