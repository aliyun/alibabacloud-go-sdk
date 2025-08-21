// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsAuthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsAuthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsAuthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse
	GetBody() *DescribeSlsAuthStatusResponseBody
}

type DescribeSlsAuthStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsAuthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsAuthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsAuthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsAuthStatusResponse) GetBody() *DescribeSlsAuthStatusResponseBody {
	return s.Body
}

func (s *DescribeSlsAuthStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetStatusCode(v int32) *DescribeSlsAuthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsAuthStatusResponse) SetBody(v *DescribeSlsAuthStatusResponseBody) *DescribeSlsAuthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsAuthStatusResponse) Validate() error {
	return dara.Validate(s)
}
