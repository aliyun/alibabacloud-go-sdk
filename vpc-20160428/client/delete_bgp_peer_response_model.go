// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBgpPeerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBgpPeerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBgpPeerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBgpPeerResponseBody) *DeleteBgpPeerResponse
	GetBody() *DeleteBgpPeerResponseBody
}

type DeleteBgpPeerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBgpPeerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBgpPeerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBgpPeerResponse) GoString() string {
	return s.String()
}

func (s *DeleteBgpPeerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBgpPeerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBgpPeerResponse) GetBody() *DeleteBgpPeerResponseBody {
	return s.Body
}

func (s *DeleteBgpPeerResponse) SetHeaders(v map[string]*string) *DeleteBgpPeerResponse {
	s.Headers = v
	return s
}

func (s *DeleteBgpPeerResponse) SetStatusCode(v int32) *DeleteBgpPeerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBgpPeerResponse) SetBody(v *DeleteBgpPeerResponseBody) *DeleteBgpPeerResponse {
	s.Body = v
	return s
}

func (s *DeleteBgpPeerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
