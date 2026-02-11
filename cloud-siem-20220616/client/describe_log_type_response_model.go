// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogTypeResponseBody) *DescribeLogTypeResponse
	GetBody() *DescribeLogTypeResponseBody
}

type DescribeLogTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogTypeResponse) GetBody() *DescribeLogTypeResponseBody {
	return s.Body
}

func (s *DescribeLogTypeResponse) SetHeaders(v map[string]*string) *DescribeLogTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogTypeResponse) SetStatusCode(v int32) *DescribeLogTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogTypeResponse) SetBody(v *DescribeLogTypeResponseBody) *DescribeLogTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeLogTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
