// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecExamplesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecExamplesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecExamplesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecExamplesResponseBody) *DescribeApisecExamplesResponse
	GetBody() *DescribeApisecExamplesResponseBody
}

type DescribeApisecExamplesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecExamplesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecExamplesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecExamplesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecExamplesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecExamplesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecExamplesResponse) GetBody() *DescribeApisecExamplesResponseBody {
	return s.Body
}

func (s *DescribeApisecExamplesResponse) SetHeaders(v map[string]*string) *DescribeApisecExamplesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecExamplesResponse) SetStatusCode(v int32) *DescribeApisecExamplesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecExamplesResponse) SetBody(v *DescribeApisecExamplesResponseBody) *DescribeApisecExamplesResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecExamplesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
