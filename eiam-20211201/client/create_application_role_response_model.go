// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationRoleResponseBody) *CreateApplicationRoleResponse
	GetBody() *CreateApplicationRoleResponseBody
}

type CreateApplicationRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationRoleResponse) GetBody() *CreateApplicationRoleResponseBody {
	return s.Body
}

func (s *CreateApplicationRoleResponse) SetHeaders(v map[string]*string) *CreateApplicationRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationRoleResponse) SetStatusCode(v int32) *CreateApplicationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationRoleResponse) SetBody(v *CreateApplicationRoleResponseBody) *CreateApplicationRoleResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
