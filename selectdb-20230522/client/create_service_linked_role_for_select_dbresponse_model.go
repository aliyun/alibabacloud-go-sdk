// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleForSelectDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForSelectDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceLinkedRoleForSelectDBResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceLinkedRoleForSelectDBResponseBody) *CreateServiceLinkedRoleForSelectDBResponse
	GetBody() *CreateServiceLinkedRoleForSelectDBResponseBody
}

type CreateServiceLinkedRoleForSelectDBResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceLinkedRoleForSelectDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceLinkedRoleForSelectDBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleForSelectDBResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) GetBody() *CreateServiceLinkedRoleForSelectDBResponseBody {
	return s.Body
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForSelectDBResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetStatusCode(v int32) *CreateServiceLinkedRoleForSelectDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) SetBody(v *CreateServiceLinkedRoleForSelectDBResponseBody) *CreateServiceLinkedRoleForSelectDBResponse {
	s.Body = v
	return s
}

func (s *CreateServiceLinkedRoleForSelectDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
