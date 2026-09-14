// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBVersionResponseBody) *DescribeDBVersionResponse
	GetBody() *DescribeDBVersionResponseBody
}

type DescribeDBVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBVersionResponse) GetBody() *DescribeDBVersionResponseBody {
	return s.Body
}

func (s *DescribeDBVersionResponse) SetHeaders(v map[string]*string) *DescribeDBVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBVersionResponse) SetStatusCode(v int32) *DescribeDBVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBVersionResponse) SetBody(v *DescribeDBVersionResponseBody) *DescribeDBVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeDBVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
