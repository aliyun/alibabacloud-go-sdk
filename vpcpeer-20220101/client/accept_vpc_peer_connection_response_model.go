// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptVpcPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AcceptVpcPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AcceptVpcPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *AcceptVpcPeerConnectionResponseBody) *AcceptVpcPeerConnectionResponse
	GetBody() *AcceptVpcPeerConnectionResponseBody
}

type AcceptVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptVpcPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s AcceptVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *AcceptVpcPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AcceptVpcPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AcceptVpcPeerConnectionResponse) GetBody() *AcceptVpcPeerConnectionResponseBody {
	return s.Body
}

func (s *AcceptVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *AcceptVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *AcceptVpcPeerConnectionResponse) SetStatusCode(v int32) *AcceptVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptVpcPeerConnectionResponse) SetBody(v *AcceptVpcPeerConnectionResponseBody) *AcceptVpcPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *AcceptVpcPeerConnectionResponse) Validate() error {
	return dara.Validate(s)
}
