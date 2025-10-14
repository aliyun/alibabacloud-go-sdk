// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarClassListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColumnarClassListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColumnarClassListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColumnarClassListResponseBody) *DescribeColumnarClassListResponse
	GetBody() *DescribeColumnarClassListResponseBody
}

type DescribeColumnarClassListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnarClassListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnarClassListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarClassListResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnarClassListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColumnarClassListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColumnarClassListResponse) GetBody() *DescribeColumnarClassListResponseBody {
	return s.Body
}

func (s *DescribeColumnarClassListResponse) SetHeaders(v map[string]*string) *DescribeColumnarClassListResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnarClassListResponse) SetStatusCode(v int32) *DescribeColumnarClassListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnarClassListResponse) SetBody(v *DescribeColumnarClassListResponseBody) *DescribeColumnarClassListResponse {
	s.Body = v
	return s
}

func (s *DescribeColumnarClassListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
