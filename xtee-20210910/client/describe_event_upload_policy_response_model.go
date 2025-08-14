// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventUploadPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventUploadPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventUploadPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventUploadPolicyResponseBody) *DescribeEventUploadPolicyResponse
	GetBody() *DescribeEventUploadPolicyResponseBody
}

type DescribeEventUploadPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventUploadPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventUploadPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventUploadPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventUploadPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventUploadPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventUploadPolicyResponse) GetBody() *DescribeEventUploadPolicyResponseBody {
	return s.Body
}

func (s *DescribeEventUploadPolicyResponse) SetHeaders(v map[string]*string) *DescribeEventUploadPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventUploadPolicyResponse) SetStatusCode(v int32) *DescribeEventUploadPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventUploadPolicyResponse) SetBody(v *DescribeEventUploadPolicyResponseBody) *DescribeEventUploadPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeEventUploadPolicyResponse) Validate() error {
	return dara.Validate(s)
}
