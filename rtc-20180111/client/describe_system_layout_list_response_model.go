// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSystemLayoutListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSystemLayoutListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSystemLayoutListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSystemLayoutListResponseBody) *DescribeSystemLayoutListResponse
	GetBody() *DescribeSystemLayoutListResponseBody
}

type DescribeSystemLayoutListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemLayoutListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemLayoutListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSystemLayoutListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemLayoutListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSystemLayoutListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSystemLayoutListResponse) GetBody() *DescribeSystemLayoutListResponseBody {
	return s.Body
}

func (s *DescribeSystemLayoutListResponse) SetHeaders(v map[string]*string) *DescribeSystemLayoutListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemLayoutListResponse) SetStatusCode(v int32) *DescribeSystemLayoutListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemLayoutListResponse) SetBody(v *DescribeSystemLayoutListResponseBody) *DescribeSystemLayoutListResponse {
	s.Body = v
	return s
}

func (s *DescribeSystemLayoutListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
