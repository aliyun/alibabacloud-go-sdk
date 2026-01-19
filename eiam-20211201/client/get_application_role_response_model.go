// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationRoleResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationRoleResponseBody) *GetApplicationRoleResponse
	GetBody() *GetApplicationRoleResponseBody
}

type GetApplicationRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRoleResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationRoleResponse) GetBody() *GetApplicationRoleResponseBody {
	return s.Body
}

func (s *GetApplicationRoleResponse) SetHeaders(v map[string]*string) *GetApplicationRoleResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationRoleResponse) SetStatusCode(v int32) *GetApplicationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationRoleResponse) SetBody(v *GetApplicationRoleResponseBody) *GetApplicationRoleResponse {
	s.Body = v
	return s
}

func (s *GetApplicationRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
