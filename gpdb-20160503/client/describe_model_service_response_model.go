// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelServiceResponseBody) *DescribeModelServiceResponse
	GetBody() *DescribeModelServiceResponseBody
}

type DescribeModelServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelServiceResponse) GetBody() *DescribeModelServiceResponseBody {
	return s.Body
}

func (s *DescribeModelServiceResponse) SetHeaders(v map[string]*string) *DescribeModelServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelServiceResponse) SetStatusCode(v int32) *DescribeModelServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelServiceResponse) SetBody(v *DescribeModelServiceResponseBody) *DescribeModelServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeModelServiceResponse) Validate() error {
	return dara.Validate(s)
}
