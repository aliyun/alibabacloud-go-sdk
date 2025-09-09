// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcPeerConnectionResponseBody) *CreateVpcPeerConnectionResponse
	GetBody() *CreateVpcPeerConnectionResponseBody
}

type CreateVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcPeerConnectionResponse) GetBody() *CreateVpcPeerConnectionResponseBody {
	return s.Body
}

func (s *CreateVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *CreateVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcPeerConnectionResponse) SetStatusCode(v int32) *CreateVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcPeerConnectionResponse) SetBody(v *CreateVpcPeerConnectionResponseBody) *CreateVpcPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateVpcPeerConnectionResponse) Validate() error {
	return dara.Validate(s)
}
