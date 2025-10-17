// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegeObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountPrivilegeObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountPrivilegeObjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountPrivilegeObjectsResponseBody) *DescribeAccountPrivilegeObjectsResponse
	GetBody() *DescribeAccountPrivilegeObjectsResponseBody
}

type DescribeAccountPrivilegeObjectsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountPrivilegeObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountPrivilegeObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegeObjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegeObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountPrivilegeObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountPrivilegeObjectsResponse) GetBody() *DescribeAccountPrivilegeObjectsResponseBody {
	return s.Body
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetHeaders(v map[string]*string) *DescribeAccountPrivilegeObjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetStatusCode(v int32) *DescribeAccountPrivilegeObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponse) SetBody(v *DescribeAccountPrivilegeObjectsResponseBody) *DescribeAccountPrivilegeObjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountPrivilegeObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
