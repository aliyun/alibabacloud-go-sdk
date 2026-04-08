// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateProjectMemberResponseBody) *CreateProjectMemberResponse
	GetBody() *CreateProjectMemberResponseBody
}

type CreateProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProjectMemberResponse) GetBody() *CreateProjectMemberResponseBody {
	return s.Body
}

func (s *CreateProjectMemberResponse) SetHeaders(v map[string]*string) *CreateProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectMemberResponse) SetStatusCode(v int32) *CreateProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectMemberResponse) SetBody(v *CreateProjectMemberResponseBody) *CreateProjectMemberResponse {
	s.Body = v
	return s
}

func (s *CreateProjectMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
