// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDhcpOptionsSetFromVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachDhcpOptionsSetFromVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachDhcpOptionsSetFromVpcResponse
	GetStatusCode() *int32
	SetBody(v *DetachDhcpOptionsSetFromVpcResponseBody) *DetachDhcpOptionsSetFromVpcResponse
	GetBody() *DetachDhcpOptionsSetFromVpcResponseBody
}

type DetachDhcpOptionsSetFromVpcResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachDhcpOptionsSetFromVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachDhcpOptionsSetFromVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachDhcpOptionsSetFromVpcResponse) GoString() string {
	return s.String()
}

func (s *DetachDhcpOptionsSetFromVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachDhcpOptionsSetFromVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachDhcpOptionsSetFromVpcResponse) GetBody() *DetachDhcpOptionsSetFromVpcResponseBody {
	return s.Body
}

func (s *DetachDhcpOptionsSetFromVpcResponse) SetHeaders(v map[string]*string) *DetachDhcpOptionsSetFromVpcResponse {
	s.Headers = v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcResponse) SetStatusCode(v int32) *DetachDhcpOptionsSetFromVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcResponse) SetBody(v *DetachDhcpOptionsSetFromVpcResponseBody) *DetachDhcpOptionsSetFromVpcResponse {
	s.Body = v
	return s
}

func (s *DetachDhcpOptionsSetFromVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
