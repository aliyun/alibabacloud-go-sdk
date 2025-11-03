// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouterInterfaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouterInterfaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouterInterfaceAttributeResponseBody) *DescribeRouterInterfaceAttributeResponse
	GetBody() *DescribeRouterInterfaceAttributeResponseBody
}

type DescribeRouterInterfaceAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouterInterfaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouterInterfaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouterInterfaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouterInterfaceAttributeResponse) GetBody() *DescribeRouterInterfaceAttributeResponseBody {
	return s.Body
}

func (s *DescribeRouterInterfaceAttributeResponse) SetHeaders(v map[string]*string) *DescribeRouterInterfaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponse) SetStatusCode(v int32) *DescribeRouterInterfaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponse) SetBody(v *DescribeRouterInterfaceAttributeResponseBody) *DescribeRouterInterfaceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
