// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationMseServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationMseServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationMseServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationMseServiceResponseBody) *DescribeApplicationMseServiceResponse
	GetBody() *DescribeApplicationMseServiceResponseBody
}

type DescribeApplicationMseServiceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationMseServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationMseServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationMseServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationMseServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationMseServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationMseServiceResponse) GetBody() *DescribeApplicationMseServiceResponseBody {
	return s.Body
}

func (s *DescribeApplicationMseServiceResponse) SetHeaders(v map[string]*string) *DescribeApplicationMseServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationMseServiceResponse) SetStatusCode(v int32) *DescribeApplicationMseServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationMseServiceResponse) SetBody(v *DescribeApplicationMseServiceResponseBody) *DescribeApplicationMseServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationMseServiceResponse) Validate() error {
	return dara.Validate(s)
}
