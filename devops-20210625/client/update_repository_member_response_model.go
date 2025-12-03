// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRepositoryMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRepositoryMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRepositoryMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRepositoryMemberResponseBody) *UpdateRepositoryMemberResponse
	GetBody() *UpdateRepositoryMemberResponseBody
}

type UpdateRepositoryMemberResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepositoryMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepositoryMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRepositoryMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRepositoryMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRepositoryMemberResponse) GetBody() *UpdateRepositoryMemberResponseBody {
	return s.Body
}

func (s *UpdateRepositoryMemberResponse) SetHeaders(v map[string]*string) *UpdateRepositoryMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryMemberResponse) SetStatusCode(v int32) *UpdateRepositoryMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryMemberResponse) SetBody(v *UpdateRepositoryMemberResponseBody) *UpdateRepositoryMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateRepositoryMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
