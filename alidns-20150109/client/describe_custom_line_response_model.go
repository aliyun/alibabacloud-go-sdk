// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomLineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomLineResponseBody) *DescribeCustomLineResponse
	GetBody() *DescribeCustomLineResponseBody
}

type DescribeCustomLineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomLineResponse) GetBody() *DescribeCustomLineResponseBody {
	return s.Body
}

func (s *DescribeCustomLineResponse) SetHeaders(v map[string]*string) *DescribeCustomLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLineResponse) SetStatusCode(v int32) *DescribeCustomLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomLineResponse) SetBody(v *DescribeCustomLineResponseBody) *DescribeCustomLineResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
