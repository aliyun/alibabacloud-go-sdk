// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddProjectMemberResponseBody) *AddProjectMemberResponse
	GetBody() *AddProjectMemberResponseBody
}

type AddProjectMemberResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *AddProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddProjectMemberResponse) GetBody() *AddProjectMemberResponseBody {
	return s.Body
}

func (s *AddProjectMemberResponse) SetHeaders(v map[string]*string) *AddProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *AddProjectMemberResponse) SetStatusCode(v int32) *AddProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddProjectMemberResponse) SetBody(v *AddProjectMemberResponseBody) *AddProjectMemberResponse {
	s.Body = v
	return s
}

func (s *AddProjectMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
