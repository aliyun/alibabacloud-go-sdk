// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecSlsProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecSlsProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecSlsProjectsResponseBody) *DescribeApisecSlsProjectsResponse
	GetBody() *DescribeApisecSlsProjectsResponseBody
}

type DescribeApisecSlsProjectsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecSlsProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecSlsProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecSlsProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecSlsProjectsResponse) GetBody() *DescribeApisecSlsProjectsResponseBody {
	return s.Body
}

func (s *DescribeApisecSlsProjectsResponse) SetHeaders(v map[string]*string) *DescribeApisecSlsProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecSlsProjectsResponse) SetStatusCode(v int32) *DescribeApisecSlsProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecSlsProjectsResponse) SetBody(v *DescribeApisecSlsProjectsResponseBody) *DescribeApisecSlsProjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecSlsProjectsResponse) Validate() error {
	return dara.Validate(s)
}
