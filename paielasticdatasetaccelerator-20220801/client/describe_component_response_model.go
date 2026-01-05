// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentResponseBody) *DescribeComponentResponse
	GetBody() *DescribeComponentResponseBody
}

type DescribeComponentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentResponse) GetBody() *DescribeComponentResponseBody {
	return s.Body
}

func (s *DescribeComponentResponse) SetHeaders(v map[string]*string) *DescribeComponentResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentResponse) SetStatusCode(v int32) *DescribeComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentResponse) SetBody(v *DescribeComponentResponseBody) *DescribeComponentResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
