// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectInfoResponseBody) *DescribeProjectInfoResponse
	GetBody() *DescribeProjectInfoResponseBody
}

type DescribeProjectInfoResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectInfoResponse) GetBody() *DescribeProjectInfoResponseBody {
	return s.Body
}

func (s *DescribeProjectInfoResponse) SetHeaders(v map[string]*string) *DescribeProjectInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectInfoResponse) SetStatusCode(v int32) *DescribeProjectInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectInfoResponse) SetBody(v *DescribeProjectInfoResponseBody) *DescribeProjectInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectInfoResponse) Validate() error {
	return dara.Validate(s)
}
