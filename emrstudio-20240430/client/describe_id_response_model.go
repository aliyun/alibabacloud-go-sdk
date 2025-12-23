// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIdResponseBody) *DescribeIdResponse
	GetBody() *DescribeIdResponseBody
}

type DescribeIdResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIdResponse) GetBody() *DescribeIdResponseBody {
	return s.Body
}

func (s *DescribeIdResponse) SetHeaders(v map[string]*string) *DescribeIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdResponse) SetStatusCode(v int32) *DescribeIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdResponse) SetBody(v *DescribeIdResponseBody) *DescribeIdResponse {
	s.Body = v
	return s
}

func (s *DescribeIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
