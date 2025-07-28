// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecAbnormalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecAbnormalsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecAbnormalsResponseBody) *DescribeApisecAbnormalsResponse
	GetBody() *DescribeApisecAbnormalsResponseBody
}

type DescribeApisecAbnormalsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecAbnormalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecAbnormalsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecAbnormalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecAbnormalsResponse) GetBody() *DescribeApisecAbnormalsResponseBody {
	return s.Body
}

func (s *DescribeApisecAbnormalsResponse) SetHeaders(v map[string]*string) *DescribeApisecAbnormalsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecAbnormalsResponse) SetStatusCode(v int32) *DescribeApisecAbnormalsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecAbnormalsResponse) SetBody(v *DescribeApisecAbnormalsResponseBody) *DescribeApisecAbnormalsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecAbnormalsResponse) Validate() error {
	return dara.Validate(s)
}
