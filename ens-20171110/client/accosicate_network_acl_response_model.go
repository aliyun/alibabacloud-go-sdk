// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccosicateNetworkAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AccosicateNetworkAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AccosicateNetworkAclResponse
	GetStatusCode() *int32
	SetBody(v *AccosicateNetworkAclResponseBody) *AccosicateNetworkAclResponse
	GetBody() *AccosicateNetworkAclResponseBody
}

type AccosicateNetworkAclResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccosicateNetworkAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AccosicateNetworkAclResponse) String() string {
	return dara.Prettify(s)
}

func (s AccosicateNetworkAclResponse) GoString() string {
	return s.String()
}

func (s *AccosicateNetworkAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AccosicateNetworkAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AccosicateNetworkAclResponse) GetBody() *AccosicateNetworkAclResponseBody {
	return s.Body
}

func (s *AccosicateNetworkAclResponse) SetHeaders(v map[string]*string) *AccosicateNetworkAclResponse {
	s.Headers = v
	return s
}

func (s *AccosicateNetworkAclResponse) SetStatusCode(v int32) *AccosicateNetworkAclResponse {
	s.StatusCode = &v
	return s
}

func (s *AccosicateNetworkAclResponse) SetBody(v *AccosicateNetworkAclResponseBody) *AccosicateNetworkAclResponse {
	s.Body = v
	return s
}

func (s *AccosicateNetworkAclResponse) Validate() error {
	return dara.Validate(s)
}
