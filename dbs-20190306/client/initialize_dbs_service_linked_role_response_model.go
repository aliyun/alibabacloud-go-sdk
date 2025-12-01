// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeDbsServiceLinkedRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeDbsServiceLinkedRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeDbsServiceLinkedRoleResponse
	GetStatusCode() *int32
	SetBody(v *InitializeDbsServiceLinkedRoleResponseBody) *InitializeDbsServiceLinkedRoleResponse
	GetBody() *InitializeDbsServiceLinkedRoleResponseBody
}

type InitializeDbsServiceLinkedRoleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeDbsServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeDbsServiceLinkedRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeDbsServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeDbsServiceLinkedRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeDbsServiceLinkedRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeDbsServiceLinkedRoleResponse) GetBody() *InitializeDbsServiceLinkedRoleResponseBody {
	return s.Body
}

func (s *InitializeDbsServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *InitializeDbsServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponse) SetStatusCode(v int32) *InitializeDbsServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponse) SetBody(v *InitializeDbsServiceLinkedRoleResponseBody) *InitializeDbsServiceLinkedRoleResponse {
	s.Body = v
	return s
}

func (s *InitializeDbsServiceLinkedRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
