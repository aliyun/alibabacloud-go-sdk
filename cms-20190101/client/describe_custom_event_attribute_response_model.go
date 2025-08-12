// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEventAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomEventAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomEventAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomEventAttributeResponseBody) *DescribeCustomEventAttributeResponse
	GetBody() *DescribeCustomEventAttributeResponseBody
}

type DescribeCustomEventAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomEventAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomEventAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEventAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEventAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomEventAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomEventAttributeResponse) GetBody() *DescribeCustomEventAttributeResponseBody {
	return s.Body
}

func (s *DescribeCustomEventAttributeResponse) SetHeaders(v map[string]*string) *DescribeCustomEventAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEventAttributeResponse) SetStatusCode(v int32) *DescribeCustomEventAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomEventAttributeResponse) SetBody(v *DescribeCustomEventAttributeResponseBody) *DescribeCustomEventAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomEventAttributeResponse) Validate() error {
	return dara.Validate(s)
}
