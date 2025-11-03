// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDhcpOptionsSetToVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachDhcpOptionsSetToVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachDhcpOptionsSetToVpcResponse
	GetStatusCode() *int32
	SetBody(v *AttachDhcpOptionsSetToVpcResponseBody) *AttachDhcpOptionsSetToVpcResponse
	GetBody() *AttachDhcpOptionsSetToVpcResponseBody
}

type AttachDhcpOptionsSetToVpcResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachDhcpOptionsSetToVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachDhcpOptionsSetToVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachDhcpOptionsSetToVpcResponse) GoString() string {
	return s.String()
}

func (s *AttachDhcpOptionsSetToVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachDhcpOptionsSetToVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachDhcpOptionsSetToVpcResponse) GetBody() *AttachDhcpOptionsSetToVpcResponseBody {
	return s.Body
}

func (s *AttachDhcpOptionsSetToVpcResponse) SetHeaders(v map[string]*string) *AttachDhcpOptionsSetToVpcResponse {
	s.Headers = v
	return s
}

func (s *AttachDhcpOptionsSetToVpcResponse) SetStatusCode(v int32) *AttachDhcpOptionsSetToVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachDhcpOptionsSetToVpcResponse) SetBody(v *AttachDhcpOptionsSetToVpcResponseBody) *AttachDhcpOptionsSetToVpcResponse {
	s.Body = v
	return s
}

func (s *AttachDhcpOptionsSetToVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
