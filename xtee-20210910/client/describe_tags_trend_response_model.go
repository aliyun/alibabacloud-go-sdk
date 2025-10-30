// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagsTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagsTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagsTrendResponseBody) *DescribeTagsTrendResponse
	GetBody() *DescribeTagsTrendResponseBody
}

type DescribeTagsTrendResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagsTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagsTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagsTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagsTrendResponse) GetBody() *DescribeTagsTrendResponseBody {
	return s.Body
}

func (s *DescribeTagsTrendResponse) SetHeaders(v map[string]*string) *DescribeTagsTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagsTrendResponse) SetStatusCode(v int32) *DescribeTagsTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagsTrendResponse) SetBody(v *DescribeTagsTrendResponseBody) *DescribeTagsTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeTagsTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
