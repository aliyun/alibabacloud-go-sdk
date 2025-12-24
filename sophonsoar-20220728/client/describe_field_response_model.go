// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFieldResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFieldResponseBody) *DescribeFieldResponse
	GetBody() *DescribeFieldResponseBody
}

type DescribeFieldResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFieldResponse) GetBody() *DescribeFieldResponseBody {
	return s.Body
}

func (s *DescribeFieldResponse) SetHeaders(v map[string]*string) *DescribeFieldResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldResponse) SetStatusCode(v int32) *DescribeFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldResponse) SetBody(v *DescribeFieldResponseBody) *DescribeFieldResponse {
	s.Body = v
	return s
}

func (s *DescribeFieldResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
