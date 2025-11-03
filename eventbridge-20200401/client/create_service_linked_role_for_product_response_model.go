// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceLinkedRoleForProductResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceLinkedRoleForProductResponseBody) *CreateServiceLinkedRoleForProductResponse
	GetBody() *CreateServiceLinkedRoleForProductResponseBody
}

type CreateServiceLinkedRoleForProductResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleForProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleForProductResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForProductResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceLinkedRoleForProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceLinkedRoleForProductResponse) GetBody() *CreateServiceLinkedRoleForProductResponseBody {
	return s.Body
}

func (s *CreateServiceLinkedRoleForProductResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForProductResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleForProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponse) SetBody(v *CreateServiceLinkedRoleForProductResponseBody) *CreateServiceLinkedRoleForProductResponse {
	s.Body = v
	return s
}

func (s *CreateServiceLinkedRoleForProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
