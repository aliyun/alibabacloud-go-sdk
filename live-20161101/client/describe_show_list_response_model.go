// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeShowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeShowListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeShowListResponseBody) *DescribeShowListResponse
	GetBody() *DescribeShowListResponseBody
}

type DescribeShowListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeShowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeShowListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListResponse) GoString() string {
	return s.String()
}

func (s *DescribeShowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeShowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeShowListResponse) GetBody() *DescribeShowListResponseBody {
	return s.Body
}

func (s *DescribeShowListResponse) SetHeaders(v map[string]*string) *DescribeShowListResponse {
	s.Headers = v
	return s
}

func (s *DescribeShowListResponse) SetStatusCode(v int32) *DescribeShowListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeShowListResponse) SetBody(v *DescribeShowListResponseBody) *DescribeShowListResponse {
	s.Body = v
	return s
}

func (s *DescribeShowListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
