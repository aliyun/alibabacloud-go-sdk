// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleAssignmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoleAssignmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoleAssignmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoleAssignmentResponseBody) *DeleteRoleAssignmentResponse
	GetBody() *DeleteRoleAssignmentResponseBody
}

type DeleteRoleAssignmentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoleAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoleAssignmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleAssignmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleAssignmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoleAssignmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoleAssignmentResponse) GetBody() *DeleteRoleAssignmentResponseBody {
	return s.Body
}

func (s *DeleteRoleAssignmentResponse) SetHeaders(v map[string]*string) *DeleteRoleAssignmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoleAssignmentResponse) SetStatusCode(v int32) *DeleteRoleAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoleAssignmentResponse) SetBody(v *DeleteRoleAssignmentResponseBody) *DeleteRoleAssignmentResponse {
	s.Body = v
	return s
}

func (s *DeleteRoleAssignmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
