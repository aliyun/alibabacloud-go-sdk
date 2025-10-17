// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDocResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDocResponseBody) *DescribeDocResponse
	GetBody() *DescribeDocResponseBody
}

type DescribeDocResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDocResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDocResponse) GoString() string {
	return s.String()
}

func (s *DescribeDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDocResponse) GetBody() *DescribeDocResponseBody {
	return s.Body
}

func (s *DescribeDocResponse) SetHeaders(v map[string]*string) *DescribeDocResponse {
	s.Headers = v
	return s
}

func (s *DescribeDocResponse) SetStatusCode(v int32) *DescribeDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDocResponse) SetBody(v *DescribeDocResponseBody) *DescribeDocResponse {
	s.Body = v
	return s
}

func (s *DescribeDocResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
