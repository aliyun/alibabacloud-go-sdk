// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationServerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationServerResponseBody) *UpdateAuthorizationServerResponse
	GetBody() *UpdateAuthorizationServerResponseBody
}

type UpdateAuthorizationServerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationServerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationServerResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationServerResponse) GetBody() *UpdateAuthorizationServerResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationServerResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationServerResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationServerResponse) SetStatusCode(v int32) *UpdateAuthorizationServerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationServerResponse) SetBody(v *UpdateAuthorizationServerResponseBody) *UpdateAuthorizationServerResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
