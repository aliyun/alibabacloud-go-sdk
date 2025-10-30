// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProjectMemberResponseBody) *UpdateProjectMemberResponse
	GetBody() *UpdateProjectMemberResponseBody
}

type UpdateProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProjectMemberResponse) GetBody() *UpdateProjectMemberResponseBody {
	return s.Body
}

func (s *UpdateProjectMemberResponse) SetHeaders(v map[string]*string) *UpdateProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectMemberResponse) SetStatusCode(v int32) *UpdateProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectMemberResponse) SetBody(v *UpdateProjectMemberResponseBody) *UpdateProjectMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateProjectMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
