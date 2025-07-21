// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountKmsStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountKmsStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountKmsStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountKmsStatusResponseBody) *DescribeAccountKmsStatusResponse
	GetBody() *DescribeAccountKmsStatusResponseBody
}

type DescribeAccountKmsStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountKmsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountKmsStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountKmsStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountKmsStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountKmsStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountKmsStatusResponse) GetBody() *DescribeAccountKmsStatusResponseBody {
	return s.Body
}

func (s *DescribeAccountKmsStatusResponse) SetHeaders(v map[string]*string) *DescribeAccountKmsStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountKmsStatusResponse) SetStatusCode(v int32) *DescribeAccountKmsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountKmsStatusResponse) SetBody(v *DescribeAccountKmsStatusResponseBody) *DescribeAccountKmsStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountKmsStatusResponse) Validate() error {
	return dara.Validate(s)
}
