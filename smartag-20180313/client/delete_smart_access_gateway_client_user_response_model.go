// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSmartAccessGatewayClientUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSmartAccessGatewayClientUserResponseBody) *DeleteSmartAccessGatewayClientUserResponse
	GetBody() *DeleteSmartAccessGatewayClientUserResponseBody
}

type DeleteSmartAccessGatewayClientUserResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSmartAccessGatewayClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSmartAccessGatewayClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayClientUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSmartAccessGatewayClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSmartAccessGatewayClientUserResponse) GetBody() *DeleteSmartAccessGatewayClientUserResponseBody {
	return s.Body
}

func (s *DeleteSmartAccessGatewayClientUserResponse) SetHeaders(v map[string]*string) *DeleteSmartAccessGatewayClientUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserResponse) SetStatusCode(v int32) *DeleteSmartAccessGatewayClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserResponse) SetBody(v *DeleteSmartAccessGatewayClientUserResponseBody) *DeleteSmartAccessGatewayClientUserResponse {
	s.Body = v
	return s
}

func (s *DeleteSmartAccessGatewayClientUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
