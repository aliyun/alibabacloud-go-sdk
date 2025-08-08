// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCurrentNodeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCurrentNodeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCurrentNodeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCurrentNodeInfoResponseBody) *DescribeCurrentNodeInfoResponse
	GetBody() *DescribeCurrentNodeInfoResponseBody
}

type DescribeCurrentNodeInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCurrentNodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCurrentNodeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCurrentNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCurrentNodeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCurrentNodeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCurrentNodeInfoResponse) GetBody() *DescribeCurrentNodeInfoResponseBody {
	return s.Body
}

func (s *DescribeCurrentNodeInfoResponse) SetHeaders(v map[string]*string) *DescribeCurrentNodeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCurrentNodeInfoResponse) SetStatusCode(v int32) *DescribeCurrentNodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCurrentNodeInfoResponse) SetBody(v *DescribeCurrentNodeInfoResponseBody) *DescribeCurrentNodeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCurrentNodeInfoResponse) Validate() error {
	return dara.Validate(s)
}
