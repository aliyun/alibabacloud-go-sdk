// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationParametersResponseBody) *DescribeApplicationParametersResponse
	GetBody() *DescribeApplicationParametersResponseBody
}

type DescribeApplicationParametersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationParametersResponse) GetBody() *DescribeApplicationParametersResponseBody {
	return s.Body
}

func (s *DescribeApplicationParametersResponse) SetHeaders(v map[string]*string) *DescribeApplicationParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationParametersResponse) SetStatusCode(v int32) *DescribeApplicationParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationParametersResponse) SetBody(v *DescribeApplicationParametersResponseBody) *DescribeApplicationParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationParametersResponse) Validate() error {
	return dara.Validate(s)
}
