// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSmartAccessGatewayUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSmartAccessGatewayUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSmartAccessGatewayUserResponse
	GetStatusCode() *int32
	SetBody(v *DisableSmartAccessGatewayUserResponseBody) *DisableSmartAccessGatewayUserResponse
	GetBody() *DisableSmartAccessGatewayUserResponseBody
}

type DisableSmartAccessGatewayUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSmartAccessGatewayUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSmartAccessGatewayUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSmartAccessGatewayUserResponse) GoString() string {
	return s.String()
}

func (s *DisableSmartAccessGatewayUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSmartAccessGatewayUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSmartAccessGatewayUserResponse) GetBody() *DisableSmartAccessGatewayUserResponseBody {
	return s.Body
}

func (s *DisableSmartAccessGatewayUserResponse) SetHeaders(v map[string]*string) *DisableSmartAccessGatewayUserResponse {
	s.Headers = v
	return s
}

func (s *DisableSmartAccessGatewayUserResponse) SetStatusCode(v int32) *DisableSmartAccessGatewayUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSmartAccessGatewayUserResponse) SetBody(v *DisableSmartAccessGatewayUserResponseBody) *DisableSmartAccessGatewayUserResponse {
	s.Body = v
	return s
}

func (s *DisableSmartAccessGatewayUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
