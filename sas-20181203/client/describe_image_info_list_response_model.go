// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageInfoListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageInfoListResponseBody) *DescribeImageInfoListResponse
	GetBody() *DescribeImageInfoListResponseBody
}

type DescribeImageInfoListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageInfoListResponse) GetBody() *DescribeImageInfoListResponseBody {
	return s.Body
}

func (s *DescribeImageInfoListResponse) SetHeaders(v map[string]*string) *DescribeImageInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageInfoListResponse) SetStatusCode(v int32) *DescribeImageInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageInfoListResponse) SetBody(v *DescribeImageInfoListResponseBody) *DescribeImageInfoListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageInfoListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
