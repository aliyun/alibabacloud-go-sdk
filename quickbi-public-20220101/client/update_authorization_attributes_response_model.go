// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationAttributesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationAttributesResponseBody) *UpdateAuthorizationAttributesResponse
	GetBody() *UpdateAuthorizationAttributesResponseBody
}

type UpdateAuthorizationAttributesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationAttributesResponse) GetBody() *UpdateAuthorizationAttributesResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationAttributesResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationAttributesResponse) SetStatusCode(v int32) *UpdateAuthorizationAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationAttributesResponse) SetBody(v *UpdateAuthorizationAttributesResponseBody) *UpdateAuthorizationAttributesResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationAttributesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
