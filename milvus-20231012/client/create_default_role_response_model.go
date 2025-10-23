// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefaultRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDefaultRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDefaultRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDefaultRoleResponseBody) *CreateDefaultRoleResponse
	GetBody() *CreateDefaultRoleResponseBody
}

type CreateDefaultRoleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDefaultRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDefaultRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDefaultRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateDefaultRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDefaultRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDefaultRoleResponse) GetBody() *CreateDefaultRoleResponseBody {
	return s.Body
}

func (s *CreateDefaultRoleResponse) SetHeaders(v map[string]*string) *CreateDefaultRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateDefaultRoleResponse) SetStatusCode(v int32) *CreateDefaultRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDefaultRoleResponse) SetBody(v *CreateDefaultRoleResponseBody) *CreateDefaultRoleResponse {
	s.Body = v
	return s
}

func (s *CreateDefaultRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
