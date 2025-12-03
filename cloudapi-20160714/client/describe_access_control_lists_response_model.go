// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessControlListsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessControlListsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessControlListsResponseBody) *DescribeAccessControlListsResponse
	GetBody() *DescribeAccessControlListsResponseBody
}

type DescribeAccessControlListsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessControlListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessControlListsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessControlListsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessControlListsResponse) GetBody() *DescribeAccessControlListsResponseBody {
	return s.Body
}

func (s *DescribeAccessControlListsResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListsResponse) SetStatusCode(v int32) *DescribeAccessControlListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListsResponse) SetBody(v *DescribeAccessControlListsResponseBody) *DescribeAccessControlListsResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessControlListsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
