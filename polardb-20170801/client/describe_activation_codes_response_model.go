// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeActivationCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeActivationCodesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeActivationCodesResponseBody) *DescribeActivationCodesResponse
	GetBody() *DescribeActivationCodesResponseBody
}

type DescribeActivationCodesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeActivationCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeActivationCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationCodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeActivationCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeActivationCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeActivationCodesResponse) GetBody() *DescribeActivationCodesResponseBody {
	return s.Body
}

func (s *DescribeActivationCodesResponse) SetHeaders(v map[string]*string) *DescribeActivationCodesResponse {
	s.Headers = v
	return s
}

func (s *DescribeActivationCodesResponse) SetStatusCode(v int32) *DescribeActivationCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeActivationCodesResponse) SetBody(v *DescribeActivationCodesResponseBody) *DescribeActivationCodesResponse {
	s.Body = v
	return s
}

func (s *DescribeActivationCodesResponse) Validate() error {
	return dara.Validate(s)
}
