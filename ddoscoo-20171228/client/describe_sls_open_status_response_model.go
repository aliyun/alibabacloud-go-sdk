// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsOpenStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsOpenStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsOpenStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsOpenStatusResponseBody) *DescribeSlsOpenStatusResponse
	GetBody() *DescribeSlsOpenStatusResponseBody
}

type DescribeSlsOpenStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsOpenStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsOpenStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsOpenStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsOpenStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsOpenStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsOpenStatusResponse) GetBody() *DescribeSlsOpenStatusResponseBody {
	return s.Body
}

func (s *DescribeSlsOpenStatusResponse) SetHeaders(v map[string]*string) *DescribeSlsOpenStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetStatusCode(v int32) *DescribeSlsOpenStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsOpenStatusResponse) SetBody(v *DescribeSlsOpenStatusResponseBody) *DescribeSlsOpenStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsOpenStatusResponse) Validate() error {
	return dara.Validate(s)
}
