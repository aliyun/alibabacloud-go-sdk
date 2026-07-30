// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualBridgeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualBridgeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualBridgeResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualBridgeResponseBody) *CreateVirtualBridgeResponse
	GetBody() *CreateVirtualBridgeResponseBody
}

type CreateVirtualBridgeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualBridgeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualBridgeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualBridgeResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualBridgeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualBridgeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualBridgeResponse) GetBody() *CreateVirtualBridgeResponseBody {
	return s.Body
}

func (s *CreateVirtualBridgeResponse) SetHeaders(v map[string]*string) *CreateVirtualBridgeResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualBridgeResponse) SetStatusCode(v int32) *CreateVirtualBridgeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualBridgeResponse) SetBody(v *CreateVirtualBridgeResponseBody) *CreateVirtualBridgeResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualBridgeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
