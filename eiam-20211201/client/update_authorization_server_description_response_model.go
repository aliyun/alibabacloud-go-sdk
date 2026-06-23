// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationServerDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationServerDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationServerDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationServerDescriptionResponseBody) *UpdateAuthorizationServerDescriptionResponse
	GetBody() *UpdateAuthorizationServerDescriptionResponseBody
}

type UpdateAuthorizationServerDescriptionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationServerDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationServerDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationServerDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationServerDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationServerDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationServerDescriptionResponse) GetBody() *UpdateAuthorizationServerDescriptionResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationServerDescriptionResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationServerDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationServerDescriptionResponse) SetStatusCode(v int32) *UpdateAuthorizationServerDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationServerDescriptionResponse) SetBody(v *UpdateAuthorizationServerDescriptionResponseBody) *UpdateAuthorizationServerDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationServerDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
