// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcPeerConnectionResponseBody) *DeleteVpcPeerConnectionResponse
	GetBody() *DeleteVpcPeerConnectionResponseBody
}

type DeleteVpcPeerConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcPeerConnectionResponse) GetBody() *DeleteVpcPeerConnectionResponseBody {
	return s.Body
}

func (s *DeleteVpcPeerConnectionResponse) SetHeaders(v map[string]*string) *DeleteVpcPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcPeerConnectionResponse) SetStatusCode(v int32) *DeleteVpcPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcPeerConnectionResponse) SetBody(v *DeleteVpcPeerConnectionResponseBody) *DeleteVpcPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcPeerConnectionResponse) Validate() error {
	return dara.Validate(s)
}
