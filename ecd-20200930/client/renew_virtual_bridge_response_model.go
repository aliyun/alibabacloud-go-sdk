// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewVirtualBridgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewVirtualBridgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewVirtualBridgeResponse
	GetStatusCode() *int32
	SetBody(v *RenewVirtualBridgeResponseBody) *RenewVirtualBridgeResponse
	GetBody() *RenewVirtualBridgeResponseBody
}

type RenewVirtualBridgeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewVirtualBridgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewVirtualBridgeResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewVirtualBridgeResponse) GoString() string {
	return s.String()
}

func (s *RenewVirtualBridgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewVirtualBridgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewVirtualBridgeResponse) GetBody() *RenewVirtualBridgeResponseBody {
	return s.Body
}

func (s *RenewVirtualBridgeResponse) SetHeaders(v map[string]*string) *RenewVirtualBridgeResponse {
	s.Headers = v
	return s
}

func (s *RenewVirtualBridgeResponse) SetStatusCode(v int32) *RenewVirtualBridgeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewVirtualBridgeResponse) SetBody(v *RenewVirtualBridgeResponseBody) *RenewVirtualBridgeResponse {
	s.Body = v
	return s
}

func (s *RenewVirtualBridgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
