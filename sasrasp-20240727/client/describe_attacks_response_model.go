// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttacksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttacksResponseBody) *DescribeAttacksResponse
	GetBody() *DescribeAttacksResponseBody
}

type DescribeAttacksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttacksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttacksResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttacksResponse) GetBody() *DescribeAttacksResponseBody {
	return s.Body
}

func (s *DescribeAttacksResponse) SetHeaders(v map[string]*string) *DescribeAttacksResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttacksResponse) SetStatusCode(v int32) *DescribeAttacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttacksResponse) SetBody(v *DescribeAttacksResponseBody) *DescribeAttacksResponse {
	s.Body = v
	return s
}

func (s *DescribeAttacksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
