// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserModelListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserModelListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserModelListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserModelListResponseBody) *DescribeUserModelListResponse
	GetBody() *DescribeUserModelListResponseBody
}

type DescribeUserModelListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserModelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserModelListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserModelListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserModelListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserModelListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserModelListResponse) GetBody() *DescribeUserModelListResponseBody {
	return s.Body
}

func (s *DescribeUserModelListResponse) SetHeaders(v map[string]*string) *DescribeUserModelListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserModelListResponse) SetStatusCode(v int32) *DescribeUserModelListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserModelListResponse) SetBody(v *DescribeUserModelListResponseBody) *DescribeUserModelListResponse {
	s.Body = v
	return s
}

func (s *DescribeUserModelListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
