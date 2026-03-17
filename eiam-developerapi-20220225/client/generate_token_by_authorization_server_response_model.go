// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTokenByAuthorizationServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTokenByAuthorizationServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTokenByAuthorizationServerResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTokenByAuthorizationServerResponseBody) *GenerateTokenByAuthorizationServerResponse
	GetBody() *GenerateTokenByAuthorizationServerResponseBody
}

type GenerateTokenByAuthorizationServerResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTokenByAuthorizationServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTokenByAuthorizationServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTokenByAuthorizationServerResponse) GoString() string {
	return s.String()
}

func (s *GenerateTokenByAuthorizationServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTokenByAuthorizationServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTokenByAuthorizationServerResponse) GetBody() *GenerateTokenByAuthorizationServerResponseBody {
	return s.Body
}

func (s *GenerateTokenByAuthorizationServerResponse) SetHeaders(v map[string]*string) *GenerateTokenByAuthorizationServerResponse {
	s.Headers = v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponse) SetStatusCode(v int32) *GenerateTokenByAuthorizationServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponse) SetBody(v *GenerateTokenByAuthorizationServerResponseBody) *GenerateTokenByAuthorizationServerResponse {
	s.Body = v
	return s
}

func (s *GenerateTokenByAuthorizationServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
