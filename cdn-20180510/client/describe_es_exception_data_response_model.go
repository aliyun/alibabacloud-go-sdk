// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEsExceptionDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEsExceptionDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEsExceptionDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEsExceptionDataResponseBody) *DescribeEsExceptionDataResponse
	GetBody() *DescribeEsExceptionDataResponseBody
}

type DescribeEsExceptionDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEsExceptionDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEsExceptionDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEsExceptionDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEsExceptionDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEsExceptionDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEsExceptionDataResponse) GetBody() *DescribeEsExceptionDataResponseBody {
	return s.Body
}

func (s *DescribeEsExceptionDataResponse) SetHeaders(v map[string]*string) *DescribeEsExceptionDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEsExceptionDataResponse) SetStatusCode(v int32) *DescribeEsExceptionDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEsExceptionDataResponse) SetBody(v *DescribeEsExceptionDataResponseBody) *DescribeEsExceptionDataResponse {
	s.Body = v
	return s
}

func (s *DescribeEsExceptionDataResponse) Validate() error {
	return dara.Validate(s)
}
