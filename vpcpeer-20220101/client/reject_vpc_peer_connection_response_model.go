// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectVpcPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectVpcPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectVpcPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *RejectVpcPeerConnectionResponseBody) *RejectVpcPeerConnectionResponse
	GetBody() *RejectVpcPeerConnectionResponseBody
}

type RejectVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectVpcPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *RejectVpcPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectVpcPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectVpcPeerConnectionResponse) GetBody() *RejectVpcPeerConnectionResponseBody {
	return s.Body
}

func (s *RejectVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *RejectVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *RejectVpcPeerConnectionResponse) SetStatusCode(v int32) *RejectVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectVpcPeerConnectionResponse) SetBody(v *RejectVpcPeerConnectionResponseBody) *RejectVpcPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *RejectVpcPeerConnectionResponse) Validate() error {
	return dara.Validate(s)
}
