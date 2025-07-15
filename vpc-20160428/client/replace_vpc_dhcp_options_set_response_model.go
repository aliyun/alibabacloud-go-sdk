// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceVpcDhcpOptionsSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceVpcDhcpOptionsSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceVpcDhcpOptionsSetResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceVpcDhcpOptionsSetResponseBody) *ReplaceVpcDhcpOptionsSetResponse
	GetBody() *ReplaceVpcDhcpOptionsSetResponseBody
}

type ReplaceVpcDhcpOptionsSetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceVpcDhcpOptionsSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceVpcDhcpOptionsSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceVpcDhcpOptionsSetResponse) GoString() string {
	return s.String()
}

func (s *ReplaceVpcDhcpOptionsSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceVpcDhcpOptionsSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceVpcDhcpOptionsSetResponse) GetBody() *ReplaceVpcDhcpOptionsSetResponseBody {
	return s.Body
}

func (s *ReplaceVpcDhcpOptionsSetResponse) SetHeaders(v map[string]*string) *ReplaceVpcDhcpOptionsSetResponse {
	s.Headers = v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetResponse) SetStatusCode(v int32) *ReplaceVpcDhcpOptionsSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetResponse) SetBody(v *ReplaceVpcDhcpOptionsSetResponseBody) *ReplaceVpcDhcpOptionsSetResponse {
	s.Body = v
	return s
}

func (s *ReplaceVpcDhcpOptionsSetResponse) Validate() error {
	return dara.Validate(s)
}
