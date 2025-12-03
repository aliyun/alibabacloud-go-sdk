// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepositoryMemberWithInheritedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepositoryMemberWithInheritedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepositoryMemberWithInheritedResponse
	GetStatusCode() *int32
	SetBody(v *ListRepositoryMemberWithInheritedResponseBody) *ListRepositoryMemberWithInheritedResponse
	GetBody() *ListRepositoryMemberWithInheritedResponseBody
}

type ListRepositoryMemberWithInheritedResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryMemberWithInheritedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryMemberWithInheritedResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepositoryMemberWithInheritedResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryMemberWithInheritedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepositoryMemberWithInheritedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepositoryMemberWithInheritedResponse) GetBody() *ListRepositoryMemberWithInheritedResponseBody {
	return s.Body
}

func (s *ListRepositoryMemberWithInheritedResponse) SetHeaders(v map[string]*string) *ListRepositoryMemberWithInheritedResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetStatusCode(v int32) *ListRepositoryMemberWithInheritedResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) SetBody(v *ListRepositoryMemberWithInheritedResponseBody) *ListRepositoryMemberWithInheritedResponse {
	s.Body = v
	return s
}

func (s *ListRepositoryMemberWithInheritedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
