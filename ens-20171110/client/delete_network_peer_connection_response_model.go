// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPeerConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkPeerConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkPeerConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkPeerConnectionResponseBody) *DeleteNetworkPeerConnectionResponse
	GetBody() *DeleteNetworkPeerConnectionResponseBody
}

type DeleteNetworkPeerConnectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkPeerConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkPeerConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPeerConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPeerConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkPeerConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkPeerConnectionResponse) GetBody() *DeleteNetworkPeerConnectionResponseBody {
	return s.Body
}

func (s *DeleteNetworkPeerConnectionResponse) SetHeaders(v map[string]*string) *DeleteNetworkPeerConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkPeerConnectionResponse) SetStatusCode(v int32) *DeleteNetworkPeerConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkPeerConnectionResponse) SetBody(v *DeleteNetworkPeerConnectionResponseBody) *DeleteNetworkPeerConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkPeerConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
