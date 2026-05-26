// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleAssignmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoleAssignmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoleAssignmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoleAssignmentResponseBody) *CreateRoleAssignmentResponse
	GetBody() *CreateRoleAssignmentResponseBody
}

type CreateRoleAssignmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoleAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleAssignmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleAssignmentResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleAssignmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoleAssignmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoleAssignmentResponse) GetBody() *CreateRoleAssignmentResponseBody {
	return s.Body
}

func (s *CreateRoleAssignmentResponse) SetHeaders(v map[string]*string) *CreateRoleAssignmentResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleAssignmentResponse) SetStatusCode(v int32) *CreateRoleAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleAssignmentResponse) SetBody(v *CreateRoleAssignmentResponseBody) *CreateRoleAssignmentResponse {
	s.Body = v
	return s
}

func (s *CreateRoleAssignmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
