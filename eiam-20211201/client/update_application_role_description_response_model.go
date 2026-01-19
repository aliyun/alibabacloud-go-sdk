// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRoleDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationRoleDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationRoleDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationRoleDescriptionResponseBody) *UpdateApplicationRoleDescriptionResponse
	GetBody() *UpdateApplicationRoleDescriptionResponseBody
}

type UpdateApplicationRoleDescriptionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationRoleDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationRoleDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRoleDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRoleDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationRoleDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationRoleDescriptionResponse) GetBody() *UpdateApplicationRoleDescriptionResponseBody {
	return s.Body
}

func (s *UpdateApplicationRoleDescriptionResponse) SetHeaders(v map[string]*string) *UpdateApplicationRoleDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationRoleDescriptionResponse) SetStatusCode(v int32) *UpdateApplicationRoleDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationRoleDescriptionResponse) SetBody(v *UpdateApplicationRoleDescriptionResponseBody) *UpdateApplicationRoleDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationRoleDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
