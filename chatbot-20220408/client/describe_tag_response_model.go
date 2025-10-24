// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagResponseBody) *DescribeTagResponse
	GetBody() *DescribeTagResponseBody
}

type DescribeTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagResponse) GetBody() *DescribeTagResponseBody {
	return s.Body
}

func (s *DescribeTagResponse) SetHeaders(v map[string]*string) *DescribeTagResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagResponse) SetStatusCode(v int32) *DescribeTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagResponse) SetBody(v *DescribeTagResponseBody) *DescribeTagResponse {
	s.Body = v
	return s
}

func (s *DescribeTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
