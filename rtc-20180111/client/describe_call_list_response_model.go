// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCallListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCallListResponseBody) *DescribeCallListResponse
	GetBody() *DescribeCallListResponseBody
}

type DescribeCallListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCallListResponse) GetBody() *DescribeCallListResponseBody {
	return s.Body
}

func (s *DescribeCallListResponse) SetHeaders(v map[string]*string) *DescribeCallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallListResponse) SetStatusCode(v int32) *DescribeCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallListResponse) SetBody(v *DescribeCallListResponseBody) *DescribeCallListResponse {
	s.Body = v
	return s
}

func (s *DescribeCallListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
