// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfaceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetworkInterfaceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetworkInterfaceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetworkInterfaceAttributeResponseBody) *DescribeNetworkInterfaceAttributeResponse
	GetBody() *DescribeNetworkInterfaceAttributeResponseBody
}

type DescribeNetworkInterfaceAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetworkInterfaceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetworkInterfaceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetworkInterfaceAttributeResponse) GetBody() *DescribeNetworkInterfaceAttributeResponseBody {
	return s.Body
}

func (s *DescribeNetworkInterfaceAttributeResponse) SetHeaders(v map[string]*string) *DescribeNetworkInterfaceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponse) SetStatusCode(v int32) *DescribeNetworkInterfaceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponse) SetBody(v *DescribeNetworkInterfaceAttributeResponseBody) *DescribeNetworkInterfaceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
