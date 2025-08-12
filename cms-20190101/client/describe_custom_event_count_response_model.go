// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomEventCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomEventCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomEventCountResponseBody) *DescribeCustomEventCountResponse
	GetBody() *DescribeCustomEventCountResponseBody
}

type DescribeCustomEventCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomEventCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomEventCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomEventCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomEventCountResponse) GetBody() *DescribeCustomEventCountResponseBody {
	return s.Body
}

func (s *DescribeCustomEventCountResponse) SetHeaders(v map[string]*string) *DescribeCustomEventCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventCountResponse) SetStatusCode(v int32) *DescribeCustomEventCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomEventCountResponse) SetBody(v *DescribeCustomEventCountResponseBody) *DescribeCustomEventCountResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomEventCountResponse) Validate() error {
	return dara.Validate(s)
}
