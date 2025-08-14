// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateStackResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateStackResponseBody) *DescribePrivateStackResponse
	GetBody() *DescribePrivateStackResponseBody
}

type DescribePrivateStackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateStackResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateStackResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateStackResponse) GetBody() *DescribePrivateStackResponseBody {
	return s.Body
}

func (s *DescribePrivateStackResponse) SetHeaders(v map[string]*string) *DescribePrivateStackResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateStackResponse) SetStatusCode(v int32) *DescribePrivateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateStackResponse) SetBody(v *DescribePrivateStackResponseBody) *DescribePrivateStackResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateStackResponse) Validate() error {
	return dara.Validate(s)
}
