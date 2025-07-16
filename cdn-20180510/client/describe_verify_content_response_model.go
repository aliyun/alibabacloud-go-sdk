// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyContentResponseBody) *DescribeVerifyContentResponse
	GetBody() *DescribeVerifyContentResponseBody
}

type DescribeVerifyContentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyContentResponse) GetBody() *DescribeVerifyContentResponseBody {
	return s.Body
}

func (s *DescribeVerifyContentResponse) SetHeaders(v map[string]*string) *DescribeVerifyContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyContentResponse) SetStatusCode(v int32) *DescribeVerifyContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyContentResponse) SetBody(v *DescribeVerifyContentResponseBody) *DescribeVerifyContentResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyContentResponse) Validate() error {
	return dara.Validate(s)
}
