// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmartAccessGatewayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmartAccessGatewayResponseBody) *DeleteSmartAccessGatewayResponse
	GetBody() *DeleteSmartAccessGatewayResponseBody
}

type DeleteSmartAccessGatewayResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmartAccessGatewayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmartAccessGatewayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmartAccessGatewayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmartAccessGatewayResponse) GetBody() *DeleteSmartAccessGatewayResponseBody {
	return s.Body
}

func (s *DeleteSmartAccessGatewayResponse) SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmartAccessGatewayResponse) SetStatusCode(v int32) *DeleteSmartAccessGatewayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmartAccessGatewayResponse) SetBody(v *DeleteSmartAccessGatewayResponseBody) *DeleteSmartAccessGatewayResponse {
	s.Body = v
	return s
}

func (s *DeleteSmartAccessGatewayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
