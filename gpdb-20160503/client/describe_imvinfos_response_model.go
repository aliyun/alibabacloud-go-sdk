// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMVInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIMVInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIMVInfosResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIMVInfosResponseBody) *DescribeIMVInfosResponse
	GetBody() *DescribeIMVInfosResponseBody
}

type DescribeIMVInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIMVInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIMVInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMVInfosResponse) GoString() string {
	return s.String()
}

func (s *DescribeIMVInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIMVInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIMVInfosResponse) GetBody() *DescribeIMVInfosResponseBody {
	return s.Body
}

func (s *DescribeIMVInfosResponse) SetHeaders(v map[string]*string) *DescribeIMVInfosResponse {
	s.Headers = v
	return s
}

func (s *DescribeIMVInfosResponse) SetStatusCode(v int32) *DescribeIMVInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIMVInfosResponse) SetBody(v *DescribeIMVInfosResponseBody) *DescribeIMVInfosResponse {
	s.Body = v
	return s
}

func (s *DescribeIMVInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
