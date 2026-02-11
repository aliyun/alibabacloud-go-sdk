// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleForDeletingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForDeletingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceLinkedRoleForDeletingResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse
	GetBody() *CheckServiceLinkedRoleForDeletingResponseBody
}

type CheckServiceLinkedRoleForDeletingResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceLinkedRoleForDeletingResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleForDeletingResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleForDeletingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceLinkedRoleForDeletingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceLinkedRoleForDeletingResponse) GetBody() *CheckServiceLinkedRoleForDeletingResponseBody {
	return s.Body
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleForDeletingResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleForDeletingResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) SetBody(v *CheckServiceLinkedRoleForDeletingResponseBody) *CheckServiceLinkedRoleForDeletingResponse {
	s.Body = v
	return s
}

func (s *CheckServiceLinkedRoleForDeletingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
