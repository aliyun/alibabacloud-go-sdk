// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNsasSuspEventTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNsasSuspEventTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNsasSuspEventTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNsasSuspEventTypeResponseBody) *DescribeNsasSuspEventTypeResponse
	GetBody() *DescribeNsasSuspEventTypeResponseBody
}

type DescribeNsasSuspEventTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNsasSuspEventTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNsasSuspEventTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNsasSuspEventTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNsasSuspEventTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNsasSuspEventTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNsasSuspEventTypeResponse) GetBody() *DescribeNsasSuspEventTypeResponseBody {
	return s.Body
}

func (s *DescribeNsasSuspEventTypeResponse) SetHeaders(v map[string]*string) *DescribeNsasSuspEventTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNsasSuspEventTypeResponse) SetStatusCode(v int32) *DescribeNsasSuspEventTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNsasSuspEventTypeResponse) SetBody(v *DescribeNsasSuspEventTypeResponseBody) *DescribeNsasSuspEventTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeNsasSuspEventTypeResponse) Validate() error {
	return dara.Validate(s)
}
