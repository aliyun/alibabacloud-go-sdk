// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualBridgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualBridgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualBridgeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualBridgeResponseBody) *DeleteVirtualBridgeResponse
	GetBody() *DeleteVirtualBridgeResponseBody
}

type DeleteVirtualBridgeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualBridgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualBridgeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualBridgeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualBridgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualBridgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualBridgeResponse) GetBody() *DeleteVirtualBridgeResponseBody {
	return s.Body
}

func (s *DeleteVirtualBridgeResponse) SetHeaders(v map[string]*string) *DeleteVirtualBridgeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualBridgeResponse) SetStatusCode(v int32) *DeleteVirtualBridgeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualBridgeResponse) SetBody(v *DeleteVirtualBridgeResponseBody) *DeleteVirtualBridgeResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualBridgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
