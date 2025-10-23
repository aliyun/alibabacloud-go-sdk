// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessControlListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessControlListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessControlListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessControlListResponseBody) *DescribeAccessControlListResponse
	GetBody() *DescribeAccessControlListResponseBody
}

type DescribeAccessControlListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessControlListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessControlListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessControlListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessControlListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessControlListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessControlListResponse) GetBody() *DescribeAccessControlListResponseBody {
	return s.Body
}

func (s *DescribeAccessControlListResponse) SetHeaders(v map[string]*string) *DescribeAccessControlListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessControlListResponse) SetStatusCode(v int32) *DescribeAccessControlListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessControlListResponse) SetBody(v *DescribeAccessControlListResponseBody) *DescribeAccessControlListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessControlListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
